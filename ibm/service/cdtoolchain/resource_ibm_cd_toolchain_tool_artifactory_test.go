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

func TestAccIBMCdToolchainToolArtifactoryBasic(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolArtifactoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolArtifactoryExists("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolArtifactoryAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolArtifactoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolArtifactoryExists("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolArtifactoryConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_artifactory.cd_toolchain_tool_artifactory",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolArtifactoryConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolArtifactoryConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_artifactory" "cd_toolchain_tool_artifactory" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				type = "npm"
				user_id = "user_id"
				token = "token"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
				repository_name = "repository_name"
				repository_url = "repository_url"
				docker_config_json = "docker_config_json"
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolArtifactoryExists(n string, obj cdtoolchainv2.GetToolByIDResponse) resource.TestCheckFunc {

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

func testAccCheckIBMCdToolchainToolArtifactoryDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_artifactory" {
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
			return fmt.Errorf("cd_toolchain_tool_artifactory still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_artifactory (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}