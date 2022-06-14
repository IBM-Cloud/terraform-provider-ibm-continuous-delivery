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

func TestAccIBMCdToolchainToolCustomDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCustomDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolCustomDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolCustomDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "name"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolCustomDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = "%s"
			parameters {
				type = "type"
				lifecyclePhase = "THINK"
				imageUrl = "imageUrl"
				documentationUrl = "documentationUrl"
				name = "name"
				dashboard_url = "dashboard_url"
				description = "description"
				additional-properties = "additional-properties"
			}
		}

		data "ibmcd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolCustomDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = "%s"
			parameters {
				type = "type"
				lifecyclePhase = "THINK"
				imageUrl = "imageUrl"
				documentationUrl = "documentationUrl"
				name = "name"
				dashboard_url = "dashboard_url"
				description = "description"
				additional-properties = "additional-properties"
			}
			name = "%s"
		}

		data "ibmcd_toolchain_tool_custom" "cd_toolchain_tool_custom" {
			toolchain_id = ibmcd_toolchain_tool_custom.cd_toolchain_tool_custom.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
