/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package utils

// Status represents the status of a unit in vertice
type Status string

func (s Status) String() string {
	return string(s)
}

type State string

func (s State) String() string {
	return string(s)
}

func (s Status) Event_type() string {
	switch s.String() {
	case LAUNCHING:
		return ONEINSTANCELAUNCHINGTYPE
	case BALANCECHECK:
		return ONEINSTANCESBALANCEVERIFYTYPE
	case INSUFFICIENT_FUND:
		return ONEINSTANCESINSUFFIENTFUNDTYPE
	case QUOTAUPDATING:
		return ONEINSTANCEUSERQUOTAUPDATING
	case QUOTAUPDATED:
		return ONEINSTANCEUSERQUOTAUPDATED
	case VMBOOTING:
		return ONEINSTANCEBOOTINGTYPE
	case LAUNCHED:
		return ONEINSTANCELAUNCHEDTYPE
	case BOOTSTRAPPING:
		return ONEINSTANCEBOOTSTRAPPINGTYPE
	case BOOTSTRAPPED:
		return ONEINSTANCEBOOTSTRAPPEDTYPE
	case STATEUPPING:
		return ONEINSTANCESTATEUPPINGTYPE
	case STATEUPPED:
		return ONEINSTANCESTATEUPPEDTYPE
	case RUNNING:
		return ONEINSTANCERUNNINGTYPE
	case STARTING:
		return ONEINSTANCESTARTINGTYPE
	case STARTED:
		return ONEINSTANCESTARTEDTYPE
	case STOPPING:
		return ONEINSTANCESTOPPINGTYPE
	case STOPPED:
		return ONEINSTANCESTOPPEDTYPE
	case UPGRADED:
		return ONEINSTANCEUPGRADEDTYPE
	case DESTROYING:
		return ONEINSTANCEDESTROYINGTYPE
	case NUKED:
		return ONEINSTANCEDELETEDTYPE
	case SNAPSHOTTING:
		return ONEINSTANCESNAPSHOTTINGTYPE
	case SNAPSHOTTED:
		return ONEINSTANCESNAPSHOTTEDTYPE
	case COOKBOOKDOWNLOADING:
		return COOKBOOKDOWNLOADINGTYPE
	case COOKBOOKDOWNLOADED:
		return COOKBOOKDOWNLOADEDTYPE
	case COOKBOOKFAILURE:
		return COOKBOOKFAILURETYPE
	case APPDEPLOYING:
		return ONEINSTANCEAPPDEPLOYING
	case APPDEPLOYED:
		return ONEINSTANCEAPPDEPLOYED
	case VNCHOSTUPDATING:
		return ONEINSTANCEVNCHOSTUPDATING
	case VNCHOSTUPDATED:
		return ONEINSTANCEVNCHOSTUPDATED
	case AUTHKEYSUPDATING:
		return AUTHKEYSUPDATINGTYPE
	case AUTHKEYSUPDATED:
		return AUTHKEYSUPDATEDTYPE
	case AUTHKEYSFAILURE:
		return AUTHKEYSFAILURETYPE
	case CHEFCONFIGSETUPSTARTING:
		return ONEINSTANCECHEFCONFIGSETUPSTARTING
	case CHEFCONFIGSETUPSTARTED:
		return ONEINSTANCECHEFCONFIGSETUPSTARTED
	case INSTANCEIPSUPDATING:
		return INSTANCEIPSUPDATINGTYPE
	case INSTANCEIPSUPDATED:
		return INSTANCEIPSUPDATEDTYPE
	case INSTANCEIPSFAILURE:
		return INSTANCEIPSFAILURETYPE
	case CONTAINERNETWORKSUCCESS:
		return CONTAINERNETWORKSUCCESSTYPE
	case CONTAINERNETWORKFAILURE:
		return CONTAINERNETWORKFAILURETYPE
	case DNSNETWORKCREATING:
		return ONEINSTANCEDNSCNAMING
	case DNSNETWORKCREATED:
		return ONEINSTANCEDNSCNAMED
	case DNSNETWORKSKIPPED:
		return ONEINSTANCEDNSNETWORKSKIPPED
	case CLONING:
		return ONEINSTANCEGITCLONING
	case CLONED:
		return ONEINSTANCEGITCLONED
	case CONTAINERLAUNCHING:
		return CONTAINERINSTANCELAUNCHINGTYPE
	case CONTAINERBOOTSTRAPPING:
		return CONTAINERINSTANCEBOOTSTRAPPING
	case CONTAINERBOOTSTRAPPED:
		return CONTAINERINSTANCEBOOTSTRAPPED
	case CONTAINERLAUNCHED:
			return CONTAINERINSTANCELAUNCHEDTYPE
	case CONTAINEREXISTS:
		return CONTAINERINSTANCEEXISTS
	case CONTAINERDELETE:
		return CONTAINERINSTANCEDELETE
	case CONTAINERSTARTING:
		return CONTAINERINSTANCESTARTING
	case CONTAINERSTARTED:
		return CONTAINERINSTANCESTARTED
	case CONTAINERSTOPPING:
		return CONTAINERINSTANCESTOPPING
	case CONTAINERSTOPPED:
		return CONTAINERINSTANCESTOPPED
	case CONTAINERRESTARTING:
		return CONTAINERINSTANCERESTARTING
	case CONTAINERRESTARTED:
		return CONTAINERINSTANCERESTARTED
	case CONTAINERUPGRADED:
		return CONTAINERINSTANCEUPGRADED
  case CONTAINERRUNNING:
		return CONTAINERINSTANCERUNNING
	case CONTAINERERROR:
		return CONTAINERINSTANCEERROR

	case WAITUNTILL:
		return ONEINSTANCEWAITING
	case LCMSTATECHECK:
		return ONEINSTANCELCMSTATECHECKING
	case VMSTATECHECK:
		return ONEINSTANCEVMSTATECHECKING
	case PENDING + ".":
		return ONEINSTANCEVMSTATEPENDING
	case HOLD + ".":
		return ONEINSTANCEVMSTATEHOLD
	case ACTIVE + ".lcm_init":
		return ONEINSTANCELCMSTATEINIT
	case ACTIVE + ".boot":
		return ONEINSTANCELCMSTATEBOOT
	case ACTIVE + ".prolog":
		return ONEINSTANCELCMMSTATEPROLOG
	case PREERROR:
		return ONEINSTANCEPREERRORTYPE
	case ERROR:
		return ONEINSTANCEERRORTYPE
	default:
		return "arrgh"
	}
}

