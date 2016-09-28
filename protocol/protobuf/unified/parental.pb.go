// Code generated by protoc-gen-go.
// source: steammessages_parental.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package unified is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package unified to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ParentalApp struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	IsAllowed        *bool   `protobuf:"varint,2,opt,name=is_allowed,json=isAllowed" json:"is_allowed,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ParentalApp) Reset()                    { *m = ParentalApp{} }
func (m *ParentalApp) String() string            { return proto.CompactTextString(m) }
func (*ParentalApp) ProtoMessage()               {}
func (*ParentalApp) Descriptor() ([]byte, []int) { return parental_fileDescriptor0, []int{0} }

func (m *ParentalApp) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *ParentalApp) GetIsAllowed() bool {
	if m != nil && m.IsAllowed != nil {
		return *m.IsAllowed
	}
	return false
}

type ParentalSettings struct {
	Steamid                *uint64        `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	ApplistBaseId          *uint32        `protobuf:"varint,2,opt,name=applist_base_id,json=applistBaseId" json:"applist_base_id,omitempty"`
	ApplistBaseDescription *string        `protobuf:"bytes,3,opt,name=applist_base_description,json=applistBaseDescription" json:"applist_base_description,omitempty"`
	ApplistBase            []*ParentalApp `protobuf:"bytes,4,rep,name=applist_base,json=applistBase" json:"applist_base,omitempty"`
	ApplistCustom          []*ParentalApp `protobuf:"bytes,5,rep,name=applist_custom,json=applistCustom" json:"applist_custom,omitempty"`
	Passwordhashtype       *uint32        `protobuf:"varint,6,opt,name=passwordhashtype" json:"passwordhashtype,omitempty"`
	Salt                   []byte         `protobuf:"bytes,7,opt,name=salt" json:"salt,omitempty"`
	Passwordhash           []byte         `protobuf:"bytes,8,opt,name=passwordhash" json:"passwordhash,omitempty"`
	IsEnabled              *bool          `protobuf:"varint,9,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	EnabledFeatures        *uint32        `protobuf:"varint,10,opt,name=enabled_features,json=enabledFeatures" json:"enabled_features,omitempty"`
	RecoveryEmail          *string        `protobuf:"bytes,11,opt,name=recovery_email,json=recoveryEmail" json:"recovery_email,omitempty"`
	XXX_unrecognized       []byte         `json:"-"`
}

func (m *ParentalSettings) Reset()                    { *m = ParentalSettings{} }
func (m *ParentalSettings) String() string            { return proto.CompactTextString(m) }
func (*ParentalSettings) ProtoMessage()               {}
func (*ParentalSettings) Descriptor() ([]byte, []int) { return parental_fileDescriptor0, []int{1} }

func (m *ParentalSettings) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *ParentalSettings) GetApplistBaseId() uint32 {
	if m != nil && m.ApplistBaseId != nil {
		return *m.ApplistBaseId
	}
	return 0
}

func (m *ParentalSettings) GetApplistBaseDescription() string {
	if m != nil && m.ApplistBaseDescription != nil {
		return *m.ApplistBaseDescription
	}
	return ""
}

func (m *ParentalSettings) GetApplistBase() []*ParentalApp {
	if m != nil {
		return m.ApplistBase
	}
	return nil
}

func (m *ParentalSettings) GetApplistCustom() []*ParentalApp {
	if m != nil {
		return m.ApplistCustom
	}
	return nil
}

func (m *ParentalSettings) GetPasswordhashtype() uint32 {
	if m != nil && m.Passwordhashtype != nil {
		return *m.Passwordhashtype
	}
	return 0
}

func (m *ParentalSettings) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *ParentalSettings) GetPasswordhash() []byte {
	if m != nil {
		return m.Passwordhash
	}
	return nil
}

func (m *ParentalSettings) GetIsEnabled() bool {
	if m != nil && m.IsEnabled != nil {
		return *m.IsEnabled
	}
	return false
}

func (m *ParentalSettings) GetEnabledFeatures() uint32 {
	if m != nil && m.EnabledFeatures != nil {
		return *m.EnabledFeatures
	}
	return 0
}

func (m *ParentalSettings) GetRecoveryEmail() string {
	if m != nil && m.RecoveryEmail != nil {
		return *m.RecoveryEmail
	}
	return ""
}

type CParental_EnableParentalSettings_Request struct {
	Password         *string           `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Settings         *ParentalSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
	Sessionid        *string           `protobuf:"bytes,3,opt,name=sessionid" json:"sessionid,omitempty"`
	Enablecode       *uint32           `protobuf:"varint,4,opt,name=enablecode" json:"enablecode,omitempty"`
	Steamid          *uint64           `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CParental_EnableParentalSettings_Request) Reset() {
	*m = CParental_EnableParentalSettings_Request{}
}
func (m *CParental_EnableParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_EnableParentalSettings_Request) ProtoMessage()    {}
func (*CParental_EnableParentalSettings_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{2}
}

func (m *CParental_EnableParentalSettings_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_EnableParentalSettings_Request) GetSettings() *ParentalSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *CParental_EnableParentalSettings_Request) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

