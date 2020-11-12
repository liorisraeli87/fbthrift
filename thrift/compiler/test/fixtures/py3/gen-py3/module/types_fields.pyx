#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cython.operator cimport dereference as deref
from libcpp.memory cimport make_unique, unique_ptr, shared_ptr

cimport thrift.py3.types
from thrift.py3.types cimport (
    reset_field as __reset_field,
    StructFieldsSetter as __StructFieldsSetter
)

from thrift.py3.types cimport const_pointer_cast


@__cython.auto_pickle(False)
cdef class __SimpleException_FieldsSetter(__StructFieldsSetter):

    @staticmethod
    cdef __SimpleException_FieldsSetter create(_module_types.cSimpleException* struct_cpp_obj):
        cdef __SimpleException_FieldsSetter __fbthrift_inst = __SimpleException_FieldsSetter.__new__(__SimpleException_FieldsSetter)
        __fbthrift_inst._struct_cpp_obj = struct_cpp_obj
        __fbthrift_inst._setters[__cstring_view(<const char*>"err_code")] = __SimpleException_FieldsSetter._set_field_0
        return __fbthrift_inst

    cdef void set_field(__SimpleException_FieldsSetter self, const char* name, object value) except *:
        cdef __cstring_view cname = __cstring_view(name)
        cdef cumap[__cstring_view, __SimpleException_FieldsSetterFunc].iterator found = self._setters.find(cname)
        if found == self._setters.end():
            raise TypeError(f"invalid field name {name.decode('utf-8')}")
        deref(found).second(self, value)

    cdef void _set_field_0(self, __fbthrift_value) except *:
        # for field err_code
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleException](deref(self._struct_cpp_obj), 0)
            return
        if not isinstance(__fbthrift_value, int):
            raise TypeError(f'err_code is not a { int !r}.')
        __fbthrift_value = <cint16_t> __fbthrift_value
        deref(self._struct_cpp_obj).err_code_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.err_code = True


@__cython.auto_pickle(False)
cdef class __OptionalRefStruct_FieldsSetter(__StructFieldsSetter):

    @staticmethod
    cdef __OptionalRefStruct_FieldsSetter create(_module_types.cOptionalRefStruct* struct_cpp_obj):
        cdef __OptionalRefStruct_FieldsSetter __fbthrift_inst = __OptionalRefStruct_FieldsSetter.__new__(__OptionalRefStruct_FieldsSetter)
        __fbthrift_inst._struct_cpp_obj = struct_cpp_obj
        __fbthrift_inst._setters[__cstring_view(<const char*>"optional_blob")] = __OptionalRefStruct_FieldsSetter._set_field_0
        return __fbthrift_inst

    cdef void set_field(__OptionalRefStruct_FieldsSetter self, const char* name, object value) except *:
        cdef __cstring_view cname = __cstring_view(name)
        cdef cumap[__cstring_view, __OptionalRefStruct_FieldsSetterFunc].iterator found = self._setters.find(cname)
        if found == self._setters.end():
            raise TypeError(f"invalid field name {name.decode('utf-8')}")
        deref(found).second(self, value)

    cdef void _set_field_0(self, __fbthrift_value) except *:
        # for field optional_blob
        if __fbthrift_value is None:
            __reset_field[_module_types.cOptionalRefStruct](deref(self._struct_cpp_obj), 0)
            return
        if not isinstance(__fbthrift_value, __iobuf.IOBuf):
            raise TypeError(f'optional_blob is not a { __iobuf.IOBuf !r}.')
        deref(self._struct_cpp_obj).optional_blob_ref().assign((<__iobuf.IOBuf?>__fbthrift_value).c_clone())
        deref(self._struct_cpp_obj).__isset.optional_blob = True


