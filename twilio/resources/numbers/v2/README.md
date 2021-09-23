
## twilio_numbers_regulatory_compliance_bundles_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**email** | string | **Required** | The email address that will receive updates when the Bundle resource changes status.
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**end_user_type** | string | Optional | The [type of End User](https://www.twilio.com/docs/phone-numbers/regulatory/api/end-user-types) of the Bundle resource.
**iso_country** | string | Optional | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Bundle&#39;s phone number country ownership request.
**number_type** | string | Optional | The type of phone number of the Bundle&#39;s ownership request. Can be &#x60;local&#x60;, &#x60;mobile&#x60;, &#x60;national&#x60;, or &#x60;toll free&#x60;.
**regulation_sid** | string | Optional | The unique string of a regulation that is associated to the Bundle resource.
**status_callback** | string | Optional | The URL we call to inform your application of status changes.
**sid** | string | *Computed* | The unique string that we created to identify the Bundle resource.
**status** | string | Optional | The verification status of the Bundle resource.

## twilio_numbers_regulatory_compliance_end_users_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;.
**attributes** | string | Optional | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the End User resource.

## twilio_numbers_regulatory_compliance_supporting_documents_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | The string that you assigned to describe the resource.
**type** | string | **Required** | The type of the Supporting Document.
**attributes** | string | Optional | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
**sid** | string | *Computed* | The unique string created by Twilio to identify the Supporting Document resource.

