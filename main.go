package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)


func doHere() string{
	var answer string



	a := []int{1,5,7,2,4}
	ReverseArray(a)
	fmt.Println(a)

	return answer
}




func main(){
	defer writer.Flush()
	Str := doHere()
	writer.WriteString(Str)
}

//---------
// Fast IO
//---------

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readLineAndSplitSpaces() []string{
	bytes, _, _ := reader.ReadLine()
	return strings.Fields(string(bytes))
}

func convertStrsToInts(strs []string) []int{
	var integers []int
	for _, s := range strs {
		n, _ := strconv.Atoi(s)
		integers = append(integers, n)
	}
	return integers
}

func parseToInt(text string) (res int) {
	v, _ := strconv.Atoi(text)
	return v
}

func convertIntstoStrs(ints []int) []string{
	var strs []string
	for _, s := range ints{
		n := strconv.Itoa(s)
		strs = append(strs, n)
	}
	return strs
}

func ArraytoOneLine(slice interface{}) string{
	var str string
    val := reflect.ValueOf(slice)
    if val.Kind() == reflect.Slice {
        for i := 0; i < val.Len(); i++ {
            str += " " + val.Index(i).Interface().(string)
        }
    }
	return strings.TrimLeft(str, " ")
}

func ArraytoMultiLine(slice interface{}) string{
	var str string
    val := reflect.ValueOf(slice)
    if val.Kind() == reflect.Slice {
        for i := 0; i < val.Len(); i++ {
            str += "\n" + val.Index(i).Interface().(string)
        }
    }
	return strings.TrimLeft(str, "\n")

}

func parseToInt64(text string) (res int64) {
	v, _ := strconv.ParseInt(text, 10, 64)
	return v
}

func Sort(ints []int) []int{
	sort.Ints(ints)
	return ints
}

func ReverseSort(ints []int) []int{
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	return ints
}

func getMax(ints []int) int{
	var largest int
	for _, v := range ints {
    if v > largest {
        largest = v
    }
	}
	return largest
}
func getMin(ints []int) int{
	var smallest int
	for _, v := range ints {
    if v < smallest {
        smallest = v
    }
	}
	return smallest
}

func ReverseArray(slice interface{}) {
    switch v := slice.(type) {
    case []int:
        for i := 0; i < len(v)/2; i++ {
            j := len(v) - i - 1
            v[i], v[j] = v[j], v[i]
        }
    case []string:
        for i := 0; i < len(v)/2; i++ {
            j := len(v) - i - 1
            v[i], v[j] = v[j], v[i]
        }
    case []rune:
        for i, j := 0, len(v)-1; i < len(v)/2; i, j = i+1, j-1 {
            v[i], v[j] = v[j], v[i]
        }
    }
}
