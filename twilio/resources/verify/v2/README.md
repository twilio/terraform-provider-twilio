
## twilio_verify_services_rate_limits_buckets_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**rate_limit_sid** | string | **Required** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**interval** | int | **Required** | Number of seconds that the rate limit will be enforced over.
**max** | int | **Required** | Maximum number of requests permitted in during the interval.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Bucket.

## twilio_verify_services_messaging_configurations_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**country** | string | **Required** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;.
**messaging_service_sid** | string | **Required** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.

## twilio_verify_services_entities_factors_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The unique SID identifier of the Service.
**identity** | string | **Required** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**factor_type** | string | **Required** | The Type of this Factor. Currently &#x60;push&#x60; and &#x60;totp&#x60; are supported. For &#x60;totp&#x60; to work, you need to contact [Twilio sales](https://www.twilio.com/help/sales) first to have the Verify TOTP feature enabled for your Twilio account.
**friendly_name** | string | **Required** | The friendly name of this Factor. This can be any string up to 64 characters, meant for humans to distinguish between Factors. For &#x60;factor_type&#x60; &#x60;push&#x60;, this could be a device name. For &#x60;factor_type&#x60; &#x60;totp&#x60;, this value is used as the “account name” in constructing the &#x60;binding.uri&#x60; property. At the same time, we recommend avoiding providing PII.
**binding_alg** | string | Optional | The algorithm used when &#x60;factor_type&#x60; is &#x60;push&#x60;. Algorithm supported: &#x60;ES256&#x60;
**binding_public_key** | string | Optional | The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;
**binding_secret** | string | Optional | The shared secret for TOTP factors encoded in Base32. This can be provided when creating the Factor, otherwise it will be generated.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**config_alg** | string | Optional | The algorithm used to derive the TOTP codes. Can be &#x60;sha1&#x60;, &#x60;sha256&#x60; or &#x60;sha512&#x60;. Defaults to &#x60;sha1&#x60;.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**config_app_id** | string | Optional | The ID that uniquely identifies your app in the Google or Apple store, such as &#x60;com.example.myapp&#x60;. It can be up to 100 characters long.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;.
**config_code_length** | int | Optional | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. The default value is defined at the service level in the property &#x60;totp.code_length&#x60;. If not configured defaults to 6.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**config_notification_platform** | string | Optional | The transport technology used to generate the Notification Token. Can be &#x60;apn&#x60; or &#x60;fcm&#x60;.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;.
**config_notification_token** | string | Optional | For APN, the device token. For FCM the registration token. It used to send the push notifications. Must be between 32 and 255 characters long.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;.
**config_sdk_version** | string | Optional | The Verify Push SDK version used to configure the factor  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;
**config_skew** | int | Optional | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. The default value is defined at the service level in the property &#x60;totp.skew&#x60;. If not configured defaults to 1.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**config_time_step** | int | Optional | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. The default value is defined at the service level in the property &#x60;totp.time_step&#x60;. Defaults to 30 seconds if not configured.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**sid** | string | *Computed* | A 34 character string that uniquely identifies this Factor.
**auth_payload** | string | Optional | The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code.

## twilio_verify_services_rate_limits_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**unique_name** | string | **Required** | Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**
**description** | string | Optional | Description of this Rate Limit
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

## twilio_verify_services_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.**
**code_length** | int | Optional | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
**custom_code_enabled** | bool | Optional | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
**do_not_share_warning_enabled** | bool | Optional | Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: &#x60;Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code&#x60;
**dtmf_input_required** | bool | Optional | Whether to ask the user to press a number before delivering the verify code in a phone call.
**lookup_enabled** | bool | Optional | Whether to perform a lookup with each verification started and return info about the phone number.
**psd2_enabled** | bool | Optional | Whether to pass PSD2 transaction parameters when starting a verification.
**push_apn_credential_sid** | string | Optional | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
**push_fcm_credential_sid** | string | Optional | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
**push_include_date** | bool | Optional | Optional configuration for the Push factors. If true, include the date in the Challenge&#39;s reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true
**skip_sms_to_landlines** | bool | Optional | Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;.
**totp_code_length** | int | Optional | Optional configuration for the TOTP factors. Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6
**totp_issuer** | string | Optional | Optional configuration for the TOTP factors. Set TOTP Issuer for this service. This will allow to configure the issuer of the TOTP URI. Defaults to the service friendly name if not provided.
**totp_skew** | int | Optional | Optional configuration for the TOTP factors. The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1
**totp_time_step** | int | Optional | Optional configuration for the TOTP factors. Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds
**tts_name** | string | Optional | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.

## twilio_verify_services_webhooks_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The unique SID identifier of the Service.
**event_types** | list(string) | **Required** | The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60;
**friendly_name** | string | **Required** | The string that you assigned to describe the webhook. **This value should not contain PII.**
**webhook_url** | string | **Required** | The URL associated with this Webhook.
**status** | string | Optional | The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60;
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Webhook resource to update.
