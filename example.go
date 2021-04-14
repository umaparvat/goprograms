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
	tenantID := "d3bFFFFFFF80-cb1e-XXXXXXX"
	oauthConfig, err := adal.NewOAuthConfig(activeDirectoryEndpoint, tenantID)
	applicationID := "UUUUU46c96-2d30--YYYYYYYY"
	callback := func(token adal.Token) error {
		// This is called after the token is acquired
		fmt.Println("token acquired successfully")
		return nil
	}
	// The resource for which the token is acquired
	resource := "https://management.core.windows.net/"
	applicationSecret := "CCCCCdLE_W34dd-9~RRRRRTTTTTT"
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

	url := "https://management.azure.com/subscriptions/TTTTTTT-37d7-UUUUUU-VVVVV-TTTTTTT/resourceGroups/network-demo/providers/Microsoft.Compute/virtualMachines/SCP-6433-VM01-00/instanceView?api-version=2020-12-01"

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
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
