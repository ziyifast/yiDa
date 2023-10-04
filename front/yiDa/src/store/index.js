import Vue from 'vue'
import Vuex from 'vuex'
import user from './user'
import createPersistedState from "vuex-persistedstate"

Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
        user,
    },
    plugins: [createPersistedState({
		storage: window.localStorage
	})]
})

export default store