#include "./include/error.hh"
#include <iostream>
#include <sstream>

using namespace sifi;

/// @brief Constructor.
/// @param code HTTP error code.
/// @param message API error message.
Error::Error(int code, string message) : code(code), error(message) {}

string
Error::String() const
{
	ostringstream os;

	os << "sifi (" << code << ") - " << error;
	return os.str();
}

/// @brief Streams the string representation of Error.
ostream &
sifi::operator<<(ostream &s, Error const &error)
{
	return s << error.String();
}

#ifdef SWIGPYTHON
string
Error::__repr__()
{
	return String();
}
#endif
