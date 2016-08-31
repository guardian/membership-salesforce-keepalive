package main

import "log"

func main() {
	envs := []Env{
		//Env{"DEV", "DEV"},
		Env{"PROD", "UAT"}, Env{"PROD", "PROD"}}

	for _, e := range envs {
		l, err := getSalesforceLogin(e)
		if err != nil {
			log.Panic(err)
		}
		err = salesforce(l)
		if err != nil {
			log.Panic(err)
		}

	}

}
