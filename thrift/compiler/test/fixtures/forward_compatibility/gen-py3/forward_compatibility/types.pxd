#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libcpp.string cimport string
from libcpp cimport bool as cbool
from cpython cimport bool as pbool
from libc.stdint cimport int8_t, int16_t, int32_t, int64_t
from libcpp.memory cimport shared_ptr, unique_ptr
from libcpp.vector cimport vector
from libcpp.set cimport set as cset
from libcpp.map cimport map as cmap, pair as cpair
from thrift.py3.exceptions cimport cTException, Error as __Error
cimport thrift.py3.types




cdef extern from "src/gen-cpp2/forward_compatibility_types_custom_protocol.h" namespace "cpp2":
    # Forward Declaration
    cdef cppclass cOldStructure "cpp2::OldStructure"
    # Forward Declaration
    cdef cppclass cNewStructure "cpp2::NewStructure"
    # Forward Declaration
    cdef cppclass cNewStructure2 "cpp2::NewStructure2"
    # Forward Declaration
    cdef cppclass cNewStructureNested "cpp2::NewStructureNested"
    # Forward Declaration
    cdef cppclass cNewStructureNestedField "cpp2::NewStructureNestedField"

cdef extern from "src/gen-cpp2/forward_compatibility_types.h" namespace "cpp2":
    cdef cppclass cOldStructure__isset "cpp2::OldStructure::__isset":
        bint features

    cdef cppclass cOldStructure "cpp2::OldStructure":
        cOldStructure() except +
        cOldStructure(const cOldStructure&) except +
        bint operator==(cOldStructure&)
        cmap[int16_t,double] features
        cOldStructure__isset __isset

    cdef cppclass cNewStructure__isset "cpp2::NewStructure::__isset":
        bint features

    cdef cppclass cNewStructure "cpp2::NewStructure":
        cNewStructure() except +
        cNewStructure(const cNewStructure&) except +
        bint operator==(cNewStructure&)
        cmap[int16_t,double] features
        cNewStructure__isset __isset

    cdef cppclass cNewStructure2__isset "cpp2::NewStructure2::__isset":
        bint features

    cdef cppclass cNewStructure2 "cpp2::NewStructure2":
        cNewStructure2() except +
        cNewStructure2(const cNewStructure2&) except +
        bint operator==(cNewStructure2&)
        cmap[int16_t,float] features
        cNewStructure2__isset __isset

    cdef cppclass cNewStructureNested__isset "cpp2::NewStructureNested::__isset":
        bint lst
        bint mp
        bint s

    cdef cppclass cNewStructureNested "cpp2::NewStructureNested":
        cNewStructureNested() except +
        cNewStructureNested(const cNewStructureNested&) except +
        bint operator==(cNewStructureNested&)
        vector[cmap[int16_t,float]] lst
        cmap[int16_t,cmap[int16_t,float]] mp
        cset[cmap[int16_t,float]] s
        cNewStructureNested__isset __isset

    cdef cppclass cNewStructureNestedField__isset "cpp2::NewStructureNestedField::__isset":
        bint f

    cdef cppclass cNewStructureNestedField "cpp2::NewStructureNestedField":
        cNewStructureNestedField() except +
        cNewStructureNestedField(const cNewStructureNestedField&) except +
        bint operator==(cNewStructureNestedField&)
        cNewStructureNested f
        cNewStructureNestedField__isset __isset


