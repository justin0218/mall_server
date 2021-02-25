// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/proto.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Sku struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	GoodsId              int64    `protobuf:"varint,4,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	Price                int64    `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Inventory            int64    `protobuf:"varint,6,opt,name=inventory,proto3" json:"inventory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sku) Reset()         { *m = Sku{} }
func (m *Sku) String() string { return proto.CompactTextString(m) }
func (*Sku) ProtoMessage()    {}
func (*Sku) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{0}
}

func (m *Sku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sku.Unmarshal(m, b)
}
func (m *Sku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sku.Marshal(b, m, deterministic)
}
func (m *Sku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sku.Merge(m, src)
}
func (m *Sku) XXX_Size() int {
	return xxx_messageInfo_Sku.Size(m)
}
func (m *Sku) XXX_DiscardUnknown() {
	xxx_messageInfo_Sku.DiscardUnknown(m)
}

var xxx_messageInfo_Sku proto.InternalMessageInfo

func (m *Sku) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Sku) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sku) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *Sku) GetGoodsId() int64 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *Sku) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Sku) GetInventory() int64 {
	if m != nil {
		return m.Inventory
	}
	return 0
}

type GetGoodsDetailReq struct {
	GoodsId              int64    `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGoodsDetailReq) Reset()         { *m = GetGoodsDetailReq{} }
func (m *GetGoodsDetailReq) String() string { return proto.CompactTextString(m) }
func (*GetGoodsDetailReq) ProtoMessage()    {}
func (*GetGoodsDetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{1}
}

func (m *GetGoodsDetailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGoodsDetailReq.Unmarshal(m, b)
}
func (m *GetGoodsDetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGoodsDetailReq.Marshal(b, m, deterministic)
}
func (m *GetGoodsDetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGoodsDetailReq.Merge(m, src)
}
func (m *GetGoodsDetailReq) XXX_Size() int {
	return xxx_messageInfo_GetGoodsDetailReq.Size(m)
}
func (m *GetGoodsDetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGoodsDetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetGoodsDetailReq proto.InternalMessageInfo

func (m *GetGoodsDetailReq) GetGoodsId() int64 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

type GetGoodsDetailRes struct {
	GoodsId              int64    `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Images               string   `protobuf:"bytes,3,opt,name=images,proto3" json:"images,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Skus                 []*Sku   `protobuf:"bytes,5,rep,name=skus,proto3" json:"skus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGoodsDetailRes) Reset()         { *m = GetGoodsDetailRes{} }
func (m *GetGoodsDetailRes) String() string { return proto.CompactTextString(m) }
func (*GetGoodsDetailRes) ProtoMessage()    {}
func (*GetGoodsDetailRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{2}
}

func (m *GetGoodsDetailRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGoodsDetailRes.Unmarshal(m, b)
}
func (m *GetGoodsDetailRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGoodsDetailRes.Marshal(b, m, deterministic)
}
func (m *GetGoodsDetailRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGoodsDetailRes.Merge(m, src)
}
func (m *GetGoodsDetailRes) XXX_Size() int {
	return xxx_messageInfo_GetGoodsDetailRes.Size(m)
}
func (m *GetGoodsDetailRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGoodsDetailRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetGoodsDetailRes proto.InternalMessageInfo

func (m *GetGoodsDetailRes) GetGoodsId() int64 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *GetGoodsDetailRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetGoodsDetailRes) GetImages() string {
	if m != nil {
		return m.Images
	}
	return ""
}

func (m *GetGoodsDetailRes) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetGoodsDetailRes) GetSkus() []*Sku {
	if m != nil {
		return m.Skus
	}
	return nil
}

type CreateOrderReq struct {
	GoodsId              int64    `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	SkuId                int64    `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	BuyNum               int64    `protobuf:"varint,3,opt,name=buy_num,json=buyNum,proto3" json:"buy_num,omitempty"`
	Uid                  int64    `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Province             string   `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	Region               string   `protobuf:"bytes,9,opt,name=region,proto3" json:"region,omitempty"`
	Addr                 string   `protobuf:"bytes,10,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderReq) Reset()         { *m = CreateOrderReq{} }
func (m *CreateOrderReq) String() string { return proto.CompactTextString(m) }
func (*CreateOrderReq) ProtoMessage()    {}
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{3}
}

func (m *CreateOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderReq.Unmarshal(m, b)
}
func (m *CreateOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderReq.Marshal(b, m, deterministic)
}
func (m *CreateOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderReq.Merge(m, src)
}
func (m *CreateOrderReq) XXX_Size() int {
	return xxx_messageInfo_CreateOrderReq.Size(m)
}
func (m *CreateOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderReq proto.InternalMessageInfo

func (m *CreateOrderReq) GetGoodsId() int64 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *CreateOrderReq) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *CreateOrderReq) GetBuyNum() int64 {
	if m != nil {
		return m.BuyNum
	}
	return 0
}

func (m *CreateOrderReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CreateOrderReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CreateOrderReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrderReq) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *CreateOrderReq) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CreateOrderReq) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *CreateOrderReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type CreateOrderRes struct {
	OrderCode            string   `protobuf:"bytes,1,opt,name=order_code,json=orderCode,proto3" json:"order_code,omitempty"`
	GoodsName            string   `protobuf:"bytes,2,opt,name=goods_name,json=goodsName,proto3" json:"goods_name,omitempty"`
	SkuName              string   `protobuf:"bytes,3,opt,name=sku_name,json=skuName,proto3" json:"sku_name,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRes) Reset()         { *m = CreateOrderRes{} }
func (m *CreateOrderRes) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRes) ProtoMessage()    {}
func (*CreateOrderRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{4}
}

