package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func salesforce(f force) error {
	u := f.URL + "/services/data/v20.0/sobjects/Contact/describe/"
	payload := strings.NewReader("")

	req, err := http.NewRequest("GET", u, payload)
	if err != nil {
		return err
	}
	req.Header.Add("authorization", "Bearer "+f.Token)
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	log.Println(f.Env.Env, res.Status)
	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		return fmt.Errorf(string(body))
	}
	return nil

}
