#ifndef __sifi_error_hh
#define __sifi_error_hh

#include <sstream>
#include <string>

using namespace std;

namespace sifi
{

	//!
	// Error represents a SIFI API error. It holds an error code (http)
	// and an error message.
	class Error
	{
	public:
		Error(int code, string message);

#ifdef SWIGPYTHON
		string __repr__();
#endif

	private:
		int    code;
		string error;

		friend ostream &operator<<(ostream &s, Error const &error);
	};

} // namespace sifi

#endif // __sifi_error_hh
