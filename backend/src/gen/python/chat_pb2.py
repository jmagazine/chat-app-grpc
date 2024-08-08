# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chat.proto
# Protobuf Python Version: 5.27.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    3,
    '',
    'chat.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nchat.proto\x12\x04\x63hat\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"O\n\x04User\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n\tfull_name\x18\x02 \x01(\tR\x08\x66ullName\x12\x1a\n\x08username\x18\x03 \x01(\tR\x08username\"k\n\x11\x43reateUserRequest\x12\x1b\n\tfull_name\x18\x01 \x01(\tR\x08\x66ullName\x12\x1a\n\x08username\x18\x02 \x01(\tR\x08username\x12\x1d\n\nhash_token\x18\x03 \x01(\tR\thashToken\"4\n\x12\x43reateUserResponse\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\n.chat.UserR\x04user\"/\n\x11\x44\x65leteUserRequest\x12\x1a\n\x08username\x18\x01 \x01(\tR\x08username\"4\n\x12\x44\x65leteUserResponse\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\n.chat.UserR\x04user\"K\n\x0eGetUserRequest\x12\x1a\n\x08username\x18\x01 \x01(\tR\x08username\x12\x1d\n\nhash_token\x18\x02 \x01(\tR\thashToken\"1\n\x0fGetUserResponse\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\n.chat.UserR\x04user\"\xc4\x01\n\x11UpdateUserRequest\x12\x1a\n\x08username\x18\x01 \x01(\tR\x08username\x12Q\n\x0eupdated_fields\x18\x02 \x03(\x0b\x32*.chat.UpdateUserRequest.UpdatedFieldsEntryR\rupdatedFields\x1a@\n\x12UpdatedFieldsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"4\n\x12UpdateUserResponse\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\n.chat.UserR\x04user\"\x14\n\x12GetAllUsersRequest\"7\n\x13GetAllUsersResponse\x12 \n\x05users\x18\x01 \x03(\x0b\x32\n.chat.UserR\x05users\"\x91\x01\n\x0b\x43hatMessage\x12\x38\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x12\x16\n\x06sender\x18\x02 \x01(\tR\x06sender\x12\x1c\n\trecipient\x18\x03 \x01(\tR\trecipient\x12\x12\n\x04text\x18\x04 \x01(\tR\x04text\"X\n\x16GetChatMessagesRequest\x12\x1b\n\tsender_id\x18\x01 \x01(\tR\x08senderId\x12!\n\x0crecipient_id\x18\x02 \x01(\tR\x0brecipientId\"H\n\x17GetChatMessagesResponse\x12-\n\x08messages\x18\x01 \x03(\x0b\x32\x11.chat.ChatMessageR\x08messages\"E\n\x16SendChatMessageRequest\x12+\n\x07message\x18\x01 \x01(\x0b\x32\x11.chat.ChatMessageR\x07message\"H\n\x17SendChatMessageResponse\x12-\n\x08messages\x18\x01 \x01(\x0b\x32\x11.chat.ChatMessageR\x08messages\"1\n\x10\x44ropTableRequest\x12\x1d\n\ntable_name\x18\x01 \x01(\tR\ttableName\"\x13\n\x11\x44ropTableResponse\"\x17\n\x15\x44\x65leteAllUsersRequest\"\x18\n\x16\x44\x65leteAllUsersResponse2\xf0\x06\n\x0b\x43hatService\x12Z\n\nCreateUser\x12\x17.chat.CreateUserRequest\x1a\x18.chat.CreateUserResponse\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/v1/CreateUser:\x01*\x12Z\n\nDeleteUser\x12\x17.chat.DeleteUserRequest\x1a\x18.chat.DeleteUserResponse\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/v1/DeleteUser:\x01*\x12Z\n\nUpdateUser\x12\x17.chat.UpdateUserRequest\x1a\x18.chat.UpdateUserResponse\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/v1/UpdateUser:\x01*\x12[\n\x0bGetAllUsers\x12\x18.chat.GetAllUsersRequest\x1a\x19.chat.GetAllUsersResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/v1/GetAllUsers\x12M\n\x07GetUser\x12\x14.chat.GetUserRequest\x1a\x15.chat.GetUserResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/v1/GetUser/*\x12m\n\x0fGetChatMessages\x12\x1c.chat.GetChatMessagesRequest\x1a\x1d.chat.GetChatMessagesResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/GetChatMessages/*\x12n\n\x0fSendChatMessage\x12\x1c.chat.SendChatMessageRequest\x1a\x1d.chat.SendChatMessageResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/SendChatMessage:\x01*\x12j\n\x0e\x44\x65leteAllUsers\x12\x1b.chat.DeleteAllUsersRequest\x1a\x1c.chat.DeleteAllUsersResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\"\x12/v1/DeleteAllUsers:\x01*\x12V\n\tDropTable\x12\x16.chat.DropTableRequest\x1a\x17.chat.DropTableResponse\"\x18\x82\xd3\xe4\x93\x02\x12\"\r/v1/DropTable:\x01*Bm\n\x08\x63om.chatB\tChatProtoP\x01Z&github.com/chat-app-grpc;chat_app_grpc\xa2\x02\x03\x43XX\xaa\x02\x04\x43hat\xca\x02\x04\x43hat\xe2\x02\x10\x43hat\\GPBMetadata\xea\x02\x04\x43hatb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chat_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\010com.chatB\tChatProtoP\001Z&github.com/chat-app-grpc;chat_app_grpc\242\002\003CXX\252\002\004Chat\312\002\004Chat\342\002\020Chat\\GPBMetadata\352\002\004Chat'
  _globals['_UPDATEUSERREQUEST_UPDATEDFIELDSENTRY']._loaded_options = None
  _globals['_UPDATEUSERREQUEST_UPDATEDFIELDSENTRY']._serialized_options = b'8\001'
  _globals['_CHATSERVICE'].methods_by_name['CreateUser']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['CreateUser']._serialized_options = b'\202\323\344\223\002\023\"\016/v1/CreateUser:\001*'
  _globals['_CHATSERVICE'].methods_by_name['DeleteUser']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['DeleteUser']._serialized_options = b'\202\323\344\223\002\023\"\016/v1/DeleteUser:\001*'
  _globals['_CHATSERVICE'].methods_by_name['UpdateUser']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['UpdateUser']._serialized_options = b'\202\323\344\223\002\023\"\016/v1/UpdateUser:\001*'
  _globals['_CHATSERVICE'].methods_by_name['GetAllUsers']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['GetAllUsers']._serialized_options = b'\202\323\344\223\002\021\022\017/v1/GetAllUsers'
  _globals['_CHATSERVICE'].methods_by_name['GetUser']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['GetUser']._serialized_options = b'\202\323\344\223\002\017\022\r/v1/GetUser/*'
  _globals['_CHATSERVICE'].methods_by_name['GetChatMessages']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['GetChatMessages']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/GetChatMessages/*'
  _globals['_CHATSERVICE'].methods_by_name['SendChatMessage']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['SendChatMessage']._serialized_options = b'\202\323\344\223\002\030\"\023/v1/SendChatMessage:\001*'
  _globals['_CHATSERVICE'].methods_by_name['DeleteAllUsers']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['DeleteAllUsers']._serialized_options = b'\202\323\344\223\002\027\"\022/v1/DeleteAllUsers:\001*'
  _globals['_CHATSERVICE'].methods_by_name['DropTable']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['DropTable']._serialized_options = b'\202\323\344\223\002\022\"\r/v1/DropTable:\001*'
  _globals['_USER']._serialized_start=83
  _globals['_USER']._serialized_end=162
  _globals['_CREATEUSERREQUEST']._serialized_start=164
  _globals['_CREATEUSERREQUEST']._serialized_end=271
  _globals['_CREATEUSERRESPONSE']._serialized_start=273
  _globals['_CREATEUSERRESPONSE']._serialized_end=325
  _globals['_DELETEUSERREQUEST']._serialized_start=327
  _globals['_DELETEUSERREQUEST']._serialized_end=374
  _globals['_DELETEUSERRESPONSE']._serialized_start=376
  _globals['_DELETEUSERRESPONSE']._serialized_end=428
  _globals['_GETUSERREQUEST']._serialized_start=430
  _globals['_GETUSERREQUEST']._serialized_end=505
  _globals['_GETUSERRESPONSE']._serialized_start=507
  _globals['_GETUSERRESPONSE']._serialized_end=556
  _globals['_UPDATEUSERREQUEST']._serialized_start=559
  _globals['_UPDATEUSERREQUEST']._serialized_end=755
  _globals['_UPDATEUSERREQUEST_UPDATEDFIELDSENTRY']._serialized_start=691
  _globals['_UPDATEUSERREQUEST_UPDATEDFIELDSENTRY']._serialized_end=755
  _globals['_UPDATEUSERRESPONSE']._serialized_start=757
  _globals['_UPDATEUSERRESPONSE']._serialized_end=809
  _globals['_GETALLUSERSREQUEST']._serialized_start=811
  _globals['_GETALLUSERSREQUEST']._serialized_end=831
  _globals['_GETALLUSERSRESPONSE']._serialized_start=833
  _globals['_GETALLUSERSRESPONSE']._serialized_end=888
  _globals['_CHATMESSAGE']._serialized_start=891
  _globals['_CHATMESSAGE']._serialized_end=1036
  _globals['_GETCHATMESSAGESREQUEST']._serialized_start=1038
  _globals['_GETCHATMESSAGESREQUEST']._serialized_end=1126
  _globals['_GETCHATMESSAGESRESPONSE']._serialized_start=1128
  _globals['_GETCHATMESSAGESRESPONSE']._serialized_end=1200
  _globals['_SENDCHATMESSAGEREQUEST']._serialized_start=1202
  _globals['_SENDCHATMESSAGEREQUEST']._serialized_end=1271
  _globals['_SENDCHATMESSAGERESPONSE']._serialized_start=1273
  _globals['_SENDCHATMESSAGERESPONSE']._serialized_end=1345
  _globals['_DROPTABLEREQUEST']._serialized_start=1347
  _globals['_DROPTABLEREQUEST']._serialized_end=1396
  _globals['_DROPTABLERESPONSE']._serialized_start=1398
  _globals['_DROPTABLERESPONSE']._serialized_end=1417
  _globals['_DELETEALLUSERSREQUEST']._serialized_start=1419
  _globals['_DELETEALLUSERSREQUEST']._serialized_end=1442
  _globals['_DELETEALLUSERSRESPONSE']._serialized_start=1444
  _globals['_DELETEALLUSERSRESPONSE']._serialized_end=1468
  _globals['_CHATSERVICE']._serialized_start=1471
  _globals['_CHATSERVICE']._serialized_end=2351
# @@protoc_insertion_point(module_scope)