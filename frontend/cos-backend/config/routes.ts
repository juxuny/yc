﻿export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            component: './user/Login',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },
  {
    path: '/welcome',
    name: 'welcome',
    icon: 'smile',
    component: './Welcome',
  },
  {
    path: '/admin',
    name: 'admin',
    icon: 'crown',
    access: 'canAdmin',
    component: './Admin',
    routes: [
      {
        path: '/admin/sub-page',
        name: 'sub-page',
        icon: 'smile',
        component: './Welcome',
      },
      {
        component: './404',
      },
    ],
  },
  {
    name: 'list.table-list',
    icon: 'table',
    path: '/list',
    component: './TableList',
  },
  {
    path: '/system',
    name: 'system',
    icon: 'setting',
    routes: [
      {
        path: '/system/user-management',
        name: 'user-management',
        icon: 'user',
        component: './system/UserManagement',
      },
      {
        path: '/system/role-management',
        name: 'role-management',
        component: './system/RoleManagement',
      },
      {
        path: '/system/permission-management',
        name: 'permission-management',
        component: './system/UserManagement',
      },
      {
        path: '/system/menu-management',
        name: 'menu-management',
        component: './system/UserManagement',
      },
    ],
  },
  {
    path: '/',
    redirect: '/welcome',
  },
  {
    component: './404',
  },
];
