<template>
  <div class="CardIcon" :class="suit">{{ card.toString() }}</div>
</template>
<script>
import Card from '../lib/Card'

export default {
  name: "CardIcon",
  props: {
    value: {
      type: [Number, Card]
    }
  },
  computed: {
    card() {
      return Card.from(this.value)
    },
    suit() {
      return ["Spade", "Heart", "Diamond", "Club"][this.card.suit]
    }
  }
}
</script>
<style>
.CardIcon {
  display: inline-block;
  font-size: 5em;
  line-height: 1em;
}
.CardIcon.Spade {
  color: black
}
.CardIcon.Heart {
  color: red
}
.CardIcon.Diamond {
  color: darkred
}
.CardIcon.Club {
  color: darkblue
}
</style>
