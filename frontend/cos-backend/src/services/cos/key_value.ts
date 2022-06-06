// @ts-ignore
/* eslint-disable */

import { doPaginationRequest, doRequest } from '@/services/cos/base';

export class KeyValue {
  static async save(body?: API.KeyValue.SaveReq, options?: { [key: string]: any }) {
    return doRequest<API.KeyValue.SaveResp>('/cos/save-value', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async list(body?: API.KeyValue.ListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.KeyValue.ListItem>('/cos/list-value', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async listAll(body?: API.KeyValue.ListAllReq, options?: { [key: string]: any }) {
    return doRequest<API.KeyValue.ListAllResp>('/cos/list-all-value', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async deleteOne(body?: API.KeyValue.SaveResp, options?: { [key: string]: any }) {
    return doRequest<API.KeyValue.DeleteResp>('/cos/delete-value', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatus(body?: API.KeyValue.UpdateStatusReq, options?: { [key: string]: any }) {
    return doRequest<API.KeyValue.UpdateStatusResp>('/cos/update-status-value', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
