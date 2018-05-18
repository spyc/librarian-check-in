// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: attendance.proto

/*
	Package modals is a generated protocol buffer package.

	It is generated from these files:
		attendance.proto

	It has these top-level messages:
		Attendance
*/
package modals

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Rank int32

const (
	Rank_fail Rank = 0
	Rank_fair Rank = 1
	Rank_good Rank = 2
)

var Rank_name = map[int32]string{
	0: "fail",
	1: "fair",
	2: "good",
}
var Rank_value = map[string]int32{
	"fail": 0,
	"fair": 1,
	"good": 2,
}

func (x Rank) String() string {
	return proto.EnumName(Rank_name, int32(x))
}
func (Rank) EnumDescriptor() ([]byte, []int) { return fileDescriptorAttendance, []int{0} }

type Attendance struct {
	Pycid    string     `protobuf:"bytes,1,opt,name=pycid,proto3" json:"pycid,omitempty"`
	CheckIn  *time.Time `protobuf:"bytes,2,opt,name=check_in,json=checkIn,stdtime" json:"check_in,omitempty"`
	CheckOut *time.Time `protobuf:"bytes,3,opt,name=check_out,json=checkOut,stdtime" json:"check_out,omitempty"`
	Rank     Rank       `protobuf:"varint,4,opt,name=rank,proto3,enum=attendance.Rank" json:"rank,omitempty"`
}

func (m *Attendance) Reset()                    { *m = Attendance{} }
func (m *Attendance) String() string            { return proto.CompactTextString(m) }
func (*Attendance) ProtoMessage()               {}
func (*Attendance) Descriptor() ([]byte, []int) { return fileDescriptorAttendance, []int{0} }

func (m *Attendance) GetPycid() string {
	if m != nil {
		return m.Pycid
	}
	return ""
}

func (m *Attendance) GetCheckIn() *time.Time {
	if m != nil {
		return m.CheckIn
	}
	return nil
}

func (m *Attendance) GetCheckOut() *time.Time {
	if m != nil {
		return m.CheckOut
	}
	return nil
}

func (m *Attendance) GetRank() Rank {
	if m != nil {
		return m.Rank
	}
	return Rank_fail
}

func init() {
	proto.RegisterType((*Attendance)(nil), "attendance.Attendance")
	proto.RegisterEnum("attendance.Rank", Rank_name, Rank_value)
}
func (this *Attendance) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Attendance)
	if !ok {
		that2, ok := that.(Attendance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Pycid != that1.Pycid {
		return false
	}
	if that1.CheckIn == nil {
		if this.CheckIn != nil {
			return false
		}
	} else if !this.CheckIn.Equal(*that1.CheckIn) {
		return false
	}
	if that1.CheckOut == nil {
		if this.CheckOut != nil {
			return false
		}
	} else if !this.CheckOut.Equal(*that1.CheckOut) {
		return false
	}
	if this.Rank != that1.Rank {
		return false
	}
	return true
}
func (m *Attendance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attendance) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Pycid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAttendance(dAtA, i, uint64(len(m.Pycid)))
		i += copy(dAtA[i:], m.Pycid)
	}
	if m.CheckIn != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAttendance(dAtA, i, uint64(types.SizeOfStdTime(*m.CheckIn)))
		n1, err := types.StdTimeMarshalTo(*m.CheckIn, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.CheckOut != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAttendance(dAtA, i, uint64(types.SizeOfStdTime(*m.CheckOut)))
		n2, err := types.StdTimeMarshalTo(*m.CheckOut, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Rank != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAttendance(dAtA, i, uint64(m.Rank))
	}
	return i, nil
}

