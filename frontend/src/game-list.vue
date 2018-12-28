<template>
  <v-container>
    <v-layout
      text-xs-center
      wrap
    >
      <v-flex xs12>
        <v-img
          :src="require('./assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        ></v-img>
        Roberto is simply amazing
      </v-flex>

      <v-flex mb-4 xs12>
        <h1 class="display-2 font-weight-bold mb-3">
          Games
        </h1>
        <game-tile v-for="(game, i) in games" :game="game"></game-tile>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import GameTile from '@/game-tile.vue';
import { Game } from '@/game-collector';

@Component({ components: { GameTile } })
export default class HelloWorld extends Vue {

  games: Game[] = [];

  mounted() {
      this.games.push({
          title: 'Super Mario Bros.', publisher: 'Nintendo', releaseDate: '1984-01-01T00:00:00Z'
      }, {
          title: 'Super Mario Bros. 2', publisher: 'Nintendo', releaseDate: '1987-01-01T00:00:00Z'
      },
      );
  }
}
</script>

<style>

</style>