func (m *CreateOrderRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRes.Unmarshal(m, b)
}
func (m *CreateOrderRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRes.Marshal(b, m, deterministic)
}
func (m *CreateOrderRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRes.Merge(m, src)
}
func (m *CreateOrderRes) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRes.Size(m)
}
func (m *CreateOrderRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRes proto.InternalMessageInfo

func (m *CreateOrderRes) GetOrderCode() string {
	if m != nil {
		return m.OrderCode
	}
	return ""
}

func (m *CreateOrderRes) GetGoodsName() string {
	if m != nil {
		return m.GoodsName
	}
	return ""
}

func (m *CreateOrderRes) GetSkuName() string {
	if m != nil {
		return m.SkuName
	}
	return ""
}

func (m *CreateOrderRes) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*Sku)(nil), "sku")
	proto.RegisterType((*GetGoodsDetailReq)(nil), "get_goods_detail_req")
	proto.RegisterType((*GetGoodsDetailRes)(nil), "get_goods_detail_res")
	proto.RegisterType((*CreateOrderReq)(nil), "create_order_req")
	proto.RegisterType((*CreateOrderRes)(nil), "create_order_res")
}

func init() { proto.RegisterFile("api/proto/proto.proto", fileDescriptor_ef32c37ea206d67b) }

