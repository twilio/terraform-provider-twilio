/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
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
	. "github.com/twilio/twilio-go/rest/studio/v2"
)

func ResourceFlowsExecutions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlowsExecutions,
		ReadContext:   readFlowsExecutions,
		UpdateContext: updateFlowsExecutions,
		DeleteContext: deleteFlowsExecutions,
		Schema: map[string]*schema.Schema{
			"flow_sid":   AsString(SchemaRequired),
			"from":       AsString(SchemaComputedOptional),
			"parameters": AsString(SchemaComputedOptional),
			"to":         AsString(SchemaComputedOptional),
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

	r, err := m.(*client.Config).Client.StudioV2.CreateExecution(flowSid, &params)
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

	err := m.(*client.Config).Client.StudioV2.DeleteExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.FetchExecution(flowSid, sid)
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

	r, err := m.(*client.Config).Client.StudioV2.UpdateExecution(flowSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceFlows() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlows,
		ReadContext:   readFlows,
		UpdateContext: updateFlows,
		DeleteContext: deleteFlows,
		Schema: map[string]*schema.Schema{
			"commit_message": AsString(SchemaComputedOptional),
			"definition":     AsString(SchemaComputedOptional),
			"friendly_name":  AsString(SchemaComputedOptional),
			"status":         AsString(SchemaComputedOptional),
			"sid":            AsString(SchemaComputed),
		},
	}
}

func createFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFlowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.StudioV2.CreateFlow(&params)
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

func deleteFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.StudioV2.DeleteFlow(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.FetchFlow(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func updateFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFlowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.UpdateFlow(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
