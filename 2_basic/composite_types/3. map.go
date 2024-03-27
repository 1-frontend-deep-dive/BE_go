package goCompositeTypes

import "fmt"

/*
	Maps are used to store data values in key:value pairs
	Maps need to be initialized before they can be used, there are 2 ways to create
	1. make(map[string]string)
	2. map[string]string{}
*/

type ShopInfo struct {
	Name string
	Address string
}

func PlayWithMap() {
	// trick: make(map[string]interface{}) to story any types for the value
	commonMenu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          2.99,
		"coffee": 				3.55,
	}

	brandA := make(map[string]ShopInfo)
	shop1 := ShopInfo{
		Name: "Shop 1",
		Address: "243 Huston",
	}
	shop2 := ShopInfo{
		Name: "Shop 2",
		Address: "557 Beverly Hills",
	}

	brandA["shop1"] = shop1
	brandA["shop2"] = shop2

	fmt.Printf("The %s has soup with price %0.2f, located at %s\n", brandA["shop1"].Name, commonMenu["soup"], brandA["shop1"].Address)
	fmt.Printf("The %s has salad with price %0.2f, located at %s\n", brandA["shop2"].Name, commonMenu["salad"], brandA["shop2"].Address)
}