import Hand from './Hand'
import Nuts from './Nuts'

export default class Player {
  constructor(id, name) {
    this.id = id
    this.name = name || id
    this.cards = []
  }
  async init(game) {
    await this.sleep()
    this.game = game
    this.cards.length = 0
  }
  get hand() {
    return new Hand(this.cards, this.game ? this.game.board : [])
  }
  get nuts() {
    return new Nuts(this.cards, this.game ? this.game.board : [], this.game ? this.game.deck : [])
  }
  sleep(time) {
    return new Promise(resolve => setTimeout(resolve, time || 10))
  }
}
