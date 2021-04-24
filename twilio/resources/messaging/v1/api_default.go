/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/twilio/common"
	. "github.com/twilio/twilio-go/twilio/rest/messaging/v1"
)

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"friendly_name":                 AsString(SchemaRequired),
			"area_code_geomatch":            AsString(SchemaOptional),
			"fallback_method":               AsString(SchemaOptional),
			"fallback_to_long_code":         AsString(SchemaOptional),
			"fallback_url":                  AsString(SchemaOptional),
			"inbound_method":                AsString(SchemaOptional),
			"inbound_request_url":           AsString(SchemaOptional),
			"mms_converter":                 AsString(SchemaOptional),
			"scan_message_content":          AsString(SchemaOptional),
			"smart_encoding":                AsString(SchemaOptional),
			"status_callback":               AsString(SchemaOptional),
			"sticky_sender":                 AsString(SchemaOptional),
			"synchronous_validation":        AsString(SchemaOptional),
			"use_inbound_webhook_on_number": AsString(SchemaOptional),
			"validity_period":               AsString(SchemaOptional),
			"account_sid":                   AsString(SchemaComputed),
			"date_created":                  AsString(SchemaComputed),
			"date_updated":                  AsString(SchemaComputed),
			"links":                         AsString(SchemaComputed),
			"sid":                           AsString(SchemaComputed),
			"url":                           AsString(SchemaComputed),
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.MessagingV1.CreateService(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.MessagingV1.DeleteService(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
