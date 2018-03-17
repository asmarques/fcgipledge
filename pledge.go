// +build !openbsd

package fcgipledge

func pledge(promises string) error {
	// no-op
	return nil
}
