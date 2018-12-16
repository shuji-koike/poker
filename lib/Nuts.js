import 'lodash.combinations'
import lodash from 'lodash'
import Rank from './Rank'
import {BOARD_SIZE, HAND_SIZE, POCKET_SIZE, FLOP_SIZE} from './Const'
import * as Pattern from './Pattern'

export default class Nuts {
  constructor(pocket, board, deck) {
    if (pocket.length < POCKET_SIZE || board.length < FLOP_SIZE + 1) return
    this.combinations = lodash.range(0, BOARD_SIZE - board.length + 1).map(a => [
      lodash.combinations(board, HAND_SIZE - POCKET_SIZE - a),
      lodash.combinations(deck, a)
    ]).map(a => a[0].map(e => a[1].map(f => pocket.concat(e, f)))).flat(2)
      .sort((a, b) => Rank.rank(a) - Rank.rank(b))
  }
  get best() {
    return this.combinations ? this.combinations[0] : []
  }
}
