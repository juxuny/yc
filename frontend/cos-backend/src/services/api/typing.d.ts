import type * as dt from '@juxuny/yc-ts-data-type/typing';

export type ModifyPasswordType = 0 | 1 | 2;

export class ModifyPasswordTypeEnum {

  static Unknown: ModifyPasswordType = 0;

  static OldPassword: ModifyPasswordType = 1;

  static Token: ModifyPasswordType = 2;

}

export type AccountType = 0 | 1;

// 账号类型
export class AccountTypeEnum {
  // 未知
  static Unknown: AccountType = 0;
  // 密码登录的账号
  static Password: AccountType = 1;

}

export type ConfigRecordType = 0 | 1 | 2 | 3 | 4 | 5;

export class ConfigRecordTypeEnum {

  static Unknown: ConfigRecordType = 0;

  static Create: ConfigRecordType = 1;

  static Modify: ConfigRecordType = 2;

  static Delete: ConfigRecordType = 3;

  static Link: ConfigRecordType = 4;

  static Unlink: ConfigRecordType = 5;

}

export type ValueType = 0 | 1 | 2 | 3 | 4;

export class ValueTypeEnum {

  static Unknown: ValueType = 0;

  static Int: ValueType = 1;

  static Float64: ValueType = 2;

  static Bool: ValueType = 3;

  static String: ValueType = 4;

}


export type HealthRequest = {

}

export type HealthResponse = {

  currentTime?: string;

}

// 登录请求
export type LoginRequest = {
  // 用户名
  userName?: string;
  // 登录密码
  password?: string;

  accountType?: AccountType;

}

export type LoginResponse = {

  userId?: number;

  name?: string;

  token?: string;

}

export type UserInfoRequest = {

  userId?: string;

}

export type UserInfoSimple = {

  userId?: string;

  name?: string;

  userName?: string;

  accountType?: AccountType;

}

export type UserInfoResponse = {

  userId?: string;

  nick?: string;

  identifier?: string;

  accountType?: AccountType;

}

export type UpdateInfoRequest = {

  userId?: string;

  nick?: string;

}

export type UpdateInfoResponse = {

}

export type ModifyPasswordRequest = {

  userId?: string;

  oldPassword?: string;

  newPassword?: string;

}

export type ModifyPasswordResponse = {

}

export type NamespaceResp = {

  id?: string;

  namespace?: string;

  createTime?: int64;

  updateTime?: int64;

  isDisabled?: bool;

  creatorId?: string;

}

export type ConfigResp = {

  id?: string;

  createTime?: int64;

  updateTime?: int64;

  deletedAt?: int64;

  configId?: string;

  isDisabled?: bool;

  creatorId?: string;

  baseId?: string;

}

export type KeyValueResp = {

  id?: string;

  createTime?: int64;

  updateTime?: int64;

  deletedAt?: int64;

  isDisabled?: bool;

  configKey?: string;

  configValue?: string;

  valueType?: ValueType;

  configId?: string;

  creatorId?: string;

  isHot?: bool;

}

export type SaveOrCreateUserRequest = {

  identifier?: string;

  credential?: string;

  accountType?: AccountType;

  userId?: string;

  nick?: string;

}

export type SaveOrCreateUserResponse = {

  userId?: string;

  isNew?: bool;

}

export type UserListItem = {

  id?: string;

  identifier?: string;

  accountType?: AccountType;

  nick?: string;

  createTime?: int64;

  updateTime?: int64;

  isDisabled?: bool;

}

export type UserListRequest = {

  pagination?: dt.Pagination;

  searchKey?: string;

  isDisabled?: string;

}

export type UserListResponse = {

  pagination?: dt.PaginationResp;

  list?: UserListItem[];

}

export type SaveNamespaceRequest = {

  namespace?: string;

  id?: string;

}

export type SaveNamespaceResponse = {

}

export type ListNamespaceRequest = {

  pagination?: dt.Pagination;

  searchKey?: string;

  isDisabled?: string;

}

export type ListNamespaceItem = {

  id?: string;

  namespace?: string;

  createTime?: int64;

  updateTime?: int64;

  isDisabled?: bool;

}