@__cython.auto_pickle(False)
cdef class __SimpleStruct_FieldsSetter(__StructFieldsSetter):

    @staticmethod
    cdef __SimpleStruct_FieldsSetter create(_module_types.cSimpleStruct* struct_cpp_obj):
        cdef __SimpleStruct_FieldsSetter __fbthrift_inst = __SimpleStruct_FieldsSetter.__new__(__SimpleStruct_FieldsSetter)
        __fbthrift_inst._struct_cpp_obj = struct_cpp_obj
        __fbthrift_inst._setters[__cstring_view(<const char*>"is_on")] = __SimpleStruct_FieldsSetter._set_field_0
        __fbthrift_inst._setters[__cstring_view(<const char*>"tiny_int")] = __SimpleStruct_FieldsSetter._set_field_1
        __fbthrift_inst._setters[__cstring_view(<const char*>"small_int")] = __SimpleStruct_FieldsSetter._set_field_2
        __fbthrift_inst._setters[__cstring_view(<const char*>"nice_sized_int")] = __SimpleStruct_FieldsSetter._set_field_3
        __fbthrift_inst._setters[__cstring_view(<const char*>"big_int")] = __SimpleStruct_FieldsSetter._set_field_4
        __fbthrift_inst._setters[__cstring_view(<const char*>"real")] = __SimpleStruct_FieldsSetter._set_field_5
        __fbthrift_inst._setters[__cstring_view(<const char*>"smaller_real")] = __SimpleStruct_FieldsSetter._set_field_6
        return __fbthrift_inst

    cdef void set_field(__SimpleStruct_FieldsSetter self, const char* name, object value) except *:
        cdef __cstring_view cname = __cstring_view(name)
        cdef cumap[__cstring_view, __SimpleStruct_FieldsSetterFunc].iterator found = self._setters.find(cname)
        if found == self._setters.end():
            raise TypeError(f"invalid field name {name.decode('utf-8')}")
        deref(found).second(self, value)

    cdef void _set_field_0(self, __fbthrift_value) except *:
        # for field is_on
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 0)
            return
        if not isinstance(__fbthrift_value, bool):
            raise TypeError(f'is_on is not a { bool !r}.')
        deref(self._struct_cpp_obj).is_on_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.is_on = True

    cdef void _set_field_1(self, __fbthrift_value) except *:
        # for field tiny_int
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 1)
            return
        if not isinstance(__fbthrift_value, int):
            raise TypeError(f'tiny_int is not a { int !r}.')
        __fbthrift_value = <cint8_t> __fbthrift_value
        deref(self._struct_cpp_obj).tiny_int_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.tiny_int = True

    cdef void _set_field_2(self, __fbthrift_value) except *:
        # for field small_int
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 2)
            return
        if not isinstance(__fbthrift_value, int):
            raise TypeError(f'small_int is not a { int !r}.')
        __fbthrift_value = <cint16_t> __fbthrift_value
        deref(self._struct_cpp_obj).small_int_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.small_int = True

    cdef void _set_field_3(self, __fbthrift_value) except *:
        # for field nice_sized_int
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 3)
            return
        if not isinstance(__fbthrift_value, int):
            raise TypeError(f'nice_sized_int is not a { int !r}.')
        __fbthrift_value = <cint32_t> __fbthrift_value
        deref(self._struct_cpp_obj).nice_sized_int_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.nice_sized_int = True

    cdef void _set_field_4(self, __fbthrift_value) except *:
        # for field big_int
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 4)
            return
        if not isinstance(__fbthrift_value, int):
            raise TypeError(f'big_int is not a { int !r}.')
        __fbthrift_value = <cint64_t> __fbthrift_value
        deref(self._struct_cpp_obj).big_int_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.big_int = True

    cdef void _set_field_5(self, __fbthrift_value) except *:
        # for field real
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 5)
            return
        if not isinstance(__fbthrift_value, (float, int)):
            raise TypeError(f'real is not a { float !r}.')
        deref(self._struct_cpp_obj).real_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.real = True

    cdef void _set_field_6(self, __fbthrift_value) except *:
        # for field smaller_real
        if __fbthrift_value is None:
            __reset_field[_module_types.cSimpleStruct](deref(self._struct_cpp_obj), 6)
            return
        if not isinstance(__fbthrift_value, (float, int)):
            raise TypeError(f'smaller_real is not a { float !r}.')
        deref(self._struct_cpp_obj).smaller_real_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.smaller_real = True


