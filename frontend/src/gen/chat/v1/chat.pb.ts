/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "..\..\fetch.pb"
import * as GoogleProtobufTimestamp from "..\..\google\protobuf\timestamp.pb"
export type User = {
  id?: string
  fullName?: string
  username?: string
}

export type CreateUserRequest = {
  fullName?: string
  username?: string
  hashToken?: string
}

export type CreateUserResponse = {
  user?: User
}

export type DeleteUserRequest = {
  username?: string
}

export type DeleteUserResponse = {
  user?: User
}

export type GetUserRequest = {
  username?: string
  hashToken?: string
}

export type GetUserResponse = {
  user?: User
}

export type UpdateUserRequest = {
  username?: string
  updatedFields?: {[key: string]: string}
}

export type UpdateUserResponse = {
  user?: User
}

export type GetAllUsersRequest = {
}

export type GetAllUsersResponse = {
  users?: User[]
}

export type ChatMessage = {
  timestamp?: GoogleProtobufTimestamp.Timestamp
  sender?: string
  recipient?: string
  text?: string
}

export type GetChatMessagesRequest = {
  senderId?: string
  recipientId?: string
}

export type GetChatMessagesResponse = {
  messages?: ChatMessage[]
}

export type SendChatMessageRequest = {
  message?: ChatMessage
}

export type SendChatMessageResponse = {
  message?: ChatMessage
}

export type DropTableRequest = {
  tableName?: string
}

export type DropTableResponse = {
}

export type DeleteAllUsersRequest = {
}

export type DeleteAllUsersResponse = {
}

export class ChatService {
  static CreateUser(req: CreateUserRequest, initReq?: fm.InitReq): Promise<CreateUserResponse> {
    return fm.fetchReq<CreateUserRequest, CreateUserResponse>(`/v1/CreateUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteUser(req: DeleteUserRequest, initReq?: fm.InitReq): Promise<DeleteUserResponse> {
    return fm.fetchReq<DeleteUserRequest, DeleteUserResponse>(`/v1/DeleteUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateUser(req: UpdateUserRequest, initReq?: fm.InitReq): Promise<UpdateUserResponse> {
    return fm.fetchReq<UpdateUserRequest, UpdateUserResponse>(`/v1/UpdateUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAllUsers(req: GetAllUsersRequest, initReq?: fm.InitReq): Promise<GetAllUsersResponse> {
    return fm.fetchReq<GetAllUsersRequest, GetAllUsersResponse>(`/v1/GetAllUsers`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetUser(req: GetUserRequest, initReq?: fm.InitReq): Promise<GetUserResponse> {
    return fm.fetchReq<GetUserRequest, GetUserResponse>(`/v1/GetUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetChatMessages(req: GetChatMessagesRequest, initReq?: fm.InitReq): Promise<GetChatMessagesResponse> {
    return fm.fetchReq<GetChatMessagesRequest, GetChatMessagesResponse>(`/v1/GetChatMessages`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static SendChatMessage(req: SendChatMessageRequest, initReq?: fm.InitReq): Promise<SendChatMessageResponse> {
    return fm.fetchReq<SendChatMessageRequest, SendChatMessageResponse>(`/v1/SendChatMessage`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteAllUsers(req: DeleteAllUsersRequest, initReq?: fm.InitReq): Promise<DeleteAllUsersResponse> {
    return fm.fetchReq<DeleteAllUsersRequest, DeleteAllUsersResponse>(`/v1/DeleteAllUsers`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DropTable(req: DropTableRequest, initReq?: fm.InitReq): Promise<DropTableResponse> {
    return fm.fetchReq<DropTableRequest, DropTableResponse>(`/v1/DropTable`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}