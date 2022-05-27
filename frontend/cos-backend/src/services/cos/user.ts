// @ts-ignore
/* eslint-disable */

import { doPaginationRequest, doRequest } from '@/services/cos/base';

export class User {
  static async userInfo(body?: API.UserInfoReq, options?: { [key: string]: any }) {
    return doRequest<API.UserInfoResp>('/cos/user-info', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async userList(body?: API.UserListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.UserListItem>('/cos/user-list', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
