// Original file: ../proto/token.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { GenerateTokenRequest as _token_GenerateTokenRequest, GenerateTokenRequest__Output as _token_GenerateTokenRequest__Output } from '../token/GenerateTokenRequest';
import type { GenerateTokenResponse as _token_GenerateTokenResponse, GenerateTokenResponse__Output as _token_GenerateTokenResponse__Output } from '../token/GenerateTokenResponse';

export interface TokenServiceClient extends grpc.Client {
  GenerateToken(argument: _token_GenerateTokenRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  GenerateToken(argument: _token_GenerateTokenRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  GenerateToken(argument: _token_GenerateTokenRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  GenerateToken(argument: _token_GenerateTokenRequest, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  generateToken(argument: _token_GenerateTokenRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  generateToken(argument: _token_GenerateTokenRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  generateToken(argument: _token_GenerateTokenRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  generateToken(argument: _token_GenerateTokenRequest, callback: grpc.requestCallback<_token_GenerateTokenResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface TokenServiceHandlers extends grpc.UntypedServiceImplementation {
  GenerateToken: grpc.handleUnaryCall<_token_GenerateTokenRequest__Output, _token_GenerateTokenResponse>;
  
}

export interface TokenServiceDefinition extends grpc.ServiceDefinition {
  GenerateToken: MethodDefinition<_token_GenerateTokenRequest, _token_GenerateTokenResponse, _token_GenerateTokenRequest__Output, _token_GenerateTokenResponse__Output>
}
