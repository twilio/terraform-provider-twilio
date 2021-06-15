/*
 * Twilio - Video
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
	. "github.com/twilio/twilio-go/rest/video/v1"
)

func ResourceCompositionHooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCompositionHooks,
		ReadContext:   readCompositionHooks,
		UpdateContext: updateCompositionHooks,
		DeleteContext: deleteCompositionHooks,
		Schema: map[string]*schema.Schema{
			"audio_sources":          AsString(SchemaOptional),
			"audio_sources_excluded": AsString(SchemaOptional),
			"enabled":                AsString(SchemaOptional),
			"format":                 AsString(SchemaOptional),
			"friendly_name":          AsString(SchemaOptional),
			"resolution":             AsString(SchemaOptional),
			"status_callback":        AsString(SchemaOptional),
			"status_callback_method": AsString(SchemaOptional),
			"trim":                   AsString(SchemaOptional),
			"video_layout":           AsString(SchemaOptional),
			"account_sid":            AsString(SchemaComputed),
			"date_created":           AsString(SchemaComputed),
			"date_updated":           AsString(SchemaComputed),
			"sid":                    AsString(SchemaComputed),
			"url":                    AsString(SchemaComputed),
		},
	}
}

func createCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCompositionHookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VideoV1.CreateCompositionHook(&params)
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

func deleteCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VideoV1.DeleteCompositionHook(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VideoV1.FetchCompositionHook(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCompositionHookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VideoV1.UpdateCompositionHook(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
