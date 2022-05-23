/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
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
	. "github.com/twilio/twilio-go/rest/events/v1"
)

func ResourceSinks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSinks,
		ReadContext:   readSinks,
		UpdateContext: updateSinks,
		DeleteContext: deleteSinks,
		Schema: map[string]*schema.Schema{
			"description":        AsString(SchemaRequired),
			"sink_configuration": AsString(SchemaRequired),
			"sink_type":          AsString(SchemaRequired),
			"sid":                AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseSinksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSinkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.EventsV1.CreateSink(&params)
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

func deleteSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.EventsV1.DeleteSink(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.FetchSink(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseSinksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSinkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.UpdateSink(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceSubscriptionsSubscribedEvents() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSubscriptionsSubscribedEvents,
		ReadContext:   readSubscriptionsSubscribedEvents,
		UpdateContext: updateSubscriptionsSubscribedEvents,
		DeleteContext: deleteSubscriptionsSubscribedEvents,
		Schema: map[string]*schema.Schema{
			"subscription_sid": AsString(SchemaRequired),
			"type":             AsString(SchemaRequired),
			"schema_version":   AsInt(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseSubscriptionsSubscribedEventsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSubscribedEventParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	subscriptionSid := d.Get("subscription_sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.CreateSubscribedEvent(subscriptionSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{subscriptionSid}
	idParts = append(idParts, (*r.Type))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	subscriptionSid := d.Get("subscription_sid").(string)
	type_ := d.Get("type").(string)

	err := m.(*client.Config).Client.EventsV1.DeleteSubscribedEvent(subscriptionSid, type_)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	subscriptionSid := d.Get("subscription_sid").(string)
	type_ := d.Get("type").(string)

	r, err := m.(*client.Config).Client.EventsV1.FetchSubscribedEvent(subscriptionSid, type_)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseSubscriptionsSubscribedEventsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected subscription_sid/type"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("subscription_sid", importParts[0])
	d.Set("type", importParts[1])

	return nil
}
func updateSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSubscribedEventParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	subscriptionSid := d.Get("subscription_sid").(string)
	type_ := d.Get("type").(string)

	r, err := m.(*client.Config).Client.EventsV1.UpdateSubscribedEvent(subscriptionSid, type_, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceSubscriptions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSubscriptions,
		ReadContext:   readSubscriptions,
		UpdateContext: updateSubscriptions,
		DeleteContext: deleteSubscriptions,
		Schema: map[string]*schema.Schema{
			"description": AsString(SchemaRequired),
			"sink_sid":    AsString(SchemaRequired),
			"types":       AsList(AsString(SchemaRequired), SchemaRequired),
			"sid":         AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseSubscriptionsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSubscriptionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.EventsV1.CreateSubscription(&params)
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

func deleteSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.EventsV1.DeleteSubscription(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.FetchSubscription(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseSubscriptionsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSubscriptionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.UpdateSubscription(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
