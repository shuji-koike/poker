all_cards_count = 45
def no_vs_no():
    outs_count = 6
    raitio = outs_count / all_cards_count + outs_count / (all_cards_count - 1 )
    print(raitio * 100)

def one_vs_no(no_hit_hands_num, flop_num):
    if all([i > sorted(flop_num)[2] for i in no_hit_hands_num[0]]):
        # 2over (imgs/one_vs_no1.png)
        outs_count = 6
        raitio = outs_count / all_cards_count + outs_count / (all_cards_count - 1 )
        print('2')
    elif any([i > sorted(flop_num)[2] for i in no_hit_hands_num[0]]):
        # 1over
        outs_count = 3
        raitio = outs_count / all_cards_count + outs_count / (all_cards_count - 1 )
        print('1')
    elif not all([i > sorted(flop_num)[2] for i in no_hit_hands_num[0]]):
        # 0over
        outs_count = 6
        raitio = outs_count / all_cards_count * (outs_count - 1) / (all_cards_count - 1 )
        print('0')
    print(raitio * 100)

def one_vs_one():
    outs_count = 5
    raitio = outs_count / all_cards_count + outs_count / (all_cards_count - 1 )
    print(raitio * 100)

def one_vs_two():
    outs_count = 2
    raitio = outs_count / all_cards_count + outs_count / (all_cards_count - 1 )
    print(raitio * 100)
