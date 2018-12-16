import 'lodash.combinations'
import lodash from 'lodash'
import Card from './Card'
import Rank from './Rank'
import {HAND_SIZE} from './Const'
import * as Pattern from './Pattern'

export default class Hand {
  constructor(pocket = [], board = []) {
    this.cards = [...pocket, ...board]
    if (this.cards.length > HAND_SIZE) {
      this.combinations = lodash.combinations(this.cards, HAND_SIZE).sort((a, b) => Rank.rank(a) - Rank.rank(b))
    } else {
      this.combinations = [this.cards]
    }
  }
  toString() {
    return `rate: ${this.rate}`
  }
  get best() {
    return this.combinations[0]
  }
  get rank() {
    return Rank.rank(this.best)
  }
  get rate() {
    return 100 - Math.floor(this.rank / Pattern.MAX * 10000) / 100
  }
  static from(arr) {
    return new Hand(arr.map(a => Card.from(a)))
  }
  static inspect(cards) {
    return {
      isRoyalFlush: this.isRoyalFlush(cards),
      isStrightFlush: this.isStrightFlush(cards),
      isFourOfAKind: this.isFourOfAKind(cards),
      isFullHouse: this.isFullHouse(cards),
      isFlush: this.isFlush(cards),
      isStright: this.isStright(cards),
      isStrightAce: this.isStrightAce(cards),
      isThreeOfAKind: this.isThreeOfAKind(cards),
      isTwoPair: this.isTwoPair(cards),
      isPair: this.isPair(cards),
      isHighCard: this.isHighCard(cards),
    }
  }
}
