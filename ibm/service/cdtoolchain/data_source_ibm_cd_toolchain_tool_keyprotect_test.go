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

func TestAccIBMCdToolchainToolKeyprotectDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolKeyprotectDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "name"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				region = "region"
				resource-group = "resource-group"
				instance-name = "instance-name"
				integration-status = "integration-status"
			}
		}

		data "ibmcd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				region = "region"
				resource-group = "resource-group"
				instance-name = "instance-name"
				integration-status = "integration-status"
			}
			name = "%s"
		}

		data "ibmcd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibmcd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
