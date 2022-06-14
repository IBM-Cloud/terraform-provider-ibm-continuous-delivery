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

func TestAccIBMCdToolchainToolPipelineDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPipelineDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPipelineDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				type = "classic"
				ui_pipeline = true
			}
		}

		data "ibmcd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolPipelineDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = "%s"
			parameters {
				name = "name"
				type = "classic"
				ui_pipeline = true
			}
			name = "%s"
		}

		data "ibmcd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibmcd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
