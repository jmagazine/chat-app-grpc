import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class User extends jspb.Message {
  getId(): string;
  setId(value: string): User;

  getFullname(): string;
  setFullname(value: string): User;

  getUsername(): string;
  setUsername(value: string): User;

  getPassword(): string;
  setPassword(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    fullname: string,
    username: string,
    password: string,
  }
}

export class CreateUserParams extends jspb.Message {
  getFullname(): string;
  setFullname(value: string): CreateUserParams;

  getUsername(): string;
  setUsername(value: string): CreateUserParams;

  getPassword(): string;
  setPassword(value: string): CreateUserParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserParams.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserParams): CreateUserParams.AsObject;
  static serializeBinaryToWriter(message: CreateUserParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserParams;
  static deserializeBinaryFromReader(message: CreateUserParams, reader: jspb.BinaryReader): CreateUserParams;
}

export namespace CreateUserParams {
  export type AsObject = {
    fullname: string,
    username: string,
    password: string,
  }
}

export class DeleteUserByUsernameParams extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): DeleteUserByUsernameParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserByUsernameParams.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserByUsernameParams): DeleteUserByUsernameParams.AsObject;
  static serializeBinaryToWriter(message: DeleteUserByUsernameParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserByUsernameParams;
  static deserializeBinaryFromReader(message: DeleteUserByUsernameParams, reader: jspb.BinaryReader): DeleteUserByUsernameParams;
}

export namespace DeleteUserByUsernameParams {
  export type AsObject = {
    username: string,
  }
}

export class LoginParams extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): LoginParams;

  getPassword(): string;
  setPassword(value: string): LoginParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginParams.AsObject;
  static toObject(includeInstance: boolean, msg: LoginParams): LoginParams.AsObject;
  static serializeBinaryToWriter(message: LoginParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginParams;
  static deserializeBinaryFromReader(message: LoginParams, reader: jspb.BinaryReader): LoginParams;
}

export namespace LoginParams {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginResponse extends jspb.Message {
  getResponseCode(): number;
  setResponseCode(value: number): LoginResponse;

  getResponse(): string;
  setResponse(value: string): LoginResponse;
  hasResponse(): boolean;
  clearResponse(): LoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    responseCode: number,
    response?: string,
  }

  export enum ResponseCase { 
    _RESPONSE_NOT_SET = 0,
    RESPONSE = 2,
  }
}

export class UpdateUserParams extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): UpdateUserParams;

  getUpdatedFieldsMap(): jspb.Map<string, string>;
  clearUpdatedFieldsMap(): UpdateUserParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserParams.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserParams): UpdateUserParams.AsObject;
  static serializeBinaryToWriter(message: UpdateUserParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserParams;
  static deserializeBinaryFromReader(message: UpdateUserParams, reader: jspb.BinaryReader): UpdateUserParams;
}

export namespace UpdateUserParams {
  export type AsObject = {
    username: string,
    updatedFieldsMap: Array<[string, string]>,
  }
}

export class GetUserByUsernameParams extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): GetUserByUsernameParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByUsernameParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByUsernameParams): GetUserByUsernameParams.AsObject;
  static serializeBinaryToWriter(message: GetUserByUsernameParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByUsernameParams;
  static deserializeBinaryFromReader(message: GetUserByUsernameParams, reader: jspb.BinaryReader): GetUserByUsernameParams;
}

export namespace GetUserByUsernameParams {
  export type AsObject = {
    username: string,
  }
}

export class GetAllUsersParams extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllUsersParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllUsersParams): GetAllUsersParams.AsObject;
  static serializeBinaryToWriter(message: GetAllUsersParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllUsersParams;
  static deserializeBinaryFromReader(message: GetAllUsersParams, reader: jspb.BinaryReader): GetAllUsersParams;
}

export namespace GetAllUsersParams {
  export type AsObject = {
  }
}

export class UsersList extends jspb.Message {
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): UsersList;
  clearUsersList(): UsersList;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UsersList.AsObject;
  static toObject(includeInstance: boolean, msg: UsersList): UsersList.AsObject;
  static serializeBinaryToWriter(message: UsersList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UsersList;
  static deserializeBinaryFromReader(message: UsersList, reader: jspb.BinaryReader): UsersList;
}

export namespace UsersList {
  export type AsObject = {
    usersList: Array<User.AsObject>,
  }
}

export class DidDeleteUserMessage extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): DidDeleteUserMessage;

  getSuccess(): boolean;
  setSuccess(value: boolean): DidDeleteUserMessage;

  getError(): string;
  setError(value: string): DidDeleteUserMessage;
  hasError(): boolean;
  clearError(): DidDeleteUserMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DidDeleteUserMessage.AsObject;
  static toObject(includeInstance: boolean, msg: DidDeleteUserMessage): DidDeleteUserMessage.AsObject;
  static serializeBinaryToWriter(message: DidDeleteUserMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DidDeleteUserMessage;
  static deserializeBinaryFromReader(message: DidDeleteUserMessage, reader: jspb.BinaryReader): DidDeleteUserMessage;
}

