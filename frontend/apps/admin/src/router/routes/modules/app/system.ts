import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const system: RouteRecordRaw[] = [
  {
    path: '/system',
    name: 'System',
    component: BasicLayout,
    meta: {
      order: 2000,
      icon: 'lucide:settings',
      title: $t('menu.system.moduleName'),
      keepAlive: true,
    },
    children: [
      {
        path: 'users',
        name: 'UserManagement',
        meta: {
          icon: 'ion:person-outline',
          title: $t('menu.system.user'),
        },
        component: () => import('#/views/app/system/users/index.vue'),
      },
      {
        path: 'users/detail/:id',
        name: 'UserDetail',
        meta: {
          hideInTab: true,
          hideInMenu: true,
          title: $t('menu.system.user_detail'),
          currentActiveMenu: '/system/user',
        },
        component: () => import('#/views/app/system/users/detail/index.vue'),
      },

      {
        path: 'menu',
        name: 'MenuManagement',
        meta: {
          icon: 'ion:menu-outline',
          title: $t('menu.system.menu'),
        },
        component: () => import('#/views/app/system/menu/index.vue'),
      },
      {
        path: 'org',
        name: 'OrganizationManagement',
        meta: {
          icon: 'ant-design:apartment-outlined',
          title: $t('menu.system.org'),
        },
        component: () => import('#/views/app/system/org/index.vue'),
      },

      {
        path: 'role',
        name: 'RoleManagement',
        meta: {
          icon: 'ant-design:team-outlined',
          title: $t('menu.system.role'),
          hideInTab: false,
        },
        component: () => import('#/views/app/system/role/index.vue'),
      },
      {
        path: 'position',
        name: 'PositionManagement',
        meta: {
          icon: 'ion:person-circle-outline',
          title: $t('menu.system.position'),
          hideInTab: false,
        },
        component: () => import('#/views/app/system/position/index.vue'),
      },
    ],
  },
];

export default system;