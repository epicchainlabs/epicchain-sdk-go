package usertest

import (
	"testing"

	"github.com/epicchainlabs/epicchain-sdk-go/crypto/test"
	"github.com/epicchainlabs/epicchain-sdk-go/user"
)

// ID returns random user.ID.
func ID(tb testing.TB) user.ID {
	var x user.ID
	s := test.RandomSignerRFC6979(tb)
	x = s.UserID()

	return x
}
