<script lang="ts" setup>
  import { computed, ref } from 'vue';

  import { useVbenModal } from '@vben/common-ui';
  import { $t } from '@vben/locales';

  import { notification } from 'ant-design-vue';

  import { useVbenForm, z } from '#/adapter/form';
  import { authorityList, defUserService } from '#/rpc';

  const data = ref();

  const getTitle = computed(() => (data.value?.create ? '创建账号' : '编辑账号'));
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
        fieldName: 'userName',
        label: '用户名',
        componentProps: {
          placeholder: $t('ui.placeholder.input'),
        },
        rules: z.string().min(1, { message: $t('ui.formRules.required') }),
      },
      // {
      //   component: 'VbenInputPassword',
      //   fieldName: 'password',
      //   label: '密码',
      //   componentProps: {
      //     passwordStrength: true,
      //     placeholder: $t('ui.placeholder.input'),
      //   },
      //   rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
      // },
      {
        component: 'Select',
        fieldName: 'authority',
        label: '权限',
        componentProps: {
          placeholder: $t('ui.placeholder.select'),
          options: authorityList,
        },
        rules: z.string().min(1, { message: $t('ui.formRules.selectRequired') }),
      },
      {
        component: 'TreeSelect',
        fieldName: 'orgId',
        label: '所属部门',
        componentProps: {
          placeholder: $t('ui.placeholder.select'),
        },
        // rules: z.string().min(1, { message: $t('authentication.orgErrorTip') }),
      },
      {
        component: 'Input',
        fieldName: 'nickName',
        label: '昵称',
        componentProps: {
          placeholder: $t('ui.placeholder.input'),
        },
        rules: z.string().min(1, { message: $t('authentication.nicknameErrorTip') }),
      },
      {
        component: 'Input',
        fieldName: 'email',
        label: '邮箱',
        componentProps: {
          placeholder: $t('ui.placeholder.input'),
        },
        rules: z.string().min(1, { message: $t('authentication.emailValidErrorTip') }),
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
          ? defUserService.CreateUser({ user: values.results })
          : defUserService.UpdateUser({
              user: values.results,
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
