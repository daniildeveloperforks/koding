# TODO: Should implement non-collaborative finder
# TODO: Should extend this class from NonCollab one.

class CollaborativeFinderPane extends CollaborativePane

  constructor: (options = {}, data) ->

    options.cssClass = "finder-pane nfinder file-container"

    super options, data

    panel             = @getDelegate()
    workspace         = panel.getDelegate()
    @sessionKey       = @getOptions().sessionKey or @createSessionKey()
    @workspaceRef     = workspace.firepadRef.child @sessionKey

    @finderController = new NFinderController
      nodeIdPath          : "path"
      nodeParentIdPath    : "parentPath"
      contextMenu         : no
      loadFilesOnInit     : yes
      useStorage          : yes
      treeControllerClass : CollaborativeFinderTreeController

    @finderController.reset()

    @finder = @finderController.getView()

    @workspaceRef.on "value", (snapshot) =>
      log "everything is something happened in host filetree", snapshot.name(), snapshot.val()

      clientData = snapshot.val()?.ClientWantsToInteractWithRemoteFileTree
      if clientData
        path             = "[#{clientData.vmName}]#{clientData.path}"
        {treeController} = @finderController
        nodeView         = treeController.nodes[path]

        treeController.openItem nodeView, (err, res) =>
          log "Host terminal done with client request", res

    # event bindings

    @finderController.on "FileTreeInteractionDone", (files) =>
      @syncContent files

    @finderController.on "OpenedAFile", (file, content) =>
      for pane in panel.panes
        if pane instanceof CollaborativeEditorPane or pane instanceof CollaborativeTabbedEditorPane
          editorPane = pane

      return  warn "could not find an editor instance to set file content" unless editorPane

      editorPane.openFile file, content

    log "i'm a host file tree and my session key is #{@sessionKey}"

  syncContent: (files) ->
    @workspaceRef.set { files }

  pistachio: ->
    """
      {{> @header}}
      {{> @finder}}
    """




class CollaborativeFinderTreeController extends NFinderTreeController

  constructor: (options = {}, data) ->

    super options, data

  addNodes: (nodes) ->
    super nodes
    @syncInteraction()

  dblClick: (nodeView, e) ->
    super nodeView, e
    log "host interacted with file tree waiting for response"

  getSnapshot: ->
    snapshot = []

    for path, node of @nodes
      nodeData = node.data

      snapshot.push
        path   : FSHelper.plainPath path
        type   : nodeData.type
        vmName : nodeData.vmName
        name   : nodeData.name

    log "finder snapshot is", snapshot

    return snapshot

  syncInteraction: ->
    @getDelegate().emit "FileTreeInteractionDone", @getSnapshot()

  toggleFolder: (nodeView, callback) ->
    super nodeView, @bound "syncInteraction"

  openFile: (nodeView) ->
    return unless nodeView
    file = nodeView.getData()
    log "host terminal is opening a file", file
    file.fetchContents (err, contents) =>
      @getDelegate().emit "OpenedAFile", file, contents
