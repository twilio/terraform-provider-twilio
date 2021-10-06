
## twilio_chat_services_channels_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Channel resource under.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**created_by** | string | Optional | The &#x60;identity&#x60; of the User that created the channel. Default is: &#x60;system&#x60;.
**date_created** | string | Optional | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this should only be used in cases where a Channel is being recreated from a backup/separate source.
**date_updated** | string | Optional | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is &#x60;null&#x60;. Note that this parameter should only be used in cases where a Channel is being recreated from a backup/separate source  and where a Message was previously updated.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**type** | string | Optional | The visibility of the channel. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the Channel resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service.
**sid** | string | *Computed* | The SID of the Channel resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to update.

## twilio_chat_services_channels_webhooks_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to create the Webhook resource under.
**channel_sid** | string | **Required** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Channel Webhook resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**type** | string | **Required** | The type of webhook. Can be: &#x60;webhook&#x60;, &#x60;studio&#x60;, or &#x60;trigger&#x60;.
**configuration_filters** | list(string) | Optional | The events that cause us to call the Channel Webhook. Used when &#x60;type&#x60; is &#x60;webhook&#x60;. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger).
**configuration_flow_sid** | string | Optional | The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in &#x60;configuration.filters&#x60; occurs. Used only when &#x60;type&#x60; is &#x60;studio&#x60;.
**configuration_method** | string | Optional | The HTTP method used to call &#x60;configuration.url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**configuration_retry_count** | int | Optional | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0.
**configuration_triggers** | list(string) | Optional | A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when &#x60;type&#x60; &#x3D; &#x60;trigger&#x60;.
**configuration_url** | string | Optional | The URL of the webhook to call using the &#x60;configuration.method&#x60;.
**sid** | string | *Computed* | The SID of the Channel Webhook resource to update.

## twilio_chat_credentials_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | The type of push-notification service the credential is for. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;.
**api_key** | string | Optional | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**certificate** | string | Optional | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**private_key** | string | Optional | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;
**sandbox** | bool | Optional | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**secret** | string | Optional | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**sid** | string | *Computed* | The SID of the Credential resource to update.

## twilio_chat_services_channels_invites_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Invite resource under.
**channel_sid** | string | **Required** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Invite resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) assigned to the new member.
**sid** | string | *Computed* | The SID of the Invite resource to fetch.

## twilio_chat_services_channels_members_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Member resource under.
**channel_sid** | string | **Required** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Member resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**date_created** | string | Optional | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source.
**date_updated** | string | Optional | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is &#x60;null&#x60;. Note that this parameter should only be used when a Member is being recreated from a backup/separate source and where a Member was previously updated.
**last_consumed_message_index** | int | Optional | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. This parameter should only be used when recreating a Member from a backup/separate source.
**last_consumption_timestamp** | string | Optional | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource).
**sid** | string | *Computed* | The SID of the Member resource to update. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value.

## twilio_chat_services_channels_messages_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Message resource under.
**channel_sid** | string | **Required** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Message resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**body** | string | Optional | The message to send to the channel. Can be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
**date_created** | string | Optional | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat&#39;s history is being recreated from a backup/separate source.
**date_updated** | string | Optional | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
**from** | string | Optional | The [Identity](https://www.twilio.com/docs/chat/identity) of the new message&#39;s author. The default value is &#x60;system&#x60;.
**last_updated_by** | string | Optional | The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.
**media_sid** | string | Optional | The SID of the [Media](https://www.twilio.com/docs/chat/rest/media) to attach to the new Message.
**sid** | string | *Computed* | The SID of the Message resource to update.

## twilio_chat_services_roles_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Role resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**permission** | list(string) | **Required** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.
**type** | string | **Required** | The type of role. Can be: &#x60;channel&#x60; for [Channel](https://www.twilio.com/docs/chat/channels) roles or &#x60;deployment&#x60; for [Service](https://www.twilio.com/docs/chat/rest/service-resource) roles.
**sid** | string | *Computed* | The SID of the Role resource to update.

## twilio_chat_services_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new resource.
**sid** | string | *Computed* | The SID of the Service resource to update.
**consumption_report_interval** | int | Optional | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
**default_channel_creator_role_sid** | string | Optional | The channel role assigned to a channel creator when they join a new channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
**default_channel_role_sid** | string | Optional | The channel role assigned to users when they are added to a channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
**default_service_role_sid** | string | Optional | The service role assigned to users when they are added to the service. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
**limits_channel_members** | int | Optional | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
**limits_user_channels** | int | Optional | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
**media_compatibility_message** | string | Optional | The message to send when a media message has no text. Can be used as placeholder message.
**notifications_added_to_channel_enabled** | bool | Optional | Whether to send a notification when a member is added to a channel. The default is &#x60;false&#x60;.
**notifications_added_to_channel_sound** | string | Optional | The name of the sound to play when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_added_to_channel_template** | string | Optional | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_invited_to_channel_enabled** | bool | Optional | Whether to send a notification when a user is invited to a channel. The default is &#x60;false&#x60;.
**notifications_invited_to_channel_sound** | string | Optional | The name of the sound to play when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_invited_to_channel_template** | string | Optional | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_log_enabled** | bool | Optional | Whether to log notifications. The default is &#x60;false&#x60;.
**notifications_new_message_badge_count_enabled** | bool | Optional | Whether the new message badge is enabled. The default is &#x60;false&#x60;.
**notifications_new_message_enabled** | bool | Optional | Whether to send a notification when a new message is added to a channel. The default is &#x60;false&#x60;.
**notifications_new_message_sound** | string | Optional | The name of the sound to play when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
**notifications_new_message_template** | string | Optional | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
**notifications_removed_from_channel_enabled** | bool | Optional | Whether to send a notification to a user when they are removed from a channel. The default is &#x60;false&#x60;.
**notifications_removed_from_channel_sound** | string | Optional | The name of the sound to play to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
**notifications_removed_from_channel_template** | string | Optional | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
**post_webhook_retry_count** | int | Optional | The number of times to retry a call to the &#x60;post_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. The default is 0, which means the call won&#39;t be retried.
**post_webhook_url** | string | Optional | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**pre_webhook_retry_count** | int | Optional | The number of times to retry a call to the &#x60;pre_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. Default retry count is 0 times, which means the call won&#39;t be retried.
**pre_webhook_url** | string | Optional | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**reachability_enabled** | bool | Optional | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;.
**read_status_enabled** | bool | Optional | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;.
**typing_indicator_timeout** | int | Optional | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds.
**webhook_filters** | list(string) | Optional | The list of webhook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**webhook_method** | string | Optional | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.

## twilio_chat_services_users_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the User resource under.
**identity** | string | **Required** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). This value is often a username or email address. See the Identity documentation for more info.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**attributes** | string | Optional | A valid JSON string that contains application-specific data.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. This value is often used for display purposes.
**role_sid** | string | Optional | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the new User.
**sid** | string | *Computed* | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

