package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	// var (
	// 	authorizer autorest.Authorizer
	// )
	var subscriptionID = "XXXXXXXX"
	var appId = "XXXXcXXX89f"
	var tenant = "XXXX-YYYYYYY"
	var appSecret = "XXXLE_W34dd-9~RRRRRR"
	// dfc := auth.NewDeviceFlowConfig(appId, tenant)
	// spToken, err := dfc.ServicePrincipalToken()
	// authorizer = autorest.NewBearerAuthorizer(spToken)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var rg = "spC-network-demo"
	vmClient, insClient := AzureAuth(subscriptionID, appId, appSecret, tenant)
	getVM(vmClient, rg, insClient)
}
func getVM(vmClient compute.VirtualMachinesClient, rg string, insClient compute.VirtualMachineExtensionsClient) {
	for vm, err := vmClient.ListComplete(context.Background(), rg); vm.NotDone(); err = vm.Next() {
		if err != nil {
			log.Print("got error while traverising RG list: ", err)
		}

		i := vm.Value()
		//instanceView := i.InstanceView
		//fmt.Println(*i.InstanceView)
		fmt.Printf("\n trying is %s,%s,%s\n", rg, *i.Name, *i.ID)

	}
	ins, err := vmClient.Get(context.Background(), rg, "SCP-6435678-VM01-00", "")
	//body, err := ioutil.ReadAll(ins.Response.Body)
	//ins.Request.Response.Body.Close().Error()
	//ins.Body.Close().Error()
	fmt.Println(*ins.Name)
	if err != nil {
		log.Print("got error while gathering the extension ", err)
	}

	// to ge the instance view
	ins_view, err := vmClient.InstanceView(context.Background(), rg, "SCP-6438883-VM01-00")
	if err != nil {
		log.Print("got error while gathering instance view ", err)
	}
	for _, ext := range *ins_view.Extensions {
		fmt.Println(*ext.Name, *ext.Type)
		// Statuses is a list of memory objects.
		for _, each_stat := range *ext.Statuses {
			fmt.Println(*each_stat.Code, *&each_stat.Level, *each_stat.DisplayStatus, *each_stat.Message)
		}
	}
}

func AzureAuth(subscriptionID string, appId string, appSecret string, tenant string) (compute.VirtualMachinesClient, compute.VirtualMachineExtensionsClient) {
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	ins_view := compute.NewVirtualMachineExtensionsClient(subscriptionID)
	//vmClient.Authorizer = authorizer
	clientAuthorizer := auth.NewClientCredentialsConfig(appId, appSecret, tenant)
	authorizer, err := clientAuthorizer.Authorizer()
	//authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Auth: Successful")
		vmClient.Authorizer = authorizer
		ins_view.Authorizer = authorizer
		fmt.Println()
	}
	return vmClient, ins_view
}
