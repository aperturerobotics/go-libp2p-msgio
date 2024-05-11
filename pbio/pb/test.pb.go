// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.6.1
// source: github.com/libp2p/go-msgio/pbio/pb/test.proto

package test_pb

import (
	base64 "encoding/base64"
	fmt "fmt"
	io "io"
	strconv "strconv"
	strings "strings"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
)

type TestRecord struct {
	unknownFields []byte
	Uint32        uint32 `protobuf:"varint,1,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64        uint64 `protobuf:"varint,2,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Bytes         []byte `protobuf:"bytes,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
	String_       string `protobuf:"bytes,4,opt,name=string,proto3" json:"string,omitempty"`
	Int32         int32  `protobuf:"varint,5,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64         int64  `protobuf:"varint,6,opt,name=int64,proto3" json:"int64,omitempty"`
}

func (x *TestRecord) Reset() {
	*x = TestRecord{}
}

func (*TestRecord) ProtoMessage() {}

func (x *TestRecord) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *TestRecord) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *TestRecord) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *TestRecord) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *TestRecord) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *TestRecord) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (m *TestRecord) CloneVT() *TestRecord {
	if m == nil {
		return (*TestRecord)(nil)
	}
	r := new(TestRecord)
	r.Uint32 = m.Uint32
	r.Uint64 = m.Uint64
	r.String_ = m.String_
	r.Int32 = m.Int32
	r.Int64 = m.Int64
	if rhs := m.Bytes; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Bytes = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *TestRecord) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *TestRecord) EqualVT(that *TestRecord) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Uint32 != that.Uint32 {
		return false
	}
	if this.Uint64 != that.Uint64 {
		return false
	}
	if string(this.Bytes) != string(that.Bytes) {
		return false
	}
	if this.String_ != that.String_ {
		return false
	}
	if this.Int32 != that.Int32 {
		return false
	}
	if this.Int64 != that.Int64 {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *TestRecord) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*TestRecord)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// MarshalProtoJSON marshals the TestRecord message to JSON.
func (x *TestRecord) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Uint32 != 0 || s.HasField("uint32") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("uint32")
		s.WriteUint32(x.Uint32)
	}
	if x.Uint64 != 0 || s.HasField("uint64") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("uint64")
		s.WriteUint64(x.Uint64)
	}
	if len(x.Bytes) > 0 || s.HasField("bytes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("bytes")
		s.WriteBytes(x.Bytes)
	}
	if x.String_ != "" || s.HasField("string") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("string")
		s.WriteString(x.String_)
	}
	if x.Int32 != 0 || s.HasField("int32") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("int32")
		s.WriteInt32(x.Int32)
	}
	if x.Int64 != 0 || s.HasField("int64") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("int64")
		s.WriteInt64(x.Int64)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the TestRecord to JSON.
func (x *TestRecord) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the TestRecord message from JSON.
func (x *TestRecord) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "uint32":
			s.AddField("uint32")
			x.Uint32 = s.ReadUint32()
		case "uint64":
			s.AddField("uint64")
			x.Uint64 = s.ReadUint64()
		case "bytes":
			s.AddField("bytes")
			x.Bytes = s.ReadBytes()
		case "string":
			s.AddField("string")
			x.String_ = s.ReadString()
		case "int32":
			s.AddField("int32")
			x.Int32 = s.ReadInt32()
		case "int64":
			s.AddField("int64")
			x.Int64 = s.ReadInt64()
		}
	})
}

// UnmarshalJSON unmarshals the TestRecord from JSON.
func (x *TestRecord) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

func (m *TestRecord) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestRecord) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *TestRecord) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Int64 != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.Int64))
		i--
		dAtA[i] = 0x30
	}
	if m.Int32 != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.Int32))
		i--
		dAtA[i] = 0x28
	}
	if len(m.String_) > 0 {
		i -= len(m.String_)
		copy(dAtA[i:], m.String_)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.String_)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bytes) > 0 {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Uint64 != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.Uint64))
		i--
		dAtA[i] = 0x10
	}
	if m.Uint32 != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.Uint32))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestRecord) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uint32 != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.Uint32))
	}
	if m.Uint64 != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.Uint64))
	}
	l = len(m.Bytes)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	l = len(m.String_)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	if m.Int32 != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.Int32))
	}
	if m.Int64 != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.Int64))
	}
	n += len(m.unknownFields)
	return n
}

func (x *TestRecord) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("TestRecord { ")
	if x.Uint32 != 0 {
		sb.WriteString(" uint32: ")
		sb.WriteString(strconv.FormatUint(uint64(x.Uint32), 10))
	}
	if x.Uint64 != 0 {
		sb.WriteString(" uint64: ")
		sb.WriteString(strconv.FormatUint(uint64(x.Uint64), 10))
	}
	if len(x.Bytes) > 0 {
		sb.WriteString(" bytes: ")
		sb.WriteString("\"")
		sb.WriteString(base64.StdEncoding.EncodeToString(x.Bytes))
		sb.WriteString("\"")
	}
	if x.String_ != "" {
		sb.WriteString(" string: ")
		sb.WriteString(strconv.Quote(x.String_))
	}
	if x.Int32 != 0 {
		sb.WriteString(" int32: ")
		sb.WriteString(strconv.FormatInt(int64(x.Int32), 10))
	}
	if x.Int64 != 0 {
		sb.WriteString(" int64: ")
		sb.WriteString(strconv.FormatInt(int64(x.Int64), 10))
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *TestRecord) String() string {
	return x.MarshalProtoText()
}
func (m *TestRecord) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: TestRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint32", wireType)
			}
			m.Uint32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uint32 |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint64", wireType)
			}
			m.Uint64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uint64 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.String_ = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32", wireType)
			}
			m.Int32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int32 |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64", wireType)
			}
			m.Int64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int64 |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
