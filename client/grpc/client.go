package grpc

import (
	"errors"

	"github.com/0glabs/0g-chain/client/grpc/query"
	"github.com/0glabs/0g-chain/client/grpc/util"
)

// ZgChainGrpcClient enables the usage of kava grpc query clients and query utils
type ZgChainGrpcClient struct {
	config ZgChainGrpcClientConfig

	// Query clients for cosmos and kava modules
	Query *query.QueryClient

	// Utils for common queries (ie fetch an unpacked BaseAccount)
	*util.Util
}

// ZgChainGrpcClientConfig is a configuration struct for a ZgChainGrpcClient
type ZgChainGrpcClientConfig struct {
	// note: add future config options here
}

// NewClient creates a new ZgChainGrpcClient via a grpc url
func NewClient(grpcUrl string) (*ZgChainGrpcClient, error) {
	return NewClientWithConfig(grpcUrl, NewDefaultConfig())
}

// NewClientWithConfig creates a new ZgChainGrpcClient via a grpc url and config
func NewClientWithConfig(grpcUrl string, config ZgChainGrpcClientConfig) (*ZgChainGrpcClient, error) {
	if grpcUrl == "" {
		return nil, errors.New("grpc url cannot be empty")
	}
	query, error := query.NewQueryClient(grpcUrl)
	if error != nil {
		return nil, error
	}
	client := &ZgChainGrpcClient{
		Query:  query,
		Util:   util.NewUtil(query),
		config: config,
	}
	return client, nil
}

func NewDefaultConfig() ZgChainGrpcClientConfig {
	return ZgChainGrpcClientConfig{}
}
