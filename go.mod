package serverheader

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"net/http"
)

// ModuleName is the name of this module.
const ModuleName = "http.handlers.serverheader"

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

	// Set the "Server" response header.
	if sh.HeaderValue != "" {
		w.Header().Set("Server", sh.HeaderValue)
	}

	return nil
}

// Interface guard
var _ caddyhttp.MiddlewareHandler = (*ServerHeader)(nil)
