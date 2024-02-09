
## twilio_notify_services_bindings_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to create the resource under.
**identity** | string | **Required** | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Up to 20 Bindings can be created for the same Identity in a given Service.
**binding_type** | string | **Required** | 
**address** | string | **Required** | The channel-specific address. For APNS, the device token. For FCM and GCM, the registration token. For SMS, a phone number in E.164 format. For Facebook Messenger, the Messenger ID of the user or a phone number in E.164 format.
**tag** | list(string) | Optional | A tag that can be used to select the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 20 tags.
**notification_protocol_version** | string | Optional | The protocol version to use to send the notification. This defaults to the value of `default_xxxx_notification_protocol_version` for the protocol in the [Service](https://www.twilio.com/docs/notify/api/service-resource). The current version is `\\\"3\\\"` for `apn`, `fcm`, and `gcm` type Bindings. The parameter is not applicable to `sms` and `facebook-messenger` type Bindings as the data format is fixed.
**credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) resource to be used to send notifications to this Binding. If present, this overrides the Credential specified in the Service resource. Applies to only `apn`, `fcm`, and `gcm` type Bindings.
**endpoint** | string | Optional | Deprecated.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Binding resource to fetch.

## twilio_notify_credentials_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | 
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**certificate** | string | Optional | [APN only] The URL-encoded representation of the certificate. Strip everything outside of the headers, e.g. `-----BEGIN CERTIFICATE-----MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A==-----END CERTIFICATE-----`
**private_key** | string | Optional | [APN only] The URL-encoded representation of the private key. Strip everything outside of the headers, e.g. `-----BEGIN RSA PRIVATE KEY-----MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR\\\\n.-----END RSA PRIVATE KEY-----`
**sandbox** | bool | Optional | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production.
**api_key** | string | Optional | [GCM only] The `Server key` of your project from Firebase console under Settings / Cloud messaging.
**secret** | string | Optional | [FCM only] The `Server key` of your project from Firebase console under Settings / Cloud messaging.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Credential resource to update.

## twilio_notify_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**apn_credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings.
**gcm_credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings.
**messaging_service_sid** | string | Optional | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/quickstart#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications.
**facebook_messenger_page_id** | string | Optional | Deprecated.
**default_apn_notification_protocol_version** | string | Optional | The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
**default_gcm_notification_protocol_version** | string | Optional | The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
**fcm_credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings.
**default_fcm_notification_protocol_version** | string | Optional | The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
**log_enabled** | bool | Optional | Whether to log notifications. Can be: `true` or `false` and the default is `true`.
**alexa_skill_id** | string | Optional | Deprecated.
**default_alexa_notification_protocol_version** | string | Optional | Deprecated.
**delivery_callback_url** | string | Optional | URL to send delivery status callback.
**delivery_callback_enabled** | bool | Optional | Callback configuration that enables delivery callbacks, default false
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.

