// store.js
import { createStore } from 'vuex';

const store = createStore({
  state() {
    return {
      username: null
    };
  },
  mutations: {
    setUsername(state, value) {
      state.username = value;
    }
  }
});

export default store;