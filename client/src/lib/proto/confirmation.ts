import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ConfirmationServiceClient as _confirmation_ConfirmationServiceClient, ConfirmationServiceDefinition as _confirmation_ConfirmationServiceDefinition } from './confirmation/ConfirmationService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  confirmation: {
    ConfirmationService: SubtypeConstructor<typeof grpc.Client, _confirmation_ConfirmationServiceClient> & { service: _confirmation_ConfirmationServiceDefinition }
    SendPhoneNumberConfirmationRequest: MessageTypeDefinition
    SendPhoneNumberConfirmationResponse: MessageTypeDefinition
    VerifyPhoneNumberConfirmationRequest: MessageTypeDefinition
    VerifyPhoneNumberConfirmationResponse: MessageTypeDefinition
  }
}

