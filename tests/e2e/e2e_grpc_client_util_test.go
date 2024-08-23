package e2e_test

import (
	"github.com/0glabs/0g-chain/chaincfg"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestGrpcClientUtil_Account() {
	// ARRANGE
	// setup 0g account
	zgAcc := suite.ZgChain.NewFundedAccount("account-test", sdk.NewCoins(chaincfg.MakeCoinForGasDenom(1e5)))

	// ACT
	rsp, err := suite.ZgChain.Grpc.BaseAccount(zgAcc.SdkAddress.String())

	// ASSERT
	suite.Require().NoError(err)
	suite.Equal(zgAcc.SdkAddress.String(), rsp.Address)
	suite.Greater(rsp.AccountNumber, uint64(1))
	suite.Equal(uint64(0), rsp.Sequence)
}
