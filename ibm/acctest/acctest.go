// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"fmt"
	"os"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var CdResourceGroupID string

func init() {
	testlogger := os.Getenv("TF_LOG")
	if testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}

	CdResourceGroupID = os.Getenv("IBM_CD_RESOURCE_GROUP_ID")
	if CdResourceGroupID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_RESOURCE_GROUP_ID for testing CD resources, CD tests will fail if this is not set")
	}
}

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = provider.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"ibmcd": TestAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := provider.Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = provider.Provider()
}

func TestAccPreCheck(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
}

func TestAccPreCheckCd(t *testing.T) {
	TestAccPreCheck(t)
	if CdResourceGroupID == "" {
		t.Fatal("IBM_CD_RESOURCE_GROUP_ID must be set for acceptance tests")
	}
}
