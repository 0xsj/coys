// Original file: ../proto/authentication.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { SignInByPhoneNumberRequest as _authentication_SignInByPhoneNumberRequest, SignInByPhoneNumberRequest__Output as _authentication_SignInByPhoneNumberRequest__Output } from '../authentication/SignInByPhoneNumberRequest';
import type { SignInByPhoneNumberResponse as _authentication_SignInByPhoneNumberResponse, SignInByPhoneNumberResponse__Output as _authentication_SignInByPhoneNumberResponse__Output } from '../authentication/SignInByPhoneNumberResponse';
import type { SignOutRequest as _authentication_SignOutRequest, SignOutRequest__Output as _authentication_SignOutRequest__Output } from '../authentication/SignOutRequest';
import type { SignOutResponse as _authentication_SignOutResponse, SignOutResponse__Output as _authentication_SignOutResponse__Output } from '../authentication/SignOutResponse';

export interface AuthenticationServiceClient extends grpc.Client {
  SignInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  SignInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  SignInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  SignInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  signInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  signInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  signInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  signInByPhoneNumber(argument: _authentication_SignInByPhoneNumberRequest, callback: grpc.requestCallback<_authentication_SignInByPhoneNumberResponse__Output>): grpc.ClientUnaryCall;
  
  SignOut(argument: _authentication_SignOutRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  SignOut(argument: _authentication_SignOutRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  SignOut(argument: _authentication_SignOutRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  SignOut(argument: _authentication_SignOutRequest, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  signOut(argument: _authentication_SignOutRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  signOut(argument: _authentication_SignOutRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  signOut(argument: _authentication_SignOutRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  signOut(argument: _authentication_SignOutRequest, callback: grpc.requestCallback<_authentication_SignOutResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface AuthenticationServiceHandlers extends grpc.UntypedServiceImplementation {
  SignInByPhoneNumber: grpc.handleUnaryCall<_authentication_SignInByPhoneNumberRequest__Output, _authentication_SignInByPhoneNumberResponse>;
  
  SignOut: grpc.handleUnaryCall<_authentication_SignOutRequest__Output, _authentication_SignOutResponse>;
  
}

export interface AuthenticationServiceDefinition extends grpc.ServiceDefinition {
  SignInByPhoneNumber: MethodDefinition<_authentication_SignInByPhoneNumberRequest, _authentication_SignInByPhoneNumberResponse, _authentication_SignInByPhoneNumberRequest__Output, _authentication_SignInByPhoneNumberResponse__Output>
  SignOut: MethodDefinition<_authentication_SignOutRequest, _authentication_SignOutResponse, _authentication_SignOutRequest__Output, _authentication_SignOutResponse__Output>
}
