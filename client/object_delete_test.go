package client

import (
	"context"
	"testing"

	cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"
	oid "github.com/epicchainlabs/epicchain-sdk-go/object/id"
	"github.com/stretchr/testify/require"
)

func TestClient_ObjectDelete(t *testing.T) {
	t.Run("missing signer", func(t *testing.T) {
		c := newClient(t, nil)

		_, err := c.ObjectDelete(context.Background(), cid.ID{}, oid.ID{}, nil, PrmObjectDelete{})
		require.ErrorIs(t, err, ErrMissingSigner)
	})
}
