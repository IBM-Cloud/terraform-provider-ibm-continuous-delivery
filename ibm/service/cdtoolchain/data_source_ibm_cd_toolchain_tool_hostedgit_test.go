// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/acctest"
)

func TestAccIBMCdToolchainToolHostedgitDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolHostedgitDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_hostedgit.cd_toolchain_tool_hostedgit", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolHostedgitDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibmcd_toolchain_tool_hostedgit" "cd_toolchain_tool_hostedgit" {
			toolchain_id = "toolchain_id"
			tool_id = "tool_id"
		}
	`)
}
