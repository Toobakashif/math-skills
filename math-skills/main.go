package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var GlobalVariance float64

func main() {
	var sumfile = []int{}
	file, _ := os.Open("data.txt")
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	var filetextscanner []string
	for filescanner.Scan() {
		filetextscanner = append(filetextscanner, filescanner.Text())
	}

	file.Close()

	length := 0
	for _, eachline := range filetextscanner {

		length = length + 1
		j, _ := strconv.Atoi(eachline)

		sumfile = append(sumfile, j)
	}

	a := Average(sumfile, length)
	fmt.Println("Average:", a)

	m := Median(sumfile, length)
	fmt.Println("Median:", m)

	s := StandardDivison(sumfile, length)
	fmt.Println("Standard Deviation", s)
	
	v := Variance(sumfile, length)
	fmt.Println("Variance", v)

}

func Average(suml []int, l int) int {
	sum := 0
	avaerage := 0
	for i := 0; i < len(suml); i++ {
		sum = sum + suml[i]
		avaerage = sum / l
	}
	return avaerage
}
func Median(s []int, length int) int {
	sort.Ints(s)
	middle := len(s) / 2
	medianValue := 0 //declare variable
	if length%2 != 0 {

		medianValue = s[middle]
	} else {

		medianValue = (s[middle-1] + s[middle]) / 2
	}
	return medianValue
}
func Variance(a []int, length int) int {
	sort.Ints(a)
	sum := 0
	div := 0.0
	sub := 0.0
	sumsqr := 0.0
	sqr := 0.0
	//sublen := 0
	for i := 0; i < length; i++ {
		sum += a[i]
		div = float64(sum / length)
	}
	for j := 0; j < length; j++ {
		sub = float64(a[j]) - div
		sqr = sub * sub
		sumsqr += sqr
	}
	//	sublen = length - 1
	variance := sumsqr / float64(length) // sublen
	GlobalVariance = variance
	return int(math.Round(variance))
}

func StandardDivison(s []int, length int) int {
	sum := 0.0
	qr := 0.0
	sqr := 0.0
	sumsqr := 0.0
	div := 0.0
	sqrl := 0
	for i := 0; i < length; i++ {
		sum += float64(s[i])
		qr = sum * sum
	}
	for j := 0; j < length; j++ {

		sqr = float64(s[j] * s[j])

		sumsqr += sqr
	}
	sqrl = length * length
	div = float64(length) * sumsqr
	div1 := div - qr
	div2 := div1 / float64(sqrl)

	var x float64 = float64(div2)
	result := math.Sqrt(x)

	result2 := math.Round(result)
	result3 := int(result2)
	//var y int = int(result)
	return result3
}
