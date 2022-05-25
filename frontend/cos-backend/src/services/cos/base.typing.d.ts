// @ts-ignore
/* eslint-disable */

declare namespace API {
  type BaseResp<Type> = {
    code?: number;
    data?: Type;
    msg?: string;
  };

  type Pagination<Type> = {
    list?: Type[];
    pageSize?: number;
    pageNum?: number;
    total?: number;
  }

  type PaginationState = {
    pageSize?: number;
    pageNum?: number;
    total?: number;
  }

  type PaginationResp<Type> = Pagination<Type>

  type PaginationReq = {
    pageNum?: number;
    pageSize?: number;
  }

  type EnableStatus = -1|0|1;

  type AccountType = 0|1;

  type SelectorItem = {
    label: string;
    value: string|number;
  }

  type ValueEnum = {[key: string]: {}}
}
