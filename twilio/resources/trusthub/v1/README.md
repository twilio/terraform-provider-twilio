
## twilio_trusthub_customer_profiles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**email** | string | **Required** | The email address that will receive updates when the Customer-Profile resource changes status.
**policy_sid** | string | **Required** | The unique string of a policy that is associated to the Customer-Profile resource.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**sid** | string | *Computed* | The unique string that we created to identify the Customer-Profile resource.
**status** | string | Optional | 

## twilio_trusthub_customer_profiles_channel_endpoint_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**customer_profile_sid** | string | **Required** | The unique string that we created to identify the CustomerProfile resource.
**channel_endpoint_type** | string | **Required** | The type of channel endpoint. eg: phone-number
**channel_endpoint_sid** | string | **Required** | The SID of an channel endpoint
**sid** | string | *Computed* | The unique string that we created to identify the resource.

## twilio_trusthub_customer_profiles_entity_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**customer_profile_sid** | string | **Required** | The unique string that we created to identify the CustomerProfile resource.
**object_sid** | string | **Required** | The SID of an object bag that holds information of the different items.
**sid** | string | *Computed* | The unique string that we created to identify the Identity resource.

## twilio_trusthub_end_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;.
**attributes** | string | Optional | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the End User resource.

## twilio_trusthub_supporting_documents_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of the Supporting Document.
**attributes** | string | Optional | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the Supporting Document resource.

## twilio_trusthub_trust_products_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**email** | string | **Required** | The email address that will receive updates when the Customer-Profile resource changes status.
**policy_sid** | string | **Required** | The unique string of a policy that is associated to the Customer-Profile resource.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**sid** | string | *Computed* | The unique string that we created to identify the Customer-Profile resource.
**status** | string | Optional | 

## twilio_trusthub_trust_products_channel_endpoint_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trust_product_sid** | string | **Required** | The unique string that we created to identify the CustomerProfile resource.
**channel_endpoint_type** | string | **Required** | The type of channel endpoint. eg: phone-number
**channel_endpoint_sid** | string | **Required** | The SID of an channel endpoint
**sid** | string | *Computed* | The unique string that we created to identify the resource.

## twilio_trusthub_trust_products_entity_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trust_product_sid** | string | **Required** | The unique string that we created to identify the TrustProduct resource.
**object_sid** | string | **Required** | The SID of an object bag that holds information of the different items.
**sid** | string | *Computed* | The unique string that we created to identify the Identity resource.

