-- 默认的角色
TRUNCATE TABLE kratos_monolithic.public.roles;
INSERT INTO kratos_monolithic.public.roles(id, parent_id, create_by, order_no, name, code, status, remark, create_time)
VALUES (1, null, 0, 1, '超级管理员', 'SYS_ADMIN', 'ON', '超级管理员拥有对系统的最高权限', now()),
       (2, null, 0, 2, '管理员', 'SYS_MANAGER', 'ON', '系统管理员拥有对整个系统的管理权限', now()),
       (3, null, 0, 3, '普通用户', 'CUSTOMER_USER', 'ON', '普通用户没有管理权限，只有设备和APP的使用权限', now()),
       (4, null, 0, 4, '游客', 'GUEST_USER', 'ON', '游客只有非常有限的数据读取权限', now());

-- 默认的组织
TRUNCATE TABLE kratos_monolithic.public.organizations;
INSERT INTO kratos_monolithic.public.organizations(id, parent_id, order_no, name, status, create_time)
VALUES (1, null, 1, '华东分部', 'ON', now()),
       (10, 1, 1, '研发部', 'ON', now()),
       (11, 1, 2, '市场部', 'ON', now()),
       (12, 1, 3, '商务部', 'ON', now()),
       (13, 1, 4, '财务部', 'ON', now()),

       (2, null, 2, '华南分部', 'ON', now()),
       (20, 2, 1, '研发部', 'ON', now()),
       (21, 2, 2, '市场部', 'ON', now()),
       (22, 2, 3, '商务部', 'ON', now()),
       (23, 2, 4, '财务部', 'ON', now()),

       (3, null, 3, '西北分部', 'ON', now()),
       (30, 3, 1, '研发部', 'ON', now()),
       (31, 3, 2, '市场部', 'ON', now()),
       (32, 3, 3, '商务部', 'ON', now()),
       (33, 3, 4, '财务部', 'ON', now())
;

-- 后台目录
TRUNCATE TABLE public.menus;
INSERT INTO public.menus(id, parent_id, type, name, path, redirect, component, status, create_time, meta)
VALUES (1, null, 'FOLDER', 'Dashboard', '/', null, 'BasicLayout', 'ON', now(), '{"order":-1, "title":"page.dashboard.title", "icon":"lucide:layout-dashboard", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (2, 1, 'MENU', 'Analytics', '/analytics', null, '#/views/dashboard/analytics/index.vue', 'ON', now(), '{"order":-1, "title":"page.dashboard.analytics", "icon":"lucide:area-chart", "affixTab": true, "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),

       (6, null, 'FOLDER', 'System', '/system', null, 'BasicLayout', 'ON', now(), '{"order":2000, "title":"menu.system.moduleName", "icon":"lucide:settings", "keepAlive":true, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (7, 6, 'MENU', 'UserManagement', 'users', null, '#/views/app/system/users/index.vue', 'ON', now(), '{"order":1, "title":"menu.system.user", "icon":"lucide:users", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (8, 6, 'MENU', 'UserDetail', 'users/detail/:id', null, '#/views/app/system/users/detail/index.vue', 'ON', now(), '{"order":2, "title":"menu.system.user_detail", "icon":"", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":true, "hideInTab":false}'),
       (9, 6, 'MENU', 'MenuManagement', 'menus', null, '#/views/app/system/menu/index.vue', 'ON', now(), '{"order":3, "title":"menu.system.menu", "icon":"lucide:square-menu", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (10, 6, 'MENU', 'OrganizationManagement', 'organizations', null, '#/views/app/system/org/index.vue', 'ON', now(), '{"order":4, "title":"menu.system.org", "icon":"lucide:network", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (11, 6, 'MENU', 'RoleManagement', 'roles', null, '#/views/app/system/role/index.vue', 'ON', now(), '{"order":5, "title":"menu.system.role", "icon":"lucide:shield-check", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}'),
       (12, 6, 'MENU', 'PositionManagement', 'positions', null, '#/views/app/system/position/index.vue', 'ON', now(), '{"order":6, "title":"menu.system.position", "icon":"lucide:id-card", "keepAlive":false, "hideInBreadcrumb":false, "hideInMenu":false, "hideInTab":false}');
