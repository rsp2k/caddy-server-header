package serverheader

import (
    "github.com/caddyserver/caddy/v2"
    "github.com/caddyserver/caddy/v2/modules/caddyhttp"
    "net/http"
    "math/rand"
)

// ModuleName is the name of this module.
const ModuleName = "http.handlers.serverheader"

func getRandomServerHeader() string {
	serverHeaders := []string{
		"Apache/1.3.26 (Unix)",
		"Apache/2.0.40 (Red Hat)",
		"Apache/2.2.8 (Win32)",
		"Apache/2.4.10 (Debian)",
		"Microsoft-IIS/4.0",
		"Microsoft-IIS/5.0",
		"Microsoft-IIS/6.0",
		"Microsoft-IIS/7.5",
		"nginx/0.5.35",
		"nginx/0.6.39",
		"nginx/1.0.6",
		"nginx/1.2.9",
		"Sun-ONE-Web-Server/6.1",
		"Zeus/4.3",
		"Zeus/3.4",
		"Jetty/4.2.9",
		"Jetty/5.1.10",
		"Lotus-Domino/5.0.12",
		"Lotus-Domino/6.5.4",
		"Tomcat/4.1.24",
		"Tomcat/5.5.23",
		"Tomcat/6.0.18",
		"OpenSSL/0.9.6c Apache/1.3.23 (Unix)",
		"OpenSSL/0.9.8k",
		"Novell-NetWare-Enterprise-Web-Server/5.1",
		"Oracle-HTTP-Server/9.0.4",
		"Oracle-HTTP-Server/10.1.2.0.2",
		"IBM_HTTP_Server/6.1.0.17 Apache/2.0.47",
		"IBM_HTTP_Server/7.0.0.0",
		"CAIDevServer/6.5",
		"Roxen/2.2",
		"Roxen/4.0.241",
		"Jigsaw/2.2.4",
		"Boa/0.94.13",
		"Boa/0.93.15",
		"AllegroServe/1.2.40",
		"CERN/3.0",
		"gws/1.12.3",
		"Cobalt/RAQ3 Apache/1.3.6",
		"WebLogic/8.1",
		"WebLogic/9.2",
		"WebLogic/10.3.6",
		"Squid/2.5.STABLE4",
		"Squid/2.7.STABLE9",
		"LiteSpeed/3.3.4",
		"LiteSpeed/4.0.10",
		"Lighttpd/1.4.15",
		"Lighttpd/1.3.16",
		"Caudium/1.4.12",
		"WN/2.4.7",
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Return a random value from the list
	return serverHeaders[rand.Intn(len(serverHeaders))]
}

func init() {
    caddy.RegisterModule(ServerHeader{})
}

// ServerHeader is a Caddy module that sets the "Server" response header.
type ServerHeader struct {
    HeaderValue string `json:"header_value,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (ServerHeader) CaddyModule() caddy.ModuleInfo {
    return caddy.ModuleInfo{
        ID:  ModuleName,
        New: func() caddy.Module { return new(ServerHeader) },
    }
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (sh ServerHeader) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
    // Call the next handler in the chain.
    if err := next.ServeHTTP(w, r); err != nil {
        return err
    }

    // Set the "Server" response heade based on value.
    if sh.HeaderValue == "random" {
        w.Header().Set("Server", getRAndomServerHeader())
    } else if sh.HeaderValue != "" {
        w.Header().Set("Server", sh.HeaderValue)
    }
    return nil
}

// Interface guard
var _ caddyhttp.MiddlewareHandler = (*ServerHeader)(nil)
