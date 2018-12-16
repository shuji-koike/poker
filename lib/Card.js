export default class Card {
  constructor(value) {
    Card.validate(value)
    this.value = value
  }
  get rank() {
    return Card.int2rank(this.value)
  }
  get suit() {
    return Card.int2suit(this.value)
  }
  toString() {
    return Card.int2unicode(this.value)
  }
  static from(value) {
    if (typeof value == 'number') {
      return new Card(value)
    } else if (value instanceof Card) {
      return value
    }
  }
  static validate(value) {
    if (typeof value != 'number') throw new Error()
    if (value < 0) throw new Error()
    if (value > 51) throw new Error()
    return true
  }
  static int2unicode(value) {
    Card.validate(value)
    return String.fromCodePoint(
      0x01F0A1 + Card.int2rank(value) + (Card.int2rank(value) > 10 ? 1 : 0)
      + 0x10 * Card.int2suit(value))
  }
  static int2string(value) {
    Card.validate(value)
    return [
      Card.rank2char(Card.int2rank(value)),
      Card.suit2char(Card.int2suit(value))
    ].join("")
  }
  static int2rank(value) {
    return value % 13
  }
  static int2suit(value) {
    return Math.floor(value / 13)
  }
  static rank2char(value) {
    switch(value) {
      case 0: return "A"
      case 9: return "T"
      case 10: return "J"
      case 11: return "Q"
      case 12: return "K"
      default: return String(value + 1)
    }
  }
  static suit2char(value) {
    switch(value) {
      case 0: return "s"
      case 1: return "h"
      case 2: return "d"
      case 3: return "c"
      default: throw new Error()
    }
  }
}
