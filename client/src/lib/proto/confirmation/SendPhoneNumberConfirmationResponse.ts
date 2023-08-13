// Original file: ../proto/confirmation.proto

import type { Long } from '@grpc/proto-loader';

export interface SendPhoneNumberConfirmationResponse {
  'retryAt'?: (number | string | Long);
}

export interface SendPhoneNumberConfirmationResponse__Output {
  'retryAt': (string);
}
