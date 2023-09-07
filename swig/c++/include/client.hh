#ifndef __sifi_client_hh
#define __sifi_client_hh

#include <string>

using namespace std;

namespace sifi
{

	class Client
	{
	public:
		Client(string host, int port);
		~Client();

		Client *SetDebug(bool debug);
		string  Get(string path);

		void Login(string username, string password);

	protected:
		char *Post(string path);

	private:
		string URL(string path);

		string     baseURL;
		string     password;
		string     username;
		string     token;
		bool       debug;
		static int instances;
	};

} // namespace sifi

#endif // __sifi_client_hh
