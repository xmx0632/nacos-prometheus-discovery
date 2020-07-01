package httputil

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Println("err:", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Println("err:", err)
	}
	return string(body), nil
}
