<template>
  <v-container>
    <v-layout
      text-xs-center
      wrap
    >
      <v-flex mb-4 xs12>
        <h1 class="display-2 font-weight-bold mb-3">
          Games
        </h1>

          <v-data-table
              hide-headers
              class="search-result-table"
              :items="games"
              :pagination.sync="pagination"
              :total-items="totalItems"
              :loading="loading"
              :rows-per-page-items='[ 20, 50, 100 ]'
          >

              <template slot="items" slot-scope="props">
                  <td>
                      <game-tile :game="props.item"></game-tile>
                  </td>
              </template>
          </v-data-table>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import GameTile from '@/game-tile.vue';
import { Game, PagedData } from '@/game-collector';
import debounce from 'debounce';
import restApi from '@/rest-api';

@Component({ components: { GameTile } })
export default class GameList extends Vue {

    @Prop({ required: true })
    filter!: string;

    games: Game[] = [];

    blazy: any = '';
    totalItems: number = 0;
    loading: boolean = true;
    pagination: any = {};

    created() {
        this.reloadTable = debounce(this.reloadTable, 750);
    }

    mounted() {
        this.reloadTable();
    }

    @Watch('pagination')
    onPaginationChanged() {
        console.log('Pagination changed');
        this.reloadTable();
    }

    @Watch('filter')
    onFilterChanged() {
        this.reloadTable();
    }

    private reloadTable() {

        this.loading = true;

        const { sortBy, descending, page, rowsPerPage } = this.pagination;

        restApi.getGames((page - 1) * rowsPerPage, rowsPerPage, this.$store.state.filter)
            .then((games: PagedData<Game>) => {
                console.log(`Setting games to: ${JSON.stringify(games)}`);
                this.games = games.data;
                this.totalItems = games.total;
            });
    }
}
</script>

<style>

</style>
