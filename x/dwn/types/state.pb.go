// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dwn/v1/state.proto

package types

import (
	_ "cosmossdk.io/orm"
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Credential struct {
	Id              []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind            string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Transports      []string `protobuf:"bytes,3,rep,name=transports,proto3" json:"transports,omitempty"`
	PublicKey       []byte   `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AttestationType string   `protobuf:"bytes,5,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	CreatedAt       uint64   `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_040a9b061177db90, []int{0}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Credential) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Credential) GetTransports() []string {
	if m != nil {
		return m.Transports
	}
	return nil
}

func (m *Credential) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Credential) GetAttestationType() string {
	if m != nil {
		return m.AttestationType
	}
	return ""
}

func (m *Credential) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Profile struct {
	Account []byte `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount  uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_040a9b061177db90, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Profile) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Credential)(nil), "dwn.v1.Credential")
	proto.RegisterType((*Profile)(nil), "dwn.v1.Profile")
}

func init() { proto.RegisterFile("dwn/v1/state.proto", fileDescriptor_040a9b061177db90) }

var fileDescriptor_040a9b061177db90 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x69, 0x4c, 0xcd, 0x20, 0xb5, 0xec, 0x41, 0x57, 0xc1, 0x18, 0xea, 0x25, 0x1e,
	0x6c, 0x28, 0xde, 0x0a, 0x1e, 0xd4, 0xa3, 0x17, 0x09, 0x9e, 0x44, 0x28, 0xdb, 0x64, 0xd5, 0xa5,
	0xcd, 0x6e, 0xd8, 0x4c, 0x5b, 0xf3, 0x12, 0xe2, 0x13, 0xf8, 0x3c, 0x7a, 0x2b, 0x78, 0xf1, 0x28,
	0xed, 0x1b, 0xf8, 0x04, 0x92, 0xb4, 0x95, 0xde, 0x76, 0xbe, 0xfd, 0xf7, 0x63, 0x66, 0x07, 0x68,
	0x32, 0x55, 0xe1, 0xa4, 0x1b, 0xe6, 0xc8, 0x51, 0x74, 0x32, 0xa3, 0x51, 0x53, 0x27, 0x99, 0xaa,
	0xce, 0xa4, 0x7b, 0xb8, 0x1f, 0xeb, 0x3c, 0xd5, 0x79, 0xa8, 0x4d, 0x5a, 0x46, 0xb4, 0x49, 0x97,
	0x81, 0xf6, 0x27, 0x01, 0xb8, 0x36, 0x22, 0x11, 0x0a, 0x25, 0x1f, 0xd1, 0x26, 0x58, 0x32, 0x61,
	0xc4, 0x27, 0xc1, 0x4e, 0x64, 0xc9, 0x84, 0x52, 0xb0, 0x87, 0x52, 0x25, 0xcc, 0xf2, 0x49, 0xe0,
	0x46, 0xd5, 0x99, 0x7a, 0x00, 0x68, 0xb8, 0xca, 0x33, 0x6d, 0x30, 0x67, 0x75, 0xbf, 0x1e, 0xb8,
	0xd1, 0x06, 0xa1, 0x47, 0x00, 0xd9, 0x78, 0x30, 0x92, 0x71, 0x7f, 0x28, 0x0a, 0x66, 0x57, 0x2e,
	0x77, 0x49, 0x6e, 0x44, 0x41, 0x4f, 0xa1, 0xc5, 0x11, 0x45, 0xd9, 0xa5, 0xd4, 0xaa, 0x8f, 0x45,
	0x26, 0xd8, 0x56, 0xa5, 0xdf, 0xdd, 0xe0, 0x77, 0x45, 0x26, 0x4a, 0x53, 0x6c, 0x04, 0x47, 0x91,
	0xf4, 0x39, 0x32, 0xc7, 0x27, 0x81, 0x1d, 0xb9, 0x2b, 0x72, 0x89, 0xbd, 0xe6, 0xef, 0xfb, 0xd7,
	0x6b, 0x7d, 0x1b, 0xec, 0x65, 0xd3, 0xed, 0x07, 0x68, 0xdc, 0x1a, 0xfd, 0x28, 0x47, 0x82, 0x32,
	0x68, 0xf0, 0x38, 0xd6, 0x63, 0x85, 0xab, 0x61, 0xd6, 0x25, 0xdd, 0x03, 0x87, 0xa7, 0xd5, 0x85,
	0x55, 0xf9, 0x56, 0x55, 0xef, 0xb8, 0x92, 0x1d, 0x80, 0xfb, 0xff, 0x92, 0xc2, 0x3a, 0xda, 0x22,
	0xcc, 0xba, 0xba, 0xf8, 0x98, 0x7b, 0x64, 0x36, 0xf7, 0xc8, 0xcf, 0xdc, 0x23, 0x6f, 0x0b, 0xaf,
	0x36, 0x5b, 0x78, 0xb5, 0xef, 0x85, 0x57, 0xbb, 0x3f, 0x79, 0x92, 0xf8, 0x3c, 0x1e, 0x74, 0x62,
	0x9d, 0x86, 0xb9, 0x56, 0xe6, 0x4c, 0xea, 0x30, 0xd6, 0x46, 0x84, 0x2f, 0x61, 0xb9, 0x92, 0x72,
	0xc2, 0x7c, 0xe0, 0x54, 0xff, 0x7d, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x50, 0x2d, 0xd6, 0xbd,
	0xa6, 0x01, 0x00, 0x00,
}

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x30
	}
	if len(m.AttestationType) > 0 {
		i -= len(m.AttestationType)
		copy(dAtA[i:], m.AttestationType)
		i = encodeVarintState(dAtA, i, uint64(len(m.AttestationType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintState(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Transports) > 0 {
		for iNdEx := len(m.Transports) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transports[iNdEx])
			copy(dAtA[i:], m.Transports[iNdEx])
			i = encodeVarintState(dAtA, i, uint64(len(m.Transports[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintState(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintState(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if len(m.Transports) > 0 {
		for _, s := range m.Transports {
			l = len(s)
			n += 1 + l + sovState(uint64(l))
		}
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.AttestationType)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovState(uint64(m.CreatedAt))
	}
	return n
}

func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovState(uint64(m.Amount))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transports", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transports = append(m.Transports, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = append(m.Account[:0], dAtA[iNdEx:postIndex]...)
			if m.Account == nil {
				m.Account = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
