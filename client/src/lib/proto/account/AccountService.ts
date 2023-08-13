// Original file: ../proto/account.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreateAccountRequest as _account_CreateAccountRequest, CreateAccountRequest__Output as _account_CreateAccountRequest__Output } from '../account/CreateAccountRequest';
import type { CreateAccountResponse as _account_CreateAccountResponse, CreateAccountResponse__Output as _account_CreateAccountResponse__Output } from '../account/CreateAccountResponse';
import type { GetAccountByIdRequest as _account_GetAccountByIdRequest, GetAccountByIdRequest__Output as _account_GetAccountByIdRequest__Output } from '../account/GetAccountByIdRequest';
import type { GetAccountByIdResponse as _account_GetAccountByIdResponse, GetAccountByIdResponse__Output as _account_GetAccountByIdResponse__Output } from '../account/GetAccountByIdResponse';
import type { GetAccountByPhoneNumberRequest as _account_GetAccountByPhoneNumberRequest, GetAccountByPhoneNumberRequest__Output as _account_GetAccountByPhoneNumberRequest__Output } from '../account/GetAccountByPhoneNumberRequest';
import type { GetAccountByPhoneNumberResponse as _account_GetAccountByPhoneNumberResponse, GetAccountByPhoneNumberResponse__Output as _account_GetAccountByPhoneNumberResponse__Output } from '../account/GetAccountByPhoneNumberResponse';

export interface AccountServiceClient extends grpc.Client {
  CreateAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  CreateAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  CreateAccount(argument: _account_CreateAccountRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  CreateAccount(argument: _account_CreateAccountRequest, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  createAccount(argument: _account_CreateAccountRequest, callback: grpc.requestCallback<_account_CreateAccountResponse__Output>): grpc.ClientUnaryCall;
  
  GetAccountById(argument: _account_GetAccountByIdRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  GetAccountById(argument: _account_GetAccountByIdRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  GetAccountById(argument: _account_GetAccountByIdRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  GetAccountById(argument: _account_GetAccountByIdRequest, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  getAccountById(argument: _account_GetAccountByIdRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  getAccountById(argument: _account_GetAccountByIdRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  getAccountById(argument: _account_GetAccountByIdRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  getAccountById(argument: _account_GetAccountByIdRequest, callback: grpc.requestCallback<_account_GetAccountByIdResponse__Output>): grpc.ClientUnaryCall;
  
  GetAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  GetAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  GetAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  GetAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  getAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  getAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  getAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  getAccountByPhoneNumber(argument: _account_GetAccountByPhoneNumberRequest, callback: grpc.requestCallback<_account_GetAccountByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface AccountServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateAccount: grpc.handleUnaryCall<_account_CreateAccountRequest__Output, _account_CreateAccountResponse>;
  
  GetAccountById: grpc.handleUnaryCall<_account_GetAccountByIdRequest__Output, _account_GetAccountByIdResponse>;
  
  GetAccountByPhoneNumber: grpc.handleUnaryCall<_account_GetAccountByPhoneNumberRequest__Output, _account_GetAccountByPhoneNumberResponse>;
  
}

export interface AccountServiceDefinition extends grpc.ServiceDefinition {
  CreateAccount: MethodDefinition<_account_CreateAccountRequest, _account_CreateAccountResponse, _account_CreateAccountRequest__Output, _account_CreateAccountResponse__Output>
  GetAccountById: MethodDefinition<_account_GetAccountByIdRequest, _account_GetAccountByIdResponse, _account_GetAccountByIdRequest__Output, _account_GetAccountByIdResponse__Output>
  GetAccountByPhoneNumber: MethodDefinition<_account_GetAccountByPhoneNumberRequest, _account_GetAccountByPhoneNumberResponse, _account_GetAccountByPhoneNumberRequest__Output, _account_GetAccountByPhoneNumberResponse__Output>
}
