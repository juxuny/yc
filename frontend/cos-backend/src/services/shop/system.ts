// @ts-ignore
/* eslint-disable */
import {doPaginationRequest, doRequest} from "@/services/cos/base";

export class System {
  static async getRoles(params?: API.SystemGetRolesReq, pagination?: API.PaginationReq, options?: {[key: string]: any}) {
    return doPaginationRequest<API.SystemGetRoleItem>('/system/get/roles', {
      method: 'POST',
      data: {
        ...(params || {}),
        ...(pagination || {}),
      },
      ...(options || {}),
    });
  }

  static async updateRoleStatus(params?: API.SystemUpdateRoleStatusReq, options?: {[key: string]: any}) {
    return doRequest<any>('/system/update/role/status', {
      method: 'POST',
      data: {
        ...(params || {}),
      },
      ...(options || {}),
    });
  }

  static async saveRole(params?: API.SystemSaveRoleReq, options?: {[key: string]: any}) {
    return doRequest<API.SystemGetRoleItem>('/system/save/role', {
      method: 'POST',
      data: {
        ...(params || {}),
      },
      ...(options || {}),
    });
  }

  // 获取角色类型列表
  static async getRoleTypeSelector(options?: {[key: string]: any}) {
    return doRequest<API.SystemGetRoleTypeSelector>('/system/get/role/type/selector', {
      method: 'POST',
      data: {},
      ...(options || {}),
    });
  }
}


