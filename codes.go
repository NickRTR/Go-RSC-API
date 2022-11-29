package main

import (
	"errors"
	"fmt"
)

type code struct {
	Code  int16  `json:"code"`
	Title string `json:"title"`
}

var codes = []code{
	{Code: 100, Title: "Continue"},
	{Code: 101, Title: "Switching Protocols"},
}

func getCodes() []code {
	return codes
}

func getJustCodes() []int16 {
	var cs []int16
	for _, cod := range codes {
		cs = append(cs, cod.Code)
	}
	return cs
}

func getTitle(c int16) (string, error) {
	for _, code := range codes {
		if code.Code == c {
			return code.Title, nil
		}
	}
	fmt.Printf("No matching Response Status Code found for %d.", c)
	return "", errors.New("No matching Response Status Code found.")
}