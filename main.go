package main

import (
	"fmt"

	"github.com/kassybas/reeu/models/resource"
)

func main() {
	fmt.Println("hello world")

	// pr := make([]resource.Resource, 3)
	// pr[0] = resource.NewResource("subtest", nil, nil, 2.4)
	// pr[1] = resource.NewResource("subtestBEEE", nil, nil, 4.3)
	// pr[2] = resource.NewResource("subXXXXtB", nil, nil, 1.4)

	// mod := make([]resource.Modifier, 1.0)
	// mod[0] = resource.Modifier{Amount: 1.1, Name: "NationalTax"}

	// r := resource.NewResource("Test", pr, mod, 0)
	r := resource.LoadResource("data/sweeden/production/malmo.yaml")
	fmt.Println(r.CollectMonthly())
	//loop.StartLoop()
}
