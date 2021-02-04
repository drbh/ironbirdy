package ironbirdy

import (
	"sort"

	"github.com/drbh/go-robinhood"
)

func BuildSpreads(c *robinhood.Client, insts []*robinhood.OptionInstrument, stockPrice float64) []SpreadLeg {
	sort.Slice(insts, func(i, j int) bool {
		return insts[i].StrikePrice < insts[j].StrikePrice
	})
	var entries = []LegEntry{}
	for i, strikeOpt := range insts {
		is, _ := c.MarketData(strikeOpt)

		if len(is) < 1 {
			continue
		}
		market := is[0]
		if strikeOpt.StrikePrice < stockPrice {
			// this call is in the money
			continue
		}
		distanceFromCurrentPrice := (strikeOpt.StrikePrice / stockPrice) - 1
		// if distanceFromCurrentPrice < 0.06 {
		// 	// this call is too close to the money
		// 	continue
		// }
		entry := LegEntry{
			ID:       i,
			Strike:   strikeOpt.StrikePrice,
			Ask:      market.AskPrice,
			Bid:      market.BidPrice,
			Fair:     market.AdjustedMarkPrice,
			Distance: distanceFromCurrentPrice,
		}
		entries = append(entries, entry)
	}

	var spreads = []SpreadLeg{}
	for i := 0; i < len(entries); i++ {
		subset := entries[i:len(entries)]
		shortEntry := subset[0]
		for j := 1; j < len(subset); j++ {
			longEntry := subset[j]
			fairDiff := shortEntry.Fair - longEntry.Fair
			fairDiffBid := shortEntry.Bid - longEntry.Ask
			fairDiffAsk := shortEntry.Ask - longEntry.Bid
			strikeDiff := longEntry.Strike - shortEntry.Strike

			spreadMark := (fairDiffAsk + fairDiffBid) / 2

			// build our spread from the long+short
			spread := SpreadLeg{
				LongStrike:    longEntry.Strike,
				ShortStrike:   shortEntry.Strike,
				LongPrice:     longEntry.Fair,
				ShortPrice:    shortEntry.Fair,
				FairDiff:      fairDiff,
				FairDiffBid:   fairDiffBid,
				FairDiffAsk:   fairDiffAsk,
				FairDiffFair:  spreadMark,
				StrikeDiff:    strikeDiff,
				PossibleRet:   spreadMark / strikeDiff,
				LongDistance:  longEntry.Distance,
				ShortDistance: shortEntry.Distance,
			}

			spreads = append(spreads, spread)

		}
	}
	return spreads
}

func BuildSpreadsPuts(c *robinhood.Client, insts []*robinhood.OptionInstrument, stockPrice float64) []SpreadLeg {
	sort.Slice(insts, func(i, j int) bool {
		return insts[i].StrikePrice > insts[j].StrikePrice
	})
	var entries = []LegEntry{}
	for i, strikeOpt := range insts {
		is, _ := c.MarketData(strikeOpt)
		if len(is) < 1 {
			continue
		}
		market := is[0]
		if strikeOpt.StrikePrice > stockPrice {
			// this call is in the money
			continue
		}
		distanceFromCurrentPrice := (strikeOpt.StrikePrice / stockPrice) - 1
		// if distanceFromCurrentPrice > -0.06 {
		// 	// this call is too close to the money
		// 	continue
		// }
		entry := LegEntry{
			ID:       i,
			Strike:   strikeOpt.StrikePrice,
			Ask:      market.AskPrice,
			Bid:      market.BidPrice,
			Fair:     market.AdjustedMarkPrice,
			Distance: distanceFromCurrentPrice,
		}
		entries = append(entries, entry)
	}

	var spreads = []SpreadLeg{}
	for i := 0; i < len(entries); i++ {
		subset := entries[i:len(entries)]
		shortEntry := subset[0]
		for j := 1; j < len(subset); j++ {
			longEntry := subset[j]
			fairDiff := shortEntry.Fair - longEntry.Fair
			// fairDiffBid := shortEntry.Bid - longEntry.Ask
			// fairDiffAsk := shortEntry.Ask - longEntry.Bid
			fairDiffBid := shortEntry.Bid - longEntry.Ask
			fairDiffAsk := shortEntry.Ask - longEntry.Bid
			strikeDiff := shortEntry.Strike - longEntry.Strike

			spreadMark := (fairDiffAsk + fairDiffBid) / 2

			// build our spread from the long+short
			spread := SpreadLeg{
				LongStrike:    longEntry.Strike,
				ShortStrike:   shortEntry.Strike,
				LongPrice:     longEntry.Fair,
				ShortPrice:    shortEntry.Fair,
				FairDiff:      fairDiff,
				FairDiffBid:   fairDiffBid,
				FairDiffAsk:   fairDiffAsk,
				FairDiffFair:  spreadMark,
				StrikeDiff:    strikeDiff,
				PossibleRet:   spreadMark / strikeDiff,
				LongDistance:  longEntry.Distance,
				ShortDistance: shortEntry.Distance,
			}

			spreads = append(spreads, spread)

		}
	}
	return spreads
}
