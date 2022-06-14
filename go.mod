module github.com/IBM-Cloud/terraform-provider-ibm-continuous-delivery

go 1.16

require (
	github.com/IBM-Cloud/bluemix-go v0.0.0-20220523145737-34645883de47
	github.com/IBM/continuous-delivery-go-sdk v0.0.5
	github.com/IBM/go-sdk-core/v5 v5.10.1
	github.com/IBM/platform-services-go-sdk v0.25.1
	github.com/apache/openwhisk-client-go v0.0.0-20200201143223-a804fb82d105
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-openapi/errors v0.20.2 // indirect
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/go-test/deep v1.0.4 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/go-cmp v0.5.8
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-version v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.16.0
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nicksnyder/go-i18n v1.10.0 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/softlayer/softlayer-go v1.0.3
	go.mongodb.org/mongo-driver v1.8.3 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/softlayer/softlayer-go v1.0.3 => github.com/IBM-Cloud/softlayer-go v1.0.5-tf

replace github.com/dgrijalva/jwt-go v3.2.0+incompatible => github.com/golang-jwt/jwt v3.2.1+incompatible
