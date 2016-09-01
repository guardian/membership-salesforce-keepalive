package main

import (
	"fmt"
	"log"
)

func main() {
	envs := []Env{
		Env{"DEV", "DEV"},
		Env{"PROD", "UAT"},
		Env{"PROD", "PROD"}}

	for _, e := range envs {
		l, err := getSalesforceLogin(e)
		if err != nil {
			fmt.Println("error fetching conf:")
			log.Panic(err)
		}
		err = salesforce(l)
		if err != nil {
			fmt.Println("error sending keepalive")
			log.Panic(err)
		}

	}

}
