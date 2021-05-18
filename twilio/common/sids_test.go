package common

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSidInterfaceTypes(t *testing.T) {
	var _ SidInterface = &Sid{}
	var _ SidInterface = &AccountSid{}
	var _ SidInterface = &StudioStateSid{}
	var _ SidInterface = &FlexFlowSid{}
	var _ SidInterface = &StudioFlowSid{}
	var _ SidInterface = &ServiceInstanceSid{}
	var _ SidInterface = &IdentityRealmSid{}
	var _ SidInterface = &IdentityRealmCertificateSid{}
	var _ SidInterface = &IdentityRealmRoleSid{}
	var _ SidInterface = &ProxyServiceSid{}
	var _ SidInterface = &PhoneNumberSid{}
	var _ SidInterface = &QuotaSid{}
	var _ SidInterface = &ChatRoleSid{}
	var _ SidInterface = &RequestSid{}
	var _ SidInterface = &TaskChannelSid{}
	var _ SidInterface = &UserSid{}
	var _ SidInterface = &ActivitySid{}
	var _ SidInterface = &WorkqueueSid{}
	var _ SidInterface = &WorkspaceSid{}
	var _ SidInterface = &WorkflowSid{}
}

func TestBasicInterfaceTypes(t *testing.T) {
	var _ DecoratedBasicTypeInterface = &Sid{}
	var _ DecoratedBasicTypeInterface = &AccountSid{}
	var _ DecoratedBasicTypeInterface = &StudioStateSid{}
	var _ DecoratedBasicTypeInterface = &FlexFlowSid{}
	var _ DecoratedBasicTypeInterface = &StudioFlowSid{}
	var _ DecoratedBasicTypeInterface = &ServiceInstanceSid{}
	var _ DecoratedBasicTypeInterface = &IdentityRealmSid{}
	var _ DecoratedBasicTypeInterface = &IdentityRealmCertificateSid{}
	var _ DecoratedBasicTypeInterface = &IdentityRealmRoleSid{}
	var _ DecoratedBasicTypeInterface = &ProxyServiceSid{}
	var _ DecoratedBasicTypeInterface = &PhoneNumberSid{}
	var _ DecoratedBasicTypeInterface = &QuotaSid{}
	var _ DecoratedBasicTypeInterface = &RequestSid{}
	var _ DecoratedBasicTypeInterface = &ChatRoleSid{}
	var _ DecoratedBasicTypeInterface = &TaskChannelSid{}
	var _ DecoratedBasicTypeInterface = &UserSid{}
	var _ DecoratedBasicTypeInterface = &ActivitySid{}
	var _ DecoratedBasicTypeInterface = &WorkqueueSid{}
	var _ DecoratedBasicTypeInterface = &WorkspaceSid{}
	var _ DecoratedBasicTypeInterface = &WorkflowSid{}
}

func TestGenericSidSet(t *testing.T) {
	sid := Sid{}
	if err := sid.Set("XY00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
	if string(sid.Prefix[:]) != "XY" {
		t.Errorf("Invalid Sid prefix %s", sid.Prefix)
	}
	if sid.Value != [16]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff} {
		t.Errorf("Invalid Sid value %v", sid.Value)
	}
}

