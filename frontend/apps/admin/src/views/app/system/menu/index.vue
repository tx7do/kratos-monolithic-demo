<script lang="ts" setup>
  import type { VxeGridProps } from '#/adapter/vxe-table';
  import type { User } from '#/rpc/api/user/service/v1/user.pb';

  import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';

  import { Button, notification, Popconfirm, Switch } from 'ant-design-vue';

  import { useVbenVxeGrid } from '#/adapter/vxe-table';
  import { $t } from '#/locales';
  import { defMenuService, makeQueryString, statusList } from '#/rpc';

  import MenuDrawer from './menu-drawer.vue';

  const formOptions: VbenFormProps = {
    // 默认展开
    collapsed: false,
    // 控制表单是否显示折叠按钮
    showCollapseButton: false,
    // 按下回车时是否提交表单
    submitOnEnter: false,
    schema: [
      {
        component: 'Input',
        fieldName: 'name',
        label: '菜单名称',
        componentProps: {
          placeholder: $t('ui.placeholder.input'),
        },
      },
      {
        component: 'Select',
        fieldName: 'status',
        label: '状态',
        componentProps: {
          options: statusList,
          placeholder: $t('ui.placeholder.select'),
        },
      },
    ],
  };

  const gridOptions: VxeGridProps<User> = {
    toolbarConfig: {
      custom: true,
      export: true,
      // import: true,
      refresh: true,
      zoom: true,
    },
    height: 'auto',
    exportConfig: {},
    pagerConfig: {},
    rowConfig: {
      isHover: true,
    },
    stripe: true,

    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          console.log('query:', formValues);
          return await defMenuService.ListMenu({
            field_mask: null,
            orderBy: [],
            query: makeQueryString(formValues),
            page: page.currentPage,
            pageSize: page.pageSize,
          });
        },
      },
    },

    columns: [
      { title: '序号', type: 'seq', width: 50 },
      { title: '菜单名称', field: 'title', slots: { default: 'title' } },
      { title: '排序', field: 'orderNo' },
      { title: '权限标识', field: 'permissionCode' },
      { title: '路由地址', field: 'path' },
      { title: '组件路径', field: 'component' },
      { title: '状态', field: 'status', slots: { default: 'status' } },
      { title: '更新时间', field: 'updateTime', formatter: 'formatDateTime' },
      { title: '备注', field: 'remark' },
      {
        title: '操作',
        field: 'action',
        fixed: 'right',
        slots: { default: 'action' },
        width: 210,
      },
    ],
  };

  const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

  const [Drawer, drawerApi] = useVbenDrawer({
    // 连接抽离的组件
    connectedComponent: MenuDrawer,
  });

  function openDrawer(create: boolean, row?: any) {
    drawerApi.setData({
      create,
      row,
    });
    drawerApi.open();
  }

  /* 创建 */
  function handleCreate() {
    console.log('创建');

    openDrawer(true);
  }

  /* 编辑 */
  function handleEdit(row: any) {
    console.log('编辑', row);
    openDrawer(false, row);
  }

  /* 删除 */
  function handleDelete(row: any) {
    console.log('删除', row);

    try {
      defMenuService.DeleteMenu({ id: row.id });

      notification.success({
        message: '删除菜单成功',
      });

      gridApi.reload();
    } catch {
      notification.error({
        message: '删除菜单失败',
      });
    }
  }

  /* 修改菜单状态 */
  async function handleStatusChanged(row: any, checked: boolean) {
    console.log('handleStatusChanged', row.status, checked);

    row.pending = true;
    row.status = checked ? 'ON' : 'OFF';

    try {
      await defMenuService.UpdateMenu({
        menu: { id: row.id, status: row.status },
        updateMask: ['id', 'status'],
      });

      notification.success({
        message: '更新菜单状态成功',
      });
    } catch {
      notification.error({
        message: '更新菜单状态失败',
      });
    } finally {
      row.pending = false;
    }
  }
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.system.menu')">
      <template #toolbar-tools>
        <Button type="primary" @click="handleCreate">创建菜单</Button>
      </template>
      <template #title="{ row }">
        <span style="marginRight: '15px'"></span>
      </template>
      <template #status="{ row }">
        <Switch
          :checked="row.status === 'ON'"
          :loading="row.pending"
          checked-children="正常"
          un-checked-children="停用"
          @change="(checked) => handleStatusChanged(row, checked as boolean)"
        />
      </template>
      <template #action="{ row }">
        <Button type="link" @click="() => handleEdit(row)">编辑</Button>
        <Popconfirm
          cancel-text="不要"
          ok-text="是的"
          title="你是否要删除掉该菜单？"
          @confirm="() => handleDelete(row)"
        >
          <Button danger type="link">删除</Button>
        </Popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
