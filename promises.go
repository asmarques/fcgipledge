package fcgipledge

import "strings"

// Common promises for pledge()
const (
	Stdio   = "stdio"
	Inet    = "inet"
	Unix    = "unix"
	DNS     = "dns"
	Rpath   = "rpath"
	Wpath   = "wpath"
	Cpath   = "cpath"
	Tmppath = "tmppath"
)

func createPromiseString(promises []string) string {
	newPromises := []string{Stdio, Unix}
	for _, promise := range promises {
		if promise != Stdio && promise != Unix {
			newPromises = append(newPromises, promise)
		}
	}
	return strings.Join(newPromises, " ")
}
