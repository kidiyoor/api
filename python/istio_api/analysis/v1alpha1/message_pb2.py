# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: analysis/v1alpha1/message.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='analysis/v1alpha1/message.proto',
  package='istio.analysis.v1alpha1',
  syntax='proto3',
  serialized_options=_b('Z\036istio.io/api/analysis/v1alpha1'),
  serialized_pb=_b('\n\x1f\x61nalysis/v1alpha1/message.proto\x12\x17istio.analysis.v1alpha1\x1a\x1cgoogle/protobuf/struct.proto\"\x90\x02\n\x13\x41nalysisMessageBase\x12?\n\x04type\x18\x01 \x01(\x0b\x32\x31.istio.analysis.v1alpha1.AnalysisMessageBase.Type\x12\x41\n\x05level\x18\x02 \x01(\x0e\x32\x32.istio.analysis.v1alpha1.AnalysisMessageBase.Level\x12\x19\n\x11\x64ocumentation_url\x18\x03 \x01(\t\x1a\"\n\x04Type\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\"6\n\x05Level\x12\x0b\n\x07UNKNOWN\x10\x00\x12\t\n\x05\x45RROR\x10\x03\x12\x0b\n\x07WARNING\x10\x08\x12\x08\n\x04INFO\x10\x0c\"\xfa\x01\n\x19\x41nalysisMessageWeakSchema\x12\x42\n\x0cmessage_base\x18\x01 \x01(\x0b\x32,.istio.analysis.v1alpha1.AnalysisMessageBase\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x10\n\x08template\x18\x03 \x01(\t\x12H\n\x04\x61rgs\x18\x04 \x03(\x0b\x32:.istio.analysis.v1alpha1.AnalysisMessageWeakSchema.ArgType\x1a(\n\x07\x41rgType\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07go_type\x18\x02 \x01(\t\"\x9b\x01\n\x16GenericAnalysisMessage\x12\x42\n\x0cmessage_base\x18\x01 \x01(\x0b\x32,.istio.analysis.v1alpha1.AnalysisMessageBase\x12%\n\x04\x61rgs\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x16\n\x0eresource_paths\x18\x03 \x03(\t\"r\n\x1cInternalErrorAnalysisMessage\x12\x42\n\x0cmessage_base\x18\x01 \x01(\x0b\x32,.istio.analysis.v1alpha1.AnalysisMessageBase\x12\x0e\n\x06\x64\x65tail\x18\x02 \x01(\tB Z\x1eistio.io/api/analysis/v1alpha1b\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,])



_ANALYSISMESSAGEBASE_LEVEL = _descriptor.EnumDescriptor(
  name='Level',
  full_name='istio.analysis.v1alpha1.AnalysisMessageBase.Level',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ERROR', index=1, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WARNING', index=2, number=8,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INFO', index=3, number=12,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=309,
  serialized_end=363,
)
_sym_db.RegisterEnumDescriptor(_ANALYSISMESSAGEBASE_LEVEL)


_ANALYSISMESSAGEBASE_TYPE = _descriptor.Descriptor(
  name='Type',
  full_name='istio.analysis.v1alpha1.AnalysisMessageBase.Type',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.analysis.v1alpha1.AnalysisMessageBase.Type.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='code', full_name='istio.analysis.v1alpha1.AnalysisMessageBase.Type.code', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=273,
  serialized_end=307,
)

_ANALYSISMESSAGEBASE = _descriptor.Descriptor(
  name='AnalysisMessageBase',
  full_name='istio.analysis.v1alpha1.AnalysisMessageBase',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='istio.analysis.v1alpha1.AnalysisMessageBase.type', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='level', full_name='istio.analysis.v1alpha1.AnalysisMessageBase.level', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='documentation_url', full_name='istio.analysis.v1alpha1.AnalysisMessageBase.documentation_url', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ANALYSISMESSAGEBASE_TYPE, ],
  enum_types=[
    _ANALYSISMESSAGEBASE_LEVEL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=91,
  serialized_end=363,
)


_ANALYSISMESSAGEWEAKSCHEMA_ARGTYPE = _descriptor.Descriptor(
  name='ArgType',
  full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.ArgType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.ArgType.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='go_type', full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.ArgType.go_type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=576,
  serialized_end=616,
)

_ANALYSISMESSAGEWEAKSCHEMA = _descriptor.Descriptor(
  name='AnalysisMessageWeakSchema',
  full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message_base', full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.message_base', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.description', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='template', full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.template', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='istio.analysis.v1alpha1.AnalysisMessageWeakSchema.args', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ANALYSISMESSAGEWEAKSCHEMA_ARGTYPE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=366,
  serialized_end=616,
)


_GENERICANALYSISMESSAGE = _descriptor.Descriptor(
  name='GenericAnalysisMessage',
  full_name='istio.analysis.v1alpha1.GenericAnalysisMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message_base', full_name='istio.analysis.v1alpha1.GenericAnalysisMessage.message_base', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='istio.analysis.v1alpha1.GenericAnalysisMessage.args', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_paths', full_name='istio.analysis.v1alpha1.GenericAnalysisMessage.resource_paths', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=619,
  serialized_end=774,
)


_INTERNALERRORANALYSISMESSAGE = _descriptor.Descriptor(
  name='InternalErrorAnalysisMessage',
  full_name='istio.analysis.v1alpha1.InternalErrorAnalysisMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message_base', full_name='istio.analysis.v1alpha1.InternalErrorAnalysisMessage.message_base', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='detail', full_name='istio.analysis.v1alpha1.InternalErrorAnalysisMessage.detail', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=776,
  serialized_end=890,
)

_ANALYSISMESSAGEBASE_TYPE.containing_type = _ANALYSISMESSAGEBASE
_ANALYSISMESSAGEBASE.fields_by_name['type'].message_type = _ANALYSISMESSAGEBASE_TYPE
_ANALYSISMESSAGEBASE.fields_by_name['level'].enum_type = _ANALYSISMESSAGEBASE_LEVEL
_ANALYSISMESSAGEBASE_LEVEL.containing_type = _ANALYSISMESSAGEBASE
_ANALYSISMESSAGEWEAKSCHEMA_ARGTYPE.containing_type = _ANALYSISMESSAGEWEAKSCHEMA
_ANALYSISMESSAGEWEAKSCHEMA.fields_by_name['message_base'].message_type = _ANALYSISMESSAGEBASE
_ANALYSISMESSAGEWEAKSCHEMA.fields_by_name['args'].message_type = _ANALYSISMESSAGEWEAKSCHEMA_ARGTYPE
_GENERICANALYSISMESSAGE.fields_by_name['message_base'].message_type = _ANALYSISMESSAGEBASE
_GENERICANALYSISMESSAGE.fields_by_name['args'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_INTERNALERRORANALYSISMESSAGE.fields_by_name['message_base'].message_type = _ANALYSISMESSAGEBASE
DESCRIPTOR.message_types_by_name['AnalysisMessageBase'] = _ANALYSISMESSAGEBASE
DESCRIPTOR.message_types_by_name['AnalysisMessageWeakSchema'] = _ANALYSISMESSAGEWEAKSCHEMA
DESCRIPTOR.message_types_by_name['GenericAnalysisMessage'] = _GENERICANALYSISMESSAGE
DESCRIPTOR.message_types_by_name['InternalErrorAnalysisMessage'] = _INTERNALERRORANALYSISMESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AnalysisMessageBase = _reflection.GeneratedProtocolMessageType('AnalysisMessageBase', (_message.Message,), {

  'Type' : _reflection.GeneratedProtocolMessageType('Type', (_message.Message,), {
    'DESCRIPTOR' : _ANALYSISMESSAGEBASE_TYPE,
    '__module__' : 'analysis.v1alpha1.message_pb2'
    # @@protoc_insertion_point(class_scope:istio.analysis.v1alpha1.AnalysisMessageBase.Type)
    })
  ,
  'DESCRIPTOR' : _ANALYSISMESSAGEBASE,
  '__module__' : 'analysis.v1alpha1.message_pb2'
  # @@protoc_insertion_point(class_scope:istio.analysis.v1alpha1.AnalysisMessageBase)
  })
_sym_db.RegisterMessage(AnalysisMessageBase)
_sym_db.RegisterMessage(AnalysisMessageBase.Type)

AnalysisMessageWeakSchema = _reflection.GeneratedProtocolMessageType('AnalysisMessageWeakSchema', (_message.Message,), {

  'ArgType' : _reflection.GeneratedProtocolMessageType('ArgType', (_message.Message,), {
    'DESCRIPTOR' : _ANALYSISMESSAGEWEAKSCHEMA_ARGTYPE,
    '__module__' : 'analysis.v1alpha1.message_pb2'
    # @@protoc_insertion_point(class_scope:istio.analysis.v1alpha1.AnalysisMessageWeakSchema.ArgType)
    })
  ,
  'DESCRIPTOR' : _ANALYSISMESSAGEWEAKSCHEMA,
  '__module__' : 'analysis.v1alpha1.message_pb2'
  # @@protoc_insertion_point(class_scope:istio.analysis.v1alpha1.AnalysisMessageWeakSchema)
  })
_sym_db.RegisterMessage(AnalysisMessageWeakSchema)
_sym_db.RegisterMessage(AnalysisMessageWeakSchema.ArgType)

GenericAnalysisMessage = _reflection.GeneratedProtocolMessageType('GenericAnalysisMessage', (_message.Message,), {
  'DESCRIPTOR' : _GENERICANALYSISMESSAGE,
  '__module__' : 'analysis.v1alpha1.message_pb2'
  # @@protoc_insertion_point(class_scope:istio.analysis.v1alpha1.GenericAnalysisMessage)
  })
_sym_db.RegisterMessage(GenericAnalysisMessage)

InternalErrorAnalysisMessage = _reflection.GeneratedProtocolMessageType('InternalErrorAnalysisMessage', (_message.Message,), {
  'DESCRIPTOR' : _INTERNALERRORANALYSISMESSAGE,
  '__module__' : 'analysis.v1alpha1.message_pb2'
  # @@protoc_insertion_point(class_scope:istio.analysis.v1alpha1.InternalErrorAnalysisMessage)
  })
_sym_db.RegisterMessage(InternalErrorAnalysisMessage)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
