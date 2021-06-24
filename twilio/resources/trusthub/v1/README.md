
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
**email** | string | **Required** | The email address that will receive updates when the Customer-Profile resource changes status.
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**policy_sid** | string | **Required** | The unique string of a policy that is associated to the Customer-Profile resource.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**sid** | string | *Computed* | The unique string that we created to identify the Customer-Profile resource.
**status** | string | Optional | The verification status of the Customer-Profile resource.

