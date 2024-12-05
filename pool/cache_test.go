package pool

import (
	"testing"

	"github.com/epicchainlabs/epicchain-sdk-go/crypto/test"
	"github.com/epicchainlabs/epicchain-sdk-go/session"
	sessiontest "github.com/epicchainlabs/epicchain-sdk-go/session/test"
	"github.com/stretchr/testify/require"
)

func TestSessionCache_GetUnmodifiedToken(t *testing.T) {
	const key = "Foo"
	target := sessiontest.Object()

	check := func(t *testing.T, tok session.Object, extra string) {
		require.False(t, tok.VerifySignature(), extra)
	}

	cache, err := newCache(defaultSessionCacheSize)
	require.NoError(t, err)

	cache.Put(key, target)
	value, ok := cache.Get(key)
	require.True(t, ok)
	check(t, value, "before sign")

	err = value.Sign(test.RandomSignerRFC6979(t))
	require.NoError(t, err)

	value, ok = cache.Get(key)
	require.True(t, ok)
	check(t, value, "after sign")
}
