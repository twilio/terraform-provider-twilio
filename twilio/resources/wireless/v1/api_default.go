/*
 * Twilio - Wireless
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
	. "github.com/twilio/twilio-go/rest/wireless/v1"
)

func ResourceRatePlans() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRatePlans,
		ReadContext:   readRatePlans,
		UpdateContext: updateRatePlans,
		DeleteContext: deleteRatePlans,
		Schema: map[string]*schema.Schema{
			"data_enabled":                     AsBool(SchemaComputedOptional),
			"data_limit":                       AsInt(SchemaComputedOptional),
			"data_metering":                    AsString(SchemaComputedOptional),
			"friendly_name":                    AsString(SchemaComputedOptional),
			"international_roaming":            AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"international_roaming_data_limit": AsInt(SchemaComputedOptional),
			"messaging_enabled":                AsBool(SchemaComputedOptional),
			"national_roaming_data_limit":      AsInt(SchemaComputedOptional),
			"national_roaming_enabled":         AsBool(SchemaComputedOptional),
			"unique_name":                      AsString(SchemaComputedOptional),
			"voice_enabled":                    AsBool(SchemaComputedOptional),
			"sid":                              AsString(SchemaComputed),
		},
	}
}

func createRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRatePlanParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.WirelessV1.CreateRatePlan(&params)
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

func deleteRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.WirelessV1.DeleteRatePlan(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.WirelessV1.FetchRatePlan(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func updateRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRatePlanParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.WirelessV1.UpdateRatePlan(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
