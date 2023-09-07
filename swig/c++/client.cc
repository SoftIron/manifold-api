#include "./include/client.hh"
#include <curl/curl.h>
#include <iostream>
#include <sstream>
#include <string>

using namespace std;
using namespace sifi;

// // LoginResponse is the response payload for /login.
// typedef struct {
// 	string token;
// } login_response;

int Client::instances = 0;

Client::Client(string host, int port) : debug(false)
{
	ostringstream os;

	os << "https://" << host << "/" << port;
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
/// @param username HyperCloud username.
/// @param password HyperCloud password.
void
Client::Login(string username, string password)
{
	this->username = username;
	this->password = password;
}

// URL returns the URL for the given path.
string
Client::URL(string path)
{
	return this->baseURL + "/" + path;
}

string
Client::Get(string path)
{
	CURL *curl;

	curl = curl_easy_init();
	if (!curl) {
		return "";
	}

	curl_easy_setopt(curl, CURLOPT_URL, URL(path).c_str());

	curl_easy_cleanup(curl);

	return "";
}

// // Login logs into the service. It is optional to call this method as any
// // Request will automatically call Login if the Client is missing its Token.
// // Call this when you want feedback right away on the acceptance of the
// // Username/Password credentials.
// func (c *Client) Login(ctx context.Context) error {
// 	c.Logger.Debug("request", "method", http.MethodGet, "url", c.URL("login"))

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL("login"), http.NoBody)
// 	if err != nil {
// 		return err
// 	}

// 	req.SetBasicAuth(c.Username, c.Password)

// 	r, err := c.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	if r.StatusCode != http.StatusOK {
// 		return fmt.Errorf("login failed: %s (%s)", r.Status, r.Request.URL.Redacted())
// 	}

// 	defer r.Body.Close()

// 	var resp LoginResponse

// 	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
// 		return err
// 	}

// 	c.Token = resp.Token

// 	return nil
// }

// // Post sends a Post request to the given URL.
// func (c *Client) Post(ctx context.Context, url string, body, resp interface{}) error {
// 	return c.Request(ctx, http.MethodPost, url, body, resp)
// }

// // Put sends a Put request to the given URL.
// func (c *Client) Put(ctx context.Context, url string, body, resp interface{}) error {
// 	return c.Request(ctx, http.MethodPut, url, body, resp)
// }

// // Patch sends a Patch request to the given URL.
// func (c *Client) Patch(ctx context.Context, url string, body, resp interface{}) error {
// 	return c.Request(ctx, http.MethodPatch, url, body, resp)
// }

// // Get sends a Get request to the given URL.
// func (c *Client) Get(ctx context.Context, url string, resp interface{}) error {
// 	return c.Request(ctx, http.MethodGet, url, nil, resp)
// }

// // Delete sends a Delete request to the given URL.
// func (c *Client) Delete(ctx context.Context, url string, resp interface{}) error {
// 	return c.Request(ctx, http.MethodDelete, url, nil, resp)
// }

// // Request sends a method request to the given URL. If the Client does not
// // have a token the Login method is called first. If a 401 response is received
// // it assumes the token has expired and will re-Login and retry the request.
// func (c *Client) Request(ctx context.Context, method, path string, in, out interface{}) error {
// 	if c == nil {
// 		return errors.New("no service")
// 	}

// 	c.Logger.Debug("request", "method", method, "path", path)

// 	if c.Token == "" {
// 		if err := c.Login(ctx); err != nil {
// 			return err
// 		}
// 	}

// 	if err := c.request(ctx, method, path, in, out); err != nil {
// 		var e Error // 401 means the token expired so automatically re-login and try again.

// 		if errors.As(err, &e) && e.Code == http.StatusUnauthorized {
// 			c.Logger.Info("token expired")

// 			if err := c.Login(ctx); err != nil {
// 				return err
// 			}

// 			return c.request(ctx, method, path, in, out)
// 		}

// 		return err
// 	}

// 	return nil
// }

// func (c *Client) request(ctx context.Context, method, path string, in, out interface{}) error {
// 	var r io.Reader = http.NoBody

// 	if in != nil { // encode body
// 		b, err := json.Marshal(in)
// 		if err != nil {
// 			return err
// 		}

// 		r = bytes.NewReader(b)
// 	}

// 	req, err := http.NewRequestWithContext(ctx, method, c.URL(path), r)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+c.Token)

// 	resp, err := c.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	if c.Debug {
// 		b, err := httputil.DumpResponse(resp, true)
// 		if err != nil {
// 			return err
// 		}

// 		fmt.Printf("RESPONSE:\n%s", string(b))
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		var b bytes.Buffer

// 		_, err := b.ReadFrom(resp.Body)
// 		if err != nil {
// 			return Error{Code: resp.StatusCode, Text: err.Error()}
// 		}

// 		return Error{Code: resp.StatusCode, Text: b.String()}
// 	}

// 	if out != nil { // parse response
// 		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
