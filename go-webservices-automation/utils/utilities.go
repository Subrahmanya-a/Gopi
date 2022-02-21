package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func GetValhallaURL(typeApi string) (string, string) {
	/*
	This function gets URL for respective valhalla API request
	:param typeApi string : API request type e.g. MANIFEST, GET
	:rtype string, string
	:return url, method : Request URL to hit, method type like MANIFEST, GET
	*/

	// Get URL from url.json
	jsonFile, _ := os.Open("../data/url.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	valhalla := result["valhalla"].(map[string]interface{})
	apiType := valhalla[typeApi].(map[string]interface{})
	url := apiType["url"].(string)
	method := apiType["method"].(string)
	jsonFile.Close()
	return url, method
}

func GetTestData(payload string) map[string]interface{} {
	/*
	This function gets test data from testData.json for respective API request
	:param reqName string : Request name to fetch test data e.g. setAutopilotPlan
	:rtype map[string]interface{}
	:return test data: Test Data for specified reqName
	*/
	// Get Test Data from testData.json
	jsonFile, _ := os.Open("../data/testData.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result[payload].(map[string]interface{})
}

func GetPayload(payload string) *bytes.Buffer {
	/*
	This function fetches required payload for respective request
	:param reqName string : Request name to fetch payload e.g. setAutopilotPlan
	:rtype *bytes.Buffer
	:return reqbody : Request Payload
	*/

	// Get request payload from payload.json
	jsonFile, _ := os.Open("../data/payload.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	graphql := result["graphql"].(map[string]interface{})
	reqPayload := graphql[payload].(map[string]interface{})
	query := reqPayload["query"].(string)
	if reqPayload["isDynamicAttr"] == true {
		testData := GetTestData(payload)
		for key, value := range testData {
			query = strings.Replace(query, key, value.(string), 1)
		}
	}
	reqbody := bytes.NewBufferString(`{
		"operationName": null,
		"variables": {},
		"query":"` + query + `"
	}`)
	return reqbody
}

// func SetRequestHeaders(req *http.Request) *http.Request {
// 	/*
// 	This function sets request headers
// 	:param req *http.Request : HTTP Request to set headers
// 	:rtype *http.Request
// 	:returns req : HTTPS Request after setting required Headers
// 	*/

// 	//Get required headers from requestHeaders.json
// 	jsonFile, _ := os.Open("../data/requestHeaders.json")
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	var result map[string]interface{}
// 	json.Unmarshal([]byte(byteValue), &result)

// 	headers := result["Headers"].(map[string]interface{})
// 	for key, value := range headers {
// 		newValue := value.(string)
// 		req.Header.Add(key, newValue)
// 	}
// 	req.Header.Add("Cookie", *config.COOKIE)
// 	return req
// }

func ResponseJson(response string) map[string]interface{} {
	/*
	This function returns response in Json format
	:param response string : Response in string
	:rype map[string]interface{}
	:return Response Json
	*/
	var res map[string]interface{}
	json.Unmarshal([]byte(response), &res)
	return res
}