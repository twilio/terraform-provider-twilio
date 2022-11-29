
## twilio_wireless_commands_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**command** | string | **Required** | The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
**sim** | string | Optional | The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
**callback_method** | string | Optional | The HTTP method we use to call `callback_url`. Can be: `POST` or `GET`, and the default is `POST`.
**callback_url** | string | Optional | The URL we call using the `callback_url` when the Command has finished sending, whether the command was delivered or it failed.
**command_mode** | string | Optional | 
**include_sid** | string | Optional | Whether to include the SID of the command in the message body. Can be: `none`, `start`, or `end`, and the default behavior is `none`. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of `start` will prepend the message with the Command SID, and `end` will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
**delivery_receipt_requested** | bool | Optional | Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to 'delivered' once the server has received a delivery receipt from the device. The default value is `true`.
**sid** | string | *Computed* | The SID of the Command resource to fetch.

## twilio_wireless_rate_plans_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It does not have to be unique.
**data_enabled** | bool | Optional | Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
**data_limit** | int | Optional | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is `1000`.
**data_metering** | string | Optional | The model used to meter data usage. Can be: `payg` and `quota-1`, `quota-10`, and `quota-50`. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
**messaging_enabled** | bool | Optional | Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource).
**voice_enabled** | bool | Optional | Deprecated.
**national_roaming_enabled** | bool | Optional | Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming).
**international_roaming** | list(string) | Optional | The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can contain: `data` and `messaging`.
**national_roaming_data_limit** | int | Optional | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info.
**international_roaming_data_limit** | int | Optional | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.
**sid** | string | *Computed* | The SID of the RatePlan resource to update.

