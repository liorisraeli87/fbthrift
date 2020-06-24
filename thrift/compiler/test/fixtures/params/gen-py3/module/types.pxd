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











cdef class List__i32(thrift.py3.types.Container):
    cdef shared_ptr[vector[cint32_t]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[cint32_t]])
    @staticmethod
    cdef shared_ptr[vector[cint32_t]] _make_instance(object items) except *

cdef class Map__i32_List__i32(thrift.py3.types.Container):
    cdef shared_ptr[cmap[cint32_t,vector[cint32_t]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[cint32_t,vector[cint32_t]]])
    @staticmethod
    cdef shared_ptr[cmap[cint32_t,vector[cint32_t]]] _make_instance(object items) except *

cdef class Set__i32(thrift.py3.types.Container):
    cdef shared_ptr[cset[cint32_t]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cset[cint32_t]])
    @staticmethod
    cdef shared_ptr[cset[cint32_t]] _make_instance(object items) except *

cdef class Map__i32_Set__i32(thrift.py3.types.Container):
    cdef shared_ptr[cmap[cint32_t,cset[cint32_t]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[cint32_t,cset[cint32_t]]])
    @staticmethod
    cdef shared_ptr[cmap[cint32_t,cset[cint32_t]]] _make_instance(object items) except *

cdef class Map__i32_i32(thrift.py3.types.Container):
    cdef shared_ptr[cmap[cint32_t,cint32_t]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[cint32_t,cint32_t]])
    @staticmethod
    cdef shared_ptr[cmap[cint32_t,cint32_t]] _make_instance(object items) except *

cdef class List__Map__i32_i32(thrift.py3.types.Container):
    cdef shared_ptr[vector[cmap[cint32_t,cint32_t]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[cmap[cint32_t,cint32_t]]])
    @staticmethod
    cdef shared_ptr[vector[cmap[cint32_t,cint32_t]]] _make_instance(object items) except *

cdef class List__Set__i32(thrift.py3.types.Container):
    cdef shared_ptr[vector[cset[cint32_t]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[cset[cint32_t]]])
    @staticmethod
    cdef shared_ptr[vector[cset[cint32_t]]] _make_instance(object items) except *

cdef class Map__i32_Map__i32_Set__i32(thrift.py3.types.Container):
    cdef shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]])
    @staticmethod
    cdef shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]] _make_instance(object items) except *

cdef class List__Map__i32_Map__i32_Set__i32(thrift.py3.types.Container):
    cdef shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]])
    @staticmethod
    cdef shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]] _make_instance(object items) except *

cdef class List__List__Map__i32_Map__i32_Set__i32(thrift.py3.types.Container):
    cdef shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]])
    @staticmethod
    cdef shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]] _make_instance(object items) except *

cdef extern from "<utility>" namespace "std" nogil:
    cdef shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]] move "std::move"(unique_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]])
    cdef shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]] move_shared "std::move"(shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]])
    cdef shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]] move "std::move"(unique_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]])
    cdef shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]] move_shared "std::move"(shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]])
    cdef shared_ptr[vector[cmap[cint32_t,cint32_t]]] move "std::move"(unique_ptr[vector[cmap[cint32_t,cint32_t]]])
    cdef shared_ptr[vector[cmap[cint32_t,cint32_t]]] move_shared "std::move"(shared_ptr[vector[cmap[cint32_t,cint32_t]]])
    cdef shared_ptr[vector[cset[cint32_t]]] move "std::move"(unique_ptr[vector[cset[cint32_t]]])
    cdef shared_ptr[vector[cset[cint32_t]]] move_shared "std::move"(shared_ptr[vector[cset[cint32_t]]])
    cdef shared_ptr[vector[cint32_t]] move "std::move"(unique_ptr[vector[cint32_t]])
    cdef shared_ptr[vector[cint32_t]] move_shared "std::move"(shared_ptr[vector[cint32_t]])
    cdef shared_ptr[cmap[cint32_t,vector[cint32_t]]] move "std::move"(unique_ptr[cmap[cint32_t,vector[cint32_t]]])
    cdef shared_ptr[cmap[cint32_t,vector[cint32_t]]] move_shared "std::move"(shared_ptr[cmap[cint32_t,vector[cint32_t]]])
    cdef shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]] move "std::move"(unique_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]])
    cdef shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]] move_shared "std::move"(shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]])
    cdef shared_ptr[cmap[cint32_t,cset[cint32_t]]] move "std::move"(unique_ptr[cmap[cint32_t,cset[cint32_t]]])
    cdef shared_ptr[cmap[cint32_t,cset[cint32_t]]] move_shared "std::move"(shared_ptr[cmap[cint32_t,cset[cint32_t]]])
    cdef shared_ptr[cmap[cint32_t,cint32_t]] move "std::move"(unique_ptr[cmap[cint32_t,cint32_t]])
    cdef shared_ptr[cmap[cint32_t,cint32_t]] move_shared "std::move"(shared_ptr[cmap[cint32_t,cint32_t]])
    cdef shared_ptr[cset[cint32_t]] move "std::move"(unique_ptr[cset[cint32_t]])
    cdef shared_ptr[cset[cint32_t]] move_shared "std::move"(shared_ptr[cset[cint32_t]])
