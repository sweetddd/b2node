// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/bridge/v1/deposit.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// DepositStatus represents the status of a deposit record.
type DepositStatus int32

const (
	// DEPOSIT_STATUS_UNSPECIFIED represents an unspecified status.
	DepositStatus_DEPOSIT_STATUS_UNSPECIFIED DepositStatus = 0
	// DEPOSIT_STATUS_PENDING represents a deposit record that is pending.
	DepositStatus_DEPOSIT_STATUS_PENDING DepositStatus = 1
	// DEPOSIT_STATUS_COMPLETED represents a deposit record that has been processed and completed.
	DepositStatus_DEPOSIT_STATUS_COMPLETED DepositStatus = 2
	// DEPOSIT_STATUS_FAILED represents a deposit record that has been processed and failed.
	DepositStatus_DEPOSIT_STATUS_FAILED DepositStatus = 3
)

var DepositStatus_name = map[int32]string{
	0: "DEPOSIT_STATUS_UNSPECIFIED",
	1: "DEPOSIT_STATUS_PENDING",
	2: "DEPOSIT_STATUS_COMPLETED",
	3: "DEPOSIT_STATUS_FAILED",
}

var DepositStatus_value = map[string]int32{
	"DEPOSIT_STATUS_UNSPECIFIED": 0,
	"DEPOSIT_STATUS_PENDING":     1,
	"DEPOSIT_STATUS_COMPLETED":   2,
	"DEPOSIT_STATUS_FAILED":      3,
}

func (x DepositStatus) String() string {
	return proto.EnumName(DepositStatus_name, int32(x))
}

func (DepositStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_426311765de7f009, []int{0}
}

// Deposit represents a record of deposit operation.
type Deposit struct {
	// tx_hash is the transaction hash of the deposit.
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// from is the sender address of the deposit.
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// to is the recipient address on the destination chain.
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// coin_type is the coin being deposited.
	CoinType string `protobuf:"bytes,4,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	// value is the amount of coins being deposited.
	Value uint64 `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	// data is the tx data of the deposit.
	Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// status is the status of the deposit record.
	// status can be: Pending, Completed, Failed.
	Status DepositStatus `protobuf:"varint,7,opt,name=status,proto3,enum=ethermint.bridge.v1.DepositStatus" json:"status,omitempty"`
	// creator is the sender of message.
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_426311765de7f009, []int{0}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Deposit) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Deposit) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Deposit) GetCoinType() string {
	if m != nil {
		return m.CoinType
	}
	return ""
}

func (m *Deposit) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Deposit) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Deposit) GetStatus() DepositStatus {
	if m != nil {
		return m.Status
	}
	return DepositStatus_DEPOSIT_STATUS_UNSPECIFIED
}

func (m *Deposit) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterEnum("ethermint.bridge.v1.DepositStatus", DepositStatus_name, DepositStatus_value)
	proto.RegisterType((*Deposit)(nil), "ethermint.bridge.v1.Deposit")
}

func init() { proto.RegisterFile("ethermint/bridge/v1/deposit.proto", fileDescriptor_426311765de7f009) }

