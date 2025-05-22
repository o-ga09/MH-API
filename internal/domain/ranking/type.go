package ranking

type Rankings []Ranking

type Rank struct{ value string }
type VoteYear struct{ value string }

func (r *Ranking) GetID() string       { return r.monsterId.Value }
func (r *Ranking) GetRank() string     { return r.ranking.value }
func (r *Ranking) GetVoteYear() string { return r.voteYear.value }
