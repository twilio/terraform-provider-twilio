
## twilio_api_accounts_addresses

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**customer_name** | string | **Required** | The name to associate with the new address.
**street** | string | **Required** | The number and street address of the new address.
**city** | string | **Required** | The city of the new address.
**region** | string | **Required** | The state or region of the new address.
**postal_code** | string | **Required** | The postal code of the new address.
**iso_country** | string | **Required** | The ISO country code of the new address.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new address. It can be up to 64 characters long.
**emergency_enabled** | bool | Optional | Whether to enable emergency calling on the new address. Can be: `true` or `false`.
**auto_correct_address** | bool | Optional | Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide.
**street_secondary** | string | Optional | The additional number and street address of the address.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Address resource to update.

## twilio_api_accounts_applications

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**api_version** | string | Optional | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is the account's default API version.
**voice_url** | string | Optional | The URL we should call when the phone number assigned to this application receives a call.
**voice_method** | string | Optional | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
**voice_fallback_method** | string | Optional | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**status_callback** | string | Optional | The URL we should call using the `status_callback_method` to send status information to your application.
**status_callback_method** | string | Optional | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`.
**voice_caller_id_lookup** | bool | Optional | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`.
**sms_url** | string | Optional | The URL we should call when the phone number receives an incoming SMS message.
**sms_method** | string | Optional | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`.
**sms_fallback_url** | string | Optional | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`.
**sms_fallback_method** | string | Optional | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`.
**sms_status_callback** | string | Optional | The URL we should call using a POST method to send status information about SMS messages sent by the application.
**message_status_callback** | string | Optional | The URL we should call using a POST method to send message status information to your application.
**friendly_name** | string | Optional | A descriptive string that you create to describe the new application. It can be up to 64 characters long.
**public_application_connect_enabled** | bool | Optional | Whether to allow other Twilio accounts to dial this applicaton using Dial verb. Can be: `true` or `false`.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Application resource to update.

## twilio_api_accounts_calls

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**to** | string | **Required** | The phone number, SIP address, or client identifier to call.
**from** | string | **Required** | The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `From` must also be a phone number.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**method** | string | Optional | The HTTP method we should use when calling the `url` parameter's value. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
**fallback_url** | string | Optional | The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored.
**fallback_method** | string | Optional | The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
**status_callback** | string | Optional | The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
**status_callback_event** | list(string) | Optional | The call progress events that we will send to the `status_callback` URL. Can be: `initiated`, `ringing`, `answered`, and `completed`. If no event is specified, we send the `completed` status. If you want to receive multiple events, specify each one in a separate `status_callback_event` parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample=code-create-a-call-resource-and-specify-a-statuscallbackevent&code-sdk-version=json). If an `application_sid` is present, this parameter is ignored.
**status_callback_method** | string | Optional | The HTTP method we should use when calling the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
**send_digits** | string | Optional | A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (`0`-`9`), '`#`', '`*`' and '`w`', to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be `ww1234#`. Remember to URL-encode this string, since the '`#`' character has special meaning in a URL. If both `SendDigits` and `MachineDetection` parameters are provided, then `MachineDetection` will be ignored.
**timeout** | int | Optional | The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is `60` seconds and the maximum is `600` seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as `15` seconds, to hang up before reaching an answering machine or voicemail.
**record** | bool | Optional | Whether to record the call. Can be `true` to record the phone call, or `false` to not. The default is `false`. The `recording_url` is sent to the `status_callback` URL.
**recording_channels** | string | Optional | The number of channels in the final recording. Can be: `mono` or `dual`. The default is `mono`. `mono` records both legs of the call in a single channel of the recording file. `dual` records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call.
**recording_status_callback** | string | Optional | The URL that we call when the recording is available to be accessed.
**recording_status_callback_method** | string | Optional | The HTTP method we should use when calling the `recording_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
**sip_auth_username** | string | Optional | The username used to authenticate the caller making a SIP call.
**sip_auth_password** | string | Optional | The password required to authenticate the user account specified in `sip_auth_username`.
**machine_detection** | string | Optional | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: `Enable` or `DetectMessageEnd`. Use `Enable` if you would like us to return `AnsweredBy` as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. If `send_digits` is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection).
**machine_detection_timeout** | int | Optional | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with `AnsweredBy` of `unknown`. The default timeout is 30 seconds.
**provider_sid** | string | Optional | 
**recording_status_callback_event** | list(string) | Optional | The recording status events that will trigger calls to the URL specified in `recording_status_callback`. Can be: `in-progress`, `completed` and `absent`. Defaults to `completed`. Separate  multiple values with a space.
**trim** | string | Optional | Whether to trim any leading and trailing silence from the recording. Can be: `trim-silence` or `do-not-trim` and the default is `trim-silence`.
**caller_id** | string | Optional | The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as `name@company.com`.
**machine_detection_speech_threshold** | int | Optional | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400.
**machine_detection_speech_end_threshold** | int | Optional | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200.
**machine_detection_silence_timeout** | int | Optional | The number of milliseconds of initial silence after which an `unknown` AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000.
**async_amd** | string | Optional | Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: `true` or `false`.
**async_amd_status_callback** | string | Optional | The URL that we should call using the `async_amd_status_callback_method` to notify customer application whether the call was answered by human, machine or fax.
**async_amd_status_callback_method** | string | Optional | The HTTP method we should use when calling the `async_amd_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
**machine_detection_engine** | string | Optional | Select answering machine detection engine. Can be: `Lumenvox` or `Asterisk`. Default: `Lumenvox`.
**machine_detection_min_word_length** | int | Optional | The minimum duration in milliseconds of voice to be considered as a word. Default: 100.
**machine_detection_max_word_length** | int | Optional | The maximum duration in milliseconds of a word to accept. Default: 5000.
**machine_detection_words_silence** | int | Optional | The minimum duration in milliseconds of silence after a word to consider the audio what follows as a new word. Default: 50.
**machine_detection_max_num_of_words** | int | Optional | The maximum number of words in the greeting. If exceeded then it's considered as MACHINE. Default: 5.
**machine_detection_silence_threshold** | int | Optional | The silence threshold. Default: 256.
**byoc** | string | Optional | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta)
**call_reason** | string | Optional | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta)
**call_token** | string | Optional | A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. Pass an incoming call's call_token value to a forwarded call via the call_token parameter when creating a new call. A forwarded call should bear the same CallerID of the original incoming call.
**recording_track** | string | Optional | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio.
**time_limit** | int | Optional | The maximum duration of the call in seconds. Constraints depend on account and configuration.
**transcribe** | bool | Optional | Whether to transcribe the call recording. The default is `false`.
**transcription_configuration** | string | Optional | The SID or (unique) friendly name of the transcription configuration object to use for transcribing.
**url** | string | Optional | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).
**twiml** | string | Optional | TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both `twiml` and `url` are provided then `twiml` parameter will be ignored. Max 4000 characters.
**application_sid** | string | Optional | The SID of the Application resource that will handle the call, if the call will be handled by an application.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Call resource to update
**status** | string | Optional | 