func (s Status) Description(name string) string {
	error_common := "Oops something went wrong on"
	switch s.String() {
	case LAUNCHING:
		return "Your " + name + " machine is initializing.."
	case BALANCECHECK:
		return "Verifying credit balance.."
	case INSUFFICIENT_FUND:
		return "Insuffient funds on you wallet to launch " + name +" machine .."
	case VMBOOTING:
		return "Virtual machine " + name + " is booting ..."
	case LAUNCHED:
		return "Machine " + name + " was initialized on cloud.."
	case QUOTAUPDATING:
		return "Machine " + name + " updating into quota.."
	case QUOTAUPDATED:
		return "Machine " + name + " updated into quota.."
	case VNCHOSTUPDATING:
		return name + " vnc_host is updating.."
	case VNCHOSTUPDATED:
		return name + " vnc_host is updated.."
	case BOOTSTRAPPING:
		return name + " was bootstrapping.."
	case BOOTSTRAPPED:
		return name + " was bootstrapped.."
	case STATEUPPING:
		return name + " is stateupping.."
	case STATEUPPED:
		return name + " is stateupped.."
	case RUNNING:
		return name + " is running.."
	case APPDEPLOYING:
		return "Your App " + name + " is deploying.."
	case APPDEPLOYED:
		return "Your App " + name + " is deployed.."
	case STARTING:
		return "Starting process initializing on " + name + ".."
	case STARTED:
		return name + " was started.."
	case STOPPING:
		return "Stopping process initializing on " + name + ".."
	case STOPPED:
		return name + " was stopped.."
	case UPGRADED:
		return name + " was upgraded.."
	case DESTROYING:
		return "Destroying process initializing on " + name + ".."
	case NUKED:
		return name + " was removed.."
	case SNAPSHOTTING:
		return "Snapshotting process initializing on " + name + ".."
	case SNAPSHOTTED:
		return name + " was snapcreated.."
	case COOKBOOKDOWNLOADING:
		return "Chef cookbooks are downloading.."
	case COOKBOOKDOWNLOADED:
		return "Chef cookbooks are successfully downloaded.."
	case COOKBOOKFAILURE:
		return error_common + "Downloading Cookbooks on " + name + ".."
	case CHEFCONFIGSETUPSTARTING:
		return "Chef config setup_starting .."
	case CHEFCONFIGSETUPSTARTED:
		return "Chef config setup_started .."
	case CLONING:
		return "Cloning your git repository .."
	case CLONED:
		return "Cloned your git repository .."
	case DNSNETWORKCREATING:
		return "DNS CNAME creating " + name + ".."
	case DNSNETWORKCREATED:
		return "DNS CNAME created successfully " + name + ".."
	case DNSNETWORKSKIPPED:
		return "DNS CNAME skipped " + name + ".."
	case AUTHKEYSUPDATING:
		return "SSH keys are updating on your " + name
	case AUTHKEYSUPDATED:
		return "SSH keys are updated on your " + name
	case AUTHKEYSFAILURE:
		return error_common + "Updating Ssh keys on " + name + ".."
	case INSTANCEIPSUPDATING:
		return "Private and public ips are updating on your " + name
	case INSTANCEIPSUPDATED:
		return "Private and public ips are updated on your " + name
	case INSTANCEIPSFAILURE:
		return error_common + "Updating private and public ips on " + name + ".."
	case CONTAINERNETWORKSUCCESS:
		return "Private and public ips are updated on your " + name
	case CONTAINERNETWORKFAILURE:
		return error_common + "Updating private and public ips on " + name + ".."
	case CONTAINERLAUNCHING:
		return "Your " + name + " container is initializing.."
	case CONTAINERBOOTSTRAPPING:
		return name + " was bootstrapping.."
	case CONTAINERBOOTSTRAPPED:
		return name + " was bootstrapped.."
	case CONTAINERLAUNCHED:
		return "Container " + name + " was initialized on cloud.."
	case CONTAINEREXISTS:
		return name + "was exists.."
	case CONTAINERDELETE:
		return name + "was deleted.."
	case CONTAINERSTARTING:
		return "Starting process initializing on " + name + ".."
	case CONTAINERSTARTED:
		return name + " was started.."
	case CONTAINERSTOPPING:
		return "Stopping process initializing on " + name + ".."
	case CONTAINERSTOPPED:
		return name + " was stopped.."
	case CONTAINERRESTARTING:
		return "Restarting process initializing on " + name + ".."
	case CONTAINERRESTARTED:
		return name + " was restarted.."
	case CONTAINERUPGRADED:
		return name + " was upgraded.."
	case CONTAINERRUNNING:
		return name + "was running.."
  case VMSTATECHECK:
    return "Machine "+ name +" state checking"
	case WAITUNTILL:
	  return "waiting 20 seconds to machine " + name +"to deploy.."
	case PENDING + ".":
		return "Machine " + name + "is in state of pending"
	case HOLD + ".":
		return "Machine " + name + "is in state of hold"
	case ACTIVE + ".lcm_init":
		return "Machine " + name + " activation initiated.."
	case ACTIVE + ".boot":
		return "Machine " + name + " activation booting.."
	case ACTIVE + ".prolog":
		return "Machine " + name + "is in state of activation prologing"
	case CONTAINERERROR:
		return error_common + name + ".."
	case ERROR:
		return error_common + name + ".."
	case PREERROR:
		return error_common + name + ".."
	default:
		return "arrgh"
	}
}
