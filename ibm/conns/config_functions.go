// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package conns

import (
	"fmt"
	"net/http"
	"net/url"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/apache/openwhisk-client-go/whisk"
)

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.functions.cloud.ibm.com"

//FunctionClient ...
func FunctionClient(c *bluemix.Config) (*whisk.Client, error) {
	baseEndpoint := getBaseURL(c.Region)
	u, err := url.Parse(fmt.Sprintf("%s/api", baseEndpoint))
	if err != nil {
		return nil, err
	}

	functionsClient, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
		Host:    u.Host,
		Version: "v1",
	})

	return functionsClient, err
}

//getBaseURL ..
func getBaseURL(region string) string {
	baseEndpoint := fmt.Sprintf(DefaultServiceURL)
	if region != "us-south" {
		baseEndpoint = fmt.Sprintf("https://%s.functions.cloud.ibm.com", region)
	}

	return baseEndpoint
}
