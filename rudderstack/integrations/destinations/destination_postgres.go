package destinations

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	c "github.com/rudderlabs/terraform-provider-rudderstack/rudderstack/configs"
)

func init() {
	supportedSourceTypes := []string{"web", "android", "ios", "unity", "reactnative", "flutter", "cordova", "amp", "cloud", "cloudSource", "shopify"}
	commonProperties, commonSchema := GetCommonConfigMeta(supportedSourceTypes)

	properties := []c.ConfigProperty{
		c.Simple("host", "host"),
		c.Simple("database", "database"),
		c.Simple("user", "user"),
		c.Simple("password", "password"),
		c.Simple("port", "port"),
		c.Simple("namespace", "namespace", c.SkipZeroValue),
		c.Simple("sslMode", "ssl_mode"),
		c.Simple("syncFrequency", "sync_frequency"),
		c.Simple("syncStartAt", "sync_start_at", c.SkipZeroValue),
		c.Simple("excludeWindow.excludeWindowStartTime", "exclude_window.0.exclude_window_start_time", c.SkipZeroValue),
		c.Simple("excludeWindow.excludeWindowEndTime", "exclude_window.0.exclude_window_end_time", c.SkipZeroValue),
		c.Simple("jsonPaths", "json_paths", c.SkipZeroValue),
		c.Simple("useRudderStorage", "use_rudder_storage", c.SkipZeroValue), // boolean
		c.Simple("bucketProvider", "bucket_provider", c.SkipZeroValue),
		c.Simple("bucketName", "bucket_name", c.SkipZeroValue),
		c.Simple("clientKey", "client_key", c.SkipZeroValue),
		c.Simple("clientCert", "client_cert", c.SkipZeroValue),
		c.Simple("serverCA", "server_ca", c.SkipZeroValue),
		c.Simple("roleBasedAuth", "role_based_auth", c.SkipZeroValue), // boolean
		c.Simple("iamRoleARN", "iam_role_arn", c.SkipZeroValue),
		c.Simple("accessKeyID", "access_key_id", c.SkipZeroValue),
		c.Simple("accessKey", "access_key", c.SkipZeroValue),
		c.Simple("accountName", "account_name", c.SkipZeroValue),
		c.Simple("accountKey", "account_key", c.SkipZeroValue),
		c.Simple("sasToken", "sas_token", c.SkipZeroValue),
		c.Simple("useSASTokens", "use_sas_tokens", c.SkipZeroValue), // boolean
		c.Simple("credentials", "credentials", c.SkipZeroValue),
		c.Simple("endPoint", "end_point", c.SkipZeroValue),
		c.Simple("secretAccessKey", "secret_access_key", c.SkipZeroValue),
		c.Simple("useSSL", "use_ssl", c.SkipZeroValue), // boolean
	}

	properties = append(properties, commonProperties...)

	schema := map[string]*schema.Schema{
		"host": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Enter the host name of your PostgreSQL database.",
		},
		"database": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Enter the name of your PostgreSQL database.",
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Enter the username of your PostgreSQL database.",
		},
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Enter the password of your PostgreSQL database.",
		},
		"port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "5432",
			Description: "Enter the port number of your PostgreSQL database.",
		},
		"namespace": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Enter the namespace of your PostgreSQL database.",
		},
		"ssl_mode": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "disable",
			Description:      "Enter the SSL mode of your PostgreSQL database.",
			ValidateDiagFunc: c.StringMatchesRegexp("^(disable|require|verify-ca)$"),
		},
		"sync_frequency": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "30",
			Description:      "Enter the frequency at which the data should be synced from your PostgreSQL database.",
			ValidateDiagFunc: c.StringMatchesRegexp("^(30|60|180|360|720|1440)$"),
		},
		"client_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Enter your Client Key Pem File",
		},
		"client_cert": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Enter your Client Cert Pem File",
		},
		"server_ca": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Enter your Server CA Pem File",
		},
		"use_rudder_storage": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Enable this setting to use RudderStack's data warehouse to store the data from your PostgreSQL database.",
		},
		"bucket_provider": {
			Type:             schema.TypeString,
			Optional:         true,
			Description:      "The cloud object storage provider to use when use_rudder_storage is disabled.",
			ValidateDiagFunc: c.StringMatchesRegexp("^(S3|GCS|AZURE_BLOB|MINIO)$"),
		},
		"bucket_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the object storage bucket.",
		},
		"access_key_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The AWS Access Key ID (for S3 bucket provider).",
		},
		"access_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The AWS Secret Access Key (for S3 bucket provider).",
		},
		"role_based_auth": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable this setting to use IAM role-based authentication.",
		},
		"iam_role_arn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The AWS IAM Role ARN (for S3 bucket provider).",
		},
		"account_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The Azure Blob Storage account name.",
		},
		"account_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The Azure Blob Storage account key.",
		},
		"credentials": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The GCS service account credentials JSON.",
		},
		"exclude_window": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Set a time window during which RudderStack will not sync data to PostgreSQL.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"exclude_window_start_time": {
						Type:             schema.TypeString,
						Optional:         true,
						ValidateDiagFunc: c.StringMatchesRegexp("^([01][0-9]|2[0-3]):[0-5][0-9]$"),
					},
					"exclude_window_end_time": {
						Type:             schema.TypeString,
						Optional:         true,
						ValidateDiagFunc: c.StringMatchesRegexp("^([01][0-9]|2[0-3]):[0-5][0-9]$"),
					},
				},
			},
		},
	}

	for key, value := range commonSchema {
		schema[key] = value
	}

	c.Destinations.Register("postgres", c.ConfigMeta{
		APIType:      "POSTGRES",
		Properties:   properties,
		ConfigSchema: schema,
	})
}
