// Original file: ../proto/account.proto

export const Status = {
  PENDING: 'PENDING',
  ACTIVE: 'ACTIVE',
  SUSPENDED: 'SUSPENDED',
} as const;

export type Status =
  | 'PENDING'
  | 0
  | 'ACTIVE'
  | 1
  | 'SUSPENDED'
  | 2

export type Status__Output = typeof Status[keyof typeof Status]