cdef extern from "<utility>" namespace "std" nogil:
    cdef shared_ptr[cOldStructure] move(unique_ptr[cOldStructure])
    cdef shared_ptr[cOldStructure] move_shared "std::move"(shared_ptr[cOldStructure])
    cdef unique_ptr[cOldStructure] move_unique "std::move"(unique_ptr[cOldStructure])
    cdef shared_ptr[cNewStructure] move(unique_ptr[cNewStructure])
    cdef shared_ptr[cNewStructure] move_shared "std::move"(shared_ptr[cNewStructure])
    cdef unique_ptr[cNewStructure] move_unique "std::move"(unique_ptr[cNewStructure])
    cdef shared_ptr[cNewStructure2] move(unique_ptr[cNewStructure2])
    cdef shared_ptr[cNewStructure2] move_shared "std::move"(shared_ptr[cNewStructure2])
    cdef unique_ptr[cNewStructure2] move_unique "std::move"(unique_ptr[cNewStructure2])
    cdef shared_ptr[cNewStructureNested] move(unique_ptr[cNewStructureNested])
    cdef shared_ptr[cNewStructureNested] move_shared "std::move"(shared_ptr[cNewStructureNested])
    cdef unique_ptr[cNewStructureNested] move_unique "std::move"(unique_ptr[cNewStructureNested])
    cdef shared_ptr[cNewStructureNestedField] move(unique_ptr[cNewStructureNestedField])
    cdef shared_ptr[cNewStructureNestedField] move_shared "std::move"(shared_ptr[cNewStructureNestedField])
    cdef unique_ptr[cNewStructureNestedField] move_unique "std::move"(unique_ptr[cNewStructureNestedField])

cdef extern from "<memory>" namespace "std" nogil:
    cdef shared_ptr[const cOldStructure] const_pointer_cast "std::const_pointer_cast<const cpp2::OldStructure>"(shared_ptr[cOldStructure])
    cdef shared_ptr[const cNewStructure] const_pointer_cast "std::const_pointer_cast<const cpp2::NewStructure>"(shared_ptr[cNewStructure])
    cdef shared_ptr[const cNewStructure2] const_pointer_cast "std::const_pointer_cast<const cpp2::NewStructure2>"(shared_ptr[cNewStructure2])
    cdef shared_ptr[const cNewStructureNested] const_pointer_cast "std::const_pointer_cast<const cpp2::NewStructureNested>"(shared_ptr[cNewStructureNested])
    cdef shared_ptr[const cNewStructureNestedField] const_pointer_cast "std::const_pointer_cast<const cpp2::NewStructureNestedField>"(shared_ptr[cNewStructureNestedField])

# Forward Definition of the cython struct
cdef class OldStructure(thrift.py3.types.Struct)

cdef class OldStructure(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cOldStructure] _cpp_obj
    cdef Map__i16_double __features

    @staticmethod
    cdef unique_ptr[cOldStructure] _make_instance(
        cOldStructure* base_instance,
        object features
    ) except *

    @staticmethod
    cdef create(shared_ptr[cOldStructure])

# Forward Definition of the cython struct
cdef class NewStructure(thrift.py3.types.Struct)

cdef class NewStructure(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cNewStructure] _cpp_obj
    cdef Map__i16_double __features

    @staticmethod
    cdef unique_ptr[cNewStructure] _make_instance(
        cNewStructure* base_instance,
        object features
    ) except *

    @staticmethod
    cdef create(shared_ptr[cNewStructure])

# Forward Definition of the cython struct
cdef class NewStructure2(thrift.py3.types.Struct)

cdef class NewStructure2(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cNewStructure2] _cpp_obj
    cdef Map__i16_float __features

    @staticmethod
    cdef unique_ptr[cNewStructure2] _make_instance(
        cNewStructure2* base_instance,
        object features
    ) except *

    @staticmethod
    cdef create(shared_ptr[cNewStructure2])

# Forward Definition of the cython struct
cdef class NewStructureNested(thrift.py3.types.Struct)

cdef class NewStructureNested(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cNewStructureNested] _cpp_obj
    cdef List__Map__i16_float __lst
    cdef Map__i16_Map__i16_float __mp
    cdef Set__Map__i16_float __s

    @staticmethod
    cdef unique_ptr[cNewStructureNested] _make_instance(
        cNewStructureNested* base_instance,
        object lst,
        object mp,
        object s
    ) except *

    @staticmethod
    cdef create(shared_ptr[cNewStructureNested])

# Forward Definition of the cython struct
cdef class NewStructureNestedField(thrift.py3.types.Struct)

