package common

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	didv1 "github.com/sonr-io/core/api/did/v1"
	dwnv1 "github.com/sonr-io/core/api/dwn/v1"
	svcv1 "github.com/sonr-io/core/api/svc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	BalanceRequest        = bankv1beta1.QueryBalanceRequest        // BalanceRequest is the request type for the Bank.Balance RPC method.
	BalanceResponse       = bankv1beta1.QueryBalanceResponse       // BalanceResponse is the response type for the Bank.Balance RPC method.
	AllBalancesRequest    = bankv1beta1.QueryAllBalancesRequest    // AllBalancesRequest is the request type for the Bank.AllBalances RPC method.
	AllBalancesResponse   = bankv1beta1.QueryAllBalancesResponse   // AllBalancesResponse is the response type for the Bank.AllBalances RPC method.
	TotalSupplyRequest    = bankv1beta1.QueryTotalSupplyRequest    // TotalSupplyRequest is the request type for the Bank.TotalSupply RPC method.
	TotalSupplyResponse   = bankv1beta1.QueryTotalSupplyResponse   // TotalSupplyResponse is the response type for the Bank.TotalSupply RPC method.
	DenomMetadataRequest  = bankv1beta1.QueryDenomMetadataRequest  // DenomMetadataRequest is the request type for the Bank.DenomMetadata RPC method.
	DenomMetadataResponse = bankv1beta1.QueryDenomMetadataResponse // DenomMetadataResponse is the response type for the Bank.DenomMetadata RPC method.
	BankParamsRequest     = bankv1beta1.QueryParamsRequest         // BankParamsRequest is the request type for the Bank.Params RPC method.
	BankParamsResponse    = bankv1beta1.QueryParamsResponse        // BankParamsResponse is the response type for the Bank.Params RPC method.
	DIDParamsRequest      = didv1.QueryRequest                     // DIDParamsRequest is the request type for the DID.Params RPC method.
	DIDParamsResponse     = didv1.QueryParamsResponse              // DIDParamsResponse is the response type for the DID.Params RPC method.
	DIDResolveResponse    = didv1.QueryResolveResponse             // DIDResolveResponse is the response type for the DID.Resolve RPC method.
	DWNParamsRequest      = dwnv1.QueryParamsRequest               // DWNParamsRequest is the request type for the DWN.Params RPC method.
	DWNParamsResponse     = dwnv1.QueryParamsResponse              // DWNParamsResponse is the response type for the DWN.Params RPC method.
	SVCParamsRequest      = svcv1.QueryParamsRequest               // SVCParamsRequest is the request type for the SVC.Params RPC method.
	SVCParamsResponse     = svcv1.QueryParamsResponse              // SVCParamsResponse is the response type for the SVC.Params RPC method.
)

func getConn(addr string) (*grpc.ClientConn, error) {
	grpcConn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return grpcConn, nil
}

func NewBankClient(addr string) (bankv1beta1.QueryClient, error) {
	conn, err := getConn(addr)
	if err != nil {
		return nil, err
	}
	return bankv1beta1.NewQueryClient(conn), nil
}

func NewDIDClient(addr string) (didv1.QueryClient, error) {
	conn, err := getConn(addr)
	if err != nil {
		return nil, err
	}
	return didv1.NewQueryClient(conn), nil
}

func NewDWNClient(addr string) (dwnv1.QueryClient, error) {
	conn, err := getConn(addr)
	if err != nil {
		return nil, err
	}
	return dwnv1.NewQueryClient(conn), nil
}

//
// func NewNodeClient(addr string) (nodev1beta1.ServiceClient, error) {
// 	conn, err := conn(addr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nodev1beta1.NewServiceClient(conn), nil
// }

func NewSVCClient(addr string) (svcv1.QueryClient, error) {
	conn, err := getConn(addr)
	if err != nil {
		return nil, err
	}
	return svcv1.NewQueryClient(conn), nil
}
