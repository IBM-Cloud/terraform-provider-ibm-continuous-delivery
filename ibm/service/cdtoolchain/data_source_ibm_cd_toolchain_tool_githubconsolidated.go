// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func DataSourceIBMCdToolchainToolGithubconsolidated() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMCdToolchainToolGithubconsolidatedRead,

		Schema: map[string]*schema.Schema{
			"toolchain_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the toolchain.",
			},
			"tool_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the tool bound to the toolchain.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool ID.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group where tool can be found.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool CRN.",
			},
			"toolchain_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of toolchain which the tool is bound to.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI representing the tool.",
			},
			"referent": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information on URIs to access this resource through the UI or API.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ui_href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI representing the this resource through the UI.",
						},
						"api_href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI representing the this resource through an API.",
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tool name.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest tool update timestamp.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parameters to be used to create the tool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"legal": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"git_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. My GitHub Enterprise Server.",
						},
						"api_root_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. https://api.github.example.com.",
						},
						"default_branch": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. main.",
						},
						"root_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e.g. https://github.example.com.",
						},
						"access_token": &schema.Schema{
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
						"blind_connection": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. Certain functionality that requires API access to the git server will be disabled. Delivery pipeline will only work using a private worker that has network access to the git server.",
						},
						"owner_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repo_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL of the repository that you are linking to.",
						},
						"source_repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type the URL of the repository that you are forking or cloning.",
						},
						"token_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration token URL.",
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_repo": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to make this repository private.",
						},
						"has_issues": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to enable GitHub Issues for lightweight issue tracking.",
						},
						"auto_init": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this checkbox to initialize this repository with a README.",
						},
						"enable_traceability": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select this check box to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.",
						},
						"authorized": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"integration_owner": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Select the user which git operations will be performed as.",
						},
						"auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"api_token": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "Personal Access Token.",
						},
					},
				},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current configuration state of the tool.",
			},
		},
	}
}

func DataSourceIBMCdToolchainToolGithubconsolidatedRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdToolchainClient, err := meta.(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

	getToolByIDOptions.SetToolchainID(d.Get("toolchain_id").(string))
	getToolByIDOptions.SetToolID(d.Get("tool_id").(string))

	getToolByIDResponse, response, err := cdToolchainClient.GetToolByIDWithContext(context, getToolByIDOptions)
	if err != nil {
		log.Printf("[DEBUG] GetToolByIDWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetToolByIDWithContext failed %s\n%s", err, response))
	}

	if *getToolByIDResponse.ToolTypeID != "githubconsolidated" {
		return diag.FromErr(fmt.Errorf("Retrieved tool is not the correct type: %s", *getToolByIDResponse.ToolTypeID))
	}

	d.SetId(*getToolByIDResponse.ID)

	if err = d.Set("id", getToolByIDResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}

	if err = d.Set("resource_group_id", getToolByIDResponse.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}

	if err = d.Set("crn", getToolByIDResponse.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	if err = d.Set("toolchain_crn", getToolByIDResponse.ToolchainCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain_crn: %s", err))
	}

	if err = d.Set("href", getToolByIDResponse.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	referent := []map[string]interface{}{}
	if getToolByIDResponse.Referent != nil {
		modelMap, err := DataSourceIBMCdToolchainToolGithubconsolidatedToolReferentToMap(getToolByIDResponse.Referent)
		if err != nil {
			return diag.FromErr(err)
		}
		referent = append(referent, modelMap)
	}
	if err = d.Set("referent", referent); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting referent %s", err))
	}

	if err = d.Set("name", getToolByIDResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(getToolByIDResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	parameters := []map[string]interface{}{}
	if getToolByIDResponse.Parameters != nil {
		modelMap := GetParametersFromRead(getToolByIDResponse.Parameters, DataSourceIBMCdToolchainToolGithubconsolidated(), nil)
		parameters = append(parameters, modelMap)
	}
	if err = d.Set("parameters", parameters); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting parameters %s", err))
	}

	if err = d.Set("state", getToolByIDResponse.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}

	return nil
}

func DataSourceIBMCdToolchainToolGithubconsolidatedToolReferentToMap(model *cdtoolchainv2.ToolReferent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UIHref != nil {
		modelMap["ui_href"] = *model.UIHref
	}
	if model.APIHref != nil {
		modelMap["api_href"] = *model.APIHref
	}
	return modelMap, nil
}
