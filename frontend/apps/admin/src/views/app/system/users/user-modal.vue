<script lang="ts" setup>
  import { computed, ref } from 'vue';

  import { useVbenModal } from '@vben/common-ui';
  import { $t } from '@vben/locales';

  import { notification } from 'ant-design-vue';

  import { useVbenForm, z } from '#/adapter/form';
  import { defUserService } from '#/rpc';

  const emit = defineEmits(['success', 'register']);

  const data = ref();

  const getTitle = computed(() => (data.value?.create ? '创建账号' : '编辑账号'));

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
        fieldName: 'userName',
        label: '用户名',
        componentProps: {
          passwordStrength: true,
          placeholder: $t('authentication.usernameTip'),
        },
        rules: z.string().min(1, { message: $t('authentication.usernameTip') }),
      },
      {
        component: 'VbenInputPassword',
        fieldName: 'password',
        label: '密码',
        componentProps: {
          passwordStrength: true,
          placeholder: $t('authentication.passwordTip'),
        },
        rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
      },
      {
        component: 'Input',
        fieldName: 'realName',
        label: '名字',
      },
      {
        component: 'Input',
        fieldName: 'email',
        label: '邮箱',
      },
      {
        component: 'Input',
        fieldName: 'phone',
        label: '手机',
      },
      {
        component: 'TreeSelect',
        fieldName: 'orgId',
        label: '所属部门',
      },
      {
        component: 'TreeSelect',
        fieldName: 'positionId',
        label: '所在岗位',
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

      modalApi.setState({ confirmLoading: true });

      console.log(getTitle.value, values);

      try {
        await (data.value?.create
          ? defUserService.CreateUser({ user: values.results })
          : defUserService.UpdateUser({
              user: { id: row.id, status: row.status },
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
        modalApi.setState({ confirmLoading: false });
      }
    },

    onOpenChange(isOpen: boolean) {
      if (isOpen) {
        data.value = modalApi.getData<Record<string, any>>();
        baseFormApi.setValues(data.value);
        modalApi.setState({ confirmLoading: false });
        console.log('onOpenChange', data.value, data.value?.create);
      }
    },
  });
</script>

<template>
  <Modal :title="getTitle">
    <BaseForm />
  </Modal>
</template>
