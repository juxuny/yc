// @ts-ignore
/* eslint-disable */

import {doRequest} from "@/services/cos/base";

export class User {
  static async userInfo(body?: API.UserInfoReq, options?: { [key: string]: any }) {
    return doRequest<API.LoginResp>('/cos/user-info', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
}
