<template>
  <div id="QuickUse">
    <AskManage />
    <AddManage />
    <AddUserPublicKey />
  </div>
</template>

<script>
import AskManage from "./AskManage.vue";
import AddManage from "./AddManage.vue";
import AddUserPublicKey from "./AddUserPublicKey.vue"

export default ({
  name: 'QuickUse',
  components: {
    AskManage, AddManage, AddUserPublicKey,
  },
});
</script>

<style scoped>
</style>
