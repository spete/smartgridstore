// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adminapi.proto

package adminapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResetAPIKeyParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetAPIKeyParams) Reset()         { *m = ResetAPIKeyParams{} }
func (m *ResetAPIKeyParams) String() string { return proto.CompactTextString(m) }
func (*ResetAPIKeyParams) ProtoMessage()    {}
func (*ResetAPIKeyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{0}
}
func (m *ResetAPIKeyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetAPIKeyParams.Unmarshal(m, b)
}
func (m *ResetAPIKeyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetAPIKeyParams.Marshal(b, m, deterministic)
}
func (dst *ResetAPIKeyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetAPIKeyParams.Merge(dst, src)
}
func (m *ResetAPIKeyParams) XXX_Size() int {
	return xxx_messageInfo_ResetAPIKeyParams.Size(m)
}
func (m *ResetAPIKeyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetAPIKeyParams.DiscardUnknown(m)
}

var xxx_messageInfo_ResetAPIKeyParams proto.InternalMessageInfo

type GetAPIKeyParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAPIKeyParams) Reset()         { *m = GetAPIKeyParams{} }
func (m *GetAPIKeyParams) String() string { return proto.CompactTextString(m) }
func (*GetAPIKeyParams) ProtoMessage()    {}
func (*GetAPIKeyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{1}
}
func (m *GetAPIKeyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAPIKeyParams.Unmarshal(m, b)
}
func (m *GetAPIKeyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAPIKeyParams.Marshal(b, m, deterministic)
}
func (dst *GetAPIKeyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAPIKeyParams.Merge(dst, src)
}
func (m *GetAPIKeyParams) XXX_Size() int {
	return xxx_messageInfo_GetAPIKeyParams.Size(m)
}
func (m *GetAPIKeyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAPIKeyParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAPIKeyParams proto.InternalMessageInfo

type APIKeyResponse struct {
	Stat                 *Status  `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	Apikey               string   `protobuf:"bytes,2,opt,name=apikey" json:"apikey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIKeyResponse) Reset()         { *m = APIKeyResponse{} }
func (m *APIKeyResponse) String() string { return proto.CompactTextString(m) }
func (*APIKeyResponse) ProtoMessage()    {}
func (*APIKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{2}
}
func (m *APIKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIKeyResponse.Unmarshal(m, b)
}
func (m *APIKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIKeyResponse.Marshal(b, m, deterministic)
}
func (dst *APIKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIKeyResponse.Merge(dst, src)
}
func (m *APIKeyResponse) XXX_Size() int {
	return xxx_messageInfo_APIKeyResponse.Size(m)
}
func (m *APIKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_APIKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_APIKeyResponse proto.InternalMessageInfo

func (m *APIKeyResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *APIKeyResponse) GetApikey() string {
	if m != nil {
		return m.Apikey
	}
	return ""
}

type ManifestAddParams struct {
	Deviceid             string          `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Metadata             []*MetaKeyValue `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ManifestAddParams) Reset()         { *m = ManifestAddParams{} }
func (m *ManifestAddParams) String() string { return proto.CompactTextString(m) }
func (*ManifestAddParams) ProtoMessage()    {}
func (*ManifestAddParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{3}
}
func (m *ManifestAddParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestAddParams.Unmarshal(m, b)
}
func (m *ManifestAddParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestAddParams.Marshal(b, m, deterministic)
}
func (dst *ManifestAddParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestAddParams.Merge(dst, src)
}
func (m *ManifestAddParams) XXX_Size() int {
	return xxx_messageInfo_ManifestAddParams.Size(m)
}
func (m *ManifestAddParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestAddParams.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestAddParams proto.InternalMessageInfo

func (m *ManifestAddParams) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

func (m *ManifestAddParams) GetMetadata() []*MetaKeyValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ManifestAddResponse struct {
	Stat                 *Status  `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestAddResponse) Reset()         { *m = ManifestAddResponse{} }
func (m *ManifestAddResponse) String() string { return proto.CompactTextString(m) }
func (*ManifestAddResponse) ProtoMessage()    {}
func (*ManifestAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{4}
}
func (m *ManifestAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestAddResponse.Unmarshal(m, b)
}
func (m *ManifestAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestAddResponse.Marshal(b, m, deterministic)
}
func (dst *ManifestAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestAddResponse.Merge(dst, src)
}
func (m *ManifestAddResponse) XXX_Size() int {
	return xxx_messageInfo_ManifestAddResponse.Size(m)
}
func (m *ManifestAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestAddResponse proto.InternalMessageInfo

func (m *ManifestAddResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type MetaKeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaKeyValue) Reset()         { *m = MetaKeyValue{} }
func (m *MetaKeyValue) String() string { return proto.CompactTextString(m) }
func (*MetaKeyValue) ProtoMessage()    {}
func (*MetaKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{5}
}
func (m *MetaKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaKeyValue.Unmarshal(m, b)
}
func (m *MetaKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaKeyValue.Marshal(b, m, deterministic)
}
func (dst *MetaKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaKeyValue.Merge(dst, src)
}
func (m *MetaKeyValue) XXX_Size() int {
	return xxx_messageInfo_MetaKeyValue.Size(m)
}
func (m *MetaKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_MetaKeyValue proto.InternalMessageInfo

func (m *MetaKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MetaKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ManifestDelParams struct {
	Deviceid             string   `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestDelParams) Reset()         { *m = ManifestDelParams{} }
func (m *ManifestDelParams) String() string { return proto.CompactTextString(m) }
func (*ManifestDelParams) ProtoMessage()    {}
func (*ManifestDelParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{6}
}
func (m *ManifestDelParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDelParams.Unmarshal(m, b)
}
func (m *ManifestDelParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDelParams.Marshal(b, m, deterministic)
}
func (dst *ManifestDelParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDelParams.Merge(dst, src)
}
func (m *ManifestDelParams) XXX_Size() int {
	return xxx_messageInfo_ManifestDelParams.Size(m)
}
func (m *ManifestDelParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDelParams.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDelParams proto.InternalMessageInfo

func (m *ManifestDelParams) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

type ManifestDelResponse struct {
	Stat                 *Status  `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestDelResponse) Reset()         { *m = ManifestDelResponse{} }
func (m *ManifestDelResponse) String() string { return proto.CompactTextString(m) }
func (*ManifestDelResponse) ProtoMessage()    {}
func (*ManifestDelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{7}
}
func (m *ManifestDelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDelResponse.Unmarshal(m, b)
}
func (m *ManifestDelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDelResponse.Marshal(b, m, deterministic)
}
func (dst *ManifestDelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDelResponse.Merge(dst, src)
}
func (m *ManifestDelResponse) XXX_Size() int {
	return xxx_messageInfo_ManifestDelResponse.Size(m)
}
func (m *ManifestDelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDelResponse proto.InternalMessageInfo

func (m *ManifestDelResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type ManifestDelPrefixParams struct {
	Deviceidprefix       string   `protobuf:"bytes,1,opt,name=deviceidprefix" json:"deviceidprefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestDelPrefixParams) Reset()         { *m = ManifestDelPrefixParams{} }
func (m *ManifestDelPrefixParams) String() string { return proto.CompactTextString(m) }
func (*ManifestDelPrefixParams) ProtoMessage()    {}
func (*ManifestDelPrefixParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{8}
}
func (m *ManifestDelPrefixParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDelPrefixParams.Unmarshal(m, b)
}
func (m *ManifestDelPrefixParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDelPrefixParams.Marshal(b, m, deterministic)
}
func (dst *ManifestDelPrefixParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDelPrefixParams.Merge(dst, src)
}
func (m *ManifestDelPrefixParams) XXX_Size() int {
	return xxx_messageInfo_ManifestDelPrefixParams.Size(m)
}
func (m *ManifestDelPrefixParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDelPrefixParams.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDelPrefixParams proto.InternalMessageInfo

func (m *ManifestDelPrefixParams) GetDeviceidprefix() string {
	if m != nil {
		return m.Deviceidprefix
	}
	return ""
}

type ManifestDelPrefixResponse struct {
	Stat                 *Status  `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	Numdeleted           uint32   `protobuf:"varint,2,opt,name=numdeleted" json:"numdeleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestDelPrefixResponse) Reset()         { *m = ManifestDelPrefixResponse{} }
func (m *ManifestDelPrefixResponse) String() string { return proto.CompactTextString(m) }
func (*ManifestDelPrefixResponse) ProtoMessage()    {}
func (*ManifestDelPrefixResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{9}
}
func (m *ManifestDelPrefixResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDelPrefixResponse.Unmarshal(m, b)
}
func (m *ManifestDelPrefixResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDelPrefixResponse.Marshal(b, m, deterministic)
}
func (dst *ManifestDelPrefixResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDelPrefixResponse.Merge(dst, src)
}
func (m *ManifestDelPrefixResponse) XXX_Size() int {
	return xxx_messageInfo_ManifestDelPrefixResponse.Size(m)
}
func (m *ManifestDelPrefixResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDelPrefixResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDelPrefixResponse proto.InternalMessageInfo

func (m *ManifestDelPrefixResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *ManifestDelPrefixResponse) GetNumdeleted() uint32 {
	if m != nil {
		return m.Numdeleted
	}
	return 0
}

type ManifestLsDevsParams struct {
	Deviceidprefix       string   `protobuf:"bytes,1,opt,name=deviceidprefix" json:"deviceidprefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestLsDevsParams) Reset()         { *m = ManifestLsDevsParams{} }
func (m *ManifestLsDevsParams) String() string { return proto.CompactTextString(m) }
func (*ManifestLsDevsParams) ProtoMessage()    {}
func (*ManifestLsDevsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{10}
}
func (m *ManifestLsDevsParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestLsDevsParams.Unmarshal(m, b)
}
func (m *ManifestLsDevsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestLsDevsParams.Marshal(b, m, deterministic)
}
func (dst *ManifestLsDevsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestLsDevsParams.Merge(dst, src)
}
func (m *ManifestLsDevsParams) XXX_Size() int {
	return xxx_messageInfo_ManifestLsDevsParams.Size(m)
}
func (m *ManifestLsDevsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestLsDevsParams.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestLsDevsParams proto.InternalMessageInfo

func (m *ManifestLsDevsParams) GetDeviceidprefix() string {
	if m != nil {
		return m.Deviceidprefix
	}
	return ""
}

type ManifestLsDevsResponse struct {
	Stat                 *Status           `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	Devices              []*ManifestDevice `protobuf:"bytes,2,rep,name=devices" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ManifestLsDevsResponse) Reset()         { *m = ManifestLsDevsResponse{} }
func (m *ManifestLsDevsResponse) String() string { return proto.CompactTextString(m) }
func (*ManifestLsDevsResponse) ProtoMessage()    {}
func (*ManifestLsDevsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{11}
}
func (m *ManifestLsDevsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestLsDevsResponse.Unmarshal(m, b)
}
func (m *ManifestLsDevsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestLsDevsResponse.Marshal(b, m, deterministic)
}
func (dst *ManifestLsDevsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestLsDevsResponse.Merge(dst, src)
}
func (m *ManifestLsDevsResponse) XXX_Size() int {
	return xxx_messageInfo_ManifestLsDevsResponse.Size(m)
}
func (m *ManifestLsDevsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestLsDevsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestLsDevsResponse proto.InternalMessageInfo

func (m *ManifestLsDevsResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *ManifestLsDevsResponse) GetDevices() []*ManifestDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

type ManifestDevice struct {
	Deviceid             string          `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Metadata             []*MetaKeyValue `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ManifestDevice) Reset()         { *m = ManifestDevice{} }
func (m *ManifestDevice) String() string { return proto.CompactTextString(m) }
func (*ManifestDevice) ProtoMessage()    {}
func (*ManifestDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{12}
}
func (m *ManifestDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDevice.Unmarshal(m, b)
}
func (m *ManifestDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDevice.Marshal(b, m, deterministic)
}
func (dst *ManifestDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDevice.Merge(dst, src)
}
func (m *ManifestDevice) XXX_Size() int {
	return xxx_messageInfo_ManifestDevice.Size(m)
}
func (m *ManifestDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDevice.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDevice proto.InternalMessageInfo

func (m *ManifestDevice) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

func (m *ManifestDevice) GetMetadata() []*MetaKeyValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Status struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminapi_ffe741b8af12354f, []int{13}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ResetAPIKeyParams)(nil), "adminapi.ResetAPIKeyParams")
	proto.RegisterType((*GetAPIKeyParams)(nil), "adminapi.GetAPIKeyParams")
	proto.RegisterType((*APIKeyResponse)(nil), "adminapi.APIKeyResponse")
	proto.RegisterType((*ManifestAddParams)(nil), "adminapi.ManifestAddParams")
	proto.RegisterType((*ManifestAddResponse)(nil), "adminapi.ManifestAddResponse")
	proto.RegisterType((*MetaKeyValue)(nil), "adminapi.MetaKeyValue")
	proto.RegisterType((*ManifestDelParams)(nil), "adminapi.ManifestDelParams")
	proto.RegisterType((*ManifestDelResponse)(nil), "adminapi.ManifestDelResponse")
	proto.RegisterType((*ManifestDelPrefixParams)(nil), "adminapi.ManifestDelPrefixParams")
	proto.RegisterType((*ManifestDelPrefixResponse)(nil), "adminapi.ManifestDelPrefixResponse")
	proto.RegisterType((*ManifestLsDevsParams)(nil), "adminapi.ManifestLsDevsParams")
	proto.RegisterType((*ManifestLsDevsResponse)(nil), "adminapi.ManifestLsDevsResponse")
	proto.RegisterType((*ManifestDevice)(nil), "adminapi.ManifestDevice")
	proto.RegisterType((*Status)(nil), "adminapi.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BTrDBAdminClient is the client API for BTrDBAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BTrDBAdminClient interface {
	// Requires Manifest capability
	ManifestAdd(ctx context.Context, in *ManifestAddParams, opts ...grpc.CallOption) (*ManifestAddResponse, error)
	ManifestDel(ctx context.Context, in *ManifestDelParams, opts ...grpc.CallOption) (*ManifestDelResponse, error)
	ManifestDelPrefix(ctx context.Context, in *ManifestDelPrefixParams, opts ...grpc.CallOption) (*ManifestDelPrefixResponse, error)
	ManifestLsDevs(ctx context.Context, in *ManifestLsDevsParams, opts ...grpc.CallOption) (*ManifestLsDevsResponse, error)
	ResetAPIKey(ctx context.Context, in *ResetAPIKeyParams, opts ...grpc.CallOption) (*APIKeyResponse, error)
	GetAPIKey(ctx context.Context, in *GetAPIKeyParams, opts ...grpc.CallOption) (*APIKeyResponse, error)
}

type bTrDBAdminClient struct {
	cc *grpc.ClientConn
}

func NewBTrDBAdminClient(cc *grpc.ClientConn) BTrDBAdminClient {
	return &bTrDBAdminClient{cc}
}

func (c *bTrDBAdminClient) ManifestAdd(ctx context.Context, in *ManifestAddParams, opts ...grpc.CallOption) (*ManifestAddResponse, error) {
	out := new(ManifestAddResponse)
	err := c.cc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ManifestDel(ctx context.Context, in *ManifestDelParams, opts ...grpc.CallOption) (*ManifestDelResponse, error) {
	out := new(ManifestDelResponse)
	err := c.cc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ManifestDelPrefix(ctx context.Context, in *ManifestDelPrefixParams, opts ...grpc.CallOption) (*ManifestDelPrefixResponse, error) {
	out := new(ManifestDelPrefixResponse)
	err := c.cc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestDelPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ManifestLsDevs(ctx context.Context, in *ManifestLsDevsParams, opts ...grpc.CallOption) (*ManifestLsDevsResponse, error) {
	out := new(ManifestLsDevsResponse)
	err := c.cc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestLsDevs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ResetAPIKey(ctx context.Context, in *ResetAPIKeyParams, opts ...grpc.CallOption) (*APIKeyResponse, error) {
	out := new(APIKeyResponse)
	err := c.cc.Invoke(ctx, "/adminapi.BTrDBAdmin/ResetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) GetAPIKey(ctx context.Context, in *GetAPIKeyParams, opts ...grpc.CallOption) (*APIKeyResponse, error) {
	out := new(APIKeyResponse)
	err := c.cc.Invoke(ctx, "/adminapi.BTrDBAdmin/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BTrDBAdminServer is the server API for BTrDBAdmin service.
type BTrDBAdminServer interface {
	// Requires Manifest capability
	ManifestAdd(context.Context, *ManifestAddParams) (*ManifestAddResponse, error)
	ManifestDel(context.Context, *ManifestDelParams) (*ManifestDelResponse, error)
	ManifestDelPrefix(context.Context, *ManifestDelPrefixParams) (*ManifestDelPrefixResponse, error)
	ManifestLsDevs(context.Context, *ManifestLsDevsParams) (*ManifestLsDevsResponse, error)
	ResetAPIKey(context.Context, *ResetAPIKeyParams) (*APIKeyResponse, error)
	GetAPIKey(context.Context, *GetAPIKeyParams) (*APIKeyResponse, error)
}

func RegisterBTrDBAdminServer(s *grpc.Server, srv BTrDBAdminServer) {
	s.RegisterService(&_BTrDBAdmin_serviceDesc, srv)
}

func _BTrDBAdmin_ManifestAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestAddParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestAdd(ctx, req.(*ManifestAddParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ManifestDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestDelParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestDel(ctx, req.(*ManifestDelParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ManifestDelPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestDelPrefixParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestDelPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestDelPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestDelPrefix(ctx, req.(*ManifestDelPrefixParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ManifestLsDevs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestLsDevsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestLsDevs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestLsDevs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestLsDevs(ctx, req.(*ManifestLsDevsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ResetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetAPIKeyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ResetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ResetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ResetAPIKey(ctx, req.(*ResetAPIKeyParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAPIKeyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).GetAPIKey(ctx, req.(*GetAPIKeyParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _BTrDBAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminapi.BTrDBAdmin",
	HandlerType: (*BTrDBAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ManifestAdd",
			Handler:    _BTrDBAdmin_ManifestAdd_Handler,
		},
		{
			MethodName: "ManifestDel",
			Handler:    _BTrDBAdmin_ManifestDel_Handler,
		},
		{
			MethodName: "ManifestDelPrefix",
			Handler:    _BTrDBAdmin_ManifestDelPrefix_Handler,
		},
		{
			MethodName: "ManifestLsDevs",
			Handler:    _BTrDBAdmin_ManifestLsDevs_Handler,
		},
		{
			MethodName: "ResetAPIKey",
			Handler:    _BTrDBAdmin_ResetAPIKey_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _BTrDBAdmin_GetAPIKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adminapi.proto",
}

func init() { proto.RegisterFile("adminapi.proto", fileDescriptor_adminapi_ffe741b8af12354f) }

var fileDescriptor_adminapi_ffe741b8af12354f = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0x95, 0xb4, 0xcd, 0x93, 0x4c, 0x9e, 0xa4, 0x8d, 0x1b, 0xd2, 0x74, 0xa1, 0x55, 0x30,
	0x08, 0x45, 0x1c, 0x12, 0x29, 0x20, 0x0e, 0x20, 0x21, 0xa5, 0x8a, 0x54, 0xa1, 0x52, 0x54, 0x2d,
	0x88, 0x03, 0xa7, 0x9a, 0x78, 0x1a, 0xad, 0xd8, 0x37, 0xd6, 0x4e, 0x44, 0x0e, 0x5c, 0x38, 0x72,
	0xe5, 0xa3, 0xf1, 0x15, 0xf8, 0x20, 0xc8, 0xf6, 0x66, 0xe3, 0xbc, 0xf0, 0xb2, 0x12, 0x37, 0xdb,
	0x33, 0xf9, 0xff, 0x66, 0xec, 0xff, 0x4e, 0xa0, 0xce, 0x78, 0xe0, 0x85, 0x2c, 0xf6, 0x7a, 0x71,
	0x12, 0xc9, 0x88, 0x94, 0x17, 0x7b, 0xe7, 0xce, 0x24, 0x8a, 0x26, 0x3e, 0xf6, 0x59, 0xec, 0xf5,
	0x59, 0x18, 0x46, 0x92, 0x49, 0x2f, 0x0a, 0x85, 0xc9, 0xa3, 0x87, 0xd0, 0x70, 0x51, 0xa0, 0x1c,
	0x5e, 0xbd, 0xb8, 0xc0, 0xf9, 0x15, 0x4b, 0x58, 0x20, 0x68, 0x03, 0xf6, 0xcf, 0xd7, 0x8e, 0x5e,
	0x41, 0xdd, 0xec, 0x5d, 0x14, 0x71, 0x14, 0x0a, 0x24, 0xf7, 0x61, 0x57, 0x48, 0x26, 0xdb, 0x85,
	0x4e, 0xa1, 0x5b, 0x1d, 0x1c, 0xf4, 0xb2, 0x02, 0x5e, 0x4b, 0x26, 0xa7, 0xc2, 0xd5, 0x51, 0xd2,
	0x82, 0x12, 0x8b, 0xbd, 0x0f, 0x38, 0x6f, 0x17, 0x3b, 0x85, 0x6e, 0xc5, 0x4d, 0x77, 0x74, 0x0c,
	0x8d, 0x4b, 0x16, 0x7a, 0x37, 0x28, 0xe4, 0x90, 0x73, 0x03, 0x21, 0x0e, 0x94, 0x39, 0xce, 0xbc,
	0x31, 0x7a, 0x5c, 0xcb, 0x56, 0xdc, 0x6c, 0x4f, 0x06, 0x50, 0x0e, 0x50, 0x32, 0xce, 0x24, 0x6b,
	0x17, 0x3b, 0x3b, 0xdd, 0xea, 0xa0, 0xb5, 0x44, 0x5e, 0xa2, 0x64, 0x17, 0x38, 0x7f, 0xcb, 0xfc,
	0x29, 0xba, 0x59, 0x1e, 0x7d, 0x06, 0x87, 0x16, 0x24, 0x5f, 0xe5, 0xf4, 0x09, 0xfc, 0x6f, 0xcb,
	0x92, 0x03, 0xd8, 0x51, 0x6d, 0x98, 0xba, 0xd4, 0x92, 0x34, 0x61, 0x6f, 0xa6, 0x42, 0x69, 0x6b,
	0x66, 0x43, 0xfb, 0xcb, 0xce, 0x46, 0xe8, 0xff, 0xb9, 0x33, 0xbb, 0xca, 0x11, 0xfa, 0x39, 0xab,
	0x1c, 0xc2, 0x91, 0x4d, 0x4b, 0xf0, 0xc6, 0xfb, 0x94, 0x32, 0x1f, 0x40, 0x7d, 0xc1, 0x88, 0xf5,
	0x79, 0x4a, 0x5e, 0x3b, 0xa5, 0x0c, 0x8e, 0x37, 0x24, 0x72, 0xbe, 0xf2, 0x29, 0x40, 0x38, 0x0d,
	0x38, 0xfa, 0x28, 0x91, 0xeb, 0xeb, 0xa8, 0xb9, 0xd6, 0x09, 0x7d, 0x0e, 0xcd, 0x05, 0xe2, 0xa5,
	0x18, 0xe1, 0x4c, 0xe4, 0x2c, 0x31, 0x81, 0xd6, 0xea, 0xef, 0x73, 0xd6, 0x37, 0x80, 0xff, 0x8c,
	0xa2, 0x48, 0xbd, 0xd3, 0xb6, 0xbc, 0x93, 0xf5, 0xae, 0x12, 0xdc, 0x45, 0x22, 0xbd, 0x86, 0xfa,
	0x6a, 0xe8, 0x9f, 0xdb, 0xb3, 0x07, 0x25, 0x53, 0x25, 0x21, 0xb0, 0x3b, 0x8e, 0x38, 0x6a, 0xd5,
	0x9a, 0xab, 0xd7, 0xca, 0x6f, 0x81, 0x98, 0xa4, 0xde, 0x52, 0xcb, 0xc1, 0xd7, 0x3d, 0x80, 0xb3,
	0x37, 0xc9, 0xe8, 0x6c, 0xa8, 0x84, 0x09, 0x42, 0xd5, 0x72, 0x37, 0xb9, 0xbd, 0xd9, 0x52, 0xf6,
	0x65, 0x39, 0x27, 0x5b, 0x83, 0x8b, 0x5b, 0xa4, 0xce, 0x97, 0xef, 0x3f, 0xbe, 0x15, 0x9b, 0x74,
	0xbf, 0x3f, 0x7b, 0xdc, 0x0f, 0xd2, 0x04, 0xc6, 0xf9, 0xd3, 0xc2, 0x43, 0x1b, 0x33, 0x42, 0x7f,
	0x1b, 0x26, 0xb3, 0xf9, 0x36, 0x8c, 0x65, 0xe9, 0xed, 0x18, 0x8e, 0xbe, 0xc2, 0x7c, 0x5e, 0xfd,
	0x6c, 0xf4, 0xbb, 0x93, 0xbb, 0xdb, 0x61, 0x96, 0xcb, 0x9d, 0x7b, 0xbf, 0x49, 0xc9, 0xc0, 0x1d,
	0x0d, 0x76, 0xe8, 0xad, 0x35, 0xb0, 0xb1, 0x97, 0xc2, 0x7f, 0x5c, 0xbe, 0xb6, 0x71, 0x18, 0x39,
	0xdd, 0x14, 0xb6, 0xbd, 0xeb, 0x74, 0x7e, 0x15, 0xcf, 0xa8, 0x27, 0x9a, 0x7a, 0x44, 0x89, 0x4d,
	0xf5, 0x05, 0xc7, 0x99, 0x50, 0xc8, 0x6b, 0xa8, 0x5a, 0xa3, 0xd7, 0xbe, 0xd8, 0x8d, 0x89, 0xec,
	0x58, 0x7e, 0x5d, 0x1d, 0xc3, 0xab, 0x77, 0x9a, 0xa8, 0x1f, 0x9a, 0x09, 0xab, 0x08, 0xef, 0xa0,
	0x92, 0xcd, 0x71, 0x72, 0xbc, 0x94, 0x38, 0xff, 0x6b, 0xf5, 0xb6, 0x56, 0x27, 0xb4, 0xa6, 0xd4,
	0x27, 0x96, 0xf6, 0xfb, 0x92, 0xfe, 0xff, 0x78, 0xf4, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xa7,
	0x81, 0xc4, 0x79, 0x06, 0x00, 0x00,
}