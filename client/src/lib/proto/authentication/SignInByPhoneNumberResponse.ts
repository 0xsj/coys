// Original file: ../proto/authentication.proto

import type { Long } from '@grpc/proto-loader';

export interface SignInByPhoneNumberResponse {
  'retryAt'?: (number | string | Long);
}

export interface SignInByPhoneNumberResponse__Output {
  'retryAt': (string);
}
