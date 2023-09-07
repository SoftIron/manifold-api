#include "./include/error.hh"
#include <iostream>

using namespace sifi;

/// @brief Constructor.
/// @param code HTTP error code.
/// @param message API error message.
Error::Error(int code, string message) : code(code), error(message) {}

/// @brief Streams the string representation of Error.
ostream &
sifi::operator<<(ostream &s, Error const &error)
{
	return s << "sifi (" << error.code << ") - " << error.error;
}

#ifdef SWIGPYTHON
string
Error::__repr__()
{
	ostringstream os;

	os << "sifi (" << code << ") - " << error;
	return os.str();
}
#endif
