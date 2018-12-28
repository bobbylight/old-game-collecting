import { expect } from 'chai';
import { shallowMount, Wrapper } from '@vue/test-utils';
import GameTile from '@/game-tile.vue';
import { Game } from '../../src/game-collector';

describe('game-tile.vue', () => {

  it('renders props.msg when passed', () => {

    const game: Game = {
      title: 'Test Game',
      publisher: 'Test Publisher',
      releaseDate: '2018-01-01T00:00:00Z'
    };

    const wrapper: Wrapper<GameTile> = shallowMount(GameTile, {
      propsData: { game },
    });
    expect(wrapper.text()).to.include(game.title);
  });
});
