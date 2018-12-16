import lodash from 'lodash'
import Deck from './Deck'
import Player from './Player'
import {POCKET_SIZE, BOARD_SIZE, FLOP_SIZE} from './Const'

export default class Game {
  constructor() {
    this.players = []
    this.board = []
    this.deck = []
  }
  async addPlayer(id) {
    await this.join(new Player(id))
  }
  async join(player) {
    await this.sleep(Math.random() * 100)
    this.players.push(player)
  }
  async init() {
    this.board.length = 0
    this.deck.length = 0
    this.players.forEach(a => a.init(this))
    this.deck.push(...Deck.random())
    for (let i = 0; i < POCKET_SIZE; i++) {
      for (let j = 0; j < this.players.length; j++) {
        await this.deal(this.players[j])
      }
    }
  }
  async deal(player) {
    await this.sleep()
    if (player) {
      player.cards.push(this.deck.pop())
    } else {
      this.board.push(this.deck.pop())
    }
  }
  async flop() {
    await this.sleep()
    if (this.board.length !== 0) {
      throw new Error()
    }
    for (let i = 0; i < FLOP_SIZE; i++) {
      await this.deal()
    }
  }
  async turn() {
    if (this.board.length < FLOP_SIZE || this.board.length >= BOARD_SIZE) {
      throw new Error()
    }
    await this.deal()
  }
  async step() {
    if (this.busy) {
      return true
    } else {
      this.busy = true
    }
    if (this.board.length === 0) {
      await this.flop()
    } else if (this.board.length >= FLOP_SIZE && this.board.length < BOARD_SIZE) {
      await this.turn()
    } else {
      this.end = true
      console.debug("end")
    }
    this.busy = false
    return !this.end
  }
  sleep(time) {
    return new Promise(resolve => setTimeout(resolve, time || 10))
  }
}
