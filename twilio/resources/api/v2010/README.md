
## twilio_api_accounts_addresses_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**city** | string | **Required** | The city of the new address.
**customer_name** | string | **Required** | The name to associate with the new address.
**iso_country** | string | **Required** | The ISO country code of the new address.
**postal_code** | string | **Required** | The postal code of the new address.
**region** | string | **Required** | The state or region of the new address.
**street** | string | **Required** | The number and street address of the new address.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.
**auto_correct_address** | bool | Optional | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide.
**emergency_enabled** | bool | Optional | Whether to enable emergency calling on the new address. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new address. It can be up to 64 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Address resource to update.

## twilio_api_accounts_applications_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**api_version** | string | Optional | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is the account&#39;s default API version.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new application. It can be up to 64 characters long.
**message_status_callback** | string | Optional | The URL we should call using a POST method to send message status information to your application.
**sms_fallback_method** | string | Optional | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**sms_fallback_url** | string | Optional | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;.
**sms_method** | string | Optional | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**sms_status_callback** | string | Optional | The URL we should call using a POST method to send status information about SMS messages sent by the application.
**sms_url** | string | Optional | The URL we should call when the phone number receives an incoming SMS message.
**status_callback** | string | Optional | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_caller_id_lookup** | bool | Optional | Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**voice_fallback_method** | string | Optional | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
**voice_method** | string | Optional | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_url** | string | Optional | The URL we should call when the phone number assigned to this application receives a call.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Application resource to update.

