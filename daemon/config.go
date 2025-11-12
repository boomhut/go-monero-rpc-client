package daemon

import "net/http"

// Config holds daemon client configuration
type Config struct {
	// Address of the monero daemon RPC
	// example: http://127.0.0.1:18081/json_rpc
	Address string
	// Custom headers to send with requests
	CustomHeaders map[string]string
	// Custom HTTP transport
	Transport http.RoundTripper
}
