import lodash from 'lodash'
import {HAND_SIZE} from './Const'
import * as Pattern from './Pattern'

export default class Rank {
  static rank(cards) {
    if (cards.length < HAND_SIZE) return Pattern.MAX
    this.cache = this.cache || new Map([])
    const key = cards.map(b => b.value).join(",")
    let value = this.cache.get(key)
    if (value !== undefined) return value
    value = 0
    if (this.isRoyalFlush(cards)) return value
    value += Pattern.ROYAL_FLUSH
    if (this.isStrightFlush(cards)) return value
    value += Pattern.STRIGHT_FLUSH
    if (this.isFourOfAKind(cards)) return value
    value += Pattern.FOUR_OF_A_KIND
    if (this.isFullHouse(cards)) return value
    value += Pattern.FULL_HOUSE
    if (this.isFlush(cards)) return value
    value += Pattern.FLUSH
    if (this.isStright(cards)) return value
    value += Pattern.STRIGHT
    if (this.isThreeOfAKind(cards)) return value
    value += Pattern.THREE_OF_A_KIND
    if (this.isTwoPair(cards)) return value
    value += Pattern.TWO_PAIR
    if (this.isPair(cards)) return value
    value += Pattern.PAIR
    this.cache.set(key, value)
    return value
  }
  static isRoyalFlush(cards) {
    return this.isStrightAce(cards) && this.isFlush(cards)
  }
  static isStrightFlush(cards) {
    return this.isFlush(cards) && this.isStright(cards)
  }
  static isFourOfAKind(cards) {
    return Object.values(this.pairs(cards)).map(a => a.length).some(a => a === 4)
  }
  static isFullHouse(cards) {
    return Object.keys(this.pairs(cards)).length === 2
  }
  static isFlush(cards) {
    return cards.every(a => a.suit === cards[0].suit)
  }
  static isStright(cards) {
    const ranks = cards.map(a => a.rank).sort((a, b) => a - b)
    if (lodash.isEqual([0, 9, 10, 11, 12], ranks)) return true
    return lodash.every(ranks, (v, i) => v === ranks[0] + i) || this.isStrightAce(cards)
  }
  static isStrightAce(cards) {
    const ranks = cards.map(a => a.rank).sort((a, b) => a - b)
    return lodash.isEqual([0, 9, 10, 11, 12], ranks)
  }
  static isThreeOfAKind(cards) {
    return Object.values(this.pairs(cards)).map(a => a.length).some(a => a === 3)
  }
  static isTwoPair(cards) {
    return Object.keys(this.pairs(cards)).length === 3
  }
  static isPair(cards) {
    return Object.keys(this.pairs(cards)).length === 4
  }
  static isHighCard(cards) {
    return Object.keys(this.pairs(cards)).length === 5
  }
  static pairs(cards) {
    return lodash.groupBy(cards, a => a.rank)
  }
}
