package ranking

import "mh-api/internal/domain/monsters"

type Ranking struct {
	monsterId monsters.MonsterId
	ranking   Rank
	voteYear  VoteYear
}

func NewRanking(monsterId, rank, voteYear string) *Ranking {
	return newRanking(
		monsters.MonsterId{Value: monsterId},
		Rank{value: rank},
		VoteYear{value: voteYear},
	)
}

func newRanking(monsterId monsters.MonsterId, ranking Rank, voteYear VoteYear) *Ranking {
	return &Ranking{
		monsterId: monsterId,
		ranking:   ranking,
		voteYear:  voteYear,
	}
}
