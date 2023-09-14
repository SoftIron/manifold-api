#ifndef __sifi_instance_hh
#define __sifi_instance_hh

#include "client.hh"

using namespace std;

namespace sifi
{
	http::Code get_instance(Client *c, int id, string &resp);
}

#endif // __sifi_instance_hh
