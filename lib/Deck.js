import lodash from 'lodash'
import Card from './Card'
export default class Deck {
  static create() {
    return lodash.range(0, 52).map(a => new Card(a))
  }
  static random() {
    return lodash.shuffle(Deck.create())
  }
  static from(arr) {
    return arr.map(a => new Card(a))
  }
}
