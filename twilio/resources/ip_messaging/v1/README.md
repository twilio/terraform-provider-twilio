
## twilio_ip_messaging_services_channels_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**attributes** | string | Optional | 
**friendly_name** | string | Optional | 
**type** | string | Optional | 
**unique_name** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_credentials_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**type** | string | **Required** | 
**api_key** | string | Optional | 
**certificate** | string | Optional | 
**friendly_name** | string | Optional | 
**private_key** | string | Optional | 
**sandbox** | bool | Optional | 
**secret** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_channels_members_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**channel_sid** | string | **Required** | 
**identity** | string | **Required** | 
**role_sid** | string | Optional | 
**sid** | string | *Computed* | 
**last_consumed_message_index** | int | Optional | 

## twilio_ip_messaging_services_channels_messages_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**channel_sid** | string | **Required** | 
**body** | string | **Required** | 
**attributes** | string | Optional | 
**from** | string | Optional | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_roles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**friendly_name** | string | **Required** | 
**permission** | list(string) | **Required** | 
**type** | string | **Required** | 
**sid** | string | *Computed* | 

## twilio_ip_messaging_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | 
**sid** | string | *Computed* | 
**consumption_report_interval** | int | Optional | 
**default_channel_creator_role_sid** | string | Optional | 
**default_channel_role_sid** | string | Optional | 
**default_service_role_sid** | string | Optional | 
**limits_channel_members** | int | Optional | 
**limits_user_channels** | int | Optional | 
**notifications_added_to_channel_enabled** | bool | Optional | 
**notifications_added_to_channel_template** | string | Optional | 
**notifications_invited_to_channel_enabled** | bool | Optional | 
**notifications_invited_to_channel_template** | string | Optional | 
**notifications_new_message_enabled** | bool | Optional | 
**notifications_new_message_template** | string | Optional | 
**notifications_removed_from_channel_enabled** | bool | Optional | 
**notifications_removed_from_channel_template** | string | Optional | 
**post_webhook_url** | string | Optional | 
**pre_webhook_url** | string | Optional | 
**reachability_enabled** | bool | Optional | 
**read_status_enabled** | bool | Optional | 
**typing_indicator_timeout** | int | Optional | 
**webhook_filters** | list(string) | Optional | 
**webhook_method** | string | Optional | 
**webhooks_on_channel_add_method** | string | Optional | 
**webhooks_on_channel_add_url** | string | Optional | 
**webhooks_on_channel_added_method** | string | Optional | 
**webhooks_on_channel_added_url** | string | Optional | 
**webhooks_on_channel_destroy_method** | string | Optional | 
**webhooks_on_channel_destroy_url** | string | Optional | 
**webhooks_on_channel_destroyed_method** | string | Optional | 
**webhooks_on_channel_destroyed_url** | string | Optional | 
**webhooks_on_channel_update_method** | string | Optional | 
**webhooks_on_channel_update_url** | string | Optional | 
**webhooks_on_channel_updated_method** | string | Optional | 
**webhooks_on_channel_updated_url** | string | Optional | 
**webhooks_on_member_add_method** | string | Optional | 
**webhooks_on_member_add_url** | string | Optional | 
**webhooks_on_member_added_method** | string | Optional | 
**webhooks_on_member_added_url** | string | Optional | 
**webhooks_on_member_remove_method** | string | Optional | 
**webhooks_on_member_remove_url** | string | Optional | 
**webhooks_on_member_removed_method** | string | Optional | 
**webhooks_on_member_removed_url** | string | Optional | 
**webhooks_on_message_remove_method** | string | Optional | 
**webhooks_on_message_remove_url** | string | Optional | 
**webhooks_on_message_removed_method** | string | Optional | 
**webhooks_on_message_removed_url** | string | Optional | 
**webhooks_on_message_send_method** | string | Optional | 
**webhooks_on_message_send_url** | string | Optional | 
**webhooks_on_message_sent_method** | string | Optional | 
**webhooks_on_message_sent_url** | string | Optional | 
**webhooks_on_message_update_method** | string | Optional | 
**webhooks_on_message_update_url** | string | Optional | 
**webhooks_on_message_updated_method** | string | Optional | 
**webhooks_on_message_updated_url** | string | Optional | 

## twilio_ip_messaging_services_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | 
**identity** | string | **Required** | 
**attributes** | string | Optional | 
**friendly_name** | string | Optional | 
**role_sid** | string | Optional | 
**sid** | string | *Computed* | 

