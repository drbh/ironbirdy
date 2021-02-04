package ironbirdy

import (
	"fmt"
	"math"

	"github.com/drbh/go-quote"
)

func BuildHistory(ticker, start, end string) []DeltaWeeks {
	spy, _ := quote.NewQuoteFromYahoo(ticker, start, end, quote.Daily, true)
	lastWeekSeen := 0
	lastYearSeen := 0
	var test = DeltaWeeks{
		MaxCloseDelta: -9999999999.0,
		MinCloseDelta: 9999999999.0,
	}
	open := false
	var rows = []DeltaWeeks{}
	for i, data := range spy.Date {

		year, week := data.ISOWeek()
		dayOfWeek := int(data.Weekday())
		iosweek := fmt.Sprintf("%d-%d", year, week)

		if week > lastWeekSeen || year > lastYearSeen {
			lastWeekSeen = week // week
			lastYearSeen = year // week

			if dayOfWeek == 1 {
				test.StartPrice = spy.Open[i]
				test.ISOWeek = iosweek
				test.MaxCloseDelta = -9999999999.0
				test.MinCloseDelta = 9999999999.0
				test.StartDate = data
				open = true
			}
		}
		delta := spy.Close[i]/test.StartPrice - 1
		if delta < test.MinCloseDelta {
			test.MinCloseDelta = delta
			test.Low = spy.Close[i]
		}
		if delta > test.MaxCloseDelta {
			if delta < math.Inf(1) {
				test.MaxCloseDelta = delta
				test.High = spy.Close[i]
			}
		}
		if dayOfWeek == 5 {
			test.EndDate = data
			if open {
				rows = append(rows, test)
			}
			open = false
		}
	}
	return rows
}