func TestGenericInvalidSidSet1(t *testing.T) {
	sid := Sid{}
	err := sid.Set("X00112233445566778899aabbccddeeff")
	if err == nil {
		t.Errorf("Set did not fail")
	}
	if !strings.Contains(err.Error(), "Invalid length for sid") {
		t.Errorf("Wrong error '%v'", err)
	}
	if sid.Valid {
		t.Errorf("Sid is valid")
	}
	if sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestGenericInvalidSidSet2(t *testing.T) {
	sid := Sid{}
	err := sid.Set("aX00112233445566778899aabbccddeeff")
	if err == nil {
		t.Errorf("Set did not fail")
	}
	if !strings.Contains(err.Error(), "Invalid prefix for sid: aX") {
		t.Errorf("Wrong error '%v'", err)
	}
	if sid.Valid {
		t.Errorf("Sid is valid")
	}
	if sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestGenericInvalidSidSet3(t *testing.T) {
	sid := Sid{}
	err := sid.Set("Xa00112233445566778899aabbccddeeff")
	if err == nil {
		t.Errorf("Set did not fail")
	}
	if !strings.Contains(err.Error(), "Invalid prefix for sid: Xa") {
		t.Errorf("Wrong error '%v'", err)
	}
	if sid.Valid {
		t.Errorf("Sid is valid")
	}
	if sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestGenericInvalidSidSet4(t *testing.T) {
	sid := Sid{}
	err := sid.Set("XYg0112233445566778899aabbccddeeff")
	if err == nil {
		t.Errorf("Set did not fail")
	}
	if !strings.Contains(err.Error(), "encoding/hex: invalid byte: U+0067 'g'") {
		t.Errorf("Wrong error '%v'", err)
	}
	if sid.Valid {
		t.Errorf("Sid is valid")
	}
	if sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestGenericSidString(t *testing.T) {
	sid := Sid{}
	if err := sid.Set("XY00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed '%v'", err)
	}
	str := sid.String()
	if str != "XY00112233445566778899aabbccddeeff" {
		t.Errorf("String() failed: result '%s'", str)
	}
}

func TestGenericJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid Sid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"XY00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
	if !m.Sid.Valid {
		t.Errorf("Sid is not valid")
	}
	if string(m.Sid.Prefix[:]) != "XY" {
		t.Errorf("Invalid Sid prefix %s", m.Sid.Prefix)
	}
	if m.Sid.Value != [16]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff} {
		t.Errorf("Invalid Sid value %v", m.Sid.Value)
	}

}

func TestGenericJsonMarshal(t *testing.T) {
	type MarshalTest struct {
		Sid Sid `json:"sid_value"`
	}

	m := MarshalTest{}
	_ = m.Sid.Set("XY00112233445566778899aabbccddeeff")
	resp, err := json.Marshal(m)
	if err != nil {
		t.Errorf("Marshal failed '%v'", err)
	}
	if string(resp) != `{"sid_value":"XY00112233445566778899aabbccddeeff"}` {
		t.Errorf("Invalid marshaling result of '%s'", string(resp))
	}
}

// ---

func TestAccountSidSet(t *testing.T) {
	sid := AccountSid{}
	if err := sid.Set("AC00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
	if string(sid.Prefix[:]) != "AC" {
		t.Errorf("Invalid Sid prefix %s", sid.Prefix)
	}
	if sid.Value != [16]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff} {
		t.Errorf("Invalid Sid value %v", sid.Value)
	}
}

func TestInvalidAccountSidSet(t *testing.T) {
	sid := AccountSid{}
	err := sid.Set("XX00112233445566778899aabbccddeeff")
	if err == nil {
		t.Errorf("Set did not fail")
	}
	if !strings.Contains(err.Error(), "Invalid prefix for sid: is XX but should be AC") {
		t.Errorf("Wrong error '%v'", err)
	}
	if sid.Valid {
		t.Errorf("Sid is valid")
	}
	if sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestAccountSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid AccountSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"AC00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
	if !m.Sid.Valid {
		t.Errorf("Sid is not valid")
	}
	if string(m.Sid.Prefix[:]) != "AC" {
		t.Errorf("Invalid Sid prefix %s", m.Sid.Prefix)
	}
	if m.Sid.Value != [16]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff} {
		t.Errorf("Invalid Sid value %v", m.Sid.Value)
	}
}

func TestInvalidAccountSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid AccountSid `json:"sid_value"`
	}

	m := MarshalTest{}
	err := json.Unmarshal([]byte(`{"sid_value":"XX00112233445566778899aabbccddeeff"}`), &m)
	if err == nil {
		t.Errorf("Unmarshal did not fail")
	}
	if !strings.Contains(err.Error(), "Invalid prefix for sid: is XX but should be AC") {
		t.Errorf("Wrong error '%v'", err)
	}
	if m.Sid.Valid {
		t.Errorf("Sid is valid")
	}
	if m.Sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestInvalidAccountSidJsonUnmarshal2(t *testing.T) {
	type MarshalTest struct {
		Sid AccountSid `json:"sid_value"`
	}

	m := MarshalTest{}
	err := json.Unmarshal([]byte(`{"sid_value":4}`), &m)
	if err == nil {
		t.Errorf("Unmarshal did not fail")
	}
	if err.Error() != "json: cannot unmarshal number into Go struct field MarshalTest.sid_value of type string" {
		t.Errorf("Wrong error '%v'", err)
	}
	if m.Sid.Valid {
		t.Errorf("Sid is valid")
	}
	if m.Sid.String() != "" {
		t.Errorf("Sid has content")
	}
}

func TestCreateAccountSid(t *testing.T) {
	sid, err := CreateAccountSid("AC00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

// Simplified tests for the rest of the specific Sid types

func TestCreateStudioStateSid(t *testing.T) {
	sid, err := CreateStudioStateSid("FF00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestStudioStateSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid StudioStateSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"FF00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

func TestStudioStateSidSet(t *testing.T) {
	sid := StudioStateSid{}
	if err := sid.Set("FF00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

// ---

func TestCreateFlexFlowSid(t *testing.T) {
	sid, err := CreateFlexFlowSid("FO00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestFlexFlowSidSet(t *testing.T) {
	sid := FlexFlowSid{}
	if err := sid.Set("FO00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestFlexFlowSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid FlexFlowSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"FO00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateStudioFlowSid(t *testing.T) {
	sid, err := CreateStudioFlowSid("FW00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestStudioFlowSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid StudioFlowSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"FW00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

func TestStudioFlowSidSet(t *testing.T) {
	sid := StudioFlowSid{}
	if err := sid.Set("FW00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

// ---

func TestCreateServiceInstanceSid(t *testing.T) {
	sid, err := CreateServiceInstanceSid("IS00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestServiceInstanceSidSet(t *testing.T) {
	sid := ServiceInstanceSid{}
	if err := sid.Set("IS00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestServiceInstanceSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid ServiceInstanceSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"IS00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateIdentityRealmSid(t *testing.T) {
	sid, err := CreateIdentityRealmSid("JB00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestIdentityRealmSidSet(t *testing.T) {
	sid := IdentityRealmSid{}
	if err := sid.Set("JB00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestIdentityRealmSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid IdentityRealmSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"JB00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateIdentityRealmCertificateSid(t *testing.T) {
	sid, err := CreateIdentityRealmCertificateSid("JC00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestIdentityRealmCertificateSidSet(t *testing.T) {
	sid := IdentityRealmCertificateSid{}
	if err := sid.Set("JC00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestIdentityRealmCertificateSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid IdentityRealmCertificateSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"JC00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateIdentityRealmRoleSid(t *testing.T) {
	sid, err := CreateIdentityRealmRoleSid("JD00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestIdentityRealmRoleSidSet(t *testing.T) {
	sid := IdentityRealmRoleSid{}
	if err := sid.Set("JD00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestIdentityRealmRoleSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid IdentityRealmRoleSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"JD00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateProxyServiceSid(t *testing.T) {
	sid, err := CreateProxyServiceSid("KS00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}
func TestProxyServiceSidSet(t *testing.T) {
	sid := ProxyServiceSid{}
	if err := sid.Set("KS00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestProxyServiceSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid ProxyServiceSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"KS00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreatePhoneNumberSid(t *testing.T) {
	sid, err := CreatePhoneNumberSid("PN00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}
func TestPhoneNumberSidSet(t *testing.T) {
	sid := PhoneNumberSid{}
	if err := sid.Set("PN00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestPhoneNumberSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid PhoneNumberSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"PN00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateQuotaSid(t *testing.T) {
	sid, err := CreateQuotaSid("QA00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestQuotaSidSet(t *testing.T) {
	sid := QuotaSid{}
	if err := sid.Set("QA00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestQuotaSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid QuotaSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"QA00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateChatRoleSid(t *testing.T) {
	sid, err := CreateChatRoleSid("RL00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestChatRoleSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid ChatRoleSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"RL00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

func TestChatRoleSidSet(t *testing.T) {
	sid := ChatRoleSid{}
	if err := sid.Set("RL00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

// ---

func TestCreateRequestSid(t *testing.T) {
	sid, err := CreateRequestSid("RQ00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestRequestSidSet(t *testing.T) {
	sid := RequestSid{}
	if err := sid.Set("RQ00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestRequestSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid RequestSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"RQ00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateTaskChannelSid(t *testing.T) {
	sid, err := CreateTaskChannelSid("TC00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestTaskChannelSidSet(t *testing.T) {
	sid := TaskChannelSid{}
	if err := sid.Set("TC00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestTaskChannelSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid TaskChannelSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"TC00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateUserSid(t *testing.T) {
	sid, err := CreateUserSid("US00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}
func TestUserSidSet(t *testing.T) {
	sid := UserSid{}
	if err := sid.Set("US00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestUserSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid UserSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"US00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateActivitySid(t *testing.T) {
	sid, err := CreateActivitySid("WA00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestActivitySidSet(t *testing.T) {
	sid := ActivitySid{}
	if err := sid.Set("WA00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestActivitySidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid ActivitySid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"WA00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateWorkqueueSid(t *testing.T) {
	sid, err := CreateWorkqueueSid("WQ00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestWorkqueueSidSet(t *testing.T) {
	sid := WorkqueueSid{}
	if err := sid.Set("WQ00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestWorkqueueSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid WorkqueueSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"WQ00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateWorkspaceSid(t *testing.T) {
	sid, err := CreateWorkspaceSid("WS00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestWorkspaceSidSet(t *testing.T) {
	sid := WorkspaceSid{}
	if err := sid.Set("WS00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestWorkspaceSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid WorkspaceSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"WS00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCreateWorkflowSid(t *testing.T) {
	sid, err := CreateWorkflowSid("WW00112233445566778899aabbccddeeff")
	if err != nil {
		t.Errorf("Create failed '%v'", err)
	}
	if !sid.Valid {
		t.Errorf("Sid is not valid")
	}
}

func TestWorkflowSidSet(t *testing.T) {
	sid := WorkflowSid{}
	if err := sid.Set("WW00112233445566778899aabbccddeeff"); err != nil {
		t.Errorf("Set failed: result '%v'", err)
	}
}

func TestWorkflowSidJsonUnmarshal(t *testing.T) {
	type MarshalTest struct {
		Sid WorkflowSid `json:"sid_value"`
	}

	m := MarshalTest{}
	if err := json.Unmarshal([]byte(`{"sid_value":"WW00112233445566778899aabbccddeeff"}`), &m); err != nil {
		t.Errorf("Unmarshal failed '%v'", err)
	}
}

// ---

func TestCanCreateRandomSid(t *testing.T) {
	rq := RequestSid{}
	assert.False(t, rq.Valid, "It's just a reference")
	_ = rq.Randomize(&rq)
	assert.True(t, rq.Valid, "It's just a reference")
}

func TestCanNotCreateRandomSidOnValue(t *testing.T) {
	rq := RequestSid{}
	assert.Panics(t, func() {
		// This panics because it expects pointer  rq.Randomize(&rq);
		// and Randomize method cannot declare Randomize(t *interface{})
		// because Cannot use 'rq' (type RequestSid) as type *interface{}
		_ = rq.Randomize(rq)
	})
}
