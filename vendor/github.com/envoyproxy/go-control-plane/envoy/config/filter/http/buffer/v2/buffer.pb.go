// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes      *types.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_buffer_15878113f6bc9bdb, []int{0}
}
func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(dst, src)
}
func (m *Buffer) XXX_Size() int {
	return m.Size()
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_buffer_15878113f6bc9bdb, []int{1}
}
func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(dst, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return m.Size()
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}
type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}
func (*BufferPerRoute_Buffer) isBufferPerRoute_Override()   {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BufferPerRoute_OneofMarshaler, _BufferPerRoute_OneofUnmarshaler, _BufferPerRoute_OneofSizer, []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func _BufferPerRoute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BufferPerRoute)
	// override
	switch x := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		t := uint64(0)
		if x.Disabled {
			t = 1
		}
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *BufferPerRoute_Buffer:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buffer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BufferPerRoute.Override has unexpected type %T", x)
	}
	return nil
}

func _BufferPerRoute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BufferPerRoute)
	switch tag {
	case 1: // override.disabled
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Override = &BufferPerRoute_Disabled{x != 0}
		return true, err
	case 2: // override.buffer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Buffer)
		err := b.DecodeMessage(msg)
		m.Override = &BufferPerRoute_Buffer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BufferPerRoute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BufferPerRoute)
	// override
	switch x := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		n += 1 // tag and wire
		n += 1
	case *BufferPerRoute_Buffer:
		s := proto.Size(x.Buffer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}
func (m *Buffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Buffer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.MaxRequestBytes.Size()))
		n1, err := m.MaxRequestBytes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BufferPerRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BufferPerRoute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Override != nil {
		nn2, err := m.Override.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BufferPerRoute_Disabled) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *BufferPerRoute_Buffer) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Buffer != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.Buffer.Size()))
		n3, err := m.Buffer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func encodeVarintBuffer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Buffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		l = m.MaxRequestBytes.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BufferPerRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Override != nil {
		n += m.Override.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BufferPerRoute_Disabled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *BufferPerRoute_Buffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Buffer != nil {
		l = m.Buffer.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	return n
}

func sovBuffer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBuffer(x uint64) (n int) {
	return sovBuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Buffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Buffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Buffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequestBytes == nil {
				m.MaxRequestBytes = &types.UInt32Value{}
			}
			if err := m.MaxRequestBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BufferPerRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BufferPerRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BufferPerRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Override = &BufferPerRoute_Disabled{b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buffer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Buffer{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Override = &BufferPerRoute_Buffer{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBuffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuffer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBuffer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBuffer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBuffer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBuffer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuffer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptor_buffer_15878113f6bc9bdb)
}

var fileDescriptor_buffer_15878113f6bc9bdb = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xeb, 0xf4, 0x8f, 0x72, 0x7d, 0xa5, 0x7b, 0xdb, 0xe8, 0x4a, 0xb7, 0x42, 0x28, 0xaa,
	0xba, 0x80, 0x3a, 0xd8, 0x28, 0x7d, 0x83, 0x88, 0x01, 0x98, 0xaa, 0x20, 0x18, 0x58, 0x2a, 0xa7,
	0xf9, 0x12, 0x8c, 0xd2, 0x38, 0x38, 0x4e, 0x68, 0x5f, 0x81, 0x87, 0xe0, 0x39, 0x10, 0x53, 0x47,
	0x46, 0x1e, 0x01, 0x75, 0xeb, 0x5b, 0xa0, 0xd8, 0x29, 0x0b, 0x03, 0x6c, 0x47, 0x39, 0x39, 0xbf,
	0x9f, 0xf5, 0x61, 0x0a, 0x59, 0x25, 0xd6, 0x74, 0x21, 0xb2, 0x98, 0x27, 0x34, 0xe6, 0xa9, 0x02,
	0x49, 0x6f, 0x95, 0xca, 0x69, 0x58, 0xc6, 0x31, 0x48, 0x5a, 0x79, 0x4d, 0x22, 0xb9, 0x14, 0x4a,
	0x38, 0x63, 0x3d, 0x20, 0x66, 0x40, 0xcc, 0x80, 0xd4, 0x03, 0xd2, 0xfc, 0x56, 0x79, 0x07, 0x6e,
	0x22, 0x44, 0x92, 0x02, 0xd5, 0x8b, 0xb0, 0x8c, 0x69, 0x54, 0x4a, 0xa6, 0xb8, 0xc8, 0x0c, 0xe3,
	0x6b, 0xff, 0x20, 0x59, 0x9e, 0x83, 0x2c, 0x9a, 0xfe, 0x7f, 0xc5, 0x52, 0x1e, 0x31, 0x05, 0x74,
	0x1f, 0x9a, 0xe2, 0x5f, 0x22, 0x12, 0xa1, 0x23, 0xad, 0x93, 0xf9, 0x3a, 0x5e, 0xe0, 0x9e, 0xaf,
	0xdd, 0xce, 0x25, 0x1e, 0x2c, 0xd9, 0x6a, 0x2e, 0xe1, 0xbe, 0x84, 0x42, 0xcd, 0xc3, 0xb5, 0x82,
	0x62, 0x88, 0x46, 0xe8, 0xf8, 0xb7, 0x77, 0x48, 0x8c, 0x94, 0xec, 0xa5, 0xe4, 0xea, 0x3c, 0x53,
	0x53, 0xef, 0x9a, 0xa5, 0x25, 0xf8, 0xbf, 0x5e, 0x76, 0x9b, 0x76, 0x67, 0x62, 0x8d, 0x5a, 0xc1,
	0xdf, 0x25, 0x5b, 0x05, 0x06, 0xe0, 0xd7, 0xfb, 0x8b, 0x8e, 0x6d, 0xf5, 0xdb, 0xe3, 0x27, 0x84,
	0xff, 0x18, 0xcb, 0x0c, 0x64, 0x20, 0x4a, 0x05, 0xce, 0x11, 0xb6, 0x23, 0x5e, 0xb0, 0x30, 0x85,
	0x48, 0x4b, 0xec, 0x06, 0x73, 0x67, 0xd9, 0xe8, 0xac, 0x15, 0x7c, 0x96, 0xce, 0x0c, 0xf7, 0xcc,
	0x71, 0x86, 0x96, 0x7e, 0xcb, 0x84, 0x7c, 0x7f, 0x44, 0x62, 0x64, 0x3e, 0xae, 0x91, 0xdd, 0x47,
	0x64, 0xf5, 0x6b, 0x66, 0xc3, 0xf1, 0x07, 0xd8, 0x16, 0x15, 0x48, 0xc9, 0x23, 0x70, 0xba, 0xcf,
	0xbb, 0x4d, 0x1b, 0xf9, 0xa7, 0xaf, 0x5b, 0x17, 0xbd, 0x6d, 0x5d, 0xf4, 0xbe, 0x75, 0x11, 0x3e,
	0xe1, 0xc2, 0x48, 0x72, 0x29, 0x56, 0xeb, 0x1f, 0xf8, 0x66, 0xe8, 0xc6, 0xaa, 0xbc, 0xb0, 0xa7,
	0xcf, 0x33, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x55, 0x9e, 0x79, 0x3f, 0x18, 0x02, 0x00, 0x00,
}