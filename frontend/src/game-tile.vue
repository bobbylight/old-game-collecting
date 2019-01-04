<template>
    <div class="game-tile">
        <table>
            <tr>
                <td>
                    <div class="box-art"></div>
                </td>
                <td class="text-xs-left">
                    <h3 class="game-tile-title">{{game.name}}</h3>
                    <div class="game-tile-publisher">{{publishersString}}</div>
                    <div class="game-tile-release-date">{{naRelDate}}</div>
                </td>
            </tr>
        </table>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Game } from './game-collector';

@Component
export default class GameTile extends Vue {

    @Prop({ required: true })
    private readonly game!: Game;

    get naRelDate(): string {

        if (!this.game.naRelDate) {
            return 'Unknown Release Date';
        }

        // toLocaleDateString() converts from UTC to user's time zone, so
        // we need to add the user's time zone as an offset
        let date: Date = new Date(this.game.naRelDate);
        date = new Date(date.getTime() + date.getTimezoneOffset() * 60000);
        return date.toLocaleDateString();
    }

    get publishersString(): string {

        if (!this.game.publishers || !this.game.publishers.length) {
            return 'Unknown Publisher';
        }

        return this.game.publishers.join(', ');
    }
}
</script>

<style lang="less">
.game-tile {

    .box-art {

        width: 60px;
        height: 100px;
        border: 1px solid #404040;
    }
}
</style>
