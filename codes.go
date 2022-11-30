package main

import (
	"errors"
	"fmt"
)

type code struct {
	Code  int16  `json:"code"`
	Title string `json:"title"`
	Description string `json:"description"`
}

var codes = []code{
	{Code: 100, Title: "Continue", Description: "This interim response indicates that the client should continue the request or ignore the response if the request is already finished."},
	{Code: 101, Title: "Switching Protocols", Description: "This code is sent in response to an Upgrade request header from the client and indicates the protocol the server is switching to."},
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

func getCode(c int16) (code, error) {
	for _, code := range codes {
		if code.Code == c {
			return code, nil
		}
	}
	message := fmt.Sprintf("No matching Response Status Code found for %d.", c)
	return code{Code: 0, Title: "", Description: ""}, errors.New(message)
}

func findCodeByTitle(input string) (code, error) {
	for _, code := range codes {
		if code.Title == input {
			return code, nil
		}
	}
	message := fmt.Sprintf("No Response Status Code found matching the title '%s'.", input)
	return code{Code: 0, Title: "", Description: ""}, errors.New(message)
}