func encodeVarintAttendance(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAttendance(r randyAttendance, easy bool) *Attendance {
	this := &Attendance{}
	this.Pycid = string(randStringAttendance(r))
	if r.Intn(10) != 0 {
		this.CheckIn = types.NewPopulatedStdTime(r, easy)
	}
	if r.Intn(10) != 0 {
		this.CheckOut = types.NewPopulatedStdTime(r, easy)
	}
	this.Rank = Rank([]int32{0, 1, 2}[r.Intn(3)])
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAttendance interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAttendance(r randyAttendance) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAttendance(r randyAttendance) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneAttendance(r)
	}
	return string(tmps)
}
func randUnrecognizedAttendance(r randyAttendance, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAttendance(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAttendance(dAtA []byte, r randyAttendance, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAttendance(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateAttendance(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateAttendance(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAttendance(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAttendance(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAttendance(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAttendance(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Attendance) Size() (n int) {
	var l int
	_ = l
	l = len(m.Pycid)
	if l > 0 {
		n += 1 + l + sovAttendance(uint64(l))
	}
	if m.CheckIn != nil {
		l = types.SizeOfStdTime(*m.CheckIn)
		n += 1 + l + sovAttendance(uint64(l))
	}
	if m.CheckOut != nil {
		l = types.SizeOfStdTime(*m.CheckOut)
		n += 1 + l + sovAttendance(uint64(l))
	}
	if m.Rank != 0 {
		n += 1 + sovAttendance(uint64(m.Rank))
	}
	return n
}

func sovAttendance(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAttendance(x uint64) (n int) {
	return sovAttendance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attendance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttendance
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
			return fmt.Errorf("proto: Attendance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attendance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pycid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttendance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAttendance
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pycid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttendance
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
				return ErrInvalidLengthAttendance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CheckIn == nil {
				m.CheckIn = new(time.Time)
			}
			if err := types.StdTimeUnmarshal(m.CheckIn, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttendance
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
				return ErrInvalidLengthAttendance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CheckOut == nil {
				m.CheckOut = new(time.Time)
			}
			if err := types.StdTimeUnmarshal(m.CheckOut, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttendance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= (Rank(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAttendance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAttendance
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
func skipAttendance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttendance
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
					return 0, ErrIntOverflowAttendance
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
					return 0, ErrIntOverflowAttendance
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
				return 0, ErrInvalidLengthAttendance
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAttendance
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
				next, err := skipAttendance(dAtA[start:])
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
	ErrInvalidLengthAttendance = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttendance   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("attendance.proto", fileDescriptorAttendance) }

var fileDescriptorAttendance = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x29, 0x49,
	0xcd, 0x4b, 0x49, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7,
	0xe7, 0xeb, 0x83, 0x95, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x2a, 0x65,
	0x8d, 0xa2, 0x3c, 0x27, 0x31, 0x2f, 0x1d, 0xa1, 0xa1, 0xa0, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xbf,
	0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x00, 0xc1, 0x82, 0x68, 0x56, 0x3a, 0xc2, 0xc8,
	0xc5, 0xe5, 0x08, 0xb7, 0x5a, 0x48, 0x84, 0x8b, 0xb5, 0xa0, 0x32, 0x39, 0x33, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0xb2, 0xe6, 0xe2, 0x48, 0xce, 0x48, 0x4d, 0xce, 0x8e,
	0xcf, 0xcc, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf,
	0x81, 0xba, 0x3e, 0xa9, 0x34, 0x4d, 0x2f, 0x04, 0x66, 0xb0, 0x13, 0xcb, 0x84, 0xfb, 0xf2, 0x8c,
	0x41, 0xec, 0x60, 0x1d, 0x9e, 0x79, 0x42, 0xb6, 0x5c, 0x9c, 0x10, 0xcd, 0xf9, 0xa5, 0x25, 0x12,
	0xcc, 0x44, 0xea, 0x86, 0xd8, 0xe7, 0x5f, 0x5a, 0x22, 0xa4, 0xc2, 0xc5, 0x52, 0x94, 0x98, 0x97,
	0x2d, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xa0, 0x87, 0x14, 0x72, 0x41, 0x89, 0x79, 0xd9,
	0x41, 0x60, 0x59, 0x2d, 0x15, 0x2e, 0x16, 0x10, 0x4f, 0x88, 0x83, 0x8b, 0x25, 0x2d, 0x31, 0x33,
	0x47, 0x80, 0x01, 0xca, 0x2a, 0x12, 0x60, 0x04, 0xb1, 0xd2, 0xf3, 0xf3, 0x53, 0x04, 0x98, 0x9c,
	0x14, 0x7e, 0x3c, 0x94, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0x71, 0xc7, 0x23, 0x39, 0xc6, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x8a, 0x2d, 0x37, 0x3f, 0x25,
	0x31, 0xa7, 0x38, 0x89, 0x0d, 0xec, 0x22, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x98,
	0xf4, 0xe2, 0xa1, 0x01, 0x00, 0x00,
}
