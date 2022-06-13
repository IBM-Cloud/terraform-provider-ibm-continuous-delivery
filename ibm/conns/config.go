// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package conns

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	gohttp "net/http"
	"os"
	"strings"
	"time"

	// Added code for the Power Colo Offering

	"github.com/IBM/go-sdk-core/v5/core"
	iamidentity "github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/apache/openwhisk-client-go/whisk"
	jwt "github.com/golang-jwt/jwt"
	slsession "github.com/softlayer/softlayer-go/session"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery/version"
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

// RetryAPIDelay - retry api delay
const RetryAPIDelay = 5 * time.Second

//BluemixRegion ...
var BluemixRegion string

var (
	errEmptyBluemixCredentials = errors.New("ibmcloud_api_key or bluemix_api_key or iam_token and iam_refresh_token must be provided. Please see the documentation on how to configure it")
)

//UserConfig ...
type UserConfig struct {
	UserID      string
	UserEmail   string
	UserAccount string
	CloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
	generation  int    `default:"2"`
}

//Config stores user provider input
type Config struct {
	//BluemixAPIKey is the Bluemix api key
	BluemixAPIKey string
	//Bluemix region
	Region string
	//Resource group id
	ResourceGroup string
	//Bluemix API timeout
	BluemixTimeout time.Duration

	//Softlayer end point url
	SoftLayerEndpointURL string

	//Softlayer API timeout
	SoftLayerTimeout time.Duration

	// Softlayer User Name
	SoftLayerUserName string

	// Softlayer API Key
	SoftLayerAPIKey string

	//Retry Count for API calls
	//Unexposed in the schema at this point as they are used only during session creation for a few calls
	//When sdk implements it we an expose them for expected behaviour
	//https://github.com/softlayer/softlayer-go/issues/41
	RetryCount int
	//Constant Retry Delay for API calls
	RetryDelay time.Duration

	// FunctionNameSpace ...
	FunctionNameSpace string

	//Riaas End point
	RiaasEndPoint string

	//Generation
	Generation int

	//IAM Token
	IAMToken string

	//TrustedProfileToken Token
	IAMTrustedProfileID string

	//IAM Refresh Token
	IAMRefreshToken string

	// Zone
	Zone          string
	Visibility    string
	EndpointsFile string
}

//Session stores the information required for communication with the SoftLayer and Bluemix API
type Session struct {
	// SoftLayerSesssion is the the SoftLayer session used to connect to the SoftLayer API
	SoftLayerSession *slsession.Session

	// BluemixSession is the the Bluemix session used to connect to the Bluemix API
	BluemixSession *bxsession.Session
}

// ClientSession ...
type ClientSession interface {
	BluemixSession() (*bxsession.Session, error)
	BluemixUserDetails() (*UserConfig, error)
	FunctionClient() (*whisk.Client, error)
	CdToolchainV2() (*cdtoolchainv2.CdToolchainV2, error)
	CdTektonPipelineV2() (*cdtektonpipelinev2.CdTektonPipelineV2, error)
}

type clientSession struct {
	session *Session

	bmxUserDetails  *UserConfig
	bmxUserFetchErr error

	functionConfigErr error
	functionClient    *whisk.Client

	bluemixSessionErr error

	// CD Toolchain
	cdToolchainClient    *cdtoolchainv2.CdToolchainV2
	cdToolchainClientErr error

	// CD Tekton Pipeline
	cdTektonPipelineClient    *cdtektonpipelinev2.CdTektonPipelineV2
	cdTektonPipelineClientErr error
}

// BluemixSession to provide the Bluemix Session
func (sess clientSession) BluemixSession() (*bxsession.Session, error) {
	return sess.session.BluemixSession, sess.bluemixSessionErr
}

// BluemixUserDetails ...
func (sess clientSession) BluemixUserDetails() (*UserConfig, error) {
	return sess.bmxUserDetails, sess.bmxUserFetchErr
}

// FunctionClient ...
func (sess clientSession) FunctionClient() (*whisk.Client, error) {
	return sess.functionClient, sess.functionConfigErr
}

var cloudEndpoint = "cloud.ibm.com"

// CD Toolchain
func (session clientSession) CdToolchainV2() (*cdtoolchainv2.CdToolchainV2, error) {
	return session.cdToolchainClient, session.cdToolchainClientErr
}

// CD Tekton Pipeline
func (session clientSession) CdTektonPipelineV2() (*cdtektonpipelinev2.CdTektonPipelineV2, error) {
	return session.cdTektonPipelineClient, session.cdTektonPipelineClientErr
}

