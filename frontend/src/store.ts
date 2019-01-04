import Vue from 'vue';
import Vuex, { Store } from 'vuex';
import { AppState } from '@/game-collector';

Vue.use(Vuex);

const store: Store<AppState> = new Store({

    state: {
        filter: ''
    },

    mutations: {
        setFilter(state: AppState, filter: string) {
            state.filter = filter;
        }
    },

    actions: {}
});

export default store;