func (m *CParental_EnableParentalSettings_Request) GetEnablecode() uint32 {
	if m != nil && m.Enablecode != nil {
		return *m.Enablecode
	}
	return 0
}

func (m *CParental_EnableParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_EnableParentalSettings_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_EnableParentalSettings_Response) Reset() {
	*m = CParental_EnableParentalSettings_Response{}
}
func (m *CParental_EnableParentalSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_EnableParentalSettings_Response) ProtoMessage()    {}
func (*CParental_EnableParentalSettings_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{3}
}

type CParental_DisableParentalSettings_Request struct {
	Password         *string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Steamid          *uint64 `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_DisableParentalSettings_Request) Reset() {
	*m = CParental_DisableParentalSettings_Request{}
}
func (m *CParental_DisableParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_DisableParentalSettings_Request) ProtoMessage()    {}
func (*CParental_DisableParentalSettings_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{4}
}

func (m *CParental_DisableParentalSettings_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_DisableParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_DisableParentalSettings_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_DisableParentalSettings_Response) Reset() {
	*m = CParental_DisableParentalSettings_Response{}
}
func (m *CParental_DisableParentalSettings_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_DisableParentalSettings_Response) ProtoMessage() {}
func (*CParental_DisableParentalSettings_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{5}
}

type CParental_GetParentalSettings_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_GetParentalSettings_Request) Reset()         { *m = CParental_GetParentalSettings_Request{} }
func (m *CParental_GetParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_GetParentalSettings_Request) ProtoMessage()    {}
func (*CParental_GetParentalSettings_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{6}
}

func (m *CParental_GetParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_GetParentalSettings_Response struct {
	Settings         *ParentalSettings `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CParental_GetParentalSettings_Response) Reset() {
	*m = CParental_GetParentalSettings_Response{}
}
func (m *CParental_GetParentalSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_GetParentalSettings_Response) ProtoMessage()    {}
func (*CParental_GetParentalSettings_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{7}
}

func (m *CParental_GetParentalSettings_Response) GetSettings() *ParentalSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type CParental_GetSignedParentalSettings_Request struct {
	Priority         *uint32 `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_GetSignedParentalSettings_Request) Reset() {
	*m = CParental_GetSignedParentalSettings_Request{}
}
func (m *CParental_GetSignedParentalSettings_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_GetSignedParentalSettings_Request) ProtoMessage() {}
func (*CParental_GetSignedParentalSettings_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{8}
}

func (m *CParental_GetSignedParentalSettings_Request) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

type CParental_GetSignedParentalSettings_Response struct {
	SerializedSettings []byte `protobuf:"bytes,1,opt,name=serialized_settings,json=serializedSettings" json:"serialized_settings,omitempty"`
	Signature          []byte `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *CParental_GetSignedParentalSettings_Response) Reset() {
	*m = CParental_GetSignedParentalSettings_Response{}
}
func (m *CParental_GetSignedParentalSettings_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_GetSignedParentalSettings_Response) ProtoMessage() {}
func (*CParental_GetSignedParentalSettings_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{9}
}

func (m *CParental_GetSignedParentalSettings_Response) GetSerializedSettings() []byte {
	if m != nil {
		return m.SerializedSettings
	}
	return nil
}

func (m *CParental_GetSignedParentalSettings_Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CParental_SetParentalSettings_Request struct {
	Password         *string           `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Settings         *ParentalSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
	NewPassword      *string           `protobuf:"bytes,3,opt,name=new_password,json=newPassword" json:"new_password,omitempty"`
	Sessionid        *string           `protobuf:"bytes,4,opt,name=sessionid" json:"sessionid,omitempty"`
	Steamid          *uint64           `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CParental_SetParentalSettings_Request) Reset()         { *m = CParental_SetParentalSettings_Request{} }
func (m *CParental_SetParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_SetParentalSettings_Request) ProtoMessage()    {}
func (*CParental_SetParentalSettings_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{10}
}

func (m *CParental_SetParentalSettings_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_SetParentalSettings_Request) GetSettings() *ParentalSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *CParental_SetParentalSettings_Request) GetNewPassword() string {
	if m != nil && m.NewPassword != nil {
		return *m.NewPassword
	}
	return ""
}

func (m *CParental_SetParentalSettings_Request) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

func (m *CParental_SetParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_SetParentalSettings_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_SetParentalSettings_Response) Reset() {
	*m = CParental_SetParentalSettings_Response{}
}
func (m *CParental_SetParentalSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_SetParentalSettings_Response) ProtoMessage()    {}
func (*CParental_SetParentalSettings_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{11}
}

