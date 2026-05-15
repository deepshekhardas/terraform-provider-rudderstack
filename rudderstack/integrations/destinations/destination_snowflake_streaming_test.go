package destinations_test

import (
	"testing"

	cmt "github.com/rudderlabs/terraform-provider-rudderstack/internal/testutil/cm"
	c "github.com/rudderlabs/terraform-provider-rudderstack/rudderstack/configs"
)

func TestDestinationResourceSnowflakeStreaming(t *testing.T) {
	cmt.AssertDestination(t, "snowflake_streaming", []c.TestConfig{
		{
			TerraformCreate: `
				account = "example-account"
				database = "example-database"
				warehouse = "example-warehouse"
				user = "example-user"
				private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----"
				namespace = "example-namespace"
			`,
			APICreate: `{
				"account": "example-account",
				"database": "example-database",
				"warehouse": "example-warehouse",
				"user": "example-user",
				"privateKey": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----",
				"namespace": "example-namespace",
				"skipTracksTable": false,
				"underscoreDivideNumbers": false,
				"allowUsersContextTraits": false
			}`,
			TerraformUpdate: `
				account = "example-account"
				database = "example-database"
				warehouse = "example-warehouse"
				user = "example-user"
				private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----"
				private_key_passphrase = "example-passphrase"
				namespace = "example-namespace"
				role = "example-role"
				skip_tracks_table = true
				json_paths = "event.properties.key1,event.properties.key2"
				underscore_divide_numbers = true
				allow_users_context_traits = true
			`,
			APIUpdate: `{
				"account": "example-account",
				"database": "example-database",
				"warehouse": "example-warehouse",
				"user": "example-user",
				"privateKey": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----",
				"privateKeyPassphrase": "example-passphrase",
				"namespace": "example-namespace",
				"role": "example-role",
				"skipTracksTable": true,
				"jsonPaths": "event.properties.key1,event.properties.key2",
				"underscoreDivideNumbers": true,
				"allowUsersContextTraits": true
			}`,
		},
	})
}

func TestDestinationResourceSnowflakeStreamingWithIceberg(t *testing.T) {
	cmt.AssertDestination(t, "snowflake_streaming", []c.TestConfig{
		{
			TerraformCreate: `
				account = "example-account"
				database = "example-database"
				warehouse = "example-warehouse"
				user = "example-user"
				private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----"
				namespace = "example-namespace"
				enable_iceberg = true
				external_volume = "EXTERNAL_VOLUME"
			`,
			APICreate: `{
				"account": "example-account",
				"database": "example-database",
				"warehouse": "example-warehouse",
				"user": "example-user",
				"privateKey": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----",
				"namespace": "example-namespace",
				"enableIceberg": true,
				"externalVolume": "EXTERNAL_VOLUME",
				"skipTracksTable": false,
				"underscoreDivideNumbers": false,
				"allowUsersContextTraits": false
			}`,
		},
	})
}

func TestDestinationResourceSnowflakeStreamingWithConnectionMode(t *testing.T) {
	cmt.AssertDestination(t, "snowflake_streaming", []c.TestConfig{
		{
			TerraformCreate: `
				account = "example-account"
				database = "example-database"
				warehouse = "example-warehouse"
				user = "example-user"
				private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----"
				namespace = "example-namespace"
				connection_mode {
					web = "cloud"
					android = "cloud"
					ios = "cloud"
				}
			`,
			APICreate: `{
				"account": "example-account",
				"database": "example-database",
				"warehouse": "example-warehouse",
				"user": "example-user",
				"privateKey": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----",
				"namespace": "example-namespace",
				"connectionMode": {
					"web": "cloud",
					"android": "cloud",
					"ios": "cloud"
				},
				"skipTracksTable": false,
				"underscoreDivideNumbers": false,
				"allowUsersContextTraits": false
			}`,
		},
	})
}

func TestDestinationResourceSnowflakeStreamingWithConsentManagement(t *testing.T) {
	cmt.AssertDestination(t, "snowflake_streaming", []c.TestConfig{
		{
			TerraformCreate: `
				account = "example-account"
				database = "example-database"
				warehouse = "example-warehouse"
				user = "example-user"
				private_key = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----"
				namespace = "example-namespace"
				consent_management {
					web = [{
						provider = "onetrust"
						resolution_strategy = "and"
						consents = ["C0001", "C0002"]
					}]
				}
			`,
			APICreate: `{
				"account": "example-account",
				"database": "example-database",
				"warehouse": "example-warehouse",
				"user": "example-user",
				"privateKey": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASC...\n-----END PRIVATE KEY-----",
				"namespace": "example-namespace",
				"consentManagement": {
					"web": [{
						"provider": "onetrust",
						"resolutionStrategy": "and",
						"consents": ["C0001", "C0002"]
					}]
				},
				"skipTracksTable": false,
				"underscoreDivideNumbers": false,
				"allowUsersContextTraits": false
			}`,
		},
	})
}
