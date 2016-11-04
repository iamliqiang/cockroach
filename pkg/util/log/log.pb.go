// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/util/log/log.proto
// DO NOT EDIT!

/*
	Package log is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/util/log/log.proto

	It has these top-level messages:
		Entry
		FileDetails
		FileInfo
*/
package log

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Severity int32

const (
	Severity_UNKNOWN Severity = 0
	Severity_INFO    Severity = 1
	Severity_WARNING Severity = 2
	Severity_ERROR   Severity = 3
	Severity_FATAL   Severity = 4
	// NONE is the end sentinel. It is never used in pracice, and must be
	// renumbered to remain at the end if new variants are added.
	Severity_NONE Severity = 5
)

var Severity_name = map[int32]string{
	0: "UNKNOWN",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "FATAL",
	5: "NONE",
}
var Severity_value = map[string]int32{
	"UNKNOWN": 0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
	"FATAL":   4,
	"NONE":    5,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}
func (Severity) EnumDescriptor() ([]byte, []int) { return fileDescriptorLog, []int{0} }

// Entry represents a cockroach structured log entry.
type Entry struct {
	Severity Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=cockroach.util.log.Severity" json:"severity,omitempty"`
	// Nanoseconds since the epoch.
	Time      int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Goroutine int64  `protobuf:"varint,6,opt,name=goroutine,proto3" json:"goroutine,omitempty"`
	File      string `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Line      int64  `protobuf:"varint,4,opt,name=line,proto3" json:"line,omitempty"`
	Message   string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{0} }

// A FileDetails holds all of the particulars that can be parsed by the name of
// a log file.
type FileDetails struct {
	Program  string   `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
	Host     string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	UserName string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Severity Severity `protobuf:"varint,4,opt,name=severity,proto3,enum=cockroach.util.log.Severity" json:"severity,omitempty"`
	Time     int64    `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	PID      int64    `protobuf:"varint,6,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (m *FileDetails) Reset()                    { *m = FileDetails{} }
func (m *FileDetails) String() string            { return proto.CompactTextString(m) }
func (*FileDetails) ProtoMessage()               {}
func (*FileDetails) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{1} }

type FileInfo struct {
	Name         string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeBytes    int64       `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	ModTimeNanos int64       `protobuf:"varint,3,opt,name=mod_time_nanos,json=modTimeNanos,proto3" json:"mod_time_nanos,omitempty"`
	Details      FileDetails `protobuf:"bytes,4,opt,name=details" json:"details"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{2} }

func init() {
	proto.RegisterType((*Entry)(nil), "cockroach.util.log.Entry")
	proto.RegisterType((*FileDetails)(nil), "cockroach.util.log.FileDetails")
	proto.RegisterType((*FileInfo)(nil), "cockroach.util.log.FileInfo")
	proto.RegisterEnum("cockroach.util.log.Severity", Severity_name, Severity_value)
}
func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Severity != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Severity))
	}
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Time))
	}
	if len(m.File) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.File)))
		i += copy(dAtA[i:], m.File)
	}
	if m.Line != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Line))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Goroutine != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Goroutine))
	}
	return i, nil
}

func (m *FileDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Program) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Program)))
		i += copy(dAtA[i:], m.Program)
	}
	if len(m.Host) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if len(m.UserName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.UserName)))
		i += copy(dAtA[i:], m.UserName)
	}
	if m.Severity != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Severity))
	}
	if m.Time != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Time))
	}
	if m.PID != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.PID))
	}
	return i, nil
}

