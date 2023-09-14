#ifndef __sifi_http_hh
#define __sifi_http_hh

#include <curl/curl.h>
#include <string>

using namespace std;

namespace http
{
	const string Delete = "DELETE";
	const string Get    = "GET";
	const string Patch  = "PATCH";
	const string Post   = "POST";
	const string Put    = "PUT";

	const int StatusOK                  = 200;
	const int StatusBadRequest          = 400;
	const int StatusUnauthorized        = 401;
	const int StatusForbidden           = 403;
	const int StatusNotFound            = 404;
	const int StatusLocked              = 423;
	const int StatusInternalServerError = 500;
	const int StatusNotImplemented      = 501;
	const int StatusServiceUnavailable  = 503;

	class Code
	{
	public:
		Code();
		bool OK();

#ifdef SWIGPYTHON
		string __repr__();
#endif

		CURLcode curl;
		int      http;

	private:
		string String() const;

		friend ostream &operator<<(ostream &s, Code const &error);
	};

} // namespace http

#endif // __sifi_http_hh
