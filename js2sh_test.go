package main

import (
	"encoding/json"
	"testing"
)

const jsExample = `{"k1": "v1", "k2": [1, true, {"kink2": [null]}]}`

func mkMap() map[string]interface{} {
	testJSON := []byte(jsExample)

	testMap := make(map[string]interface{})

	if err := json.Unmarshal(testJSON, &testMap); err != nil {
		panic(err)
	}

	return testMap
}

func doTest(t *testing.T, msg string) {
	testMap := mkMap()

	if varType("", testMap) == false {
		t.Errorf("%v conversion failed (%s)", testMap, msg)
	}
}

func TestVarDefault(t *testing.T) {
	doTest(t, "default")
}

func TestVarNoup(t *testing.T) {
	*noup = true

	doTest(t, "noup")

	*noup = false
}

func TestVarSeparator(t *testing.T) {
	*separator = "/"

	doTest(t, "separator")

	*separator = "_"
}

func TestVarFilter(t *testing.T) {

	*filter = "K1"

	doTest(t, "filter")

	*filter = ""
}

func TestVarAll(t *testing.T) {
	*noup = true
	*separator = "/"
	*filter = "k2"

	doTest(t, "all flags")
}
