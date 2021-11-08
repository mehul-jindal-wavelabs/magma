// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/mconfig/mconfigs.proto

package mconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protos "magma/orc8r/lib/go/protos"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//===========================================================================
// SharedMConfig is a wrapper around configs that are shared amongst services
//===========================================================================
type SharedMconfig struct {
	SentryConfig         *SharedSentryConfig `protobuf:"bytes,1,opt,name=sentry_config,json=sentryConfig,proto3" json:"sentry_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SharedMconfig) Reset()         { *m = SharedMconfig{} }
func (m *SharedMconfig) String() string { return proto.CompactTextString(m) }
func (*SharedMconfig) ProtoMessage()    {}
func (*SharedMconfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{0}
}

func (m *SharedMconfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedMconfig.Unmarshal(m, b)
}
func (m *SharedMconfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedMconfig.Marshal(b, m, deterministic)
}
func (m *SharedMconfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedMconfig.Merge(m, src)
}
func (m *SharedMconfig) XXX_Size() int {
	return xxx_messageInfo_SharedMconfig.Size(m)
}
func (m *SharedMconfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedMconfig.DiscardUnknown(m)
}

var xxx_messageInfo_SharedMconfig proto.InternalMessageInfo

func (m *SharedMconfig) GetSentryConfig() *SharedSentryConfig {
	if m != nil {
		return m.SentryConfig
	}
	return nil
}

// --------------------------------------------------------------------------
// SentryConfig stores the network-wide Sentry.io configuration
// TODO(hcgatewood): rename, duplicated in lte/.../mconfigs
// --------------------------------------------------------------------------
type SharedSentryConfig struct {
	// dsn_python initializes the Sentry Python SDK and sets the remote URL.
	// If set to empty string, Sentry Python SDK will not be initialized.
	DsnPython string `protobuf:"bytes,1,opt,name=dsn_python,json=dsnPython,proto3" json:"dsn_python,omitempty"`
	// dsn_native initializes the Sentry Native SDK for C/C++ and sets the
	// remote URL. If set to empty string, Sentry Native SDK will not be
	// initialized.
	DsnNative string `protobuf:"bytes,2,opt,name=dsn_native,json=dsnNative,proto3" json:"dsn_native,omitempty"`
	// upload_mme_log decides whether MME service log file (/var/log/mme.log)
	// is uploaded along with MME crashreports
	UploadMmeLog bool `protobuf:"varint,3,opt,name=upload_mme_log,json=uploadMmeLog,proto3" json:"upload_mme_log,omitempty"`
	// sample_rate sets the rate at which Python error events are sampled.
	// sample_rate should be a number between 0 (0% of errors sent) and 1 (100%
	// of errors sent)
	SampleRate           float32  `protobuf:"fixed32,4,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSentryConfig) Reset()         { *m = SharedSentryConfig{} }
func (m *SharedSentryConfig) String() string { return proto.CompactTextString(m) }
func (*SharedSentryConfig) ProtoMessage()    {}
func (*SharedSentryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{1}
}

func (m *SharedSentryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSentryConfig.Unmarshal(m, b)
}
func (m *SharedSentryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSentryConfig.Marshal(b, m, deterministic)
}
func (m *SharedSentryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSentryConfig.Merge(m, src)
}
func (m *SharedSentryConfig) XXX_Size() int {
	return xxx_messageInfo_SharedSentryConfig.Size(m)
}
func (m *SharedSentryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSentryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSentryConfig proto.InternalMessageInfo

func (m *SharedSentryConfig) GetDsnPython() string {
	if m != nil {
		return m.DsnPython
	}
	return ""
}

func (m *SharedSentryConfig) GetDsnNative() string {
	if m != nil {
		return m.DsnNative
	}
	return ""
}

func (m *SharedSentryConfig) GetUploadMmeLog() bool {
	if m != nil {
		return m.UploadMmeLog
	}
	return false
}

func (m *SharedSentryConfig) GetSampleRate() float32 {
	if m != nil {
		return m.SampleRate
	}
	return 0
}

//------------------------------------------------------------------------------
// Control Proxy configs
//------------------------------------------------------------------------------
type ControlProxy struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ControlProxy) Reset()         { *m = ControlProxy{} }
func (m *ControlProxy) String() string { return proto.CompactTextString(m) }
func (*ControlProxy) ProtoMessage()    {}
func (*ControlProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{2}
}

func (m *ControlProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlProxy.Unmarshal(m, b)
}
func (m *ControlProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlProxy.Marshal(b, m, deterministic)
}
func (m *ControlProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlProxy.Merge(m, src)
}
func (m *ControlProxy) XXX_Size() int {
	return xxx_messageInfo_ControlProxy.Size(m)
}
func (m *ControlProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ControlProxy proto.InternalMessageInfo

func (m *ControlProxy) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

type ImageSpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Order                int64    `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageSpec) Reset()         { *m = ImageSpec{} }
func (m *ImageSpec) String() string { return proto.CompactTextString(m) }
func (*ImageSpec) ProtoMessage()    {}
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{3}
}

func (m *ImageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSpec.Unmarshal(m, b)
}
func (m *ImageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSpec.Marshal(b, m, deterministic)
}
func (m *ImageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSpec.Merge(m, src)
}
func (m *ImageSpec) XXX_Size() int {
	return xxx_messageInfo_ImageSpec.Size(m)
}
func (m *ImageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSpec proto.InternalMessageInfo

func (m *ImageSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageSpec) GetOrder() int64 {
	if m != nil {
		return m.Order
	}
	return 0
}

type MagmaD struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	// Interval for the gateways to send checkin rpc calls to the cloud.
	CheckinInterval int32 `protobuf:"varint,2,opt,name=checkin_interval,json=checkinInterval,proto3" json:"checkin_interval,omitempty"`
	// Checkin rpc timeout
	CheckinTimeout int32 `protobuf:"varint,3,opt,name=checkin_timeout,json=checkinTimeout,proto3" json:"checkin_timeout,omitempty"`
	// Enables autoupgrading of the magma package
	AutoupgradeEnabled bool `protobuf:"varint,4,opt,name=autoupgrade_enabled,json=autoupgradeEnabled,proto3" json:"autoupgrade_enabled,omitempty"`
	// Interval to poll for package upgrades
	AutoupgradePollInterval int32 `protobuf:"varint,5,opt,name=autoupgrade_poll_interval,json=autoupgradePollInterval,proto3" json:"autoupgrade_poll_interval,omitempty"`
	// The magma package version the gateway should upgrade to
	PackageVersion string `protobuf:"bytes,6,opt,name=package_version,json=packageVersion,proto3" json:"package_version,omitempty"`
	// List of upgrade images
	Images []*ImageSpec `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	// For streamer, should be left unused by gateway
	TierId       string          `protobuf:"bytes,8,opt,name=tier_id,json=tierId,proto3" json:"tier_id,omitempty"`
	FeatureFlags map[string]bool `protobuf:"bytes,9,rep,name=feature_flags,json=featureFlags,proto3" json:"feature_flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of dynamic_services
	DynamicServices      []string `protobuf:"bytes,10,rep,name=dynamic_services,json=dynamicServices,proto3" json:"dynamic_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MagmaD) Reset()         { *m = MagmaD{} }
func (m *MagmaD) String() string { return proto.CompactTextString(m) }
func (*MagmaD) ProtoMessage()    {}
func (*MagmaD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{4}
}

func (m *MagmaD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MagmaD.Unmarshal(m, b)
}
func (m *MagmaD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MagmaD.Marshal(b, m, deterministic)
}
func (m *MagmaD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MagmaD.Merge(m, src)
}
func (m *MagmaD) XXX_Size() int {
	return xxx_messageInfo_MagmaD.Size(m)
}
func (m *MagmaD) XXX_DiscardUnknown() {
	xxx_messageInfo_MagmaD.DiscardUnknown(m)
}

var xxx_messageInfo_MagmaD proto.InternalMessageInfo

func (m *MagmaD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *MagmaD) GetCheckinInterval() int32 {
	if m != nil {
		return m.CheckinInterval
	}
	return 0
}

func (m *MagmaD) GetCheckinTimeout() int32 {
	if m != nil {
		return m.CheckinTimeout
	}
	return 0
}

func (m *MagmaD) GetAutoupgradeEnabled() bool {
	if m != nil {
		return m.AutoupgradeEnabled
	}
	return false
}

func (m *MagmaD) GetAutoupgradePollInterval() int32 {
	if m != nil {
		return m.AutoupgradePollInterval
	}
	return 0
}

func (m *MagmaD) GetPackageVersion() string {
	if m != nil {
		return m.PackageVersion
	}
	return ""
}

func (m *MagmaD) GetImages() []*ImageSpec {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *MagmaD) GetTierId() string {
	if m != nil {
		return m.TierId
	}
	return ""
}

func (m *MagmaD) GetFeatureFlags() map[string]bool {
	if m != nil {
		return m.FeatureFlags
	}
	return nil
}

func (m *MagmaD) GetDynamicServices() []string {
	if m != nil {
		return m.DynamicServices
	}
	return nil
}

//------------------------------------------------------------------------------
// EventD configs
//------------------------------------------------------------------------------
type EventD struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	// The verbosity level for events.
	// All events less than or equal to this verbosity will be logged.
	EventVerbosity       int32    `protobuf:"varint,2,opt,name=event_verbosity,json=eventVerbosity,proto3" json:"event_verbosity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventD) Reset()         { *m = EventD{} }
func (m *EventD) String() string { return proto.CompactTextString(m) }
func (*EventD) ProtoMessage()    {}
func (*EventD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{5}
}

func (m *EventD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventD.Unmarshal(m, b)
}
func (m *EventD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventD.Marshal(b, m, deterministic)
}
func (m *EventD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventD.Merge(m, src)
}
func (m *EventD) XXX_Size() int {
	return xxx_messageInfo_EventD.Size(m)
}
func (m *EventD) XXX_DiscardUnknown() {
	xxx_messageInfo_EventD.DiscardUnknown(m)
}

var xxx_messageInfo_EventD proto.InternalMessageInfo

func (m *EventD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *EventD) GetEventVerbosity() int32 {
	if m != nil {
		return m.EventVerbosity
	}
	return 0
}

type DirectoryD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DirectoryD) Reset()         { *m = DirectoryD{} }
func (m *DirectoryD) String() string { return proto.CompactTextString(m) }
func (*DirectoryD) ProtoMessage()    {}
func (*DirectoryD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{6}
}

func (m *DirectoryD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryD.Unmarshal(m, b)
}
func (m *DirectoryD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryD.Marshal(b, m, deterministic)
}
func (m *DirectoryD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryD.Merge(m, src)
}
func (m *DirectoryD) XXX_Size() int {
	return xxx_messageInfo_DirectoryD.Size(m)
}
func (m *DirectoryD) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryD.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryD proto.InternalMessageInfo

func (m *DirectoryD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

//------------------------------------------------------------------------------
// MetricsD configs
//------------------------------------------------------------------------------
type MetricsD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MetricsD) Reset()         { *m = MetricsD{} }
func (m *MetricsD) String() string { return proto.CompactTextString(m) }
func (*MetricsD) ProtoMessage()    {}
func (*MetricsD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{7}
}

func (m *MetricsD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsD.Unmarshal(m, b)
}
func (m *MetricsD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsD.Marshal(b, m, deterministic)
}
func (m *MetricsD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsD.Merge(m, src)
}
func (m *MetricsD) XXX_Size() int {
	return xxx_messageInfo_MetricsD.Size(m)
}
func (m *MetricsD) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsD.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsD proto.InternalMessageInfo

func (m *MetricsD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

//------------------------------------------------------------------------------
// State configs
//------------------------------------------------------------------------------
type State struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	SyncInterval         uint32          `protobuf:"varint,2,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{8}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *State) GetSyncInterval() uint32 {
	if m != nil {
		return m.SyncInterval
	}
	return 0
}

//------------------------------------------------------------------------------
// Fluent Bit configs
//------------------------------------------------------------------------------
type FluentBit struct {
	ExtraTags            map[string]string `protobuf:"bytes,1,rep,name=extra_tags,json=extraTags,proto3" json:"extra_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ThrottleRate         uint32            `protobuf:"varint,10,opt,name=throttle_rate,json=throttleRate,proto3" json:"throttle_rate,omitempty"`
	ThrottleWindow       uint32            `protobuf:"varint,11,opt,name=throttle_window,json=throttleWindow,proto3" json:"throttle_window,omitempty"`
	ThrottleInterval     string            `protobuf:"bytes,12,opt,name=throttle_interval,json=throttleInterval,proto3" json:"throttle_interval,omitempty"`
	FilesByTag           map[string]string `protobuf:"bytes,20,rep,name=files_by_tag,json=filesByTag,proto3" json:"files_by_tag,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FluentBit) Reset()         { *m = FluentBit{} }
func (m *FluentBit) String() string { return proto.CompactTextString(m) }
func (*FluentBit) ProtoMessage()    {}
func (*FluentBit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{9}
}

func (m *FluentBit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FluentBit.Unmarshal(m, b)
}
func (m *FluentBit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FluentBit.Marshal(b, m, deterministic)
}
func (m *FluentBit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FluentBit.Merge(m, src)
}
func (m *FluentBit) XXX_Size() int {
	return xxx_messageInfo_FluentBit.Size(m)
}
func (m *FluentBit) XXX_DiscardUnknown() {
	xxx_messageInfo_FluentBit.DiscardUnknown(m)
}

var xxx_messageInfo_FluentBit proto.InternalMessageInfo

func (m *FluentBit) GetExtraTags() map[string]string {
	if m != nil {
		return m.ExtraTags
	}
	return nil
}

func (m *FluentBit) GetThrottleRate() uint32 {
	if m != nil {
		return m.ThrottleRate
	}
	return 0
}

func (m *FluentBit) GetThrottleWindow() uint32 {
	if m != nil {
		return m.ThrottleWindow
	}
	return 0
}

func (m *FluentBit) GetThrottleInterval() string {
	if m != nil {
		return m.ThrottleInterval
	}
	return ""
}

func (m *FluentBit) GetFilesByTag() map[string]string {
	if m != nil {
		return m.FilesByTag
	}
	return nil
}

//------------------------------------------------------------------------------
// OpenVPN client configs
//------------------------------------------------------------------------------
type OpenVPN struct {
	EnableShellAccess    bool     `protobuf:"varint,1,opt,name=enable_shell_access,json=enableShellAccess,proto3" json:"enable_shell_access,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenVPN) Reset()         { *m = OpenVPN{} }
func (m *OpenVPN) String() string { return proto.CompactTextString(m) }
func (*OpenVPN) ProtoMessage()    {}
func (*OpenVPN) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{10}
}

func (m *OpenVPN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenVPN.Unmarshal(m, b)
}
func (m *OpenVPN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenVPN.Marshal(b, m, deterministic)
}
func (m *OpenVPN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenVPN.Merge(m, src)
}
func (m *OpenVPN) XXX_Size() int {
	return xxx_messageInfo_OpenVPN.Size(m)
}
func (m *OpenVPN) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenVPN.DiscardUnknown(m)
}

var xxx_messageInfo_OpenVPN proto.InternalMessageInfo

func (m *OpenVPN) GetEnableShellAccess() bool {
	if m != nil {
		return m.EnableShellAccess
	}
	return false
}

//------------------------------------------------------------------------------
// CtraceD configs
//------------------------------------------------------------------------------
type CtraceD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CtraceD) Reset()         { *m = CtraceD{} }
func (m *CtraceD) String() string { return proto.CompactTextString(m) }
func (*CtraceD) ProtoMessage()    {}
func (*CtraceD) Descriptor() ([]byte, []int) {
	return fileDescriptor_9618f358f05ec5b9, []int{11}
}

func (m *CtraceD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CtraceD.Unmarshal(m, b)
}
func (m *CtraceD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CtraceD.Marshal(b, m, deterministic)
}
func (m *CtraceD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CtraceD.Merge(m, src)
}
func (m *CtraceD) XXX_Size() int {
	return xxx_messageInfo_CtraceD.Size(m)
}
func (m *CtraceD) XXX_DiscardUnknown() {
	xxx_messageInfo_CtraceD.DiscardUnknown(m)
}

var xxx_messageInfo_CtraceD proto.InternalMessageInfo

func (m *CtraceD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func init() {
	proto.RegisterType((*SharedMconfig)(nil), "magma.mconfig.SharedMconfig")
	proto.RegisterType((*SharedSentryConfig)(nil), "magma.mconfig.SharedSentryConfig")
	proto.RegisterType((*ControlProxy)(nil), "magma.mconfig.ControlProxy")
	proto.RegisterType((*ImageSpec)(nil), "magma.mconfig.ImageSpec")
	proto.RegisterType((*MagmaD)(nil), "magma.mconfig.MagmaD")
	proto.RegisterMapType((map[string]bool)(nil), "magma.mconfig.MagmaD.FeatureFlagsEntry")
	proto.RegisterType((*EventD)(nil), "magma.mconfig.EventD")
	proto.RegisterType((*DirectoryD)(nil), "magma.mconfig.DirectoryD")
	proto.RegisterType((*MetricsD)(nil), "magma.mconfig.MetricsD")
	proto.RegisterType((*State)(nil), "magma.mconfig.State")
	proto.RegisterType((*FluentBit)(nil), "magma.mconfig.FluentBit")
	proto.RegisterMapType((map[string]string)(nil), "magma.mconfig.FluentBit.ExtraTagsEntry")
	proto.RegisterMapType((map[string]string)(nil), "magma.mconfig.FluentBit.FilesByTagEntry")
	proto.RegisterType((*OpenVPN)(nil), "magma.mconfig.OpenVPN")
	proto.RegisterType((*CtraceD)(nil), "magma.mconfig.CtraceD")
}

func init() {
	proto.RegisterFile("orc8r/protos/mconfig/mconfigs.proto", fileDescriptor_9618f358f05ec5b9)
}

var fileDescriptor_9618f358f05ec5b9 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x6f, 0x6b, 0xe3, 0x46,
	0x10, 0xc6, 0x71, 0x9c, 0x38, 0xf6, 0xc4, 0x76, 0x92, 0xbd, 0x2b, 0xa7, 0x0b, 0x94, 0xfa, 0x94,
	0x42, 0x5c, 0x0a, 0x76, 0x49, 0x29, 0x5c, 0x8f, 0x5e, 0xff, 0x24, 0x17, 0x43, 0x4a, 0x72, 0x0d,
	0x72, 0xc8, 0x41, 0xdf, 0xa8, 0x6b, 0x69, 0x2c, 0x2f, 0x59, 0xed, 0x9a, 0xdd, 0xb5, 0xef, 0xf4,
	0x49, 0x4a, 0x3f, 0x43, 0xbf, 0x64, 0xd9, 0x5d, 0xc9, 0x75, 0x72, 0x2d, 0xa5, 0x7e, 0x65, 0xcd,
	0x6f, 0x9e, 0x99, 0xd5, 0x8c, 0x1e, 0x4b, 0x70, 0x2c, 0x55, 0xf2, 0x52, 0x0d, 0xe7, 0x4a, 0x1a,
	0xa9, 0x87, 0x79, 0x22, 0xc5, 0x94, 0x65, 0xd5, 0xaf, 0x1e, 0x38, 0x4e, 0x3a, 0x39, 0xcd, 0x72,
	0x3a, 0x28, 0xe9, 0xd1, 0xf3, 0x07, 0x35, 0x89, 0xcc, 0x73, 0x29, 0xbc, 0x32, 0x7c, 0x07, 0x9d,
	0xf1, 0x8c, 0x2a, 0x4c, 0xaf, 0xbd, 0x96, 0x8c, 0xa0, 0xa3, 0x51, 0x18, 0x55, 0xc4, 0x1e, 0x04,
	0xb5, 0x5e, 0xad, 0xbf, 0x77, 0xfa, 0x62, 0xf0, 0xa0, 0xe5, 0xc0, 0x17, 0x8d, 0x9d, 0xf2, 0xdc,
	0xa1, 0xa8, 0xad, 0xd7, 0xa2, 0xf0, 0x8f, 0x1a, 0x90, 0x8f, 0x45, 0xe4, 0x53, 0x80, 0x54, 0x8b,
	0x78, 0x5e, 0x98, 0x99, 0x14, 0xae, 0x77, 0x2b, 0x6a, 0xa5, 0x5a, 0xdc, 0x38, 0x50, 0xa5, 0x05,
	0x35, 0x6c, 0x89, 0xc1, 0xd6, 0x2a, 0xfd, 0xd6, 0x01, 0xf2, 0x39, 0x74, 0x17, 0x73, 0x2e, 0x69,
	0x1a, 0xe7, 0x39, 0xc6, 0x5c, 0x66, 0x41, 0xbd, 0x57, 0xeb, 0x37, 0xa3, 0xb6, 0xa7, 0xd7, 0x39,
	0x5e, 0xc9, 0x8c, 0x7c, 0x06, 0x7b, 0x9a, 0xe6, 0x73, 0x8e, 0xb1, 0xa2, 0x06, 0x83, 0xed, 0x5e,
	0xad, 0xbf, 0x15, 0x81, 0x47, 0x11, 0x35, 0x18, 0x9e, 0x41, 0xfb, 0x5c, 0x0a, 0xa3, 0x24, 0xbf,
	0x51, 0xf2, 0x43, 0x41, 0x4e, 0xa1, 0xc5, 0x65, 0x16, 0x73, 0x5c, 0x22, 0x77, 0xf7, 0xd4, 0x3d,
	0xfd, 0xa4, 0x9c, 0xd7, 0x6d, 0x6e, 0x70, 0x25, 0xb3, 0x2b, 0x9b, 0x8c, 0x9a, 0xbc, 0xbc, 0x0a,
	0xbf, 0x81, 0xd6, 0x65, 0x4e, 0x33, 0x1c, 0xcf, 0x31, 0x21, 0x04, 0xb6, 0x05, 0xcd, 0xb1, 0x9c,
	0xc7, 0x5d, 0x93, 0xa7, 0xb0, 0x23, 0x55, 0x8a, 0xca, 0x4d, 0x51, 0x8f, 0x7c, 0x10, 0xfe, 0xb9,
	0x0d, 0x8d, 0x6b, 0xdb, 0xf9, 0xcd, 0x26, 0xa7, 0x92, 0x2f, 0xe0, 0x20, 0x99, 0x61, 0x72, 0xcf,
	0x44, 0xcc, 0x84, 0x41, 0xb5, 0xa4, 0xdc, 0xf5, 0xdf, 0x89, 0xf6, 0x4b, 0x7e, 0x59, 0x62, 0x72,
	0x02, 0x15, 0x8a, 0x0d, 0xcb, 0x51, 0x2e, 0x8c, 0x5b, 0xd6, 0x4e, 0xd4, 0x2d, 0xf1, 0xad, 0xa7,
	0x64, 0x08, 0x4f, 0xe8, 0xc2, 0xc8, 0xc5, 0x3c, 0x53, 0x34, 0xc5, 0x18, 0x05, 0x9d, 0x70, 0x4c,
	0xdd, 0xda, 0x9a, 0x11, 0x59, 0x4b, 0x5d, 0xf8, 0x0c, 0x79, 0x05, 0xcf, 0xd7, 0x0b, 0xe6, 0x92,
	0xf3, 0xbf, 0xef, 0x66, 0xc7, 0x9d, 0xf1, 0x6c, 0x4d, 0x70, 0x23, 0x39, 0x5f, 0xbf, 0xab, 0x39,
	0x4d, 0xee, 0x69, 0x86, 0xf1, 0x12, 0x95, 0x66, 0x52, 0x04, 0x0d, 0xb7, 0xb4, 0x6e, 0x89, 0xef,
	0x3c, 0x25, 0x5f, 0x41, 0x83, 0xd9, 0xfd, 0xea, 0x60, 0xb7, 0x57, 0xef, 0xef, 0x9d, 0x06, 0x8f,
	0x0c, 0xb8, 0x5a, 0x7e, 0x54, 0xea, 0xc8, 0x33, 0xd8, 0x35, 0x0c, 0x55, 0xcc, 0xd2, 0xa0, 0xe9,
	0x5a, 0x36, 0x6c, 0x78, 0x99, 0x92, 0x2b, 0xe8, 0x4c, 0x91, 0x9a, 0x85, 0xc2, 0x78, 0xca, 0x69,
	0xa6, 0x83, 0x96, 0xeb, 0x78, 0xf2, 0xa8, 0xa3, 0x7f, 0x2c, 0x83, 0x91, 0x97, 0x8e, 0xac, 0xf2,
	0xc2, 0x3a, 0x37, 0x6a, 0x4f, 0xd7, 0x90, 0x7d, 0x04, 0x69, 0x21, 0x68, 0xce, 0x92, 0x58, 0xa3,
	0x5a, 0xb2, 0x04, 0x75, 0x00, 0xbd, 0x7a, 0xbf, 0x15, 0xed, 0x97, 0x7c, 0x5c, 0xe2, 0xa3, 0x1f,
	0xe0, 0xf0, 0xa3, 0x6e, 0xe4, 0x00, 0xea, 0xf7, 0x58, 0x94, 0x56, 0xb1, 0x97, 0xd6, 0x29, 0x4b,
	0xca, 0x17, 0xde, 0xef, 0xcd, 0xc8, 0x07, 0xaf, 0xb6, 0x5e, 0xd6, 0x42, 0x84, 0xc6, 0xc5, 0x12,
	0x85, 0xd9, 0xcc, 0x2c, 0x27, 0xb0, 0x8f, 0xb6, 0xda, 0x6e, 0x7a, 0x22, 0x35, 0x33, 0x45, 0xe9,
	0x95, 0xae, 0xc3, 0x77, 0x15, 0x0d, 0x7f, 0x04, 0x78, 0xc3, 0x14, 0x26, 0x46, 0xaa, 0x62, 0xa3,
	0xa3, 0xc2, 0xef, 0xa1, 0x79, 0x8d, 0x46, 0xb1, 0x44, 0x6f, 0x56, 0xff, 0x1b, 0xec, 0x8c, 0x0d,
	0x35, 0xb8, 0xd1, 0x9c, 0xc7, 0xd0, 0xd1, 0x85, 0x48, 0x1e, 0xfe, 0x23, 0x3a, 0x51, 0xdb, 0xc2,
	0xca, 0x78, 0xe1, 0xef, 0x75, 0x68, 0x8d, 0xf8, 0x02, 0x85, 0x39, 0x63, 0x86, 0x8c, 0x00, 0xf0,
	0x83, 0x51, 0x34, 0x36, 0xd6, 0x0f, 0xb5, 0x7f, 0xf4, 0xc3, 0x4a, 0x3d, 0xb8, 0xb0, 0xd2, 0xdb,
	0x95, 0x1f, 0x5a, 0x58, 0xc5, 0xf6, 0x68, 0x33, 0x53, 0xd2, 0x98, 0xea, 0x65, 0x03, 0xfe, 0xe8,
	0x0a, 0xda, 0xd7, 0x8d, 0x7d, 0x0e, 0x2b, 0xd1, 0x7b, 0x26, 0x52, 0xf9, 0x3e, 0xd8, 0x73, 0xb2,
	0x6e, 0x85, 0xdf, 0x39, 0x4a, 0xbe, 0x84, 0xc3, 0x95, 0x70, 0x35, 0x4c, 0xdb, 0x19, 0xe5, 0xa0,
	0x4a, 0xac, 0xfe, 0x49, 0x3f, 0x43, 0x7b, 0xca, 0x38, 0xea, 0x78, 0x52, 0xd8, 0x29, 0x82, 0xa7,
	0x6e, 0x88, 0xfe, 0xbf, 0x0e, 0x31, 0xb2, 0xe2, 0xb3, 0xe2, 0x96, 0x66, 0x7e, 0x0a, 0x98, 0xae,
	0xc0, 0xd1, 0x77, 0xd0, 0x7d, 0x38, 0xe3, 0x7f, 0xb9, 0xb4, 0xb5, 0xe6, 0xd2, 0xa3, 0xd7, 0xb0,
	0xff, 0xa8, 0xf9, 0xff, 0x29, 0x0f, 0xbf, 0x85, 0xdd, 0x5f, 0xe6, 0x28, 0xee, 0x6e, 0xde, 0x92,
	0x01, 0x3c, 0xf1, 0xaf, 0x9f, 0x58, 0xcf, 0x90, 0xf3, 0x98, 0x26, 0x09, 0x6a, 0xed, 0xda, 0x34,
	0xa3, 0x43, 0x9f, 0x1a, 0xdb, 0xcc, 0x4f, 0x2e, 0x11, 0xbe, 0x86, 0xdd, 0x73, 0xa3, 0x68, 0x82,
	0x1b, 0xb9, 0xee, 0xec, 0xf8, 0xd7, 0x17, 0x4e, 0x31, 0xf4, 0xdf, 0x47, 0xce, 0x26, 0xc3, 0x4c,
	0x3e, 0xfa, 0xb4, 0x4e, 0x1a, 0x2e, 0xfe, 0xfa, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x20,
	0x8a, 0xbe, 0x79, 0x07, 0x00, 0x00,
}
