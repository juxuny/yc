// @ts-ignore
/* eslint-disable */

declare namespace API {
  declare namespace Config {
    type SaveReq = {
      id?: string;
      namespaceId: string;
      configId: string;
      baseId?: string;
    };

    type SaveResp = {};

    type ListReq = {
      pagination: PaginationReq;
      searchKey?: string;
      isDisabled?: boolean;
    };

    type ListItem = {
      id: string;
      namespaceId: string;
      configId: string;
      isDisabled: boolean;
      createTime: number;
      updateTime: number;
    };

    type DeleteReq = {
      id: string;
    };

    type DeleteResp = {};

    type UpdateStatusReq = {
      id: string;
      isDisabled: boolean;
    };

    type UpdateStatusResp = {};
  }
}
