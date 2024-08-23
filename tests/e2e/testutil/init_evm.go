package testutil

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/0glabs/0g-chain/tests/e2e/contracts/greeter"
	evmutiltypes "github.com/0glabs/0g-chain/x/evmutil/types"
)

// InitZgChainEvmData is run after the chain is running, but before the tests are run.
// It is used to initialize some EVM state, such as deploying contracts.
func (suite *E2eTestSuite) InitZgChainEvmData() {
	whale := suite.ZgChain.GetAccount(FundedAccountName)

	// ensure funded account has nonzero erc20 balance
	balance := suite.ZgChain.GetErc20Balance(suite.DeployedErc20.Address, whale.EvmAddress)
	if balance.Cmp(big.NewInt(0)) != 1 {
		panic(fmt.Sprintf(
			"expected funded account (%s) to have erc20 balance of token %s",
			whale.EvmAddress.Hex(),
			suite.DeployedErc20.Address.Hex(),
		))
	}

	// expect the erc20 to be enabled for conversion to sdk.Coin
	params, err := suite.ZgChain.Grpc.Query.Evmutil.Params(context.Background(), &evmutiltypes.QueryParamsRequest{})
	if err != nil {
		panic(fmt.Sprintf("failed to fetch evmutil params during init: %s", err))
	}
	found := false
	erc20Addr := suite.DeployedErc20.Address.Hex()
	for _, p := range params.Params.EnabledConversionPairs {
		if common.BytesToAddress(p.ZgChainERC20Address).Hex() == erc20Addr {
			found = true
			suite.DeployedErc20.CosmosDenom = p.Denom
		}
	}
	if !found {
		panic(fmt.Sprintf("erc20 %s must be enabled for conversion to cosmos coin", erc20Addr))
	}
	suite.ZgChain.RegisterErc20(suite.DeployedErc20.Address)

	// deploy an example contract
	greeterAddr, _, _, err := greeter.DeployGreeter(
		whale.evmSigner.Auth,
		whale.evmSigner.EvmClient,
		"what's up!",
	)
	suite.NoError(err, "failed to deploy a contract to the EVM")
	suite.ZgChain.ContractAddrs["greeter"] = greeterAddr
}

// FundZgChainErc20Balance sends the pre-deployed ERC20 token to the `toAddress`.
func (suite *E2eTestSuite) FundZgChainErc20Balance(toAddress common.Address, amount *big.Int) EvmTxResponse {
	// funded account should have erc20 balance
	whale := suite.ZgChain.GetAccount(FundedAccountName)
	res, err := whale.TransferErc20(suite.DeployedErc20.Address, toAddress, amount)
	suite.NoError(err)
	return res
}
