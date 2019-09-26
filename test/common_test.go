package test

import (
	"testing"

	. "github.com/maretol/twister/logic"
)

func TestGetKeysFromMap(t *testing.T) {

	testMap := make(map[string]interface{})
	testMap["hoge"] = 0
	testMap["fuga"] = 1
	result2 := GetKeysFromInterfaceMap(testMap)
	if len(result2) != 2 {
		t.Fatalf("Failed result length is not correct\n")
	}
	t.Log("Test Successed")
}
