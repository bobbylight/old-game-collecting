<template>
    <div>
        <h3 class="game-tile-title">{{game.title}}</h3>
        <div class="game-tile-publisher">{{game.publisher}}</div>
        <div class="game-tile-release-date">{{game.releaseDate}}</div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Game } from './game-collector';

@Component
export default class GameTile extends Vue {

    @Prop({ required: true })
    private readonly game!: Game;
}
</script>

<style lang="less">
.game-tile {
}
</style>