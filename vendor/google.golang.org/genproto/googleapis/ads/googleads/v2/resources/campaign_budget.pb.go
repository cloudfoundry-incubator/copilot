// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_budget.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A campaign budget.
type CampaignBudget struct {
	// The resource name of the campaign budget.
	// Campaign budget resource names have the form:
	//
	// `customers/{customer_id}/campaignBudgets/{budget_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the campaign budget.
	//
	// A campaign budget is created using the CampaignBudgetService create
	// operation and is assigned a budget ID. A budget ID can be shared across
	// different campaigns; the system will then allocate the campaign budget
	// among different campaigns to get optimum results.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the campaign budget.
	//
	// When creating a campaign budget through CampaignBudgetService, every
	// explicitly shared campaign budget must have a non-null, non-empty name.
	// Campaign budgets that are not explicitly shared derive their name from the
	// attached campaign's name.
	//
	// The length of this string must be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The amount of the budget, in the local currency for the account.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit. Monthly spend is capped at 30.4 times this amount.
	AmountMicros *wrappers.Int64Value `protobuf:"bytes,5,opt,name=amount_micros,json=amountMicros,proto3" json:"amount_micros,omitempty"`
	// The lifetime amount of the budget, in the local currency for the account.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,10,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// The status of this campaign budget. This field is read-only.
	Status enums.BudgetStatusEnum_BudgetStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.BudgetStatusEnum_BudgetStatus" json:"status,omitempty"`
	// The delivery method that determines the rate at which the campaign budget
	// is spent.
	//
	// Defaults to STANDARD if unspecified in a create operation.
	DeliveryMethod enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod `protobuf:"varint,7,opt,name=delivery_method,json=deliveryMethod,proto3,enum=google.ads.googleads.v2.enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod" json:"delivery_method,omitempty"`
	// Specifies whether the budget is explicitly shared. Defaults to true if
	// unspecified in a create operation.
	//
	// If true, the budget was created with the purpose of sharing
	// across one or more campaigns.
	//
	// If false, the budget was created with the intention of only being used
	// with a single campaign. The budget's name and status will stay in sync
	// with the campaign's name and status. Attempting to share the budget with a
	// second campaign will result in an error.
	//
	// A non-shared budget can become an explicitly shared. The same operation
	// must also assign the budget a name.
	//
	// A shared campaign budget can never become non-shared.
	ExplicitlyShared *wrappers.BoolValue `protobuf:"bytes,8,opt,name=explicitly_shared,json=explicitlyShared,proto3" json:"explicitly_shared,omitempty"`
	// The number of campaigns actively using the budget.
	//
	// This field is read-only.
	ReferenceCount *wrappers.Int64Value `protobuf:"bytes,9,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	// Indicates whether there is a recommended budget for this campaign budget.
	//
	// This field is read-only.
	HasRecommendedBudget *wrappers.BoolValue `protobuf:"bytes,11,opt,name=has_recommended_budget,json=hasRecommendedBudget,proto3" json:"has_recommended_budget,omitempty"`
	// The recommended budget amount. If no recommendation is available, this will
	// be set to the budget amount.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	//
	// This field is read-only.
	RecommendedBudgetAmountMicros *wrappers.Int64Value `protobuf:"bytes,12,opt,name=recommended_budget_amount_micros,json=recommendedBudgetAmountMicros,proto3" json:"recommended_budget_amount_micros,omitempty"`
	// Period over which to spend the budget. Defaults to DAILY if not specified.
	Period enums.BudgetPeriodEnum_BudgetPeriod `protobuf:"varint,13,opt,name=period,proto3,enum=google.ads.googleads.v2.enums.BudgetPeriodEnum_BudgetPeriod" json:"period,omitempty"`
	// The estimated change in weekly clicks if the recommended budget is applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyClicks *wrappers.Int64Value `protobuf:"bytes,14,opt,name=recommended_budget_estimated_change_weekly_clicks,json=recommendedBudgetEstimatedChangeWeeklyClicks,proto3" json:"recommended_budget_estimated_change_weekly_clicks,omitempty"`
	// The estimated change in weekly cost in micros if the recommended budget is
	// applied. One million is equivalent to one currency unit.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyCostMicros *wrappers.Int64Value `protobuf:"bytes,15,opt,name=recommended_budget_estimated_change_weekly_cost_micros,json=recommendedBudgetEstimatedChangeWeeklyCostMicros,proto3" json:"recommended_budget_estimated_change_weekly_cost_micros,omitempty"`
	// The estimated change in weekly interactions if the recommended budget is
	// applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyInteractions *wrappers.Int64Value `protobuf:"bytes,16,opt,name=recommended_budget_estimated_change_weekly_interactions,json=recommendedBudgetEstimatedChangeWeeklyInteractions,proto3" json:"recommended_budget_estimated_change_weekly_interactions,omitempty"`
	// The estimated change in weekly views if the recommended budget is applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyViews *wrappers.Int64Value `protobuf:"bytes,17,opt,name=recommended_budget_estimated_change_weekly_views,json=recommendedBudgetEstimatedChangeWeeklyViews,proto3" json:"recommended_budget_estimated_change_weekly_views,omitempty"`
	// The type of the campaign budget.
	Type                 enums.BudgetTypeEnum_BudgetType `protobuf:"varint,18,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.BudgetTypeEnum_BudgetType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CampaignBudget) Reset()         { *m = CampaignBudget{} }
func (m *CampaignBudget) String() string { return proto.CompactTextString(m) }
func (*CampaignBudget) ProtoMessage()    {}
func (*CampaignBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2bb05963aea82a6, []int{0}
}

func (m *CampaignBudget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBudget.Unmarshal(m, b)
}
func (m *CampaignBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBudget.Marshal(b, m, deterministic)
}
func (m *CampaignBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBudget.Merge(m, src)
}
func (m *CampaignBudget) XXX_Size() int {
	return xxx_messageInfo_CampaignBudget.Size(m)
}
func (m *CampaignBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBudget proto.InternalMessageInfo

func (m *CampaignBudget) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignBudget) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CampaignBudget) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CampaignBudget) GetAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.AmountMicros
	}
	return nil
}

