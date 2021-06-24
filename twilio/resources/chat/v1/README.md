
## twilio_chat_services_channels_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**type** | string | Optional | The visibility of the channel. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Channel resource to update.

## twilio_chat_credentials_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | The type of push-notification service the credential is for. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;.
**api_key** | string | Optional | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**certificate** | string | Optional | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**private_key** | string | Optional | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----&#x60;
**sandbox** | bool | Optional | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**secret** | string | Optional | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Credential resource to update.

## twilio_chat_services_channels_members_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**channel_sid** | string | **Required** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new member belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services).
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Member resource to update.
**last_consumed_message_index** | int | Optional | The index of the last [Message](https://www.twilio.com/docs/api/chat/rest/messages) that the Member has read within the [Channel](https://www.twilio.com/docs/api/chat/rest/channels).

## twilio_chat_services_channels_messages_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**channel_sid** | string | **Required** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**body** | string | **Required** | The message to send to the channel. Can also be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**from** | string | Optional | The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message&#39;s author. The default value is &#x60;system&#x60;.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Message resource to update.

## twilio_chat_services_roles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**permission** | list(string) | **Required** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60; and are described in the documentation.
**type** | string | **Required** | The type of role. Can be: &#x60;channel&#x60; for [Channel](https://www.twilio.com/docs/chat/api/channels) roles or &#x60;deployment&#x60; for [Service](https://www.twilio.com/docs/chat/api/services) roles.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Role resource to update.

## twilio_chat_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.
**consumption_report_interval** | int | Optional | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
**default_channel_creator_role_sid** | string | Optional | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**default_channel_role_sid** | string | Optional | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**default_service_role_sid** | string | Optional | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**limits_channel_members** | int | Optional | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
**limits_user_channels** | int | Optional | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
**notifications_added_to_channel_enabled** | bool | Optional | Whether to send a notification when a member is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**notifications_added_to_channel_template** | string | Optional | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_invited_to_channel_enabled** | bool | Optional | Whether to send a notification when a user is invited to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**notifications_invited_to_channel_template** | string | Optional | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_new_message_enabled** | bool | Optional | Whether to send a notification when a new message is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**notifications_new_message_template** | string | Optional | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
**notifications_removed_from_channel_enabled** | bool | Optional | Whether to send a notification to a user when they are removed from a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**notifications_removed_from_channel_template** | string | Optional | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
**post_webhook_url** | string | Optional | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**pre_webhook_url** | string | Optional | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**reachability_enabled** | bool | Optional | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;.
**read_status_enabled** | bool | Optional | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;.
**typing_indicator_timeout** | int | Optional | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds.
**webhook_filters** | list(string) | Optional | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**webhook_method** | string | Optional | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**webhooks_on_channel_add_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_channel_add.url&#x60;.
**webhooks_on_channel_add_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_add&#x60; event using the &#x60;webhooks.on_channel_add.method&#x60; HTTP method.
**webhooks_on_channel_added_method** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event&#x60;.
**webhooks_on_channel_added_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_added.method&#x60; HTTP method.
**webhooks_on_channel_destroy_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroy.url&#x60;.
**webhooks_on_channel_destroy_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_destroy&#x60; event using the &#x60;webhooks.on_channel_destroy.method&#x60; HTTP method.
**webhooks_on_channel_destroyed_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroyed.url&#x60;.
**webhooks_on_channel_destroyed_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_destroyed.method&#x60; HTTP method.
**webhooks_on_channel_update_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_channel_update.url&#x60;.
**webhooks_on_channel_update_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_update&#x60; event using the &#x60;webhooks.on_channel_update.method&#x60; HTTP method.
**webhooks_on_channel_updated_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;.
**webhooks_on_channel_updated_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method.
**webhooks_on_member_add_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_member_add.url&#x60;.
**webhooks_on_member_add_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_member_add&#x60; event using the &#x60;webhooks.on_member_add.method&#x60; HTTP method.
**webhooks_on_member_added_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;.
**webhooks_on_member_added_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method.
**webhooks_on_member_remove_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_member_remove.url&#x60;.
**webhooks_on_member_remove_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_member_remove&#x60; event using the &#x60;webhooks.on_member_remove.method&#x60; HTTP method.
**webhooks_on_member_removed_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_member_removed.url&#x60;.
**webhooks_on_member_removed_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_member_removed&#x60; event using the &#x60;webhooks.on_member_removed.method&#x60; HTTP method.
**webhooks_on_message_remove_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_message_remove.url&#x60;.
**webhooks_on_message_remove_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_remove&#x60; event using the &#x60;webhooks.on_message_remove.method&#x60; HTTP method.
**webhooks_on_message_removed_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_message_removed.url&#x60;.
**webhooks_on_message_removed_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_removed&#x60; event using the &#x60;webhooks.on_message_removed.method&#x60; HTTP method.
**webhooks_on_message_send_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_message_send.url&#x60;.
**webhooks_on_message_send_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_send&#x60; event using the &#x60;webhooks.on_message_send.method&#x60; HTTP method.
**webhooks_on_message_sent_method** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event&#x60;.
**webhooks_on_message_sent_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event using the &#x60;webhooks.on_message_sent.method&#x60; HTTP method.
**webhooks_on_message_update_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_message_update.url&#x60;.
**webhooks_on_message_update_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_update&#x60; event using the &#x60;webhooks.on_message_update.method&#x60; HTTP method.
**webhooks_on_message_updated_method** | string | Optional | The HTTP method to use when calling the &#x60;webhooks.on_message_updated.url&#x60;.
**webhooks_on_message_updated_url** | string | Optional | The URL of the webhook to call in response to the &#x60;on_message_updated&#x60; event using the &#x60;webhooks.on_message_updated.method&#x60; HTTP method.

## twilio_chat_services_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details.
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. This value is often used for display purposes.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the User resource to update.

