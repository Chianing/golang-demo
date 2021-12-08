package scene

import (
	"fmt"
	"sort"
)

// TestMapPrintOrder
// https://www.cnblogs.com/noKing/p/11661567.html
func TestMapPrintOrder() {
	m := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}

	printByRange(m)
	printByRange(m)

	fmt.Println("--------------")

	printBySortedKey(m)
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
