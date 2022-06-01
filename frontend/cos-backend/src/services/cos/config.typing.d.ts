// @ts-ignore
/* eslint-disable */

declare namespace API {
  declare namespace Config {
    type SaveReq = {
      id?: string;
      namespace: string;
    }

    type SaveResp = {}

    type ListReq = {
      pagination: PaginationReq;
      searchKey?: string;
      isDisabled?: boolean;
    }

    type ListItem = {
      id: string;
      namespace: string;
      isDisabled: boolean;
      createTime: number;
      updateTime: number;
    }

    type DeleteReq = {
      id: string;
    }

    type DeleteResp = {}

    type UpdateStatusReq = {
      id: string;
      isDisabled: boolean;
    }

    type UpdateStatusResp = {}

  }
}
