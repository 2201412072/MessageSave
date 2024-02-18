<template>
    <AskManage />
    <AddManage />
</template>

<script>
import AskManage from "./AskManage.vue";
import AddManage from "./AddManage.vue";

export default ({
  name: 'QuickUse',
  components: {
    AskManage, AddManage,
  },
});
</script>

<style scoped>

</style>
