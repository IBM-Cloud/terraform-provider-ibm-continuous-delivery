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

func TestAccIBMCdToolchainToolPrivateworkerDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPrivateworkerDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				workerQueueCredentials = "workerQueueCredentials"
				workerQueueIdentifier = "workerQueueIdentifier"
			}
		}

		data "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