// ClientSession configures and returns a fully initialized ClientSession
func (c *Config) ClientSession() (interface{}, error) {
	sess, err := newSession(c)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Configured Region: %s\n", c.Region)
	session := clientSession{
		session: sess,
	}

	if sess.BluemixSession == nil {
		//Can be nil only  if bluemix_api_key is not provided
		log.Println("Skipping Bluemix Clients configuration")
		session.bluemixSessionErr = errEmptyBluemixCredentials
		session.functionConfigErr = errEmptyBluemixCredentials
		session.bmxUserFetchErr = errEmptyBluemixCredentials

		return session, nil
	}

	if sess.BluemixSession.Config.BluemixAPIKey != "" {
		err = authenticateAPIKey(sess.BluemixSession)
		if err != nil {
			for count := c.RetryCount; count >= 0; count-- {
				if err == nil || !isRetryable(err) {
					break
				}
				time.Sleep(c.RetryDelay)
				log.Printf("Retrying IAM Authentication %d", count)
				err = authenticateAPIKey(sess.BluemixSession)
			}
			if err != nil {
				session.bmxUserFetchErr = fmt.Errorf("[ERROR] Error occured while fetching auth key for account user details: %q", err)
				session.functionConfigErr = fmt.Errorf("[ERROR] Error occured while fetching auth key for function: %q", err)
			}
		}
		err = authenticateCF(sess.BluemixSession)
		if err != nil {
			for count := c.RetryCount; count >= 0; count-- {
				if err == nil || !isRetryable(err) {
					break
				}
				time.Sleep(c.RetryDelay)
				log.Printf("Retrying CF Authentication %d", count)
				err = authenticateCF(sess.BluemixSession)
			}
			if err != nil {
				session.functionConfigErr = fmt.Errorf("[ERROR] Error occured while fetching auth key for function: %q", err)
			}
		}
	}

	if c.IAMTrustedProfileID == "" && sess.BluemixSession.Config.IAMAccessToken != "" && sess.BluemixSession.Config.BluemixAPIKey == "" {
		err := RefreshToken(sess.BluemixSession)
		if err != nil {
			for count := c.RetryCount; count >= 0; count-- {
				if err == nil || !isRetryable(err) {
					break
				}
				time.Sleep(c.RetryDelay)
				log.Printf("Retrying refresh token %d", count)
				err = RefreshToken(sess.BluemixSession)
			}
			if err != nil {
				return nil, fmt.Errorf("[ERROR] Error occured while refreshing the token: %q", err)
			}
		}

	}
	userConfig, err := fetchUserDetails(sess.BluemixSession, c.RetryCount, c.RetryDelay)
	if err != nil {
		session.bmxUserFetchErr = fmt.Errorf("[ERROR] Error occured while fetching account user details: %q", err)
	}
	session.bmxUserDetails = userConfig

	if sess.SoftLayerSession != nil && sess.SoftLayerSession.IAMToken != "" {
		sess.SoftLayerSession.IAMToken = sess.BluemixSession.Config.IAMAccessToken
		sess.SoftLayerSession.IAMRefreshToken = sess.BluemixSession.Config.IAMRefreshToken
	}

	session.functionClient, session.functionConfigErr = FunctionClient(sess.BluemixSession.Config)

	BluemixRegion = sess.BluemixSession.Config.Region
	var fileMap map[string]interface{}
	if f := EnvFallBack([]string{"IBMCLOUD_ENDPOINTS_FILE_PATH", "IC_ENDPOINTS_FILE_PATH"}, c.EndpointsFile); f != "" {
		jsonFile, err := os.Open(f)
		if err != nil {
			log.Fatalf("Unable to open Endpoints File %s", err)
		}
		defer jsonFile.Close()
		bytes, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatalf("Unable to read Endpoints File %s", err)
		}
		err = json.Unmarshal([]byte(bytes), &fileMap)
		if err != nil {
			log.Fatalf("Unable to unmarshal Endpoints File %s", err)
		}
	}

	iamURL := iamidentity.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			iamURL = ContructEndpoint(fmt.Sprintf("private.%s.iam", c.Region), cloudEndpoint)
		} else {
			iamURL = ContructEndpoint("private.iam", cloudEndpoint)
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		iamURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_IAM_API_ENDPOINT", c.Region, iamURL)
	}

	var authenticator core.Authenticator

	if c.BluemixAPIKey != "" || sess.BluemixSession.Config.IAMRefreshToken != "" {
		if c.BluemixAPIKey != "" {
			authenticator = &core.IamAuthenticator{
				ApiKey: c.BluemixAPIKey,
				URL:    EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL),
			}
		} else {
			// Construct the IamAuthenticator with the IAM refresh token.
			authenticator = &core.IamAuthenticator{
				RefreshToken: sess.BluemixSession.Config.IAMRefreshToken,
				ClientId:     "bx",
				ClientSecret: "bx",
				URL:          EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL),
			}
		}
	} else if strings.HasPrefix(sess.BluemixSession.Config.IAMAccessToken, "Bearer") {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken[7:],
		}
	} else {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken,
		}
	}

	// Construct an "options" struct for creating the service client.
	var cdToolchainClientURL string
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		cdToolchainClientURL, err = cdtoolchainv2.GetServiceURLForRegion("private." + c.Region)
		if err != nil && c.Visibility == "public-and-private" {
			cdToolchainClientURL, err = cdtoolchainv2.GetServiceURLForRegion(c.Region)
		}
	} else {
		cdToolchainClientURL, err = cdtoolchainv2.GetServiceURLForRegion(c.Region)
	}
	if err != nil {
		cdToolchainClientURL = cdtoolchainv2.DefaultServiceURL
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		cdToolchainClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_TOOLCHAIN_ENDPOINT", c.Region, cdToolchainClientURL)
	}
	cdToolchainClientOptions := &cdtoolchainv2.CdToolchainV2Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_TOOLCHAIN_ENDPOINT"}, cdToolchainClientURL),
	}

	// Construct the service client.
	session.cdToolchainClient, err = cdtoolchainv2.NewCdToolchainV2(cdToolchainClientOptions)
	if err == nil {
		// Enable retries for API calls
		session.cdToolchainClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.cdToolchainClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm-continuous-delivery/%s", version.Version)},
		})
	} else {
		session.cdToolchainClientErr = fmt.Errorf("Error occurred while configuring Toolchain service: %q", err)
	}

	// Construct an "options" struct for creating the tekton pipeline service client.
	var cdTektonPipelineClientURL string
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		cdTektonPipelineClientURL, err = cdtektonpipelinev2.GetServiceURLForRegion("private." + c.Region)
		if err != nil && c.Visibility == "public-and-private" {
			cdTektonPipelineClientURL, err = cdtektonpipelinev2.GetServiceURLForRegion(c.Region)
		}
	} else {
		cdTektonPipelineClientURL, err = cdtektonpipelinev2.GetServiceURLForRegion(c.Region)
	}
	if err != nil {
		cdTektonPipelineClientURL = cdtektonpipelinev2.DefaultServiceURL
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		cdTektonPipelineClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_TEKTON_PIPELINE_ENDPOINT", c.Region, cdTektonPipelineClientURL)
	}
	cdTektonPipelineClientOptions := &cdtektonpipelinev2.CdTektonPipelineV2Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_TEKTON_PIPELINE_ENDPOINT"}, cdTektonPipelineClientURL),
	}
	// Construct the service client.
	session.cdTektonPipelineClient, err = cdtektonpipelinev2.NewCdTektonPipelineV2(cdTektonPipelineClientOptions)
	if err == nil {
		// Enable retries for API calls
		session.cdTektonPipelineClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.cdTektonPipelineClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm-continuous-delivery/%s", version.Version)},
		})
	} else {
		session.cdTektonPipelineClientErr = fmt.Errorf("Error occurred while configuring CD Tekton Pipeline service: %q", err)
	}

	if os.Getenv("TF_LOG") != "" {
		logDestination := log.Writer()
		goLogger := log.New(logDestination, "", log.LstdFlags)
		core.SetLogger(core.NewLogger(core.LevelDebug, goLogger, goLogger))
	}
	return session, nil
}