export type ListNamespaceResponse = {

  pagination?: dt.PaginationResp;

  list?: ListNamespaceItem[];

}

export type DeleteNamespaceRequest = {

  id?: string;

}

export type DeleteNamespaceResponse = {

}

export type SaveConfigRequest = {

  id?: string;

  namespaceId?: string;

  configId?: string;

  baseId?: string;

}

export type SaveConfigResponse = {

}

export type DeleteConfigRequest = {

  id?: string;

}

export type DeleteConfigResponse = {

}

export type ListConfigRequest = {

  pagination?: dt.Pagination;

  namespaceId?: string;

  searchKey?: string;

  isDisabled?: string;

}

export type ListConfigItem = {

  id?: string;

  createTime?: int64;

  updateTime?: int64;

  baseId?: string;

  namespaceId?: string;

  configId?: string;

  linkCount?: number;

  isDisabled?: bool;

}

export type ListConfigResponse = {

  pagination?: dt.PaginationResp;

  list?: ListConfigItem[];

}

export type CloneConfigRequest = {

  id?: string;

  newConfigId?: string;

}

export type CloneConfigResponse = {

}

export type SaveValueRequest = {

  configId?: string;

  configKey?: string;

  configValue?: string;

  isHot?: bool;

  valueType?: ValueType;

}

export type SaveValueResponse = {

}

export type ListValueRequest = {

  pagination?: dt.Pagination;

  configId?: string;

  searchKey?: string;

  isDisabled?: string;

}

export type ListValueResponse = {

  pagination?: dt.PaginationResp;

  list?: KeyValueResp[];

}

export type DeleteValueRequest = {

  id?: string;

  key?: string;

}

export type DeleteValueResponse = {

}

export type DisableValueRequest = {

  id?: string;

  isDisabled?: bool;

}

export type DisableValueResponse = {

}

export type UpdateStatusValueRequest = {

  id?: string;

  isDisabled?: bool;

}

export type UpdateStatusValueResponse = {

}

export type ListAllValueRequest = {

  configId?: string;

  isDisabled?: string;

  isHot?: string;

  searchKey?: string;

}

export type ListAllValueResponse = {

  list?: KeyValueResp[];

}

export type UserUpdateStatusRequest = {

  userId?: string;

  isDisabled?: bool;

}

export type UserUpdateStatusResponse = {

}

export type UserDeleteRequest = {

  userId?: string;

}

export type UserDeleteResponse = {

}

export type UpdateStatusNamespaceRequest = {

  id?: string;

  isDisabled?: bool;

}

export type UpdateStatusNamespaceResponse = {

}

export type UpdateStatusConfigRequest = {

  id?: string;

  isDisabled?: bool;

}

export type UpdateStatusConfigResponse = {

}

export type SelectorRequest = {

  isDisabled?: string;

}

export type SelectorItem = {

  label?: string;

  value?: string;

}

export type SelectorResponse = {

  list?: SelectorItem[];

}

export type AccessKeyItem = {

  id?: string;

  createTime?: int64;

  updateTime?: int64;

  isDisabled?: bool;

  deletedAt?: int64;

  userId?: string;

  accessKey?: string;

  hasValidity?: bool;

  validStartTime?: int64;

  validEndTime?: int64;

  remark?: string;

}

export type AccessKeyListRequest = {

  searchKey?: string;

  pagination?: dt.Pagination;

  isDisabled?: string;

}

export type AccessKeyListResponse = {

  pagination?: dt.PaginationResp;

  list?: AccessKeyItem[];

}

export type CreateAccessKeyRequest = {

  remark?: string;

  hasValidity?: bool;

  validStartTime?: int64;

  validEndTime?: int64;

}

export type CreateAccessKeyResponse = {

  accessKey?: string;

  remark?: string;

  secret?: string;

}

export type UpdateStatusAccessKeyRequest = {

  id?: string;

  isDisabled?: bool;

}

export type UpdateStatusAccessKeyResponse = {

}

export type DeleteAccessKeyRequest = {

  id?: string;

}

export type DeleteAccessKeyResponse = {

}

export type SetAccessKeyRemarkRequest = {

  id?: string;

  remark?: string;

}

export type SetAccessKeyRemarkResponse = {

}

