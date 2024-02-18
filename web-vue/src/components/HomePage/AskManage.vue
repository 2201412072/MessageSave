<template>
    <div id="AskManage" title="Search the password app wanted.">
        <span>Search</span>
        <div id="input-div" class="input-container">
          <el-input class="input" type="text" v-model="key_word" placeholder="please input the key word." ></el-input>
          <el-button class="input-button" @click="ask">Ask</el-button>
        </div>
        <div id="text-show" class="text-container">
          <div class="scrollable-container">
            <el-scrollbar class="scrollbar-wrapper">
              <div class="scrollable-content">
              {{ textContent }}
              </div>
            </el-scrollbar>
          </div>
          <el-button class="copy-button" @click="copy">Copy</el-button>
        </div>
    </div>
</template>

<script>

export default ({
  name: 'AskManage',
  components: {
  },
  data() {
    return {
      key_word: '',
      textContent: '',
    };
  },
  methods: {
    ask() {
      console.log(this.key_word);
    },
    copy(){
      console.log(this.textContent);
    }
  },
});
</script>

<style scoped>
.input-container{
  display: flex;
  align-items: center;
  width: 100%;
}

.text-container{
  display: flex;
  align-items: center;
  width: 100%;
}

.scrollable-container {
  flex: 1;
  height: 200px; /* 设置容器的高度 */
  overflow: auto; /* 添加滚动条 */
  border: 1px solid;
}

.scrollbar-wrapper{
  height: 100%;
  overflow-x: hidden !important;
}

.scrollable-content {
  padding: 10px; /* 可根据需求设置内容的内边距 */
  color: #8a8a8a;
}

.copy-button{
  margin-left: auto;
}
</style>
