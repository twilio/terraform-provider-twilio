
## twilio_conversations_configuration_addresses_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | 
**address** | string | **Required** | The unique address to be configured. The address can be a whatsapp address or phone number
**friendly_name** | string | Optional | The human-readable name of this configuration, limited to 256 characters. Optional.
**auto_creation_enabled** | bool | Optional | Enable/Disable auto-creating conversations for messages to this address
**auto_creation_type** | string | Optional | 
**auto_creation_conversation_service_sid** | string | Optional | Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
**auto_creation_webhook_url** | string | Optional | For type `webhook`, the url for the webhook request.
**auto_creation_webhook_method** | string | Optional | 
**auto_creation_webhook_filters** | list(string) | Optional | The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
**auto_creation_studio_flow_sid** | string | Optional | For type `studio`, the studio flow SID where the webhook should be sent to.
**auto_creation_studio_retry_count** | int | Optional | For type `studio`, number of times to retry the webhook request
**sid** | string | *Computed* | The SID of the Address Configuration resource. This value can be either the `sid` or the `address` of the configuration

## twilio_conversations_conversations_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**friendly_name** | string | Optional | The human-readable name of this conversation, limited to 256 characters. Optional.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
**date_created** | string | Optional | The date that this resource was created.
**date_updated** | string | Optional | The date that this resource was last updated.
**messaging_service_sid** | string | Optional | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**attributes** | string | Optional | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**state** | string | Optional | 
**timers_inactive** | string | Optional | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**timers_closed** | string | Optional | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

## twilio_conversations_conversations_messages_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**conversation_sid** | string | **Required** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**author** | string | Optional | The channel specific identifier of the message's author. Defaults to `system`.
**body** | string | Optional | The content of the message, can be up to 1,600 characters long.
**date_created** | string | Optional | The date that this resource was created.
**date_updated** | string | Optional | The date that this resource was last updated. `null` if the message has not been edited.
**attributes** | string | Optional | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**media_sid** | string | Optional | The Media SID to be attached to the new Message.
**content_sid** | string | Optional | The unique ID of the multi-channel [Rich Content](https://www.twilio.com/docs/content-api) template, required for template-generated messages.  **Note** that if this field is set, `Body` and `MediaSid` parameters are ignored.
**content_variables** | string | Optional | A structurally valid JSON string that contains values to resolve Rich Content template variables.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_conversations_conversations_participants_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**conversation_sid** | string | **Required** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**identity** | string | Optional | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
**messaging_binding_address** | string | Optional | The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
**messaging_binding_proxy_address** | string | Optional | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
**date_created** | string | Optional | The date that this resource was created.
**date_updated** | string | Optional | The date that this resource was last updated.
**attributes** | string | Optional | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**messaging_binding_projected_address** | string | Optional | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
**role_sid** | string | Optional | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.
**last_read_message_index** | int | Optional | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
**last_read_timestamp** | string | Optional | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.

## twilio_conversations_conversations_webhooks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**conversation_sid** | string | **Required** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**target** | string | **Required** | 
**configuration_url** | string | Optional | The absolute url the webhook request should be sent to.
**configuration_method** | string | Optional | 
**configuration_filters** | list(string) | Optional | The list of events, firing webhook event for this Conversation.
**configuration_triggers** | list(string) | Optional | The list of keywords, firing webhook event for this Conversation.
**configuration_flow_sid** | string | Optional | The studio flow SID, where the webhook should be sent to.
**configuration_replay_after** | int | Optional | The message index for which and it's successors the webhook will be replayed. Not set by default
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_conversations_credentials_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | 
**friendly_name** | string | Optional | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**certificate** | string | Optional | [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A== -----END CERTIFICATE-----`.
**private_key** | string | Optional | [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----`.
**sandbox** | bool | Optional | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production.
**api_key** | string | Optional | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**secret** | string | Optional | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_conversations_roles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**type** | string | **Required** | 
**permission** | list(string) | **Required** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`.
**sid** | string | *Computed* | The SID of the Role resource to update.

## twilio_conversations_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The human-readable name of this service, limited to 256 characters. Optional.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_conversations_services_conversations_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**chat_service_sid** | string | **Required** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**friendly_name** | string | Optional | The human-readable name of this conversation, limited to 256 characters. Optional.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
**attributes** | string | Optional | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**messaging_service_sid** | string | Optional | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**date_created** | string | Optional | The date that this resource was created.
**date_updated** | string | Optional | The date that this resource was last updated.
**state** | string | Optional | 
**timers_inactive** | string | Optional | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**timers_closed** | string | Optional | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

## twilio_conversations_services_conversations_messages_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**chat_service_sid** | string | **Required** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**conversation_sid** | string | **Required** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**author** | string | Optional | The channel specific identifier of the message's author. Defaults to `system`.
**body** | string | Optional | The content of the message, can be up to 1,600 characters long.
**date_created** | string | Optional | The date that this resource was created.
**date_updated** | string | Optional | The date that this resource was last updated. `null` if the message has not been edited.
**attributes** | string | Optional | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**media_sid** | string | Optional | The Media SID to be attached to the new Message.
**content_sid** | string | Optional | The unique ID of the multi-channel [Rich Content](https://www.twilio.com/docs/content-api) template, required for template-generated messages.  **Note** that if this field is set, `Body` and `MediaSid` parameters are ignored.
**content_variables** | string | Optional | A structurally valid JSON string that contains values to resolve Rich Content template variables.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_conversations_services_conversations_participants_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**chat_service_sid** | string | **Required** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**conversation_sid** | string | **Required** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**identity** | string | Optional | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
**messaging_binding_address** | string | Optional | The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
**messaging_binding_proxy_address** | string | Optional | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
**date_created** | string | Optional | The date that this resource was created.
**date_updated** | string | Optional | The date that this resource was last updated.
**attributes** | string | Optional | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**messaging_binding_projected_address** | string | Optional | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
**role_sid** | string | Optional | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.
**last_read_message_index** | int | Optional | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
**last_read_timestamp** | string | Optional | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.

## twilio_conversations_services_conversations_webhooks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**chat_service_sid** | string | **Required** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**conversation_sid** | string | **Required** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**target** | string | **Required** | 
**configuration_url** | string | Optional | The absolute url the webhook request should be sent to.
**configuration_method** | string | Optional | 
**configuration_filters** | list(string) | Optional | The list of events, firing webhook event for this Conversation.
**configuration_triggers** | list(string) | Optional | The list of keywords, firing webhook event for this Conversation.
**configuration_flow_sid** | string | Optional | The studio flow SID, where the webhook should be sent to.
**configuration_replay_after** | int | Optional | The message index for which and it's successors the webhook will be replayed. Not set by default
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_conversations_services_roles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**chat_service_sid** | string | **Required** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to create the Role resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**type** | string | **Required** | 
**permission** | list(string) | **Required** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`.
**sid** | string | *Computed* | The SID of the Role resource to update.

## twilio_conversations_services_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**chat_service_sid** | string | **Required** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.
**identity** | string | **Required** | The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**friendly_name** | string | Optional | The string that you assigned to describe the resource.
**attributes** | string | Optional | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
**role_sid** | string | Optional | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
**sid** | string | *Computed* | The SID of the User resource to update. This value can be either the `sid` or the `identity` of the User resource to update.

## twilio_conversations_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**identity** | string | **Required** | The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**friendly_name** | string | Optional | The string that you assigned to describe the resource.
**attributes** | string | Optional | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
**role_sid** | string | Optional | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
**sid** | string | *Computed* | The SID of the User resource to update. This value can be either the `sid` or the `identity` of the User resource to update.

