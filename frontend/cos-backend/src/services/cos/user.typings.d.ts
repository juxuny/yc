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
}