var fileDescriptor_426311765de7f009 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x6f, 0xaa, 0x40,
	0x14, 0x65, 0x50, 0x41, 0x27, 0x79, 0x86, 0xcc, 0xfb, 0x9a, 0xe7, 0x6b, 0x88, 0x75, 0x65, 0xbb,
	0x80, 0xd8, 0xee, 0xba, 0xb3, 0x82, 0x2d, 0x89, 0x55, 0x22, 0xb8, 0xe9, 0x86, 0xa0, 0x4e, 0x85,
	0xa4, 0x38, 0x04, 0x46, 0xa2, 0xfb, 0xfe, 0x80, 0xfe, 0xac, 0x2e, 0x5d, 0x76, 0xd9, 0xc8, 0x1f,
	0x69, 0x00, 0x6d, 0x53, 0xd3, 0xdd, 0xbd, 0xe7, 0x9c, 0x7b, 0x72, 0x72, 0x0f, 0x3c, 0x25, 0xcc,
	0x23, 0x51, 0xe0, 0x2f, 0x99, 0x3a, 0x8d, 0xfc, 0xf9, 0x82, 0xa8, 0x49, 0x47, 0x9d, 0x93, 0x90,
	0xc6, 0x3e, 0x53, 0xc2, 0x88, 0x32, 0x8a, 0x7e, 0x7e, 0x48, 0x94, 0x42, 0xa2, 0x24, 0x9d, 0x56,
	0x0a, 0xa0, 0xa8, 0x15, 0x32, 0xf4, 0x17, 0x8a, 0x6c, 0xed, 0x78, 0x6e, 0xec, 0x61, 0xd0, 0x04,
	0xed, 0xda, 0x58, 0x60, 0xeb, 0x5b, 0x37, 0xf6, 0x10, 0x82, 0xe5, 0x87, 0x88, 0x06, 0x98, 0xcf,
	0xd1, 0x7c, 0x46, 0x75, 0xc8, 0x33, 0x8a, 0x4b, 0x39, 0xc2, 0x33, 0x8a, 0xfe, 0xc3, 0xda, 0x8c,
	0xfa, 0x4b, 0x87, 0x6d, 0x42, 0x82, 0xcb, 0x39, 0x5c, 0xcd, 0x00, 0x7b, 0x13, 0x12, 0xf4, 0x0b,
	0x56, 0x12, 0xf7, 0x71, 0x45, 0x70, 0xa5, 0x09, 0xda, 0xe5, 0x71, 0xb1, 0x64, 0xb6, 0x73, 0x97,
	0xb9, 0x58, 0x28, 0x6c, 0xb3, 0x19, 0x5d, 0x41, 0x21, 0x66, 0x2e, 0x5b, 0xc5, 0x58, 0x6c, 0x82,
	0x76, 0xfd, 0xa2, 0xa5, 0x7c, 0x93, 0x5a, 0xd9, 0x27, 0xb6, 0x72, 0xe5, 0x78, 0x7f, 0x81, 0x30,
	0x14, 0x67, 0x11, 0x71, 0x19, 0x8d, 0x70, 0x35, 0xb7, 0x3c, 0xac, 0xe7, 0x4f, 0x00, 0xfe, 0xf8,
	0x72, 0x83, 0x64, 0xd8, 0xd0, 0x74, 0x73, 0x64, 0x19, 0xb6, 0x63, 0xd9, 0x5d, 0x7b, 0x62, 0x39,
	0x93, 0xa1, 0x65, 0xea, 0x3d, 0xa3, 0x6f, 0xe8, 0x9a, 0xc4, 0xa1, 0x06, 0xfc, 0x73, 0xc4, 0x9b,
	0xfa, 0x50, 0x33, 0x86, 0x37, 0x12, 0x40, 0x27, 0x10, 0x1f, 0x71, 0xbd, 0xd1, 0x9d, 0x39, 0xd0,
	0x6d, 0x5d, 0x93, 0x78, 0xf4, 0x0f, 0xfe, 0x3e, 0x62, 0xfb, 0x5d, 0x63, 0xa0, 0x6b, 0x52, 0xe9,
	0xba, 0xf7, 0xb2, 0x93, 0xc1, 0x76, 0x27, 0x83, 0xb7, 0x9d, 0x0c, 0x9e, 0x53, 0x99, 0xdb, 0xa6,
	0x32, 0xf7, 0x9a, 0xca, 0xdc, 0xfd, 0xd9, 0xc2, 0x67, 0xde, 0x6a, 0xaa, 0xcc, 0x68, 0xa0, 0x92,
	0x24, 0xa0, 0xb1, 0xfa, 0xd9, 0xe7, 0xfa, 0xd0, 0x68, 0xf6, 0xda, 0x78, 0x2a, 0xe4, 0x6d, 0x5e,
	0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x3b, 0x6e, 0x39, 0xf2, 0x01, 0x00, 0x00,
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if m.Status != 0 {
		i = encodeVarintDeposit(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if m.Value != 0 {
		i = encodeVarintDeposit(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x28
	}
	if len(m.CoinType) > 0 {
		i -= len(m.CoinType)
		copy(dAtA[i:], m.CoinType)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.CoinType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.CoinType)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovDeposit(uint64(m.Value))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovDeposit(uint64(m.Status))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= DepositStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
