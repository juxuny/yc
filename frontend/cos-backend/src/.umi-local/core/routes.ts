// @ts-nocheck
import React from 'react';
import { ApplyPluginsType, dynamic } from '/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/node_modules/@umijs/runtime';
import * as umiExports from './umiExports';
import { plugin } from './plugin';
import LoadingComponent from '@ant-design/pro-layout/es/PageLoading';

export function getRoutes() {
  const routes = [
  {
    "path": "/",
    "component": dynamic({ loader: () => import(/* webpackChunkName: '.umi-local__plugin-layout__Layout' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/.umi-local/plugin-layout/Layout.tsx'), loading: LoadingComponent}),
    "routes": [
      {
        "path": "/user",
        "layout": false,
        "routes": [
          {
            "path": "/user",
            "routes": [
              {
                "name": "login",
                "path": "/user/login",
                "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__user__Login' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/user/Login'), loading: LoadingComponent}),
                "exact": true
              }
            ]
          },
          {
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__404' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/404'), loading: LoadingComponent}),
            "exact": true
          }
        ]
      },
      {
        "path": "/welcome",
        "name": "welcome",
        "icon": "smile",
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__Welcome' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/Welcome'), loading: LoadingComponent}),
        "exact": true
      },
      {
        "path": "/admin",
        "name": "admin",
        "icon": "crown",
        "access": "canAdmin",
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__Admin' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/Admin'), loading: LoadingComponent}),
        "routes": [
          {
            "path": "/admin/sub-page",
            "name": "sub-page",
            "icon": "smile",
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__Welcome' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/Welcome'), loading: LoadingComponent}),
            "exact": true
          },
          {
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__404' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/404'), loading: LoadingComponent}),
            "exact": true
          }
        ]
      },
      {
        "name": "list.table-list",
        "icon": "table",
        "path": "/list",
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__TableList' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/TableList'), loading: LoadingComponent}),
        "exact": true
      },
      {
        "path": "/system",
        "name": "system",
        "icon": "setting",
        "routes": [
          {
            "path": "/system/user-management",
            "name": "user-management",
            "icon": "user",
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__system__UserManagement' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/system/UserManagement'), loading: LoadingComponent}),
            "exact": true
          },
          {
            "path": "/system/role-management",
            "name": "role-management",
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__system__RoleManagement' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/system/RoleManagement'), loading: LoadingComponent}),
            "exact": true
          },
          {
            "path": "/system/permission-management",
            "name": "permission-management",
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__system__UserManagement' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/system/UserManagement'), loading: LoadingComponent}),
            "exact": true
          },
          {
            "path": "/system/menu-management",
            "name": "menu-management",
            "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__system__UserManagement' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/system/UserManagement'), loading: LoadingComponent}),
            "exact": true
          }
        ]
      },
      {
        "path": "/index.html",
        "redirect": "/welcome",
        "exact": true
      },
      {
        "path": "/",
        "redirect": "/welcome",
        "exact": true
      },
      {
        "component": dynamic({ loader: () => import(/* webpackChunkName: 'p__404' */'/Users/juxuny/pro/yuan/cloud/yc/frontend/cos-backend/src/pages/404'), loading: LoadingComponent}),
        "exact": true
      }
    ]
  }
];

  // allow user to extend routes
  plugin.applyPlugins({
    key: 'patchRoutes',
    type: ApplyPluginsType.event,
    args: { routes },
  });

  return routes;
}
