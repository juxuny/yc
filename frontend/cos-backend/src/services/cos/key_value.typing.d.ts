// @ts-ignore
/* eslint-disable */

declare namespace API {
  declare namespace KeyValue {

    type ValueType = 1 | 2 | 3;

    type SaveReq = {
      id?: string;
      configId: string;
      configKey: string;
      configValue: string;
      isHot: boolean;
      valueType: ValueType;
    };

    type SaveResp = {};

    type ListReq = {
      pagination: PaginationReq;
      configId: string;
      searchKey?: string;
      isHot?: boolean;
      isDisabled?: boolean;
    };

    type ListItem = {
      id: string;
      configId: string;
      isDisabled: boolean;
      createTime: number;
      updateTime: number;
      configKey: string;
      configValue: string;
      isHot: boolean;
      valueType: ValueType;
    };

    type ListAllReq = {
      configId: string | number;
      isDisabled?: boolean;
      isHot?: boolean;
    }

    type ListAllResp = {
      list: ListItem[];
    }

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
