#include "./include/client.hh"
#include "./include/http.hh"
#include <curl/curl.h>
#include <iostream>
#include <sstream>
#include <string>

using namespace std;
using namespace sifi;

static size_t
writeCallback(char *ptr, size_t size, size_t nmemb, void *userdata)
{
	string *dest = (string *)userdata;

	dest->append(ptr, size * nmemb);
	return size * nmemb;
}

static string
get_token(const string &body)
{
	// {"token":"eyJhbGciOiJIUzI1NiIsInRIzVuf_b748XVFwmpuhVZ0D_eNToEl8jbI8lk"}
	//
	// This is the only JSON payload that we need to decode, and we know the
	// format is as above. Simple state machine to parse this and allow for
	// whitespace.

	int state = 0, start, end;

	// search to the first non-whitespace character after the opening brace
	for (start = 0; start < body.size(); start++) {
		//		cout << "state:" << state << " start:" << start << " c:" << body[start] << endl;
		if (isspace(body[start])) {
			continue;
		}
		if (state == 0 && body[start] == '{') {
			state++;
			continue;
		}
		if (state == 1 && body[start] == '"') {
			state++;
			continue;
		}
		if (state == 2 && body[start] == 't') {
			start = body.find(':');
			state++;
			continue;
		}
		if (state == 3 && body[start] == '"') {
			start++; // consume starting quote
			state++;
			break;
		}
		break;
	}

	if (state != 4) {
		return "invalid token";
	}

	// search from the end to the first non-whitespace character after the
	// opening brace
	for (state = 0, end = body.rfind('}'); end > 0; end--) {
		//		cout << "state:" << state << " end:" << end << " c:" << body[end] << endl;
		if (isspace(body[start])) {
			continue;
		}
		if (state == 0 && body[end] == '}') {
			state++;
			continue;
		}
		if (state == 1 && body[end] == '"') {
			state++;
		}
		break;
	}

	if (state != 2) {
		return "invalid token";
	}

	return body.substr(start, end - start);
}

int Client::instances = 0;

Client::Client(string address) : debug(false)
{
	ostringstream os;

	os << "https://" << address;
	this->baseURL = os.str();

	if (instances++ == 0) {
		curl_global_init(CURL_GLOBAL_ALL);
	}
}

Client::~Client()
{
	if (--instances == 0) {
		curl_global_cleanup();
	}
}

/// @brief Turn debugging on or off.
Client *
Client::SetDebug(bool debug)
{
	this->debug = debug;
	return this;
}

/// @brief Set the SIFI API Login information.
/// @param username Manifold username.
/// @param password Manifold password.
Client *
Client::SetLogin(string username, string password)
{
	this->username = username;
	this->password = password;
	return this;
}

// URL returns the URL for the given path.
string
Client::URL(string path)
{
	return this->baseURL + "/" + path;
}

/// Logs into the service. It is optional to call this method as any
/// Request will automatically call Login if the Client is missing its Token.
/// Call this when you want feedback right away on the acceptance of the
/// Username/Password credentials.
http::Code
Client::Login()
{
	if (debug) {
		cout << "request method=" << http::Get << " url=" << URL("login") << endl;
	}

	CURL      *curl = curl_easy_init();
	http::Code code;

	if (!curl) {
		code.curl = CURLE_FAILED_INIT;
		return code;
	}

	curl_easy_setopt(curl, CURLOPT_URL, URL("login").c_str());
	curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, 0);
	curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, 0);
	curl_easy_setopt(curl, CURLOPT_HTTPAUTH, CURLAUTH_BASIC); // basic auth to get token
	curl_easy_setopt(curl, CURLOPT_USERNAME, username.c_str());
	curl_easy_setopt(curl, CURLOPT_PASSWORD, password.c_str());
	// if (debug) {
	// 	curl_easy_setopt(curl, CURLOPT_VERBOSE, 1L);
	// }

	string response;

	curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writeCallback);
	curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);

	code.curl = curl_easy_perform(curl);
	if (code.curl == CURLE_OK) {
		long response_code;

		curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &response_code);
		code.http = 0xffffffff & response_code;
	}

	curl_easy_cleanup(curl);

	if (code.http == http::StatusOK) {
		token = get_token(response);
	}

	return code;
}

/// Sends a method request to the given URL. If the Client does not have a token
/// the Login method is called first. If a 401 response is received it assumes
/// the token has expired and will re-Login and retry the request.
http::Code
Client::RequestWrapper(string method, string path, const string &body, string &response)
{
	http::Code code;

	if (token == "") {
		code = Login();
		if (code.http != http::StatusOK) {
			return code;
		}
	}

	code = Request(method, URL(path), body, response);

	if (code.http == http::StatusUnauthorized) {
		code = Login();
		if (code.http != http::StatusOK) {
			return code;
		}

		return Request(method, URL(path), body, response);
	}

	return code;
}

http::Code
Client::Request(string method, string url, const string &body, string &response)
{
	if (debug) {
		cout << "request method=" << method << " url=" << url << endl;
	}

	http::Code code;
	CURL      *curl;

	curl = curl_easy_init();
	if (!curl) {
		code.curl = CURLE_FAILED_INIT;
		return code;
	}

	curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
	curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, method.c_str());
	curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, 0);
	curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, 0);
	// if (debug) {
	// 	curl_easy_setopt(curl, CURLOPT_VERBOSE, 1L);
	// }

	struct curl_slist *header = NULL;

	if (!body.empty()) {
		header = curl_slist_append(header, "Content-Type: application/json");

		curl_easy_setopt(curl, CURLOPT_POSTFIELDS, body.c_str());
		curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE, body.size() + 1);
	}

	stringstream auth;
	auth << "Authorization: Bearer " << token;
	header = curl_slist_append(header, auth.str().c_str());

	curl_easy_setopt(curl, CURLOPT_HTTPHEADER, header);
	curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writeCallback);
	curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);

	code.curl = curl_easy_perform(curl);
	if (code.curl == CURLE_OK) {
		long response_code;

		curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &response_code);
		code.http = 0xffffffff & response_code;
	}

	curl_easy_cleanup(curl);

	return code;
}