## twilio_api_accounts_calls_feedback_summary

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**start_date** | string | **Required** | Only include feedback given on or after this date. Format is `YYYY-MM-DD` and specified in UTC.
**end_date** | string | **Required** | Only include feedback given on or before this date. Format is `YYYY-MM-DD` and specified in UTC.
**path_account_sid** | string | Optional | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**include_subaccounts** | bool | Optional | Whether to also include Feedback resources from all subaccounts. `true` includes feedback from all subaccounts and `false`, the default, includes feedback from only the specified account.
**status_callback** | string | Optional | The URL that we will request when the feedback summary is complete.
**status_callback_method** | string | Optional | The HTTP method (`GET` or `POST`) we use to make the request to the `StatusCallback` URL.
**sid** | string | *Computed* | A 34 character string that uniquely identifies this resource.

## twilio_api_accounts_calls_recordings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**call_sid** | string | **Required** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**recording_status_callback_event** | list(string) | Optional | The recording status events on which we should call the `recording_status_callback` URL. Can be: `in-progress`, `completed` and `absent` and the default is `completed`. Separate multiple event values with a space.
**recording_status_callback** | string | Optional | The URL we should call using the `recording_status_callback_method` on each recording event specified in  `recording_status_callback_event`. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback).
**recording_status_callback_method** | string | Optional | The HTTP method we should use to call `recording_status_callback`. Can be: `GET` or `POST` and the default is `POST`.
**trim** | string | Optional | Whether to trim any leading and trailing silence in the recording. Can be: `trim-silence` or `do-not-trim` and the default is `do-not-trim`. `trim-silence` trims the silence from the beginning and end of the recording and `do-not-trim` does not.
**recording_channels** | string | Optional | The number of channels used in the recording. Can be: `mono` or `dual` and the default is `mono`. `mono` records all parties of the call into one channel. `dual` records each party of a 2-party call into separate channels.
**play_beep** | bool | Optional | 
**recording_track** | string | Optional | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Recording resource to update.
**status** | string | Optional | 
**pause_behavior** | string | Optional | Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`.

## twilio_api_accounts_conferences_recordings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**conference_sid** | string | **Required** | The Conference SID that identifies the conference associated with the recording.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**recording_status_callback_event** | list(string) | Optional | The recording status events on which we should call the `recording_status_callback` URL. Can be: `in-progress`, `completed` and `absent` and the default is `completed`. Separate multiple event values with a space.
**recording_status_callback** | string | Optional | The URL we should call when the recording events specified in parameter `recording_status_callback_event` occur.
**recording_status_callback_method** | string | Optional | The HTTP method we should use to call `recording_status_callback`. Can be: `GET` or `POST` and the default is `POST`.
**trim** | string | Optional | Whether to trim any leading and trailing silence in the recording. Can be: `trim-silence` or `do-not-trim` and the default is `do-not-trim`. `trim-silence` trims the silence from the beginning and end of the recording and `do-not-trim` does not.
**play_beep** | bool | Optional | 
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use `Twilio.CURRENT` to reference the current active recording.
**status** | string | Optional | 
**pause_behavior** | string | Optional | Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`.

