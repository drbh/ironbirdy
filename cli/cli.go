package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/drbh/go-robinhood"
	"github.com/drbh/ironbirdy"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func calculate(configPath string, tickerSymbol string, expireDateNext string, historyStart string, historyEnd string, minDistanceFromPrice float64) []ironbirdy.IronCondor {

	fmt.Printf("🎬 %s\n", "Starting")

	fmt.Printf("🐓 %s: %s \n", "Ticker", tickerSymbol)
	fmt.Printf("🐓 %s: %s and %s \n", "Bewteen", historyStart, historyEnd)

	f, err := os.Open(configPath)
	if err != nil {
		spew.Dump(err)
	}
	defer f.Close()

	var cfg ironbirdy.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		spew.Dump(err)
	}

	o := &robinhood.OAuth{
		Username: cfg.Account.Email,
		Password: cfg.Account.Password,
	}

	fmt.Printf("🤖 %s from %s\n", "Account loaded", configPath)

	c, err := robinhood.Dial(o)

	quote, _ := c.GetQuote(tickerSymbol)
	stockPrice := quote[0].LastTradePrice

	spew.Dump(quote)

	ch := ironbirdy.GetChains(c, tickerSymbol)

	for i, date := range ch[0].ExpirationDates {

		t, _ := time.Parse("2006-01-02", date)
		currentTime := time.Now()
		// add 24 for end of day calc
		daysUntilExpiration := (t.Sub(currentTime).Hours() + 24) / 24
		// fmt.Println(i, daysUntilExpiration, date)

		if float64(i) == daysUntilExpiration {

		}
	}

	fmt.Printf("🏹 %s, target is %s \n", "Chains loaded", expireDateNext)

	// date := "2021-02-05"

	dateInts := strings.Split(expireDateNext, "-")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	y, err := strconv.Atoi(dateInts[0])
	m, err := strconv.Atoi(dateInts[1])
	d, err := strconv.Atoi(dateInts[2])

	insts, err := ch[0].GetInstrument(ctx, "call", robinhood.NewDate(y, m, d))
	// spew.Dump(insts)
	if err != nil {
		spew.Dump(err)
	}

	fmt.Printf("☎️  %s with count: %d\n", "Calls loaded", len(insts))

	bearCallSpread := ironbirdy.BuildSpreads(c, insts, stockPrice)
	// spew.Dump(len(bearCallSpread))

	fmt.Printf("💪 %s with count: %d\n", "Calls spreads built", len(bearCallSpread))

	insts, err = ch[0].GetInstrument(ctx, "put", robinhood.NewDate(y, m, d))
	// spew.Dump(insts)
	if err != nil {
		spew.Dump(err)
	}

	fmt.Printf("🔰 %s with count: %d\n", "Puts loaded", len(insts))

	bullPutSpreads := ironbirdy.BuildSpreadsPuts(c, insts, stockPrice)
	// spew.Dump(len(bullPutSpreads))

	fmt.Printf("💪 %s with count: %d\n", "Put spreads built", len(bullPutSpreads))

	fmt.Printf("📚 %s\n", "Building history")

	hist := ironbirdy.BuildHistory(tickerSymbol, historyStart, historyEnd)
	// spew.Dump(len(hist))

	fmt.Printf("💪 %s with count: %d \n", "Building history", len(hist))

	fmt.Printf("👾 %s\n", "Adjusting returns by risk")

	condors := ironbirdy.AdjustRisk(bearCallSpread, bullPutSpreads, hist, minDistanceFromPrice)

	fmt.Printf("✅  %s\n", "Done")

	return condors

}

