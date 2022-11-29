
## twilio_sync_services_documents_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the Sync Document
**data** | string | Optional | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
**ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (the Sync Document's time-to-live).
**sid** | string | *Computed* | The SID of the Document resource to update. Can be the Document resource's `sid` or its `unique_name`.
**if_match** | string | Optional | The If-Match HTTP request header

## twilio_sync_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**friendly_name** | string | Optional | A string that you assign to describe the resource.
**webhook_url** | string | Optional | The URL we should call when Sync objects are manipulated.
**reachability_webhooks_enabled** | bool | Optional | Whether the service instance should call `webhook_url` when client endpoints connect to Sync. The default is `false`.
**acl_enabled** | bool | Optional | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
**reachability_debouncing_enabled** | bool | Optional | Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event.
**reachability_debouncing_window** | int | Optional | The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the `webhook_url` is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to `webhook_url`.
**webhooks_from_rest_enabled** | bool | Optional | Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`.
**sid** | string | *Computed* | The SID of the Service resource to update.

## twilio_sync_services_lists_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
**ttl** | int | Optional | Alias for collection_ttl. If both are provided, this value is ignored.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
**sid** | string | *Computed* | The SID of the Sync List resource to update. Can be the Sync List resource's `sid` or its `unique_name`.

## twilio_sync_services_lists_items_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in.
**list_sid** | string | **Required** | The SID of the Sync List to add the new List Item to. Can be the Sync List resource's `sid` or its `unique_name`.
**data** | string | **Required** | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**ttl** | int | Optional | An alias for `item_ttl`. If both parameters are provided, this value is ignored.
**item_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted.
**index** | int | *Computed* | The index of the Sync List Item resource to update.
**if_match** | string | Optional | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).

## twilio_sync_services_maps_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource.
**ttl** | int | Optional | An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
**sid** | string | *Computed* | The SID of the Sync Map resource to update. Can be the Sync Map's `sid` or its `unique_name`.

## twilio_sync_services_maps_items_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in.
**map_sid** | string | **Required** | The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource's `sid` or its `unique_name`.
**key** | string | **Required** | The unique, user-defined key for the Map Item. Can be up to 320 characters long.
**data** | string | **Required** | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**ttl** | int | Optional | An alias for `item_ttl`. If both parameters are provided, this value is ignored.
**item_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted.
**if_match** | string | Optional | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).

## twilio_sync_services_streams_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
**ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
**sid** | string | *Computed* | The SID of the Stream resource to update.