## twilio_api_accounts_incoming_phone_numbers

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**api_version** | string | Optional | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`.
**friendly_name** | string | Optional | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number.
**sms_application_sid** | string | Optional | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application.
**sms_fallback_method** | string | Optional | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
**sms_fallback_url** | string | Optional | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`.
**sms_method** | string | Optional | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`.
**sms_url** | string | Optional | The URL we should call when the new phone number receives an incoming SMS message.
**status_callback** | string | Optional | The URL we should call using the `status_callback_method` to send status information to your application.
**status_callback_method** | string | Optional | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**voice_application_sid** | string | Optional | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
**voice_caller_id_lookup** | bool | Optional | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`.
**voice_fallback_method** | string | Optional | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
**voice_method** | string | Optional | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`.
**voice_url** | string | Optional | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set.
**emergency_status** | string | Optional | 
**emergency_address_sid** | string | Optional | The SID of the emergency address configuration to use for emergency calling from the new phone number.
**trunk_sid** | string | Optional | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
**identity_sid** | string | Optional | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
**address_sid** | string | Optional | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
**voice_receive_mode** | string | Optional | 
**bundle_sid** | string | Optional | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
**phone_number** | string | Optional | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**area_code** | string | Optional | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an `area_code` or a `phone_number`.** (US and Canada only).
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update.
**account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).

## twilio_api_accounts_incoming_phone_numbers_assigned_add_ons

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**resource_sid** | string | **Required** | The SID of the Phone Number to assign the Add-on.
**installed_add_on_sid** | string | **Required** | The SID that identifies the Add-on installation.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the resource to fetch.

## twilio_api_accounts_messages

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**to** | string | **Required** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**status_callback** | string | Optional | The URL we should call using the `status_callback_method` to send status information to your application. If specified, we POST these message status changes to the URL: `queued`, `failed`, `sent`, `delivered`, or `undelivered`. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including `MessageSid`, `MessageStatus`, and `ErrorCode`. If you include this parameter with the `messaging_service_sid`, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed.
**application_sid** | string | Optional | The SID of the application that should receive message status. We POST a `message_sid` parameter and a `message_status` parameter with a value of `sent` or `failed` to the [application](https://www.twilio.com/docs/usage/api/applications)'s `message_status_callback`. If a `status_callback` parameter is also passed, it will be ignored and the application's `message_status_callback` parameter will be used.
**max_price** | float | Optional | The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds `max_price`, the message will fail and a status of `Failed` is sent to the status callback. If `MaxPrice` is not set, the message cost is not checked.
**provide_feedback** | bool | Optional | Whether to confirm delivery of the message. Set this value to `true` if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is `false` by default.
**attempt** | int | Optional | Total number of attempts made ( including this ) to send out the message regardless of the provider used
**validity_period** | int | Optional | How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds.
**max_rate** | string | Optional | 
**force_delivery** | bool | Optional | Reserved
**provider_sid** | string | Optional | 
**content_retention** | string | Optional | 
**address_retention** | string | Optional | 
**smart_encoded** | bool | Optional | Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: `true` or `false`.
**persistent_action** | list(string) | Optional | Rich actions for Channels Messages.
**interactive_data** | string | Optional | 
**force_opt_in** | bool | Optional | 
**shorten_urls** | bool | Optional | Determines the usage of Click Tracking. Setting it to `true` will instruct Twilio to replace all links in the Message with a shortened version based on the associated Domain Sid and track clicks on them. If this parameter is not set on an API call, we will use the value set on the Messaging Service. If this parameter is not set and the value is not configured on the Messaging Service used this will default to `false`.
**domain_sid** | string | Optional | Sets the Domain object used to shorten links with Click Tracking. Links in the message body will be replaced with shortened links on the specified domain. This parameter will override what is configured on the Messaging Service used. If this parameter is not set and the value is not set on the Messaging Service then shortening cannot be used.
**click_tracking_callback** | string | Optional | Sets the URL where Twilio will send a POST request each time an HTTP request is received to the shortened links in this message. If you include this parameter we will override what is configured on the Messaging Service. URLs must contain a valid hostname and underscores are not allowed.
**schedule_type** | string | Optional | 
**send_at** | string | Optional | The time that Twilio will send the message. Must be in ISO 8601 format.
**invoice_tag** | string | Optional | 
**send_as_mms** | bool | Optional | If set to True, Twilio will deliver the message as a single MMS message, regardless of the presence of media.
**content_sid** | string | Optional | The SID of the Content object returned at Content API content create time (https://www.twilio.com/docs/content-api/create-and-send-your-first-content-api-template#create-a-template). If this parameter is not specified, then the Content API will not be utilized.
**content_variables** | string | Optional | Key-value pairs of variable names to substitution values, used alongside a content_sid. If not specified, Content API will default to the default variables defined at create time.
**message_intent** | string | Optional | 
**tags** | string | Optional | A string containing a JSON map of key value pairs of tags to be recorded as metadata for the message. The object may contain up to 10 tags. Keys and values can each be up to 128 characters in length.
**from** | string | Optional | A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using `messaging_service_sid`, this parameter must be empty.
**messaging_service_sid** | string | Optional | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the `from` parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the `from` phone number for delivery.
**body** | string | Optional | The text of the message you want to send. Can be up to 1,600 characters in length.
**media_url** | list(string) | Optional | The URL of the media to send with the message. The media can be of type `gif`, `png`, and `jpeg` and will be formatted correctly on the recipient's device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple `media_url` parameters in the POST request. You can include up to 10 `media_url` parameters per message. You can send images in an SMS message in only the US and Canada.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Message resource to update.
**status** | string | Optional | 

## twilio_api_accounts_keys

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Key resource to update.

## twilio_api_accounts_signing_keys

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**sid** | string | *Computed* | 

## twilio_api_accounts_conferences_participants

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**conference_sid** | string | **Required** | The SID of the participant's conference.
**from** | string | **Required** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `from` must also be a phone number. If `to` is sip address, this value of `from` should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
**to** | string | **Required** | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as `sip:name@company.com`. Client identifiers are formatted `client:name`. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**status_callback** | string | Optional | The URL we should call using the `status_callback_method` to send status information to your application.
**status_callback_method** | string | Optional | The HTTP method we should use to call `status_callback`. Can be: `GET` and `POST` and defaults to `POST`.
**status_callback_event** | list(string) | Optional | The conference state changes that should generate a call to `status_callback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`.
**label** | string | Optional | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant.
**timeout** | int | Optional | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between `5` and `600`, inclusive. The default value is `60`. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds.
**record** | bool | Optional | Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`.
**muted** | bool | Optional | Whether the agent is muted in the conference. Can be `true` or `false` and the default is `false`.
**beep** | string | Optional | Whether to play a notification beep to the conference when the participant joins. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`.
**start_conference_on_enter** | bool | Optional | Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
**end_conference_on_exit** | bool | Optional | Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`.
**wait_url** | string | Optional | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
**wait_method** | string | Optional | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
**early_media** | bool | Optional | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: `true` or `false` and defaults to `true`.
**max_participants** | int | Optional | The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`.
**conference_record** | string | Optional | Whether to record the conference the participant is joining. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`.
**conference_trim** | string | Optional | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`.
**conference_status_callback** | string | Optional | The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored.
**conference_status_callback_method** | string | Optional | The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**conference_status_callback_event** | list(string) | Optional | The conference state changes that should generate a call to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `modify`, `speaker`, and `announcement`. Separate multiple values with a space. Defaults to `start end`.
**recording_channels** | string | Optional | The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`.
**recording_status_callback** | string | Optional | The URL that we should call using the `recording_status_callback_method` when the recording status changes.
**recording_status_callback_method** | string | Optional | The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**sip_auth_username** | string | Optional | The SIP username used for authentication.
**sip_auth_password** | string | Optional | The SIP password for authentication.
**region** | string | Optional | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`.
**conference_recording_status_callback** | string | Optional | The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available.
**conference_recording_status_callback_method** | string | Optional | The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**recording_status_callback_event** | list(string) | Optional | The recording state changes that should generate a call to `recording_status_callback`. Can be: `started`, `in-progress`, `paused`, `resumed`, `stopped`, `completed`, `failed`, and `absent`. Separate multiple values with a space, ex: `'in-progress completed failed'`.
**conference_recording_status_callback_event** | list(string) | Optional | The conference recording state changes that generate a call to `conference_recording_status_callback`. Can be: `in-progress`, `completed`, `failed`, and `absent`. Separate multiple values with a space, ex: `'in-progress completed failed'`
**coaching** | bool | Optional | Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined.
**call_sid_to_coach** | string | Optional | The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`.
**jitter_buffer_size** | string | Optional | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant's audio is mixed into the conference. Can be: `off`, `small`, `medium`, and `large`. Default to `large`.
**byoc** | string | Optional | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta)
**caller_id** | string | Optional | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `callerId` must also be a phone number. If `to` is sip address, this value of `callerId` should be a username portion to be used to populate the From header that is passed to the SIP endpoint.
**call_reason** | string | Optional | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta)
**recording_track** | string | Optional | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is sent from Twilio. `both` records the audio that is received and sent by Twilio.
**time_limit** | int | Optional | The maximum duration of the call in seconds. Constraints depend on account and configuration.
**machine_detection** | string | Optional | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: `Enable` or `DetectMessageEnd`. Use `Enable` if you would like us to return `AnsweredBy` as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. If `send_digits` is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection).
**machine_detection_timeout** | int | Optional | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with `AnsweredBy` of `unknown`. The default timeout is 30 seconds.
**machine_detection_speech_threshold** | int | Optional | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400.
**machine_detection_speech_end_threshold** | int | Optional | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200.
**machine_detection_silence_timeout** | int | Optional | The number of milliseconds of initial silence after which an `unknown` AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000.
**amd_status_callback** | string | Optional | The URL that we should call using the `amd_status_callback_method` to notify customer application whether the call was answered by human, machine or fax.
**amd_status_callback_method** | string | Optional | The HTTP method we should use when calling the `amd_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
**machine_detection_engine** | string | Optional | Select answering machine detection engine. Can be: `Lumenvox` or `Asterisk`. Default: `Lumenvox`.
**machine_detection_min_word_length** | int | Optional | The minimum duration in milliseconds of voice to be considered as a word. Default: 100.
**machine_detection_max_word_length** | int | Optional | The maximum duration in milliseconds of a word to accept. Default: 5000.
**machine_detection_words_silence** | int | Optional | The minimum duration in milliseconds of silence after a word to consider the audio what follows as a new word. Default: 50.
**machine_detection_max_num_of_words** | int | Optional | The maximum number of words in the greeting. If exceeded, then it's considered as MACHINE. Default: 5.
**machine_detection_silence_threshold** | int | Optional | The silence threshold. Default: 256.
**call_sid** | string | *Computed* | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.
**hold** | bool | Optional | Whether the participant should be on hold. Can be: `true` or `false`. `true` puts the participant on hold, and `false` lets them rejoin the conference.
**hold_url** | string | Optional | The URL we call using the `hold_method` for music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.
**hold_method** | string | Optional | The HTTP method we should use to call `hold_url`. Can be: `GET` or `POST` and the default is `GET`.
**announce_url** | string | Optional | The URL we call using the `announce_method` for an announcement to the participant. The URL may return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.
**announce_method** | string | Optional | The HTTP method we should use to call `announce_url`. Can be: `GET` or `POST` and defaults to `POST`.
**beep_on_exit** | bool | Optional | Whether to play a notification beep to the conference when the participant exits. Can be: `true` or `false`.

## twilio_api_accounts_queues

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you created to describe this resource. It can be up to 64 characters long.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**max_size** | int | Optional | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the Queue resource to update

## twilio_api_accounts_sip_domains_auth_calls_credential_list_mappings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_sid** | string | **Required** | The SID of the SIP domain that will contain the new resource.
**credential_list_sid** | string | **Required** | The SID of the CredentialList resource to map to the SIP domain.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

## twilio_api_accounts_sip_domains_auth_calls_ip_access_control_list_mappings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_sid** | string | **Required** | The SID of the SIP domain that will contain the new resource.
**ip_access_control_list_sid** | string | **Required** | The SID of the IpAccessControlList resource to map to the SIP domain.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch.

## twilio_api_accounts_sip_domains_auth_registrations_credential_list_mappings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_sid** | string | **Required** | The SID of the SIP domain that will contain the new resource.
**credential_list_sid** | string | **Required** | The SID of the CredentialList resource to map to the SIP domain.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

## twilio_api_accounts_sip_credential_lists_credentials

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**credential_list_sid** | string | **Required** | The unique id that identifies the credential list to include the created credential.
**username** | string | **Required** | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio's challenge of the initial INVITE. It can be up to 32 characters long.
**password** | string | **Required** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)
**path_account_sid** | string | Optional | The unique id of the Account that is responsible for this resource.
**sid** | string | *Computed* | The unique id that identifies the resource to update.

## twilio_api_accounts_sip_credential_lists

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A human readable descriptive text that describes the CredentialList, up to 64 characters long.
**path_account_sid** | string | Optional | The unique id of the Account that is responsible for this resource.
**sid** | string | *Computed* | The credential list Sid that uniquely identifies this resource

## twilio_api_accounts_sip_domains_credential_list_mappings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_sid** | string | **Required** | A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped.
**credential_list_sid** | string | **Required** | A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain.
**path_account_sid** | string | Optional | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**sid** | string | *Computed* | A 34 character string that uniquely identifies the resource to fetch.

## twilio_api_accounts_sip_domains

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_name** | string | **Required** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`.
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**friendly_name** | string | Optional | A descriptive string that you created to describe the resource. It can be up to 64 characters long.
**voice_url** | string | Optional | The URL we should when the domain receives a call.
**voice_method** | string | Optional | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`.
**voice_fallback_method** | string | Optional | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**voice_status_callback_url** | string | Optional | The URL that we should call to pass status parameters (such as call ended) to your application.
**voice_status_callback_method** | string | Optional | The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`.
**sip_registration** | bool | Optional | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not.
**emergency_calling_enabled** | bool | Optional | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
**secure** | bool | Optional | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
**byoc_trunk_sid** | string | Optional | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
**emergency_caller_sid** | string | Optional | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the SipDomain resource to update.

