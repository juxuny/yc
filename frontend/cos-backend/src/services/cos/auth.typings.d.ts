// @ts-ignore
/* eslint-disable */

declare namespace API {
  type LoginReq = {
    accountType: API.AccountType;
    userName?: string;
    password?: string;
  };

  type LoginResp = {
    userId?: string;
    token?: string;
    name?: string;
  };
}
