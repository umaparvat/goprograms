package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	subscriptionID := "GGGG730f4-3ea1-45f5-969e-XXXXXX"
	appId := "YYYYYc96-2d30-4f0f-bcee-XXXXX"
	tenant := "YYYYYY-cb1e-40f7-b59a-XXXXX"
	appSecret := "GGGGGG-9~KK9tIV4Lli9XXXXX"
	clientAuthorizer := auth.NewClientCredentialsConfig(appId, appSecret, tenant)
	authorizer, err := clientAuthorizer.Authorizer()
	tagClient := resources.NewTagsClient(subscriptionID)
	subClient := subscriptions.NewClient()
	if err != nil {
		fmt.Println(err)
	} else {

		tagClient.Authorizer = authorizer
		subClient.Authorizer = authorizer

	}
	t := make(map[string]string)
	// working list complete
	for list, err := tagClient.ListComplete(context.Background()); list.NotDone(); err = list.Next() {
		if err != nil {
			fmt.Println(err, "error traverising List result")
		}
		tagDetail := list.Value()
		fmt.Println(*tagDetail.TagName, *tagDetail.Values)
		fmt.Println("----------------------###########-----")

	}
	//a single page value
	list, err := tagClient.List(context.Background())
	for _, tagDetails :=  range list.Values() {
		fmt.Println(*tagDetails.TagName, *tagDetails.Values)
	}


	// To create subscription tag
	tagName := "test_r"
	value := "checking"
	_, cerr := tagClient.CreateOrUpdate(context.Background(), tagName)
	if cerr != nil {
			fmt.Errorf("checking for existence of tags by tagName %q: %+v", tagName, cerr)

	}
	existingtagsval, err := tagClient.CreateOrUpdateValue(context.Background(), tagName, value)
	if err != nil {
		 fmt.Errorf("Adding tag value %q for tags by tagName %q: %+v", value, tagName, err)

	}
	fmt.Println("\ncreateor update tag key and value\n", *existingtagsval.ID, *existingtagsval.TagValue)
	fmt.Println("\n status code of tagvalue", *&existingtagsval.StatusCode)
	//To create tag at subscriptionscope
	tvalue := "tag-value-3"
	tval := make(map[string]*string)
	tval["tagKey3"] = &tvalue
	fmt.Println("map tag", tval)
	rtags := resources.Tags{
		Tags: tval,
	}
	fmt.Println("\nresource tags", rtags)
	tagParameter := resources.TagsResource{Properties: &rtags}
	fmt.Printf("Tag resource parameter", tagParameter)
	// to create a tag at subscription scope
	crtags, crerr := tagClient.CreateOrUpdateAtScope(context.Background(), "subscriptions/"+subscriptionID, tagParameter)
	if crerr != nil {
		fmt.Errorf("retireving tags from subscription %q: %+v", subscriptionID, crerr)

	}
	fmt.Println("printing properties:\n", *&crtags.StatusCode)
	// To update at a subscription Scope for delete
	tagPatchParamter := resources.TagsPatchResource{Operation: "Merge", Properties: &rtags}
	fmt.Println("\n\n TagPatch Operation\n\n", tagPatchParamter)
	uptags, urerr := tagClient.UpdateAtScope(context.Background(), "subscriptions/"+subscriptionID, tagPatchParamter)
	if urerr != nil {
		fmt.Errorf("updating tags from subscription %q: %+v", subscriptionID, urerr)

	}
	fmt.Println("printing properties:\n", *&uptags.StatusCode)

	// to get a subscriptionscope
	subTags, serr := tagClient.GetAtScope(context.Background(), "subscriptions/"+subscriptionID)
	if serr != nil {
		fmt.Errorf("retireving tags from subscription %q: %+v", subscriptionID, serr)
	}
	fmt.Println("get at scope output\n", *subTags.ID, *subTags.Name, Flatten(*&subTags.Properties.Tags))
	//To get the subscription tags alone
	//SubscriptionTagRead(subClient, context.Background(), subscriptionID)

}

func SubscriptionTagRead(subClient subscriptions.Client, ctx context.Context, subscriptionID string) error {
	existingSub, err := subClient.Get(ctx, subscriptionID)
	if err != nil {
		fmt.Errorf("could not read existing Subscription %q", subscriptionID)
	}
	t := existingSub.Tags
	fmt.Println("sub output\n", Flatten(t))
	return nil
}

func Flatten(tagMap map[string]*string) map[string]interface{} {
	// If tagsMap is nil, len(tagsMap) will be 0.
	output := make(map[string]interface{}, len(tagMap))

	for i, v := range tagMap {
		if v == nil {
			continue
		}

		output[i] = *v
	}

	return output
}
