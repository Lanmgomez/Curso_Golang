package main

import (
	"fmt"
	"std/app/api"
)

func main() {

	res := api.FetchApi()

	fmt.Println("Ok:", res)
}
