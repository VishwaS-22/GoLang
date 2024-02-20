package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=abcdefghijkl"

func main() {
	fmt.Println("Handling URL's in GO!")
	fmt.Println(myurl)

	res, _ := url.Parse(myurl)

	fmt.Println(res.Scheme)

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	newURL := partsofURL.String()
	fmt.Println(newURL)

}
