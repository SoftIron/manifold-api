import logging

import requests

from requests.packages.urllib3.exceptions import InsecureRequestWarning

log = logging.getLogger(__name__)


class ManifoldBaseError(requests.RequestException):
    """base exception with a data field"""

    def __init__(self, message, data):
        super().__init__(message)
        self.data = data


class ManifoldError(ManifoldBaseError):
    """an error returned by Manifold"""


class ManifoldSession:
    def __init__(
        self,
        user,
        passwd,
        api_host,
        api_port="443",
        api_path="manifold-api",
        api_ver="v2",
        verify=False,
    ):
        self._session = requests.Session()
        self._user = user
        self._passwd = passwd
        self._session.verify = verify
        self._api_host = api_host.strip("/")
        self._api_port = api_port
        self._api_path = api_path.strip("/")
        self._api_ver = api_ver.strip("/")
        self._connection_attempts = 0
        self._timeout = 5.1
        self.api_url = (
            f"{self._api_host}:{self._api_port}/{self._api_path}/{self._api_ver}/"
        )

        log.debug("ManifoldSession instantiating")
        log.debug("manifold server: %s", self.api_url)
        # we'll store the most recent request response here
        self.resp = None

        # finally, if we're ignoring bad tls, let's actually ignore it
        if not verify:
            log.debug("TLS verify set to False")
            requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

    def login(self):
        self._session.auth = (self._user, self._passwd)

        log.debug("attempting login...")
        log.debug(
            "login: verb=GET, endpoint=%s, user=%s",
            f"{self._api_host}:{self._api_port}/{self._api_path}/login",
            self._user,
        )
        try:
            self.resp = self._session.get(
                f"{self._api_host}:{self._api_port}/{self._api_path}/login",
                timeout=self._timeout,
            )
        except:  # requests.exceptions.ConnectionError:
            raise

        self._connection_attempts += 1

        if not self.resp.ok or self._connection_attempts > 3:
            raise ManifoldError(
                "Too many connection attempts, or possibly bad password",
                data=self.resp.text,
            )

        try:
            tok = self.resp.json()["token"]
        except (ValueError, requests.InvalidJSONError):
            raise ManifoldError(
                "Server returned OK but unable to decode JSON response",
                data=self.resp.text,
            )
        except KeyError:
            raise ManifoldError(
                "Server returned OK but response has no auth token", data=self.resp.text
            )

        log.debug("... success")
        self._session.headers.update({"Authorization": f"Bearer: {tok}"})

        # requests seems to want to reuse `auth` if it exists but then we get 401's,
        # so null it here and reset the attempt counter
        self._session.auth = None
        self._connection_attempts = 0

    def get(self, *args, **kwargs):
        return self.httpdo("get", *args, **kwargs)

    def post(self, *args, **kwargs):
        return self.httpdo("post", *args, **kwargs)

    def patch(self, *args, **kwargs):
        return self.httpdo("patch", *args, **kwargs)

    def httpdo(self, verb, endpoint, json=None, **kwargs):
        log.debug("httpdo: verb=%s, endpoint=%s, json=%s", verb, endpoint, json)
        url = self.api_url + endpoint.strip("/")
        http_call = getattr(self._session, verb)

        self.resp = http_call(url, json=json, timeout=self._timeout, **kwargs)

        log.debug("resp: [%s] %s", self.resp.status_code, self.resp.text)
        if self.resp.status_code == 401:
            self.login()
            self.httpdo(verb, endpoint, json)

        if self.resp.status_code == 404:
            return {}

        if not self.resp.ok:
            log.debug(
                "response not understood: [%s] %s",
                self.resp.status_code,
                self.resp.text,
            )

            # see if we can get json out of the response...
            try:
                results = self.resp.json()
            except (ValueError, requests.InvalidJSONError):
                results = self.resp.text

            raise ManifoldError(f"API error: {results}", data=results)

        try:
            results = self.resp.json()
        except (ValueError, requests.InvalidJSONError):
            raise ManifoldError(
                "Server returned OK but unable to decode JSON response",
                data=self.resp.text,
            )
            # all manifold endpoints return json, so this means manifold returned
            # "ok", but not valid json?

        return results
