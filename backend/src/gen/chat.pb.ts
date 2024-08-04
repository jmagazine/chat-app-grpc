/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from ".\fetch.pb"
import * as GoogleProtobufTimestamp from ".\google\protobuf\timestamp.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type User = {
  id?: string
  fullName?: string
  username?: string
}

export type CreateUserParams = {
  fullName?: string
  username?: string
  hashToken?: string
}

export type DeleteUserByUsernameParams = {
  username?: string
}

export type GetUserParams = {
  username?: string
  hashToken?: string
}

export type UpdateUserParams = {
  username?: string
  updatedFields?: {[key: string]: string}
}

export type GetAllUsersParams = {
}

export type UsersList = {
  users?: User[]
}


type BaseDidDeleteUserMessage = {
  username?: string
  success?: boolean
}

export type DidDeleteUserMessage = BaseDidDeleteUserMessage
  & OneOf<{ error: string }>

export type GetChatMessagesParams = {
  senderId?: string
  recipientId?: string
}

export type SendChatMessageParams = {
  message?: ChatMessage
}

export type ChatMessage = {
  timestamp?: GoogleProtobufTimestamp.Timestamp
  sender?: string
  recipient?: string
  text?: string
}

export type ChatMessageList = {
  messages?: ChatMessage[]
}

export type DropTableParams = {
  tableName?: string
}

export type DropTableMessage = {
  success?: boolean
}

export type DeleteAllUsersParams = {
}

export type DidDeleteAllUsers = {
}

export class ChatService {
  static CreateUser(req: CreateUserParams, initReq?: fm.InitReq): Promise<User> {
    return fm.fetchReq<CreateUserParams, User>(`/v1/createUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteUserByUsername(req: DeleteUserByUsernameParams, initReq?: fm.InitReq): Promise<DidDeleteUserMessage> {
    return fm.fetchReq<DeleteUserByUsernameParams, DidDeleteUserMessage>(`/v1/DeleteUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateUser(req: UpdateUserParams, initReq?: fm.InitReq): Promise<User> {
    return fm.fetchReq<UpdateUserParams, User>(`/v1/UpdateUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAllUsers(req: GetAllUsersParams, initReq?: fm.InitReq): Promise<UsersList> {
    return fm.fetchReq<GetAllUsersParams, UsersList>(`/v1/GetAllUsers`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetUser(req: GetUserParams, initReq?: fm.InitReq): Promise<User> {
    return fm.fetchReq<GetUserParams, User>(`/v1/GetUser`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetChatMessages(req: GetChatMessagesParams, initReq?: fm.InitReq): Promise<ChatMessageList> {
    return fm.fetchReq<GetChatMessagesParams, ChatMessageList>(`/v1/GetChatMessages`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static SendChatMessage(req: SendChatMessageParams, initReq?: fm.InitReq): Promise<ChatMessage> {
    return fm.fetchReq<SendChatMessageParams, ChatMessage>(`/v1/SendChatMessage`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteAllUsers(req: DeleteAllUsersParams, initReq?: fm.InitReq): Promise<DidDeleteAllUsers> {
    return fm.fetchReq<DeleteAllUsersParams, DidDeleteAllUsers>(`/v1/DeleteAllUsers`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DropTable(req: DropTableParams, initReq?: fm.InitReq): Promise<DropTableMessage> {
    return fm.fetchReq<DropTableParams, DropTableMessage>(`/v1/DropTable`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}