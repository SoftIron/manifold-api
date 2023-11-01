#include "./include/instance.hh"
#include "./include/sifi.hh"
#include "include/http.hh"
#include <sstream>

using namespace std;
using namespace sifi;

http::Code
sifi::get_instance(sifi::Client *c, int id, string &resp)
{
	stringstream path;

	path << CloudPath << "/" << CloudInstancePath << "/" << id;

	return c->Get(path.str(), resp);
}
