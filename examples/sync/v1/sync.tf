terraform {
  required_providers {
    twilio = {
      source  = "twilio.com/twilio/twilio"
      version = ">=0.2.0"
    }
  }
}

provider "twilio" {
  //  account_sid defaults to TWILIO_ACCOUNT_SID env var
  //  auth_token  defaults to TWILIO_AUTH_TOKEN env var
}

# Step 1: Create Sync Services for dev, stage and prod
resource "twilio_sync_services_v1" "service_dev" {
  friendly_name                   = "Dev Sync"
  acl_enabled                     = true
  webhook_url                     = "https://example.com/dev"
  reachability_webhooks_enabled   = true
  reachability_debouncing_enabled = true
  reachability_debouncing_window  = 20000
}

resource "twilio_sync_services_v1" "service_stage" {
  friendly_name                   = "Stage Sync"
  acl_enabled                     = true
  webhook_url                     = "https://example.com/stage"
  reachability_webhooks_enabled   = true
  reachability_debouncing_enabled = true
  reachability_debouncing_window  = 20000
}

resource "twilio_sync_services_v1" "service_prod" {
  friendly_name                   = "Prod Sync"
  acl_enabled                     = true
  webhook_url                     = "https://example.com/prod"
  reachability_webhooks_enabled   = true
  reachability_debouncing_enabled = true
  reachability_debouncing_window  = 20000
}

# Step 2: Create Sync Document for each service
resource "twilio_sync_services_documents_v1" "doc_dev" {
  service_sid = twilio_sync_services_v1.service_dev.id
  data = jsonencode({
    state : "new",
    count : 1
  })
  ttl         = 3600
  unique_name = "sync_doc_dev"
}

resource "twilio_sync_services_documents_v1" "doc_stage" {
  service_sid = twilio_sync_services_v1.service_stage.id
  data = jsonencode({
    state : "new",
    count : 1
  })
  ttl         = 3600
  unique_name = "sync_doc_stage"
}

resource "twilio_sync_services_documents_v1" "doc_prod" {
  service_sid = twilio_sync_services_v1.service_prod.id
  data = jsonencode({
    state : "new",
    count : 1
  })
  ttl         = 3600
  unique_name = "sync_doc_prod"
}

# Step 3: Create Sync List for each of the services
resource "twilio_sync_services_lists_v1" "list_dev" {
  service_sid    = twilio_sync_services_v1.service_dev.id
  collection_ttl = 3600
  ttl            = 7200
  unique_name    = "sync_list_dev"
}

resource "twilio_sync_services_lists_v1" "list_stage" {
  service_sid    = twilio_sync_services_v1.service_stage.id
  collection_ttl = 3600
  ttl            = 7200
  unique_name    = "sync_list_stage"
}

resource "twilio_sync_services_lists_v1" "list_prod" {
  service_sid    = twilio_sync_services_v1.service_prod.id
  collection_ttl = 3600
  ttl            = 7200
  unique_name    = "sync_list_prod"
}

# Create two items in the list
resource "twilio_sync_services_lists_items_v1" "list_items_1_dev" {
  service_sid    = twilio_sync_services_v1.service_dev.id
  list_sid       = twilio_sync_services_lists_v1.list_dev.id
  collection_ttl = 3600
  data = jsonencode({
    state : "init"
  })
  item_ttl = 600
  ttl      = 7200
}

resource "twilio_sync_services_lists_items_v1" "list_items_2_dev" {
  service_sid    = twilio_sync_services_v1.service_dev.id
  list_sid       = twilio_sync_services_lists_v1.list_dev.id
  collection_ttl = 3600
  data = jsonencode({
    state : "in-progress"
  })
  item_ttl = 600
  ttl      = 7200
}

resource "twilio_sync_services_lists_items_v1" "list_items_stage" {
  service_sid    = twilio_sync_services_v1.service_stage.id
  list_sid       = twilio_sync_services_lists_v1.list_stage.id
  collection_ttl = 3600
  data = jsonencode({
    state : "init"
  })
  item_ttl = 600
  ttl      = 7200
}

