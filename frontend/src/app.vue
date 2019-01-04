<template>
    <v-app>
        <NavBar/>

        <v-content>
            <GameList :filter="$store.state.filter"/>
        </v-content>
    </v-app>
</template>

<script>
import GameList from '@/game-list.vue';
import NavBar from '@/navbar.vue';

export default {
    name: 'App',
    components: {
        GameList,
        NavBar
    },
    data() {
        return {
            //
        };
    }
};
</script>