type CParental_ValidateToken_Request struct {
	UnlockToken      *string `protobuf:"bytes,1,opt,name=unlock_token,json=unlockToken" json:"unlock_token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ValidateToken_Request) Reset()         { *m = CParental_ValidateToken_Request{} }
func (m *CParental_ValidateToken_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidateToken_Request) ProtoMessage()    {}
func (*CParental_ValidateToken_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{12}
}

func (m *CParental_ValidateToken_Request) GetUnlockToken() string {
	if m != nil && m.UnlockToken != nil {
		return *m.UnlockToken
	}
	return ""
}

type CParental_ValidateToken_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_ValidateToken_Response) Reset()         { *m = CParental_ValidateToken_Response{} }
func (m *CParental_ValidateToken_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidateToken_Response) ProtoMessage()    {}
func (*CParental_ValidateToken_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{13}
}

type CParental_ValidatePassword_Request struct {
	Password            *string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Session             *string `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	SendUnlockOnSuccess *bool   `protobuf:"varint,3,opt,name=send_unlock_on_success,json=sendUnlockOnSuccess" json:"send_unlock_on_success,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *CParental_ValidatePassword_Request) Reset()         { *m = CParental_ValidatePassword_Request{} }
func (m *CParental_ValidatePassword_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidatePassword_Request) ProtoMessage()    {}
func (*CParental_ValidatePassword_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{14}
}

func (m *CParental_ValidatePassword_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_ValidatePassword_Request) GetSession() string {
	if m != nil && m.Session != nil {
		return *m.Session
	}
	return ""
}

func (m *CParental_ValidatePassword_Request) GetSendUnlockOnSuccess() bool {
	if m != nil && m.SendUnlockOnSuccess != nil {
		return *m.SendUnlockOnSuccess
	}
	return false
}

type CParental_ValidatePassword_Response struct {
	Token            *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ValidatePassword_Response) Reset()         { *m = CParental_ValidatePassword_Response{} }
func (m *CParental_ValidatePassword_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidatePassword_Response) ProtoMessage()    {}
func (*CParental_ValidatePassword_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{15}
}

func (m *CParental_ValidatePassword_Response) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type CParental_LockClient_Request struct {
	Session          *string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_LockClient_Request) Reset()                    { *m = CParental_LockClient_Request{} }
func (m *CParental_LockClient_Request) String() string            { return proto.CompactTextString(m) }
func (*CParental_LockClient_Request) ProtoMessage()               {}
func (*CParental_LockClient_Request) Descriptor() ([]byte, []int) { return parental_fileDescriptor0, []int{16} }

func (m *CParental_LockClient_Request) GetSession() string {
	if m != nil && m.Session != nil {
		return *m.Session
	}
	return ""
}

type CParental_LockClient_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_LockClient_Response) Reset()                    { *m = CParental_LockClient_Response{} }
func (m *CParental_LockClient_Response) String() string            { return proto.CompactTextString(m) }
func (*CParental_LockClient_Response) ProtoMessage()               {}
func (*CParental_LockClient_Response) Descriptor() ([]byte, []int) { return parental_fileDescriptor0, []int{17} }

type CParental_RequestRecoveryCode_Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_RequestRecoveryCode_Request) Reset()         { *m = CParental_RequestRecoveryCode_Request{} }
func (m *CParental_RequestRecoveryCode_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_RequestRecoveryCode_Request) ProtoMessage()    {}
func (*CParental_RequestRecoveryCode_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{18}
}

type CParental_RequestRecoveryCode_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_RequestRecoveryCode_Response) Reset() {
	*m = CParental_RequestRecoveryCode_Response{}
}
func (m *CParental_RequestRecoveryCode_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_RequestRecoveryCode_Response) ProtoMessage()    {}
func (*CParental_RequestRecoveryCode_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{19}
}

type CParental_DisableWithRecoveryCode_Request struct {
	RecoveryCode     *uint32 `protobuf:"varint,1,opt,name=recovery_code,json=recoveryCode" json:"recovery_code,omitempty"`
	Steamid          *uint64 `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_DisableWithRecoveryCode_Request) Reset() {
	*m = CParental_DisableWithRecoveryCode_Request{}
}
func (m *CParental_DisableWithRecoveryCode_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_DisableWithRecoveryCode_Request) ProtoMessage()    {}
func (*CParental_DisableWithRecoveryCode_Request) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{20}
}

func (m *CParental_DisableWithRecoveryCode_Request) GetRecoveryCode() uint32 {
	if m != nil && m.RecoveryCode != nil {
		return *m.RecoveryCode
	}
	return 0
}

func (m *CParental_DisableWithRecoveryCode_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_DisableWithRecoveryCode_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_DisableWithRecoveryCode_Response) Reset() {
	*m = CParental_DisableWithRecoveryCode_Response{}
}
func (m *CParental_DisableWithRecoveryCode_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_DisableWithRecoveryCode_Response) ProtoMessage() {}
func (*CParental_DisableWithRecoveryCode_Response) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{21}
}

