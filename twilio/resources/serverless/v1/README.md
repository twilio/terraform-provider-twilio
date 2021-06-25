
## twilio_serverless_services_assets_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Asset resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.
**sid** | string | *Computed* | The SID that identifies the Asset resource to update.

## twilio_serverless_services_functions_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Function resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.
**sid** | string | *Computed* | The SID of the Function resource to update.

## twilio_serverless_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters.
**unique_name** | string | **Required** | A user-defined string that uniquely identifies the Service resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the Service resource. This value must be 50 characters or less in length and be unique.
**include_credentials** | bool | Optional | Whether to inject Account credentials into a function invocation context. The default value is &#x60;true&#x60;.
**ui_editable** | bool | Optional | Whether the Service&#39;s properties and subresources can be edited via the UI. The default value is &#x60;false&#x60;.
**sid** | string | *Computed* | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to update.

## twilio_serverless_services_environments_variables_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Variable resource under.
**environment_sid** | string | **Required** | The SID of the Environment in which the Variable resource exists.
**key** | string | **Required** | A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
**value** | string | **Required** | A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.
**sid** | string | *Computed* | The SID of the Variable resource to update.

