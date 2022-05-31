export default [
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
    path: '/account',
    name: 'account',
    icon: 'user',
    routes: [
      {
        path: '/account/settings',
        name: 'settings',
        icon: 'settings',
        component: './account/Settings'
      }
    ]
  },
  {
    path: '/config',
    name: 'Config',
    icon: 'user',
    routes: [
      {
        path: '/config/namespace',
        name: 'Namespace',
        icon: 'settings',
        component: './config/Namespace'
      },
      {
        path: '/config/config-management',
        name: 'ConfigManagement',
        icon: 'settings',
        component: './config/ConfigManagement'
      }
    ]
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
      }
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
