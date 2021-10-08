
## twilio_notify_services_bindings_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to create the resource under.
**address** | string | **Required** | The channel-specific address. For APNS, the device token. For FCM and GCM, the registration token. For SMS, a phone number in E.164 format. For Facebook Messenger, the Messenger ID of the user or a phone number in E.164 format.
**binding_type** | string | **Required** | The transport technology to use for the Binding. Can be: &#x60;apn&#x60;, &#x60;fcm&#x60;, &#x60;gcm&#x60;, &#x60;sms&#x60;, or &#x60;facebook-messenger&#x60;.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Up to 20 Bindings can be created for the same Identity in a given Service.
**credential_sid** | string | Optional | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) resource to be used to send notifications to this Binding. If present, this overrides the Credential specified in the Service resource. Applies to only &#x60;apn&#x60;, &#x60;fcm&#x60;, and &#x60;gcm&#x60; type Bindings.
**endpoint** | string | Optional | Deprecated.
**notification_protocol_version** | string | Optional | The protocol version to use to send the notification. This defaults to the value of &#x60;default_xxxx_notification_protocol_version&#x60; for the protocol in the [Service](https://www.twilio.com/docs/notify/api/service-resource). The current version is &#x60;\\\&quot;3\\\&quot;&#x60; for &#x60;apn&#x60;, &#x60;fcm&#x60;, and &#x60;gcm&#x60; type Bindings. The parameter is not applicable to &#x60;sms&#x60; and &#x60;facebook-messenger&#x60; type Bindings as the data format is fixed.
**tag** | list(string) | Optional | A tag that can be used to select the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 20 tags.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Binding resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of related resources
**tags** | list(string) | *Computed* | The list of tags associated with this Binding
**url** | string | *Computed* | The absolute URL of the Binding resource

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
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the Credential resource

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
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of the resources related to the service
**url** | string | *Computed* | The absolute URL of the Service resource