func (m *CampaignBudget) GetTotalAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TotalAmountMicros
	}
	return nil
}

func (m *CampaignBudget) GetStatus() enums.BudgetStatusEnum_BudgetStatus {
	if m != nil {
		return m.Status
	}
	return enums.BudgetStatusEnum_UNSPECIFIED
}

func (m *CampaignBudget) GetDeliveryMethod() enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod {
	if m != nil {
		return m.DeliveryMethod
	}
	return enums.BudgetDeliveryMethodEnum_UNSPECIFIED
}

func (m *CampaignBudget) GetExplicitlyShared() *wrappers.BoolValue {
	if m != nil {
		return m.ExplicitlyShared
	}
	return nil
}

func (m *CampaignBudget) GetReferenceCount() *wrappers.Int64Value {
	if m != nil {
		return m.ReferenceCount
	}
	return nil
}

func (m *CampaignBudget) GetHasRecommendedBudget() *wrappers.BoolValue {
	if m != nil {
		return m.HasRecommendedBudget
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetAmountMicros
	}
	return nil
}

func (m *CampaignBudget) GetPeriod() enums.BudgetPeriodEnum_BudgetPeriod {
	if m != nil {
		return m.Period
	}
	return enums.BudgetPeriodEnum_UNSPECIFIED
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyClicks() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyClicks
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyCostMicros
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyInteractions() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyInteractions
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyViews() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyViews
	}
	return nil
}

func (m *CampaignBudget) GetType() enums.BudgetTypeEnum_BudgetType {
	if m != nil {
		return m.Type
	}
	return enums.BudgetTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignBudget)(nil), "google.ads.googleads.v2.resources.CampaignBudget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_budget.proto", fileDescriptor_b2bb05963aea82a6)
}

