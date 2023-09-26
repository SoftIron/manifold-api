#include "./include/http.hh"
#include <curl/curl.h>
#include <iostream>
#include <sstream>

using namespace http;

Code::Code() : curl(CURLE_OK), http(0) {}

string
Code::String() const
{
	if (curl != CURLE_OK) {
		ostringstream os;

		os << curl_easy_strerror(curl);
		if (http != 0) {
			os << ":" << http;
		}

		return os.str();
	}

	return to_string(http);
}

bool
Code::OK()
{
	return curl == CURLE_OK && http == http::StatusOK;
}

/// @brief Streams the string representation of Code.
ostream &
http::operator<<(ostream &s, Code const &code)
{
	return s << code.String();
}

#ifdef SWIGPYTHON
string
Code::__repr__()
{
	return String();
}
#endif
