package twilio

import (
	"errors"
	"fmt"

	"github.com/twilio/terraform-provider-twilio/util"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceFlexFlow() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceFlexFlowCreate,
		Read:   resourceFlexFlowRead,
		Update: resourceFlexFlowUpdate,
		Delete: resourceFlexFlowDelete,
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"chat_service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"channel_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"contact_identity": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"integration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workspace_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"flow_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"channel": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"creation_on_message": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"retry_count": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"long_lived": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"janitor_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFlexFlowCreate(d *schema.ResourceData, m interface{}) error {
	params, err := expandFlexFlowParams(d)
	if err != nil {
		return err
	}
	f, err := m.(*Config).Client.FlexFlow.Create(params)
	if err != nil {
		return fmt.Errorf("could not create FlexFlow: %s", err)
	}
	d.SetId(*f.SID)
	return resourceFlexFlowRead(d, m)
}

func resourceFlexFlowRead(d *schema.ResourceData, m interface{}) error {
	f, err := m.(*Config).Client.FlexFlow.Fetch(d.Id())
	if err != nil {
		return fmt.Errorf("could not read FlexFlow: %s", err)
	}
	i := flattenIntegration(f.Integration)
	d.Set("account_sid", f.AccountSID)
	d.Set("date_created", f.DateCreated)
	d.Set("date_updated", f.DateUpdated)
	d.Set("sid", f.SID)
	d.Set("friendly_name", f.FriendlyName)
	d.Set("chat_service_sid", f.ChatServiceSID)
	d.Set("channel_type", f.ChannelType)
	d.Set("contact_identity", f.ContactIdentity)
	d.Set("enabled", f.Enabled)
	d.Set("integration_type", f.IntegrationType)
	d.Set("integration", i)
	d.Set("long_lived", f.LongLived)
	d.Set("janitor_enabled", f.JanitorEnabled)
	d.Set("url", f.URL)
	return nil
}

func resourceFlexFlowUpdate(d *schema.ResourceData, m interface{}) error {
	params, err := expandFlexFlowParams(d)
	if err != nil {
		return fmt.Errorf("could not expand FlexFlow params: %s", err)
	}

	if _, err = m.(*Config).Client.FlexFlow.Update(d.Id(), params); err != nil {
		return fmt.Errorf("could not update FlexFlow: %s", err)
	}

	return resourceFlexFlowRead(d, m)
}

func resourceFlexFlowDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.FlexFlow.Delete(d.Id()); err != nil {
		return fmt.Errorf("could not delete FlexFlow: %s", err)
	}
	return nil
}

func expandIntegration(base interface{}) (*types.Integration, error) {
	if base == nil {
		return nil, nil
	}
	integrationL := base.([]interface{})

	if len(integrationL) > 1 {
		return nil, errors.New("cannot specify more than new message integration")
	}

	integration := new(types.Integration)

	for _, n := range integrationL {
		i := n.(map[string]interface{})
		integration.WorkspaceSID = util.String(i["workspace_sid"].(string))
		integration.FlowSID = util.String(i["flow_sid"].(string))
		integration.Channel = util.String(i["channel"].(string))
		integration.Timeout = util.Int(i["timeout"].(int))
		integration.Priority = util.Int(i["priority"].(int))
		integration.CreationOnMessage = util.Bool(i["creation_on_message"].(bool))
		integration.RetryCount = util.Int(i["retry_count"].(int))
	}

	return integration, nil
}

func flattenIntegration(i *types.Integration) []interface{} {
	values := map[string]interface{}{}

	if i.WorkspaceSID != nil {
		values["workspace_sid"] = *i.WorkspaceSID
	}

	if i.WorkflowSID != nil {
		values["workflow_sid"] = *i.WorkflowSID
	}

	if i.Channel != nil {
		values["channel"] = *i.Channel
	}

	if i.Timeout != nil {
		values["timeout"] = *i.Timeout
	}

	if i.Priority != nil {
		values["priority"] = *i.Priority
	}

	if i.CreationOnMessage != nil {
		values["creation_on_message"] = *i.CreationOnMessage
	}

	if i.RetryCount != nil {
		values["retry_count"] = *i.RetryCount

	}

	return []interface{}{values}
}

func expandFlexFlowParams(d *schema.ResourceData) (*types.FlexFlowParams, error) {
	i, err := expandIntegration(d.Get("integration"))
	if err != nil {
		return nil, fmt.Errorf("could not expand Integration: %s", err)
	}

	params := &types.FlexFlowParams{
		FriendlyName:    util.String(d.Get("friendly_name").(string)),
		ChatServiceSID:  util.String(d.Get("chat_service_sid").(string)),
		ChannelType:     util.String(d.Get("channel_type").(string)),
		ContactIdentity: util.String(d.Get("contact_identity").(string)),
		Enabled:         util.Bool(d.Get("enabled").(bool)),
		IntegrationType: util.String(d.Get("integration_type").(string)),
		Integration:     i,
		LongLived:       util.Bool(d.Get("long_lived").(bool)),
		JanitorEnabled:  util.Bool(d.Get("janitor_enabled").(bool)),
		URL:             util.String(d.Get("url").(string)),
	}
	return params, nil
}