## twilio_api_accounts_calls_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**from** | string | **Required** | The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;From&#x60; must also be a phone number.
**to** | string | **Required** | The phone number, SIP address, or client identifier to call.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**application_sid** | string | Optional | The SID of the Application resource that will handle the call, if the call will be handled by an application.
**async_amd** | string | Optional | Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**async_amd_status_callback** | string | Optional | The URL that we should call using the &#x60;async_amd_status_callback_method&#x60; to notify customer application whether the call was answered by human, machine or fax.
**async_amd_status_callback_method** | string | Optional | The HTTP method we should use when calling the &#x60;async_amd_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**byoc** | string | Optional | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta)
**call_reason** | string | Optional | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta)
**call_token** | string | Optional | A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. this field should be populated by the incoming call&#39;s call_token to make this outgoing call as a forwarded call of incoming call. A forwarded call should bear the same caller-id of incoming call.
**caller_id** | string | Optional | The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as &#x60;name@company.com&#x60;.
**fallback_method** | string | Optional | The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
**fallback_url** | string | Optional | The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
**machine_detection** | string | Optional | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: &#x60;Enable&#x60; or &#x60;DetectMessageEnd&#x60;. Use &#x60;Enable&#x60; if you would like us to return &#x60;AnsweredBy&#x60; as soon as the called party is identified. Use &#x60;DetectMessageEnd&#x60;, if you would like to leave a message on an answering machine. If &#x60;send_digits&#x60; is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection).
**machine_detection_silence_timeout** | int | Optional | The number of milliseconds of initial silence after which an &#x60;unknown&#x60; AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000.
**machine_detection_speech_end_threshold** | int | Optional | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200.
**machine_detection_speech_threshold** | int | Optional | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400.
**machine_detection_timeout** | int | Optional | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with &#x60;AnsweredBy&#x60; of &#x60;unknown&#x60;. The default timeout is 30 seconds.
**method** | string | Optional | The HTTP method we should use when calling the &#x60;url&#x60; parameter&#39;s value. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
**record** | bool | Optional | Whether to record the call. Can be &#x60;true&#x60; to record the phone call, or &#x60;false&#x60; to not. The default is &#x60;false&#x60;. The &#x60;recording_url&#x60; is sent to the &#x60;status_callback&#x60; URL.
**recording_channels** | string | Optional | The number of channels in the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60;. The default is &#x60;mono&#x60;. &#x60;mono&#x60; records both legs of the call in a single channel of the recording file. &#x60;dual&#x60; records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call.
**recording_status_callback** | string | Optional | The URL that we call when the recording is available to be accessed.
**recording_status_callback_event** | list(string) | Optional | The recording status events that will trigger calls to the URL specified in &#x60;recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Defaults to &#x60;completed&#x60;. Separate  multiple values with a space.
**recording_status_callback_method** | string | Optional | The HTTP method we should use when calling the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**recording_track** | string | Optional | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio.
**send_digits** | string | Optional | A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (&#x60;0&#x60;-&#x60;9&#x60;), &#39;&#x60;#&#x60;&#39;, &#39;&#x60;*&#x60;&#39; and &#39;&#x60;w&#x60;&#39;, to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be &#x60;ww1234#&#x60;. Remember to URL-encode this string, since the &#39;&#x60;#&#x60;&#39; character has special meaning in a URL. If both &#x60;SendDigits&#x60; and &#x60;MachineDetection&#x60; parameters are provided, then &#x60;MachineDetection&#x60; will be ignored.
**sip_auth_password** | string | Optional | The password required to authenticate the user account specified in &#x60;sip_auth_username&#x60;.
**sip_auth_username** | string | Optional | The username used to authenticate the caller making a SIP call.
**status_callback** | string | Optional | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
**status_callback_event** | list(string) | Optional | The call progress events that we will send to the &#x60;status_callback&#x60; URL. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. If no event is specified, we send the &#x60;completed&#x60; status. If you want to receive multiple events, specify each one in a separate &#x60;status_callback_event&#x60; parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample&#x3D;code-create-a-call-resource-and-specify-a-statuscallbackevent&amp;code-sdk-version&#x3D;json). If an &#x60;application_sid&#x60; is present, this parameter is ignored.
**status_callback_method** | string | Optional | The HTTP method we should use when calling the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
**time_limit** | int | Optional | The maximum duration of the call in seconds. Constraints depend on account and configuration.
**timeout** | int | Optional | The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is &#x60;60&#x60; seconds and the maximum is &#x60;600&#x60; seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as &#x60;15&#x60; seconds, to hang up before reaching an answering machine or voicemail.
**trim** | string | Optional | Whether to trim any leading and trailing silence from the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;trim-silence&#x60;.
**twiml** | string | Optional | TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both &#x60;twiml&#x60; and &#x60;url&#x60; are provided then &#x60;twiml&#x60; parameter will be ignored.
**url** | string | Optional | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Call resource to update
**status** | string | Optional | The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;completed&#x60;. Specifying &#x60;canceled&#x60; will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying &#x60;completed&#x60; will attempt to hang up a call even if it&#39;s already in progress.

## twilio_api_accounts_calls_recordings_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**call_sid** | string | **Required** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**recording_channels** | string | Optional | The number of channels used in the recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. &#x60;mono&#x60; records all parties of the call into one channel. &#x60;dual&#x60; records each party of a 2-party call into separate channels.
**recording_status_callback** | string | Optional | The URL we should call using the &#x60;recording_status_callback_method&#x60; on each recording event specified in  &#x60;recording_status_callback_event&#x60;. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback).
**recording_status_callback_event** | list(string) | Optional | The recording status events on which we should call the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60; and the default is &#x60;completed&#x60;. Separate multiple event values with a space.
**recording_status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**recording_track** | string | Optional | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio.
**trim** | string | Optional | Whether to trim any leading and trailing silence in the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;do-not-trim&#x60;. &#x60;trim-silence&#x60; trims the silence from the beginning and end of the recording and &#x60;do-not-trim&#x60; does not.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Recording resource to update.
**status** | string | Optional | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;.
**pause_behavior** | string | Optional | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;.

## twilio_api_accounts_incoming_phone_numbers_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**address_sid** | string | Optional | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
**api_version** | string | Optional | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;.
**area_code** | string | Optional | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an &#x60;area_code&#x60; or a &#x60;phone_number&#x60;.** (US and Canada only).
**bundle_sid** | string | Optional | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
**emergency_address_sid** | string | Optional | The SID of the emergency address configuration to use for emergency calling from the new phone number.
**emergency_status** | string | Optional | The configuration status parameter that determines whether the new phone number is enabled for emergency calling.
**friendly_name** | string | Optional | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number.
**identity_sid** | string | Optional | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
**phone_number** | string | Optional | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**sms_application_sid** | string | Optional | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application.
**sms_fallback_method** | string | Optional | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**sms_fallback_url** | string | Optional | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
**sms_method** | string | Optional | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**sms_url** | string | Optional | The URL we should call when the new phone number receives an incoming SMS message.
**status_callback** | string | Optional | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**trunk_sid** | string | Optional | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
**voice_application_sid** | string | Optional | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
**voice_caller_id_lookup** | bool | Optional | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
**voice_fallback_method** | string | Optional | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
**voice_method** | string | Optional | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**voice_receive_mode** | string | Optional | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;.
**voice_url** | string | Optional | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update.
**account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).

