package test

import (
//	"go-webservices-automation/config"
	"go-webservices-automation/utils"
	//"os"
	"testing"
	"github.com/dailymotion/allure-go"
	"github.com/stretchr/testify/assert"
	
)



func TestValhallaManifest(t *testing.T) {
	/*
	This test verifies 'MANIFEST' request for File Server
	*/
	allure.Test(t, allure.Action(func() {
		url, method := utils.GetValhallaURL("MANIFEST")
		statusCode, response := utils.ValhallaManifest(url, method)
		res := utils.ResponseJson(response)
		allure.Step(allure.Description("Verifying Status code and Response code"),
							allure.Action(func() {
								assert.Equal(t, 200, statusCode)
								assert.Equal(t, float64(3), res["total"], "Number of files returned are incorrect")
							}))
	}))
}
