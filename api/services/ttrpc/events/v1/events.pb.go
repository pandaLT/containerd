// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/ttrpc/events/v1/events.proto

package events

import (
	context "context"
	fmt "fmt"
	_ "github.com/containerd/containerd/protobuf/plugin"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	github_com_containerd_typeurl "github.com/containerd/typeurl"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type ForwardRequest struct {
	Envelope             *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ForwardRequest) Reset()      { *m = ForwardRequest{} }
func (*ForwardRequest) ProtoMessage() {}
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19f98672016720b5, []int{0}
}
func (m *ForwardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForwardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardRequest.Merge(m, src)
}
func (m *ForwardRequest) XXX_Size() int {
	return m.Size()
}
func (m *ForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardRequest proto.InternalMessageInfo

type Envelope struct {
	Timestamp            *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Namespace            string           `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Topic                string           `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                *types.Any       `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Envelope) Reset()      { *m = Envelope{} }
func (*Envelope) ProtoMessage() {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_19f98672016720b5, []int{1}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return m.Size()
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ForwardRequest)(nil), "containerd.services.events.ttrpc.v1.ForwardRequest")
	proto.RegisterType((*Envelope)(nil), "containerd.services.events.ttrpc.v1.Envelope")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/ttrpc/events/v1/events.proto", fileDescriptor_19f98672016720b5)
}

var fileDescriptor_19f98672016720b5 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x0e, 0xd2, 0x40,
	0x10, 0x65, 0x15, 0x10, 0xd6, 0xc4, 0xc3, 0x86, 0x98, 0x5a, 0x4d, 0x21, 0x78, 0x21, 0x26, 0xee,
	0x06, 0xb8, 0x18, 0xbd, 0xa8, 0x11, 0x13, 0x3d, 0x36, 0xc6, 0x83, 0x26, 0xc6, 0xa5, 0x0c, 0x65,
	0x93, 0x76, 0x77, 0x6d, 0xb7, 0x35, 0xdc, 0xf8, 0x1e, 0xff, 0xc0, 0x3f, 0xe0, 0xe8, 0xd1, 0xa3,
	0xf4, 0x4b, 0x0c, 0xdb, 0x96, 0x2a, 0x1c, 0x24, 0xf1, 0xf6, 0x3a, 0xef, 0xbd, 0x79, 0x33, 0xd3,
	0xc5, 0x6f, 0x43, 0x61, 0x36, 0xd9, 0x92, 0x06, 0x2a, 0x66, 0x81, 0x92, 0x86, 0x0b, 0x09, 0xc9,
	0xea, 0x4f, 0xc8, 0xb5, 0x60, 0x29, 0x24, 0xb9, 0x08, 0x20, 0x65, 0xc6, 0x24, 0x3a, 0x60, 0x90,
	0x83, 0x34, 0x29, 0xcb, 0xa7, 0x15, 0xa2, 0x3a, 0x51, 0x46, 0x91, 0x87, 0x8d, 0x8b, 0xd6, 0x0e,
	0x5a, 0x29, 0xac, 0x91, 0xe6, 0x53, 0xf7, 0xf9, 0x3f, 0x03, 0x6d, 0xb3, 0x65, 0xb6, 0x66, 0x3a,
	0xca, 0x42, 0x21, 0xd9, 0x5a, 0x40, 0xb4, 0xd2, 0xdc, 0x6c, 0xca, 0x18, 0x77, 0x10, 0xaa, 0x50,
	0x59, 0xc8, 0x8e, 0xa8, 0xaa, 0xde, 0x0b, 0x95, 0x0a, 0x23, 0x68, 0xdc, 0x5c, 0x6e, 0x2b, 0xea,
	0xfe, 0x39, 0x05, 0xb1, 0x36, 0x35, 0x39, 0x3c, 0x27, 0x8d, 0x88, 0x21, 0x35, 0x3c, 0xd6, 0xa5,
	0x60, 0xfc, 0x11, 0xdf, 0x79, 0xad, 0x92, 0xaf, 0x3c, 0x59, 0xf9, 0xf0, 0x25, 0x83, 0xd4, 0x90,
	0x37, 0xb8, 0x07, 0x32, 0x87, 0x48, 0x69, 0x70, 0xd0, 0x08, 0x4d, 0x6e, 0xcf, 0x1e, 0xd3, 0x2b,
	0x56, 0xa7, 0x8b, 0xca, 0xe4, 0x9f, 0xec, 0xe3, 0x6f, 0x08, 0xf7, 0xea, 0x32, 0x79, 0x82, 0xfb,
	0xa7, 0xf0, 0xaa, 0xb1, 0x4b, 0xcb, 0xf1, 0x68, 0x3d, 0x1e, 0x7d, 0x57, 0x2b, 0xfc, 0x46, 0x4c,
	0x1e, 0xe0, 0xbe, 0xe4, 0x31, 0xa4, 0x9a, 0x07, 0xe0, 0xdc, 0x18, 0xa1, 0x49, 0xdf, 0x6f, 0x0a,
	0x64, 0x80, 0x3b, 0x46, 0x69, 0x11, 0x38, 0x37, 0x2d, 0x53, 0x7e, 0x90, 0x47, 0xb8, 0x63, 0x07,
	0x74, 0xda, 0x36, 0x69, 0x70, 0x91, 0xf4, 0x42, 0x6e, 0xfd, 0x52, 0xf2, 0xb4, 0xbd, 0xfb, 0x3e,
	0x44, 0xb3, 0xcf, 0xb8, 0xbb, 0xb0, 0x2b, 0x91, 0xf7, 0xf8, 0x56, 0x75, 0x13, 0x32, 0xbf, 0x6a,
	0xf5, 0xbf, 0x2f, 0xe8, 0xde, 0xbd, 0x08, 0x5b, 0x1c, 0x7f, 0xc9, 0xcb, 0x4f, 0xfb, 0x83, 0xd7,
	0xfa, 0x79, 0xf0, 0x5a, 0xbb, 0xc2, 0x43, 0xfb, 0xc2, 0x43, 0x3f, 0x0a, 0x0f, 0xfd, 0x2a, 0x3c,
	0xf4, 0xe1, 0xd5, 0x7f, 0xbd, 0xd3, 0x67, 0x25, 0x5a, 0x76, 0x6d, 0xde, 0xfc, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x87, 0xc7, 0x48, 0x3b, 0xf6, 0x02, 0x00, 0x00,
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *Envelope) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	// unhandled: timestamp
	case "namespace":
		return string(m.Namespace), len(m.Namespace) > 0
	case "topic":
		return string(m.Topic), len(m.Topic) > 0
	case "event":
		decoded, err := github_com_containerd_typeurl.UnmarshalAny(m.Event)
		if err != nil {
			return "", false
		}

		adaptor, ok := decoded.(interface{ Field([]string) (string, bool) })
		if !ok {
			return "", false
		}
		return adaptor.Field(fieldpath[1:])
	}
	return "", false
}
func (m *ForwardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForwardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Envelope != nil {
		{
			size, err := m.Envelope.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Envelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
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
func (m *ForwardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Envelope != nil {
		l = m.Envelope.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Envelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovEvents(uint64(l))
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
func (this *ForwardRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ForwardRequest{`,
		`Envelope:` + strings.Replace(this.Envelope.String(), "Envelope", "Envelope", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Envelope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Envelope{`,
		`Timestamp:` + strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Timestamp", "types.Timestamp", 1) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Topic:` + fmt.Sprintf("%v", this.Topic) + `,`,
		`Event:` + strings.Replace(fmt.Sprintf("%v", this.Event), "Any", "types.Any", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEvents(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

type EventsService interface {
	Forward(ctx context.Context, req *ForwardRequest) (*types.Empty, error)
}

func RegisterEventsService(srv *github_com_containerd_ttrpc.Server, svc EventsService) {
	srv.Register("containerd.services.events.ttrpc.v1.Events", map[string]github_com_containerd_ttrpc.Method{
		"Forward": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req ForwardRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Forward(ctx, &req)
		},
	})
}

type eventsClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewEventsClient(client *github_com_containerd_ttrpc.Client) EventsService {
	return &eventsClient{
		client: client,
	}
}

func (c *eventsClient) Forward(ctx context.Context, req *ForwardRequest) (*types.Empty, error) {
	var resp types.Empty
	if err := c.client.Call(ctx, "containerd.services.events.ttrpc.v1.Events", "Forward", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (m *ForwardRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ForwardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Envelope == nil {
				m.Envelope = &Envelope{}
			}
			if err := m.Envelope.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *Envelope) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
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
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
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
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &types.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
