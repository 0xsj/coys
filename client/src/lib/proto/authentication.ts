import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { AuthenticationServiceClient as _authentication_AuthenticationServiceClient, AuthenticationServiceDefinition as _authentication_AuthenticationServiceDefinition } from './authentication/AuthenticationService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  authentication: {
    AuthenticationService: SubtypeConstructor<typeof grpc.Client, _authentication_AuthenticationServiceClient> & { service: _authentication_AuthenticationServiceDefinition }
    ConfirmPhoneNumberRequest: MessageTypeDefinition
    ConfirmPhoneNumberResponse: MessageTypeDefinition
    RefreshTokenRequest: MessageTypeDefinition
    RefreshTokenResponse: MessageTypeDefinition
    SignInByPhoneNumberRequest: MessageTypeDefinition
    SignInByPhoneNumberResponse: MessageTypeDefinition
    SignOutRequest: MessageTypeDefinition
    SignOutResponse: MessageTypeDefinition
    VerifyAccessRequest: MessageTypeDefinition
    VerifyAccessResponse: MessageTypeDefinition
  }
}

