/*
 * Twilio - Studio
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
	. "github.com/twilio/twilio-go/rest/studio/v1"
)

func ResourceFlowsExecutions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlowsExecutions,
		ReadContext:   readFlowsExecutions,
		UpdateContext: updateFlowsExecutions,
		DeleteContext: deleteFlowsExecutions,
		Schema: map[string]*schema.Schema{
			"flow_sid":   AsString(SchemaRequired),
			"from":       AsString(SchemaRequired),
			"to":         AsString(SchemaRequired),
			"parameters": AsString(SchemaComputedOptional),
			"sid":        AsString(SchemaComputed),
			"status":     AsString(SchemaComputedOptional),
		},
	}
}

func createFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateExecutionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.CreateExecution(flowSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Sid))
	d.Set("sid", *r.Sid)

	return updateFlowsExecutions(ctx, d, m)
}

func deleteFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.StudioV1.DeleteExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.FetchExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func updateFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateExecutionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.UpdateExecution(flowSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
