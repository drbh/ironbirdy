package ironbirdy

import "sort"

func calcAdjusted(a, b, j, k, z, m float64) float64 {
	// a = countBelow(this.state.putSpread.ShortDistance),
	// b = countAbove(this.state.callSpread.ShortDistance),
	// j = this.state.callSpread.PossibleRet,
	// k = this.state.putSpread.PossibleRet,
	// z = this.state.putSpread.StrikeDiff,
	// m = historicalData.length;
	return (-a - b) * (1.0 + (j+k)/z) / m
	// return ((-1 * (a + b)) / m) * (1 + (j + k) / z);
}

func Max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func countAbove(hist []DeltaWeeks, digit float64) float64 {
	var count int = 0
	for _, h := range hist {
		if h.MaxCloseDelta >= digit {
			count++
		}
	}
	return float64(count)
}

func countBelow(hist []DeltaWeeks, digit float64) float64 {
	var count int = 0
	for _, h := range hist {
		if h.MinCloseDelta <= digit {
			count++
		}
	}
	return float64(count)
}

func AdjustRisk(bears []SpreadLeg, bulls []SpreadLeg, hist []DeltaWeeks, minDistanceFromPrice float64) []IronCondor {
	// fmt.Printf("%d %d\n", len(bulls), len(bears))
	// fmt.Printf("%d\n", len(hist))
	var condors = []IronCondor{}
	for _, bearCallSpread := range bears {
		for _, bullPutSpread := range bulls {

			if (bullPutSpread.PossibleRet <= 0.0) || (bearCallSpread.PossibleRet <= 0.0) {
				continue
			}
			if (bullPutSpread.StrikeDiff != 1.0) || (bearCallSpread.StrikeDiff != 1.0) {
				continue
			}
			if (bullPutSpread.ShortDistance >= -minDistanceFromPrice) || (bearCallSpread.ShortDistance <= minDistanceFromPrice) {
				continue
			}

			largerSpread := Max(bearCallSpread.StrikeDiff, bullPutSpread.StrikeDiff)

			below := countBelow(hist, bullPutSpread.ShortDistance)
			above := countAbove(hist, bearCallSpread.ShortDistance)

			rar := calcAdjusted(
				below,
				above,
				bearCallSpread.PossibleRet,
				bullPutSpread.PossibleRet,
				largerSpread,
				float64(len(hist)),
			)
			ic := IronCondor{
				CallLeg:    bearCallSpread,
				PutLeg:     bullPutSpread,
				Rar:        rar,
				Prem:       bullPutSpread.PossibleRet + bearCallSpread.PossibleRet,
				Collateral: largerSpread,
				Ret:        (bullPutSpread.PossibleRet + bearCallSpread.PossibleRet) / largerSpread,

				CountBelow: below,
				CountAbove: above,
				CRP:        bearCallSpread.PossibleRet,
				PRP:        bullPutSpread.PossibleRet,
				Diff:       largerSpread,
				Len:        float64(len(hist)),
				Fail:       (below + above) / float64(len(hist)),
			}
			condors = append(condors, ic)
		}
	}

	// sort.Slice(condors, func(i, j int) bool {
	// 	// return condors[i].Rar > condors[j].Rar
	// 	if condors[i].Rar != condors[j].Rar {
	// 		return condors[i].Rar > condors[j].Rar
	// 	}
	// 	return condors[i].Ret > condors[j].Ret
	// })

	sort.Slice(condors, func(i, j int) bool {
		if condors[i].Ret != condors[j].Ret {
			return condors[i].Ret > condors[j].Ret
		}
		if condors[i].Rar != condors[j].Rar {
			return condors[i].Rar > condors[j].Rar
		}
		return condors[i].Fail < condors[j].Fail
	})

	return condors

}
