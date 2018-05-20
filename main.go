package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	s "strings"
)

const (
	usageStr = `Usage:
	%s /path/to/file.json
	echo '{"key": "value"}' | %s
`
	noupUsage   = "don't upper-case variables names"
	filterUsage = "show only entries matching this filter"
	sepUsage    = "key separator"
)

var noup *bool = flag.Bool("n", false, noupUsage)
var filter *string = flag.String("f", "", filterUsage)
var separator *string = flag.String("s", "_", sepUsage)

func fileExists(f string) bool {
	if _, err := os.Stat(flag.Arg(0)); err == nil {
		return true
	}
	return false
}

func hasPrev(prev string) string {
	if prev != "" {
		prev += *separator
	}
	return prev
}

func numFmt(num float64) string {
	if num == float64(int(num)) {
		return "%.0f"
	}
	return "%f"
}

func doUp(str string) string {
	if *noup == false {
		return s.ToUpper(str)
	}
	return str
}

func filterOut(str string) {
	if *filter == "" || s.Contains(str, *filter) {
		fmt.Printf(str)
	}
}

func varType(prev string, v interface{}) bool {

	switch v.(type) {
	case map[string]interface{}:
		for key, val := range v.(map[string]interface{}) {
			varType(hasPrev(prev)+key, val)
		}
	case []interface{}:
		for i, v := range v.([]interface{}) {
			varType(hasPrev(prev)+strconv.Itoa(i), v)
		}
	case float64:
		num := v.(float64)
		fmtstr := "%s=\"" + numFmt(num) + "\"\n"
		filterOut(fmt.Sprintf(fmtstr, doUp(prev), num))
	case string:
		filterOut(fmt.Sprintf("%s=\"%s\"\n", doUp(prev), v.(string)))
	case bool:
		filterOut(fmt.Sprintf("%s=\"%t\"\n", doUp(prev), v.(bool)))
	case nil:
		filterOut(fmt.Sprintf("%s=\"null\"\n", doUp(prev)))
	default:
		return false
	}

	return true
}

func main() {

	var data []byte
	var err error
	var f interface{}

	flag.Usage = func() {
		fmt.Printf(usageStr, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 || !fileExists(flag.Arg(0)) {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(flag.Arg(0))
	}

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &f); err != nil {
		panic(err)
	}

	if varType("", f) != true {
		os.Exit(1)
	}
}
