
## twilio_chat_services_channels_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. This value must be 64 characters or less in length and be unique within the Service.
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**type** | string | Optional | 
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Channel resource to update.

## twilio_chat_credentials_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | 
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**certificate** | string | Optional | [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A== -----END CERTIFICATE-----`
**private_key** | string | Optional | [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----`
**sandbox** | bool | Optional | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production.
**api_key** | string | Optional | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**secret** | string | Optional | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Credential resource to update.

## twilio_chat_services_channels_invites_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**channel_sid** | string | **Required** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to.
**identity** | string | **Required** | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new member.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Invite resource to fetch.

## twilio_chat_services_channels_members_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**channel_sid** | string | **Required** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new member belongs to. Can be the Channel resource's `sid` or `unique_name`.
**identity** | string | **Required** | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services).
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Member resource to update.
**last_consumed_message_index** | int | Optional | The index of the last [Message](https://www.twilio.com/docs/api/chat/rest/messages) that the Member has read within the [Channel](https://www.twilio.com/docs/api/chat/rest/channels).

## twilio_chat_services_channels_messages_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**channel_sid** | string | **Required** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to. Can be the Channel resource's `sid` or `unique_name`.
**body** | string | **Required** | The message to send to the channel. Can also be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
**from** | string | Optional | The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message's author. The default value is `system`.
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Message resource to update.

## twilio_chat_services_roles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**type** | string | **Required** | 
**permission** | list(string) | **Required** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type` and are described in the documentation.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Role resource to update.

## twilio_chat_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Service resource to update.
**default_service_role_sid** | string | Optional | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**default_channel_role_sid** | string | Optional | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**default_channel_creator_role_sid** | string | Optional | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**read_status_enabled** | bool | Optional | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is `true`.
**reachability_enabled** | bool | Optional | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is `false`.
**typing_indicator_timeout** | int | Optional | How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds.
**consumption_report_interval** | int | Optional | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
**notifications_new_message_enabled** | bool | Optional | Whether to send a notification when a new message is added to a channel. Can be: `true` or `false` and the default is `false`.
**notifications_new_message_template** | string | Optional | The template to use to create the notification text displayed when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
**notifications_added_to_channel_enabled** | bool | Optional | Whether to send a notification when a member is added to a channel. Can be: `true` or `false` and the default is `false`.
**notifications_added_to_channel_template** | string | Optional | The template to use to create the notification text displayed when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
**notifications_removed_from_channel_enabled** | bool | Optional | Whether to send a notification to a user when they are removed from a channel. Can be: `true` or `false` and the default is `false`.
**notifications_removed_from_channel_template** | string | Optional | The template to use to create the notification text displayed to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
**notifications_invited_to_channel_enabled** | bool | Optional | Whether to send a notification when a user is invited to a channel. Can be: `true` or `false` and the default is `false`.
**notifications_invited_to_channel_template** | string | Optional | The template to use to create the notification text displayed when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
**pre_webhook_url** | string | Optional | The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**post_webhook_url** | string | Optional | The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**webhook_method** | string | Optional | The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**webhook_filters** | list(string) | Optional | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**webhooks_on_message_send_url** | string | Optional | The URL of the webhook to call in response to the `on_message_send` event using the `webhooks.on_message_send.method` HTTP method.
**webhooks_on_message_send_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_message_send.url`.
**webhooks_on_message_update_url** | string | Optional | The URL of the webhook to call in response to the `on_message_update` event using the `webhooks.on_message_update.method` HTTP method.
**webhooks_on_message_update_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_message_update.url`.
**webhooks_on_message_remove_url** | string | Optional | The URL of the webhook to call in response to the `on_message_remove` event using the `webhooks.on_message_remove.method` HTTP method.
**webhooks_on_message_remove_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_message_remove.url`.
**webhooks_on_channel_add_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_add` event using the `webhooks.on_channel_add.method` HTTP method.
**webhooks_on_channel_add_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_channel_add.url`.
**webhooks_on_channel_destroy_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_destroy` event using the `webhooks.on_channel_destroy.method` HTTP method.
**webhooks_on_channel_destroy_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_channel_destroy.url`.
**webhooks_on_channel_update_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_update` event using the `webhooks.on_channel_update.method` HTTP method.
**webhooks_on_channel_update_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_channel_update.url`.
**webhooks_on_member_add_url** | string | Optional | The URL of the webhook to call in response to the `on_member_add` event using the `webhooks.on_member_add.method` HTTP method.
**webhooks_on_member_add_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_member_add.url`.
**webhooks_on_member_remove_url** | string | Optional | The URL of the webhook to call in response to the `on_member_remove` event using the `webhooks.on_member_remove.method` HTTP method.
**webhooks_on_member_remove_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_member_remove.url`.
**webhooks_on_message_sent_url** | string | Optional | The URL of the webhook to call in response to the `on_message_sent` event using the `webhooks.on_message_sent.method` HTTP method.
**webhooks_on_message_sent_method** | string | Optional | The URL of the webhook to call in response to the `on_message_sent` event`.
**webhooks_on_message_updated_url** | string | Optional | The URL of the webhook to call in response to the `on_message_updated` event using the `webhooks.on_message_updated.method` HTTP method.
**webhooks_on_message_updated_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_message_updated.url`.
**webhooks_on_message_removed_url** | string | Optional | The URL of the webhook to call in response to the `on_message_removed` event using the `webhooks.on_message_removed.method` HTTP method.
**webhooks_on_message_removed_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_message_removed.url`.
**webhooks_on_channel_added_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_added.method` HTTP method.
**webhooks_on_channel_added_method** | string | Optional | The URL of the webhook to call in response to the `on_channel_added` event`.
**webhooks_on_channel_destroyed_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_destroyed.method` HTTP method.
**webhooks_on_channel_destroyed_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_channel_destroyed.url`.
**webhooks_on_channel_updated_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method.
**webhooks_on_channel_updated_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_channel_updated.url`.
**webhooks_on_member_added_url** | string | Optional | The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method.
**webhooks_on_member_added_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_channel_updated.url`.
**webhooks_on_member_removed_url** | string | Optional | The URL of the webhook to call in response to the `on_member_removed` event using the `webhooks.on_member_removed.method` HTTP method.
**webhooks_on_member_removed_method** | string | Optional | The HTTP method to use when calling the `webhooks.on_member_removed.url`.
**limits_channel_members** | int | Optional | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
**limits_user_channels** | int | Optional | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.

## twilio_chat_services_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**identity** | string | **Required** | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User.
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. This value is often used for display purposes.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the User resource to update.