export namespace DidDeleteUserMessage {
  export type AsObject = {
    username: string,
    success: boolean,
    error?: string,
  }

  export enum ErrorCase { 
    _ERROR_NOT_SET = 0,
    ERROR = 3,
  }
}

export class GetChatMessagesParams extends jspb.Message {
  getSenderId(): string;
  setSenderId(value: string): GetChatMessagesParams;

  getRecipientId(): string;
  setRecipientId(value: string): GetChatMessagesParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetChatMessagesParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetChatMessagesParams): GetChatMessagesParams.AsObject;
  static serializeBinaryToWriter(message: GetChatMessagesParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetChatMessagesParams;
  static deserializeBinaryFromReader(message: GetChatMessagesParams, reader: jspb.BinaryReader): GetChatMessagesParams;
}

export namespace GetChatMessagesParams {
  export type AsObject = {
    senderId: string,
    recipientId: string,
  }
}

export class SendChatMessageParams extends jspb.Message {
  getMessage(): ChatMessage | undefined;
  setMessage(value?: ChatMessage): SendChatMessageParams;
  hasMessage(): boolean;
  clearMessage(): SendChatMessageParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendChatMessageParams.AsObject;
  static toObject(includeInstance: boolean, msg: SendChatMessageParams): SendChatMessageParams.AsObject;
  static serializeBinaryToWriter(message: SendChatMessageParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendChatMessageParams;
  static deserializeBinaryFromReader(message: SendChatMessageParams, reader: jspb.BinaryReader): SendChatMessageParams;
}

export namespace SendChatMessageParams {
  export type AsObject = {
    message?: ChatMessage.AsObject,
  }
}

export class ChatMessage extends jspb.Message {
  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): ChatMessage;
  hasTimestamp(): boolean;
  clearTimestamp(): ChatMessage;

  getSender(): string;
  setSender(value: string): ChatMessage;

  getRecipient(): string;
  setRecipient(value: string): ChatMessage;

  getText(): string;
  setText(value: string): ChatMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessage): ChatMessage.AsObject;
  static serializeBinaryToWriter(message: ChatMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessage;
  static deserializeBinaryFromReader(message: ChatMessage, reader: jspb.BinaryReader): ChatMessage;
}

export namespace ChatMessage {
  export type AsObject = {
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    sender: string,
    recipient: string,
    text: string,
  }
}

export class ChatMessageList extends jspb.Message {
  getMessagesList(): Array<ChatMessage>;
  setMessagesList(value: Array<ChatMessage>): ChatMessageList;
  clearMessagesList(): ChatMessageList;
  addMessages(value?: ChatMessage, index?: number): ChatMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessageList.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessageList): ChatMessageList.AsObject;
  static serializeBinaryToWriter(message: ChatMessageList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessageList;
  static deserializeBinaryFromReader(message: ChatMessageList, reader: jspb.BinaryReader): ChatMessageList;
}

export namespace ChatMessageList {
  export type AsObject = {
    messagesList: Array<ChatMessage.AsObject>,
  }
}

export class DropTableParams extends jspb.Message {
  getTablename(): string;
  setTablename(value: string): DropTableParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DropTableParams.AsObject;
  static toObject(includeInstance: boolean, msg: DropTableParams): DropTableParams.AsObject;
  static serializeBinaryToWriter(message: DropTableParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DropTableParams;
  static deserializeBinaryFromReader(message: DropTableParams, reader: jspb.BinaryReader): DropTableParams;
}

export namespace DropTableParams {
  export type AsObject = {
    tablename: string,
  }
}

export class DropTableMessage extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): DropTableMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DropTableMessage.AsObject;
  static toObject(includeInstance: boolean, msg: DropTableMessage): DropTableMessage.AsObject;
  static serializeBinaryToWriter(message: DropTableMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DropTableMessage;
  static deserializeBinaryFromReader(message: DropTableMessage, reader: jspb.BinaryReader): DropTableMessage;
}

export namespace DropTableMessage {
  export type AsObject = {
    success: boolean,
  }
}

export class GetServerParams extends jspb.Message {
  getPassword(): string;
  setPassword(value: string): GetServerParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetServerParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetServerParams): GetServerParams.AsObject;
  static serializeBinaryToWriter(message: GetServerParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetServerParams;
  static deserializeBinaryFromReader(message: GetServerParams, reader: jspb.BinaryReader): GetServerParams;
}

export namespace GetServerParams {
  export type AsObject = {
    password: string,
  }
}

export class Server extends jspb.Message {
  getPort(): string;
  setPort(value: string): Server;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Server.AsObject;
  static toObject(includeInstance: boolean, msg: Server): Server.AsObject;
  static serializeBinaryToWriter(message: Server, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Server;
  static deserializeBinaryFromReader(message: Server, reader: jspb.BinaryReader): Server;
}

export namespace Server {
  export type AsObject = {
    port: string,
  }
}

