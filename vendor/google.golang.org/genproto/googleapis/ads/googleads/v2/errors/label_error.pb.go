// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/label_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible label errors.
type LabelErrorEnum_LabelError int32

const (
	// Enum unspecified.
	LabelErrorEnum_UNSPECIFIED LabelErrorEnum_LabelError = 0
	// The received error code is not known in this version.
	LabelErrorEnum_UNKNOWN LabelErrorEnum_LabelError = 1
	// An inactive label cannot be applied.
	LabelErrorEnum_CANNOT_APPLY_INACTIVE_LABEL LabelErrorEnum_LabelError = 2
	// A label cannot be applied to a disabled ad group criterion.
	LabelErrorEnum_CANNOT_APPLY_LABEL_TO_DISABLED_AD_GROUP_CRITERION LabelErrorEnum_LabelError = 3
	// A label cannot be applied to a negative ad group criterion.
	LabelErrorEnum_CANNOT_APPLY_LABEL_TO_NEGATIVE_AD_GROUP_CRITERION LabelErrorEnum_LabelError = 4
	// Cannot apply more than 50 labels per resource.
	LabelErrorEnum_EXCEEDED_LABEL_LIMIT_PER_TYPE LabelErrorEnum_LabelError = 5
	// Labels from a manager account cannot be applied to campaign, ad group,
	// ad group ad, or ad group criterion resources.
	LabelErrorEnum_INVALID_RESOURCE_FOR_MANAGER_LABEL LabelErrorEnum_LabelError = 6
	// Label names must be unique.
	LabelErrorEnum_DUPLICATE_NAME LabelErrorEnum_LabelError = 7
	// Label names cannot be empty.
	LabelErrorEnum_INVALID_LABEL_NAME LabelErrorEnum_LabelError = 8
	// Labels cannot be applied to a draft.
	LabelErrorEnum_CANNOT_ATTACH_LABEL_TO_DRAFT LabelErrorEnum_LabelError = 9
	// Labels not from a manager account cannot be applied to the customer
	// resource.
	LabelErrorEnum_CANNOT_ATTACH_NON_MANAGER_LABEL_TO_CUSTOMER LabelErrorEnum_LabelError = 10
)

var LabelErrorEnum_LabelError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "CANNOT_APPLY_INACTIVE_LABEL",
	3:  "CANNOT_APPLY_LABEL_TO_DISABLED_AD_GROUP_CRITERION",
	4:  "CANNOT_APPLY_LABEL_TO_NEGATIVE_AD_GROUP_CRITERION",
	5:  "EXCEEDED_LABEL_LIMIT_PER_TYPE",
	6:  "INVALID_RESOURCE_FOR_MANAGER_LABEL",
	7:  "DUPLICATE_NAME",
	8:  "INVALID_LABEL_NAME",
	9:  "CANNOT_ATTACH_LABEL_TO_DRAFT",
	10: "CANNOT_ATTACH_NON_MANAGER_LABEL_TO_CUSTOMER",
}

var LabelErrorEnum_LabelError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"CANNOT_APPLY_INACTIVE_LABEL": 2,
	"CANNOT_APPLY_LABEL_TO_DISABLED_AD_GROUP_CRITERION": 3,
	"CANNOT_APPLY_LABEL_TO_NEGATIVE_AD_GROUP_CRITERION": 4,
	"EXCEEDED_LABEL_LIMIT_PER_TYPE":                     5,
	"INVALID_RESOURCE_FOR_MANAGER_LABEL":                6,
	"DUPLICATE_NAME":                                    7,
	"INVALID_LABEL_NAME":                                8,
	"CANNOT_ATTACH_LABEL_TO_DRAFT":                      9,
	"CANNOT_ATTACH_NON_MANAGER_LABEL_TO_CUSTOMER":       10,
}

func (x LabelErrorEnum_LabelError) String() string {
	return proto.EnumName(LabelErrorEnum_LabelError_name, int32(x))
}

func (LabelErrorEnum_LabelError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51e059200f18f70d, []int{0, 0}
}

// Container for enum describing possible label errors.
type LabelErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelErrorEnum) Reset()         { *m = LabelErrorEnum{} }
func (m *LabelErrorEnum) String() string { return proto.CompactTextString(m) }
func (*LabelErrorEnum) ProtoMessage()    {}
func (*LabelErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e059200f18f70d, []int{0}
}

func (m *LabelErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelErrorEnum.Unmarshal(m, b)
}
func (m *LabelErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelErrorEnum.Marshal(b, m, deterministic)
}
func (m *LabelErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelErrorEnum.Merge(m, src)
}
func (m *LabelErrorEnum) XXX_Size() int {
	return xxx_messageInfo_LabelErrorEnum.Size(m)
}
func (m *LabelErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LabelErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.LabelErrorEnum_LabelError", LabelErrorEnum_LabelError_name, LabelErrorEnum_LabelError_value)
	proto.RegisterType((*LabelErrorEnum)(nil), "google.ads.googleads.v2.errors.LabelErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/label_error.proto", fileDescriptor_51e059200f18f70d)
}

