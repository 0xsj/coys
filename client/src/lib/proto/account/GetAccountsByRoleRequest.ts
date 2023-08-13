// Original file: ../proto/account.proto

import type { Role as _account_Role, Role__Output as _account_Role__Output } from '../account/Role';
import type { Long } from '@grpc/proto-loader';

export interface GetAccountsByRoleRequest {
  'role'?: (_account_Role);
  'skip'?: (number | string | Long);
  'limit'?: (number | string | Long);
}

export interface GetAccountsByRoleRequest__Output {
  'role': (_account_Role__Output);
  'skip': (string);
  'limit': (string);
}
