[![GitHub Actions](https://img.shields.io/github/actions/workflow/status/softiron/manifold-api/build.yaml?branch=main)](https://github.com/softiron/manifold-api/actions?query=workflow%3Abuild)
[![Go Version](https://img.shields.io/github/go-mod/go-version/softiron/manifold-api)](https://img.shields.io/github/go-mod/go-version/softiron/manifold-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/softiron/manifold-api)](https://goreportcard.com/report/github.com/softiron/manifold-api)
[![Go Reference](https://pkg.go.dev/badge/github.com/softiron/manifold-api.svg)](https://pkg.go.dev/github.com/softiron/manifold-api)


# Manifold API Client Bindings

 View the API [swagger](https://softiron.github.io/manifold-doc/) doc.

## Go Bindings

See [example](example/main.go) for sample client.

# Manual Use

## Login

The API uses [Basic
Authentication](https://datatracker.ietf.org/doc/html/rfc7617) where the client
does a `GET`
[login](https://softiron.github.io/manifold-doc/#/default/get_login) and
includes a username and password in the `Authorization` header.

Included in the response is an access token in the form of a JSON Web Token with
a 10 minute TTL.

## Access Token

All subsequest requests must used the access token in the `Authorization`
header, for example:

```
Authorization: Bearer dht467bv4570flw2r
```

If a request uses an expired token the request will return HTTP code 401
(unauthorized). This is a hint to the client to re-login and obtain a new access
token.

### Refresh Token

The following is true for development releases of the API only, and is not a
released feature.

In addition to returning an access token, a login will also return a refresh
token in the form of a JSON Web Token with a one year TTL. This token can be
used in place of Basic Authentication to re-login and obtain a new access and
refresh token.

Refresh tokens are single use, so once used the server will reject it (401) even
if the TTL has not expired. If the refresh token has expired or was lost a
re-login must use Basic Authentication with a username and password.

Usage of refresh tokens is optional, and the client can instead re-login using
Basic Authentication.



