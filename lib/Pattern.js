export const MAX = (52 * 51 * 50 * 49 * 48) / (5 * 4 * 3 * 2 * 1)
export const ROYAL_FLUSH = 4
export const STRIGHT_FLUSH = 4 * 9
export const FOUR_OF_A_KIND = 13 * 48
export const FULL_HOUSE = 52 * 72
export const FLUSH = ((13 * 12 * 11 * 10 * 9) / (5 * 4 * 3 * 2 * 1) - 10) * 4
export const STRIGHT = (4 * 4 * 4 * 4 * 4 - 4) * 10
export const THREE_OF_A_KIND = ((48 * 47) / (2 * 1) - 72) * 52
export const TWO_PAIR = (13 * 12) / (2 * 1) * 6 * 6 * 44
export const PAIR = (12 * 11 * 10) / (3 * 2 * 1) * 4 * 4 * 4 * 78
export const ALL = ROYAL_FLUSH
  + STRIGHT_FLUSH
  + FOUR_OF_A_KIND
  + FULL_HOUSE
  + FLUSH
  + STRIGHT
  + THREE_OF_A_KIND
  + TWO_PAIR
  + PAIR
