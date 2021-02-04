package ironbirdy

import "time"

// Config allows us to parse/read out secrets file
type Config struct {
	Account struct {
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
	} `yaml:"account"`
}

// LegEntry is a single option option
type LegEntry struct {
	ID       int
	Strike   float64
	Ask      float64
	Bid      float64
	Fair     float64
	Distance float64
}

// SpreadLeg is two options that make bull put or bear call spread
type SpreadLeg struct {
	LongStrike    float64
	LongPrice     float64
	ShortStrike   float64
	ShortPrice    float64
	FairDiff      float64
	FairDiffBid   float64
	FairDiffAsk   float64
	FairDiffFair  float64
	StrikeDiff    float64
	PossibleRet   float64
	LongDistance  float64
	ShortDistance float64
}

// DeltaWeeks
type DeltaWeeks struct {
	ISOWeek       string
	StartDate     time.Time
	EndDate       time.Time
	StartPrice    float64
	MaxCloseDelta float64
	MinCloseDelta float64
	Low           float64
	High          float64
}

type IronCondor struct {
	CallLeg    SpreadLeg
	PutLeg     SpreadLeg
	Rar        float64
	Prem       float64
	Collateral float64
	Ret        float64

	CountBelow float64
	CountAbove float64
	CRP        float64
	PRP        float64
	Diff       float64
	Len        float64
	Fail       float64
}
