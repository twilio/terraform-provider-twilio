
## twilio_ip_messaging_services_channels_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**friendly_name** | string | Optional | 
**unique_name** | string | Optional | 
**attributes** | string | Optional | 
**type** | string | Optional | 
**date_created** | string | Optional | 
**date_updated** | string | Optional | 
**created_by** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_channels_webhooks_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**channel_sid** | string | **Required** | 
**type** | string | **Required** | 
**configuration_url** | string | Optional | 
**configuration_method** | string | Optional | 
**configuration_filters** | list(string) | Optional | 
**configuration_triggers** | list(string) | Optional | 
**configuration_flow_sid** | string | Optional | 
**configuration_retry_count** | int | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_credentials_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | 
**friendly_name** | string | Optional | 
**certificate** | string | Optional | 
**private_key** | string | Optional | 
**sandbox** | bool | Optional | 
**api_key** | string | Optional | 
**secret** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_channels_invites_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**channel_sid** | string | **Required** | 
**identity** | string | **Required** | 
**role_sid** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_channels_members_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**channel_sid** | string | **Required** | 
**identity** | string | **Required** | 
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**role_sid** | string | Optional | 
**last_consumed_message_index** | int | Optional | 
**last_consumption_timestamp** | string | Optional | 
**date_created** | string | Optional | 
**date_updated** | string | Optional | 
**attributes** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_channels_messages_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**channel_sid** | string | **Required** | 
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**from** | string | Optional | 
**attributes** | string | Optional | 
**date_created** | string | Optional | 
**date_updated** | string | Optional | 
**last_updated_by** | string | Optional | 
**body** | string | Optional | 
**media_sid** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_roles_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**friendly_name** | string | **Required** | 
**type** | string | **Required** | 
**permission** | list(string) | **Required** | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | 
**sid** | string | *Computed* | 
**default_service_role_sid** | string | Optional | 
**default_channel_role_sid** | string | Optional | 
**default_channel_creator_role_sid** | string | Optional | 
**read_status_enabled** | bool | Optional | 
**reachability_enabled** | bool | Optional | 
**typing_indicator_timeout** | int | Optional | 
**consumption_report_interval** | int | Optional | 
**notifications_new_message_enabled** | bool | Optional | 
**notifications_new_message_template** | string | Optional | 
**notifications_new_message_sound** | string | Optional | 
**notifications_new_message_badge_count_enabled** | bool | Optional | 
**notifications_added_to_channel_enabled** | bool | Optional | 
**notifications_added_to_channel_template** | string | Optional | 
**notifications_added_to_channel_sound** | string | Optional | 
**notifications_removed_from_channel_enabled** | bool | Optional | 
**notifications_removed_from_channel_template** | string | Optional | 
**notifications_removed_from_channel_sound** | string | Optional | 
**notifications_invited_to_channel_enabled** | bool | Optional | 
**notifications_invited_to_channel_template** | string | Optional | 
**notifications_invited_to_channel_sound** | string | Optional | 
**pre_webhook_url** | string | Optional | 
**post_webhook_url** | string | Optional | 
**webhook_method** | string | Optional | 
**webhook_filters** | list(string) | Optional | 
**limits_channel_members** | int | Optional | 
**limits_user_channels** | int | Optional | 
**media_compatibility_message** | string | Optional | 
**pre_webhook_retry_count** | int | Optional | 
**post_webhook_retry_count** | int | Optional | 
**notifications_log_enabled** | bool | Optional | 

## twilio_ip_messaging_services_users_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**identity** | string | **Required** | 
**x_twilio_webhook_enabled** | string | Optional | The X-Twilio-Webhook-Enabled HTTP request header
**role_sid** | string | Optional | 
**attributes** | string | Optional | 
**friendly_name** | string | Optional | 
**sid** | string | *Computed* | 

