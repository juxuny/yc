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

  static async saveOrCreateUser(body?: API.SaveUserInfoReq, options?: { [key: string]: any }) {
    return doRequest<API.SaveUserInfoResp>('/cos/save-or-create-user', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatus(body?: API.UserUpdateStatusReq, options?: { [key: string]: any }) {
    return doRequest<API.UserUpdateStatusResp>('/cos/user-update-status', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async delete(body?: API.UserDeleteReq, options?: { [key: string]: any }) {
    return doRequest<API.UserDeleteResp>('/cos/user-delete', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async modifyPassword(body?: API.UserModifyPasswordReq, options?: { [key: string]: any }) {
    return doRequest<API.UserModifyPasswordResp>('/cos/modify-password', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
