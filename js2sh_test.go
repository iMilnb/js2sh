package main

import (
	"encoding/json"
	"testing"
)

func mkMap() map[string]interface{} {
	testJSON := []byte(`{"key1": "val1", "key2": [1, 2, {"kink2": [3, 4]}]}`)

	testMap := make(map[string]interface{})

	if err := json.Unmarshal(testJSON, &testMap); err != nil {
		panic(err)
	}

	return testMap
}

func TestVarDefault(t *testing.T) {
	testMap := mkMap()

	if varType("", testMap) == false {
		t.Errorf("%v conversion failed", testMap)
	}
}

func TestVarNoup(t *testing.T) {
	testMap := mkMap()

	*noup = true

	if varType("", testMap) == false {
		t.Errorf("%v conversion failed with noup", testMap)
	}

	*noup = false
}

func TestVarSeparator(t *testing.T) {
	testMap := mkMap()

	*separator = "/"

	if varType("", testMap) == false {
		t.Errorf("%v conversion failed with separator", testMap)
	}

	*separator = "_"
}

func TestVarFilter(t *testing.T) {
	testMap := mkMap()

	*filter = "KEY"

	if varType("", testMap) == false {
		t.Errorf("%v conversion failed with filter", testMap)
	}

	*filter = ""
}

func TestVarAll(t *testing.T) {
	testMap := mkMap()

	*noup = true
	*separator = "/"
	*filter = "key"

	if varType("", testMap) == false {
		t.Errorf("%v conversion failed with all flags", testMap)
	}
}
