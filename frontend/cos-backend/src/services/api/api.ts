// Code generated by yc@v0.0.1. DO NOT EDIT.
import type * as typing from './typing';
import {doRequest} from '@juxuny/yc-ts-data-type';
import { prefix } from './index'

export class cos {
  // health check
  static health = (body: typing.HealthRequest, options?: { [key: string]: any }) => doRequest<typing.HealthResponse>(prefix + '/cos/health', { method: 'POST', data: body, ...(options || {}) });

  static login = (body: typing.LoginRequest, options?: { [key: string]: any }) => doRequest<typing.LoginResponse>(prefix + '/cos/login', { method: 'POST', data: body, ...(options || {}) });

  static userInfo = (body: typing.UserInfoRequest, options?: { [key: string]: any }) => doRequest<typing.UserInfoResponse>(prefix + '/cos/user-info', { method: 'POST', data: body, ...(options || {}) });

  static updateInfo = (body: typing.UpdateInfoRequest, options?: { [key: string]: any }) => doRequest<typing.UpdateInfoResponse>(prefix + '/cos/update-info', { method: 'POST', data: body, ...(options || {}) });

  static modifyPassword = (body: typing.ModifyPasswordRequest, options?: { [key: string]: any }) => doRequest<typing.ModifyPasswordResponse>(prefix + '/cos/modify-password', { method: 'POST', data: body, ...(options || {}) });

  static saveOrCreateUser = (body: typing.SaveOrCreateUserRequest, options?: { [key: string]: any }) => doRequest<typing.SaveOrCreateUserResponse>(prefix + '/cos/save-or-create-user', { method: 'POST', data: body, ...(options || {}) });

  static userList = (body: typing.UserListRequest, options?: { [key: string]: any }) => doRequest<typing.UserListResponse>(prefix + '/cos/user-list', { method: 'POST', data: body, ...(options || {}) });

  static userUpdateStatus = (body: typing.UserUpdateStatusRequest, options?: { [key: string]: any }) => doRequest<typing.UserUpdateStatusResponse>(prefix + '/cos/user-update-status', { method: 'POST', data: body, ...(options || {}) });

  static userDelete = (body: typing.UserDeleteRequest, options?: { [key: string]: any }) => doRequest<typing.UserDeleteResponse>(prefix + '/cos/user-delete', { method: 'POST', data: body, ...(options || {}) });

  static accessKeyList = (body: typing.AccessKeyListRequest, options?: { [key: string]: any }) => doRequest<typing.AccessKeyListResponse>(prefix + '/cos/access-key-list', { method: 'POST', data: body, ...(options || {}) });

  static createAccessKey = (body: typing.CreateAccessKeyRequest, options?: { [key: string]: any }) => doRequest<typing.CreateAccessKeyResponse>(prefix + '/cos/create-access-key', { method: 'POST', data: body, ...(options || {}) });

  static updateStatusAccessKey = (body: typing.UpdateStatusAccessKeyRequest, options?: { [key: string]: any }) => doRequest<typing.UpdateStatusAccessKeyResponse>(prefix + '/cos/update-status-access-key', { method: 'POST', data: body, ...(options || {}) });

  static deleteAccessKey = (body: typing.DeleteAccessKeyRequest, options?: { [key: string]: any }) => doRequest<typing.DeleteAccessKeyResponse>(prefix + '/cos/delete-access-key', { method: 'POST', data: body, ...(options || {}) });

  static setRemarkAccessKey = (body: typing.SetAccessKeyRemarkRequest, options?: { [key: string]: any }) => doRequest<typing.SetAccessKeyRemarkResponse>(prefix + '/cos/set-remark-access-key', { method: 'POST', data: body, ...(options || {}) });

  static saveNamespace = (body: typing.SaveNamespaceRequest, options?: { [key: string]: any }) => doRequest<typing.SaveNamespaceResponse>(prefix + '/cos/save-namespace', { method: 'POST', data: body, ...(options || {}) });

  static listNamespace = (body: typing.ListNamespaceRequest, options?: { [key: string]: any }) => doRequest<typing.ListNamespaceResponse>(prefix + '/cos/list-namespace', { method: 'POST', data: body, ...(options || {}) });

  static deleteNamespace = (body: typing.DeleteNamespaceRequest, options?: { [key: string]: any }) => doRequest<typing.DeleteNamespaceResponse>(prefix + '/cos/delete-namespace', { method: 'POST', data: body, ...(options || {}) });

  static updateStatusNamespace = (body: typing.UpdateStatusNamespaceRequest, options?: { [key: string]: any }) => doRequest<typing.UpdateStatusNamespaceResponse>(prefix + '/cos/update-status-namespace', { method: 'POST', data: body, ...(options || {}) });

  static selectorNamespace = (body: typing.SelectorRequest, options?: { [key: string]: any }) => doRequest<typing.SelectorResponse>(prefix + '/cos/selector-namespace', { method: 'POST', data: body, ...(options || {}) });

  static saveConfig = (body: typing.SaveConfigRequest, options?: { [key: string]: any }) => doRequest<typing.SaveConfigResponse>(prefix + '/cos/save-config', { method: 'POST', data: body, ...(options || {}) });

  static deleteConfig = (body: typing.DeleteConfigRequest, options?: { [key: string]: any }) => doRequest<typing.DeleteConfigResponse>(prefix + '/cos/delete-config', { method: 'POST', data: body, ...(options || {}) });

  static listConfig = (body: typing.ListConfigRequest, options?: { [key: string]: any }) => doRequest<typing.ListConfigResponse>(prefix + '/cos/list-config', { method: 'POST', data: body, ...(options || {}) });

  static cloneConfig = (body: typing.CloneConfigRequest, options?: { [key: string]: any }) => doRequest<typing.CloneConfigResponse>(prefix + '/cos/clone-config', { method: 'POST', data: body, ...(options || {}) });

  static updateStatusConfig = (body: typing.UpdateStatusConfigRequest, options?: { [key: string]: any }) => doRequest<typing.UpdateStatusConfigResponse>(prefix + '/cos/update-status-config', { method: 'POST', data: body, ...(options || {}) });

  static saveValue = (body: typing.SaveValueRequest, options?: { [key: string]: any }) => doRequest<typing.SaveValueResponse>(prefix + '/cos/save-value', { method: 'POST', data: body, ...(options || {}) });

  static deleteValue = (body: typing.DeleteValueRequest, options?: { [key: string]: any }) => doRequest<typing.DeleteValueRequest>(prefix + '/cos/delete-value', { method: 'POST', data: body, ...(options || {}) });

  static listValue = (body: typing.ListValueRequest, options?: { [key: string]: any }) => doRequest<typing.ListValueResponse>(prefix + '/cos/list-value', { method: 'POST', data: body, ...(options || {}) });

  static disableValue = (body: typing.DisableValueRequest, options?: { [key: string]: any }) => doRequest<typing.DisableValueResponse>(prefix + '/cos/disable-value', { method: 'POST', data: body, ...(options || {}) });

  static listAllValue = (body: typing.ListAllValueRequest, options?: { [key: string]: any }) => doRequest<typing.ListAllValueResponse>(prefix + '/cos/list-all-value', { method: 'POST', data: body, ...(options || {}) });

  static updateStatusValue = (body: typing.UpdateStatusValueRequest, options?: { [key: string]: any }) => doRequest<typing.UpdateStatusValueResponse>(prefix + '/cos/update-status-value', { method: 'POST', data: body, ...(options || {}) });

}