resource "twilio_sync_services_lists_items_v1" "list_items_prod" {
  service_sid    = twilio_sync_services_v1.service_prod.id
  list_sid       = twilio_sync_services_lists_v1.list_prod.id
  collection_ttl = 3600
  data = jsonencode({
    state : "init"
  })
  item_ttl = 600
  ttl      = 7200
}

# Step 4: Create Sync Map
resource "twilio_sync_services_maps_v1" "map_dev" {
  service_sid    = twilio_sync_services_v1.service_dev.id
  collection_ttl = 3600
  ttl            = 7200
  unique_name    = "sync_map_dev"
}


resource "twilio_sync_services_maps_v1" "map_stage" {
  service_sid    = twilio_sync_services_v1.service_stage.id
  collection_ttl = 3600
  ttl            = 7200
  unique_name    = "sync_map_stage"
}

resource "twilio_sync_services_maps_v1" "map_prod" {
  service_sid    = twilio_sync_services_v1.service_prod.id
  collection_ttl = 3600
  ttl            = 7200
  unique_name    = "sync_map_prod"
}

resource "twilio_sync_services_maps_items_v1" "map_items_1_dev" {
  service_sid = twilio_sync_services_v1.service_dev.id
  map_sid     = twilio_sync_services_maps_v1.map_dev.id
  data = jsonencode({
    field1 : "something"
    field2 : {
      field2_1 : "something else"
    }
  })
  key            = "key_dev_item_1"
  collection_ttl = 3600
  item_ttl       = 1800
  ttl            = 1800
}

resource "twilio_sync_services_maps_items_v1" "map_items_1_stage" {
  service_sid = twilio_sync_services_v1.service_stage.id
  map_sid     = twilio_sync_services_maps_v1.map_stage.id
  data = jsonencode({
    field1 : "something"
    field2 : {
      field2_1 : "something else"
    }
  })
  key            = "key_stage_item_1"
  collection_ttl = 3600
  item_ttl       = 1800
  ttl            = 1800
}

resource "twilio_sync_services_maps_items_v1" "map_items_2_stage" {
  service_sid = twilio_sync_services_v1.service_stage.id
  map_sid     = twilio_sync_services_maps_v1.map_stage.id
  data = jsonencode({
    field1 : "something"
    field2 : {
      field2_1 : "something else"
    }
  })
  key            = "key_stage_item_2"
  collection_ttl = 3600
  item_ttl       = 1800
  ttl            = 1800
}

resource "twilio_sync_services_maps_items_v1" "map_items_1_prod" {
  service_sid = twilio_sync_services_v1.service_prod.id
  map_sid     = twilio_sync_services_maps_v1.map_prod.id
  data = jsonencode({
    field1 : "something"
    field2 : {
      field2_1 : "something else"
    }
  })
  key            = "key_prod_item_1"
  collection_ttl = 3600
  item_ttl       = 1800
  ttl            = 1800
}

# Step 5: Create Sync Streams
resource "twilio_sync_services_streams_v1" "stream_dev" {
  service_sid = twilio_sync_services_v1.service_dev.id
  ttl         = 3600
  unique_name = "sync_streams_dev"
}

resource "twilio_sync_services_streams_v1" "stream_stage" {
  service_sid = twilio_sync_services_v1.service_stage.id
  ttl         = 3600
  unique_name = "sync_streams_stage"
}

resource "twilio_sync_services_streams_v1" "stream_prod" {
  service_sid = twilio_sync_services_v1.service_prod.id
  ttl         = 3600
  unique_name = "sync_streams_prod"
}

output "service_dev" {
  value = twilio_sync_services_v1.service_dev
}

output "service_stage" {
  value = twilio_sync_services_v1.service_stage
}

output "service_prod" {
  value = twilio_sync_services_v1.service_prod
}
