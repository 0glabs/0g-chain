package e2e_test

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/0glabs/0g-chain/chaincfg"
	evmutiltypes "github.com/0glabs/0g-chain/x/evmutil/types"
)

func (suite *IntegrationTestSuite) TestGrpcClientQueryCosmosModule_Balance() {
	// ARRANGE
	// setup 0g account
	funds := chaincfg.MakeCoinForGasDenom(1e5)
	zgAcc := suite.ZgChain.NewFundedAccount("balance-test", sdk.NewCoins(funds))

	// ACT
	rsp, err := suite.ZgChain.Grpc.Query.Bank.Balance(context.Background(), &banktypes.QueryBalanceRequest{
		Address: zgAcc.SdkAddress.String(),
		Denom:   funds.Denom,
	})

	// ASSERT
	suite.Require().NoError(err)
	suite.Require().Equal(funds.Amount, rsp.Balance.Amount)
}

func (suite *IntegrationTestSuite) TestGrpcClientQueryKavaModule_EvmParams() {
	// ACT
	rsp, err := suite.ZgChain.Grpc.Query.Evmutil.Params(
		context.Background(), &evmutiltypes.QueryParamsRequest{},
	)

	// ASSERT
	suite.Require().NoError(err)
	suite.Require().GreaterOrEqual(len(rsp.Params.AllowedCosmosDenoms), 1)
	suite.Require().GreaterOrEqual(len(rsp.Params.EnabledConversionPairs), 1)
}
