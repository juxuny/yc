// @ts-ignore
/* eslint-disable */

import { doPaginationRequest, doRequest } from '@/services/cos/base';

export class User {
  static async userInfo(body?: API.User.InfoReq, options?: { [key: string]: any }) {
    return doRequest<API.User.InfoResp>('/cos/user-info', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async userList(body?: API.User.ListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.User.ListItem>('/cos/user-list', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async saveOrCreateUser(body?: API.User.SaveReq, options?: { [key: string]: any }) {
    return doRequest<API.User.SaveResp>('/cos/save-or-create-user', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatus(body?: API.User.UpdateStatusReq, options?: { [key: string]: any }) {
    return doRequest<API.User.UpdateStatusResp>('/cos/user-update-status', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async delete(body?: API.User.DeleteReq, options?: { [key: string]: any }) {
    return doRequest<API.User.DeleteResp>('/cos/user-delete', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async modifyPassword(body?: API.User.ModifyPasswordReq, options?: { [key: string]: any }) {
    return doRequest<API.User.ModifyPasswordResp>('/cos/modify-password', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async accessKeyList(body?: API.User.AccessKeyListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.User.AccessKeyListItem>('/cos/access-key-list', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async createAccessKey(body?: API.User.CreateAccessKeyReq, options?: { [key: string]: any }) {
    return doRequest<API.User.CreateAccessKeyResp>('/cos/create-access-key', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatusAccessKey(body?: API.User.UpdateStatusAccessKeyReq, options?: { [key: string]: any }) {
    return doRequest<API.User.UpdateStatusAccessKeyResp>('/cos/update-status-access-key', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async deleteAccessKey(body?: API.User.DeleteAccessKeyReq, options?: { [key: string]: any }) {
    return doRequest<API.User.DeleteAccessKeyResp>('/cos/delete-access-key', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async setRemarkAccessKey(body?: API.User.SetRemarkAccessKeyReq, options?: { [key: string]: any }) {
    return doRequest<API.User.SetRemarkAccessKeyResp>('/cos/set-remark-access-key', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
