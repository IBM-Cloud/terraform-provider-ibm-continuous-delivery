// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"os"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/service/cdtektonpipeline"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/service/cdtoolchain"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/ibm/validate"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bluemix_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Bluemix API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_API_KEY", "BLUEMIX_API_KEY"}, nil),
				Deprecated:  "This field is deprecated please use ibmcloud_api_key",
			},
			"bluemix_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Bluemix API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_TIMEOUT", "BLUEMIX_TIMEOUT"}, nil),
				Deprecated:  "This field is deprecated please use ibmcloud_timeout",
			},
			"ibmcloud_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM Cloud API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_API_KEY", "IBMCLOUD_API_KEY"}, nil),
			},
			"ibmcloud_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_TIMEOUT", "IBMCLOUD_TIMEOUT"}, 60),
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_REGION", "IBMCLOUD_REGION", "BM_REGION", "BLUEMIX_REGION"}, "us-south"),
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM cloud Region zone (for example 'us-south-1') for power resources.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_ZONE", "IBMCLOUD_ZONE"}, ""),
			},
			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Resource group id.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_RESOURCE_GROUP", "IBMCLOUD_RESOURCE_GROUP", "BM_RESOURCE_GROUP", "BLUEMIX_RESOURCE_GROUP"}, ""),
			},
			"softlayer_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SoftLayer API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_API_KEY", "SOFTLAYER_API_KEY"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_api_key",
			},
			"softlayer_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SoftLayer user name",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_USERNAME", "SOFTLAYER_USERNAME"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_username",
			},
			"softlayer_endpoint_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Softlayer Endpoint",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_ENDPOINT_URL", "SOFTLAYER_ENDPOINT_URL"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_endpoint_url",
			},
			"softlayer_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any SoftLayer API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_TIMEOUT", "SOFTLAYER_TIMEOUT"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_timeout",
			},
			"iaas_classic_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Classic Infrastructure API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_API_KEY"}, nil),
			},
			"iaas_classic_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Classic Infrastructure API user name",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_USERNAME"}, nil),
			},
			"iaas_classic_endpoint_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Classic Infrastructure Endpoint",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_ENDPOINT_URL"}, "https://api.softlayer.com/rest/v3"),
			},
			"iaas_classic_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Classic Infrastructure API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_TIMEOUT"}, 60),
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The retry count to set for API calls.",
				DefaultFunc: schema.EnvDefaultFunc("MAX_RETRIES", 10),
			},
			"function_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM Cloud Function namespace",
				DefaultFunc: schema.EnvDefaultFunc("FUNCTION_NAMESPACE", nil),
				Deprecated:  "This field will be deprecated soon",
			},
			"riaas_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The next generation infrastructure service endpoint url.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"RIAAS_ENDPOINT"}, nil),
				Deprecated:  "This field is deprecated use generation",
			},
			"generation": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Generation of Virtual Private Cloud. Default is 2",
				//DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_GENERATION", "IBMCLOUD_GENERATION"}, nil),
				Deprecated: "The generation field is deprecated and will be removed after couple of releases",
			},
			"iam_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM Trusted Profile Authentication token",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_IAM_PROFILE_ID", "IBMCLOUD_IAM_PROFILE_ID"}, nil),
			},
			"iam_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM Authentication token",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_IAM_TOKEN", "IBMCLOUD_IAM_TOKEN"}, nil),
			},
			"iam_refresh_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM Authentication refresh token",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_IAM_REFRESH_TOKEN", "IBMCLOUD_IAM_REFRESH_TOKEN"}, nil),
			},
			"visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private", "public-and-private"}),
				Description:  "Visibility of the provider if it is private or public.",
				DefaultFunc:  schema.MultiEnvDefaultFunc([]string{"IC_VISIBILITY", "IBMCLOUD_VISIBILITY"}, "public"),
			},
			"endpoints_file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path of the file that contains private and public regional endpoints mapping",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_ENDPOINTS_FILE_PATH", "IBMCLOUD_ENDPOINTS_FILE_PATH"}, nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			// // Added for Toolchain
			"ibm_cd_toolchain":                          cdtoolchain.DataSourceIBMCdToolchain(),
			"ibm_cd_toolchain_tool_keyprotect":          cdtoolchain.DataSourceIBMCdToolchainToolKeyprotect(),
			"ibm_cd_toolchain_tool_secretsmanager":      cdtoolchain.DataSourceIBMCdToolchainToolSecretsmanager(),
			"ibm_cd_toolchain_tool_bitbucketgit":        cdtoolchain.DataSourceIBMCdToolchainToolBitbucketgit(),
			"ibm_cd_toolchain_tool_githubintegrated":    cdtoolchain.DataSourceIBMCdToolchainToolGithubintegrated(),
			"ibm_cd_toolchain_tool_githubconsolidated":  cdtoolchain.DataSourceIBMCdToolchainToolGithubconsolidated(),
			"ibm_cd_toolchain_tool_githubpublic":        cdtoolchain.DataSourceIBMCdToolchainToolGithubpublic(),
			"ibm_cd_toolchain_tool_gitlab":              cdtoolchain.DataSourceIBMCdToolchainToolGitlab(),
			"ibm_cd_toolchain_tool_hostedgit":           cdtoolchain.DataSourceIBMCdToolchainToolHostedgit(),
			"ibm_cd_toolchain_tool_artifactory":         cdtoolchain.DataSourceIBMCdToolchainToolArtifactory(),
			"ibm_cd_toolchain_tool_custom":              cdtoolchain.DataSourceIBMCdToolchainToolCustom(),
			"ibm_cd_toolchain_tool_pipeline":            cdtoolchain.DataSourceIBMCdToolchainToolPipeline(),
			"ibm_cd_toolchain_tool_devopsinsights":      cdtoolchain.DataSourceIBMCdToolchainToolDevopsinsights(),
			"ibm_cd_toolchain_tool_slack":               cdtoolchain.DataSourceIBMCdToolchainToolSlack(),
			"ibm_cd_toolchain_tool_sonarqube":           cdtoolchain.DataSourceIBMCdToolchainToolSonarqube(),
			"ibm_cd_toolchain_tool_hashicorpvault":      cdtoolchain.DataSourceIBMCdToolchainToolHashicorpvault(),
			"ibm_cd_toolchain_tool_securitycompliance":  cdtoolchain.DataSourceIBMCdToolchainToolSecuritycompliance(),
			"ibm_cd_toolchain_tool_privateworker":       cdtoolchain.DataSourceIBMCdToolchainToolPrivateworker(),
			"ibm_cd_toolchain_tool_appconfig":           cdtoolchain.DataSourceIBMCdToolchainToolAppconfig(),
			"ibm_cd_toolchain_tool_jenkins":             cdtoolchain.DataSourceIBMCdToolchainToolJenkins(),
			"ibm_cd_toolchain_tool_jira":                cdtoolchain.DataSourceIBMCdToolchainToolJira(),
			"ibm_cd_toolchain_tool_nexus":               cdtoolchain.DataSourceIBMCdToolchainToolNexus(),
			"ibm_cd_toolchain_tool_pagerduty":           cdtoolchain.DataSourceIBMCdToolchainToolPagerduty(),
			"ibm_cd_toolchain_tool_rationalteamconcert": cdtoolchain.DataSourceIBMCdToolchainToolRationalteamconcert(),
			"ibm_cd_toolchain_tool_saucelabs":           cdtoolchain.DataSourceIBMCdToolchainToolSaucelabs(),

			// Added for Tekton Pipeline
			"ibm_cd_tekton_pipeline_definition":       cdtektonpipeline.DataSourceIBMTektonPipelineDefinition(),
			"ibm_cd_tekton_pipeline_trigger_property": cdtektonpipeline.DataSourceIBMTektonPipelineTriggerProperty(),
			"ibm_cd_tekton_pipeline_property":         cdtektonpipeline.DataSourceIBMTektonPipelineProperty(),
			"ibm_cd_tekton_pipeline_trigger":          cdtektonpipeline.DataSourceIBMTektonPipelineTrigger(),
			"ibm_cd_tekton_pipeline":                  cdtektonpipeline.DataSourceIBMTektonPipeline(),
		},

		ResourcesMap: map[string]*schema.Resource{
			// // Added for Toolchain
			"ibm_cd_toolchain":                          cdtoolchain.ResourceIBMCdToolchain(),
			"ibm_cd_toolchain_tool_keyprotect":          cdtoolchain.ResourceIBMCdToolchainToolKeyprotect(),
			"ibm_cd_toolchain_tool_secretsmanager":      cdtoolchain.ResourceIBMCdToolchainToolSecretsmanager(),
			"ibm_cd_toolchain_tool_bitbucketgit":        cdtoolchain.ResourceIBMCdToolchainToolBitbucketgit(),
			"ibm_cd_toolchain_tool_githubintegrated":    cdtoolchain.ResourceIBMCdToolchainToolGithubintegrated(),
			"ibm_cd_toolchain_tool_githubconsolidated":  cdtoolchain.ResourceIBMCdToolchainToolGithubconsolidated(),
			"ibm_cd_toolchain_tool_githubpublic":        cdtoolchain.ResourceIBMCdToolchainToolGithubpublic(),
			"ibm_cd_toolchain_tool_gitlab":              cdtoolchain.ResourceIBMCdToolchainToolGitlab(),
			"ibm_cd_toolchain_tool_hostedgit":           cdtoolchain.ResourceIBMCdToolchainToolHostedgit(),
			"ibm_cd_toolchain_tool_artifactory":         cdtoolchain.ResourceIBMCdToolchainToolArtifactory(),
			"ibm_cd_toolchain_tool_custom":              cdtoolchain.ResourceIBMCdToolchainToolCustom(),
			"ibm_cd_toolchain_tool_pipeline":            cdtoolchain.ResourceIBMCdToolchainToolPipeline(),
			"ibm_cd_toolchain_tool_devopsinsights":      cdtoolchain.ResourceIBMCdToolchainToolDevopsinsights(),
			"ibm_cd_toolchain_tool_slack":               cdtoolchain.ResourceIBMCdToolchainToolSlack(),
			"ibm_cd_toolchain_tool_sonarqube":           cdtoolchain.ResourceIBMCdToolchainToolSonarqube(),
			"ibm_cd_toolchain_tool_hashicorpvault":      cdtoolchain.ResourceIBMCdToolchainToolHashicorpvault(),
			"ibm_cd_toolchain_tool_securitycompliance":  cdtoolchain.ResourceIBMCdToolchainToolSecuritycompliance(),
			"ibm_cd_toolchain_tool_privateworker":       cdtoolchain.ResourceIBMCdToolchainToolPrivateworker(),
			"ibm_cd_toolchain_tool_appconfig":           cdtoolchain.ResourceIBMCdToolchainToolAppconfig(),
			"ibm_cd_toolchain_tool_jenkins":             cdtoolchain.ResourceIBMCdToolchainToolJenkins(),
			"ibm_cd_toolchain_tool_jira":                cdtoolchain.ResourceIBMCdToolchainToolJira(),
			"ibm_cd_toolchain_tool_nexus":               cdtoolchain.ResourceIBMCdToolchainToolNexus(),
			"ibm_cd_toolchain_tool_pagerduty":           cdtoolchain.ResourceIBMCdToolchainToolPagerduty(),
			"ibm_cd_toolchain_tool_rationalteamconcert": cdtoolchain.ResourceIBMCdToolchainToolRationalteamconcert(),
			"ibm_cd_toolchain_tool_saucelabs":           cdtoolchain.ResourceIBMCdToolchainToolSaucelabs(),

			// // Added for Tekton Pipeline
			"ibm_cd_tekton_pipeline_definition":       cdtektonpipeline.ResourceIBMTektonPipelineDefinition(),
			"ibm_cd_tekton_pipeline_trigger_property": cdtektonpipeline.ResourceIBMTektonPipelineTriggerProperty(),
			"ibm_cd_tekton_pipeline_property":         cdtektonpipeline.ResourceIBMTektonPipelineProperty(),
			"ibm_cd_tekton_pipeline_trigger":          cdtektonpipeline.ResourceIBMTektonPipelineTrigger(),
			"ibm_cd_tekton_pipeline":                  cdtektonpipeline.ResourceIBMTektonPipeline(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var globalValidatorDict validate.ValidatorDict
var initOnce sync.Once

func init() {
	validate.SetValidatorDict(Validator())
}

// Validator return validator
func Validator() validate.ValidatorDict {
	initOnce.Do(func() {
		globalValidatorDict = validate.ValidatorDict{
			ResourceValidatorDictionary: map[string]*validate.ResourceValidator{
				// // Added for Toolchains
				"ibm_cd_toolchain":                          cdtoolchain.ResourceIBMCdToolchainValidator(),
				"ibm_cd_toolchain_tool_keyprotect":          cdtoolchain.ResourceIBMCdToolchainToolKeyprotectValidator(),
				"ibm_cd_toolchain_tool_secretsmanager":      cdtoolchain.ResourceIBMCdToolchainToolSecretsmanagerValidator(),
				"ibm_cd_toolchain_tool_bitbucketgit":        cdtoolchain.ResourceIBMCdToolchainToolBitbucketgitValidator(),
				"ibm_cd_toolchain_tool_githubintegrated":    cdtoolchain.ResourceIBMCdToolchainToolGithubintegratedValidator(),
				"ibm_cd_toolchain_tool_githubconsolidated":  cdtoolchain.ResourceIBMCdToolchainToolGithubconsolidatedValidator(),
				"ibm_cd_toolchain_tool_githubpublic":        cdtoolchain.ResourceIBMCdToolchainToolGithubpublicValidator(),
				"ibm_cd_toolchain_tool_gitlab":              cdtoolchain.ResourceIBMCdToolchainToolGitlabValidator(),
				"ibm_cd_toolchain_tool_hostedgit":           cdtoolchain.ResourceIBMCdToolchainToolHostedgitValidator(),
				"ibm_cd_toolchain_tool_artifactory":         cdtoolchain.ResourceIBMCdToolchainToolArtifactoryValidator(),
				"ibm_cd_toolchain_tool_custom":              cdtoolchain.ResourceIBMCdToolchainToolCustomValidator(),
				"ibm_cd_toolchain_tool_pipeline":            cdtoolchain.ResourceIBMCdToolchainToolPipelineValidator(),
				"ibm_cd_toolchain_tool_slack":               cdtoolchain.ResourceIBMCdToolchainToolSlackValidator(),
				"ibm_cd_toolchain_tool_devopsinsights":      cdtoolchain.ResourceIBMCdToolchainToolDevopsinsightsValidator(),
				"ibm_cd_toolchain_tool_sonarqube":           cdtoolchain.ResourceIBMCdToolchainToolSonarqubeValidator(),
				"ibm_cd_toolchain_tool_hashicorpvault":      cdtoolchain.ResourceIBMCdToolchainToolHashicorpvaultValidator(),
				"ibm_cd_toolchain_tool_securitycompliance":  cdtoolchain.ResourceIBMCdToolchainToolSecuritycomplianceValidator(),
				"ibm_cd_toolchain_tool_privateworker":       cdtoolchain.ResourceIBMCdToolchainToolPrivateworkerValidator(),
				"ibm_cd_toolchain_tool_appconfig":           cdtoolchain.ResourceIBMCdToolchainToolAppconfigValidator(),
				"ibm_cd_toolchain_tool_jenkins":             cdtoolchain.ResourceIBMCdToolchainToolJenkinsValidator(),
				"ibm_cd_toolchain_tool_jira":                cdtoolchain.ResourceIBMCdToolchainToolJiraValidator(),
				"ibm_cd_toolchain_tool_nexus":               cdtoolchain.ResourceIBMCdToolchainToolNexusValidator(),
				"ibm_cd_toolchain_tool_pagerduty":           cdtoolchain.ResourceIBMCdToolchainToolPagerdutyValidator(),
				"ibm_cd_toolchain_tool_rationalteamconcert": cdtoolchain.ResourceIBMCdToolchainToolRationalteamconcertValidator(),
				"ibm_cd_toolchain_tool_saucelabs":           cdtoolchain.ResourceIBMCdToolchainToolSaucelabsValidator(),

				// // Added for Tekton Pipeline
				"ibm_cd_tekton_pipeline_definition":       cdtektonpipeline.ResourceIBMTektonPipelineDefinitionValidator(),
				"ibm_cd_tekton_pipeline_trigger_property": cdtektonpipeline.ResourceIBMTektonPipelineTriggerPropertyValidator(),
				"ibm_cd_tekton_pipeline_property":         cdtektonpipeline.ResourceIBMTektonPipelinePropertyValidator(),
				"ibm_cd_tekton_pipeline_trigger":          cdtektonpipeline.ResourceIBMTektonPipelineTriggerValidator(),
			},
			DataSourceValidatorDictionary: map[string]*validate.ResourceValidator{},
		}
	})
	return globalValidatorDict
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var bluemixAPIKey string
	var bluemixTimeout int
	var iamToken, iamRefreshToken, iamTrustedProfileId string
	if key, ok := d.GetOk("bluemix_api_key"); ok {
		bluemixAPIKey = key.(string)
	}
	if key, ok := d.GetOk("ibmcloud_api_key"); ok {
		bluemixAPIKey = key.(string)
	}
	if itoken, ok := d.GetOk("iam_token"); ok {
		iamToken = itoken.(string)
	}
	if rtoken, ok := d.GetOk("iam_refresh_token"); ok {
		iamRefreshToken = rtoken.(string)
	}
	if ttoken, ok := d.GetOk("iam_profile_id"); ok {
		iamTrustedProfileId = ttoken.(string)
	}
	var softlayerUsername, softlayerAPIKey, softlayerEndpointUrl string
	var softlayerTimeout int
	if username, ok := d.GetOk("softlayer_username"); ok {
		softlayerUsername = username.(string)
	}
	if username, ok := d.GetOk("iaas_classic_username"); ok {
		softlayerUsername = username.(string)
	}
	if apikey, ok := d.GetOk("softlayer_api_key"); ok {
		softlayerAPIKey = apikey.(string)
	}
	if apikey, ok := d.GetOk("iaas_classic_api_key"); ok {
		softlayerAPIKey = apikey.(string)
	}
	if endpoint, ok := d.GetOk("softlayer_endpoint_url"); ok {
		softlayerEndpointUrl = endpoint.(string)
	}
	if endpoint, ok := d.GetOk("iaas_classic_endpoint_url"); ok {
		softlayerEndpointUrl = endpoint.(string)
	}
	if tm, ok := d.GetOk("softlayer_timeout"); ok {
		softlayerTimeout = tm.(int)
	}
	if tm, ok := d.GetOk("iaas_classic_timeout"); ok {
		softlayerTimeout = tm.(int)
	}

	if tm, ok := d.GetOk("bluemix_timeout"); ok {
		bluemixTimeout = tm.(int)
	}
	if tm, ok := d.GetOk("ibmcloud_timeout"); ok {
		bluemixTimeout = tm.(int)
	}
	var visibility string
	if v, ok := d.GetOk("visibility"); ok {
		visibility = v.(string)
	}
	var file string
	if f, ok := d.GetOk("endpoints_file_path"); ok {
		file = f.(string)
	}

	resourceGrp := d.Get("resource_group").(string)
	region := d.Get("region").(string)
	zone := d.Get("zone").(string)
	retryCount := d.Get("max_retries").(int)
	wskNameSpace := d.Get("function_namespace").(string)
	riaasEndPoint := d.Get("riaas_endpoint").(string)

	wskEnvVal, err := schema.EnvDefaultFunc("FUNCTION_NAMESPACE", "")()
	if err != nil {
		return nil, err
	}
	//Set environment variable to be used in DiffSupressFunction
	if wskEnvVal.(string) == "" {
		os.Setenv("FUNCTION_NAMESPACE", wskNameSpace)
	}

	config := conns.Config{
		BluemixAPIKey:        bluemixAPIKey,
		Region:               region,
		ResourceGroup:        resourceGrp,
		BluemixTimeout:       time.Duration(bluemixTimeout) * time.Second,
		SoftLayerTimeout:     time.Duration(softlayerTimeout) * time.Second,
		SoftLayerUserName:    softlayerUsername,
		SoftLayerAPIKey:      softlayerAPIKey,
		RetryCount:           retryCount,
		SoftLayerEndpointURL: softlayerEndpointUrl,
		RetryDelay:           conns.RetryAPIDelay,
		FunctionNameSpace:    wskNameSpace,
		RiaasEndPoint:        riaasEndPoint,
		IAMToken:             iamToken,
		IAMRefreshToken:      iamRefreshToken,
		Zone:                 zone,
		Visibility:           visibility,
		EndpointsFile:        file,
		IAMTrustedProfileID:  iamTrustedProfileId,
	}

	return config.ClientSession()
}
