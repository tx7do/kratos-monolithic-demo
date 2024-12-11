<script lang="ts" setup>
  import { onMounted, ref } from 'vue';

  import { ListOrganization } from '/@/api/app/organization';
  import { BasicTree, TreeItem } from '/@/components/Tree';

  const emit = defineEmits(['select']);
  const treeData = ref<TreeItem[]>([]);

  async function fetch() {
    const orgData = (await ListOrganization({})) || [];
    treeData.value = orgData.items as unknown as TreeItem[];
  }

  function handleSelect(keys) {
    emit('select', keys[0]);
  }

  onMounted(() => {
    fetch();
  });
</script>

<template>
  <div class="m-4 mr-0 overflow-hidden bg-white">
    <BasicTree
      title="部门列表"
      toolbar
      search
      :click-row-to-expand="false"
      :tree-data="treeData"
      :field-names="{ key: 'id', title: 'name' }"
      @select="handleSelect"
    />
  </div>
</template>
