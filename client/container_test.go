package client

import (
	"context"
	"testing"

	"github.com/epicchainlabs/epicchain-sdk-go/container"
	cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"
	"github.com/epicchainlabs/epicchain-sdk-go/eacl"
	"github.com/stretchr/testify/require"
)

func TestClient_Container(t *testing.T) {
	c := newClient(t, nil)
	ctx := context.Background()

	t.Run("missing signer", func(t *testing.T) {
		tt := []struct {
			name       string
			methodCall func() error
		}{
			{
				"put",
				func() error {
					_, err := c.ContainerPut(ctx, container.Container{}, nil, PrmContainerPut{})
					return err
				},
			},
			{
				"delete",
				func() error {
					return c.ContainerDelete(ctx, cid.ID{}, nil, PrmContainerDelete{})
				},
			},
			{
				"set_eacl",
				func() error {
					return c.ContainerSetEACL(ctx, eacl.Table{}, nil, PrmContainerSetEACL{})
				},
			},
		}

		for _, test := range tt {
			t.Run(test.name, func(t *testing.T) {
				require.ErrorIs(t, test.methodCall(), ErrMissingSigner)
			})
		}
	})
}