cdef extern from "<utility>" nogil:
    pass  
    shared_ptr[vector[cint32_t]] reference_shared_ptr_Map__i32_List__i32 "thrift::py3::reference_shared_ptr<std::vector<int32_t>>"(...)
    shared_ptr[cset[cint32_t]] reference_shared_ptr_Map__i32_Set__i32 "thrift::py3::reference_shared_ptr<std::set<int32_t>>"(...)
    shared_ptr[cmap[cint32_t,cint32_t]] reference_shared_ptr_List__Map__i32_i32 "thrift::py3::reference_shared_ptr<std::map<int32_t,int32_t>>"(...)
    shared_ptr[cset[cint32_t]] reference_shared_ptr_List__Set__i32 "thrift::py3::reference_shared_ptr<std::set<int32_t>>"(...)
    shared_ptr[cmap[cint32_t,cset[cint32_t]]] reference_shared_ptr_Map__i32_Map__i32_Set__i32 "thrift::py3::reference_shared_ptr<std::map<int32_t,std::set<int32_t>>>"(...)
    shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]] reference_shared_ptr_List__Map__i32_Map__i32_Set__i32 "thrift::py3::reference_shared_ptr<std::map<int32_t,std::map<int32_t,std::set<int32_t>>>>"(...)
    shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]] reference_shared_ptr_List__List__Map__i32_Map__i32_Set__i32 "thrift::py3::reference_shared_ptr<std::vector<std::map<int32_t,std::map<int32_t,std::set<int32_t>>>>>"(...)
cdef extern from "<memory>" namespace "std" nogil:
    cdef shared_ptr[const vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]] const_pointer_cast "std::const_pointer_cast<const std::vector<std::vector<std::map<int32_t,std::map<int32_t,std::set<int32_t>>>>>>"(shared_ptr[vector[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]]])
    cdef shared_ptr[const vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]] const_pointer_cast "std::const_pointer_cast<const std::vector<std::map<int32_t,std::map<int32_t,std::set<int32_t>>>>>"(shared_ptr[vector[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]]])
    cdef shared_ptr[const vector[cmap[cint32_t,cint32_t]]] const_pointer_cast "std::const_pointer_cast<const std::vector<std::map<int32_t,int32_t>>>"(shared_ptr[vector[cmap[cint32_t,cint32_t]]])
    cdef shared_ptr[const vector[cset[cint32_t]]] const_pointer_cast "std::const_pointer_cast<const std::vector<std::set<int32_t>>>"(shared_ptr[vector[cset[cint32_t]]])
    cdef shared_ptr[const vector[cint32_t]] const_pointer_cast "std::const_pointer_cast<const std::vector<int32_t>>"(shared_ptr[vector[cint32_t]])
    cdef shared_ptr[const cmap[cint32_t,vector[cint32_t]]] const_pointer_cast "std::const_pointer_cast<const std::map<int32_t,std::vector<int32_t>>>"(shared_ptr[cmap[cint32_t,vector[cint32_t]]])
    cdef shared_ptr[const cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]] const_pointer_cast "std::const_pointer_cast<const std::map<int32_t,std::map<int32_t,std::set<int32_t>>>>"(shared_ptr[cmap[cint32_t,cmap[cint32_t,cset[cint32_t]]]])
    cdef shared_ptr[const cmap[cint32_t,cset[cint32_t]]] const_pointer_cast "std::const_pointer_cast<const std::map<int32_t,std::set<int32_t>>>"(shared_ptr[cmap[cint32_t,cset[cint32_t]]])
    cdef shared_ptr[const cmap[cint32_t,cint32_t]] const_pointer_cast "std::const_pointer_cast<const std::map<int32_t,int32_t>>"(shared_ptr[cmap[cint32_t,cint32_t]])
    cdef shared_ptr[const cset[cint32_t]] const_pointer_cast "std::const_pointer_cast<const std::set<int32_t>>"(shared_ptr[cset[cint32_t]])

