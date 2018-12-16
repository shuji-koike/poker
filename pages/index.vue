<template>
  <section class="container">
    <h1>poker</h1>
    <div>
      <button @click="game.init()">Init</button>
      <button @click="game.step()">Step</button>
    </div>
    <PlayerView v-for="(player, index) in players" :key="index" :player="player" />
    <div>
      <CardIcon v-for="(card, index) in board" :key="index" :value="card" />
    </div>
    <div>
      <CardIcon v-for="(card, index) in deck" :key="index" :value="card" />
    </div>
    <div>
      <CardIcon v-for="(card, index) in [0, 13, 26, 10, 9]" :key="index" :value="card" />
    </div>
  </section>
</template>
<script>
import Hand from '../lib/Hand'
import Game from '../lib/Game'
import CardIcon from '../components/CardIcon'
import PlayerView from '../components/PlayerView'

export default {
  components: {
    CardIcon,
    PlayerView,
  },
  data() {
    const game = new Game()
    return {
      seats: 5,
      game,
      players: game.players,
      board: game.board
    }
  },
  computed: {
    deck() {
      return Array.from(this.game.deck).sort((a, b) => a.value - b.value)
    }
  },
  async mounted() {
    console.log(Hand.from([0, 12, 11, 10, 9]).rate)
    console.log(Hand.from([12, 11, 10, 9, 8]).rate)
    console.log(Hand.from([0, 13, 11, 10, 9]).rate)
    await Promise.all(Array.from(Array(this.seats), (_, i) => this.game.addPlayer(i)))
    await this.game.init()
    await this.game.step()
    // await this.game.step()
    // await this.game.step()
    // while(await this.game.step()) {}
    console.log(this.game)
    console.log(this.players[0].nuts)
  }
}
</script>
