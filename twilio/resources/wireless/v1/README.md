
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

