<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer, z } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { useVbenForm } from '#/adapter/form';
import { menuTypeList, statusList } from '#/rpc';

const data = ref();

const getTitle = computed(() => (data.value?.create ? '创建菜单' : '编辑菜单'));
// const isCreate = computed(() => data.value?.create);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  // 所有表单项共用，可单独在表单内覆盖
  commonConfig: {
    // 所有表单项
    componentProps: {
      class: 'w-full',
    },
  },
  schema: [
    {
      component: 'RadioGroup',
      fieldName: 'type',
      label: '菜单类型',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: menuTypeList,
      },
    },

    {
      component: 'Input',
      fieldName: 'name',
      label: '菜单名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
    },
    {
      component: 'TreeSelect',
      fieldName: 'parentId',
      label: '上级菜单',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'orderNo',
      label: '排序',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
    },
    {
      component: 'IconPicker',
      fieldName: 'icon',
      label: '图标',
    },
    {
      component: 'Input',
      fieldName: 'path',
      label: '路由地址',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
    },
    {
      component: 'Input',
      fieldName: 'component',
      label: '组件路径',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
    },
    {
      component: 'Input',
      fieldName: 'permissionCode',
      label: '权限标识',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
    },
    {
      component: 'RadioGroup',
      fieldName: 'status',
      label: '状态',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: statusList,
      },
      rules: 'selectRequired',
    },
    {
      component: 'Switch',
      fieldName: 'isExt',
      label: '是否外链',
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'keepAlive',
      label: '是否缓存',
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'show',
      label: '是否显示',
      componentProps: {
        class: 'w-auto',
      },
    },
  ],
});

const [Drawer, drawerApi] = useVbenDrawer({
  onCancel() {
    drawerApi.close();
  },
  onConfirm() {
    console.log('onConfirm');
  },
  onOpenChange(isOpen) {
    if (isOpen) {
      data.value = drawerApi.getData<Record<string, any>>();
      baseFormApi.setValues(data.value?.row);

      // setLoading(true);
      setLoading(false);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>
<template>
  <Drawer :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