var fileDescriptor_b2bb05963aea82a6 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x18, 0x86, 0xe5, 0x6c, 0x29, 0xec, 0x6c, 0x9a, 0x6e, 0xbd, 0x08, 0x59, 0x65, 0x41, 0x59, 0xd0,
	0x4a, 0x95, 0x16, 0xd9, 0x5d, 0x83, 0x5a, 0x64, 0x38, 0x20, 0x49, 0xab, 0xaa, 0xd0, 0xa2, 0xc8,
	0xad, 0x82, 0x84, 0x22, 0x59, 0x53, 0xcf, 0x57, 0x67, 0x54, 0x7b, 0xc6, 0x9a, 0x19, 0xa7, 0xe4,
	0x0c, 0x21, 0x8e, 0x90, 0x90, 0xb8, 0x06, 0x0e, 0xb9, 0x14, 0x2e, 0x85, 0xab, 0x40, 0x1e, 0xff,
	0x34, 0x3f, 0x14, 0x27, 0x67, 0xf3, 0xf3, 0xbd, 0x4f, 0xde, 0x79, 0xe7, 0x73, 0x06, 0x1d, 0x47,
	0x9c, 0x47, 0x31, 0x38, 0x98, 0x48, 0xa7, 0x18, 0xe6, 0xa3, 0xa9, 0xeb, 0x08, 0x90, 0x3c, 0x13,
	0x21, 0x48, 0x27, 0xc4, 0x49, 0x8a, 0x69, 0xc4, 0x82, 0x9b, 0x8c, 0x44, 0xa0, 0xec, 0x54, 0x70,
	0xc5, 0xcd, 0x57, 0x45, 0xb5, 0x8d, 0x89, 0xb4, 0x6b, 0xa1, 0x3d, 0x75, 0xed, 0x5a, 0xb8, 0xef,
	0x3d, 0xc6, 0x06, 0x96, 0x25, 0xd2, 0x29, 0x70, 0x01, 0x81, 0x98, 0x4e, 0x41, 0xcc, 0x82, 0x04,
	0xd4, 0x84, 0x93, 0x02, 0xbf, 0xff, 0x76, 0x2d, 0x6d, 0x0a, 0x82, 0x6e, 0x28, 0x91, 0x0a, 0xab,
	0x4c, 0x96, 0x12, 0x67, 0x2d, 0x89, 0x9a, 0xa5, 0x50, 0x0a, 0x3e, 0x2e, 0x05, 0x7a, 0x76, 0x93,
	0xdd, 0x3a, 0xf7, 0x02, 0xa7, 0x29, 0x88, 0x0a, 0xf8, 0xb2, 0x02, 0xa6, 0xd4, 0xc1, 0x8c, 0x71,
	0x85, 0x15, 0xe5, 0xac, 0xdc, 0xfd, 0xe4, 0xb7, 0x36, 0xea, 0x0c, 0xca, 0x34, 0xfb, 0x9a, 0x6d,
	0x7e, 0x8a, 0x76, 0xaa, 0xc0, 0x02, 0x86, 0x13, 0xb0, 0x8c, 0xae, 0x71, 0xf0, 0xd4, 0x6f, 0x57,
	0x8b, 0xdf, 0xe3, 0x04, 0xcc, 0x37, 0xa8, 0x45, 0x89, 0xf5, 0xa4, 0x6b, 0x1c, 0x3c, 0x73, 0x3f,
	0x2c, 0xd3, 0xb6, 0x2b, 0x0b, 0xf6, 0x39, 0x53, 0x47, 0x5f, 0x8c, 0x70, 0x9c, 0x81, 0xdf, 0xa2,
	0xc4, 0x3c, 0x44, 0x5b, 0x1a, 0xb4, 0xa5, 0xcb, 0x5f, 0xae, 0x94, 0x5f, 0x29, 0x41, 0x59, 0x54,
	0xd4, 0xeb, 0x4a, 0xf3, 0x1b, 0xb4, 0x83, 0x13, 0x9e, 0x31, 0x15, 0x24, 0x34, 0x14, 0x5c, 0x5a,
	0xef, 0x34, 0xff, 0x52, 0xbb, 0x50, 0x5c, 0x6a, 0x81, 0xf9, 0x1d, 0x7a, 0xa1, 0xb8, 0xc2, 0x71,
	0xb0, 0xc8, 0x41, 0xcd, 0x9c, 0x3d, 0xad, 0xeb, 0xcd, 0xc3, 0xae, 0xd1, 0x76, 0x71, 0x49, 0xd6,
	0x76, 0xd7, 0x38, 0xe8, 0xb8, 0x5f, 0xdb, 0x8f, 0xb5, 0x9a, 0xbe, 0x25, 0xbb, 0x48, 0xf2, 0x4a,
	0x4b, 0x4e, 0x59, 0x96, 0x2c, 0x2c, 0xf8, 0x25, 0xcb, 0xcc, 0xd0, 0xee, 0x52, 0xa7, 0x59, 0xef,
	0x6a, 0xfc, 0xc5, 0x5a, 0xf8, 0x93, 0x52, 0x7b, 0xa9, 0xa5, 0x73, 0x3f, 0xb3, 0xb8, 0xe1, 0x77,
	0xc8, 0xc2, 0xdc, 0x3c, 0x43, 0x7b, 0xf0, 0x53, 0x1a, 0xd3, 0x90, 0xaa, 0x78, 0x16, 0xc8, 0x09,
	0x16, 0x40, 0xac, 0xf7, 0x74, 0x2e, 0xfb, 0x2b, 0xb9, 0xf4, 0x39, 0x8f, 0x8b, 0x58, 0x9e, 0x3f,
	0x88, 0xae, 0xb4, 0xc6, 0x3c, 0x41, 0xbb, 0x02, 0x6e, 0x41, 0x00, 0x0b, 0x21, 0x08, 0xf3, 0xb8,
	0xac, 0xa7, 0xcd, 0xf1, 0x76, 0x6a, 0xcd, 0x20, 0x97, 0x98, 0x43, 0xf4, 0xc1, 0x04, 0xcb, 0x40,
	0x40, 0xc8, 0x93, 0x04, 0x18, 0x01, 0x52, 0x7e, 0xd5, 0xd6, 0xb3, 0x46, 0x4f, 0xef, 0x4f, 0xb0,
	0xf4, 0x1f, 0x84, 0x65, 0x03, 0x13, 0xd4, 0x5d, 0xa5, 0x2d, 0xf5, 0x41, 0xbb, 0xd9, 0xe8, 0x47,
	0x62, 0x99, 0xbc, 0xdc, 0x13, 0xc5, 0xb7, 0x6e, 0xed, 0x6c, 0xd0, 0x13, 0x43, 0x2d, 0x99, 0xbb,
	0xac, 0x62, 0xc1, 0x2f, 0x59, 0xe6, 0xaf, 0x06, 0x7a, 0xfb, 0x1f, 0xe6, 0x41, 0x2a, 0x9a, 0x60,
	0x05, 0x24, 0x08, 0x27, 0x98, 0x45, 0x10, 0xdc, 0x03, 0xdc, 0xc5, 0xb3, 0x20, 0x8c, 0x69, 0x78,
	0x27, 0xad, 0x4e, 0xf3, 0x69, 0x3e, 0x5b, 0x39, 0xcd, 0x69, 0xc5, 0x1c, 0x68, 0xe4, 0x0f, 0x9a,
	0x38, 0xd0, 0x40, 0xf3, 0x77, 0x03, 0x1d, 0x6d, 0x62, 0x83, 0xcb, 0x3a, 0xd9, 0xdd, 0x66, 0x2f,
	0x87, 0x6b, 0x7a, 0xe1, 0xb2, 0x0a, 0xfb, 0x0f, 0x03, 0x1d, 0x6f, 0xe0, 0x87, 0x32, 0x05, 0x02,
	0x87, 0xfa, 0x8f, 0xce, 0x7a, 0xde, 0x6c, 0xc8, 0x5d, 0xcf, 0xd0, 0xf9, 0x1c, 0xd6, 0xfc, 0xc5,
	0x40, 0x87, 0x1b, 0x58, 0x9a, 0x52, 0xb8, 0x97, 0xd6, 0x5e, 0xb3, 0x97, 0x37, 0xeb, 0x79, 0x19,
	0xe5, 0x3c, 0xf3, 0x02, 0x6d, 0xe5, 0x4f, 0x81, 0x65, 0xea, 0x16, 0xfc, 0x72, 0xad, 0x16, 0xbc,
	0x9e, 0xa5, 0x30, 0xd7, 0x80, 0xf9, 0xd4, 0xd7, 0x94, 0xfe, 0xcf, 0x2d, 0xf4, 0x3a, 0xe4, 0x89,
	0xdd, 0xf8, 0x8e, 0xf6, 0x5f, 0x2c, 0xbe, 0x19, 0xc3, 0xfc, 0x1c, 0x43, 0xe3, 0xc7, 0x6f, 0x4b,
	0x65, 0xc4, 0x63, 0xcc, 0x22, 0x9b, 0x8b, 0xc8, 0x89, 0x80, 0xe9, 0x53, 0x56, 0x8f, 0x59, 0x4a,
	0xe5, 0xff, 0xbc, 0xec, 0x5f, 0xd5, 0xa3, 0x3f, 0x5b, 0x4f, 0xce, 0x7a, 0xbd, 0xbf, 0x5a, 0xaf,
	0xce, 0x0a, 0x64, 0x8f, 0x48, 0xbb, 0x18, 0xe6, 0xa3, 0x91, 0x6b, 0xfb, 0x55, 0xe5, 0xdf, 0x55,
	0xcd, 0xb8, 0x47, 0xe4, 0xb8, 0xae, 0x19, 0x8f, 0xdc, 0x71, 0x5d, 0xf3, 0x4f, 0xeb, 0x75, 0xb1,
	0xe1, 0x79, 0x3d, 0x22, 0x3d, 0xaf, 0xae, 0xf2, 0xbc, 0x91, 0xeb, 0x79, 0x75, 0xdd, 0xcd, 0xb6,
	0x36, 0xfb, 0xf9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xf3, 0x79, 0x50, 0x85, 0x08, 0x00,
	0x00,
}
