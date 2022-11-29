
## twilio_numbers_regulatory_compliance_bundles_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**email** | string | **Required** | The email address that will receive updates when the Bundle resource changes status.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**regulation_sid** | string | Optional | The unique string of a regulation that is associated to the Bundle resource.
**iso_country** | string | Optional | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Bundle's phone number country ownership request.
**end_user_type** | string | Optional | 
**number_type** | string | Optional | The type of phone number of the Bundle's ownership request. Can be `local`, `mobile`, `national`, or `toll free`.
**sid** | string | *Computed* | The unique string that we created to identify the Bundle resource.
**status** | string | Optional | 

## twilio_numbers_regulatory_compliance_end_users_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | 
**attributes** | string | Optional | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the End User resource.

## twilio_numbers_regulatory_compliance_bundles_item_assignments_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**bundle_sid** | string | **Required** | The unique string that we created to identify the Bundle resource.
**object_sid** | string | **Required** | The SID of an object bag that holds information of the different items.
**sid** | string | *Computed* | The unique string that we created to identify the Identity resource.

## twilio_numbers_regulatory_compliance_supporting_documents_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of the Supporting Document.
**attributes** | string | Optional | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the Supporting Document resource.

