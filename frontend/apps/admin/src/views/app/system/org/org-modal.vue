<script lang="ts" setup>
  import { computed, ref } from 'vue';

  import { useVbenModal } from '@vben/common-ui';
  import { $t } from '@vben/locales';

  import { notification } from 'ant-design-vue';

  import { useVbenForm, z } from '#/adapter/form';
  import { defOrganizationService, statusList } from '#/rpc';

  const data = ref();

  const getTitle = computed(() => (data.value?.create ? '创建部门' : '编辑部门'));
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
        label: '部门名称',
        componentProps: {
          placeholder: $t('ui.placeholder.input'),
        },
        rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
      },
      {
        component: 'TreeSelect',
        fieldName: 'parentId',
        label: '上级部门',
        componentProps: {
          placeholder: $t('ui.placeholder.select'),
        },
        rules: z.string().min(1, { message: $t('authentication.orgErrorTip') }),
      },
      {
        component: 'InputNumber',
        fieldName: 'orderNo',
        label: '排序',
        componentProps: {
          placeholder: $t('ui.placeholder.input'),
        },
        rules: z.string().min(1, { message: $t('authentication.nicknameErrorTip') }),
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

  const [Modal, modalApi] = useVbenModal({
    onCancel() {
      modalApi.close();
    },

    async onConfirm() {
      console.log('onConfirm');

      const values = await baseFormApi.validate();
      if (!values.valid) {
        return;
      }

      setLoading(true);

      console.log(getTitle.value, values);

      try {
        await (data.value?.create
          ? defOrganizationService.CreateOrganization({ org: values.results })
          : defOrganizationService.UpdateOrganization({
              org: values.results,
              updateMask: ['id', 'status'],
            }));

        notification.success({
          message: `${getTitle.value}成功`,
        });
      } catch {
        notification.success({
          message: `${getTitle.value}失败`,
        });
      } finally {
        modalApi.close();
        setLoading(false);
      }
    },

    onOpenChange(isOpen: boolean) {
      if (isOpen) {
        data.value = modalApi.getData<Record<string, any>>();
        baseFormApi.setValues(data.value?.row);
        setLoading(false);
        console.log('onOpenChange', data.value, data.value?.create);
      }
    },
  });

  function setLoading(loading: boolean) {
    modalApi.setState({ confirmLoading: loading });
  }
</script>

<template>
  <Modal :title="getTitle">
    <BaseForm />
  </Modal>
</template>
