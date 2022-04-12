### balanced range
#### situation 1
+ hero has pocket TsTd
+ assume that opponent's range consists of AA and 22
    + then opponent's consists of 50% range AA combos and 50%s range 22 combos
    + more specifically, opponent's range consists of the following holecard:
    + 8.3333%AsAc, 8.3333%AsAd, 8.3333%AsAh, 8.3333%AcAd, 8.3333%AcAh, 8.3333%AdAh, 8.3333% 2s2c, 8.3333% 2s2d, 8.3333% 2s2h, 8.3333% 2c2d, 8.3333% 2c2h, 8.3333% 2d2h
+ hero's TT versus each of these holecards equity are
    + 19.01% vs AsAc, 19.01% vs AsAd, 19.01% vs AsAh, 19.01% vs AcAd, 19.01% vs AcAh, 19.01% vs AdAh, 81.87% vs 2s2c, 81.87% vs 2s2d, 81.87% vs 2s2h, 81.87% vs 2c2d, 81.87% vs 2c2h, 81.87% vs 2d2h
+ then hero's TT vs opponent's range AA, 22 overall equity can be calculated:
    + (19.01*12.5%)*4 + 81.87*12.5%*4 = 50.44%

#### situation 2
+ hero has pocket TsTd
+ assume that opponent's range consists of AA, AKs, 22, A5s
    + then opponent's range consists of 6 AA combos, 4 AKs combos, 6 22 combos, 4 A5s combos
    + that is: 30% AA, 20% AKs, 30% 22, 20% A5s
+ hero's TT versus each of these range equity are
    + 19.01% vs AA, 53.92% vs AKs, 81.87% vs 22, 66.95% vs A5s
+ then hero's TT vs opponent's range AA, 22, AKs, A5s overall equity can be calculated:
    + 19.01*30% + 53.92*20% + 81.87*30% + 66.95*20% = 54.438%

### unbalanced range
#### situation 3
+ hero has pocket TsTd
+ assume that opponent's range consists of AA, 22
+ but hero thinks opponent's range consists of uneven combos of AA,22 according to opponent's tendency or actions
+ assume that hero thinks opponent's range consists of 80% AA, 20% 22
+ then hero's TT vs opponent's range 80% AA, 20% 22 overall equity can be calculated:
    + 19.01 * 80% + 81.87 * 20% = 31.582%

#### situation 4
+ hero has pocket TsTd
+ assume that opponent's range consists of AA, AKs, 22, A5s
+ but hero thinks opponent's range consists of uneven combos of AA, AKs, 22, A5s according to opponent's tendency or actions
+ assume that hero thinks opponent's range consists of 80% (AA, AKs), 20% (22, A5s)
+ then hero's TT vs opponent's range 80% (AA,AKs), 20% (22,A5s) overall equity can be calculated:
    + (19.01 * 6/10 + 53.92 * 4/10) * 80% + (81.87*6/10 + 66.95*4/10) * 20% = 41.56%