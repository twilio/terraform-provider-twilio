
## twilio_accounts_credentials_aws_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**credentials** | string | **Required** | A string that contains the AWS access credentials in the format `<AWS_ACCESS_KEY_ID>:<AWS_SECRET_ACCESS_KEY>`. For example, `AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**account_sid** | string | Optional | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the AWS resource to update.

## twilio_accounts_credentials_public_keys_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**public_key** | string | **Required** | A URL encoded representation of the public key. For example, `-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----`
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**account_sid** | string | Optional | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the PublicKey resource to update.

