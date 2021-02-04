package ironbirdy

import (
	"fmt"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"

	"strconv"

	"github.com/drbh/go-robinhood"
)

// this convets thing to floats
func makeFloat(val string) float64 {
	b2, _ := strconv.ParseFloat(val, 64)

	return b2
}

// GetChains we want to get all of the option chains.
func GetChains(c *robinhood.Client, tickerSymbol string) []*robinhood.OptionChain {
	// fmt.Println(tickerSymbol, balance, stockPrice)
	i, err := c.GetInstrumentForSymbol(tickerSymbol)
	if err != nil {
		spew.Dump(err)
	}
	ch, err := c.GetOptionChains(i)
	if err != nil {
		spew.Dump(err)
	}
	return ch
}

var floatType = reflect.TypeOf(float64(0))

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

func WriteCondorsToFile(condors []IronCondor, name string) {
	file, fileErr := os.Create(name)
	if fileErr != nil {
		spew.Dump(fileErr)
	}
	names := structs.Names(condors[0])

	putNames := structs.Names(condors[0].PutLeg)

	for i := 0; i < len(names); i++ {

		currentName := names[i]

		if currentName == "PutLeg" {
			for ni, n := range putNames {
				fmt.Fprintf(file, "PutLeg.%s", n)

				if ni < (len(putNames) - 1) {
					fmt.Fprintf(file, ",")
				}

			}

		} else if currentName == "CallLeg" {
			for ni, n := range putNames {
				fmt.Fprintf(file, "CallLeg.%s", n)

				if ni < (len(putNames) - 1) {
					fmt.Fprintf(file, ",")
				}

			}
		} else {
			fmt.Fprintf(file, "%s", currentName)
		}

		if i < (len(names) - 1) {
			fmt.Fprintf(file, ",")
		}
	}
	fmt.Fprintf(file, "\n")
	for i := range condors {
		k := condors[i]
		v := reflect.ValueOf(k)
		for i := 0; i < v.NumField(); i++ {

			_, flerr := getFloat(v.Field(i).Interface())

			if flerr != nil {
				// fmt.Println(fflt)

				v1 := reflect.ValueOf(v.Field(i).Interface())
				for i := 0; i < v1.NumField(); i++ {
					fmt.Fprintf(file, "%.4f", v1.Field(i).Interface())
					if i < (v1.NumField() - 1) {
						fmt.Fprintf(file, ",")
					}
				}

			} else {
				fmt.Fprintf(file, "%.4f", v.Field(i).Interface())
			}

			if i < (v.NumField() - 1) {
				fmt.Fprintf(file, ",")
			}
		}
		fmt.Fprintf(file, "\n")
	}
}
