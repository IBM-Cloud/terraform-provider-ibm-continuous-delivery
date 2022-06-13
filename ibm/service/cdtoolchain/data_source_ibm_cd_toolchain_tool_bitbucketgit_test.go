// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/acctest"
)

func TestAccIBMCdToolchainToolBitbucketgitDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolBitbucketgitDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolBitbucketgitDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibmcd_toolchain_tool_bitbucketgit" "cd_toolchain_tool_bitbucketgit" {
			toolchain_id = "toolchain_id"
			tool_id = "tool_id"
		}
	`)
}