func (m *FileInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.SizeBytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.SizeBytes))
	}
	if m.ModTimeNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.ModTimeNanos))
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintLog(dAtA, i, uint64(m.Details.Size()))
	n1, err := m.Details.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeFixed64Log(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Log(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Entry) Size() (n int) {
	var l int
	_ = l
	if m.Severity != 0 {
		n += 1 + sovLog(uint64(m.Severity))
	}
	if m.Time != 0 {
		n += 1 + sovLog(uint64(m.Time))
	}
	l = len(m.File)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Line != 0 {
		n += 1 + sovLog(uint64(m.Line))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Goroutine != 0 {
		n += 1 + sovLog(uint64(m.Goroutine))
	}
	return n
}

func (m *FileDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.Program)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Severity != 0 {
		n += 1 + sovLog(uint64(m.Severity))
	}
	if m.Time != 0 {
		n += 1 + sovLog(uint64(m.Time))
	}
	if m.PID != 0 {
		n += 1 + sovLog(uint64(m.PID))
	}
	return n
}

func (m *FileInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovLog(uint64(m.SizeBytes))
	}
	if m.ModTimeNanos != 0 {
		n += 1 + sovLog(uint64(m.ModTimeNanos))
	}
	l = m.Details.Size()
	n += 1 + l + sovLog(uint64(l))
	return n
}

func sovLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			m.Severity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Severity |= (Severity(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			m.Line = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Line |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goroutine", wireType)
			}
			m.Goroutine = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Goroutine |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
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
func (m *FileDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: FileDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Program", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Program = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			m.Severity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Severity |= (Severity(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PID", wireType)
			}
			m.PID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
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
func (m *FileInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: FileInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModTimeNanos", wireType)
			}
			m.ModTimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModTimeNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
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
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
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
				return 0, ErrInvalidLengthLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLog
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
				next, err := skipLog(dAtA[start:])
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
	ErrInvalidLengthLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/util/log/log.proto", fileDescriptorLog) }

