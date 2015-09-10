// Code generated by protoc-gen-gogo.
// source: fuzz3.proto
// DO NOT EDIT!

package fuzztests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NinOptNative3 struct {
	Field1  float64 `protobuf:"fixed64,1,opt,name=Field1,proto3" json:"Field1,omitempty"`
	Field2  float32 `protobuf:"fixed32,2,opt,name=Field2,proto3" json:"Field2,omitempty"`
	Field3  int32   `protobuf:"varint,3,opt,name=Field3,proto3" json:"Field3,omitempty"`
	Field4  int64   `protobuf:"varint,4,opt,name=Field4,proto3" json:"Field4,omitempty"`
	Field5  uint32  `protobuf:"varint,5,opt,name=Field5,proto3" json:"Field5,omitempty"`
	Field6  uint64  `protobuf:"varint,6,opt,name=Field6,proto3" json:"Field6,omitempty"`
	Field7  int32   `protobuf:"zigzag32,7,opt,name=Field7,proto3" json:"Field7,omitempty"`
	Field8  int64   `protobuf:"zigzag64,8,opt,name=Field8,proto3" json:"Field8,omitempty"`
	Field9  uint32  `protobuf:"fixed32,9,opt,name=Field9,proto3" json:"Field9,omitempty"`
	Field10 int32   `protobuf:"fixed32,10,opt,name=Field10,proto3" json:"Field10,omitempty"`
	Field11 uint64  `protobuf:"fixed64,11,opt,name=Field11,proto3" json:"Field11,omitempty"`
	Field12 int64   `protobuf:"fixed64,12,opt,name=Field12,proto3" json:"Field12,omitempty"`
	Field13 bool    `protobuf:"varint,13,opt,name=Field13,proto3" json:"Field13,omitempty"`
	Field14 string  `protobuf:"bytes,14,opt,name=Field14,proto3" json:"Field14,omitempty"`
	Field15 []byte  `protobuf:"bytes,15,opt,name=Field15,proto3" json:"Field15,omitempty"`
}

func (m *NinOptNative3) Reset()         { *m = NinOptNative3{} }
func (m *NinOptNative3) String() string { return proto.CompactTextString(m) }
func (*NinOptNative3) ProtoMessage()    {}

type NinRepNative3 struct {
	Field1  []float64 `protobuf:"fixed64,1,rep,name=Field1" json:"Field1,omitempty"`
	Field2  []float32 `protobuf:"fixed32,2,rep,name=Field2" json:"Field2,omitempty"`
	Field3  []int32   `protobuf:"varint,3,rep,name=Field3" json:"Field3,omitempty"`
	Field4  []int64   `protobuf:"varint,4,rep,name=Field4" json:"Field4,omitempty"`
	Field5  []uint32  `protobuf:"varint,5,rep,name=Field5" json:"Field5,omitempty"`
	Field6  []uint64  `protobuf:"varint,6,rep,name=Field6" json:"Field6,omitempty"`
	Field7  []int32   `protobuf:"zigzag32,7,rep,name=Field7" json:"Field7,omitempty"`
	Field8  []int64   `protobuf:"zigzag64,8,rep,name=Field8" json:"Field8,omitempty"`
	Field9  []uint32  `protobuf:"fixed32,9,rep,name=Field9" json:"Field9,omitempty"`
	Field10 []int32   `protobuf:"fixed32,10,rep,name=Field10" json:"Field10,omitempty"`
	Field11 []uint64  `protobuf:"fixed64,11,rep,name=Field11" json:"Field11,omitempty"`
	Field12 []int64   `protobuf:"fixed64,12,rep,name=Field12" json:"Field12,omitempty"`
	Field13 []bool    `protobuf:"varint,13,rep,name=Field13" json:"Field13,omitempty"`
	Field14 []string  `protobuf:"bytes,14,rep,name=Field14" json:"Field14,omitempty"`
	Field15 [][]byte  `protobuf:"bytes,15,rep,name=Field15" json:"Field15,omitempty"`
}

func (m *NinRepNative3) Reset()         { *m = NinRepNative3{} }
func (m *NinRepNative3) String() string { return proto.CompactTextString(m) }
func (*NinRepNative3) ProtoMessage()    {}

