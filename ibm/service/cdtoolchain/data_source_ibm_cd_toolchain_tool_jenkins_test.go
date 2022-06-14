// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/acctest"
)

func TestAccIBMCdToolchainToolJenkinsDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolJenkinsDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "name"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolJenkinsDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				webhook_url = "webhook_url"
				api_user_name = "api_user_name"
				api_token = "api_token"
			}
		}

		data "ibmcd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolJenkinsDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				webhook_url = "webhook_url"
				api_user_name = "api_user_name"
				api_token = "api_token"
			}
			name = "%s"
		}

		data "ibmcd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibmcd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
