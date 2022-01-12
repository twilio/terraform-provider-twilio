/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/voice/v1"
)

func ResourceByocTrunks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createByocTrunks,
		ReadContext:   readByocTrunks,
		UpdateContext: updateByocTrunks,
		DeleteContext: deleteByocTrunks,
		Schema: map[string]*schema.Schema{
			"cnam_lookup_enabled":    AsBool(SchemaComputedOptional),
			"connection_policy_sid":  AsString(SchemaComputedOptional),
			"friendly_name":          AsString(SchemaComputedOptional),
			"from_domain_sid":        AsString(SchemaComputedOptional),
			"status_callback_method": AsString(SchemaComputedOptional),
			"status_callback_url":    AsString(SchemaComputedOptional),
			"voice_fallback_method":  AsString(SchemaComputedOptional),
			"voice_fallback_url":     AsString(SchemaComputedOptional),
			"voice_method":           AsString(SchemaComputedOptional),
			"voice_url":              AsString(SchemaComputedOptional),
			"sid":                    AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseByocTrunksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createByocTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateByocTrunkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VoiceV1.CreateByocTrunk(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteByocTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VoiceV1.DeleteByocTrunk(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readByocTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.FetchByocTrunk(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseByocTrunksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateByocTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateByocTrunkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.UpdateByocTrunk(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceConnectionPolicies() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConnectionPolicies,
		ReadContext:   readConnectionPolicies,
		UpdateContext: updateConnectionPolicies,
		DeleteContext: deleteConnectionPolicies,
		Schema: map[string]*schema.Schema{
			"friendly_name": AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseConnectionPoliciesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createConnectionPolicies(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateConnectionPolicyParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VoiceV1.CreateConnectionPolicy(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteConnectionPolicies(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VoiceV1.DeleteConnectionPolicy(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readConnectionPolicies(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.FetchConnectionPolicy(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseConnectionPoliciesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateConnectionPolicies(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateConnectionPolicyParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.UpdateConnectionPolicy(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceConnectionPoliciesTargets() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConnectionPoliciesTargets,
		ReadContext:   readConnectionPoliciesTargets,
		UpdateContext: updateConnectionPoliciesTargets,
		DeleteContext: deleteConnectionPoliciesTargets,
		Schema: map[string]*schema.Schema{
			"connection_policy_sid": AsString(SchemaRequired),
			"target":                AsString(SchemaRequired),
			"enabled":               AsBool(SchemaComputedOptional),
			"friendly_name":         AsString(SchemaComputedOptional),
			"priority":              AsInt(SchemaComputedOptional),
			"weight":                AsInt(SchemaComputedOptional),
			"sid":                   AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseConnectionPoliciesTargetsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createConnectionPoliciesTargets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateConnectionPolicyTargetParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	connectionPolicySid := d.Get("connection_policy_sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.CreateConnectionPolicyTarget(connectionPolicySid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{connectionPolicySid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteConnectionPoliciesTargets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connectionPolicySid := d.Get("connection_policy_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VoiceV1.DeleteConnectionPolicyTarget(connectionPolicySid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readConnectionPoliciesTargets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connectionPolicySid := d.Get("connection_policy_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.FetchConnectionPolicyTarget(connectionPolicySid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseConnectionPoliciesTargetsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected connection_policy_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("connection_policy_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateConnectionPoliciesTargets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateConnectionPolicyTargetParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	connectionPolicySid := d.Get("connection_policy_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.UpdateConnectionPolicyTarget(connectionPolicySid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceIpRecords() *schema.Resource {
	return &schema.Resource{
		CreateContext: createIpRecords,
		ReadContext:   readIpRecords,
		UpdateContext: updateIpRecords,
		DeleteContext: deleteIpRecords,
		Schema: map[string]*schema.Schema{
			"ip_address":         AsString(SchemaRequired),
			"cidr_prefix_length": AsInt(SchemaComputedOptional),
			"friendly_name":      AsString(SchemaComputedOptional),
			"sid":                AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseIpRecordsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createIpRecords(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateIpRecordParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VoiceV1.CreateIpRecord(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteIpRecords(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VoiceV1.DeleteIpRecord(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readIpRecords(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.FetchIpRecord(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseIpRecordsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateIpRecords(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateIpRecordParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.UpdateIpRecord(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceSourceIpMappings() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSourceIpMappings,
		ReadContext:   readSourceIpMappings,
		UpdateContext: updateSourceIpMappings,
		DeleteContext: deleteSourceIpMappings,
		Schema: map[string]*schema.Schema{
			"ip_record_sid":  AsString(SchemaRequired),
			"sip_domain_sid": AsString(SchemaRequired),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseSourceIpMappingsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createSourceIpMappings(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSourceIpMappingParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VoiceV1.CreateSourceIpMapping(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteSourceIpMappings(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VoiceV1.DeleteSourceIpMapping(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readSourceIpMappings(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.FetchSourceIpMapping(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseSourceIpMappingsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateSourceIpMappings(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSourceIpMappingParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VoiceV1.UpdateSourceIpMapping(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
