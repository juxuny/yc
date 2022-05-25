// @ts-ignore
/* eslint-disable */

import {doRequest} from "@/services/cos/base";

export class User {
  static async getInfo(options?: { [key: string]: any }) {
    return doRequest<API.UserGetInfoResp>('/user/get/info', {
      method: 'GET',
      data: {},
      ...(options || {}),
    });
  }
}
