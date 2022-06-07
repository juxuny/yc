
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
    path: '/account',
    name: 'account',
    icon: 'user',
    routes: [
      {
        path: '/account/settings',
        name: 'settings',
        component: './account/Settings'
      },
      {
        path: '/account/access-key',
        name: 'access-key',
        component: './account/AccessKeyManagement'
      }
    ]
  },
  {
    path: '/config',
    name: 'config',
    icon: 'database',
    routes: [
      {
        path: '/config/namespace',
        name: 'namespace',
        component: './config/Namespace'
      },
      {
        path: '/config/config-management',
        name: 'config-management',
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
