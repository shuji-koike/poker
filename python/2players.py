# -*- coding: utf-8 -*-

import numpy as np
from calculation_2players import *

# marks spade = 0, club = 1, heart = 2, diamond = 3
# Ace = 100
#all_cards = [list(range(1, 14, 1)), list(range(1, 14, 1)), list(range(1, 14, 1)), list(range(1, 14, 1))]

# one vs two
hands_num = [[2, 2], [3, 12]]
# one vs no(2over)
#hands_num = [[12, 8], [100, 13]]
# one vs no(1over)
#hands_num = [[9, 7], [10, 2]]
# one vs no(0over)
#hands_num = [[13, 7], [4, 2]]
# one vs one
#hands_num = [[13, 7], [12, 2]]

hands_mark = [[0, 2], [3, 1]]
flop_num = [3, 7, 12]
flop_mark = [2, 3, 0]
rank = [10,10]
is_possible_straight = [False, False]
is_possible_flash = [[False, False, False, False,], [False, False, False, False,]]

for i, hand in enumerate(hands_num):
    # ペアの有無
    if not hand[0] == hand[1]:
        if hand[0] in flop_num and hand[1] in flop_num:
            rank[i] = 8
        elif hand[0] in flop_num or hand[1] in flop_num:
            rank[i] = 9
        else:
            rank[i] = 10
    else:
        rank[i] = 7 if hand[0] in flop_num else 9

    # ストレートの有無
    if rank[i] == 10:
        flop_find_straight = np.append(flop_num, [hand[0], hand[1]])
    elif rank[i] == 9:
        flop_find_straight = np.append(flop_num, [hand[0]])
    elif rank[i] == 8:
        flop_find_straight = np.array(flop_num)
    flop_find_straight = np.sort(flop_find_straight)
    hello1 = np.where((flop_find_straight >= flop_find_straight[0]) & (flop_find_straight <= flop_find_straight[0] + 4))
    is_possible_straight[i] = False if hello1[0].size < 3 else True
    hello1 = np.where((flop_find_straight >= flop_find_straight[1]) & (flop_find_straight <= flop_find_straight[1] + 4))
    is_possible_straight[i] = False if hello1[0].size < 3 else True
    if rank[0] == 10:
        hello1 = np.where((flop_find_straight >= flop_find_straight[2]) & (flop_find_straight <= flop_find_straight[2] + 4))
        is_possible_straight[i] = False if hello1[0].size < 3 else True

    # フラッシュの有無
    flop_find_flash = np.append(flop_mark, [hands_mark[i][0], hands_mark[i][1]])
    for k in range(4):
        if np.sum(flop_find_flash == k) >= 3:
            is_possible_flash[i][k] = True

# アウツを見つける
# ストレートもフラッシュもない場合
if not any(is_possible_flash[0]) and not any(is_possible_flash[1]) and not is_possible_straight[0] and not is_possible_straight[1]:
    print("no-no")
    # ノーヒット vs ノーヒット (imgs/no_vs_no.png)
    if all([i == 10 for i in rank]):
        no_vs_no()
        print("no-no-no")

    #  one vs no (a.png)
    if sum([i == 10 for i in rank]) == len(rank) -1 and sum([i == 9 for i in rank]) == 1:
        no_hit = [i for i, x in enumerate(rank) if x == 10]
        no_hit_hands_num = []
        for i in no_hit:
            no_hit_hands_num.append(hands_num[i])
        one_vs_no(no_hit_hands_num, flop_num)
        print("one-no")

    # one vs one
    if rank == [9, 9]:
        one_vs_one()
        print("one-one")

    #  one vs two (imgs/a.png)
    if rank == [8, 9] or rank == [9, 8]:
        one_vs_two()
        print("one-two")

# Gutshot

# Buble Gutshot

# Open End

# Flush Draw

# Open End Flush Draw

rank1 = [10,9,10]
hands_num1 = [[12, 8], [1, 13], [2, 3]]

print(rank)
print(is_possible_straight)
print(is_possible_flash)