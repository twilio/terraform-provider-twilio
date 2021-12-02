/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
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
	. "github.com/twilio/twilio-go/rest/autopilot/v1"
)

func ResourceAssistants() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistants,
		ReadContext:   readAssistants,
		UpdateContext: updateAssistants,
		DeleteContext: deleteAssistants,
		Schema: map[string]*schema.Schema{
			"callback_events":   AsString(SchemaComputedOptional),
			"callback_url":      AsString(SchemaComputedOptional),
			"defaults":          AsString(SchemaComputedOptional),
			"friendly_name":     AsString(SchemaComputedOptional),
			"log_queries":       AsBool(SchemaComputedOptional),
			"style_sheet":       AsString(SchemaComputedOptional),
			"unique_name":       AsString(SchemaComputedOptional),
			"sid":               AsString(SchemaComputed),
			"development_stage": AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateAssistantParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.AutopilotV1.CreateAssistant(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateAssistants(ctx, d, m)
}

func deleteAssistants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteAssistant(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchAssistant(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateAssistants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateAssistantParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateAssistant(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceAssistantsTasksFields() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsTasksFields,
		ReadContext:   readAssistantsTasksFields,
		DeleteContext: deleteAssistantsTasksFields,
		Schema: map[string]*schema.Schema{
			"assistant_sid": AsString(SchemaForceNewRequired),
			"task_sid":      AsString(SchemaForceNewRequired),
			"field_type":    AsString(SchemaForceNewRequired),
			"unique_name":   AsString(SchemaForceNewRequired),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsTasksFieldsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsTasksFields(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFieldParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateField(assistantSid, taskSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid, taskSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsTasksFields(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteField(assistantSid, taskSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsTasksFields(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchField(assistantSid, taskSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsTasksFieldsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/task_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("task_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func ResourceAssistantsFieldTypes() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsFieldTypes,
		ReadContext:   readAssistantsFieldTypes,
		UpdateContext: updateAssistantsFieldTypes,
		DeleteContext: deleteAssistantsFieldTypes,
		Schema: map[string]*schema.Schema{
			"assistant_sid": AsString(SchemaRequired),
			"unique_name":   AsString(SchemaRequired),
			"friendly_name": AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsFieldTypesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsFieldTypes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFieldTypeParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateFieldType(assistantSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsFieldTypes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteFieldType(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsFieldTypes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchFieldType(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsFieldTypesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateAssistantsFieldTypes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFieldTypeParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateFieldType(assistantSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceAssistantsFieldTypesFieldValues() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsFieldTypesFieldValues,
		ReadContext:   readAssistantsFieldTypesFieldValues,
		DeleteContext: deleteAssistantsFieldTypesFieldValues,
		Schema: map[string]*schema.Schema{
			"assistant_sid":  AsString(SchemaForceNewRequired),
			"field_type_sid": AsString(SchemaForceNewRequired),
			"language":       AsString(SchemaForceNewRequired),
			"value":          AsString(SchemaForceNewRequired),
			"synonym_of":     AsString(SchemaForceNewOptional),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsFieldTypesFieldValuesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsFieldTypesFieldValues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFieldValueParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	fieldTypeSid := d.Get("field_type_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateFieldValue(assistantSid, fieldTypeSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid, fieldTypeSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsFieldTypesFieldValues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	fieldTypeSid := d.Get("field_type_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteFieldValue(assistantSid, fieldTypeSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsFieldTypesFieldValues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	fieldTypeSid := d.Get("field_type_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchFieldValue(assistantSid, fieldTypeSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsFieldTypesFieldValuesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/field_type_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("field_type_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func ResourceAssistantsModelBuilds() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsModelBuilds,
		ReadContext:   readAssistantsModelBuilds,
		UpdateContext: updateAssistantsModelBuilds,
		DeleteContext: deleteAssistantsModelBuilds,
		Schema: map[string]*schema.Schema{
			"assistant_sid":   AsString(SchemaRequired),
			"status_callback": AsString(SchemaComputedOptional),
			"unique_name":     AsString(SchemaComputedOptional),
			"sid":             AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsModelBuildsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsModelBuilds(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateModelBuildParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateModelBuild(assistantSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsModelBuilds(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteModelBuild(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsModelBuilds(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchModelBuild(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsModelBuildsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateAssistantsModelBuilds(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateModelBuildParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateModelBuild(assistantSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceAssistantsQueries() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsQueries,
		ReadContext:   readAssistantsQueries,
		UpdateContext: updateAssistantsQueries,
		DeleteContext: deleteAssistantsQueries,
		Schema: map[string]*schema.Schema{
			"assistant_sid": AsString(SchemaRequired),
			"language":      AsString(SchemaRequired),
			"query":         AsString(SchemaRequired),
			"model_build":   AsString(SchemaComputedOptional),
			"tasks":         AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
			"sample_sid":    AsString(SchemaComputedOptional),
			"status":        AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsQueriesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsQueries(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateQueryParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateQuery(assistantSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateAssistantsQueries(ctx, d, m)
}

func deleteAssistantsQueries(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteQuery(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsQueries(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchQuery(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsQueriesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateAssistantsQueries(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateQueryParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateQuery(assistantSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceAssistantsTasksSamples() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsTasksSamples,
		ReadContext:   readAssistantsTasksSamples,
		UpdateContext: updateAssistantsTasksSamples,
		DeleteContext: deleteAssistantsTasksSamples,
		Schema: map[string]*schema.Schema{
			"assistant_sid":  AsString(SchemaRequired),
			"task_sid":       AsString(SchemaRequired),
			"language":       AsString(SchemaRequired),
			"tagged_text":    AsString(SchemaRequired),
			"source_channel": AsString(SchemaComputedOptional),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsTasksSamplesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsTasksSamples(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSampleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateSample(assistantSid, taskSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid, taskSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsTasksSamples(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteSample(assistantSid, taskSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsTasksSamples(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchSample(assistantSid, taskSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsTasksSamplesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/task_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("task_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateAssistantsTasksSamples(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSampleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateSample(assistantSid, taskSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceAssistantsTasks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsTasks,
		ReadContext:   readAssistantsTasks,
		UpdateContext: updateAssistantsTasks,
		DeleteContext: deleteAssistantsTasks,
		Schema: map[string]*schema.Schema{
			"assistant_sid": AsString(SchemaRequired),
			"unique_name":   AsString(SchemaRequired),
			"actions":       AsString(SchemaComputedOptional),
			"actions_url":   AsString(SchemaComputedOptional),
			"friendly_name": AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsTasksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTaskParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateTask(assistantSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteTask(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchTask(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsTasksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateAssistantsTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateTaskParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateTask(assistantSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceAssistantsWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAssistantsWebhooks,
		ReadContext:   readAssistantsWebhooks,
		UpdateContext: updateAssistantsWebhooks,
		DeleteContext: deleteAssistantsWebhooks,
		Schema: map[string]*schema.Schema{
			"assistant_sid":  AsString(SchemaRequired),
			"events":         AsString(SchemaRequired),
			"unique_name":    AsString(SchemaRequired),
			"webhook_url":    AsString(SchemaRequired),
			"webhook_method": AsString(SchemaComputedOptional),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseAssistantsWebhooksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createAssistantsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.CreateWebhook(assistantSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{assistantSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteAssistantsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AutopilotV1.DeleteWebhook(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readAssistantsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.FetchWebhook(assistantSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseAssistantsWebhooksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected assistant_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("assistant_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateAssistantsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AutopilotV1.UpdateWebhook(assistantSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
