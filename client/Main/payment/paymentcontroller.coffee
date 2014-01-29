class PaymentController extends KDController

  fetchPaymentMethods: (callback) ->

    { dash } = Bongo

    methods = null
    preferredPaymentMethod = null
    appStorage = new AppStorage 'Account', '1.0'
    queue = [

      -> appStorage.fetchStorage ->
        preferredPaymentMethod = appStorage.getValue 'preferredPaymentMethod'
        queue.fin()

      => KD.whoami().fetchPaymentMethods (err, paymentMethods) ->
        methods = paymentMethods
        queue.fin err
    ]

    dash queue, (err) -> callback err, {
      preferredPaymentMethod
      methods
      appStorage
    }

  observePaymentSave: (modal, callback) ->
    modal.on 'PaymentInfoSubmitted', (paymentMethodId, updatedPaymentInfo) =>
      @updatePaymentInfo paymentMethodId, updatedPaymentInfo, (err, savedPaymentInfo) =>
        return callback err  if err
        callback null, savedPaymentInfo
        @emit 'PaymentDataChanged'

  removePaymentMethod: (paymentMethodId, callback) ->
    { JPayment } = KD.remote.api
    JPayment.removePaymentMethod paymentMethodId, (err) =>
      return callback err  if err
      @emit 'PaymentDataChanged'

  fetchSubscription: do ->
    findActiveSubscription = (subscriptions, planCode, callback) ->
      for own paymentMethodId, subs of subscriptions
        for sub in subscriptions
          if sub.planCode is planCode and sub.status in ['canceled', 'active']
            return callback null, sub

      callback null

    fetchSubscription = (type, planCode, callback) ->
      { JPaymentSubscription } = KD.remote.api

      if type is 'group'
        KD.getGroup().checkPayment (err, subs) =>
          findActiveSubscription subs, planCode, callback
      else
        JPaymentSubscription.fetchUserSubscriptions (err, subs) ->
          findActiveSubscription subs, planCode, callback

  fetchPlanByCode: (planCode, callback) ->

    { JPaymentPlan } = KD.remote.api

    JPaymentPlan.fetchPlanByCode planCode, callback

  fetchPaymentInfo: (type, callback) ->

    { JPaymentPlan } = KD.remote.api

    switch type
      when 'group', 'expensed'
        KD.getGroup().fetchPaymentInfo callback
      when 'user'
        JPaymentPlan.fetchAccountDetails callback

  updatePaymentInfo: (paymentMethodId, paymentMethod, callback) ->

    { JPayment } = KD.remote.api

    JPayment.setPaymentInfo paymentMethodId, paymentMethod, callback

  createPaymentInfoModal: -> new PaymentFormModal

  createUpgradeForm: (tag, options = {}) ->

    { dash } = Bongo

    form = new PlanUpgradeForm { tag }

    KD.getGroup().fetchProducts 'plan', tags: tag, (err, plans) =>
      return  if KD.showError err

      queue = plans.map (plan) -> ->
        plan.fetchProducts (err, products) ->
          return  if KD.showError err

          plan.childProducts = products
          queue.fin()

      subscription = null
      queue.push =>
        @fetchSubscriptionsWithPlans tags: tag, (err, [subscription_]) ->
          subscription = subscription_
          queue.fin()

      dash queue, ->
        form.setPlans plans
        form.setCurrentSubscription subscription, options  if subscription

    return form

  createUpgradeWorkflow: (options = {}) ->
    {tag, productForm, confirmForm} = options

    productForm or= @createUpgradeForm tag, options
    confirmForm or= new PlanUpgradeConfirmForm
    workflow      = new PaymentWorkflow {productForm, confirmForm}

    productForm
      .on 'PlanSelected', (plan, planOptions) ->
        callback = ->
          workflow.collectData productData: { plan, planOptions }

        {oldSubscription} = workflow.collector.data
        unless oldSubscription
        then callback()
        else
          spend = oldSubscription?.usage ? {}
          oldSubscription.checkQuota {spend, multiplyFactor: 1}, (err) ->
            return  if KD.showError err
            callback()

      .on 'CurrentSubscriptionSet', (oldSubscription) ->
        workflow.collectData { oldSubscription }

    workflow
      .on 'DataCollected', (data) =>
        @transitionSubscription data, (err, subscription) ->
          return  if KD.showError err
          workflow.emit 'Finished', data, subscription
      .enter()

    workflow

  confirmReactivation: (subscription, callback) ->
    modal = KDModalView.confirm
      title       : 'Inactive subscription'
      description :
        """
        Your existing subscription for this plan has been canceled.  Would
        you like to reactivate it?
        """
      subView     : new SubscriptionView {}, subscription
      ok          :
        title     : 'Reactivate'
        callback  : -> subscription.resume (err) ->
          return callback err  if err

          modal.destroy()

          callback null, subscription

  createSubscription: (options, callback) ->
    { plan, planOptions, promotionType, email, paymentMethod, createAccount } = options
    { paymentMethodId, billing } = paymentMethod
    { planApi } = planOptions

    throw new Error "Must provide a plan API!"  unless planApi?

    options = {
      planOptions
      promotionType
      paymentMethodId
      planCode: plan.planCode
    }

    planApi.subscribe options, (err, subscription) =>
      if err?.short is 'existing_subscription'
        { existingSubscription } = err

        if existingSubscription.status is 'active'
          new KDNotificationView
            title: "You are already subscribed to this plan!"
          KD.getSingleton('router').handleRoute '/Account/Subscriptions'

        else
          existingSubscription.plan = plan
          @confirmReactivation existingSubscription, callback

      else if createAccount
        { JUser } = KD.remote.api

        { cardFirstName: firstName, cardLastName: lastName } = billing

        JUser.convert { firstName, lastName, email }, (err) ->
          return callback err  if err

          JUser.logout (err) ->
            return callback err  if err

            callback null
      else
        callback err, subscription

  transitionSubscription: (formData, callback) ->
    { productData, oldSubscription, promotionType, paymentMethod, createAccount, email } = formData
    { plan, planOptions } = productData
    { planCode } = plan
    { paymentMethodId } = paymentMethod
    if oldSubscription
      oldSubscription.transitionTo { planCode, paymentMethodId }, callback
    else
      @createSubscription {
        plan
        planOptions
        promotionType
        email
        paymentMethod
        createAccount
      }, callback

  debitSubscription: (subscription, pack, callback) ->
    subscription.debit pack, (err, nonce) =>
      return  if KD.showError err

      @emit 'SubscriptionDebited', subscription

      callback null, nonce

  fetchSubscriptionsWithPlans: (options, callback) ->
    [callback, options] = [options, callback]  unless callback

    options ?= {}

    KD.whoami().fetchPlansAndSubscriptions options, (err, plansAndSubs) =>
      return callback err  if err

      { subscriptions } = @groupPlansBySubscription plansAndSubs

      callback null, subscriptions

  groupPlansBySubscription: (plansAndSubscriptions = {}) ->

    { plans, subscriptions } = plansAndSubscriptions

    plansByCode = plans.reduce( (memo, plan) ->
      memo[plan.planCode] = plan
      memo
    , {})

    for subscription in subscriptions
      subscription.plan = plansByCode[subscription.planCode]

    { plans, subscriptions }
