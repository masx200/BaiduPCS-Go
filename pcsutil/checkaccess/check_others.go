//go:build aix || plan9
// +build aix plan9

package checkaccess

func AccessRDWR(path string) bool {
	return true
}
