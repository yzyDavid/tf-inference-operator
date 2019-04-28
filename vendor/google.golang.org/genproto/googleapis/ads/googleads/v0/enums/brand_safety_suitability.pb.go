// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/brand_safety_suitability.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 3-Tier brand safety suitability control.
type BrandSafetySuitabilityEnum_BrandSafetySuitability int32

const (
	// Not specified.
	BrandSafetySuitabilityEnum_UNSPECIFIED BrandSafetySuitabilityEnum_BrandSafetySuitability = 0
	// Used for return value only. Represents value unknown in this version.
	BrandSafetySuitabilityEnum_UNKNOWN BrandSafetySuitabilityEnum_BrandSafetySuitability = 1
	// This option lets you show ads across all inventory on YouTube and video
	// partners that meet our standards for monetization. This option may be an
	// appropriate choice for brands that want maximum access to the full
	// breadth of videos eligible for ads, including, for example, videos that
	// have strong profanity in the context of comedy or a documentary, or
	// excessive violence as featured in video games.
	BrandSafetySuitabilityEnum_EXPANDED_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 2
	// This option lets you show ads across a wide range of content that's
	// appropriate for most brands, such as popular music videos, documentaries,
	// and movie trailers. The content you can show ads on is based on YouTube's
	// advertiser-friendly content guidelines that take into account, for
	// example, the strength or frequency of profanity, or the appropriateness
	// of subject matter like sensitive events. Ads won't show, for example, on
	// content with repeated strong profanity, strong sexual content, or graphic
	// violence.
	BrandSafetySuitabilityEnum_STANDARD_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 3
	// This option lets you show ads on a reduced range of content that's
	// appropriate for brands with particularly strict guidelines around
	// inappropriate language and sexual suggestiveness; above and beyond what
	// YouTube's advertiser-friendly content guidelines address. The videos
	// accessible in this sensitive category meet heightened requirements,
	// especially for inappropriate language and sexual suggestiveness. For
	// example, your ads will be excluded from showing on some of YouTube's most
	// popular music videos and other pop culture content across YouTube and
	// Google video partners.
	BrandSafetySuitabilityEnum_LIMITED_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 4
)

var BrandSafetySuitabilityEnum_BrandSafetySuitability_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EXPANDED_INVENTORY",
	3: "STANDARD_INVENTORY",
	4: "LIMITED_INVENTORY",
}
var BrandSafetySuitabilityEnum_BrandSafetySuitability_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"EXPANDED_INVENTORY": 2,
	"STANDARD_INVENTORY": 3,
	"LIMITED_INVENTORY":  4,
}

func (x BrandSafetySuitabilityEnum_BrandSafetySuitability) String() string {
	return proto.EnumName(BrandSafetySuitabilityEnum_BrandSafetySuitability_name, int32(x))
}
func (BrandSafetySuitabilityEnum_BrandSafetySuitability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_brand_safety_suitability_9619f3de7a720efd, []int{0, 0}
}

// Container for enum with 3-Tier brand safety suitability control.
type BrandSafetySuitabilityEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrandSafetySuitabilityEnum) Reset()         { *m = BrandSafetySuitabilityEnum{} }
func (m *BrandSafetySuitabilityEnum) String() string { return proto.CompactTextString(m) }
func (*BrandSafetySuitabilityEnum) ProtoMessage()    {}
func (*BrandSafetySuitabilityEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_brand_safety_suitability_9619f3de7a720efd, []int{0}
}
func (m *BrandSafetySuitabilityEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrandSafetySuitabilityEnum.Unmarshal(m, b)
}
func (m *BrandSafetySuitabilityEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrandSafetySuitabilityEnum.Marshal(b, m, deterministic)
}
func (dst *BrandSafetySuitabilityEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrandSafetySuitabilityEnum.Merge(dst, src)
}
func (m *BrandSafetySuitabilityEnum) XXX_Size() int {
	return xxx_messageInfo_BrandSafetySuitabilityEnum.Size(m)
}
func (m *BrandSafetySuitabilityEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BrandSafetySuitabilityEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BrandSafetySuitabilityEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BrandSafetySuitabilityEnum)(nil), "google.ads.googleads.v0.enums.BrandSafetySuitabilityEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.BrandSafetySuitabilityEnum_BrandSafetySuitability", BrandSafetySuitabilityEnum_BrandSafetySuitability_name, BrandSafetySuitabilityEnum_BrandSafetySuitability_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/brand_safety_suitability.proto", fileDescriptor_brand_safety_suitability_9619f3de7a720efd)
}