@__cython.auto_pickle(False)
cdef class __ComplexStruct_FieldsSetter(__StructFieldsSetter):

    @staticmethod
    cdef __ComplexStruct_FieldsSetter create(_module_types.cComplexStruct* struct_cpp_obj):
        cdef __ComplexStruct_FieldsSetter __fbthrift_inst = __ComplexStruct_FieldsSetter.__new__(__ComplexStruct_FieldsSetter)
        __fbthrift_inst._struct_cpp_obj = struct_cpp_obj
        __fbthrift_inst._setters[__cstring_view(<const char*>"structOne")] = __ComplexStruct_FieldsSetter._set_field_0
        __fbthrift_inst._setters[__cstring_view(<const char*>"structTwo")] = __ComplexStruct_FieldsSetter._set_field_1
        __fbthrift_inst._setters[__cstring_view(<const char*>"an_integer")] = __ComplexStruct_FieldsSetter._set_field_2
        __fbthrift_inst._setters[__cstring_view(<const char*>"name")] = __ComplexStruct_FieldsSetter._set_field_3
        __fbthrift_inst._setters[__cstring_view(<const char*>"an_enum")] = __ComplexStruct_FieldsSetter._set_field_4
        __fbthrift_inst._setters[__cstring_view(<const char*>"some_bytes")] = __ComplexStruct_FieldsSetter._set_field_5
        __fbthrift_inst._setters[__cstring_view(<const char*>"sender")] = __ComplexStruct_FieldsSetter._set_field_6
        __fbthrift_inst._setters[__cstring_view(<const char*>"cdef_")] = __ComplexStruct_FieldsSetter._set_field_7
        __fbthrift_inst._setters[__cstring_view(<const char*>"bytes_with_cpp_type")] = __ComplexStruct_FieldsSetter._set_field_8
        return __fbthrift_inst

    cdef void set_field(__ComplexStruct_FieldsSetter self, const char* name, object value) except *:
        cdef __cstring_view cname = __cstring_view(name)
        cdef cumap[__cstring_view, __ComplexStruct_FieldsSetterFunc].iterator found = self._setters.find(cname)
        if found == self._setters.end():
            raise TypeError(f"invalid field name {name.decode('utf-8')}")
        deref(found).second(self, value)

    cdef void _set_field_0(self, __fbthrift_value) except *:
        # for field structOne
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 0)
            return
        if not isinstance(__fbthrift_value, _module_types.SimpleStruct):
            raise TypeError(f'structOne is not a { _module_types.SimpleStruct !r}.')
        deref(self._struct_cpp_obj).structOne_ref().assign(deref((<_module_types.SimpleStruct?> __fbthrift_value)._cpp_obj))
        deref(self._struct_cpp_obj).__isset.structOne = True

    cdef void _set_field_1(self, __fbthrift_value) except *:
        # for field structTwo
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 1)
            return
        if not isinstance(__fbthrift_value, _module_types.SimpleStruct):
            raise TypeError(f'structTwo is not a { _module_types.SimpleStruct !r}.')
        deref(self._struct_cpp_obj).structTwo_ref().assign(deref((<_module_types.SimpleStruct?> __fbthrift_value)._cpp_obj))
        deref(self._struct_cpp_obj).__isset.structTwo = True

    cdef void _set_field_2(self, __fbthrift_value) except *:
        # for field an_integer
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 2)
            return
        if not isinstance(__fbthrift_value, int):
            raise TypeError(f'an_integer is not a { int !r}.')
        __fbthrift_value = <cint32_t> __fbthrift_value
        deref(self._struct_cpp_obj).an_integer_ref().assign(__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.an_integer = True

    cdef void _set_field_3(self, __fbthrift_value) except *:
        # for field name
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 3)
            return
        if not isinstance(__fbthrift_value, str):
            raise TypeError(f'name is not a { str !r}.')
        deref(self._struct_cpp_obj).name_ref().assign(cmove(bytes_to_string(__fbthrift_value.encode('utf-8'))))
        deref(self._struct_cpp_obj).__isset.name = True

    cdef void _set_field_4(self, __fbthrift_value) except *:
        # for field an_enum
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 4)
            return
        if not isinstance(__fbthrift_value, _module_types.AnEnum):
            raise TypeError(f'field an_enum value: {repr(__fbthrift_value)} is not of the enum type { _module_types.AnEnum }.')
        deref(self._struct_cpp_obj).an_enum_ref().assign(<_module_types.cAnEnum><int>__fbthrift_value)
        deref(self._struct_cpp_obj).__isset.an_enum = True

    cdef void _set_field_5(self, __fbthrift_value) except *:
        # for field some_bytes
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 5)
            return
        if not isinstance(__fbthrift_value, bytes):
            raise TypeError(f'some_bytes is not a { bytes !r}.')
        deref(self._struct_cpp_obj).some_bytes_ref().assign(cmove(bytes_to_string(__fbthrift_value)))
        deref(self._struct_cpp_obj).__isset.some_bytes = True

    cdef void _set_field_6(self, __fbthrift_value) except *:
        # for field sender
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 6)
            return
        if not isinstance(__fbthrift_value, str):
            raise TypeError(f'sender is not a { str !r}.')
        deref(self._struct_cpp_obj).sender_ref().assign(cmove(bytes_to_string(__fbthrift_value.encode('utf-8'))))
        deref(self._struct_cpp_obj).__isset.sender = True

    cdef void _set_field_7(self, __fbthrift_value) except *:
        # for field cdef_
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 7)
            return
        if not isinstance(__fbthrift_value, str):
            raise TypeError(f'cdef_ is not a { str !r}.')
        deref(self._struct_cpp_obj).cdef__ref().assign(cmove(bytes_to_string(__fbthrift_value.encode('utf-8'))))
        deref(self._struct_cpp_obj).__isset.cdef_ = True

    cdef void _set_field_8(self, __fbthrift_value) except *:
        # for field bytes_with_cpp_type
        if __fbthrift_value is None:
            __reset_field[_module_types.cComplexStruct](deref(self._struct_cpp_obj), 8)
            return
        if not isinstance(__fbthrift_value, bytes):
            raise TypeError(f'bytes_with_cpp_type is not a { bytes !r}.')
        deref(self._struct_cpp_obj).bytes_with_cpp_type_ref().assign(_module_types.foo_Bar(cmove(<string>__fbthrift_value)))
        deref(self._struct_cpp_obj).__isset.bytes_with_cpp_type = True


