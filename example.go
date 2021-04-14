package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/go-autorest/autorest/adal"
)

type Response struct {
	ComputerName string `json:"computerName"`
	OsName       string `json:"osName"`
	OsVersion    string `json:"osVersion"`
	VMAgent      struct {
		VMAgentVersion string `json:"vmAgentVersion"`
		Statuses       []struct {
			Code          string `json:"code"`
			Level         string `json:"level"`
			DisplayStatus string `json:"displayStatus"`
			Message       string `json:"message"`
			Time          string `json:"time"`
		} `json:"statuses"`
		ExtensionHandlers []struct {
			Type               string `json:"type"`
			TypeHandlerVersion string `json:"typeHandlerVersion"`
			Status             struct {
				Code          string `json:"code"`
				Level         string `json:"level"`
				DisplayStatus string `json:"displayStatus"`
				Message       string `json:"message"`
			} `json:"status"`
		} `json:"extensionHandlers"`
	} `json:"vmAgent"`
	Disks []struct {
		Name     string `json:"name"`
		Statuses []struct {
			Code          string `json:"code"`
			Level         string `json:"level"`
			DisplayStatus string `json:"displayStatus"`
			Time          string `json:"time"`
		} `json:"statuses"`
	} `json:"disks"`
	BootDiagnostics struct {
		ConsoleScreenshotBlobURI string `json:"consoleScreenshotBlobUri"`
		SerialConsoleLogBlobURI  string `json:"serialConsoleLogBlobUri"`
	} `json:"bootDiagnostics"`
	Extensions []struct {
		Name               string `json:"name"`
		Type               string `json:"type"`
		TypeHandlerVersion string `json:"typeHandlerVersion"`
		Statuses           []struct {
			Code          string `json:"code"`
			Level         string `json:"level"`
			DisplayStatus string `json:"displayStatus"`
			Message       string `json:"message"`
		} `json:"statuses"`
	} `json:"extensions"`
	HyperVGeneration string `json:"hyperVGeneration"`
	Statuses         []struct {
		Code          string `json:"code"`
		Level         string `json:"level"`
		DisplayStatus string `json:"displayStatus"`
		Message       string `json:"message"`
	} `json:"statuses"`
}