// CreateVersionDate requires mandatory version attribute. Any date from 2019-12-13 up to the currentdate may be provided. Specify the current date to request the latest version.
func CreateVersionDate() *string {
	version := time.Now().Format("2006-01-02")
	return &version
}

func newSession(c *Config) (*Session, error) {
	ibmSession := &Session{}

	softlayerSession := &slsession.Session{
		Endpoint:  c.SoftLayerEndpointURL,
		Timeout:   c.SoftLayerTimeout,
		UserName:  c.SoftLayerUserName,
		APIKey:    c.SoftLayerAPIKey,
		Debug:     os.Getenv("TF_LOG") != "",
		Retries:   c.RetryCount,
		RetryWait: c.RetryDelay,
	}

	if c.IAMToken != "" {
		log.Println("Configuring SoftLayer Session with token")
		softlayerSession.IAMToken = c.IAMToken
		softlayerSession.IAMRefreshToken = c.IAMRefreshToken
	}
	if c.SoftLayerAPIKey != "" && c.SoftLayerUserName != "" {
		log.Println("Configuring SoftLayer Session with API key")
		softlayerSession.APIKey = c.SoftLayerAPIKey
		softlayerSession.UserName = c.SoftLayerUserName
	}
	softlayerSession.AppendUserAgent(fmt.Sprintf("terraform-provider-ibm-continuous-delivery/%s", version.Version))
	ibmSession.SoftLayerSession = softlayerSession

	if c.IAMTrustedProfileID == "" && (c.IAMToken != "" && c.IAMRefreshToken == "") || (c.IAMToken == "" && c.IAMRefreshToken != "") {
		return nil, fmt.Errorf("iam_token and iam_refresh_token must be provided")
	}
	if c.IAMTrustedProfileID != "" && c.IAMToken == "" {
		return nil, fmt.Errorf("iam_token and iam_profile_id must be provided")
	}

	if c.IAMToken != "" {
		log.Println("Configuring IBM Cloud Session with token")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			IAMAccessToken:  c.IAMToken,
			IAMRefreshToken: c.IAMRefreshToken,
			//Comment out debug mode for v0.12
			Debug:         os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
			Visibility:    c.Visibility,
			EndpointsFile: c.EndpointsFile,
			UserAgent:     fmt.Sprintf("terraform-provider-ibm-continuous-delivery/%s", version.Version),
		}
		sess, err := bxsession.New(bmxConfig)
		if err != nil {
			return nil, err
		}
		ibmSession.BluemixSession = sess
	}

	if c.BluemixAPIKey != "" {
		log.Println("Configuring IBM Cloud Session with API key")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			BluemixAPIKey: c.BluemixAPIKey,
			//Comment out debug mode for v0.12
			Debug:         os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
			Visibility:    c.Visibility,
			EndpointsFile: c.EndpointsFile,
			UserAgent:     fmt.Sprintf("terraform-provider-ibm-continuous-delivery/%s", version.Version),
		}
		sess, err := bxsession.New(bmxConfig)
		if err != nil {
			return nil, err
		}
		ibmSession.BluemixSession = sess
	}

	return ibmSession, nil
}

