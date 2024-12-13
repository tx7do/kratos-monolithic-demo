<script lang="ts" setup>
import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { useVbenForm, z } from '#/adapter/form';

const data = ref();

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
      component: 'VbenInputPassword',
      fieldName: 'old_password',
      label: '当前密码',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
    },
    {
      component: 'VbenInputPassword',
      fieldName: 'new_password',
      label: '新密码',
      componentProps: {
        passwordStrength: true,
        placeholder: $t('ui.placeholder.input'),
      },
      rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
    },
    {
      component: 'VbenInputPassword',
      fieldName: 'confirm_password',
      label: '确认密码',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      rules: z.string().min(1, { message: $t('authentication.passwordTip') }),
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
  <Modal title="修改密码">
    <BaseForm />
  </Modal>
</template>
