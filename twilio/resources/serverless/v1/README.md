
## twilio_serverless_services_assets_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Asset resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.
**sid** | string | *Computed* | The SID that identifies the Asset resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the Asset resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Asset resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Asset resource was last updated
**links** | string | *Computed* | The URLs of the Asset resource&#39;s nested resources
**url** | string | *Computed* | The absolute URL of the Asset resource

## twilio_serverless_services_builds_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Build resource under.
**asset_versions** | list(string) | Optional | The list of Asset Version resource SIDs to include in the Build.
**dependencies** | string | Optional | A list of objects that describe the Dependencies included in the Build. Each object contains the &#x60;name&#x60; and &#x60;version&#x60; of the dependency.
**function_versions** | list(string) | Optional | The list of the Function Version resource SIDs to include in the Build.
**runtime** | string | Optional | The Runtime version that will be used to run the Build resource when it is deployed.
**sid** | string | *Computed* | The SID of the Build resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the Build resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Build resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Build resource was last updated
**links** | string | *Computed* | 
**status** | string | *Computed* | The status of the Build
**url** | string | *Computed* | The absolute URL of the Build resource

## twilio_serverless_services_environments_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Environment resource under.
**unique_name** | string | **Required** | A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters.
**domain_suffix** | string | Optional | A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters.
**sid** | string | *Computed* | The SID of the Environment resource to fetch.
**account_sid** | string | *Computed* | The SID of the Account that created the Environment resource
**build_sid** | string | *Computed* | The SID of the build deployed in the environment
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Environment resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Environment resource was last updated
**domain_name** | string | *Computed* | The domain name for all Functions and Assets deployed in the Environment
**links** | string | *Computed* | The URLs of the Environment resource&#39;s nested resources
**url** | string | *Computed* | The absolute URL of the Environment resource

## twilio_serverless_services_functions_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Function resource under.
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.
**sid** | string | *Computed* | The SID of the Function resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the Function resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Function resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Function resource was last updated
**links** | string | *Computed* | The URLs of nested resources of the Function resource
**url** | string | *Computed* | The absolute URL of the Function resource

## twilio_serverless_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | **Required** | A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters.
**unique_name** | string | **Required** | A user-defined string that uniquely identifies the Service resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the Service resource. This value must be 50 characters or less in length and be unique.
**include_credentials** | bool | Optional | Whether to inject Account credentials into a function invocation context. The default value is &#x60;true&#x60;.
**ui_editable** | bool | Optional | Whether the Service&#39;s properties and subresources can be edited via the UI. The default value is &#x60;false&#x60;.
**sid** | string | *Computed* | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the Service resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Service resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Service resource was last updated
**domain_base** | string | *Computed* | The base domain name for this Service, which is a combination of the unique name and a randomly generated string
**links** | string | *Computed* | The URLs of the Service&#39;s nested resources
**url** | string | *Computed* | The absolute URL of the Service resource

## twilio_serverless_services_environments_variables_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the Service to create the Variable resource under.
**environment_sid** | string | **Required** | The SID of the Environment in which the Variable resource exists.
**key** | string | **Required** | A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
**value** | string | **Required** | A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.
**sid** | string | *Computed* | The SID of the Variable resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the Variable resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the Variable resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the Variable resource was last updated
**url** | string | *Computed* | The absolute URL of the Variable resource

