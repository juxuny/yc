// @ts-ignore
/* eslint-disable */

import { doPaginationRequest, doRequest } from '@/services/cos/base';

export class Namespace {
  static async save(body?: API.Namespace.SaveReq, options?: { [key: string]: any }) {
    return doRequest<API.Namespace.SaveResp>('/cos/save-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async list(body?: API.Namespace.ListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.Namespace.ListItem>('/cos/list-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async deleteNamespace(body?: API.Namespace.DeleteReq, options?: { [key: string]: any }) {
    return doRequest<API.Namespace.DeleteResp>('/cos/delete-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatus(
    body?: API.Namespace.UpdateStatusReq,
    options?: { [key: string]: any },
  ) {
    return doRequest<API.Namespace.UpdateStatusResp>('/cos/update-status-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async selector(body?: API.SelectorReq, options?: { [key: string]: any }) {
    return doRequest<API.SelectorResp>('/cos/selector-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
