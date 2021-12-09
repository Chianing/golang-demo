package scene

import (
	"fmt"
	"golang-demo/model"
	"sort"
)

// TestMapPrintOrder
// https://www.cnblogs.com/noKing/p/11661567.html
func TestMapPrintOrder() {
	m := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}

	printByRange(m)
	printByRange(m)

	// fmt.Println("--------------")

	// printBySortedKey(m)
}

func TestChangeMapValueField() {
	m := make(map[string]model.TestModel)
	testModel := initTestModel()

	m["k1"] = testModel
	// m["k1"].Name = "name2"
	// testModel.Name = "name2"

	fmt.Println("m[k1]:", m["k1"])

	// changeMapValue(m)
	// fmt.Println("changeMapValue, m[k1]:", m["k1"])

	// fmt.Printf("TestChangeMapValueField %p\n", m)

	// changeMapToNil(m)
	// fmt.Println("changeMapToNil, m:", m)
}

func TestChangeMapPtrValueField() {
	m := make(map[string]*model.TestModel)
	testModel := initTestModel()

	m["k1"] = &testModel
	m["k1"].Name = "name2"

	fmt.Println("m[k1]:", *m["k1"])

}

func changeMapValue(m map[string]model.TestModel) {
	testModel := initTestModel()
	testModel.Name = "name2"
	m["k1"] = testModel

	fmt.Printf("changeMapValue %p\n", m)

}

func changeMapToNil(m map[string]model.TestModel) {
	changeMapValue(m)

	fmt.Printf("changeMapToNil %p\n", m)
	m = nil
}

func printBySortedKey(m map[string]string) {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("key:", k, "value:", m[k])
	}
}

func printByRange(m map[string]string) {
	for key := range m {
		fmt.Print(key, ",")
	}
	fmt.Println()
}

func initTestModel() model.TestModel {
	return model.TestModel{Name: "name1"}
}
