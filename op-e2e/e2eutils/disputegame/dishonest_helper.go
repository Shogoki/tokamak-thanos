package disputegame

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/types"
)

type dishonestClaim struct {
	ParentIndex int64
	IsAttack    bool
	Valid       bool
}

type DishonestHelper struct {
	*OutputGameHelper
	*OutputHonestHelper
	claims   map[dishonestClaim]bool
	defender bool
}

func newDishonestHelper(g *OutputGameHelper, correctTrace *OutputHonestHelper, defender bool) *DishonestHelper {
	return &DishonestHelper{g, correctTrace, make(map[dishonestClaim]bool), defender}
}

func (t *DishonestHelper) Attack(ctx context.Context, claimIndex int64) {
	c := dishonestClaim{claimIndex, true, false}
	if t.claims[c] {
		return
	}
	t.claims[c] = true
	t.OutputGameHelper.Attack(ctx, claimIndex, common.Hash{byte(claimIndex)})
}

func (t *DishonestHelper) Defend(ctx context.Context, claimIndex int64) {
	c := dishonestClaim{claimIndex, false, false}
	if t.claims[c] {
		return
	}
	t.claims[c] = true
	t.OutputGameHelper.Defend(ctx, claimIndex, common.Hash{byte(claimIndex)})
}

func (t *DishonestHelper) AttackCorrect(ctx context.Context, claimIndex int64) {
	c := dishonestClaim{claimIndex, true, true}
	if t.claims[c] {
		return
	}
	t.claims[c] = true
	t.OutputHonestHelper.Attack(ctx, claimIndex)
}

func (t *DishonestHelper) DefendCorrect(ctx context.Context, claimIndex int64) {
	c := dishonestClaim{claimIndex, false, true}
	if t.claims[c] {
		return
	}
	t.claims[c] = true
	t.OutputHonestHelper.Defend(ctx, claimIndex)
}

// ExhaustDishonestClaims makes all possible significant moves (mod honest challenger's) in a game.
// It is very inefficient and should NOT be used on games with large depths
func (d *DishonestHelper) ExhaustDishonestClaims(ctx context.Context, rootClaim *ClaimHelper) {
	depth := d.MaxDepth(ctx)
	splitDepth := d.SplitDepth(ctx)

	move := func(claimIndex int64, claimData ContractClaim) {
		// dishonest level, valid attack
		// dishonest level, invalid attack
		// dishonest level, valid defense
		// dishonest level, invalid defense
		// honest level, invalid attack
		// honest level, invalid defense

		pos := types.NewPositionFromGIndex(claimData.Position)
		if pos.Depth() == depth {
			return
		}

		d.LogGameData(ctx)
		d.OutputGameHelper.t.Logf("Dishonest moves against claimIndex %d", claimIndex)
		agreeWithLevel := d.defender == (pos.Depth()%2 == 0)
		if !agreeWithLevel {
			d.AttackCorrect(ctx, claimIndex)
			if claimIndex != 0 && pos.Depth() != splitDepth+1 {
				d.DefendCorrect(ctx, claimIndex)
			}
		}
		d.Attack(ctx, claimIndex)
		if claimIndex != 0 && pos.Depth() != splitDepth+1 {
			d.Defend(ctx, claimIndex)
		}
	}

	numClaimsSeen := rootClaim.index
	for {
		// Use a short timeout since we don't know the challenger will respond,
		// and this is only designed for the alphabet game where the response should be fast.
		newCount, err := d.waitForNewClaim(ctx, numClaimsSeen, 30*time.Second)
		if errors.Is(err, context.DeadlineExceeded) {
			// we assume that the honest challenger has stopped responding
			// There's nothing to respond to.
			break
		}
		d.OutputGameHelper.require.NoError(err)

		for ; numClaimsSeen < newCount; numClaimsSeen++ {
			claimData := d.getClaim(ctx, numClaimsSeen)
			move(numClaimsSeen, claimData)
		}
	}
}
