package main

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
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
// Fast IO
//---------

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readLineAndSplitSpaces() []string {
	bytes, _, _ := reader.ReadLine()
	return strings.Fields(string(bytes))
}

func convertStrsToInts(strs []string) []int {
	var integers []int
	for _, s := range strs {
		n, _ := strconv.Atoi(s)
		integers = append(integers, n)
	}
	return integers
}

func doubleArraytoLines(double [][]interface{}) string {
	var str string
	for _, v := range double {
		str += ArrayToOneLine(v)
		str += "\n"
	}
	return str
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

func ArrayToOneLine(data interface{}) string {
	var result string
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			switch v.(type) {
			case int:
				result += strconv.Itoa(v.(int)) + " "
			case float64:
				result += strconv.FormatFloat(v.(float64), 'f', -1, 64) + " "
			case string:
				result += v.(string) + " "
			case bool:
				if v == true {
					result += "Yes "
				} else {
					result += "No "
				}
			}
		}
	}
	return strings.TrimRight(result, " ")
}

func ArraytoMultiLine(data interface{}) string {
	var result string
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			switch v.(type) {
			case int:
				result += strconv.Itoa(v.(int)) + "\n"
			case float64:
				result += strconv.FormatFloat(v.(float64), 'f', -1, 64) + "\n"
			case string:
				result += v.(string) + "\n"
			case bool:
				if v == true {
					result += "Yes\n"
				} else {
					result += "No\n"
				}
			}
		}
	}
	return strings.TrimRight(result, "\n")
}

func SortSlice(slice interface{}) {
	switch v := slice.(type) {
	case []int:
		sort.Ints(v)
	case []float64:
		sort.Float64s(v)
	}
}

func SortSliceDescending(slice interface{}) {
	switch v := slice.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(v)))
	}
}

func getMax(slice interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}
	max := s.Index(0).Interface()
	for i := 1; i < s.Len(); i++ {
		v := s.Index(i).Interface()
		switch v := v.(type) {
		case int:
			if v > max.(int) {
				max = v
			}
		case float64:
			if v > max.(float64) {
				max = v
			}
		}
	}
	return max
}

func getMin(slice interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}
	min := s.Index(0).Interface()
	for i := 1; i < s.Len(); i++ {
		v := s.Index(i).Interface()
		switch v := v.(type) {
		case int:
			if v < min.(int) {
				min = v
			}
		case float64:
			if v < min.(float64) {
				min = v
			}
		}
	}
	return min
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
