// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package module

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift-go"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

var GoUnusedProtection__ int;

type ContainerTypedef = map[int16]string

func ContainerTypedefPtr(v ContainerTypedef) *ContainerTypedef { return &v }

// Attributes:
//  - IntValue
//  - StringValue
//  - IntListValue
//  - StringListValue
//  - TypedefValue
//  - StringRef
type ComplexUnion struct {
  IntValue *int64 `thrift:"intValue,1" db:"intValue" json:"intValue,omitempty"`
  IntListValue []int64 `thrift:"intListValue,2" db:"intListValue" json:"intListValue,omitempty"`
  StringListValue []string `thrift:"stringListValue,3" db:"stringListValue" json:"stringListValue,omitempty"`
  // unused field # 4
  StringValue *string `thrift:"stringValue,5" db:"stringValue" json:"stringValue,omitempty"`
  // unused fields # 6 to 8
  TypedefValue ContainerTypedef `thrift:"typedefValue,9" db:"typedefValue" json:"typedefValue,omitempty"`
  // unused fields # 10 to 13
  StringRef *string `thrift:"stringRef,14" db:"stringRef" json:"stringRef,omitempty"`
}

func NewComplexUnion() *ComplexUnion {
  return &ComplexUnion{}
}

var ComplexUnion_IntValue_DEFAULT int64
func (p *ComplexUnion) GetIntValue() int64 {
  if !p.IsSetIntValue() {
    return ComplexUnion_IntValue_DEFAULT
  }
return *p.IntValue
}
var ComplexUnion_StringValue_DEFAULT string
func (p *ComplexUnion) GetStringValue() string {
  if !p.IsSetStringValue() {
    return ComplexUnion_StringValue_DEFAULT
  }
return *p.StringValue
}
var ComplexUnion_IntListValue_DEFAULT []int64

func (p *ComplexUnion) GetIntListValue() []int64 {
  return p.IntListValue
}
var ComplexUnion_StringListValue_DEFAULT []string

func (p *ComplexUnion) GetStringListValue() []string {
  return p.StringListValue
}
var ComplexUnion_TypedefValue_DEFAULT ContainerTypedef

func (p *ComplexUnion) GetTypedefValue() ContainerTypedef {
  return p.TypedefValue
}
var ComplexUnion_StringRef_DEFAULT string
func (p *ComplexUnion) GetStringRef() string {
  if !p.IsSetStringRef() {
    return ComplexUnion_StringRef_DEFAULT
  }
return *p.StringRef
}
func (p *ComplexUnion) CountSetFieldsComplexUnion() int {
  count := 0
  if (p.IsSetIntValue()) {
    count++
  }
  if (p.IsSetStringValue()) {
    count++
  }
  if (p.IsSetStringRef()) {
    count++
  }
  return count

}

func (p *ComplexUnion) IsSetIntValue() bool {
  return p.IntValue != nil
}

func (p *ComplexUnion) IsSetStringValue() bool {
  return p.StringValue != nil
}

func (p *ComplexUnion) IsSetIntListValue() bool {
  return p.IntListValue != nil
}

func (p *ComplexUnion) IsSetStringListValue() bool {
  return p.StringListValue != nil
}

func (p *ComplexUnion) IsSetTypedefValue() bool {
  return p.TypedefValue != nil
}

func (p *ComplexUnion) IsSetStringRef() bool {
  return p.StringRef != nil
}

func (p *ComplexUnion) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 9:
      if err := p.ReadField9(iprot); err != nil {
        return err
      }
    case 14:
      if err := p.ReadField14(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ComplexUnion)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.IntValue = &v
}
  return nil
}

func (p *ComplexUnion)  ReadField5(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.StringValue = &v
}
  return nil
}

func (p *ComplexUnion)  ReadField2(iprot thrift.Protocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]int64, 0, size)
  p.IntListValue =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 int64
    if v, err := iprot.ReadI64(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.IntListValue = append(p.IntListValue, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ComplexUnion)  ReadField3(iprot thrift.Protocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.StringListValue =  tSlice
  for i := 0; i < size; i ++ {
var _elem1 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem1 = v
}
    p.StringListValue = append(p.StringListValue, _elem1)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ComplexUnion)  ReadField9(iprot thrift.Protocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(ContainerTypedef, size)
  p.TypedefValue =  tMap
  for i := 0; i < size; i ++ {
var _key2 int16
    if v, err := iprot.ReadI16(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key2 = v
}
var _val3 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val3 = v
}
    p.TypedefValue[_key2] = _val3
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *ComplexUnion)  ReadField14(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 14: ", err)
} else {
  p.StringRef = &v
}
  return nil
}