@__cython.auto_pickle(False)
cdef class __BinaryUnionStruct_FieldsSetter(__StructFieldsSetter):

    @staticmethod
    cdef __BinaryUnionStruct_FieldsSetter create(_module_types.cBinaryUnionStruct* struct_cpp_obj):
        cdef __BinaryUnionStruct_FieldsSetter __fbthrift_inst = __BinaryUnionStruct_FieldsSetter.__new__(__BinaryUnionStruct_FieldsSetter)
        __fbthrift_inst._struct_cpp_obj = struct_cpp_obj
        __fbthrift_inst._setters[__cstring_view(<const char*>"u")] = __BinaryUnionStruct_FieldsSetter._set_field_0
        return __fbthrift_inst

    cdef void set_field(__BinaryUnionStruct_FieldsSetter self, const char* name, object value) except *:
        cdef __cstring_view cname = __cstring_view(name)
        cdef cumap[__cstring_view, __BinaryUnionStruct_FieldsSetterFunc].iterator found = self._setters.find(cname)
        if found == self._setters.end():
            raise TypeError(f"invalid field name {name.decode('utf-8')}")
        deref(found).second(self, value)

    cdef void _set_field_0(self, __fbthrift_value) except *:
        # for field u
        if __fbthrift_value is None:
            __reset_field[_module_types.cBinaryUnionStruct](deref(self._struct_cpp_obj), 0)
            return
        if not isinstance(__fbthrift_value, _module_types.BinaryUnion):
            raise TypeError(f'u is not a { _module_types.BinaryUnion !r}.')
        deref(self._struct_cpp_obj).u_ref().assign(deref((<_module_types.BinaryUnion?> __fbthrift_value)._cpp_obj))
        deref(self._struct_cpp_obj).__isset.u = True
