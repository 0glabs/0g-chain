package pricefeed

import (
	"time"

	"github.com/0glabs/0g-chain/x/pricefeed/keeper"
	"github.com/0glabs/0g-chain/x/pricefeed/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.SetCurrentPricesForAllMarkets(ctx)
}