var fileDescriptor_brand_safety_suitability_9619f3de7a720efd = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xa4, 0xa2, 0xc4, 0xbc, 0x94, 0xf8, 0xe2, 0xc4, 0xb4, 0xd4, 0x92, 0xca,
	0xf8, 0xe2, 0xd2, 0xcc, 0x92, 0xc4, 0xa4, 0xcc, 0x9c, 0xcc, 0x92, 0x4a, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x59, 0x88, 0x16, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x6e, 0xbd, 0x32, 0x03,
	0x3d, 0xb0, 0x6e, 0xa5, 0xd9, 0x8c, 0x5c, 0x52, 0x4e, 0x20, 0x13, 0x82, 0xc1, 0x06, 0x04, 0x23,
	0xf4, 0xbb, 0xe6, 0x95, 0xe6, 0x2a, 0xd5, 0x72, 0x89, 0x61, 0x97, 0x15, 0xe2, 0xe7, 0xe2, 0x0e,
	0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0xe2, 0xe6, 0x62,
	0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x14, 0x12, 0xe3, 0x12, 0x72, 0x8d, 0x08,
	0x70, 0xf4, 0x73, 0x71, 0x75, 0x89, 0xf7, 0xf4, 0x0b, 0x73, 0xf5, 0x0b, 0xf1, 0x0f, 0x8a, 0x14,
	0x60, 0x02, 0x89, 0x07, 0x87, 0x38, 0xfa, 0xb9, 0x38, 0x06, 0x21, 0x8b, 0x33, 0x0b, 0x89, 0x72,
	0x09, 0xfa, 0x78, 0xfa, 0x7a, 0x86, 0xa0, 0x28, 0x67, 0x71, 0xfa, 0xc0, 0xc8, 0xa5, 0x98, 0x9c,
	0x9f, 0xab, 0x87, 0xd7, 0x0f, 0x4e, 0xd2, 0xd8, 0x9d, 0x18, 0x00, 0xf2, 0x7f, 0x00, 0x63, 0x94,
	0x13, 0x54, 0x77, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a,
	0x1e, 0x38, 0x74, 0x60, 0xe1, 0x59, 0x90, 0x59, 0x8c, 0x23, 0x78, 0xad, 0xc1, 0xe4, 0x22, 0x26,
	0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0xb2, 0xee, 0x10, 0xa3, 0x1c, 0x53, 0x8a, 0xf5, 0x20, 0x4c,
	0x10, 0x2b, 0xcc, 0x40, 0x0f, 0x14, 0x58, 0xc5, 0xa7, 0x60, 0xf2, 0x31, 0x8e, 0x29, 0xc5, 0x31,
	0x70, 0xf9, 0x98, 0x30, 0x83, 0x18, 0xb0, 0xfc, 0x2b, 0x26, 0x45, 0x88, 0xa0, 0x95, 0x95, 0x63,
	0x4a, 0xb1, 0x95, 0x15, 0x5c, 0x85, 0x95, 0x55, 0x98, 0x81, 0x95, 0x15, 0x58, 0x4d, 0x12, 0x1b,
	0xd8, 0x61, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x3c, 0x85, 0xe6, 0xf6, 0x01, 0x00,
	0x00,
}
