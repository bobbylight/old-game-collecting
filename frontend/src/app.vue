<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>Vuetify</span>
        <span class="font-weight-light">Old Man Games</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        flat
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
      >
        <span class="mr-2">Latest Release Robert</span>
      </v-btn>
    </v-toolbar>

    <v-content>
      <GameList/>
    </v-content>
  </v-app>
</template>

<script>
import GameList from '@/game-list.vue';

export default {
  name: 'App',
  components: {
    GameList
  },
  data () {
    return {
      //
    };
  }
};
</script>
