# KC50 Close 5% vs 10% Move Summary

Basis: start from each date's close, then scan later close prices for first threshold breach. Natural days include the start date.

## Metric Comparison

| metric | 5% close move | 10% close move |
| --- | ---: | ---: |
| sample_with_breach | 1551 | 1541 |
| sample_without_breach | 3 | 13 |
| mean_days | 21.51 | 55.95 |
| median_days | 15.00 | 43.00 |
| p10_days | 5.00 | 14.00 |
| p25_days | 8.00 | 24.00 |
| p75_days | 28.00 | 81.00 |
| p90_days | 47.00 | 117.00 |
| p95_days | 61.50 | 138.00 |
| min_days | 2 | 2 |
| max_days | 108 | 286 |

## Bucket Distribution Comparison

| bucket | 5% count | 5% percentage | 10% count | 10% percentage |
| --- | ---: | ---: | ---: | ---: |
| 1-10 | 501 | 32.24% | 82 | 5.28% |
| 11-20 | 452 | 29.09% | 225 | 14.48% |
| 21-30 | 259 | 16.67% | 214 | 13.77% |
| 31-40 | 135 | 8.69% | 184 | 11.84% |
| 41-50 | 71 | 4.57% | 173 | 11.13% |
| 51-60 | 51 | 3.28% | 114 | 7.34% |
| 61-70 | 38 | 2.45% | 76 | 4.89% |
| 71-80 | 23 | 1.48% | 83 | 5.34% |
| 81-90 | 10 | 0.64% | 84 | 5.41% |
| 91-100 | 8 | 0.51% | 69 | 4.44% |
| 101-110 | 3 | 0.19% | 46 | 2.96% |
| 111-120 | 0 | 0.00% | 53 | 3.41% |
| 121-130 | 0 | 0.00% | 40 | 2.57% |
| 131-140 | 0 | 0.00% | 33 | 2.12% |
| 141-150 | 0 | 0.00% | 28 | 1.80% |
| 151-160 | 0 | 0.00% | 11 | 0.71% |
| 161-170 | 0 | 0.00% | 5 | 0.32% |
| 171-180 | 0 | 0.00% | 3 | 0.19% |
| 181-190 | 0 | 0.00% | 9 | 0.58% |
| 191-200 | 0 | 0.00% | 7 | 0.45% |
| 201-210 | 0 | 0.00% | 1 | 0.06% |
| 211-220 | 0 | 0.00% | 0 | 0.00% |
| 221-230 | 0 | 0.00% | 0 | 0.00% |
| 231-240 | 0 | 0.00% | 0 | 0.00% |
| 241-250 | 0 | 0.00% | 0 | 0.00% |
| 251-260 | 0 | 0.00% | 0 | 0.00% |
| 261-270 | 0 | 0.00% | 0 | 0.00% |
| 271-280 | 0 | 0.00% | 0 | 0.00% |
| 281-290 | 0 | 0.00% | 1 | 0.06% |
| no_breach_yet | 3 | 0.19% | 13 | 0.84% |

## Trading Read

- 5% close move: median 15 days, p75 28 days. Better fit short-cycle double-buy target.
- 10% close move: median 43 days, p75 81 days. Needs longer tenor or rolling position.
- 5% move within 30 days: 77.99%.
- 10% move within 30 days: 33.53%.
- 10% move within 90 days: 79.47%.
- KC50 reaches 5% close move much faster than 10% close move, gap is large in first 30-60 days.