type NinRepPackedNative3 struct {
	Field1  []float64 `protobuf:"fixed64,1,rep,packed,name=Field1" json:"Field1,omitempty"`
	Field2  []float32 `protobuf:"fixed32,2,rep,packed,name=Field2" json:"Field2,omitempty"`
	Field3  []int32   `protobuf:"varint,3,rep,packed,name=Field3" json:"Field3,omitempty"`
	Field4  []int64   `protobuf:"varint,4,rep,packed,name=Field4" json:"Field4,omitempty"`
	Field5  []uint32  `protobuf:"varint,5,rep,packed,name=Field5" json:"Field5,omitempty"`
	Field6  []uint64  `protobuf:"varint,6,rep,packed,name=Field6" json:"Field6,omitempty"`
	Field7  []int32   `protobuf:"zigzag32,7,rep,packed,name=Field7" json:"Field7,omitempty"`
	Field8  []int64   `protobuf:"zigzag64,8,rep,packed,name=Field8" json:"Field8,omitempty"`
	Field9  []uint32  `protobuf:"fixed32,9,rep,packed,name=Field9" json:"Field9,omitempty"`
	Field10 []int32   `protobuf:"fixed32,10,rep,packed,name=Field10" json:"Field10,omitempty"`
	Field11 []uint64  `protobuf:"fixed64,11,rep,packed,name=Field11" json:"Field11,omitempty"`
	Field12 []int64   `protobuf:"fixed64,12,rep,packed,name=Field12" json:"Field12,omitempty"`
	Field13 []bool    `protobuf:"varint,13,rep,packed,name=Field13" json:"Field13,omitempty"`
}

func (m *NinRepPackedNative3) Reset()         { *m = NinRepPackedNative3{} }
func (m *NinRepPackedNative3) String() string { return proto.CompactTextString(m) }
func (*NinRepPackedNative3) ProtoMessage()    {}

type NinOptStruct3 struct {
	Field1  float64        `protobuf:"fixed64,1,opt,name=Field1,proto3" json:"Field1,omitempty"`
	Field2  float32        `protobuf:"fixed32,2,opt,name=Field2,proto3" json:"Field2,omitempty"`
	Field3  *NinOptNative3 `protobuf:"bytes,3,opt,name=Field3" json:"Field3,omitempty"`
	Field4  *NinOptNative3 `protobuf:"bytes,4,opt,name=Field4" json:"Field4,omitempty"`
	Field6  uint64         `protobuf:"varint,6,opt,name=Field6,proto3" json:"Field6,omitempty"`
	Field7  int32          `protobuf:"zigzag32,7,opt,name=Field7,proto3" json:"Field7,omitempty"`
	Field8  *NinOptNative3 `protobuf:"bytes,8,opt,name=Field8" json:"Field8,omitempty"`
	Field13 bool           `protobuf:"varint,13,opt,name=Field13,proto3" json:"Field13,omitempty"`
	Field14 string         `protobuf:"bytes,14,opt,name=Field14,proto3" json:"Field14,omitempty"`
	Field15 []byte         `protobuf:"bytes,15,opt,name=Field15,proto3" json:"Field15,omitempty"`
}

func (m *NinOptStruct3) Reset()         { *m = NinOptStruct3{} }
func (m *NinOptStruct3) String() string { return proto.CompactTextString(m) }
func (*NinOptStruct3) ProtoMessage()    {}

func (m *NinOptStruct3) GetField3() *NinOptNative3 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinOptStruct3) GetField4() *NinOptNative3 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinOptStruct3) GetField8() *NinOptNative3 {
	if m != nil {
		return m.Field8
	}
	return nil
}

type NinRepStruct3 struct {
	Field1  []float64        `protobuf:"fixed64,1,rep,name=Field1" json:"Field1,omitempty"`
	Field2  []float32        `protobuf:"fixed32,2,rep,name=Field2" json:"Field2,omitempty"`
	Field3  []*NinOptNative3 `protobuf:"bytes,3,rep,name=Field3" json:"Field3,omitempty"`
	Field4  []*NinOptNative3 `protobuf:"bytes,4,rep,name=Field4" json:"Field4,omitempty"`
	Field6  []uint64         `protobuf:"varint,6,rep,name=Field6" json:"Field6,omitempty"`
	Field7  []int32          `protobuf:"zigzag32,7,rep,name=Field7" json:"Field7,omitempty"`
	Field8  []*NinOptNative3 `protobuf:"bytes,8,rep,name=Field8" json:"Field8,omitempty"`
	Field13 []bool           `protobuf:"varint,13,rep,name=Field13" json:"Field13,omitempty"`
	Field14 []string         `protobuf:"bytes,14,rep,name=Field14" json:"Field14,omitempty"`
	Field15 [][]byte         `protobuf:"bytes,15,rep,name=Field15" json:"Field15,omitempty"`
}

