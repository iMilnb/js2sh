package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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
	t := reflect.ValueOf(v)

	switch t.Kind() {
	case reflect.Map:
		for _, k := range t.MapKeys() {
			i := k.Interface().(string)
			varType(hasPrev(prev)+i, v.(map[string]interface{})[i])
		}
	case reflect.Slice:
		for i, v := range v.([]interface{}) {
			varType(hasPrev(prev)+strconv.Itoa(i), v)
		}
	case reflect.Float64:
		num := v.(float64)
		fmtstr := "%s=\"" + numFmt(num) + "\"\n"
		fmt.Printf(fmtstr, s.ToUpper(prev), num)
	case reflect.String:
		fmt.Printf("%s=\"%s\"\n", s.ToUpper(prev), v.(string))
	}
}

func main() {

	var data []byte
	var err error
	var f interface{}

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
