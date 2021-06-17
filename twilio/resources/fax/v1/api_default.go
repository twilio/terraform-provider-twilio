/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
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
	. "github.com/twilio/twilio-go/rest/fax/v1"
)

func ResourceFaxes() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFaxes,
		ReadContext:   readFaxes,
		UpdateContext: updateFaxes,
		DeleteContext: deleteFaxes,
		Schema: map[string]*schema.Schema{
			"from":              AsString(SchemaComputedOptional),
			"media_url":         AsString(SchemaComputedOptional),
			"quality":           AsString(SchemaComputedOptional),
			"sip_auth_password": AsString(SchemaComputedOptional),
			"sip_auth_username": AsString(SchemaComputedOptional),
			"status_callback":   AsString(SchemaComputedOptional),
			"store_media":       AsBool(SchemaComputedOptional),
			"to":                AsString(SchemaComputedOptional),
			"ttl":               AsInt(SchemaComputedOptional),
			"account_sid":       AsString(SchemaComputed),
			"api_version":       AsString(SchemaComputed),
			"date_created":      AsString(SchemaComputed),
			"date_updated":      AsString(SchemaComputed),
			"direction":         AsString(SchemaComputed),
			"duration":          AsInt(SchemaComputed),
			"links":             AsString(SchemaComputed),
			"media_sid":         AsString(SchemaComputed),
			"num_pages":         AsInt(SchemaComputed),
			"price":             AsFloat(SchemaComputed),
			"price_unit":        AsString(SchemaComputed),
			"sid":               AsString(SchemaComputed),
			"status":            AsString(SchemaComputedOptional),
			"url":               AsString(SchemaComputed),
		},
	}
}

func createFaxes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFaxParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.FaxV1.CreateFax(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Sid))
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteFaxes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.FaxV1.DeleteFax(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readFaxes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FaxV1.FetchFax(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateFaxes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFaxParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FaxV1.UpdateFax(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
