<template>
    <v-toolbar app>
        <v-toolbar-title class="headline text-uppercase">
            <span>Game Collector</span>
            <span class="font-weight-light">NES</span>
        </v-toolbar-title>

        <v-toolbar-items>
            <v-text-field
                v-model="filter"
                label="Filter"
                outline
            ></v-text-field>
        </v-toolbar-items>

        <v-spacer></v-spacer>

        <v-toolbar-items class="hidden-sm-and-down">
            <v-btn
                flat
                href="https://github.com/vuetifyjs/vuetify/releases/latest"
                target="_blank">
                <span class="mr-2">My Games</span>
            </v-btn>
            <v-btn
                flat
                href="https://github.com/vuetifyjs/vuetify/releases/latest"
                target="_blank">
                <span class="mr-2">All Games</span>
            </v-btn>
        </v-toolbar-items>
    </v-toolbar>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';

@Component({ components: { } })
export default class NavBar extends Vue {

    filter: string = '';

    @Watch('filter')
    onFilterChanged(newValue: string) {
        this.$store.commit('setFilter', newValue);
    }
}
</script>
