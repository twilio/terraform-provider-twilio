
## twilio_notify_credentials_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | The Credential type. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;.
**api_key** | string | Optional | [GCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging.
**certificate** | string | Optional | [APN only] The URL-encoded representation of the certificate. Strip everything outside of the headers, e.g. &#x60;-----BEGIN CERTIFICATE-----MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D;-----END CERTIFICATE-----&#x60;
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**private_key** | string | Optional | [APN only] The URL-encoded representation of the private key. Strip everything outside of the headers, e.g. &#x60;-----BEGIN RSA PRIVATE KEY-----MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR\\\\n.-----END RSA PRIVATE KEY-----&#x60;
**sandbox** | bool | Optional | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**secret** | string | Optional | [FCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Credential resource to update.

## twilio_notify_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**alexa_skill_id** | string | Optional | Deprecated.
**apn_credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings.
**default_alexa_notification_protocol_version** | string | Optional | Deprecated.
**default_apn_notification_protocol_version** | string | Optional | The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
**default_fcm_notification_protocol_version** | string | Optional | The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
**default_gcm_notification_protocol_version** | string | Optional | The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
**delivery_callback_enabled** | bool | Optional | Callback configuration that enables delivery callbacks, default false
**delivery_callback_url** | string | Optional | URL to send delivery status callback.
**facebook_messenger_page_id** | string | Optional | Deprecated.
**fcm_credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**gcm_credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings.
**log_enabled** | bool | Optional | Whether to log notifications. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
**messaging_service_sid** | string | Optional | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/send-messages#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.

