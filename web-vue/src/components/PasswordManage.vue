<template>
    <div>
        <h2>{{name}}</h2>
        <DataTable :page="name"></DataTable>
    </div>
</template>

<script>
import DataTable from './DataTable.vue'

export default {
  name: 'PasswordManage',
  components:{
    DataTable,
  },
  data(){
    return {
      name:'PasswordManage',
    }
  },
};
</script>

<style scoped>

</style>
