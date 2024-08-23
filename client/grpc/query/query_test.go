package query_test

import (
	"testing"

	"github.com/0glabs/0g-chain/client/grpc/query"
	"github.com/stretchr/testify/require"
)

func TestNewQueryClient_InvalidGprc(t *testing.T) {
	t.Run("valid connection", func(t *testing.T) {
		conn, err := query.NewQueryClient("http://localhost:1234")
		require.NoError(t, err)
		require.NotNil(t, conn)
	})

	t.Run("non-empty url", func(t *testing.T) {
		_, err := query.NewQueryClient("")
		require.ErrorContains(t, err, "unknown grpc url scheme")
	})

	t.Run("invalid url scheme", func(t *testing.T) {
		_, err := query.NewQueryClient("ftp://localhost:1234")
		require.ErrorContains(t, err, "unknown grpc url scheme")
	})
}

func TestNewQueryClient_ValidClient(t *testing.T) {
	t.Run("all clients are created", func(t *testing.T) {
		client, err := query.NewQueryClient("http://localhost:1234")
		require.NoError(t, err)
		require.NotNil(t, client)

		// validate cosmos clients
		require.NotNil(t, client.Tm)
		require.NotNil(t, client.Tx)
		require.NotNil(t, client.Auth)
		require.NotNil(t, client.Authz)
		require.NotNil(t, client.Bank)
		require.NotNil(t, client.Distribution)
		require.NotNil(t, client.Evidence)
		require.NotNil(t, client.Gov)
		require.NotNil(t, client.GovBeta)
		require.NotNil(t, client.Mint)
		require.NotNil(t, client.Params)
		require.NotNil(t, client.Slashing)
		require.NotNil(t, client.Staking)
		require.NotNil(t, client.Upgrade)
		require.NotNil(t, client.Consensus)

		// validate 3rd party clients
		require.NotNil(t, client.Evm)
		require.NotNil(t, client.Feemarket)
		require.NotNil(t, client.IbcClient)
		require.NotNil(t, client.IbcTransfer)

		// validate kava clients
		require.NotNil(t, client.Bep3)
		require.NotNil(t, client.Committee)
		require.NotNil(t, client.Evmutil)
		require.NotNil(t, client.Issuance)
		require.NotNil(t, client.Pricefeed)
	})
}
