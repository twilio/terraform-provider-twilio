
## twilio_wireless_commands_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**command** | string | **Required** | The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
**callback_method** | string | Optional | The HTTP method we use to call &#x60;callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;, and the default is &#x60;POST&#x60;.
**callback_url** | string | Optional | The URL we call using the &#x60;callback_url&#x60; when the Command has finished sending, whether the command was delivered or it failed.
**command_mode** | string | Optional | The mode to use when sending the SMS message. Can be: &#x60;text&#x60; or &#x60;binary&#x60;. The default SMS mode is &#x60;text&#x60;.
**delivery_receipt_requested** | bool | Optional | Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to &#39;delivered&#39; once the server has received a delivery receipt from the device. The default value is &#x60;true&#x60;.
**include_sid** | string | Optional | Whether to include the SID of the command in the message body. Can be: &#x60;none&#x60;, &#x60;start&#x60;, or &#x60;end&#x60;, and the default behavior is &#x60;none&#x60;. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of &#x60;start&#x60; will prepend the message with the Command SID, and &#x60;end&#x60; will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
**sim** | string | Optional | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
**sid** | string | *Computed* | The SID of the Command resource to fetch.

## twilio_wireless_rate_plans_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**data_enabled** | bool | Optional | Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
**data_limit** | int | Optional | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is &#x60;1000&#x60;.
**data_metering** | string | Optional | The model used to meter data usage. Can be: &#x60;payg&#x60; and &#x60;quota-1&#x60;, &#x60;quota-10&#x60;, and &#x60;quota-50&#x60;. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It does not have to be unique.
**international_roaming** | list(string) | Optional | The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can be: &#x60;data&#x60;, &#x60;voice&#x60;, and &#x60;messaging&#x60;.
**international_roaming_data_limit** | int | Optional | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.
**messaging_enabled** | bool | Optional | Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource).
**national_roaming_data_limit** | int | Optional | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info.
**national_roaming_enabled** | bool | Optional | Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming).
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource.
**voice_enabled** | bool | Optional | Whether SIMs can make and receive voice calls.
**sid** | string | *Computed* | The SID of the RatePlan resource to update.