func main() {
	var configPath string
	var tickerSymbol string
	var expirationDateNext string
	var writeToCSV string
	var historyStart string
	var historyEnd string
	var minDistanceFromPrice float64

	var generateBearCallSpreadsCommand = &cobra.Command{
		Use:   "bearcall",
		Short: "Generate all bear call spreads from the current market prices",
		// Long: ``,
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// configPath := args[0]
			// tickerSymbol := args[1]
			// expireDateNext := args[2]

			fmt.Printf("🎬 %s\n", "Starting")

			fmt.Printf("🐓 %s: %s \n", "Ticker", tickerSymbol)

			f, err := os.Open(configPath)
			if err != nil {
				spew.Dump(err)
			}
			defer f.Close()

			var cfg ironbirdy.Config
			decoder := yaml.NewDecoder(f)
			err = decoder.Decode(&cfg)
			if err != nil {
				spew.Dump(err)
			}

			o := &robinhood.OAuth{
				Username: cfg.Account.Email,
				Password: cfg.Account.Password,
			}

			fmt.Printf("🤖 %s from %s\n", "Account loaded", configPath)

			c, err := robinhood.Dial(o)

			quote, _ := c.GetQuote(tickerSymbol)
			stockPrice := quote[0].LastTradePrice

			spew.Dump(quote)

			ch := ironbirdy.GetChains(c, tickerSymbol)

			for i, date := range ch[0].ExpirationDates {

				t, _ := time.Parse("2006-01-02", date)
				currentTime := time.Now()
				// add 24 for end of day calc
				daysUntilExpiration := (t.Sub(currentTime).Hours() + 24) / 24
				// fmt.Println(i, daysUntilExpiration, date)

				if float64(i) == daysUntilExpiration {

				}
			}

			fmt.Printf("🏹 %s, target is %s \n", "Chains loaded", expirationDateNext)

			// date := "2021-02-05"

			dateInts := strings.Split(expirationDateNext, "-")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			y, err := strconv.Atoi(dateInts[0])
			m, err := strconv.Atoi(dateInts[1])
			d, err := strconv.Atoi(dateInts[2])

			insts, err := ch[0].GetInstrument(ctx, "call", robinhood.NewDate(y, m, d))
			// spew.Dump(insts)
			if err != nil {
				spew.Dump(err)
			}

			fmt.Printf("☎️  %s with count: %d\n", "Calls loaded", len(insts))

			bearCallSpread := ironbirdy.BuildSpreads(c, insts, stockPrice)

			if writeToCSV != "" {
				fmt.Println("Write To CSV")
			} else {
				spew.Dump(bearCallSpread[0])
			}

			// spew.Dump(bearCallSpread[0])

		},
	}

	generateBearCallSpreadsCommand.Flags().StringVarP(&configPath, "config", "c", "/Users/drbh/.robintools/config.yml", "path to config file")
	generateBearCallSpreadsCommand.Flags().StringVarP(&tickerSymbol, "ticker", "t", "IWM", "underlying symbol")
	generateBearCallSpreadsCommand.Flags().StringVarP(&expirationDateNext, "expire", "x", "2021-02-05", "expiration date")
	generateBearCallSpreadsCommand.Flags().StringVarP(&writeToCSV, "output", "o", "", "write csv [filename]")
	generateBearCallSpreadsCommand.MarkFlagRequired("ticker")

	var generateBullPutSpreadsCommand = &cobra.Command{
		Use:   "bullput",
		Short: "Generate all bull put spreads from the current market prices",
		// 		Long: `print is for printing anything back to the screen.
		// For many years people have printed back to the screen.`,
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// configPath := args[0]
			// tickerSymbol := args[1]
			// expireDateNext := args[2]

			fmt.Printf("🎬 %s\n", "Starting")

			fmt.Printf("🐓 %s: %s \n", "Ticker", tickerSymbol)

			f, err := os.Open(configPath)
			if err != nil {
				spew.Dump(err)
			}
			defer f.Close()

			var cfg ironbirdy.Config
			decoder := yaml.NewDecoder(f)
			err = decoder.Decode(&cfg)
			if err != nil {
				spew.Dump(err)
			}

			o := &robinhood.OAuth{
				Username: cfg.Account.Email,
				Password: cfg.Account.Password,
			}

			fmt.Printf("🤖 %s from %s\n", "Account loaded", configPath)

			c, err := robinhood.Dial(o)

			quote, _ := c.GetQuote(tickerSymbol)
			stockPrice := quote[0].LastTradePrice

			spew.Dump(quote)

			ch := ironbirdy.GetChains(c, tickerSymbol)

			for i, date := range ch[0].ExpirationDates {

				t, _ := time.Parse("2006-01-02", date)
				currentTime := time.Now()
				// add 24 for end of day calc
				daysUntilExpiration := (t.Sub(currentTime).Hours() + 24) / 24
				// fmt.Println(i, daysUntilExpiration, date)

				if float64(i) == daysUntilExpiration {

				}
			}

			fmt.Printf("🏹 %s, target is %s \n", "Chains loaded", expirationDateNext)

			// date := "2021-02-05"

			dateInts := strings.Split(expirationDateNext, "-")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			y, err := strconv.Atoi(dateInts[0])
			m, err := strconv.Atoi(dateInts[1])
			d, err := strconv.Atoi(dateInts[2])

			insts, err := ch[0].GetInstrument(ctx, "put", robinhood.NewDate(y, m, d))
			// spew.Dump(insts)
			if err != nil {
				spew.Dump(err)
			}

			fmt.Printf("☎️  %s with count: %d\n", "Calls loaded", len(insts))

			bullPutSpread := ironbirdy.BuildSpreadsPuts(c, insts, stockPrice)

			if writeToCSV != "" {
				fmt.Println("Write To CSV", writeToCSV)
			} else {
				spew.Dump(bullPutSpread[0])
			}
		},
	}

	generateBullPutSpreadsCommand.Flags().StringVarP(&configPath, "config", "c", "/Users/drbh/.robintools/config.yml", "path to config file")
	generateBullPutSpreadsCommand.Flags().StringVarP(&tickerSymbol, "ticker", "t", "IWM", "underlying symbol")
	generateBullPutSpreadsCommand.Flags().StringVarP(&expirationDateNext, "expire", "x", "2021-02-05", "expiration date")
	generateBullPutSpreadsCommand.Flags().StringVarP(&writeToCSV, "output", "o", "", "write csv [filename]")
	generateBullPutSpreadsCommand.MarkFlagRequired("ticker")

	var generateHistoricalVolitilityCommand = &cobra.Command{
		Use:   "histvol [string to print]",
		Short: "Use historical daily pricing to find max postive and negative deltas intra week",
		// 		Long: `print is for printing anything back to the screen.
		// For many years people have printed back to the screen.`,
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// tickerSymbol := args[0]
			// historyStart := args[1]
			// historyEnd := args[2]

			hist := ironbirdy.BuildHistory(tickerSymbol, historyStart, historyEnd)
			// spew.Dump(len(hist))
			if writeToCSV != "" {
				fmt.Println("Write To CSV", writeToCSV)
			} else {
				spew.Dump(hist)
			}

		},
	}

	generateHistoricalVolitilityCommand.Flags().StringVarP(&tickerSymbol, "ticker", "t", "IWM", "underlying symbol")
	generateHistoricalVolitilityCommand.Flags().StringVarP(&historyStart, "start", "s", "1800-01-01", "start date")
	generateHistoricalVolitilityCommand.Flags().StringVarP(&historyEnd, "end", "e", "2021-02-04", "end date")
	generateHistoricalVolitilityCommand.Flags().StringVarP(&writeToCSV, "output", "o", "", "write csv [filename]")
	generateHistoricalVolitilityCommand.MarkFlagRequired("ticker")

	var generateCondorCommand = &cobra.Command{
		Use:   "condors",
		Short: "Generate iron condors",
		Long: `Generate Iron Condors above a minimum distance
and ones that have spread with postive returns and only have a strike difference of $1`,
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// tickerSymbol := args[0]
			// historyStart := args[1]
			// historyEnd := args[2]
			// configPath := args[3]
			// expireDateNext := args[4]
			// minDistanceFromPrice := makeFloat(args[5])

			condors := calculate(configPath, tickerSymbol, expirationDateNext, historyStart, historyEnd, minDistanceFromPrice)

			if writeToCSV != "" {
				// fmt.Println("Write To CSV", writeToCSV)
				ironbirdy.WriteCondorsToFile(condors, writeToCSV)
			} else {
				spew.Dump(condors[0])
			}

		},
	}

	generateCondorCommand.Flags().StringVarP(&tickerSymbol, "ticker", "t", "IWM", "underlying symbol")
	generateCondorCommand.Flags().StringVarP(&configPath, "config", "c", "/Users/drbh/.robintools/config.yml", "path to config file")
	generateCondorCommand.Flags().StringVarP(&expirationDateNext, "expire", "x", "2021-02-05", "expiration date")
	generateCondorCommand.Flags().StringVarP(&historyStart, "start", "s", "1800-01-01", "start date")
	generateCondorCommand.Flags().StringVarP(&historyEnd, "end", "e", "2021-02-04", "end date")
	generateCondorCommand.Flags().StringVarP(&writeToCSV, "output", "o", "", "write csv [filename]")
	generateCondorCommand.Flags().Float64VarP(&minDistanceFromPrice, "distance", "d", 0.05, "min distance from price")
	generateCondorCommand.MarkFlagRequired("ticker")

	var rootCmd = &cobra.Command{Use: "ironbirdy"}
	rootCmd.AddCommand(generateBearCallSpreadsCommand)
	rootCmd.AddCommand(generateBullPutSpreadsCommand)
	rootCmd.AddCommand(generateHistoricalVolitilityCommand)
	rootCmd.AddCommand(generateCondorCommand)
	rootCmd.Execute()

}

// this convets thing to floats
func makeFloat(val string) float64 {
	b2, _ := strconv.ParseFloat(val, 64)

	return b2
}
