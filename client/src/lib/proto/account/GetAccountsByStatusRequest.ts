// Original file: ../proto/account.proto

import type { Status as _account_Status, Status__Output as _account_Status__Output } from '../account/Status';
import type { Long } from '@grpc/proto-loader';

export interface GetAccountsByStatusRequest {
  'status'?: (_account_Status);
  'skip'?: (number | string | Long);
  'limit'?: (number | string | Long);
}

export interface GetAccountsByStatusRequest__Output {
  'status': (_account_Status__Output);
  'skip': (string);
  'limit': (string);
}
