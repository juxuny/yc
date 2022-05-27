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
    userId: string;
    nick: string;
    identifier: string;
    accountType: AccountType;
    createTime: number;
    updateTime: number;
  };

  type UserListReq = {
    pagination: PaginationReq;
    searchKey?: string;
    isDisabled?: boolean;
  };

  type SaveUserInfoReq = {
    userId: string;
    nick: string;
    identifier: string;
  };
}