var fileDescriptor_ef32c37ea206d67b = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x3d, 0x8f, 0x13, 0x31,
	0x10, 0x3d, 0x67, 0xbf, 0xb2, 0x03, 0x42, 0xc1, 0xba, 0x80, 0x2f, 0x02, 0x29, 0xda, 0x2a, 0x55,
	0x10, 0x87, 0xc4, 0x0f, 0x38, 0xaa, 0x6b, 0x52, 0x6c, 0x49, 0xb3, 0x72, 0xd6, 0x56, 0xb0, 0x36,
	0xbb, 0x0e, 0xfe, 0x88, 0x94, 0x12, 0x7e, 0x00, 0xfc, 0x5a, 0x7a, 0xe4, 0x71, 0x14, 0x72, 0x77,
	0x11, 0xd7, 0x38, 0xf3, 0xde, 0xcc, 0xc4, 0x7e, 0x6f, 0x66, 0x61, 0xca, 0x77, 0xea, 0xc3, 0xce,
	0x68, 0xa7, 0xe3, 0xb9, 0xc4, 0xb3, 0xfa, 0x45, 0x20, 0xb1, 0x9d, 0xa7, 0xaf, 0x60, 0xa4, 0x04,
	0x23, 0x73, 0xb2, 0x48, 0xea, 0x91, 0x12, 0x94, 0x42, 0x3a, 0xf0, 0x5e, 0xb2, 0xd1, 0x9c, 0x2c,
	0xca, 0x1a, 0x63, 0x7a, 0x0d, 0x59, 0xab, 0xf7, 0xd2, 0xb0, 0x04, 0xc9, 0x08, 0xe8, 0x0d, 0x8c,
	0x37, 0x5a, 0x0b, 0xdb, 0x28, 0xc1, 0x52, 0xec, 0x2f, 0x10, 0xdf, 0x8b, 0xd0, 0xb0, 0x33, 0xaa,
	0x95, 0x2c, 0x43, 0x3e, 0x02, 0xfa, 0x0e, 0x4a, 0x35, 0xec, 0xe5, 0xe0, 0xb4, 0x39, 0xb0, 0x1c,
	0x33, 0xff, 0x88, 0xea, 0x23, 0x5c, 0x6f, 0xa4, 0x6b, 0xe2, 0x5f, 0x0a, 0xe9, 0xb8, 0xda, 0x36,
	0x46, 0x7e, 0x7f, 0x70, 0x0d, 0x79, 0x70, 0x4d, 0xf5, 0x9b, 0x5c, 0xec, 0xb1, 0xff, 0xe9, 0xb9,
	0xa8, 0xef, 0x0d, 0xe4, 0xaa, 0xe7, 0x1b, 0x69, 0x8f, 0x02, 0x8f, 0x28, 0xf0, 0xd6, 0x71, 0xe7,
	0x2d, 0xea, 0xcb, 0xea, 0x23, 0xa2, 0x0c, 0x52, 0xdb, 0x79, 0xcb, 0xb2, 0x79, 0xb2, 0x78, 0x71,
	0x9b, 0x2e, 0x6d, 0xe7, 0x6b, 0x64, 0xaa, 0x3f, 0x04, 0x26, 0xad, 0x91, 0xdc, 0xc9, 0x46, 0x1b,
	0x21, 0xcd, 0x33, 0x0a, 0xe8, 0x14, 0x72, 0xdb, 0xf9, 0x90, 0x18, 0x45, 0xa7, 0x6c, 0xe7, 0xef,
	0x05, 0x7d, 0x0b, 0xc5, 0xda, 0x1f, 0x9a, 0xc1, 0xf7, 0xf8, 0xa2, 0xa4, 0xce, 0xd7, 0xfe, 0xb0,
	0xf2, 0x3d, 0x9d, 0x40, 0xe2, 0x4f, 0x76, 0x87, 0x10, 0xad, 0xfe, 0xa6, 0x87, 0x68, 0x75, 0x59,
	0x47, 0x70, 0x52, 0x99, 0x9f, 0xa9, 0x9c, 0xc1, 0x78, 0x67, 0xf4, 0x5e, 0x0d, 0xad, 0x64, 0x05,
	0xf2, 0x27, 0x1c, 0xea, 0x5b, 0xe5, 0x0e, 0x6c, 0x1c, 0xeb, 0x43, 0x1c, 0xd4, 0x1b, 0xb9, 0x51,
	0x7a, 0x60, 0x65, 0x74, 0x25, 0xa2, 0x50, 0xcb, 0x85, 0x30, 0x0c, 0x62, 0x6d, 0x88, 0xab, 0x1f,
	0x4f, 0x75, 0x5b, 0xfa, 0x1e, 0x20, 0x82, 0x56, 0x0b, 0x89, 0xca, 0xcb, 0xba, 0x44, 0xe6, 0x8b,
	0x16, 0x32, 0xa4, 0xa3, 0x2d, 0x67, 0xf3, 0x28, 0x91, 0x59, 0x85, 0xe7, 0xde, 0xc0, 0x38, 0x58,
	0x83, 0xc9, 0x38, 0x96, 0xc2, 0x76, 0x7e, 0x75, 0xdc, 0xc7, 0xb8, 0x5e, 0xe9, 0xd9, 0x7a, 0xdd,
	0xfe, 0x24, 0x90, 0xf6, 0x7c, 0xbb, 0xa5, 0x77, 0x30, 0x79, 0xbc, 0x15, 0x74, 0xba, 0xbc, 0xb4,
	0x5c, 0xb3, 0x8b, 0xb4, 0xad, 0xae, 0xe8, 0x67, 0x78, 0x79, 0xae, 0x87, 0xbe, 0x5e, 0x3e, 0x1e,
	0xeb, 0xec, 0x09, 0x65, 0xab, 0xab, 0xbb, 0xe2, 0x6b, 0x86, 0xdf, 0xd7, 0x3a, 0xc7, 0x9f, 0x4f,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0x3b, 0xec, 0x81, 0x7f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MallClient is the client API for Mall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MallClient interface {
	GetGoodsDetail(ctx context.Context, in *GetGoodsDetailReq, opts ...grpc.CallOption) (*GetGoodsDetailRes, error)
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error)
}

type mallClient struct {
	cc *grpc.ClientConn
}

func NewMallClient(cc *grpc.ClientConn) MallClient {
	return &mallClient{cc}
}

func (c *mallClient) GetGoodsDetail(ctx context.Context, in *GetGoodsDetailReq, opts ...grpc.CallOption) (*GetGoodsDetailRes, error) {
	out := new(GetGoodsDetailRes)
	err := c.cc.Invoke(ctx, "/mall/get_goods_detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error) {
	out := new(CreateOrderRes)
	err := c.cc.Invoke(ctx, "/mall/create_order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MallServer is the server API for Mall service.
type MallServer interface {
	GetGoodsDetail(context.Context, *GetGoodsDetailReq) (*GetGoodsDetailRes, error)
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error)
}

// UnimplementedMallServer can be embedded to have forward compatible implementations.
type UnimplementedMallServer struct {
}

func (*UnimplementedMallServer) GetGoodsDetail(ctx context.Context, req *GetGoodsDetailReq) (*GetGoodsDetailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (*UnimplementedMallServer) CreateOrder(ctx context.Context, req *CreateOrderReq) (*CreateOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

func RegisterMallServer(s *grpc.Server, srv MallServer) {
	s.RegisterService(&_Mall_serviceDesc, srv)
}

func _Mall_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodsDetail(ctx, req.(*GetGoodsDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mall",
	HandlerType: (*MallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_goods_detail",
			Handler:    _Mall_GetGoodsDetail_Handler,
		},
		{
			MethodName: "create_order",
			Handler:    _Mall_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/proto.proto",
}