cdef class NewStructureNestedField(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cNewStructureNestedField] _cpp_obj
    cdef NewStructureNested __f

    @staticmethod
    cdef unique_ptr[cNewStructureNestedField] _make_instance(
        cNewStructureNestedField* base_instance,
        object f
    ) except *

    @staticmethod
    cdef create(shared_ptr[cNewStructureNestedField])


cdef class Map__i16_double:
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cmap[int16_t,double]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[int16_t,double]])
    @staticmethod
    cdef unique_ptr[cmap[int16_t,double]] _make_instance(object items) except *

cdef class Map__i16_float:
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cmap[int16_t,float]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[int16_t,float]])
    @staticmethod
    cdef unique_ptr[cmap[int16_t,float]] _make_instance(object items) except *

cdef class List__Map__i16_float:
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[vector[cmap[int16_t,float]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[vector[cmap[int16_t,float]]])
    @staticmethod
    cdef unique_ptr[vector[cmap[int16_t,float]]] _make_instance(object items) except *

cdef class Map__i16_Map__i16_float:
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cmap[int16_t,cmap[int16_t,float]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cmap[int16_t,cmap[int16_t,float]]])
    @staticmethod
    cdef unique_ptr[cmap[int16_t,cmap[int16_t,float]]] _make_instance(object items) except *

cdef class Set__Map__i16_float:
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cset[cmap[int16_t,float]]] _cpp_obj
    @staticmethod
    cdef create(shared_ptr[cset[cmap[int16_t,float]]])
    @staticmethod
    cdef unique_ptr[cset[cmap[int16_t,float]]] _make_instance(object items) except *

cdef extern from "<utility>" namespace "std" nogil:
    cdef shared_ptr[cmap[int16_t,double]] move(unique_ptr[cmap[int16_t,double]])
    cdef unique_ptr[cmap[int16_t,double]] move_unique "std::move"(unique_ptr[cmap[int16_t,double]])
    cdef shared_ptr[cmap[int16_t,float]] move(unique_ptr[cmap[int16_t,float]])
    cdef unique_ptr[cmap[int16_t,float]] move_unique "std::move"(unique_ptr[cmap[int16_t,float]])
    cdef shared_ptr[vector[cmap[int16_t,float]]] move(unique_ptr[vector[cmap[int16_t,float]]])
    cdef unique_ptr[vector[cmap[int16_t,float]]] move_unique "std::move"(unique_ptr[vector[cmap[int16_t,float]]])
    cdef shared_ptr[cmap[int16_t,cmap[int16_t,float]]] move(unique_ptr[cmap[int16_t,cmap[int16_t,float]]])
    cdef unique_ptr[cmap[int16_t,cmap[int16_t,float]]] move_unique "std::move"(unique_ptr[cmap[int16_t,cmap[int16_t,float]]])
    cdef shared_ptr[cset[cmap[int16_t,float]]] move(unique_ptr[cset[cmap[int16_t,float]]])
    cdef unique_ptr[cset[cmap[int16_t,float]]] move_unique "std::move"(unique_ptr[cset[cmap[int16_t,float]]])
cdef extern from "<memory>" namespace "std" nogil:
    cdef shared_ptr[const cmap[int16_t,double]] const_pointer_cast "std::const_pointer_cast"(shared_ptr[cmap[int16_t,double]])

    cdef shared_ptr[const cmap[int16_t,float]] const_pointer_cast "std::const_pointer_cast"(shared_ptr[cmap[int16_t,float]])

    cdef shared_ptr[const vector[cmap[int16_t,float]]] const_pointer_cast "std::const_pointer_cast"(shared_ptr[vector[cmap[int16_t,float]]])

    cdef shared_ptr[const cmap[int16_t,cmap[int16_t,float]]] const_pointer_cast "std::const_pointer_cast"(shared_ptr[cmap[int16_t,cmap[int16_t,float]]])

    cdef shared_ptr[const cset[cmap[int16_t,float]]] const_pointer_cast "std::const_pointer_cast"(shared_ptr[cset[cmap[int16_t,float]]])