func main() {


	// deviceConfig := auth.NewDeviceFlowConfig(clientId, tenant_id)
	// authorizer, err := deviceConfig.Authorizer()
	//fmt.Println(authorizer)
	const activeDirectoryEndpoint = "https://login.microsoftonline.com/"
	tenantID := "d3bFFFFFFF80-cb1e-40f7-b59a-154105743342"
	oauthConfig, err := adal.NewOAuthConfig(activeDirectoryEndpoint, tenantID)
	applicationID := "UUUUU46c96-2d30-4f0f-bcee-6fcbbbdd289f"
	callback := func(token adal.Token) error {
		// This is called after the token is acquired
		fmt.Println("token acquired successfully")
		return nil
	}
	// The resource for which the token is acquired
	resource := "https://management.core.windows.net/"
	applicationSecret := "CCCCCdLE_W34dd-9~KK9tIV4Lli9CG5cFUHZS"
	spt, err := adal.NewServicePrincipalToken(*oauthConfig,
		applicationID,
		applicationSecret,
		resource,
		callback)
	if err != nil {
		fmt.Println("unable to acquire lock , err is ", err)
	}

	// Acquire a new access token
	mytoken := ""
	err = spt.Refresh()
	if err == nil {
		mytoken = spt.Token().AccessToken
		//fmt.Println(mytoken)
	}

	url := "https://management.azure.com/subscriptions/259b19e3-37d7-UUUUUU-VVVVV-5df86937a7f5/resourceGroups/spoke-network-demo/providers/Microsoft.Compute/virtualMachines/SACP-6433-VM01-00/instanceView?api-version=2020-12-01"

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6Im5PbzNaRHJPRFhFSzFqS1doWHNsSFJfS1hFZyIsImtpZCI6Im5PbzNaRHJPRFhFSzFqS1doWHNsSFJfS1hFZyJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldC8iLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC9kM2JjMjE4MC1jYjFlLTQwZjctYjU5YS0xNTQxMDU3NDMzNDIvIiwiaWF0IjoxNjEzMjA5NjM0LCJuYmYiOjE2MTMyMDk2MzQsImV4cCI6MTYxMzIxMzUzNCwiYWlvIjoiRTJaZ1lFZzNWZG1mNmhzcHYvdVBNYmZmOC9NQ0FBPT0iLCJhcHBpZCI6IjAxZDQ2Yzk2LTJkMzAtNGYwZi1iY2VlLTZmY2JiYmRkMjg5ZiIsImFwcGlkYWNyIjoiMSIsImdyb3VwcyI6WyI5MmM5NDIyNy01MWQ3LTQ4OGItOWY1Yy1lZjM2MGI3OTNhMDEiLCJhMjk5NmNiOS03ZjhkLTQ5Y2MtYmUzZi0wY2RmNWY0NWZkYWIiLCI5YjNmNzU3Yi0wYjVmLTQxYzktYWNlOC05MzMwMmQwODcyNjQiLCIyMzRlZDM0Ni1kNzg3LTRjNjQtOWNlZC00MTkzOGExYzliMTciLCJkZmE5NWU4ZS01MmMzLTQ3YzMtYjFlMS1iZDQzZTk2NzgxN2UiXSwiaWRwIjoiaHR0cHM6Ly9zdHMud2luZG93cy5uZXQvZDNiYzIxODAtY2IxZS00MGY3LWI1OWEtMTU0MTA1NzQzMzQyLyIsIm9pZCI6ImUyZjQwMTI5LWJmNTAtNDRhNi04YzE4LTQzOGFhM2VjNjMxZiIsInJoIjoiMC5BUjhBZ0NHODB4N0w5MEMxbWhWQkJYUXpRcFpzMUFFd0xROVB2TzV2eTd2ZEtKOGZBQUEuIiwic3ViIjoiZTJmNDAxMjktYmY1MC00NGE2LThjMTgtNDM4YWEzZWM2MzFmIiwidGlkIjoiZDNiYzIxODAtY2IxZS00MGY3LWI1OWEtMTU0MTA1NzQzMzQyIiwidXRpIjoiaE8wNVlHZHJxRTZBQm5yU01uZGZBQSIsInZlciI6IjEuMCIsInhtc190Y2R0IjoxNTI4NzMyMjgxfQ.g6jm3eigMCvKIJMUJCgGzrKjwuJqf7tjyvQIcYrbgV95EcxeinO67itJqbEoq3DCxD1rB2xbQ0XAQoJ7nupH9DfkK02NCquX9zOlQktQDjNQtGUjL8OMO62bnZ6Kvlu8stZlzF1jN1FU3-pMKbbc3Ff76SDqyYwjj8sJQVJPKZ7_6g0g6j21FPEpiB-d9vxIinNSvT265MIQAlwmNmMhD7-pOp5L2-XmPVm4uR49usVlvCRougRKD3Wu-pywXwVC7hGalBUUWTpDRutfAPO7qptk2ORtW7rSJtpGSTtCj5LtLcBEAEit40fq88rh_0GmnmWEvBSzOSTZjl3YTd26tw")
	var bearer_token = fmt.Sprintf("Bearer %s", mytoken)
	//fmt.Println(bearer_token)
	req.Header.Add("Authorization", bearer_token)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	var responseObject Response
	json.Unmarshal(body, &responseObject)
	//fmt.Println(responseObject)
	//fmt.Println(responseObject.Extensions)
	for _, s := range responseObject.Extensions {
		fmt.Println(s.Name, s.Statuses[0].Message)
	}

}
