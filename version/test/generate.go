package versiontest

import (
	"math/rand"

	"github.com/epicchainlabs/epicchain-sdk-go/version"
)

// Version returns random version.Version.
func Version() (v version.Version) {
	v.SetMajor(rand.Uint32())
	v.SetMinor(rand.Uint32())
	return v
}
