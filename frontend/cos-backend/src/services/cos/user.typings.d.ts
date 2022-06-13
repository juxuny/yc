// @ts-ignore
/* eslint-disable */

declare namespace API {
  declare namespace User {
    type InfoReq = {
      userId: string;
    }

    type InfoResp = {
      userId?: string;
      nick?: string;
      identifier?: string;
      accountType: AccountType;
    }

    type ListItem = {
      id: string;
      nick: string;
      identifier: string;
      accountType: AccountType;
      createTime: number;
      updateTime: number;
      isDisabled: boolean;
    }

    type ListReq = {
      pagination: PaginationReq;
      searchKey?: string;
      isDisabled?: boolean;
    }

    type SaveReq = {
      userId?: string;
      nick: string;
      identifier: string;
      credential?: string;
      accountType?: AccountType;
    }

    type SaveResp = {
      id?: string;
    }

    type UpdateStatusReq = {
      userId: string;
      isDisabled: boolean;
    }

    type UpdateStatusResp = {}

    type DeleteReq = {
      userId: string;
    }

    type DeleteResp = {}

    type ModifyPasswordReq = {
      userId?: string;
      oldPassword: string;
      newPassword: string;
    }

    type ModifyPasswordResp = {}

    type AccessKeyListReq = {
      searchKey?: string;
      pagination: API.PaginationReq;
      isDisabled?: boolean;
    }

    type AccessKeyListItem = {
      id: number | string;
      createTime: number;
      updateTime: number;
      isDisabled: boolean;
      userId: string | number;
      accessKey: string;
      hasValidity: boolean;
      validStartTime: number;
      validEndTime: number;
      remark: string;
    }

    type CreateAccessKeyReq = {
      remark: string;
      hasValidity: boolean;
      validStartTime: number;
      validEndTime: number;
    }

    type CreateAccessKeyResp = {
      remark: string;
      accessKey: string;
      secret: string;
    }

    type UpdateStatusAccessKeyReq = {
      id: number | string;
      isDisabled: boolean;
    }

    type UpdateStatusAccessKeyResp = {}

    type DeleteAccessKeyReq = {
      id: number | string;
    }

    type DeleteAccessKeyResp = {}

    type SetRemarkAccessKeyReq = {
      id: number | string;
      remark: string;
    }

    type SetRemarkAccessKeyResp = {}
  }
}
