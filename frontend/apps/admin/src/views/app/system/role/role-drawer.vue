<script lang="ts" setup>
  import { computed, ref } from 'vue';

  import { useVbenDrawer, z } from '@vben/common-ui';
  import { $t } from '@vben/locales';

  import { useVbenForm } from '#/adapter/form';
  import { statusList } from '#/rpc';

  const data = ref();

  const getTitle = computed(() => (data.value?.create ? '创建角色' : '编辑角色'));
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
        component: 'Input',
        fieldName: 'name',
        label: '角色名称',
        rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
      },
      {
        component: 'Input',
        fieldName: 'code',
        label: '角色值',
        rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
      },

      {
        component: 'RadioGroup',
        fieldName: 'status',
        label: '状态',
        componentProps: {
          isButton: true,
          class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
          options: statusList,
        },
      },
      {
        component: 'Textarea',
        fieldName: 'remark',
        label: '备注',
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
