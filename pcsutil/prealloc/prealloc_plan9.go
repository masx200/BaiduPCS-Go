//go:build plan9
// +build plan9

package prealloc

func PreAlloc(fd uintptr, length int64) error {
	return nil
}
