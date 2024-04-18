package signer

import (
	"context"
)

func (n *OracleNode) MtA(ctx context.Context, request *MtARequest) (*MtAResponse, error) {

	result := n.signerNode.MtA(request.K, request.Index)

	return &MtAResponse{B: result}, nil

}
