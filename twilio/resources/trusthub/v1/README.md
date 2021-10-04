
## twilio_trusthub_customer_profiles_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**email** | string | **Required** | The email address that will receive updates when the Customer-Profile resource changes status.
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**policy_sid** | string | **Required** | The unique string of a policy that is associated to the Customer-Profile resource.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**sid** | string | *Computed* | The unique string that we created to identify the Customer-Profile resource.
**status** | string | Optional | The verification status of the Customer-Profile resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of the Assigned Items of the Customer-Profile resource
**url** | string | *Computed* | The absolute URL of the Customer-Profile resource
**valid_until** | string | *Computed* | The ISO 8601 date and time in GMT when the resource will be valid until.

## twilio_trusthub_customer_profiles_channel_endpoint_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**customer_profile_sid** | string | **Required** | The unique string that we created to identify the CustomerProfile resource.
**channel_endpoint_sid** | string | **Required** | The SID of an channel endpoint
**channel_endpoint_type** | string | **Required** | The type of channel endpoint. eg: phone-number
**sid** | string | *Computed* | The unique string that we created to identify the resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**url** | string | *Computed* | The absolute URL of the Identity resource

## twilio_trusthub_customer_profiles_entity_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**customer_profile_sid** | string | **Required** | The unique string that we created to identify the CustomerProfile resource.
**object_sid** | string | **Required** | The SID of an object bag that holds information of the different items.
**sid** | string | *Computed* | The unique string that we created to identify the Identity resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**url** | string | *Computed* | The absolute URL of the Identity resource

## twilio_trusthub_end_users_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;.
**attributes** | string | Optional | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the End User resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the End User resource

## twilio_trusthub_supporting_documents_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of the Supporting Document.
**attributes** | string | Optional | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the Supporting Document resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**mime_type** | string | *Computed* | The image type of the file
**status** | string | *Computed* | The verification status of the Supporting Document resource
**url** | string | *Computed* | The absolute URL of the Supporting Document resource

## twilio_trusthub_trust_products_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**email** | string | **Required** | The email address that will receive updates when the Customer-Profile resource changes status.
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**policy_sid** | string | **Required** | The unique string of a policy that is associated to the Customer-Profile resource.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**sid** | string | *Computed* | The unique string that we created to identify the Customer-Profile resource.
**status** | string | Optional | The verification status of the Customer-Profile resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of the Assigned Items of the Customer-Profile resource
**url** | string | *Computed* | The absolute URL of the Customer-Profile resource
**valid_until** | string | *Computed* | The ISO 8601 date and time in GMT when the resource will be valid until.

## twilio_trusthub_trust_products_channel_endpoint_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trust_product_sid** | string | **Required** | The unique string that we created to identify the CustomerProfile resource.
**channel_endpoint_sid** | string | **Required** | The SID of an channel endpoint
**channel_endpoint_type** | string | **Required** | The type of channel endpoint. eg: phone-number
**sid** | string | *Computed* | The unique string that we created to identify the resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**url** | string | *Computed* | The absolute URL of the Identity resource

## twilio_trusthub_trust_products_entity_assignments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**trust_product_sid** | string | **Required** | The unique string that we created to identify the TrustProduct resource.
**object_sid** | string | **Required** | The SID of an object bag that holds information of the different items.
**sid** | string | *Computed* | The unique string that we created to identify the Identity resource.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**url** | string | *Computed* | The absolute URL of the Identity resource

