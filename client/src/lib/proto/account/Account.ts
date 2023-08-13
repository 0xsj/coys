// Original file: ../proto/account.proto

import type { Role as _account_Role, Role__Output as _account_Role__Output } from '../account/Role';
import type { Status as _account_Status, Status__Output as _account_Status__Output } from '../account/Status';
import type { Long } from '@grpc/proto-loader';

export interface Account {
  'id'?: (string);
  'phoneNumber'?: (string);
  'role'?: (_account_Role);
  'status'?: (_account_Status);
  'createdAt'?: (number | string | Long);
}

export interface Account__Output {
  'id': (string);
  'phoneNumber': (string);
  'role': (_account_Role__Output);
  'status': (_account_Status__Output);
  'createdAt': (string);
}