func (p *ComplexUnion) Write(oprot thrift.Protocol) error {
  if c := p.CountSetFieldsComplexUnion(); c != 1 {
    return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
  }
  if err := oprot.WriteStructBegin("ComplexUnion"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField5(oprot); err != nil { return err }
  if err := p.writeField9(oprot); err != nil { return err }
  if err := p.writeField14(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ComplexUnion) writeField1(oprot thrift.Protocol) (err error) {
  if p.IsSetIntValue() {
    if err := oprot.WriteFieldBegin("intValue", thrift.I64, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:intValue: ", p), err) }
    if err := oprot.WriteI64(int64(*p.IntValue)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.intValue (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:intValue: ", p), err) }
  }
  return err
}

func (p *ComplexUnion) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetIntListValue() {
    if err := oprot.WriteFieldBegin("intListValue", thrift.LIST, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:intListValue: ", p), err) }
    if err := oprot.WriteListBegin(thrift.I64, len(p.IntListValue)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.IntListValue {
      if err := oprot.WriteI64(int64(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:intListValue: ", p), err) }
  }
  return err
}

func (p *ComplexUnion) writeField3(oprot thrift.Protocol) (err error) {
  if p.IsSetStringListValue() {
    if err := oprot.WriteFieldBegin("stringListValue", thrift.LIST, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:stringListValue: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRING, len(p.StringListValue)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.StringListValue {
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:stringListValue: ", p), err) }
  }
  return err
}

func (p *ComplexUnion) writeField5(oprot thrift.Protocol) (err error) {
  if p.IsSetStringValue() {
    if err := oprot.WriteFieldBegin("stringValue", thrift.STRING, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:stringValue: ", p), err) }
    if err := oprot.WriteString(string(*p.StringValue)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.stringValue (5) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:stringValue: ", p), err) }
  }
  return err
}

func (p *ComplexUnion) writeField9(oprot thrift.Protocol) (err error) {
  if p.IsSetTypedefValue() {
    if err := oprot.WriteFieldBegin("typedefValue", thrift.MAP, 9); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:typedefValue: ", p), err) }
    if err := oprot.WriteMapBegin(thrift.I16, thrift.STRING, len(p.TypedefValue)); err != nil {
      return thrift.PrependError("error writing map begin: ", err)
    }
    for k, v := range p.TypedefValue {
      if err := oprot.WriteI16(int16(k)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteMapEnd(); err != nil {
      return thrift.PrependError("error writing map end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 9:typedefValue: ", p), err) }
  }
  return err
}

func (p *ComplexUnion) writeField14(oprot thrift.Protocol) (err error) {
  if p.IsSetStringRef() {
    if err := oprot.WriteFieldBegin("stringRef", thrift.STRING, 14); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 14:stringRef: ", p), err) }
    if err := oprot.WriteString(string(*p.StringRef)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.stringRef (14) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 14:stringRef: ", p), err) }
  }
  return err
}

func (p *ComplexUnion) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ComplexUnion(%+v)", *p)
}

// Attributes:
//  - IntListValue
//  - StringListValue
type ListUnion struct {
  // unused field # 1
  IntListValue []int64 `thrift:"intListValue,2" db:"intListValue" json:"intListValue,omitempty"`
  StringListValue []string `thrift:"stringListValue,3" db:"stringListValue" json:"stringListValue,omitempty"`
}

func NewListUnion() *ListUnion {
  return &ListUnion{}
}

var ListUnion_IntListValue_DEFAULT []int64

func (p *ListUnion) GetIntListValue() []int64 {
  return p.IntListValue
}
var ListUnion_StringListValue_DEFAULT []string

func (p *ListUnion) GetStringListValue() []string {
  return p.StringListValue
}
func (p *ListUnion) IsSetIntListValue() bool {
  return p.IntListValue != nil
}

func (p *ListUnion) IsSetStringListValue() bool {
  return p.StringListValue != nil
}

func (p *ListUnion) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ListUnion)  ReadField2(iprot thrift.Protocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]int64, 0, size)
  p.IntListValue =  tSlice
  for i := 0; i < size; i ++ {
var _elem4 int64
    if v, err := iprot.ReadI64(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem4 = v
}
    p.IntListValue = append(p.IntListValue, _elem4)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ListUnion)  ReadField3(iprot thrift.Protocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.StringListValue =  tSlice
  for i := 0; i < size; i ++ {
var _elem5 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem5 = v
}
    p.StringListValue = append(p.StringListValue, _elem5)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ListUnion) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("ListUnion"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ListUnion) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetIntListValue() {
    if err := oprot.WriteFieldBegin("intListValue", thrift.LIST, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:intListValue: ", p), err) }
    if err := oprot.WriteListBegin(thrift.I64, len(p.IntListValue)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.IntListValue {
      if err := oprot.WriteI64(int64(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:intListValue: ", p), err) }
  }
  return err
}

func (p *ListUnion) writeField3(oprot thrift.Protocol) (err error) {
  if p.IsSetStringListValue() {
    if err := oprot.WriteFieldBegin("stringListValue", thrift.LIST, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:stringListValue: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRING, len(p.StringListValue)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.StringListValue {
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:stringListValue: ", p), err) }
  }
  return err
}

func (p *ListUnion) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ListUnion(%+v)", *p)
}

// Attributes:
//  - BinaryData
//  - StringData
type DataUnion struct {
  BinaryData []byte `thrift:"binaryData,1" db:"binaryData" json:"binaryData,omitempty"`
  StringData *string `thrift:"stringData,2" db:"stringData" json:"stringData,omitempty"`
}

func NewDataUnion() *DataUnion {
  return &DataUnion{}
}

var DataUnion_BinaryData_DEFAULT []byte

func (p *DataUnion) GetBinaryData() []byte {
  return p.BinaryData
}
var DataUnion_StringData_DEFAULT string
func (p *DataUnion) GetStringData() string {
  if !p.IsSetStringData() {
    return DataUnion_StringData_DEFAULT
  }
return *p.StringData
}
func (p *DataUnion) CountSetFieldsDataUnion() int {
  count := 0
  if (p.IsSetStringData()) {
    count++
  }
  return count

}

func (p *DataUnion) IsSetBinaryData() bool {
  return p.BinaryData != nil
}

func (p *DataUnion) IsSetStringData() bool {
  return p.StringData != nil
}

func (p *DataUnion) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *DataUnion)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.BinaryData = v
}
  return nil
}

func (p *DataUnion)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.StringData = &v
}
  return nil
}

func (p *DataUnion) Write(oprot thrift.Protocol) error {
  if c := p.CountSetFieldsDataUnion(); c != 1 {
    return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
  }
  if err := oprot.WriteStructBegin("DataUnion"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *DataUnion) writeField1(oprot thrift.Protocol) (err error) {
  if p.IsSetBinaryData() {
    if err := oprot.WriteFieldBegin("binaryData", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:binaryData: ", p), err) }
    if err := oprot.WriteBinary(p.BinaryData); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.binaryData (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:binaryData: ", p), err) }
  }
  return err
}

func (p *DataUnion) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetStringData() {
    if err := oprot.WriteFieldBegin("stringData", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stringData: ", p), err) }
    if err := oprot.WriteString(string(*p.StringData)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.stringData (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stringData: ", p), err) }
  }
  return err
}

func (p *DataUnion) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("DataUnion(%+v)", *p)
}

// Attributes:
//  - StrVal
//  - IntVal
//  - TypedefValue
type Val struct {
  StrVal string `thrift:"strVal,1" db:"strVal" json:"strVal"`
  IntVal int32 `thrift:"intVal,2" db:"intVal" json:"intVal"`
  // unused fields # 3 to 8
  TypedefValue ContainerTypedef `thrift:"typedefValue,9" db:"typedefValue" json:"typedefValue"`
}

func NewVal() *Val {
  return &Val{}
}


func (p *Val) GetStrVal() string {
  return p.StrVal
}

func (p *Val) GetIntVal() int32 {
  return p.IntVal
}

func (p *Val) GetTypedefValue() ContainerTypedef {
  return p.TypedefValue
}
func (p *Val) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 9:
      if err := p.ReadField9(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Val)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.StrVal = v
}
  return nil
}

func (p *Val)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.IntVal = v
}
  return nil
}

func (p *Val)  ReadField9(iprot thrift.Protocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(ContainerTypedef, size)
  p.TypedefValue =  tMap
  for i := 0; i < size; i ++ {
var _key6 int16
    if v, err := iprot.ReadI16(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key6 = v
}
var _val7 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val7 = v
}
    p.TypedefValue[_key6] = _val7
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *Val) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("Val"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField9(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Val) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("strVal", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:strVal: ", p), err) }
  if err := oprot.WriteString(string(p.StrVal)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.strVal (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:strVal: ", p), err) }
  return err
}

func (p *Val) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("intVal", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:intVal: ", p), err) }
  if err := oprot.WriteI32(int32(p.IntVal)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.intVal (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:intVal: ", p), err) }
  return err
}

func (p *Val) writeField9(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("typedefValue", thrift.MAP, 9); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:typedefValue: ", p), err) }
  if err := oprot.WriteMapBegin(thrift.I16, thrift.STRING, len(p.TypedefValue)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.TypedefValue {
    if err := oprot.WriteI16(int16(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 9:typedefValue: ", p), err) }
  return err
}

func (p *Val) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Val(%+v)", *p)
}

// Attributes:
//  - V1
//  - V2
type ValUnion struct {
  V1 *Val `thrift:"v1,1" db:"v1" json:"v1,omitempty"`
  V2 *Val `thrift:"v2,2" db:"v2" json:"v2,omitempty"`
}

func NewValUnion() *ValUnion {
  return &ValUnion{}
}

var ValUnion_V1_DEFAULT *Val
func (p *ValUnion) GetV1() *Val {
  if !p.IsSetV1() {
    return ValUnion_V1_DEFAULT
  }
return p.V1
}
var ValUnion_V2_DEFAULT *Val
func (p *ValUnion) GetV2() *Val {
  if !p.IsSetV2() {
    return ValUnion_V2_DEFAULT
  }
return p.V2
}
func (p *ValUnion) CountSetFieldsValUnion() int {
  count := 0
  if (p.IsSetV1()) {
    count++
  }
  if (p.IsSetV2()) {
    count++
  }
  return count

}

func (p *ValUnion) IsSetV1() bool {
  return p.V1 != nil
}

func (p *ValUnion) IsSetV2() bool {
  return p.V2 != nil
}

func (p *ValUnion) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ValUnion)  ReadField1(iprot thrift.Protocol) error {
  p.V1 = NewVal()
  if err := p.V1.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.V1), err)
  }
  return nil
}

func (p *ValUnion)  ReadField2(iprot thrift.Protocol) error {
  p.V2 = NewVal()
  if err := p.V2.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.V2), err)
  }
  return nil
}

func (p *ValUnion) Write(oprot thrift.Protocol) error {
  if c := p.CountSetFieldsValUnion(); c != 1 {
    return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
  }
  if err := oprot.WriteStructBegin("ValUnion"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ValUnion) writeField1(oprot thrift.Protocol) (err error) {
  if p.IsSetV1() {
    if err := oprot.WriteFieldBegin("v1", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:v1: ", p), err) }
    if err := p.V1.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.V1), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:v1: ", p), err) }
  }
  return err
}

func (p *ValUnion) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetV2() {
    if err := oprot.WriteFieldBegin("v2", thrift.STRUCT, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:v2: ", p), err) }
    if err := p.V2.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.V2), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:v2: ", p), err) }
  }
  return err
}

func (p *ValUnion) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ValUnion(%+v)", *p)
}

// Attributes:
//  - ThingOne
//  - ThingTwo
type VirtualComplexUnion struct {
  ThingOne *string `thrift:"thingOne,1" db:"thingOne" json:"thingOne,omitempty"`
  ThingTwo *string `thrift:"thingTwo,2" db:"thingTwo" json:"thingTwo,omitempty"`
}

func NewVirtualComplexUnion() *VirtualComplexUnion {
  return &VirtualComplexUnion{}
}

var VirtualComplexUnion_ThingOne_DEFAULT string
func (p *VirtualComplexUnion) GetThingOne() string {
  if !p.IsSetThingOne() {
    return VirtualComplexUnion_ThingOne_DEFAULT
  }
return *p.ThingOne
}
var VirtualComplexUnion_ThingTwo_DEFAULT string
func (p *VirtualComplexUnion) GetThingTwo() string {
  if !p.IsSetThingTwo() {
    return VirtualComplexUnion_ThingTwo_DEFAULT
  }
return *p.ThingTwo
}
func (p *VirtualComplexUnion) CountSetFieldsVirtualComplexUnion() int {
  count := 0
  if (p.IsSetThingOne()) {
    count++
  }
  if (p.IsSetThingTwo()) {
    count++
  }
  return count

}

func (p *VirtualComplexUnion) IsSetThingOne() bool {
  return p.ThingOne != nil
}

func (p *VirtualComplexUnion) IsSetThingTwo() bool {
  return p.ThingTwo != nil
}

func (p *VirtualComplexUnion) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *VirtualComplexUnion)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ThingOne = &v
}
  return nil
}

func (p *VirtualComplexUnion)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ThingTwo = &v
}
  return nil
}

func (p *VirtualComplexUnion) Write(oprot thrift.Protocol) error {
  if c := p.CountSetFieldsVirtualComplexUnion(); c != 1 {
    return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
  }
  if err := oprot.WriteStructBegin("VirtualComplexUnion"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *VirtualComplexUnion) writeField1(oprot thrift.Protocol) (err error) {
  if p.IsSetThingOne() {
    if err := oprot.WriteFieldBegin("thingOne", thrift.STRING, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:thingOne: ", p), err) }
    if err := oprot.WriteString(string(*p.ThingOne)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.thingOne (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:thingOne: ", p), err) }
  }
  return err
}

func (p *VirtualComplexUnion) writeField2(oprot thrift.Protocol) (err error) {
  if p.IsSetThingTwo() {
    if err := oprot.WriteFieldBegin("thingTwo", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:thingTwo: ", p), err) }
    if err := oprot.WriteString(string(*p.ThingTwo)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.thingTwo (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:thingTwo: ", p), err) }
  }
  return err
}

func (p *VirtualComplexUnion) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("VirtualComplexUnion(%+v)", *p)
}

