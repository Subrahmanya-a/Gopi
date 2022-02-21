package test

import (
	"errors"
	"go-webservices-automation/utils"
	"testing"
	"github.com/dailymotion/allure-go"
)

func TestValhallaGet(t *testing.T) {
	/*
	This test verifies 'GET' request for File Server
	*/
	url, method := utils.GetValhallaURL("GET")
	statusCode, _ := utils.ValhallaGet(url, method)
	
	allure.Test(t, allure.Action(func() {
		allure.Step(allure.Description("Verifying Status code"), allure.Action(func() {
			 if statusCode != 200 {
				t.Errorf("Expected 200, but got %d ", statusCode)
			 	allure.Fail(errors.New("status code not as expected"))
			
			 }				
		 }))
	}))
}

