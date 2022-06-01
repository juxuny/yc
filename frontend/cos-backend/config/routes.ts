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
        icon: 'settings',
        component: './account/Settings'
      }
    ]
  },
  {
    path: '/config',
    name: 'config',
    icon: 'user',
    routes: [
      {
        path: '/config/namespace',
        name: 'namespace',
        icon: 'settings',
        component: './config/Namespace'
      },
      {
        path: '/config/config-management',
        name: 'config-management',
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
