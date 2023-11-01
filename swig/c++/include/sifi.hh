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

	const int    PortNumber = 7434;
	const string Version    = "SIFI_VERSION";

	// from  cloud/cloud.go

	const string CloudPath                  = "v2/cloud";
	const string CloudAccessControlListPath = "acl";
	const string CloudClusterPath           = "cluster";
	const string CloudDatastorePath         = "datastore";
	const string CloudDocumentPath          = "document";
	const string CloudGroupPath             = "group";
	const string CloudHookPath              = "hook";
	const string CloudHostPath              = "host";
	const string CloudImagePath             = "image";
	const string CloudInstanceGroupPath     = "instance-group";
	const string CloudInstancePath          = "instance";
	const string CloudMarketPath            = "market";
	const string CloudSecurityGroupPath     = "security-group";
	const string CloudSystemPath            = "system";
	const string CloudTemplatePath          = "template";
	const string CloudUserPath              = "user";
	const string CloudVirtualDataCenterPath = "vdc";
	const string CloudVirtualNetworkPath    = "vnet";
	const string CloudVirtualRouterPath     = "vrouter";
	const string CloudZonePath              = "zone";

	// from metal/metal.go

	const string MetalPath     = "v2/metal";
	const string MetalHostPath = "host";

	// from snapshot/snapshot.go

	const string SnapshotPath        = "v2/snapper";
	const string SnapshotArchivePath = "archive";
	const string SnapshotManualPath  = "manual";
	const string SnapshotRemovePath  = "remote";
	const string SnapshotStackPath   = "stack";
	const string SnapshotStatusPath  = "status";
} // namespace sifi

#endif // __sifi_hh
