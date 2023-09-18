#ifndef __sifi_hh
#define __sifi_hh

#include <string>

#ifndef SIFI_VERSION
#	define SIFI_VERSION "undefined"
#endif

using namespace std;

namespace sifi
{
	// from sifi/sifi.go

	const string PathPrefix = "v2";
	const int    PortNumber = 7434;
	const string Version    = "SIFI_VERSION";

	// from  cloud/cloud.go

	const string AccessControlListPath = "acl";
	const string ClusterPath           = "cluster";
	const string ComputePath           = "compute/host";
	const string DatastorePath         = "datastore";
	const string DocumentPath          = "document";
	const string GroupPath             = "group";
	const string HookPath              = "hook";
	const string HostPath              = "host";
	const string ImagePath             = "image";
	const string InstanceGroupPath     = "instance-group";
	const string InstancePath          = "instance";
	const string MarketPath            = "market";
	const string PingPath              = "ping";
	const string SecurityGroupPath     = "security-group";
	const string SystemPath            = "system";
	const string TemplatePath          = "template";
	const string UserPath              = "user";
	const string VirtualDataCenterPath = "vdc";
	const string VirtualNetworkPath    = "vnet";
	const string VirtualRouterPath     = "vrouter";
	const string ZonePath              = "zone";
} // namespace sifi

#endif // __sifi_hh
