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

func TestAccIBMCdToolchainToolSecretsmanagerDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSecretsmanagerDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "name"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = "%s"
		}

		data "ibmcd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolSecretsmanagerDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibmcd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				region = "region"
				resource-group = "resource-group"
				instance-name = "instance-name"
				integration-status = "integration-status"
			}
		}

		data "ibmcd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager" {
			toolchain_id = ibmcd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
