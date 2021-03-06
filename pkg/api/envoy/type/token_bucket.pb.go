// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/token_bucket.proto

package envoy_type

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Configures a token bucket, typically used for rate limiting.
type TokenBucket struct {
	// The maximum tokens that the bucket can hold. This is also the number of tokens that the bucket
	// initially contains.
	MaxTokens uint32 `protobuf:"varint,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// The number of tokens added to the bucket during each fill interval. If not specified, defaults
	// to a single token.
	TokensPerFill *types.UInt32Value `protobuf:"bytes,2,opt,name=tokens_per_fill,json=tokensPerFill,proto3" json:"tokens_per_fill,omitempty"`
	// The fill interval that tokens are added to the bucket. During each fill interval
	// `tokens_per_fill` are added to the bucket. The bucket will never contain more than
	// `max_tokens` tokens.
	FillInterval         *types.Duration `protobuf:"bytes,3,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TokenBucket) Reset()         { *m = TokenBucket{} }
func (m *TokenBucket) String() string { return proto.CompactTextString(m) }
func (*TokenBucket) ProtoMessage()    {}
func (*TokenBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_04e870b436501a80, []int{0}
}
func (m *TokenBucket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenBucket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenBucket.Merge(m, src)
}
func (m *TokenBucket) XXX_Size() int {
	return m.Size()
}
func (m *TokenBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenBucket.DiscardUnknown(m)
}

var xxx_messageInfo_TokenBucket proto.InternalMessageInfo

func (m *TokenBucket) GetMaxTokens() uint32 {
	if m != nil {
		return m.MaxTokens
	}
	return 0
}

func (m *TokenBucket) GetTokensPerFill() *types.UInt32Value {
	if m != nil {
		return m.TokensPerFill
	}
	return nil
}

func (m *TokenBucket) GetFillInterval() *types.Duration {
	if m != nil {
		return m.FillInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenBucket)(nil), "envoy.type.TokenBucket")
}

func init() { proto.RegisterFile("envoy/type/token_bucket.proto", fileDescriptor_04e870b436501a80) }

