package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"go-webservices-automation/config"
)

func ValhallaManifest(url string, method string, ) (int, string) {
	/*
	This function hits file server MANIFEST request and return error/respone body
	
	:param url string : Request URL to hit e.g. https://35.235.73.188/sites/80ccc9fc-cd6a-447f-b69f-6a24bccd6ccf/environments/dev/files
	:param method string : Request method type e.g. MANIFEST
	:rtype int, string
	:return statusCode, response : Status Code and Response body from received response
	*/

	// Injects certificate for authorization
	client := config.CertsInjectionsPem()
	manifestPath := "?limit=3"
	req, _ := http.NewRequest(method, url+manifestPath, nil)
	// Send Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return resp.StatusCode, err.Error()
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return 1, err.Error()
	}
	fmt.Println(string(data))
	fmt.Println(resp.StatusCode)

	// Return statusCode and response body
	return resp.StatusCode, string(data)
	
}

func ValhallaGet(url string, method string) (int, string) {
	/*
	This function hits file server GET request and return error/respone body
	
	:param url string : Request URL to hit e.g. https://35.235.73.188/sites/80ccc9fc-cd6a-447f-b69f-6a24bccd6ccf/environments/dev/files
	:param method string : Request method type e.g. GET
	:rtype int, string
	:return statusCode, response : Status Code and Response body from received response
	*/

	// Injects certificate for authorization
	client := config.CertsInjectionsPem()
	get := "/t2"
	req, _ := http.NewRequest(method, url+get, nil)

	// Send Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return resp.StatusCode, err.Error()
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return 1, err.Error()
	}
	fmt.Println(string(data))
	fmt.Println(resp.StatusCode)

	// Return statusCode and response body
	return resp.StatusCode, string(data)
}