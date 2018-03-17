package fcgipledge

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

// Config represents the server configuration
type Config struct {
	Promises []string
}

// ListenAndServe starts a FastCGI server listening on a Unix domain socket
// and restricts further operations using OpenBSD's pledge() mechanism. The
// supplied handler will be used to handle incoming requests. If handler is nil
// the http.DefaultServeMux will be used. By default the process will be unable
// to perform any syscalls except for those required to communicate via the
// created socket (equivalent to supplying "stdio" and "unix" promises).
func ListenAndServe(name string, handler http.Handler, config *Config) error {
	err := os.Remove(name)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error removing socket: %s", err)
	}

	// create socket and trigger go's network stack probe before restricting process
	l, err := net.Listen("unix", name)
	if err != nil {
		return fmt.Errorf("error creating socket: %s", err)
	}

	var promises []string
	if config != nil {
		promises = config.Promises
	}

	p := createPromiseString(promises)
	if err = pledge(p); err != nil {
		return fmt.Errorf("pledge error: %s", err)
	}

	err = fcgi.Serve(l, handler)
	if err != nil {
		return fmt.Errorf("error starting server: %s", err)
	}

	return nil
}
