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

func TestAccIBMCdToolchainToolSlackBasic(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolSlackDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolSlackExists("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", conf),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSlackAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolSlackDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolSlackExists("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", conf),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSlackConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibmcd_toolchain_tool_slack.cd_toolchain_tool_slack",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSlackConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibmcd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = "%s"
			parameters {
				api_token = "api_token"
				channel_name = "channel_name"
				team_url = "team_url"
				pipeline_start = true
				pipeline_success = true
				pipeline_fail = true
				toolchain_bind = true
				toolchain_unbind = true
			}
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolSlackConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibmcd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
			toolchain_id = "%s"
			parameters {
				api_token = "api_token"
				channel_name = "channel_name"
				team_url = "team_url"
				pipeline_start = true
				pipeline_success = true
				pipeline_fail = true
				toolchain_bind = true
				toolchain_unbind = true
			}
			name = "%s"
		}
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolSlackExists(n string, obj cdtoolchainv2.GetToolByIDResponse) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolSlackDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibmcd_toolchain_tool_slack" {
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
			return fmt.Errorf("cd_toolchain_tool_slack still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_slack (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
