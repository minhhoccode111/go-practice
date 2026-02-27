package main

/*
Starting with Go 1.6, the http package has transparent support for the HTTP/2
protocol when using HTTPS. Programs that must disable HTTP/2 can do so by
setting [Transport.TLSNextProto] (for clients) or [Server.TLSNextProto] (for
servers) to a non-nil, empty map. Alternatively, the following GODEBUG settings
are currently supported:

GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
Please report any issues before disabling HTTP/2 support: https://golang.org/s/http2bug

The http package's Transport and Server both automatically enable HTTP/2 support
for simple configurations. To enable HTTP/2 for more complex configurations, to
use lower-level HTTP/2 features, or to use a newer version of Go's http2 package,
import "golang.org/x/net/http2" directly and use its ConfigureTransport and/or
ConfigureServer functions. Manually configuring HTTP/2 via the
golang.org/x/net/http2 package takes precedence over the net/http package's
built-in HTTP/2 support.
*/

func main() {

}
