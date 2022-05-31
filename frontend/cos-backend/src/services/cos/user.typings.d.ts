// @ts-ignore
/* eslint-disable */

declare namespace API {
  type UserInfoReq = {
    userId: string;
  };

  type UserInfoResp = {
    userId?: string;
    nick?: string;
    identifier?: string;
    accountType: AccountType;
  };

  type UserListItem = {
    id: string;
    nick: string;
    identifier: string;
    accountType: AccountType;
    createTime: number;
    updateTime: number;
    isDisabled: boolean;
  };

  type UserListReq = {
    pagination: PaginationReq;
    searchKey?: string;
    isDisabled?: boolean;
  };

  type SaveUserInfoReq = {
    userId?: string;
    nick: string;
    identifier: string;
    credential?: string;
    accountType?: AccountType;
  };

  type SaveUserInfoResp = {
    id?: string;
  };

  type UserUpdateStatusReq = {
    userId: string;
    isDisabled: boolean;
  };

  type UserUpdateStatusResp = {};

  type UserDeleteReq = {
    userId: string;
  };

  type UserDeleteResp = {};

  type UserModifyPasswordReq = {
    userId?: string;
    oldPassword: string;
    newPassword: string;
  };

  type UserModifyPasswordResp = {};
}
