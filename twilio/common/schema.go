package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	SchemaRequired = &options{
		Required: true,
	}

	SchemaOptional = &options{
		Optional: true,
	}

	SchemaComputed = &options{
		Computed: true,
	}

	SchemaForceNewRequired = &options{
		Required: true,
		ForceNew: true,
	}

	SchemaForceNewOptional = &options{
		Optional: true,
		ForceNew: true,
	}
)

type options struct {
	Optional bool
	Required bool
	Computed bool
	ForceNew bool
}

type SchemaPlus struct {
	*schema.Schema
}

func AsList(obj interface{}, conf *options) *schema.Schema {
	var ret *schema.Schema
	if item, ok := obj.(*schema.Schema); ok {
		// scalar type
		// validation is not supported for list items
		ret = &schema.Schema{
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: item.Type,
			},
		}
	}
	if item, ok := obj.(map[string]*schema.Schema); ok {
		ret = &schema.Schema{
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: item,
			},
		}
	}
	return addSchemaOptions(ret, conf)
}

func AsString(conf *options) *schema.Schema {
	ret := &schema.Schema{
		Type: schema.TypeString,
	}
	return addSchemaOptions(ret, conf)
}

func AsMap(conf *options) *schema.Schema {
	ret := &schema.Schema{
		Type: schema.TypeMap,
	}
	return addSchemaOptions(ret, conf)
}

func AsBool(conf *options) *schema.Schema {
	ret := &schema.Schema{
		Type: schema.TypeBool,
	}
	return addSchemaOptions(ret, conf)

}

func AsInt(conf *options) *schema.Schema {
	ret := &schema.Schema{
		Type: schema.TypeInt,
	}
	return addSchemaOptions(ret, conf)
}

func AsFloat(conf *options) *schema.Schema {
	ret := &schema.Schema{
		Type: schema.TypeFloat,
	}
	return addSchemaOptions(ret, conf)
}

func AsSid(sid SidInterface, conf *options) *schema.Schema {
	ret := &schema.Schema{
		Type: schema.TypeString,
	}
	if !conf.Computed {
		ret.ValidateFunc = func(val interface{}, key string) (warns []string, errs []error) {
			if err := sid.Set(val); err != nil {
				return []string{}, []error{err}
			}
			return []string{}, []error{}
		}
	}
	return addSchemaOptions(ret, conf)
}

// helper method for schema definition, for each schema item, one of optional, required, or computed must be set
func addSchemaOptions(s *schema.Schema, conf *options) *schema.Schema {
	if conf.Computed {
		if conf.Required || conf.Optional {
			panic(fmt.Errorf("Computed parameter can't be Required/Optional"))
		}
		if conf.ForceNew {
			panic(fmt.Errorf("Computed parameter can't be ForceNew"))
		}
	} else {
		if conf.Optional == conf.Required {
			panic(fmt.Errorf("Computed and Required are mutually exclusive"))
		}
	}
	s.Optional = conf.Optional
	s.Required = conf.Required
	s.Computed = conf.Computed
	s.ForceNew = conf.ForceNew
	return s
}