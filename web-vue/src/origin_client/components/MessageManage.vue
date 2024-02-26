<template>
    <div>
        <h2>{{this.name}}</h2>
        <DataTable :page="name"></DataTable>
    </div>
</template>

<script>
import DataTable from './DataTable.vue'

export default {
  name: 'MessageManage',
  components:{
    DataTable,
  },
  data(){
      return {
        name:'MessageManage',
      };
  },
  methods:{
    mounted(){
      alert(this.name);
      debugger;
    },
    created(){
      alert(this.name);
    },
  },
};
</script>

<style scoped>

</style>