var fileDescriptor_04e870b436501a80 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x6a, 0x32, 0x31,
	0x1c, 0xc5, 0x8d, 0x9f, 0xe8, 0xd7, 0x58, 0xa9, 0xcc, 0xa6, 0xd3, 0x52, 0x07, 0xe9, 0xa2, 0x88,
	0x8b, 0x04, 0xf4, 0x06, 0xa1, 0x14, 0x2c, 0x14, 0x44, 0xda, 0x6e, 0x87, 0x58, 0xa3, 0x04, 0x63,
	0x12, 0x32, 0x89, 0xd5, 0x5d, 0xcf, 0xd3, 0x23, 0xf4, 0x04, 0x2e, 0xed, 0x0d, 0xca, 0x1c, 0xc3,
	0x55, 0x99, 0xcc, 0x0c, 0x16, 0xdc, 0x05, 0x7e, 0xef, 0xbd, 0xff, 0xcb, 0x83, 0x1d, 0x26, 0xd7,
	0x6a, 0x8b, 0xed, 0x56, 0x33, 0x6c, 0xd5, 0x92, 0xc9, 0x78, 0xea, 0xde, 0x96, 0xcc, 0x22, 0x6d,
	0x94, 0x55, 0x01, 0xf4, 0x18, 0x65, 0xf8, 0x3a, 0x5a, 0x28, 0xb5, 0x10, 0x0c, 0x7b, 0x32, 0x75,
	0x73, 0x3c, 0x73, 0x86, 0x5a, 0xae, 0x64, 0xae, 0x3d, 0xe5, 0xef, 0x86, 0x6a, 0xcd, 0x4c, 0x52,
	0xf0, 0x8e, 0x9b, 0x69, 0x8a, 0xa9, 0x94, 0xca, 0x7a, 0x5b, 0x82, 0x13, 0x4b, 0xad, 0x2b, 0xf1,
	0xe5, 0x9a, 0x0a, 0x3e, 0xa3, 0x96, 0xe1, 0xf2, 0x91, 0x83, 0xdb, 0x3d, 0x80, 0xcd, 0xe7, 0xac,
	0x1a, 0xf1, 0xcd, 0x82, 0x3b, 0x08, 0x57, 0x74, 0x13, 0xfb, 0xb6, 0x49, 0x08, 0xba, 0xa0, 0xd7,
	0x22, 0x8d, 0x03, 0xa9, 0xf5, 0xab, 0xdd, 0xca, 0xe4, 0x6c, 0x45, 0x37, 0x5e, 0x9c, 0x04, 0x4f,
	0xf0, 0x22, 0xd7, 0xc4, 0x9a, 0x99, 0x78, 0xce, 0x85, 0x08, 0xab, 0x5d, 0xd0, 0x6b, 0x0e, 0x6e,
	0x50, 0xde, 0x14, 0x95, 0x4d, 0xd1, 0xcb, 0x48, 0xda, 0xe1, 0xe0, 0x95, 0x0a, 0xc7, 0x8e, 0x51,
	0xad, 0xdc, 0x3d, 0x66, 0xe6, 0x81, 0x0b, 0x11, 0x3c, 0xc2, 0x56, 0x96, 0x11, 0x73, 0x69, 0x99,
	0x59, 0x53, 0x11, 0xfe, 0xf3, 0x61, 0x57, 0x27, 0x61, 0xf7, 0xc5, 0x2c, 0x04, 0x1e, 0x48, 0xe3,
	0x13, 0xd4, 0xfe, 0x83, 0x7e, 0x65, 0x72, 0x9e, 0x79, 0x47, 0x85, 0x95, 0x90, 0x5d, 0x1a, 0x81,
	0x7d, 0x1a, 0x81, 0x9f, 0x34, 0x02, 0x5f, 0x1f, 0xbb, 0xef, 0x7a, 0xb5, 0x5d, 0x81, 0x21, 0x57,
	0xc8, 0xef, 0xad, 0x8d, 0xda, 0x6c, 0xd1, 0x71, 0x7a, 0xd2, 0xfe, 0xf3, 0xff, 0x71, 0x76, 0x67,
	0x0c, 0xa6, 0x75, 0x7f, 0x70, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x2a, 0x11, 0x70, 0xc2,
	0x01, 0x00, 0x00,
}

func (m *TokenBucket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenBucket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenBucket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.FillInterval != nil {
		{
			size, err := m.FillInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTokenBucket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TokensPerFill != nil {
		{
			size, err := m.TokensPerFill.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTokenBucket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxTokens != 0 {
		i = encodeVarintTokenBucket(dAtA, i, uint64(m.MaxTokens))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenBucket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenBucket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenBucket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTokens != 0 {
		n += 1 + sovTokenBucket(uint64(m.MaxTokens))
	}
	if m.TokensPerFill != nil {
		l = m.TokensPerFill.Size()
		n += 1 + l + sovTokenBucket(uint64(l))
	}
	if m.FillInterval != nil {
		l = m.FillInterval.Size()
		n += 1 + l + sovTokenBucket(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTokenBucket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenBucket(x uint64) (n int) {
	return sovTokenBucket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenBucket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenBucket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenBucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenBucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTokens", wireType)
			}
			m.MaxTokens = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTokens |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensPerFill", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTokenBucket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TokensPerFill == nil {
				m.TokensPerFill = &types.UInt32Value{}
			}
			if err := m.TokensPerFill.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FillInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTokenBucket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FillInterval == nil {
				m.FillInterval = &types.Duration{}
			}
			if err := m.FillInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenBucket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokenBucket
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTokenBucket
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
func skipTokenBucket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenBucket
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
					return 0, ErrIntOverflowTokenBucket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTokenBucket
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
			if length < 0 {
				return 0, ErrInvalidLengthTokenBucket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenBucket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenBucket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenBucket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenBucket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenBucket = fmt.Errorf("proto: unexpected end of group")
)
