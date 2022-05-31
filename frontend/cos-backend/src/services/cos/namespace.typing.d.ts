// @ts-ignore
/* eslint-disable */

declare namespace API {
  declare namespace Namespace {
    type NamespaceSaveReq = {
      id?: string;
      namespace: string;
    }

    type NamespaceSaveResp = {}

    type NamespaceListReq = {
      pagination: PaginationReq;
      searchKey?: string;
      isDisabled?: boolean;
    }

    type NamespaceListItem = {
      id: string;
      namespace: string;
      isDisabled: boolean;
    }

    type NamespaceDeleteReq = {
      id: string;
    }

    type NamespaceDeleteResp = {}

    type NamespaceUpdateStatusReq = {
      id: string;
      isDisabled: boolean;
    }

    type NamespaceUpdateStatusResp = {}

  }
}
