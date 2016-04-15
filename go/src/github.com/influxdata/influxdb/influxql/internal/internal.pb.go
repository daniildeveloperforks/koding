// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package influxql is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Point
	Aux
	IteratorOptions
	Measurements
	Measurement
	Interval
	IteratorStats
	Series
	SeriesList
*/
package influxql

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Point struct {
	Name             *string        `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Tags             *string        `protobuf:"bytes,2,req,name=Tags" json:"Tags,omitempty"`
	Time             *int64         `protobuf:"varint,3,req,name=Time" json:"Time,omitempty"`
	Nil              *bool          `protobuf:"varint,4,req,name=Nil" json:"Nil,omitempty"`
	Aux              []*Aux         `protobuf:"bytes,5,rep,name=Aux" json:"Aux,omitempty"`
	Aggregated       *uint32        `protobuf:"varint,6,opt,name=Aggregated" json:"Aggregated,omitempty"`
	FloatValue       *float64       `protobuf:"fixed64,7,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64         `protobuf:"varint,8,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string        `protobuf:"bytes,9,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool          `protobuf:"varint,10,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	Stats            *IteratorStats `protobuf:"bytes,11,opt,name=Stats" json:"Stats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *Point) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Point) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *Point) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Point) GetNil() bool {
	if m != nil && m.Nil != nil {
		return *m.Nil
	}
	return false
}

func (m *Point) GetAux() []*Aux {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *Point) GetAggregated() uint32 {
	if m != nil && m.Aggregated != nil {
		return *m.Aggregated
	}
	return 0
}

func (m *Point) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Point) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Point) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Point) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Point) GetStats() *IteratorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Aux struct {
	DataType         *int32   `protobuf:"varint,1,req,name=DataType" json:"DataType,omitempty"`
	FloatValue       *float64 `protobuf:"fixed64,2,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64   `protobuf:"varint,3,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string  `protobuf:"bytes,4,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool    `protobuf:"varint,5,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Aux) Reset()                    { *m = Aux{} }
func (m *Aux) String() string            { return proto.CompactTextString(m) }
func (*Aux) ProtoMessage()               {}
func (*Aux) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Aux) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *Aux) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Aux) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Aux) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Aux) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

type IteratorOptions struct {
	Expr             *string        `protobuf:"bytes,1,opt,name=Expr" json:"Expr,omitempty"`
	Aux              []string       `protobuf:"bytes,2,rep,name=Aux" json:"Aux,omitempty"`
	Sources          []*Measurement `protobuf:"bytes,3,rep,name=Sources" json:"Sources,omitempty"`
	Interval         *Interval      `protobuf:"bytes,4,opt,name=Interval" json:"Interval,omitempty"`
	Dimensions       []string       `protobuf:"bytes,5,rep,name=Dimensions" json:"Dimensions,omitempty"`
	Fill             *int32         `protobuf:"varint,6,opt,name=Fill" json:"Fill,omitempty"`
	FillValue        *float64       `protobuf:"fixed64,7,opt,name=FillValue" json:"FillValue,omitempty"`
	Condition        *string        `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
	StartTime        *int64         `protobuf:"varint,9,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,10,opt,name=EndTime" json:"EndTime,omitempty"`
	Ascending        *bool          `protobuf:"varint,11,opt,name=Ascending" json:"Ascending,omitempty"`
	Limit            *int64         `protobuf:"varint,12,opt,name=Limit" json:"Limit,omitempty"`
	Offset           *int64         `protobuf:"varint,13,opt,name=Offset" json:"Offset,omitempty"`
	SLimit           *int64         `protobuf:"varint,14,opt,name=SLimit" json:"SLimit,omitempty"`
	SOffset          *int64         `protobuf:"varint,15,opt,name=SOffset" json:"SOffset,omitempty"`
	Dedupe           *bool          `protobuf:"varint,16,opt,name=Dedupe" json:"Dedupe,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *IteratorOptions) Reset()                    { *m = IteratorOptions{} }
func (m *IteratorOptions) String() string            { return proto.CompactTextString(m) }
func (*IteratorOptions) ProtoMessage()               {}
func (*IteratorOptions) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *IteratorOptions) GetExpr() string {
	if m != nil && m.Expr != nil {
		return *m.Expr
	}
	return ""
}

func (m *IteratorOptions) GetAux() []string {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *IteratorOptions) GetSources() []*Measurement {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *IteratorOptions) GetInterval() *Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *IteratorOptions) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *IteratorOptions) GetFill() int32 {
	if m != nil && m.Fill != nil {
		return *m.Fill
	}
	return 0
}

func (m *IteratorOptions) GetFillValue() float64 {
	if m != nil && m.FillValue != nil {
		return *m.FillValue
	}
	return 0
}

func (m *IteratorOptions) GetCondition() string {
	if m != nil && m.Condition != nil {
		return *m.Condition
	}
	return ""
}

func (m *IteratorOptions) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *IteratorOptions) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *IteratorOptions) GetAscending() bool {
	if m != nil && m.Ascending != nil {
		return *m.Ascending
	}
	return false
}

func (m *IteratorOptions) GetLimit() int64 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *IteratorOptions) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IteratorOptions) GetSLimit() int64 {
	if m != nil && m.SLimit != nil {
		return *m.SLimit
	}
	return 0
}

func (m *IteratorOptions) GetSOffset() int64 {
	if m != nil && m.SOffset != nil {
		return *m.SOffset
	}
	return 0
}

func (m *IteratorOptions) GetDedupe() bool {
	if m != nil && m.Dedupe != nil {
		return *m.Dedupe
	}
	return false
}

type Measurements struct {
	Items            []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Measurements) Reset()                    { *m = Measurements{} }
func (m *Measurements) String() string            { return proto.CompactTextString(m) }
func (*Measurements) ProtoMessage()               {}
func (*Measurements) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

func (m *Measurements) GetItems() []*Measurement {
	if m != nil {
		return m.Items
	}
	return nil
}

type Measurement struct {
	Database         *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy  *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex            *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget         *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Measurement) Reset()                    { *m = Measurement{} }
func (m *Measurement) String() string            { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()               {}
func (*Measurement) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Measurement) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Measurement) GetRetentionPolicy() string {
	if m != nil && m.RetentionPolicy != nil {
		return *m.RetentionPolicy
	}
	return ""
}

func (m *Measurement) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Measurement) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Measurement) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

type Interval struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=Duration" json:"Duration,omitempty"`
	Offset           *int64 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Interval) Reset()                    { *m = Interval{} }
