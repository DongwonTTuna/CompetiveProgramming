package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func doHere() string {
	var answer string

	return answer
}

func main() {
	defer writer.Flush()
	Str := doHere()
	writer.WriteString(Str)
}

//---------
// Functions for competition Programming - DongwonTTuna
//---------

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readLineAndSplitSpaces() []string {
	bytes, _, _ := reader.ReadLine()
	return strings.Fields(string(bytes))
}

func parseStringsToNumbers(strs []string, isInteger bool) interface{} {
	if isInteger {
		var ints []int
		for _, str := range strs {
			n, err := strconv.Atoi(str)
			if err != nil {
				return nil
			}
			ints = append(ints, n)
		}
		return ints
	} else {
		var floats []float64
		for _, str := range strs {
			n, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return nil
			}
			floats = append(floats, n)
		}
		return floats
	}
}

func doubleArraytoLines(double [][]interface{}) string {
	var builder strings.Builder
	for _, v := range double {
		builder.WriteString(ArrayToOneLine(v))
		builder.WriteString("\n")
	}
	return builder.String()
}

func parseToInt(text string) (res int) {
	v, _ := strconv.Atoi(text)
	return v
}

func parseIntstoStrs(ints []int) []string {
	var strs []string
	for _, s := range ints {
		n := strconv.Itoa(s)
		strs = append(strs, n)
	}
	return strs
}

func ArraytoMultiLine(data interface{}) string {
	switch v := data.(type) {
	case []int:
		return MakeLineInts(v, "\n")
	case []float64:
		return MakeLineFloat64(v, "\n")
	case []string:
		return MakeLineString(v, "\n")
	case []bool:
		return MakeLineBool(v, "\n")
	default:
		return ""
	}
}

func ArrayToOneLine(data interface{}) string {
	switch v := data.(type) {
	case []int:
		return MakeLineInts(v, " ")
	case []float64:
		return MakeLineFloat64(v, " ")
	case []string:
		return MakeLineString(v, " ")
	case []bool:
		return MakeLineBool(v, " ")
	default:
		return ""
	}
}

func MakeLineInts(val []int, str string) string {
	var builder strings.Builder
	for _, v := range val {
		builder.WriteString(strconv.Itoa(v))
		builder.WriteString(str)
	}
	return strings.TrimRight(builder.String(), str)
}
func MakeLineFloat64(val []float64, str string) string {
	var builder strings.Builder
	for _, v := range val {
		builder.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		builder.WriteString(str)
	}
	return strings.TrimRight(builder.String(), str)
}
func MakeLineString(val []string, str string) string {
	var builder strings.Builder
	for _, v := range val {
		builder.WriteString(v)
		builder.WriteString(str)
	}
	return strings.TrimRight(builder.String(), str)
}
func MakeLineBool(val []bool, str string) string {
	var builder strings.Builder
	for _, v := range val {
		if v {
			builder.WriteString("Yes")
			builder.WriteString(str)
		} else {
			builder.WriteString("No")
			builder.WriteString(str)
		}
	}
	return strings.TrimRight(builder.String(), str)
}

func SortSlice(slice interface{}) {
	switch v := slice.(type) {
	case []int:
		sort.Ints(v)
	case []float64:
		sort.Float64s(v)
	case []string:
		sort.Strings(v)
	}
}

func SortSliceDescending(slice interface{}) {
	switch v := slice.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(v)))
	case []string:
		sort.Sort(sort.Reverse(sort.StringSlice(v)))
	}
}

func getMax(slice interface{}) interface{} {
	switch s := slice.(type) {
	case []int:
		max := getMaxInt(s)
		return max
	case []float64:
		max := getMaxFloat64(s)
		return max
	default:
		return nil
	}
}

func getMaxInt(slice []int) int {
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func getMaxFloat64(slice []float64) float64 {
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func getMin(slice interface{}) interface{} {
	switch s := slice.(type) {
	case []int:
		max := getMinInt(s)
		return max
	case []float64:
		max := getMinFloat64(s)
		return max
	default:
		return nil
	}
}

func getMinInt(slice []int) int {
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}

func getMinFloat64(slice []float64) float64 {
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}

func getAverage(slice interface{}) float64 {
	switch s := slice.(type) {
	case []int:
		average := getAverageInt(s)
		return average
	case []float64:
		average := getAverageFloat64(s)
		return average
	default:
		log.Fatal("slice is not []int or []float64")
		return -1
	}
}

func getAverageInt(slice []int) float64 {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return float64(sum) / float64(len(slice))
}

func getAverageFloat64(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum / float64(len(slice))
}

func Reverse(slice interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return
	}
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		tmp := s.Index(i).Interface()
		s.Index(i).Set(s.Index(j))
		s.Index(j).Set(reflect.ValueOf(tmp))
	}
}

func AlphabetToNumber(alpha rune) int {
	if unicode.IsLetter(alpha) {
		return int(alpha)
	} else {
		return -1
	}
}
func NumberToAlphabet(number int) string {
	if (number >= 65 && number <= 90) || (number >= 97 && number <= 122) {
		return string(rune(number))
	} else {
		return ""
	}
}
