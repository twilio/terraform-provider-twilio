
## twilio_sync_services_documents_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in.
**data** | string | Optional | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
**ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (the Sync Document&#39;s time-to-live).
**unique_name** | string | Optional | An application-defined string that uniquely identifies the Sync Document
**sid** | string | *Computed* | The SID of the Document resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**if_match** | string | Optional | The If-Match HTTP request header
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**created_by** | string | *Computed* | The identity of the Sync Document&#39;s creator
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_expires** | string | *Computed* | The ISO 8601 date and time in GMT when the Sync Document expires
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of resources related to the Sync Document
**revision** | string | *Computed* | The current revision of the Sync Document, represented by a string identifier
**url** | string | *Computed* | The absolute URL of the Document resource

## twilio_sync_services_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**acl_enabled** | bool | Optional | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
**friendly_name** | string | Optional | A string that you assign to describe the resource.
**reachability_debouncing_enabled** | bool | Optional | Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event.
**reachability_debouncing_window** | int | Optional | The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the &#x60;webhook_url&#x60; is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to &#x60;webhook_url&#x60;.
**reachability_webhooks_enabled** | bool | Optional | Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;.
**webhook_url** | string | Optional | The URL we should call when Sync objects are manipulated.
**webhooks_from_rest_enabled** | bool | Optional | Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;.
**sid** | string | *Computed* | The SID of the Service resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of related resources
**unique_name** | string | *Computed* | An application-defined string that uniquely identifies the resource
**url** | string | *Computed* | The absolute URL of the Service resource

## twilio_sync_services_lists_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
**ttl** | int | Optional | Alias for collection_ttl. If both are provided, this value is ignored.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.
**sid** | string | *Computed* | The SID of the Sync List resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**created_by** | string | *Computed* | The identity of the Sync List&#39;s creator
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_expires** | string | *Computed* | The ISO 8601 date and time in GMT when the Sync List expires
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of the Sync List&#39;s nested resources
**revision** | string | *Computed* | The current revision of the Sync List, represented as a string
**url** | string | *Computed* | The absolute URL of the Sync List resource

## twilio_sync_services_lists_items_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in.
**list_sid** | string | **Required** | The SID of the Sync List to add the new List Item to. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**data** | string | **Required** | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.
**item_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**ttl** | int | Optional | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.
**index** | int | *Computed* | The index of the Sync List Item resource to update.
**if_match** | string | Optional | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**created_by** | string | *Computed* | The identity of the List Item&#39;s creator
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_expires** | string | *Computed* | The ISO 8601 date and time in GMT when the List Item expires
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**revision** | string | *Computed* | The current revision of the item, represented as a string
**url** | string | *Computed* | The absolute URL of the List Item resource

## twilio_sync_services_maps_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
**ttl** | int | Optional | An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored.
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.
**sid** | string | *Computed* | The SID of the Sync Map resource to update. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**created_by** | string | *Computed* | The identity of the Sync Map&#39;s creator
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_expires** | string | *Computed* | The ISO 8601 date and time in GMT when the Sync Map expires
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of the Sync Map&#39;s nested resources
**revision** | string | *Computed* | The current revision of the Sync Map, represented as a string
**url** | string | *Computed* | The absolute URL of the Sync Map resource

## twilio_sync_services_maps_items_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in.
**map_sid** | string | **Required** | The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**data** | string | **Required** | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**key** | string | **Required** | The unique, user-defined key for the Map Item. Can be up to 320 characters long.
**collection_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.
**item_ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**ttl** | int | Optional | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.
**if_match** | string | Optional | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**created_by** | string | *Computed* | The identity of the Map Item&#39;s creator
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_expires** | string | *Computed* | The ISO 8601 date and time in GMT when the Map Item expires
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**revision** | string | *Computed* | The current revision of the Map Item, represented as a string
**url** | string | *Computed* | The absolute URL of the Map Item resource

## twilio_sync_services_streams_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**service_sid** | string | **Required** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in.
**ttl** | int | Optional | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
**unique_name** | string | Optional | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.
**sid** | string | *Computed* | The SID of the Stream resource to update.
**account_sid** | string | *Computed* | The SID of the Account that created the resource
**created_by** | string | *Computed* | The Identity of the Stream&#39;s creator
**date_created** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was created
**date_expires** | string | *Computed* | The ISO 8601 date and time in GMT when the Message Stream expires
**date_updated** | string | *Computed* | The ISO 8601 date and time in GMT when the resource was last updated
**links** | string | *Computed* | The URLs of the Stream&#39;s nested resources
**url** | string | *Computed* | The absolute URL of the Message Stream resource

