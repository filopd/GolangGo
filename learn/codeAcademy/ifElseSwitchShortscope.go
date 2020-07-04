package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var v1Logical bool = true
	if v1Logical {
		fmt.Println("v1Logical is True.")
	}
	var v2Logical bool = false
	if !v2Logical {
		fmt.Println("v2Logical is False.")
	}
	if !v1Logical {
		fmt.Println("v1Logical is False.")
	} else if v2Logical {
		fmt.Println("v2Logical is True")
	} else {
		fmt.Println("v1Logical True & v2Logical False")
	}
	fmt.Println("#########################")
	v1String := "Hello"
	switch v1String {
	case "Hi":
		fmt.Println("Case 1 is Hi.")
	case "Hello":
		fmt.Println("Case 2 is Hello.")
	default:
		fmt.Println("Case 3 is default.")
	}
	fmt.Println("#########################")
	if v3Logical := true; v3Logical {
		fmt.Println("This is short scoped", v3Logical)
	}
	fmt.Println("#########################")
	switch v1Case := "NoCase"; v1Case {
	case "case1":
		fmt.Println("This is case 1.")
	default:
		fmt.Println("This is default.")
	}
	fmt.Println("#########################")
	fmt.Println("Random 1 number bet 0 to 100 is always:", rand.Intn(100))
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random 2 number bet 0 to 100 by Seed is:", rand.Intn(100))

}