## twilio_api_accounts_messages_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**to** | string | **Required** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**address_retention** | string | Optional | Determines if the address can be stored or obfuscated based on privacy settings
**application_sid** | string | Optional | The SID of the application that should receive message status. We POST a &#x60;message_sid&#x60; parameter and a &#x60;message_status&#x60; parameter with a value of &#x60;sent&#x60; or &#x60;failed&#x60; to the [application](https://www.twilio.com/docs/usage/api/applications)&#39;s &#x60;message_status_callback&#x60;. If a &#x60;status_callback&#x60; parameter is also passed, it will be ignored and the application&#39;s &#x60;message_status_callback&#x60; parameter will be used.
**attempt** | int | Optional | Total number of attempts made ( including this ) to send out the message regardless of the provider used
**body** | string | Optional | The text of the message you want to send. Can be up to 1,600 characters in length.
**content_retention** | string | Optional | Determines if the message content can be stored or redacted based on privacy settings
**force_delivery** | bool | Optional | Reserved
**from** | string | Optional | A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using &#x60;messaging_service_sid&#x60;, this parameter must be empty.
**max_price** | float | Optional | The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds &#x60;max_price&#x60;, the message will fail and a status of &#x60;Failed&#x60; is sent to the status callback. If &#x60;MaxPrice&#x60; is not set, the message cost is not checked.
**media_url** | list(string) | Optional | The URL of the media to send with the message. The media can be of type &#x60;gif&#x60;, &#x60;png&#x60;, and &#x60;jpeg&#x60; and will be formatted correctly on the recipient&#39;s device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple &#x60;media_url&#x60; parameters in the POST request. You can include up to 10 &#x60;media_url&#x60; parameters per message. You can send images in an SMS message in only the US and Canada.
**messaging_service_sid** | string | Optional | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the &#x60;from&#x60; parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the &#x60;from&#x60; phone number for delivery.
**persistent_action** | list(string) | Optional | Rich actions for Channels Messages.
**provide_feedback** | bool | Optional | Whether to confirm delivery of the message. Set this value to &#x60;true&#x60; if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is &#x60;false&#x60; by default.
**smart_encoded** | bool | Optional | Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**status_callback** | string | Optional | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If specified, we POST these message status changes to the URL: &#x60;queued&#x60;, &#x60;failed&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, or &#x60;undelivered&#x60;. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including &#x60;MessageSid&#x60;, &#x60;MessageStatus&#x60;, and &#x60;ErrorCode&#x60;. If you include this parameter with the &#x60;messaging_service_sid&#x60;, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed.
**validity_period** | int | Optional | How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Message resource to update.

## twilio_api_accounts_keys_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Key resource to update.

## twilio_api_accounts_signing_keys_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**sid** | string | *Computed* | 

## twilio_api_accounts_conferences_participants_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**conference_sid** | string | **Required** | The SID of the participant&#39;s conference.
**from** | string | **Required** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;from&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;from&#x60; should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
**to** | string | **Required** | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as &#x60;sip:name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**beep** | string | Optional | Whether to play a notification beep to the conference when the participant joins. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;.
**byoc** | string | Optional | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta)
**call_reason** | string | Optional | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta)
**call_sid_to_coach** | string | Optional | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;.
**caller_id** | string | Optional | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;callerId&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;callerId&#x60; should be a username portion to be used to populate the From header that is passed to the SIP endpoint.
**coaching** | bool | Optional | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined.
**conference_record** | string | Optional | Whether to record the conference the participant is joining. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;.
**conference_recording_status_callback** | string | Optional | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available.
**conference_recording_status_callback_event** | list(string) | Optional | The conference recording state changes that generate a call to &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, and &#x60;absent&#x60;. Separate multiple values with a space, ex: &#x60;&#39;in-progress completed failed&#39;&#x60;
**conference_recording_status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**conference_status_callback** | string | Optional | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored.
**conference_status_callback_event** | list(string) | Optional | The conference state changes that should generate a call to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;modify&#x60;, &#x60;speaker&#x60;, and &#x60;announcement&#x60;. Separate multiple values with a space. Defaults to &#x60;start end&#x60;.
**conference_status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**conference_trim** | string | Optional | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;.
**early_media** | bool | Optional | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;.
**end_conference_on_exit** | bool | Optional | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
**jitter_buffer_size** | string | Optional | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant&#39;s audio is mixed into the conference. Can be: &#x60;off&#x60;, &#x60;small&#x60;, &#x60;medium&#x60;, and &#x60;large&#x60;. Default to &#x60;large&#x60;.
**label** | string | Optional | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant.
**max_participants** | int | Optional | The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;.
**muted** | bool | Optional | Whether the agent is muted in the conference. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**record** | bool | Optional | Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**recording_channels** | string | Optional | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;.
**recording_status_callback** | string | Optional | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes.
**recording_status_callback_event** | list(string) | Optional | The recording state changes that should generate a call to &#x60;recording_status_callback&#x60;. Can be: &#x60;started&#x60;, &#x60;in-progress&#x60;, &#x60;paused&#x60;, &#x60;resumed&#x60;, &#x60;stopped&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, and &#x60;absent&#x60;. Separate multiple values with a space, ex: &#x60;&#39;in-progress completed failed&#39;&#x60;.
**recording_status_callback_method** | string | Optional | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**recording_track** | string | Optional | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is sent from Twilio. &#x60;both&#x60; records the audio that is received and sent by Twilio.
**region** | string | Optional | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;.
**sip_auth_password** | string | Optional | The SIP password for authentication.
**sip_auth_username** | string | Optional | The SIP username used for authentication.
**start_conference_on_enter** | bool | Optional | Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
**status_callback** | string | Optional | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**status_callback_event** | list(string) | Optional | The conference state changes that should generate a call to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. Separate multiple values with a space. The default value is &#x60;completed&#x60;.
**status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; and &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**time_limit** | int | Optional | The maximum duration of the call in seconds. Constraints depend on account and configuration.
**timeout** | int | Optional | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between &#x60;5&#x60; and &#x60;600&#x60;, inclusive. The default value is &#x60;60&#x60;. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds.
**wait_method** | string | Optional | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
**wait_url** | string | Optional | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
**call_sid** | string | *Computed* | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.
**announce_method** | string | Optional | The HTTP method we should use to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**announce_url** | string | Optional | The URL we call using the &#x60;announce_method&#x60; for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60; commands.
**beep_on_exit** | bool | Optional | Whether to play a notification beep to the conference when the participant exits. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**hold** | bool | Optional | Whether the participant should be on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; puts the participant on hold, and &#x60;false&#x60; lets them rejoin the conference.
**hold_method** | string | Optional | The HTTP method we should use to call &#x60;hold_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;GET&#x60;.
**hold_url** | string | Optional | The URL we call using the &#x60;hold_method&#x60; for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60; or &#x60;&lt;Redirect&gt;&#x60; commands.

