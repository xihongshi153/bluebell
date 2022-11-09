package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

var APIURL = "http://127.0.0.1:8000/register"

func TestRegister(t *testing.T) {
	response, err := http.PostForm(APIURL, url.Values{
		"username": {"123q"},
		"password": {"456q"},
		"email":    {"789@qq.com"}})

	//okay, moving on...
	if err != nil {
		//handle postform error
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		//handle read response error
	}

	fmt.Printf("%s\n", string(body))
}
