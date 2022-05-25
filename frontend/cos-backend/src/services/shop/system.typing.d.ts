// @ts-ignore
/* eslint-disable */


declare namespace API {

  type SystemGetRolesReq = {
    status: API.EnableStatus;
    type: number;
  }

  type SystemGetRoleItem = {
    id?: number;
    type?: number;
    typeName?: string;
    name?: string;
    status?: API.EnableStatus;
    createTime?: string;
    updateTime?: string;
    remark?: string;
  }

  type SystemUpdateRoleStatusReq = {
    id?: number;
    status?: API.EnableStatus;
  }

  type SystemSaveRoleReq = {
    id?: number;
    type?: number;
    name?: string;
    remark?: string;
  }

  type SystemGetRoleTypeSelector = {
    list?: SelectorItem[];
  }
}