func (m *Interval) String() string            { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()               {}
func (*Interval) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

func (m *Interval) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Interval) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

type IteratorStats struct {
	SeriesN          *int64 `protobuf:"varint,1,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN           *int64 `protobuf:"varint,2,opt,name=PointN" json:"PointN,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IteratorStats) Reset()                    { *m = IteratorStats{} }
func (m *IteratorStats) String() string            { return proto.CompactTextString(m) }
func (*IteratorStats) ProtoMessage()               {}
func (*IteratorStats) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

func (m *IteratorStats) GetSeriesN() int64 {
	if m != nil && m.SeriesN != nil {
		return *m.SeriesN
	}
	return 0
}

func (m *IteratorStats) GetPointN() int64 {
	if m != nil && m.PointN != nil {
		return *m.PointN
	}
	return 0
}

type Series struct {
	Name             *string  `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Tags             []byte   `protobuf:"bytes,2,opt,name=Tags" json:"Tags,omitempty"`
	Aux              []uint32 `protobuf:"varint,3,rep,name=Aux" json:"Aux,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Series) Reset()                    { *m = Series{} }
func (m *Series) String() string            { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()               {}
func (*Series) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

func (m *Series) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Series) GetTags() []byte {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Series) GetAux() []uint32 {
	if m != nil {
		return m.Aux
	}
	return nil
}

type SeriesList struct {
	Items            []*Series `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SeriesList) Reset()                    { *m = SeriesList{} }
func (m *SeriesList) String() string            { return proto.CompactTextString(m) }
func (*SeriesList) ProtoMessage()               {}
func (*SeriesList) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{8} }

func (m *SeriesList) GetItems() []*Series {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "influxql.Point")
	proto.RegisterType((*Aux)(nil), "influxql.Aux")
	proto.RegisterType((*IteratorOptions)(nil), "influxql.IteratorOptions")
	proto.RegisterType((*Measurements)(nil), "influxql.Measurements")
	proto.RegisterType((*Measurement)(nil), "influxql.Measurement")
	proto.RegisterType((*Interval)(nil), "influxql.Interval")
	proto.RegisterType((*IteratorStats)(nil), "influxql.IteratorStats")
	proto.RegisterType((*Series)(nil), "influxql.Series")
	proto.RegisterType((*SeriesList)(nil), "influxql.SeriesList")
}

var fileDescriptorInternal = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0xdf, 0x6e, 0xda, 0x3c,
	0x14, 0x57, 0x48, 0x53, 0x92, 0x13, 0x52, 0xf8, 0xfc, 0x6d, 0x22, 0xda, 0xcd, 0x50, 0x34, 0x4d,
	0x5c, 0x6c, 0x6c, 0x43, 0x7b, 0x01, 0x36, 0x5a, 0x09, 0xa9, 0xa3, 0x55, 0x41, 0xbb, 0xf7, 0xc0,
	0x44, 0x96, 0x4c, 0xcc, 0x6c, 0x67, 0xa2, 0x8f, 0xb8, 0x67, 0xd9, 0x4b, 0xec, 0xd8, 0x49, 0x0a,
	0x54, 0xec, 0x2e, 0xe7, 0x77, 0x8e, 0x73, 0x7e, 0x7f, 0x6c, 0xe8, 0xf3, 0xc2, 0x30, 0x55, 0x50,
	0xf1, 0xa1, 0xf9, 0x18, 0xed, 0x94, 0x34, 0x92, 0x84, 0xbc, 0xd8, 0x88, 0x72, 0xff, 0x53, 0x64,
	0x7f, 0x3c, 0x08, 0xee, 0x25, 0xb6, 0x49, 0x07, 0x2e, 0xe6, 0x74, 0xcb, 0x52, 0x6f, 0xd0, 0x1a,
	0x46, 0xb6, 0x5a, 0xd2, 0x5c, 0xa7, 0xad, 0xa7, 0x8a, 0x63, 0xcf, 0xc7, 0xca, 0x27, 0x31, 0xf8,
	0x73, 0x2e, 0xd2, 0x0b, 0x2c, 0x42, 0xf2, 0x0a, 0xfc, 0x49, 0xb9, 0x4f, 0x83, 0x81, 0x3f, 0x8c,
	0xc7, 0xc9, 0xa8, 0xf9, 0xf1, 0x08, 0x41, 0x42, 0x00, 0x26, 0x79, 0xae, 0x58, 0x4e, 0x0d, 0x5b,
	0xa7, 0x97, 0x03, 0x6f, 0x98, 0x58, 0xec, 0x46, 0x48, 0x6a, 0xbe, 0x53, 0x51, 0xb2, 0xb4, 0x8d,
	0x98, 0x47, 0x5e, 0x40, 0x67, 0x86, 0x04, 0x73, 0xa6, 0x2a, 0x34, 0x44, 0xd4, 0x27, 0xff, 0x43,
	0xbc, 0x30, 0x8a, 0x17, 0x79, 0x05, 0x46, 0x08, 0x46, 0x76, 0xf4, 0x8b, 0x94, 0x82, 0xd1, 0xa2,
	0x42, 0x01, 0xd1, 0x90, 0xbc, 0x85, 0x60, 0x61, 0xa8, 0xd1, 0x69, 0x8c, 0x65, 0x3c, 0xee, 0x1f,
	0x68, 0xcc, 0x50, 0x37, 0x35, 0x52, 0xb9, 0x76, 0x26, 0x1c, 0x59, 0xd2, 0x83, 0x70, 0x4a, 0x0d,
	0x5d, 0x3e, 0xee, 0x2a, 0xb9, 0xc1, 0x33, 0x56, 0xad, 0xb3, 0xac, 0xfc, 0x73, 0xac, 0x2e, 0xce,
	0xb2, 0x0a, 0x2c, 0xab, 0xec, 0x77, 0x0b, 0xba, 0xcd, 0xfe, 0xbb, 0x9d, 0xe1, 0xb2, 0xd0, 0xd6,
	0xc9, 0xeb, 0xfd, 0x4e, 0xe1, 0x5a, 0x7b, 0x2e, 0xae, 0xcc, 0x6b, 0xa1, 0x79, 0x11, 0x8a, 0x68,
	0x2f, 0x64, 0xa9, 0x56, 0x4c, 0xe3, 0x2a, 0xeb, 0xe6, 0xcb, 0x83, 0x8c, 0x6f, 0x8c, 0xea, 0x52,
	0xb1, 0x2d, 0xc3, 0xa0, 0xde, 0x40, 0x68, 0x79, 0xa9, 0x5f, 0x54, 0xb8, 0xf5, 0xf1, 0x98, 0x1c,
	0xe9, 0xad, 0x3b, 0x56, 0xd1, 0x14, 0x23, 0x2b, 0xb4, 0x5d, 0xeb, 0xe2, 0x71, 0x31, 0xde, 0x70,
	0x21, 0x5c, 0x12, 0x01, 0xf9, 0x0f, 0x22, 0x5b, 0x1d, 0x07, 0x81, 0xd0, 0x57, 0x59, 0xac, 0xb9,
	0xe5, 0xea, 0x52, 0x88, 0x2c, 0x84, 0xde, 0x29, 0xe3, 0xf2, 0x8f, 0x9c, 0x05, 0x5d, 0x68, 0x5f,
	0x17, 0x6b, 0x07, 0x80, 0x03, 0x70, 0x66, 0xa2, 0x57, 0x0c, 0x0f, 0x16, 0xb9, 0x8b, 0x20, 0x24,
	0x09, 0x04, 0xb7, 0x7c, 0xcb, 0x4d, 0xda, 0x71, 0x13, 0x57, 0x70, 0x79, 0xb7, 0xd9, 0x68, 0x66,
	0xd2, 0xa4, 0xa9, 0x17, 0x55, 0xff, 0xaa, 0xf9, 0xe5, 0xa2, 0x1e, 0xe8, 0x36, 0x03, 0x53, 0xb6,
	0x2e, 0x31, 0xa0, 0x9e, 0xf3, 0xf2, 0x33, 0x74, 0x8e, 0x3c, 0xd0, 0x68, 0x42, 0x80, 0xd6, 0x6e,
	0x35, 0x1a, 0xf9, 0x6f, 0xab, 0xb2, 0x1c, 0xe2, 0x63, 0xe7, 0xea, 0xdc, 0x7f, 0x50, 0xcd, 0xea,
	0x00, 0xfa, 0xd0, 0x7d, 0x60, 0x06, 0x7b, 0x28, 0xf8, 0x5e, 0x0a, 0xbe, 0x7a, 0x74, 0xe1, 0x47,
	0x4f, 0xaf, 0xc1, 0x77, 0x15, 0xaa, 0x79, 0xc0, 0x8b, 0xb0, 0xaf, 0xe3, 0xc6, 0xff, 0xcc, 0xf4,
	0x92, 0xaa, 0x1c, 0xe9, 0x56, 0x51, 0xbf, 0x3b, 0x64, 0xe2, 0xb6, 0x94, 0x18, 0xba, 0xf5, 0xd0,
	0x7b, 0xa6, 0xde, 0xfe, 0xdc, 0xcf, 0x3e, 0x42, 0x72, 0x72, 0x2f, 0x9d, 0x7c, 0xa6, 0x38, 0xd3,
	0xf3, 0xc3, 0x09, 0xf7, 0x2a, 0xe7, 0xf5, 0x89, 0x4f, 0xe8, 0x97, 0x1b, 0x38, 0x7a, 0xa6, 0xde,
	0xc9, 0x33, 0xf5, 0x86, 0x9d, 0xe6, 0x3a, 0xd9, 0xdb, 0x93, 0x64, 0xef, 0x01, 0xaa, 0x23, 0xb7,
	0x5c, 0x1b, 0xf2, 0xfa, 0xd4, 0xaf, 0xde, 0xc1, 0xaf, 0x6a, 0xe8, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x58, 0x08, 0xa6, 0x2c, 0x04, 0x00, 0x00,
}