package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("App web request")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Res type is %T\n", res)
	defer res.Body.Close()

	db, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	c := string(db)
	fmt.Println(c)
}
