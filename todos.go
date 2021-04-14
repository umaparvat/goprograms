package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyJsonResponse struct {
	Networks struct {
		LandingZone struct {
			Cidrs            []string `json:"cidrs"`
			ProviderSpecific struct {
				Dnsservers    []string `json:"dnsservers"`
				ResourceGroup string   `json:"resource_group"`
				Subscription  string   `json:"subscription"`
			} `json:"provider_specific"`
			Tags struct {
				Environment  string `json:"Environment"`
				Network      string `json:"Network"`
				Type         string `json:"Type"`
				Subscription string `json:"subscription"`
				Test         string `json:"test"`
			} `json:"tags"`
		} `json:"landing_zone"`
	} `json:"networks"`
	SecurityGroups struct {
		Rules struct {
			One0_48_144_136_29 []struct {
				Cidrs    []string `json:"cidrs"`
				Name     string   `json:"name"`
				Ports    []string `json:"ports"`
				Protocol string   `json:"protocol"`
			} `json:"10.48.144.136/29"`
			One0_48_144_144_29 []struct {
				Cidrs    []string `json:"cidrs"`
				Name     string   `json:"name"`
				Ports    []string `json:"ports"`
				Protocol string   `json:"protocol"`
			} `json:"10.48.144.144/29"`
			One0_48_144_160_27 []struct {
				Cidrs    []string `json:"cidrs"`
				Name     string   `json:"name"`
				Ports    []string `json:"ports"`
				Protocol string   `json:"protocol"`
			} `json:"10.48.144.160/27"`
		} `json:"rules"`
		Tags struct {
			Environment  string `json:"Environment"`
			Network      string `json:"Network"`
			Type         string `json:"Type"`
			Subscription string `json:"subscription"`
			Test         string `json:"test"`
		} `json:"tags"`
	} `json:"securityGroups"`
	Subnets struct {
		One0_48_144_136_29 struct {
			Cidr             string `json:"cidr"`
			Name             string `json:"name"`
			ProviderSpecific struct {
				LinkEndpointPolicies bool `json:"link_endpoint_policies"`
			} `json:"provider_specific"`
			Tags     struct{} `json:"tags"`
			ZoneName string   `json:"zoneName"`
			ZoneType string   `json:"zoneType"`
		} `json:"10.48.144.136/29"`
		One0_48_144_144_29 struct {
			Cidr             string `json:"cidr"`
			Name             string `json:"name"`
			ProviderSpecific struct {
				LinkEndpointPolicies bool `json:"link_endpoint_policies"`
			} `json:"provider_specific"`
			Tags     struct{} `json:"tags"`
			ZoneName string   `json:"zoneName"`
			ZoneType string   `json:"zoneType"`
		} `json:"10.48.144.144/29"`
		One0_48_144_160_27 struct {
			Cidr             string `json:"cidr"`
			Name             string `json:"name"`
			ProviderSpecific struct {
				LinkEndpointPolicies bool `json:"link_endpoint_policies"`
			} `json:"provider_specific"`
			Tags     struct{} `json:"tags"`
			ZoneName string   `json:"zoneName"`
			ZoneType string   `json:"zoneType"`
		} `json:"10.48.144.160/27"`
	} `json:"subnets"`
}

func getdatas(todos string) (MyJsonResponse, error) {
	var responseObject MyJsonResponse
	url := "https://jsonplaceholder.typicode.com/" + todos + "/1"
	// how to construct the above url when todos is variable
	furl := fmt.Sprintf("https://jsonplaceholder.typicode.com/%s/1", todos) //will this work?
	fmt.Println(url, furl)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return responseObject, err
	}
	req.Header.Add("Cookie", "something")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return responseObject, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return responseObject, err
	}
	fmt.Println(string(body))

	json.Unmarshal(body, &responseObject)
	return responseObject, err
}

func main() {
	res, _ := getdatas("apple")
	fmt.Println(res)
}
