// Original file: ../proto/account.proto

export const Role = {
  ADMIN: 'ADMIN',
  CREATOR: 'CREATOR',
  USER: 'USER',
} as const;

export type Role =
  | 'ADMIN'
  | 0
  | 'CREATOR'
  | 1
  | 'USER'
  | 2

export type Role__Output = typeof Role[keyof typeof Role]
