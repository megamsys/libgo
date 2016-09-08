package utils

const (

	//keys for watchers
	MAILGUN = "mailgun"
	SLACK   = "slack"
	INFOBIP = "infobip"
	SCYLLA  = "scylla"
	META    = "meta"
	WHMCS   = "whmcs"
	SENDER  = "sender"

	//config keys by watchers
	TOKEN          = "token"
	CHANNEL        = "channel"
	USERNAME       = "username"
	PASSWORD       = "password"
	APPLICATION_ID = "application_id"
	MESSAGE_ID     = "message_id"
	API_KEY        = "api_key"
	DOMAIN         = "domain"
	PIGGYBANKS     = "piggybanks"
	VERTICE_EMAIL  = "vertice_email"
	VERTICE_APIKEY = "vertice_apikey"
	USERMAIL       = "email"
	STATUS         = "status"


	HOME           = "home"
	DIR            = "dir"
	SCYLLAHOST     = "scylla_host"
	SCYLLAKEYSPACE = "scylla_keyspace"

	EVENT_TYPE  = "event_type"
	ACCOUNT_ID  = "account_id"
	ASSEMBLY_ID = "assembly_id"
	DATA        = "data"
	CREATED_AT  = "created_at"


	//args for notification
	NILAVU    = "nilavu"
	LOGO      = "logo"
	NAME      = "name"
	VERTNAME  = "appname"
	TEAM      = "team"
	VERTTYPE  = "type"
	EMAIL     = "email"
	DAYS      = "days"
	COST      = "cost"
	STARTTIME = "starttime"
	ENDTIME   = "endtime"
	//STATUS    = "status"
	//DESCRIPTION = "description"
	IPV4PUB = "ipv4public"
	IPV4PRI = "ipv4private"
	IPV6PRI = "ipv6private"
	IPV6PUB = "ipv6public"

	EventMachine   = "machine"
	EventContainer = "container"
	EventBill      = "bill"
	EventUser      = "user"
	EventStatus    = "status"

	BILLMGR = "bill"
	ADDONS  = "addons"

	PROVIDER        = "provider"
	PROVIDER_ONE    = "one"
	PROVIDER_DOCKER = "docker"

	LAUNCHING     = "launching"
	LAUNCHED      = "launched"
	BOOTSTRAPPED  = "bootstrapped"
	BOOTSTRAPPING = "bootstrapping"

	STATEUPPING   = "stateup_starting"
	STATEUPPED    = "stateup_started"
	RUNNING      = "running"
	STARTING     = "starting"
	STARTED      = "started"
	STOPPING     = "stopping"
	STOPPED      = "stopped"
	RESTARTING     = "restarting"
	RESTARTED      = "restarted"
	UPGRADED     = "upgraded"
	DESTROYING   = "destroying"
	NUKED        = "nuked"
	ERROR        = "error"
	SNAPSHOTTING   = "snapshotting"
	SNAPSHOTTED    = "snapshotted"

	VNCHOSTUPDATING = "vnchostupdating"
	VNCHOSTUPDATED = "vnchostupdated"
	DNSNETWORKCREATING = "dnscnamecreating"
	DNSNETWORKCREATED  = "dnscnamecreated"
	DNSNETWORKSKIPPED  = "dnscnameskipped"
	CLONING = "gitcloning"
	CLONED = "gitcloned"
	CHEFSOLOSTARTING = "chefsolostarting"
	CHEFSOLOFINISHED = "chefsolofinished"
	BUILDSTARTING = "buildstarting"
	BUILDSTOPPED = "buildstopped"
	SERVICESTARTING = "servicestarting"
	SERVICESTOPPED = "servicestopped"

	COOKBOOKDOWNLOADING = "cookbook_downloading"
	COOKBOOKDOWNLOADED = "cookbook_downloaded"
	COOKBOOKFAILURE = "cookbook_failure"
	AUTHKEYSUPDATING = "authkeys_updating"
	AUTHKEYSUPDATED = "authkeys_updated"
	AUTHKEYSFAILURE = "authkeys_failure"
	INSTANCEIPSUPDATING = "ips_updating"
	INSTANCEIPSUPDATED = "ips_updated"
	INSTANCEIPSFAILURE = "ips_failure"
	CHEFCONFIGSETUPSTARTING = "chefconfigsetup_starting"
	CHEFCONFIGSETUPSTARTED = "chefconfigsetup_started"
	CONTAINERNETWORKSUCCESS = "container_network_success"
	CONTAINERNETWORKFAILURE = "container_network_failure"



	// StatusLaunching is the initial status of a box
	// it should transition shortly to a more specific status
	StatusLaunching = Status(LAUNCHING)

	// StatusLaunched is the status for box after launched in cloud.
	StatusLaunched = Status(LAUNCHED)

	// StatusBootstrapped is the status for box after being booted by the agent in cloud
	StatusBootstrapped  = Status(BOOTSTRAPPED)
	StatusBootstrapping = Status(BOOTSTRAPPING)

	// Stateup is the status of the which is moving up in the state in cloud.
	// Sent by vertice to gulpd when it received StatusBootstrapped.
	StatusStateupping = Status(STATEUPPING)
	StatusStateupped = Status(STATEUPPED)

	StatusVncHostUpdating = Status(VNCHOSTUPDATING)
	StatusVncHostUpdated = Status(VNCHOSTUPDATED)
	//fully up instance
	StatusRunning = Status(RUNNING)

	StatusStarting = Status(STARTING)
	StatusStarted  = Status(STARTED)

	StatusStopping = Status(STOPPING)
	StatusStopped  = Status(STOPPED)

	StatusRestarting = Status(RESTARTING)
	StatusRestarted  = Status(RESTARTED)

	StatusUpgraded = Status(UPGRADED)

	StatusDestroying = Status(DESTROYING)
	StatusNuked      = Status(NUKED)

	StatusSnapCreating = Status(SNAPSHOTTING)
	StatusSnapCreated  = Status(SNAPSHOTTED)

	// StatusError is the status for units that failed to start, because of
	// a box error.
	StatusError = Status(ERROR)

	StatusCookbookDownloading = Status(COOKBOOKDOWNLOADING)
	StatusCookbookDownloaded = Status(COOKBOOKDOWNLOADED)
	StatusCookbookFailure = Status(COOKBOOKFAILURE)
	StatusAuthkeysUpdating = Status(AUTHKEYSUPDATING)
	StatusAuthkeysUpdated = Status(AUTHKEYSUPDATED)
	StatusAuthkeysFailure = Status(AUTHKEYSFAILURE)

	StatusIpsUpdating = Status(INSTANCEIPSUPDATING)
	StatusIpsUpdated = Status(INSTANCEIPSUPDATED)
	StatusIpsFailure = Status(INSTANCEIPSFAILURE)
	StatusChefConfigSetupping = Status(CHEFCONFIGSETUPSTARTING)
	StatusChefConfigSetupped = Status(CHEFCONFIGSETUPSTARTED)
	StatusChefsoloStarting = Status(CHEFSOLOSTARTING)
	StatusChefsoloFinished = Status(CHEFSOLOFINISHED)

	StatusNetworkCreating = Status(DNSNETWORKCREATING)
	StatusNetworkCreated  = Status(DNSNETWORKCREATED)
	StatusNetworkSkipped  = Status(DNSNETWORKSKIPPED)
	StatusCloning = Status(CLONING)
	StatusCloned  = Status(CLONED)
	StatusBuildStarting = Status(BUILDSTARTING)
	StatusBuildStoped = Status(BUILDSTOPPED)
	StatusServiceStarting = Status(SERVICESTARTING)
	StatusServiceStopped = Status(SERVICESTOPPED)

	StatusContainerNetworkSuccess = Status(CONTAINERNETWORKSUCCESS)
	StatusContainerNetworkFailure = Status(CONTAINERNETWORKFAILURE)

	ONEINSTANCELAUNCHINGTYPE = "compute.instance.launching"
	ONEINSTANCEVNCHOSTUPDATING = "compute.instance.vnchostupdating"
	ONEINSTANCEVNCHOSTUPDATED = "compute.instance.vnchostupdated"
	ONEINSTANCECHEFCONFIGSETUPSTARTING = "compute.instance.chefconfigsetupstarting"
	ONEINSTANCECHEFCONFIGSETUPSTARTED = "compute.instance.chefconfigsetupstarted"
	ONEINSTANCEGITCLONING = "compute.instance.gitcloning"
	ONEINSTANCEGITCLONED = "compute.instance.gitcloned"
	ONEINSTANCECHEFSOLOSTARTING = "compute.instance.chefsolostarting"
	ONEINSTANCECHEFSOLOFINISHED = "compute.instance.chefsolofinished"
	ONEINSTANCEBUILDSTARTING = "compute.instance.buildstarting"
	ONEINSTANCEBUILDSTOPPED = "compute.instance.buildstopped"
	ONEINSTANCESERVICESTARTING = "compute.instance.servicestarting"
	ONEINSTANCESERVICESTOPPED = "compute.instance.servicestopped"
	ONEINSTANCEDNSCNAMING    = "compute.instance.dnscnaming"
	ONEINSTANCEDNSCNAMED     = "compute.instance.dnscnamed"
	ONEINSTANCEDNSNETWORKSKIPPED = "compute.instance.dnscnameskipped"
	ONEINSTANCEBOOTSTRAPPINGTYPE = "compute.instance.bootstrapping"
	ONEINSTANCEBOOTSTRAPPEDTYPE  = "compute.instance.bootstrapped"
	ONEINSTANCESTATEUPPINGTYPE   = "compute.instance.stateupstarting"
	ONEINSTANCESTATEUPPEDTYPE    = "compute.instance.stateupstarted"
	ONEINSTANCERUNNINGTYPE       = "compute.instance.running"
	ONEINSTANCELAUNCHEDTYPE      = "compute.instance.launched"
	ONEINSTANCEEXISTSTYPE        = "compute.instance.exists"
	ONEINSTANCEDESTROYINGTYPE    = "compute.instance.destroying"
	ONEINSTANCEDELETEDTYPE       = "compute.instance.deleted"
	ONEINSTANCESTARTINGTYPE      = "compute.instance.starting"
	ONEINSTANCESTARTEDTYPE       = "compute.instance.started"
	ONEINSTANCESTOPPINGTYPE      = "compute.instance.stopping"
	ONEINSTANCESTOPPEDTYPE       = "compute.instance.stopped"
	ONEINSTANCERESTARTINGTYPE    = "compute.instance.restarting"
	ONEINSTANCERESTARTEDTYPE     = "compute.instance.restarted"
	ONEINSTANCEUPGRADEDTYPE      = "compute.instance.upgraded"
	ONEINSTANCERESIZINGTYPE      = "compute.instance.resizing"
	ONEINSTANCERESIZEDTYPE       = "compute.instance.resized"
	ONEINSTANCEERRORTYPE         = "compute.instance.error"
	ONEINSTANCESNAPSHOTTINGTYPE  = "compute.instance.snapshotting"
	ONEINSTANCESNAPSHOTTEDTYPE   = "compute.instance.snapshotted"

	COOKBOOKDOWNLOADINGTYPE      = "compute.instance.cookbook_downloading"
	COOKBOOKDOWNLOADEDTYPE      = "compute.instance.cookbook_downloaded"
	COOKBOOKFAILURETYPE          = "compute.instance.cookbook_download_failure"
	AUTHKEYSUPDATINGTYPE         = "compute.instance.authkeysupdating"
	AUTHKEYSUPDATEDTYPE          = "compute.instance.authkeysupdated"
	AUTHKEYSFAILURETYPE          = "compute.instance.authkeysfailure"
	INSTANCEIPSUPDATINGTYPE       = "compute.instance.ip_updating"
	INSTANCEIPSUPDATEDTYPE       = "compute.instance.ip_updated"
	INSTANCEIPSFAILURETYPE       = "compute.instance.ip_updatefailure"

	CONTAINERNETWORKSUCCESSTYPE      = "net.container.ip_allocate_success"
	CONTAINERNETWORKFAILURETYPE      = "net.container.ip_allocate_failure"
)