## twilio_api_accounts_sip_ip_access_control_lists

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A human readable descriptive text that describes the IpAccessControlList, up to 255 characters long.
**path_account_sid** | string | Optional | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**sid** | string | *Computed* | A 34 character string that uniquely identifies the resource to udpate.

## twilio_api_accounts_sip_domains_ip_access_control_list_mappings

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**domain_sid** | string | **Required** | A 34 character string that uniquely identifies the SIP domain.
**ip_access_control_list_sid** | string | **Required** | The unique id of the IP access control list to map to the SIP domain.
**path_account_sid** | string | Optional | The unique id of the Account that is responsible for this resource.
**sid** | string | *Computed* | A 34 character string that uniquely identifies the resource to fetch.

## twilio_api_accounts_sip_ip_access_control_lists_ip_addresses

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**ip_access_control_list_sid** | string | **Required** | The IpAccessControlList Sid with which to associate the created IpAddress resource.
**friendly_name** | string | **Required** | A human readable descriptive text for this resource, up to 255 characters long.
**ip_address** | string | **Required** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
**path_account_sid** | string | Optional | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**cidr_prefix_length** | int | Optional | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
**sid** | string | *Computed* | A 34 character string that identifies the IpAddress resource to update.

## twilio_api_accounts_usage_triggers

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**callback_url** | string | **Required** | The URL we should call using `callback_method` when the trigger fires.
**trigger_value** | string | **Required** | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as `+30` to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a `+` as `%2B`.
**usage_category** | string | **Required** | 
**path_account_sid** | string | Optional | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**callback_method** | string | Optional | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**recurring** | string | Optional | 
**trigger_by** | string | Optional | 
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update.

