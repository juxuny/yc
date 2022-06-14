// @ts-ignore
/* eslint-disable */

import request from 'umi-request';
import { StorageKey, LocalStorage } from '@/storage';
import { message } from 'antd';
import { UUID } from 'uuid-generator-ts';

const apiPrefix = '/api';

const uuid = new UUID();

const genReqId = () => {
  return uuid.getDashFreeUUID();
};

request.use(async (ctx, next) => {
  await next();
  if (ctx.res.code !== 0) {
    console.log(ctx.res.code);
    message.error(ctx.res.msg);
    return;
  }
});

export function doRequest<Type>(path: string, options?: { [key: string]: any }) {
  let token = LocalStorage.getItem(StorageKey.TOKEN);
  return request<API.BaseResp<Type>>(apiPrefix + path, {
    headers: {
      'X-Rpc-Token': token || '',
      'Client-Request-Id': genReqId(),
    },
    errorHandler: (error) => {
      if (error.response.status === 504) {
        message.error('Gateway Timeout');
      } else {
        console.log(error.data?.msg);
        message.error(error.data?.msg);
      }
    },
    ...(options || {}),
  });
}

export function doPaginationRequest<Type>(path: string, options?: { [key: string]: any }) {
  return doRequest<API.PaginationResp<Type>>(path, options);
}
