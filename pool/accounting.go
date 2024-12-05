package pool

import (
	"context"

	"github.com/epicchainlabs/epicchain-sdk-go/accounting"
	"github.com/epicchainlabs/epicchain-sdk-go/client"
)

// BalanceGet requests current balance of the NeoFS account.
//
// See details in [client.Client.BalanceGet].
func (p *Pool) BalanceGet(ctx context.Context, prm client.PrmBalanceGet) (accounting.Decimal, error) {
	c, err := p.sdkClient()
	if err != nil {
		return accounting.Decimal{}, err
	}

	return c.BalanceGet(ctx, prm)
}
