// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/acctest"
)

func TestAccIBMCdToolchainToolGithubpublicDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolGithubpublicDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "href"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibmcd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolGithubpublicDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibmcd_toolchain_tool_githubpublic" "cd_toolchain_tool_githubpublic" {
			toolchain_id = "toolchain_id"
			tool_id = "tool_id"
		}
	`)
}
