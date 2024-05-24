package bonds

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	faultTypes "github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/types"
	"github.com/tokamak-network/tokamak-thanos/op-service/sources/batching"
	"golang.org/x/exp/maps"
)

var noBond = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))

type BondContract interface {
	GetCredits(ctx context.Context, block batching.Block, recipients ...common.Address) ([]*big.Int, error)
}

// CalculateRequiredCollateral determines the minimum balance required for a fault dispute game contract in order
// to pay the outstanding bonds and credits.
// It returns the sum of unpaid bonds from claims, plus the sum of allocated but unclaimed credits.
func CalculateRequiredCollateral(ctx context.Context, contract BondContract, blockHash common.Hash, claims []faultTypes.Claim) (*big.Int, error) {
	unpaidBonds := big.NewInt(0)
	recipients := make(map[common.Address]bool)
	for _, claim := range claims {
		if noBond.Cmp(claim.Bond) != 0 {
			unpaidBonds = new(big.Int).Add(unpaidBonds, claim.Bond)
		}
		recipients[claim.Claimant] = true
		if claim.CounteredBy != (common.Address{}) {
			recipients[claim.CounteredBy] = true
		}
	}

	credits, err := contract.GetCredits(ctx, batching.BlockByHash(blockHash), maps.Keys(recipients)...)
	if err != nil {
		return nil, fmt.Errorf("failed to load credits: %w", err)
	}
	for _, credit := range credits {
		unpaidBonds = new(big.Int).Add(unpaidBonds, credit)
	}
	return unpaidBonds, nil
}
