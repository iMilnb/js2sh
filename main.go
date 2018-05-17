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

func fileExists(f string) bool {
	if _, err := os.Stat(flag.Arg(0)); err == nil {
		return true
	}
	return false
}

func hasPrev(prev string) string {
	if prev != "" {
		prev += "_"
	}
	return prev
}

func numFmt(num float64) string {
	if num == float64(int(num)) {
		return "%.0f"
	}
	return "%f"
}

func varType(prev string, v interface{}) {

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
		fmt.Printf(fmtstr, s.ToUpper(prev), num)
	case string:
		fmt.Printf("%s=\"%s\"\n", s.ToUpper(prev), v.(string))
	}
}

func main() {

	var data []byte
	var err error
	var f interface{}

	flag.Usage = func() {
		fmt.Printf("Usage:\n\t%s /path/to/file.json\n", os.Args[0])
		fmt.Printf("\tcat /path/to/file.json | %s\n", os.Args[0])
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

	varType("", f)
}