## twilio_api_accounts_queues_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you created to describe this resource. It can be up to 64 characters long.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**max_size** | int | Optional | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Queue resource to update

## twilio_api_accounts_sipcredential_lists_credentials_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**credential_list_sid** | string | **Required** | The unique id that identifies the credential list to include the created credential.
**password** | string | **Required** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;)
**username** | string | **Required** | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio&#39;s challenge of the initial INVITE. It can be up to 32 characters long.
**path_account_sid** | string | Optional | The unique id of the Account that is responsible for this resource.
**sid** | string | *Computed* | The unique id that identifies the resource to update.

## twilio_api_accounts_sipcredential_lists_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A human readable descriptive text that describes the CredentialList, up to 64 characters long.
**path_account_sid** | string | Optional | The unique id of the Account that is responsible for this resource.
**sid** | string | *Computed* | The credential list Sid that uniquely identifies this resource

## twilio_api_accounts_sipdomains_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_name** | string | **Required** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**byoc_trunk_sid** | string | Optional | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
**emergency_caller_sid** | string | Optional | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
**emergency_calling_enabled** | bool | Optional | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
**friendly_name** | string | Optional | A descriptive string that you created to describe the resource. It can be up to 64 characters long.
**secure** | bool | Optional | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
**sip_registration** | bool | Optional | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not.
**voice_fallback_method** | string | Optional | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;.
**voice_method** | string | Optional | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_status_callback_url** | string | Optional | The URL that we should call to pass status parameters (such as call ended) to your application.
**voice_url** | string | Optional | The URL we should when the domain receives a call.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the SipDomain resource to update.

## twilio_api_accounts_sipip_access_control_lists_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long.
**path_account_sid** | string | Optional | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**sid** | string | *Computed* | A 34 character string that uniquely identifies the resource to udpate.

## twilio_api_accounts_sipip_access_control_lists_ip_addresses_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**ip_access_control_list_sid** | string | **Required** | The IpAccessControlList Sid with which to associate the created IpAddress resource.
**friendly_name** | string | **Required** | A human readable descriptive text for this resource, up to 64 characters long.
**ip_address** | string | **Required** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
**path_account_sid** | string | Optional | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**cidr_prefix_length** | int | Optional | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
**sid** | string | *Computed* | A 34 character string that identifies the IpAddress resource to update.

## twilio_api_accounts_usage_triggers_v2010

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**callback_url** | string | **Required** | The URL we should call using &#x60;callback_method&#x60; when the trigger fires.
**trigger_value** | string | **Required** | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as &#x60;+30&#x60; to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a &#x60;+&#x60; as &#x60;%2B&#x60;.
**usage_category** | string | **Required** | The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**callback_method** | string | Optional | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**recurring** | string | Optional | The frequency of a recurring UsageTrigger.  Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT.
**trigger_by** | string | Optional | The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is &#x60;usage&#x60;.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update.

