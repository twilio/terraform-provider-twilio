package common

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
)

type Sid struct {
	Valid  bool
	Prefix [2]byte
	Value  [16]byte
}

type SidInterface interface {
	Set(value interface{}) error
	Get() (interface{}, bool)
	UnmarshalJSON(buffer []byte) error
	MarshalJSON() ([]byte, error)
}

type AccountSid struct {
	Sid `prefix:"AC"`
}

type StudioStateSid struct {
	Sid `prefix:"FF"`
}

type FlexFlowSid struct {
	Sid `prefix:"FO"`
}

type StudioFlowSid struct {
	Sid `prefix:"FW"`
}

type ServiceInstanceSid struct {
	Sid `prefix:"IS"`
}

type IdentityRealmSid struct {
	Sid `prefix:"JB"`
}

type IdentityRealmCertificateSid struct {
	Sid `prefix:"JC"`
}

type IdentityRealmRoleSid struct {
	Sid `prefix:"JD"`
}

type ProxyServiceSid struct {
	Sid `prefix:"KS"`
}

type PhoneNumberSid struct {
	Sid `prefix:"PN"`
}

type QuotaSid struct {
	Sid `prefix:"QA"`
}

type ChatRoleSid struct {
	Sid `prefix:"RL"`
}

type RequestSid struct {
	Sid `prefix:"RQ"`
}

type TaskChannelSid struct {
	Sid `prefix:"TC"`
}

type UserSid struct {
	Sid `prefix:"US"`
}

type ActivitySid struct {
	Sid `prefix:"WA"`
}

type WorkqueueSid struct {
	Sid `prefix:"WQ"`
}

type WorkspaceSid struct {
	Sid `prefix:"WS"`
}

type WorkflowSid struct {
	Sid `prefix:"WW"`
}

func CreateSid(str string) (Sid, error) {
	ret := Sid{}
	return ret, ret.Set(str)
}

func (sid *Sid) Set(value interface{}) error {
	return sid.setWithType(value, nil)
}

// Stringer
func (sid Sid) String() string {
	if !sid.Valid {
		return ""
	}
	return string(sid.Prefix[:]) + hex.EncodeToString(sid.Value[:])
}

// DecoratedBasicTypeInterface
func (sid Sid) Get() (interface{}, bool) {
	if !sid.Valid {
		return nil, false
	}
	return string(sid.Prefix[:]) + hex.EncodeToString(sid.Value[:]), true
}

func (sid Sid) GetNativePresentation() (interface{}, bool) {
	return sid.Get()
}

// json marshaling
func (sid *Sid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, nil)
}

func (sid Sid) MarshalJSON() ([]byte, error) {
	if str, ok := sid.GetNativePresentation(); ok {
		return json.Marshal(str)
	}
	return []byte("null"), nil
}

// Randomize sid value (useful for creating new)
// PS! t must be an interface
// t cannot be a t *interface{} because you
// Cannot use any other sid (type RequestSid) as type *interface{}
func (sid *Sid) Randomize(t interface{}) error {
	buf := [16]byte{}
	if _, err := rand.Read(buf[:]); err != nil {
		return err
	}
	prefix := reflect.TypeOf(t).Elem().Field(0).Tag.Get("prefix")
	sidVal := fmt.Sprintf("%s%s", prefix[:], hex.EncodeToString(buf[:]))
	return sid.Set(sidVal)
}

// private helpers
func (sid *Sid) setWithType(value interface{}, t interface{}) error {
	sid.Valid = false
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return CreateErrorGeneric("Sid input value is not string")
	}
	// default value for TF is empty
	if len(str) == 0 {
		return nil
	}
	if len(str) != 34 {
		return CreateErrorGeneric("Invalid length for sid")
	}

	if t != nil {
		prefix := reflect.TypeOf(t).Elem().Field(0).Tag.Get("prefix")
		if prefix != str[:2] {
			return CreateErrorGeneric("Invalid prefix for sid: is " + str[:2] + " but should be " + prefix)
		}
	} else {
		isPrefixLetter := func(letter byte) bool {
			return letter >= 'A' && letter <= 'Z'
		}

		if !isPrefixLetter(str[0]) || !isPrefixLetter(str[1]) {
			return CreateErrorGeneric("Invalid prefix for sid: " + str[:2])
		}
	}

	if _, err := hex.Decode(sid.Value[:], []byte(str[2:])); err != nil {
		return err
	}
	copy(sid.Prefix[:2], str[:2])
	sid.Valid = true
	return nil
}

