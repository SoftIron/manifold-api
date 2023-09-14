#include "./include/client.hh"
#include "./include/instance.hh"
#include "./include/sifi.hh"
#include "include/http.hh"
#include <curl/curl.h>
#include <getopt.h>
#include <iostream>

using namespace std;

int
main(int argc, char *argv[])
{
	string address = "localhost:" + to_string(sifi::PortNumber);
	string username, password;

	while (true) {
		static int                 c;
		static const struct option options[] = {
			{"address",  required_argument, 0, 'a'},
			{"password", required_argument, 0, 'p'},
			{"username", required_argument, 0, 'u'},
			{0,          0,                 0, 0  },
		};

		int option_index = 0;

		c = getopt_long(argc, argv, "a:p:u:", options, &option_index);

		if (c == -1) {
			break; // done, end of options
		}

		switch (c) {
		case 'a':
			address = optarg;
			break;
		case 'p':
			password = optarg;
			break;
		case 'u':
			username = optarg;
			break;
		default:
			exit(-1);
		}
	}

	if (username.empty() || password.empty()) {
		cerr << "username and password required" << endl;
		exit(-1);
	}

	sifi::Client *client = new sifi::Client(address);

	client->SetDebug(true)->SetLogin(username, password);

	http::Code code = client->Login();
	if (!code.OK()) {
		cerr << "Error: " << code << endl;
		exit(-1);
	}

	string resp;

	code = sifi::get_instance(client, 0, resp);
	if (!code.OK()) {
		cerr << "Error: " << code << endl;
		exit(-1);
	}

	cout << resp << endl;

	return 0;
}
