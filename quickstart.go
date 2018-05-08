package quickstart

import (
	"fmt"
)

// VersionMajor defines the major version number
const VersionMajor int = 0

// VersionMinor defines the minor version number
const VersionMinor int = 9

// VersionString returns the full version string
func VersionString() string {
	return fmt.Sprintf("%d.%d", VersionMajor, VersionMinor)
}
