// @ts-ignore
/* eslint-disable */

import {doRequest} from "@/services/cos/base";

export class Auth {
  static async login(body: API.LoginReq, options?: { [key: string]: any }) {
    return doRequest<API.LoginResp>('/auth/login', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