func (sid *Sid) unmarshalJSONWithType(buffer []byte, t interface{}) error {
	var str string
	if err := json.Unmarshal(buffer, &str); err != nil {
		return err
	}
	return sid.setWithType(str, t)
}

// Specific Sid functions

func CreateAccountSid(str string) (AccountSid, error) {
	ret := AccountSid{}
	return ret, ret.Set(str)
}

func (sid *AccountSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *AccountSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateStudioStateSid(str string) (StudioStateSid, error) {
	ret := StudioStateSid{}
	return ret, ret.Set(str)
}

func (sid *StudioStateSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *StudioStateSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateFlexFlowSid(str string) (FlexFlowSid, error) {
	ret := FlexFlowSid{}
	return ret, ret.Set(str)
}

func (sid *FlexFlowSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *FlexFlowSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateStudioFlowSid(str string) (StudioFlowSid, error) {
	ret := StudioFlowSid{}
	return ret, ret.Set(str)
}

func (sid *StudioFlowSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *StudioFlowSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateServiceInstanceSid(str string) (ServiceInstanceSid, error) {
	ret := ServiceInstanceSid{}
	return ret, ret.Set(str)
}

func (sid *ServiceInstanceSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *ServiceInstanceSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateIdentityRealmSid(str string) (IdentityRealmSid, error) {
	ret := IdentityRealmSid{}
	return ret, ret.Set(str)
}

func (sid *IdentityRealmSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *IdentityRealmSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateIdentityRealmCertificateSid(str string) (IdentityRealmCertificateSid, error) {
	ret := IdentityRealmCertificateSid{}
	return ret, ret.Set(str)
}

func (sid *IdentityRealmCertificateSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *IdentityRealmCertificateSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateIdentityRealmRoleSid(str string) (IdentityRealmRoleSid, error) {
	ret := IdentityRealmRoleSid{}
	return ret, ret.Set(str)
}

func (sid *IdentityRealmRoleSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *IdentityRealmRoleSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateProxyServiceSid(str string) (ProxyServiceSid, error) {
	ret := ProxyServiceSid{}
	return ret, ret.Set(str)
}

func (sid *ProxyServiceSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *ProxyServiceSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreatePhoneNumberSid(str string) (PhoneNumberSid, error) {
	ret := PhoneNumberSid{}
	return ret, ret.Set(str)
}

func (sid *PhoneNumberSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *PhoneNumberSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateQuotaSid(str string) (QuotaSid, error) {
	ret := QuotaSid{}
	return ret, ret.Set(str)
}

func (sid *QuotaSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *QuotaSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateChatRoleSid(str string) (ChatRoleSid, error) {
	ret := ChatRoleSid{}
	return ret, ret.Set(str)
}

func (sid *ChatRoleSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *ChatRoleSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateRequestSid(str string) (RequestSid, error) {
	ret := RequestSid{}
	return ret, ret.Set(str)
}

func (sid *RequestSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *RequestSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateTaskChannelSid(str string) (TaskChannelSid, error) {
	ret := TaskChannelSid{}
	return ret, ret.Set(str)
}

func (sid *TaskChannelSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *TaskChannelSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateUserSid(str string) (UserSid, error) {
	ret := UserSid{}
	return ret, ret.Set(str)
}

func (sid *UserSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *UserSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateActivitySid(str string) (ActivitySid, error) {
	ret := ActivitySid{}
	return ret, ret.Set(str)
}

func (sid *ActivitySid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *ActivitySid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateWorkqueueSid(str string) (WorkqueueSid, error) {
	ret := WorkqueueSid{}
	return ret, ret.Set(str)
}

func (sid *WorkqueueSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *WorkqueueSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateWorkspaceSid(str string) (WorkspaceSid, error) {
	ret := WorkspaceSid{}
	return ret, ret.Set(str)
}

func (sid *WorkspaceSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *WorkspaceSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}

func CreateWorkflowSid(str string) (WorkflowSid, error) {
	ret := WorkflowSid{}
	return ret, ret.Set(str)
}

func (sid *WorkflowSid) Set(value interface{}) error {
	return sid.setWithType(value, sid)
}

func (sid *WorkflowSid) UnmarshalJSON(buffer []byte) error {
	return sid.unmarshalJSONWithType(buffer, sid)
}
