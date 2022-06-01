// @ts-ignore
/* eslint-disable */

import { doPaginationRequest, doRequest } from '@/services/cos/base';

export class Config {
  static async save(body?: API.Config.SaveReq, options?: { [key: string]: any }) {
    return doRequest<API.Config.SaveResp>('/cos/save-config', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async list(body?: API.Config.ListReq, options?: { [key: string]: any }) {
    return doPaginationRequest<API.Config.ListItem>('/cos/list-config', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async deleteOne(body?: API.Config.SaveResp, options?: { [key: string]: any }) {
    return doRequest<API.Config.DeleteResp>('/cos/delete-config', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }

  static async updateStatus(body?: API.Config.UpdateStatusReq, options?: { [key: string]: any }) {
    return doRequest<API.Config.UpdateStatusResp>('/cos/update-status-config', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
