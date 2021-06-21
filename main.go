package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	retry "github.com/avast/retry-go"
)

const (
	url = "https://www.yahoo.com/"
)

func mian() {
	option := retry.LastErrorOnly(true)
	option
	retry.RandomDelay(3, nil, config)
}

func main1() {
	var body []byte
	err := retry.Do(func() error {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return nil
	})
	if err == nil {
		fmt.Printf(string(body))
	}

}
