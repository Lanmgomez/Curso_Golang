package api

import (
	"fmt"
	"net/http"
)

func FetchApi() *http.Response {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Erro:", err)
		return nil
	}

	fmt.Println("Ok", res)
	return res
}
