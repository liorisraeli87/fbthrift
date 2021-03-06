/**
 * Autogenerated by Thrift for src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#pragma once

#include <thrift/lib/cpp2/gen/module_data_h.h>

#include "thrift/compiler/test/fixtures/complex-struct/gen-cpp2/module_types.h"

namespace apache { namespace thrift {

template <> struct TEnumDataStorage<::cpp2::MyEnum> {
  using type = ::cpp2::MyEnum;
  static constexpr const std::size_t size = 5;
  static constexpr const std::array<type, size> values = {{
    type::MyValue1,
    type::MyValue2,
    type::MyValue3,
    type::MyValue4,
    type::MyValue5,
  }};
  static constexpr const std::array<folly::StringPiece, size> names = {{
    "MyValue1",
    "MyValue2",
    "MyValue3",
    "MyValue4",
    "MyValue5",
  }};
};

template <> struct TEnumDataStorage<::cpp2::MyUnion::Type> {
  using type = ::cpp2::MyUnion::Type;
  static constexpr const std::size_t size = 6;
  static constexpr const std::array<type, size> values = {{
    type::myEnum,
    type::myStruct,
    type::myDataItem,
    type::complexNestedStruct,
    type::longValue,
    type::intValue,
  }};
  static constexpr const std::array<folly::StringPiece, size> names = {{
    "myEnum",
    "myStruct",
    "myDataItem",
    "complexNestedStruct",
    "longValue",
    "intValue",
  }};
};

template <> struct TEnumDataStorage<::cpp2::MyUnionFloatFieldThrowExp::Type> {
  using type = ::cpp2::MyUnionFloatFieldThrowExp::Type;
  static constexpr const std::size_t size = 4;
  static constexpr const std::array<type, size> values = {{
    type::myEnum,
    type::setFloat,
    type::myDataItem,
    type::complexNestedStruct,
  }};
  static constexpr const std::array<folly::StringPiece, size> names = {{
    "myEnum",
    "setFloat",
    "myDataItem",
    "complexNestedStruct",
  }};
};


template <> struct TStructDataStorage<::cpp2::MyStructFloatFieldThrowExp> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 4;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "myLongField",
    "MyByteField",
    "myStringField",
    "myFloatField",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I64,
    TType::T_BYTE,
    TType::T_STRING,
    TType::T_FLOAT,
  }};
};


template <> struct TStructDataStorage<::cpp2::MyStructMapFloatThrowExp> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 2;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "myLongField",
    "mapListOfFloats",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I64,
    TType::T_MAP,
  }};
};


template <> struct TStructDataStorage<::cpp2::MyDataItem> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 0;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
  }};
};


template <> struct TStructDataStorage<::cpp2::MyStruct> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 28;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "MyIntField",
    "MyStringField",
    "MyDataField",
    "myEnum",
    "MyBoolField",
    "MyByteField",
    "MyShortField",
    "MyLongField",
    "MyDoubleField",
    "lDouble",
    "lShort",
    "lInteger",
    "lLong",
    "lString",
    "lBool",
    "lByte",
    "mShortString",
    "mIntegerString",
    "mStringMyStruct",
    "mStringBool",
    "mIntegerInteger",
    "mIntegerBool",
    "sShort",
    "sMyStruct",
    "sLong",
    "sString",
    "sByte",
    "mListList",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
    12,
    13,
    14,
    15,
    16,
    17,
    18,
    19,
    20,
    21,
    22,
    23,
    24,
    25,
    26,
    27,
    28,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I64,
    TType::T_STRING,
    TType::T_STRUCT,
    TType::T_I32,
    TType::T_BOOL,
    TType::T_BYTE,
    TType::T_I16,
    TType::T_I64,
    TType::T_DOUBLE,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_SET,
    TType::T_SET,
    TType::T_SET,
    TType::T_SET,
    TType::T_SET,
    TType::T_MAP,
  }};
};


template <> struct TStructDataStorage<::cpp2::SimpleStruct> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 2;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "age",
    "name",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I64,
    TType::T_STRING,
  }};
};


template <> struct TStructDataStorage<::cpp2::ComplexNestedStruct> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 18;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "setOfSetOfInt",
    "listofListOfListOfListOfEnum",
    "listOfListOfMyStruct",
    "setOfListOfListOfLong",
    "setOfSetOfsetOfLong",
    "mapStructListOfListOfLong",
    "mKeyStructValInt",
    "listOfMapKeyIntValInt",
    "listOfMapKeyStrValList",
    "mapKeySetValLong",
    "mapKeyListValLong",
    "mapKeyMapValMap",
    "mapKeySetValMap",
    "NestedMaps",
    "mapKeyIntValList",
    "mapKeyIntValSet",
    "mapKeySetValInt",
    "mapKeyListValSet",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
    12,
    13,
    14,
    15,
    16,
    17,
    18,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_SET,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_SET,
    TType::T_SET,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
  }};
};


template <> struct TStructDataStorage<::cpp2::MyUnion> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 6;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "myEnum",
    "myStruct",
    "myDataItem",
    "complexNestedStruct",
    "longValue",
    "intValue",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
    5,
    6,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I32,
    TType::T_STRUCT,
    TType::T_STRUCT,
    TType::T_STRUCT,
    TType::T_I64,
    TType::T_I32,
  }};
};


template <> struct TStructDataStorage<::cpp2::defaultStruct> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 22;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "myLongDFset",
    "myLongDF",
    "portDFset",
    "portNum",
    "myBinaryDFset",
    "myBinary",
    "myByteDFSet",
    "myByte",
    "myDoubleDFset",
    "myDoubleDFZero",
    "myDouble",
    "field3",
    "myList",
    "mySet",
    "simpleStruct",
    "listStructDFset",
    "myUnion",
    "listUnionDFset",
    "mapNestlistStructDfSet",
    "mapJavaTypeDFset",
    "emptyMap",
    "enumMapDFset",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    12,
    13,
    14,
    15,
    16,
    17,
    18,
    19,
    20,
    21,
    22,
    23,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I64,
    TType::T_I64,
    TType::T_I32,
    TType::T_I32,
    TType::T_STRING,
    TType::T_STRING,
    TType::T_BYTE,
    TType::T_BYTE,
    TType::T_DOUBLE,
    TType::T_DOUBLE,
    TType::T_DOUBLE,
    TType::T_MAP,
    TType::T_LIST,
    TType::T_SET,
    TType::T_STRUCT,
    TType::T_LIST,
    TType::T_STRUCT,
    TType::T_LIST,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_MAP,
  }};
};


template <> struct TStructDataStorage<::cpp2::MyStructTypeDef> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 9;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "myLongField",
    "myLongTypeDef",
    "myStringField",
    "myStringTypedef",
    "myMapField",
    "myMapTypedef",
    "myListField",
    "myListTypedef",
    "myMapListOfTypeDef",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I64,
    TType::T_I64,
    TType::T_STRING,
    TType::T_STRING,
    TType::T_MAP,
    TType::T_MAP,
    TType::T_LIST,
    TType::T_LIST,
    TType::T_MAP,
  }};
};


template <> struct TStructDataStorage<::cpp2::MyUnionFloatFieldThrowExp> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 4;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "myEnum",
    "setFloat",
    "myDataItem",
    "complexNestedStruct",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_I32,
    TType::T_LIST,
    TType::T_STRUCT,
    TType::T_STRUCT,
  }};
};


template <> struct TStructDataStorage<::cpp2::TypeRemapped> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 4;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "lsMap",
    "ioMap",
    "BigInteger",
    "binaryTestBuffer",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_MAP,
    TType::T_MAP,
    TType::T_I32,
    TType::T_STRING,
  }};
};


template <> struct TStructDataStorage<::cpp2::emptyXcep> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 0;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
  }};
};


template <> struct TStructDataStorage<::cpp2::reqXcep> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 2;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "message",
    "errorCode",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_STRING,
    TType::T_I32,
  }};
};


template <> struct TStructDataStorage<::cpp2::optXcep> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 2;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "message",
    "errorCode",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_STRING,
    TType::T_I32,
  }};
};


template <> struct TStructDataStorage<::cpp2::complexException> {
 private:
  using TType = apache::thrift::protocol::TType;

 public:
  static constexpr const std::size_t fields_size = 6;
  static constexpr std::array<folly::StringPiece, fields_size> fields_names = {{
    "message",
    "listStrings",
    "errorEnum",
    "unionError",
    "structError",
    "lsMap",
  }};
  static constexpr std::array<int16_t, fields_size> fields_ids = {{
    1,
    2,
    3,
    4,
    5,
    6,
  }};
  static constexpr std::array<TType, fields_size> fields_types = {{
    TType::T_STRING,
    TType::T_LIST,
    TType::T_I32,
    TType::T_STRUCT,
    TType::T_STRUCT,
    TType::T_MAP,
  }};
};


}} // apache::thrift
