package pricefeed_test

import (
	"testing"

	"github.com/0glabs/0g-chain/x/pricefeed"
	"github.com/0glabs/0g-chain/x/pricefeed/keeper"
	"github.com/0glabs/0g-chain/x/pricefeed/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestEndBlocker_UpdatesMultipleMarkets(t *testing.T) {
	testutil.SetCurrentPrices_PriceCalculations(t, func(ctx sdk.Context, keeper keeper.Keeper) {
		pricefeed.EndBlocker(ctx, keeper)
	})

	testutil.SetCurrentPrices_EventEmission(t, func(ctx sdk.Context, keeper keeper.Keeper) {
		pricefeed.EndBlocker(ctx, keeper)
	})
}
