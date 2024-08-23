package keeper_test

import (
	"math/big"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/0glabs/0g-chain/chaincfg"
	"github.com/0glabs/0g-chain/x/evmutil/keeper"
	"github.com/0glabs/0g-chain/x/evmutil/testutil"
	"github.com/0glabs/0g-chain/x/evmutil/types"
)

type evmBankKeeperTestSuite struct {
	testutil.Suite
}

func (suite *evmBankKeeperTestSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_ReturnsSpendable() {
	startingCoins := sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10))
	startingEvmDenom := sdkmath.NewInt(100)

	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	bacc := authtypes.NewBaseAccountWithAddress(suite.Addrs[0])
	vacc := vesting.NewContinuousVestingAccount(bacc, startingCoins, now.Unix(), endTime.Unix())
	suite.AccountKeeper.SetAccount(suite.Ctx, vacc)

	err := suite.App.FundAccount(suite.Ctx, suite.Addrs[0], startingCoins)
	suite.Require().NoError(err)
	err = suite.Keeper.SetBalance(suite.Ctx, suite.Addrs[0], startingEvmDenom)
	suite.Require().NoError(err)

	coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.EvmDenom)
	suite.Require().Equal(startingEvmDenom, coin.Amount)

	ctx := suite.Ctx.WithBlockTime(now.Add(12 * time.Hour))
	coin = suite.EvmBankKeeper.GetBalance(ctx, suite.Addrs[0], chaincfg.EvmDenom)
	suite.Require().Equal(sdkmath.NewIntFromUint64(5_000_000_000_100), coin.Amount)
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_NotEvmDenom() {
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.GasDenom)
	})
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "busd")
	})
}