var fileDescriptorLog = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0xcc, 0xd6, 0x76, 0x63, 0xbf, 0xa0, 0xca, 0x5a, 0x71, 0x30, 0xd0, 0x3a, 0x51, 0xc4, 0x21,
	0xe2, 0xe0, 0x48, 0xe5, 0xc2, 0x0d, 0x25, 0x6a, 0x82, 0x22, 0xd0, 0x06, 0x2d, 0x45, 0x95, 0xb8,
	0x44, 0x6e, 0xb2, 0x75, 0x57, 0xb5, 0xbd, 0x91, 0xbd, 0x41, 0x0a, 0x5f, 0xc1, 0x47, 0xf0, 0x0b,
	0x7c, 0x00, 0xb7, 0x1c, 0x39, 0x72, 0xaa, 0xc0, 0xfc, 0x08, 0x7a, 0xeb, 0xba, 0x14, 0xc1, 0x89,
	0x83, 0xa5, 0xd9, 0xd9, 0x79, 0xfb, 0x66, 0x46, 0x86, 0xde, 0x52, 0x2d, 0xaf, 0x0a, 0x15, 0x2f,
	0x2f, 0x87, 0xeb, 0xab, 0x64, 0xb8, 0xd1, 0x32, 0x1d, 0xa6, 0x2a, 0xc1, 0x2f, 0x5a, 0x17, 0x4a,
	0x2b, 0x4a, 0x6f, 0x15, 0x11, 0xde, 0x46, 0xa9, 0x4a, 0x1e, 0xde, 0x4f, 0x54, 0xa2, 0xcc, 0xf5,
	0x10, 0x51, 0xad, 0xec, 0x7f, 0x26, 0xe0, 0x4c, 0x72, 0x5d, 0x6c, 0xe9, 0x33, 0x70, 0x4b, 0xf1,
	0x5e, 0x14, 0x52, 0x6f, 0x03, 0xd2, 0x23, 0x83, 0x83, 0xe3, 0xc3, 0xe8, 0xef, 0x67, 0xa2, 0x37,
	0x37, 0x1a, 0x7e, 0xab, 0xa6, 0x14, 0x6c, 0x2d, 0x33, 0x11, 0xec, 0xf5, 0xc8, 0xc0, 0xe2, 0x06,
	0x23, 0x77, 0x21, 0x53, 0x11, 0x58, 0x3d, 0x32, 0xf0, 0xb8, 0xc1, 0xc8, 0xa5, 0x32, 0x17, 0x81,
	0x5d, 0xeb, 0x10, 0xd3, 0x00, 0xda, 0x99, 0x28, 0xcb, 0x38, 0x11, 0x81, 0x63, 0xa4, 0xcd, 0x91,
	0x1e, 0x82, 0x97, 0xa8, 0x42, 0x6d, 0x34, 0x8e, 0xec, 0x9b, 0x91, 0xdf, 0x44, 0xff, 0x0b, 0x81,
	0xce, 0x54, 0xa6, 0xe2, 0x44, 0xe8, 0x58, 0xa6, 0x25, 0xbe, 0xb3, 0x2e, 0x54, 0x52, 0xc4, 0x99,
	0x31, 0xef, 0xf1, 0xe6, 0x88, 0x5b, 0x2f, 0x55, 0xa9, 0x8d, 0x3b, 0x8f, 0x1b, 0x4c, 0x1f, 0x81,
	0xb7, 0x29, 0x45, 0xb1, 0xc8, 0xe3, 0xac, 0xb1, 0xe8, 0x22, 0xc1, 0xe2, 0x4c, 0xfc, 0x51, 0x84,
	0xfd, 0x5f, 0x45, 0x38, 0x77, 0x8a, 0x78, 0x00, 0xd6, 0x5a, 0xae, 0xea, 0x00, 0xe3, 0x76, 0x75,
	0xdd, 0xb5, 0x5e, 0xcf, 0x4e, 0x38, 0x72, 0xfd, 0x4f, 0x04, 0x5c, 0xcc, 0x30, 0xcb, 0x2f, 0x14,
	0xce, 0x1a, 0x37, 0xb5, 0x7b, 0x83, 0xe9, 0x11, 0x40, 0x29, 0x3f, 0x88, 0xc5, 0xf9, 0x56, 0x8b,
	0xf2, 0xa6, 0x5e, 0x0f, 0x99, 0x31, 0x12, 0xf4, 0x31, 0x1c, 0x64, 0x6a, 0xb5, 0xc0, 0x35, 0x8b,
	0x3c, 0xce, 0x55, 0x69, 0xa2, 0x58, 0xfc, 0x5e, 0xa6, 0x56, 0xa7, 0x32, 0x13, 0x0c, 0x39, 0xfa,
	0x1c, 0xda, 0xab, 0xba, 0x24, 0x93, 0xa6, 0x73, 0xdc, 0xfd, 0x57, 0x9a, 0x3b, 0x5d, 0x8e, 0xed,
	0xdd, 0x75, 0xb7, 0xc5, 0x9b, 0xa9, 0x27, 0x0c, 0xdc, 0x26, 0x2b, 0xed, 0x40, 0xfb, 0x2d, 0x7b,
	0xc9, 0xe6, 0x67, 0xcc, 0x6f, 0x51, 0x17, 0xec, 0x19, 0x9b, 0xce, 0x7d, 0x82, 0xf4, 0xd9, 0x88,
	0xb3, 0x19, 0x7b, 0xe1, 0xef, 0x51, 0x0f, 0x9c, 0x09, 0xe7, 0x73, 0xee, 0x5b, 0x08, 0xa7, 0xa3,
	0xd3, 0xd1, 0x2b, 0xdf, 0x46, 0x31, 0x9b, 0xb3, 0x89, 0xef, 0x8c, 0x8f, 0x76, 0x3f, 0xc2, 0xd6,
	0xae, 0x0a, 0xc9, 0xd7, 0x2a, 0x24, 0xdf, 0xaa, 0x90, 0x7c, 0xaf, 0x42, 0xf2, 0xf1, 0x67, 0xd8,
	0x7a, 0x67, 0xa5, 0x2a, 0x39, 0xdf, 0x37, 0x3f, 0xe6, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x3e, 0x89, 0x15, 0xe6, 0x02, 0x00, 0x00,
}
