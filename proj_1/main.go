package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

type Hotel struct {
	eanHotelID    int
	name          string
	address1      string
	city          string
	stateProvince string
	postalCode    string
	latitude      float64
	longitude     float64
	starRating    float64
	highRate      float64
	lowRate       float64
}

func (e *Hotel) getFieldString(field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func (e *Hotel) getFieldInteger(field string) int {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func (e *Hotel) getFieldFloat(field string) float64 {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Float()
}

//=======================================================================

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func toFloat(s string) float64 {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return math.NaN()
	}
	return x
}
func toInt(s string) int {
	x, err := strconv.Atoi(s)
	checkError(err)
	return x
}

func getNumericColumn(table map[int]Hotel, name string) map[int]float64 {
	n := len(table)
	column := make(map[int]float64, n)
	for i, rec := range table {
		column[i] = rec.getFieldFloat(name)
	}
	return column
}

func main() {

	// 1. Open the file
	f, err := os.Open("new_york_hotels.csv")
	checkError(err)

	// 2. Initialize the reader
	reader := csv.NewReader(f)
	// 3. Read all the records
	records, _ := reader.ReadAll()
	// 4. Iterate through the records as you wish

	nRows := len(records) - 1
	Hotels := make(map[int]Hotel, nRows)
	for rownum, rec := range records {
		// rownum 0 is the header
		if rownum > 0 {
			Hotels[rownum] = Hotel{
				eanHotelID:    toInt(rec[0]),
				name:          rec[1],
				address1:      rec[2],
				city:          rec[3],
				stateProvince: rec[4],
				postalCode:    rec[5],
				latitude:      toFloat(rec[6]),
				longitude:     toFloat(rec[7]),
				starRating:    toFloat(rec[8]),
				highRate:      toFloat(rec[9]),
				lowRate:       toFloat(rec[10]),
			}
		}
	}
	fmt.Println(Hotels[10])
	t := Hotels[3]                        // returns a Hotel struct
	fmt.Println(t.getFieldString("name")) // fetches t.name
	// hRate is a column vector, well not exactly (slice)
	hRate := getNumericColumn(Hotels, "highRate")
	fmt.Println(len(hRate), hRate[3])

	x := 42
	y := float32(43.3)
	z := "hello"

	xt := reflect.TypeOf(x).Kind()
	yt := reflect.TypeOf(y).Kind()
	zt := reflect.TypeOf(z).Kind()

	fmt.Printf("%T: %s\n", xt, xt)
	fmt.Printf("%T: %s\n", yt, yt)
	fmt.Printf("%T: %s\n", zt, zt)

	fmt.Println(yt)

}
