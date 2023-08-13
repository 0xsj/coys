// Original file: ../proto/account.proto

import type { Role as _account_Role, Role__Output as _account_Role__Output } from '../account/Role';

export interface CreateAccountRequest {
  'phoneNumber'?: (string);
  'role'?: (_account_Role);
}

export interface CreateAccountRequest__Output {
  'phoneNumber': (string);
  'role': (_account_Role__Output);
}
