# CSI1000 High/Low 5% vs 10% Move Summary

Basis: start from each date's close, then scan later daily high/low for first threshold breach. Natural days include the start date.

## Metric Comparison

| metric | 5% high/low move | 10% high/low move |
| --- | ---: | ---: |
| sample_with_breach | 5191 | 5167 |
| sample_without_breach | 7 | 31 |
| up_breach | 2439 | 2509 |
| down_breach | 2751 | 2658 |
| both_side_same_day_breach | 1 | 0 |
| mean_days | 19.80 | 59.38 |
| median_days | 14.00 | 41.00 |
| p10_days | 4.00 | 11.00 |
| p25_days | 7.00 | 21.00 |
| p75_days | 26.00 | 79.00 |
| p90_days | 44.00 | 134.00 |
| p95_days | 59.00 | 179.00 |
| min_days | 2 | 2 |
| max_days | 140 | 400 |

## Bucket Distribution Comparison

| bucket | 5% count | 5% percentage | 10% count | 10% percentage |
| --- | ---: | ---: | ---: | ---: |
| 1-10 | 2007 | 38.61% | 509 | 9.79% |
| 11-20 | 1437 | 27.65% | 754 | 14.51% |
| 21-30 | 760 | 14.62% | 743 | 14.29% |
| 31-40 | 374 | 7.20% | 568 | 10.93% |
| 41-50 | 206 | 3.96% | 466 | 8.96% |
| 51-60 | 167 | 3.21% | 330 | 6.35% |
| 61-70 | 101 | 1.94% | 305 | 5.87% |
| 71-80 | 62 | 1.19% | 227 | 4.37% |
| 81-90 | 32 | 0.62% | 222 | 4.27% |
| 91-100 | 20 | 0.38% | 147 | 2.83% |
| 101-110 | 9 | 0.17% | 122 | 2.35% |
| 111-120 | 14 | 0.27% | 118 | 2.27% |
| 121-130 | 0 | 0.00% | 103 | 1.98% |
| 131-140 | 2 | 0.04% | 91 | 1.75% |
| 141-150 | 0 | 0.00% | 92 | 1.77% |
| 151-160 | 0 | 0.00% | 40 | 0.77% |
| 161-170 | 0 | 0.00% | 40 | 0.77% |
| 171-180 | 0 | 0.00% | 37 | 0.71% |
| 181-190 | 0 | 0.00% | 38 | 0.73% |
| 191-200 | 0 | 0.00% | 34 | 0.65% |
| 201-210 | 0 | 0.00% | 33 | 0.63% |
| 211-220 | 0 | 0.00% | 27 | 0.52% |
| 221-230 | 0 | 0.00% | 18 | 0.35% |
| 231-240 | 0 | 0.00% | 15 | 0.29% |
| 241-250 | 0 | 0.00% | 7 | 0.13% |
| 251-260 | 0 | 0.00% | 17 | 0.33% |
| 261-270 | 0 | 0.00% | 18 | 0.35% |
| 271-280 | 0 | 0.00% | 11 | 0.21% |
| 281-290 | 0 | 0.00% | 13 | 0.25% |
| 291-300 | 0 | 0.00% | 4 | 0.08% |
| 301-310 | 0 | 0.00% | 3 | 0.06% |
| 311-320 | 0 | 0.00% | 3 | 0.06% |
| 321-330 | 0 | 0.00% | 1 | 0.02% |
| 331-340 | 0 | 0.00% | 3 | 0.06% |
| 341-350 | 0 | 0.00% | 0 | 0.00% |
| 351-360 | 0 | 0.00% | 2 | 0.04% |
| 361-370 | 0 | 0.00% | 2 | 0.04% |
| 371-380 | 0 | 0.00% | 0 | 0.00% |
| 381-390 | 0 | 0.00% | 0 | 0.00% |
| 391-400 | 0 | 0.00% | 4 | 0.08% |
| no_breach_yet | 7 | 0.13% | 31 | 0.60% |

## Trading Read

- 5% high/low move: median 14 days, p75 26 days. This is short-cycle double-buy target.
- 10% high/low move: median 41 days, p75 79 days. This needs longer tenor or rolling position.
- 5% move within 30 days: 80.88%.
- 10% move within 30 days: 38.59%.
- 10% move within 90 days: 78.27%.
- Down breaches are slightly more frequent than up breaches in both thresholds.