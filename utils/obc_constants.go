package utils

const (
  	HOST_IP     = "host_ip"

    FETCHINGHOSTINFOS = "hostinfos.starting"
    FETCHEDHOSTINFOS = "hostinfos.completed"
    FETCHINGHOSTCPU    = "hostinfos.cpu.starting"
    FETCHINGHOSTRAM    = "hostinfos.memory.starting"
    FETCHINGHOSTOS     = "hostinfos.os.starting"
    FETCHINGHOSTDISK   = "hostinfos.disk.starting"
    FETCHINGHOSTNAME   = "hostinfos.hostname.starting"
    FETCHEDHOSTCPU    = "hostinfos.cpu.completed"
    FETCHEDHOSTRAM    = "hostinfos.memory.completed"
    FETCHEDHOSTOS     = "hostinfos.os.completed"
    FETCHEDHOSTDISK   = "hostinfos.disk.completed"
    FETCHEDHOSTNAME   = "hostinfos.hostname.completed"

    INSTALLINGNILAVU = "nilavuinstall"
    INSTALLINGGATEWAY = "gatewayinstall"
    INSTALLINGVERTICE = "verticeinstall"
    INSTALLINGGULPD = "gulpdinstall"
    INSTALLINGCOMMON = "commoninstall"
    INSTALLINGCASSSANDRA = "cassandrainstall"
    INSTALLEDNILAVU = "nilavuinstalled"
    INSTALLEDGATEWAY = "gatewayinstalled"
    INSTALLEDVERTICE = "verticeinstalled"
    INSTALLEDGULPD = "gulpdinstalled"
    INSTALLEDCOMMON = "commoninstalled"
    INSTALLEDCASSSANDRA = "cassandrainstalled"


    //Getting informations (cpu, ram, disk and etc.,) about the Host
    //HOSTINFOS templates statuses
    StatusGettingInfos = Status(FETCHINGHOSTINFOS)
    StatusGotInfos = Status(FETCHEDHOSTINFOS)
    StatusGettingCPU   = Status(FETCHINGHOSTCPU)
    StatusGettingRAM   = Status(FETCHINGHOSTRAM)
    StatusGettingHostName = Status(FETCHINGHOSTNAME)
    StatusGettingDisk  = Status(FETCHINGHOSTDISK)
    StatusGettingOS    = Status(FETCHINGHOSTDISK)
    StatusGotCPU   =  Status(FETCHEDHOSTCPU)
    StatusGotRAM   = Status(FETCHEDHOSTRAM)
    StatusGotHostName =  Status(FETCHEDHOSTNAME)
    StatusGotDisk  =Status(FETCHEDHOSTDISK)
    StatusGotOS    = Status(FETCHEDHOSTOS)


    //Installing Vertice Packages statuses



   //Templates HostInfos status types
    OBCHOSTINFOSFETCHING =  "obc.hostinfos.starting"
    OBCHOSTINFOSCPUPARSING = "obc.hostinfos.cpu.parsing"
    OBCHOSTINFOSCPUPARSED = "obc.hostinfos.cpu.parsed"
    OBCHOSTINFOSRAMPARSING = "obc.hostinfos.ram.parsing"
    OBCHOSTINFOSRAMPARSED = "obc.hostinfos.ram.parsed"
    OBCHOSTINFOSOSPARSING = "obc.hostinfos.os.parsing"
    OBCHOSTINFOSOSPARSED = "obc.hostinfos.os.parsed"
    OBCHOSTINFOSDISKPARSING = "obc.hostinfos.disk.parsing"
    OBCHOSTINFOSDISKPARSED = "obc.hostinfos.disk.parsed"
    OBCHOSTINFOSHOSTNAMEPARSING = "obc.hostinfos.hostname.parsing"
    OBCHOSTINFOSHOSTNAMEPARSED = "obc.hostinfos.hostname.parsed"
    OBCHOSTINFOSSUCCESS  =  "obc.hostinfos.complated.successfully"
)
