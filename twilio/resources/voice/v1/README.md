
## twilio_voice_byoc_trunks_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**cnam_lookup_enabled** | bool | Optional | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**connection_policy_sid** | string | Optional | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**from_domain_sid** | string | Optional | The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;.
**status_callback_method** | string | Optional | The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**status_callback_url** | string | Optional | The URL that we should call to pass status parameters (such as call ended) to your application.
**voice_fallback_method** | string | Optional | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_fallback_url** | string | Optional | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;.
**voice_method** | string | Optional | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**voice_url** | string | Optional | The URL we should call when the BYOC Trunk receives a call.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT that the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT that the resource was last updated
**url** | string | *Computed* | The absolute URL of the resource

## twilio_voice_connection_policies_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**sid** | string | *Computed* | The unique string that we created to identify the Connection Policy resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of related resources
**url** | string | *Computed* | The absolute URL of the resource

## twilio_voice_connection_policies_targets_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**connection_policy_sid** | string | **Required** | The SID of the Connection Policy that owns the Target.
**target** | string | **Required** | The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported.
**enabled** | bool | Optional | Whether the Target is enabled. The default is &#x60;true&#x60;.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**priority** | int | Optional | The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target.
**weight** | int | Optional | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority.
**sid** | string | *Computed* | The unique string that we created to identify the Target resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT when the resource was last updated
**url** | string | *Computed* | The absolute URL of the resource

## twilio_voice_ip_records_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**ip_address** | string | **Required** | An IP address in dotted decimal notation, IPv4 only.
**cidr_prefix_length** | int | Optional | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.
**friendly_name** | string | Optional | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IP Record resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT that the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT that the resource was last updated
**url** | string | *Computed* | The absolute URL of the resource

## twilio_voice_source_ip_mappings_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**ip_record_sid** | string | **Required** | The Twilio-provided string that uniquely identifies the IP Record resource to map from.
**sip_domain_sid** | string | **Required** | The SID of the SIP Domain that the IP Record should be mapped to.
**sid** | string | *Computed* | The Twilio-provided string that uniquely identifies the IP Record resource to update.
**date_created** | string | *Computed* | The RFC 2822 date and time in GMT that the resource was created
**date_updated** | string | *Computed* | The RFC 2822 date and time in GMT that the resource was last updated
**url** | string | *Computed* | The absolute URL of the resource