var fileDescriptor_51e059200f18f70d = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x40, 0x49, 0x02, 0x2d, 0x6c, 0xa5, 0xd6, 0xda, 0x03, 0x07, 0x28, 0x05, 0x72, 0xe0, 0x82,
	0x64, 0x43, 0x10, 0x17, 0x73, 0x9a, 0xd8, 0x13, 0xb3, 0xc2, 0x59, 0x5b, 0x9b, 0xb5, 0xa1, 0x28,
	0xd2, 0xca, 0xc5, 0x91, 0x15, 0x29, 0xf5, 0x46, 0x76, 0xe8, 0x0f, 0xf0, 0x0f, 0x7c, 0x00, 0x47,
	0x3e, 0x85, 0x4f, 0xe9, 0x57, 0x20, 0x7b, 0xe3, 0x94, 0x48, 0xb4, 0x27, 0x8f, 0x47, 0xef, 0xcd,
	0xcc, 0x6a, 0x86, 0xbc, 0x29, 0xb4, 0x2e, 0x56, 0x0b, 0x27, 0xcb, 0x6b, 0xc7, 0x84, 0x4d, 0x74,
	0x35, 0x72, 0x16, 0x55, 0xa5, 0xab, 0xda, 0x59, 0x65, 0x17, 0x8b, 0x95, 0x6a, 0x7f, 0xec, 0x75,
	0xa5, 0x37, 0x9a, 0x9e, 0x19, 0xcc, 0xce, 0xf2, 0xda, 0xde, 0x19, 0xf6, 0xd5, 0xc8, 0x36, 0xc6,
	0x93, 0xd3, 0xae, 0xe2, 0x7a, 0xe9, 0x64, 0x65, 0xa9, 0x37, 0xd9, 0x66, 0xa9, 0xcb, 0xda, 0xd8,
	0xc3, 0x9f, 0x03, 0x72, 0x1c, 0x36, 0x35, 0xb1, 0xa1, 0xb1, 0xfc, 0x7e, 0x39, 0xfc, 0x31, 0x20,
	0xe4, 0x26, 0x45, 0x4f, 0xc8, 0x51, 0xc2, 0x67, 0x31, 0x7a, 0x6c, 0xc2, 0xd0, 0xb7, 0xee, 0xd1,
	0x23, 0x72, 0x98, 0xf0, 0x4f, 0x3c, 0xfa, 0xcc, 0xad, 0x1e, 0x7d, 0x4e, 0x9e, 0x7a, 0xc0, 0x79,
	0x24, 0x15, 0xc4, 0x71, 0x78, 0xae, 0x18, 0x07, 0x4f, 0xb2, 0x14, 0x55, 0x08, 0x63, 0x0c, 0xad,
	0x3e, 0x7d, 0x4f, 0xde, 0xee, 0x01, 0x6d, 0x5e, 0xc9, 0x48, 0xf9, 0x6c, 0x06, 0xe3, 0x10, 0x7d,
	0x05, 0xbe, 0x0a, 0x44, 0x94, 0xc4, 0xca, 0x13, 0x4c, 0xa2, 0x60, 0x11, 0xb7, 0x06, 0xb7, 0x6b,
	0x1c, 0x03, 0x68, 0x1b, 0xfc, 0x47, 0xbb, 0x4f, 0x5f, 0x92, 0x67, 0xf8, 0xc5, 0x43, 0xf4, 0xd1,
	0xdf, 0x2a, 0x21, 0x9b, 0x32, 0xa9, 0x62, 0x14, 0x4a, 0x9e, 0xc7, 0x68, 0x3d, 0xa0, 0xaf, 0xc8,
	0x90, 0xf1, 0x14, 0x42, 0xe6, 0x2b, 0x81, 0xb3, 0x28, 0x11, 0x1e, 0xaa, 0x49, 0x24, 0xd4, 0x14,
	0x38, 0x04, 0x28, 0xb6, 0x83, 0x1f, 0x50, 0x4a, 0x8e, 0xfd, 0x24, 0x0e, 0x99, 0x07, 0x12, 0x15,
	0x87, 0x29, 0x5a, 0x87, 0xf4, 0x31, 0xa1, 0x9d, 0x6b, 0xaa, 0xb7, 0xf9, 0x87, 0xf4, 0x05, 0x39,
	0xed, 0xa6, 0x95, 0x12, 0xbc, 0x8f, 0xff, 0xbc, 0x52, 0xc0, 0x44, 0x5a, 0x8f, 0xa8, 0x43, 0x5e,
	0xef, 0x13, 0x3c, 0xe2, 0xfb, 0x2d, 0x1b, 0xda, 0x4b, 0x66, 0x32, 0x9a, 0xa2, 0xb0, 0xc8, 0xf8,
	0xba, 0x47, 0x86, 0xdf, 0xf4, 0xa5, 0x7d, 0xf7, 0x76, 0xc7, 0x27, 0x37, 0x9b, 0x8a, 0x9b, 0x85,
	0xc6, 0xbd, 0xaf, 0xfe, 0x56, 0x29, 0xf4, 0x2a, 0x2b, 0x0b, 0x5b, 0x57, 0x85, 0x53, 0x2c, 0xca,
	0x76, 0xdd, 0xdd, 0x49, 0xad, 0x97, 0xf5, 0x6d, 0x17, 0xf6, 0xc1, 0x7c, 0x7e, 0xf5, 0x07, 0x01,
	0xc0, 0xef, 0xfe, 0x59, 0x60, 0x8a, 0x41, 0x5e, 0xdb, 0x26, 0x6c, 0xa2, 0x74, 0x64, 0xb7, 0x2d,
	0xeb, 0x3f, 0x1d, 0x30, 0x87, 0xbc, 0x9e, 0xef, 0x80, 0x79, 0x3a, 0x9a, 0x1b, 0xe0, 0xba, 0x3f,
	0x34, 0x59, 0xd7, 0x85, 0xbc, 0x76, 0xdd, 0x1d, 0xe2, 0xba, 0xe9, 0xc8, 0x75, 0x0d, 0x74, 0x71,
	0xd0, 0x4e, 0xf7, 0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xa2, 0x2c, 0x32, 0xfe, 0x02,
	0x00, 0x00,
}