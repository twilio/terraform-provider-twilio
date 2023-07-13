
## twilio_numbers_hosted_number_authorization_documents_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**address_sid** | string | **Required** | A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument.
**email** | string | **Required** | Email that this AuthorizationDocument will be sent to for signing.
**contact_phone_number** | string | **Required** | The contact phone number of the person authorized to sign the Authorization Document.
**hosted_number_order_sids** | list(string) | **Required** | A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio's platform.
**contact_title** | string | Optional | The title of the person authorized to sign the Authorization Document for this phone number.
**cc_emails** | list(string) | Optional | Email recipients who will be informed when an Authorization Document has been sent and signed.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this AuthorizationDocument.

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

## twilio_numbers_hosted_number_orders_v2

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**phone_number** | string | **Required** | The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format
**contact_phone_number** | string | **Required** | The contact phone number of the person authorized to sign the Authorization Document.
**address_sid** | string | **Required** | Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number.
**email** | string | **Required** | Optional. Email of the owner of this phone number that is being hosted.
**account_sid** | string | Optional | This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to.
**friendly_name** | string | Optional | A 64 character string that is a human readable text that describes this resource.
**cc_emails** | list(string) | Optional | Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to.
**sms_url** | string | Optional | The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource.
**sms_method** | string | Optional | The HTTP method that should be used to request the SmsUrl. Must be either `GET` or `POST`.  This will be copied onto the IncomingPhoneNumber resource.
**sms_fallback_url** | string | Optional | A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource.
**sms_capability** | bool | Optional | Used to specify that the SMS capability will be hosted on Twilio's platform.
**sms_fallback_method** | string | Optional | The HTTP method that should be used to request the SmsFallbackUrl. Must be either `GET` or `POST`. This will be copied onto the IncomingPhoneNumber resource.
**status_callback_url** | string | Optional | Optional. The Status Callback URL attached to the IncomingPhoneNumber resource.
**status_callback_method** | string | Optional | Optional. The Status Callback Method attached to the IncomingPhoneNumber resource.
**sms_application_sid** | string | Optional | Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a `SmsApplicationSid` is present, Twilio will ignore all of the SMS urls above and use those set on the application.
**contact_title** | string | Optional | The title of the person authorized to sign the Authorization Document for this phone number.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this HostedNumberOrder.

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

