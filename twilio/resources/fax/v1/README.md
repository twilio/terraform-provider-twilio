
## twilio_fax_faxes_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**media_url** | string | **Required** | The URL of the PDF that contains the fax. See our [security](https://www.twilio.com/docs/usage/security) page for information on how to ensure the request for your media comes from Twilio.
**to** | string | **Required** | The phone number to receive the fax in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the recipient&#39;s SIP URI.
**from** | string | Optional | The number the fax was sent from. Can be the phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the SIP &#x60;from&#x60; value. The caller ID displayed to the recipient uses this value. If this is a phone number, it must be a Twilio number or a verified outgoing caller id from your account. If &#x60;to&#x60; is a SIP address, this can be any alphanumeric string (and also the characters &#x60;+&#x60;, &#x60;_&#x60;, &#x60;.&#x60;, and &#x60;-&#x60;), which will be used in the &#x60;from&#x60; header of the SIP request.
**quality** | string | Optional | The [Fax Quality value](https://www.twilio.com/docs/fax/api/fax-resource#fax-quality-values) that describes the fax quality. Can be: &#x60;standard&#x60;, &#x60;fine&#x60;, or &#x60;superfine&#x60; and defaults to &#x60;fine&#x60;.
**sip_auth_password** | string | Optional | The password to use with &#x60;sip_auth_username&#x60; to authenticate faxes sent to a SIP address.
**sip_auth_username** | string | Optional | The username to use with the &#x60;sip_auth_password&#x60; to authenticate faxes sent to a SIP address.
**status_callback** | string | Optional | The URL we should call using the &#x60;POST&#x60; method to send [status information](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-callback) to your application when the status of the fax changes.
**store_media** | bool | Optional | Whether to store a copy of the sent media on our servers for later retrieval. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
**ttl** | int | Optional | How long in minutes from when the fax is initiated that we should try to send the fax.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Fax resource to update.
**status** | string | Optional | The new [status](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-values) of the resource. Can be only &#x60;canceled&#x60;. This may fail if transmission has already started.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**api_version** | string | *Computed* | The API version used to transmit the fax
**date_created** | string | *Computed* | The ISO 8601 formatted date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 formatted date and time in GMT when the resource was last updated
**direction** | string | *Computed* | The direction of the fax
**duration** | int | *Computed* | The time it took to transmit the fax
**links** | string | *Computed* | The URLs of the fax&#39;s related resources
**media_sid** | string | *Computed* | The SID of the FaxMedia resource that is associated with the Fax
**num_pages** | int | *Computed* | The number of pages contained in the fax document
**price** | float | *Computed* | The fax transmission price
**price_unit** | string | *Computed* | The ISO 4217 currency used for billing
**url** | string | *Computed* | The absolute URL of the fax resource

