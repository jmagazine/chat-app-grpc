/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.20.1
// source: chat.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as chat_pb from './chat_pb'; // proto import: "chat.proto"


export class ChatServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreateNewUser = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/CreateNewUser',
    grpcWeb.MethodType.UNARY,
    chat_pb.CreateUserParams,
    chat_pb.User,
    (request: chat_pb.CreateUserParams) => {
      return request.serializeBinary();
    },
    chat_pb.User.deserializeBinary
  );

  createNewUser(
    request: chat_pb.CreateUserParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.User>;

  createNewUser(
    request: chat_pb.CreateUserParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.User) => void): grpcWeb.ClientReadableStream<chat_pb.User>;

  createNewUser(
    request: chat_pb.CreateUserParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/CreateNewUser',
        request,
        metadata || {},
        this.methodDescriptorCreateNewUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/CreateNewUser',
    request,
    metadata || {},
    this.methodDescriptorCreateNewUser);
  }

  methodDescriptorDeleteUserByUsername = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/DeleteUserByUsername',
    grpcWeb.MethodType.UNARY,
    chat_pb.DeleteUserByUsernameParams,
    chat_pb.DidDeleteUserMessage,
    (request: chat_pb.DeleteUserByUsernameParams) => {
      return request.serializeBinary();
    },
    chat_pb.DidDeleteUserMessage.deserializeBinary
  );

  deleteUserByUsername(
    request: chat_pb.DeleteUserByUsernameParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.DidDeleteUserMessage>;

  deleteUserByUsername(
    request: chat_pb.DeleteUserByUsernameParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.DidDeleteUserMessage) => void): grpcWeb.ClientReadableStream<chat_pb.DidDeleteUserMessage>;

  deleteUserByUsername(
    request: chat_pb.DeleteUserByUsernameParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.DidDeleteUserMessage) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/DeleteUserByUsername',
        request,
        metadata || {},
        this.methodDescriptorDeleteUserByUsername,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/DeleteUserByUsername',
    request,
    metadata || {},
    this.methodDescriptorDeleteUserByUsername);
  }

  methodDescriptorUpdateUser = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/UpdateUser',
    grpcWeb.MethodType.UNARY,
    chat_pb.UpdateUserParams,
    chat_pb.User,
    (request: chat_pb.UpdateUserParams) => {
      return request.serializeBinary();
    },
    chat_pb.User.deserializeBinary
  );

  updateUser(
    request: chat_pb.UpdateUserParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.User>;

  updateUser(
    request: chat_pb.UpdateUserParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.User) => void): grpcWeb.ClientReadableStream<chat_pb.User>;

  updateUser(
    request: chat_pb.UpdateUserParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/UpdateUser',
        request,
        metadata || {},
        this.methodDescriptorUpdateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/UpdateUser',
    request,
    metadata || {},
    this.methodDescriptorUpdateUser);
  }

  methodDescriptorGetAllUsers = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/GetAllUsers',
    grpcWeb.MethodType.UNARY,
    chat_pb.GetAllUsersParams,
    chat_pb.UsersList,
    (request: chat_pb.GetAllUsersParams) => {
      return request.serializeBinary();
    },
    chat_pb.UsersList.deserializeBinary
  );

  getAllUsers(
    request: chat_pb.GetAllUsersParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.UsersList>;

  getAllUsers(
    request: chat_pb.GetAllUsersParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.UsersList) => void): grpcWeb.ClientReadableStream<chat_pb.UsersList>;

  getAllUsers(
    request: chat_pb.GetAllUsersParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.UsersList) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/GetAllUsers',
        request,
        metadata || {},
        this.methodDescriptorGetAllUsers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/GetAllUsers',
    request,
    metadata || {},
    this.methodDescriptorGetAllUsers);
  }

  methodDescriptorGetUserByUsername = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/GetUserByUsername',
    grpcWeb.MethodType.UNARY,
    chat_pb.GetUserByUsernameParams,
    chat_pb.User,
    (request: chat_pb.GetUserByUsernameParams) => {
      return request.serializeBinary();
    },
    chat_pb.User.deserializeBinary
  );

  getUserByUsername(
    request: chat_pb.GetUserByUsernameParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.User>;

  getUserByUsername(
    request: chat_pb.GetUserByUsernameParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.User) => void): grpcWeb.ClientReadableStream<chat_pb.User>;

  getUserByUsername(
    request: chat_pb.GetUserByUsernameParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/GetUserByUsername',
        request,
        metadata || {},
        this.methodDescriptorGetUserByUsername,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/GetUserByUsername',
    request,
    metadata || {},
    this.methodDescriptorGetUserByUsername);
  }

  methodDescriptorLogin = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/Login',
    grpcWeb.MethodType.UNARY,
    chat_pb.LoginParams,
    chat_pb.LoginResponse,
    (request: chat_pb.LoginParams) => {
      return request.serializeBinary();
    },
    chat_pb.LoginResponse.deserializeBinary
  );

  login(
    request: chat_pb.LoginParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.LoginResponse>;

  login(
    request: chat_pb.LoginParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<chat_pb.LoginResponse>;

  login(
    request: chat_pb.LoginParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/Login',
        request,
        metadata || {},
        this.methodDescriptorLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/Login',
    request,
    metadata || {},
    this.methodDescriptorLogin);
  }

  methodDescriptorGetChatMessages = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/GetChatMessages',
    grpcWeb.MethodType.UNARY,
    chat_pb.GetChatMessagesParams,
    chat_pb.ChatMessageList,
    (request: chat_pb.GetChatMessagesParams) => {
      return request.serializeBinary();
    },
    chat_pb.ChatMessageList.deserializeBinary
  );

  getChatMessages(
    request: chat_pb.GetChatMessagesParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.ChatMessageList>;

  getChatMessages(
    request: chat_pb.GetChatMessagesParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.ChatMessageList) => void): grpcWeb.ClientReadableStream<chat_pb.ChatMessageList>;

  getChatMessages(
    request: chat_pb.GetChatMessagesParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.ChatMessageList) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/GetChatMessages',
        request,
        metadata || {},
        this.methodDescriptorGetChatMessages,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/GetChatMessages',
    request,
    metadata || {},
    this.methodDescriptorGetChatMessages);
  }

  methodDescriptorSendChatMessage = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/SendChatMessage',
    grpcWeb.MethodType.UNARY,
    chat_pb.SendChatMessageParams,
    chat_pb.ChatMessage,
    (request: chat_pb.SendChatMessageParams) => {
      return request.serializeBinary();
    },
    chat_pb.ChatMessage.deserializeBinary
  );

  sendChatMessage(
    request: chat_pb.SendChatMessageParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.ChatMessage>;

  sendChatMessage(
    request: chat_pb.SendChatMessageParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.ChatMessage) => void): grpcWeb.ClientReadableStream<chat_pb.ChatMessage>;

  sendChatMessage(
    request: chat_pb.SendChatMessageParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.ChatMessage) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/SendChatMessage',
        request,
        metadata || {},
        this.methodDescriptorSendChatMessage,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/SendChatMessage',
    request,
    metadata || {},
    this.methodDescriptorSendChatMessage);
  }

  methodDescriptorDropTable = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/DropTable',
    grpcWeb.MethodType.UNARY,
    chat_pb.DropTableParams,
    chat_pb.DropTableMessage,
    (request: chat_pb.DropTableParams) => {
      return request.serializeBinary();
    },
    chat_pb.DropTableMessage.deserializeBinary
  );

  dropTable(
    request: chat_pb.DropTableParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.DropTableMessage>;

  dropTable(
    request: chat_pb.DropTableParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.DropTableMessage) => void): grpcWeb.ClientReadableStream<chat_pb.DropTableMessage>;

  dropTable(
    request: chat_pb.DropTableParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.DropTableMessage) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/DropTable',
        request,
        metadata || {},
        this.methodDescriptorDropTable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/DropTable',
    request,
    metadata || {},
    this.methodDescriptorDropTable);
  }

  methodDescriptorGetServer = new grpcWeb.MethodDescriptor(
    '/chat.ChatService/GetServer',
    grpcWeb.MethodType.UNARY,
    chat_pb.GetServerParams,
    chat_pb.Server,
    (request: chat_pb.GetServerParams) => {
      return request.serializeBinary();
    },
    chat_pb.Server.deserializeBinary
  );

  getServer(
    request: chat_pb.GetServerParams,
    metadata?: grpcWeb.Metadata | null): Promise<chat_pb.Server>;

  getServer(
    request: chat_pb.GetServerParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: chat_pb.Server) => void): grpcWeb.ClientReadableStream<chat_pb.Server>;

  getServer(
    request: chat_pb.GetServerParams,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: chat_pb.Server) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.ChatService/GetServer',
        request,
        metadata || {},
        this.methodDescriptorGetServer,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.ChatService/GetServer',
    request,
    metadata || {},
    this.methodDescriptorGetServer);
  }

}

