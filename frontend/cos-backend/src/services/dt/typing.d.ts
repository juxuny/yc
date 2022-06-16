type Pagination = {
  pageNum: number;
  pageSize: number;
}

type PaginationResp = {
  pageNum: number;
  pageSize: number;
  total: number;
}

type Error = {
  code: number;
  msg: string;
  data: Map<string, string>;
}

type BaseResp<Type> = {
  code: number;
  data?: Type;
  msg?: string;
}
