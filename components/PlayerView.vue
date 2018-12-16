<template>
  <section class="PlayerView">
    <span>{{ player.name }}</span>
    <CardIcon v-for="(card, index) in cards" :key="index" :value="card" />
    <div>
      <div>{{ handName }}</div>
      <div>
        <button>Fold</button>
        <button>Check</button>
        <button>Raise</button>
        <button @click="debug()">Debug</button>
      </div>
    </div>
    <div>
      <CardIcon v-for="(card, j) in player.hand.best" :key="j" :value="card" />
    </div>
    <div v-if="player.nuts">
      <CardIcon v-for="(card, j) in player.nuts.best" :key="j" :value="card" />
    </div>
  </section>
</template>
<script>
import Player from '../lib/Player'
import CardIcon from '../components/CardIcon.vue'

export default {
  name: "PlayerView",
  components: {
    CardIcon,
  },
  props: {
    player: Player
  },
  computed: {
    cards() {
      return this.player.cards
    },
    handName() {
      return this.player.hand.toString()
    }
  },
  methods: {
    debug() {
      console.debug(this);
      console.debug(this.player.nuts)
    }
  }
}
</script>
<style>
.PlayerView > div {
  display: inline-block;
}
</style>
