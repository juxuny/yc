// @ts-ignore
/* eslint-disable */

declare namespace API {
  declare namespace User {
    type InfoReq = {
      userId: string;
    };

    type InfoResp = {
      userId?: string;
      nick?: string;
      identifier?: string;
      accountType: AccountType;
    };

    type ListItem = {
      id: string;
      nick: string;
      identifier: string;
      accountType: AccountType;
      createTime: number;
      updateTime: number;
      isDisabled: boolean;
    };

    type ListReq = {
      pagination: PaginationReq;
      searchKey?: string;
      isDisabled?: boolean;
    };

    type SaveReq = {
      userId?: string;
      nick: string;
      identifier: string;
      credential?: string;
      accountType?: AccountType;
    };

    type SaveResp = {
      id?: string;
    };

    type UpdateStatusReq = {
      userId: string;
      isDisabled: boolean;
    };

    type UpdateStatusResp = {};

    type DeleteReq = {
      userId: string;
    };

    type DeleteResp = {};

    type ModifyPasswordReq = {
      userId?: string;
      oldPassword: string;
      newPassword: string;
    };

    type ModifyPasswordResp = {};
  }
}
