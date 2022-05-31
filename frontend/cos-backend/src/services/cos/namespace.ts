// @ts-ignore
/* eslint-disable */

import { doPaginationRequest, doRequest } from '@/services/cos/base';

export class Namespace {
  static async save(body?: API.Namespace.NamespaceSaveReq, options?: { [key: string]: any }) {
    return doRequest<API.Namespace.NamespaceSaveResp>('/cos/save-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async list(body?: API.Namespace.NamespaceListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.Namespace.NamespaceListItem>('/cos/list-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }


  static async deleteNamespace(body?: API.Namespace.NamespaceSaveResp, options?: { [key: string]: any }) {
    return doRequest<API.Namespace.NamespaceDeleteResp>('/cos/delete-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatus(body?: API.Namespace.NamespaceUpdateStatusReq, options?: { [key: string]: any }) {
    return doRequest<API.Namespace.NamespaceUpdateStatusResp>('/cos/update-status-namespace', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

}
