#ifndef __sifi_client_hh
#define __sifi_client_hh

#include "http.hh"
#include <string>

using namespace std;

namespace sifi
{

	class Client
	{
	public:
		Client(string address);
		~Client();

		Client *SetDebug(bool debug);
		Client *SetLogin(string username, string password);

		http::Code Login();

		http::Code Delete(string path, string &response);
		http::Code Get(string path, string &response);
		http::Code Patch(string path, const string &body, string &response);
		http::Code Post(string path, const string &body, string &response);
		http::Code Put(string path, const string &body, string &response);

	private:
		string URL(string path);

		http::Code RequestWrapper(string method, string path, const string &body, string &response);
		http::Code Request(string method, string path, const string &body, string &response);

		string baseURL;
		string password;
		string username;
		string token;
		bool   debug;

		static int instances;
	};

	inline http::Code
	Client::Delete(string path, string &response)
	{
		return RequestWrapper(http::Delete, path, "", response);
	}

	inline http::Code
	Client::Get(string path, string &response)
	{
		return RequestWrapper(http::Get, path, "", response);
	}

	inline http::Code
	Client::Patch(string path, const string &body, string &response)
	{
		return RequestWrapper(http::Patch, path, body, response);
	}

	inline http::Code
	Client::Post(string path, const string &body, string &response)
	{
		return RequestWrapper(http::Post, path, body, response);
	}

	inline http::Code
	Client::Put(string path, const string &body, string &response)
	{
		return RequestWrapper(http::Put, path, body, response);
	}

} // namespace sifi

#endif // __sifi_client_hh
