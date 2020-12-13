// +build openbsd

package fcgipledge

import (
	"golang.org/x/sys/unix"
)

func pledge(promises string) error {
	return unix.Pledge(promises, "")
}
