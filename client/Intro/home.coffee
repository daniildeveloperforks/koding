class HomePage extends JView

  constructor:(options = {}, data)->

    options.domId = 'home-page'

    super options, data

    @pricingButton = new KDButtonView
      title       : "<a href='mailto:sales@koding.com?subject=Koding, white label' target='_self'>Get your own Koding for your team<cite>Contact us for details</cite></a>"
      cssClass    : 'solid green shadowed pricing'
      icon        : 'yes'
      iconClass   : 'dollar'
      click       : (event)-> KD.utils.stopDOMEvent event

    @registerForm = new HomeRegisterForm
      callback    : (formData)-> @doRegister formData

    @registerFormBottom = new HomeRegisterForm
      callback    : (formData)-> @doRegister formData

    @githubLink   = new KDCustomHTMLView
      tagName     : "a"
      partial     : "Or you can sign up using <strong>GitHub</strong>"
      click       : ->
        KD.singletons.oauthController.openPopup "github"

    @markers = new MarkerController

  show:->

    @unsetClass 'out'
    document.body.classList.add 'intro'
    KD.utils.defer => @markers.reset()

    super

  hide:->

    @setClass 'out'
    document.body.classList.remove 'intro'

    super

  viewAppended:->

    super

    vmMarker = @markers.create 'vms',
      client    : '#home-page .laptop .teamwork'
      container : this
      wait      : 1000
      message   : 'ACCESS YOUR VMS ONLINE'
      offset    :
        top     : 150
        left    : 50

    navMarker = @markers.create 'nav',
      client    : '#home-page .laptop .teamwork'
      container : this
      wait      : 1300
      message   : 'EASY ACCESS TO YOUR APPS'
      offset    :
        top     : -30
        left    : 240

    chatMarker = @markers.create 'chat',
      client    : '#home-page .laptop .teamwork'
      container : this
      wait      : 1600
      message   : 'WORK TOGETHER, HAVE FUN!'
      offset    :
        top     : 150
        left    : 700

    playMarker = @markers.create 'play',
      client    : '#home-page .laptop .teamwork'
      container : this
      wait      : 1900
      message   : 'INSTANTLY SPIN UP PLAYGROUNDS'
      offset    :
        top     : 275
        left    : 500

    logoMarker = @markers.create 'logo',
      client    : '#home-page .browser'
      container : this
      wait      : 2200
      message   : 'WHITE-LABEL KODING'
      offset    :
        top     : 25
        left    : 25


  pistachio:->

    """
      <header id='home-header'>
        <div class="inner-container">
          <a href="/" class="logo"><cite></cite></a>
          <a href="/" class="logotype">Koding</a>
          <a href="/Login" class="login fr">LOGIN</a>
        </div>
      </header>
      <main>
        <div class="clearfix">
          <div class="headings-container">
            <h1 class='big-header'>CODING ENVIRONMENT<br/>FROM THE FUTURE</h1>
            <h2>Social development in your browser, sign up to join a great community and code on powerful VMs.</h2>
          </div>
          <div class="register-container">
            {{> @registerForm}}
            <h3>{{> @githubLink}}</h3>
          </div>
        </div>
      </main>
      <figure class='laptop'>
        <section class='teamwork'></section>
      </figure>
      <section id='home-features' class='clearfix'>
        <div class='appstore clearfix'>
          <span class='icon'></span>
          <article>
            <h4>APPSTORE</h4>Speed up with user contributed apps, or create your own app, Koding has a great toolset to interact with VMs and to build UIs around.
          </article>
        </div>
        <div class='teamwork clearfix'>
          <span class='icon'></span>
          <article>
            <h4>TEAMWORK</h4>Collaborative development environment for lecture groups, pair programming, or simply for sharing what you're doing with a total stranger.
          </article>
        </div>
        <div class='social clearfix'>
          <span class='icon'></span>
          <article>
            <h4>SOCIAL</h4>Share with the community, learn from the experts or help those who have yet to start coding. Socialize with like minded people and have fun.
          </article>
        </div>
      </section>
      <section id='home-groups'>
        <h2 class='big-header'>GROUPS, HAVE YOUR OWN KODING</h2>
        <h3>Have all your development needs in a single private space.</h3>
        <figure class='education'></figure>
        <figure class='browser'></figure>
        <div class='group-features clearfix'>
          <div class='white-label clearfix'>
            <span class='icon'></span>
            <article>
              <h4>WHITE-LABEL KODING</h4>
              You can have your private Koding in the cloud, with your rules, your apps and your own members. Please <a id='home-contact-link' href='mailto:education@koding.com?subject=Koding, white label' target='_self'>contact us</a> for further information.
            </article>
          </div>
          <div class='school clearfix'>
            <span class='icon'></span>
            <article>
              <h4>USE IT IN YOUR SCHOOL</h4>
              Koding in the classroom, prepare your files online, share them with the whole class instantly. Collaborate live or just make your students watch what you're doing.
            </article>
          </div>
          <div class='project clearfix'>
            <span class='icon'></span>
            <article>
              <h4>CREATE PROJECT GROUPS</h4>
              Want to work on a project with your buddies and use the same resources and running instances, share a VM between your fellow developers.
            </article>
          </div>
        </div>
        {{> @pricingButton}}
      </section>
      <section id='home-bottom'>
        <h2 class='big-header'>IF YOU ARE READY TO GO, LET’S DO THIS</h2>
        <h3 class='hidden'>Something super simple and super descriptive goes here</h3>
        {{> @registerFormBottom}}
      </section>
      <footer class='clearfix'>
        <div class='fl'>
          <a href="/" class="logo"><cite></cite></a>
          <address>
          #{(new Date).getFullYear()} © Koding, Inc. </br>358 Brannan Street, San Francisco, CA, 94107
          </address>
        </div>
        <nav>
          <a class="hidden" href="/About">About</a>
          <a href="mailto:hello@koding.com">Contact</a>
          <a href="http://learn.koding.com/">University</a>
          <a href="http://koding.github.io/jobs/">Jobs</a>
          <a href="http://blog.koding.com">Blog</a>
        </nav>
      </footer>
    """

KD.introView = new HomePage
KD.introView.appendToDomBody()