func (suite *evmBankKeeperTestSuite) TestGetBalance() {
	tests := []struct {
		name           string
		startingAmount sdk.Coins
		expAmount      sdkmath.Int
	}{
		{
			"gas denom with evm denom",
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 100),
				sdk.NewInt64Coin(chaincfg.GasDenom, 10),
			),
			sdkmath.NewInt(10_000_000_000_100),
		},
		{
			"just evm denom",
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 100),
				sdk.NewInt64Coin("busd", 100),
			),
			sdkmath.NewInt(100),
		},
		{
			"just gas denom",
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.GasDenom, 10),
				sdk.NewInt64Coin("busd", 100),
			),
			sdkmath.NewInt(10_000_000_000_000),
		},
		{
			"no gas denom or evm denom",
			sdk.NewCoins(),
			sdk.ZeroInt(),
		},
		{
			"with avaka that is more than 1 gas denom",
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 20_000_000_000_220),
				sdk.NewInt64Coin(chaincfg.GasDenom, 11),
			),
			sdkmath.NewInt(31_000_000_000_220),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithZgChain(suite.Addrs[0], tt.startingAmount)
			coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.EvmDenom)
			suite.Require().Equal(tt.expAmount, coin.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromModuleToAccount() {
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin(chaincfg.EvmDenom, 200),
		sdk.NewInt64Coin(chaincfg.GasDenom, 100),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		startingAccBal sdk.Coins
		expAccBal      sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_000_000_000_010)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 10),
				sdk.NewInt64Coin(chaincfg.GasDenom, 12),
			),
			false,
		},
		{
			"send less than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 122)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 122),
				sdk.NewInt64Coin(chaincfg.GasDenom, 0),
			),
			false,
		},
		{
			"send an exact amount of gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 98_000_000_000_000)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 0o0),
				sdk.NewInt64Coin(chaincfg.GasDenom, 98),
			),
			false,
		},
		{
			"send no evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 0)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 0),
				sdk.NewInt64Coin(chaincfg.GasDenom, 0),
			),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total evm denom to cover",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough gas denom to cover",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts receiver's evm denom to gas denom if there's enough evm denom after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 99_000_000_000_200)),
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 999_999_999_900),
				sdk.NewInt64Coin(chaincfg.GasDenom, 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 100),
				sdk.NewInt64Coin(chaincfg.GasDenom, 101),
			),
			false,
		},
		{
			"converts all of receiver's evm denom to gas denom even if somehow receiver has more than 1 gas denom of evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_000_000_000_100)),
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 5_999_999_999_990),
				sdk.NewInt64Coin(chaincfg.GasDenom, 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 90),
				sdk.NewInt64Coin(chaincfg.GasDenom, 19),
			),
			false,
		},
		{
			"swap 1 gas denom for evm denom if module account doesn't have enough evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 99_000_000_001_000)),
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 200),
				sdk.NewInt64Coin(chaincfg.GasDenom, 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin(chaincfg.EvmDenom, 1200),
				sdk.NewInt64Coin(chaincfg.GasDenom, 100),
			),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithZgChain(suite.Addrs[0], tt.startingAccBal)
			suite.FundModuleAccountWithZgChain(evmtypes.ModuleName, startingModuleCoins)

			// fund our module with some gas denom to account for converting extra evm denom back to gas denom
			suite.FundModuleAccountWithZgChain(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10)))

			err := suite.EvmBankKeeper.SendCoinsFromModuleToAccount(suite.Ctx, evmtypes.ModuleName, suite.Addrs[0], tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check gas denom
			GasDenomSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.GasDenom)
			suite.Require().Equal(tt.expAccBal.AmountOf(chaincfg.GasDenom).Int64(), GasDenomSender.Amount.Int64())

			// check evm denom
			actualEvmDenom := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expAccBal.AmountOf(chaincfg.EvmDenom).Int64(), actualEvmDenom.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromAccountToModule() {
	startingAccCoins := sdk.NewCoins(
		sdk.NewInt64Coin(chaincfg.EvmDenom, 200),
		sdk.NewInt64Coin(chaincfg.GasDenom, 100),
	)
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_000),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		expSenderCoins sdk.Coins
		expModuleCoins sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_000_000_000_010)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 190), sdk.NewInt64Coin(chaincfg.GasDenom, 88)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_010), sdk.NewInt64Coin(chaincfg.GasDenom, 12)),
			false,
		},
		{
			"send less than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 122)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 78), sdk.NewInt64Coin(chaincfg.GasDenom, 100)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_122), sdk.NewInt64Coin(chaincfg.GasDenom, 0)),
			false,
		},
		{
			"send an exact amount of gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 98_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 200), sdk.NewInt64Coin(chaincfg.GasDenom, 2)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_000), sdk.NewInt64Coin(chaincfg.GasDenom, 98)),
			false,
		},
		{
			"send no evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 0)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 200), sdk.NewInt64Coin(chaincfg.GasDenom, 100)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_000), sdk.NewInt64Coin(chaincfg.GasDenom, 0)),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin(chaincfg.EvmDenom, 12_000_000_000_000),
				sdk.NewInt64Coin(chaincfg.EvmDenom, 2_000_000_000_000),
			},
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total evm denom to cover",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough gas denom to cover",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts 1 gas denom to evm denom if not enough evm denom to cover",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 99_001_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 999_000_000_200), sdk.NewInt64Coin(chaincfg.GasDenom, 0)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 101_000_000_000), sdk.NewInt64Coin(chaincfg.GasDenom, 99)),
			false,
		},
		{
			"converts receiver's evm denom to gas denom if there's enough evm denom after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 5_900_000_000_200)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_000_000_000), sdk.NewInt64Coin(chaincfg.GasDenom, 94)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 200), sdk.NewInt64Coin(chaincfg.GasDenom, 6)),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundAccountWithZgChain(suite.Addrs[0], startingAccCoins)
			suite.FundModuleAccountWithZgChain(evmtypes.ModuleName, startingModuleCoins)

			err := suite.EvmBankKeeper.SendCoinsFromAccountToModule(suite.Ctx, suite.Addrs[0], evmtypes.ModuleName, tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check sender balance
			GasDenomSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.GasDenom)
			suite.Require().Equal(tt.expSenderCoins.AmountOf(chaincfg.GasDenom).Int64(), GasDenomSender.Amount.Int64())
			actualEvmDenom := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expSenderCoins.AmountOf(chaincfg.EvmDenom).Int64(), actualEvmDenom.Int64())

			// check module balance
			moduleAddr := suite.AccountKeeper.GetModuleAddress(evmtypes.ModuleName)
			GasDenomSender = suite.BankKeeper.GetBalance(suite.Ctx, moduleAddr, chaincfg.GasDenom)
			suite.Require().Equal(tt.expModuleCoins.AmountOf(chaincfg.GasDenom).Int64(), GasDenomSender.Amount.Int64())
			actualEvmDenom = suite.Keeper.GetBalance(suite.Ctx, moduleAddr)
			suite.Require().Equal(tt.expModuleCoins.AmountOf(chaincfg.EvmDenom).Int64(), actualEvmDenom.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestBurnCoins() {
	startingGasDenom := sdkmath.NewInt(100)
	tests := []struct {
		name          string
		burnCoins     sdk.Coins
		expGasDenom   sdkmath.Int
		expEvmDenom   sdkmath.Int
		hasErr        bool
		evmDenomStart sdkmath.Int
	}{
		{
			"burn more than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_021_000_000_002)),
			sdkmath.NewInt(88),
			sdkmath.NewInt(100_000_000_000),
			false,
			sdkmath.NewInt(121_000_000_002),
		},
		{
			"burn less than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 122)),
			sdkmath.NewInt(100),
			sdkmath.NewInt(878),
			false,
			sdkmath.NewInt(1000),
		},
		{
			"burn an exact amount of gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 98_000_000_000_000)),
			sdkmath.NewInt(2),
			sdkmath.NewInt(10),
			false,
			sdkmath.NewInt(10),
		},
		{
			"burn no evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 0)),
			startingGasDenom,
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"errors if burning other coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 500), sdk.NewInt64Coin("busd", 1000)),
			startingGasDenom,
			sdkmath.NewInt(100),
			true,
			sdkmath.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin(chaincfg.EvmDenom, 12_000_000_000_000),
				sdk.NewInt64Coin(chaincfg.EvmDenom, 2_000_000_000_000),
			},
			startingGasDenom,
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"errors if burn amount is negative",
			sdk.Coins{sdk.Coin{Denom: chaincfg.EvmDenom, Amount: sdkmath.NewInt(-100)}},
			startingGasDenom,
			sdkmath.NewInt(50),
			true,
			sdkmath.NewInt(50),
		},
		{
			"errors if not enough evm denom to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100_999_000_000_000)),
			sdkmath.NewInt(0),
			sdkmath.NewInt(99_000_000_000),
			true,
			sdkmath.NewInt(99_000_000_000),
		},
		{
			"errors if not enough gas denom to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 200_000_000_000_000)),
			sdkmath.NewInt(100),
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"converts 1 gas denom to evm denom if not enough evm denom to cover",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_021_000_000_002)),
			sdkmath.NewInt(87),
			sdkmath.NewInt(980_000_000_000),
			false,
			sdkmath.NewInt(1_000_000_002),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			startingCoins := sdk.NewCoins(
				sdk.NewCoin(chaincfg.GasDenom, startingGasDenom),
				sdk.NewCoin(chaincfg.EvmDenom, tt.evmDenomStart),
			)
			suite.FundModuleAccountWithZgChain(evmtypes.ModuleName, startingCoins)

			err := suite.EvmBankKeeper.BurnCoins(suite.Ctx, evmtypes.ModuleName, tt.burnCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check gas denom
			GasDenomActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, chaincfg.GasDenom)
			suite.Require().Equal(tt.expGasDenom, GasDenomActual.Amount)

			// check evm denom
			evmDenomActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.expEvmDenom, evmDenomActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestMintCoins() {
	tests := []struct {
		name          string
		mintCoins     sdk.Coins
		GasDenomCnt   sdkmath.Int
		evmDenomCnt   sdkmath.Int
		hasErr        bool
		evmDenomStart sdkmath.Int
	}{
		{
			"mint more than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_021_000_000_002)),
			sdkmath.NewInt(12),
			sdkmath.NewInt(21_000_000_002),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint less than 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 901_000_000_001)),
			sdk.ZeroInt(),
			sdkmath.NewInt(901_000_000_001),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint an exact amount of gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 123_000_000_000_000_000)),
			sdkmath.NewInt(123_000),
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint no evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 0)),
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"errors if minting other coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.ZeroInt(),
			sdkmath.NewInt(100),
			true,
			sdkmath.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin(chaincfg.EvmDenom, 12_000_000_000_000),
				sdk.NewInt64Coin(chaincfg.EvmDenom, 2_000_000_000_000),
			},
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"errors if mint amount is negative",
			sdk.Coins{sdk.Coin{Denom: chaincfg.EvmDenom, Amount: sdkmath.NewInt(-100)}},
			sdk.ZeroInt(),
			sdkmath.NewInt(50),
			true,
			sdkmath.NewInt(50),
		},
		{
			"adds to existing evm denom balance",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 12_021_000_000_002)),
			sdkmath.NewInt(12),
			sdkmath.NewInt(21_000_000_102),
			false,
			sdkmath.NewInt(100),
		},
		{
			"convert evm denom balance to gas denom if it exceeds 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 10_999_000_000_000)),
			sdkmath.NewInt(12),
			sdkmath.NewInt(1_200_000_001),
			false,
			sdkmath.NewInt(1_002_200_000_001),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundModuleAccountWithZgChain(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10)))
			suite.FundModuleAccountWithZgChain(evmtypes.ModuleName, sdk.NewCoins(sdk.NewCoin(chaincfg.EvmDenom, tt.evmDenomStart)))

			err := suite.EvmBankKeeper.MintCoins(suite.Ctx, evmtypes.ModuleName, tt.mintCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check gas denom
			GasDenomActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, chaincfg.GasDenom)
			suite.Require().Equal(tt.GasDenomCnt, GasDenomActual.Amount)

			// check evm denom
			evmDenomActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.evmDenomCnt, evmDenomActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestValidateEvmCoins() {
	tests := []struct {
		name      string
		coins     sdk.Coins
		shouldErr bool
	}{
		{
			"valid coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 500)),
			false,
		},
		{
			"dup coins",
			sdk.Coins{sdk.NewInt64Coin(chaincfg.EvmDenom, 500), sdk.NewInt64Coin(chaincfg.EvmDenom, 500)},
			true,
		},
		{
			"not evm coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 500)),
			true,
		},
		{
			"negative coins",
			sdk.Coins{sdk.Coin{Denom: chaincfg.EvmDenom, Amount: sdkmath.NewInt(-500)}},
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := keeper.ValidateEvmCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertOneGasDenomToEvmDenomIfNeeded() {
	evmDenomNeeded := sdkmath.NewInt(200)
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
		success       bool
	}{
		{
			"not enough gas denom for conversion",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100)),
			false,
		},
		{
			"converts 1 gas denom to evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10), sdk.NewInt64Coin(chaincfg.EvmDenom, 100)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 9), sdk.NewInt64Coin(chaincfg.EvmDenom, 1_000_000_000_100)),
			true,
		},
		{
			"conversion not needed",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10), sdk.NewInt64Coin(chaincfg.EvmDenom, 200)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10), sdk.NewInt64Coin(chaincfg.EvmDenom, 200)),
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithZgChain(suite.Addrs[0], tt.startingCoins)
			err := suite.EvmBankKeeper.ConvertOneGasDenomToEvmDenomIfNeeded(suite.Ctx, suite.Addrs[0], evmDenomNeeded)
			moduleZgChain := suite.BankKeeper.GetBalance(suite.Ctx, suite.AccountKeeper.GetModuleAddress(types.ModuleName), chaincfg.GasDenom)
			if tt.success {
				suite.Require().NoError(err)
				if tt.startingCoins.AmountOf(chaincfg.EvmDenom).LT(evmDenomNeeded) {
					suite.Require().Equal(sdk.OneInt(), moduleZgChain.Amount)
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(sdk.ZeroInt(), moduleZgChain.Amount)
			}

			evmDenomCnt := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf(chaincfg.EvmDenom), evmDenomCnt)
			GasDenomCoin := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.GasDenom)
			suite.Require().Equal(tt.expectedCoins.AmountOf(chaincfg.GasDenom), GasDenomCoin.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertEvmDenomToGasDenom() {
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
	}{
		{
			"not enough gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 100), sdk.NewInt64Coin(chaincfg.GasDenom, 0)),
		},
		{
			"converts evm denom for 1 gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10), sdk.NewInt64Coin(chaincfg.EvmDenom, 1_000_000_000_003)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 11), sdk.NewInt64Coin(chaincfg.EvmDenom, 3)),
		},
		{
			"converts more than 1 gas denom of evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10), sdk.NewInt64Coin(chaincfg.EvmDenom, 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 18), sdk.NewInt64Coin(chaincfg.EvmDenom, 123)),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			err := suite.App.FundModuleAccount(suite.Ctx, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 10)))
			suite.Require().NoError(err)
			suite.FundAccountWithZgChain(suite.Addrs[0], tt.startingCoins)
			err = suite.EvmBankKeeper.ConvertEvmDenomToGasDenom(suite.Ctx, suite.Addrs[0])
			suite.Require().NoError(err)
			evmDenomCnt := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf(chaincfg.EvmDenom), evmDenomCnt)
			GasDenomCoin := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], chaincfg.GasDenom)
			suite.Require().Equal(tt.expectedCoins.AmountOf(chaincfg.GasDenom), GasDenomCoin.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSplitEvmDenomCoins() {
	tests := []struct {
		name          string
		coins         sdk.Coins
		expectedCoins sdk.Coins
		shouldErr     bool
	}{
		{
			"invalid coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 500)),
			nil,
			true,
		},
		{
			"empty coins",
			sdk.NewCoins(),
			sdk.NewCoins(),
			false,
		},
		{
			"gas denom & evm denom coins",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 8), sdk.NewInt64Coin(chaincfg.EvmDenom, 123)),
			false,
		},
		{
			"only evm denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 10_123)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 10_123)),
			false,
		},
		{
			"only gas denom",
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.EvmDenom, 5_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin(chaincfg.GasDenom, 5)),
			false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			GasDenomCoin, evmDenomCnt, err := keeper.SplitEvmDenomCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tt.expectedCoins.AmountOf(chaincfg.GasDenom), GasDenomCoin.Amount)
				suite.Require().Equal(tt.expectedCoins.AmountOf(chaincfg.EvmDenom), evmDenomCnt)
			}
		})
	}
}

func TestEvmBankKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(evmBankKeeperTestSuite))
}

func makeBigIntByString(s string) *big.Int {
	i, _ := new(big.Int).SetString(s, 10)
	return i
}
