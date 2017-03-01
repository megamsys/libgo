package utils

const (

	//keys for watchers
	MAILGUN = "mailgun"
	SLACK   = "slack"
	INFOBIP = "infobip"
	SCYLLA  = "scylla"
	META    = "meta"
	WHMCS   = "WHMCS"
	SENDER  = "sender"

	SCYLLAMGR = "scylladb"
	//config keys by watchers
	TOKEN          = "token"
	CHANNEL        = "channel"
	USERNAME       = "username"
	PASSWORD       = "password"
	WHMCS_PASSWORD = "whmcs_password"
	WHMCS_APIKEY   = "whmcs_apikey"
	APPLICATION_ID = "application_id"
	MESSAGE_ID     = "message_id"
	API_KEY        = "api_key"
	DOMAIN         = "domain"
	PIGGYBANKS     = "piggybanks"
	VERTICE_EMAIL  = "vertice_email"
	VERTICE_APIKEY = "vertice_apikey"
	VERTICE_ORGID  = "vertice_orgid"
	USERMAIL       = "email"
	STATUS         = "status"
	HOST_IP        = "host_ip"
	VERTICEAPI     = "verticegateway"

	VONE_CLOUD     = "vone_cloud"
	HOME           = "home"
	DIR            = "dir"
	ORG_ID         = "org_id"
	MASTER_KEY     = "master_key"
	API_URL        = "url"
	SCYLLAHOST     = "scylla_host"
	SCYLLAKEYSPACE = "scylla_keyspace"
	SCYLLAUSERNAME = "scylla_username"
	SCYLLAPASSWORD = "scylla_password"

	ASSEMBLIES_ID  = "assemblies_id"
	EVENT_TYPE     = "event_type"
	ACCOUNT_ID     = "account_id"
	ASSEMBLY_ID    = "assembly_id"
	ASSEMBLY_NAME  = "assembly_name"
	QUOTA_ID       = "quota_id"
	DATA           = "data"
	CREATED_AT     = "created_at"
	REGION         = "region"
	DATASTORE      = "datastore"
	RAW_IMAGE_ID   = "raw_image_id"
	IMAGE_NAME     = "image_name"
	IMAGE_ID       = "image_id"
	INSTANCE_ID    = "instance_id"
	MARKETPLACE_ID = "marketplace_id"

	// events keys
	ACCOUNTID        = "AccountId"
	ASSEMBLYID       = "AssemblyId"
	ASSEMBLYNAME     = "AssemblyName"
	CONSUMED         = "Consumed"
	START_TIME       = "StartTime"
	END_TIME         = "EndTime"
	RESOURCES        = "Resources"
	BILL_TYPE        = "BillType"
	SOFT_ACTION      = "SoftAction"
	SOFT_LIMIT       = "SoftLimit"
	SOFT_GRACEPERIOD = "SoftGracePeriod"
	HARD_GRACEPERIOD = "HardGracePeriod"
	HARD_ACTION      = "HardAction"
	HARD_LIMIT       = "HardLimit"
	SKEWS_TYPE       = "SkewsType"
	QUOTAID          = "QuotaId"

	ACTION              = "action"
	ACTION_TRIGGERED_AT = "action_trigger_at"
	NEXT_ACTION_DUE_AT  = "next_action_due_at"
	NEXT_ACTION         = "next_action"

	//args for notification
	NILAVU       = "nilavu"
	LOGO         = "logo"
	NAME         = "name"
	VERTNAME     = "appname"
	TEAM         = "team"
	VERTTYPE     = "type"
	EMAIL        = "email"
	DAYS         = "days"
	COST         = "cost"
	STARTTIME    = "starttime"
	ENDTIME      = "endtime"
	IPV4PUB      = "ipv4public"
	IPV4PRI      = "ipv4private"
	IPV6PRI      = "ipv6private"
	IPV6PUB      = "ipv6public"
	STORAGE_TYPE = "storage_hddtype"
	STORAGE      = "storage"
	CPU          = "cpu"
	RAM          = "ram"

	YES              = "yes"
	FALSE            = "false"
	TRUE             = "true"
	ENABLED          = "enabled"
	EventMachine     = "machine"
	EventMarketplace = "marketplaces"
	EventContainer   = "container"
	EventBill        = "bill"
	EventUser        = "user"
	EventStatus      = "status"

	BILLMGR = "bill"
	ADDONS  = "addons"

	PROVIDER         = "provider"
	PROVIDER_ONE     = "one"
	PROVIDER_DOCKER  = "docker"
	PROVIDER_RANCHER = "rancher"

	LAUNCHING         = "launching"
	INITIALIZING      = "initializing"
	BALANCECHECK      = "balance_check"
	INSUFFICIENT_FUND = "insufficient_fund"
	INITIALIZED       = "initialized"
	LAUNCHED          = "launched"
	VMBOOTING         = "vm_booting"
	BOOTSTRAPPED      = "bootstrapped"
	BOOTSTRAPPING     = "bootstrapping"

	CONTAINERINITIALIZING = "containerinitializing"
	CONTAINERINITIALIZED  = "containerinitialized"

	STATEUPPING    = "stateup_starting"
	STATEUPPED     = "stateup_started"
	RUNNING        = "running"
	STARTING       = "starting"
	STARTED        = "started"
	STOPPING       = "stopping"
	STOPPED        = "stopped"
	RESTARTING     = "restarting"
	RESTARTED      = "restarted"
	UPGRADED       = "upgraded"
	DESTROYING     = "destroying"
	DESTROYED      = "destroyed"
	NUKED          = "nuked"
	SNAPSHOTTING   = "snapshot_creating"
	SNAPSHOTTED    = "snapshot_created"
	SNAPDESTORYING = "snapshot_deleting"
	SNAPDESTORYED  = "snapshot_deleted"
	RESTORINGSNAP  = "snapshot_restoring"
	SNAPRESTORED   = "snapshot_restored"
	ACTIVESNAP     = "active"
	DEACTIVESNAP   = "deactive"

	IMAGECREATING = "image_creating"
	IMAGECREATED  = "image_created"
	IMAGEREMOVED  = "image_removed"
	IMAGEREMOVING = "image_removing"

	DISKATTACHING = "disk_attaching"
	DISKATTACHED  = "disk_attached"
	DISKDETACHING = "disk_detaching"
	DISKDETACHED  = "disk_detached"

	QUOTAUPDATING = "quota_updating"
	QUOTAUPDATED  = "quota_updated"
	QUOTA_UNPAID  = "quota_unpaid"

	LCMSTATECHECK = "check_lcmstate"
	VMSTATECHECK  = "check_vmstate"
	WAITUNTILL    = "waituntill"

	VNCHOSTUPDATING    = "vnchostupdating"
	VNCHOSTUPDATED     = "vnchostupdated"
	DNSNETWORKCREATING = "dnscnamecreating"
	DNSNETWORKCREATED  = "dnscnamecreated"
	DNSNETWORKSKIPPED  = "dnscnameskipped"
	CLONING            = "gitcloning"
	CLONED             = "gitcloned"
	APPDEPLOYING       = "appdeploying"
	APPDEPLOYED        = "appdeployed"
	BUILDSTARTING      = "buildstarting"
	BUILDSTOPPED       = "buildstopped"
	SERVICESTARTING    = "servicestarting"
	SERVICESTOPPED     = "servicestopped"
	RESETPASSWORD      = "reset_password"

	COOKBOOKDOWNLOADING     = "cookbook_downloading"
	COOKBOOKDOWNLOADED      = "cookbook_downloaded"
	COOKBOOKFAILURE         = "cookbook_failure"
	AUTHKEYSUPDATING        = "authkeys_updating"
	AUTHKEYSUPDATED         = "authkeys_updated"
	AUTHKEYSFAILURE         = "authkeys_failure"
	INSTANCEIPSUPDATING     = "ips_updating"
	INSTANCEIPSUPDATED      = "ips_updated"
	INSTANCEIPSFAILURE      = "ips_failure"
	CHEFCONFIGSETUPSTARTING = "chefconfigsetup_starting"
	CHEFCONFIGSETUPSTARTED  = "chefconfigsetup_started"
	CONTAINERNETWORKSUCCESS = "container_network_success"
	CONTAINERNETWORKFAILURE = "container_network_failure"

	CONTAINERLAUNCHING     = "containerlaunching"
	CONTAINERBOOTSTRAPPING = "containerbootstrapping"
	CONTAINERBOOTSTRAPPED  = "containerbootstrapped"
	CONTAINERLAUNCHED      = "containerlaunched"
	CONTAINEREXISTS        = "containerexists"
	CONTAINERDELETE        = "containerdelete"
	CONTAINERSTARTING      = "containerstarting"
	CONTAINERSTARTED       = "containerstarted"
	CONTAINERSTOPPING      = "containerstopping"
	CONTAINERSTOPPED       = "containerstopped"
	CONTAINERRESTARTING    = "containerrestarting"
	CONTAINERRESTARTED     = "containerrestarted"
	CONTAINERUPGRADED      = "containerupgraded"
	CONTAINERRUNNING       = "containerrunning"
	CONTAINERERROR         = "containererror"
	DATABLOCK_CREATING     = "datablock_creating"
	DATABLOCK_CREATED      = "datablock_created"
	IMAGE_SAVING           = "image_saving"
	IMAGE_SAVED            = "image_saved"
	IMAGE_READY            = "image_ready"

	ERROR            = "error"
	PREERROR         = "preerror"
	NETWORK_ERROR    = "netwroking_error"
	VMLAUNCH_ERROR   = "launching_error"
	VMBOOT_ERROR     = "vmboot_error"
	BOOTSTRAPE_ERROR = "bootstrapping_error"
	STATEUP_ERROR    = "stateup_error"
	STATEDOWN_ERROR  = "statedown_error"
	VMRESUME_ERROR   = "vmresume_error"
	VMPOWEROFF_ERROR = "vmshutdown_error"

	PREDEPLOY_ERROR = "preerror"
	POST_ERROR      = "posterror"
	PARKED          = "parked"

	CREATING  = "creating"
	ACTIVE    = "active"
	DEACTIVE  = "deactive"
	PENDING   = "pending"
	INIT      = "init"
	HOLD      = "hold"
	SUSPENDED = "suspended"

	StateContainerInitializing = State(CONTAINERINITIALIZING)
	StateContainerInitialized  = State(CONTAINERINITIALIZED)
	StateContainerBootstrapped = State(CONTAINERBOOTSTRAPPED)
	StateContainerRunning      = State(CONTAINERRUNNING)
	StateDestroying            = State(DESTROYING)
	StateDestroyed             = State(DESTROYED)
	StateMachineParked         = State(PARKED)
	// StateLaunched is the milestone state for box after launched in cloud.
	StateInitializing = State(INITIALIZING)

	StateInitialized = State(INITIALIZED)

	// StateBootstrapped is the milestone state for box after being booted by the agent in cloud
	StateBootstrapped = State(BOOTSTRAPPED)
	//StateRunning is the milestone for fully up instance
	StateRunning = State(RUNNING)

	StateStopped = State(STOPPED)

	StatePreError  = State(PREDEPLOY_ERROR)
	StatePostError = State(POST_ERROR)

	// StatusLaunching is the initial status of a box
	// it should transition shortly to a more specific status
	StatusLaunching = Status(LAUNCHING)

	StatusBalanceVerified = Status(BALANCECHECK)

	// StatusLaunched is the status for box after launched in cloud.
	StatusLaunched = Status(LAUNCHED)

	StatusQuotaUpdating = Status(QUOTAUPDATING)
	StatusQuotaUpdated  = Status(QUOTAUPDATED)

	StatusInsufficientFund = Status(INSUFFICIENT_FUND)
	StatusQuotaUnpaid      = Status(QUOTA_UNPAID)

	StatusVMBooting = Status(VMBOOTING)

	// StatusBootstrapped is the status for box after being booted by the agent in cloud
	StatusBootstrapped  = Status(BOOTSTRAPPED)
	StatusBootstrapping = Status(BOOTSTRAPPING)

	// Stateup is the status of the which is moving up in the state in cloud.
	// Sent by vertice to gulpd when it received StatusBootstrapped.
	StatusStateupping = Status(STATEUPPING)
	StatusStateupped  = Status(STATEUPPED)

	StatusVncHostUpdating = Status(VNCHOSTUPDATING)
	StatusVncHostUpdated  = Status(VNCHOSTUPDATED)
	//fully up instance
	StatusRunning  = Status(RUNNING)
	StatusActive   = Status(ACTIVE)
	StatusDeactive = Status(DEACTIVE)
	StatusCreating = Status(CREATING)

	StatusResetPassword = Status(RESETPASSWORD)
	StatusStarting      = Status(STARTING)
	StatusStarted       = Status(STARTED)
	StatusStopping      = Status(STOPPING)
	StatusStopped       = Status(STOPPED)
	StatusRestarting    = Status(RESTARTING)
	StatusRestarted     = Status(RESTARTED)
	StatusUpgraded      = Status(UPGRADED)
	StatusDestroying    = Status(DESTROYING)
	StatusDestroyed     = Status(DESTROYED)
	StatusNuked         = Status(NUKED)

	StatusSnapCreating   = Status(SNAPSHOTTING)
	StatusSnapCreated    = Status(SNAPSHOTTED)
	StatusSnapDeleting   = Status(SNAPDESTORYING)
	StatusSnapDeleted    = Status(SNAPDESTORYED)
	StatusSnapRestoring  = Status(RESTORINGSNAP)
	StatusSnapRestored   = Status(SNAPRESTORED)
	StatusBackupCreated  = Status(IMAGECREATED)
	StatusBackupCreating = Status(IMAGECREATING)
	StatusBackupDeleting = Status(IMAGEREMOVING)
	StatusBackupDeleted  = Status(IMAGEREMOVED)

	StatusDiskAttaching    = Status(DISKATTACHING)
	StatusDiskAttached     = Status(DISKATTACHED)
	StatusDiskDetaching    = Status(DISKDETACHING)
	StatusDiskDetached     = Status(DISKDETACHED)
	StatusLcmStateChecking = Status(LCMSTATECHECK)
	StatusWaitUntill       = Status(WAITUNTILL)
	StatusVmStateChecking  = Status(VMSTATECHECK)

	StatusCookbookDownloading = Status(COOKBOOKDOWNLOADING)
	StatusCookbookDownloaded  = Status(COOKBOOKDOWNLOADED)
	StatusCookbookFailure     = Status(COOKBOOKFAILURE)
	StatusAuthkeysUpdating    = Status(AUTHKEYSUPDATING)
	StatusAuthkeysUpdated     = Status(AUTHKEYSUPDATED)
	StatusAuthkeysFailure     = Status(AUTHKEYSFAILURE)

	StatusIpsUpdating         = Status(INSTANCEIPSUPDATING)
	StatusIpsUpdated          = Status(INSTANCEIPSUPDATED)
	StatusIpsFailure          = Status(INSTANCEIPSFAILURE)
	StatusChefConfigSetupping = Status(CHEFCONFIGSETUPSTARTING)
	StatusChefConfigSetupped  = Status(CHEFCONFIGSETUPSTARTED)
	StatusAppDeploying        = Status(APPDEPLOYING)
	StatusAppDeployed         = Status(APPDEPLOYED)

	StatusNetworkCreating = Status(DNSNETWORKCREATING)
	StatusNetworkCreated  = Status(DNSNETWORKCREATED)
	StatusNetworkSkipped  = Status(DNSNETWORKSKIPPED)
	StatusCloning         = Status(CLONING)
	StatusCloned          = Status(CLONED)
	StatusBuildStarting   = Status(BUILDSTARTING)
	StatusBuildStoped     = Status(BUILDSTOPPED)
	StatusServiceStarting = Status(SERVICESTARTING)
	StatusServiceStopped  = Status(SERVICESTOPPED)

	StatusContainerLaunching      = Status(CONTAINERLAUNCHING)
	StatusContainerBootstrapping  = Status(CONTAINERBOOTSTRAPPING)
	StatusContainerBootstrapped   = Status(CONTAINERBOOTSTRAPPED)
	StatusContainerLaunched       = Status(CONTAINERLAUNCHED)
	StatusContainerExists         = Status(CONTAINEREXISTS)
	StatusContainerDelete         = Status(CONTAINERDELETE)
	StatusContainerStarting       = Status(CONTAINERSTARTING)
	StatusContainerStarted        = Status(CONTAINERSTARTED)
	StatusContainerStopping       = Status(CONTAINERSTOPPING)
	StatusContainerStopped        = Status(CONTAINERSTOPPED)
	StatusContainerRestarting     = Status(CONTAINERRESTARTING)
	StatusContainerRestarted      = Status(CONTAINERRESTARTED)
	StatusContainerUpgraded       = Status(CONTAINERUPGRADED)
	StatusContainerRunning        = Status(CONTAINERRUNNING)
	StatusContainerNetworkSuccess = Status(CONTAINERNETWORKSUCCESS)
	StatusContainerNetworkFailure = Status(CONTAINERNETWORKFAILURE)
	StatusContainerError          = Status(CONTAINERERROR)

	// StatusError is the status for units that failed to start, because of
	// a box error.
	StatusError          = Status(ERROR)
	StatusPreError       = Status(PREERROR)
	StatusNetworkError   = Status(NETWORK_ERROR)
	StatusVmLaunchError  = Status(VMLAUNCH_ERROR)
	StatusVmBootError    = Status(VMBOOT_ERROR)
	StatusBootstrapError = Status(BOOTSTRAPE_ERROR)
	StatusStateupError   = Status(STATEUP_ERROR)
	StatusStatedownError = Status(STATEDOWN_ERROR)
	StatusVmStartError   = Status(VMRESUME_ERROR)
	StatusVmStopError    = Status(VMPOWEROFF_ERROR)

	StatusDataBlockCreating = Status(DATABLOCK_CREATING)
	StatusDataBlockCreated  = Status(DATABLOCK_CREATED)
	StatusImageSaving       = Status(IMAGE_SAVING)
	StatusImageSaved        = Status(IMAGE_SAVED)
	StatusImageReady        = Status(IMAGE_READY)

	ONEINSTANCELAUNCHINGTYPE           = "compute.instance.launching"
	ONEINSTANCEBOOTINGTYPE             = "compute.instance.booting"
	ONEINSTANCEVNCHOSTUPDATING         = "compute.instance.vnchostupdating"
	ONEINSTANCEVNCHOSTUPDATED          = "compute.instance.vnchostupdated"
	ONEINSTANCECHEFCONFIGSETUPSTARTING = "compute.instance.chefconfigsetupstarting"
	ONEINSTANCECHEFCONFIGSETUPSTARTED  = "compute.instance.chefconfigsetupstarted"
	ONEINSTANCEGITCLONING              = "compute.instance.gitcloning"
	ONEINSTANCEGITCLONED               = "compute.instance.gitcloned"
	ONEINSTANCEAPPDEPLOYING            = "compute.instance.appdeploying"
	ONEINSTANCEAPPDEPLOYED             = "compute.instance.appdeployed"
	ONEINSTANCEBUILDSTARTING           = "compute.instance.buildstarting"
	ONEINSTANCEBUILDSTOPPED            = "compute.instance.buildstopped"
	ONEINSTANCESERVICESTARTING         = "compute.instance.servicestarting"
	ONEINSTANCESERVICESTOPPED          = "compute.instance.servicestopped"
	ONEINSTANCEDNSCNAMING              = "compute.instance.dnscnaming"
	ONEINSTANCEDNSCNAMED               = "compute.instance.dnscnamed"
	ONEINSTANCEDNSNETWORKSKIPPED       = "compute.instance.dnscnameskipped"
	ONEINSTANCEBOOTSTRAPPINGTYPE       = "compute.instance.bootstrapping"
	ONEINSTANCEBOOTSTRAPPEDTYPE        = "compute.instance.bootstrapped"
	ONEINSTANCESTATEUPPINGTYPE         = "compute.instance.stateupstarting"
	ONEINSTANCESTATEUPPEDTYPE          = "compute.instance.stateupstarted"
	ONEINSTANCERUNNINGTYPE             = "compute.instance.running"
	ONEINSTANCELAUNCHEDTYPE            = "compute.instance.launched"
	ONEINSTANCEEXISTSTYPE              = "compute.instance.exists"
	ONEINSTANCEDESTROYINGTYPE          = "compute.instance.destroying"
	ONEINSTANCEDELETEDTYPE             = "compute.instance.deleted"
	ONEINSTANCESTARTINGTYPE            = "compute.instance.starting"
	ONEINSTANCESTARTEDTYPE             = "compute.instance.started"
	ONEINSTANCESTOPPINGTYPE            = "compute.instance.stopping"
	ONEINSTANCESTOPPEDTYPE             = "compute.instance.stopped"
	ONEINSTANCERESTARTINGTYPE          = "compute.instance.restarting"
	ONEINSTANCERESTARTEDTYPE           = "compute.instance.restarted"
	ONEINSTANCEUPGRADEDTYPE            = "compute.instance.upgraded"
	ONEINSTANCERESIZINGTYPE            = "compute.instance.resizing"
	ONEINSTANCERESIZEDTYPE             = "compute.instance.resized"
	ONEINSTANCEERRORTYPE               = "compute.instance.posterror"
	ONEINSTANCEPREERRORTYPE            = "compute.instance.preerror"
	ONEINSTANCESNAPSHOTTINGTYPE        = "compute.instance.snapshotting"
	ONEINSTANCESNAPSHOTTEDTYPE         = "compute.instance.snapshotted"
	ONEINSTANCESBALANCEVERIFYTYPE      = "compute.instance.balance_check"
	ONEINSTANCESINSUFFIENTFUNDTYPE     = "compute.instance.insufficient_fund"
	ONEINSTANCEUSERQUOTAUPDATING       = "compute.instance.quota_updating"
	ONEINSTANCEUSERQUOTAUPDATED        = "compute.instance.quota_updated"
	ONEINSTANCELCMSTATECHECKING        = "compute.instance.lcmstate_checking"
	ONEINSTANCEVMSTATECHECKING         = "compute.instance.vmstate_checking"
	ONEINSTANCESQUOTAUNPAID            = "compute.instance.quota_unpaid"

	ONEINSTANCEWAITING        = "compute.instance.waituntill"
	ONEINSTANCEVMSTATEPENDING = "compute.instance.pending"
	ONEINSTANCEVMSTATEPROLOG  = "compute.instance.active"
	ONEINSTANCEVMSTATEHOLD    = "compute.instance.hold"

	ONEINSTANCELCMMSTATEPROLOG = "compute.instance.active_prolog"
	ONEINSTANCELCMSTATEBOOT    = "compute.instance.active_boot"
	ONEINSTANCELCMSTATERUNNING = "compute.instance.active_running"
	ONEINSTANCELCMSTATEINIT    = "compute.instance.active_lcm_init"

	COOKBOOKDOWNLOADINGTYPE  = "compute.instance.cookbook_downloading"
	COOKBOOKDOWNLOADEDTYPE   = "compute.instance.cookbook_downloaded"
	COOKBOOKFAILURETYPE      = "compute.instance.cookbook_download_failure"
	AUTHKEYSUPDATINGTYPE     = "compute.instance.authkeysupdating"
	AUTHKEYSUPDATEDTYPE      = "compute.instance.authkeysupdated"
	AUTHKEYSFAILURETYPE      = "compute.instance.authkeysfailure"
	INSTANCEIPSUPDATINGTYPE  = "compute.instance.ip_updating"
	INSTANCEIPSUPDATEDTYPE   = "compute.instance.ip_updated"
	INSTANCEIPSFAILURETYPE   = "compute.instance.ip_updatefailure"
	INSTANCERESETOLDPASSWORD = "compute.instance.reset_password"

	CONTAINERINSTANCELAUNCHINGTYPE = "compute.container.launching"
	CONTAINERINSTANCEBOOTSTRAPPING = "compute.container.bootstrapping"
	CONTAINERINSTANCEBOOTSTRAPPED  = "compute.container.bootstrapped"
	CONTAINERINSTANCELAUNCHEDTYPE  = "compute.container.launched"
	CONTAINERINSTANCEEXISTS        = "compute.container.exists"
	CONTAINERINSTANCEDELETE        = "compute.container.delete"
	CONTAINERINSTANCESTARTING      = "compute.container.starting"
	CONTAINERINSTANCESTARTED       = "compute.container.started"
	CONTAINERINSTANCESTOPPING      = "compute.container.stopping"
	CONTAINERINSTANCESTOPPED       = "compute.container.stopped"
	CONTAINERINSTANCERESTARTING    = "compute.container.restarting"
	CONTAINERINSTANCERESTARTED     = "compute.container.restarted"
	CONTAINERINSTANCEUPGRADED      = "compute.container.upgraded"
	CONTAINERINSTANCERUNNING       = "compute.container.running"
	CONTAINERNETWORKSUCCESSTYPE    = "compute.container.ip_allocate_success"
	CONTAINERNETWORKFAILURETYPE    = "compute.container.ip_allocate_failure"
	CONTAINERINSTANCEERROR         = "compute.container.posterror"

	MARKETPLACEINSTANCELAUNCHINGTYPE = "marketplaces.instance.launching"
	MARKETPLACEINSTANCELAUNCHEDTYPE  = "marketplaces.instance.launched"
	MARKETPLACEBLOCKCREATING         = "marketplaces.instance.block_creating"
	MARKETPLACEBLOCKCREATED          = "marketplaces.instance.block_created"
	MARKETPLACELCMSTATECHECKING      = "marketplaces.instance.lcmstate_checking"
	MARKETPLACEVMSTATECHECKING       = "marketplaces.instance.vmstate_checking"
	MARKETPLACEWAITING               = "marketplaces.instance.waituntill"
	MARKETPLACEVMSTATEPENDING        = "marketplaces.instance.pending"
	MARKETPLACEVMSTATEPROLOG         = "marketplaces.instance.active"
	MARKETPLACEVMSTATEHOLD           = "marketplaces.instance.hold"
	MARKETPLACELCMSTATEPROLOG        = "marketplaces.instance.active_prolog"
	MARKETPLACELCMSTATEBOOT          = "marketplaces.instance.active_boot"
	MARKETPLACELCMSTATERUNNING       = "marketplaces.instance.active_running"
	MARKETPLACELCMSTATEINIT          = "marketplaces.instance.active_lcm_init"
	MARKETPLACEBOOTINGTYPE           = "marketplaces.instance.booting"
	MARKETPLACEVNCHOSTUPDATING       = "marketplaces.instance.vnchostupdating"
	MARKETPLACEVNCHOSTUPDATED        = "marketplaces.instance.vnchostupdated"
	MARKETPLACEPREERRORTYPE          = "marketplaces.instance.preerror"
	MARKETPLACEIMAGESAVING           = "marketplaces.image.saving"
	MARKETPLACEIMAGESAVED            = "marketplaces.image.saved"
	MARKETPLACEIMAGEREADY            = "marketplaces.image.ready"
)
