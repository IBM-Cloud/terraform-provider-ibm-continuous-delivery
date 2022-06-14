// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolHashicorpvaultBasic(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolHashicorpvaultDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolHashicorpvaultExists("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", conf),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolHashicorpvaultAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolHashicorpvaultDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolHashicorpvaultExists("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", conf),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHashicorpvaultConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibmcd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolHashicorpvaultConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibmcd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				server_url = "server_url"
				authentication_method = "token"
				token = "token"
				role_id = "role_id"
				secret_id = "secret_id"
				dashboard_url = "dashboard_url"
				path = "path"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolHashicorpvaultConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibmcd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				server_url = "server_url"
				authentication_method = "token"
				token = "token"
				role_id = "role_id"
				secret_id = "secret_id"
				dashboard_url = "dashboard_url"
				path = "path"
				secret_filter = "secret_filter"
				default_secret = "default_secret"
				username = "username"
				password = "password"
			}
			name = "%s"
		}
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolHashicorpvaultExists(n string, obj cdtoolchainv2.GetToolByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		getToolByIDResponse, _, err := cdToolchainClient.GetToolByID(getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *getToolByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolHashicorpvaultDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibmcd_toolchain_tool_hashicorpvault" {
			continue
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetToolByID(getToolByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_hashicorpvault still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_hashicorpvault (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
