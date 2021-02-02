// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: events.proto

package result

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResultCreatedEvent struct {
	ResultID             string   `protobuf:"bytes,1,opt,name=ResultID,proto3" json:"ResultID,omitempty"`
	CheckpointID         string   `protobuf:"bytes,2,opt,name=CheckpointID,proto3" json:"CheckpointID,omitempty"`
	SportsmenID          string   `protobuf:"bytes,3,opt,name=SportsmenID,proto3" json:"SportsmenID,omitempty"`
	TimeStart            int64    `protobuf:"varint,4,opt,name=TimeStart,proto3" json:"TimeStart,omitempty"`
	Version              uint32   `protobuf:"varint,255,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultCreatedEvent) Reset()         { *m = ResultCreatedEvent{} }
func (m *ResultCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*ResultCreatedEvent) ProtoMessage()    {}
func (*ResultCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}
func (m *ResultCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResultCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResultCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResultCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultCreatedEvent.Merge(m, src)
}
func (m *ResultCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ResultCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ResultCreatedEvent proto.InternalMessageInfo

func (m *ResultCreatedEvent) GetResultID() string {
	if m != nil {
		return m.ResultID
	}
	return ""
}

func (m *ResultCreatedEvent) GetCheckpointID() string {
	if m != nil {
		return m.CheckpointID
	}
	return ""
}

func (m *ResultCreatedEvent) GetSportsmenID() string {
	if m != nil {
		return m.SportsmenID
	}
	return ""
}

func (m *ResultCreatedEvent) GetTimeStart() int64 {
	if m != nil {
		return m.TimeStart
	}
	return 0
}

func (m *ResultCreatedEvent) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ResultFinishedEvent struct {
	ResultID             string   `protobuf:"bytes,1,opt,name=ResultID,proto3" json:"ResultID,omitempty"`
	TimeFinish           int64    `protobuf:"varint,2,opt,name=TimeFinish,proto3" json:"TimeFinish,omitempty"`
	Version              uint32   `protobuf:"varint,255,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultFinishedEvent) Reset()         { *m = ResultFinishedEvent{} }
func (m *ResultFinishedEvent) String() string { return proto.CompactTextString(m) }
func (*ResultFinishedEvent) ProtoMessage()    {}
func (*ResultFinishedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}
func (m *ResultFinishedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResultFinishedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResultFinishedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResultFinishedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultFinishedEvent.Merge(m, src)
}
func (m *ResultFinishedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ResultFinishedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultFinishedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ResultFinishedEvent proto.InternalMessageInfo

func (m *ResultFinishedEvent) GetResultID() string {
	if m != nil {
		return m.ResultID
	}
	return ""
}

func (m *ResultFinishedEvent) GetTimeFinish() int64 {
	if m != nil {
		return m.TimeFinish
	}
	return 0
}

func (m *ResultFinishedEvent) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*ResultCreatedEvent)(nil), "result.ResultCreatedEvent")
	proto.RegisterType((*ResultFinishedEvent)(nil), "result.ResultFinishedEvent")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x4b, 0xcd,
	0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29,
	0x51, 0x5a, 0xcf, 0xc8, 0x25, 0x14, 0x04, 0x66, 0x3a, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0xb8,
	0x82, 0x54, 0x09, 0x49, 0x71, 0x71, 0x40, 0x44, 0x3d, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xe0, 0x7c, 0x21, 0x25, 0x2e, 0x1e, 0xe7, 0x8c, 0xd4, 0xe4, 0xec, 0x82, 0xfc, 0xcc, 0x3c,
	0x90, 0x3c, 0x13, 0x58, 0x1e, 0x45, 0x4c, 0x48, 0x81, 0x8b, 0x3b, 0xb8, 0x20, 0xbf, 0xa8, 0xa4,
	0x38, 0x37, 0x35, 0xcf, 0xd3, 0x45, 0x82, 0x19, 0xac, 0x04, 0x59, 0x48, 0x48, 0x86, 0x8b, 0x33,
	0x24, 0x33, 0x37, 0x35, 0xb8, 0x24, 0xb1, 0xa8, 0x44, 0x82, 0x45, 0x81, 0x51, 0x83, 0x39, 0x08,
	0x21, 0x20, 0x24, 0xc9, 0xc5, 0x1e, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27, 0xf1, 0x1f, 0x64,
	0x3f, 0x6f, 0x10, 0x8c, 0xaf, 0x94, 0xc3, 0x25, 0x0c, 0x71, 0x8a, 0x5b, 0x66, 0x5e, 0x66, 0x71,
	0x06, 0x31, 0x2e, 0x96, 0xe3, 0xe2, 0x02, 0x19, 0x0d, 0xd1, 0x00, 0x76, 0x2f, 0x73, 0x10, 0x92,
	0x08, 0x1e, 0xdb, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x01, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x96, 0xde, 0xf5, 0x6d, 0x50, 0x01, 0x00, 0x00,
}

func (m *ResultCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResultCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResultCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Version != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0xf
		i--
		dAtA[i] = 0xf8
	}
	if m.TimeStart != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TimeStart))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SportsmenID) > 0 {
		i -= len(m.SportsmenID)
		copy(dAtA[i:], m.SportsmenID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SportsmenID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CheckpointID) > 0 {
		i -= len(m.CheckpointID)
		copy(dAtA[i:], m.CheckpointID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.CheckpointID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ResultID) > 0 {
		i -= len(m.ResultID)
		copy(dAtA[i:], m.ResultID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ResultID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResultFinishedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResultFinishedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResultFinishedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Version != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0xf
		i--
		dAtA[i] = 0xf8
	}
	if m.TimeFinish != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TimeFinish))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ResultID) > 0 {
		i -= len(m.ResultID)
		copy(dAtA[i:], m.ResultID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ResultID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResultCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResultID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.CheckpointID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SportsmenID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.TimeStart != 0 {
		n += 1 + sovEvents(uint64(m.TimeStart))
	}
	if m.Version != 0 {
		n += 2 + sovEvents(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResultFinishedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResultID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.TimeFinish != 0 {
		n += 1 + sovEvents(uint64(m.TimeFinish))
	}
	if m.Version != 0 {
		n += 2 + sovEvents(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResultCreatedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ResultCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckpointID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CheckpointID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SportsmenID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SportsmenID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStart", wireType)
			}
			m.TimeStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStart |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 255:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *ResultFinishedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ResultFinishedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultFinishedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeFinish", wireType)
			}
			m.TimeFinish = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeFinish |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 255:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
