// @ts-ignore
/* eslint-disable */

declare namespace API {
  type QueryParams<Type> = Type & {
    current?: number | undefined;
    pageSize?: number | undefined;
    keyword?: string | undefined;
  };

  type BaseResp<Type> = {
    code?: number;
    data?: Type;
    msg?: string;
  };

  type Pagination<Type> = {
    list?: Type[];
    pagination: PaginationState;
  };

  const createPagination = (params: {
    current?: number | undefined;
    pageSize?: number | undefined;
  }) => {
    return {
      pageNum: params.current,
      pageSize: params.pageSize,
    };
  };

  type PaginationState = {
    pageSize?: number;
    pageNum?: number;
    total?: number;
  };

  type PaginationResp<Type> = Pagination<Type>;

  type PaginationReq = {
    pageNum?: number;
    pageSize?: number;
  };

  type EnableStatus = -1 | 0 | 1;

  type AccountType = 0 | 1;

  type SelectorItem = {
    label: string;
    value: string | number | undefined;
  };

  type ValueEnum = { [key: string]: {} };

  type SelectorReq = {
    isDisabled?: boolean;
  };

  type SelectorResp = {
    list: SelectorItem[];
  };
}