func (m *NinRepStruct3) Reset()         { *m = NinRepStruct3{} }
func (m *NinRepStruct3) String() string { return proto.CompactTextString(m) }
func (*NinRepStruct3) ProtoMessage()    {}

func (m *NinRepStruct3) GetField3() []*NinOptNative3 {
	if m != nil {
		return m.Field3
	}
	return nil
}

func (m *NinRepStruct3) GetField4() []*NinOptNative3 {
	if m != nil {
		return m.Field4
	}
	return nil
}

func (m *NinRepStruct3) GetField8() []*NinOptNative3 {
	if m != nil {
		return m.Field8
	}
	return nil
}

type NinNestedStruct3 struct {
	Field1 *NinOptStruct3   `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2 []*NinRepStruct3 `protobuf:"bytes,2,rep,name=Field2" json:"Field2,omitempty"`
}

func (m *NinNestedStruct3) Reset()         { *m = NinNestedStruct3{} }
func (m *NinNestedStruct3) String() string { return proto.CompactTextString(m) }
func (*NinNestedStruct3) ProtoMessage()    {}

func (m *NinNestedStruct3) GetField1() *NinOptStruct3 {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *NinNestedStruct3) GetField2() []*NinRepStruct3 {
	if m != nil {
		return m.Field2
	}
	return nil
}

type Nil3 struct {
}

func (m *Nil3) Reset()         { *m = Nil3{} }
func (m *Nil3) String() string { return proto.CompactTextString(m) }
func (*Nil3) ProtoMessage()    {}

type NestedDefinition3 struct {
	Field1 int64                                              `protobuf:"varint,1,opt,name=Field1,proto3" json:"Field1,omitempty"`
	NNM    *NestedDefinition3_NestedMessage3_NestedNestedMsg3 `protobuf:"bytes,3,opt,name=NNM" json:"NNM,omitempty"`
	NM     *NestedDefinition3_NestedMessage3                  `protobuf:"bytes,4,opt,name=NM" json:"NM,omitempty"`
}

func (m *NestedDefinition3) Reset()         { *m = NestedDefinition3{} }
func (m *NestedDefinition3) String() string { return proto.CompactTextString(m) }
func (*NestedDefinition3) ProtoMessage()    {}

func (m *NestedDefinition3) GetNNM() *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	if m != nil {
		return m.NNM
	}
	return nil
}

func (m *NestedDefinition3) GetNM() *NestedDefinition3_NestedMessage3 {
	if m != nil {
		return m.NM
	}
	return nil
}

type NestedDefinition3_NestedMessage3 struct {
	NestedField1 uint64                                             `protobuf:"fixed64,1,opt,name=NestedField1,proto3" json:"NestedField1,omitempty"`
	NNM          *NestedDefinition3_NestedMessage3_NestedNestedMsg3 `protobuf:"bytes,2,opt,name=NNM" json:"NNM,omitempty"`
}

func (m *NestedDefinition3_NestedMessage3) Reset()         { *m = NestedDefinition3_NestedMessage3{} }
func (m *NestedDefinition3_NestedMessage3) String() string { return proto.CompactTextString(m) }
func (*NestedDefinition3_NestedMessage3) ProtoMessage()    {}

func (m *NestedDefinition3_NestedMessage3) GetNNM() *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	if m != nil {
		return m.NNM
	}
	return nil
}

type NestedDefinition3_NestedMessage3_NestedNestedMsg3 struct {
	NestedNestedField1 string `protobuf:"bytes,10,opt,name=NestedNestedField1,proto3" json:"NestedNestedField1,omitempty"`
}

func (m *NestedDefinition3_NestedMessage3_NestedNestedMsg3) Reset() {
	*m = NestedDefinition3_NestedMessage3_NestedNestedMsg3{}
}
func (m *NestedDefinition3_NestedMessage3_NestedNestedMsg3) String() string {
	return proto.CompactTextString(m)
}
func (*NestedDefinition3_NestedMessage3_NestedNestedMsg3) ProtoMessage() {}

type NestedScope3 struct {
	A *NestedDefinition3_NestedMessage3_NestedNestedMsg3 `protobuf:"bytes,1,opt,name=A" json:"A,omitempty"`
	C *NestedDefinition3_NestedMessage3                  `protobuf:"bytes,3,opt,name=C" json:"C,omitempty"`
}

func (m *NestedScope3) Reset()         { *m = NestedScope3{} }
func (m *NestedScope3) String() string { return proto.CompactTextString(m) }
func (*NestedScope3) ProtoMessage()    {}

func (m *NestedScope3) GetA() *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *NestedScope3) GetC() *NestedDefinition3_NestedMessage3 {
	if m != nil {
		return m.C
	}
	return nil
}

func NewPopulatedNinOptNative3(r randyFuzz3, easy bool) *NinOptNative3 {
	this := &NinOptNative3{}
	this.Field1 = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	this.Field2 = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	this.Field3 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field3 *= -1
	}
	this.Field4 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field4 *= -1
	}
	this.Field5 = uint32(r.Uint32())
	this.Field6 = uint64(uint64(r.Uint32()))
	this.Field7 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field7 *= -1
	}
	this.Field8 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field8 *= -1
	}
	this.Field9 = uint32(r.Uint32())
	this.Field10 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field10 *= -1
	}
	this.Field11 = uint64(uint64(r.Uint32()))
	this.Field12 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field12 *= -1
	}
	this.Field13 = bool(bool(r.Intn(2) == 0))
	this.Field14 = randStringFuzz3(r)
	v1 := r.Intn(100)
	this.Field15 = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Field15[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNinRepNative3(r randyFuzz3, easy bool) *NinRepNative3 {
	this := &NinRepNative3{}
	v2 := r.Intn(100)
	this.Field1 = make([]float64, v2)
	for i := 0; i < v2; i++ {
		this.Field1[i] = float64(r.Float64())
		if r.Intn(2) == 0 {
			this.Field1[i] *= -1
		}
	}
	v3 := r.Intn(100)
	this.Field2 = make([]float32, v3)
	for i := 0; i < v3; i++ {
		this.Field2[i] = float32(r.Float32())
		if r.Intn(2) == 0 {
			this.Field2[i] *= -1
		}
	}
	v4 := r.Intn(100)
	this.Field3 = make([]int32, v4)
	for i := 0; i < v4; i++ {
		this.Field3[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field3[i] *= -1
		}
	}
	v5 := r.Intn(100)
	this.Field4 = make([]int64, v5)
	for i := 0; i < v5; i++ {
		this.Field4[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Field4[i] *= -1
		}
	}
	v6 := r.Intn(100)
	this.Field5 = make([]uint32, v6)
	for i := 0; i < v6; i++ {
		this.Field5[i] = uint32(r.Uint32())
	}
	v7 := r.Intn(100)
	this.Field6 = make([]uint64, v7)
	for i := 0; i < v7; i++ {
		this.Field6[i] = uint64(uint64(r.Uint32()))
	}
	v8 := r.Intn(100)
	this.Field7 = make([]int32, v8)
	for i := 0; i < v8; i++ {
		this.Field7[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field7[i] *= -1
		}
	}
	v9 := r.Intn(100)
	this.Field8 = make([]int64, v9)
	for i := 0; i < v9; i++ {
		this.Field8[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Field8[i] *= -1
		}
	}
	v10 := r.Intn(100)
	this.Field9 = make([]uint32, v10)
	for i := 0; i < v10; i++ {
		this.Field9[i] = uint32(r.Uint32())
	}
	v11 := r.Intn(100)
	this.Field10 = make([]int32, v11)
	for i := 0; i < v11; i++ {
		this.Field10[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field10[i] *= -1
		}
	}
	v12 := r.Intn(100)
	this.Field11 = make([]uint64, v12)
	for i := 0; i < v12; i++ {
		this.Field11[i] = uint64(uint64(r.Uint32()))
	}
	v13 := r.Intn(100)
	this.Field12 = make([]int64, v13)
	for i := 0; i < v13; i++ {
		this.Field12[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Field12[i] *= -1
		}
	}
	v14 := r.Intn(100)
	this.Field13 = make([]bool, v14)
	for i := 0; i < v14; i++ {
		this.Field13[i] = bool(bool(r.Intn(2) == 0))
	}
	v15 := r.Intn(10)
	this.Field14 = make([]string, v15)
	for i := 0; i < v15; i++ {
		this.Field14[i] = randStringFuzz3(r)
	}
	v16 := r.Intn(100)
	this.Field15 = make([][]byte, v16)
	for i := 0; i < v16; i++ {
		v17 := r.Intn(100)
		this.Field15[i] = make([]byte, v17)
		for j := 0; j < v17; j++ {
			this.Field15[i][j] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNinRepPackedNative3(r randyFuzz3, easy bool) *NinRepPackedNative3 {
	this := &NinRepPackedNative3{}
	v18 := r.Intn(100)
	this.Field1 = make([]float64, v18)
	for i := 0; i < v18; i++ {
		this.Field1[i] = float64(r.Float64())
		if r.Intn(2) == 0 {
			this.Field1[i] *= -1
		}
	}
	v19 := r.Intn(100)
	this.Field2 = make([]float32, v19)
	for i := 0; i < v19; i++ {
		this.Field2[i] = float32(r.Float32())
		if r.Intn(2) == 0 {
			this.Field2[i] *= -1
		}
	}
	v20 := r.Intn(100)
	this.Field3 = make([]int32, v20)
	for i := 0; i < v20; i++ {
		this.Field3[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field3[i] *= -1
		}
	}
	v21 := r.Intn(100)
	this.Field4 = make([]int64, v21)
	for i := 0; i < v21; i++ {
		this.Field4[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Field4[i] *= -1
		}
	}
	v22 := r.Intn(100)
	this.Field5 = make([]uint32, v22)
	for i := 0; i < v22; i++ {
		this.Field5[i] = uint32(r.Uint32())
	}
	v23 := r.Intn(100)
	this.Field6 = make([]uint64, v23)
	for i := 0; i < v23; i++ {
		this.Field6[i] = uint64(uint64(r.Uint32()))
	}
	v24 := r.Intn(100)
	this.Field7 = make([]int32, v24)
	for i := 0; i < v24; i++ {
		this.Field7[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field7[i] *= -1
		}
	}
	v25 := r.Intn(100)
	this.Field8 = make([]int64, v25)
	for i := 0; i < v25; i++ {
		this.Field8[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Field8[i] *= -1
		}
	}
	v26 := r.Intn(100)
	this.Field9 = make([]uint32, v26)
	for i := 0; i < v26; i++ {
		this.Field9[i] = uint32(r.Uint32())
	}
	v27 := r.Intn(100)
	this.Field10 = make([]int32, v27)
	for i := 0; i < v27; i++ {
		this.Field10[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field10[i] *= -1
		}
	}
	v28 := r.Intn(100)
	this.Field11 = make([]uint64, v28)
	for i := 0; i < v28; i++ {
		this.Field11[i] = uint64(uint64(r.Uint32()))
	}
	v29 := r.Intn(100)
	this.Field12 = make([]int64, v29)
	for i := 0; i < v29; i++ {
		this.Field12[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Field12[i] *= -1
		}
	}
	v30 := r.Intn(100)
	this.Field13 = make([]bool, v30)
	for i := 0; i < v30; i++ {
		this.Field13[i] = bool(bool(r.Intn(2) == 0))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNinOptStruct3(r randyFuzz3, easy bool) *NinOptStruct3 {
	this := &NinOptStruct3{}
	this.Field1 = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	this.Field2 = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	if r.Intn(10) != 0 {
		this.Field3 = NewPopulatedNinOptNative3(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Field4 = NewPopulatedNinOptNative3(r, easy)
	}
	this.Field6 = uint64(uint64(r.Uint32()))
	this.Field7 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field7 *= -1
	}
	if r.Intn(10) != 0 {
		this.Field8 = NewPopulatedNinOptNative3(r, easy)
	}
	this.Field13 = bool(bool(r.Intn(2) == 0))
	this.Field14 = randStringFuzz3(r)
	v31 := r.Intn(100)
	this.Field15 = make([]byte, v31)
	for i := 0; i < v31; i++ {
		this.Field15[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNinRepStruct3(r randyFuzz3, easy bool) *NinRepStruct3 {
	this := &NinRepStruct3{}
	v32 := r.Intn(100)
	this.Field1 = make([]float64, v32)
	for i := 0; i < v32; i++ {
		this.Field1[i] = float64(r.Float64())
		if r.Intn(2) == 0 {
			this.Field1[i] *= -1
		}
	}
	v33 := r.Intn(100)
	this.Field2 = make([]float32, v33)
	for i := 0; i < v33; i++ {
		this.Field2[i] = float32(r.Float32())
		if r.Intn(2) == 0 {
			this.Field2[i] *= -1
		}
	}
	if r.Intn(10) != 0 {
		v34 := r.Intn(10)
		this.Field3 = make([]*NinOptNative3, v34)
		for i := 0; i < v34; i++ {
			this.Field3[i] = NewPopulatedNinOptNative3(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v35 := r.Intn(10)
		this.Field4 = make([]*NinOptNative3, v35)
		for i := 0; i < v35; i++ {
			this.Field4[i] = NewPopulatedNinOptNative3(r, easy)
		}
	}
	v36 := r.Intn(100)
	this.Field6 = make([]uint64, v36)
	for i := 0; i < v36; i++ {
		this.Field6[i] = uint64(uint64(r.Uint32()))
	}
	v37 := r.Intn(100)
	this.Field7 = make([]int32, v37)
	for i := 0; i < v37; i++ {
		this.Field7[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.Field7[i] *= -1
		}
	}
	if r.Intn(10) != 0 {
		v38 := r.Intn(10)
		this.Field8 = make([]*NinOptNative3, v38)
		for i := 0; i < v38; i++ {
			this.Field8[i] = NewPopulatedNinOptNative3(r, easy)
		}
	}
	v39 := r.Intn(100)
	this.Field13 = make([]bool, v39)
	for i := 0; i < v39; i++ {
		this.Field13[i] = bool(bool(r.Intn(2) == 0))
	}
	v40 := r.Intn(10)
	this.Field14 = make([]string, v40)
	for i := 0; i < v40; i++ {
		this.Field14[i] = randStringFuzz3(r)
	}
	v41 := r.Intn(100)
	this.Field15 = make([][]byte, v41)
	for i := 0; i < v41; i++ {
		v42 := r.Intn(100)
		this.Field15[i] = make([]byte, v42)
		for j := 0; j < v42; j++ {
			this.Field15[i][j] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNinNestedStruct3(r randyFuzz3, easy bool) *NinNestedStruct3 {
	this := &NinNestedStruct3{}
	if r.Intn(10) != 0 {
		this.Field1 = NewPopulatedNinOptStruct3(r, easy)
	}
	if r.Intn(10) != 0 {
		v43 := r.Intn(10)
		this.Field2 = make([]*NinRepStruct3, v43)
		for i := 0; i < v43; i++ {
			this.Field2[i] = NewPopulatedNinRepStruct3(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNil3(r randyFuzz3, easy bool) *Nil3 {
	this := &Nil3{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNestedDefinition3(r randyFuzz3, easy bool) *NestedDefinition3 {
	this := &NestedDefinition3{}
	this.Field1 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	if r.Intn(10) != 0 {
		this.NNM = NewPopulatedNestedDefinition3_NestedMessage3_NestedNestedMsg3(r, easy)
	}
	if r.Intn(10) != 0 {
		this.NM = NewPopulatedNestedDefinition3_NestedMessage3(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNestedDefinition3_NestedMessage3(r randyFuzz3, easy bool) *NestedDefinition3_NestedMessage3 {
	this := &NestedDefinition3_NestedMessage3{}
	this.NestedField1 = uint64(uint64(r.Uint32()))
	if r.Intn(10) != 0 {
		this.NNM = NewPopulatedNestedDefinition3_NestedMessage3_NestedNestedMsg3(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNestedDefinition3_NestedMessage3_NestedNestedMsg3(r randyFuzz3, easy bool) *NestedDefinition3_NestedMessage3_NestedNestedMsg3 {
	this := &NestedDefinition3_NestedMessage3_NestedNestedMsg3{}
	this.NestedNestedField1 = randStringFuzz3(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNestedScope3(r randyFuzz3, easy bool) *NestedScope3 {
	this := &NestedScope3{}
	if r.Intn(10) != 0 {
		this.A = NewPopulatedNestedDefinition3_NestedMessage3_NestedNestedMsg3(r, easy)
	}
	if r.Intn(10) != 0 {
		this.C = NewPopulatedNestedDefinition3_NestedMessage3(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyFuzz3 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFuzz3(r randyFuzz3) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFuzz3(r randyFuzz3) string {
	v44 := r.Intn(100)
	tmps := make([]rune, v44)
	for i := 0; i < v44; i++ {
		tmps[i] = randUTF8RuneFuzz3(r)
	}
	return string(tmps)
}
func randUnrecognizedFuzz3(r randyFuzz3, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFuzz3(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFuzz3(data []byte, r randyFuzz3, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFuzz3(data, uint64(key))
		v45 := r.Int63()
		if r.Intn(2) == 0 {
			v45 *= -1
		}
		data = encodeVarintPopulateFuzz3(data, uint64(v45))
	case 1:
		data = encodeVarintPopulateFuzz3(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFuzz3(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFuzz3(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFuzz3(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFuzz3(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