func authenticateAPIKey(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent":            []string{http.UserAgent()},
			"X-Original-User-Agent": []string{config.UserAgent},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func authenticateCF(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewUAARepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent":            []string{http.UserAgent()},
			"X-Original-User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func fetchUserDetails(sess *bxsession.Session, retries int, retryDelay time.Duration) (*UserConfig, error) {
	config := sess.Config
	user := UserConfig{}
	var bluemixToken string

	if strings.HasPrefix(config.IAMAccessToken, "Bearer") {
		bluemixToken = config.IAMAccessToken[7:len(config.IAMAccessToken)]
	} else {
		bluemixToken = config.IAMAccessToken
	}

	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		if retries > 0 {
			if config.BluemixAPIKey != "" {
				time.Sleep(retryDelay)
				log.Printf("Retrying authentication for user details %d", retries)
				_ = authenticateAPIKey(sess)
				return fetchUserDetails(sess, retries-1, retryDelay)
			}
		}
		return &user, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.UserEmail = email.(string)
	}
	user.UserID = claims["id"].(string)
	user.UserAccount = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.CloudName = "bluemix"
	} else {
		user.CloudName = "staging"
	}
	user.cloudType = "public"

	user.generation = 2
	return &user, nil
}

func RefreshToken(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent":            []string{http.UserAgent()},
			"X-Original-User-Agent": []string{config.UserAgent},
		},
	})
	if err != nil {
		return err
	}
	_, err = tokenRefresher.RefreshToken()
	return err
}

func EnvFallBack(envs []string, defaultValue string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return defaultValue
}
func fileFallBack(fileMap map[string]interface{}, visibility, key, region, defaultValue string) string {
	if val, ok := fileMap[key]; ok {
		if v, ok := val.(map[string]interface{})[visibility]; ok {
			if r, ok := v.(map[string]interface{})[region]; ok && r.(string) != "" {
				return r.(string)
			}
		}
	}
	return defaultValue
}

// DefaultTransport ...
func DefaultTransport() gohttp.RoundTripper {
	transport := &gohttp.Transport{
		Proxy:               gohttp.ProxyFromEnvironment,
		DisableKeepAlives:   true,
		MaxIdleConnsPerHost: -1,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	return transport
}

func isRetryable(err error) bool {
	if bmErr, ok := err.(bmxerror.RequestFailure); ok {
		switch bmErr.StatusCode() {
		case 408, 504, 599, 429, 500, 502, 520, 503:
			return true
		}
	}

	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}

	if netErr, ok := err.(*net.OpError); ok && netErr.Timeout() {
		return true
	}

	if netErr, ok := err.(net.UnknownNetworkError); ok && netErr.Timeout() {
		return true
	}

	return false
}

func ContructEndpoint(subdomain, domain string) string {
	endpoint := fmt.Sprintf("https://%s.%s", subdomain, domain)
	return endpoint
}
