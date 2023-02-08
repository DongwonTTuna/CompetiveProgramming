package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func doHere() string {

	return ""
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
		ints := make([]int, len(strs))
		for i, v := range strs {
			ints[i], _ = strconv.Atoi(v)
		}
		return ints
	} else {
		floats := make([]float64, len(strs))
		for i, v := range strs {
			floats[i], _ = strconv.ParseFloat(v, 64)
		}
		return floats
	}
}

func doubleArraytoLines(double interface{}) string {
	var builder strings.Builder
	val := reflect.ValueOf(double)

	for i := 0; i < val.Len(); i++ {
		builder.WriteString(ArrayToOneLineNoWhiteSpace(val.Index(i).Interface()))
		builder.WriteString("\n")
	}
	return builder.String()
}

func parseToInt(text string) (res int) {
	v, _ := strconv.Atoi(text)
	return v
}

func parseNumberstoStrs(slice interface{}) []string {
	a := reflect.ValueOf(slice)
	res := make([]string, a.Len())
	switch v := slice.(type) {
	case []int:
		for i, v := range v {
			res[i] = strconv.Itoa(v)
		}
	case []float64:
		for i, f := range v {
			res[i] = strconv.FormatFloat(f, 'f', -1, 64)
		}
	default:
		log.Fatal("parseNumberstoStrs : Type is not supported")
	}
	return res
}

func ArraytoMultiLine(data interface{}) string {
	return MakeLine(data, "\n")
}

func ArrayToOneLine(data interface{}) string {
	return MakeLine(data, " ")
}

func ArrayToOneLineNoWhiteSpace(data interface{}) string {
	return MakeLine(data, "")
}

func MakeLine(val interface{}, separator string) string {
	var builder strings.Builder
	switch v := val.(type) {
	case []int:
		for _, i := range v {
			builder.WriteString(strconv.Itoa(i))
			builder.WriteString(separator)
		}
	case []float64:
		for _, f := range v {
			builder.WriteString(strconv.FormatFloat(f, 'f', -1, 64))
			builder.WriteString(separator)
		}
	case []string:
		for _, s := range v {
			builder.WriteString(s)
			builder.WriteString(separator)
		}
	case []rune:
		for _, r := range v {
			builder.WriteRune(r)
			builder.WriteString(separator)
		}
	case []bool:
		for _, b := range v {
			if b {
				builder.WriteString("Yes")
			} else {
				builder.WriteString("No")
			}
			builder.WriteString(separator)
		}
	default:
		log.Fatal("MakeLine : Type is not supported")
	}
	return strings.TrimRight(builder.String(), separator)
}

func SortSlice(slice interface{}) {
	switch slice.(type) {
	case []int:
		sort.Ints(slice.([]int))
	case []float64:
		sort.Float64s(slice.([]float64))
	case []string:
		sort.Strings(slice.([]string))
	case []rune:
		sort.Slice(slice.([]rune), func(i, j int) bool {
			return slice.([]rune)[i] < slice.([]rune)[j]
		})
	}
}

func SortSliceDescending(slice interface{}) {
	switch slice.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(slice.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(slice.([]float64))))
	case []string:
		sort.Sort(sort.Reverse(sort.StringSlice(slice.([]string))))
	case []rune:
		sort.Slice(slice.([]rune), func(i, j int) bool {
			return slice.([]rune)[i] > slice.([]rune)[j]
		})
	}
}

func getMax(slice interface{}) interface{} {
	switch s := slice.(type) {
	case []int:
		sort.Ints(s)
		return s[len(s)-1]
	case []float64:
		sort.Float64s(s)
		return s[len(s)-1]
	default:
		return nil
	}
}

func getMin(slice interface{}) interface{} {
	switch s := slice.(type) {
	case []int:
		sort.Ints(s)
		return s[0]
	case []float64:
		sort.Float64s(s)
		return s[0]
	default:
		return nil
	}
}

func getAverage(slice interface{}) float64 {
	switch s := slice.(type) {
	case []int:
		var sum int
		for i := 0; i < len(s); i++ {
			sum += s[i]
		}
		return float64(sum) / float64(len(s))
	case []float64:
		var sum float64
		for i := 0; i < len(s); i++ {
			sum += s[i]
		}
		return sum / float64(len(s))
	default:
		log.Fatal("slice is not []int or []float64")
		return -1
	}
}

func reverseArray(arr interface{}) {
	val := reflect.ValueOf(arr)

	len := val.Len()

	for i := 0; i < len/2; i++ {
		opp := len - i - 1

		left := val.Index(i).Interface()
		right := val.Index(opp).Interface()

		val.Index(i).Set(reflect.ValueOf(right))
		val.Index(opp).Set(reflect.ValueOf(left))
	}
}

func removeIndex(arr interface{}, index int) interface{} {
	val := reflect.ValueOf(arr)

	len := val.Len()

	if index < 0 || index >= len {
		log.Fatal("Index out of range.")
	}

	newSlice := reflect.MakeSlice(val.Type(), 0, len-1)
	newSlice = reflect.AppendSlice(newSlice, val.Slice(0, index))
	newSlice = reflect.AppendSlice(newSlice, val.Slice(index+1, len))

	return newSlice.Interface()
}