type CParental_ParentalSettingsChange_Notification struct {
	SerializedSettings []byte  `protobuf:"bytes,1,opt,name=serialized_settings,json=serializedSettings" json:"serialized_settings,omitempty"`
	Signature          []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Password           *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Sessionid          *string `protobuf:"bytes,4,opt,name=sessionid" json:"sessionid,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *CParental_ParentalSettingsChange_Notification) Reset() {
	*m = CParental_ParentalSettingsChange_Notification{}
}
func (m *CParental_ParentalSettingsChange_Notification) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_ParentalSettingsChange_Notification) ProtoMessage() {}
func (*CParental_ParentalSettingsChange_Notification) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{22}
}

func (m *CParental_ParentalSettingsChange_Notification) GetSerializedSettings() []byte {
	if m != nil {
		return m.SerializedSettings
	}
	return nil
}

func (m *CParental_ParentalSettingsChange_Notification) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CParental_ParentalSettingsChange_Notification) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_ParentalSettingsChange_Notification) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

type CParental_ParentalUnlock_Notification struct {
	Password         *string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Sessionid        *string `protobuf:"bytes,2,opt,name=sessionid" json:"sessionid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ParentalUnlock_Notification) Reset()         { *m = CParental_ParentalUnlock_Notification{} }
func (m *CParental_ParentalUnlock_Notification) String() string { return proto.CompactTextString(m) }
func (*CParental_ParentalUnlock_Notification) ProtoMessage()    {}
func (*CParental_ParentalUnlock_Notification) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{23}
}

func (m *CParental_ParentalUnlock_Notification) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_ParentalUnlock_Notification) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

