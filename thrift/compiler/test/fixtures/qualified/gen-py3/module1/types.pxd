#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libc.stdint cimport (
    int8_t as cint8_t,
    int16_t as cint16_t,
    int32_t as cint32_t,
    int64_t as cint64_t,
    uint32_t as cuint32_t,
)
from libcpp.string cimport string
from libcpp cimport bool as cbool, nullptr, nullptr_t
from cpython cimport bool as pbool
from libcpp.memory cimport shared_ptr, unique_ptr
from libcpp.vector cimport vector
from libcpp.set cimport set as cset
from libcpp.map cimport map as cmap, pair as cpair
from thrift.py3.exceptions cimport cTException
cimport folly.iobuf as __iobuf
cimport thrift.py3.exceptions
cimport thrift.py3.types
from thrift.py3.types cimport bstring, move, optional_field_ref
from folly.optional cimport cOptional


cdef extern from "gen-cpp2/module1_types.h" namespace "::module1":
    cdef cppclass cEnum "::module1::Enum":
        bint operator==(cEnum&)
        bint operator!=(cEnum&)
    cEnum Enum__ONE "::module1::Enum::ONE"
    cEnum Enum__TWO "::module1::Enum::TWO"
    cEnum Enum__THREE "::module1::Enum::THREE"




cdef class Enum(thrift.py3.types.CompiledEnum):
    pass


cdef cEnum Enum_to_cpp(Enum value)



cdef extern from "gen-cpp2/module1_types_custom_protocol.h" namespace "::module1":
    cdef cppclass cStruct__isset "::module1::Struct::__isset":
        bint first
        bint second

    cdef cppclass cStruct "::module1::Struct":
        cStruct() except +
        cStruct(const cStruct&) except +
        bint operator==(cStruct&)
        bint operator!=(cStruct&)
        bint operator<(cStruct&)
        bint operator>(cStruct&)
        bint operator<=(cStruct&)
        bint operator>=(cStruct&)
        cint32_t first
        string second
        cStruct__isset __isset


cdef extern from "<utility>" namespace "std" nogil:
    cdef shared_ptr[cStruct] move(unique_ptr[cStruct])
    cdef shared_ptr[cStruct] move_shared "std::move"(shared_ptr[cStruct])
    cdef unique_ptr[cStruct] move_unique "std::move"(unique_ptr[cStruct])

cdef extern from "<memory>" namespace "std" nogil:
    cdef shared_ptr[const cStruct] const_pointer_cast "std::const_pointer_cast<const ::module1::Struct>"(shared_ptr[cStruct])



cdef class Struct(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cStruct] _cpp_obj

    @staticmethod
    cdef unique_ptr[cStruct] _make_instance(
        cStruct* base_instance,
        bint* __isNOTSET,
        object first,
        str second
    ) except *

    @staticmethod
    cdef create(shared_ptr[cStruct])


cdef class List__Enum(thrift.py3.types.Container):
    cdef shared_ptr[vector[cEnum]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[cEnum]])
    @staticmethod
    cdef shared_ptr[vector[cEnum]] _make_instance(object items) except *

cdef extern from "<utility>" namespace "std" nogil:
    cdef shared_ptr[vector[cEnum]] move "std::move"(unique_ptr[vector[cEnum]])
    cdef shared_ptr[vector[cEnum]] move_shared "std::move"(shared_ptr[vector[cEnum]])
cdef extern from "<utility>" nogil:
    pass  
cdef extern from "<memory>" namespace "std" nogil:
    cdef shared_ptr[const vector[cEnum]] const_pointer_cast "std::const_pointer_cast<const std::vector<::module1::Enum>>"(shared_ptr[vector[cEnum]])

cdef extern from "gen-cpp2/module1_constants.h" namespace "::module1":
    cdef cStruct cc1 "::module1::module1_constants::c1"()
    cdef vector[cEnum] ce1s "::module1::module1_constants::e1s"()
