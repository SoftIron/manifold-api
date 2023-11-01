#ifndef __sifi_instance_hh
#define __sifi_instance_hh

#include "client.hh"
#include <time.h>

using namespace std;

namespace sifi
{
	http::Code del_instance_disk(Client *c, int id, int disk);
	http::Code del_instance_disk_snapshot(Client *c, int id, int disk, int snapshot);
	http::Code del_instance_nic(Client *c, int id, int nic);
	http::Code del_instance_schedule(Client *c, int id, int schedule);
	http::Code del_instance_security_group(Client *c, int id, int nic);
	http::Code del_instance_snapshot(Client *c, int id, int snapshot);
	http::Code get_instance(Client *c, int id, string &resp);
	http::Code get_instance_set(Client *c, int ids[], string &resp);
	http::Code get_instances(Client *c, string state, bool extended, string &resp);
	http::Code get_instance_monitoring_data(Client *c, int id, string &resp);
	http::Code get_instances_monitoring_data(Client *c, string &resp);
	http::Code get_instances_accounting_data(Client *c, string filter, time_t start, time_t end);
} // namespace sifi

#endif // __sifi_instance_hh