type CParental_ParentalLock_Notification struct {
	Sessionid        *string `protobuf:"bytes,1,opt,name=sessionid" json:"sessionid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ParentalLock_Notification) Reset()         { *m = CParental_ParentalLock_Notification{} }
func (m *CParental_ParentalLock_Notification) String() string { return proto.CompactTextString(m) }
func (*CParental_ParentalLock_Notification) ProtoMessage()    {}
func (*CParental_ParentalLock_Notification) Descriptor() ([]byte, []int) {
	return parental_fileDescriptor0, []int{24}
}

func (m *CParental_ParentalLock_Notification) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

func init() {
	proto.RegisterType((*ParentalApp)(nil), "ParentalApp")
	proto.RegisterType((*ParentalSettings)(nil), "ParentalSettings")
	proto.RegisterType((*CParental_EnableParentalSettings_Request)(nil), "CParental_EnableParentalSettings_Request")
	proto.RegisterType((*CParental_EnableParentalSettings_Response)(nil), "CParental_EnableParentalSettings_Response")
	proto.RegisterType((*CParental_DisableParentalSettings_Request)(nil), "CParental_DisableParentalSettings_Request")
	proto.RegisterType((*CParental_DisableParentalSettings_Response)(nil), "CParental_DisableParentalSettings_Response")
	proto.RegisterType((*CParental_GetParentalSettings_Request)(nil), "CParental_GetParentalSettings_Request")
	proto.RegisterType((*CParental_GetParentalSettings_Response)(nil), "CParental_GetParentalSettings_Response")
	proto.RegisterType((*CParental_GetSignedParentalSettings_Request)(nil), "CParental_GetSignedParentalSettings_Request")
	proto.RegisterType((*CParental_GetSignedParentalSettings_Response)(nil), "CParental_GetSignedParentalSettings_Response")
	proto.RegisterType((*CParental_SetParentalSettings_Request)(nil), "CParental_SetParentalSettings_Request")
	proto.RegisterType((*CParental_SetParentalSettings_Response)(nil), "CParental_SetParentalSettings_Response")
	proto.RegisterType((*CParental_ValidateToken_Request)(nil), "CParental_ValidateToken_Request")
	proto.RegisterType((*CParental_ValidateToken_Response)(nil), "CParental_ValidateToken_Response")
	proto.RegisterType((*CParental_ValidatePassword_Request)(nil), "CParental_ValidatePassword_Request")
	proto.RegisterType((*CParental_ValidatePassword_Response)(nil), "CParental_ValidatePassword_Response")
	proto.RegisterType((*CParental_LockClient_Request)(nil), "CParental_LockClient_Request")
	proto.RegisterType((*CParental_LockClient_Response)(nil), "CParental_LockClient_Response")
	proto.RegisterType((*CParental_RequestRecoveryCode_Request)(nil), "CParental_RequestRecoveryCode_Request")
	proto.RegisterType((*CParental_RequestRecoveryCode_Response)(nil), "CParental_RequestRecoveryCode_Response")
	proto.RegisterType((*CParental_DisableWithRecoveryCode_Request)(nil), "CParental_DisableWithRecoveryCode_Request")
	proto.RegisterType((*CParental_DisableWithRecoveryCode_Response)(nil), "CParental_DisableWithRecoveryCode_Response")
	proto.RegisterType((*CParental_ParentalSettingsChange_Notification)(nil), "CParental_ParentalSettingsChange_Notification")
	proto.RegisterType((*CParental_ParentalUnlock_Notification)(nil), "CParental_ParentalUnlock_Notification")
	proto.RegisterType((*CParental_ParentalLock_Notification)(nil), "CParental_ParentalLock_Notification")
}

func init() { proto.RegisterFile("steammessages_parental.steamclient.proto", parental_fileDescriptor0) }

var parental_fileDescriptor0 = []byte{
	// 1458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x57, 0xcf, 0x6f, 0xd4, 0xc6,
	0x17, 0x97, 0x49, 0x80, 0x64, 0x76, 0x13, 0xf2, 0x1d, 0x10, 0x5f, 0xb3, 0x22, 0x30, 0x18, 0x08,
	0x01, 0x82, 0x91, 0x82, 0xd4, 0x22, 0x21, 0xb5, 0x4a, 0x16, 0x4a, 0x51, 0x29, 0x50, 0x2f, 0x2d,
	0x95, 0x38, 0x58, 0x8e, 0x3d, 0xd9, 0x98, 0xec, 0xce, 0x6c, 0x3d, 0xb3, 0xa4, 0xa9, 0x54, 0xa9,
	0xea, 0xa1, 0xb7, 0xaa, 0x97, 0x1e, 0x7b, 0xaf, 0xd4, 0x9e, 0x51, 0x7b, 0xac, 0xd4, 0x53, 0x0f,
	0x95, 0xaa, 0xfe, 0x15, 0xfd, 0x33, 0xfa, 0x3c, 0x33, 0xde, 0xb5, 0xf7, 0xa7, 0xb7, 0x70, 0xf3,
	0xbc, 0xf7, 0xe6, 0xcd, 0xe7, 0xbd, 0xf7, 0x99, 0x79, 0xcf, 0x68, 0x5d, 0x48, 0x1a, 0xb4, 0xdb,
	0x54, 0x88, 0xa0, 0x49, 0x85, 0xdf, 0x09, 0x12, 0xca, 0x64, 0xd0, 0x72, 0x95, 0x38, 0x6c, 0xc5,
	0xb0, 0x72, 0x3b, 0x09, 0x97, 0xbc, 0xb6, 0x51, 0xb4, 0xec, 0xb2, 0x78, 0x37, 0xa6, 0x91, 0xbf,
	0x13, 0x08, 0x3a, 0x6c, 0xed, 0x6c, 0xa3, 0xca, 0x13, 0xe3, 0x6b, 0xab, 0xd3, 0xc1, 0xa7, 0xd0,
	0xd1, 0xa0, 0xd3, 0x89, 0x23, 0xdb, 0x22, 0xd6, 0xfa, 0x92, 0xa7, 0x17, 0x78, 0x15, 0xa1, 0x58,
	0xf8, 0x41, 0xab, 0xc5, 0x0f, 0x68, 0x64, 0x1f, 0x01, 0xd5, 0x82, 0xb7, 0x18, 0x8b, 0x2d, 0x2d,
	0x70, 0xfe, 0x98, 0x43, 0x2b, 0x99, 0x93, 0x06, 0x95, 0x32, 0x66, 0x4d, 0x81, 0x6d, 0x74, 0x5c,
	0x9d, 0x66, 0x7c, 0x1d, 0xf3, 0xb2, 0x25, 0x5e, 0x43, 0x27, 0xc0, 0x6d, 0x2b, 0x16, 0x52, 0x81,
	0xf2, 0x63, 0xed, 0x72, 0xc9, 0x5b, 0x32, 0xe2, 0x6d, 0x90, 0x3e, 0x88, 0xf0, 0x6d, 0x64, 0x17,
	0xec, 0x22, 0x2a, 0xc2, 0x24, 0xee, 0xc8, 0x98, 0x33, 0x7b, 0x0e, 0x36, 0x2c, 0x7a, 0xa7, 0x73,
	0x1b, 0xee, 0xf6, 0xb5, 0xf8, 0x26, 0xaa, 0xe6, 0x77, 0xda, 0xf3, 0x64, 0x6e, 0xbd, 0xb2, 0x59,
	0x75, 0x73, 0x91, 0x7a, 0x95, 0xdc, 0x5e, 0x7c, 0x0b, 0x2d, 0x67, 0x1b, 0xc2, 0xae, 0x90, 0xbc,
	0x6d, 0x1f, 0x1d, 0xb1, 0x25, 0xc3, 0x57, 0x57, 0x26, 0xf8, 0x1a, 0x5a, 0xe9, 0x04, 0x42, 0x1c,
	0xf0, 0x24, 0xda, 0x0b, 0xc4, 0x9e, 0x3c, 0xec, 0x50, 0xfb, 0x98, 0x0a, 0x64, 0x48, 0x8e, 0x31,
	0x9a, 0x17, 0x41, 0x4b, 0xda, 0xc7, 0x41, 0x5f, 0xf5, 0xd4, 0x37, 0x76, 0x50, 0x35, 0x6f, 0x67,
	0x2f, 0x28, 0x5d, 0x41, 0x66, 0x32, 0x4f, 0x59, 0xb0, 0xd3, 0x82, 0xcc, 0x2f, 0x66, 0x99, 0xbf,
	0xa7, 0x05, 0xf8, 0x2a, 0x5a, 0x31, 0x3a, 0x7f, 0x97, 0x06, 0xb2, 0x9b, 0x50, 0x61, 0x23, 0x05,
	0xe1, 0x84, 0x91, 0xbf, 0x67, 0xc4, 0xf8, 0x32, 0x5a, 0x4e, 0x68, 0xc8, 0x5f, 0xd2, 0xe4, 0xd0,
	0xa7, 0xed, 0x20, 0x6e, 0xd9, 0x15, 0x95, 0xc3, 0xa5, 0x4c, 0x7a, 0x2f, 0x15, 0x3a, 0x7f, 0x5a,
	0x68, 0xbd, 0x9e, 0x05, 0xed, 0xeb, 0x73, 0x06, 0x6b, 0xeb, 0x7b, 0xf4, 0xb3, 0x2e, 0x15, 0x12,
	0xd7, 0xd0, 0x42, 0x86, 0x56, 0x15, 0x79, 0xd1, 0xeb, 0xad, 0xf1, 0x0d, 0xb4, 0x20, 0x8c, 0xbd,
	0x2a, 0x6f, 0x65, 0xf3, 0x7f, 0xee, 0xa0, 0x23, 0xaf, 0x67, 0x82, 0xcf, 0xa2, 0x45, 0x01, 0x94,
	0x85, 0xea, 0x01, 0x1d, 0x74, 0x75, 0xfb, 0x02, 0x7c, 0x0e, 0x21, 0x1d, 0x4f, 0xc8, 0xa3, 0xb4,
	0x9c, 0x69, 0x84, 0x39, 0x49, 0x9e, 0x6c, 0xa8, 0x40, 0x36, 0xe7, 0x3a, 0xba, 0x5a, 0x22, 0x1c,
	0xd1, 0xe1, 0x4c, 0x50, 0x27, 0xc8, 0x1b, 0xdf, 0x8d, 0xc5, 0x7f, 0x0e, 0x7e, 0x3c, 0x9e, 0x0d,
	0x74, 0xad, 0xcc, 0x11, 0x06, 0xd0, 0x16, 0xba, 0xdc, 0xb7, 0xbe, 0x4f, 0xe5, 0x58, 0x30, 0xe3,
	0x0f, 0x7c, 0x86, 0xd6, 0xa6, 0xb9, 0xd0, 0x87, 0x15, 0x2a, 0x66, 0x4d, 0xad, 0x98, 0xf3, 0x00,
	0x5d, 0x2f, 0x38, 0x6e, 0xc4, 0x4d, 0x46, 0xa3, 0x89, 0xe9, 0x4a, 0x62, 0x9e, 0xc4, 0xf2, 0xd0,
	0x3c, 0x2e, 0xbd, 0xb5, 0xf3, 0x25, 0xda, 0x28, 0xe7, 0xca, 0x20, 0xbd, 0x89, 0x4e, 0x0a, 0x9a,
	0xc4, 0x41, 0x2b, 0xfe, 0x02, 0x98, 0x5f, 0x00, 0x5d, 0xf5, 0x70, 0x5f, 0xd5, 0xc8, 0xb3, 0x0b,
	0x7c, 0xaa, 0xab, 0xa0, 0xd8, 0x58, 0xf5, 0xfa, 0x82, 0x94, 0xf3, 0xb9, 0x34, 0x37, 0x26, 0xa4,
	0xf9, 0x0d, 0x12, 0xfe, 0x02, 0xaa, 0x32, 0x7a, 0xe0, 0xf7, 0xdc, 0x69, 0xce, 0x57, 0x40, 0xf6,
	0x24, 0xf3, 0x58, 0xb8, 0x13, 0xf3, 0x83, 0x77, 0x62, 0x7c, 0xc9, 0xd7, 0xf3, 0x25, 0x6f, 0x4c,
	0x28, 0xb9, 0x73, 0x17, 0x9d, 0xef, 0x5b, 0x7e, 0x02, 0x49, 0x8b, 0x02, 0x49, 0x9f, 0xf2, 0x7d,
	0xca, 0x7a, 0x21, 0x03, 0xce, 0x2e, 0x6b, 0xf1, 0x70, 0xdf, 0x97, 0xa9, 0xdc, 0x84, 0x5d, 0xd1,
	0x32, 0x65, 0xea, 0x38, 0x88, 0x8c, 0xf7, 0x62, 0x4e, 0xfa, 0xce, 0x42, 0xce, 0xb0, 0x51, 0x16,
	0x6a, 0xe9, 0x4b, 0xa5, 0xa3, 0x57, 0xf9, 0x5d, 0xf4, 0xb2, 0x25, 0x3c, 0xdf, 0xa7, 0x05, 0x65,
	0x91, 0x6f, 0x80, 0x72, 0xe6, 0x8b, 0x6e, 0x18, 0x82, 0x52, 0x65, 0x75, 0xc1, 0x3b, 0x99, 0x6a,
	0x3f, 0x56, 0xca, 0xc7, 0xac, 0xa1, 0x55, 0xce, 0x1d, 0x74, 0x71, 0x22, 0x20, 0xc3, 0x35, 0xe8,
	0x88, 0xf9, 0xc0, 0xf5, 0xc2, 0xb9, 0x8d, 0xce, 0xf6, 0x37, 0x3f, 0x04, 0xbf, 0x75, 0xd5, 0x54,
	0x0b, 0xf7, 0xd1, 0x60, 0xb5, 0x0a, 0x58, 0x9d, 0xf3, 0x68, 0x75, 0xcc, 0x4e, 0x93, 0xa9, 0x2b,
	0x79, 0x32, 0x1a, 0x7f, 0x9e, 0x79, 0xa3, 0xeb, 0xf0, 0xd8, 0x65, 0xb2, 0x62, 0x99, 0x47, 0x1b,
	0x1a, 0x97, 0x2f, 0x46, 0xbc, 0x6b, 0xcf, 0x62, 0xb9, 0x37, 0xca, 0x2d, 0xbe, 0x88, 0x7a, 0x2d,
	0xc1, 0x57, 0xcf, 0xad, 0xbe, 0xad, 0xd5, 0x24, 0x67, 0x3c, 0xe3, 0x03, 0x37, 0xe2, 0x2c, 0x83,
	0xec, 0x57, 0x0b, 0xdd, 0xe8, 0x9b, 0x0f, 0x12, 0xb5, 0xbe, 0x17, 0xb0, 0x26, 0xf5, 0x1f, 0x71,
	0x09, 0x33, 0x4c, 0x18, 0x98, 0xde, 0xfe, 0x26, 0xef, 0x7e, 0x81, 0x70, 0x73, 0x03, 0x84, 0x9b,
	0x78, 0xff, 0xa0, 0x59, 0x5c, 0x1e, 0x46, 0xae, 0x49, 0x56, 0x44, 0x3c, 0x89, 0xd3, 0x85, 0x23,
	0x8e, 0x0c, 0x1e, 0x51, 0xcf, 0x53, 0x34, 0xfb, 0x78, 0x38, 0x74, 0x40, 0xc1, 0x89, 0x35, 0xe0,
	0x64, 0xf3, 0xef, 0x65, 0xb4, 0x90, 0xed, 0xc5, 0x7f, 0x59, 0xe8, 0xf4, 0xe8, 0x2e, 0x88, 0xaf,
	0xba, 0x65, 0xfb, 0x7e, 0xed, 0x9a, 0x5b, 0xbe, 0xa7, 0xfa, 0x5f, 0xbf, 0xb2, 0x9f, 0x6b, 0x23,
	0x92, 0x4d, 0xad, 0x24, 0x2b, 0x1b, 0xd9, 0xe5, 0x09, 0x91, 0x7b, 0x94, 0xb4, 0x78, 0xb3, 0x49,
	0x23, 0x12, 0x33, 0x12, 0x84, 0x21, 0xef, 0x32, 0xb9, 0x41, 0xb8, 0x9a, 0xe5, 0x60, 0xe8, 0x3c,
	0xcc, 0xcc, 0x95, 0x65, 0xd8, 0x4d, 0x52, 0x27, 0x3d, 0x17, 0xf8, 0x67, 0x0b, 0xfd, 0x7f, 0x4c,
	0x23, 0xc5, 0x79, 0xa0, 0x53, 0xfa, 0x79, 0xed, 0xba, 0x3b, 0x43, 0x63, 0x7e, 0x1b, 0xa2, 0xba,
	0x65, 0xac, 0x66, 0x09, 0x0b, 0xff, 0x68, 0xa1, 0x93, 0x23, 0xba, 0x30, 0x5e, 0x73, 0x4b, 0x35,
	0xfa, 0xda, 0x15, 0xb7, 0x5c, 0x37, 0x77, 0xde, 0x05, 0x84, 0x77, 0xc0, 0xa2, 0x90, 0xb4, 0x59,
	0x90, 0xfe, 0x63, 0xa1, 0x33, 0x63, 0x7b, 0x31, 0xde, 0x70, 0x67, 0x68, 0xfe, 0xb5, 0x1b, 0xee,
	0x2c, 0xfd, 0xdd, 0x61, 0x80, 0xfd, 0xc5, 0x6b, 0x60, 0x57, 0x9f, 0xa9, 0xba, 0x0d, 0xfa, 0x40,
	0x92, 0x30, 0x60, 0x64, 0xe7, 0x90, 0xc0, 0xf3, 0xa3, 0xfe, 0x85, 0xd2, 0xef, 0x74, 0x1f, 0xbc,
	0x68, 0x34, 0x06, 0xa1, 0x2a, 0x4a, 0x63, 0x4a, 0x51, 0x1a, 0x25, 0x8b, 0xd2, 0x98, 0x5a, 0x94,
	0xc6, 0x6b, 0x14, 0x05, 0x90, 0x2e, 0x15, 0x3a, 0x2c, 0x26, 0xee, 0x94, 0x0e, 0x5e, 0xbb, 0xe0,
	0x4e, 0xed, 0xce, 0x1f, 0x01, 0xae, 0x0f, 0xeb, 0x7b, 0x34, 0xdc, 0x27, 0xf1, 0xae, 0x3a, 0xba,
	0x09, 0x89, 0x61, 0x7d, 0x68, 0xba, 0xb3, 0x12, 0xd5, 0xfc, 0x48, 0x2c, 0x48, 0xc8, 0x01, 0x7b,
	0x28, 0x27, 0x20, 0xfd, 0xc5, 0x42, 0x2b, 0x83, 0x5d, 0x15, 0x5f, 0x74, 0xa7, 0xcf, 0x00, 0xb5,
	0x4b, 0x6e, 0x89, 0xbe, 0xec, 0x7c, 0x0a, 0x90, 0x9f, 0x66, 0x6a, 0x85, 0xa1, 0xd3, 0x0a, 0x62,
	0x26, 0xe9, 0xe7, 0x69, 0x46, 0xb5, 0xf5, 0x04, 0x86, 0x04, 0x2c, 0x82, 0xfa, 0x43, 0x0b, 0x00,
	0x11, 0x2b, 0x84, 0x87, 0x61, 0x54, 0x41, 0xfd, 0xc6, 0x8c, 0x57, 0xdd, 0x49, 0x9d, 0xbe, 0x76,
	0xce, 0x9d, 0xdc, 0xce, 0xb7, 0x01, 0xe7, 0x3b, 0xea, 0xb9, 0x3e, 0x84, 0xac, 0x31, 0x06, 0x59,
	0x03, 0x24, 0xfa, 0x1f, 0x5c, 0x68, 0x7a, 0x06, 0x44, 0x9d, 0x0e, 0x3f, 0x7c, 0x84, 0x87, 0x8a,
	0x15, 0x1a, 0x2a, 0xd9, 0x49, 0xf8, 0x01, 0x74, 0x34, 0xfc, 0x3b, 0xf0, 0x73, 0x44, 0x83, 0x2f,
	0xf0, 0x73, 0xc2, 0xa4, 0x50, 0xe0, 0xe7, 0xc4, 0x41, 0xe1, 0x39, 0x80, 0x7d, 0x66, 0x2c, 0xe0,
	0xfc, 0xac, 0xe5, 0x93, 0x74, 0x0e, 0x20, 0x3b, 0x14, 0x18, 0x0a, 0xa9, 0x93, 0x3c, 0xbb, 0x3d,
	0x5a, 0xa9, 0xfe, 0x26, 0x49, 0x10, 0x45, 0xf0, 0x9b, 0xd9, 0xe7, 0xae, 0xe8, 0xd0, 0x50, 0xdf,
	0xb6, 0x8c, 0x11, 0x3f, 0xf5, 0x1f, 0xea, 0xc1, 0x81, 0x60, 0xd4, 0x43, 0x3d, 0x6e, 0x40, 0x19,
	0xf5, 0x50, 0x8f, 0x1f, 0x30, 0xde, 0x82, 0x88, 0x36, 0xb7, 0xa4, 0xa4, 0xed, 0x4e, 0x21, 0x22,
	0x53, 0x71, 0xce, 0x46, 0x83, 0xad, 0xad, 0xc2, 0xbe, 0x33, 0x4f, 0x86, 0x6e, 0x26, 0x54, 0xe3,
	0x65, 0x1c, 0xd2, 0xcd, 0x6f, 0xe6, 0xd1, 0x72, 0xa6, 0x35, 0x3c, 0xf9, 0xc1, 0x42, 0xa7, 0x74,
	0x9d, 0x8b, 0xe3, 0x0b, 0x76, 0xdd, 0x99, 0x26, 0x9c, 0x5a, 0xc5, 0x7d, 0xc4, 0x7b, 0xf8, 0xef,
	0x03, 0x8e, 0x7a, 0x5e, 0x4d, 0x76, 0x13, 0xde, 0x56, 0x38, 0x68, 0x92, 0x96, 0x42, 0x33, 0x89,
	0xf0, 0x5d, 0x88, 0x2e, 0x54, 0xde, 0x52, 0xee, 0x0c, 0xbd, 0x29, 0xf8, 0x7b, 0x0b, 0x55, 0x35,
	0x3c, 0x3d, 0xa3, 0x14, 0xc8, 0x33, 0x61, 0x7c, 0x29, 0xc2, 0x79, 0x0c, 0x70, 0x3e, 0x28, 0x01,
	0x47, 0xf3, 0xba, 0x77, 0xb1, 0xc6, 0x53, 0xfb, 0x5b, 0xb8, 0x6c, 0x1a, 0x56, 0x7a, 0x79, 0xf0,
	0x25, 0xb7, 0xc4, 0xc0, 0x53, 0x84, 0xf4, 0x10, 0x20, 0xbd, 0x5f, 0x1a, 0xd2, 0x94, 0xab, 0x56,
	0x73, 0xc1, 0xdb, 0xda, 0x70, 0xdd, 0x8d, 0x0f, 0x96, 0x3b, 0x46, 0xfc, 0xf6, 0xca, 0x3e, 0xb2,
	0x3d, 0xf7, 0x95, 0x65, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xef, 0x59, 0xe8, 0x52, 0x9d, 0x13,
	0x00, 0x00,
}
