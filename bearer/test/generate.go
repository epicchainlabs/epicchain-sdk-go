package bearertest

import (
	"testing"

	"github.com/epicchainlabs/epicchain-sdk-go/bearer"
	eacltest "github.com/epicchainlabs/epicchain-sdk-go/eacl/test"
	usertest "github.com/epicchainlabs/epicchain-sdk-go/user/test"
)

// Token returns random bearer.Token.
//
// Resulting token is unsigned.
func Token(t testing.TB) (tok bearer.Token) {
	tok.SetExp(3)
	tok.SetNbf(2)
	tok.SetIat(1)
	tok.ForUser(usertest.ID(t))
	tok.SetEACLTable(eacltest.Table(t))

	return tok
}
