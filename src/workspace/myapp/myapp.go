package main

import (
	"fmt"

	"example.com/iftest"
	"example.com/maptest"
	"example.com/mypkg"
	"example.com/vartest"
)

func main() {
	a,b,c :=vartest.Vartest()
	mapresponse :=maptest.Maptest()
	text,ok := maptest.Maptest()["z"]
	for key, value := range maptest.Maptest() {
		fmt.Printf("%s=%d\n", key, value)
	}
    fmt.Println(mypkg.Hello())
	fmt.Println(a,b,c)
	fmt.Println(len(mapresponse))
	fmt.Println(len(mapresponse))
	fmt.Println(text)
	fmt.Println(ok)
	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not exist")
	}
	fmt.Println(iftest.Iftest(100,10))
}