## Page 1

*[Image OCR]
WU
HEDGING

MANAGING
VANILLA AND
CAUTIO OPTIONS

Noceim Taloh
[End OCR]*

## Page 2

*[Image OCR]
Dynamic
Hedging
MANAGING VANILLA AND

EXOTIC OPTIONS

Nassim Taleb

JOHN WILEY & SONS, INC.
bane ¢ i

New York © Chichester ¢ Bris to ¢ Singapore * Weinheim
[End OCR]*

## Page 3

*[Image OCR]
Preface

After closing about 200,000 option transactions! (that is separate option
tickets) over 12 years and studying about 70,000 risk management reports, I
felt that I needed to sit down and reflect on the thousands of mishedges I
had committed.

I clambered up to my attic where, during 6 entire months, I spent 14
hours a day, 7 days a week, immersed in probability theory, numerical
analysis, and mathematical statistics (at a Ph.D. level). Then I began to
write this book.

Like George Soros’, I believed in a greater uncertainty principle (more
acute than Heisenberg’s) that largely invalidates social science theories
based on physics-like methodology and weakens the notion of modeling
outside of the natural sciences. It ain’t physics, I kept warning my trainees
throughout my career.

My other argument against being scientific was that, even if it were a
“science,” option theory (while perhaps on the right track) would be too
young to be reliable. I then needed to warn the public (and the regulators)
against taking an unseasoned and new field and applying some of its still
misspecified models to reality. Many of the market risks that have been
well known to traders since imperial Rome (like squeezes and the snow-
balling liquidity. holes) have not yet been rediscovered by the scientific risk
managers. I am convinced that the financial system is largely threatened
by the proliferation of risk management advisory services run by former
scientists who bullied their way into financial markets. My intention was to
downgrade hedging and risk management from the status of science to that
of a craft, until further notice.

This book is about hedging the risks of standard and exotic options, as
part of the larger framework of risk management. No road map was avail-
able since little has been written on this subject (in contrast to the extensive
literature for valuation).

Dynamic hedging is more like medicine than biology. It is learned by
gaining practical experience as well as by studying published research. The
wrinkles of the marketplace often dominate other complex issues, which can
lead option theoreticians onto a wrong path. Traders’ lore can only be trans-
mitted through practice. This book will meld matters of practical (not neces-
sarily anecdotal) importance with fundamental theory.

The major theme is to present traders and risk managers with the tools
to navigate around the difficult notion of manufacturing financial products
[End OCR]*

## Page 4

*[Image OCR]
vi Preface

through book-running. This book will introduce the arcane world of dy-
namic monitoring of risks. The core of dynamic hedging includes:

¢ The need for a methodology for the implementation of the Black-
Scholes-Merton? replicating process for options or any other nonlin-
ear security under the constraints imposed by the marketplace.

* The need to generalize the Black-Scholes-Merton framework to cover
other parameters than the underlying security in the replicating pro-
cess (like volatility or interest rates).

¢ The awareness that transaction costs and frequency can cause a de-
parture from the canons of continuous time finance.

e The awareness that distributions are unstable and hard to model.

Much of the common option literature has been concerned with details
of the pricing of instruments (some of which remain untractable).* These
works often provide insignificant answers to insignificant problems, such
as the search for precision in the pricing of American options with constant
volatility or interest rates (penny wisdom and dollar insensitivity). In addi-
tion, the nontheoretical option literature, departing from the Black-Scholes-
Merton framework, has been ensconced in static risk measurement. Most
documents introducing traders to conventional risks show only the static,
not the dynamic, risks. A derivatives ‘position that is dynamically hedged
will be subjected to an entirely different risk profile and, ‘given the limita-
tions of such hedging, will therefore be subjected to path dependence (a key
word for an option manufacturer).

Readers will not find a magazine-type proliferation of traded exotic
structures that delineate infinite variations and combinations; instead, the
analysis is limited to the nondecomposable structures (the SDF, smallest
decomposable fragment). A structure that is the addition of two products
will be therefore excluded (except in few cases of nonadditivity, where the
combination has some merit for the dynamic hedger). The objective of this
book is to provide the traders and risk managers with the tools, not the
ramifications.

Readers should use this book like a roadmap, searching out topics that
interest them and moving freely from topic to topic. The formal definitions
serve as anchors between categories.

More advanced mathematical topics are relegated to. the modules at
the back of the book. An attempt has been made to avoid mathematical
language and to explain issues in plain English. Formulas do not appear
until Chapter 22. In addition, in the presentation of mathematical ele-
ments, the book avoids the measure-theoretic framework (required for
most proofs in probability theory) and follows an intuitive path. Most
[End OCR]*

## Page 5

*[Image OCR]
Preface vii

mathematical concepts surrounding the topic can be explained with an in-
tuitive verbal description accompanied by graphical hints.

Option Wizards provide a lighter note for many serious topics. Since
these sections are designed to be read independently, readers can flip be-
tween them at their discretion.

Finally, throughout this book, the pronoun he is used as a stylistic con-
vention for ease of reading. This use should always be construed as gender
neutral.

e Part I (Chapters 1-6) defines market microstructure and products.
The markets are viewed from the vantage point of broker speaker
boxes and market pits, but also are defined in the formal setting of
market microstructure theory.

¢ Part II (Chapters 7-16) defines the basics of vanilla option risk and
presents measurement tools.

¢ Part III (Chapters 17-23) describes the risks of exotic options.

¢ Part IV (Modules A to G) presents more quantitative tools of analy-
sis and bridges a practitioner’s world with option theory. These
modules, however, should not be construed as appendixes: Most of
their content belong to the core of the text.

NOTES ON THE TEXT

Given that I did not initially learn about options in the literature but di-
rectly from the market (through observation and experimentation), most of
my reasoning remains highly intuitive. I apologize to people with more
scholastic tastes who may not be used to such a presentation. Most exam-
ples in this book are presented as generic situations. The volatility will be
defined as 15.7% (to make one standard deviation equal to 1% daily move).
The markets will be scaled to trade at 100.° For the purposes of pure option
situations, the forward is equal to spot, and the financial carry is insignifi-
cant (except in the rare difficult cases where it matters). All options will be
European style except for the exotic options where a term structure may
be introduced if it becomes relevant for the exercise.

Showing the profile of a butterfly, for example, will involve using
98/100/102 generic calls and studying a calendar spread looking at the
three-month 100 against the 6-month 100 calls, and so on.

The creation of generic examples will standardize all cases and help
in equating situations throughout the book. When dealing with a purely
conceptual option problem, it is necessary to strip out the underlying par-
ticularities. Optionality transcends the details in most cases. Where these
[End OCR]*

## Page 6

*[Image OCR]
viii Preface

particularities are essential, we will revert to a singular example taken from
a specific market.

Notations

P/L Profit and loss on a position

VorF Reserved for the value of a generic derivative security

K Strike price

S, or U, Underlying asset at time 0

H Outstrike (barrier strike)

H,,andH, High and low outstrikes for a double barrier

r Interest rate for the numeraire (also r,)

r,ord Rates for the counter-currency or the dividend payout for
a stock

y, Zero-coupon yield period t

t Time between present until expiration (except in cases
where ft, is the present, where it becomes t — t,)

O; Standard deviation of the natural log of the prices of the
asset i

Pi; Correlation of the natural log of the prices between
assets i and j

EJS, Conditional expectation at times 0 of the price of the
asset at times t

Jargon

Many terms may be linguistically ambiguous for people from outside the
industry. Even option books written for practitioners do not seem to pick
up our vernacular.

The same designation delta is used for both the rate and the total equiv-
alent exposure (rate times face value). The same for gamma, vega, theta, and
other Greeks.

By “volatility” is always meant implied volatility, not historical. By “15
volatility” read 15% implied volatility for the instrument in annualized
terms.

By “underlying” is meant underlying asset, which is also called spot,
cash price (as opposed to forward or future).
[End OCR]*

## Page 7

*[Image OCR]
Preface ix

By “50 cents price” for an option is always meant .5% of face value. By “1
dollar” is meant 1%. By “tick” or “pip” is meant .01% of face value.

By “long the 100” read long the 100 strike.

By “shorter” and “longer” option read “option with shorter time to ex-
piration.”

By “leg” is meant one side of a trade in a strategy.

By “Black-Scholes-Merton” is meant the Black-Scholes-Merton option
valuation model, as well as its extensions to more complex securities.

By “stopping time” or “first exit” read expected stopping time or ex-
pected first exit time.

By “high correlation matrix” read a correlation matrix with most para-
meters close to 1.

By “integral” is often meant stochastic integral. By “sensitivity to a para-
meter” is meant comparative static sensitivity to a parameter change. By “the
delta vanishes asymptotically” is meant that the delta vanishes asymptoti-
cally to the asset price.

By x/2 y is meant (x/2) y and a + b/2 means a + (b/2).

Finally, the term derivative, can mean either a derivative security or the
mathematical rate of change. When possible, the text will specify if it is a
mathematical derivative; otherwise it should be construed as being a security.

Nassim TALEB

Larchmont, New York
November 1996
[End OCR]*

## Page 8

*[Image OCR]
Contents

Introduction Dynamic Hedging 1

Principles of Real World Dynamic Hedging 1
General Risk Management 3

Part I

MARKETS, INSTRUMENTS, PEOPLE

1 Introduction to the Instruments 9

Derivatives 9
Synthetic Securities 12
Time-Dependent Linear Derivatives 13
Noncontingent Time-Dependent Nonlinear Derivatives 16
Options and Other Contingent Claims 16
Simple Options 18
Hard and Soft Optionality 20
Basic Rules of Options Equivalence 20
Mirror Image Rule 22
American Options, Early Exercise, and Other Headaches
(Advanced Topic) 24
Soft American Options 24
Hard American Options 25
A Brief Warning about Early Exercise Tests 27
Forwards, Futures, and Forward-Forwards
(Advanced Topic) 29
Credit 30
Marks-to-Market Differences 30
The Correlation between the Future and
the Financing (Advanced Issue) 31
Forward-Forward 32
Core Risk Management: Distinction between Primary and
Secondary Risks 32
Applying the Framework to Specific Instruments 35

xi
[End OCR]*

## Page 9

*[Image OCR]
xil

2

Contents

The Generalized Option 38

Step 1. The Homogeneity of the Structure 38

Step 2. The Type of Payoff: Continuous
and Discontinuous 41

Step 3. Barriers 43

Step 4. Dimension of the Structure and the Number
of Assets 43

Step 5. Order of the Options 45

Step 6. Path Dependence 46

Market Making and Market Using 48

Book Runners versus Price Takers 48
Commoditized and Nonstandard Products 50
Trading Risks in Commoditized Products 51
Profitability 53
Proprietary Departments 54
Tacit Rules in Market Making 56
Market Making and the Price for Immediacy 57
Market Making and Autocorrelation of Price Changes 58
Market Making and the Illusion of Profitability 58
Adverse Selection, Signaling, and the Risk Management
of Market Makers 60
Value Trading versus the Greater Fool Theory 62
Monkeys ona Typewriter 64
The Statistical Value of Track Records 64
More Modern Methods of Monitoring Traders 65
The Fair Dice and the Dubins-Savage Optimal
Strategy 65 ,
The ArcSine Law of the P/L 66

Liquidity and Liquidity Holes 68

Liquidity 68

Liquidity Holes 69

Liquidity and Risk Management 70

Stop Orders and the Path of Illiquidity 70
Barrier Options and the Liquidity Vacuum 72
One-Way Liquidity Traps 73

Holes, Black-Scholes, and the Ills of Memory 73
Limits and Market Failures 74

Reverse Slippage 74

Liquidity and Triple Witching Hour 75
[End OCR]*

## Page 10

*[Image OCR]
Contents xili

Portfolio Insurance 75
Liquidity and Option Pricing 77

Arbitrage and the Arbitrageurs 80

A Trader's Definition 80

Mechanical versus Behavioral Stability 81

The Deterministic Relationships 82

Passive Arbitrage 83

An Absorbing Barrier Called the “Squeeze” 84
Duration of the Arbitrage 84

Arbitrage and the Accounting Systems 85
Other Nonmarket Forms of Arbitrage 86
Arbitrage and the Variance of Returns 87

Volatility and Correlation 88

Calculating Historical Volatility and Correlation 92
Centering around the Mean 92

Introducing Filtering 95

There Is No Such Thing as Constant Volatility and
Correlation 97

The Parkinson Number and the Variance Ratio Method 101

Part Il
MEASURING OPTION RISKS

The Real World and the Black-Scholes-Merton
Assumptions 109 ’
Black-Scholes-Merton as an Almost Nonparametric

Pricing System 109

Adapting Black-Scholes-Merton: The Delta 115

Characteristics of a Delta 116
The Continuous Time Delta Is Not Always a Hedge Ratio 116
Delta as a Measure for Risk 121
Confusion: Delta by the Cash or by the Forward 123
Delta for Linear Instruments : 123
Delta for a Forward 123
Delta for a Forward-Forward 125
Delta fora Future 125
Delta and the Barrier Options 126
[End OCR]*

## Page 11

*[Image OCR]
xiv Contents

10

Delta and the Bucketing 127
Delta in the Value at Risk 127
Delta, Volatility, and Extreme Volatility 127

Gamma and Shadow Gamma 132

Simple Gamma _ 132
Gamma Imperfections fora Book 133
Correction for the Gamma of the Back Month 136
First Adjustment 137
Second Adjustment 138
Shadow Gamma _ 138
Shadow Gamma and the Skew 142
GARCH Gamma _ 142
Advanced Shadow Gamma _ 142
Case Study in Shadow Gamma: The Syldavian Elections 145

Vega and the Volatility Surface 147

Vega and Modified Vega 147
Vega and the Gamma _ 149
The Modified Vega 150 .
How to Compute the Simple Weightings 151
Advanced Method: The Covariance Bucket Vega 153
Forward Implied Volatilities 154
Computing Forward Implied Volatility 154
Multifactor Vega 158
Volatility Surface 164
The Method of Squares for Risk Management 164

Theta and Minor Greeks 167

Theta and the Modified Theta 167
Modifying the Theta 167
Theta fora Bet 169
Theta, Interest Carry, and Self-Financing Strategies 169
Shadow Theta 170
Weakness of the Theta Measure 171
Minor Greeks 171
Rho, Modified Rho 171
Omega (Option Duration) 174
Alpha 178
[End OCR]*

## Page 12

*[Image OCR]
11

12

13

Table of Greeks 181
Stealth and Health 182
Convexity, Modified Convexity 183
The “Double Bubble” 190

The Greeks and Their Behavior

The Bleed: Gamma and Delta Bleed (Holding
Volatility Constant) 191
Bleed with Changes in Volatility 195
Going into the Expiration of a Vanilla Option 196
Ddeltadvol (Stability Ratio) 200
Test 1 of Stability 200
Test 2 of Stability: The Asymptotic Vega Test 201
Moments of an Option Position 202
Ignoring Higher Greeks: The Lock Delta 204

Fungibility, Convergence, and Stacking

Fungibility 208
Ranking of Fungibility 209
Fungibility and the Term Structure of Prices:
The Cash-and-Carry Line | 210
Fungibility and Option Arbitrage 212
Changes in the Rules of the Game 212
Convergence 213
‘Mapping Convergence 215
Convergence and Convexity 216
Levels of Convergence Trading 216
Volatility and Convergence 216
Convergence and Biased Assets 216
Stacking Techniques 217
Other Stacking Applications 220

Some Wrinkles of Option Markets

Expiration Pin Risks 222

Sticky Strikes 223

Market Barriers 224
A Currency Band: Is It a Barrier? 225
The Absent Barrier 226

What Flat Means 226
Primary and Secondary Exposures 228

Contents xv

191

208

222
[End OCR]*

## Page 13

*[Image OCR]
xvi

14

15

16

17

Contents

Bucketing and Topography 229

Static Straight Bucketing 229
‘American and Path-Dependent Options 231
Advanced Topic: The Forward or “Forward-Forward”

Bucket 231

Topography 232
Strike Topography (or Static Topography) 233
Dynamic Topography (Local Volatility Exposure) 235
Barrier Payoff Topography 237

Beware the Distribution 238

The Tails 238

Random Volatility 238

Histograms from the Markets 242
The Skew and Biased Assets 245

Biased Assets 248

Nonparallel Accounting 249

Value Linked to Price 250

Currencies as Assets 250

Reverse Assets 251

Volatility Regimes 251

Correlation between Interest Rates and Carry 252
More Advanced Put-Call Parity Rules 252

Option Trading Concepts 256

Initiation to Volatility Trading: Vega versus Gamma 260
Soft versus Hard Deltas 262 :
Volatility Betting 263

Higher Moment Bets 264
Case Study: Path Dependence of a Regular Option 265
Simple Case Study: The “Worst Case” Scenario 270

PART III

TRADING AND HEDGING Exotic OPTIONS

Binary Options: European Style 273

European Binary Options 273
Hedging witha Vanilla 275
Definition of the Bet: Forward and Spot Bets 278
[End OCR]*

## Page 14

*[Image OCR]
18

19

Contents xvii

Pricing with the Skew 279
A Formal Pricing on the Skew 281
The Skew Paradox 282
Difference between the Binary and the Delta: The Delta
Paradox Revisited 284
First Hedging Consequences 286
The Delta Is a Dirac Delta 286
Gamma for a Bet 287
Conclusion: Statistical Trading versus Dynamic Hedging 289
Case Study in Binary Packages—Contingent Premium
Options 290
Recommended Use: Potential Devaluations 291
Case Study: The Betspreads 292
Advanced Case Study: Multiasset Bets 294

Binary Options: American Style 295

American Single Binary Options 295
Hedging an American Binary: Fooled by the Greeks 298
Case Study: National Vega Bank 298
The Ravages of Time 299
Understanding the Vega Convexity 303
Trading Methods 305 .
Case Study: At-Settlement American Binary Options 306
Other Greeks 307
American Double Binary Options 307
Vegas of the Double Binary 308
Other Applications of American Barriers 309
Credit Risk 311

Barrier Options (I) 312

Barrier Options (Regular) 312
Knock-Out Options 312
Knock-In Options 317
Effects of Volatility 319
Adding the Drift: Complexity of the Forward Line 321
Risk Reversals 323
Put/Call Symmetry and the Hedging of Barrier
Options 323
Barrier Decomposition under Skew Environments 331
The Reflection Principle 335
Girsanov 339
Effect of Time on Knock-Out Options 339
[End OCR]*

## Page 15

*[Image OCR]
xviii Contents

20

21

22

First Exit Time and Its Risk-Neutral Expectation 340
Issues in Pricing Barrier Options 343
The Single Volatility Fudge 343
A More Accurate Method: The Dupire-Derman-Kani
Technique 344
Additional Pricing Complexity: The Variance Ratios 345
Exercise: Adding the Puts 346

Barrier Options (II) 347

Reverse Barrier Options 347
Reverse Knock-Out Options 347
Case Study: The Knock-Out Box 348
Hedging Reverse Knock-Outs: A Graphical

Case Study 356

Double Barrier Options 362
Rebate 363
Exercise: Adding the Knock-In 363
Alternative Barrier Options 363
The Exploding Option 364
Capped Index Option 365

Reading a Risk Management Report 368
Gaps and Gap Reports 374

Compound, Choosers, and Higher Order Options 376

Vega Convexity: The Costs of Dynamic Hedging 378
Uses of Compound Options: Hedging Barrier Vega 379
Chooser Options 380

A Few Applications of the Higher Order Options 382

ee

Multiasset Options 383

Choice between Assets: Rainbow Options 384
Correlated and Uncorrelated Greeks 387
Linear Combinations 390
Basket Options 391
Lognormality 391
Correlation Issues 392
Composite Underlying Securities 395
Quantitative Case Study: Indexed Notes 395
Background 396
Terms of the Note 396
Where Is the Underlying? 397
Triangular Decomposition 398
[End OCR]*

## Page 16

*[Image OCR]
Contents xix

23 Minor Exotics: Lookback and Asian Options 403

Lookback and Ladder Options 403
The Rollover Option 404
A Footnote on Basket Options: Asian Options 408

Parr IV

MOoDULES

Module A Brownian Motion on a Spreadsheet, a Tutorial 415

The Classical One-Asset Random Walk 415

Some Questions 417

A Two-Asset Random Walk: An Introduction to the
Effects of Correlation 420

Extension: A Three-Asset Random Walk 424

Module B_ Risk Neutrality Explained 426

Step 1. Probabilistic Fairness, the “Fair Dice” and
the Skew 426

Step 2. Adding the Real World: The Risk-Neutral
Argument 427
The Drift 427

Module C Numeraire Relativity and the Two-Country Paradox 431

Extension: The Two-Country Paradox 433

Conclusion 435 ~~

Mathematical Note 436
Conclusion 437

Module D Correlation Triangles: A Graphical Case Study 438

Correlation Triangle Rule 441
Calculating an Implied Correlation Curve 444

Module E The Value-at-Risk 445

Simplified Examples 446
Example 1. No Diversification 447
Example 2. A Cross-Position 447
Example 3. Two Possible Trades 448
[End OCR]*

## Page 17

*[Image OCR]
xx Contents

Module F_ Probabilistic Rankings in Arbitrage 453

Ranking of Securities 453
European Option Rules 453
Calendar Rules 454
Barrier and Digital Rules 454
Correlation Rules 455

Correlation Convexity Rules 457

General Convexity Rules 458

Module G Option Pricing 459

Ito’s Lemma Explained 459
Ito’s Lemma for Two Assets 462
Black-Scholes Equation 463
The Risk-Neutral Argument 463
Stochastic Volatility Model 464
Multiasset Options 466
Rainbow Options 466
Outperformance Options 467
Spread Options 467
Compound and Chooser Order Options 467
Compound Options 468
Chooser Options 468
Barrier Options 468
The Reflection Principle 469
' Girsanov’s Theorem 469
Pricing Barriers 470
Numerical Stochastic Integration: A Sample 477
A Mathematica™ Program 477

Notes 479
Bibliography 490

Index 49
[End OCR]*

## Page 18

*[Image OCR]
Introduction

Dynamic Hedging

As these events are beyond our understanding let us fake being their instigator.
Jean Cocteau

PRINCIPLES OF REAL WorRLD DYNAMIC HEDGING

@ Rebalancing the gamma corresponds to buying and selling the underly-
ing security in order to replicate the payoff of the option.

Even if traders knew the exact future volatility but hedged themselves (re-
balanced the gamma) at discretely spaced increments, they would still have
difficulty predicting the final P/L. Option pricing eliminated, by necessity,
the transaction costs. If however, they traded every millionth of a second at
the screen mid-market they would get a P/L with certainty.

Increasing the frequency of adjustments would compress the results as
shown in the following chart:

Probability
density

Replicating price

Mathematically, the shape of the distribution is simply determined by
o/Vn (functional central limit theorem). An intuitive explanation is that
an average will tend to the mean at a speed of 1/V number of draws. In this
case, the mean is the Black-Scholes-Merton value of the security. In Chap-
ter 16, a case study shows the magnitude of the adjustment luck in the
tracking of the P/L.
[End OCR]*

## Page 19

*[Image OCR]
2 Introduction

The increase in transaction frequency would have the effect of increas-
ing the dynamic hedging costs (owing to transaction costs) by moving the
center of the distribution to the left, as seen in the following chart:

Market makers hence face the continuous dilemma between:

¢ Variance of returns on one hand. (Option replication is not a risk-free
proposition.)
e Transaction costs on the other.

When an option trader sells a structure that is worth 5.00 on his theo-
retical value sheets, he can expect, baring transaction costs, to manufacture
it for somewhere between 4.00 and 6.00. However, the more parameters he
needs to be protected against, the higher the management costs.

Risk Management Rule: The more volatile parameters the option
trader needs to be protected against, the more allowance is neces-

sary against both the volatility of the parameters and the expected
transaction costs.

The thesis of this book is that the more parameters fly around (interest
rates, volatility term structures, etc), the more arduous the dynamic hedg-
ing. Unlike the Black-Scholes-Merton world our trading environment re-
quires us to hedge more than the gamma. Every possible second derivative
will be costly. We must hedge the vega convexity, the exposures to the rates
and so on as they are not stable. Chapter 15 shows the effect of stochastic
volatility and provides the intuition for the vega convexity. Chapter 10 in-
cludes a description of convexity broadly defined.

As Black-Scholes-Merton includes the cost of gamma in computing the
time value of the option, so will we need to add the convexity of all
non linear parameters to it (an Ito term, see Module G). A position that is
short gamma on interest rates in addition to the Black-Scholes-Merton
gamma needs to be priced accordingly, with the proper markup (or down).
The Option Wizard, The Contamination (or Convexity) Principle, provides an
[End OCR]*

## Page 20

*[Image OCR]
Dynantic Hedging 3

intuitive explanation of the issue. Problems will occur where the moves in
the parameters might be correlated.

Positions that require vega neutrality in a concave vega (short volatility
of volatility) will then be inferior in value to others that do not require such
dynamic hedge owing to linear vega and so on.

The need to understand the replicating process more thoroughly with
exotic options reflects their greater hedging costs.

GENERAL RISK MANAGEMENT

In the immature modern financial jargon, risk management is usually inter-
preted as either coping with the financial risks inherent to an otherwise
nonfinancial businesses or the market risks incurred by providers of finan-
cial instruments (such as trading firms, exchange traders, banks, and other
dealers). The latter will be the concern of this book. Books on risk manage-
ment abound by people who try to explain sophisticated financial tools to
treasurers of corporations or pension funds managers. Because providers
face daily risks in the manufacturing of the financial products and instru-
ments, this book focuses on the risk transfer between a static product (cor-
porate risk) and a dynamic one (the market maker’s), the delta-neutral
operator in Black-Scholes-Merton’s world.

There are two levels of risk management: micromanagement and macro-
management. ,

Micromanagement occurs at the level of one product line or one coher-
ent book run by a unit (e.g., exchange local or “upstairs” trading desk).
This encompasses the intimate knowledge of the behavior of every de-
rivative product with respect to time and market movement and the
thorough thinking in multiple dimensions required by the derivatives
trader’s function (the fruit of lengthy apprenticeship).

Macromanagement takes place at the level of a general firm. This is a
generally more quantitative, more theoretical function perfected by a
general watchdog. Somehow by magic, the sum of the risks does not add
up to the total risks. Some books match one another, while the risks of
others become additive. However, the measurement tools, like diversifi-
cation and correlation, so celebrated in modern finance appear to be
still imprecise and of weak predictive powers. Their use can danger-
ously lull the risk managers into a false sense of statistical security.

The macromanager’s function is to allocate risk (defined as total possi-
ble variance) across products and units to maximize the utilization of a
trading firm’s most valuable resource: risk. Figure J.1 illustrates the rela-
tionship of dynamic hedging to risk management.
[End OCR]*

## Page 21

*[Image OCR]
4 Introduction

Global Risk Management

DYNAMIC HEDGING
Macro Market Risks
Micro Market Risks

Primary Market Secondary Market Residual Risks
Risks Risks

Fraud Risks Credit Risks Technological Operational Legal Risks
Risks Risks

Figure I.1 The larger risk management framework.

There has been a growth in the number of “risk management advisors,”
an industry sometimes populated by people with an amateurish knowledge
of risk. Using some form of shallow technical skills, these advisors emit
pronouncements on such matters as “risk management” without a true un-
derstanding of the distribution. Such inexperience and weakness become
most apparent with the value-at-risk fad or the outpouring of books on risk
management by authors who never traded a contract.

There is no shortcut method of risk management. Every person who
needs to be acquainted with the activity should work hard at trying to
make sense of the dynamic interaction of portfolio components relative to
time and combinations of market movements. Every derivative user or
trader needs to be able to account for the effect of either the passage of
time or the movement of the underlying assets on the portfolio. Issues
of pricing are therefore overstated: Most money is made or lost because of
market movement, not because of mispricing. Often the cause is mishedg-
ing. Most commonly, losses result from a poor understanding of liquidity
and the shape of the statistical distribution.
[End OCR]*

## Page 22

*[Image OCR]
Dynamic Hedging 5

Option Wizard: Culture Shocks

There are spectacular communication problems between traders and quants
(research assistants) owing to many factors:

* People who have completed a Ph.D. program develop a tendency to be
thorough.

¢ Traders, by contrast, are impatient and need brief simplistic descriptions,
something someone who spent most of his adult life solving complex prob-
lems is often unable to provide. Traders tend to see matters in “flashes,” a
type of nonlinear epistemology. Such desire for the highest level of possible
abstraction is sometimes mistaken for “short attention span.”

This leads to Brecht-like amusing exercises in incommunicability between
traders and their quants. One trader (himself a former quant) built the following
strategy in dealing with the research department people. When communicat-
ing with him, they were advised to do the following:

1. In a single sentence, explain the conclusion, before discussing the subject
matter.

2. Ina single sentence, explain the subject matter.

3. If unable to perform 1 or 2, then abandon the entire project.

This method helps the quant understand the major difference between the
less challenging matters of the real world and the thrills of science. In addition,
it provides him with a way to learn to focus like a businessperson.

Richard Feynman,* one of the greatest minds of our times, was comfort-
able enough with his subject matter to write an intuitive book on quantum
physics without a single formula.

* Richard Feynman, QED (1985) Princeton: Princeton University Press.
[End OCR]*

## Page 23

*[Image OCR]
PART I

MARKETS, INSTRUMENTS, PEOPLE
[End OCR]*

## Page 24

*[Image OCR]
Chapter 1

Introduction to the Instruments

Understanding a theory means (...) understanding it as an attempt to solve a
certain problem.
Sir Karl Popper

In this chapter, we will rapidly but formally define the instruments and
present their major characteristics. It is recommended that all readers, even
those knowledgeable in this area, study the following definitions, as they
will provide the framework of analysis used in the book.

DERIVATIVES

@ A derivative is a security whose price ultimately depends on that of an-
other asset (called underlying). There are different categories of deriva-
tives, ranging from something as simple as a future to something as
complex as an exotic option, with all the shades in between.

The best way to look at derivatives is to separate them into two broad
categories: linear and nonlinear derivatives. A linear derivative is easy to
hedge and lock in completely, whereas a nonlinear one will present serious
instability and require dynamic hedging.

@ A nonlinear derivative with respect to a parameter is one that presents
a second derivative (or partial derivative with respect to that parameter)
different from 0.

The option wizard presents a graphic linear or nonlinear derivative view of
the concept of nonlinearity.

Risk Management Rule: All nonlinear derivatives are time-

dependent in their price.

This rule is explored in the contamination principle and will be dis-
cussed throughout this book. For now, it is enough to state that nonlinearity
[End OCR]*

## Page 25

*[Image OCR]
10 Markets, Instruments, People

Option Wizard: The Greeks

The “Greeks,” as option traders call them, denote the sensitivity of the option
price with respect to several parameters. The following are basic definitions for
use in Part |. These terms will be explored in greater detail later in this book.

Delta Sensitivity of the option price to the change in the underlying asset
price.

Gamma _ Sensitivity of the option delta to the change in the underlying
asset price.

Vega Sensitivity of the option price to the change in implied volatility.

Theta Expected change in the option price with the passage of time as-
suming risk-neutral growth in the asset.

Rho = Sensitivity of the option price to interest rates or dividend payout.

“Long gamma or vega” means a positive sensitivity to the Greek (a higher
P/L at a higher parameter).

General Assets

Time dependent Nonlinear
Noncontingent Contingent
Derivatives Claims

(Mostly) Linear
Synthetic

Securities

Linear Options
Baskets Spreads | | or “Quasi-Linear” Nonlinear Categories

Swaps (fixed Squared
income,commod- forward payoffs
ities, etc.) , Cubed forward
Currency payoffs

Payoff is the
(weighted)

Payoff is the
difference
between
securities

average of
securities

Forwards
FRA
Eurodollars

Figure 1.1 Classification of derivatives.
[End OCR]*

## Page 26

*[Image OCR]
Introduction to the Instruments 11

Option Wizard: The Hedger’s Viewpoint

Throughout this book, think of a derivative in terms of its replication costs. For
this purpose, the world is divided into two categories: the user and the manu-
facturer. Their utility function and even results will be entirely different. The
user is involved in terminal value (usually), while the manufacturer engages in
dynamic hedging (when he is doing his job right), which markedly alters the
product.

A dynamic hedger will not be interested in whether he owns a put or a call
(first-order hedges will make them identical). What matters is the strike and
the expiration.

is gamma (or more generally called convexity) and that gamma needs to be
accompanied by time decay (the “rent”).

Derivatives are not always linear, convex, or concave across all moves
(see Figures 1.2A—D). A test of local linearity of a derivative security (that
is a function of the underlying asset) between asset prices S, and S,, with
0<A<1, will satisfy the following equality :

V(AS, + (1 — A)S,) = A V(S,) + 1 — A)VIS,)
It will be convex between S, and S, if:

V(AS, + (1 — A)S,) SA V(S,) + A — ANV(S,)
It will be concave if:

VAS, + (1 — A)S,) = A V(S,) + (1 — AYV(S,)

Option Wizard: Linear and Nonlinear Securities

' Although we are initially considering linearity with respect to the underlying
asset, this notion will be later extended to other parameters such as interest
rates and volatility.

As the graphical representation at the top of page 12 shows, a linear secu-
rity constantly behaves like a line. In option parlance, it will have a delta but
no other Greeks,* and certainly no curvature. Linear securities require little or
no dynamic hedging. .

*The Greeks initially represented the various derivatives of the Black-Scholes-Merton formula.
By extension, it became any sensitivity of a derivative security with respect to a particular mar-
ket parameter.
[End OCR]*

## Page 27

*[Image OCR]
12. Markets, Instruments, People

A Linear Product
shows literally a

line (or close to a
line)

Derivative
Price

Underlying Asset

Figure 1.2A Linearity.

Derivative
Price

Underlying Asset

Figure 1.2C A nonlinear derivative:
Concave security.

Derivative
Price

Underlying Asset

Figure 1.2B A nonlinear derivative:
Convex security.

Derivative
Price

Underlying Asset

Figure 1.2D A mixed nonlinear
derivative. ,

Many securities exhibit some linearity until the test of fire. These in-
struments are called “quasi-linear” securities. Convexity affects many fi-
nancial instruments, even those least suspected.

The contamination principle, we will see, dictates that every nonlinear
security commands time value, positive if the security is convex and nega-
tive if it is concave.

SYNTHETIC SECURITIES

@ A synthetic security is a linear combination of two or more primary in-
struments in the markets.

A basket’s price is derived from a weighted combination of existing primary
instruments. For example, Standard & Poor’s 500 (SP500) contract is a
weighted average of the components. It can therefore be exactly replicated
with a mixture of the components (for those who have the time and patience
to leg the 500 stocks). A European currency unit (ECU) is another example
[End OCR]*

## Page 28

*[Image OCR]
Introduction to the Instruments 13

Option Wizard: The First Derivative Trade

The earliest option trade on record in Western literature was a bet on future
crop by Thales of Miletus, which Aristotle recounted with great pride in his
Politics.* To benefit from a better than expected olive crop, Thales put a de-
posit on every olive press in the vicinity of Miletus. As demand for these grew,
he sublet the facilities for profit, mostly to make the point that philosophers
who so desire can achieve material success. The dichotomy between the
“MIT-smart” and the “Brooklyn-smart,” today prevalent on Wall Street, was al-
ready apparent in fifth-century Asia Minor. Thales used the first derivative in-
strument, actually an option on a future, a second-order derivative at that! He
did not trade olives, which he would have had to sell short, but chose to buy
the equivalent of a call on the olive presses, for fall delivery, with the knowl-
edge that all he could lose was his deposit.

*See Russell (1945).

of an arithmetic average. A basket composed of 20% stock A and 80% V
stock B can easily be replicated with the purchase of every component.

Synthetic securities are not always linear, but exceptions are rare enough
for us not to bother with them. For example, when the average of the instru-
ments is not arithmetic, some oddities can result. The U.S. dollar index,
traded on FINEX will present some convexity owing to the geometric nature
of the averaging and will therefore trade at a premium to the underlying se-
curities for that reason. The nonlinearity that results from the convexity can
cause a neutral person to benefit from a market move either way, which can
make one side of the arbitrage more desirable. Cohvexity, which will be de-
fined later on, usually commands a price since Wall Street rarely grants free
lunches.

Example: The Elevator Bank issues its own “mother of all baskets” and
emits some notes, the payoff of which is indexed off the basket. The offi-
cial reason is that the basket effectively tracks inflation or some other
indicator. The true reason is that the basket will have a lower volatility
than the sum of the instruments and is believed to be easy to hedge.

TIME-DEPENDENT LINEAR DERIVATIVES

@ Time-dependent linear derivatives are instruments separated from the
original asset through time.
[End OCR]*

## Page 29

*[Image OCR]
14 Markets, Instruments, People

Option Wizard: The Contamination (or Convexity) Principle

The most important notion in option hedging and trading is the contamination
principle: It is the fundamental principle of dynamic hedging. It means roughly
that if there is a possible spot in time and space capable of bringing a profit,
then the areas surrounding it need to account for that effect.

The contamination principle is similar to the notion of heat transfer.* If a
spot is located near a source of heat, then its temperature will rise accordingly.
If the asset price nears a level that would bring a sizable profit to the portfolio
then the area surrounding it should cause a modicum of profit as well.

in the following chart, a derivative will pay $1 million should some event
take place in the market. Starting at a given point, the security pays on the
node to the right. It would then be unreasonable to think that such a security
would be worth nothing at Point So. One standard deviation move should result
in a payoff of $1,000,000 (or part of that sum).

The Contamination Principle

Derivative

Payoff 1 St Dev

Payoff
$1 million

Payoff

So Asset Price

Every experienced trader, seeing the payoff and the probability attached to
it, would buy the derivative security. The derivative would then be worth more
than 0, and it is easy to see time value taking shape in the following chart:

Derivative

Payoff 1 St Dev Payoff

<q si niillion

Possible
Time Decay

Payoff

So Asset Price
However, time value means time decay. As shown in the following chart,

with everything remaining the same, the option would decrease in time value

at the standard deviations as the payoff becomes less likely:

(Continued)
[End OCR]*

## Page 30

*[Image OCR]
Introduction to the Instruments 15

1 St Dev

Payoff

<b
Payoff = $.4 million
0

The final chart shows that when the points are far from each other (as in
option prices), a convex line forms.

Convexity

Asset Price

This example suggests why an option has “convexity.”

* The origin of the designation contamination resides in the price of an option “contaminating”
those around it. When an option becomes expensive, those around it need to follow. Likewise,
one can see that if an Arrow-Debreu price rises those around it need to increase as well. They
are then said to contaminate each other.

A swap is a linear (or quasi-linear) derivative, as the second-order derivative
with respect to the asset price is equal (or close to) zero. In trader’s par-
lance, its hedge ratio is not supposed to change relative to movements in the
underlying asset, although it can be argued that there is no such thing as a
purely linear derivative.

Time-dependent linear derivatives include:

¢ Forwards. They are agreements to swap some proceeds in the future.

° Floating Rate Agreements (FRAs), Eurodollars. They will, for the
purpose of this book, be forward-forwards that can be broken down
into strips of products that start period t and end period ¢ + 1.

¢ Swaps. Whatever their end use (this is not of any concern here),
they can be composed into a combination of Eurodollars or FRAs.
Most of their complexity comes from being detail-heavy, but they
otherwise exhibit well-behaved features.'

Aside from the correlation between a future price and its financing,
most of the difficulties of these assets lie in the problem of interpolation
[End OCR]*

## Page 31

*[Image OCR]
16 Markets, Instruments, People

between two points. The way time treats them is not always easy to ascer-
tain. A lesser, but no small, difficulty resides in the multiplication of
minute details such as the conventions on the delivery or the 360/365-day
basis and other rules that no trader was ever known to memorize.

NONCONTINGENT TIME-DEPENDENT
NONLINEAR DERIVATIVES

@ Noncontingent time-dependent nonlinear derivatives represent instru-
ments that are convex, concave, or mixed (with respect to the underly-
ing asset) but that are not options (i-e., noncontingent).

The infamous LIBOR-square (London Interbank Offer Rate) for example,
provides a lurid example of such convexity. The LIBOR cube will have a
third derivative (a convexity of the curvature) but the product does not ap-
pear very likely to sell heavily. Despite their payoffs, however, these instru-
ments do not constitute an option since both parties are obligated to swap
the proceeds. The strangeness of their payoff is that they are generally con-
vex above a point and concave below, or vice versa. The acceleration of the
positive payoffs on one side is balanced by the acceleration of the negative
payoffs on the other. .
Example: The Elevator Bank sells to its customers in the area sur-
rounding Cincinnati, Ohio, a note paying to its holder the square of the
interest rate move (between inception and some predetermined time in
the future). The customer, stressed out by low interest rates, would thus
be compensated in an accelerated way against further rate drops. The
note looks like an option, and being (on the surface) arduous to hedge, it
will sell for higher than “fair” value.

OPTIONS AND OTHER CONTINGENT CLAIMS

The price of options depends on contingent events. They represent the bulk of
this study. They are the culprit, the topic of this book..A swap is a linear de-
rivative while a path-dependent swap will present uncertain payoffs. In the
past, operators defined these instruments as tools with “optionality,” where
one party had a right to choose and the other one was under an obligation.

We will study options at two levels. Level 1 is a conventional presenta-
tion of the product, as if the exotic option markets did not exist. Level 2 will
go into a more generalized presentation of the option markets that would
encompass both vanilla and exotics (Chapter 2).
[End OCR]*

## Page 32

*[Image OCR]
Introduction to the Instruments 17

Option Wizard: The Contamination Principle and LIBOR Square

LIBOR square is a contract with a mixed payoff: convex in ore zone and con-

cave in another. At the origin (where the contract is set), it is neither. According

to the contamination principle, the contract needs to be higher than its value on

the line in areas where it is long gamma and lower in areas where it is short.*
The contract pays:

q(x — x,)? if x > x,
—q(x — x,)? if x < x,
with q the quantity, x, the origin, x the present LIBOR price.
Therefore, its delta is
2q(x — x,) if x > x,
—2q(x — x,) if x < x,
And the gamma is
2q if x > x,

—2q if x< x,

The following chart illustrates the valuation process:

Likely value

LIBOR square
(contract value) gamma

LIBOR

The discussion of option trading will clarify the concept of cells contami-
nation: The arbitrage derived value of any security will depend on its delta-
neutral replication.

*LIBOR square is a contract that pays to one party the square of the difference between the
origin and a higher price and obligates it to pay the square of the difference between the lower
price and the origin.
[End OCR]*

## Page 33

*[Image OCR]
18 Markets, Instruments, People

Simple Options

Options are contingent claims, and thus distinguish themselves from other
products in being a potential asset for one party and a potential liability for
another. This contingency in their value subjects them to probability the-
ory. All option pricing consists therefore in dealing with probability.

@ A put is the right to sell an instrument at a certain price (the strike
price) within some time frame. A call is the right to buy the instrument.

With the opening of many new instruments, puts or calls can be
confused. They depend on the numeraire. For a currency pair, a put
on Mark/Dollar (the right to sell Marks and buy dollars) is a call on
Dollar/ Mark. Similarly, a put on yields is a call on bonds, a matter of some
confusing importance since one illustrious exchange defined a contract on
yield (thus inventing an asset), while most cash instruments are traded ac-
cording to their price.

An interesting extension is that a call on the SP500 is a put on cash for
someone whose P/L is computed in SP500 units (the case of indexed fund
managers). More on that later.

@ A European option can only be exercised on the last day. An American
option can be exercised any time between its inception and the end
date. A hybrid, the Bermudan option can be exercised on a set number
of days between inception and expiration.

M Intrinsic value, for a call, or the in-the-money part of an option is the
difference between the asset value and the strike price if that difference
is positive. It is zero if the difference is negative. For a put, it is the re-
verse. For a European option (an option that is only exercisable on one
date), the intrinsic value is typically expressed by traders (by conven- _
tion) as the difference between the strike and the corresponding for-
ward. Since the option is not exercised before expiration, the only price
that matters is the term market price of the asset (for delivery on the ex-
piration of the option).

The best way intuitively to test whether any particular instrument is an
option is to see if the payoff is asymmetric and if there is a strike price (Fig-
ure 1.3). A call option is priced at expiration as the

Max (S — K, 0)

with S the asset price at expiration and K the strike price. It is read as the
greater of the difference between the asset price at expiration and the strike or zero.
[End OCR]*

## Page 34

*[Image OCR]
Introduction to the Instruments 19

This simple formula means in English that the owner collects some
amount (intrinsic) or nothing, whichever is greatest. 5 — K is the difference
between spot and strike. When it is negative, the operator would prefer to
receive nothing.

A put option will be expressed as

Max (K — S, 0)

@ Forwards and futures are contracts to unconditionally exchange an
asset at a predetermined date for an agreed-on price.

They are straight claims, with assets and liabilities on both sides of the
fence (as opposed to options where one party has an asset and the other a li-
ability). They also distinguish themselves from options because the payoff
does not give any party the element of choice. As will be explained later,
some minor technicalities in the definition of the futures contracts (truly,
very minor) such as the bond futures* spawned an entire cottage industry of
arbitrageurs, analysts, and the like.

Call Value at Expiration

Asset Price

Put Value at Expiration

Asset Price

By the contamination principle...
Time value of a call

Asset Price

Figure 1.3 Option value.
[End OCR]*

## Page 35

*[Image OCR]
20 Markets, Instruments, People

Another intuitive way to understand an option is to think of it as half
the forward or future: The long partakes of the upside (above a certain
point, called the strike price) and the short partakes of the downside, for a
fee. It is no wonder that being long a put and short a call replicates entirely
the future (of course, there are questions of early exercise that we will ig-
nore initially).

Options proliferate in daily life, some of which people are long (they
own the choice) for some of which people are short. A chief executive officer
owns, by virtue of his position, an option. The manager can partake of the
upside of the company and get a bonus when the performance is acceptable.
Should the industry experience a downturn and the company go bankrupt,
the CEO’s sole risk would be his job or some unpleasant but otherwise fi-
nancially harmless castigation. He will not be asked to appear before the
board with his checkbook in hand. It is easy to see that the strike price for
the bonus is the required performance and that the shareholders sold (or,
rather, gave) that option to the manager.

Forwards, options, and futures represent the bulk of derivatives instru-
ments. A later section in this chapter will present a more advanced version
of options.

HARD AND SOFT OPTIONALITY
HH Optionality is a broad term used by traders to describe a nonlinearity
in the payoff of an instrument. It is often applied to convex instruments
or to situations like a “stop loss” or a known order in the market.

As an extension of the contamination principle, every item with op-
tionality needs to trade at a premium and the shrinkage of premium
with the passage of time (owing to the narrowing of the probability
measure of events) will mean necessarily that every item with optional-
ity will have time decay.

It is convenient to call hard optionality the situation where a contract has a
strike price and soft optionality the situation where the contract has a built-
in convexity but no real strike price. Soft optionality presents generally
milder gamma and other Greeks but will present more stable features
across time.

Basic RULES OF OPTIONS EQUIVALENCE

Below are the basic rules of what traders call “option algebra.”

¢ Put-call parity for a European option: Long call/short put = Long
forward, provided they are all of the same strike.
[End OCR]*

## Page 36

*[Image OCR]
Introduction to the Instruments 21

Warning Expiration: “Pin Risks” (described in Chapter 13) cause put-
call parity not to hold for listed options.

Replacing a long with a “+” sign and the short with a “ — ” sign allows
for the following simple arithmetic:
Position equivalence:

+C — P =F. A long forward is equal to a long call/short put of the
same strike.

Hence:

+C = P + F. Long a call is equivalent to long a put/long a future.
And

+P =C — F. Long a put is equivalent to long a call/short a future.

¢ For a “soft” American option (see definition later in this chapter),
the put-call parity rules hold but with a weaker equivalence.

¢ For “hard” American options, the rule becomes more slippery. More
complex rules are described in Chapter 15. It is recommended for the
nonspecialist to completely ignore put-call parity.

Example: Assume that the 3-month 102 Put trades for $1.975, that the
3-month 102 Call trades for $2.9625, and that the forward for the exact
delivery date for both trades for 101.00.

At expiration, assuming financing at 5%, the call will cost
$1.975 < .05 X (90/360) = .025 and the put will cost .0375. We will
then have the P/L shown in Table 1.1. .

Table 1.1 Static Put/Cali Parity

Forward Total Forward
Asset Price 102CallP/L ContractP/L 102PutP/L P/L+102 Put P/L

106 2 5 -3 2
105 1 4 -3 1
104 0 3 —3 0
103 ~1 2 -3 —1
102 —2 1 -3 —2
101 -2 0 —2 —2
100 —2 -1 -1 —2

99 —2 —2 0 —2

98 —2 —3 1 —2

97 —2 —4 2 4
[End OCR]*

## Page 37

*[Image OCR]
22. Markets, Instruments, People

Table 1.1 shows the profile at expiration. If two trades are identical at ex-
piration everywhere on the map of possible prices, and if they expire on the
same day, then the trades will have the same risk and profit/loss profile
during their life.

The following rule applies to all markets, properly rescaled, provided
they are European and present liquid forwards. To make that rule accept-
able for the future, proper allowance needs to be made for the “tailing” (see
Chapter 7).

Mirror Image Rule

1 unit of a put + x% of the unit in forward = 1 unit of a call + (100 — x)%
of the unit in forward, all of the same strike and expiration.

This rule is obvious: A put delta neutral at 30% is equal to P + 3F =
(C — F) + .3F = C — .7F. Ifa 103 put has a forward delta of 30% the 103 call
will have a forward delta of 70%. This formula uses forward delta, not cash
deltas that need to be adjusted for. Most risk management systems disclose
the cash, not the forward delta, as does the canonical Black-Scholes-Merton
formula. ,

Warning: European options need to be hedged with the forward,‘ not
cash. However, most commercial pricing systems tend to disclose the
spot hedge instead, which often can be misleading. Traders often have
recourse to cash for short-term hedges owing to the lack of liquidity in
the forwards. This habit generally leads them to forget the exact condi-
tions for adequate put/call parity.

@ A forward delta for a European option is the equivalent cash position
with the same delivery date as the underlying asset.

For an American option, a forward delta is typically of uncertain dura-
tion. Such a duration, however, is generally calculated and called the
“omega” (discussed later in this chapter), but will be too unstable for us to
use for adequate equivalence.

Consequently, a straddle will be equal to two calls delta neutral or two
puts delta neutral (of the same strike). Assume that the forward delta of a
put is 30%:

Straddle = 2P + .6F = 2(C — F) + .6F = 2C — 2F + .6F = 2C — 14F

Consequently, a call calendar spread will have the same profile as a put
calendar spread (assuming interest rates constant).
[End OCR]*

## Page 38

*[Image OCR]
Introduction to the Instruments 23

A put butterfly will have the same price as a call butterfly. Examine the
98/100/102 butterfly (buy one 98 call, sell two 100 cali, buy one 102 call):

198C — 2100C + 1102C = 1 (98P + F) — 2 (100P + F) + 1 (102P + F)
=198P -2100P +1102P+1F~-2F+1F
= Put Butterfly.

A call butterfly 98-100-102, a put butterfly 98-100-102, and a condor 98-
100-102 will have the same exposure.

A 98-100-102 condor is defined as long a 98 put, long a 102 call, and
short the 100 straddle:

= 98P — 100P — 100C + 102C

= 98P — 100P — (100P + F) + (102P + F)

= 98P — 100P — 100P + 102P -F +F

= 98P — 2100P + 102P a put butterfly, which we established is
equivalent to the call butterfly.

To gain an intuitive feel the reader can verify that Figure 1.4 shows the
same P/L profile for the following:

Long 98, long 102, short twice the 100, all calls or all puts.
Long 98 puts, long 102 calls, short the 100 straddle.
Long the 98 calls, long the 102 puts, short the 100 straddle.

A result of these rules is that the volatility of an out-of-the-money put.
should be exactly equal to that of a corresponding in-the-money call of the
same strike.

2.5

Butterfly
2.0
1.5

1.0

intrinsic

0.5
0.0
—-0.5
96 98 100 102 104
Asset Price

Figure 1.4 Butterfly profile.
[End OCR]*

## Page 39

*[Image OCR]
24 Markets, Instruments, People

Risk Management Rules:

Traders should never carry put-call parity rules outside of a strike.

Some of the preceding rules can be used with soft American op-
tions, except when the delta of the option becomes too high.

AMERICAN OPTIONS, EARLY EXERCISE, AND
OTHER HEADACHES (ADVANCED Topic)

An American option poses more problems than European options because
the path followed by the underlying asset can lead to possible early exer-
cise. With a European option, pricing is a simple matter: One can just dis-
count the final payoffs on expiration day.

Without getting immersed in the pricing complications of American op-
tions, it is safe to say that the complexity arises because uncertainty about
the date of occurrence of the early exercise makes it difficult to model. The
rules depend on time and the amount of intrinsic value, which makes the
early exercise rules too uncertain.

Such early exercise is generally determined two ways: the soft (or easy)
rule and the hard rule.

Soft American Options

@ A soft American option (also called a pseudo-European option) is only
subjected to early exercise from the standpoint of the financing of the
intrinsic value.

An extension of this definition is that only one interest rate, that affect-
ing the financing of the premium for the operator, impacts the decision to
early exercise.

For risk management and trading purposes, soft American options will
be largely similar to the European options, except when interest rates be-
come very high relative to volatility. The reason they are often called
pseudo-European options is that they behave in general like European op-
tions, except when very deep in the money. The test of early exercise is
whether the total option value is less than the time value of the money be-
tween the time of consideration and expiration.

Example: Assume that an asset trades at $100, with interest rates at 6%
(annualized) and volatility at 15.7%. Assume also that the 3-month 80
[End OCR]*

## Page 40

*[Image OCR]
Introduction to the Instruments 25

call is worth $20, at least if it is American. Forgoing early exercise would
create an opportunity cost of 20 x 90/360 x .06 = .30 cents, the financ-
ing of the $20 premium for 3 months. The time value of the equivalent
put is close to zero (by put-call parity), so the intelligent operator can
swap the call into the underlying asset and buy the put to replicate the
same initial structure at a better cost. He would end up long the put and
long the underlying asset.

Hard American Options

@ A hard American option is an option subjected to early exercise tests
from the standpoint of both the financing of the intrinsic value and the
carry costs of the underlying asset until the nominal expiration.

By extension, two rates, that of the premium for the operator, and that
of the carry of the underlying asset impact the decision to early exercise.

An early exercise can thus be attributed, in addition to the soft Ameri-
can rule, to the yield benefits of an early position in the asset. The following
filter needs to be taken into account at all times during the life of the op-
tion: Would the operator do better if he owned an interest-bearing asset
than if he owned the equivalent position through the options?

Example: The operator owns the same call as earlier, but the underly-
ing asset is a currency that pays 20% interest while domestic rates are
6%. The option trader has the additional benefit of exercising the call
because, on top of the financing, he can own the currency that pays the
high interest rates against his home currency, which costs him only 6%
to short, therefore earning approximately 14% annualized on the same
value. The benefits of early exercise from the asset ownership are much
higher than in the previous case: .14 X 100 (face value) < 90/360 = $3.5.
In addition to that, the trader has the value of getting the cash much ear-
lier, which was computed in the previous case at $.30. The operator
should perhaps have exercised the option much earlier to benefit from
this extra kicker.

Likewise the put on the high-yielding currency will not be subjected
to early exercise. Take the following rule: A market trades at 100. The
120 put with $20 of intrinsic value will be better held till expiration
since the operator would have to pay $3.50 to hold the equivalent short
position in financing differential. True, the 30 cents of financing the
premium would be deducted from that value in the cost-benefits analy-
sis. So the put would not be early exercised by an optimal operator. It
will trade and will be considered for all intents and purposes as a Euro-
pean instrument.®
[End OCR]*

## Page 41

*[Image OCR]
26 Markets, Instruments, People

These rules of hard early exercise extend to the following instru-
ments:

¢ Bonds with a Positive Carry. American calls are early exercisable
when the bond is financed more cheaply than its yield. American puts
will therefore be similar to (but not quite the same as) European puts.

¢ Bonds with a Negative Carry. The reverse is true.

© Equities. When the equities are negative carry (and the dividend
payout is known), the calls resemble European ones and vice versa.

Never trust the price for an American option’ on a cash instrument.
There may be changes in the parameters throughout the life of the in-
strument (nondividend paying stocks starting to pay, changes in inter-
est rates, reversal of the interest rate differential) that can affect its
future value. Parameters, unfortunately, are not frozen.

Options on futures, therefore, will be considered subject to the soft
exercise rules. We will generally fold them within the category of Euro-
pean instruments for the purpose of our analysis. At all times, the test
will be that of the financing of the premium; differences only become
pertinent between European and American options for deep in-the-
money instruments.

European instruments tend to prevail where there are differences
in the pricing mechanism. The market generally seems to go to the most
liquid instrument, and operators—by shying away from complicated
options—make them less attractive to trade. European options domi-
nate with the currencies, and represent close to 99% of the volume of
those with a high differential.

e

Option Wizard: The Simplicity Rule

A major rule should be taken into account: The market always tends to flow to
simplicity. Complexity is generally costlier to both monitor and produce, and
somehow in the long run, demand moves away from the complex in favor of
the simple. New, elaborate contracts certainly attract people, but typically the
novelty wanes. Operators will then try to satisfy their needs for protection by
seeking the cheapest possible way.

Complex products cost more to replicate. An optimal operator becomes
cost conscious and avoids enriching the financial institutions when he can sat-
isfy his interests more economically. This becomes noticeable with the life or
death of listed financial instruments. It is the recipe for the survival of exchange
contracts. That rule will be discussed in the study of exotic options.
[End OCR]*

## Page 42

*[Image OCR]
Introduction to the Instruments 27

A Brief Warning about Early Exercise Tests

Most operators have their risk management system flag an early exercise
and routinely terminate the option. This is often less than optimal (and
downright dangerous) for the following reasons:

¢ One should refine the test by using a proper volatility for the corre-
sponding out-of-the-money option of the same strike. Most risk man-
agement systems do not apply the proper volatility smiles correctly.
For example, if the option early exercise test is done at a volatility of
16% and the strike price concerned trades at 20%, the early exercise
flagged by the system would be erroneous.

¢ An additional refinement should be the use of a volatility term struc-
ture by retesting with a longer or shorter duration, inputting the
highest possible volatility between time zero and the expiration of
the option.

* When dealing with a large-size position, dealers, upon exercise, syn-
thetically become short an out-of-the-money option. It is recom-
mended to test for the liquidity of such strike in the market and
compute the replacement costs.

¢ A war story: The day before the stock market crash (of 1987) X, a mar-
ket maker in the Eurodollar options, found a deep-in-the-money “re-
versal” in his books (a reversal means that the trader is long a put, short a
call, short the future). According to his risk management system, the put
was early exercisable. He exercised the put staying naked the calls.
The following day, the market experienced a 10 standard deviation
rally on the open, putting him out of business and, worse, causing his
story to become a legend. Ironically, X belonged to the category of
“wings” buyers (people who always own out-of-the-money options).
He never shorted “wing” options and ended up hurt by a synthetic
(and entirely accidental) short.

Consequence: Smile-Calendars (Advanced Topic). The volatility of op-
tions of the same strike needs to be equal to allow for put-call parity equiv-
alence. While this rule applies unconditionally for European options, many
operators mistakenly apply it to the American variety. In some cases, the
rules can be applied, and in other cases they need to be lifted.

When a strong skew is impacting the market, the put-call parity can be
weakened considerably by the following:

¢ A rising volatility curve could separate the put and calls from each
other because the nonexercisable leg would follow the nominal matu-
rity, whereas the exercisable one would have a considerably shorter
expected life.
[End OCR]*

## Page 43

*[Image OCR]
28 Markets, Instruments, People

Example: Assume that 3-month options trade at 15.7% volatility but
that the 1-day options traded at 13%. An exercisable 80 call would be
priced at 13%, whereas the nonexercisable leg (the 80 put) would trade
at 20%.

e A strong smile (defined as an implied volatility that is a function of
the strike price and time to maturity) can present worse results.

Example: As shown in Figure 1.5, assume that the 80 strike for the
3-month option trades at 20% volatility (point A), while the at-the-
money trades at 15.7 for 3-month (point B) and 13% for 1-day options
(point C). The 3-month 80 call will trade at 13% volatility (point C),
while the 80 put would trade at 20%.

Risk Management Rule: Hard American options are more valu-
able than European options of the same expiration, because they

harbor a compound option on interest rates or volatility. The differ-
ence between them increases with either the volatility of volatility or
the level and volatility of interest (or carry) rates.

The rule is easy to explain to traders who had to live in highly unstable
interest rates or fluctuating implied volatility (vvol). The American option
gives the owner the right to pay the carry to “extend” the option one
additional day and therefore make a bet that the option may no longer be

20%

0,
15.70% 3-Month Curve

13%
1-Day: Curve

80 90 100

Strike

Figure 1.5 Smile curves.
[End OCR]*

## Page 44

*[Image OCR]
Introduction to the Instruments 29

exercisable. As such it becomes an extendible option (for a discussion of
compound options see Chapter 21).

American Options Nobody Ever Exercises. Some contracts for American
options do not present any early exercise value. They are the options on fu-
tures where there is a marks-to-market daily. For example: London Inter-
bank Financial Future Exchange (LIFFE) option products are margined as a
future. The profits can be taken out daily and can thus earn interest. Buying
a deep-in-the-money option requires a much smaller outlay of cash than
other markets.

FORWARDS, FUTURES, AND FORWARD-FORWARDS
(ADVANCED Topic)

@ A forward contract between two parties obligates both of them to buy
and sell a given asset on a specified date, or to exchange payments ac-
cording to a formula.

-@ A future is a standardized forward contract listed on an exchange
with set maturities where the exchange clearing house is the counter-
party. Trades take place in an open outcry system and the liquidity is
improved by the standardization and the interchangeability between
contracts. * :

There is a difference in hedge ratios between the forward and the fu-
ture. In the forward, the exchange of payments takes place at a terminal
date, while with the future there is a variation margin system of pay-as-
you-go. The parties need to exchange payments that correspond to their
daily profit or loss. A winner can thus earn interest on the profits. Such dif-
ference can be substantial for long-dated instruments.

To hedge a forward with a future, the trader needs to “tail,”* to adjust
for the difference between a present value claim and a cash one. The hedge
(number of units of futures for one unit of forward) is:

h=e

where f is the number of years and r is the zero-rate until the forward
delivery.’

Example: A 3-month listed future on the Chicago Merchantile Ex-
change (CME) in the Deutsche Mark (DEM) against the U.S. Dollar
(USD) trades at .70 cents per dollar (one whose expiration happens to
fall exactly 90 days from today). The 90-day forward, by arbitrage, also
trades at .70 (expressed in spot at 1.4286).
[End OCR]*

## Page 45

*[Image OCR]
30 Markets, Instruments, People

What is the equivalent Deutsche mark (DEM) amount to 100 futures
(a future on CME represents 125,000 DEM)? Assume 90-day interest
rates are 6%.

Total face value is DEM 12,500,000. The hedge ratio h is Exp
[ —.06 x .25] = .985, number of futures for one forward. The hedge in
the forward is therefore: 12,500,000/.985 = 1,269,000.

If the DEM future immediately moves to .71 the profit on the future
leg will be $125,000 deliverable immediately. The profits in the forward
will be $126,900 but will be only a mark-to-market realizable in 90 days.

Credit

Credit is another difference between futures and the forwards. Generic
swaps, caps, and floors (with futures equivalent dates) are easy to hedge
with strips and Euro options. Going to the over-the-counter market would
create a mismatch that consumes two-way credit facilities. If Credit Syl-
davia engages in a swap with Banca Nazionale del Lavoro, both parties will
reduce their credit lines to each other and swell their contingent balance
sheets. Should both unwind their side of the trade at different times with
different counterparties who in turn trade with each other, the numbers
would multiply: Each of the four parties would reduce the size of its books
and have its lines to the rest of the world reduced as well. Had all the trades
taken place on a standardized exchange, everyone would be flat.

Typically, the way market makers operate is to use the exchange when
initiating a trade (when they act as a customer) and to trade a forward
when they are acting as a market maker, with other parties calling them
and dealing on their prices.

Marks-to-Market Differences

The marks-to-market rules create non-trivial differences between the for-
wards (and similar instruments) and the futures (see Table 1.2). Forwards are
self-marked by operators according to some convention. A few points are up-
dated by the trader and the back office and all the intermediate points on the .
curve are computed by whatever algorithm is used by the system.

The future, on the other hand, is marked by the exchange through some
well-defined rules. The Eurodollar curve, about 40 contracts, is not gener-,
ated through computer algorithms but with the last trade or bid/offer in
mind, even if the resulting curve becomes jagged. An operator who marks
the same curve with fewer points will generally have a smoother result.

Another issue lies in the timing. It is called the nonsynchronous marked-to-
market problem. Many operators carry positions on one exchange offset
against positions on another. Exchanges do not have the same settlement time.
A future marks-to-market would not accurately portray the resulting P/L.
[End OCR]*

## Page 46

*[Image OCR]
Table 1.2 Differences between Forward and Futures

Marks to Market
Variation Margins
Credit Risks

Instrument Hedge

Trading Risks

Trading costs

Sensitivity to Financing

Introduction to the Instruments 31

Forward

Institutional, self-created
None

Depends on the
institution

Delta is present valued
(see Greeks)

Higher exposure to
illiquidity as the
contracts rapidly move
away from the liquid
maturities

Direct costs are gener-
ally low

Indirect costs are higher,
as the bid /offer spread
generally larger

Nonsensitive to the
correlation with the
financing rate.

*Except, of course, in “emerging” markets.

Example:

Future

Official settlement
Daily
Almost nonexistent:*

systemic rather than
limited to one institution

Delta is “raw,” not
present valued

Higher liquidity but
fewer “pillars” of trading

Direct costs are high:
commission, clearing
charges, exchange fees
Indirect costs are lower:
“spread” is tighter

Sensitive to the
correlation between price
and the financing rate

An arbitrageur plays the forwards in USD-JPY (Japanese

yen) by trading forwards in the over-the-counter market against EuroYen
futures in Singapore and the Eurodollars in the United States. The bank
marks the positions at 4:30 p.m. New York time, while the Singapore fu-
tures are marked before he starts his day in New York and the Eurodol-
lar at 3:00 p.m. New York time. His resultant P/L will never reflect the
accurate liquidating value of his position.

The Correlation between the Future and
the Financing (Advanced Issue)

In the preceding situations, there is independence between the financing
rate of the P/L stemming from the future and the expected moves in the
[End OCR]*

## Page 47

*[Image OCR]
32 Markets, Instruments, People

future. A correlation between the financing and the future will translate
into a convexity or concavity of the future compared with the forward.

Risk Management Rule: When there is a positive correlation be-
tween the financing rate r and a future contract F" (on any possi-
ble underlying asset) subjected to marks-to-market rules, the
future will be convex and will trade above the forward of the same
delivery.

Converse: Whenever there is a negative correlation between
the financing rate r and a future contract F, the future will be
concave and will then need to trade at a lower rate than a corre-
sponding forward.

An illustration of the rule in provided in Chapter 10.

Forward-Forward

@ A forward-forward is a contract to exchange an asset at one period
against the reverse trade at a later period.

For quasi-linear derivatives, such as fixed-income instruments, it is the
price ratio between two forwards.

FF(t1, #2) = F(t2)/F(t1)

The forward-forward rate is determined by the existing rates in the
market interpolated to solve for the break-even rate for the period between
f1 and #2. .

The Eurodollar futures are forward-forward deposits, and ironically,
the forward-forward often sets the spot (the tail wagging the dog, as often
repeated in future circles). ,

For options, the forward-forward is computed with the nonlinearity of
time in mind. This will be dealt with in Chapter 9.

Core Risk MANAGEMENT: DISTINCTION BETWEEN
PRIMARY AND SECONDARY RISKS

Market risks can be primary or secondary, but sometimes the distinction
can be counterintuitive; some instruments and markets present more dan-
ger in the fringe risks than in their primary exposures.
[End OCR]*

## Page 48

*[Image OCR]
Introduction to the Instruments 33

A primary risk is one that constitutes the bulk of the variance of profits
and losses (the P/L), and is where most hedging efforts should be concen-
trated. Since markets move rapidly, it is easy to see the primary risks as
those a trader would cover first, and the residual risks are those that can
usually wait until the end of the trading session.

The following classification of risks excludes the party-spoiling corre-
lation-based products (like options of one of two instruments and other
goodies).

¢ For an equity derivatives portfolio, market risks are almost entirely
directional, with all the possible permutations concerning the deriv-
atives of the effect of the underlying equities. Matters pertaining to
interest rates are deemed secondary, as the thrust of the positioning
is not at the level of these parameters. Someone can be hurt from the
indirect effects of an interest rate change on a long dated stock posi-
tion (through the effect on the pricing of the forwards), but such vari-
ance would be insignificant compared with that caused by the price
action and the changes in volatility. Such analysis does not reflect
the possible effect of the interest rates move on the equity markets,
just their effect on the time structure or equity prices. The health
risks of passive smoking are not seriously significant to a man al-
ready diagnosed with cancer.

¢ For a fixed income derivatives book, however, both the underlying
asset (that is the cash flow schedules as discounted by the interest
rates) and the general structure of the interest rate curve need to be
taken into account as primary risks. This is mostly attributable to the
consideration that fixed-income positions react to a schedule of
prices spread over time, not just to one price like an equity or a cur-
rency. Every coupon paying fixed income security is a simple sum of
smaller zero-coupon securities with different expirations. A 10-year
swap’s price will principally be affected by the 10-year rates. It will
also be affected by the 4-year rates, everything else being equal, be-
cause of the effect on the reinvestment over the period. A term struc-
ture becomes an integral part of the risks and rewards of every
instrument and should therefore be considered a primary risk.

¢ For a currency book, the exchange rate of the currency pair and the
volatility are the primary risks. But both the interest rate differen-
tials and the term structures of interest rates are preeminent; they
become part of the primary risks with some categories of developing
markets that experience a high interest rate volatility.

Such classifications need to adapt to most possible instruments: Those
who have tried to come up with a generalized theory of risk management
[End OCR]*

## Page 49

*[Image OCR]
34 Markets, Instruments, People

have so far failed in their laborious attempts. No two instruments will be
equal. Traders, therefore, must search initially for common points between
instruments.

Tables 1.3 and 1.4 will apply to the SDF (smallest decomposable
fragment).

Liquidity is preeminent as a risk (the invisible risk as traders concede). It
will be examined in this book at all levels without fitting it in any category.
Liquidity is the source and the cause of everything in trading and it should
remain in the back of every risk manager’s mind.

Since all options are basically alike this text adheres, whenever possi-
ble, to a standardized example. Most examples in the book depict option
risks for a generic asset with a flat forward. This approach simplifies the
effect of the projection into the future and isolates the pure option risks of
the instruments.

In other words we rescale everything to a numeraire. Hence the forward
will trade at 100% of cash regardless of the expiration. Such simplification
allows a focus on the option risks, without invalidating the real issues.
Where the simplification leads to inaccuracies, a tradable instrument will
be used taking into account its particularities. The drift and the notion of
Girsanov change of probability will be introduced in cases where it affects
the hedging technique.

Initially, readers will measure the risks of the derivatives, with simple
cases where a “pure” portfolio is created. The basic “Greeks” (vega, gamma,

Table 1.3 Primary and Secondary Risks by Market

Market Primary Market Risks Secondary Market Risks
Equities The underlying equity The domestic interest rate
Volatility ° Dividend payout
Volatility term structure
Fixed Income Rates Higher order derivatives
Term structure of rates Risks of pricing formula
Volatility Volatility term structure*

Stability of the covari-
ances between maturities

Currencies Price Rates ineach currency *
Volatility Volatility term structure

Commodities Price The domestic interest rate
Volatility Storage costs

Term structure of prices _—- Volatility term structure

* It is assumed that a cap or a floor are decomposed into caplets and floorlets.
[End OCR]*

## Page 50

*[Image OCR]
Table 1.4 Primary and Secondary Risks by Instrument

Instrument

Barrier Options

Introduction to the Instruments 35

Additional Primary
Market Risks

Term structure of rates
Term structure of vol-

Additional Secondary
Market Risks

Skew
Volatility of volatility

atility
Liquidity at the barrier

Asian Options Term structure of rates Model Risk
Term structure of vol- Skew

atility
Variance ratio

Lookback Options Mean reversion Volatility of volatility
Arcsine law

“Driftwood” effect

Compound Options Volatility of volatility Mean reversion

Multidimensional Correlation measure

Options Volatility of volatility
(hence correlations)

time decay) will be examined before readers modify their measures as prac-
titioners do routinely, whether they are conscious of these transformations
or not. This will be followed by a thorough presentation of the modified
delta, the modified vega, the modified theta, and the modified gamma.
This blanket method should be possible to apply to all instruments.
The generating blocks of some exotic options—bets, barriers, and corre-
lation—will be the final objects of study. Readers will be shown a road map
‘for extending these principles to other, complex instruments. To apply
these techniques to a book of specific instruments, it is important to learn
to decompose the option risk from the residual risks proper to the instru-
ment. This is best done through the decomposition of the instrument into
liquid traded segments (see Figure 1.7).

APPLYING THE FRAMEWORK TO SPECIFIC INSTRUMENTS

A swap is simply a multiasset instrument composed of correlated seg-
ments often entirely decomposable into very liquid Eurodollar strips. The
- framework on multiasset options and the techniques of stacking with cor-
relation matrices should be sufficient for the trader to adapt the principles
of dynamic hedging to swaps. As for index-amortizing swaps, the use of
[End OCR]*

## Page 51

*[Image OCR]
36 Markets, Instruments, People

General
Case
Options stripped of the particularities of the
assets

Pure

Second Order Option Risks
Option Risks: Delta
Vega Stability Vega (modified)
Skew Gamma (modified)
Theta (modified)

Spatial Dimension ,
Topography, Barriers

Time Dimension
Bucketing, Rho
Homogeneity

Time and Space
Relationships:
Bleed
Higher Derivatives
Path Dependency

Parameter Convexity

Figure 1.6 Options risks defined.
[End OCR]*

## Page 52

*[Image OCR]
Currencies
Forwards

Introduction to the Instruments 37

Residual Risks
Particular Cases

Option on
Futures

Figure 1.7. An analysis of residual risk by groups of instruments.

American digital options should perform the task, mixed with some
knowledge of compound options where applicable. The problem with most
fixed-income traders is that they tend to limit their knowledge to the intri-.
cacies of the instrument rather than focus on the generating blocks. This
method is considerably easier than the holistic Heath-Jarrow-Morton

approach.
[End OCR]*

## Page 53

*[Image OCR]
Chapter 2

The Generalized Option

A drunk man will find his way home. A drunk bird may get lost forever.
Kakutani'

Most traders distinguish between vanilla and exotic options. The vanilla op-
tion and vanilla forward are easily priced in the market and benefit from liq-
uid markets. A nonconventional structure would then be called exotic. On
the Chicago Board Options Exchange (CBOE), where puts were introduced
in 1978, Jim Piper? recounts that puts were then exotic and calls vanilla.
There is no true functional difference between the products, except that

pricing exotic options often seems to be more complicated. For risk man-_

agement purposes, six dimensions of analysis are crucial for understanding
and pricing options (see Figure 2.1).

Step 1. THE HOMOGENEITY OF THE STRUCTURE

@ A time-homogeneous structure refers to a payoff structure that does
not contractually change through time. The contract is said to be uni-
form with respect to time, between inception and expiration.

Some structures have restrictions on them as the payoff depends
on whether some event took place before or after a specified date. So the
first question to ask is whether the structure has any time-changing char-
acteristics.

Example: The characteristics of a deferred strike option will depend
heavily on whether an event takes place before the day when the strike
is set or after. More practically, the option will appear to have no delta
and no gamma,’ but will have some vega between inception and the
strike setting time. After the strike setting, generally the settled spot
price on a given date, it becomes a normal vanilla (boring) option. The
contract for such an option defines two periods—Period 1 when the op-
tion will be (on paper) gamma-less and so on, and Period 2 when the op-
tion will present vanilla features.

38

x
[End OCR]*

## Page 54

*[Image OCR]
The Generalized Option 39

General Option
Time Non-Time
Homogeneous Homogeneous

Digital Ramp
Payoff Payoff

a)

Barrier No Barrier

Number of Assets

i

Order of the Option
First Order, Compound...

Path Dependent Path Independent
Memor less Hard Path
y Dependent

Figure 2.1 The general option.
[End OCR]*

## Page 55

*[Image OCR]
40 Markets, Instruments, People

Option Wizard: The Smallest Decomposable Fragment of a Structure

Every structure for risk management purpose needs to be decomposed into
smaller units until it is no longer deemed decomposable.
The term decomposable means “composed of structures added together.”

¢ A double barrier option is not the addition of two barrier options.
° A strangle, however, is the addition of a put and a call.

¢ A European double bet is decomposable into two European single bets. An
American double bet (“either or”) is not decomposable.

¢ An option on a basket is not decomposable into smaller options on the com-
ponents. Likewise an option on a spread is not decomposable into two op-
tions on the components.

* Cap or floor options are decomposable into “caplets” or “floorlets.” An op-
tion on a swap, called swaption by some, is not decomposable (it is akin to
an option on a basket).

¢ A European barrier option will not be covered in this book as it is a simple
addition of a digital and a call spread.

¢ Some options are limit decomposable: They can be broken down into
smaller options but in either infinite amounts or in strike pairs that are infi-
nitely narrow. A lookback option is decomposable into an infinite series of
knock-in and knock-out with strike prices infinitely close together. A binary
call is limit decomposable into an infinite amount of infinitely narrow call
spreads, and that knowledge is useful in trading and pricing although no-
body ever tried to perform the replication.

A reminder: For a risk-conscious dynamic hedger (and subjected to trans-
action costs), the returns of a sum of SDF in a portfolio will be markedly dif-
ferent from the sum of the returns of independently managed SDF.

All options intrinsically change over time. The question here is whether
the contract is written in such a way as to alter the payoff after certain
dates. Figure 2.2 shows the time to expiration “sliced” in periods. Each pe-
riod will have its own associated payoff. For example, a window option is a ,
barrier option where the barriers are lifted after a certain period. With a
flexible barrier option, the barriers can widen by a percentage point every
month. First month: The barrier is at 109 (with, as usual, the spot at 100);
second month, the barrier is at 110; third month, at 111; and so on.

Nontime-homogeneous structures can be quite difficult to manage.
Later chapters will describe their risks and how to handle Greeks that can
accommodate the changes over time.
[End OCR]*

## Page 56

*[Image OCR]
The Generalized Option 41

Figure 2.2 Slicing time.

They also can be the most difficult to price, though not always. The
model needs to take into account that the payoff changes according to the
period. Quants often have difficulties with options that are very easy to
hedge (like American options) and find it often trivial to deal with options
that are arduous to trade (like bets and binary options).

Step 2. THE Type OF PAYOFF:
CONTINUOUS AND DISCONTINUOUS

@ Digital (discontinuous) and ramp (continuous) payoffs are the two
kinds of payoffs (Figures 2.3 and 2.4). Ramp payoffs are continuous be-
tween points, and digital payoffs pay “all or nothing,” with no shade in
between. They are akin to bets between two people on the outcome of

an event.

Derivative Price
Terminal Value

Underlying Asset

Figure 2.3 Digital payoff.
[End OCR]*

## Page 57

*[Image OCR]
42 Markets, Instruments, People

Derivative Price
Terminal Value
Underlying Asset

Figure 2.4 Ramp payoff.

Example: Assume that a person bets on the outcome of a well-publicized
murder trial, as traders often do, in the amount of $1,000. There will be
no intermediate payoff between the loss of $1,000 and the gain of $1,000.
The payoff is then said to be discontinuous.

There are two (good news/bad news) points to consider when dealing with
the discontinuous payoff options.

1. The bad news is that almost all available hedges in the market are
continuous payoff products, therefore creating imperfect or unstable
hedges. There may be constructions that provide an accurate hedge
(such as vertical spreads), but these constructions are too costly to
execute and generally are unavailable.

2. The good news is that the bet option has a small bite. It is a relatively
harmless product for those who trade-it as it should be traded—as a
bet. Dynamic hedging is to be avoided with these situations.

This distinction between payoff types is ever present in structures.
There are bet options that are American and pay “if touched” and bet op-
tions that pay if some condition is met on expiration. They also provide us
with good grounds for the transition to barriers.

Example—European: A three-year dollar denominated note pays 7%
every coupon date when dollar-yen is above 100 and 4% otherwise. This
note could be constructed with a regular note of the same maturity and
a strip of digital call options struck at 100 expiring on the coupon dates.

Example—American: The same note as before except that it pays 8%
unless the yen trades at 110 at some point during its three-year life.
[End OCR]*

## Page 58

*[Image OCR]
The Generalized Option 43

Figure 2.5 Crossing a barrier.

Step 3. BARRIERS

@ A barrier is a price level in the market (called a trigger) that, when
reached, markedly alters the payoff of the structure.

Perhaps the most entertaining aspect of option trading is the barrier. Many
traders have witnessed “trigger” prices being crossed with heavy heart-
beat. They are also interesting for the researcher because they call on a col-
orful branch of mathematics (probability theory).

The trader must ask whether the structure has a barrier terminating it
(knock-out) or initiating it. Figure 2.5 shows an asset price crossing the up
barrier.

@ A high barrier (or up barrier) is a barrier located above the present spot.
A low barrier (or down barrier) is a barrier located below the spot.

Step 4. DIMENSION OF THE STRUCTURE
AND THE NUMBER OF ASSETS

‘WM The dimension of the structure is 1 + the number of variables affecting
its value (one is for time).

As shown in the tutorial introduction to the random walk in Module A, the
conventional “vanilla” option has two dimensions: the asset and time. With
two pieces of information: the asset price (y axis) and the time (x axis), the
information can be represented in a two-dimensional graph. The market
could be illustrated by the classical random walk of a drunk man down
Madison avenue.
[End OCR]*

## Page 59

*[Image OCR]
44 Markets, Instruments, People

When more than one variable is involved in the structure, traders need
to deal with the notions of correlation and independence.

There are many types of higher dimensional structures: In the “hard” 3D
structure, there is more than one strike price; in a “soft” 3D structure (e.g., a
convertible), the level of interest rates matters in the construction itself.

An American currency option is particularly sensitive to interest rates
and the curve. It is therefore considered a structure of higher dimension.

The custom of derivatives traders to hedge themselves in liquid instru-
ments causes all the instruments to become multiasset (higher dimension
by construction).

@ A multiasset structure by construction is one that results from the
hedge of an illiquid structure with liquid options. When an option in an
illiquid market is hedged with a “similar” option in a liquid one, the
“similarity” becomes a risk.

This increase in the dimension of the structure is generally necessary with
strange instruments.

Example—Multiasset 3D by Construction: A trader at a bank sold to a
customer (oh no!) a DEM-AUD (Deutsche mark-Australian dollar) cross
option. The option is clearly hedgeable with the underlying asset (the
DEM-AUD spot and forward markets). However, the spot market in that
country is not very liquid. So the trader breaks up the structure into two
units: USD-DEM and AUD-USD. He will hedge his spot and forward ex-
posure there, thus getting into some first order (delta) neutrality. The
‘reader will find in Module C, the meaning of the orderings between two
currencies or between a currency and the numeraire “base currency.”
The trader, being wise, requires further hedging to reduce the
volatility risks of the book. He thereforé buys some of the volatility in
USD-DEM and some in AUD-USD, as these are liquid in the market.
(The proportion will be analyzed later.) He will now have a multiasset
portfolio, with the following prices: USD-DEM and AUD-USD, whereas
previously he only had one of them to consider. The concept is hard to
understand but the two different assets will yield three different volatil-
ities (one for each resulting pair) linked together by two correlations.

Example—Multiasset 4D by Design: ‘he trader in the earlier example.
received a phone call with bad news and, as a result, had to trade a Greek
Drachma-Australian dollar large option position (symbol GDR-AUD).

He broke up the volatility risk into liquid subcomponents: GDR-DEM,

USD-DEM, AUD-USD and traded each one of those against it (3 assets +

time = 4 dimensions).
[End OCR]*

## Page 60

*[Image OCR]
The Generalized Option 45

Option Wizard: How to Get in Trouble Hedging

Financial markets history is littered with stories of traders who bought an in-
strument and tried to reduce the risks by selling a “highly correlated” one
against it, only to discover that they doubled their risks. The trader after seeing
on the screen the price of one of the two instruments go down (the one he is
long, of course) and the other go up (the one he is short) will blame markets for
not being well behaved.

A deeper analysis would show that, typically, some of the instruments that
are easy to buy against easy-to-sell “correlated siblings” are invitations for
trouble. They act as a trap that will attract many hedgers and arbitrage traders
then force them into noisy liquidations. This effect can lead to disastrous re-
sults on gullible managers using such apparatus as the “value at risk.”

Example—Swaps and Bonds: Swaps would be easy instruments to
trade, except that the exact maturity and coupon are never as liquid as
operators wish. These swaps will need to be matched, alternatively,
with other swaps that, with a different coupon, will trade on their own.
When hedged with strip (Eurodollar futures added up in a weighted
manner), swaps will become sensitive to the correlation between the
components .

STEP 5. ORDER OF THE OPTIONS

@ Higher order options refer to contingent options. An option on another
option is a second-order option. An option on a second order option is a
third-order option, and so on. ‘

Higher order options are becoming more and more common with financial
products. Many options, called extendible, have built into them a renewal of
the contract at a certain price. They can be extended, say five times, for a
predetermined $1 each. However, should the option not be renewed at any
extension period, the entire structure expires worthless or the underlying
asset is delivered at the strike price like a regular delivery.

As the trader goes into higher and higher orders, he becomes con-
fronted with the contingency of a contingency. Positions will move further
out-of-the-money, and the risks will become less and less linear. The com-
pounding of a probability with a second-order option reduces the certainty
of the payoff and will therefore require less hedge than a regular option.
The delta will be lower but very unstable, to the point of carrying little risk
[End OCR]*

## Page 61

*[Image OCR]
46 Markets, Instruments, People

disclosure. Delta and fist-order Greeks will be less meaningful and their in-
stability will require the risk manager to have recourse-to higher order de-
rivatives, such as the third and fourth moment, to assess the portfolio risks.

Example: A common use of compound options is with warrant issues.
The price of the warrant is set at some point, say at $2, but the investor has
a few hours or days to make up his mind as to whether he wants to buy
into it or not. The option dealer is thus short an option on the warrant.

Strep 6. PATH DEPENDENCE

M Path-independent options are options whose payoff depends exclu-
sively on the events upon expiration, regardless of the route taken. A
path-dependent option depends, in addition to the final price, on at
least one price in the path taken to get there.

A European option’s price for the buyer depends exclusively on the final
price. If someone buys the 104 calls, he stands to benefit if the security gets
above 104 upon expiration; otherwise, he would lose his premium. How the
spot gets to that point is theoretically irrelevant. The option is therefore
said to be path-independent. As will be shown, such path independence is a
myth. Every security is path-dependent for a person, including an end user, |
who has a marked-to-market and the power of a decision throughout its life.

Figure 2.6 shows two paths leading to the same terminal value.

There are two types of path dependent categories: the soft and the hard
path dependent options.

@ The soft-path-dependent option (also called “memoryless path depen-
dent) depends on one piece of information. The hard-path-dependent
option (also called “path-dependent with strong memory”) will present

Figure 2.6 Two paths leading to the same terminal value.
[End OCR]*

## Page 62

*[Image OCR]
The Generalized Option 47

a payoff that is entirely dependent on the path taken during some inter-
val. It is therefore truly path-dependent because all the roads do not
lead to Rome in the same fashion.

A barrier option is a soft-path-dependent option because at any time it de-
pends on whether the asset traded at or through the barrier or not.

A lookback option is another soft-path-dependent option because it de-
pends on one piece of information from the path: the extremum. What suc-
cession of events lead there is not relevant (the lookback options are
somewhat related to the barriers in the sense that they can be decomposed
into series of knock-in and knock-out options).

Options on averages are hard-path-dependent. Every observation will
count and the value will depend on every step taken by the security. Den-
sity, however, affects how strong the memory is, and what constitutes a
path dependence. A dense path-dependent option cares about every piece
of information, whereas a lighter one cares about infrequent ones.

Example: An Asian option will present dense path-dependent fea-
tures if it is based on hourly prices and light features if the sampling is
monthly or quarterly. The effect this has on the final hedging will be
discussed later in the book.
[End OCR]*

## Page 63

*[Image OCR]
Chapter 3

Market Making and Market Using

Anybody can buy and sell.
James C. Powers

The goal for this chapter is to describe the modus operandi of the various
participants in the market (see Figure 3.1).

Pit traders’ jargon identifies a three-tier classification of the market par-
ticipants: the local, the paper, and the arb.

@ The local is a floor market maker. A pure price-sensitive lot, locals are
entirely liquidity driven in their trades. They are the party benefiting
from the time and space advantage.

@ Paper means end customers. Their name is a reminder that in the past a
runner would carry their order (literally running) through the floor of
the exchange written on a piece of distinctive paper. Locals are always
eager to trade with paper because they feel that they have the edge. A
local who wants to trade with other locals generally needs to disguise
his order to make it look like paper.

@ Arbs are arbitrageurs or a mixture between market makers and end
users. They are the ones carrying the infamous program trading in the
SP500 (a cash-future arbitrage) and transfer liquidity from one market
to another (for a small fee, of course).

Book RUNNERS VERSUS PRICE TAKERS

Book runners, also called market makers, perform this task with a product
against an inventory, generally called a “book.” Book runners generally
post some price at which they are willing to buy (their bid) and some price
at which they would be consenting sellers (the offer). They consider that
they can benefit from positive expected returns, the fair odds, because they
hope to earn a large part of the difference between the middle of the market
and the bid or offer, thus achieving transaction positivity (while most other

48
[End OCR]*

## Page 64

*[Image OCR]
Market Making and Market Using 49

Traders

End Users
“Buy Side”

Book Runners
“Sell Side”

Hedge Funds Corporate Unleveraged |.

Treasuries Funds

Proprietary
Traders

ZO

Self-
Proclaimed

“Directional” True Hedging Speculation

“Arbitrage”
Traders

(overt or
disguised)

Traders

Figure 3.1 Classification of market participants.

participants in the market incur transaction costs in proportion to the liq-
uidity of the instruments involved).

Price takers are market users, a function diametrically opposed to that
of the book runner. A market maker acts as a liquidity conduit, standing be-
tween end users, for a compensation. The principal difference in attitude
between a book runner and an investor lies in their relative price sensitiv-
ity. A market maker’s willingness to take on a position will increase as the
price becomes more attractive; not so for the market user (except in the
event of liquidity holes). A market maker reflecting a strong view in all of
his trades would ultimately bankrupt his firm as the volume of his transac-
tions would then become one-sided and cause an extravagant accumulation
of positions.

Most of the academic literature distinguishes between “upstairs” and
floor traders. “Upstairs” traders refers to price takers. Nowadays market
makers in complex products are located “upstairs.”

No operator is entirely a market maker (that is always devoid of opin-
ion) and no investor is entirely price insensitive. Most investors try to min-
imize transaction costs, sometimes even earn the spread by working limit
[End OCR]*

## Page 65

*[Image OCR]
50 Markets, Instruments, People

orders. Effective market makers “lean” on a position, that is try to keep a
core long or short position and skew their prices. Typically their bias is de-
fensive. In a trending bull market, book runners intuitively fear their get-
ting short from the overabundance of buy orders and reflect the preference
for a long position by shading their bids and offers upward (and conversely
in a bear market). Similarly price takers might delay entering a trade as the
price becomes less attractive although occasionally they take an opposite
position than intended as a sole result of a price adjustment. Sometimes
price sensitivity may be negative: An operator with a belief in trends might
be enticed to buy when prices get higher or sell when prices drop.

Finally the difference between a market maker and a market user lies in
that the former needs to be defensive whereas the latter has the freedom to
be aggressive.

COMMODITIZED AND NONSTANDARDIZED PRODUCTS

There are two types of markets: “commoditized” products that have en-
tered their liquid phase and “nonstandardized” ones where the instru-
ments are still immature and new (see Figure 3.2). Commoditized products
exhibit a high volume and high degree of liquidity but relatively small edge,
a narrow bid-to-offer spread, and a general overcrowding of talent. Their
advantage principally is that the market makers can make up for the narrow
spread with the high volume—“a fast nickel is better than a slow dollar,” as
the adage goes. They are so competitive that they require a certain sophisti-
cation in addition to market sense, a notion that most traders use and under-
stand but that nobody has been able to define accurately.

Commoditized products have standard agreements in place, eliminate
most surprises, and typically trade between dealers where constant match-
ing of risks takes place. The existence of an interbank market is the test of
standardization. They rank from the very simple cash products to some
lower forms of exotic options (e.g., single barrier knock-outs). Commodi-
tized products are price sensitive, but the volume is high and traders can
fine-tune their portfolios in a way to capture their edge.

Nonstandardized products, like structures, have payoffs that are pecu-
liar to the instrument itself and require special pricing capabilities, such as
an on-staff mathematician. In contrast, the commoditized products can be
priced and managed with the aid of commercially available software prod-
ucts (generally faulty). It can become necessary to design programs for
every trade, with a higher incidence of pricing “bugs.” An option with a
payoff attached to several assets, with a barrier that is reset six times and an
uncertain expiration date (it can be extended) will not be easily booked ina
commercial risk management system. It will require a special computer
subroutine to track its “Greeks.”
[End OCR]*

## Page 66

*[Image OCR]
Market Making and Market Using 51

Complex
Structures

Commoditized
Products

Linear “Volatility Nonstandardized
Derivatives Spreaders”

~Swaps dealers
-Forward
traders

Figure 3.2 Different categories of derivatives book running: Commoditized and
nonstandard products.

The real difference between commoditized and nonstandard products
is that one type is tailor-made, with a high price tag attached to it and
smaller traffic, while the other has the features of a discount store with
standard sizes and prices, but a higher volume. Table 3.1 lists these and
other major differences. All successful nonstandard products become com-
moditized. When many of us started trading options, these were mostly
nonstandard but the markets moved rapidly from one category to the other.
Swaps structures that were difficult to price only five years ago are now
given to junior traders.

Trading Risks in Commoditized Products

Cash Products. Thanks to the competition, these products represent less
rewards as they are usually more liquid and require no sophisticated risk
management tools. However, they allow the traders to completely offset
the risk of their trades and avoid the necessity of warehousing contracts for
any length of time. A cash trader will hold a position over his horizon by
choice only.
[End OCR]*

## Page 67

*[Image OCR]
52 Markets, Instruments, People

Table 3.1 Commoditized and Nonstandard Products

Commoditized Products Nonstandard Products
Price sensitive Product designation sensitive
Supermarket-like products Tailor-made products
Narrow bid /offer spread Wide spread, even absence
of two-way markets
Trade interbank, often Only trade between a financial
through a broker institution and an end user
Standard pricing formulas Unseasoned pricing
Trading talent and experience Pricing ability required, mathematical
required sophistication
High volume Low volume
Standard risk management systems Tailor-made systems, needs on-staff
programmers

Option Wizard: Trader Decay

As products increase in sophistication, traders become rapidly obsolete. This is
becoming significant in most markets.

Unlike practitioners of most professions, traders are not given time e (or do
not give themselves enough time) to adapt to the new products or learn new
techniques. Traders are constantly under profitability pressures that force them
to focus on “producing” lest they join the ranks of laid-off Wall Streeters. The
trader finds himself a niche and milks it until its (and his) termination.

In addition, holding a trading position requires such an intellectual com-
mitment that the trader may feel consumed to’the point that other functions of
the mind slow down. One comparison is the RAM of a personal computer that
is heavily burdened by a program running in the background.

Firms on Wail Street that in the past invested in employee development
have been squeezed into “cleanups.” A trader is generally “as good as his last
trade.”

This explains partly the segregation between commoditized products and
complex derivatives in most firms. Newcomers present inexpensive sophistica-
tion but no trading experience. They, in turn will rapidly decay and will be re-
placed by the new PhD crop.

S, a thoughtful trader with some intellectual curiosity, once told the author
about his dilemma: taking a learning sabbatical and giving up his present job
with the possibility of not being able to return, or remaining on the job and be-
coming richer but obsolete.
[End OCR]*

## Page 68

*[Image OCR]
Market Making and Market Using 53

Linear (or quasi-linear) Derivatives. Traders hedging a forward or a swap
become subjected to maturity mismatches. Their hedge ratios, however, do
not change with time and pricing levels. They rarely incur transformation
in hedging parameters (assuming no meaningful convexity differences) ex-
cept for some extrinsic ones (e.g., correlation coefficients). A swap trader on
a deserted island can always summarize his position (but not his P/L),
whereas an option trader needs to check volatility, spot, interest rates, and
so on to know his.

The sensitivity of these portfolios remains constant across time or vari-
ation in the market prices. The fact that the hedges might not work in con-
sequence to yield curve shifts or other phenomena does not necessarily
make these derivatives nonlinear. The trading complexities reside in the
difficulties in gauging the hedge ratios and the relative volatilities and co-
variances of the different maturities. In addition, swaps are an area where
the accuracy of the pricing and interpolation are evident.

Nonlinear Derivatives. These are options-related instruments and fixed
income securities with high convexity that change in sensitivity with the
passage of time or with market levels (see Chapter 11 for a discussion).

Book running in these derivatives is less accommodating than that in
cash or linear products: The trader is ready to warehouse a trade for a given
amount of time, generally until expiration, against a fee. Because the hedge
ratios and the particularities of the instruments change according to time
or market movement, the book runner needs to engage in dynamic manage-
ment of the position. Novice option traders are taught that there is only one
way to hedge the sale of an option: by buying back the same one. The inher-
ent instability of derivative positions and their dependence on time and
market levels makes it clearly impossible to define oneself as completely
“flat,” except when the book is empty of trades, a rare occurrence.

Profitability

Derivatives market making in both standard and nonstandard products is
an attractive occupation because although every option is relatively illiquid,
the market as a whole for the Greeks is very liquid. Unlike a market maker
in an illiquid commodity that is uncorrelated to a more liquid one, the op-
tion market maker can easily find cheap alternative hedges.

The compensation for traders comes from offsetting the Greek risks
(delta, vega, gamma) more rapidly than predicted. That privilege arises be-
cause market makers have decided to run a book and incur the costs associ-
ated with the book management. The barriers to entry become their running
the book as an ongoing business. Option locals who warehouse strike risks
can easily find ways to reduce the variance of the portfolio with combina-
tions of available options.
[End OCR]*

## Page 69

*[Image OCR]
54 Markets, Instruments, People

PROPRIETARY DEPARTMENTS

The distinction between market makers and customers exists within most
trading firms and banks. The firms usually benefit from the “franchises”
or “customer flows,” two buzzwords connected with the book-running
function that represent its original mission of trying to service the cus-
tomer for a fee (except that here the service becomes the risk and the fee
becomes the bid/ask spread). Banks saw the success of their hedge fund
customers in the early 1990s and tried to replicate their customers’ suc-
cesses by creating internal hedge funds. Many failed and had to liquidate
rather rapidly, moving their orientation back toward their conventional
way of plying the book-running trade.

The reasons to start a proprietary function were obvious. They had an
overspill of back-office capability, which made the marginal cost of manu-
facturing trading profits very low. They were privy to very fresh informa-
tion on order flows into the market whether or not it turned out to be
relevant in trading. They also had, like everybody, excessive capital freed up
from conventional uses such as lending against the receivables of iron nails
manufacturing companies.

Their performance, however, did not match expectations for the follow-
ing reasons: They had a different utility curve than the funds, their execu-
tives and managers were operating under a different utility function than
the fund managers; and they had different expected returns (the funds were
riskier traders). The funds managers had a collection of funds that each pre-
sented the features of an option. Bank managers had one option on a collec-
tion of generally uncorrelated assets. According to the Basket Rule (to be
discussed later in this book), the expected returns from such a basket would
be considerably lower. In addition, banks sought “diversification,” some-
thing that seemed to be a good fashionable idea, but that lowered the price of
the manager’s option. .

Ina totally random but “fair” return (the expected profits are zero, a mar-
tingale, or “fair dice”), one option on a volatile instrument is more valuable or,
at worst, equal to a portfolio of options. Such a rule would apply even if the re-
turns were not a martingale, to a point. It resembles the pricing of an option
that is more sensitive to the volatility than the expected drift (Figure 3.3).

A trader with a “synthetic capital” of $10 million (i.e., the virtual capital
allocated to him by the bank) has often more earning power than a manager
with a higher amount of capital under him, because the manager will have °
to face the “diversification”—the offsetting returns from traders.

Example: A trader has $10 million “synthetic capital.” His yearly
“swing” is 25%: He will either make $2.5 million or lose $2.5 million.
His draw, in addition to his salary will contractually specify 20%. The
trader cannot be thrown out before the termination of the contract.
[End OCR]*

## Page 70

*[Image OCR]
Market Making and Market Using 55

December 31
P/L = $2.5 mil Expected Value
50% Good Year —-» Bonus = $.5 mil 50% of $.5 mil
A 250,000
January 1
A contract
$10 mil Capital P/L =($2.5 mil) Expected Value
Ne 50% Bad Year-————+» Bonus = 0 0
Total Option Present Value of
$250,000

Figure 3.3. Expected returns from a proprietary position: A mediocre trader.

The expected returns for the trader will increase with volatility.

In comparing the good trader to the bad trader, what determines the
price of the option is the volatility, not the expected returns from the
position. Of course, the example is simplistic: Bad traders do not have
easy access to capital; one needs (sometimes, but not always) something
that remotely resembles a track record (Figure 3.4). —

The manager in a financial institution, in contrast, runs a large
group of traders. He is paid 10% of a much larger pie: he will hence have
the illusion of partaking of a larger upside. Wrong. The net sum of
traders, particularly if they are not proficient, will be worth less than the
P/L of one person. What is the solution? Avoid diversification at all
costs. Hire one volatile trader, bad or good (the distinction is of small

2

December 31

P/L = $5 mil Expected Value
40% Good Year ——» Bonus = $1 mil 40% of $1 mil
400,000
January 1
A contract
$10 mil Capital P/L = ($5 mil) Expected Value
> 60% Bad Year» Bonus = 0 0
Total Option Present Value of
$400,000

Figure 3.4 Expected returns from a proprietary position: A bad trader.
[End OCR]*

## Page 71

*[Image OCR]
56 Markets, Instruments, People

relevance). If needs be, make sure ali the traders carry the same position,
without anyone offsetting the other. It is quite dispiriting for a manager
to see traders carry positions on opposite sides of the market: While one
of them will be profitable and would buy a new car, the net will lock ina
loss because of hedging costs. A firm cannot ask its traders to pay for
their losses unless they are losing back some of their profits. A deep-in-
the-money option resembles a future: A trader with profits loses the op-
tionality since his compensation would decrease with losses.

Compare that with a fund manager who has three independent
funds. He can draw 20% of each fund, regardless of the outcome of the
others. Unlike the manager of a trading desk, he benefits from the sum
of the three options. He can enhance his returns by being diversified
across separate funds, perhaps even negatively correlated between each
one of them. The only way a bank trading manager could achieve an
equivalent position would be to maintain three jobs simultaneously at
three different banks with a percentage payout in each.

This example shows why proprietary trading departments ended up clos-
ing down while fund managers thrived.

Tacit RULES IN MARKET MAKING

Market making in derivatives is subjected to a host of unwritten rules that
vary according to the market for each instrument. These rules evolve con-
tinuously as the market reaches a semblance of maturity. They are neces-
sary for dealers to coexist and correspond to the needs of an ambiguous
relationship between them: Dealers are at once accomplices and competi-
tors; they need each other as trading counterparties but compete for cus-
tomer business. Generally, if a trade is too large for a dealer, proper risk
management forces the dealer to split it with the rest of the market, for a fee.
The need to ensure liquidity and the adequate spread of information for one
another (without weakening any of the parties) leads to well-defined rules
of behavior.

A market maker in a cash product, for example, by asking another one
for a price, is giving up information about his position against the consum-
mation of a trade that would reduce that position. The trade-off will be fur-
ther enforced because he will now have an accomplice. .

Example: A large corporation buys 500 million British pounds from
Witibank. The Witibank traders, unwilling to hold such a large position,
“call out” for prices, contacting other dealers for £20 million each. Deal-
ers need to produce the price without knowing beforehand whether
[End OCR]*

## Page 72

*[Image OCR]
Market Making and Market Using 57

Witibank is a buyer or a seller. They would realize that Witi has a rela-
tively large amount to execute and know that by getting £20 million
they would be handicapped as the market might be flooded by the trade
for a certain period. However, they cannot escape their obligation to
quote Witibank as they would require the same favor should they be on
the other side quoting large amounts to their own customers. By trad-
ing, they will also gather a valuable piece of information, the suspicion
of a large order weighed against the possibility of a bluff. By trading
with Witibank they will then become partners in a secret.

Very liquid products have thus highly developed rules, down to how
long a person can wait before the quote becomes invalidated. Also, it is not
acceptable to come back to the same dealer after “passing” (avoiding the
trade) on a market, to quote too much without trading, or to pass on a
“choice” market; that is, when the counterpart shows the same price as a bid
and as an offer.

On the floor of the exchanges regulated by the Commodity Futures
Trading Commission, market making for futures and options is usually
subjected to the open outcry rules. The trader needs to continuously refresh
his bid or offer and specify a quantity for which his market is valid; other-
wise, he could be asked to honor his price for larger amounts than his capi-
tal would allow. Traders have no obligation to respond to the brokers nor to
other traders: They hawk their wares to generate business and create adver-
tising for themselves. ,

In the “upstairs” markets, there are subdivisions that depend on the
level of liquidity of the product. The more established and liquid ones, such
as short-term currency options, trade like cash, with mature rules and tacit
obligations between dealers; whereas structured notes are not an interbank
market and have no clear rules. A trader has neither the obligation—nor the
need—to quote them to another dealer. :

MARKET MAKING AND THE PRICE FOR IMMEDIACY

A theory by Sanford Grossman and Merton Miller? states that the market
‘maker function is to provide the price for immediacy. The price taker pays
the market maker in order not to have to wait and incur the risks that rise
with the square root of time. The price for immediacy is then reached as an
equilibrium between buyers and suppliers of immediacy.

Started by Garman in 1976, Market Microstructure is still a young branch
of financial theory. Most of its motivations are grounded on the awareness
. that prices do not reach equilibrium immediately in a big Walrassian auction
- meeting. Customers need market makers to fill gaps arising from imperfect
[End OCR]*

## Page 73

*[Image OCR]
58 Markets, Instruments, People

synchronization between buyers and sellers. Bid/offer spreads will then
depend on competition between market makers.’ Some of the theories de-
veloped will be explored later in the book.

MARKET MAKING AND
AUTOCORRELATION OF PRICE CHANGES

Briefly, a positive autocorrelation of a price change means that an up-move
is more likely than not to be followed by an up-move and a down-move by a
down-move. A negative autocorrelation means the opposite: The following
move is more likely to be of the opposite sign than the previous one.

The horizon of such effect also needs to be defined. A first-order auto-
correlation means that the effect is immediately a function of the move that
just happened. In simple terms, a higher order one means that there might
be a lag before such an effect is felt.

Many studies report the existence a first-order negative autocorrelation
of price changes. Such autocorrelation of prices reflects the “positive edge”
of market makers.* Most studies report an autocorrelation for price changes
when the observations are very frequent, as well as the vanishing of the cor-
relation when observations are separated by more than.a few minutes. The
most recent, Guillaume et al. (1995) reports a correlation between prices in
the dollar/DM currency pair extremely significant up to four minutes for
data observed between 1987 and 1993. Figure 3.5 shows the correlation of
the changes (the logs of the difference) and the rapidity with which such
autocorrelation washes out. Because the 4-minute price volatility is lower
than the bid/offer spread, it becomes conceivable that only a market maker
can capture such advantage. An average dollar/DM trader will execute a
trade every 18 seconds° during liquid times, which falls within the zone of
high negative autocorrelation. ‘

This first-order autocorrelation of prices (i.e., a short-term memory for
events preceding it) does not force us to depart from the notion of a “fair
dice,” or martingale, once we translate the difference into some positive
transaction costs for a market maker. Only a market maker can benefit from
a submartingale (that is a favorable dice) and in a limited way at that.

The issue of autocorrelation of returns will be further discussed in the
context of the variance ratios in Chapter 6.

MARKET MAKING AND THE ILLUSION OF PROFITABILITY
Market making in complex products can show immediate rewards and future

thorns. Mostly profits are shown after the trade because the mark-to-model
process generally implemented derives security prices off a mid-market.
[End OCR]*

## Page 74

*[Image OCR]
Market Making and Market Using 59

0.02
0.00
—0.02
—0.04
—0.06

—0.08

Autocorrelation

~-0.10
0.12
0.14

0.16
0 10 20 30 40 50 60

Time Lag (minutes)

Figure 3.5 Olsen & Associates autocorrelogram of changes, 1987-1993.

Traders usually find a fair value and mark the derivative up, then recognize
most or all of the difference between the fair value and the selling price as a
profit. This method does not allow for the future costs incurred in the man-
agement of an option book. The more complex the security, the more subse-
quent dynamic hedging will be required, and that in more than one market (a
cross-currency structured note needs to be readjusted in the Euro-Strips of
both currencies and in options on the cross currency, options need to be
rolled as they come close to expiration, etc.). Heavy slippage will be incurred
when the position attains some large size. All this would eat up the initial
profit, but would often be concealed from the trader who receives a steady
flow of such profit-booking trades. It would be difficult for the trader to es-
pouse a more aggressive stance and reduce transaction costs by observing the
different behavior of the multifaceted subcomponents: Market makers are
crippled by the size of their positions in relation to their appetite for risk.
Complex derivatives usually come in batches causing either the initial
firm or the competitors to sell more of a successful product. Ideas that sell
are rapidly imitated, and customers buy them with the comfort that oth-
ers are doing the same. That bandwagon makes all the incremental hedges
of the product more costly because their liquidity dries up. A costly exam-
ple is the Mexican peso range notes that make the issuing bank short out-
of-the-money puts on the currency. Their demand was widespread before
[End OCR]*

## Page 75

*[Image OCR]
60 Markets, Instruments, People

Option Wizard: Market Making Rie nme et ts

Option book runners constantly want to retire into proprietary trading depart-
ments. The strains of a position can be unrelenting, particularly when the
traders become driven by their books; they then manage positions so defen-
sively as to pay the price of liquidity for each trade. Typically, market makers
are given structures by their salesforce and asked to manage the risk and pro-
tect the “profits.” In distinguishing between primary and secondary trades
(the first being a trade that corresponds to the initiation of a position, the sec-
ond a defensive, dynamic hedge), accelerating the number of trades of the
second kind generally causes both fatigue and an erosion of profits. The un-
predictability of the results, added to the rarity of a vacation (usually heavy in
long-distance calls) causes the eventual combat fatigue.

While a hedge fund manager can exit all his positions with one phone call,
the option book runner can only put Novocain™ on the wound; uprooting the
problem is impossible. Personal stress usually results from the trader’s inability
to set his pace: the difference between the hunter and the prey. Most traders
are hunter types, possessing an aggressive streak that is hardly compatible
with the behavior of a prey.

Exchange traders usually do not suffer the same fate because they have a
better control over their inventory and the smail number of expirations in their
books. It is not uncommon to see exchange traders completely liquidating their
book before a telephone-free trekking trip to Nepal.

Another advantage to exchange traders is their freedom to make markets
(some exchanges, however, force them into a token obligation). They can
rapidly identify a regime shift and widen markets, something a trader at an in-
stitution cannot afford to do lest the customers call the boss. Typically, they
can refrain from showing one side of the market when markets heat up (as they
did during the stock market crashes of 1987 and 1989).

the 1995 devaluation and dealers found it difficult, even before the scare,
to repurchase the puts or hedge in the synthetic forward without heavy
costs.

ADVERSE SELECTION, SIGNALING, AND THE
Risk MANAGEMENT OF MARKET MAKERS

Adverse selection is a condition well known by insurance companies. They
have a tendency to get the less desirable policies since a dying person is
more likely to buy insurance than a healthy one. It is also difficult to distin-
guish between customers since a sick man can withhold information about
[End OCR]*

## Page 76

*[Image OCR]
Market Making and Market Using 61

his health. Likewise, traders have a tendency of moaning that they tend to
acquire the garbage from the market (bad distribution).

Traders experience another type of problem: a counterparty who buys
from them may have information they do not have. We will examine, later in
the chapter, the effects such information can have on fair-value pre-trade
and fair-value post-trade. Such signaling has been investigated extensively
in microeconomics and financial theory (see O’Hara, 1995, for a survey).
But unlike an insurance company, a trader might find the signal valuable. A
trader who traded with a central bank will certainly lose on that particular
trade. He will, however, end up winning if he uses the information to
“piggy back” the central bank.

Many traders compare market making to being short an option. Typi-
cally, leaving the price “out” guarantees the trader a trade if his bid becomes
highest or his offer lowest owing to market moves. Even though the market
maker's price is valid for only a few seconds, market makers can thus be ex-
posed to being “picked off” when information comes out during that time
and the market maker cannot officially change his quotes.

It is well known that market makers lose money chronically when mar-
kets experience shocks but produce profits in normal conditions when they
can earn the spread without much risk. In that respect, market makers ap-
pear to be short volatility.

Option Wizard: Comparative Lifestyles of Traders

X traded options “upstairs” for 10 years, mostly as a book runner or arbi-
trageur, and for a few years on the floors of the Chicago Mercantile Exchange
in currencies, SP500, and minor soft commodity contracts. Both positions had
their advantages. As an upstairs trader, he kept his mind occupied and was able
to watch a variety of markets, but he had no flexibility and, owing to the size of
his positions, was saddled with difficult (but profitable) positions all the time,
which prevented him from taking real vacations. Every day for him was like
maneuvering a large aircraft carrier, but his job was intellectually rewarding.

As a local, he felt the exhilaration of freedom but was physically limited
with markets and instruments. It felt like a blue-collar job but with a higher
payout. In addition, he no longer got the respect large traders usually com-
mand: Brokers who in the past had toadied up to him now hardly gave him the
time of day-—not a great loss as he had matured by then.

The best lifestyle X ever experienced was that of a carefree wind-surfing
and bicycling fanatic who “scalped” options in the pit only on days when the
wind was unaccommodating. X closed his positions in the winter and drove
from Chicago to California. What was surprising was that the sports bum was,
besides, an excellent pit trader.
[End OCR]*

## Page 77

*[Image OCR]
62 Markets, Instruments, People

What considerably mitigates the risks of a market maker is the manage-
ment of the liquidity of the market available to book runners. They typi-
cally have enough outlays to offset the risks so that being stuck with a
position for a long duration becomes exceptional.

VALUE TRADING VERSUS THE
GREATER FOOL THEORY

There is a running dichotomy among option traders whether it is preferable
to warehouse a trade because of its intrinsic attractiveness or to trade ac-
cording to other people’s desires. Here are the two parties:

1. The school of value trading favors only accumulating positions that,
warehoused until expiration, represent a positive expected return. The
traders should be capable of sustaining a lack of trading for an extensive pe-
riod. As these value trades generally go against the market, the trader needs
to have a certain measure of indifference to marks-to-market losses (e.g.,
buying a “cheap” calendar spread where the back month is heavily dis-
counted owing to structural supply; buying out-of-the-money calls in equity
markets owing to the strong discount brought about by the oversupply of
covered-writes).

2. The school of the greater fool theory suggests a trading attitude
more accommodative to the marketplace. It reflects the belief that traders
are conduits. Value resides as much in the discounting of the final payoff (a
net payoff of the trade minus the cost of hedging a book) as in the possibil-
ity that someone may “pay up” for the option. A trader would not mind
holding an expensive option provided that a greater fool could come to re-
lieve him from it. Playing greater fool is often indispensable with market
makers of liquid products owing to their large volume and frequency of
trades. Because market makers hardly ever carry trades for a significant pe-
riod, they find it interesting to value options in terms of their extrinsic
rather than their intrinsic value.

This theory is best described in Schwager’s New Market Wizards (1992),
where Jeff Yass, an outstanding market maker, describes his approach to
option trading in terms of conditional probability. Somehow “value” is con-
ditional on the most recent information rather than the overall picture. If a
market operator is given a position by someone else, he receives more than a
position: He gets information, and that information would necessarily erad-
icate the edge. So there are two values: (a) value that was there before the
trade (prior value) and (b) value that is conditional on the trade, which will
be necessarily higher (if the operator is short) or lower (if the operator is
long) than the original one (post value).
[End OCR]*

## Page 78

*[Image OCR]
Market Making and Market Using 63

Table 3.2. Value Trading versus Greater Fool Methods

Value Trading

Only inventory net option positions

that harbor a positive expected return.

Focuses on intrinsic value.

Fund manager’s approach: slow
volume, smal! turnover.

A market maker with new products or
a fund manager with liquid products.

Time is expected to work in favor of
the position, as the value is amortized
as expiration nears.

The positions do not “mark” well, that
is typically the initial mark-to-market
works against the trader.

Greater Fool Trading

Only buy the options that might turn
out to be in great demand. Only sell
the options that someone else might
want to get rid of.

Focuses on extrinsic value, the signal
from the market.

Market maker’s approach: high volume
and turnover.

A market maker with liquid products.

Time works against the position.
Sometimes the trader is long expensive
options or short cheap options and the
gamma differential could be heavy (see
topic).

Positions mark generally well.

Table 3.2 lists the attributes of the two methods.

Example: Two option traders would eventually make money on two

diametrically opposite trades, provided they remained consistent with
their trading styles. The first one, a value trader, would go against the
skew in biased assets markets by buying calls and selling puts for a
volatility differential that is not justified by the real skew of the distri-
bution. However, he needs to hedge his deltas as if he were going to
stick to his trade until expiration, without changes in the process that
would make him sell futures in dips as if volatility were higher.

The second one would buy puts and sell calls, knowing that the dif-
ference is not warranted by the distribution. He is aware, however, that
someone along the way is very likely to panic at the first dip and buy the
expensive puts from him at a higher price. He also can often detect a
massive anxiety in the market once it starts to feel heavy. His danger,
however, is that warehousing such a trade until expiration would cause
him to lose money.

“An expensive option does not come alone,” the saying goes. When
an option trader sells an option in the market, he would tend to find
that the series would become more expensive, owing to the structure of
[End OCR]*

## Page 79

*[Image OCR]
64 Markets, Instruments, People

imbalances. The best way for greater fool traders to exit is for them to
see the series of options start softening up.

MONKEYS ON A TYPEWRITER

The Statistical Value of Track Records

There is an application of a mathematical lemma called the Borel-Cantelli
Lemma’: If one puts an infinite number of monkeys in front of (strongly built)
typewriters and lets them clap away (without destroying the machinery),
there is a certainty that one of them will come out with an exact version of
the Iliad. Once that hero among monkeys is found, would any reader invest
his life’s savings on a bet that the monkey would write the Odyssey next?

The same applies to traders. There is tendency on the part of people with
a small knowledge of probability laws to select proprietary traders and fund
managers based on the following principle: It is very unlikely for a trader to
be up money ina consistent fashion without his doing something right. Track
records, therefore, become preeminent. So they call on the rule of the likeli-
hood of such a successful run and tell themselves that if someone performed
better than the rest in the past then he has a great chance of performing bet-
ter than the crowd in future days. I have never seen anyone count the mon-
keys. Nobody counts the traders in the market to calculate—instead of the
probability of success—the conditional probability of successful runs given
the number of traders in operation over a finite market history.

One of the key properties of probability laws is that they are counterin-
tuitive. People look at a casino seeing how a small “edge” translates into
virtually certain profits and infer that the same rules apply to a man rolling
the dice. The sums of random walks (net profits from the sum of roulette ta-
bles) act differently from the final sum of a random walk (a gambler’s net
results at the end of one session on one roulette table). A modest advantage
on a sum of random walks translates into certain profits. The same advan-
tage of one roulette table is generally drowned by the volatility of the ran-
dom walk itself. The reader could test the difference by generating series of
tosses on a spreadsheet using a random number generator and looking at
the discrepancies in the results. Even with a small advantage, a counterin-
tuitively high number of negative runs would mar the party. The only way
to reduce the frequency of negative runs would be to increase the number of
trades and make the bets smaller.

One could think of a roulette table and view the edge of a trader who has
negative odds with a system that gives him a 45% chance of winning. What is
the probability of the person betting $1 being in the black after 30 throws?
Using the cumulative binomial] distribution, he has at least a 35% chance of

«
[End OCR]*

## Page 80

*[Image OCR]
Market Making and Market Using 65

ending up profitable. If the gambler sliced the bet into pieces of $0.10 and
compensated with 10 times more bets, the final number of positive results
would be 4.6%. This provides a vivid illustration of the reasons to use caution
when considering the notions of “edge” and “skilled traders.” In a long Wall
Street career, I have often heard the adage: “You are as good as your last
trade.” Often the behavior of the managers of traders is dictated by their lat-
est run. Investors somehow are not as easily fooled.”

More Modern Methods of Monitoring Traders

After the ravages of February 1994, people wised up to statistics and started
examining new methods to distinguish the truly skilled trader from the
lucky monkey on a typewriter. Bill Fung,’ a consultant to hedge funds, in-
troduced a nonparametric method of assessing the relative performance of
a trader to the market. It takes into account that some markets trend and
(remember 1993) that some traders were simply lucky to be long those secu-
tities that happened to trend. In addition, he monitors the transaction fre-
quency to ascertain the fit between the trader’s intention and the actual
events in the market: The higher the frequency of trades, the higher the sta-
tistical significance of the trader’s performance. This method is grounded
on the same awareness of the fit of the distribution as the one described in
the dependence of a dynamic hedger on the transaction frequency. An op-
tion trader who frequently rebalances his book would capture the true dis-
tribution with more certainty (for the better or the: worse) than his
colleagues who hedges himself only occasionally.

Where the fund manager “sits” on a long or short position, therefore de-
pending entirely on the underlying market, the investor can save himself
considerable fees, a cut in the expected return, and some headaches, by buy-
ing a liquid future on the commodity itself and completely bypassing the
fund manager.

The Fair Dice and the Dubins-Savage Optimal Strategy

How aggressive a trader needs to be depends highly on his edge, or ex-
pected return from the game:?

¢ When the edge is positive,’° (the trader has a positive expected re-
turn from the game, as is the case with most market makers), it is al-
ways best to take the minimum amount of risk and let central limit
slowly push the position into profitability. That is the recommended
method for market makers to progressively increase the stakes, in
proportion to the accumulated profits. In probability terms, it is bet-
ter to minimize the volatility to cash-in on the drift.
[End OCR]*

## Page 81

*[Image OCR]
66 Markets, Instruments, People

¢ When the edge is negative,'! it is best to be exposed as little as possi-
ble to the negative drift. The operator should optimize by taking as
much risk as possible. Betting small would ensure a slow and certain
death by letting central limit catch up on him.

¢ Atany rate, when the trader owns an option on his profits (i.e., gets a
cut of the upside without having to share into the downside), as is
the case with traders using other people’s money, it is always optimal
to take as much risk as possible. An option is worth the most when
the volatility is highest.

The ArcSine Law of the P/L

Another counterintuitive phenomenon affecting traders concerns the dis-
tribution of their P/L. One would think that in a fair game (a totally ran-
dom P/L), the most likely time spent in the “black” or in the “red” would
hover around 6-month per annum. In reality, the probability of spend-
ing 1 month or 11 months in either is markedly higher. As a matter of fact,
the probability of spending 6 months in the black or in the red is the low-
est of all!

That striking law is called the ArcSine law for a random walk” and af-
fects us in more than one instance. It also applies to the distribution of the
maximum and the minimum and will be covered in Chapter 23.

Figure 3.6 shows the expected amount of time a “fair” trader is expected
to spend in the black (or the red). Notice that the drift does not alter the dis-
tribution until it becomes very strong. The arcsine law makes commonly used
methods of comparing the cumulative performance of traders erroneous. It
shows how unjustified most of the conventional measurements can be.

Risk Management Rules

¢ When judging a nonmarket-making trader, remember that a
large share of his profitability is attributable to luck.
The ratio of luck to skills decreases with the transaction fre-
quency. For a proprietary trader, it is very high, for a market
maker it will be very low.

In general, traders become extremely arrogant and difficult to
manage when they are very profitable. A trading manager must
be able to face wrongly profitable traders without being influ-
enced by their P/L.

Do not use Sharpe ratios with nonlinear products, particularly
options, or with linear products traded in a nonlinear manner
(such as a stop loss)."
[End OCR]*

## Page 82

*[Image OCR]
Market Making and Market Using 67

Probability
density

Months

2 4 6 8 10 12

Figure 3.6 Arcsine law for the P/L.

We end on this final problem: A trader is given a stop-loss of $100,000 in any
given month (i.e., he would have to close the books and go home until the
end of the month). Another trader is allowed $200,000. Assuming that both
traded the same size in one uniformly volatile market (constant volatility)
and that their profits and losses were (as with most traders) random, which
one has a higher expected return?

Answer: They have the same expected profitability. They share the same
expected returns but the trader with the smaller stop loss will incur more
frequent losses. His P/L will have a strong, positive, “skew”: frequent small

losses and infrequent large gains.
[End OCR]*

## Page 83

*[Image OCR]
Chapter 4

Liquidity and Liquidity Holes

The market is a large movie theater with a small door.
An options veteran

This chapter should be read prior to studying barrier options.

LIQUIDITY

Liquidity! in a market is defined by the ease with which an operator can
enter and exit it for a given block of securities.

If one were to summarize what trading (as opposed to investing) is

about, the best answer would be adequate management (and understand-
ing) of liquidity. Liquidity is the source of everything related to markets.

Slippage is a practitioner’s measurement for liquidity. It is computed
(for a given quantity to execute) by taking the variation between the av-
erage execution price and the initial middle point of the bid and the
offer. Slippage is not always a precise measure of liquidity for a particu-
lar commodity, but it provides a reliable comparative measurement of
liquidity between markets.

Academics offer a number of measurements for liquidity, mostly using

the numbers of transactions or the bid /offer spread.

68

Example. Slippage: A fund manager needs to go long the JPY currency
against the dollar. The 115 calls on the Chicago Mercantile Exchange are
quoted 88/92 which means that the middle market in theory is at 90. He
would then assume that for buying 4,000 contracts, he would have to
pay 92 for the first 1,000, 93 for the next 1,000, and up to 98 for the bal-
ance of 2,000 as he would drive the currency itself higher thanks to op-
tion traders hedging their deltas. His weighted average will then be
95.25 and he will count an overall slippage of 5.25 ticks for his execu
tion. He should make the same allowance for his unwinding the trade,
provided he picks the same time of the day. Slippage would be more im-
portant when the market gets more volatile (which is when the fund
[End OCR]*

## Page 84

*[Image OCR]
Liquidity and Liquidity Holes 69

manager might find it more interesting to get involved or to unwind the
position) or when the overlapping of time zones reduces the total num-
ber of market participants.

Slippage is the principal reason for leverage hedge fund traders to limit
the amount of capital under management. While their overall dollar returns
might increase (in the event of their being profitable), the percentage re-
turns decline accordingly.

Ironically, slippage costs often disappear with large blocks: Many large
traders can use their buying power to prop up the market in which they ac-
cumulate a position. Often the market rallies and stabilizes at the price at
which the last buy trade was executed. It is in exiting the trade that slippage
costs will invariably be incurred.

LiouIpiIry HOLES

@ A liquidity hole or a black hole is a temporary event in the market that
suspends the regular mechanics of equilibrium attainment. It is an in-
formational glitch in the mechanism of free markets, one that can cause
considerable damage to firms. In practice, it can be seen when lower
prices bring accelerated supply and higher prices accelerated demand.

Liquidity holes are attributable to the way information initially affects a
market. Typically, liquidity holes occur when operators are aware of a major
piece of information (an event, or a size order in the market), but cannot
gauge its size and possible impact. Information causes anxiety, conspiracy
theory, and price conflicts. Most often, operators need to interpret informa-
tion. Markets are supposed to move to the extent of the difference between
news and initial expectations but the latter ate usually unknown and diffi-
cult to estimate precisely.

A political announcement or the release of an economic figure can cre-
ate a liquidity hole by leading to disruption in the normal price determina-
tion process.

When a stop-loss is executed, as in the following example, operators
freeze as they become aware of the order without further information as to
its magnitude. Many suspend their trading temporarily, thus causing the
market to gyrate even further.

Liquidity holes would not be very dangerous except that with the large
open interest in nonlinear derivatives, some operators have large contin-
gent orders that need to be executed regardless of the market makers’
spreads.

Often barrier options (knock-in, knock-out) are to blame for the liquid-
ity holes and the intraday volatility that results from them.
[End OCR]*

## Page 85

*[Image OCR]
70 Markets, Instruments, People

LIQUIDITY AND Risk MANAGEMENT

It cannot be stressed enough that liquidity is the most serious risk manage-
ment problem. A substantial part of unforeseen losses is due either to mar-
ket jumps caused by illiquidity or to liquidation costs that substantially
move the markets against one’s position. Liquidation costs tend to be usu-
ally underestimated since operators usually “fade” when someone is forced
into a market action.

Another element to consider in the value-at-risk (see Module E) method
is that most simulations exclude the liquidation costs at the stopping bar-
rier. The market is merciless with operators who start closing down a posi-
tion, particularly when the liquidating party has no choice. Since forced
liquidations take place in markets under duress, one can imagine the effects
of reduced liquidity on the unwiding of a portfolio.

STOP ORDERS AND THE PATH OF ILLIQUIDITY

M Stop-loss orders are instructions to buy or sell a given quantity when
the price prints on the screen, unconditional of the filling price. Buy-
stop orders are generally higher than the market and sell-stop orders
are lower than the market.

® Limit orders are instructions to buy or sell a given quantity at a prede-
termined price in the market. Market rules guarantee that no price
would trade “through” the order without the fill being performed.

These two basic definitions will be illustrated with examples later in
the chapter.

Trying to measure the markets in accordance with their volume is a de-
ceiving enterprise. Markets are not linear and do not reach equilibrium in
an orderly way. Information can be nefarious as well as helpful. The cer-
tainty of a path should certain conditions be satisfied (a conditional path)
can cause such gaps.

The mechanics of a stop-loss can shed some light on the issue of path-
certainty. Stop-losses in cash products are buy or sell orders triggered by
the market reaching a certain level. As such they resemble trigger options.
A stop-loss in the SP500 to buy 500 lots at 624.00 will be triggered by the
broker being obligated to start buying that amount at any price should there
be an official print at that level. Customers usually leave a stop in the mar-
ketplace to protect profits, limit losses, or sometimes to enter the market if
some magic level is attained. Most trainees are obligated to leave a stop for
every open position as their bosses fear their breach of discipline.

Since stop losses guarantee an order, it becomes understandable
that some operators can have an incentive to trigger the large stops. The
[End OCR]*

## Page 86

*[Image OCR]
Liquidity and Liquidity Holes 71

following example shows the execution of a stop-loss in a relatively less
liquid market.

Example: A large European fund manager gives a Canadian bank a
stop-loss for $300 million against the Canadian dollar before going
home, at noon New York time. The Canadian dollar presently trades at
1.38 against the USD and he diligently places his stop at 1.3750, a level
satisfactorily far given the low volatility of the Canadian currency. The
trader in possession of the order knows that should the dollar drop he
might have an interesting asset in his hands.

Nothing happens for a few hours until the New York trader notices
that the Canadian dollar has dropped to 1.3780. He then makes the
judgment that he could push the market toward the stop, provided he
runs into no real buy order anywhere. He can then sell a large block of
dollars given the certainty that he can buy them back from his own
order at 1.3750. He will then as clumsily as possible sell $300 million,
hoping the market remains vulnerable, until the customer gets filled in
his stop-loss. His aggressive selling, instead of causing more buyers to
come into existence, as implied by normal market dynamics, will make
operators stay on the sidelines because of the suspicion of a piece of in-
formation they do not have. Figure 4.1 depicts the liquidity hole associ-
ated with this stop-loss.

Most FX traders are
at the bar
1.38 ~™
Fee | signs
1.378 Fee |
Trader Ne cheamer | hits
bids “for Ne cheamer |

No buyers MF re |
LIQUIDITY
HOLE
Eee | 37 5()

| mee | sell order
| mee | [Heavy selling |

Order filled: trader buys from
customer at 1.3725 to cover
what he sold between 1.3780
and 1.375

1.3725

Figure 4.1 Liquidity hole.
[End OCR]*

## Page 87

*[Image OCR]
72 Markets, Instruments, People

Most traders on exchanges operate under the same principle: that of
contingent forced large trades. They do not need the exact information: It is
easy to sense where the stops are located from the behavior of the brokers
as their levels get closer.

BARRIER OPTIONS AND THE LIQUIDITY VACUUM

Much of the recent market volatility was caused by the triggering of barrier
options. As these are starting to abound, so are unexplainable gaps in the
markets beginning to appear. In the previous example, we saw that a large
stop was a free option. However, a stop can be canceled, or given to the
trader too late for him to be able to utilize it properly. A trigger option will
offer none of these mitigating features: It will be there offered to the market
maker far in advance and by its nature is noncancelable.

The option trader with a knock-out on his books will place an order to
unwind his deltas, conditional upon the termination of an option. He then
would not mind being filled on the stop-loss, as the small casualty would be
offset by terminating a liability (the short option) ahead of schedule. He
would also benefit from the execution of his own stop order, if it is handled

properly.

Example: Sterling is trading at 1.6000 USD/GBP (one needs 1.60 dollars
to buy one pound). A trader is short the 1.60 Sterling calls for a 3-month
expiration with a knock-out at 1.59, in 200 million pounds (equivalent to
US$320 million). He knows that should 1.5900 print and trade in one or
more places on the street his option would be terminated. The option
delta is 78%, and he holds 156 million British pounds against it.

The option trader makes a quick judgment as to what it would take
to make the currency drop to below 1.5900. The stakes for him are large
and he has an incentive to trigger 1.5900. In addition to losing his op-
tion, he could make some money in the process, by playing with the
stop-loss information, as in the previous example, except that the stop-
loss would be his own.

He faces two options. He could “sell” his order to the spot desk
against .10 pips ($200,000) and ask the desk to manipulate the markets
for him. He could also trigger his own stop. He would then wait to see
the first down-move in sterling. Stops are more easily triggered after
the Bank of England (also called the “old lady”) goes home, in the New
York afternoon. The best time would be after New York closes and the
market rotates to Sydney, Australia, and Wellington, New Zealand.
However, the sterling becomes illiquid enough for that purpose after
Europe’s close.

The trader would sell most of his hedge as the sterling drops: 30 at
1.5945, to soften the market, then 20 at 1.5930, 25 at 1.5920, 20 at 1.5910.
[End OCR]*

## Page 88

*[Image OCR]
Liquidity and Liquidity Holes 73

The current total is 95 (out of the 156 to unwind) and he could be in
trouble if the market rallies. It would perhaps be better if he put pres-
sure on the market with the remaining 65 by aggressively offering the
pounds. He would sell 30 at 1.5900 and would prefer.to see 1.5890 trade
to ensure that the trigger fill cannot be contested. He would then sell
the balance at 1.5895 and 1.5890 by breaking them up between several
counterparties to make his fill incontrovertible. Some counterparties
could legally contest the trigger if there is no sufficient proof of trading.
By consummating the trade with two reputable firms, he can defend
himself against such allegations.

This influence on the path of the commodity reduces the Markovian
component of the Brownian motion: The process now has a memory.
Barriers and stops can cause the next price movement to depend on the
previous one. Down moves away from 1.60 led to an acceleration.

A similar, but milder effect took place during the portfolio insurance fi-
asco: Locals detected a large player whose next move would depend on the
past market movement and benefited from it.

OnE-Way Liguipiry TRAPS

Many markets, particularly biased agsets (markets subject to frequent pan-
ics), can present an asymmetrical behavior of prices. Operators experience
a deceiving liquidity on the way in, and discover to their sorrow that get-
ting out is another matter, particularly when liquidity holes occur. Traders
experience such imbalance with out-of-the-money puts on biased assets
that are in chronic shortage (see Chapter 15).

Hotes, BLACK-SCHOLES, AND THE ILLS OF MEMORY

The principle behind the arbitrage derivation method in modern option pric-
ing theory is the memoryless Brownian motion, with no remembrance of
where the spot came from to the last price. This assumption seems to be
under the onslaught of all manner of financial theorists and traders as new
theories about some form of dependence between past information and mar-
ket price action seize the day.” Without getting into the chaos theory fad,
there seem to be very respectable theories based on nonlinear dependence
heretofore undetected by the conventional statistical machinery.* Markets
may, after all, have some form of memory.

At the option trader level, market efficiency does not appear to become
attainable despite the increases in the traded volume of most instruments.
What is relevant is that the liquidity holes display the weakness of the
Markov‘ assumption as the path followed to the last price becomes quite
[End OCR]*

## Page 89

*[Image OCR]
74 Markets, Instruments, People

meaningful. Conditional upon the market touching some level, there will be a pre-
dictable price action by some barrier option liquidator that will precipitate the mar-
ket in some known direction.

Even the most liquid markets display acute weaknesses at the most liq-
uid of times. It is not fully explainable why a market that moves one trillion
a day could be influenced by an order that represents 0.2% of the total vol-
ume. It is also unexplainable why the experience of liquidity holes has not
provided immunity as the market would use the holes as known events and
would adjust for their occurrence through adequate expectation.

A fact that shows why markets have memory is the noise level and frenzy
activity that surrounds a new low or new high in any market, regardless of
the actual volatility that led there. Markets hitting new highs and new lows
start moving in gaps. Often stop-losses need to be taken care of.

LIMITS AND MARKET FAILURES

The exchanges seem to understand that free markets exhibit failures and
have stepped in to suspend trading after a certain accumulation of events.
The “circuit breakers” in the SP500 (the mother of all liquidity holes) stop
the market for five minutes if it opens within a certain range, then gradu-
ally establish wider limits of trading. The circuit breakers appear to have
been effective since the 1987 stock market crash, at least as a psychological
buffer and a means to allow traders to take a deep breath and gauge the
available information before the resumption of trading. They have been ac-
tivated routinely and appear to be a successful experiment.

REVERSE SLIPPAGE

Some traders are expert at quietly accumulating securities, causing a short-
age; then noticing that the balance or the order flow has shifted to their ad-
vantage, they can cause the market to move their way. They incur positive
slippage: ironically, it can be profitable to execute a large order in an illiquid
market and unprofitable to execute a smaller one.

Reverse slippage takes place when the accumulation is done at continu-
ously higher prices in the event of a buying splurge, and at lower prices in the
event of a selling one. Every day, the marks-to-market of the accumulating en-
tity will be improved during that phase. Typically, the accumulating party
“digs” the market at the end of the period, by starting to buy clumsily and
causing the prices to move higher in such a way as to bring stop-loss orders
into play.

However, the dangers are mostly in getting out of the trade. The trader
cannot engage in the same pushing of a market as he did during the position
[End OCR]*

## Page 90

*[Image OCR]
Liquidity and Liquidity Holes 75

building phase. He would be causing his own collapse. In illiquid markets,
particularly exchanges, the activity will be detected rapidly. Traders can tell
the imprint of a source of activity no matter how well disguised. They will be
merciless with a trader getting out as they have an idea of his size. An un-
winding party does not have the freedom to stop midway as an accumulating
one does.

Reverse slippage has been profitable for fund managers who created
squeezes large enough to download a large portion of their entire inventory
at one price.

LIQUIDITY AND TRIPLE WITCHING Hour

The mechanics of stock index settlement can lead to much abuse by special-
ists and stock market makers who can thus induce a temporary gap in the
market and create a liquidity hole to their advantage.

Example: A basket arbitrageur holds cash and shorts SP500 futures as
a full hedge. The contract will be cash settled (the trader gets a compen-
sation of the difference between the previous day’s settlement and the
official closing price). But after the settlement, the trader will still have
to close his long cash position. He will need to secure a bid not far from
his settlement price for the arbitrage to be effective.

The settlement terms are that the index will be computed off the of-
ficial close of all the components. So the best solution for the program
trader is to unwind the stocks at the “market on close” feature in the
marketplace. Therefore, the tacit unspoken contract he can establish
with the stock specialists is, “I do not care where I get my fill as long as
that number becomes the one used to compute the price of the index.”

The specialists will then use their time and place advantage to move
the market to where the imbalances are located. Small imbalances
caused by a slight excess of sellers over buyers would bring about a “sell
expiration.” The specialist knows that there is no chance in the small
time imparted to see a buyer come bid better than him, and he would
therefore obtain the stock at a very attractive price. The entire operation
takes place very quickly. What is remarkable is that it has been taking
place so long without efficient markets adjusting to the pattern.

PORTFOLIO INSURANCE
A major illustration of a liquidity hole caused by forced trading is the port-

folio insurance episode.® It illustrates the case of a panic in that a program
that covered about 75 to 100 billion dollars, less than 3% of assets under
[End OCR]*

## Page 91

*[Image OCR]
76 Markets, Instruments, People

management, could cause a liquidity hole of such a massive scale. That port-
folio insurance caused the crash remains disputable; that it exacerbated the
market movement is a certainty.

Portfolio insurance was initially marketed as a protection against ad-
verse market movements for asset holders through a dynamic option repli-
cation technique. The absence in liquidity in long-term options justified
the technique of replication in place of simple option buying. Human na-
ture is such that a potentially free lunch appears easier to sell to a fund
manager than a real option. The selling technique can reach a more gener-
alized level as shown in the Option Wizard, How They Fool Their Customers
(Chapter 20).

As a method of duplicating the put through delta hedging, portfolio in-
surance was a negative gamma trade. It could have been argued that, even if
portfolio insurance had been executed through actual options, the market
makers would have been short the same amount of gamma thus exacerbat-
ing the movement with their negative curvature. The answer is that the
“free” feature of the option attracted capital to the asset itself thus causing
people to buy the stock in larger quantities than if the program did not
exist. For a discussion of the difference between a true option and its dy-
namically replicated equivalent see Grossman (1988).’

This author spent one year in the SP500 pit after the crash trying to re-
constitute the episode of October 1987, as a lesson in liquidity management.
To understand market-making dynamics, every trader should spend a few
months in such a pit. At first, the tumultuous pace conceals. the real activity.
After a while, some order starts to emanate from the pit’s behavior. That is
when the person becomes part of this large body that is a marketplace.

My conclusion is that locals detected a pattern of forced selling as a re-
sult of lower prices in the market. It was easily detectable that the freebie
called a stop-loss was available in such a manner. The traders rapidly real-
ized that, as in a stop-loss, the decision to sell was not relieved by a lower
market price. In fact, it turned out to be compounded by the lower prices.
The classical “front running” of the stop-loss order was made easy because
only one forum was available for the program and that most of it was con-
centrated among few brokers.

A sixth sense, or what locals call “pit chemistry,” allows them to iden-
tify the origin of a market order without much effort. Living in the midst of
a centralized marketplace, they can identify patterns and recognize the sig-
nature of every player.

It has been argued that portfolio insurance was the result of the worst
possible case of market equilibrium attainment: imperfect information. A
large measure of anxiety was caused by the lack of knowledge about infor-
mation in the markét. It is difficult for a market maker to show a bid for a
future without the existence of order in the marketplace. One can also
[End OCR]*

## Page 92

*[Image OCR]
Liquidity and Liquidity Holes 77

conclude that portfolio insurance became the hostage to its own initial
success: smaller amounts under the program would not have caused such
a snowballing effect.

LIQUIDITY AND OPTION PRICING

There has been a growing awareness in the option pricing literature about
the existence of transaction costs. Several option theories have been devel-
oped in an attempt to conciliate market microstructure theories and option
pricing. Leland (1985) introduced the notion of a break-even volatility in-
corporating an adjustment for transaction costs. Such adjustment, which he
calls A, is added to the volatility entered into the risk-neutral portfolio
replication of the Black-Scholes-Merton equation. The resulting break-even
volatility for a short seller is the augmented volatility:

o,=o0V(1 + A)

[2 k
A= |—*xX
7 oVit

o represents the volatility (vanilla volatility), k is the round-trip transaction
costs (in percentage), and 8t is the time interval between successive adjust-
ments of the portfolio.

However, the trader needs to take into account that an option book will
not be monotonically long or monotonically short volatility, but that the
break-even volatility will be a markup or markdown according to his posi-
tion. The short premium player needs a higher volatility to break even
while the long player needs a lower volatility. The short gamma pays the
negative spread with stop-losses, while the long premium operator has
_ stop-profits that could absorb some profits.

Whalley and Wilmott (1993) make the resulting break-even volatility a
simple function of the sign of the gamma. As earlier, it will be oV(1 + A)
_ when the operator is negative gamma and oV(1 — A) when the operator
is positive gamma. Hence o({[) = oV(1 + A sign(I)). This creates a situation
of variable volatility with the resulting complexities.
Perhaps the area in option research that will witness the largest growth
will be the dynamic analysis of option risks under variable transaction
- costs. Replication costs for an option market maker depend on more factors
than the sign of the gamma as with Whalley and Wilmott (1993).
In dealing with transaction costs, option traders have massive economies
of scale because the total net gamma in their portfolio sometimes will be equal

where
[End OCR]*

## Page 93

*[Image OCR]
78 Markets, Instruments, People

to several thousand times the aggregate gross gamma in absolute value (ie.,
the sum of the gamma in absolute value for every option in the portfolio). This
makes the management costs of any particular option minuscule.

Another factor, more difficult to assess, is that of the asymmetry in op-
tion costs. The option seller and the option buyer are not quite mirror im-
ages, given the difference in utility function between the two.

Risk Management Rule: Option trader lore states that when long

gamma, use limit orders. When short gamma, use stop orders.

An option trader who “sits on the bid or the offer” when short gamma is
said to be penny-wise and pound-foolish. He would later have to “chase”
the market. Conversely, an option trader who pays spreads when long
gamma is fooling himself. “Let the market come to you” is the experienced
trader’s recommendation. Table 4.1 shows the characteristics of short and
long gamma.

An option seller has no choice whether to hedge or not, hence all of his
adjustments need to be done while incurring transaction costs, regardless of
his status as a market maker. A short gamma operator needs to adjust his
gamma using stop orders, which are triggered with the printing of a set price
in the marketplace. A stop-loss with the SP500 trading at 455.00 will cause
the order filler to buy at best in the marketplace once 455.00 officially trades.
The stop will be filled at 455.00 plus bid/offer spread, hence transaction
costs will be incurred. In addition, the bid/offer spread typically widens
when the market moves, and such moves generally accompany the need to
rebalance.

Table 4.1 Execution of Secondary Trades

Long Gamma Short Gamma

Buy Orders Can buy when the market goes | Needs to buy when the market
lower goes higher
Posts a bid in the market that Leaves a stop order that
can guarantee a trade if the becomes a market order
market goes through his level triggered by the official printing
of a given price
Sell Orders Cansell when the market goes — Needs to sell when the market

higher goes lower

Posts an offer in the market
that can guarantee a trade if the
market goes through his level

Leaves a stop order that
becomes a market order
triggered by the official printing
of a given price
[End OCR]*

## Page 94

*[Image OCR]
Liquidity and Liquidity Holes 79

An option buyer incurs less transaction costs, if any. A seasoned long
gamma player will have orders that go against the market and will know
to place them in the marketplace instead of paying the spread unnecessar-
ily. On an official exchange, the order cannot be violated by the market,
which means that the market cannot trade through the order without a fill
being guaranteed. It is also the case for orders put through the over-the-
counter broker in currencies. Some option traders (including the author)
who dislike being short gamma advocate that a skillful long premium
trader will in fact earn the transaction costs, as he can become a secondary
market maker in the cash market. Assume that, as in the preceding case,
the SP500 trades at 449.80. The option market maker being long gamma
can leave an order at 450.00. The market would not be able to trade without
his order being filled. Assume that the market without him is 449.50,
450.05. By offering 450.00, he then can hope to earn part of the spread.
True, the market makers would be able to match, but a certain measure of
pit etiquette gives the “paper” the right to get a fill along with the market
makers. In a marketplace like the Chicago pits, the brokers are powerful
and can command better fills than locals. So the transaction costs are lim-
ited to the commission, a negligible cost for the large players.

At a more advanced level, every known market experiences the negative
autocorrelation of prices for short-term observations, which can also be inter-
preted as mean reversion (see correllogram in Figure 3.3). That effect should
be considered in any discussion about implicit bid-offer spread,’ which is
wider than the nominal spread as the latter can only satisfy small transac-
tions. In some less liquid markets, such as the individual equities, the explicit
spread is wider than the implicit spread since the market makers display a
market wider than their intentions to trade given that their posted price rep-
resents a short option. This mean-reversion will cause the market makers to
derive an “edge” and by extension allows skillful positive gamma players to
profit in the execution of their hedges. ,

When it comes to barrier options, a skilled market maker who is short
the knock-out or long the knock-in can manipulate the market to improve
on his fill.
[End OCR]*

## Page 95

*[Image OCR]
Chapter 5

Arbitrage and the Arbitrageurs

Commerce is sordid... if it is a small affair in which one buys only to resell im-
mediately what one has bought.
Cicero

A TRADER’S DEFINITION!

M Traders define arbitrage as a form of trading that makes a bet on a dif-
ferential between instruments, generally with the belief that the re-
turns will be attractive in relation to the risks incurred. Arbitrageurs
believe in capturing mispricings between instruments or markets.
Traders tend to ascribe the termi “arbitrage” for a wide array of trading
activities of varying levels of risks.”

In more formal academic literature arbitrage means that a linear
combination of securities costing 0 can have the possibility of turning
up with a positive value without ever having a negative value.
Traders’ definition of arbitrage is less restrictive and can be phrased
as follows: arbitrage means that the expected value of a self-financing
portfolio (that can be negative) is positive. The issue of risk neutrality
and martingale measure (irrelevant to most traders) is relegated to
Module B.

The Oxford English Dictionary (in 1971) defines it as: “The similar traffic
in stocks, so as to take advantage of the difference of prices at which the
same stock may be quoted at the same time in the exchange markets of
distant places.”

There are different levels of arbitrage as shown in Table 5.1. Some trad-
ing activities that escape the classification in Table 5.1 carry the designation
“arbitrage” but do not qualify by any stretch of the imagination. “Risk arbi-
trage” in the equity markets has nothing to do with arbitrage; it refers to
takeover where the trader buys the stock of the company being acquired and
sells that of the acquiring one (thus doubling the risk).

80
[End OCR]*

## Page 96

*[Image OCR]
Arbitrage and the Arbitrageurs 81

Table 5.1 Orders of Trader Arbitrage

Degree
First Order

Second Order

Second Order

Third Order

Definition

A strong, locked-in mechanical
relationship, in same
instrument

Different instruments, same
underlying securities

Different (but related)
underlying securities, same
instrument

Different securities, different
instruments, deemed to
behave in related manner
(correlation-based hedging)

Examples

Currency triangular arbitrage
Location arbitrage
Conversions and reversals for
European options

“Crush” or “Crack”

Cash-future arbitrage
Program trading
Delivery arbitrage
Distributional arbitrage
(option spreading)
Stripping

“Value” trading

Bond arbitrage
Forward trading
Volatility trading

Bond against swaps (the asset
spread) -
Cross-market relationships
Cross-volatility plays
Cross-currency yield curve
arbitrage :

MECHANICAL VERSUS BEHAVIORAL STABILITY

A bet on the behavioral stability of a relationship is a weaker form of arbi-
trage than one that attempts to capture values mechanically linked.

@ Mechanical stability means that two instruments have an identifiable
link between them and such link is possible to reproduce artificially.
Behavioral stability exists when one has to marshal historical records
to establish the a posteriori link between two instruments.°

¢ Examples of mechanical stability include locked-in cost-of-carry,
a cross-currency that is the result of two components, a forward-
forward box, an instrument that converges exactly into another.

e Examples of behavioral stability are U.S. versus Canadian swap
curves, German versus Swiss interest rates (until the Berlin Wall fell,
[End OCR]*

## Page 97

*[Image OCR]
82 Markets, Instruments, People

and with it the stability of the relationship), and corporate bonds. Be-
cause there are no true, identifiable links between the commodities,
it is not possible to bridge them mechanically, but history shows a re-
markable price stability, to such extent that one could replace the
other in a hedge. These apparent relationships provide the market
with booby traps.

THE DETERMINISTIC RELATIONSHIPS

In arbitrageur’s parlance, deterministic relationships are ones that can be
matched entirely, creating perfect, or close to perfect arbitrages. Hedging a
deterministic relationship will produce an ironclad protection that should
resist all market moves and remain stable over time.

Deterministic relationships are opposed to behavioral, or stochastic re-
lationships, where there is a correlation between instruments or some de-
pendence but that cannot be expected to resist all kinds of eventualities.

The following are examples of deterministic relationships:

¢ A basket is only the sum of its components. It can therefore be repli-
cated by buying or selling every component.

¢ A cross-currency, like a German mark/Japanese yen cross-currency
exposure can be replicated exactly with dollar versus German mark
and dollar versus Japanese yen.

¢ A knock-out call can be replicated with a vanilla and a knock-in call,
all having the same strikes and out-strikes.

The following are examples of behavioral relationships:

¢ Hedging a bond exposure with a swap is taking advantage of a sta-
ble, not certain, relationship. The Italian swap market stops tracking
the bond market when a crisis causes a divergence between the cred-
itworthiness of the government and that of financial institutions.
The U.S. “flight to quality” can also cause a divergence between the
different issuers.

¢ The common practice of hedging a currency exposure with another
one (e.g., “stacking” exposure to European exchange rates by concen-
trating the hedge in the German mark, owing to its liquidity) can
sometimes bring surprises.

¢ Using all the value-at-risk numbers as exposure presents the same
uncertainty as the preceding hedges.
[End OCR]*

## Page 98

*[Image OCR]
Arbitrage and the Arbitrageurs 83

PASSIVE ARBITRAGE

M@ Passive Arbitrage‘ is the capital neutral swapping of an asset against an-
other similar one that presents more attractive features.

Many operators can engage in arbitrage activities while carrying only one
leg of the spread on their books. Investors who own a security can replace it
with a similar one that commands a higher return or carries a lower risk,
and thus help increase their potential earnings.

Investors who need a certain exposure to the stock market can replace
their baskets with futures when the prices become more attractive. They
could either carry the futures until expiration (and get back to the market
by buying cash on the last day of trading at the settlement price) or switch
again when the price differential flips in favor of the baskets.

Bond portfolio holders willing to maintain a set exposure can roll their
positions along the yield curve to the most attractive issues. They can
search for the “value” point that offers the highest yield or the fastest “roll”
on the curve (see the discussion of convergence in Chapter 12). They can also
switch back and forth between issues as the relationships switch. On-the-
run, liquid issues command a higher price (lower yield), usually to compen-
sate for their liquidity. In some bond markets like that in Italy, the
difference between liquid and illiquid bonds can be significant. Such pre-
mium for liquidity should be of small relevance to the long-term holder and
active switching between issues presents rewards. :

Companies with a set level of financing needs can play the swap market
to improve their costs without significant alteration in the level of the risks
incurred.

Companies with foreign currency exposures can select the most attrac-
tive periods to establish their long-term contracts, They can also actively
move their exposure up and down the curves to the optimal hedging point.

Most fund managers have their performance indexed to an industry av-
erage. The needs to beat the average force them to act like short-term
traders. They can attain better value through indirect arbitrage.

Sometimes the passive arbitrageur has distinct advantages over the ac-
tive one. The holder of Japanese stocks would have better luck swapping
them into convertible securities when these are priced to the onerous stock
borrowing costs of the marginal arbitrageur. Convertible securities would
track the stock thanks to arbitrage, but if stocks are expensive to borrow,
they will be discounted to compensate for the additional costs. If a fund
manager needs to own Japanese stocks, however, it may be more optimal to
swap his holdings into convertibles and hedge the residual interest rate risk.
His expected return will therefore be enhanced by the cheapness of the
warrants that the arbitrageurs are unable to capture.
[End OCR]*

## Page 99

*[Image OCR]
84 Markets, Instruments, People

Examples of such advantages abound. Expensive bonds are difficult to
short when their repo cost is high. The passive arbitrageur, however, would
not need to worry about the repo when he owns the bond and can replace it
with one with a better yield.

AN ABSORBING BARRIER CALLED THE “SQUEEZE”

One of the pernicious aspects of arbitrage is that, while a trade could prove
totally and mechanically hedged, which means that the profits will be unde-
niably realized on expiration regardless of what happens during delivery, op-
erators could still go bankrupt. One example is program trading in which the
short future/long stock will have to pay in cash for the losses on the future
leg, without collecting anything from the paper profits on the stock side. A
sharp rise would easily deplete the trader’s capital: Ironically, it is the most
undeniable form of arbitrage that can be the most biting. The operator, obvi-
ously, will not be alone in such a situation, and the trade will widen, making
the sale of stocks against the purchase of the future unattractive. This is the
case of a classical squeeze: Only the most capitalized can survive; those who
have “deep pockets” can buy the inventories of the distressed smaller firms.

The same risks can affect the conversions and reversals in markets
subject to the future marks-to-market rules, where profits and losses on
the futures legs are immediately collected and option profits are only to
be realized at expiration.

DURATION OF THE ARBITRAGE

@ The duration of arbitrage (for a portfolio constituted of path-indepen-
dent products) is the weighted average time to expiration value of the
absolute amounts in a book. This number needs not to be weighted by
any other factor as it discloses information that affects not necessarily
the P/L risks but a spate of intangible risks, such as personnel risks.

For path dependent products it is advisable to use stopping time, as
explained in Chapter 19. (It is often called first-exit time.)

This measure works well with linear products. If the operator is long the
one year and short the two years in equal amounts the duration of the arbi-
trage is one and a half years.

The institutional structure of Wall Street and the frequency of the ac-
counting period create a predilection for immediate profits. Guaranteed
profits tomorrow against profits today would not do: The manager might lose
his job or (even worse) be penalized in his compensation. Traders are not
noted for their tolerance of pain. The accounting period punctuates Wall
[End OCR]*

## Page 100

*[Image OCR]
Arbitrage and the Arbitrageurs 85

Street’s horizon. Some aggressive brokerage firms have been penalized by
the market for the “proprietary trading losses” in one quarter when they
were holding high-quality convergence trades (such as the Italian interest
rate swaps against government paper). Two results come to mind: (1) The
analysis of any payoff partakes of a larger utility framework and (2) the same
arbitrage varies in attractiveness between traders owing to the differences in
utility curves. There should be a constant shifting of trades between the nar-
row-minded and the infinitely patient.

Such predilection for short-term profits makes the efficient market
frontier unenforceable as the returns for securities would have additional
path-dependent restrictions put on them.

ARBITRAGE AND THE ACCOUNTING SYSTEMS

It is often easier to arbitrage one’s accounting system than the market. The
derivatives markets are growing faster than the systems to properly ac-
count for them. Here is a short list of examples:

¢ Some complex option-related products, such as index-amortizing
swaps or most correlation products are not seasoned enough for the
market to agree on a uniform pricing formula. The absence of market
price allows firms to mark to, their parameters and delude them-
selves in that way. Correlation (not an observable asset in the market)
could be derived on a three-month historical data at one shop and
one-month at another. Both would be wrong as correlation might be
so unstable as not really to exist. The absence of norms would allow
two traders at different houses to book different profits or losses on
the same trade. Unlike listed products, there is no benchmark to use
as a control figure to correct one’s estimation. There is no escape
from theoretical marking to market and the theory is not robust
enough yet.

¢ Money market and swap traders often arbitrage the credit rating of
their employer. They can borrow money better than they can place it
at for the same maturity, and their accounting system could easily
pick up the projected cash-flow differentials as profits, as few sys-
tems discount the creditworthiness of the counterpart. There are
counterpart limits, but no counterpart pricing schedules in ordinary
systems. The money is ultimately earned (barring bankruptcies), but
it belongs to the institution’s credit department, which ultimately
bears the risk, not to the trader’s skills.

* Traders in some houses do not pay interest on the realized losses.
They earn interest, though, on the unrealized profits, under the dis-
counting method (e.g., an option trader is long future against a long
[End OCR]*

## Page 101

*[Image OCR]
86 Markets, Instruments, People

put; the profits in the puts will converge using the Exp(-r t) process
while the future losses are not accounted for in the funding). Traders
could, in a limited way, take advantage of the system.

* Most profits on complex options show immediately on booking the
trade. A trader selling a complex note will be marked at some “fair
value” that approximates the theoretical break-even point but ina fric-
tionless market, that is, one where the subsequent adjustments to the
trade took place without the costs associated with hedging (spreads
and commissions). If the trader were to add up the costs of dynamic
hedging the product over the life of the option, he would be surprised
at the decrease in the potential profits. Trades that require the adjust-
ments with Eurodollar futures would suffer the most. Currencies, es-
pecially the less liquid ones, require onerous costs of rebalancing in
cash, rollovers, and options. Usually, the trader stays ahead by contin-
uously booking new “profitable” trades, then pays the price on liqui-
dation of the position or when the stream of new deals slows to a
trickle (as it did in 1996).

OTHER NONMARKET FORMS OF ARBITRAGE

Many types of arbitrages, including the following, do not fall within the
scope of this book because of their nonmarket orientation:

¢ Credit Arbitrage. It denotes acquiring debt in one security at a
cheaper price than the debt in another. This is generally done as a pas-
sive arbitrage activity, defined earlier. It is usually difficult to short a
corporate bond because of the occasional difficulties in borrowing
them. However, a bonds portfolio manager can improve his expected
return and maintain the overall rating in his portfolio by optimizing
the mix between issues and swapping the expensive ones for better
value bonds. Credit rating is not a market function, and as a trader first
and last, the author recommends listening to the market rather than to
the credit agencies in the definition of ranking. Such skepticism would
therefore invalidate many perceived “value trades” and arbitrages.

¢ Tax Arbitrage. Many equity swaps are due to the privileged tax
treatment of one party in the market. That party can then arbitrage
its condition by transferring the tax to the other party for a profit.
For example, German tax laws impose a dividend withholding tax on
foreign, but not domestic, investors. A foreigner can replicate the po-
sition by entering in a transaction with a domestic German entity
that matches the payoff of the equity swap. The German party can ar-
bitrage the situation by buying the equities and selling the forward
to foreign counterparties.
[End OCR]*

## Page 102

*[Image OCR]
Arbitrage and the Arbitrageurs 87

¢ Legal Arbitrage. Some parties are disallowed by their authorities or
their bylaws to engage in some specific transactions. Domestic French
residents, to give an example, were banned at some point in the 1980s
from buying puts or shorting their currency. Another case of legal ar-
bitrage is the structured note market where the payoff is tied to the
performance of some market, hence includes an option, but is flower-
ing with fund managers who are not allowed to buy options. The law
is skirted through the veil of the note and the fund manager is then
“trapped” since he cannot deconstruct the note himself and needs to
go to the bank to secure a market for the option.

ARBITRAGE AND THE VARIANCE OF RETURNS

The definition of arbitrage is becoming controversial as many forms of
trading defined as arbitrage often carry a higher variance of returns than
outright directional trading. This is partly because the average arbi-
trageur carries larger amounts on his books than the average speculative
trader. It is also due to the accumulation of positions by like-minded arbi-
trageurs, which puts the pressure on relationships. When a security be-
comes perceived as expensive, there will be a rush of traders shorting it
and buying a similar instrument. If the security stays so for a longer time
owing to a specific buyer, the accumulation will turn too large for the ar-
bitrage community to handle and traders will reach their limits. As the
arbitrage community reaches its saturation level, pressure on the relation-
ship between what is deemed expensive and what is deemed cheap will
cause severe marks-to-market losses. Liquidation of the less capitalized
arbitrageurs will ensue.

“Inefficiencies in the market will last longer than traders can remain solvent,”
an option trader once said. .

More advanced notions of arbitrage and stochastic dominance are pre-
sented in Module F.
[End OCR]*

## Page 103

*[Image OCR]
Chapter 6

Volatility and Correlation

What traders and historians share is an ingrained distrust of the notion of
correlation.

An option veteran

This chapter introduces the notion of volatility using minimum mathemat-
ics. The reader should try to develop a sense of where the notion of volatil-
ity can be ambiguous.

BM Volatility is best defined as the amount of variability in the returns of a
particular asset. (As will be discussed, there are many variations in the
methods of measurement):

88

Actual volatility is the actual movement experienced by the market.
It is often called historical, sométimes historical actual.

Implied volatility is the volatility parameter derived from the option
prices for a given maturity. Operators use the Black-Scholes-Merton
formula (and its derivatives) as a benchmark. It is therefore customary
to equate the option prices to their solution using the Black-Scholes-
Merton method, even if one believes that it is inappropriate and faulty,
rather than try to solve for a more advanced pricing formula.

Correlation refers to the least-square measured association between
two random variables. It identifies the degree of certainty with which a
person can predict the move in one random variable as a result of a
change in the other variable. The random variables concerned are the
logarithmic returns of the assets [or Log (Price Period t) — Log (Price
Period t — 1)].

Actual correlation is the amount of actual association between the
moves of two markets. It is often called historical, sometimes histor-
ical actual.

Implied correlation is the correlation parameter derived from option
prices of the components. There will be as many implied correlations
as existing maturities. (Module D includes an explanation for calcu-
lating the implied correlation using the triangle.)
[End OCR]*

## Page 104

*[Image OCR]
Volatility and Correlation 89

Option Wizard: Correlation and Volatility

Traders often make the simple blunder of assuming that a 100% correlation be-
tween two assets A and B means that asset A should move by 1% in response
to asset B moving up by 1% (in the same direction). This is not true: They could
be correlated by 100% and asset B can move by 2% for every 1% move in
asset B, if asset B has twice the volatility of asset A.

What correlation really means is the expected ratio of the moves between
assets divided by their corresponding volatility.

The assumption behind most random walk models is that of lognormal-
ity, which is a necessary solution to the restriction that assets cannot (in
theory) have a negative price. Module A describes the process of a random
walk and describes correlation. The most dangerous assumption is that of a
constant correlation. Traders beware: Markets have adjusted to the fact that
volatility is not constant, not yet to the fact that correlation is volatile.

Figure 6.1 shows the cumulative distribution of one standard deviation
over time. Assuming 15.7% annualized volatility, the market is expected
to move 1% per day, 2.23% every 5 days (1% x V5 ), 4.47% every 20 days
(1% X V20), 15.7% every 248 days—a business year (1% X V248).

Theorists have discussed at length whether the market follows a geo-
metric Brownian motion or an arithmetic one. (See Table 6.1, Figures 6.2
and 6.3.) The difference can easily be explained as follows:

¢ A geometric Brownian motion means (roughly)! that the market

maintains a constant expected percentage move, say 1%. This simpli-
fies a complicated issue, and can explain why volatility in dollar

Asset Price

Years

0.2 0.4 0.6 0.8 1

Figure 6.1 One standard deviation as a function of time.
[End OCR]*

## Page 105

*[Image OCR]
90 Markets, Instruments, People

Table 6.1 Constant Dollar and
Percentage Moves

Constant Constant
Dollar Percentage

Asset S Move Move
80 19.84% 0.80
82 19.35 0.82
84 18.89 0.84
86 18.45 0.86
88 18.03 0.88
90 17.63 0.90
92 17.25 0.92
94 16.88 0.94
96 16.53 0.96
98 16.19 0.98
100 15.87 1.00
102 15.56 1.02
104 15.26 1.04
106 14.97 1.06
108 14.69 1.08
110 14.43 1.10
112 14.17 1.12
114 13.92 1.14
116 13.68 1.16
118 13.45 ° 1.18

120 13.23 1.20

20.0%
18.0%
Za
= 16.0%
3
14.0%
12.0% +4+-+44+4++44-4- 44-4 4H + HHH HH
Qo oO N ao vt °o ©
eo wo o o 2 = =
Asset Price

Figure 6.2 Constant dollar move 1% per business day.
[End OCR]*

## Page 106

*[Image OCR]
Volatility and Correlation 91

1.20

g

° 1.10

=

D> 1.00

o 1.

5

8 0.90

a

a
0.80

oCUC«wmUUC NLC THK CUDTCOCCD
on ODO DO DMD OG TF =
- +.

—

Asset Price

Figure 6.3 Percentage move at 15.7% volatility.

Option Wizard: “Perhaps Bachelier Was Right”

Many of the caps/floor book runners were hurt by the rally in Eurodollars (drop
in yield) between 1991 and 1993. With interest rates at 9%, it was remarkable
that historical volatility was close to about 9 basis points per day on average.
Market markers, however, were selling out-of the money calls with the belief
that “percentage” volatility would be unchanged in the rally, which means that
at 4.5% Eurodollars, the “tick volatility,” or “dollar volatility,” would be close
to 4.5 basis points. They were proven wrong as the tick volatility remained al-
most constant, and their positions in out-of-the-money calls made them short
volatility at very low levels. .

The rally in the Eurodeposits resulted in a rise in implied and historical
volatility as defined by option theory (i.e., lognormal returns). Somehow the
price of the at-the-money straddle in ticks remained close to the same level
throughout. This prompted a famous market thinker to confide to the author:
“Perhaps, after all, that man Bachelier was right.” He was indicating Louis
Bachelier, the French mathematician, who wrote his doctoral dissertation in
1900, 73 years before Black and Scholes, pricing options with an arithmetic
Brownian motion. Bachelier was mistreated by history for a long time.*

Euroyen traders experienced considerably more trauma. As short-term in-
terest rates in 1995 reached .10%, a 1 basis point move corresponded to
160% volatility. Volatility moved during the rally from the “social” levels
{around 20%) to 200% as the Eurodeposit traders were not informed that the
market needs to be lognormal. in addition, Bachelier’s model seemed perfectly
credible as the market started paying for 100 calls (puts struck at zero) based
on the belief that the possibility of the market going negative could not be
ruled out.

Moral: Traders should take theoretical dogma less seriously.

*See Bachelier (1900).
[End OCR]*

## Page 107

*[Image OCR]
92 Markets, Instruments, People

terms needs to be higher at higher prices. It also explains why volatil-
ity would drop as the market weakens, a dampening that would pre-
vent the market from reaching negative levels. As the market drops
from 100 to 1, 15.7%, volatility would represent .01 ticks. At .10 it
would be .001 and so on.

e An arithmetic Brownian motion means that the market maintains
constant expected dollar moves, regardless of its level. This assump-
tion is refuted by academics because it can lead to possible negative
asset prices. However, traders take it seriously as they believe that for
medium-size increments, the assumption holds. In addition, money
market instruments have unexpectedly flirted with negative prices
throughout history.*

¢ The interesting results of a mixed process usually witnessed in fi-
nancial markets—arithmetic in the short term and geometric in the
long run—will be discussed later in this book.

CALCULATING HISTORICAL VOLATILITY
AND CORRELATION

Centering around the Mean

For ease of calculation as well as theoretical reasons, volatility of returns
for a dynamic hedger could be calculated as the square root of the sum of
the squares of the movements, not the deviations from an average. If a mar-
ket moved 1% a day every day in the same direction (say upward) for an
entire month, the conventional measurement of volatility would put it at
0% since all of the moves were at the mean move. This clashes with an op-
tion trader’s instinct. An option trader would stjll consider it to be 1% a
day and in this case would prefer to buy volatility when it is offered
cheaper than 16%.

Statistically, given the absence of well-pronounced trends, or trends
that dominate the variance, both measures present similar results.*

Formulas. Taking the return in any period for a given tradable pair in the

market to be
bos 5)
x, = Log
' P, -1

The (natural) log of the two prices will correspond (roughly enough for
traders) to the percentage return.

These prices need to be sampled periodically and consistently. If the fre-
quency is every Thursday at midnight, the returns all need to correspond to
[End OCR]*

## Page 108

*[Image OCR]
Volatility and Correlation 93

such a period. Many operators use daily official close to close while others
take a price during the day when all markets are very liquid.
Volatility of x for any given period is generally estimated as

1 n _
Ox” \ (2-1) 2 (9

Note that (1 — 1) is used in place of n because of the loss of one degree of
freedom in the estimation of the mean x.

n
x= >
t=1

Ditching the x, the mean return, means there is one parameter less to
estimate. It also makes the estimation of volatility closer to what would af-
fect a trader’s P/L.* So the volatility (nonweighted) noncentered o’ could be
expressed as:

x

= (2

, 1 n 3
>

To annualize and express the volatilities using the trader’s conventions,
multiply by the annualization factor. If the returns from calculating volatil-
ity are daily and the year is considered to be 248 days, the trader would
multiply the figure by V248. If the returns are weekly, he would then mul-
tiply by V52, and so on.

Note that the trader could multiply by V365 provided he added the
weekends in the data series; that would add two zeros per week to the com-
putation and would be roughly equal with some variance for the shorter
dates.

Table 6.2 provides an example of historical volatility calculation. It
shows daily prices (for an imaginary pair) in succession for 22 days. This
yields 21 returns. Column 3 calculates the natural logarithm of the return
and Column 4 calculates the square of the Log(S,/S,_ ,). The annualized
volatility is computed as the square root of the sum of Column 4 /21 multi-
plied by the square root of 248 = 10.73%.

The correlation between the returns x and y (both being pairs not nec-
essarily of the same numeraire) is classically defined as:

So, - Hy, -D
1 =1

pw n-1 0,9,

but it is permissible to ignore the mean of the returns and to use:
n

1 et

pn o' 0"

y
[End OCR]*

## Page 109

*[Image OCR]
94 Markets, Instruments, People

Table 6.2 Computation of Historical Volatility

Day Price Log(S,/S, _ ,) {Log(S,/S,_ ,)?
1 99.36
2 98.73 —0.006372 0.000040608
3 98.97 0.002412 0.000005820
4 99.94 0.009795 0.000095941
5 100.20 0.002608 0.000006801
6 101.47 0.012560 0.000157765
7 102.06 0.005819 0.000033856
8 101.50 —0.005514 0.000030409
9 102.45 0.009293 0.000086363
10 102.89 0.004357 0.000018979
11 103.09 0.001886 0.000003558
12 103.71 0.006020 0.000036234
13 104.04 0.003142 0.000009873
14 105.17 0.010855 0.000117827
15 104.52 —0.006257 0.000039155
16 104.32 —0.001886 0.000003556
17 103.31 —0.009714 0.000094366
18 102.27 —0.010156 0.000103154
19 101.63 —0.006251 0.000039072
20 101.22 —0.004074 0.000016596
21 101.44 0.002185 0.000004776
22 101.04 —0.003941
V Average Average
Annualized
Volatility (%) 10.73

Table 6.3 provides an example of correlation calculations.

o’, = .006762 (using daily volatilities)
‘3 = -006363

D Ret, Ret, = .000019
Correlation = .000019/(21 .006363 .006762) = .02

tl

q
tl

The links between implied volatility and implied correlation will be exam-
ined later in the book.
[End OCR]*

## Page 110

*[Image OCR]
Volatility and Correlation 95

Table 6.3 Correlation Calculations

Ret, Ret,

Day S, Sp Log (S4/Sq-1) Log (S5,/S,,_,)  (Ret,)” (Ret,)’ Ret, Ret,

1 99.36 100.32
2 98.73 100.95 —0.006372 0.006246 0.000041 0.000039 —0.000040
3 98.97 101.40 0.002412 0.004448 0.000006 0.000020 0.000011
4 99.94 102.93 0.009795 0.014965 0.000096 0.000224 0.000147
5 100.20 102.96 0.002608 0.000281 0.000007 0.000000 0.000001
6 101.47 103.06 0.012560 0.000991 0.000158 0.000001 0.000012
7 102.06 102.83 0.005819 —0.002253 0.000034 0.000005 -—0.000013
8 101.50 103.04 ~0.005514 0.002091 0.000030 0.000004 —0.000012
9 102.45 103.16 0.009293 0.001113 0.000086 0.000001 0.000010
10 102.89 103.85 0.004357 0.006734 0.000019 0.000045 0.000029
11 103.09 103.36 0.001886 —0.004779 0.000004 9.000023 —0.000009
12 103.71 103.75 0.006020 0.003798 0.000036 0.000014 0.000023
13 104.04 103.76 0.003142 0.000128 0.000010 0.000000 0.000000
14 105.17 103.29 0.010855 —0.004544 0.000118 0.000021 —0.000049
15 104.52 104.93 —0.006257 0.015749 0.000039 0.000248 —0.000099
16 104.32 104.28 —0.001886 —0.006274 0.000004 0.000039 0.000012
17 103.31 103.56 —0.009714 —0.006898 0.000094 0.000048 0.000067
18 102.27. 104.18 —0.010156 0.005926 0.000103 0.000035 -—0.000060
19 101.63 104.85 —0.006251 0.006463 0.000039 0.000042 —-0.000040
20 101.22 104.75 —0.004074 —0.000952 0.000017 0.000001 0.000004
21 101.44 105.32 0.002185 0.005442 0.000005 0.000030 0.000012
22 101.04 104.99 —0.003941 —0.003193 0.000016 0.000010 0.000013

> 0.000960 0.000850 0.000019
V Average 0.006762 0.006363

INTRODUCING FILTERING

What volatility should be used? 10 days? 100 days”... There is, fortunately
for traders, no straight answer, and such disagreements help create a mar-
ket. Some traders with a long memory prefer to go back many years, while
others are victims of some form of market amnesia.

Filtering is a simple method for taking into account that events in the
past need to have uneven weightings (see Figure 6.4). The following text
presents a simplified version of the Kalman filter: the exponential decay.°

It behooves the trader to assign importance to recent events in propor-
tion to their distance away from the present. But the trader should be flexi-
ble enough not to take the measurement for gospel, either because he may
have information about a predicted riot tomorrow that would make all past
information academic, or because some structural changes took place that
need to have an effect on the weighting. The last thing a trader needs to do
is be scientific or play econometrician during business hours.
[End OCR]*

## Page 111

*[Image OCR]
96 Markets, Instruments, People

Figure 6.4 Decay factors: Volatility filtering.

Table 6.4 Computation of Volatility

Day Sl Log(S,/S,_ ,) {Log(S,/S,_,)}??
1 99.36 -
2 98.73 —0.006372 0.000041
3 98.97 0.002412 0.000006
4 99.94 0.009795 0.000096
5 100.20 0.002608 0.000007
6 101.47 0.012560 0.000158
7 102.06 0.005819 0.000034
8 101.50 —0.005514 0.000030
9 102.45 0.009293 0.000086
10 102.89 0.004357 0.000019
11 103.09 0.001886 0.000004
12 103.71 0.006020 0.000036
13 104.04 0.003142 0.000010
14 105.17 0.010855 0.000118
15 104.52 —0.006257 0.000039
16 104.32 —0.001886 0.000004
17 103.31 —0.009714 0.000094
18 102.27 —0.010156 0.000103
19 101.63 —0.006251 0.000039
20 101.22 —0.004074 0.000017
21 101.44 0.002185 0.000005
22 101.04 —0.003930 0.000015

SUM

nN

0.5275
0.5438
0.5606
0.578
0.5958
0.6143
0.6333
0.6528
0.673
0.6938
0.7153
0.7374
0.7602
0.7837
0.808
0.833
0.8587
0.8853
0.9127
0.9409
0.97
15.278

2.14E-05
3.16E-06
5.38E-05
3.93E-06

9.4E-05
2.08E-05
1.93E-05
5.64E-05
1.28E-05
2.47E-06
2.59E-05
7.28E-06
8.96E-05
3.07E-05
2.87E-06
7.86E-05
8.86E-05
3.46E-05
1.51E-05
4.49E-06

1.5E-05
[End OCR]*

## Page 112

*[Image OCR]
Volatility and Correlation 97

The symbol A is used for the decay factor. It replaces the number of days
to go backward in the previous regime. In place of lengthening memory, one
could bring the weights closer to 1.

O=>Y M=Q4N24084... 44) 2
f=1 1-—A
when 7 is very large (higher than 1,000 observations as a general rule) and
since ) is, by definition, less than 1.
The same decay factor could be applied to the correlation, in the same
manner, to cancel the effect of recent memory.

pay SNe
t=1
Using the same example, Table 6.4 shows the computation of the volatility
using a A of .97.

Column 1 presents the number of days; Column 2 presents the asset
moves; Column 3, the logarithmic returns (natural logs); Column 4, the
square of Column 3; and Column 5 has the \ to the power of the number of
days. It is apparent in moving up the columns that the weightings vanish in
importance. We use \ on the first day (day 22 since we are computing
volatility on day 22 or later), \? on the’second day (day 21), A? on the third
day (day 20), and so on.

Column 6 shows the result of multiplying Columns 5 and 4, and is the
weighted squared return for the period.

The daily weighted volatility is the square root of (sum of Column 6
/21)/sum of Column 5 = .006675. The annualized percentage volatility is
obtained by multiplying it by V248 = 10.51%.

The reader could try the same exercise with the weighted correlations.

THERE Is No SucH THING AS CONSTANT
VOLATILITY AND CORRELATION

Figures 6.5 through 6.9 will show the trader to beware the notion of con-
stant volatility and, even more, that of constant correlation.

Figure 6.5 shows implied volatility in the USD-DEM for the one-month
options during a part of 1992.

Figure 6.6 shows actual two-week volatility in the SP500. The volatility
of the SP500 shot up to 120% the week of the crash of October 1987, which
showed a peak into the graph boundaries. Measuring the volatility of vola-
tility would show it to be even.more volatile than that of the underlying
[End OCR]*

## Page 113

*[Image OCR]
98 Markets, Instruments, People

Option Wizard: Rubber Time Explained

The notion of economic versus actual time crops up once in the while in dis-
cussions between option traders. It is a difficult matter but seems to be ad-
justed for properly by the markets.

Trading on the weekend does not take place too often. Markets are as-
sumed to be open about 247 days (except for European countries with a
plethora of holidays).

Traders price the process in volatility of 365 days but actually adjust back
by removing holidays. For example, if they decided to price a Tuesday option
on a Friday, they would use the number of business days (2) and would manip-
ulate the formula that uses 5 days to pretend that the volatility over the next 5
days would be low enough to accommodate for the actual moves. If they as-
sumed that volatility is supposed to be 15.7% in the market per actual number
of days, it would be 1% for 1 day, 1.41% for 2 days, 15.7% for 1 year, and
so on... The operator on a Friday would use the volatility that would give
such result of the expected move of 1.41% over 4 days, namely 11.08%
(1.41%/V4 = .71 XV 247 = 11.08%).

More advanced methods give a weight per actual day and assume that
Sundays can exist as a source of volatility since nasty revolutions can take
place at such idle times. Many operators count Saturday and Sundays as quar-
ter days in the winter and even less in the summer.

Some traders actually price options “on the clock” with a real almanac of
expected market moves that lower the intensity of lunchtime in Tokyo when
markets go into a slumber.

The important intuition traders need to grasp is that volatility and square
root of time have the same effect on the dispersion of a random walk.

23.00%
21.00%
19.00%
17.00%

15.00%

Volatility

13.00%

11.00%

9.00%

7.00%

30/92
TIW92
TAN
7eahs2
8/7/92
a/18/92
9/16/92
9/25/02
10/6/92
10/18/92

10/26/92

11/4/92

6/10/92
6/19/92
9/27/92
9/7192
11/13/92
11/2492

Time

Figure 6.5 USD-DEM Implied volatility is not constant.
[End OCR]*

## Page 114

*[Image OCR]
Volatility and Correlation 99

Volatility
om SGaBReRS

HTH

Time
Figure 6.6 SP500 2 week historical volatility 1985-1995.

security itself, particularly if it were broken up into nonoverlapping peri-
ods of short duration.

When measuring the volatility of volatility traders should be careful to
avoid running overlapping data. In addition, the length of the period matters,
and it is recommended to select as short a period between sampling points as
possible. Figures 6.6 and 6.7 show that historical volatility moves even more
than implied.

Correlation is even more volatile than both volatility and the underly-
ing itself. Figures 6.8 and 6.9 show the correlations of the daily moves in
nonoverlapping periods of two weeks.

One does not have to be a statistician to understand that correlation
moves all the time.

Chapter 15 will provide an in-depth discussion of distributions, from
the vantage point of option pricing.

Figure 6.10 shows the volatility of historical volatility in nonoverlap-
ping periods, and Figure 6.11 depicts implied volatility.

Volatility
a SR 8

=
a o
———}

Time
Figure 6.7. Yen 2 week historical volatility 1985-1995.
[End OCR]*

## Page 115

*[Image OCR]
100 Markets, Instruments, People

0.4

-0.6
Time

Figure 6.8 Correlation between sterling against USD and FRF against USD, 1985-
1995.

1

eq he
|

-0.2
-0.4
-0.6

Corr

Time

Figure 6.9 Correlation between sterling against USD’and JPY against USD, 1985-
1995.

Time

Figure 6.10 Volatility of historical volatility, 1985-1995.
[End OCR]*

## Page 116

*[Image OCR]
Volatility and Correlation 101

Time

Figure 6.11 Volatility of implied volatility, 1985-1995.

THE PARKINSON NUMBER AND THE
VARIANCE Ratio METHOD

Among the most significant contributions to option traders are the Parkinson
number and the variance ratio method.

The Parkinson number, generated by the physicist Michael Parkinson in
1980, aims at estimating the volatility of returns for a (geometric) random
walk using only the high and low in any particular period. This section will
show how to use it selectively, most often in reverse, to derive the distribu-
tion of the maximum or minimum in any day knowing the historical volatil-
ity.© The Parkinson number will be called the P:

1 n 1 S 2
P=,/ — log —H
| (ny 2, dlog2 8S
We have S,, and S, the close-to-close registered high and the registered low
respectively in any particular time frame.

¢ All traders respect the integrity of the close-to-close number because
of its official nature but are cautious about taking the high or low be-
cause of the misprints and manipulations attending its registration.
Sometimes fictitious trades take place at one level owing to manipu-
lation while at other times computer errors mar the data.

¢ In addition, there are many instances in cash over-the-counter prod-
ucts, such as illiquid foreign exchange markets, where a new high or
a new low could escape the screens and remain the sole knowledge of
the traders involved in the trade.

¢ Finally the Parkinson number applies to a 24-hour period assuming
no closing of the market and no discontinuities (an even number of
[End OCR]*

## Page 117

*[Image OCR]
102 Markets, Instruments, People

Option Wizard: “GARCH in Their Head”

Option traders can easily understand ARCH modeling precisely because they
“have GARCH in their heads” as said one trader to whom the concept was ex-
plained on a yellow sandwich bag. Indeed, they have a much better one than
that presently available to econometricians.

Sometime in 1982, Engle* discovered heteroskedasticity (i.e., changing
volatility) in times series on British inflation, and started the first attempt at
modeling it. It resulted in ARCH (autoregressive conditional heteroskedastic-
ity), which was an effort to estimate the volatility process using techniques
from times series analysis called autoregressive models. It assumed that future
volatility was linked to its past realizations, each realization having its own
weight. Intuitively, it can be described as a process that can be forecast using
past volatility data in a decreasing manner (as with the filtering technique).
“Volatility begets volatility,” as the saying goes. A volatile day is more likely to
be followed by a volatile day and a quiet one by a quiet one. Inexpensive soft-
ware packages now help the researcher in “fitting” the parameters to come up
with an elaborate (but alas fragile) prediction tool.*

After Engle, a deluge of models started hitting the academic world, attempt-
ing to make it more potent. Bollerslev and Engel generalized it to GARCH (both
autoregressive and moving average). The most interesting development came
with Nelson’s E-GARCH (exponential), which considered that the volatility
shocks down (when the market drops).are the significant part in determining fu-
ture volatility, a fact familiar to equity traders and Mexican peso investors.

Later more complex GARCH models were built on the realization that
shocks that generally matter are those where some threshold is broken, those
that depend on a regime condition, and so on. Then came N-GARCH (nonlin-
ear), SWARCH (switching regime), and others. H-GARCH (heterogeneous) used
a time scale that is similar to the concept of rubber time for an option trader. The
models are getting closer to a seasoned option trader’s opinion, which could be
that past information about the price action impacts future volatility, but in
some complex way. Subtle sets of rules in the trader’s mind help him form his
own opinion in a more potent way than these data-processing techniques. This
might explain why ARCH and siblings (even the neural network extensions)
never made an impact in the trading room: They are still at this point but a weak
imitation of a veteran trader’s mind. Implied volatility, basically the trader’s
opinion, outperforms GARCH predictions except where. squeezes take place
that make the implied reach abnormal proportions. Otherwise, it would be used
with more success (see Taleb, 1996a).

Finally, one reason GARCH never made it is that implied volatility contains
information that is not available in past prices (e.g., scheduled elections,
economic releases) but that are significant in determining future volatility.
A sudden announcement of a meeting between two trade ministers would

(Continued)
[End OCR]*

## Page 118

*[Image OCR]
Volatility and Correlation 103

not impact the currency pair but traders know that the outcome will result in a
jump (up or down, pending the nature of the agreement); such information is
not contained in past prices. As a matter of fact, volatility would freeze, as the
market will be numb to any information on the currency prior to the meeting.
So GARCH would predict a drop in volatility.

In addition, some information that was available in past prices (e.g., jumps
due to elections) does not appear to bear any significance for future volatility
owing to its nonrecurring nature. A trader knows how to filter them in his
volatility prediction in a way that econometricians do not. As a community,
traders benefit from close to three decades of oral tradition and war stories.

The conclusion is we conclude that in the present state of the contest
between the computer and the trader, the trader still has, markedly, the upper
hand.

*See Engle (1982).
*The GARCH(1,1) can be characterized by the following process for o, the volatility period ¢.

2= 2
o a + a,€

2
' -1+B,¢

t t-1

e = 0/7 Z?, Z, follows a normal distribution with 0 mean and unit variance.

a, > 0, (generally of the order of .01 for daily observations), «, + B, <1 (but generally
very close to 1).

What is remarkable about GARCH is that both the asset price and its volatility share the
same “innovation” Z with the difference that Z, becomes squared for volatility.

We can see that the B, represents the “constant volatility” coefficient, while a, provides a
factor of the GARCH innovation, the stochastic element. The higher the a, the fatter the tails of
the distribution of the asset. In a way, volatility follows a Chi-square process and plotting the
process generated by it shows a pronounced right tail. In short, the distribution of volatility pre-
sents a positive skew (third moment) and that of the asset presents a high kurtosis (that is the
ratio of the fourth moment to the square of the second moment).

transactions). Otherwise, it is preferable to use the Garman and Klass
(1980) estimation of volatility through the combination of close-to-
close and high/low. The Garman-Klass estimator GK is:

GK se 1 | 3 Si |
= 5ilog —* | — .39| log
S, S

t-1

An important use of the Parkinson number is the assessment of the dis-
tribution of prices during the day as well as a better understanding of mar-
ket dynamics. Comparing the Parkinson number and the periodically
sampled volatility helps traders understand the mean reversion in the mar-
ket as well as the distribution of stop-losses. Some clear rules can be de-
rived from that information.

Comparing the Parkinson number P with the definition of periodically
sampled historical volatility gives this result:

P = 1.670’
[End OCR]*

## Page 119

*[Image OCR]
104 Markets, Instruments, People

It means that the volatility of the market as observed through the 24
hours or 1 week, or any stable sampling, period should be related (through
the distribution of the maximum and the minimum) to the volatility as
measured by the extremes.

Warning: Such measurement cannot be used to compare close-to-
close volatility with intraday high/low. It can compare 24-hour
high/low to data sampled every day at the same time. For markets,
like most equities, which trade during the day only, it is better to use
open-to-close volatility.

This estimator can give meaningful information with the following sit-
uations:

¢ Pricing Barrier Options (and the related American digital and look-
back options). They get triggered through the printing of a price on
the screen; therefore, the distribution of the extremes is most impor-
tant. The barrier option trader needs only one information: the high
or the low to see if his option was knocked-in or knocked-out. How
this extreme is distributed matters more than the close-to-close or
any other estimator of volatility. The Parkinson number is the sole in-
formation. If there is a bias making the P consistently higher than
1.67 o’ , the trader knows that the probability of hitting the trigger is
higher. This will be discussed later in the book with reference to bar-
rier options.

* General Option Delta Adjustments. The comparison between the
Parkinson number and the periodically sampled volatility can reveal
serious information about the mean reversion of a particular market
and allow the trader to set his frequency of adjustment accordingly.
If P is higher than 1.67 o' the trader needs to hedge a long gamma
more frequently. Otherwise, he could lag the needed adjustment, a
technique called “letting the gammas run.””

¢ General Trading Strategies. The market maker edge is strongest in
cases where P is higher than 1.670’. It is otherwise better to follow a
trend. This manifests itself (see Chapter 4) in a negative short-term
autocorrelation of prices.

Figure 6.12 shows the Parkinson number ratio to the volatility during
the same period for the U.S. Treasury bond futures over a period spanning
almost three years until May 1995. The results are strikingly convincing:
There seems to be a clear bias in favor of a wider high/low range than as-
sumed by a random walk. Additional testing by the author® shows the bias
[End OCR]*

## Page 120

*[Image OCR]
Volatility and Correlation 105

3
2.8
2.6
2.4
Q
S 22
F< af
2
1.8
1.6 Theoretical
1.4
SBRSSSSARESHESES SA
SSsStERESRsagSeSRasSBE

Figure 6.12 Ratio of the Parkinson statistic and open-close volatility in U.S. bond
futures, 8/1992-5/1995.

to be permanent in close to the 20 markets surveyed. The reader can draw
his own conclusions.

Another seminal piece of work (for traders) concerns the frequency of
sampling. An interesting study for option traders is A. Lo and A. C. MacKin-
ley’s paper? on a particular application of a technique called the variance ra-
tios. The authors intended to prove that stock prices have a memory through
a simple test of variance relative to sampling frequency.!°

Briefly, if the volatility on an hourly sampled basis turns out to be higher
than the volatility on a daily sampled basis, the market can be considered as
mean-reverting. On the other hand, if the market showed a higher volatility
at a wider frequency between the dates then it can be concluded that there is
a trend. More powerful tests were later invented!’ but this simple variance
ratio is simple enough for traders to use and understand.

For example, if the market moved by 1% per day, it would be expected
that over 20 business days it moved by V20 = 4.47%. Otherwise something
suspicious would be deemed to be taking place, such as the market moving
more often one way than the other.

What traders usually notice is the higher volatility when sampled
hourly, especially in such markets as the SP500 and the Eurodollars. This
happens regardless of the ratio between the volatilities for longer periods,
such as 1 day and higher.

The variance ratio is well known by traders, even by those who have never
heard of the method (see Figure 6.13). Often stock traders wonder why the
broad market moved during 1995 by close to 35% with historical volatility
close to 10%, or why the dollar lost periodically 20% of its value every year
during the 1980s without volatility being any higher than the teens.’?
[End OCR]*

## Page 121

*[Image OCR]
106 Markets, Instruments, People

Historical

Mean Reversion
Volatility

SN

ee
I >
1 hour 1 day

Frequency of Observations

Historical Trend
Volatility (rare Intraday)

a

—_———— —_—

_—

1 hour ° 1 day
Frequency of Observations

Figure 6.13 Variance ratio method.

0.3
0.25
0.2

0.15

Volatility

0.1 Daily

0.05

1/20/93
3/31/93
6/10/93
8/19/93
10/28/93
1/7/94
3/18/94
6/8/94
8/17/94
12/29/94
3/13/95
§/22/95

Figure 6.14 Smallest possible time increment volatility.
[End OCR]*

## Page 122

*[Image OCR]
Volatility and Correlation 107

Figure 6.14 shows the high frequency tick data as a source of informa-
tion about both the transaction cost function and the mean reversion of
markets as expressed by the edge of the market maker in cash. The volatility
was computed as measured by tick increments to show the reader how ex-
treme an effect the sampling of volatility can have over the price. It shows
the “instantaneous” volatility (as required by Black-Scholes-Merton for the
replication of the option) as twice the daily measured one.
[End OCR]*

## Page 123

*[Image OCR]
part LI
MEASURING OPTION RISKS

THE REAL WORLD AND THE
BLACK-SCHOLES-MERTON ASSUMPTIONS

Risk Management Rule: Dynamic Hedging with Black-Scholes-Merton.
It is better to improve on a simple but seasoned model than operate

with a more advanced but newer model.
It is better to use a model with the smallest number of parame-
ters to estimate.

Better models often create nightmares. Every options trader and risk man-
ager, while reaping the benefits of the contributions of Black, Scholes, and
Merton, needs to spend his time working around the assumptions that
needed to be made for the model to stand.’ Indeed, all experienced traders
operate under the same model and are comfortable with it because they
have learned the necessary tricks to make it work. Despite the criticisms of
the formula, traders have refused alternatives because they have learned its
limitations. No experienced trader would willingly trade Black-Scholes-
Merton for another pricing tool.

Black-Scholes-Merton as an Almost Nonparametric
Pricing System

With the model, we are able to get a unified measure of the price of an op-
tion. Traders can build from there by introducing the additional parameters
as an informal tacked-on improvement, such as the jump factor, volatility of
volatility, correlation between asset price and interest rates, and different
statistical estimates of the underlying process. With most exotic options,
the need to parametrize may be more pronounced.

109
[End OCR]*

## Page 124

*[Image OCR]
110 Measuring Option Risks

Option Wizard: The Essence of Black-Scholes-Merton

The essence of the model* is risk-neutral replication of securities in a market
that is said to be complete. It does not mean, however, that every security
needs to be replicated by every arbitrageur, as has often been misinterpreted.

For example, with a warrant on a Japanese stock, the operators involved in
such arbitrage cannot dynamically replicate it because of frequent difficulties
in shorting the stock. It does not mean, however that fair value needs to be dif-
ferent since the replication costs hamper arbitrageurs. Black-Scholes-Merton
does not require for every operator to be delta neutral in infinitesimal changes
in the stock price. Some invisible passive replicator would trade the warrant for
the stock back and forth.

Another problem, however is that an operator may need to hold a security
until expiration. To reach his decision to buy the security, he will have to use a
term variance that is different from the infinitesimal variance of the security.

Just as the value of some options can be thought of as limit decomposition
of others with the full knowledge that such value would not be attainable
through replication, we can consider the Black-Scholes-Merton fair price as the
limit of the risk-neutral arbitrage.

*Advanced topic.

With risk management, another issue appears: that of the accuracy of
the Black-Scholes-Merton predicted changes in price. Traders have to work
around derivatives of the formula, building modifications that range from
the simple rules of thumb to the more complete multifactor models.

Black-Scholes-Merton is based on five assumptions. The first two are
fundamental (are philosophically of the essence. of the model); the next
three are parametric or distributional (can be altered without any signifi-
cant change in the model):

1. Ito Process (fundamental). \t is characterized by a random component
that is independent and identically distributed (as described ear-
lier—the Brownian motion). Its principal characteristic is that it needs
to be memoryless. The notion of continuous prices is not as relevant as
it was made out to be by early detractors of Black-Scholes-Merton:
Continuous finance is a tool, not a philosophical statement.

2. Frictionless Markets (fundamental). There are no transaction costs, no
costs of adjustment, no stamp tax, or exchange controls. This as-
sumption added to the previous one implies that the operator can
buy and sell in large quantities to adjust the delta. It leads to total
[End OCR]*

## Page 125

*[Image OCR]
Measuring Option Risks 111

absence of the impact of utility functions. The existence of transac-
tion costs would necessarily change the argument for a hedging pol-
icy of an isolated operator, but it would not impact fair value.

3. Constant Volatility (parametric assumption). It implies that the daily
variations are drawn from the same distribution and that the vari-
ance is known. It leads to a constant correlation between different
assets.?

4. Geometric Brownian Motion (distributional assumption). It implies (see
Module B and Chapter 10) that the motion of assets is “geomet-
ric’—that the expected variance of the logarithms of the returns
remains constant.

5. Constant (and known) Drift (parametric assumption). In trader’s terms,
the structure of the forwards slope is constant.

A series of corrections will be made to reduce the impact of these assump-
tions. First, it is necessary to gauge the cost to traders:

¢ Assumption 1 needs to be lifted in cases of serious path dependency,
as with liquidity holes, particularly if the manufacturing of the op-
tion might impact the path of the underlying asset (e.g., as for port-
folio insurance).

¢ Assumption 2, sometimes weakened, means that option traders can-
not adjust their delta every microtick change in the price of the un-
derlying. Less frequent rebalancing of the portfolio implies less
tracking of theoretical value per trade but only in the long run. Thus
the operator would have to start divorcing theoretical value and dy-
namic replication. It means that value can be obtained from continu-
ous time finance but the delta needs to be discretely computed.?

¢ Assumption 3 is perhaps easiest to correct. Volatility is unstable since
traders make markets in it (one can buy or sell volatility). This leads
to a divorce between historical and implied volatility. The instability
of volatility causes the delta to lose its quality of hedge ratio (unless
properly modified), and the gamma its predictability of changes in
delta. However, sometimes the trader’s work will be made easy by
awareness of some link between volatility changes and magnitudes
of market moves. Such information is included in this text. Moreover,
because the volatility is not constant, a volatility curve results and as-
sumptions must be made about the behavior of such curve.

¢ Assumption 4 will be discussed in Chapter 11. In some cases, the dis-
tribution might be considered arithmetic, as with Bachelier’s early
work.
[End OCR]*

## Page 126

*[Image OCR]
112 Measuring Option Risks

Greeks and Their Shortcomings

Greek
Delta

Gamma

Theta

Vega

Definition

The sensitivity of a
derivative price to
the underlying
security.

The rate of change
or “curvature” in
the delta.

The sensitivity of a
portfolio of options
to time.

The sensitivity of an
option to volatility.

Shortcomings

- Continuous time

hedging is for the
textbook.

The delta does not
work on a portfolio
of options that mixes
longs and shorts.

It is an extremely
weak measure of
risks.

It is meaningless for
a portfolio of
options.

It does not take into
account changes in
volatility when the
market moves.

It does not take into
account the changes
in volatility with the
shortening of
expiration.

It does not take into
account that
volatility drops
when markets do not
move at all.

It is extremely
misleading in a
portfolio that has
calendar spreads, as
different maturities
do not have the
same volatility of
volatility.

Modification

Use a discrete delta
with increments.
Shadow delta adds
some vegas and
gammas to it.

As with a delta, use
discrete increments.
Examine two
different numbers:
Up gamma and
Down gamma.
Shadow gamma takes
into account the
volatility smile.
Skew gamma
measures the impact
of nonlinear changes
in volatility arising
from an up-move or
a down-move.

Use a term structure
of volatilities to
reprice the portfolio
one day hence—only
when volatilities are
deemed to converge.
Use shadow theta to
accelerate time decay
in stable markets
and reduce it in
volatile conditions.

Use a simple
volatility curve
model for the
weightings.

More advanced: Use
a covariance matrix
for the buckets.*
[End OCR]*

## Page 127

*[Image OCR]
Continued

Greek
Rhol

Rho2

Convexity

Omega

Alpha

Correlation
Vegat

Definition

The sensitivity of an
option to domestic
costs of carry.

The sensitivity of an
option to dividends,
coupon, or the
foreign rate.

A general blanket
term to define
nonlinearity ina
derivative
instrument.

Shortcut method to
compute the
expected life of an
American vanilla, or
the expected exit
time for an
American binary or
a knock-out option.

Method to compute
the cost of gamma.

Method to compute
the sensitivity of
the price of a
structure ina
portfolio to various
correlations.

Measuring Option Risks

Shortcomings

It assumes,
unrealistically,
parallel shifts in the
curve.

Same as Rhol.

Convexity, a much
abused term,
assumes parallel
shifts and moves of
equal volatility.

It does not take into
account local
volatility between
two points in time
and space (i.e.,
forward volatility
and skew slope).

7

*A bucket vega is the vega exposure between two dates.
*The correlation vega is covered in Chapter 22.

113

Modification

Use a simple term
structure model for
the expected
position.

Same as Rhol.

Use a simple yield
curve model (a
variance-based one
factor model) for the
weightings.

e Assumption 5 has two results. The first is that indeed the rate of drift
moves (interest rates are far from constant), and it is often correlated
with the movement in the asset-prices (in “biased” assets like Mexico
the link is obvious). The second result is that the drift does not move
in parallel, but in a predictable manner.

The table on pages 112-113 lists shortcomings of the Greeks
and simple modifications to correct them. This will enable use of the
[End OCR]*

## Page 128

*[Image OCR]
114 Measuring Option Risks

Option Wizard: Why Good Models Die

Most new models that attempted to correct the failings of Black-Scholes-
Merton did not survive. An admirable and realistic one, Merton’s jump diffu-
sion pricing tool, is rarely implemented for the simple reason that it requires
the estimation of two additional parameters, the Poisson jump size and its fre-
quency. Stochastic volatility techniques (see Hull and White, 1987) also were
undeservedly consigned to the dustbin of business school libraries because of
the need to estimate two additional parameters, volatility of volatility and the
correlation between the volatility and some indicator of the asset price. The
same problems mar the implementation of potent yield curve models, such as
those fitting the Heath-Jarrow-Morton framework; traders tend to avoid them
despite the insistence of their research staff, in favor of the simple Black-
Scholes-Merton that they know how to trick.

Traders are not fooled by the Black-Scholes-Merton formula: The existence
of a “volatility surface” is one such adaptation. But they find it preferable to
fudge a parameter, namely volatility, and make it a function of time to expiration
and strike price, rather than have to precisely estimate another.

various derivatives of the Black-Scholes-Merton formula in the light
of the previously outlined reservations.

Technical details about the Black-Scholes-Merton derivation are
provided in Module G. For a very pedagogical derivation the reader
should refer to Hull (1993), Cox and Rubinstein (1985), or Jarrow and
Rudd (1983).
[End OCR]*

## Page 129

*[Image OCR]
Chapter 7

Adapting Black-Scholes-Merton:
The Delta

It is always preferable to be roughly hedged against a broad set of eventualities
than exactly hedged against a narrow parameter.

Marty O’Connell

@ Delta means the sensitivity of a derivative price to the movement in the
underlying asset. It is either expressed in percentages or in total
amounts. A 50% delta is supposed to mean that the derivative is half
as sensitive as the asset and that one needs two dollars in face value of
the derivative to replicate the behavior of one dollar of the asset.

A delta is expressed as the first mathematical derivative of the product with
respect to the underlying asset. It means that it is the hedge ratio of the
asset for an infinitely small move. Somehow, when the portfolio includes more
than one option, with a combination of shorts and longs, delta and hedge
ratio start parting ways. "

The delta is not necessarily limited to options and contingent claims. It
can be used for forwards, futures, and other linear products, where its ac-
curacy is greater. Linear means nondynamically hedged, and dynamic
hedging is what makes the delta very murky. The delta for a forward would
take into account the discounting of the cash flow to achieve equivalence to
a cash product.

Perhaps the largest misconception in the financial markets attends the
definition and meaning of the delta. Every operator instinctively knows
that hedging in continuous time will never be possible. The difference between
discrete and continuous swells when one looks at special situations such as
a risk-reversal or barrier options. r

Attempting to give the delta a meaning in terms of risk management is
denying the dynamic interaction of parameters. It partakes of the desire
by nonprofessionals to get a numerical exposure at no cost. Nothing can be
schematized in dynamic markets. Delta loses in its significance when
traders move from a simple option to a portfolio of longs and shorts. Many
traders to this day live under a delta limit rather than a more global sce-
nario analysis.

115
[End OCR]*

## Page 130

*[Image OCR]
116 Measuring Option Risks

Risk Management Rule: Continuous time models should be used

for pricing and getting a benchmark fair value, not to hedge.

CHARACTERISTICS OF A DELTA

Figure 7.1 shows the delta for a simple strategy. The call with 100 strike
price is priced at 15.7% volatility, the delta is originally 50. It appears to be
the tangent to the option price, and shows the slope of the price change at
the origin, 100.

In the case of Figure 7.2 it is easy to ascertain that the delta in fact is a
hedge about the origin, 100, and that the hedge would need some adjust-
ment where we see a gap between the slope and the price curve.

THE CONTINUOUS TIME DE Ta Is Not
ALWAYS A HEDGE RaTiIo

The delta generally interpreted to mean an equivalent spot position be-
comes inapplicable to the management of a portfolio. Seasoned profession-
als ignore the current definition of the delta as both a measure of risk and
an indication of a position equivalence (although their research depart-
ments and textbook-trained risk managers usually do not). For anything be-
yond a medium-length stable option close to the money, a delta does not
reflect anything meaningful.

Price of a simple call

9.00
8.00
7.00
6.00
5.00
4.00
3.00
2.00
1.00
0.00

P/L

Cash Price

Figure 7.1 Delta for a single option.
[End OCR]*

## Page 131

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 117

200

100

50

PAL

-100

-150

Cash Price

Figure 7.2. Delta and other lies.

Option Wizard: The Half Billion-Dollar Delta

The problems with the meaning of a delta are compounded with barrier
options. . .

An interesting illustration of this microscopic change is an argument that
reportedly took place (at a major institution) between a risk manager and a
group of traders concerning barrier options. The risk manager (characteristi-
cally suspicious of traders’ explanations) was vetoing the trade on grounds that
the delta reached 10,000% at the barrier. According to his logic, it would put
the position at much above any limit allowable by the firm, a $5 million trade
hitting the equivalent position of half a billion! So he prevented the trader from
engaging in the transaction.

The trader was infuriated by such logic. His maximum risk on the trade was
less than $400,000 if everything went wrong. True, the delta jumped to abnor-
mal proportions between 1.399999999 and 1.40 but without being branded
as a heretic it was difficult to explain to the semi-mathematician-risk manager
that deltas sometimes were an irrelevant waste of computer power.

According to the formula, the trader needed to buy a half billion dollars
around the barrier and then sell them back at or above it to make the transition
smooth. If this maneuver were feasible, it would make barrier options better in-
struments to trade. Leaving the trade alone, and considering it as a single bet
with a positive expected return, appears to be a more conservative approach.
in that light, the delta bears no true meaning. Some traders trick the delta of a
barrier by changing the expiration date, as we will see in Chapter 20.
[End OCR]*

## Page 132

*[Image OCR]
118 Measuring Option Risks

Orthodox definition of the delta:
OF
Delta = ——
elta Ti

F is the derivative F(U,t)
U is the underlying security.

It is the derivative of the option price to the underlying. In plain En-
glish, it would correspond to changes in the option price stemming from in-
finitely small changes in the underlying asset. This concept is extremely
useful for mathematical derivations. From a standpoint of trading, it offers
no significance, for the following reasons:

e There is no such thing as an infinitely small move in the market.

¢ If there were such microscopic moves they would be nobody’s con-
cern.

To make the mathematical continuous-time finance delta relevant, it
needs to be accompanied by the second derivative, the gamma, and at least
a third one, the DgammaDspot. In addition, because volatility tends to
change when markets move, adding the vega to the exposure would be nec-
essary. .

In the real world, the modified delta notion appears to be a more ade-
quate measure. It is defined as:
AF

Modified delta = AU

with the A as the change in option price stemming from a set change in the
underlying security. If the call price picks up .05 points when the underly-
ing asset moves from 100 to 100.1 then its delta will be .05/.10 = .50 or 50%.

However, moving the asset in one direction (up) is not necessarily a
great approximation of the behavior of the function on the way south as
well. A more powerful tool is to use the following:

Delta = —
elta = >

AU™ Au?

>(au- + au)

with AU~ and AU, respectively, down-moves and up-moves in the under-
lying asset.

One can see the results of such derivation: Delta would then depend on
the magnitude of the changes in the underlying security, that is the AU. The
increment becomes then at the discretion of the operator. It could be a func-
tion of either his utility curve or his estimation of future volatility.
[End OCR]*

## Page 133

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 119

The advantages of using the discrete delta is that it incorporates a little
of the second and third derivatives that should complete the mathematical
delta in any form of analysis.

Example. A Misleading Delta: Assume there is a flat yield curve and
that the forward is equivalent to spot.
All options are European and have a one-month maturity.

The option trader has the following position:

¢ He is long $1 million of the 96 calls (delta .824, total continuous delta
$824,000 long). ,

¢ He is short $1 million of the 104 calls (delta .198, total continuous
delta $198,000 short).

¢ His total continuous delta is long $626,000.

¢ He could hedge it by selling $626,000 of forward.

Table 7.1 shows his performance. Table 7.1 shows the delta of the position
positive everywhere except around 100.

It is apparent that the P/L on display in Figure 7.2 is similar to that of a
simple long position, except around the origin—and the origin is something

Table 7.1 Inapplicability of the
Continuous-Time Delta

Asset P/L Delta
Price 000 000

92.5 —122 * 420
93.5 ~ 84 349
95.5 —30 194
97 -9 90
98 -2 39
99 0 8
100 0 0
101 1 16
102 4 53 -
103 12 106
104 26 171
105 46 . 241
106 74 311

107 108 377
108.5 171 461
[End OCR]*

## Page 134

*[Image OCR]
120 Measuring Option Risks

Cash Price

Figure 7.3. Modified delta hedge.

nobody cares about. At an option seminar, when the author asked the
crowd whether they considered the graph to reflect a long, square, or short
position, all the traders answered “long.” Characteristically, the persons
without trading experience answered “square.”

The following paradox is taking shape: The trader in the example buys
$550,000 cash instead of $626,000. The answers of the seminar participants
who were shown Figure 7.3 were distributed between “square,” “short,”
and a few “long”:

¢ Those looking at the extremes (asset price 93 or 107) saw a positive
P/L in a rally and a negative one in a selloff. The position in their
eyes “felt long.”

¢ Those looking at a narrow range (asset price ranging between 99 and
101) “felt short.”

¢ Those looking at a middle range (asset price 95 or 103) “felt squarish.”

It can be seen that the hedge shown in Figure 7.3 and Table 7.2 is more
effective than the first one for wider variations: Maximum loss drops to
—65 from —122. Maximum profit drops to 107 from 171. Using a broad set
of assumptions, Position 2 is more delta neutral than Position 1 although it
seems delta short. For a A of 2.5 points, the maximum is entirely neutral.
Taking this position thus widened the neutrality increment to what ap-
peared to be a decent margin.
[End OCR]*

## Page 135

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 121

Table 7.2 A Modified Delta Hedge

P/L Delta
000 000
92.5 —65 344
94 —22 235
95 —2 157
96 10 81
97 14 14
98 13 —37
99 7 —68
100 0 -76
101 -7 -60
102 —il —23
103 —11 30
104 5 95
105 8 165
106 28 235

108.5 107 385

Had the author asked the seminar participants: “How would you feel if
there was a threat of a nuclear attack?” they would no doubt have answered
“long.” And to the question “If your time frame were six minutes?” the an-
swer would have been “short.” This leads to the following rule:

Risk Management Rule: A delta depends on the operator’s per-
ception of future volatility and his utility, as well as his possible

frequency of adjustments.
These elements matter less and less when the market gains in
liquidity.

Operators are forced to define the delta as a function of increments that
matter to them. The preceding position would appear square if the P/L cor-
responded to pennies, but it would require attention if it corresponded to
millions. Somehow, operators are obligated to let the utility curve creep up
into their trading.*

DELTA AS A MEASURE FOR RISK

Delta fails adequately to measure risks even taking a simple position. It pro-
vides the same measure to an extremely risky position and an extremely
safe one.
[End OCR]*

## Page 136

*[Image OCR]
122 Measuring Option Risks

Cash Price

Figure 7.4 P/L for a long/low delta call.

Example: Two positions express the same view: One is long option, the
other short. Initially they have the same delta.

Case 1: A trader is long 1,000 calls
Case 2: A trader is short 1,000 puts of the same delta

Figures 7.4 and 7.5 reveal the difference.

Cash Price

Figure 7.5 P/L for a short/low delta put.
[End OCR]*

## Page 137

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 123

The trades displayed in Figures 7.4 and 7.5 both have the same delta:
$200,000 approximately. It is easy to see why the delta is a poor approxima-
tion of risk.

CONFUSION: DELTA BY THE CASH OR BY THE FORWARD

The delta as expressed by the Black-Scholes-Merton formula concerns the
amount of cash the operator needs to execute to offset an option position.
For all European options, however, the real exposure lies in the forward.

Nevertheless, operators prefer to see the cash delta as they generally
hedge themselves with it. It is easier to monitor on a screen and quote in the
market. When they deal with options on futures or use the futures as a
hedge, they need to use a different delta fit to the exact period in the future.
The difference between the two is sometimes far from trivial. Operators
often must deal with questions like this one: An option that is close to the
money in the forward trades at 50% delta. What is the cash delta?

The answer is to transform the potential future exposure into a cash ex-
posure through the delta of a forward. This can be done by discounting the
forward exposure using the cash-future growth rate as a discounting factor.
Therefore, the delta of the cash will be the discounted value of that number.
The discounting method will depend on the underlying security of the op-
tion, as will be described. .

DELTA FOR LINEAR INSTRUMENTS

It is necessary to examine the deltas for linear instruments both as a hedge
between futures/forward, or even futures and.futures, or as a hedge for
any of the two with options or vice versa.

For a view of the cash and carry arbitrage relationships, traders can
read DeRosa (1992; 1996).

Delta for a Forward

A Non-Interest-Bearing Asset. The general formula for the computation
of a forward is:

F=et'S
for a non-interest-paying asset.

F is the forward, S the spot, t the time to expiration (a shortcut for t — £,),
r the domestic rate, rf the foreign rate, d the dividend rate.
[End OCR]*

## Page 138

*[Image OCR]
124 Measuring Option Risks

Foreign Currency Forwards. The covered interest parity arbitrage for-
mula is (typically the first and often only formula a trader learns—and
needs):

F=et—fts

However, the forward does not immediately deliver the profits and
losses. The operator has to wait until the settlement day to earn or pay the
realized sum. It is easy to see that a profit on a one-year forward will turn
to cash one year from now; therefore, the value of the nominal asset should
be discounted like a zero-coupon bond.

Thus, the P/L generated by the forward needs to be discounted back to
the cash using the usual e~".

Therefore, using AS as move-in spot:

P/L of Forward = ee” ~'f# AS = eft AS

Delta of a Forward = e~f!

Ironically, only the foreign rate matters. Even more ironically, the delta
being the discounted foreign rate, there is some relativism as far as the
hedging. Two people looking at the same forward would use two different
hedge ratios based on their home currency. Module C on the numeraire
problem, provides a discussion of the.issue.

Foreign Cross-Currency Pairs. A foreign cross-currency pair is a contract
to exchange one currency for another at a determined point in the future. An
example would be sterling/DM for an American citizen, or sterling / Dollar
for a German citizen. Not much volume is generally transacted between
these pairs, but by virtue of the multiplication of the combinations, there are
too many of them. Some traders quote Drachma/Australian dollar or
Lira/New Zealand dollar though not routinely so.

Example: Assume there is Currency 1 and Currency 2 with the spot ex-
pressed in units of Currency 2 per Currency 1. The forward therefore is:

F= etl — r2)t Ss

with rl the rates in Currency 1 and r2 the rates in Currency 2.

Now the trader is confronted with the currency in which the profits
and losses are computed, as they will determine how he would discount
the hedge. The trader has the choice of using Currency 1 or Currency 2
as a discounting one. This difficulty arises because unrealized profits
and losses from a sterling/DM trade can be either in sterling or in DM.
The trader cannot translate unrealized profits and losses into his home
[End OCR]*

## Page 139

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 125

currency without incurring a foreign exchange exposure in one of the
currencies.
Discounting the hedge would give us:

Delta of a Forward = e ~"1~ 7)
ry is either rl or 12, whichever the trader chooses as his “anchor” cur-
rency (the currency in which the profits and losses are computed).

Stock or Stock Index. For a stock or stock index that continuously pays
dividends (something that naturally does not exist but that is an assump-
tion), the delta of a forward is the same as with the foreign currency except
that the trader uses the dividend rate instead. The result is:

a)t dt

Delta of Forward = e te" —% = e~
with d the dividend rate.
When the dividend rate is not continuous, one needs to use the exact

payout to the end date.

Delta for a Forward-Forward

The forward-forward can be easily defined as an arithmetic sum of one
long, one short position in the same instrument at two different dates. The
complexity sometimes arises from the lack of equivalency between amounts
(2 years vs. 3 years will need to be “tailed” properly to account for the pre-
sent value).

Example: The delta is the difference between both individual deltas.
Take F1 and F2 as two forwards with delivery t1 and #2. The forward-
forward will therefore be called F(f1,#2). 3

The delta of F(#1,t2) will be equal to Delta F (#2) — Delta F (1).

The profits to t1 need to be discounted at a different rate than the
profits to t2. In addition, this gives rise to interpolation issues that will
cause some difficulties in determining the exact rl given f1 and 12
given £2.

A more complex approach would be to weigh the deltas of each for-
ward since it is not exactly the same position.

Delta for a Future
The general formula for the computation of a future is:
Future = e” Spot

for a non-interest, non-dividend-paying asset.
[End OCR]*

## Page 140

*[Image OCR]
126 Measuring Option Risks

However there is a serious difference between futures and forward as
the futures settle daily in cash a variation margin, which eliminates the dis-
counting the trader applies to the forward.

¢ For a foreign currency, the covered interest parity arbitrage formula
is:

Future = Fe? ~ S's,

Delta of Future =e” ~"/?!

* For a foreign cross-currency, the same holds. Assume there is Cur-
rency 1 and Currency 2 with the spot expressed in units of Currency
2 per Currency 1. The forward therefore is:

Delta of a Future = e~

¢ For a stock or stock index dividends, continuously paying the delta
of a forward is similar to that of the foreign currency except that the
dividend rate is used instead. The result is:

Delta of Future = e*~”!, with d the dividend rate.

When the dividend rate is not continuous, the exact payout to the end date
is used.

Stability of a Delta for a Linear Derivative. The deltas of futures, for-
wards, swaps, and the like are stable (since they are called linear deriva-
tives). Their second derivatives are almost equal to zero (i.e., no gamma,
except for a small convexity) and they have only one meaningful first deriv-
ative, the delta. No modification is therefore necessary to compensate for
the shortcomings of such approximation, except in cases of extreme convex-
ity in very volatile markets.

DELTA AND THE BARRIER OPTIONS

Often dealers have their palms sweating when confronted with the deltas
close to 10,000% when they come close to the barrier nearing expiration day.
In some cases, numbers swell beyond the screen’s ability to display the
numbers. This issue is examined more closely in Chapters 19 and 20, but
often this jump is only applicable for a small increment.

It is indeed erroneous to try to hedge according to that schedule. Get-
ting rid of that notion makes barrier options easier to trade.
[End OCR]*

## Page 141

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 127

Barrier options present an extreme case of discontinuity that is very in-
structive for convincing anyone about the flaws of using a continuous time
framework for option analysis.

DELTA AND THE BUCKETING

@ A bucket is the bundling of exposures by groups of neighboring
maturities.

Another limitation to the delta measure is that it does not display the risks
in the “basis” or the risk between buckets. The delta in the cash is some-
times insignificant when one is confronted with a volatile basis or a long
time separating the cash from the future. The risk manager therefore needs
to see the “bucketing.” This is due to the volatility of both the volatility (it
can impact back-month deltas) and the interest and carry rates.

DELTA IN THE VALUE AT Risk

The reader needs to review a more thorough explanation of the value at risk
(VAR) concepts in Module E. A (poor) method of incorporating of the deriv-
atives risk in the value at risk method is the following”:

Equivalent position = face value X (delta x expected move
+ .5 gamma (expected move)*)

In the case of the risk reversal in the Figure 7.1, the total exposure for
such a position, according to the VAR, would then be zero. Including a
gamma did not help much since the gamma is nil.

Some risk managers go the extra mile and add the vega to the analysis.
Would it improve the risk measure in this example? Of course not: Despite
all its hidden risks, the trade is vega neutral. In addition, the measurement
needs to include the shadow gamma to be complete, in order to incorporate
the vega of the position. At the time of this writing, current methods for as-
sessing the risks of option portfolios resort to the repricing of an option
portfolio at simulated large movements.

DELTA, VOLATILITY, AND EXTREME VOLATILITY

All operators in options learn that a rise in volatility would cause the delta of
an out-of-the-money call to rise and that of an in-the-money call to drop,
[End OCR]*

## Page 142

*[Image OCR]
128 Measuring Option Risks

Option Wizard: Delta and the Probability of Being in the Money

(This issue is examined in greater detail in Chapter 17.)

The delta is the “risk-neutral” (see Module B) replication of the option. The
value corresponds to the integral of the payoffs between the strike price and
infinity for a call (or zero for a put) assuming the underlying process is risk-
neutral. In discrete terms, it would be the sum of the payoff of every eventuality
multiplied by its risk-neutral probability. The delta is the sensitivity of that value
to the changes in the underlying asset.

More practically, it corresponds to the ratio of the asset the trader must
carry to avoid having any instantaneous P/L from the micromoves.

The probability of ending in the money is simply the discounted probability
divorced from the associated payoff. In the study of binary options, it will be
shown that the delta and the probability of being in the money would be the
same at very low volatility and in the absence of a skew, and would start di-
verging in the presence of either a skew or a high volatility (as it could cause a
right-side skew, owing to lognormality, see Figure 7.7).

Barrier option traders who experience deltas in 500% magnitudes are fully
aware, albeit intuitively, that delta means replicating quantity not probability.
They know that the scary delta corresponds to the amounts to buy or sell to
protect their book from losses and maintain their comfortable lifestyle.

We can also extend some notions from the two-country paradox: The
probability of being in the money using one part of a pair as a numeraire is
the delta for the party using the other side as a numeraire. So in a vanilla
DEM-USD option, a delta for a USD-based person is the probability of being
in the money for a DEM-based person. This striking paradox is discussed in
Chapter 19 and Module C.

therefore bringing deltas closer to 50% (or the present value of 50%, to be pre-
cise),

In the formula, as shown here, asset S follows the geometric Brownian
motion:

S, = S, exp{(w ~ 50?) t + 0 Vt Z}

S, is the price of the stock at time 0 (the present), the (risk-neutral) drift
(interest rate differential or numeraire rates less carry), o the volatility and
t time until expiration.

Z follows a reduced centered normal distribution such as p(z = x) =
exp(—x?/2) /V2 7.

We have the conditional expectation E(S,) at time 0 = S, exp( t).
[End OCR]*

## Page 143

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 129

Table 7.3. Asset Values One Period Ahead
a
o (%) -3 —2 —1 0 1 2 3
10.00 90.11 93.28 96.55 99.94 103.45 107.08 110.84
15.70 84.88 89.60 94.59 99.85 105.41 111.28 117.47
25.00 76.91 83.84 91.40 99.85 108.60 118.39 129.05

50.00 58.72 69.77 82.91 98.52 117.07 139.12 165.31
100.00 33.47 47.26 66.73 94.22 133.04 187.86 265.27

The rise in volatility has the tendency to increase the expected final
price of the stock through compounding, while the drift would make it de-
crease by .5 a t.

Table 7.3 shows exactly what takes place when we calculate the asset
price on a spreadsheet. He assumes that t = 12% of one year and plugs in
.12. We also assume that w = O and that S, = 100. The volatility (the first col-
umn) is expressed in annualized terms (to match f, which is expressed in
annualized terms as well).

The Z values share the same probability. So if the center column were at
100, in all cases the expected values of the final asset S would be higher
than S,. There is a correcting term—.5 o? t to satisfy the martingale prop-
erty: Every cell multiplied by its probability needs to sum up to S,. Note
how —.5 o? t pulls down the market inthe center column (Z = 0).

The compounding is caused by what is called the geometric return:
Since returns compound, a higher volatility would raise such compounding
and would increase the divergence between an arithmetic process (where
the returns are constant) and a geometric one (where the returns depend on
the level of the asset). The net effect between these two counterbalancing
factors is the thickness of the right tails. This effect will be examined again
in Chapters 17 and 18. Figure 7.6 shows the graph of the terminal values of
S against their probability.

At a higher volatility, as shown in Figure 7.7, the distribution develops
an increasing skew to the right. The mean, however, remains the same: The
surfaces on each side of 100 are equal. It means that the median needs to
slide left in an amount commensurate to the volatility level. This shift will
have an effect on the delta.

Example: The forward trades at 100, spot at 100. Table 7.4 lists the
deltas of the 110 calls, the 90 calls, and 90 puts, all with 180 days until
expiration. Assume interest is at zero for simplification.

Volatility does not often rise to 180%. An nonequity option trader would
see an actively liquid instrument do so only once every few years. However,
these instances provide the option trader with an accelerated tutorial.
[End OCR]*

## Page 144

*[Image OCR]
130 Measuring Option Risks

Probability
Density

0 100 200 300
Asset Price

Figure 7.6 Asset distribution at medium volatility.

The following is a list of recent triple-digit volatility cases.

1987: Silver, Equity Indices, Eurodollars.

1990: Oil and related markets (Gulf War).

1992: Short sterling.

1995: Mexico (short-term options traded reportedly at 250%), Euro-yen.

1992-1993-1995: PIBOR (Paris Interbank Borrowed Rate) (French Euro-
deposits experience routine panics to the point of making junior option
traders extremely aware of la lognormalité).

This hints at a common Wall Street issue: A call on a bond can also be
viewed as a put on yield. For pricing purposes, traders use a geometric

4

Probability |
Density

+ t 4
0 500 1000

Asset Price

Figure 7.7. Asset distribution at high volatility.
[End OCR]*

## Page 145

*[Image OCR]
Adapting Black-Scholes-Merton: The Delta 131

Table 7.4 Deltas with Extreme Shifts in Volatility

90 110 90

Put Call Call

VOL Delta Delta Delta
10 —0.06 0.09 0.94
15 —0.15 0.20 0.85
20 —0.21 0.27 0.79
30 —0.27 0.36 0.73
40 —0.30 0.42 0.70
50 —0.32 0.46 0.68
60 ~0.32 0.49 0.68
80 ~0.32 0.54 0.68
100 —0.31 0.59 0.69
120 —0.29 0.62 0.71
140 —0.27 0.65 0.73
160 —0.26 0.68 0.74

180 —0.24 0.71 0.76

Brownian motion for both, which is incoherent: Both yields and prices can-
not be lognormal and have a fat right tail at times of high volatility. Aug-
menting this contradiction is that traders often mix the same instruments
in the same book. Bond futures, for example (baring the embedded outper-
formance option), priced on a log-Brownian price, are mixed with bonds
priced on a log-Brownian yield and put together as hedges in the same
book. This incorrect mixing can carry serious consequences for a large book
at a high volatility, as the delta difference becomes acute.

The delta and partial-delta concepts for multivariate options will be
dealt with in Chapter 22.
[End OCR]*

## Page 146

*[Image OCR]
Chapter 8

Gamma and Shadow Gamma

One day in mid-1994, the dealing rooms in the United States were rocked with the
news of the bankruptcy of a hedge fund, costing a minimum of $600 million to
their investors. What worried the community was that the blown-up fund was
meant to be “market neutral.” Market neutrality, it had appeared, would be a
great panacea in the volatile world of tightening policy and distorted relationship:
In theory the fund would warehouse cheap securities, hedge them, and achieve
above-average returns for the Florida residents.

One trader was asked by his manager to explain the results. He shouted:
“That guy did not get the second derivative right.”

SmmPpLeE GAMMA

@ The gamma is the second (mathematical) derivative of the derivative
with respect to the asset price. It is easily calculated analytically with
the following:

a 2E
ol?

The unevenness of the gamma in space (i.e., with time moving and the asset
price remaining constant) has these effects:

¢ For an at-the-money option, the gamma is maximum when the op-
tion nears expiration.

¢ For an out-of-the-money option, the gamma is maximum when the
option is far away from expiration.

This time dependence of the gamma has some consequences for calendar
spreading, as shown in Figure 8.1. If a trader buys option A and sells option
B, the gamma will be positive at the money (Figure 8.1 shows the A line
higher than the B line), but the lines cross at some level. The best comparison
is between a short-distance runner and a marathon runner. The marathon
runner would win the long-distance race. The short-distance runner would
win a 100-meter dash. In between, there is a race of a certain length where
they would be of equal speed.

132
[End OCR]*

## Page 147

*[Image OCR]
Gamma and Shadow Gamma 133

90 100 110 120

Figure 8.1 Gamma changes with time. A is a 3-month option; B is a 6-month op-
tion. Both have a 100 strike.

Figure 8.1 suggests there are important pitfalls for option replication
and portfolio stabilization: Often operators hedge their gamma with an op-
tion trade that takes care of their immediate need (defined in a narrow re-
gion around the origin) but does not provide long-term stability to the
position. They look at their gamma reports and buy the exact, necessary
protection; such protection might be short-lived in a bursting market.

Risk Management Rule: A range needs to be associated with

every gamma measurement.

GAMMA IMPERFECTIONS FOR A BOOK

More even so than the delta, the gamma measure is often too narrow to dis-
play the results of the actual move in a logical increment. For a book, this
measure loses in accuracy because the multiplication of options causes the
gamma to increase in “locality” —to depend more and more on a particular
spot range. A gamma could be long at 100, short at 101.65, long again higher
and so forth, according to which structure dominates at a given spot point.

A practically sound way to measure the gamma is. to vary the underly-
ing price and calculate the actual change in hedge ratio over the increment.
The operation needs to be performed twice:

1. To display the “up-gamma” by moving the price north and comput-
ing the change in delta.

2. To display the “down-gamma” by moving the price south and com-
puting the change in delta.
[End OCR]*

## Page 148

*[Image OCR]
134 Measuring Option Risks

Option Wizard: More Pitfalls for Risk Managers

Risk managers without market experience commonly blame traders for incor-
porating the third derivative in the analysis. However, traders need to use as
powerful a measure at all stages of the game even if their designation of
“gamma” could embrace more than its theoretical scope. Traders call delta
something that has a fittle gamma in it and gamma something that has a little
Ddelta-Dvol in it.

Table 8.1 Up-Gamma and Down-Gamma

Up- Down-
Delta Gamma Gamma
93.5 —36 0 0
94 —35 1 0
94.5 —35 1 0
95 —35 1 1
95.5 —34 2 1
96 —33 3 1
96.5 ~32 , 4 2
97 —30 5 3
97.5 —28 7 4
98 —25 9 5
98.5 —21 12 7
99 —16 16 9
99.5 —9 20 12
100.0 0 25 16
100.5 11 30 20
101 25 37 25
101.5 42 44 30
102 62 52 37
102.5 86 61 44
103 114 71 52
103.5 147 81 61
104 185 91 71
104.5 228 101 81
105 276 111 91
105.5 329 121 101
106 387 131 111
106.5 450 139 121
107 518 146 131

107.5 589 152 139
[End OCR]*

## Page 149

*[Image OCR]
Gamma and Shadow Gamma 135

@ The up-gamma (or right gamma) is the discretely computed change in
delta should the asset price move higher by some definéd increment. The
down-gamma (or left gamma) is the same with the asset moving lower.

Averaging both and coming up with a total aggregate is deceiving. Ina
risk reversal (see following definition) situation, the gamma will be positive
one way and negative the other. Netting the two would show the gamma
to be deceivingly square. Using the up-gamma and down-gamma would
display the third-order risk.

@ Risk reversal is any position that has an up-gamma over some incre-
ment of a different sign than the down-gamma.

Example. A Regular Option: The trade is long $2,000M of the 110 calls
with 1 month to go. The initial delta is $36M and the trader goes delta
neutral (a small 1.8%). Volatility is 15.7%. Table 8.1 shows the up-gamma
and down-gamma changes. Anyone looking at Figure 8.2 can see the in-
stability of the rate of change of the delta. This is manifested for every
point by a difference between an up-gamma and a down-gamma.

Example. Case (Risk Reversal): The classical example shown in Figure
8.3 shows that the gamma needs to be measured with respect to the ori-
gin (Table 8.2).

Delta for an Option Strategy

Delta

Cash Price

Figure 8.2 An unstable delta for a single option.
[End OCR]*

## Page 150

*[Image OCR]
136 Measuring Option Risks

Table 8.2 Up-Gamma and Down-Gamma
for a Risk Reversal

Up- Down-
Delta Gamma Gamma
93.5 449 —140 —165
94 375 —126 —153
94.5 309 —112 —140
95 249 —98 —126
95.5 197 —84 —112
96 151 —71 —98
96.5 113 —59 —84
97 80 —47 —71
97.5 54 —36 —59
98 33 —26 —47
98.5 18 —16 —36
99 7 —-7 —26
99.5 2 2 -16
100.0 0 11 —7
100.5 3 20 2
101 11 29 11
101.5 23 38 20
102 40 48 29
102.5 61 58 38
103 87 - 68 48
103.5 119 79 58
104 156 90 68
104.5 198 100 79
105 245 111 90
105.5 298 121 100
106 356 130 111
106.5 419 139 . 121
107 486 146 130

107.5 558 152 139

CORRECTION FOR THE GAMMA OF THE
Back MONTH

Often the calendar spreading gives rise to two different levels of gamma: a
long gamma in one maturity against a short gamma in another one. This
may be stable except that the two maturities might not have the same vari-
ance, since the “basis,” or the difference between cash and futures might
be positively correlated to the cash, which would lead to unequal moves
between forwards. In that case, the static analysis of the gamma would be
misleading.
[End OCR]*

## Page 151

*[Image OCR]
Gamma and Shadow Gamma 137

Delta for an Option Strategy

Cash Price

Figure 8.3. Gamma for a risk reversal.

A correction is necessary because the Black-Scholes-Merton does not
allow for the rates moving. The following simplified example shows such
adjustment.

Example: This is a simple calendar spread between two different ma-
turities in the SP500 contracts on the Chicago Mercantile Exchange
(CME):

True Gamma (in front
Position Delta Gamma _ contract equivalent)

Sep (90 days) Long 2000 1030 100 100
contracts

March (270 days) Short 2000 (1054) (58) (65)
contracts

First Adjustment

The back month could have a lower or higher volatility exposure than the
front owing to the present valuing. When the future curve trades at a dis-
count (i.e., the back month is lower than the front), one unit of the commod-
ity in 6-month maturity would be less than one unit of the commodity in
[End OCR]*

## Page 152

*[Image OCR]
138 Measuring Option Risks

3-month maturity; one contract in 6 months’ worth of exposure becomes
smaller than the 3-month one. When the future trades at a premium, the
opposite effect is expected. However, such effect is weak and swamped by
other factors affecting the volatility of the back-month contract.

Second Adjustment

The back month can also have a higher or lower gamma owing to the stabil-
ity of the basis (the cash-future relationship).

There are multiple methods to check the volatility of the back month
relative to the front:

e Single-factor method: It consists in using the relative volatilities of
every month. The 6-month can be more or less volatile than the 1-
month and the trader would discount the gamma accordingly.

¢ Another method, more complicated, can be derived by construction.
It involves the use of a covariance matrix of the forward curve. Its
basic methodology is covered in Chapter 12.

In the preceding example, the operator discovered empirically that
March had 12% more volatility than September. It means that whenever
September moved by 1 point, March moved by 1.12 points. The operator
needed to be hedged for such discrepancy. He therefore computed the
gamma by multiplying it by 1.12 to get 65 gammas instead of 58.

Risk Management Rule: Gammas of different maturities cannot

be compared without proper adjustment.

This point will be discussed further in Chapter 12.

SHADOW GAmMma!

Often the gamma itself means nothing as the position is sensitive to volatil-
ity changes (or skew prices) and requires richer analysis techniques. This
necessitates the embedding of the expected effect on other factors deter-
mined by the move (such as volatility or sometimes interest rates).

Most practitioners commit the basic mistake of running their positions
without taking into account that moves in the underlying asset are linked
to changes in other elements in the market. Jumps in the market invariably
lead to jumps in volatility. A gamma number that does not factor this ele-
ment would be meaningless.
[End OCR]*

## Page 153

*[Image OCR]
Gamma and Shadow Gamma 139

Risk Management Rule: Whatever can be predicted with confi-

dence needs to be incorporated in the risk analysis. Ignoring them
would make the Greek measures an entirely theoretical workout.

This does not conflict with the desire to keep Black-Scholes-Merton.
Predicting, not modeling, is recommended since predicting allows the
trader to change his mind while modeling freezes the opinion into an inex-
tricable apparatus of estimators.

M@ Shadow gamma is the computation of the forecast changes in delta tak-
ing into account the changes in volatility and its impact on the position
(Figure 8.4, Table 8.3). The position is then reevaluated using new
volatility parameters.

Example: A trader is long the wings (i.e., long out of money options).
He estimates that volatility should rise if the market moves away drasti-
cally, in any direction. (For simplification, assume that the asset is par-
allel and that up-moves in the underlying would cause the same
changes in parameters as down-moves.) He would like to capitalize on
that effect when the option markets are closed by trading the gamma
overnight. ,

Make the initial assumption that at 98 or 102 the volatility will be
one point higher.

At 98.00 overnight in the future, the trader can buy 731 units of un-
derlying instead of 645. At 102 overnight, he can sell 698 units of the

underlying instead of 612.
: Shadow
Portfolio Gamma
Value }

Figure 8.4 Comparing P/L from gamma and shadow gamma.
[End OCR]*

## Page 154

*[Image OCR]
140 Measuring Option Risks

Table 8.3. Shadow Gamma

P/L P/L Delta Delta
Higher Higher

Unchanged Volatility Unchanged Volatility Delta

Asset Volatility (100 b.p) Volatility (100 b.p) Difference
96 279 352 —1,605 —1,730 — 126
96.5 206 273 —1,323 —1,445 ~122
97 146 207 —1,071 —1,185 —114
97.5 99 154 —846 —947 —102
98 61 112 — 645 —731 —86
98.5 34 81 ~464 —532 —67
99 15 59 —300 —346 —47
99.5 4 46 —146 ~171 —25
100 0 39 0 -1 —2
100.5 4 46 145 167 22
101 15 58 292 337 44
101.5 33 80 447 513 66
102 59 10 612 698 86
102.5 94 150 793 896 103
103 139 200 991 1,109 118
103.5 194 261 1,209 1,339 130
104 260 334 1,450 1,588 137

So the trader can sell more futures in the rally, buy more in the sell-
off. The real gamma is stronger than in a one-factor matrix. While the
conventional matrix shows an up-gamma of 292, the shadow gamma is
337. While the conventional down-gamma is 300, the shadow down-
gamma is 346.

The shadow gamma works both ways: By sticking to a tighter future re-
balancing, the short gamma hedger can thus more accurately hedge the
changes in the P/L.

For a change x% in the underlying? if it is assumed that volatility picks
up by Sig(x) (without complicating matters any further by drawing a poly-
nomial).

At point x, with v’<x)<v

Shadow up-gamma (x,) = (Delta (x, V + Sig(x)) — Delta (x, V))/(x — v0)
Shadow down-gamma (x,) = (Delta (x,, V) — Delta (x’, V + Sig(x’)))/(x,- v’)
A more advanced shadow gamma (see Table 8.3) consists in using an

exact association between the forwards and the cash and project more ac-
centuated cash moves.
[End OCR]*

## Page 155

*[Image OCR]
Gamma and Shadow Gamma 141

Example (advanced). Grid of Expected Volatility/Spot Dependence: The
seasoned option trader can look into his memory and assess the impact
on volatility of a market move. The result is the volatility map (Table
8.4). Although this is nothing but a forecast, it is generally better than
the common methods of looking at moves with constant volatility.

It is undeniable that large moves, especially when they take place after a
quiet period, result in a higher volatility. It is also undeniable that volatility
behavior could be predicted for a “biased” asset (asymmetrical assets that
cause anxiety during price drops) as the violent moves down cause an in-
considerable measure of panic. The trader, therefore, should recompute his
P/L at those prices and try to cover with the delta and the gamma positions
that satisfy his risk appetite.

Note. The preceding map linking expected volatility to various asset
prices could be derived from option prices in the market with the skew and
smile analysis techniques. The market cannot be fooled for too long, and it
discovered that large moves are generally accompanied with jumps in
volatility. It also believes that some moves (typically rallies in stock indices)
generally take place with excruciating slowness and will indicate that be-
havior through option prices.*

Table 8.4 Map of Volatility at Different Price Levels

Resulting
Price Volatility
Starting One Day Change
Price Hence (3 Month Options)
105 1
104 * 0.5
103 unch*
102 —0.5
101 —0.2
100 100 unch
99 unch
98 0.5
97 1
96 2
95 3
94 4
93 7 (panic)
92 10
91 10
90 10

*Unchanged.
[End OCR]*

## Page 156

*[Image OCR]
142 Measuring Option Risks

Threshold Shadow Gamma is a variation of the shadow gamma method. It
corresponds to the belief that the expected change in volatility is not a lin-
ear function of the move and requires the use of schedules of moves. (In-
stead of using mathematical methods, it is always recommended to ask
traders to draw scenario analyses.)

SHADOW GAMMA AND THE SKEW

With biased assets (see Chapter 15) the gamma needs to take into account the
behavior of volatility and the movement along the “skew curve.” If there is
an asymmetry in option volatilities between the upside and downside
strikes, the gamma needs to take into account that the volatility of an at-the-
money option may increase or decrease. An indication of future volatility
can be drawn from where the out-of-the-money calls and puts are presently
trading.
Skew gamma is also called “asymmetrical shadow gamma.”

GARCH GAMMA

ARCH is a volatility modeling method that traders avoid talking about. It
corresponds to the realization by econometricians that volatility moves in
clusters. A large move would cause another large move, and a quiet day is
likely to be followed by a quiet day. To traders, there was nothing remark-
able about such insights. However, ARCH provided the framework for het-
eroskedastic thinking in academia.

GARCH gamma by Engle and Rosenberg (1995) is the first academic dis-
covery of shadow gamma. As the markets move, so will the future volatility,
and this information needs to be incorporated into the future delta. The dif-
ference between present and future delta is called GARCH gamma.

Superficially, GARCH gamma sounds like shadow gamma. However,
shadow gamma does not make any statement as to the behavior of the
future market’s actual volatility, whereas GARCH gamma predicts both fu-
ture historical and implied volatility. Shadow gamma is a simple heuristi-
cally obtained forecast of the price of options as a deterministic function of
the path taken by the underlying securities.

ADVANCED SHADOW GAMMA

M Advanced shadow gamma takes into account the trader’s expected
volatility and interest (or carry) rate moves that accompany the changes
[End OCR]*

## Page 157

*[Image OCR]
Gamma and Shadow Gamma 143

in the asset price. In addition, both volatility and interest rate curves are
expected to shift in a nonparallel way.

Advanced shadow gamma could be easily used in a commodity that has
both its price change and its volatility correlated to the interest rates. Cur-
rency pairs in which there is a “weak” side often present such difficulties.
The Mexican peso hardly sells off against the OECD‘ currencies without a
corresponding defensive rise in interest rates and the corresponding in-
crease in volatility. The same applies to the currencies in a “band” (e.g.,
the former ERM)° where the weak parties need to be defended with pro-
hibitively high interest rates to deter the speculators and other scourges of
central banks. This leads to the necessity of running the forward using
more complex factors. The analysis would take place as follows.

If the cash currency drops by 10% in a “biased” asset, in addition to
the rise in the volatility the operator needs to forecast some increase in the
interest rate differential. That widening in the differential would cause
the back month to move further than the front month. Such imbalance cre-
ates additional gamma that could be interpreted as higher or lower.

Should the currency experience the (unlikely) rally in the asset price,
the interest rate differential would narrow but not in the same amount.

Example: Assume the existence of the imaginary currency of Syldavia.
SYL-USD is the symbol for the currency pair.

US Dollar Rd (one year) = 6%.
Syldavian currency Rf = 20%.
Spot price = 100.

The one-year forward price, satisfying the covered interest rate par-
ity formula (100 < Exp(.06 — .20)), would trade at 86.93.

In a crisis, the currency drops to 90. Such a drop would cause the in-
terest rates in the Syldavian capital to rise by 2000 basis points (in the
one-year) (Figure 8.5). As a result, the forward is now at 90 X Exp (.06 —
.40) = 62.79. So while the spot fell by 10 points (10%), the one-year for-
ward took a hit by 24.14 points (27.75%). ,

The assumption in Table 8.5 is that the trader had only one position in
his portfolio, one-year options.

More precision could be brought into the model by repeating the exer-
cise for every expiration (e.g., the 6-month and the 2-years), making the
same forecast then assessing the total impact on the portfolio.
[End OCR]*

## Page 158

*[Image OCR]
144 Measuring Option Risks

Forward Price
A Biased Assets

Term Structure

100
Initial Cash/Future

90 Line

80

70 Anxiety Cash/Future
Line

60

Time to Delivery

Figure 8.5 Behavior of the yield curve at times of stress.

Table 8.5 Map of Volatility and Interest Rate Differentials at Different
Price Levels .

Resulting
Resulting Interest
Volatility Rate
Change Differential
Starting Price One (1 Year Option) (1 Year Forward)
Price Day Hence (3 Month Options), (3 Month Differential)*
105 1
104 0.5 —.75%
103 unch* -5
102 —-0.5 —5
101 —0.2 0
100 100 unch 0
98 unch 0
96 0.5 2
94 1 4
92 2 (panic) 10
90 3 20
88 ? ?
86 ? 2

*Unchanged.
t Trader unable to forecast.
[End OCR]*

## Page 159

*[Image OCR]
Gamma and Shadow Gamma 145

CASE STUDY IN SHADOW GAMMA:
THE SYLDAVIAN ELECTIONS

Syldavia is a country that faces decisive elections as it will be choosing be-
tween an anarchist regime or a Western style capitalist one. Six months be-
fore the elections were announced, when Syldavia was a quiet country,
volatility was at a boring 14% for all maturities. It is expected to return to
that level should the anarchists lose. For some reason, financial markets
tend to prefer Western capitalism and dislike the anarchists.

The volatility is presently at 20% in the one-month options, declining all
the way to 16% in the one-year.

To simplify the case study, the position assumed will be exclusively in
one-month options. The trader has in his book a mixture of strikes, as
shown in Table 8.6. The results of the ballots are due in one hour and the
trader, trusting no polls, imparts no true probability to each event. The Syl-
davian currency presently trades at 100 to the dollar with interest rates at
14% against 6% for the USD.

The trader met with his peers and came up with a road map concerning
the outcome of the elections (Figure 8.6). Using this road map, the trader
will have to ignore the conventional Greeks and limit himself to the analysis

Table 8.6 Risk Matrix and the Outcomes of the Syldavian Elections

P/L Delta
Asset V=14.7% V=20% V = 29% V=14.7% V=20% V=29%
93 - 14 —86 ~—225 334 383 462
—181
94 40 —52 Scenario A 189 282 406
95 52 —29 —143 65» 193 357
96 54 —13 —i10 —32 118 317
97 47 —5 —80 —104 60 287
98 34 —1 —52 —150 20 267
99 18 0 —26 -173 1 258
0 0
100 -8 Starting 15 —175 Starting 260
Point Point

101 -17 1 26 —156 19 272
102 ~-31 4 55 —118 55 293
103 —40 12 85 —57 108 324
104 —41 26 120 25 174 362
105 —34 48 158 127 252 407
—15 339 457

106 Scenario B 77 201 247

107 16 116 250 379 431 510
[End OCR]*

## Page 160

*[Image OCR]
146 Measuring Option Risks

Anarchists
Win
Spot=94
Vol=29%
rd=6%
f=23%
Fwd
92.68

Anarchists
Lose
Spot=106
Vol=14.7
rd=6%
f=10%
Fwd
105.65

Figure 8.6 Outcomes of Syldavian elections.

of his P/L should either of the two states occur. The markets will go from
the preelection state to either of the two outcomes without any trading in
between so the idea of delta rebalancing will have to be forgotten for now.

Table 8.6 allows the trader to look at the conventional delta/gamma
matrix and compare it with the “true” distribution.

According to the conventional theoretical Black-Scholes-Merton delta
notion, the trader starts with neutrality at 20%. Fhe delta shows accordingly
0 at the delta column that corresponds to the combination of 100 in the asset
price and 20% in the volatility.

There is no possible trading in between the states so the trader will have
to derive the shadow deltas and gammas from the exercise. Scenario A cor-
responds to the expected P/L should the Anarchists win and Scenario B to
that if the Anarchists lose. Scenario A presents a high volatility lower secu-
rity price while scenario B presents lower volatility higher security price.

The first discovery traders can make about the position is that although
it presents a Black-Scholes-Merton gamma that is flat (mildly positive in the
rally and negative in the selloff), it is short shadow gamma. A rally causes a
loss of 15 and a selloff a loss of 181.

A more advanced series of assumptions for option trading will be cov-
ered in Chapter 16.

¥
[End OCR]*

## Page 161

*[Image OCR]
Chapter 9

Vega and the Volatility Surface

A loud Italian option trader was known to start singing “volare, volare,” an Ital-
ian popular song, every time the market experienced a panic, to the great annoy-
ance of his peers. When asked for the reason (by the angry author) he explained
that volatility came from the Latin volare, which means “to fly.”

For reasons of convenience, this chapter will cover all issues relating to
the evolution in the volatility curves and forward volatility as well as the
forward-start and other non-time-homogeneous options. The best way to
study a forward-start option is by understanding the sensitivity to for-
ward or local volatility of any structure.

VEGA AND MODIFIED VEGA

M@ The vega (also called zeta, kappa by nontraders) is the sensitivity of an
option to the changes in the implied volatility for a maturity equal to its
stopping time. A nonvanilla option will be more precisely sensitive to
the forward volatility in the market between its inception and its stop-
ping time. Any convex structure will have a vega.

This section of the chapter will focus on the simple vega of a European op-
tion of known maturity.

It is expressed as:

oF
Vega = 0
o is the implied volatility for the maturity matching that of the option, F is
the derivative security. The best way to ascertain it numerically is by repric-
ing the instrument at different levels of volatilities.

Usually the vega is expressed as a discrete measure (i.e., for a discrete
move in the volatility). In addition, many multiply it by the level of volatility
to make it correspond to a set percentage move in volatility level. For exam-
ple, if volatility is 18% and the vega is .5, the option (or the structure) will
pick up 50 cents when the volatility rises by one percentage point to 19%
(and so on...).

147
[End OCR]*

## Page 162

*[Image OCR]
148 Measuring Option Risks

Some people would derive the vega for 1.8 moves in implied volatility
(10%) so they can compare their position risks across instruments by as-
suming that other markets had a 10% vega move as well. Their vega will
then be 1.8 times .5 = .90.

Most vegas decrease with time except that of a lookback and reverse
knock-out whose vegas increase with time under some conditions.

As with the gamma and the theta, it is easy to see that the vega follows
a bell shape, with the maximum reached when the option is at the money
(by the forward) (see Figure 9.1).

Figure 9.2 shows how the rise in volatility would lengthen the tails. It
also gives a hint on the convexity: The vega of at-the-money options stays
the same when volatility rises but the vega of the options away from the
money rises.

Risk Management Rule: The vega of at-the money options is stable
to volatility. Options that are away from the money (in-the-money

or out-of-the-money) are convex with respect to volatility for the
owner and concave for the seller.

The rule is easy to verify: The second derivative of an option price with re-
spect to volatility equals 0 when the strike price equals the forward and be-
comes increasingly positive when thé strike price is away from it.

Warning: The raw vega may be relevant for an individual option but
means little for the risks of an option book. That effect will be explored
in great depth with the study of the term structure of volatility.

Figure 9.1 Vega with time.
[End OCR]*

## Page 163

*[Image OCR]
Vega and the Volatility Surface 149

Asset 129 0.05

Figure 9.2. Vega and volatility.

The reader should by now be familiar with this book’s leitmotiv: what
works for one single option will not work for a book. The conventional
training of people, which consists of toying with the conventional deriva-
tives of the Black-Scholes formula, has a negative effect on their operating
style. Trading an option bears little relevance to trading a book.

In mathematical terms, the book, neutral in its lower moments (see Chap-
ter 11), can easily lose its stability in the higher moments. An option book, we
will see, is not as “compact” as mathematicians believe. It will generally be
neutral in the lower moments and exposed to various risks in the higher mo-
ments. A simple option, though, loses exposure in higher moments.

Vega and the Gamma

Vega is related to the gamma in a strange way since they seem to evolve in
different ways. The vega is the integral of the gamma profits (i.e., expected
gamma rebalancing P/L) over the duration of the option at one volatility
minus the same integral at a different volatility. Intuitively, the vega P/L
that results from the volatility going higher for a long option holder should
be equal to the expected sum of the gamma profits over the period should
the market go his way. It is precisely that difference for a gamma hedger
that should correspond to a higher P/L.

Example: A straddle owner sees his 3-month straddle increase by
$100,000 owing to the volatility in the market moving from 15% to
16%. This means, precisely, that if the new, higher volatility prevailed,
his gamma profits should yield over the next period the same exact
$100,000—barring slippage.
[End OCR]*

## Page 164

*[Image OCR]
150 Measuring Option Risks

Mathematically, it is:
Vega = o t S* Gamma

with S the asset price, t the time left to expiration, and o the volatility.

The Modified Vega

Risk Management Rule: One should not compare, net, or add the

vegas of two options of two different maturities without any
weighting in the measure.

This concept will be explained later in the chapter.

M@ The modified vega is a simplified one-factor model using the variance
of the volatilities broken up by maturities as an indicator of hedging
precision.

The modified vega, while a more potent indicator of risk than the un-
weighted vega, is not recommended as strongly the more advanced modified
forward vega, which will also be described. The reader will be able to choose
between the two methods.

Since all maturities do not react the same way to changes in the percep-
tion of future volatility, there is a need to correct the vega both for hedging
and risk management purposes.

To intuitively understand why the vega exposure for a book needs to be
weighted, consider a position long one-month options for an equivalent
$100,000 vegas and short two-year options for the same vega amount (which
means that the book has a smaller number of two-year options). Assume
that a shock hits the market. Would it be unreasonable to assume that the
volatility of the one-month option would increase more than the back (un-
less information on structural changes indicates that the volatility spurt
would be sustained)?

In other words, if the stock market were to crash tomorrow, one-month
options would be expected to rise considerably but such volatility would
not be expected to be sustained for an entire year, and one-year options
would rise in volatility, but less than shorter term options. The market is ex-
pected to settle down after a while and a longer term option would only be
affected in its early phase. On the other hand, should the market stay dor-
mant for a brief period, it would be foolish to believe that the structure of
volatility changed so that the market lost all of its vigor.
[End OCR]*

## Page 165

*[Image OCR]
Vega and the Volatility Surface 151

The modified vega’ corresponds to the sensitivity of the options portfolio
to nonparallel changes in the general level of volatility. Shorter maturities are
usually more sensitive to volatility barring temporal information. But opera-
tors need to be open-minded: Should they discover that for some reason the
vegas act in reverse, they should calibrate their weightings accordingly.

Ht
Modified vega = >’ VF,
i=1

with V; the vega for the maturity bucket, F, the volatility weight. The compli-
cations in computing a bucket vega are discussed later in the chapter.

The convention is to use a volatility factor pillared around the 3-month,
a medium-term horizon that is liquid enough to satisfy a risk manager. All
options would be compared with that maturity in their sensitivity.

How to Compute the Simple Weightings

The most important step in volatility risk management is the awareness of
the existence of some weights. Most methods used by operators, despite
different levels of complexity, yield similar results.

“Theoretical” Weighting. Volatility weightings by the square root of the
respective nominal duration of every maturity are often called “theoreti-
cal” by traders. It means that long-term volatility is constant and that op-
tions revert back to that known long-term level at the speed of the inverse of
the square root of time (1/V't). Such thinking is grounded on the awareness
of a constant long-term volatility as reflected by the price of the longest op-
tion on the board. Operators are often proved wrong in markets because the
volatility of the SP100 and the SP500 indices has been undergoing constant
dampening. .
Calculation Method. The operator selects a pillar, say the 3-month options
(generally the most liquid maturity) and weighs the exposures in the other

months using a factor of duration V (90/days to expiration):

Example: An exposure in 1-month option will have V (90/30) = 1.73
times the importance of a vega in the 3-month option. An exposure in
the 1-year will have V (90/365) = .50 times the vega in the 3-month. It
means that $100,000 in vega exposure in the 1-month is equivalent to
$173,000 exposure in the 3-month and $346,000 in the 1-year.

Empirical Weightings. Volatility weightings that are derived from the ob-
served behavior of prices in the market are called empirical.
[End OCR]*

## Page 166

*[Image OCR]
152. Measuring Option Risks

Calculation Method. The operator selects a pillar, say the 3-month op-
tions and weighs the exposures in the other months using the relative
volatility of each period. Typical relative volatility can be obtained by tak-
ing the following ratio:

Sum of the absolute changes in one period

Sum of the absolute changes in the other period

Measurements by the two methods tend to be similar, with a hitch: the
weightings can exhibit considerable instability. There is some nonlinearity
to take into account. At times of shock, the front options tend to overreact
and lead: The back-month options usually wait to see if there are structural
changes or if the change in volatility was a simple blip. Smaller shocks can
cause the opposite effect.

Table 9.1 lists some results of a study performed by the author? that
processed one-day changes in volatility (1988-1994, 1400 observations),
using over-the-counter closing data. The study computed the ratio of the
move in the period’s implied volatility to that of the 3-month.

Using 10-day nonoverlapping changes (1988-1994, 1390 observations),
the results shown in Table 9.2 were obtained (in the shape of the ratio of the
move in the period’s implied volatility to that of the 3-month) (Figure 9.3).

So the unbelievable appeared to be true: There is such a strong mean re-
version that the long-term mean seems to drop after a large move upward
and vice versa, as betrayed by the relative stability of the one-year options.

Using currencies in the studies presented some significant advantages
over other financial instruments. There is in the major currencies a devel-
oped liquid over-the-counter currency option market where instruments
are quoted in rolling periods of one month, two months, and so on. In addi-
tion currency options are quoted in implied volatility parameters for an at-
the-money European option, which reduces the risk of a “bad print.” A
“bad print” is a wrong settlement price for an option that distorts the im-
plied volatility calculations. Listed markets tend to have fixed expirations
that complicate the study because it is arduous to follow a constant date. So
on some days, the “front”-listed option would have 15 days until expiration,
whereas on others it would be 33 days, a matter that can hinder the study.

Table 9.1 Volatility Weightings (Daily Observations)

Period DM JY Square Root
1 month 1.84 1.75 1.73
2 months 1.30 1.31 1.22
3 months 1.00 1.00 1.00

6 months 0.60 0.64 0.71
1 year 0.36 0.39 0.50
[End OCR]*

## Page 167

*[Image OCR]
Vega and the Volatility Surface 153

Table 9.2 Cumulative (10-Day) Volatility Weightings

Period DM JY Square Root
1 month 1.68 1.73 1.73
2 months 1.26 1.29 1.22
3 months 1.00 1.00 1.00
6 months 0.65 0.66 0.71
1 year 0.37 0.44 0.50

Warning: While this weighting scheme is applicable to swaptions and
bond options, operators should take proper care in the measurement of
vegas for forward-forward options such as Euros (EuroMarks, Eurodol-
lars, Eurolira, etc.). The buckets correspond to different underlying in-
struments with their own volatility and their own volatility regimes.

A shortcoming is that the weightings may not be stable enough for
the trader to consider them “square.” This leads to a numerical expo-
sure that would take into account the tracking fluctuations between
maturities.

Advanced Method: The Covariance Bucket Vega

A more modern method to assess the vega risks consists in studying the co-
variance matrix of the nonoverlapping forward buckets. Prior to starting
the analysis, it is necessary to define the forward volatility.

ge 2 8
a &

8

Figure 9.3. Volatility cone at long-term volatility 12%.

232.5
382.5
420
[End OCR]*

## Page 168

*[Image OCR]
154 Measuring Option Risks

FORWARD IMPLIED VOLATILITIES
@ A forward (also called forward-forward) implied volatility between
two dates (t1 and t2) represents the expected volatility between the two

periods inferred from option prices.

We will discuss further down the generalization of forward implied to the
Dupire “local” volatility.

Risk Management Rule: it is imperative for traders involved in

path-dependent or deferred-start options to examine their vega
risks in forward buckets.

Computing Forward Implied Volatility

We start by decomposing time in non-overlapping slices that correspond to
traded dates in the market (Figure 9.4).

We define the earliest possible starting date as t,, which, when it is the
present moment, becomes 0. We define o,, ,. as the volatility between two
points f, and t,. So the volatility between f, and t,, would be that convention-
ally quoted in the market as the volatility for the period. Accordingly, the
volatility of a 90-day option would then be expressed a8 9 9) * OG), Would
then be the variance.

If we used equal time slices:

2 22 2 2
Goin = Gon + ase b+ + On — in

[tj be) [3] [4) [5]

Figure 9.4 Time slices. V(3,4) would be the forward volatility between three
months and four months.
[End OCR]*

## Page 169

*[Image OCR]
Vega and the Volatility Surface 155

Or using unequal time slices n, through n, (more adequate since the
markets price narrow buckets for very short dates and wider ones for longer
dates):

1
2 —_— a _ 2 _ 2 _ 2
Fotn — + —f ((E, ton + (ty — E)OR yy +--+ Ew EDO = tn
n 0
With oP “4 the annualized variance between t, _, and t,. The difference

between two points could be months or simply minutes. One can even con-
sider the frame where t, — t, is one hour and t, — t, is one month.

It is therefore possible to infer the “local” volatility between two points
f, and t, while knowing the volatility between 0 and f, and that between ft,
and t,.

on 1

a a _ _ 2
2 fF ((t, t Gon + (f, — £)o;
2 *9

12)

Hence, choosing periods ¢, _, and f, such that there are prices for op-
tions expiring on these dates available in the market, the following can be
derived:

2 _— — —
_ Gio nt, ty) 6 tn —a (t,, -—a to)
tn — a, in (t — +t} )
R . n- @

which traders call forward volatility (or sometimes “forward-forward
volatility”) between t, _, and ft...

Example 1: Derivation of a forward curve.

t, = 0 (the present moment)

ft, = 180 days

a=90sot,_. = 90 days

The volatilities 90 days and 180 days in the market are 17% and 15.5%,

respectively. The notation would use o,, = .17 and 0,,, = .155 which
leads to:

90

_ | .1552(180) — .172(90) __
So9 180 | ee OO 1884

The market is pricing a 13.84% volatility for the period between 3
months and six months.
[End OCR]*

## Page 170

*[Image OCR]
156 Measuring Option Risks

Table 9.3. Forward Curve

Forward

DM Volatility

Vols between

Days (0, days) Buckets
30 16.00 16.00
60 15.30 14.57
90 14.70 13.42
180 13.60 12.40
360 12.85 12.05
720 12.20 11.51

Example 2. Generation of a Forward Volatility Curve: Starting witha
spot volatility curve as in Table 9.3, one can compute the forward
volatility. Figures 9.5 and 9.6 show the contrast between spot and for-
ward curves.

Sometimes the existence of arbitrage is revealed by a significantly lower
or higher forward volatility for a given bucket: Because options are traded
in spot volatility terms rather than forward volatility, it is more difficult to
spot arbitrages and easier for markets to get out of line.

Warning: The forward volatility method is only applicable to options
on fungible assets, such as foreign exchange or equities. Eurodollar op-
tions, for example, are non-overlapping instruments and the use of the
preceding method of analysis on vegas emanating from a Eurodollar
book could distort the risks.

16
15
> 14
eB 13
@
6 12
> 41
10
9 t ++ + | i
oO Oo oO So So So Oo
o ive) Oo © o NW
ba wo ™
Days to Expiration

Figure 9.5 Spot volatility curve.
[End OCR]*

## Page 171

*[Image OCR]
Vega and the Volatility Surface 157

Forward Curve

Volatility

360+

ta

Days

90-180
180-360

Figure 9.6 Forward volatility curve.

A bucket vega is a more potent risk management tool for vanilla op-
tions. However, with path-dependent options, it is the only possible way of analyz-
ing exposures. As will be discussed, vanilla options present a linear bucket
sensitivity to variance whereas path-dependent options present uneven ex-
posures in the time buckets. Some of them, like knock-out options, could be
long vega one bucket and short in another.

Example. How to Translate Regular Vegas into Forward Buckets: Note:
This technique is only applicable to vanilla products.

Assume the following exposure in the DEM options, read from a
“straight” bucket.

Maturities Exposure

0-30 100
0-60 — 148
0-90 17
0-180 167
0-360 —233
Total —97

The preceding numbers resemble those that come out of commer-
cially available risk management systems. Should the trader buy a 3-
month option, the entire position vega would show in the bucket that
ends with the 3-month, in the case above the 0-90.

The operator needs, however, to correct them in some rapid manner.
Assume that path-independent options have exposures that are linearly
distributed through time (e.g., a 3-month option would have a third of
[End OCR]*

## Page 172

*[Image OCR]
158 Measuring Option Risks

its exposure to vega in the first 0-30 bucket, one third in the second
30-60 bucket, and one third in the 60—90 bucket).

This can easily be theoretically proven. It suffices to say that the
gamma is strongest at the end but that the odds of being at the strongest
point are slim, while the gamma is more even, but weaker, at the begin-
ning. It can also be viewed that the expectation of gamma profit and loss
does not vary with the period, that the difference in integrals at two dif-
ferent levels of volatility would correspond to the vega.

The operator must take into account the unevenness of the buckets.
Since he has decided to present narrower slices at the beginning and
wider ones at the end, a one-year option will have:

¢ 50% of its exposure in the 180-360 bucket.

° 25% of its exposure in the 90-180 bucket.

© 8.33% of its exposure in each of the 0-30, 30-60 and 60-90 buckets.
Table 9.4 decomposes the exposure.

Multifactor Vega

This advanced vega method includes the establishment of a forward volatil-
ity correlation matrix. It could be done on a spot volatility exposure pro-
vided the position does not include path-dependent or deferred-start
components.

Without elaborating on the method (it is described at length in Chap-
ter 12), it is possible to establish the notion of expected risk based on nor-
mal behavior of the buckets.

Step 1. The operator builds a correlation matrix of the percentage moves be-
tween forward-forward buckets, say by slicing time into 0-30, 30-60, 60-90,

Table 9.4 Decomposition Steps for Conventional Buckets Exposure

(000) 0-30 0-60 0-90 0-180 0-360
Exposure 100 —148 17 167 ~233
Break-even Break-even Break-even Break-even Break-even Total
0-30 30-60 60-90 90-180 180-360 Exposure
0-30 100 —74 5.67 27.83 —19.42 40.00
30-60 —74 5.67 27.83 —19.42 —60.00
60-90 5.67 27.83 —19.42 14.00
90~180 83.50 ~58.25 25.00

180-360 —116.50 —116.50
[End OCR]*

## Page 173

*[Image OCR]
Vega and the Volatility Surface 159

Table 9.5 Correlation Matrix for Volatility Buckets, DEM
Options 1994-1995 (284 business days)

Bucket 0-30 30-60 60-90 90-180 180-360

0-30 1 0.33 0.25 0.174 0.098
30-60 1 —0.33 0.16 0.154
60-90 1 -0.14 —0.057
90-180 1 —0.19

180-360 1

90-180, 180-270, 270-365, 365-730, and so on. Using historical analysis, the
operator then fills in the correlations between the relative periods (see Tables
9.5 and 9.6).

Statistical issues will appear, as with any correlation matrix of the sort.
The principal one relates to how far back the operator must go to get a ma-
trix that carries any significance. A lengthy sampling period probably
would include periods so far back that the behavior of volatility might have
been entirely different from the present. A shorter one incurs the risk of
missing statistical significance. Another difficulty relates to the choice of
linear methods of statistical association (i.e., correlation) in markets that
may not be tractable with the least-square models.

Step 2. The final exposure is computed with the aid of some simple ma-
trix? algebra, yielding one single number. Such an analysis would disclose,
in addition to the vega weightings, the stability of the exposures. Neighbor-
ing buckets would present a stronger compatibility so the exposures be-
tween them would cancel out more thoroughly than exposures in remotely
spaced buckets. .

Table 9.6 Annualized Volatility of Changes
(per bucket), 1994-1995

Annualized Daily
Volatility Volatility

of Volatility of Volatility
Bucket (%) (%)
0-30 48.80 3.07
30-60 38.74 2.44
60-90 40.47 2.55
90-180 18.65 1.17

180-360 13.53 0.85
[End OCR]*

## Page 174

*[Image OCR]
160 Measuring Option Risks

Formulas. Traders do not need to be daunted by the matrix computations as

these are usually provided by spreadsheet packages. It will be helpful to set

down the formulas prior to using an example that fits the previous matrix.
The volatilities of the forward buckets are:

V1 = Vol(0-30)
V2 = Vol(30-60)
a. V3 = Vol(60-90)
V (daily volatilities) =
V4 = Vol(90-180)
V5 = Vol(180-360)

V6 = Vol(360-730)

The trader then builds the exposure vector, using for input the vegas for
10 percent change in volatility (ie., P/L resulting from volatility moving
from 15% to 16.5%) and multiplying every one by 10 times the correspond-
ing V (daily volatility). Owing to the nonlinearity of the effect of volatility
on a book, the exposure can be better approximated by using a local band,
here 10%, rather than move volatility to zero and obtain a “face value vega”
that would be grossly distorted.

El = Exp(0-30)
E2 = Exp(30-60)
E3 = Exp(60-90)
E4 = Exp(90-180)
E5 = Exp(180-360)
E6 = Exp(360-730)
The operator then builds the forward volatility correlation matrix: The

unit used for the variance corresponds to the percentage changes in volatil-
ity, using the differences of the natural logarithms of the periods:

1 Corr12 Corri6
Corr12 1
M= .
Corr16 1

The transpose of E is multiplied by V, and the result is multiplied by E.
The resulting number would express the net P/L that is expected from a
move of one standard deviation in the general volatility levels,
[End OCR]*

## Page 175

*[Image OCR]
Vega and the Volatility Surface 161

Table 9.7. Bucket Exposure

Daily
Bucket
Vega Volatility

Bucket Exposure (%)
0-30 40,000 3.07
30-60 —60,000 2.44
60-90 14,000 2.55
90-180 25,000 1.17
180-360 —116,500 0.85

Example: Take the previous exposure, as shown in Table 9.7.

The exposure, using the one-factor “weighted” volatility framework,
would show the position to be flat. (Gee Table 9.8.) The trader can analyze
the position using the conventional method by reconstituting the “raw”
buckets. Thus vega(0,90) can be reconstituted with the intermediate
buckets:

Vega(0,90) = Vega(0,30) + Vega(30,60) + Vega(60,90)

The same position presents a value at risk of $21,000. So, using accurate
historical data, such a position is supposed to move by an average of $21,000
per day in spite of being square.

An even more interesting approach is to incorporate, in addition to the
time buckets, possible skew shifts.

Table 9.8 One Factor Bucket Exposure

“Weighted”

Bucket Vega Weightings Exposure
0-30 100,000 1.73 173,000
0-60 — 148,000 1.21 — 179,080
0-90 17,000 1 17,000
0-180 167,000 71 118,570
0-360 — 233,000 55 — 128,150

Total 1,340

Weighted
[End OCR]*

## Page 176

*[Image OCR]
162 Measuring Option Risks

Option Wizard: Curve and Surface Shifts

Any curve (volatility, interest rate, forward prices) can experience the shifts
shown in the following charts:

First Order: Parallel Shift

Price

Time

Second Order: Rotation

Price

Third Order: Convexity Changes

Price

Time
[End OCR]*

## Page 177

*[Image OCR]
Vega and the Volatility Surface 163

A curve therefore can experience any combination of the preceding factors.
The same notions can apply to a volatility surface:

First Order: Parallel Shift

Second Order: Rotation

A Combination

Higher Order: More interesting deformations are experienced, some of
which can be quite ludicrous—but not impossible—as shown in the following
chart:
[End OCR]*

## Page 178

*[Image OCR]
164 Measuring Option Risks

Volatility Surface

M Volatility surface (also called a Dupire-Derman-Kani Surface by
traders) is the representation of the values of implied volatility in the
market represented in function of the time to expiration (horizontally*)
and the strike price (vertically). It is used to derive:

¢ A spot curve displaying the volatility from 0 to time t.

¢ A “local” volatility curve showing the instantaneous volatility at dif-
ferent possible asset price levels at different points in the future.°

It is easily seen that the local volatility is to forward volatility what a di-
agonal spread is to a calendar. The local volatility is the generalization by
Dupire, Derman, and Kani of the forward volatility between two points.®

The Table 9.9 shows the implied volatility as a function of the number of
days to expiration and the strike price of the listed SP500 options as shown
on the Chicago Mercantile Exchange. To prevent the chronic inaccuracies as-
sociated with the settlement price system, these were provided by a market
maker who drew an accurate snapshot of where options were quoted in the
pit at the most liquid times.

Needless to say that put/call parity is respected since puts and calls
typically trade at the same volatility. The American feature is “weak” (see
Chapter 1) and market makers typically use the out-of-the-money leg of any
strike as a benchmark. If the put is out of the money, its volatility will be
used, not that of the corresponding call.

Instead of using the raw strike, it is more meaningful to use the “per-
centage of moniness,” usually defined as the log (asset/strike).’

The first future trades at 585 in this table.

Figure 9.7 shows the volatility surface. It is called the spot surface, not
the forward volatility surface, as described earlier in the chapter.

Figure 9.8 corresponds to a function estimating the volatility as a func-
tion of strike and time to expiration.®

THE METHOD OF SQUARES FOR Risk MANAGEMENT

The same method used to compute the multifactor vegas with the aid of a
volatility matrix by isolating the exposures in buckets and testing the sta-
bility of the hedges can be applied by cutting the position in squares of
strikes and time to expiration. The advantage of such a method is that it dis-
closes the sensitivity of the position to various deformations, such as the
skew stability (Figure 9.9).

The method of squares consists of cutting the position into units com-
posed of small squares and estimating the vega per square. Such a method
[End OCR]*

## Page 179

*[Image OCR]
Vega and the Volatility Surface 165

Table 9.9 Implied Volatility Surface

SP500
SKEW — Skew 11/15/95
Days to December January February March
Expiration 30 67 121 219
Strike  Log(S/K) 30 67 121 219
525 0.11 21.5 18.45 17.43 17.22
530 0.10 20.7 17.95 16.95 16.92
535 0.09 19.7 17.45 16.6 16.60
540 0.08 18.9 16.95 16.05 16.35
545 0.07 18.15 16.45 16.1 16.00
550 0.06 17.25 15.85 15.75 15.8
555 0.05 16.4 15.45 15.4 15.5
560 0.04 15.55 14.95 15 15.25
565 0.03 14.7 14.45 14.5 14.97
570 0.03 14.05 13.95 14.25 14.7
575 0.02 13.45 13.45 13.9 14.42
580 0.01 12.85 13 13.55 14.15
585 0.00 12.35 12.6 13.2 13.92
590 —0.01 11.85 12.2 12.9 13.7
595 —0.02 11.35 11.8 12.6 13.4
600 —0.03 10.85 11.45 12.3 13.25
605 —0.03 10.45 11.15 12.05 13.05
610 —0.04 10.05 | 10.9 11.8 12.85
615 —0.05 9.86 10.55 11.5 12.65
620 —0.06 9.86 10.3 11 12.45
625 —0.07 9.86 10.15 11 12.45
. 630 —0.07 9.86 10 11 12.45
25
20 m 20-25
(1 15-20
15 £ |010-15
%S |m5-10
10 S$ |mo-s
0.14
0.0.
Log(S/K)

. 219
124
Days to Expiration
(Scaled)

Figure 9.7 Volatility surface.
[End OCR]*

## Page 180

*[Image OCR]
166 Measuring Option Risks

Volatility

Log (Fwa/Str}

Figure 9.8 Volatility as a function of the ratio of forward to strike.

is required in “skewed” markets such as the equity indices or the beastly
Euro-Yen where the risks of the shape of the skew are often superior to
those of volatility.

The trader would next run the correlation matrix between the implied
volatilities of the different squares to obtain a clear picture of the risks. This
method is not necessary in well-behaved markets but imperative in skew
products. In the DEM-USD pair, for example, the square s17 is very corre-
lated with the square s57 or any other in the t7 bucket. This is not the case of
the SP where skew shifts can cause some rotation of the volatility surface.

The final vega exposure can be obtained by drawing a giant matrix
and repeating the exercise of the multifactor vega described earlier in this
chapter.

Strike

115}sit |s12|s13 ]si4 |s15 }s16 |s17 |s18 |
110
105
100|s41 |s42 |s43 |s44 |s45 |s46 |s47 1548 |
95|s51 |s52 |s53 [s54 |s55 |s56 |s57 [58 |
90]s61 ]s62 |s63 }s64 |s65 |s66 ]s67 |s68 |
85
80{s81 |s82 |s83 }s84 [585 |s86 [587 [s88 |
[s91 ]s92 }s93 ]s94 1595 [396 |s97 [598 |

tl 2 3 t4 t5 t6 t7 t8 Time

Figure 9.9 Decomposition of the vega by squares.
[End OCR]*

## Page 181

*[Image OCR]
Chapter 1 0
Theta and Minor Greeks

Almost one century ago, a young French mathematician named Gaston Bachelier
had the insight (among other surprising intuitions) to write, in his doctoral
thesis, see Bachelier (1900), that the expected price tomorrow of a call value
was today’s. He gave the right answer fo what most option beginners fail to un-
derstand: that time decay is not the expected P/L from an option. If the option is
priced at the right volatility (assuming interest rates are 0), time decay will be
expected to be 0.

THETA AND THE MODIFIED THETA

M The theta is a loss in time value of an option portfolio that results from
the passage of time.

Theta is called rent by traders. Some traders never go home paying theta:
Many former option traders incur time decay phobia.

The pricing is straightforward: The trader can use the difference be-
tween the price of an option today and the same on the next day, keeping
everything else constant.

One way to look at the representation of theta is that it goes hand in
hand with gamma. The alpha (gamma per theta ratio) will be the same re-
gardless of the number of days to expiration, and so on. Selling very short-
term options, a sport that is periodically practiced by newcomers, would be
an attractive breadwinner except that the risks are exactly the same as sell-
ing longer options, unless the trader sells an expensive strike.

Figures 10.1 and 10.2 show the time decay for an at-the-money and an
out-of-the-money option.

Modifying the Theta

Theta corresponds to the repricing of a portfolio with one day less to expi-
ration and checking the difference between the two prices. However, what if
the volatility and other parameters for period t+1 were different than
those of period tf?

167
[End OCR]*

## Page 182

*[Image OCR]
168 Measuring Option Risks
2.00 < .
ion price

1.50

1.00

Option Price

0.50

0.00

° Nn © Oo FT DBD N ©
~ ony 8 aea aes 8

Days to expiration

Figure 10.1 Changes in at-the-money option price with time.

Example: Thirty-day options trade at 16%; 29-day options trade at
15.8%. Should the trader assume that a position in a 30-day option
would “converge” to 15.8% volatility tomorrow when the position will
have 29 days of shelf life? Or should he assume that the 29-day implied
volatility tomorrow will be at 16%?

There is always the question of whether, like interest rates, forward
volatility is a predictor of future volatility or whether a volatility curve
exists for structural reasons. Should the reader believe that the term struc-
ture of volatility reflects future volatility, that there is valuable informa-
tion embedded in the curve, there is no reason to modify the theta. In
many cases, however, it makes sense to do so. Sometimes, information
about a meeting causes the curve to price at a higher volatility in the days

0.14
0.12
0.10
0.08
0.06
0.04
0.02

0.00
° un © Oo ~~ DO NY ©
ver 2 & A & 8 BS

optign price

Option Price

Days to expiration

Figure 10.2 Changes in out-of-the-money option price with time.
[End OCR]*

## Page 183

*[Image OCR]
Theta and Minor Greeks 169

immediately following the event and the predictive powers of the volatil-
ity curve become incontrovertible. Most times, there is no reason for such
a curve other than the belief in the mean reversion of volatility. In those
cases, modification of the theta should be entirely subjective and left to
the discretion of the trader.

H Modified theta is the price of an option today using current volatility
less the price of the same option tomorrow using the volatility of one
day shorter option.

It is usually interpolated using a square root of time method, to take into ac-
count the possible drop on the volatility curve.

Theta for a Bet

As markets come close to the barrier, theta can be overwhelming (see Chap-
ters 17 and 18).

Theta, Interest Carry, and Self-Financing Strategies

Traders eliminate the interest costs of holding the premium to compute the
theta because the carry of an option should be neutral to a trader who funds
himself. In other words, if the trader incurs carry costs, the price of the theta
will be increased by the interest paid on the premium, making it totally neu-
tral. Thus, if interest rates are 20%, theta will be lower by 20% of the total pre-
mium (the option will have a lower price because of the present-value effect)
but the carry costs of holding the option will offset these savings. It is as-
sumed that the trader has borrowed the money to buy the option and that he
would pay the difference in higher interest.

Many traders erroneously factor the prenmaium costs in the theta
computation.

As a convention, in this book, theta includes no premium costs.

In the derivation of an option value through a self-financing strategy,
the seller of the option is supposed to buy an interest-yielding instrument.
This is equivalent to the option trader funding himself from his firm by
paying for negative balances and earning interest on positive ones.

More controversial is the notion of the forward. When the trader moves
the portfolio up one day for reevaluation, should he move up the spot to the
forward, considering that the expected spot is the forward price? Many
traders do not do so, as it matters little. However, in markets where the
cash-future line is steep, this makes a considerable difference. It seems best
to deal with this issue using the trader’s expectation. If the trader feels that
the expectations built into the forward are erroneous, he needs to analyze
the delta with such belief in mind.
[End OCR]*

## Page 184

*[Image OCR]
170 Measuring Option Risks

The expectation of an asset S, period t, at period 0 was:
EAS,) = S,(e”)

This issue will be dealt with at length in Chapter 12.

Shadow Theta

Losses from theta are generally worse than predicted, as premium buyers
become the victims of additional marks-to-market losses. This is often be-
cause losses from theta, caused by quiet markets, are compounded by losses
from the drop in implied volatility that accompanies quiet markets (as mar-
kets stop moving, the expectation of future movement declines). This is not
always applicable to structures and needs to be measured using appropriate
care and subjective judgment. A quiet market ahead of a political announce-
ment or a refunding session can lead to situations of “reverse decay.” Also,
option prices are sometimes so undervalued that the forecast of the quiet
period is built in.

Example: Thirty-day options trade at a volatility of 16%. It is conceiv-
able that, 7 days from now, if nothing happens, the 23-day option would
be priced at 14%. The shadow theta is the regular theta plus the price
impact of the expected decline in volatility.

Theta is a number that, taken on its own, has little meaning. It theoreti-
cally assumes that nothing else but time moves (volatility and asset price
constant). Traders, however, may have information about the behavior of
volatility if the spot does not move, and the changes in volatility should spot
moves.

Theta becomes much lower when the markets move. This is the same
reasoning as the one used in the shadow gamma analysis. Again, it relies on
an entirely subjective perception of the behavior of implied volatility under
a certain set of conditions.

Risk Management Rule: If the trader believes that volatility may

drop should the market remain frozen such information needs to
be included in the time decay.

Figure 10.3 shows the phenomenon, with the two curves, both produc-
ing the profits and losses at some point in the future. The first curve (“un-
changed volatility”) shows the regular theta. The second curve (“variable
volatility”) shows the shadow theta and a combination between shadow
[End OCR]*

## Page 185

*[Image OCR]
Theta and Minor Greeks 171

Variable
Volatility

P/L

Unchanged
Volatility

Asset Price

Figure 10.3 Theta and shadow theta.

theta and shadow gamma as it is generally accepted that volatility changes
when the market moves.

In the past, the author has tried to measure this effect with some success.
It helps in determining the optimal gamma hedges when the market whip-
saws between levels and thus allows for better capturing of the gamma.
Tighter rebalancing always helps mitigate the accelerated time decay.

Weakness of the Theta Measure -

The major weakness of theta is that it is a path-independent measure. No
true measure of time decay between period t and period t + 1 (ie., t plus 1
day) can forecast the effect of the “whipsaw,” the path where the market
snaps away from its origin then returns to it during the time interval. Such
an event is theoretically improbable but rather frequent in reality.

MINOR GREEKS

Rho, Modified Rho

@ The rho is the sensitivity of an option position to the interest rates in
the numeraire currency as well as the rate determining the cash/future
relationship.

The most misunderstood option risk is the rho. Most of the confusion oc-
curs because it assumes a parallel shift in the curve, and parallel shifts gen-
erally take place exclusively in the imagination of rookie risk managers.

The rho needs to be broken up into its components: the rhop, the rhol,
and the rho2.
[End OCR]*

## Page 186

*[Image OCR]
Option Wizard: The Joys of Interpolation

Yield curves are not given in the market as one simple continuous line. The se-
curities from which the curve is derived correspond to fixed and known matu-
rities. Maturities that fall in between, for mark-to-market purposes in dealer’s
inventories need to be interpolated in some way or another.

The way “broken dates” are interpolated is a matter of significance, partic-
ularly for nonyield-curve option instruments that depend on some interest rate
for input. A currency option requires the risk-neutral yield curves for both
sides of the currency pair. An equity option requires, in addition to the nu-
meraire currency yield curve, the curve for the dividend payout (something so
difficult that equity option traders prefer not to talk about it).

This book will not discuss the techniques of interpolation as they have
been the subject of Byzantine hairsplitting polemics between quants with free
time on their hands and some patience for details.

The first principle used for interpolation is eliminate jaggedness. Visualize
an extreme case of a seesaw yield curve as in the chart on the left:

Maturity

Maturity .

interpolating straight-line on the yield curve in this chart causes some
sharp local peaks. True, the yield curve might be irrational but there are signif-
icant reasons to smooth out the peak: The market cannot be expected to expe-
rience such drops between two dates (except, of course, if there is some
information in between). :

One method, among many others, used to smooth out the curve is called
cubic spline. \t is used by road-building engineers to make sure that their roads
do not experience very sharp turns anywhere and cause accidents.* Applying
such technique would result in the curve on the right, more pleasant to the eye.

The second curve appears smoother but may present some unsatisfactory
peaks that the trader can then correct again. There are mixtures of techniques
that the research assistant would gladly explain at length.

This leads us to the second principle: Adapt to the market. If the market
happens to use some irrational technique for interpolation the trader needs to
conform to it for pricing, though not for the decision-making process.

*The method simply consists of fitting cubic polynomials between nodal points by making sure
that the derivative of each polynomial is equal, at each node, to that of ithe derivatives of both
the preceding and the following polynomial.

172
[End OCR]*

## Page 187

*[Image OCR]
Theta and Minor Greeks 173

@ The rhop corresponds to the risks associated with the financing of the
premium of the option.

® The rhol corresponds to the sensitivity of a portfolio of options to the
base currency interest rates in which the P/L is computed. For a dollar-
denominated equity option, it is the USD rate. For a yen-demoninated
equity option, it is the yen rate. In a currency option pair, it is the base
currency or, if the pair is in two different currencies, say DEM-JPY, the
rhol is calculated for the currency in which the unrealized P/L is ini-
tially accounted.

Rhop is included in rho1.

@ The rho2 corresponds to the sensitivity of the portfolio to the risk-neutral
rate of return of the underlying asset.! For an equity, it is the dividend
payout. For a currency, it is the foreign interest rate, or “counter rate.”

The yield curve moves in parallel only by accident. It could undergo moves
that change its shape in addition to the moves in the market prices. Typi-
cally, yield curve models are built to ascertain the speed of variation of the
relative maturities. Simple models take into account the slope of the curve
as a indicator of future relative speed of movement, while more complex
ones take into account the deformations of the curve (in higher and higher
degrees of deformation). Because this text is not about primary yield curve
risks (it is principally concerned with volatility, not the underlying process)
but has a focus on secondary yield curve risks (the risks are incidental to
the position), simpler methods of calculating the true rho can be used with-
out much loss of precision. The section on stacking in Chapter 12 shows
more sophisticated risk measurement techniques.

@ The modified rho is the application of any simple one-factor yield curve
model, (the relative variance of each maturity) to calculate the true sen-
sitivity of an option book to the interest rates in the market.

A brief examination of interest rates would show that the shorter rates
are often more volatile than the longer ones. While a purblind person could
see that moves are generally higher in the front than in the back, a more
thorough analysis is required to find the factors determining the relative
exposure of one contract compared with another. Yield curve models gener-
ally use some assumptions about the “reversion,” itself a function of the
shape of the curve.

The modification of the Rho should take into account the relative
volatility of the underlying instruments and their correlations.
[End OCR]*

## Page 188

*[Image OCR]
174 Measuring Option Risks

As with volatility weightings, there are several methods for computing
such behavior, ranking from the very simple to those inspired by the richer
Heath Jarrow Morton (1987) general framework of models. Like the vega
modification, a simple method of relative weightings outperforms the un-
weighted rho by a large measure.

Simple weightings are calculated by measuring the average move of one
maturity compared with another.

Reminder: A benchmark is required. The benchmark is a liquid matu-
rity to which the rest of the curve exposures are equated. Traders gener-
ally use the 3-month as the benchmark. The rho and modified rho of the
benchmark are equal.

Example: In computing a simplified modified rho, select the simplest
method of finding “weightings” to each maturity.

Step 1: Breaking up the exposure in “buckets.”

Step 2: Repricing the portfolio with higher interest rates.? The exposure
is in dollars per rise of 100 basis points in interest rates.

Benchmark: 3-month.
0-3 3-6 6-9 - 9-12 12-18 18-24 24-48
months months months months months months months

Rhoi (raw) $100,000 $300,000 $400,000 $600,000 -$750,000 $420,000 — $1,000,000
Modification

(weights) 1.22 1.22 1 94 21 .82 75
Modified
Rho $120,000 $366,000 $400,000 $564,000 — $682,500 $344,400 — $750,000

Total Rhol (raw): $70,000
Total Rho1 (modified): $362,000. Quite a difference.

Omega (Option Duration) *

@ The option duration, also called omega, is the expected life of a soft
path-dependent option. For a barrier option and American binary
structure, the option duration is generally called first exit, or stop-
ping, time.

The option duration is conceptually different for an American option than
for a barrier, because the barrier is an exit price known in advance,
whereas an American option does not provide for a predetermined exer-
cise price or time.
[End OCR]*

## Page 189

*[Image OCR]
Theta and Minor Greeks 175

American options have an expected life that is equal to or shorter than
that of an equivalent European option of the same maturity. So do barrier
options as these instruments can be terminated earlier than the end date.

This aspect of American options is particularly meaningful when one
has recourse to a steeply rising or a steeply declining volatility term struc-
ture. The exercise of an American option depends on the two following de-
cision rules (see Chapter 1):

¢ The Soft Rule: A comparison of the premium over intrinsic value
(determined by the time value of the option) and the earnings from
the reinvested premium in a money market instrument. If the op-
tion premium costs $2 in financing until the final expiration date
and the extra premium of the option is only $1, it becomes optimal]
to exercise the option.

¢ The Hard Rule: A comparison of the premium over intrinsic value
and the costs of maintaining the underlying delta until expiration.
The position in the underlying asset can have a “negative carry” (i.e.,
the currency pair one is short has one currency, the short, yielding a
higher interest rate than the long leg, or a stock paying an extremely
high dividend, much higher than the financing rates in the market).
For example, being long USD-MXP (long dollars against short the
Mexican peso) with the annual interest rates in Mexico for the term
at 50% and those in the United States at 6%, will cost 44% per annum
in carry. A deep in-the-money call on Mexico would force the opera-
tor into the following choice:

Sell Mexico against the deep in-the-money call and incur the costs
of borrowing the currency at 50% and lending dollars at 6%.

Abandon the option as the game isn’t worth the costs. It is more ad-
vantageous to forgo the returns from the possible future volatility.

As shown earlier, the rules of premature exercise depend on two para-
meters that can vary: volatility and carry. While testing the premium ver-
sus carry rules, operators need to keep in mind that the game can change
and the optimal decision reversed should one of the two undergo changes.
This makes a more complex method of decision making necessary when
either interest rates are volatile or the yield curve presents a steep slope. A
nested binomial (or trinomial) tree is one such best technique, although
some operators resort to the heavier Monte Carlo methods.

The strength of a binomial tree is that it can include within itself some in-
formation on future volatility and interest rates, through both the forward
price and its process. The volatility does not need to be constant between
[End OCR]*

## Page 190

*[Image OCR]
176 Measuring Option Risks

Option Wizard: Discrete States of the Parameters on Trees*

There are two levels of methods involving “local” volatility: the basic and the
stochastic parameter binomial/trinomial.

The basic binomial tree prices options using information available from
the market and applies it to the growth between nodes. Such information
could be the forward volatility, the interest rates forwards, the skew, and so on.

The following chart shows that the parameters used correspond to both
forward volatility and forward interest rates rather than spot rates. Each state
between S(O) and S(3) depends on a forward volatility for the magnitude of the
u or d moves (the size of the steps) and the carry rd-d (interest rate differential
or dividend nétof financial carry). Forward volatility affects the absolute size of
both u and d while the carry determines the ratio u/d.

See also Module B.

S(3)

uuu

wey wae
5(0) u us a
oS

vol(1,2) | vol(2,3)
rd(1,2) rd(2,3)

For reasons beyond the scope of this book, a trinomial tree is more adapt-
able to the processes of forward-forward parameters. The trinomial presents the,
advantages of allowing for changes in volatility through the middle node (bino-
mial trees often do not recombine when operators try to include a “smile”).

The stochastic parameter tree is an adaptation of the preceding tree with
another branching for volatility and/or interest rates at every node.

The following figure shows a simplified state of the volatility curve at
node 1. We can define Vu as volatility going up and assign a term structure to

(Continued)
[End OCR]*

## Page 191

*[Image OCR]
Theta and Minor Greeks 177

Probability

Su(1)Vu 50%
u Su(1)Vd 50%
d Sd(1)Vu 50%
Sd(1)Vd 50%

the state, as well as a probability of each move,* with Vu and Vd states in the
term structure of volatility. The intuition of the evolution can be visualized
with the following chart of the spot curve (a forward curve is generally de-
rived from it at every node):

S(0)

19
Z 14 Vu
%
& 9 Vd
4

Time To Expiration

This methodology has been applied to interest rates yield curve evolution
but can easily be applied to the volatility curve.

* The reader unfamiliar with binomial trees would gain from reading an introductory text on bi-
nomial option pricing, such as Cox and Rubinstein (1985) or Hull (1993).

* Most traders who discussed the point with the author throughout the years showed a marked
preference for simulating and modeling a “skewed” process: higher probability of declining
volatility and lower probability of rising volatility, both states having equal expected payoffs
(higher up-moves).

nodes. The operator can thus defer the early exercise if the tree disclosed that
he would have to regret it should the following day bring a higher volatility.

Another way to view it is that the American option needs all kind
of possible information between time zero and the expiration. A closed
form formula would not carry such information: The Black-Scholes-
Merton method only requires operators to plug into their computer the
market parameters on the expiration date without more information in
between.
[End OCR]*

## Page 192

*[Image OCR]
178 Measuring Option Risks

A shortcut method used by the author is the “Rho fudge.” Its purpose is
to find the right duration (i.e., expected time to termination) for an Ameri-
can option. Knowing the right maturity allows for:

¢ The use of the right parameter for the pricing of the option (ie.,
volatility and term interest and carry rates).

e A better test for early exercise.

It involves deriving the sensitivity to interest rates of the path-depen-
dent option by assuming that the expected life of the option would match
that of the synthetic interest rate instrument thus derived.

Rho2 of an American option

= Nominal duration x
Omega onmnatearaon ( Rho2 of a European option

The reason Rho2 (it corresponds to the foreign rate or the dividend pay-
out) is used instead of Rhol is because it excludes the financing cost of the
premium. The trader also can compute the Rho2 from the delta by assuming
that the delta is a zero-coupon bond of the nominal maturity provided.

Example: Assume that the 1-year American option is equivalent to a
European one (zero probability of early exercise): The omega should
then be one year. .

Rho2 of a 1-year option = delta X 1 (i.e., a rise in 100 basis points in
the foreign rate or the dividend payout would correspond to a 100 basis
points rise in the option price).

Delta X 1

Omega =
Delta

-= 1 (years).

Now assume that the structure reacts only by .73

Delta X .73

oO =
mega Delta

= .73 years.

This method, while far from thorough, provides a rapid tool to com-
pute the expected life of a knock-out or any American binary option.

Alpha

@ The alpha (also called gamma rent) is the theta per gamma ratio for an
option position.

It reflects the gamma quality in terms of the “rent” and thus is the best in-
dicator of the quality of the earnings from gamma for one dollar risked
overnight. A high alpha means that the premium owner is not adequately
[End OCR]*

## Page 193

*[Image OCR]
Theta and Minor Greeks 179

compensated for the costs of the time decay. A low alpha means that the
trader is risking little theta for the gamma.

The premium seller who does not collect enough premium for the risks
taken will not survive for long. Typically, spreading options (see Chapter
16) helps achieve a better alpha (that is a lower one for the longs and a
higher one for the shorts), through the purchase of a cheap option and the
sale of an expensive one.

It is computed as follows:

Alpha = Decay/Gamma

A more advanced computation of the alpha would use the modified
theta, instead of the analytical theta, as a rate of decay.

The modified theta takes into account the slide on the curve. It can be
computed as:

Alpha = Modified decay/Gamma

Gamma is also more accurate when discretely computed. It is always
preferable to calculate the shadow gamma when deriving the alpha, to get the
full picture of both the risks and the returns.

The fair value alpha according to a schedule of volatility levels is pre-
sented in Table 10.1. Assuming interest rates are zero (or that the pre-
mium account would be eventually charged back to the trader in his cash
account)*

Theta = 4% Gamma X (Volatility)* (Asset price)”

Hence

_ Theta _ 1 36)
Alpha = Gamma 2° S
which is insensitive to time to expiration (at constant volatility, or flat
volatility term structure). Most beginning option traders (and some un-
skilled veterans) think that by selling short-term options they may be cap-
turing more time decay for the risks involved. The gamma of a longer
option is lower than that of the front but so is time decay. -

The alpha could be dependent on time to expiration if the volatility
level depended on the number of days to go. If the volatility of the 3-month
traded lower than the 6-month, then the alpha of the 6-month would be
different from the alpha of the 3-month. Otherwise, it would be the same.

Eliminating the interest carry from the equation, as explained earlier in
this chapter, allows the trader to get a pure theta number untainted with
interest rates.
[End OCR]*

## Page 194

*[Image OCR]
180 Measuring Option Risks

Table 10.1 Fair Value Alpha for a Given

Volatility and No Interest Rates Effect

Volatility Alpha
2.0 56
4.0 225
6.0 506
8.0 899
9.0 1138
10.0 1405
11.0 1700
12.0 2023
14.0 2754
16.0 3596
18.0 4551
20.0 5618
22.0 6796
24.0 8086
26.0 9488
28.0 11002
30.0 12627
32.0 14363
34.0 16210
36.0 18168
38.0 20237
40.0 22416
42.0 24706
44.0 27105

Risk Management Rule:

alpha for a short gamma position or higher than a fair value alpha

for a long gamma position will result in long-term losses (by the
law of large numbers).°

Table 10.1 shows that at 10% volatility, a gamma should cost $1,405 re-

gardless of the expiration.

Example: The alpha for a portfolio could yield unexpected results. Ina
calendar spread where the 1-month trades at 11% annualized volatility
and the 3-month trades at 12%, the trader buys $100 million 1-month

An alpha that is lower than the fair value
[End OCR]*

## Page 195

*[Image OCR]
Theta and Minor Greeks 181

“

at-the-money calls at 11 vols: total 13.5 positive gammas, at a cost of
$22,950 per day.

He sells $100 million 3-month at-the-money calls: total 7.5 negative
gammas for a total cost of $15,172 per day.

The net position will be long 6 gammas at a cost of $7,780 per day.

The alpha is therefore $1,297 per gamma. According to the scale in
Table 10.1, the trader’s costs are such that he will break even at close to
9.5 volatility, which represents an average of .6% moves per day.

TABLE OF GREEKS

The calculations shown in Tables 10.2 and 10.3 are usually memorized by
market makers. They can thus rapidly compare options without using cal-
culators.

For simplification, no interest rates are assumed. The reader should pre-
sent value the vega using e~"', t being the period in fraction of years.

Weights: Derived using the ratios of the square root of time of the pe-
riod. These are the weights for those who believe that volatility mean-
reverts at a Vt speed.

Vega: Classical unmodified at-the-money vega. The reader can multiply
the period length by his own factor to’ get the modified vega.

Table 10.2 Vega, Gamma, and Theta

Face = 1 r= 0% Vol
million

. 15
Days Maturity Weights Vega Gamma Theta
3 3d 0.36 2.93 —1297
7 lw 361% 0.55 1.92 — 668
14 2w 255 0.78 1.36 —443
21 3w 208 0.96 1.11 —355
30 Im 174 1.14 0.93 —293
61 2m 122 1.60 0.65 —203
91 3m 100 2.00 0.53 —165
182 6m 71 2.83 0.46 —143
373 9m 58 3.46 0.41 —128
365 12m 50 3.98 0.38 -116
456 15m 45 4.46 0.35 —108

547 18m 41 4.88 0.33 -101
730 2 year 35 5.63 0.31 —95
[End OCR]*

## Page 196

*[Image OCR]
182 Measuring Option Risks

Table 10.3 Vega Ratios for Out-of-the-
Money Options

Forward Delta Vega Ratio %
50 100
45 98
40 95
35 91
30 85
25 77
20 68
15 57
10 43

5 25

Gamma: Computed for a volatility at 15%. For a different volatility, use
the inverse ratio: gamma is inversely proportional to the volatility level.

Gamma at 10% = Gamma at 15% X (15/10) = 1.5 k Gamma at 15%

Theta: Classical unmodified theta, for a volatility at 15%. To compute
theta for a different volatility, one should use the ratio of the volatilities:

Theta at 10% = Theta at 15% X (10/15) = .667 X Theta at 15%

Vega Ratio: Table 10.3 shows the vega relationship to the at-the-money
option (by the forward). The vega of a 3-month at-the-money option is
known from the vega column in Table 10.2. The reader can thus derive
the vega of an out-of-the-money option provided he knows the delta.

For a deep in-the-money option, the vega ratio is 100 minus the delta of
the out-of-the-money corresponding instrument of the same strike. Put-call
parity allows for vegas of European options and all the second derivatives
to be the same for puts and calls of the same strike.

Because interest rates of zero are assumed, it becomes necessary to be
careful with the relationship when interest rates rise: The difference between
forward delta and cash deltas increases. Unfortunately, most option traders
measure their deltas initially in the cash.

Stealth and Health

These simplified measures are used by barrier option traders.° Both mea-
sures have their limitations as they do not take into account the volatility of
the underlying asset.’
[End OCR]*

## Page 197

*[Image OCR]
Theta and Minor Greeks 183

M Stealth corresponds to the percentage difference between the strike and
the trigger. Health is the percentage difference between the current
spot (not the forward, as will be explained) and the trigger.

Stealth is used as an indicator of how much the option resembles a
vanilla. The higher the stealth, the closer the option to a regular option in
both price and risk profile:

e¢ The market trades at 100. The 100 call with a knock-out at .00001 will
in all respects be similar to a 100 call plain vanilla. The 100 call with
a knock-out at 300 (the nasty reverse knock-out) also will be priced
like a regular call: Its difference will not appear until close to 300.

¢ The 100 call knock-in at .01 will be priced like no option. Likewise, a
100 call knock-in at 300 will not behave like much.

Health is an indicator of the execution risk. Because options traders
need to unwind their hedges at the barrier (and need to come as close to
that point as possible), it is a necessary risk management tool that makes
them aware of the proximity of the execution. Typically, a warning flag
comes up when that number drops below 1 standard deviation. Some people
mistake the health for the difference between the trigger and the forward:
The options do not knock out at the forward. This would only be applicable
to the rare forward knock-out/ins. ;

A much more potent measure than both is to use the expected stopping
time (also called expected time of arrival or first exit time, as described in
Chapter 18).

Convexity, Modified Convexity®

Convexity, also called curvature, is examined here in-a conventional frame-
work in an effort to incorporate the idea of a general yield curve behavior in
the measurement. Traders will learn here that:

* Convexity needs to be associated with the volatility of the parameter
concerned.

* Convexity exists everywhere, alas for the dynamic hedger.

B Convexity for an instrument describes the nonlinearity of the payoff of
the instrument with respect to a parallel move in one parameter. The in-
strument is then called convex with respect to that parameter.

Although convexity originally referred to bonds, it was discovered
that the concept needed to be applied to all financial instruments and to
their sensitivity with respect to more than one parameter. The convexity
[End OCR]*

## Page 198

*[Image OCR]
184 Measuring Option Risks

of derivative securities—since their pricing depends on many parame-
ters—will be defined as their second partial derivative relative to a partic-
ular parameter. The convexity of an option vis-a-vis spot is its gamma.
Out-of-the-money options will present a convexity relative to volatility
and interest rates.

Convexity of a Bond to Interest Rates (Simplified)

y is the annual yield expressed in continuously compounded rate equiv-
alent.

t is the unit period (1 year in this case).

iis the number of years.

C, is the cash flow of the period i.

B is the Bond price (the discounted coupons or payments).

B= Cie 7%

1

[Ms

t

a°B .
Convexity is ae SC, tPe ~¥i
Yo G1

Convexity of an Option to the Underlying Asset. The gamma is 0?V/0S*
for a derivative security. The measure will depend on the structure with V
the derivative price and S the underlying asset.

Vega Convexity. The vega convexity is 0?V/d0? with o the implied volatil-
ity in the market.

Price

3.00
4.00
5.00
6.00

88 8
yr Oo

10.00
11.00

Yield

Figure 10.4 Convexity for a fixed income instrument (zero coupon 30-year bond).
[End OCR]*

## Page 199

*[Image OCR]
Theta and Minor Greeks 185

Table 10.4 Case of an Exactly at-the-

Money Call
Value Vega
1 0.28 0.28
2 0.56 0.28
3 0.84 0.28
4 1.12 0.28
5 1.40 0.28
6 1.68 0.28
7 1.96 0.28
8 2.24 0.28
9 2.52 0.28
10 2.80 0.28
11 3.08 0.28
12 3.36 0.28
13 3.64 0.28
14 3.92 0.28
15 4.20 0.28
16 4.48 0.28
17 4.76 0.28
18 5.04 0.28
19 5.32 0.28
20 5.60 0.28
21 5.88 . 0.28

22 6.16 0.28

Convexity in Figure 10.4 is the effect of a nonlinear change in payoff re-
sulting from the changes in yield. It is magnified in this example by the use
of a 30-year maturity and no coupon.

Convexity for a Nonlinear Derivative. Table 10.4 is an example of a case of
an exactly at-the-money call. It is European and exactly at the money by the
forward. We can see in Figure 10.5 that its vega is linear. Tables 10.5, 10.6,
and Figure 10.6 show the vega convexity for an option that is away from the

Option Wizard: Middlebrow Convexity versus Modified Convexity

Middlebrow convexity is a measure of the sensitivity of a structure to a parame-
ter. The technique is erroneous for comparative purposes as the convexity of a
2-year bond (with respect to yield) could hardly be compared with that of
a 10-year bond since the two rates exhibit different volatilities.

Modified convexity is an attempt to rectify that by rescaling to the relative
volatilities.
[End OCR]*

## Page 200

*[Image OCR]
186 Measuring Option Risks

10.00
9.00
8.00 S
7.00
6.00
5.00
4.00
3.00
2.00
1.00

0.00
Mn Oo OO FR 2 ®
7 2 FF a & Bf B

Price/Vega

Volatility

Figure 10.5 Volatility sensitivity for an at-the-money option.

Table 10.5 Case of a 180 Day 15% Out-of-
the-Money Call (or 15% In-the-Money Put)

Value Vega

5 0.00 0.00
6 0.00 0.00
7 0.00 0.01
8 0.01- 0.02
9 0.03 0.03
10 0.07 0.05
11 0.12 0.07
12 0.18 0.08
13 0.27 0.10
14 0.37 0.12
15 0.49 . 0.13
16 0.62 0.14
17 0.76 0.16
18 0.92 0.17
19 1.09 0.18
20 1.26 0.19
21 1.45 0.20
22 1.65 0.20
23 1.85 0.21
24 2.06 0.22
25 2.27 0.22
26 2.49 0.23
27 2.72 0.23
28 2.95 0.23
29 3.18 0.24

30 3.42 0.24
[End OCR]*

## Page 201

*[Image OCR]
Theta and Minor Greeks

Table 10.6 A Simplified Modified Convexity

Modified

Maturity Factor Convexity Convexity
2 1.23 0.04 0.05
3 1.2 0.09 0.11
4 1.16 0.15 0.17
5 1.13 0.22 0.25
6 1.1 0.30 0.33
7 1.08 0.38 0.42
8 1.05 0.48 0.50
9 1.03 0.58 0.60
10 1 0.69 0.69
15 0.91 1.27 1.15
20 0.85 1.87 1.59
30 0.84 2.98 2.50

187

money. Chapter 15 provides a discussion for the impact of vega convexity on
option pricing.

Note. When volatility rises excessively, all options become concave with
respect to volatility because every option is capped at the price of the un-
derlying asset. If the underlying asset is worth 100, an option price cannot
exceed 100 or rational] traders (there are a few) would resort to buying the
asset instead of the call. a

Viewed differently, at a very high volatility, an asset becomes a call on
itself since it can rise infinitely but only go down in a limited way. There-
fore, there will be a convergence between the price of the asset and that of

the call option.

4.50
4.00
3.50
3.00
2.50
2.00
1.50
1.00
0.50
0.00

Price/Vega

Volatility

Figure 10.6 Volatility convexity for an out-of-the-money (or in-the-money) option.
[End OCR]*

## Page 202

*[Image OCR]
188 Measuring Option Risks

Modified Interest Rate Convexity. While most instruments that offer con-
vexity will command a premium for the extra kicker, traders should care-
fully measure the behavior of the underlying parameters. Many bond traders
find that instruments with the highest convexity, the long-term low-coupon
bonds, become positioned on the yield curve at the least volatile point, thus
making the benefits of their convexity smaller than expected. It becomes er-
roneous to compare the convexity of a 5-year with that of a 30-year bond
without the use of factors similar to the volatility weightings.

The same applies to options. Often the highest convexity is found in back
month out-of-the-money options with the most stable implied volatility.

Information about convexity multiplied in the middlebrow literature of
the 1980s and early 1990s. Assuming parallel moves the number is generally
useless. Using a trader’s intuition is more helpful than computing indica-
tors that are misleading to the public. What is remarkable is that these
methods ignored the academic developments pioneered by Vasicek (1977),
Cox-Ingersoll-Ross (1985a) and others who developed methods of gauging
the changes in the shapes of the yield curve.

Much of the convexity-based hedging implemented by the middlebrow
fund managers was therefore based on faulty assumptions.

The easiest way to modify convexity is to take into account relative
volatility.

Example: A one-factor model used to correct convexity.

Data: May 1993—May 1995

Maturity 2y Sy 10y 30y
Sensitivity 1.231335 1.144669 1 0.839248

Sensitivity is the relative sensitivity of the yield curve (expressed in
yield) to that of a 10-year Treasury bond. A one-factor model, in essence, se-
lects one instrument as a “pillar” and assumes a 100 percent correlation
with others.

Using a polynomial interpolation, it is possible to construct the quick-
and-dirty approximated function? (Table 10.6).

The convexity used in this example is defined as a dollar difference be-
tween the absolute change from a 100 basis points move up and a 100 basis
points move down, for an instrument trading at par with interest rates at
6%, which is close to the general yield curve in June 1995. This quick-and-
dirty method could be refined further by measuring exact duration.

Example. Eurodollars Are Convex: This example illustrates the differ-
ence between forwards and futures where there is a correlation between
the price of the instrument and its rate of financing. Its simplicity
[End OCR]*

## Page 203

*[Image OCR]
Theta and Minor Greeks 189

allows for easier understanding of the notion. It illustrates the risk
management rule on convexity and concavity of futures that was dis-
cussed in Chapter 1.

The “boost” (not to be confused with the name of a double-barrier
linked product) is an implicit convexity to being short Eurodollar futures
because one can reinvest profits at a higher rate and can finance losses at
a lower rate.

Eurodollar futures are of the marked-to-market category; operators
receive the variation margin at the closing of a profitable day and need
to pay it after a losing one. There is a connection between financing
costs and profitability: When the market goes down, the short Eurodol-
lar futures position earns money. The money is wired to the investor
and can thus be invested at high rate (since the Eurodollars went down).
Conversely, the trader incurs losses when the market goes up, but these
losses can be financed at a lower rate.

Assume that the position shown in Table 10.7 is short 400 Eurodollar
contracts on the Chicago Mercantile Exchange for a contract that expires
in one year. In addition, assume for pedagogical reasons that the yield

Table 10.7 Illustration of the Boost

Future Rate Profit Reinvestment Total
88 0.12 6,000,000 720,000 6,720,000
88.5 0.115 5,500,000 632,500 6,132,500
89 0.11 5,000,000 550,000 5,550,000
89.5 0.105 4,500,000 472,500 4,972,500
90 0.1 4,000,000 400,000 4,400,000
90.5 0.095 3,500,000 332,500 3,832,500
91 0.09 3,000,000 270,000 3,270,000
91.5 0.085 2,500,000 212,500 2,712,500
92 0.08 2,000,000 160,000 2,160,000
92.5 0.075 1,500,000 112,500 1,612,500
93 0.07 1,000,000 70,000 1,070,000
93.5 0.065 500,000 32,500 532,500
94 0.06 — — —
94.5 0.055 (500,000) (27,500) (527,500)
95 0.05 (1,000,000) (50,000) (1,050,000)
95.5 0.045 (1,500,000) (67,500) (1,567,500)
96 0.04 (2,000,000) (80,000) (2,080,000)
96.5 0.035 (2,500,000) (87,500) (2,587,500)
97 0.03 (3,000,000) (90,000) (3,090,000)
97.5 0.025 (3,500,000) (87,500) (3,587,500)
98 0.02 (4,000,000) (80,000) (4,080,000)
98.5 0.015 (4,500,000) (67,500) (4,567,500)
[End OCR]*

## Page 204

*[Image OCR]
190 Measuring Option Risks

curve is flat and moves in parallel (a 100% correlation between the fi-
nancing of the position and its profitability).

When the Eurodollars go to 9300, the trader is able to reinvest his
positive variation margin at 7%, thus earning $70,000 over the year until
expiration. When Eurodollars go to 9500, however, the trader incurring
a loss will only finance it for $50,000. This difference of $20,000 repre-
sents a minute convexity differential, but needs to be taken into account
by traders.

A boost becomes significant when dealing with back months, since
profits in 5-year Eurodollars will be reinvested over the next 5 years. So
the boost increases when maturities lengthen.

The boost accelerates when rates are high, as the reinvestment rate
is higher.'°

The “Double Bubble”

An interesting aspect of the boost is that being short a Eurodollar strip pre-
sents a positive convexity. This would not be remarkable except that being
short Eurodollar strips presents a remarkable hedge to being long zero
coupons. Since zero coupons are convex as well, the trade would offer dou-
ble convexity, or what one famous arbitrageur dubs “the double bubble.”
[End OCR]*

## Page 205

*[Image OCR]
Chapter 1 1
The Greeks and Their Behavior

’ Time is elastic.
Marcel Proust

Using a dynamic framework, this chapter presents an analysis of the
changes of the Greeks in time and space (the relation between the asset and
the strike). Only these two dimensions are necessary because time and
volatility exert the same effect on the first and second derivatives, though
not on the value of the portfolio. Volatility and time are one and the same
element in a Brownian motion, since variance (the square of the volatility) is
exactly a proportion of time. Where the time exerts its own independent ef-
fect is on the e" and e* (carry rate and discount rate), which generally do
not carry any undue significance.
The analysis will consider the following effects:

¢ The various changes of Greeks with time that traders typically call
the bleed (the metaphor is that of a slow loss of vital blood—option
- traders dislike decay).

¢ The various changes of parameters with space that traders call the
speed or the “D”: DdeltaDspot, DgammaDspot, DvegaDspot. For a
mathematician they are various third derivatives.

THE BLEED: GAMMA AND DELTA BLEED
(HOLDING VOLATILITY CONSTANT)

@ The bleed is the change in the delta and the gamma of an option posi-
tion with the passage of time.

The delta bleed and gamma bleed are computed by repricing the portfolio
one day shorter and measuring the difference.

Delta bleed = Delta today — Delta next period

Gamma bleed = Gamma today — Gamma next period

191
[End OCR]*

## Page 206

*[Image OCR]
192 Measuring Option Risks

Option Wizard: The Difficult Boss Cette Story)

One option trader recounts that he was given a limit by the management in the
form of a certain amount of maximum allowable delta, vega, and theta expo-
sures for his position. So he knew he needed to leave a wide margin between
his exposures and the position limit.

One day, he was long out-of-the-money options and the market exploded
higher, taking him late in the day into an area where his long vega increased
owing to the position of his out-of-the-money options. He was then called into
his boss’s office to be informed that the fine-tuning clerk whose job consisted
in measuring the trader’s compliance with the sacrosanct limits had detected a
trespassing of some gravity. So the following dialogue took place:

Boss: You are outside the limit.

Trader: | know, but the market moved rapidly late in the day. Anyway,
my delta was computed off a lower volatility so I’m protected.

Boss: You mean you agree that you are outside the limit.

Trader: | repeat, the dynamics of the position took me there. When the
option markets closed, | was inside the limit.

Boss: So you tried to cheat by being outside the limit after the close of
the market.

Trader: If | knew what the market was going to do, | would have taken a
straight position. | had a strong convexity to my vega that made it increase
at a higher volatility. You have to look at both third derivatives with re-
spect to spot and with respect to volatility.

(silence . . .)
Boss: Ummph. This time, | will sign off on the sheet. Next time, don’t
cheat again.

Every option trader has such war stories. The most ludicrous ones involve a
limit in the shape of a delta restriction.

Time pushes all the out-of-the-money options further out-of-the-money,
therefore reducing their deltas. Time also pushes all the in-the-money op-
tions further in-the-money, therefore increasing their deltas. Although this
seems trivial, traders should not overlook the following confusion: Gamma
is said to increase with time, but for an out-of-the-money option being
pushed further, out-of-the-money would cause some gamma loss. The re-
sults are therefore mixed.

Book runners are always miffed at the alterations brought about by
time. Because portfolios of book runners always include a mixture of
[End OCR]*

## Page 207

*[Image OCR]
The Greeks and Their Behavior 193

strikes, with strong concentration away from the money, a meaningful daily
bleed is guaranteed.

Risk Management Rule: An option position that is long up-gamma

(see definition, Chapter 8) and short down-gamma will bleed into
shorter delta over time, and vice versa.

The rule is simple to understand: (1) Such a position is long out-of-the-
money calls short out-of-the-money puts. The calls will decrease in deltas
(shorter delta) since the delta of out-of-the-money options decreases with
time and the puts will decrease in deltas for the same reason (longer deltas);
(2) the same position can be constructed (put/call parity) with deep-in-the-
money calls and puts. The same result can be proven by the put-call parity
rules (see the discussion on option equivalence in Chapter 1).

Example: For simplification, assume that the book has two trades:
Short $1,000,000 in face value of a 100 call expiring in 60 days (OPT1 in
Table 11.1).

Long $5,040,000 in face value of a 106 call expiring in 30 days (OPT2 in
Table 11.1).

Option Wizard: Thank God It’s Friday

The bleed is always accentuated for traders ahead of a weekend when, after
Friday lunch, they start examining their Monday exposures, not always without
a shock. The endless dialogues between traders as to when to set the clock
usually take place then: whether to use Monday “sheets,” Friday sheets, or
somewhere in between (sheets are printed reports showing the price of the op-
tion at different levels of asset price and volatility).

Traders commonly use risk management systems with discrete time. Options
are not usually priced on a continuous clock. Time to expiration is not updated
every second, but in 24 hour increments. Most operate on a 7-day-week basis,
365 days a year. While this jump function is of small consequence during the
week, it becomes an endless source of confusion when, on Friday, the traders
issue portfolio simulations with a Monday date.

When pit traders perform “covered” trades on a Friday afternoon, where
they trade the option covered with the delta, disagreements occur as to the
amount of deltas to swap. It is then easy to see that one part of the pit uses
Friday deltas, another one Saturday, another one Sunday, and so on.
[End OCR]*

## Page 208

*[Image OCR]
194 Measuring Option Risks

Implied volatility as usual is at 15.7% annualized (for both options) and
the market is initially at 100.

The trade starts delta neutral as both legs match each other.

Table 11.1 shows that the deltas (delta 1 column) of the at-the-money 100
calls (OPT1) remains constant (almost, except for a total loss rounded to 2
deltas) while the 106 calls (OPT2) lose their deltas (delta 2 column) rapidly
(from to 512 to 175). They “bleed” fast.

The table shows the acceleration of gammas for the 100 calls offset
(“positive gamma bleed”) by a deceleration for the 106 calls, with an impor-
tant effect on the overall gamma of the portfolio.

The graphs in Figure 11.1 (A) and (B) show that the delta bleed is topical
(it depends on the location of the spot). It reverses above the strike that
caused the bleed. This analysis leads to the fourth derivative: how the bleed
behaves with respect to spot. Most nonoption traders, particularly those with
an engineering education, believe that fourth derivatives are of small impor-
tance in practical life. They are trained to cut everything beyond the second
derivative, which, for options, would be a gamma. This attitude is incompati-
ble with the trading of a portfolio of options where the risks change abruptly
between levels. Many option traders report disagreements with Taylor-
expansion-minded semimathematicians who deride their attempt at looking
at the heart of the position with fourth (and higher) derivatives.

Table 11.1 Effect of Time over the Deltas and Gammas

Days Total ‘Total
Hence OPT1 Deltal Gammal OPT2 Delta2 Gamma2 Delta Gamma

0 2.54 —513 —63 0.21 512 » 199 0 136
1 2.52 —513 -63 0.20 493 197 —20 133
2 2.50 ~512 —64 0.19 472 194 —40 130
3 2.47 —512 —64 0.17 452 191 ~61 127
4 2.45 —512 —65 0.16 430 188 —82 123
5 2.43 —512 —65 0.15 409 184 -103 119
6 2.41 —512 —66 0.14 387 180 —125 114
7 2.39 —512 —67 0.12 364 176 ~147 109
8 2.36 —512 ~-67 0.11 342 171 —170 104
9 2.34 —512 —68 0.10 319 166 ~193 98
10 2.32 —512 —69 0.09 295 160 —217 92
11 2.29 —511 —69 0.08 271 154 —240 85
12 2.27 —511 —70 0.07 247 147 —264 77
13 2.25 —511 -71 0.06 223 139 ~-288 68
14 2.22 —511 ~72 0.05 199 131 —312 59
15 2.20 -511 —72 0.04 175 122 —336 49
[End OCR]*

## Page 209

*[Image OCR]
The Greeks and Their Behavior 195

Today + 15 days

Deita

Asset Price

A. Delta change over time for a ratio diagonal

1200
1000

800

600

P/L (000)

day + 15d
400 y ays

-200 3

Asset Price

B. P/L change over time for a ratio diagonal

Figure 11.1 Effects of time over a ratio diagonal spread.

Fourth derivatives will not be explored in any depth here, except to
warn traders of the effect of change in bleed at different levels.

Bleed with Changes in Volatility

Changes in implied volatility do not overly complicate the bleed. Volatility
shifts can operate like a time acceleration, with an effect that is linked to
the length of the option (more exactly the square of the volatility shift is
proportional to the length of the option). Every option beginner knows that
the shorter the option, the greater the effect of time on the structure. A
longer dated option, say a one-year call with 20 deltas, would lose, at 16%
[End OCR]*

## Page 210

*[Image OCR]
196 Measuring Option Risks

volatility, .04 deltas (the delta would go from 20% to 19.96%). A change in
volatility from 16% to 15% would move the delta from 20% to 18.2%.

Warning: The barrier options trader needs to be aware of the two-way
bleed (see Option Wizard, Forward and Backward Bleed). A rise in volatil-
ity can initially lengthen the delta, while a further rise could decrease
it. Barrier options are thus considered extremely topical: What works
on one point of the possible map does not necessarily work on others.
The third, fourth, fifth, and sixth moments (see definition later in this
chapter) are always there to confuse operators.

Going into the Expiration of a Vanilla Option

Bleed moving into an expiration is so rapid that handling it properly re-
quires a great deal of experience. That is where the difference between the
experience of listed traders with discrete expirations and that of over-the-
counter operators with almost continuous ones becomes most apparent.
Bleed then becomes so fast that volatility changes cannot carry too large an
effect on it.

Many operators start trading an expiration as a binary one, with al] the
bleed taken on the report generated the previous day. At 5 o’clock the previ-
ous day, they generate a run with the deltas of the option in binary form: zero
if the option is out of the money and 100% if it is in the money. However, it
does not mean that they will adjust their position in a binary way. Most still
trade it in a smooth manner, with progressive adjustments. Others use more
sophisticated devices and trade “on the clock,” with their risk management
system adjusting the delta by repricing the portfolio continuously.

Example. An Expiration Day Report: Assume that the market trades at
100 and that the operator is short the 101 calls expiring in the over-the-
counter market at 10:00 a.m. the same day.

Figures 11.2 and 11.3 show the option delta and price the day before expi-
ration. Figures 11.4 and 11.5 show the same elements on the following day.

So the operator needs to flip between Report I (Figures 11.2 and 11.3) and
Report I (Figures 11.4 and 11.5) by making sure that the convergence takes
place in a smooth manner. Ideally, the following would occur: The operator
would subscribe to the smoothness rule by ascertaining that the transition
between two points does not take place abruptly. So Report I would hold at all
times but with a continuous rescaling in accordance to the passage of time.

One day before expiration, the trader would look at Report I. Then he
would need to try to maintain the same smooth profile between two points
by narrowing down the scale, as by a microscope in which the magnifying
power is gradually increased.
[End OCR]*

## Page 211

*[Image OCR]
The Greeks and Their Behavior 197

100 4.00
B80 3.00
3 io 8 2.00
a 40 =o
20 "1.00
Q
& Ss ~ & ~ F 6 Sg
= = & 8 8
Asset Price Asset Price
Figure 11.2. Report I: Single option Figure 11.3 Report I: Single option
delta 1 day before expiration. price 1 day before expiration.
100 4.00
80 3.00
& 60 8
20 1.00
0 " 0.00
> mM MW aw w mm ®2 © NHN wW
5" 8 F 8 5° g =F Bg
Asset Price Asset Price

Figure 11.4 Report II: Single option Figure 11.5 Report II: Single option
delta at expiration. price at expiration.

If the trader has no transaction costs, expirations‘-would be completely
smooth, as by the Black-Scholes-Merton formula. However, traders need to
restrict their trading to a number of discrete interventions. [t does not mean
that they need to adjust their delta on expiration in a binary form. They
must remember the golden rule that hedging is smoothing. Depending on
the traders’ perception of their own risks, they could narrow the scale in a
continuous way, as follows:

1. Day Before 2. Half a Day to Go 3. 1 Hour to Go
Delta Asset Delta Asset Delta Asset
1 98 1 98.84 1 100.37
25 100.44 25 100.61 25 100.89
50 101.00 50 101.00 50 101.00
75 101.65 75 101.40 75 101.11

99 102.95 99 102.38 99 101.40
[End OCR]*

## Page 212

*[Image OCR]
198 Measuring Option Risks

Option Wizard (Advanced): A Modicum of Ito Calculus

When the author was being interviewed for his first job, a senior foreign ex-
change trader showed him the following scheme on a tomato-sauce-stained
restaurant napkin:

One should sell a call option on USD-DEM for the hefty premium it then
commanded. Then he would leave a stop-loss to buy the full amount at the
strike price. If the spot recrossed, then he would do the reverse. Given this
brilliant strategy, the premium in the market appeared unjustified.

The author tried to register some doubt about this plan but could not
quite explain stochastic integration to a head option trader (besides this was
1983 and option traders concealed any degree other than an MBA to avoid
the “nerd-discrimination” that prevailed at the time).

Most option traders report that they have had to explain to a boss why
they do not follow such a strategy. The reason is as follows:

Local time* never becomes zero, owing to the non-differentiability of the
Brownian motion. We know from the spreadsheet exercise that the Brown-
ian motion never becomes smooth. Making the time intervals smaller does
not make it more differentiable; otherwise, there would be a rebalancing
strategy that would outperform others. If traders rebalanced at the interval
where the function became smooth, they would keep the entire premium
and there would be no need for this book.

The frequency of rebalancing does not affect the fair value of the option.
Assume that traders operate in the Black-Scholes-Merton world of no transac-
tion costs. The rebalancing P/L does not depend on the rebalancing frequency.
It does not matter whether they rebalance over the time interval At at the in-
crement At or at increments At/2 or At/4 or At/n, with corresponding spot
moves AS, AS/2, and so on. The expected P/L would be more precise the
smaller the adjustment period, but it certainly would not vary.

The key to understanding Ito’s calculus is knowing that there is a lag be-
tween the decision to rebalance and the final execution price and such lag is
never eradicated.

In fact, the costs of hedging are V2/m V (AS)? times the amount of time
spent swinging between S and S + AS, the resulting volatility, since the trader
hedges himself after the move AS and thus loses the absolute value of AS times
the face value.

This figure illustrates the rebalancing strategy issue. In practical terms, the
amount of times, he would have to pay the transaction costs of one tick between
1.5000 and 1.5001 in USD-DEM is equal to the amount of time he would pay
the transaction costs of 100 ticks between 1.5000 and 1.5100. For the market
will snap more often between 1.5000 and 1.5001 in a compensating way.

(Continued)
[End OCR]*

## Page 213

*[Image OCR]
The Greeks and Their Behavior 199

150.00 150.10

Sell

Loss 10 pips

Narrower Increments

150.00 150.01

@
Sell

Loss 1 pip

Intuitively, the trader can truly derive an option price as the expected
rebalancing at the strike. An out-of-the-money option is less likely to whip
around the strike because the strike is away from the center of the distribution.

First Conclusion: The costs of replicating an option (assuming no transac-
tion costs) by buying the entire face value at the strike or by continuously
rebalancing the delta are expected to be the same. They will correspond to
the Black-Scholes-Merton value.

Second Conclusion (food for thought): In some cases, the wrinkles of the
market favor the strategy of hedging with discrete increments and execut-
ing the full face value. Markets supported by central banks will cross
the intervention level and not return. Likewise, in serially correlated mar-
kets, some chartists and voodoo traders enjoy being short delta below
some level and long above it.

*See Karatzas and Shreve (1991).
[End OCR]*

## Page 214

*[Image OCR]
200 Measuring Option Risks

DDELTADVOL (STABILITY RATIO)

@ Ddeltadvol corresponds to the changes in deltas resulting from changes
in volatility levels.

The reader can decide whether weighted or unweighted volatility needs to
be used. The concept is similar to that of the bleed except that it can move in
two directions. Volatility moving up exerts the effect of the reversing of
time on an option portfolio. Likewise volatility moving down causes the
time to shorten and carries exactly the same effect as the bleed. The effect of
the moves in volatility over the portfolio is called the forward and back-
ward bleed. At a more advanced level; a portfolio, made up of spreads can-
not be projected through time without the testing of the effect of both the
bleed and the Ddeltadvol. The trader can thus only run a scenario of a series
of paths with changes in implied volatility and reprice the portfolio as if he
remained risk neutral in hewing to a delta-neutral strategy. Because the
delta neutrality will depend on the volatility at the rebalancing point, the
exercise is called a path-dependent portfolio testing.

Test 1 of Stability

The trader should raise the volatility and examine the first derivatives, par-
ticularly the delta. An increase in the deltas means that the position be-
comes increasingly longer vegas in a rally and shorter in a sell-off. It is
called a positive Ddeltadvol. It means that the book is net short options
below the money and net long option above the money. A negative Ddelta-
dvol means the reverse.

Example: A simple vertical spread is the most likely candidate for in-
stability, curiously. If the operator is long a 100 call and short a 110 call,
the deltas will shorten with Test 1. Higher volatility would raise the
deltas of the out-of-the-money calls and keep the at-the-money calls
constant.

Option Wizard: Forward and Backward Bleed

The forward bleed is the effect of time with the expiration becoming
shorter. It consists in moving the portfolio one day hence.

The backward bleed is the effect of time in reverse, as if the book went
back in time. It is similar to the notion of reverse time decay. The increase
in volatility exerts such an effect on the non-interest component of an op-
tion price.
[End OCR]*

## Page 215

*[Image OCR]
The Greeks and Their Behavior 201

This test needs to be run routinely on books in emerging market
currencies, the biased assets and other products where it not advisable to
be short volatility when the market decreases. It presents the benefit of
being simpler to perform than those tests that require matrices. The
reader should note that this test tells little about the gamma; it can only
explain the sensitivity of the vega to the skew.

Test 2 of Stability: The Asymptotic Vega Test

Repeating the preceding test at different levels of the underlying asset
could show the different exposures at various levels of the topography.
Sometimes the vega “flips” when the market rallies above or drops below a
cutting point. Such reversal is usually caused by the clustering of options in
one location dominating a particular area. That cluster, in turn would be
dominated by another one further away from the money. Figure 11.6 shows
such behavior.

Example: One is long the 110 calls in $100 million, short the 130 calls
in 300 million. When the volatility is sufficiently low and the market re-
mains between 100 and 115, the book would exhibit long vegas. This
would appear in a positive Ddeltadvol Test. At higher volatility levels,
however, the test would reverse and show a negative Ddeltadvol.

In the preceding example, the vegas appear jagged. An-abrupt rise in
volatility, however, would turn the vegas to monotonously short across the
entire spectrum. This is called the asymptotic vega test, which is only per-

formed for stress testing as it gives way out-of-the-money options undue
importance.

130

foo]
Low Positive Flat Negative
Volatility Vega Vega Vega
Negative DdeltaDvol and Vega

Figure 11.6 Option clusters.

High
Volatility
[End OCR]*

## Page 216

*[Image OCR]
202 Measuring Option Risks

MOMENTS OF AN OPTION POSITION

Moments for a distribution are ways mathematicians characterize some of
its higher order behavior. What the author calls moments for a position cor-
responds to the sensitivity of an option book to the higher moments of the
distribution. The modern financial approach of relying on a normal distri-
bution allows theoreticians to ignore its higher moments. The reason is that
the normal distribution is entirely characterized by its first two moments—
the mean and volatility. A compact distribution is defined (see Ingersoll,
1986) as a distribution the higher moments of which become increasingly
small in relation to the second moment. “Compact support” can be seen in
the tails of the bell curve that come rapidly close to 0. Therefore, in theory,
an option trader will only be exposed to the market direction (via the delta)
and the market volatility (through the gamma). Higher derivatives should
not truly matter.

Too bad. Given that the true world distributions do not entirely look nor-
mal (except by accident) we need to worry about fat tails, positive and nega-
tive skew, jumps, and other annoying matters. We will therefore stretch the
mathematical language by using the designation moment for a position in-
stead of a distribution.

The moments of an option position’ represent the sensitivity to some
order of change in the underlying security. The more information about the
higher moments, the more able the trader would be to track the position
and predict the changes over a wider range of prices.

The way operators look at the moments in their framework of analysis
concerns the orders of move in the asset price with time kept constant. So
the method of moments reveals little about the vega and its changes and the
sensitivity to time. . .

e First Moment. The delta, which indicates an exposure to mean of
the distribution.

¢ Second Moment. The delta of the delta, the gamma.

¢ Third Moment. The delta of the gamma, which is called the skew. It
is the gamma exposure in function of the asset price and becomes
naturally the level of asymmetry between them. If the gamma be-
comes more positive in the rally and more negative in the sell-off, the
third moment will be positive. It will be otherwise negative. (It be-
comes apparent that odd moments are indicators of symmetry while
even moments are indicators of convexity.)

¢ Fourth Moment. The tails. It corresponds to the gamma of the
gamma as it is the second moment of the second moment. When the
fourth moment is positive, the position is convex and therefore can
allow some sleep at night. In practical terms, it means that the
[End OCR]*

## Page 217

*[Image OCR]
The Greeks and Their Behavior 203

gamma increases when the market moves away from the center and
decreases when volatility moves down.

Probabilists usually ignore higher moments since those vanish very
rapidly. However, option traders, surprisingly, do not. It is their business to
deal with the fine tuning of the option position at higher moments. So one
probabilist choked at a mathematical finance conference while listening to
two option traders argue during lunch about the seventh moment of a dis-
tribution (indeed a true story).

An option book being spread out between items at all the possible
time and space points of the distribution will become very sensitive to the
moments of moments, and endlessly so, even without adding the beastly
compound option structures. This is because traders tend to maintain a
delta-neutral, gamma-neutral book with thousand of strikes netting out.
Canceling the lower moments (locally) is rather easy (one phone call to
take care of the first two moments). Canceling the next 5,000 moments,
however, would necessitate closing down the book. Going down the risk
management positions shows why option traders make good probability
experts:

¢ The Fifth Moment. It becomes the asymmetry sensitivity of the
fourth moment. A portfolio solely containing a short American bar-
rier option hedged with a short vanilla option will therefore be ex-
tremely sensitive to the fifth moment. It will increase in concavity as
the market comes closer to the barrier and will act more linear as the
market moves away from it.

¢ The Sixth Moment. Except for compound options, even moments do
not seem to be very beastly.

© The Seventh Moment. It is the sign of the convexity changes as the
underlying asset moves up or down. Typically, convexity positions
constructed with higher order out-of-the-money options one side of
the market (only out-of-the-money calls on calls) against at-the-
money options of an inferior order (e.g., a vanilla at-the-money) will
have a seventh moment that is very pronounced. In the call on calls
case versus at-the-money, the convexity will slide down and hit zero.
Where someone buys out-of-the-money calls on calls and sells out-of-
the-money puts on calls (delta neutral, of course), the seventh mo-
ment will be very scary.

With compound options, traders can thus reach seriously higher orders of
moments. An installment option (a fifth-order compound option), we will
see, requires an analysis that takes into account at least nine moments for
hedging stability.
[End OCR]*

## Page 218

*[Image OCR]
204 Measuring Option Risks

Risk Management Rule: Positions that seem neutral in the lower

moments and have an increasing exposure in the higher moments
will present trading difficulties.

For example, a delta-neutral risk reversal will appear to be harmless for
anyone stopping at the delta and gamma. However, the skew can have a
drastic effect on the P/L. The skew can be more volatile than the second mo-
ment, and this is also true for fourth moment positions.

IGNORING HIGHER GREEKS: THE Lock DELTA

Stress-testing a portfolio is a must in risk management. Although it is a ca-
sual, uncomplicated method, it proves effective for micromanagement of
the risks when one looks into only a single asset in an independent portfo-
lio of options.

@ The lock delta, or asymptotic delta method is a measurement of the
risks in a derivatives portfolio at extreme boundaries, generally, zero or
the infinite. The principal use of the asymptotic delta is to display the
structure of the position that is hidden by scenario analysis.

Option Wizard: Parametric and Nonparametric Tests

The market parameters being considerably unstable, it is always necessary to
perform tests that are independent of the distribution. The Lebanese lira, for
example, is pegged to the dollar most of the time and often jumps after some
political mishap. Even the notions of fat tails and skew are hardly applicable
owing to the extreme non-normality of the market. Using Greeks such as
gamma represent dangers. The tests performed on a portfolio containing such
assets need to be entirely street-smart, with simple scenarios (e.g., If the cur-
rency moves down to such a level, what is my P/L?).

A nonparametric test is defined as one where no assumptions are made
about the distribution (and its parameters). There is an entire branch of statis-
tics called “robust” that deals with such issues. Stress tests fall into the non-
parametric category.

It is necessary in poorly behaved instruments like emerging market debt
where the distribution and its parameters (especially volatility) are so unstable
as to make the conventional risk management tools inadequate.
[End OCR]*

## Page 219

*[Image OCR]
The Greeks and Their Behavior 205

Other partial derivatives than the delta are by definition zero at these
boundaries: The portfolio then ceases to behave like a derivatives portfolio.
As such, its sensitivity to volatility will be nil: Its gamma and vega will be
zero. In that respect, calling the measure a delta can be an exaggeration.
The measure is the netting of the underlying if all options were exercised. It
can literally be done on a napkin for positions that do not include exotic op-
tions and correlation products. One can count the net number of calls, add
them to the cash and futures, present value them, and translate the number
into asymptotic maximum up-deltas. Conversely, one can do the same with
puts and come up with the maximum down-deltas.

Risk Management Rule: The firms that survive large shocks in the
market (and the breakdown of convential associations between in-

struments) are the ones lucky enough not to have “scientific” risk
managers among their staff.

Experienced risk managers do not take the heavy models too seriously.
The key for survival resides in distinguishing between risk managers who
believe in models, and those confident enough not to.

Example 1. Lock Delta for a Covered Write: In all examples, the market
trades at 100. ,

A portfolio is long 100 futures, short 100 3-month calls struck at 110.
The upside asymptotic delta is zero delta. The downside asymptotic
delta is 100 long deltas as shown in Figure 11.7.

90
80
70
60
50
40
30
20
10

Delta

on oOo tt OD
- rr er TF

Cash Price

Figure 11.7 Lock delta for a covered write.
[End OCR]*

## Page 220

*[Image OCR]
206 Measuring Option Risks

Example 2. Lock Delta for a Strangle Write: A portfolio is short 100
out-of-the-money calls struck at 104 and short 100 out-of-the-money
puts struck at 96 (a 96-104 strangle). The upside asymptotic delta is 100
short deltas. The downside asymptotic delta is 100 long deltas as shown
in Figure 11.8.

Example 3. Lock Delta for a Ratio Write: A portfolio is long 100 calls
struck at 100, short 200 calls struck at 104 (a ratio write). The upside as-
ymptotic delta is short 100 deltas, the downside asymptotic delta is zero
as shown in Figure 11.9.

Before option market makers discovered the simulation matrices, the sys-
tem was principally asymptotic delta-related: The clearing firm measured
the net difference between upside and downside risk and charged the cus-
tomer a margin in accordance with the net residual exposure, which is that of
the asymptotic delta. A covered writer was absolved of margin, but the sale of
an additional call (a one-by-two covered write) caused the margin to become
that of a naked call. The clearing firms then were not interested in the math-
ematics of option pricing and did not concern themselves, to their protection,
with Greek derivatives.

The mentality of nonmathematical equity option traders was then en-
tirely pragmatic: Stocks used to experience sudden harrowing jumps on
takeover rumors and experienced unstable volatility, with brief discontinu-
ous movements following company news and announcements. Survival in-
stinct prevented the early participants from putting much trust in formulas

100

Delta
Oo

-100

Cash Price

Figure 11.8 Lock delta for a strangle write.
[End OCR]*

## Page 221

*[Image OCR]
The Greeks and Their Behavior 207

Delta
&

-100

Cash Price

Figure 11.9 Lock delta for a ratio write.

and volatility-based measurements when the distribution displayed acute
fat tails.

Over-the-counter (OTC) dealers and institutions not being subjected to
margins on their OTC portfolio use the regular parametric scenario analy-
sis for the management of the exposures. They therefore do not have the
watchdog protection of a narrow-minded exchange to protect them from
the nonstatistical risks.

The asymptotic delta risk loses its significance in the case of dynamic
hedging, as it does not take into account the hedging of the position be-
tween points. Conversely, the asymptotic delta risk measurement does not
reveal the risk of whipsaw that could range far in excess of the amounts dis-
closed. In brief, it only discloses the risk from events where the trader does
not have time to hedge.
[End OCR]*

## Page 222

*[Image OCR]
Chapter 12

Fungibility, Convergence, and Stacking

Find me an arbitrage in commodities and I will show you a squeeze.
An option proverb

FUNGIBILITY

M@ Fungibility refers to the degree of specificity required for the satisfac-
tion of the deliverability obligation. A highly fungible commodity is an
abstract one and could be created electronically (like a swap). A non-
fungible one is specific and cannot be created at will, such as livestock
that requires the delivery of a cow at a specific place at a specific time.

From the vantage point of financial theory, fungibility reflects the feasibil-
ity of risk-neutral replication of a derivative security.

Fungibility of a commodity represents the degree of dependence on time
and space for the delivery process. This issue prevents theoreticians from
reaching a generalized derivatives-pricing model, as it shows how deep
differences between products could be. Some commodities are totally
theoretical and intangible (an inflation index), others are palpable and
specific.

Whenever there are two instruments in the same category, the market
for the more fungible one will be considerably larger and more liquid. Bond
futures seem to be more interesting than cash bonds because the operators
do not have the specificity of a particular bond to borrow (despite the con-
version optionality). Likewise, swaps are better arbitraged than bonds.

Example: A contract specifying for the deliverability of calves of
some detailed species becomes subjected to the dynamics of supply
and demand for that particular specification. Sometimes open interest
exceeds the number of animals alive of that specification, causing an
imbalance between longs and shorts. The owners of the underlying
asset, typically seasoned traders with a vast experience of the squeez-
ing process, will accumulate the physical and force delivery of the

208
[End OCR]*

## Page 223

*[Image OCR]
Fungibility, Convergence, and Stacking 209

underlying. Within the narrow time window available, it becomes im-
possible to create animals from scratch, regardless of the money avail-
able. The owners of the physical typically know it and would not fail to
benefit from it. The squeeze will] last until open interest numbers drop
to the physical ceiling.

Ranking of Fungibility

Commodities can be classified between two extremes: totally fungible and
delivery specific.

¢« Anextremely fungible commodity is a currency that a bank can create
electronically and wire anywhere in the world. The only restriction
becomes exchange control, the establishment of hindrances to convert-
ibility, something no longer common with the currencies of the indus-
trialized world. Being able to wire currency instantaneously without
any restriction removes it from the time or place restrictions. This has
a momentous effect on the term structure of prices.

¢ At the other extreme lies a physical commodity that needs to be deliv-
ered at one specific location. The location specificity subjects arbi-
trage to shipping considerations. It is therefore possible to see the
same commodity with the same grade trading at two different prices
at two different locations at the same moment because of. the shipment
costs and time. Spot oil can trade for 18.23 for 3-day delivery in the
New York harbor and 18.75 for the same delivery date of the grade in
Rotterdam. Since moving inventory at an acceptable price within the
deadline of the delivery is impossible, a certain margin between secu-
rities will be allowed. For the right price, however, someone would
charter a Concorde and fly any shipment anywhere.

Nonfungibility makes elements of one grade trade at a different
price than those of another grade. Time and place restrictions oper-
ate in more than one way. Storage can involve such costs that those
holding a product in the ground would dictate the marginal carry of
the commodity.

Example: Oil in the ground does not really cause the owner to incur
any direct physical carrying cost. If the owner is a government that is
not chartered to sell a chunk of the homeland, no financial carry is usu-
ally taken into consideration. On the other hand an individual who
owns the land could always make a comparison between its future mar-
ket value and the investment of the proceeds in some treasury bill.
Someone with no marks-to-market has no real financing cost to take
into consideration.
[End OCR]*

## Page 224

*[Image OCR]
210 Measuring Option Risks

So there is no true force to prop up the price of the forward, except,
of course, the behavior of the real users who can buy long dated for-
wards to lock in the price of future needs, which, for trading purposes,
can be deemed at best negligible. The forward price can therefore trade
at a much lower price than cash, and the difference between the cash
and the future can fluctuate wildly without any trader having the guts
to confront the market.

Example: Often the commodity can escape all rules when it comes to a
perishable or live animal. No financial power can electronically create a
live animal (not yet given the state of genetic engineering) or extend its
life beyond term to make a long-term forward delivery.

Fungibility and the Term Structure of Prices:
The Cash-and-Carry Line

Risk Management Rule: A calendar spread in options on cash as-

sets in less fungible commodities needs to be decomposed into its
different components and analyzed separately.

The prices follow the following curve: If a commodity one year hence be-
comes too expensive relative to the spot, an arbitrageur can buy the cash,
compute proper carrying costs (both financial and physical storage), then
sell the future to lock in a profit. This puts an upper boundary on the future
costs, provided the commodities do not go through deterioration through
time. This upper boundary is called the maximum contango (see Figure 12.1).

33.5
31.5
29.5
8 27.5
& 25.5
23.5
21.5
19.5

- Qo Oo Q Oo eo So o oO =)

S8S3egBeR sg &

Days

Figure 12.1 Boundary for a commodity’s forward prices: Maximum contango.
[End OCR]*

## Page 225

*[Image OCR]
Fungibility, Convergence, and Stacking 211

Example: A commodity is trading at $20.00 per unit, incurring storage
(and insurance) costs of 8% annualized and fixed. The maximum con-
tango line will espouse the shape of the yield curve for interest rates plus
the storage costs. It is not possible for a commodity to exceed the line, but
the yield curve can take any shape underneath it (see Figure 12.1).

So commodities are split into:

¢ Those that can only be arbitraged one-way on the time line, like most
physical commodities.

¢ Those that can be arbitraged two-way, like currencies and liquid eq-
uities and fixed income instruments.

¢ Those that are not arbitrageable on the time line, like all perishable
products.

. Table 12.1 classifies commodities according to their fungibility.

Table 12.1 Classifications of Products by Their Degree of Fungibility

Commodity
Currencies, gold

Cash bonds (United
States, Germany)

Euro (Eurodollars,
Euromarks, etc... .)

Stocks

Agricultural
commodities

Fungibility
High in the absence of
exchange controls

Medium except for
“tight” issues (that are
difficult to borrow in
the repo market)

None, since the
contracts do not overlap

High in most liquid U.S.
issues

Low in Japan and
Germany

Impossible when the
back months trade at a
discount, or when the
products are highly
perishable

Low in cases of cash-
and-carry

Stability of Term—
Structure of Volatility

High

Medium

Low for the front
contracts, higher for the
back months

High even in case of
unstable borrowing

None
[End OCR]*

## Page 226

*[Image OCR]
212 Measuring Option Risks

Fungibility and Option Arbitrage

Option arbitrage consists principally in synthetically moving the liquidity of
one instrument against another, that is trading between the liquidity of the
products. A fungible product will subscribe to the line F(t) = F(0) + Carry,
with most Greeks of the forward fungible with those of the cash, plus or
minus a variance. The initial Greeks such as delta and gamma can be used for
the cash or the future interchangeably.

Another way is to consider the structure of correlation between the
products. A high correlation matrix will lead to a spreadable volatility
curve. A low correlation matrix requires the bucketing of expirations as
semiindependent products.

The stability of second derivatives will be higher than first ones, and so
on, as one goes into higher order differentials. This means that if the deltas
of the 9-month does not properly hedge that of the 3-month, the gamma
hedges will be more stable. The DgammaDspot will be even more stable and
so forth. A trader should initially square up his delta in each expiration,
then concentrate on the gammas, and so on.

A trader who is long a 1-year call and short a 2-year call in oil should
first do a 1-year/2-year swap or its equivalent. He will end up with resid-
ual delta in the rallies as he will get longer the 1-year forward and short
the 2-year. He then will have to square up the incremental deltas back and
forth.

This leads to the following rule:

Risk Management Rule: Option arbitrage cannot be performed in
markets that do not have a risk-neutral forward market. The for-

ward needs to be perfectly arbitrageable for the option operator to
feel confident about his activities there.

Changes in the Rules of the Game

Currencies appear to be arbitrageable on both sides of the forward line. At
times of crisis, however, a currency under siege can command high interest
rates or might turn out to be unobtainable. An operator will still be able to
deliver it and incur a liability but covering that liability by borrowing the
currency overnight could prove onerous. The Irish punt commanded 4000%
overnight interest rates during the 1992 financial crisis. The central bank’s
stated objective was to make it as difficult as possible to borrow, “to give the
speculators a lesson.” The Spanish peseta after a weakening against the
other European currencies will become tight as many short sellers need to
[End OCR]*

## Page 227

*[Image OCR]
Fungibility, Convergence, and Stacking 213

Price

Dec- <Apr- Aug- Dec-
95 96 96 96

Figure 12.2 Live cattle prices on Friday October 27, 1995, Chicago Mercantile
Exchange.

borrow it to cover. Sometimes the central banks create a wedge between the
domestic and the offshore (or Euro) markets.

Example. Live Cattle: Figure 12.2 represents the curve of cattle prices.
As expected, the curve does not exhibit any smoothness: It is difficult to
transfer a live animal into December 1996, when the animal could be-
come sick.

The volatility is difficult to gauge owing to the illiquidity of the
back month. Whereas it is the rule for a fungible product that the price
of a back month at-the-money (forward) straddle to be at least that of the
price of a straddle of a shorter expiration (minus financing), such a re-
striction does not apply to such instruments. The straddle in the back
month could trade at a lower price than that in the front without any
possible arbitrage.

CONVERGENCE

M Convergence trading! is any form of speculative position taking based
on the belief that the (risk-adjusted) term structure of a security or its
derivatives (interest rate yields, volatility, gamma costs, etc.) does not
represent a completely unbiased predictor of the future.

Securities are supposed to include some bias in proportion to their risk (in
the form of covariance with the market or liquidity constraints). Typically,
financial market equilibrium would, according to theory, represent a bal-
ance between the drift of a security and its risk premium. However, Wall
Street does not hold a coherent and stable belief in risk premiums and
would therefore consider a carry to be whatever disagrees with their vision
of the future price.
[End OCR]*

## Page 228

*[Image OCR]
214 Measuring Option Risks

Example: An imaginary commodity future with an extremely high
drift:

Spot price: 100.

1-month price: 110

2-month price: 125

3-month price: 145

The structure of the financial markets indicates that (assuming no risk
premium for the commodity) the price of the instrument should pick up 10
cents in the next month. This is called the “ride up” on the curve. It implies
that the expected price 1 month from now is 110, 2 months from now 125,
and so on.

¢ First-Order Convergence. An operator who shorts the one-month fu-
ture can expect that if the security does not drop by $10.00 a month,
.33 a day, he will benefit from the positive convergence of his trade.
This operation is called first-order convergence, as it seems risky.

¢ Second-Order Convergence. The operator who prefers not to incur
the risks of a naked short position would engage in a spread. He
would buy the 1-month forward at 110 and sell the 2-month at 125.
That way he can consider that he reduced his directional risk (as-
suming the two instruments are highly correlated).

¢ Butterfly Convergence. A shrewder operator, provided he can ac-
cess markets without incurring high transaction costs, would look at
the previous trade and consider the trader rather naive. The trade
leaves the person exposed to the shifts in the curve generally result-
ing from imperfect correlation between instruments. So he adds a
twist: a butterfly. ;

He buys the second month twice and sells the first and third
month against them, thus achieving some relative safety from moves
in the marketplace. This is called two-factor immunization: The first
one is the direction, the second one is the slope of the curve. The op-
erator leaves himself only exposed to the convexity in the shape of
the curve.

There is a Wall Street continuous dilemma between those who believe that
the future converges to spot and those who believe that the spot “decon-
verges” to the future. In other words, the question is whether the curve cor-
responds to some utility function, some structural condition, or to market
expectations. Many operators make different judgments at different times
depending on the information available in the market. They may believe in
convergence at some times, say in the absence of a central bank policy and
subscribe to the expectation belief at other times.
[End OCR]*

## Page 229

*[Image OCR]
Fungibility, Convergence, and Stacking 215

Option Wizard: Carry versus Convergence

One should not confuse carry and convergence. Convergence includes carry.
Carry corresponds to the cash flow differential between instruments or the dif-
ference between an instrument and the cost of holding it (or the “risk-free”
rate according to financial theory). Convergence is calculated by repricing the
instrument with a decrease on the term structure in addition to the carry.

Most sophisticated convergence calculations take into account seasonal
factors, such as year-end liquidity.

Mapping Convergence

For a Fixed-Income Instrument. Dollar convergence is equivalent to:
Carry + Drop on the curve in basis points

The drop on the curve should be computed in by breaking down the cash
flows of every instrument and repricing them on the zero-curve with one
day shorter to go.

For a Eurodollar Future. There is no carry. The drop on the curve can be
more easily computed in two ways:

1. The easy way is by interpolating the straight line difference between
contracts and repricing the instrument as if it were one day shorter.

2. The more accurate method is to use the polynomial approximation for
the time function or the spline in order to avoid abrupt changes in
convergence. The price of the securities would be replaced by the sea-
sonal adjusted one. The December future, for example, would be
purged of the liquidity premium. (Year-end interest rates are often
higher because of balance sheet requirements. Operators need to take
such effect into account.) ,

For a Currency Cash Position. For a currency cash position, carry is
the convergence. It is estimated by checking the overnight interest rate
differential.

For a Currency Forward. A currency forward is only a currency spot po-
sition plus two fixed income zero-coupon positions, one long and the
other short.
[End OCR]*

## Page 230

*[Image OCR]
216 Measuring Option Risks

The convergence would be priced as the spot carry plus the difference in
the effects in the drop on the curve in each of the two instruments.

Convergence and Convexity

Often time decay can be mistaken for convergence. Instruments that pre-
sent a high convexity will often have a theta built into them.

The bond future includes an option to choose between a variety of de-
liverables. Estimating its convergence would require some adjustment to
take into account the “time decay” of the embedded option.

Securities of the Government National Mortgage Association (GNMAs)
need to be priced on an option-adjusted basis. The owner of the structure
has a high level of negative convexity in his books.

Levels of Convergence Trading

¢ First-Order Convergence Trading. The positive carry trade. Owning
a bond because of its carry.

¢ Second-Order Convergence Trading. The forward-forward. Playing
the carry differential between two maturities.

° Third-Order Convergence Trading. Butterflies, barbells.

¢ Fourth-Order Convergence Trading. Double butterflying, running a
minimum variance stacking program. (See next section.)

Volatility and Convergence

Many speculators (called “carry hogs”) are attracted by “positive carry” in-
struments. Typically, these instruments will exhibit a skew behavior with a
strong volatility in the event of a sell-off. Often the holder’s premium is
only a compensation for the holding risks.

Convergence and Biased Assets ’

There was a corny theory on Wall Street that assets, particularly those of
foreign countries that carried a high return were attractive on a risk-
adjusted basis. True, these assets were deemed risky, but they presented a
better than average risk/return. Returns were deemed to be convergence.
Risks were interpreted as their historical volatility. Traders who traded in
that manner were in their own eyes sellers of expensive insurance. The
method extended to diversification techniques where uncorrelated high-
yielding products were bundled in a package that would present what ap-
peared to be abnormal returns.

Needless to say, those traders disappeared. The use of the sacrosanct
correlation methods caused the buildup of positions in currency pairs that
[End OCR]*

## Page 231

*[Image OCR]
Fungibility, Convergence, and Stacking 217

appeared to be highly correlated but presented a very high carry/conver-
gence. Speculators flocked to funds that did the following trades:

¢ They bought the Italian lira, which commanded a high interest rate,
and. borrowed Dmarks, which carried a low interest rate, and were
lulled by the protection of the European Monetary System.

¢ They bought the Scandinavian currencies against the DM for the
same reasons. Historical volatility appeared to be low.

¢ They invested in Mexican fixed income assets comforted by some
“pacto” where the Mexican government committed to stabilize the
currency within a band.

These assets proved to be heteroskedastic (i.e., had a changing volatility
structure: 3-month historical moved from 2% to 50%) and presented a
skewed behavior in the bargain. Massive losses ensued, both in September
1992 and in the Mexican debacle of 1995. Except for brief periods, however,
the fund managers and traders presented outstanding “Sharpe ratios,” a
measure that is useless with skew and fat tails.

The traders collected fees and bonuses during their previous years
thanks to the optionality of their payoff.

STACKING TECHNIQUES

Stacking is a short-term hedging technique that aims at minimizing the ex-
ecution of a multiple leg hedge by concentrating on a few liquid instru-
ments that track the position best.

Stacking is best used as a transitory hedge. It is used by cap/floor and
swap traders who can thus hedge the residual deltas stemming from a
market move (the secondary, gamma-related hedges). It is also prized by
market makers unwilling to commit time to the fine tuning of a position
they may not carry on the books for long. It is also practiced in basket trad-
ing. The risks of stacking increase with time: Relationships and hedge
ratios change. The technique of stacking is not meant to be a permanent
hedge, only a temporary maneuver around the liquidity of a market.
Stacking a position permanently can cause serious troubles. It is to be re-
membered that Metalgessellshaft, the large German concern that experi-
enced a ten-digit loss from an oil hedge gone awry, was market neutral but
had stacked all its exposure in the front future.”

Example. The Quick-and-Dirty Hedge: A trader needs to rapidly hedge
a two year strip in the Eurodollar markets, by selling the equivalent of 95
of each of the first expiration (for the purpose of simplification, assume
that the hedge ratio is equal in each maturity). Given the volatility of
the Eurodollar markets, he may need to do it quickly.
[End OCR]*

## Page 232

*[Image OCR]
218 Measuring Option Risks

Table 12.2 Simple Stack

Hedge Net
Position 1 Exposure

Eurol —95 -—95
Euro2 —95 —95
Euro3 —95 —95
Euro4 -95 —95
Euro5 —95 760 665
Euro6 —95 —95
Euro7 —95 —95
Euro8 —95 —95
Unweighted
Total —760 760 0

¢ A Market-Neutral Stack. An easy method would be to compute the
amount of total exposure, that is 760 contracts, and then sell them all
in one expiration to ensure market neutrality. The trader would pick
the fourth or the fifth expiration and sell 760 contracts.

¢ A Butterflying Stack. A stack that would offer some protection for
the shifts in the shape in the yield curves.

In Table 12.2, the net exposure corresponds to the risks the trader has on
the books after stacking. -

Traders need to hedge as an emergency. Working eight orders at the
same time can cause confusion, partial fills, and headaches as no broker can
concentrate on all the expirations at once. Even when there is a market for a
strip as a spread, bid /offer are wider than every individual leg.

The trader needs to hedge the primary risk as rapidly as possible, then
fine-tune the resulting mismatch by working spreads (Table 12.3).

The trader could further lower the risks by putting spread orders in the
market and waiting for fills to come at more attractive prices than if he exe-
cuted it as a strip (or less unattractive price, as is usually the case). It could
be done the following way: Trader can buy 95 Eurol against Euro3, buy 95
Euro5 against Euro7, and so on, thus reducing every leg until the position
becomes entirely flat.

Example. Minimum Variance Stacking: As some. securities do not
closely track each other, there are methods of optimizing a hedge. The
correlation matrix can be used to search for the best combinations or
the most appropriate stack.

In the preceding example, the front Eurodollar future does not track the
market very well. The back-month futures, on the other hand, are almost in-
terchangeable.
[End OCR]*

## Page 233

*[Image OCR]
Fungibility, Convergence, and Stacking 219

Table 12.3 Butterfly Stack

Hedge Net
Position 2 Exposure

Eurol —95 —95
Euro2 —95 —95
Euro3 —95 380 285
Euro4 —95 —95
Euro5 —95 0 —95
Euro6 —95 —95
Euro7 —95 380 285
Euro8& —95 —95
Unweighted

Total —760 760 0

The computations in Table 12.4 are based on one year of data as of May
15, 1995. ED1 is the perpetual front-month Eurodollar contract, ED2 the
perpetual second, and so on.

Volatility is the annualized standard deviation based on daily log-
returns.

A single stack, as shown in Table 12.6, reduces the position to 28% of the
initial risks.

A butterfly stack (Table 12.7) reduces the risks to 14% of the initial
position. ,

A stack taking advantage of the full power of the correlation between
periods reduced the risks to 5% of the initial position with three trades
only: The trader bought 95 ED1, 285 ED3, and 380 ED4. (See Table 12.8.)
Eliminating the first month, as it does not correlate to the rest and does not
have a similar variance, corresponds to the bulk of the reduction. Because

Table 12.4 Eurodollar Correlation Matrix*
ED1 ED2 ED3 ED4 ED5 ED6 ED7 ED8
>

EDI 1.00 0.90 0.84 0.78 0.76 0.74 0.72 0.70
ED2 0.90 1.00 0.97 0.93 0.90 0.88 0.86 0.84
ED3 0.84 0.97 1.00 0.97 0.96 0.94 0.92 0.90
ED4 0.78 0.93 0.97 1.00 0.98 0.97 0.95 0.94
ED5 0.76 0.90 0.96 0.98 1.00 0.99 0.98 0.97
ED6 0.74 0.88 0.94 0.97 0.99 1.00 0.99 0.99
ED7 0.72 0.86 0.92 0.95 0.98 0.99 1.00 1.00
ED8 0.70 0.84 0.904 0.94 0.97 0.99 1.00 1.00
Vol% 10.96 15.90 18.50 18.90 18.23 17.10 16.00 15.30

*Eurodeposit’s correlation matrices show the same general pattern. Eurolira, Euroyen,
Short Sterling, PIBOR, and so on, show the same high correlation in the back more and
higher relative independence of the front contract.
[End OCR]*

## Page 234

*[Image OCR]
220 Measuring Option Risks

Table 12.5 Initial Position

Standard Net

Contract Deviation (%) Exposure
ED1 10.96 —95
ED2 15.88 —95
ED3 18.50 —95
ED4 18.85 —95
ED5 18.23 —95
ED6 17.13 —95
ED7 16.16 —95
ED8 15.32 —95
NET

VAR(000)*

112.99

*NET value at risk is defined as the expected
profit or loss for 1 standard deviation in the
market. The formula for its computation is pro-
vided in Module E.

we found the combination that minimized the variance (to the close-to-
optimal level), such method is called minimum variance stacking.

Other Stacking Applications

° Basket Trading. When trading the SP500 cash-future arbitrage, pro-
gram traders can create sub-portfolios of liquid instruments to concen-
trate their hedges. Hedging a short future with all of the 500 stocks
will not be feasible immediately. It may not be necessary given the

Table 12.6 Single Stack

Vol. (%) Pos (Mil)
ED1 10.96 -95
ED2 15.88 -95
ED3 18.50 —95
ED4 18.85 -95
EDS 18.23 665
ED6 17.13 —95
ED7 16.16 —95
ED8 15.32 —95
NET

VAR(000)

32.16
[End OCR]*

## Page 235

*[Image OCR]
Fungibility, Convergence, and Stacking 221

Table 12.7 Butterfly Stack

Vol (%) Pos (Mil)
ED1 10.96 —95
ED2 15.88 ~95
ED3 18.50 283
ED4 18.85 —95
ED5 18.23 —95
ED6 17.13 ~95
ED? 16.16 288
ED8 15.32 ~95
NET

VAR(000)
17.76

Table 12.8 The “Smart Stack”

Vol (%) Pos (Mil)
ED1 10.96 0
ED2 15.88 —95
ED3 18.50 190
ED4 18.85 —95
EDS5 18.23. * ~95
ED6 17.13 —95
ED7 16.16 285
ED8 15.32 —95
0
NET

VAR(000)
7.36

speed with which index arbitrageurs reverse their basket position.
They identify beforehand a number of liquid stocks and stack their po-
sition in a subbasket. In a trending bull market, as basket traders find
themselves chronically short the future, they can perform “pairs trad-
ing” to reduce the risks.

Index Replication. Traders short USD-ECU rapidly replicate the po-
sition by stacking USD-DM (United States dollar against the Deutsche
Mark). They will later cross the position from the USD-DEM into
other components by hedging the less volatile crosses such as DEM-
Drachma, DEM-Guilder, DEM-FRF (French Franc).
[End OCR]*

## Page 236

*[Image OCR]
Chapter 13
Some Wrinkles of Option Markets

Sell [volatility] on Thursday and buy on Friday.
An old option proverb (about volatility selling prior to a weekend)

EXPIRATION PIN Risks

@ The expiration pin risk is the expiration variance for an options posi-
tion. It results from the absence of timely information about the out-
come of the option assignment lottery.

There is a lag between the close of an option market and the notification to
the parties short options about whether or not they were assigned. The lag is
necessary for processing reasons, but when the option is too close to the
money, there could be a great deal of uncertainty about the outcome. The
pin prolongs the uncertainty and thus the life of the option, but sometimes
creates a contingent claim that is not rewarded for.

Risk Management Rule: Put/Call parity does not hold for non-

cash-settled listed options (both European and American).

The expiration pin risk could be significant when the operator is involved in
a conversion or a reversal. The notification of the exercises is due at least 24
hours before the answer about the results of the assignment process.

Example: A market on a listed exchange is closing at 100 exactly. The
trader is short the 100 calls, long 100 puts and long the underlying fu-
ture—what a text book would call a perfect hedge (the so-called conver-
sion). The trader does not know whether to exercise or not. Should he
exercise all his puts and deliver the futures, he would hope that the
calls get entirely abandoned. Should he abandon the puts, he would
wish to see all his calls exercised so he can deliver his futures. The
party on the other side will be in the same quandary.

222
[End OCR]*

## Page 237

*[Image OCR]
Some Wrinkles of Option Markets 223

What makes matters even more complicated is that a closing price on a
future is not a real indication as to where the market will be at the times of
exercise. The future closing does not usually terminate the option: There is a
gap of a few hours in which information can still hit the markets with opera-
tors guessing the general effect of such information without more precision.

In the preceding example, assume that the future closes at 101 but that
after the close news of a government scandal hits the market. The operator
may wonder what to do with the 100 puts, officially out-of-the-money. He
knows with some certainty that the market would be lower if it were open
to trade. His risks are that he may abandon the puts and not be assigned on
the calls with the market opening markedly lower. The other risks are that
he may exercise the puts and still be assigned on his calls, with the market
unchanged if the news is subsequently denied after the deadline. Often
some of the counterparties with a large open interest can manipulate the
markets and turn the uncertainty in their favor, as when they become in
possession of a large sell order and decide to exercise options deemed out-
of-the-money.

Some solutions have been offered by exchanges to reduce pin risks, none
of them satisfactory. One of them is to allow conversion-reversals to be net-
ted provided one can find a counterparty that has exactly opposite sides.

The only markets where a conversion or a reversal is a pure arbitrage is
where the underlying is a cash-settled future that expires at the same time
as the option. Examples: Eurodollars, Euremarks, Pibor (the quarterly op-
tions), SP500 (quarterly’).

Warning: Traders, like poets, often give the same designation to two dif-
ferent items. “Pin risk” is also used to refer to the P/L swing at the barrier
for a portfolio containing barrier options.

Sticky STRIKES

B Sticky strikes are over-the-counter or listed strikes in which the
buildup of a large open interest alters the behavior of the market around
the strike price near expiration.

Sticky strikes are usually magnified by the concentration of the long or
short open interest in the hands of one or more parties that do not delta
hedge. The pressure on the underlying occurs because the option traders’
behavior will diverge from that of the customer.

For example, if a large coverered writer is short the SP100 450 calls, op-
tion traders burdened with long gamma will tend to buy and sell around
the strike, creating to their detriment some stability around their worst-
case-scenario price. There will be bids underneath the strike and offers
above. Meanwhile, the party short the largest portion remains indifferent to
[End OCR]*

## Page 238

*[Image OCR]
224 Measuring Option Risks

the market being only interested in the terminal state of in-the-moneyness
or out-of-the-moneyness.

The concentration in the hands of a nondelta hedger will cause such
stickiness. As expiration nears, locals will need to buy all their deltas below
and sell all of them above and the underlying will tend to stick around the
strike until some fresh supply or demand pushes it away. If the same open
interest were distributed in the hands of market makers, the stickiness
would be greatly reduced as there would be two-way orders by long match-
ing short gamma.

Typically when such a strike is reached on expiration day (and there are
high odds of having a strike crossed), it acts as an absorbing state for the
market. There will be large sellers above it if approached from below and
large buyers underneath it if approached from above.

These conditions where the option’s open interest determines the path
of the underlying are starting to emerge as many customers are picking de-
rivatives in place of the primary securities. Warrant traders are most sub-
jected to that effect as the ratio of secondary delta hedgers to total traders
there is greatest.

Note: Cash traders complain that it is often the case of the tail wagging
the dog. Very little research has been done to ascertain the impact of expira-
tions on the behavior of the spot.

There is a saying in currency options markets, where daily expirations
tend to affect the behavior of the spot, that one always has to gain from sell-
ing a cheap overnight option or buying an expensive one. When overnight op-
tions are cheap, the market will be stabilized by the traders sucked into it.
The market at 10:00 a.m. New York will converge to the strike should spot be
in the vicinity. Conversely, the market will whip when options are expensive.

Risk Management Rule: When dynamic hedgers are long a strike

(and consequently, static hedgers short it) the strike will be sticky.
It will whip otherwise.

MARKET BARRIERS

m@ A market barrier is a level that is supposed to cause some stickiness be-
cause of institutional constraints.?

Market barriers should not be confused with barrier options. The following
are examples of market barriers:

¢ Acurrency band where central banks limit the market from trading
through some level by intervention.
[End OCR]*

## Page 239

*[Image OCR]
Some Wrinkles of Option Markets 225

e A floor for the price of an agricultural commodity guaranteed by the
government.

Market limits that, when reached, cause the exchange to shut down repre-
sent a weaker form of market barrier.

A Currency Band: Is It a Barrier?

Some economists consider that market barriers should cause the volatility
to be dampened in their proximity. Somehow the claim in the literature
that a target zone is heteroskedastic proved to be true, but exactly in
reverse.?

Figure 13.1 shows the “rubber trees” shrinking as the market comes
close to the barrier. Obviously, a market at, say, 20 pips from a limit cannot
go up more than that. To compensate for that fact (assuming constant
skew), it should be limited to a move down by 20 pips, and so on. The
volatility drops as one comes close to the barrier, and the asset starts “past-
ing” around it.

A more complex analysis would increase the skew, by saying that the
market on the way down would be able to drop by more than 20 cents but
that the probability attached to it should be low to satisfy the following
“fair” game (assuming for simplification that interest rates are negligible):

pu = (1 ~ pid

with p the probability for the market to go up, u the size of the up-move,
(1 — p) the probability for the market to go down and d the size of the
down-move.

Traders do not buy the argument: The volatility of the forward should in-
crease dramatically as one comes close to the band. Every time the markets
approach such a band, some warfare erupts between speculators and central
banks. In addition, as by some phenomenon physicists call hysteresis, the mar-
kets have a tendency to snap through the barrier with a vengeance. »

Te Tee]

Figure 13.1 The rubber tree around the band.
[End OCR]*

## Page 240

*[Image OCR]
226 Measuring Option Risks

The Absent Barrier

Practitioners discovered (at great expense) that the reflecting /absorbing
barriers set as boundaries in the analysis of currency bands turned out to
be optical illusions. Before the explosion of the European Rate Mechanism,
there was a great deal of literature on the topic.

Somehow when a currency reaches a band in the spot, which is the only
observable portion of the trading, the forward continues to trade unfet-
tered. What appears to be extreme interest rate volatility is nothing but the
volatility of the currency being translated in interest rate terms.

Spot is not necessarily what is visible on the screen. Spot for delivery in
a week is a sum of spot + forward “points,” themselves function of the in-
terest rate differential. The trader buys or sells spot and then buys or sell
the forward points. There is no real market for nonsynthetic forward out-
right. The method of calculation is as follows:

Ss =Fx eft2 > rit

S is the spot, F the forward.

If one freezes S, F will continue trading. For S to remain constant,
(r2 — r1) will adjust. Generally, it is the currency’s interest rate, here r2
that will take the brunt of the adjustment.*

The impact on an option’s position is that a European option, or one
that is not for cash delivery, will trade off the forward, totally unconcerned
with spot. Notice that 99% of vanilla currency options are European.° Such
an option will go through the barrier like a knife in warm butter. Its price
wil] adjust to the final F. In pricing these instruments, it is therefore neces-
sary to eliminate the idea of a barrier.

Risk Management Rule: A currency band, even when fully en-
forced, does not represent a flawless barrier for a European option

unless interest rates for both elements of the pair are banded as
well.

Wuat FLat MEANS

A flat position is a position that does not present any market risk. For cash
and short-term products, a flat position is a matched position. A derivatives
position is presented as flat only vis-a-vis a particular mathematical partial
derivative. One could be flat delta (i.e., the first derivative function of spot)
but not gamma or vega. Flat gamma is local and one could be exposed to the
[End OCR]*

## Page 241

*[Image OCR]
Some Wrinkles of Option Markets 227

Option Wizard: An EMS War Story

A trader who survived the EMS (European Monetary System) breakup of 1992
recounts the following. In the events leading up to the turmoil of September
1992, he had to execute a large option order in sterling versus German marks
for a customer. He needed to buy quantities of out-of-the-money puts on the
sterling, calls on the mark, struck 10% outside the official government band.
The customer was a conspiracy theorist fund manager who believed in the im-
minent breakup of the monetary order. He belonged to the small coterie of
traders who took on the Bank of England later that month.

The trader, being risk averse, decided to do what most of his headache-free
peers do: call a large market maker, add on a margin, and earn the difference.
He called two of the largest option dealers in the world asking for their selling
price and received the following answers:

1. W., based in the United States, and a former CME pit trader, told him that
they were reluctant to show a price “because the strike is outside the band”
and the options were “too risky.” They would accommodate him if necessary,
but at very expensive implied volatility and only for a moderate amount.

2. B., based in Europe, literally laughed at him. “But Zey are outside the band,
if | am not mistaken,” he was told. “How many do you vant? | can sell you
all you need. You should give your money to charity instead.”

W. was a former pit trader surrounded by former pit traders. He is now
the head of worldwide foreign exchange for a large bank. The second dealer,
B., was a graduate of a prestigious European school of engineering. He is now
back to engineering and other precise activities after his desk was decimated
by disastrous losses in September 1992. After all, in the physical world, barri-
ers are barriers, bridges are bridges, and horses are horses.

a eS SS

s

third derivative of the security price to spot, usually the indicator of the
risk reversal, commonly called DdeltaDspot, and so on.

The need to understand the relativity of the flat concept needs to be
stressed for noninitiated managers of derivatives traders, or users of the
_ product who, unable to exactly offset a position for lack of liquidity, start
» running trades to reduce the risk. Running trades against a book to offset a
Greek (i.e., a first-order or second-order partial derivative) causes complica-
tions for the non-market maker.

Path-dependent and barrier options are never flat with respect to par-
tial derivatives for long. The pin risk for barriers makes them only hedge-
able dynamically. Binary options are never flat.
[End OCR]*

## Page 242

*[Image OCR]
228 Measuring Option Risks

Primary and Secondary Exposures

A primary exposure for an option trader is one that is caused by the initial
trade. A market maker selling a cap has as an initial exposure the delta
equivalents in the FRAs or Eurodollar futures. A foreign exchange option
trader has to hedge both his initial delta and forward exposure.

A secondary exposure is the one that results from changes in one or
more parameters that determine the price of the derivative. It can come
from the convexity of the book with respect to volatility or the underlying.
The trader may have to fight the market if he is long gamma and would have
to chase the market otherwise. He may have to adjust the volatility expo-
sure as the curve shifts and causes both the shorts and the longs to be
weighted again.

Secondary exposures can also stem from the bleed, the overnight changes
in the deltas, gammas, rhol1, and rho2.
[End OCR]*

## Page 243

*[Image OCR]
Chapter 14
Bucketing and Topography

Never hire a very well dressed option trader.
T.G.

This chapter will cover common methods of risk management as used by in-
stitutions. They tie in with some of the analyses discussed in Chapter 9 on
vega. These should be mandatory reading for persons who understand op-
tions but have never examined the risks of a book taken as a whole.

STATIC STRAIGHT BUCKETING

@ Bucketing represents the breaking down of the risks of the position
with respect to a particular parameter in time intervals. Operators thus
obtain the delta buckets, vega buckets, gamma buckets, and so on.

It is a static risk management method because the markets are assumed to
be constant. No convexity truly shows; Vega and rho, for example, can vary
in some combination of instruments with the asset level or, sometimes, may
increase or decrease as the parameter changes. The effect is assumed to be
beastly with compound, barrier, and other exotic options.

It can already be detected that the bucketing does not show higher mo-
ments of the sensitivity of the portfolio to the asset price.

@ Straight bucketing is a simple method that shows the exposure be-
tween cash and the expiration time t. It is only applicable to products
that have a known and certain expiration date and start immediately.

Example. A Bucketing Method for a European Option: Start with a cur-
rency position in GBP-USD (Great Britain pound against the United
States dollar).

Spot is 1.6050.

229
[End OCR]*

## Page 244

*[Image OCR]
230 Measuring Option Risks

The trader buys a 6-month call (183 days) in the amount of GBP
100 mil. The price is 4.578% and the trader pays $4,578,000 for it. Its
hedge is 50 deltas. The following parameters hold:

183-day USD rate 5.8438% (basis 360 days)
GBP rate 6.915% (basis 360 days)

[1 + (183/360 X .058438)]
{1 + (183/360 x .06915)]

Table 14.1 shows deltas, gammas, vegas, Rhol, and Rho2 (all in USD).
All measures are unweighted. The bucketing is called “straight” because
the exposure in the 6-month corresponds to an exposure starting at time 0
and ending in 6 months. It is easy to see that the delta as shown in Table 14.1
corresponds to the present value of the forward delta. Thus the spot delta is
equal (in 000) to (82656) / [1 + (1.83/360 X .06915)] = (79850).

All the other elements are straightforward.

The gamma of GBP 8,355,000 means that the position should increase by
such amount should the market rally 1% and decrease likewise in a sell-off.
This is an approximation: The position will not actually pick up such an
amount in a rally owing to the third derivative (it is now at maximum
gamma being at-the-money).

The vega is straightforward: $438,000 for a rise of 1 volatility percentage
point. .
The rho domestic corresponds to the sensitivity of the option to a
change of 100 basis points in the U.S. rates. The Rho2 foreign corresponds to
the same with the foreign rate. However, the Rhol is lower because it also is
sensitive to the discounting of the premium the other way. The methods for
calculating rhol and rho2 are as follows:

The forward = Spot X = 1.5973

¢ Foreign Rho2. The procedure to obtain rho2 is to multiply the delta
by the rate and interpolate for 183 days. It is (183/360) x 82656 x 100
bp = $420,168. The present value of that using the U.S. rate is 408,000

(assuming no convexity). ,
* Domestic Rho1. itis the same number minus the exposure to the fi-

nancing of the premium. The premium was $4.57M. The effect of the

Table 14.1 Buckets
(000) All lw im 2m 3m 6m 9m ly 2y S5yt+

Delta (79850) (82656)
Gamma 8355 8355
Vega 438 438
Rhol 385 385

Rho2 (408) (408)
[End OCR]*

## Page 245

*[Image OCR]
Bucketing and Topography 231

rise of the U.S. rates would be to cost him more in the daily financing
of his premium, by 4.57 X 183/360 = $23,200. Roughly, the net Rhol
is 408,000 — 23,200 = 384,800.

Note: The effect of convexity over this position would be to lower the nega-
tive P/L stemming from a rise in foreign rates and increase the positive P/L
stemming from a drop, and vice versa for domestic rates.

American and Path-Dependent Options

The bucketing method works in general better for European options be-
cause the buckets are only partially unstable. They could be increased or
decreased (through the gain or loss of gamma or vega when the option
moves in on out of the money), but it is not possible for the exposure of one
particular option to change buckets.

With American and path-dependent options, the exact duration of the
exposure is uncertain. This makes the bucketing method often inappropri-
ate. A barrier that increases in probability of becoming extinguished will
move in vega, gamma, and Greeks closer to the short-term buckets. Likewise
when the market moves away from the strike, American and path-dependent
options start resembling a European option and would have an exposure
closer to the nominal maturity.

Volatility shifts would also exert the same effect on the barrier matu-
rity. A higher volatility shortens the duration of a barrier knock-out option.
Paradoxically, it would lengthen that of an American option.

Some options like the knock-in structures or deferred strike options are
impossible to fit in a straight bucket because their exposure starts at one
point and ends at another, therefore making them dependent on a forward-
forward bucket. They depend more on events that take place during some
time segments than in others. Worse, some structures (like barriers) are
calendar spreads in disguise. The trader can be seriously misled by the net
bucket exposure disclosed by the straight system.

The only way to see the risks of these instruments is to engage in dy-
namic bucketing, which requires the trader to review the position at various
states of the underlying at volatilities and try to match the risks accordingly.

Advanced Topic: The Forward or “Forward-Forward” Bucket

Although rarely implemented, owing to the computational complications,
the need to see a forward-forward option risk appears to be essential with
path-dependent options, options on products such as knock-in options, and
non-time-homogeneous structures such as variable bets. The building of such
a bucketing system obsessed the author for a long time, until the most recent
accretions in numerical techniques.
[End OCR]*

## Page 246

*[Image OCR]
232 Measuring Option Risks

Table 14.2 A Forward-Forward Bucket
(000) 0-lw lw-iIm Im-2m 2m-3m 3m-6m

Delta 82 82 82 82 82
Gamma 8 8 8 8 8
Vega A A A A 4
Rhol 4 A 4 A A

Rho2 (4) (4) (4) (.4) (.4)

The forward-forward bucketing would appear as shown in Table 14.2.

The computer’s method for filling the cells is quite complicated when
the portfolio includes path-dependent options. As there are no analytical
methods for that, it needs to move the cells in such a way as to leave the pa-
rameters affecting the other parts constant.

The forward-forward bucketing needs to be constructed by breaking up
the exposures in distinct volatility periods and running the following sim-
ulation on the total position:

¢ Volatility is raised for one particular bucket but not others. The bino-
mial tree is run with the new higher volatility between the nodes for
that period but is maintained constant for other periods. The differ-
ence in price stemming from that would be the vega.

¢ If the portfolio consisted only of European options, it could be as-
sumed that the volatility does not depend on the location of the
shock but on the general effect over the entire portfolio. Every bucket
would therefore share an equal portion of the total gamma.

¢ The vega and the gamma will be bucket dependent for non-European
instruments. >

This method is reviewed extensively in Chapter 9.

TOPOGRAPHY

™@ Position Topography for an options book is a risk management method
that displays the distribution of the exposures of the portfolio across
time and possible asset prices.

¢ There are two kind of topography reports: strike (or static) topogra-
phy and gamma (or dynamic) topography.

¢ Topography has considerable advantages as it allows for transcend-
ing the Greeks as shown in Tabie 14.3.
[End OCR]*

## Page 247

*[Image OCR]
Bucketing and Topography 233

Table 14.3 Old Days Strike Topography

March June
Strike P Cc P Cc

65 26

70 —20

75 ~20

80 21

85 5 1

90 11 2

95 —5 —19 11 —2
100 —72 -71
105 2 70
110 1 23 11
115 -9
120
125

Strike Topography (or Static Topography)

H Static Topography is a two-dimensional map of the position displaying
the distribution of the exposure in face value horizontally (across expi-
rations) and vertically (across strikes).

When option trading was still at a primitive state and option traders were
dealing with only one or, at the most, two expirations, they used to main-
tain their open positions on a card provided by the clearing house so they
could rapidly see what they had in their inventory. The clearing houses
handed out the cards without charge. Table 14.3 depicts a sample card.

The trader could thus examine the position beyond the deltas and gam-
mas and trade his position instead of trading his Greeks. Such a method, in-
deed, has an admirable tutorial value as it forces the trader to learn the
intricacies of option trading without having recourse to pseudo-mathemat-
ical methods like the Greek reports. Instead of waking up at night with cold
sweats and shouting the abstract, “I do not enjoy my being short 22.23 gam-
mas and 71 weighted vegas,” the trader would phrase his worries about the
position in more precise terms such as, “I am short so many of the 105 calls
that I need serious protection.”

Modern book runners, however, do not have such luxury. As there are
no set expiration dates and no set strikes, the card in Table 14.2 would need
a few thousand lines and columns to reflect the topography of a large over-
the-counter portfolio. In addition, it does not allow for any option structure
beyond the vanilla.
[End OCR]*

## Page 248

*[Image OCR]
234 Measuring Option Risks

So over-the-counter books need a design that allows traders to bundle
the exposures in time and space points and examine their concentration
risks. The strike topography is generally designed as shown in Table 14.4.

The over-the-counter strike topography displays the face value exposure
per bidimensional bucket. The report shows horizontally (strike-wise) the
net of the face value exposures between midpoints on the grid and verti-
cally (time periods) the net of face value between points. So the 1W/100
bucket displays the net of all the trades between two days and one week and
every option struck between 99 and 101.

The strike topography should deliberately exclude most nonvanilla op-
tions, which can be remedied in a separate report. American options, in ad-
dition should be handled with care, because their nominal expiration is not
exactly the expected one. Only soft American instruments should be in-
cluded on the matrix. A more thorough treatment of American options cor-
responds to shifting their nominal maturity into the “omega” or real time
to expiration.

An application of the topography method is the method of squares, ex-
plained in Chapter 9.

Adding Correlated Instruments. It is possible to incorporate more than
one similar commodity by summing them up, that is assuming 100% cor-
relation. Such a method aims only at examining the strike concentration or
the origin of gammas and vegas and can thus simplify without deluding
the trader about his position. For example, if one thinks that it would be
necessary to examine the risks of the French franc and the German mark
in one report for topography purposes only, then all the USD-FRF strikes
would be converted into USD-DEM. The real improvement, however,

Table 14.4 Over-the-Counter Strike Topography .
Spot 80- 85 90 93 96 98 100 102 104 107 110 115 120+

1d -83 86 -12 ~—33 -41 -10 -9 90 97 16 -49 56 15
2d 47 9 18 68 -15 -73 20 54 13 -54 -67 -2™4% -4
lw -33 66 —-18 ~25 45 66 -35 50 —-71 18 27 23 —58
2w 38 34 12 44 ‘55 54 ~53 -41 47 64 -28 -9 37
Im 35 -17 -55 34 3 52 43 7 -~8 -15 30 —27 13
2m 12. 45 2 ~-25 33 38 ~—20 15 5 -21 1 -26 —34
3m -14 ~-27 21 #13 -28 -5 22 -6 -35 13 -24 39 6

6m -9 -11 1 23 20 -28 28 6 ~-11 -29 29 -15 -18
9m -2 -14 -1 8 -6 —6 0 -17 5 1 3 1 3
ly 6 -7 -6 2 1-3 -2 5 9 1 10 9 5
2y 2 5 -2 4 -1 2 4 -1 -2 5 -2 —-4 3
3y —4 3 1 0 -5 -2 0 -2 -3 -5 3. -4 5
5y+ -1 4 0 1 -5 -4 0 4 -2 -4 -5 -4 ~-2
[End OCR]*

## Page 249

*[Image OCR]
Bucketing and Topography 235

would come from the hardly practicable three-dimensional map that al-
lows for the slopes of the correlation.

Scaled Strike Topography. Since the difference between strikes that are
close to the money is more meaningful for short-term options than for back-
month structures, it is necessary to scale. For all intents and purposes the
“gap” between the 100 and the 101 strikes is insignificant for a 5-year option
and very annoying for an overnight structure. The method as shown in
Table 14.5 used to account for the risk difference is the scaling method that,
in place of strikes, examines standard deviations. At 15.7 volatility, a 1 stan-
dard deviation difference between strikes is 1 percentage point, that is, be-
tween 100 and 101. In one year, it is the square root of 252 (252 days of
business), namely 15.7. So the equivalent to an overnight gap of 1 point is
the one-year gap of 15.7 points. Hence, assuming the spot traded at 100 (and
no drift), the report would put in the same column the 101 overnight strike
and the 115.7 one-year strike.

Dynamic Topography (Local Volatility Exposure)

Dynamic topography is a potent method of analysis as it can reveal the sta-
bility of a position through time. It consists in keeping the position constant
and running simulation one day ahead, two days ahead, one week ahead,
and so on, and disclosing the map of the resulting Greeks.

The needs for the dynamic topography are accentuated in an environ-
ment where the positions change rapidly and frequently. A conventional
gamma or vega matrix can obscure the real position if strong changes take
place through time. For example, if the position is long one-week options
and short two-week options in large quantities, the position report would
mask the real issues, which is the risk one week from today for one entire
week. .

Computation Technique. The report is not concerned with the deltas but
simply with the gammas. The operator starts by running today’s gammas
for different asset price levels. Then he would block the options expiring on
the following day and run the gammas one day hence for all of these levels,
and so on. Moving the calendar one day hence means running the global
position as if the calendar date was one day forward.

Table 14.5 Scaled Strike Topography

Standard
Deviations ~4 -3 -2 -15 -1 —-.5 0 5 1 #15 2 3 4
1d 100 -150 50

2d
[End OCR]*

## Page 250

*[Image OCR]
236 Measuring Option Risks

The following report would not perform very well for exotic options
since their path dependency makes every report conditional. In other
words, the exposure on a 100 call 103 knock-out depends on whether or not
the market traded through 103 the previous day.

The map shown in Table 14.6 would result. It may appear similar to the
previous ones but is truly a different animal, as it eliminates the positions
that expire as time goes by.

The preceding topography technique could be further improved with
the use of correlations. The operator would need to run the reports using
the possible changes through time of the correlations between markets.

Shortcomings. There is no known way to derive from the dynamic topog-
raphy the exposures to path-dependent options. A trader might look at the
gamma of the book should the market reach 102 in one month. Whatever
number he looked at would not be accurate if the position included knock-
out options. Say that he had a structure that knocked out at 104.5. Because
the market ended at 102 does not necessarily mean that it did not transit
through 104.5, in which case the gamma exposure would be markedly dif-
ferent. The same applies to a lookback: A pulldown from a high level would
result in a lower gamma than a situation where the market trades at its ex-
tremum (its maximum or minimum).

The complications of path dependence are that many more possible
paths are available than end points and there is no rapid way to visualize
the position taking those into account.

Table 14.6 Gamma Topography (Dynamic)
Spot 80- 85 90 93 96 98 100 102 104 107 110 115 120+

1d -83 86 -12 ~-33 -41 -10 -9 90 97 16 49 56 15
2d 47 9 18 68 -15 -73 20 54 13 -54 -67 -24 -4
lw -33 66 -18 —-25 45 66 -35 50 -71 18 27 283 -—58
2w —-38 34 12 44 55 54 -—53 -41 47 64 -28 —-9 37
im 35 -17 -55 34 3 52 43 7 -8 -15 30 -27 13
2m 12 «45 2 -25 33 38 -20 15 5 —21 1 -26 -—34
3m —-14 -27 21 13 -28 -5 22 -6 -35 13 -24 39 6
6m -9 -11 1 23 20 ~-28 28 6 -11 -29 29 -15 —-18
9m ~2 -14 -1 8 -6 —-6 0 -17 5 1 3 1 3
ly 6 -7 -6 2 1 -3 -2 5 9 1 10 9 5
2y 2 5 -2 4 —-1 2 4 -1 -2 5 -2 —4 3
3y —4 3 1 0 -5 -2 0 -2 -3 -5 3. 4 5
sy+ ~1 4 0 1 -5 —-4 0 4 -2 -4 -5 -4 -2
[End OCR]*

## Page 251

*[Image OCR]
Bucketing and Topography 237

= =

Oe

Time

Figure 14.1 Payoff topography.

Barrier Payoff Topography

A payoff topography is a map showing where a large sum of money could be
lost on a “pin,” a worst-case scenario for a bet option or a reverse knock-out
(see Chapters 17-20). Somehow, the pin risks, which can be massive, do not
show far in advance. Moreover, no true expiration date is attached to the
pin, which makes the expected stopping time a fuzzy function of volatility,
rates, and so on.

There is, however, the need to see the worst-case scenario. This disclo-
sure is greatly simplified because the worst possible case for a barrier lies
on the strike at expiration. So it would suffice to establish the map of the
nominal maturity of the payoffs to prevent the trader from adding to a neg-
ative “zone.” Figure 14.1 shows such a map.
[End OCR]*

## Page 252

*[Image OCR]
Chapter 15

Beware the Distribution

I prefer the judgment of a 55-year old trader to that of a 25-year old mathematician.
Alan Greenspan

M@ Take one maturity and measure the implied volatility differential be-
tween the at-the-money options and the out-of-the-money “wings”
priced off the Black-Scholes model. The tails will represent the general
pricing of the wings while the implied skew will measure the degree of
asymmetry of the distribution.

This notion has already been discussed relative to the volatility surface.
This chapter focuses on the reasons for the existence of the tails and the
skew.

Risk Management Rule: Option traders are likely to get the “bad”
distribution through their market making. Somehow, the anxiety
of markets gets reflected in their position. Economists call this phe-
nomenon adverse selection.

THE TAILs

y
While the “skew” is an evasive issue to most traders, the “tails” are easily
explained. The tails are the volatility of the out-of-the-money options rela-
tive to those at the money using the conventional Black-Scholes formula.

Random Volatility

The prime reason for the higher price of in-the-money and out-of-the-
money options than the Black-Scholes value, is the phenomenon traders call
the “volatility of volatility,” or vvol. It is related to the fourth moment, or
kurtosis.’ There are other related reasons, like the convexity of the vega.
In the old days, however, traders used to disbelieve the phenomenon and

238
[End OCR]*

## Page 253

*[Image OCR]
Beware the Distribution 239

attribute the inflated prices of out-of-the-money options to the “lottery ef-
fect,” which meant that investors were ready to spend small sums to get a
large payoff and that such attraction of the large payoff would blind them to
real value. The lottery effect, in the long run, seemed to have favored the in-
vestors owing to the convexity for which they paid so little.

The following simple explanation of option pricing uses an adaptation
of the Black-Scholes-Merton formula for a variable volatility inspired from
the Hull and White model.’

Assume that volatility follows a process like that of the asset itself,
without any drift. This is a simplifying method aimed at assessing the dam-
age from variable volatility rather than hoping to model volatility itself.
Most techniques of volatility modeling have, at the time of writing, led to
little if any convincing results. The model is further explained in Module G.

Table 15.1 shows the impact of actual volatility swings on a 90-day at-
the-money option priced at 15.7% volatility. Three cases are considered.
Case 1, with vvol = 0.25, means that volatility changes on average about
0.16 per day (i.e., between 15.7 and 15.86), randomly, just like the asset itself.
Case 2, with vvol = 0.5, means that volatility changes twice as much.
Vvol = 0 corresponds to the conventional Black-Scholes case, in which
volatility remains constant and the entire sample is drawn from the same
volatility of 15.7%.

Each case leads to different option prices. It is easily seen that an at-the-
money option is entirely unaffected by volatility shifts, and that out-of-the-
money options benefit the most from it. The last three columns show the
Black-Scholes implied volatility using the conventional method of practition-
ers, Which shifts from a particular price to the implied volatility equivalent.
The trader can thus confirm that the Black-Scholes values are identical to the
vvol = 0 case.

The stochastic volatility model used for Figure 15.1 gives the results of
fat tails in proportion to the volatility of volatility. Since the volatility of
volatility is not known and can hardly be estimated, the reader should re-
tain from the exercise a set of rules rather than a modeling framework.

The results show an increase in values where there is convexity to
volatility (i.e., away from the money). At-the-money options imperceptibly
lose in volatility owing to their minuscule concavity (a call is capped at the
price of the asset).

If the trader were to include correlation in the model between asset price
(or its changes) and volatility (or its changes), he would have significant re-
sults: The smile would tilt left or right, with the following characteristics:

The smile would not change in its convex shape but would pivot to accom-
modate asymmetry between upside and downside. A negative correlation
between AS/S (changes in asset) and Aa/o (changes in volatility) would
result in upside option prices that are cheaper than downside strikes and
[End OCR]*

## Page 254

*[Image OCR]
240 Measuring Option Risks

Table 15.1 The Impact of Volatility of Volatility
Implied Using Black-

Case 1 Case 2 Case 3 Scholes-Merton
vvol vvol vvol Delta (Black-

Strike 0.25 0.5 0 Scholes-Merton) 1 2 3
86 0.13 0.26 0.08 0.03 17.3 20.1 15.0
87 0.17 0.31 0.11 0.04 17.1 19.7 15.7
88 0.22 0.38 0.16 0.05 16.9 19.2 15.7
89 0.29 0.45 0.23 0.07 16.7 18.8 15.7
90 0.37 0.53 0.31 0.09 16.6 18.4 15.7
91 0.48 0.64 0.42 0.12 16.5 18.0 15.7
92 0.61 0.76 0.56 0.15 16.3 17.6 15.7
93 0.78 0.91 0.73 0.19 16.2 17.3 15.7
94 0.98 1.09 0.93 0.23 16.1 16.9 15.7
95 1.21 1.31 1.19 0.27 16.0 16.6 15.7
96 1.50 — 1.56 1.47 0.31 15.9 16.3 15.7
97 1.83 1.87 1.82 0.36 15.9 16.1 15.7
98 2.21 2.23 2.20 0.41 15.8 15.9 15.7
99 2.64 2.65 2.64 0.46 15.8 15.8 15.7

100 3.13 3.13 3.13 0.52 15.8 15.8 15.7
101 2.67 2.68 2.67 0.57 15.8 15.9 15.7
102 2.27 2.29 2.27 0.62 15.8 15.9 15.7
103 1.91 1.95 1.90 0.66 15.9 16.1 15.7
104 1.60 1.67 1.58 . 0.71 15.9 16.3 15.7
105 1.34 1.43 1.33 0.75 16.0 16.5 15.7
106 1.11 1.23 1.07 0.78 16.1 16.8 15.7
107 0.92 1.06 0.88 0.82 16.2 17,1 15.7
108 0.76 0.92 0.70 0.85 16.2 17.4 15.7
109 0.63 0.79 0.57 0.87 16.4 17.7 15.7
110 0.52 0.69 0.46 0.90 16.5 18.0 15.7
111 0.42 0.60 0.35 0.92 , 16.6 18.4 15.7
112 0.35 0.53 0.28 0.93 16.7 18.7 15.7
113 0.28 0.46 0.21 0.95 16.8 19.0 15.7
114 0.23 0.41 0.16 0.96 17.0 19.3 15.7
115 0.19 0.36 0.13 0.97 17.1 19.9 15.7
116 0.16 “0.32 0.09 0.97 17.2 20.0 15.7
117 0.13 0.28 0.07 0.98 17.4 20.3 15.7
118 0.10 0.25 0.05 0.98 17.5 20.6 15.7

119 0.08 0.22 0.04 0.99 17.7 21.0 15.7
120 0.07 0.20 0.03 0.99 17.8 21.3 15.7
[End OCR]*

## Page 255

*[Image OCR]
Beware the Distribution 241

22.0 +
21.0 +

20.0

19.0 + Vol 50%

18.0 +
125%

17.0

16.0 + Vvol 0%

N © Oo
= _

-

Black-Scholes-Merton Volatility

+r 2
- ol

oO
wo
Strike Price

Figure 15.1 Smile effect, 90-day option, stochastic volatility model.

could explain the skew. A positive correlation between them would lead to
positive skew.

Warning: It is difficult to establish a dependence between asset price
and volatility. To show the behavior of asset prices in a skewed envi-
ronment this exercise should be appropriately altered for the following
reasons: -

1. Asset returns and volatility may be correlated, but in a nonlinear
way. Typically, the correlation holds from small moves but may re-
verse for a large one. A clear example is the SP500 futures: Volatility
drops after a small rally but increases after a large one.

2. It is clear to traders who have studied data that the volatility, if any-
thing, is correlated (or, more appropriately, dssociated) to the range in
asset prices, not the asset price or its variation. It is not surprising to
see that the behavior is marked by “thresholds.” The market drops
would not increase volatility if they were to take place withjn the
confines of a known range, particularly after a rally. Such a relation-
ship is not easy to model (in spite of all the ARCH-style attempts).

A frequent question is, Why does a variable volatility cause fat tails??

¢ The primary explanation is that of likelihood of asset prices condi-
tional on states of volatility. Assume random volatility and random
asset prices. It is easy to understand that, conditional on being in the
tails, the most likely state is one of high volatility. High volatility can
more easily take the market to the tails than a lower one. Thus the
tails will have the thickness of the higher volatility.
[End OCR]*

## Page 256

*[Image OCR]
242 Measuring Option Risks

¢ The lower volatility is more likely to keep the market in the middle,
the “peak.” Thus the distribution, when in the middle (i.e., no mean-
ingful changes in the market prices), is more likely to take the shape
of a lower volatility.

The concept is even easier to understand graphically. Figure 15.2 shows two
distributions, one of which has four times the volatility of the other. The
higher volatility distribution has fatter tails than the lower one. The lower
volatility distribution has a higher peak than the higher one. It becomes con-
ceivable that in mixing the two distributions, the higher volatility would
dominate the tails, the lower one would dominate the body.

Histograms from the Markets

@ A histogram shows the relative frequency of a certain magnitude of the
returns (or more frequently the logarithms of the price changes) during
a certain time interval. Typically, a histogram shows daily changes. Fre-
quency is obtained by bucketing the moves and counting the percentage
(or total number) of occurrences in every bucket.

Figures 15.3 through 15.6 show the actual distribution of the following as-
sets: the Japanese yen, SP500 and the 30-year U.S. government bond yields,
and the Euromark futures. The first three cover 10 years and the last one 3
years. They show the histograms of the frequency of the differences of logs
of prices (the returns) plotted against a normal distribution of the same
overall volatility. All of the figures display the high peak syndrome.

The astute reader might observe “high peaks” in place of “fat tails.” In
fact, traders betting against the fat tails typically make bets against the
peak: They try to make some profits when nothing happens rather than
during extreme moves.

s

——— Lower volatility:
dominating peak

Higher volatility:
dominating tails

Figure 15.2. Mixing normal distributions.
[End OCR]*

## Page 257

*[Image OCR]
Beware the Distribution 243

07
40+ y
apanese Yen
40+ Jap
HD+
= a+
g
3 20+
© 20
me 150 Theoretical
Distribution
40
sl
0
oO oo
$BSREREESSLSLRBSEY
ooogocooe eeeeo see 8g
S9SGGGSGSCKCOC OSG COS

Log Daily Changes

Figure 15.3 Japanese yen, sampling of 2,703 returns from January 1985 to June
1995.

SP500

Frequency
8
at

Wy;
. Theoretical
ol ;
Orient aH THITIH
Be BF gf of af gf ge ge af af af af at a at HF 2k
DANSSTHNGSSONDS SONS
Log Daily Changes

Figure 15.4 Standard & Poor’s 500 cash index. The series include all returns from
January 1985 to June 1995.
[End OCR]*

## Page 258

*[Image OCR]
244 Measuring Option Risks

400 +
$50 7
300 +
250 +
a
= Seriesi
¥ 200 +
> Series3
xu 30-year yield
i t50 4
100 ¢
50 +
O Arras t TH HM EE
Eee eeaeege es ee eRE
¥¥ 9 Vr +r GF oe oS KF N BO +

Log Daily Changes

Figure 15.5 U.S. government 30-year yield. Distribution of the natural logarithms
of the yields on a perpetual 30-year bond (new issue rolled in), from January 1985
to June 1995.

140
1
10 s
> Euromarks
we
§ — Sees
e Series
i

se 3s 3 3 sf s& 3 3 3 3 B B Bf

2S 55 SR SSR S HK

FV RP? VT POST NOT +
Log Daily Changes

Figure 15.6 EuroMark futures, 1992-1995. The constant first contract was used.
[End OCR]*

## Page 259

*[Image OCR]
Beware the Distribution 245

Heteroskedastic 5-20 %

Frequency

omoskedastic

Log Daily Changes

Figure 15.7 Changing volatility regimes histogram.

Example. A Tailor-Made fat-Tailed Distribution: For amusement (ona
summer day that qualifies as “high peak”), the author conceived the fol-
lowing distribution: three regimes in the market, each one of which car-
ries a different volatility. The regimes are as follows:

s1: Volatility is 15% (normal market conditions).

s2: Volatility is 5% (holiday mood and summer slumber).
s3: Volatility is 20% (anxiety).

Each regime is equally probable.

The Monte Carlo simulator, sampling between n them, produced the
histogram shown in Figure 15.7.

THE SKEW AND BIASED ASSETS ,

The skew is the assymetry in the distribution. Take a daily move
x, = Log(S,/S, _ ,), with o’ its noncentered volatility (assuming mean re-
turn of 0 as explained in Chapter 6). The (noncentered) skew will be:

1 n x?

nN. a’
which will be positive if there is a positive correlation between x, and x,?
and negative if there is a negative correlation. Intuitively the skew ex-
presses the correlation between the move of a random variable (x,) and
its volatility (x,7).
[End OCR]*

## Page 260

*[Image OCR]
Option Wizard: Graphical Representation of Pareto-Levy*

Many traders hear about the Pareto-Levy family of stable distributions. This is

a broad class of distributions said to be “stable” because they can be shifted.

The trader needs to know no further than that the bell-shaped distribution is

but one particular case of this large and unhappy family. A description follows.
The characteristic function of the Pareto-Levy distribution:

Log f() = i5t— yl tl® (1+ iB(t/Itl) tan(an/2))

where i is the imaginary number V— 1 . When a = 2, it becomes the Fourier
transform of the normal distribution (with also 8 = 0, y = 1,8 = 1), exp (-f).

The function has no second moment (that is no variance) when a < 2 and
no first moment (that is no mean) for a < 1. It means that the function has an
infinite variance, an implication that is scary for anyone involved in the mar-
kets. Physically, the left and right tail never come close to the origin. No more
compact support.

As shown in the following chart, the inverse characteristic function (i.e.,
the density function stemming from it) would have the following shape: As a
goes down from 2 toward 0, the tails get thicker and the convergence of the
density toward the zero probability mark gets slower. As a gets smaller, a
wider and wider graph is needed to show the distribution. That illustrates the
notion of infinite volatility: It is impossible to fit the graph into the confines of
a visual frame.*

Probability
Density
(different
scales)

Even scarier: When the distribution has no mean, no peak is visible, but
lines go up to the ceiling without ever meeting.

Many chaos theorists have complained of the misuse of the Pareto-Levy
class of distributions in the common down-market literature.

it is easier for traders to think of fat tails as a result of changing volatility
than the product of some blowing-out variance. In fact, Pareto-Levy is the fruit
of very high volatility of volatility on a very small volatility.

*This option wizard can be skipped at first reading.
‘Feller 11 (1971) provides a trick to numerically invert the Fourier transform and derive the den-
sity function.
[End OCR]*

## Page 261

*[Image OCR]
Beware the Distribution 247

Figure 15.8 A highly skewed distribution.

Figures 15.8 and 15.9 show two different degrees of skewness in the dis-
tribution. The skew is not easy to translate into skewed volatility surface.
The linear skew measure presented above is too weak a statistic to properly
explain the true dependence between volatility and asset prices. There are
many cases where a dynamic hedger would make money owning expensive
puts and shorting cheap calls in spite of a close to symmetric histogram.

A histogram does not show the path. Suppose that the Syldavian cur-
rency initially goes up and down symmetrically by 1%. As it goes down,
however, three days in a row the volatility increases. Such increase in volatil-
ity is likely to push it down further but can as well bring it back to the origin.
The histogram would look almost balanced. In reality, it would show a
slightly bulging left tail, but nothing serious enough to be detected. What the
histogram. shows well is the classical scenario of a market that only goes
down by, say, $4 (with 20% probability) and goes up by $1 (with 80% proba-
bility). It does not very well show the correlation between asset price and
volatility, only the correlation between asset changes.and volatility changes.

Figure 15.9 A medium skewed distribution.
[End OCR]*

## Page 262

*[Image OCR]
248 Measuring Option Risks

Risk Management Rule: The histogram does not appropriately

show the dependence between price level and market volatility.

How this dependence is to be translated into the distribution of implied
volatility between strike prices is uncertain. The customary method has
been to look at the final histogram, derive a probability distribution and try
to compare it with that implied by option prices. The trader can generate a
density from an option price by taking the second derivative of the option
with respect to the strike price. (See Breeden and Liechtenberger, 1978.) As
a trader, the author is extremely suspicious of this technique as it ignores
the path dependency that arises from changes in volatility. As explained
with the greater fool theory (Chapter 3), the behavior of implied volatility is
a determining factor in option trading. Out-of-the-money puts therefore be-
come great assets not only because the market might drop violently, but
mostly because the behavior of the market at lower levels might benefit their
owner. This leads to the following rule:

Risk Management Rule: Path Dependence. The value of the skew
for a dynamic hedger resides more in the behavior of implied

volatility along the path leading to a terminal value than in the |
probability of the asset ending up on such terminal value.

Biased Assets

The old trader’s adage, “Up the escalator, down the chute,” seems to describe
the behavior of a variety of assets. ‘

M@ Biased assets are assets with an asymmetrical distribution, character-
ized with an increased volatility in the sell-off and, to a lesser extent, a
diminished volatility in the rally.

The structure of their market is such that a sell-off causes anxiety while ral-
lies cause euphoria. Such anxiety would cause severe volatility on the way
down while euphoria generally leads to orderly markets on the way up. An
example that easily comes to mind is the Mexican currency. Table 15.2 pre-
sents a simplified world composed of two poles of extreme regimes for bi-
ased assets. There are naturally shades in between as well as transitional
periods between states.

There are many explanations for biased assets, some of which will be
described in the following sections.
[End OCR]*

## Page 263

*[Image OCR]
Beware the Distribution 249

Table 15.2 A Simplified Two-Regime World for Biased Assets

Characteristic

Market condition

Historical volatility
Implied volatility

Skew

Serial correlation

Correlation with other
assets

Type 1 Regime

Type 2 Regime

A severe break ina
market following a
protracted rally (ora
quiet period)

Increases

Increases markedly,
often overshoots
historical

Flatter skew but higher
volatility

Often negative
autocorrelation,
“whipping” markets

Total breakdown of
correlations. L6w
correlations increase.
High correlations
decrease

A normal condition
where the market offers
high returns through
high “carry” or positive
“trend”

Generally low

Low, generally close to
historical. Out-of-the-
money calls trade
usually lower than
recent historical

High skew at lower
volatility generally from
call selling. The lower
the volatility, the higher
the skew.

Positive autocorrelation,
a slow, “quiet trend”

Medium, stable
correlation with similar
assets.

Nonparallel Accounting

That some assets have owners but no one is consciously short them is one of
the miracles of modern finance. In other words, there is an accounting dis-
crepancy between the owner who is subjected to some rules and the seller
who does not have any marks to market.

Bonds issued by the government of Italy are marked to market by the
owners who show a measure of happiness when the prices go up (i.e., yield
comes down). Paradoxically, the government of Italy too will exhibit happi-
ness when the prices go up because their financing costs in the future will
be lower. While the owner’s wealth increased in the move, that of the Italian
Republic did not decrease. If the government of Italy had a P/L showing
that they issued bonds at 100, now trading at 105, and that they would lose
a few trillion lires should they buy back their issues, the matter would be al-
together different.
[End OCR]*

## Page 264

*[Image OCR]
250 Measuring Option Risks

Stocks are issued by corporations. At the end of a day when the stocks
are up, the wealth of the country increases by the jump in capitalization.
Wealth was created out of nothing. But corporations who issued the stocks
should not begrudge the value created at their expense: The owners can
issue even more stock, manager’s stock options have improved, employees
now are richer in their retirement plan, and so on. This is a miracle where
barring a few exceptions, just about every person in the world benefits from
a rally. The unhappy minority is composed of (1) those who do not own
stocks and see their neighbor driving away in a new car (they could be con-
sidered the true shorts)—only rarely do these become the majority; (2) the
uncovered short, the commodity traders, the hedge funds hoping to make a
killing in the next market crash.

Value Linked to Price

A stock price that goes up often brings stability to the company because it
can now have access to easier financing. When the market value of a com-
pany goes up, the debt-to-equity ratio improves from both an accounting
and economic standpoint.’ From an accounting point of view, the company
can raise cash and retire debt thus increasing the left-hand side of the bal-
ance sheet. From an economic one, the total capitalization increases. Such
increase can lead to a better credit rating and ensure a wider (and cheaper)
ability to borrow. Likewise a drop in price makes it riskier, therefore more
volatile. The same applies to governments. When an emerging country has
higher asset prices, the government can easily fulfill its debt obligations by
issuing debt. But when the asset prices are trashed, governments might find
it increasingly difficult to borrow, thus precipitating a vicious cycle and
creating more uncertainty.

Currencies as Assets

Veteran option traders can easily detect if the currency pair is biased. Typi-
cally, currencies that are “parallel” exhibit a symmetrical behavior between
sell-off and rallies. They are the ones where the trading of the currency pair
is dominated by flows based on the commercial exchange of products. On
the other hand, the currencies that act as investment assets vis-a-vis one an-
other are going to behave in an asymmetrical manner. The German mark
against the Italian lira, the Spanish peseta, or the Greek drachma represents
such behavior.
The following are some rules to remember:

* Generally, currencies that act as assets present a strong correlation
between their price and their interest rates. Interest rates are raised
either to “defend” the currency or because of the capital flights in a
sell-off.
[End OCR]*

## Page 265

*[Image OCR]
Beware the Distribution 251

¢ Currencies that are trade vehicles will have an independence be-
tween their price and interest rate.

Reverse Assets

Some assets like gold, and to a lesser extent the Swiss franc, the German
mark, and the yen, will exhibit the mirror image of others, often when they
behave exactly as the opposite of a regular asset. They become the recipient
of the capital flights away from regular investment currencies and equities.

Volatility Regimes

The high volatility for a short span of time followed by a low volatility for a
long period shows a “fat tailed” histogram of the distribution. Many appli-
cations of the Poisson arrival time have shown consistent results with the
actual histograms, with such methods as the jump diffusion process.

Typically, the histogram will also show a “skew”; the distribution will
be asymmetrical with fat tails on the left and thin tails to the right. As de-
scribed earlier, however, the histogram conceals the sequence of events. The
conditional skew (conditional on the previous regime being Type 2) is very
high.

These facts can appear rather well in a histogram once one isolates the
Type 1 from Type 2 regimes. .

Figure 15.10 shows regime Type 1 including the first move down in the
markets that ensures the passage into a Type 2 market.

For a Type 1 regime plus transition, a trading strategy should be set up
to benefit from the following scenarios:

Rally Sell-Off

Slow Likely Unlikely
Fast Unlikely Likely

Figure 15.10 Regime type 1 including a transitional buffer.
[End OCR]*

## Page 266

*[Image OCR]
252 Measuring Option Risks

Figure 15.11 Type 2 regime (after the crash).

Correlation between Interest Rates and Carry

The major feature of biased assets is the strong correlation between interest
rates and the carry. More of it is explained in Chapter 8, which is the chap-
ter devoted to gamma and shadow gamma.

More ADVANCED PuT-CALL PAriry RULES

In general, for European and soft American options, the delta-neutral call
will replicate the risk profile of a delta-neutral put. For that reason, the
volatility of an in-the-money call needs to be exactly the same as that of
the offsetting out-of-the-money put. Hence when mentioning trades “for a
credit” or for a “debit,” the trader examines the transaction by stripping
every intrinsic portion (adjusted for present value) out of the total costs.

Therefore, the contrast between in-the-mongy and out-of-the-money
becomes nonexistent (except for delta differences) since both have an
equivalent time value and each can replicate the other synthetically
using a deterministic instrument. The difference becomes at the level
of the “moniness”: how far in distance between the strike dnd the
asset price.

Accordingly, this analysis will consider options as far as their “upside
moniness” (how far the strike is above the current asset price), and their
“downside moniness” (how far the strike is below it).

Barrier Products. A short regular knock-out option long the vanilla will
replicate the long knock-in option, both of the same strike and outstrike.

—-KO+V=KI
hence: KI — V = —KO
[End OCR]*

## Page 267

*[Image OCR]
Beware the Distribution 253

For example, a 102 KO call with a 98 outstrike can be replicated by buying a
102 call knock-in at 98 and selling a vanilla 102 call. At the barrier, both op-
tions will be square.

This applies to reverse knock-out options as well.

American Binary Options. These can synthetically be constructed using
KO with rebates corresponding to the size of the bet. Synthetically, a binary
“call” bet (if touched) at 104 with a payoff of $2 is a 104 knock-out call with
both strikes and outstrikes of 104 (the option will never be in the money) and
a rebate of $2. The rebate can be interpreted as an American binary, with
some possible complications attending the time value on the payment date.

European Binary Products. These have the same put/call parity rules as
regular options, but with a twist: A long call is the mirror of a short put, in
vegas and all, unlike a regular option that has a rule of long call + delta =
long put + deltas. This is called the reversibility rule and can be con-
structed as follows:

At the limit, a binary is a spread. Establish first the equivalence of
spreads by taking two theoretical strikes 96 and 94:

Long a 96 P short a 94 P long a net .13 deltas will replicate the short
94 C long a 96 C but long .13 deltas, surprisingly. Unlike a simple option,
the replication of a spread is done through deltas of the same side and
the same magnitude.

By extension, a long binary 95 puts long .13 deltas (in the forward) is
equivalent to a short binary 95 calls long .13 deltas.

Rainbow Options (Dual or Multiple Strike). The put-call parity rules do
not fully hold. For example, a put on IBM or Microsoft cannot be replicated
with a call on IBM or Microsoft and the underlying assets. Traders are just
thankful that the financial markets do not present too many of these. Most
traders hate these instruments because of their deception. ,

Compound Options. Put-call parity rules hold at the second option order
level. A long (European) call option on a call can be replicated with a long
put option on a call and long the call, all of the same strike and outstrike. A
short put option on a put can be replicated by shorting the call on the same
option and shorting the option, and so on. Even third-order options, like a
call on a compound option can be hedged by using a put on the same order
and the option one order below it.

The breakdown of the put-call parity rules for an American option leads
to the notion of omega, the expected life of the option (different from the
nominal duration). It is therefore preferable to consider that American
[End OCR]*

## Page 268

*[Image OCR]
254 Measuring Option Risks

Option Wizard: A Vexed Question

There was a middlebrow swaps head trader at a famous institution who gave
the fear-shaken MBAs he interviewed the following question:

“The Eurodollars are capped at 100.00, and cannot trade any higher. When
the market goes to 100.00 what is the price of the 100 strike calls?”

“Zero,” the interviewee would answer (proudly).
“Right. This is a good answer. How about the put?”

“...” (silence, a lengthy silence).

“Well” the swaps trader would announce, “you failed my test, the put
should trade for zero. Don’t you understand put-call parity rules? Perhaps
you should get a job in manufacturing or in the accounting department of
some corporation.”

One day, the swaps trader met his match in the person of a friend of the
author, a veteran quantitative trader (quite a rare breed). The swaps trader
started his usual bullying, and the veteran answered:

“Your question makes no sense. If your axiom is that the Eurodoilar market
cannot possibly go to 100 (rates go to zero), then there is an incoherence
in the reasoning. The market could approach 0, but the volatility would be
such that it would take eternity to.get there.

“Second, there is a worse conclusion: A market that reaches 100 (assuming
it is bounded there) would be dead and offer no volatility whatsoever.

“There are many explanations for that. An explanation in trader’s terms is
that a market that goes to 100 would be in itself a free option. So traders
can sell at 100.00 knowing that it would be free money: They can partake
of a sell-off without a possible rally! Knowing that if the market has any
volatility in it, traders would sell it at 99.99 thus paying a tick for the per-
petual option, and so forth. The market then would settle at a price where
your call would be worth more than zero.

’
“In mathematical language, a market trading at 100 would be degenerate,
which then would make your question about puts and calls superfluous.”

The quantitative trader did not get the job (to his great luck). One result of
the conversation, however, is that the swaps head trader stopped asking his
question.
[End OCR]*

## Page 269

*[Image OCR]
Beware the Distribution 255

options need to be priced on a term structure of volatility that is shorter
than the nominal European one. At all times, an option on one strike will be
close to the European (the out-of-the-money one with longer expected life)
while the counter option will be priced as an American.

At a more advanced level, one needs to examine the skew effect. Put-call
parity being suspended, the deep in-the-money option would obey the
early exercise rules, which would translate, in addition to the term struc-
ture effect, into a different skew.
[End OCR]*

## Page 270

*[Image OCR]
Chapter 16
Option Trading Concepts

Hiring a trader is like selling volatility. If he does very well or very poorly, you are
out of a job.
An old option proverb

This chapter provides the necessary definitions of option trading concepts
covered elsewhere in the book.

B Officially, option replication corresponds to a self-financing method to
replicate the payoffs of an option instrument with some other instru-
ments (Figure 16.1). In practice, option replication is a broader concept
that covers all operations done through options.

@ Static replication is a risk management strategy that consists in finding
a match for an option position that does not require continuous rebal-
ancing. Static replication aims at both reducing the P/L variance from
the trade and minimizing the transaction costs.

The first and easiest static replication is the put-call-asset arbitrage where a
call is synthetically made into a put and so on.

The reader should be warned against the static replication of instru-
ments that have a stopping time (i.e., an unstable duration) with instruments
that have a constant duration. This topic will be covered in the discussion of
binary and barrier options. ’

Finally, there is the need to distinguish between decomposition and
replication. Many trades require operators to decompose them as a value
discovery method and a hedging orientation. Decomposition can reveal
the skew risks, for example. In most cases, however, the replication will be
impractical.

Static replication via the rolling back of the risks on a binomial tree
(Derman’s method)! consists in taking the final payoff of the security, con-
structing a binomial tree around it, and trying to replicate the payoff at al-
most every node of the tree. Such a method can be helpful in hedging the
risks of some barrier options. We will show an adaptation to such method
where, in place of the payoff, both the payoffs and the Greeks are matched.

256
[End OCR]*

## Page 271

*[Image OCR]
Option Trading Concepts 257

Option Replication

First Order: (Delta, Rho)
rst Order Portfolio Insurance

European
Put-Call
Parity

Second Order
(Gamma,etc.)

Weak Static
Equivalence

Nth Order:
Moment Matching
Aggressive Replication

Replication

of Path- Through Market
Dependent Making
Duplication Options

(Ramp
Replication)

Figure 16.1 Option replication.

The recursive option replication method used by the author with a great
measure of success consists in taking a series of states for the structure and
finding trades that minimize the exposure to the states (see Figure 16.2). A
state is defined as a possible asset price in the future.

Greeks to Match on Every State
Delta.

Modified Gamma.

Modified Vega.

Theta.

Modified Rho (Rho1, Rho2).
Bleed.

Correlation delta (if any).

The best static hedge is the one that matches the Greeks on every state
across all the nodes.” A shortcoming is that often operators need to spend a
[End OCR]*

## Page 272

*[Image OCR]
258 Measuring Option Risks

Final Payoff
94 days 40 days 17 days 7 days Present
State 1
“State 1
State 2 — State t
State 2 “> State 1

State 3 we State 2 — °—————~». Initial State
State 3 State2 J

State 4 —. State 3
Stated

State § —_—
Figure 16.2 Recursive option replication.

great deal of the income from the trade paying bid /offer spreads to match-
ing the Greeks as it may require such a combination of trades in oft-illiquid
markets. The answer is very often dynamic hedging.

BM Dynamic hedging implies sticking to a minimum Greek exposure and
rebalancing continuously to achieve a certain neutrality. It is opposed to
static hedging where the trade is looked on as one combination with
some expiration outcome.

The difference with the operation in Figure 16.2 is that a dynamic hedger
does not look for trades that completely neutralize the tree.

Dynamic hedging concerns all the Greeks in the book. It starts with the
rebalancing of the deltas (as the market moves or as the delta bleeds with
time). As gammas change, it involves the adjustment through options to re-
duce or increase gammas and the consequent time decay. As markets move,
Rhol and Rho2 need to be adjusted and so on.

One matter worthy of mentioning concerning dynamic hedging: It
makes every option become path dependent. °

@ Neutral spreading is an option trading technique that consists in buy-
ing a certain quantity of options against selling another, all of different
strikes and expiration. It is a form of dynamic hedging that allows for
some Greek risk provided there is proper compensation for it.

In floor trading lore, a spreader is someone who matches the “red” tickets
with the “blue” ones (originally on the exchanges red tickets were sales and
blue tickets were buys).

While there are varying degrees of neutral spreading, it is generally
assumed that if someone buys an option and sells another, provided their
strikes and expirations are sufficiently near one another, the ability to
capture “value” is greatly enhanced. The study of biased assets will show
[End OCR]*

## Page 273

*[Image OCR]
Option Trading Concepts 259

that the spreader needs some additional constraints in some markets
where the skew makes “downside” strikes remarkably dissimilar to “up-
side” ones.

Most successful option trading firms were started by spreaders, and, to
this day, the most successful one ever, Swiss Bank Corporation (in its asso-
ciation with O’Connor and Company, originally a successful Chicago Board
Options Exchange [CBOE] market maker firm), requires floor experience
from its traders to hone their spreading skills.

Sophistication in spreading is usually obtained only by dint of experi-
ence. How some instruments marry with each other can certainly be deter-
mined theoretically, but as with most games, only thorough practice would
allow the person to be able to match an option with another.

Classical spreading techniques can be easily extended to the exotic op-
tions arena: Many of the rules are common. However, the distance between
strikes and expirations becomes more difficult to assess.

The value of spreading as a market-making technique is greatly appreci-
ated by the exchanges and the financial backers, since such traders com-
monly have easy access to capital and the exchanges allow them great
leverage.

The advantages of spreading include:

¢ Capturing of Central Limit. The market maker can reduce the vari-
ance (read luck) and maximize the drift (read skills). The notion is
described in Chapter 3.

° Insulation from the Risks of the Formula. Generally, trading options
against options is a serious way to be protected from the imperfec-
tions of Black-Scholes-Merton (or others). Assume that, with the mar-
ket trading at 100, the at-the-money options puts and calls struck at
100 show an implied Black-Scholes-Merton volatility of 15.7%. Re-
gardless of the effect of the formula and the distribution, put-call
parity arbitrage will make puts and calls of the same strike and expi-
ration trade at the same time value (hence the same volatility).

Suppose that the market for the 104 calls shows a volatility of 16.2
and that the market of the 104.5 calls shows a volatility of 15.5. Sell-
ing the 104 calls and buying the 104.5 calls while adjusting the
amounts to satisfy a low residual gamma and delta neutrality seems
to be a good idea. Should one discover (as many.did) that Black-
Scholes-Merton underprices out-of-the-money options, there will be
no material effect on the trade: Both options would be affected. Con-
versely, should one discover that Black-Scholes-Merton overprices
the upside (which would be true in biased assets), the same will be
true. The strikes are sufficiently close together for the spread to be
impervious to the formula.
[End OCR]*

## Page 274

*[Image OCR]
260 Measuring Option Risks

¢ Realization of Theoretical Edge. Trading options against options
proves to be a safe way to hedge against the host of second, third de-
rivatives, and so on that plague the trader:

First Level: Delta neutrality (including the rhol, rho2).
Second Level: Gamma neutrality.
Third Level: Vega neutrality.

@ Theoretical edge is the difference between the trading price of one in-
strument and some defined fair value.

It is safe to see how a spread sometimes confers such a theoretical edge:

1. Insulation from the Broad Set of Parameters. There was a market
maker on the floor of the CBOE who attributes his survival to the
fact that he only traded calls.

2. An Effectual Market-Making Device. Options are generally indi-
vidually illiquid while the general market is very liquid. That is the
reason, as mentioned earlier, that the active market makers can se-
cure easy financing. The more active the trader, the more profitable
he will be at the end if he keeps transacting “above” or “below” the-
oretical value. Some market maker firms impose a volume require-
ment on their traders to force them into activity. By forcing them to
trade actively under severe gamma and theta constraints, they can
ensure that the traders are satisfying some budget requirements.

INITIATION TO VOLATILITY TRADING:
VEGA VERSUS GAMMA

Risk Management Rule: “Long” or “short” volatility carries no

true risk management meaning. One needs to qualify the gammas
(with a range) and the vegas. ’

The expression “long and short volatility” is not an expression that can
be used by a dynamic hedger. In the early 1980s, it was difficult for traders
to explain to the bleary-eyed bosses the difference between gamma and
vega. Typically for one option, the position would be straightforward since
it will be long or short both.

X once worked (briefly) for a boss who attributed his expertise in op-
tion trading to the fact that he had been extremely profitable once (acciden-
tally) on an option position. X was short calendars and was asked whether
X was long or short volatility. X was short volatility but that he wanted the
[End OCR]*

## Page 275

*[Image OCR]
Option Trading Concepts 261

Option Wizard: Basic Forms of Option Strategies

Simple trades with simple products:
Straddles, strangles, butterflies, volatility bets.
Complex trades with simple products:

Long leptokurtosis (fourth moment bet).
Playing the volatility term structure.
Calendar/diagonal spreading.

Long vega convexity.

Long the Eurodollar “dampening” effect.

Distributional arbitrage: skew trading.

Simple trades with complex products:
Playing the variance ratio with barrier options.
Bets and the reflecting barriers.

Complex trades with complex products:

Distributional arbitrage through contingent premium options.

Playing the second order convergence with barriers: arbitraging the
slope of the curve.

Playing the reverse knock-out convexity against ramp options.

Arbitraging higher moments of the distribution with a combination of
bet and compound options.

market to move. It was immediately easy to see the confusion in the boss’s
eyes. It was difficult for X to explain that he was long gamma and short
vega, and after he ran out of breath trying to explain the theory of option
pricing to him, he resolved never to work for a bad listener.

In a two-option book, the following structures are used (Table 16.1):

¢ Option A is a short term at-the-money option.
¢ Option B is a medium term at-the-money option.

Risk Management Rule: ‘Long or short gamma needs to be quali-
fied with a range (as shown in Chapter 8). A trader needs to explain

if he is continuously long gamma, long up-gamma, down-gamma,
or if the gamma flips somewhere.
[End OCR]*

## Page 276

*[Image OCR]
262 Measuring Option Risks

A calendar spread has a gamma that reverses away from the money. The rea-
son is that a calendar seller can only hope to make a limited sum of money
from the calendar (generally the credit for an at-the-money calendar). An op-
erator willing to bet against the wings has several options: Either play the
out-of-the-money options against the at-the-money, or, more easily play the
calendar.

By applying a ratio to the time spread, the operator can achieve a
gamma that would remain neutral to mildly short at-the-money but that be-
comes rapidly very positive in the wings.

SOFT VERSUS HARD DELTAS

M@ A soft delta execution is a delta hedge done through an option, that, ac-
cordingly, vanishes asymptotically (to the asset price). A hard delta
does not vanish at the limit of the move in the underlying asset.

Soft deltas are generally used to cover secondary deltas (i.e., deltas arising
from market moves) that cannot be hedged with a hard delta without in-
creasing the risks in the extremes.

Example: The trader has the follewing position on the books: long at-
the-money option and short larger quantities of out-of-the-money calls:

Asset 100 102 104 106 108 110 12 114
Delta 0 10 15 17 15 12 —10 —50
Gamma Long Long Long Square Short Short Very Short Very Short

At 106, the trader has some profits stemming from the long deltas. Sell-
ing hard deltas would cause his negative gamma to turn worse in the
tails:

Asset 100 102 104 106 108 110 112 114
Delta —17 —7 —2 0 —2 —5 —27 -67

By selling deltas through options, with the purchase of a put for exam-
ple, or a put spread, the following would be achieved:

Asset 100 102 104 106 108 110 112 114
Delta —2 -1 -8 0 2 0 —10 -50

Sometimes the best way to sell deltas is by shifting strikes (i.e., moving the
center of the gamma closer to the short zone). At the beginning, the clumps
[End OCR]*

## Page 277

*[Image OCR]
Option Trading Concepts 263

of long gamma were centered around 100 and the short gamma around 112.
The rally should give the trader the opportunity to shift the center from 100
to 106, thus stabilizing the position. The best possible hedge in this case is a
soft delta (buying options struck in the 106 zone) added to a selling options
struck in the 100 zone. The next step, should the rally continue, would be to
move the longs into the 112 zone by repeating the operation.

Recommendation. In the rebalancing of an option position, the trader
needs to perform the following test: Does the gamma flip the other way? If
the long gamma flips into a negative gamma, it will be very dangerous to
sell “hard” deltas (i.e., cash or futures) because the P/L would then turn
out to be worse in the tails of the matrix.

VOLATILITY BETTING

These trades can involve more than one option. All options considered in
this section are vanilla.

M A first-order volatility trade using vanilla options is a trade that satis-
fies the following conditions:

¢ It is monotonically long or short volatility though without necessar-
ily having a constant gamma.

¢ Its vegas and gammas are on the same side of the market.
* Itis intended to be delta neutral.

These characteristics apply to simple option trades: long straddles, long
strangles, and so on. Any form of trading where all options are long or short
would satisfy such rules (Table 16.1).

First-order volatility trades present the advantages of being tractable,
with a P/L that is easy to forecast (Figure 16.3). They represent positions
with a smaller degree of complexity.

Table 16.1 The Volatility Rules

Long Vega Short Vega
Long Gamma Long A, Long B Long A, Short B
Long A,
Long B
Short Gamma Short A, Long B Short A, Short B
Short A,
Short B

Weightings are not used in this introductory framework.
Option B has longer time to expiration than option A.
[End OCR]*

## Page 278

*[Image OCR]
264 Measuring Option Risks

P/L P/L
Asset | — Asset
Price Price

[ Symmetrical | Assymmetrical

Figure 16.3 First-order volatility trade.

@ A second-order volatility trade is an option trade that includes long and
short options of different strikes or different maturities but within the
confines of one product. It is characterized by the following:

e The gamma always flip from positive to negative on the map.

¢ Its raw vega can change sign somewhere along the lines (this is not
an essential condition).

¢ Its weighted vega always reverses somewhere on the map.

* Owing to its complexity, graphical representation in two dimensions
often becomes irrelevant and of weak revealing powers.

A call spread (one by one), such as buying the 102 calls selling the 106
calls in equal amounts, will exhibit the characteristics of a single call when
the market is at 100. Should the market rally to 104, however, it will then
have the zero-gamma, positive down-gamma, negative up-gamma that char-
acterizes a risk reversal.

A ratio spread where the operator buys the 102 calls and sells twice as
much of the 106 calls will present opposite characteristics.

Higher Moment Bets

Mm A third moment bet is a form of distribution arbitrage where,a bet is
made on the correlation between the volatility of a particular market
and the asset price.

@ A fourth moment bet is long or short the volatility of volatility. It could
be achieved either with out-of-the-money options or with calendars.

Example: A ratio “backspread” or reverse spread is a method that in-
cludes the buying of out-of-the-money options in large amounts and the
selling of smaller amounts of at-the-money but making sure the trade
satisfies the “credit” rule (i.e., the trade initially generates a positive
cash flow). The credit rule is more difficult to interpret when one uses
[End OCR]*

## Page 279

*[Image OCR]
Option Trading Concepts 265

in-the-money options. In that case, one should deduct the present value
of the intrinsic part of every option using the put-call parity rule to
equate them with out-of-the-money.

The trade shown in Figure 16.4 was accomplished with the purchase of both
out-of-the-money puts and out-of-the-money calls and the selling of smaller
amounts of at-the-money straddles of the same maturity.

Figure 16.5 shows the second method, which entails the buying of 60-
day options in some amount and selling 20-day options on 80% of the
amount,

Both trades show the position benefiting from the fat tails and the high
peaks. Both trades, however, will have different vega sensitivities, but close
to flat modified vega.

CASE STuDy: PATH DEPENDENCE OF A
REGULAR OPTION

The following illustrates the degree to which option replication could be
path dependent.

This case study uses a string of returns‘, including a beginning price
and the end price, and its volatility. With these numbers, it is possible to
reshuffle a high number of sequences while the beginning price, the end
price, and the volatility remain the same.

Ratio Backspread

60 +

40

20 7 s

PA
°

Today+3 days

Asset Price

Figure 16.4 Trading the wings, first method.
[End OCR]*

## Page 280

*[Image OCR]
266 Measuring Option Risks

Calendar Spread

dday+3 days

PR

Asset Price

Figure 16.5 Trading the wings, second method.

To show how to create the price movement, a sequence of percentage re-
turns is taken, say, 1%, 5%, —1%, —5%. A security starts at 100. The next
price is 100 (1 + .01) = 101. The next price is 101 x (1 + .05) = 1.0605. The
next price is 106.05 X (1 — .01) = 1.0499. The next price is 1.0499 x (1 —..05)
= 99.74 (the final price). The price sequence is then 100, 101, 1.0605, 1.0499,
99.74. Mix the preceding returns, and get 5%, —5%, —1%, 1%. The price se-
quence will therefore be 100, 105, 99.75, 98.75, 99.74, quite a different path.

In the following example, the goal is to replicate 252 trading days by
taking an arbitrary distribution with a 15.6% volatility (average movement
1% per business day) and reshuffle it into 106 different paths.

All the prices start at 100 and end at 98.6. Figure 16.6 shows eight sam-
ple paths and illustrates the discrepancy in the paths that can result: One
path shows a high of 127, the other shows a high of 105. It is again counter-
intuitive that both have the same mean and the same volatility. Indeed,
some of the shuffled results have a yearly range of 10% while others have a
yearly range of 60%.

Figure 16.6 shows only eight of the paths used in the analysis for the
sake of clarity.

The trader buys $10 million of a one-year European call option struck at
100 and rebalances the delta daily. European options are said to be path in-
dependent, therefore only the final outcome matters for them. Table 16.2
shows the profit and loss from a strategy that would buy volatility at the
exact 15.6% and hedge, without transaction costs, at the close of every day.

This author once, for his amusement, posed the problem to three cate-
gories of people: the junior trader, the experienced trader, and the trading
[End OCR]*

## Page 281

*[Image OCR]
Option Trading Concepts 267

130
125 +
120 ——-Path1 .
9 M7 A —— Path2
2 110 + Le
a. 105 | | Path3
a 100 i : Path4
< 95 . ——— Path5
90 —— Path6
85
SO -brma ache aoe en acnsinensiin Path7
1 25 49 73 97 121 145 169 193 217 241 —— Paths

Business Days

Figure 16.6 Eight paths of the same volatility and same final return.

manager. Most junior traders and trading managers rehashed what they
knew from the training packages that volatility was volatility and the order
of events did not matter. Experienced traders explained that we were deal-
ing with a discrete, not continuous time world and that there would be a
variance that would be a function of transaction frequency. The interesting
result is that all of the experienced traders got the answer right and all the
trading managers got it wrong. The financial community is still largely un-
aware of the poor tracking of the risks through dynamic hedging.

Most experienced traders explained that it was preferable, when long
gamma, to have the large moves when gamma is at the maximum and the
small moves when the market is furthest away from the strike.

Table 16.2 Results of the Dynamic Hedge .

P/L Frequency P/L Frequency P/L Frequency
~72000 1 ~ 20000 1 32000 6
~ 68000 0 — 16000 12 36000 4
~64000 1 — 12000 13 40000 4
~ 60000 0 —8000 4 44000 4
—56000 0 —4000 2 48000 3
-—52000 0 0 6 52000 1
~—48000 4) 4000 5 56000 1
—44000 1 8000 7 60000 0
—40000 0 12000 2 64000 0
—36000 1 16000 6 68000 2
—32000 1 20000 5 72000 + 1
~—28000 2 24000 7 Total 106
—24000 1 28000 2
[End OCR]*

## Page 282

*[Image OCR]
268 Measuring Option Risks

Table 16.3 Extreme Path Dependence: A Risk Reversal

P/L Frequency P/L Frequency P/L Frequency
— 600,000 2 — 150,000 8 300,000 7

~— 550,000 0 — 100,000 5 350,000 7

— 500,000 0 —50,000 11 400,000 0

— 450,000 1 0 8 450,000 0
—400,000 2 50,000 9 500,000 1
—350,000 0 100,000 6 550,000 1
—300,000 4 150,000 9 600,000 0

~ 250,000 4 200,000 7 650,000 2

— 200,000 3 250,000 9

The results show how path dependent the world can be for a dynamic
hedger. The preceding example shows a simple position that is monotoni-
cally long gamma. The next position is initially close to flat gamma but pre-
sents mixed features.

Table 16.3 represents the following risk reversal: Long $100 million of
the 90 call, short $100 million of the 110 call and delta neutral. The delta, as
before, is rebalanced daily. The results are quite unsettling.

The distribution that caused such P/L had no skew; generally there was
no difference between down-volatility and up-volatility. There was also no
correlation between the volatility ofthe market and the level of the underly-
ing asset. In short, a clean distribution, extremely theoretical.

This suggests the following trade-off: The only way to reduce path de-
pendence is to increase the transaction frequency. Such increase in the
transaction frequency will raise the costs of trading. The impact can be fur-
ther compounded because a negative gamma trader incurs higher such
costs than a positive gamma counterpart.

Many heads of trading rooms, underestimating the impact of the path,
will issue comments such as, “On balance, in the long run, I will eventually
capture the ‘edge’ in the market.” Most of them are unaware of the follow-
ing rule: In options, variance of P/L is usually underestimated. One needs
more diversification. One of the rules of diversification, which is diversifi-
cation through time, does not properly work owing to the leverage and the
fact that traders are continuously monitored by a nonstatistical person
breathing down their neck. Moreover, the life of a trader is too short for ad-
equate time diversification.

In the real market, the results will certainly be worse, for the following
reasons:

e The trader usually will hedge himself, when short gamma, with stop
losses, which increase the costs and causes a higher incidence of
whipsaws.
[End OCR]*

## Page 283

*[Image OCR]
Option Trading Concepts 269

Most markets exhibit some form of a skew.

Traders have some form of absorbing barrier in their P/L. Many of
the 106 runs had losses in excess of the final worst-case apparent re-
sult of $600,000. If we stopped a trader at $300,000, we would have
many more negative runs.

Perhaps the element that would make the preceding spread widest is
implied volatility. It is assumed that the trade is only adjusted with
the asset, not with options. Seeing the P/L of every path and analyz-
ing the P/L volatility would show that some of them are scary.

Table 16.4 The Nasty Path

Days to Asset Option Option P/L P/L Cum P/L
Expiration Price Price Delta Option Future ($000)
30 100.00 0.02 26
29 100.13 0.02 28 14 —3 —2.0
28 100.26 0.02 30 ~ 15 —4 —4,1
27 100.39 0.02 32 1.6 —-4 —6.4
26 100.50 0.03 34 1.0 —4 -9.0
25 100.65 0.03 38 2.5 -5 11.6
24 100.78 0.03 41 1.9 -5 -14.6
23 100.91 0.03 44 2.1 -5 —17.8
22 101.04 0.03 47 - 2.2 -6 —21.3
21 101.17 0.04 51 2.3 —6 —25.1
20 101.30 0.04 55 2.5 -7 —29.3
19 101.43 0.04 60 2.6 -7 —33.8
18 101.56 0.04 64 2.8 —8 ~38.8
17 101.69 0.05 70 2.9 —-8 —44.3
16 101.82 0.05 75 3.1 ~9 —50.2
15 101.95 0.05 81 3.2: -10 —56.7
14 102.08 0.06 88 3.4 -11 —63.9
13 102.21 0.06 96 3.5 —11 —71.9
12 102.34 0.06 104 3.7 -12 —80.6
11 102.47 0.07 112 3.8 -13 +90.3
10 102.60 0.07 122 3.8 -15 ~101.1
9 102.73 0.08 133 3.8 -—16 —113.1
8 102.86 0.08 145 3.7 -17 —126.7
7 102.99 0.08 158 3.5 —19 —142.0
6 103.12 0.09 173 3.1 —21 —159.5
5 103.25 0.09 190 2.3 —22 179.7
4 103.38 0.09 208 0.9 —25 — 203.5
3 103.51 0.09 229 —1.6 —27 —232.1
2 103.64 0.08 253 —6.3 —30 ~ 268.2
1 103.77 0.06 273 17.2 —33 —318.2
0 103.98 0.00 0 —63.8 —57 —439.4
[End OCR]*

## Page 284

*[Image OCR]
270 Measuring Option Risks

The preceding position becomes long or short vega at some level, and
adjustments would worsen the final spread.

The following basic example was designed by the author in his forma-
tive years to satisfy a bet with one of his colleagues.

SIMPLE CASE STupDy: THE “WorRSsT CASE” SCENARIO

Inexperienced risk managers typically describe an option risk as being the
premium when one is long and conversely when one is short. This does not
apply for a dynamic hedger. This case study, with a simple example will
show how a premium seller can earn more than the initial premium col-
lected and a premium buyer can lose more than the expenditure.

An out-of-the-money call has a 20% delta.

"Asset price: 100.5
Strike price: 104.
Days to expiration: 30.
Price: .19.
Amount: $10,000,000.
Initial premium: $190,000.

It would appear that a buyer of the call would only lose the premium,
$190,000, not more.

It is apparently a very safe trade. Look at the worst-case scenario for a
delta hedger: Initial delta: $280,000. The hedger sells the entire amount.

Look at the P/L table over the next 30 days: The path followed by the un-
derlying is particularly vicious as it rallies daily (see Table 16.4).

It is easy to see that dynamic hedging got our man in trouble, causing
losses of $440,000 where they should have been limited to $190,000. Such
events are rather common (though rarely to such an extreme) with trends
where out-of-the-money options on the side of the trend decay mercilessly
and delta hedgers end up losing on both the option and the delta hedge.
[End OCR]*

## Page 285

*[Image OCR]
parT III

TRADING AND HEDGING
EXOTIC OPTIONS

In Part III of this book, the focus is on the risks of the major exotic options
from the standpoint of the dynamic hedger (and the informed customer). In
place of an anecdotal enumeration of all the possible instruments, the discus-
sion will concentrate on the hedging techniques through decomposition of
the blocks of risks.

Risk managers who never traded before will learn that:

¢ The risk of any soft path-dependent option can only be understood
through thorough grounding in American binary options and stop-
ping times. Typically, they commit thé tragic mistake of loading their
minds with peripheral notions, most of which are related to pricing.

e¢ Understanding multi-asset options needs to come from the notion of
matrix analysis and cross-gammas by intuitively looking at the covari-
ance matrix as one volatility. In addition, every structure can become
multi-assets through correlation hedging.

¢ Understanding non-time-homogeneous risks needs to come from the
grounding in calendar spreading and the learning about the notions
of forward volatility, not through the dissection of deferred options
and other minor varieties. ’

In this text, Asian options and lookbacks are subjects of minor impor-
tance. Asian options are simply an application of a basket methodology (the
basket is in time), and lookbacks can be thought of as a simple footnote on
barrier options.

Finally it is necessary to stress that training in exotic options is often
better performed with a good basic and structured education in vanilla
puts and calls. Conventional option trading provides a good training in the
failings of the distributions. This refers to what has been described as com-
plex trading with simple instruments as opposed to simple trading with
complex instruments.

271
[End OCR]*

## Page 286

*[Image OCR]
272 Trading and Hedging Exotic Options

| Road Map :

Understand discontinuous

payoff, “pin” risks

e Understand third moment
risks

¢ Learn to trade simple

options

Understand Dirac Delta

European Binary Options
(mandatory)

Understand stopping time
¢ Understand fourth mo-
ment risks

¢ Understand vega convexity:
¢ Understand lookback op-
tions and other structures

American Binary Option
Barrier Options
(mandatory)

A general framework for a
multitude of structures:
shout, extendible options,
etc.

¢ Learn more about Ameri-
can options

Higher Order Options

Learn to trade and gener-
alize any structure in
higher dimension

¢ Learn about basket op-
tions, hence Asians

Multiasset Options
(mandatory)

i:

Asian and Lookback Options
(optional)
[End OCR]*

## Page 287

*[Image OCR]
Chapter 17

Binary Options: European Style

Options that are trivial to price (like binary options) are difficult to hedge. Op-
tions that are difficult to price (like Asian options) are trivial to hedge.

Howard Savery, Exotic Options Trader

The next two chapters are devoted to the trading issues related to bet op-
tions. The reader needs to complete these chapters before attempting to un-
derstand barrier options. To make the progression easier, the discussion
will start with European binary options, as the concept of the “pin” needs
to be well understood before proceeding with American binary options,
which add the complexity of an unknown and highly unstable “duration,”
or stopping time.

The concept of the pin is so central to understanding both option theory
and practice that this author started originally writing the book with that
topic in mind and had to compose all the previous chapters to provide the
reader with the tools to master the risks of the barrier structure. For a hard-
working person, however, a thorough knowledge of barriers and digitals
provides appropriate training in handling the risks of what exotic traders
call the “junior stuff.”

Binary options are perhaps the best training ground for a trader as they
can teach more about advanced book management. At the heart of most ex-
otic structure and every bet resides a binary. It also provides an advanced
test of the skills of any risk manager: Many experienced risk managers who
fail the binary exam should perhaps consider some serious retraining.

Barrier options are also difficult to comprehend for risk managers with a
superficial knowledge of options. Ironically, traders with shallow experience
commonly misunderstand these options even though they dabble in them;
whereas the often-more-qualified trainees master the barrier options. They
invariably allowed the author to distinguish between the true experienced
tisk manager and those who needed further training.

EUROPEAN BINARY OPTIONS

These options are also called digital derivatives, bet options, and gap puts
and calls.

273
[End OCR]*

## Page 288

*[Image OCR]
274 Trading and Hedging Exotic Options

@ A binary option pays a single sum upon the satisfaction of the obliga-
tion. It is called binary because, like 0 or 1, it either pays the full amount
or zero, with nothing in between.

This discontinuous payoff makes binary options particularly hard to hedge
using regular structures that present a continuous payoff.

The opposite of a binary or gap payoff is a ramp, or conventional option.
Binary options can be either American or European.

@ A European binary option is a bet on the asset being higher or lower
than a certain level at expiration. American bets are “if touched” types
that will terminate the bet upon an event at any moment between incep-
tion and expiration and are thus more difficult to hedge.

Some variations that can be both American and European will be examined.
European binaries can thus have only one strike, unlike American ones that
can be of the either-or type. There are no double bets per se, as a European
double bet can be constructed as the sum of two independent bets.

Binary options are present in many structures, as a European or Amer-
ican type. Contingent premium options are a simple construction of Euro-
pean vanilla options plus a bet that pays the initial premium if the option is
in the money.

Common trading beliefs hold that a European bet is easy to hedge ex-
cept close to expiration, where the delta becomes explosive as spot enters
the neighborhood of the strike. It will be shown that it is difficult to hedge
in both circumstances. As a matter of fact, much money was unwittingly
lost in the daily management of a digital with a long time to its expiration.

One reason traders get an accelerated training with bet options is that
their delta imitate the gamma of a regular option.

Often European bet options exhibit parameter linearity problems that
make hedging their vega particularly difficult for a dynamic hedger.

Figure 17.1 shows that a binary option is short gamma and earning theta
(but moderately) when it is in the money and long gamma and paying time
decay when it is out of the money. Exactly at the money, it loses all its option-
_ality and acts like a future. This risk-reversal feature for a barrier appears to
facilitate their hedging with a few skewed instruments. It will be shown that
the appropriate hedge for a barrier is a narrow spread, done in such a size as
to compound the skew effect. Most of all, barriers need to be priced with the
skew structure of the market in mind. In a commodity that becomes volatile
after a sell-off and loses steam after a rally, such a structure can indeed be
very favorable unless the price of the barrier includes a compensation for such
imbalance (with the volatility differential we call the skew).

Look carefully in the microscope at the part that is long gamma (the left
side of the figure).
[End OCR]*

## Page 289

*[Image OCR]
Binary Options: European Style 275

100
80

Short Term

60 Long Term
40
20

0

o wt
no D>

104
109

_
rc

Figure 17.1 Mixed convexity of a European binary.

What causes the long gamma is the local convexity. A binary will pre-
sent a positive leverage when it is out of the money (the trader is risking less
to make more) and negative leverage otherwise (the trader is risking more
to make less). If the bet pays one dollar at expiration and is currently worth
one penny, it will present very large convexity features. It can thus pick up
99 cents and lose 1 only. This would look like a long gamma structure.

Figure 17.2 shows no true difficulty in protecting the vega and gamma
of the binary using ordinary vanilla options. As expiration nears, however,
the profile would start sliding toward the pin exposure. Figure 17.3 shows
the price sensitivity nearing expiration. The graph will become more and
more vertical for a narrow segment and flat otherwise.

Hedging with a Vanilla

Look at the shape of Figure 17.3 and try to find the profile of a structure
that would imitate such a payoff. The first option to come to mind is the call
spread.

Option Wizard: Risk Reversal (revisited)

A risk reversal, as defined in this book, corresponds to situations whtre the
gamma and/or vegas flip from positive to negative across one point.

Initially, a risk reversal corresponded to fences (out-of-the-money cail pur-
chase financed with the sale of a put) but rapidly option traders started refer-
ring to it as a description of asymmetric risk in an option position. It became
the third moment.

A risk reversal for a book manager is the switch in risk across one point.
The best description of a risk reversal to date*: “It is like buying drought insur-
ance and financing it with flood insurance.”

*Lenny Dendunnen.
[End OCR]*

## Page 290

*[Image OCR]
276 Trading and Hedging Exotic Options

{  Qubotahe-money binary

Volatility
12 #15 18 241 24 27 30

Figure 17.2 Price of a 30-day, 15 percent out-of-the-money option and its sensi-
tivity to volatility.

Figure 17.4 presents a call spread where the operator is long the 100.00
strike and short the 100.01.

The difference between Figure 17.3 and Figure 17.4 lies in the scale. The
shapes are very similar. This shows that the binary is really a call spread in
disguise. It presents all such features: long gamma where the proximity of
the long leg makes it dominate (below the long) and short gamma where the
proximity of the short leg makes it dominate (above the short). So the mi-
croscopic call spread is long gamma below 100 and short gamma above. Like
a conventional option, the gamma weakens away from the strike, and the
combination peaks at a certain spot. In between, it looks like a risk reversal.

Option
Price

100

Shorter
Maturity

Longer
Maturity

Asset

o t o Price
oo fo) o

-
”

104
109

Figure 17.3 Binary option nearing expiration.
[End OCR]*

## Page 291

*[Image OCR]
Binary Options: European Style 277

Shorter
Maturity | Longer

Maturity

Price

°
9
94.5
96.5
98.5

400.5

102.5

104.5

106.5

108.5

Cash Price

Figure 17.4 A cali spread nearing expiration.

One additional piece of information, critical to understand the similar-
ity between the two positions: They are both European, path-independent,
and share a common expiration date. As the difference resides merely in the
scale, the increase in the size of one should perfectly replicate the other.

The next step is to compute the face value of the call spread that is fit to
hedge the barrier.

Assume the payoff of the barrier is $100 per unit if the bet is satisfied
and none if the bet is not satisfied. So the amount of the call spread should
be such that one tick represents $100. So the trader would need the miserly
amount of 10,000 units of the call spread. Such replication, however, only be-
comes effective everywhere on the map of risk-neutral possible outcomes if there
is no zone in which it fails to track the other leg. In this example, the trader
selected a market that moved in indivisible “ticks” of .01. Should the market
trade in discrete increments of .05, his task would be made easier as the
replicating amounts could thus be smaller (2,000 units).

Option Wizard: Dividing a Tick

The exchanges make it illegal to shave a “tick” into smaller increments.

Traders could otherwise trade in smaller units by dividing the trade. Imag-
ine that the market moved in minimum increments of .01..Buying 50 lots at
100.00 and 50 lots at 100.01 from the same person would accomplish a
“ginzy,” as the average would be 100.005.

The exchanges did not make buying at two different prices illegal, they
made the trading with the intention of breaking up the tick into smaller prices
an illegal operation. What is illegal is telling a trader: “I will buy some from you
at 100.01 if you also will sell me some at 100.00.”
[End OCR]*

## Page 292

*[Image OCR]
278 ‘Trading and Hedging Exotic Options

In continuous time finance, the replicating trade would need to be done
in infinitely large amounts for an infinitely small difference between |
strikes to properly hedge the book. But markets do not trade continuously,
be it only for the sake of convention, as there is a set minimum price incre-
ment called the “tick.” Such an increment is formal on exchanges as trading
in smaller fractions is disallowed (see Option Wizard, Dividing a Tick). It is
the product of etiquette and convention in over-the-counter markets, as
most screens and systems are not set up to handle any number smaller than
a set fraction. Traders in some markets consider it downright impolite to
display a price in too precise an increment (more than 1/100).

Definition of the Bet: Forward and Spot Bets

It is important to see if the bet contract defines satisfaction as “higher” or
“at least” the bet strike price, which, in mathematical symbol, is > or =. The
difference is minor but affects our replication. If the bet terms specify “I
pay if the asset is 100 or higher” then the call spread needs to be defined as
long 99.99 short 100. If the bet says “higher than 100” then the replication
would necessitate the long 100 short 100.01.

Another important feature concerns the timing of payment. Some bets,
particularly when they are out of the money, are sold as options where the
“buyer” pays some premium immediately and gets some payment if the
asset satisfies the bet at expiration~Others are denominated as a bet where
the party that loses pays the money. The only difference between the two
lies in the discounting of the premium.

mM A spot bet is where one of the parties outlays the premium initially and
would receive payment if the conditions of the bet are satisfied.

A forward bet is one where the parties agree to exchange payments at the
end. Thus, the party paying the initial premium would be considered the
“buyer,” arbitrarily. The following are examples:
.
¢ A bet where one of the parties pays $50 if the spot is higher than 100
on the Friday after Thanksgiving and would receive $50 is a forward
bet.

¢ A bet on the outcome of a tennis game is typically a forward bet.

¢ Anote that may pay 6% interest if Mexico stays above a range and 5%
otherwise is a spot bet. The owner of the note, through the reduced
coupon, is paying in advance for the bet.

¢ A contract where one pays 10 cents and may receive $1 if the fed
eases on the next meeting is a spot bet.
[End OCR]*

## Page 293

*[Image OCR]
Binary Options: European Style 279

So there is no trick worthy of analysis other than:

e The forward bet is forward valued.

¢ The spot bet has one leg present valued (the initial payment) and one
leg forward valued (the payment received).

This book tries to minimize dwelling on simple interest rate arithmetic. So
the reader should assume that interest rates are zero (for a while) and get on
with bigger and better things.

By convention, a call bet is the term used for the agreement where a
party is betting that the price will be higher than the strike on expiration
date. A call bet for one party is a put bet for another.

By convention also, the bet is quoted in percentage of the total at stake.
If the final payment is $1 then the bet priced at 25 cents corresponds to a
spot or forward expenditure of 25 cents. If the bet payment is $500, then it
becomes $125.

PRICING WITH THE SKEW

Assuming replication is important to probe the structure of an option and
get a feel for its theoretical value. Replicating, however, is rather impractical
if one takes into account the slightest amount of transaction costs. This au-
thor has never seen anyone replicate a binary option with narrow spreads.

As the skew exists in most markets and since the binary is sensitive to
the skew, proper allowance should be made in the pricing. If the minimum
tick size is .01, the option will be sensitive to the difference in skew be-
tween the 100 and the 100.01. This is not a joke: If there is a microscopic
skew volatility difference between the strikes owing to the large amounts
involved on the spread, such skew will bear a meaningful effect on its theo-
retical value.

The way to price the skew effect is to look at the wider replicating portfo-
lio and the amounts involved. Assume that there is a difference of .5 volatility
points between the 99.5 and the 100.5 strikes of the same maturity as the bi-
nary. The skew effect on the structure should be equal to the dollar effect of
the skew on the replicating portfolio. It is computed as follows:

Assume that the trader has a 3-month bet paying $100 if the market is
higher than 100 (for simplicity, assume that interest rates are nil and that
volatility is 15.7). The Black-Scholes-Merton value is 49.6. The 99.50/100.50
replicating spread will be approximately 100 times the bet if, in this make-
believe schematic world, the market moves in very large increments of 1
point. This assumption is necessary to eliminate the possibility of the mar-
ket settling between 99.50 and 100.50, in which case the replicating portfolio
[End OCR]*

## Page 294

*[Image OCR]
280 Trading and Hedging Exotic Options

would not match the binary option. Taking the Black-Scholes-Merton values
of the call spread, the value of the bet should be close to 49.6 (this discus-
sion will not get into the notion of the small risk-neutral difference between
the values of both instruments). So in this world of only one-dollar moves,
the trader appears to be hedged.

But there is the skew in the market, as can be seen in an exaggerated
example:

Assume that the call spread trades for .69 because the 99.5 call trades
one volatility point higher than the 100.50. This means that the binary op-
tion needs to be more expensive, and must trade at $69 (as opposed to $49.6
using Black-Scholes-Merton) to make up for the skew. Would anyone bet so
much that the spot would be higher on expiration? The numbers show how
exaggerated the skew effect can be on the barrier.

Now the conflict: A bet trades at the probability of the asset price end-
ing up higher than the bet level. The difference between the delta and the
bet will be explained later; it is important to note that the skew—any
skew—makes the two elements diverge. That divergence will lead to the
more complicated notion of the skew paradox.

In practice, however, the skew would not be as pronounced as in the pre-
ceding example except perhaps in biased assets or SP100 options after an
anxiety attack. The normal skew ranges (in well-behaved assets) between .5
and 3 volatility points for the difference between a .25 delta call and .25
delta put (i.e., between a 75 delta call and a 25 delta call). Interpolating a 2-
point skew, a 3-month option shows about .12 skew at 15.7 volatility for the
difference between 99.5 and 100.5 (since the difference is .06 deltas and .5
deltas command a 2-point skew). The price effect of the .12 skew is .025 on

Option Wizard: The Replication Map

Two path-independent trades should be equal in price if they present the same
payoff everywhere on the map of possible prices (by the rule of stochatic dom-
inance). ,

Path-independent means that they could be treated by looking at expira-
tion only. That statement was qualified earlier with the explanation that path-
independent options become path dependent when hedged dynamically. The
present case involves static, not dynamic, hedges.

Replicating a binary with a call spread makes it necessary to look at limit-
ing cases. The call spread will work outside the strikes, but not inside. The nar-
rower the call spread, therefore, the most perfect the replication.

The call spread, as a limit decomposition of the binary, will always be
higher in price, by a small amount that vanishes as the increments narrow.
[End OCR]*

## Page 295

*[Image OCR]
Binary Options: European Style 281

the call spread, making the difference between strikes .524 instead of .49. So
through the replicating portfolio, the price of the bet is found to be worth
approximately $52.4 instead of $49.6. This is still meaningful enough to
alert traders about potential mispricing.

Verification: Widening the increments to cover a spread that would be
50 deltas apart (i.e., the one quoted by equivalence with the 25 delta risk
reversal) would reveal the following strikes: 95.70 and 105.4 that is 9.7
points apart. It would then be necessary to execute 9.52 spreads on the
hedge to make $100 if the bet is satisfied (if the market settled above
105.4) and lose the premium paid for the spread if the market settles
below 95.70. Priced on the skew, the value of the vanilla call spread
would be $52.30, whereas the same one, Black-Scholes-Merton-priced
(using the same volatility on both legs), would yield $49, similar to the
fair value of the bet without the skew.

A Formal Pricing on the Skew

Intuitively, the price of the binary can be defined as the risk-neutral proba-
bility (i.e., taking the mean of the asset returns out of the equation thanks
to delta neutrality) of ending up in the money. If one takes the earlier argu-
ment of the perfect replication through the call spread, the following equa-
tion results, using the call price C, the strike price of the binary K, and h the
difference between the strikes of the replicating call spread:

C(K + h) — C(K)
h

Thus it resembles the derivative of the call with respect to the strike."

Furthermore, if the trader determined that the volatility of the call was
- afunction of the h, he would therefore create a skew function and establish
a pricing mechanism.

Taking the skew as a function of K, that is o(K), and the skew slope at
point K the derivative of the volatility with respect to the strike as shown in
Figure 17.5. The bet value becomes 8C/8K + 8C/8o0 X 80/8K = Bet (no
skew) + vega of the vanilla call of the same strike X “skew slope.””

Binary call = lim when h 0

Example: Spot trades at 100. Zero interest rates are assumed, for peda-
gogical reasons. Consider the 3-month 100 bet call. The skew for a
3-month maturity increases by .5 volatility between 99.50 and 100.50;
the skew slope would then be .5/1 = .5. The vega of the at-the-money
call for the 3-month maturity is .190 per vol point. The value of the bet is
therefore .496 (the Black-Scholes-Merton value at 15.7% volatility) + .19
x 5 = .591 for one unit paid at expiration in case the bet is satisfied.
[End OCR]*

## Page 296

*[Image OCR]
282 Trading and Hedging Exotic Options

Volatility

Derivative dv/ds at point 100

96 98 100 104 Strike

Figure 17.5 The skew slope for a given expiration.

The Skew Paradox

Quiz: The reader is told that a market went up on any day only by steps of
$1 and went down by steps of $9, with no other possible price change. Such
process is shown in Figure 17.6. Out of 10 steps, 9 will be up and 1 down.
What would be the reader pay for a one-day at-the-money call bet paying $1
if the market closes tomorrow higher than the present level?

The answer is $.90 The key is that the bet does not depend on the ex-
pected return but on the expected number of times the spot would be
higher. That the market drops in large amounts is irrelevant. The payoff of a
bet option is the same whether the market drops by $1 or by $50.

Either by put-call parity rules or by use of the same argument, one ar-
rives at a value of .1 for the put. This is an intuitive explanation for the skew.
It also shows the difference between the bet and the delta: The delta, unlike
the bet, is concerned with the magnitude of the moves since the trader needs
to be protected for such eventuality.

Graphically, it could be represented as shown in Figure 17.7. The area A
needs to be equal to area B. Financial markets impose the constraint on
every security that the left integral be equal to the right one plus the risk-
neutral drift, which results in a mean of the risk-neutral drift, m.

Price Change Price Probability

_— 1 101 90%
100
98 10%
Expected Payoff 0

Figure 17.6
[End OCR]*

## Page 297

*[Image OCR]
Binary Options: European Style 283

Figure 17.7. A skewed distribution.

Hence m is set in such a way that
| foopOddx = | fixp()dx + risk-neutral drift

with f(x) the payoff and p(x) the probability attached to it.

This does not mean that an equal number of observations should fall on
both sides of the fence. A bet is only concerned with the number of observa-
tions, not the expected value of each.

Therefore the bet that x exceeds m on a certain date is simply

I - p(x)dx.

The skew by increasing the potential payoff on the left integral needs to
be compensated with a shift of the mean to the right to prevent the markets
from giving the short seller any higher expected return than the long holder
(this is what is called the fair dice argument) (Figure 17.8).

Delta
0.!
0.
0.
0.
Bet

100%
Volatility

Figure 17.8 Delta and bet with volatility rising.
[End OCR]*

## Page 298

*[Image OCR]
284 Trading and Hedging Exotic Options

Difference between the Binary and the Delta:
The Delta Paradox Revisited

Why isn’t the delta the probability of exercise?

Simply because the delta takes into account the payoff.

In the graph of the skewed distribution (Figure 17.7), the delta is simply
the right integral for a call and the left one for a put. The probability of
being in the money is, to put it bluntly, the binary option.

The difference is subtle. In dealing with a geometric Brownian motion,
the distribution is shifted to the right (the compounding effect of lognor-
mality was discussed earlier in Chapter 7), and the higher the volatility the
higher the shift as the asset grows in constant percentage. The higher
the volatility the bigger the shift to the right. This results in the increase of
the delta as a hedging protection.

This increase in the volatility is accompanied with an increase in the
right side of the distribution. The distribution at higher volatility will ex-
hibit a bulging right side that would illustrate the lognormality effect. Ac-
cording to the preceding principle, this will result in an increase in the
risk-neutral frequency of observations to the left to maintain the fair dice
condition in the environment (see Module B). Such increase in observations

Option Wizard: Pricing the European Bet

dt = (LoglS e-"/K en") + rt + 0/20 V0)

d2 = (LogS e-"/K e-"] + rt — 0t/2)o V0

A vanilla call = x N[d1] — k Exp[ —ré] Nfd2]

Delta = Nid}

Binary cash “call” = e~" N[d2]

Binary cash “put” = e7"' (1—N{a2})

Binary forward “call” = N[d2]

Binary forward “put” = (1 -N[d2])

S is the underlying.

ris the interest rate of the numeraire currency.

rf is the return of the asset concerned (foreign currency rate or dividend payout).
K is the strike price. |

t is time to expiration.
[End OCR]*

## Page 299

*[Image OCR]
Binary Options: European Style 285

Option Wizard (Advanced): The Dirac Delta

The Dirac delta is commonly used for the impulse function. It could be easy for
traders to visualize the behavior of volatility at one point in time surrounding a
special announcement. The daily volatility could still remain at, say, 16% but
the forward-forward volatility for that second would be immensely high (in the
hundreds or thousands), which would make it very difficult to measure.

The delta function can be simplified as follows: Take e, a small number, the
smallest possible unit that can be imagined. Then the rectangle defined by the
sides 1/e and e would have an area of 1 while all the areas surrounding it would
have a measure of 0.

A simplified definition: 8() = 1/e for 0 < t< e and 8() = O elsewhere.

Dirac Delta Simplified
x

1/e

The Dirac delta is often used for the gamma at expiration: The cumulative
gamma would be the delta, but it is seemingly very large at the very narrowly
defined expiration instant. It also applies to the delta of a binary option close
to extinction.

s

is needed to compensate for the difference between the payoff to the right
and the payoff to the left. The value of the bet will therefore drop.

The delta is there to account for both the possible payoffs and their fre-
quency while the binary option only accounts for the frequency.

Mathematical Note: A feature of the two-currency paradox is that the
delta for one party is the binary of the other and vice versa. This is due
to the Jensen’s inequality: The expectation of &« (the inverted price) is
not the inverse of the expectation of x, as discussed in Module C. Also,
for one particular strike, the delta of the put is equal to the price of the
[End OCR]*

## Page 300

*[Image OCR]
286 Trading and Hedging Exotic Options

binary call and vice versa. This can lead to the binary paradox—an in-
credible situation:

A bet in dollars for a dollar-based person on USD-DEM is different in
price from the translation into German marks of a bet in German
marks on USD-DEM of the same strike and expiration. Such a differ-
ence will increase through both time to expiration and an increase in
volatility.

The reason is straightforward: Define N(d2) as the price of a bet for a
dollar-based person. N(d1) will be the price of the bet for the person
based in DEM, by the numeraire inversion.

The extension of the paradox is a little unsettling: Two positions on
two different continents cannot be marked at the same price. Fur-
thermore, two traders performing a trade on opposite sides of the
fence would thus both show a profit or a loss on the same leg.

First Hedging Consequences

Bets need to be priced with the skew of the market taken into ac-
count using the preceding method.

The trader should not be deceived by the apparent lack of gamma and
vegas at the origin.

The best replication for a digital is a wide risk reversal (that would in-
clude any protection against skew). There will be a trade-off between
transaction costs and optimal hedges. The trader needs to shrink the

. difference between the strikes as time progresses until expiration, ata

gradual pace. As such an optimal approach consumes transaction
costs, there is a need for infrequent hedging.

When the bet option is away from expiration, the real risks are the
skew. As it nears expiration, the risks transfer to the pin. In practice,
the skew is hedgeable, the pin is not.

The Delta Is a Dirac Delta

As shown in Figure 17.9, the delta for a bet looks similar to the gamma of a
regular option. In addition, it will behave and “bleed” like the gamma of an
option. This is because a delta is almost a bet (in the risk-neutral universe,
to be correct). The delta of the bet will therefore be the gamma.

A familiarity with the Dirac delta function is helpful in understanding
the way the delta moves. The Dirac delta is an interesting function that has
zero value everywhere except for one point, but with the integral over the
map equal to 1. Likewise the delta at expiration is zero everywhere, but the
integral of the delta over the possible moves will equal the bet’s face value.
[End OCR]*

## Page 301

*[Image OCR]
Binary Options: European Style 287

Delta

90 100 110 120
Asset Price

Figure 17.9 Delta of a bet (strike 100): Three months and six months to
expiration.

The gamma of a vanilla option follows a Dirac delta on expiration second;
however, few traders take expiration gamma seriously, because it is only the
derivative of the exposure. With binary options, however, the awkward delta
is of some concern because many traders hedge it with cash when the market
is liquid enough for them to do so.

A delta can be viewed as the quantity that is necessary to purchase to
break even over a certain move by the cash. Because the trader knows how
much money he needs to make over one particular move, the notion becomes
almost trivial. However, in coming closer and closer to expiration, the quan-
tity that needs to be purchased will be close to nil in areas that do not cross
the “strike” of the bet. The delta for the bet, however, will become close to in-
finite in one very narrow point close to the strike point. In practice, the ex-
ample of the perfect replication holds: If the market were allowed to move by
only one-tick increments, between say 100.00 and 100.01, then the delta be-
tween 100.00 and 100.01 would be 100 times the bet payout size (if the trade is
of the at least variety; otherwise the trader would have to hedge between
99.99 and 100.00). A trader betting $100 would have to get a position of a face
value of 100/.01 = $10,000 to satisfy his risk.

Figure 17.9 shows how the delta concentrates around the strike as the
option nears expiration. Figure 17.10 show the “step function” of the bet
price very close to expiration time. Figure 17.11 show the Dirac delta, the
derivative of the function shown in Figure 17.10.

Gamma for a Bet

Because the delta of a bet resembles the gamma of an option, it becomes con-
ceivable that the gamma of a bet would resemble the DgammaDspot, or third
[End OCR]*

## Page 302

*[Image OCR]
288 Trading and Hedging Exotic Options

Price
(% payoff)
100

Asset Price

90 100 110 120

Figure 17.10 Price jump on expiration day.

Delta
(%)
17.5

15
12.5
10

wou Wl

Asset Price
90 100 110 120

Figure 17.11 Delta for a bet a few seconds before expiration.

z

Gamma

——te Asset Price
90 100 110 0

Figure 17.12 Gamma of a bet.
[End OCR]*

## Page 303

*[Image OCR]
Binary Options: European Style 289

Asset Price

120

Figure 17.13 Gamma in time and space.

derivative with respect to spot. It is interesting to see how trading exotic op-

tions can help a person learn about the behavior of regular instruments.
Like a vanilla option, the gamma of the bet is flatter and more stable

when the option has longer until expiration (Figures 17.12 and 17.13).

CONCLUSION: STATISTICAL TRADING VERSUS
Dynamic HEDGING

In summary, we deemed that it was extremely difficult to hedge a binary
with an option that has a continuous payoff owing to the onerous trans-
action costs. Even where there are minimum tick increments, a static hedge
exists that perfectly matches a binary payout but is impractical and impos-
sible to execute. So it is best to leave the hedging to structures that are likely
to bite, especially as they come closer to expiration.

Hedging the bet with deltas only would be easier and sometimes the
variance of the P/L could be reduced, but by no means could it be elimi-
nated. In some cases, it might be increased. However, it is easy to be “shoot-
ing” the delta, that is overhedging it, in zones where the trader would be
exposed to losses beyond his risk appetite. So if the 102 bet causes him some
headaches, he could buy more deltas at 101 and take the risks of having the
market drop back down. Buying way ahead of the barrier would then give
him some extra P/L at 102 that could be spent in transaction costs whipping
around the bet strike. “Shooting” deltas neither increases nor decreases the
total return. It only smoothes the P/L around the barrier at the expense of
slightly worsening it away from it.

Banks are sometimes best left to run some of the binary options like bets,
much as insurance companies can live with a certain amount of local risk that
[End OCR]*

## Page 304

*[Image OCR]
290 Trading and Hedging Exotic Options

Dollar Variance

Low Variance
High Number
of Strikes of Strikes

Statistical
Replication Mode
Mode Convergence
(Dynamic to Law of
Hedging) Large
Numbers
Risk-Neutral
Valuation Subjective
Valuation

Figure 17.14 Dollar variance test.

is possible to diversify away. The good news about binary options is that their
worst-case scenarios are bounded and diversifiable. Their vega vanishes as
they become close to the money presenting a meaningful protection. A hitch
is that in such cases, because of the absence of delta neutrality, traders need
to value the bets using their own subjective distribution, not the risk-neutral
one. Figure 17.14 shows the different hedging and valuation policies available
to institutions.

Many of the nonfinancial bets, like bets on political events, are of a bi-
nary nature and leave no choice: They do not lend themselves to continuous
hedging. Many banks issue notes with a payoff linked to a political event or
the determination of federal reserve policy (e.g., a note that bets on a dis-
count rate cut). Precisely because binary options have a discontinuous pay-
off, they can be easily adapted to such nonfinancial and nontradable bets.

So the trader can assume the pin risk when the amounts at stake are not
life-threatening. The object is to have enough of them to become “statistical,”
that is, let the variance be dominated by the number of bets. Otherwise, the
trader should be less statistical and more replicating.

CASE STUDY IN BINARY PACKAGES—
CONTINGENT PREMIUM OPTIONS

This combination is being covered despite being a simple construction to il-
lustrate the use of binary options in packages.
[End OCR]*

## Page 305

*[Image OCR]
Binary Options: European Style 291

M@ Contingent premium options are vanilla options where the option
buyer only pays the price of the option in the event of it being in the
money.

These options can thus be constructed with simple options plus a forward
bet for the premium amount.

This sounds pretty much like a free lunch except for the hitch: It creates
an area of negative profitability around the strike as shown in Figure 17.15.
The owner of the option will still have to pay the premium of his option if it
is slightly in the money, and such amount is, if the option is initially at the
money, twice what he would have paid for a regular option. In other words,
the only way to lose money on such a trade is to be slightly right.

Mechanically, it is generally composed of a regular option plus a binary
where the face value of the bet corresponds to the initial premium.

Recommended Use: Potential Devaluations

A generally good use of a contingent premium option is in distributional ar-
bitrage. The trader can have the contingent option struck at a price where he
believes there is a reflecting barrier, like a currency band, or an officially de-
fended limit to the move of an asset. The idea is that, should the price break
through such a level, the market will no longer be supported by the authori-
ties and a devaluation will cause the option to be clearly in the money.

Figure 17.16 shows a market mapping through an intervention level
with a vengeance.

Another interesting use is with strongly biased assets where down
moves are. rare but accentuated. The trader can structure a position to

Cash Price

Figure 17.15 Expiration P/L for a contingent premium option.
[End OCR]*

## Page 306

*[Image OCR]
292 Trading and Hedging Exotic Options

Asset
Price

Time

“aT ™ A

Figure 17.16 Bouncing against the intervention level.

benefit if the market precipitously falls with an at-the-money contingent
premium put.

A final use for such structure is in markets with the “fat tails” where
moves in either direction are accentuated, making intermediate moves
quite rare.

CASE STuDyY: THE BETSPREADS

Betspreads are simple binary call or put spreads. There is no mystery to
them as they are obtained by construction. A betspread is a structure that
collects a certain amount if the share price ends up between strike 1.and
strike 2 and collects nothing elsewhere.

From a trading standpoint, they resemble butterflies and condors (a
butterfly with 4 strikes) in the sense that if a binary bet resembles a call
spread, a betspread resembles a condor.

While this book is not concerned with single strategies, betspreads
merit some attention. The study of betspreads can enlighten the trader in
the management of a pure binary options book. Because a binary option is
by construction a risk reversal, a betspread is therefore a double risk rever-
sal. This is encouraging as the position risk can be shifted down by one mo-
ment of the distribution, as shown in the following analysis. ,

A six-month betspread paying $1 if the asset ends up between 100 and
105 would look as depicted in Figure 17.17.

With time, the price of the betspread would naturally converge to the
final payoff, and the risk would become topical. Figures 17.18 and 17.19 il-
lustrate the effect of the passage of time on the price of the structure.

From that, it becomes easy to conclude that perhaps the best hedge for a
binary option is a similar binary option.

Examine the gammas. Figure 17.20 shows the gamma of a betspread.
Compare it with Figure 17.12 (gamma of a bet). The risk reversal (third mo-
ment) aspect seems to be replaced by a fifth moment instability, which is, in
this case, preferable.
[End OCR]*

## Page 307

*[Image OCR]
Binary Options: European Style 293

Spread

Asset
100 105 110 1 120

95

Figure 17.17 Price of the betspread.

Spread

Oo CoO Oo 8
oOo & DBD DW F

Asset
95 100 105 110 115 120

Figure 17.18 Betspread one month to one minute to expiration.

Time to
Expiration

Figure 17.19 The effect of time on price.
[End OCR]*

## Page 308

*[Image OCR]
294 Trading and Hedging Exotic Options

Gamma

- Asset Price
10 9 10 110 3=115

Figure 17.20 Gamma of a betspread.

Asset Price

Figure 17.21 Gammas of a series of betspreads.

The accordion shown in Figure 17.21 is a series of betspreads.
Conclusion: the bets are easier to hedge than regular options as the
asymmetry of the vegas can be reduced with the trading of another bet.

Advanced Case Study: Multiasset Bets

A bet that pays if either of the two commodities ends up higher’or lower
than a certain level at expiration is called a multiasset bet.

The reader can, as an exercise, study the exposure to the correlation of a
structure that presents two bets, of the either-or variety, using the preced-
ing methodology and that presented in Chapter 22. To price it intuitively, he
can imagine the value of the limit decomposition of a rainbow call spread.
[End OCR]*

## Page 309

*[Image OCR]
Chapter 18

Binary Options: American Style

Adam K., one of the greatest option traders of all times (and one of the fastest
minds the author ever met) called the author with the following request:

“T hear you are giving a 4-day seminar on hedging exotic options. Can you
give it to me during lunch time tomorrow? All I need is the beef. You see, I don't
have patience for details.”

AMERICAN SINGLE BINARY OPTIONS

@ American binary (or digital) option is a common bet option with the
difference that it pays when the price is “touched” rather than if the
price settles above or below a certain price.

As defined, the American binary is path dependent (Figure 18.1). A trader
cannot go on an extended trekking vacation and come back on expiration
day to satisfy the bet obligation. The bet can thus reach expiration any time
before.

The first noticeable difference between the American and the European
bet is the price. Because the option has at least equal chances of being
knocked out, it is therefore a larger outlay than the European one for the
“buyer” at all times (generally twice).' The buyer is defined as the person
who receives the payment if the price is touched. The American bet will
thus be more expensive with the consequence that the risk of the lack of sat-
isfaction for the seller will be much larger.

Option Wizard: The Contamination Principle and Barrier Options

If there is a point, namely the barrier, that can create some positive P/L (the
termination of the bet is positive for the buyer), then the position will be long
gamma in smoothly decreasing amounts in areas surrounding it. It will there-
fore be long vega from that position.

If, for reasons of carry (see Girsanov), the bet termination results in nega-
tive P/L, then, again the areas close to such point will carry a negative gamma.

295
[End OCR]*

## Page 310

*[Image OCR]
296 Trading and Hedging Exotic Options

Price

0.34
0.24
0.14 Volatility

0.04

nL
© o
o> -

oo SG &
Oo _ CO
- oO
Asset Price ~

Figure 18.1 American digital.

Example: Spot trades at 100. Assume volatility at 15.7%. A European

bet that spot will be higher than 105 at expiration 100 days hence will be

worth 26% of the face value. An American bet will be worth 51% of the
face value of the payoff.

Another difference is that because, unlike the European, it is not possi-
ble to cross the fence (i.e., move “in” the money), the American bet will al-
ways be long volatility for the seller. Intuitively, it is easier to sense that,
unlike another structure, the movement would benefit the seller if it takes
the option closer to the money. If movement one way can translate into
higher profits, by the contamination principle the movement should help in
the other direction.

It was established in Chapter 17 that the European binary is a risk
reversal. The American binary is generally monotonic in gamma, though
not always (Figure 18.2).

Risk Management Rule: With no drift (forward trades at flat with

cash), the American binary never changes in the sign of the vegas
and gammas: It remains a pocket of localized long vega.

The issue of interest rates will be tackled in Chapters 19 and 20. A posi-
tive curve can lead to negative gamma near the bet price for an American bi-
nary and positive gamma away from it. The rule is as follows:
[End OCR]*

## Page 311

*[Image OCR]
Binary Options: American Style 297

Risk Management Rule: The American binary will be positive
gamma every where (for the “owner,” the person who earns the pay-

off when the barrier is hit) when the delta hedge against the option
incurs a negative carry.

For a definition of negative carry in an asset, traders can easily use the fol-
lowing rule: forward price for the period t + 1 is higher than the forward
price for period f (the spot).

Risk Management Rule: The profile of the American binary will
look like that of a risk reversal (i.e., a third position moment differ-

ent from 0) if the delta hedge of the owner earns a positive carry su-
perior to the time decay of the same binary on the same asset
without drift.

A peculiar aspect of the American binary that leads to the notion of
stopping time is that the distribution of the expected time to extinction is
what matters. In a way, an American binary is a bet on time.

The expected exit time has been around for a while in probability the-
ory. Module G provides technical details and references. The issue also will
be discussed further in Chapter 19.

0.28
0.16

0.04 Volatility

ed oO
se] Oo
oO -
—

108.5

Asset Price

Figure 18.2 Volatility sensitivity of an American digital.
[End OCR]*

## Page 312

*[Image OCR]
298 Trading and Hedging Exotic Options

Hedging an American Binary: Fooled by the Greeks

This section provides a simplified tutorial for the management of the vega
risk of any nonvanilla structure with a concave vega.

Case Study: National Vega Bank

This case study describes the attempt to hedge the following option: a 105
“if touched” bet paying $1 on termination per unit. The trader at National
Vega Bank becomes, due to an ugly structure that is otherwise stripped and
hedged, long the 105 American bet in $10,000,000. Spot is at 100 and, as
usual, the forward trades at flat. Volatility, as in most examples, is at 15.7%
which corresponds to 1% a day move 252 days a year. The bet expires ex-
actly 100 days from now. “Fair” value (better called unfair value) is 54.7%
(i.e., $5,470,000).

Table 18.1 P/L of the Three Legs of the Trade

Price Price Option P/L Hedge P/L Total P/L
94.8 0.21119 —3353 4160 807
95.2 0.23034 ~3162 3840 678
95.6 0.2506 —2959 3520 561
96.0 0.27201 —2745 3200 455
96.4 0.29455 —2520 ~ 2880 360
96.8 0.31823 —2283 2560 277
97.2 0.34304 —2035 2240 205
97.6 0.36896 —1775 1920 © 145
98.0 0.39598 -1505 1600 95
98.4 0.42407 —1224 1280 56
98.8 0.45321 —933 960 27
99.2 0.48335 —631 *640 9
99.6 0.51447 —320 320 0

100.0 0.5465 0 0 0

100.4 0.57941 329 —320 9

100.8 0.61313 666 —640 26°

101.2 0.64761 1011 —960 51

101.6 0.68279 1363 —1280 83

102.0 0.7186 1721 —1600 121

102.4 0.75497 2085 -1920 .. 165

102.8 0.79184 2453 —2240 213

103.2 0.82911 2826 —2560 266

103.6 0.86673 3202 —2880 322

104.0 0.90462 3581 —3200 381

104.4 0.9427 3962 —3520 442

104.8 0.98089 4344 —3840 504

105.2 0 Closed Closed 504
[End OCR]*

## Page 313

*[Image OCR]
Binary Options: American Style 299

By “buying,” the trader has a similar position as being long a call-
spread, which is long delta. In addition, as he is betting on the spot moving,
he will be long volatility.

The first thing that comes to mind is the delta. The trader will hedge it
by selling the equivalent of 8% of the face value of the security, namely
$800,000.

Table 18.1 and Figure 18.3 show the resulting three legs of the trade:
The P/L for the binary, that from the delta, and the final net. As can be
seen in Figure 18.3, the position is mildly long gamma. However, the
gamma disappears beyond the barrier. In addition, the trader would need
to rapidly unwind his position. As shown in Figure 18.3 and discussed in
the Option Wizard, The Gap Delta, there is a discontinued payoff. Should
the trader avoid catching the market on time, the hedge would be off by a
slippage factor.

Figure 18.4 shows the effect of time at 2 asset price levels. The gamma
would increase, as time flows, near the barrier and decrease everywhere
else. The slope of the decay at 103, 2 points away from the barrier, is consid-
erably steeper than the slope at 100, 5 points way. As decay and gamma go
in pair, the reader can infer an increase in the gamma near the barrier with
the passage of time.

The Ravages of Time

Figure 18.5 depicts the initial position with fewer days to go, showing the
effect of time. Needless to say, the position would increase in gamma as

P/L

Hedge portfolio

Figure 18.3 Gamma position: P/L of the three legs of the trade.
[End OCR]*

## Page 314

*[Image OCR]
300 Trading and Hedging Exotic Options

Option Wizard: The Gap Delta (1)

The gap delta is the difference in deltas near a particular barrier. The trader
needs to unwind a certain amount of deltas that were used to hedge a barrier
structure by dumping that amount in the market at a given price.

The trader has no guarantee of getting the exact target price, owing to liq-
uidity matters. Trying to execute ahead of the barrier is equally dangerous as it
raises the dangers of the “fake trigger,” a situation where the market trades
near the barrier then pulls back without hitting it.

The difference between the price the unwinding was executed at and the
barrier is called the slippage (as defined in Chapter 4). Many traders get negative
surprises as markets near the barrier. {liquid markets can be vicious as they un-
dergo liquidity holes (as defined in Chapter 4) near a particular barrier.

The gap delta receives extensive treatment in other parts of the book, as it
deserves as much attention as possible.

time comes closer to expiration. As time goes by, the following problem
takes place: the stakes get bigger close to the barrier.

The P/L cannot be read between two different dates on the graph on
Figure 18.5. It corresponds to the recentering of the position at 100 the day
of the run. The P/L lines are all made to be zero at 100, which prevents the
user from comparing two dates, but allows visualizing the P/L within one
date. Such an issue of noncomparability is explained in the Option Wizard,
Looking at a Graph through Time.

Note. It is assumed that the trader did not unwind his delta in the mar-
ket after the barrier is hit. This shows the precipitous drop in P/L as the bi-
nary, being terminated ceases to accumulate positive P/L while the delta
continues to build a negative P/L.

From Figure 18.5, it is evident that the delta becomes larger close to the
strike next to expiration. So the trader needs to sell the asset in the rally.

’
0.8

0.6 103 Spot

0.4

Price

02 100 Spot

0
1 7 13 19 25 31 37 43 49 55

Days to Expiration

Figure 18.4 American binary 105 calls: Over time.
[End OCR]*

## Page 315

*[Image OCR]
Binary Options: American Style 301

Option Wizard: Looking at a Graph through Time

Unlike many functions that one can easily project through time, an option po-
sition needs to be analyzed while keeping in mind that markets are not just
frozen between periods and that traders are not stuffed birds. The positions
will change and traders react accordingly.

A graph representing the same position one month later would not cor-
rectly forecast the P/L as it does not include the appropriate adjustments an
option position continuously undergoes.

However, the following takes place: These deltas present as many potential
problems to deal with, as the knocking-off of the binary would necessitate
unwinding larger and larger amounts of money. The trader would need to
buy back what he sold. The higher the delta, the more he would need to sell,
the more he would have to buy back. From whom? Most probably from
those he initially sold to.

In addition, look at the potential danger: What if the binary were not
knocked off? The P/L accumulated on the buildup of time value in the bi-
nary would deflate rapidly, causing considerable time decay.

Look then at the same graph without recentering it (Figure 18.6). It is
easy to infer that the position increases in gamma through time and, conse-
quently, in time decay. As can be seen, the time decay becomes a serious
matter close to the end. The P/L should be interpreted from the origin of
100 to the next point. It assumes delta hedging at the date and at the price of
100 without subsequent hedging until the barrier is hit. So the 10-day posi-
tion was made delta neutral on the delta of the tenth day, considerably lower

~~ 30 days
70 days

Asset Price

Figure 18.5 Binary option: 70 days through 10 days.
[End OCR]*

## Page 316

*[Image OCR]
302 Trading and Hedging Exotic Options

0.3 +
0.2 +
0.1 +

O fel det be ee tet d beet

P/L

-0.1
-0.2
-0.3

Ts.

Figure 18.6 P/L through time.

than that of the seventieth day. Otherwise, the price of the bet would ex-
hibit different features.

Another piece of information can be gathered from the same graph but
without delta hedging.

There is a subtle difference between Figures 18.5 and 18.6. In the previ-
ous one, the 30-day line came below both the 10-day and the 70-day, owing
to the delta neutrality at 100.

Figure 18.7 provides the intuition of the true risks. The longer the time
to expiration, the more horizontal the slope of the asset price to the option
value, the lower the risks. With little time left on the clock, the trade is safe
away from the strike and risky near it. Figure 18.8 shows the corresponding
gamma.

Option Price

xn ~ @
tk Oo
on oO

100.8

102
103.2
104.4
105.6

Asset Price

Figure 18.7 American binary option: Price through time.
[End OCR]*

## Page 317

*[Image OCR]
Binary Options: American Style 303

Gamma

RPNWPR UD ~I

Asset Price
92 94 96 98 100102104

Figure 18.8 Gamma as near expiration.

Understanding the Vega Convexity

In Figure 18.9 the first noticeable effect is that the difference becomes thin-
ner at the extremes. While every option trader knows than an option’s vega
diminishes away from the strike, this seems to be quite exaggerated. In ad-
dition, the vega vanishes completely at the barrier.

This effect is not trivial. It raises the problem of how to hedge the vega
of an American binary option with instruments that do not vanish at the bar-
rier. Many attempts at the so-called static option replication have been made
without much success at finding a structure that can imitate the payoff
everywhere and cancel the pin risk.

The notion of vega convexity is best demonstrated by creating a portfo-
lio that is long vega through the binary and short vega with other instru-
ments in the market that present a linear sensitivity to implied volatility.

First, look at the effect of volatility changes on one structure, as shown
in Figure 18.10.

American Binary

Price

~> 95 97.5 100 102.5 10
Asset Price

Figure 18.9 Volatility effect.
[End OCR]*

## Page 318

*[Image OCR]
304 Trading and Hedging Exotic Options

Price

80
60
40
20

Volatility
0.050.10.150.20.250.30.35

Figure 18.10 Sensitivity of an American digital to volatility.

The reader can see that the function is concave. Higher volatility causes
the price to increase, but at a decreasing rate. It means that should the
trader go vega neutral against the structure, he will end up short volatility
of volatility, or to use a different language, short the fourth moment. From a
trader’s experience, the worst moment to be short is the fourth moment be-
cause, unlike the third moment (the skew), there is generally no clear com-
pensation for the risks incurred. The prices in the market, using constant
volatility:models, do not factor such exposure.

Starting at 15.7 volatility, the trader decides to hedge the vega using an
instrument of the same official duration. So he sells enough volatility to be
“vega neutral,” that is for small moves in volatility.

Figure 18.11 shows the volatility concavity. This concavity is mostly due
to the shortening of the first exit time as volatility rises. The vega sensitiv-
ity shrinks as the trader comes near the trigger. Such concavity would
change according to the position of the market relative to the barrier. Figure
18.12 shows the concavity close to the barrier which is more pronounced
than the one shown in Figure 18.11.

P/L

Volatility
0.150 20.250.30.35

Figure 18.11 Volatility concavity of an American binary.
[End OCR]*

## Page 319

*[Image OCR]
Binary Options: American Style 305

P/L

Volatility
0.30.35

-5

-10

-15

-20

Figure 18.12 Concavity close to the barrier.

So far, the following risks have been shown:

¢ Duration Risks. The barrier is a bet on duration. Vanillas do not
hedge in a shifting volatility curve against the shifts, lengthening, and
shortening of the structure. So much for static option replication.

* Gap Risks. The risk of having to pay up to unwind the deltas close
to the bet level.

¢ Vega Risk (after the barrier). Vega hedges need to take into account
that the exposure vanishes past the bet price. It means that whatever
hedge the trader performs should shrink accordingly past that level.
Good luck in finding one.

* Vega Concavity (for the owner). Whatever structure is used as a
hedge needs to have a positive fourth moment (i.e., long volatility of
volatility) to be able to offset the risks of the structure.

These risks, inherent to the American binary, are present in some shape or
another in all barrier options.

Trading Methods

s

American binary options are truly options on time rather than options on the
asset. As such, they can only be hedgeable with instruments that are simi-
larly options on time. These instruments are necessarily other barriers, with
similar payoff, at some trigger point not too far removed from the bet itself.

Hedging a long position in American bets with European vanilla options
is a true losing proposition, except in cases where one prices the fourth mo-
ment risk into the structure. Hedging them with binary options of the Euro-
pean variety creates an illusion of a hedge: The European bet is altogether
another instrument, except on expiration minute.

The greatest danger traders face in their activity comes from the reading
of the risk on a computer sheet. Vegas stemming from a European vanilla
[End OCR]*

## Page 320

*[Image OCR]
306 Trading and Hedging Exotic Options

with its clear and well-defined duration are more reliable than those arising
from the American binary options with unknown duration and unstable mo-
ments. So traders who navigate with the Greeks are going to be subjected to
serious dangers. “Non-Greek” traders (i.e., nonparametric traders) who ab-
sorb as much information about their strikes fare better. That is why it is best
recommended to divorce a barrier book from a vanilla position to better see
the strikes and navigate without the superficial information about the lower
Greeks.

Case Study: At-Settlement American Binary Options

Some binary options, as if the American binary feature was not deleterious
enough, present the additional difficulty of a conditional feature. Besides
the “if touched” variety, there is the “if settled” type that makes the regular
American binary seem easy to trade by comparison.

An if settled binary only pays the bet value if the underlying asset offi-
cially settles through the trigger. This means that the asset price needs to
settle higher than the trigger if the bet is a call and lower than the trigger if
the asset is a put.

Many traders believe that such a feature comes into play on the last day.
But the last day for an American binary is going to be every possible day.

This feature acts on the price of the binary in a strange way: It creates a
negative gamma hole around the price of the barrier. Earlier, it was seen that
the important feature of an American binary bet is that (except for some cases
of a high carry on the delta), it remains long gamma for the “owner” until ter-
mination. This feature is attributable to the fact that the position could re-
main only on one side of the barrier and could not cross it at all without
termination, which would make the vega and gamma dimension mute.

An American if settled binary would therefore act like a European binary
during the day and like an American between days as shown in Figure 18.13.
This causes serious complications with the unwinding of the gap delta.

s

Negative

105 Gamma

Buy back
sell a

sol

Positive Gamma

Figure 18.13 Mixed gamma around the if settled American barrier.
[End OCR]*

## Page 321

*[Image OCR]
Binary Options: American Style 307

Example: A 105 Call if settled American bet pays $1 on termination. The
trader sells deltas against the position. As the market rallies, he would
need to sell more to benefit from the probability of cashing in on the bet.
However, he would have to buy everything back if the strike is crossed.
Should the market cross the gap early in the day, the uncertainty would
last the longest. The trader has no way of knowing whether he would re-
main in the terminating zone or if he would cross again to reenter the op-
posite area. What if the market moved back down and settled on the
“negative” side of the bet? This constitutes a negative gamma position
around the barrier that vanishes in the extremes, high above or low
below, as depicted in Figures 18.13 and 18.14.

Other Greeks

Studying the behavior of other Greeks such as the rho is generally futile be-
cause of the instability of the measure. Having established that the gamma
was unusual, there is no need to go into secondary, less significant Greeks.

AMERICAN DOUBLE BINARY OPTIONS

@ An American double bet is a bet whose condition is satisfied if the mar-
ket touches either of two levels during the life of the option. These levels
are generally called high barrier and low barrier.

A particularity of American double bets is that they are not the sum of two
bets. A European double bet is SDF decomposable, which means that one
could add two European binary options and get the resulting European dou-
ble bet, which needs not be examined here for that reason. An American dou-
ble bet is a structure that terminates whenever one of the legs is touched,
which is a serious difference. The SDF is explained in the Option Wizard,
Smallest Decomposable Fragment, in Chapter 2.

*

Bet Price

100
95
90
85
80
Asset Price
90 95 100 105 110

Figure 18.14 American bet.
[End OCR]*

## Page 322

*[Image OCR]
308 Trading and Hedging Exotic Options

One feature to remember about the American double bets is that, coun-
terintuitively, they will trade at a very narrow discount of the bet value (it is
usually expressed in percentages; a bet priced at 96% pays 100 for every 96
dollars invested in it). Prices such as 98% or 99% are common for options
that are a few weeks away with barriers straddling the market by only 2 or 3
daily standard deviations.

An interesting feature of American bets is how short their expected exit
time may be. A double bet trading at 80% of face value implies, roughly, as-
suming no drift and no interest rates, that it has about 20% of the nominal
time to live.

The double bet feature is often present in the so-called range notes
where the investor is sold financial assets and notes that pay their coupons
according to whether the market remains within some range or not. In that
case, the assumption is that the coupon is the face value of the bet. Figure
18.15 shows the profile of the double binary with time.

Vegas of the Double Binary

The barrier reacts to an increase in volatility by shortening its expected
time of arrival. This presents a very concave way of being long vega. For
that reason, even hedging by selling other concave vegas (like a single bar-
rier) might not be sufficient. The instrument trading at 80 or 90 can only
reach 100 as volatility increases, which limits its vega power. The concav-
ity, as a rule of thumb, can be ascertained from the following: As the
amounts to possibly lose from a trade become considerably higher than
the amounts to possibly earn, the concavity feature intensifies.

It is easy to see that above an implied volatility of 20% the bet reaches a
price close to 100%.

Price

0.8

0.6
0.4 Time (Years)

100 0.2
Asset Price 110

Figure 18.15 American double binary.
[End OCR]*

## Page 323

*[Image OCR]
Binary Options: American Style 309

Table 18.2 shows the trade of a double bet at 16% volatility (Figures 18.16
and 18.17) and an attempt to hedge the vega exposure with a linear vega. The
amounts are expressed per unit of a 1% bet.

Other Applications of American Barriers

Amortizing Interest Rate Securities

@ The maturity, coupon, or in some structures, the principal, of amortiz-
ing interest rate securities, varies according to whether some prede-
fined rate was reached in the market. They can be defined in any
possible manner but those that are path dependent can irreversibly de-
cline in coupon payments or principal if some interest rate level is
reached.

As such, it is safe to assume that the amortizing swap is a zero coupon in-
strument plus a series of forward American bets on every coupon payment,
each bet being a specific level in the market. This assumption requires no
special feature like requirements for the market to remain under some level
for a prolonged period.

They can be forward bets because each bet is timed like the payment of
a regular coupon. Again every structure varies from pricing purposes but
the trading intuition remains that of a sum of American bets.

Amortizing swaps represent another situation where the traders were
long American digital options and discovered themselves to be long gamma
until the barrier. The traders sold options to flatten out their gamma and
rapidly discovered the notions of vega convexity and that of a vanishing leg
of a hedge. As these securities were fashionable for a brief period, the build
up in inventory caused a few spectacular losses when the dealers rushed to
buy back their gamma. ,

Bet price

100

©
2

~ Asset Price
90 95 100 105 110

Figure 18.16 Volatility effect of 12, 16, and 18% on a one-year 90/110 American
bet.
[End OCR]*

## Page 324

*[Image OCR]
310 Trading and Hedging Exotic Options

Table 18.2 A Vega Hedge for a Double Binary Option

Volatility (%) Price (%) P/L Hedge P/L Total
6 24.4 —70.5 23.71 —46.81
7 35.2 —59.7 21.34 —38.38
8 46.0 —48.9 18.96 —29,.92
9 56.2 —38.7 16.59 —22.12

10 65.3 —29.6 14.22 —15.38
11 73.2 —21.7 11.85 —9.86
12 79.8 —15.1 9.48 —5.63
13 85.1 —9.8 7.11 —2.66
14 89.3 —5.6 4.74 —0.83
15 92.5 —2.4 2.37 0.00
16 94.9 0.0 0.00 0.00
17 96.6 1.7 —2.37 —0.66
18 97.8 2.9 —4,74 —1.84
19 98.6 3.7 —7.11 —3.40
20 99.1 4.2 —9.48 —5.23
21 99.5 4.6 —11.85 —7.26
22 99.7 4.8 —14.22 —9 Al
23 99.8 4.9 —16.59 —11.65
24 99.9 5.0 —18.96 —13.95
25 99.9 5.1 —21.34 —16.28
26 100.0 5.1 —23.71 — 18.63
27 100.0 5.1 —26.08 —21.00
28 100.0 5.1 — 28.45 —23.37
29 100.0 5.1 —30.82 —25.69

10

’
0 bebe EH
| & g& € 3 SX S
tof © &5 8 6 BERK
oS -20

-30

-40

-59 1

Volatility

Figure 18.17 Nasty vega neutrality.
[End OCR]*

## Page 325

*[Image OCR]
Binary Options: American Style 311

Option Wizard: Too Much Hedging Is Bad for You

This chapter will end with the following war story: An aggressive derivatives
house lost a considerable amount of money on a trade where they managed to
buy “cheap” volatility with a series of knock-out options using rebates, in fact
obtaining a strip of American binary options. They had a considerable “margin”
in the trade, which means that the trader went home with some theoretical
marks-to-market profits and thought highly of himself.

The trader was asked by his boss to manage the risk and reduce the expo-
sure. “Cash-in on the money,” he was told. This was a reference to the money
they believed they had earned off the customer. So the trader examined the
Greeks and sold the vega, to perfect the vega neutrality in a manner that would
befit such a sophisticated house as his.

A few weeks later, the trader was out of a job. Volatility exploded and he
lost considerable money due to the difference in convexity between the vegas
he owned and the vegas he so massively sold. His boss, not very recipient
about notions of fourth moments and other nerdy matters, fired him on
grounds that the trader did not properly offset the exposure. Actually, the
trader, hedging a double barrier with out-of-the-money options, was even short
the sixth moment.

The trader, like most derivatives traders who lose their jobs, landed himself
a better position (he had gained valuable experience) and concluded with the
following wisdom: Hedging increases your risks.

Risk Management Rule: It is necessary to avoid hedging a discon-
tinuous exposure with continuous one.

Credit Risk

s

Assume that a trader buys a USD denominated Mongolian government
note. As the note will trade at an interest rate differential with the default-
free USD rate, it is convenient to view the payment differential as a form of
forward American bet on the default by the Mongolian government. Again,
. the American bet is a general framework: One could estimate that the face
value of the bet is the total paper minus some recovery value.
[End OCR]*

## Page 326

*[Image OCR]
Chapter 19

Barrier Options (1)

A true trader is a human being endowed with the rare gift of a positively sloping
learning curve.

This chapter opens with a discussion of knock-in and knock-out options (of
the regular variety that knock in or out when they are out-of-the-money).’
These options represent the simplest forms of barrier trading and hedging.
Reverse barrier options will be covered in Chapter 20. The terminology “re-
verse” indicates the unusual nature of the transaction, as well as the hedg-
ing difficulties incurred in its management. Table 19.1 shows the different
barrier categories and rates their trading complexity.

With reverse barriers, the payoff at the barrier is significant. The trader
needs familiarity and experience with regular knock-outs before muddling
in the waters of reverse barriers. Prior understanding of American binary
options is a must. .

Rebates will be covered in the next chapter, as those make a barrier op-
tion closer to an American binary. The next chapter will also discuss double
barriers as well as the major possible variations around the barrier theme.

Unless otherwise mentioned, the underlying options concerned are
European style.

BARRIER OPTIONS (REGULAR)

Regular barrier options include calls down and out, puts up and ut, calls
down and in, puts up and in. This section is required to understand all
other barrier structures.

Knock-Out Options

@ A knock-out option is a regular option with a second strike price, called
“trigger,” “outstrike,” or “barrier.” The option is considered expired
when the second strike, the barrier, is crossed.

wou

The terminology varies (as with most new products), but knock-outs
are defined as having the trigger placed in such a way that the option dies

312
[End OCR]*

## Page 327

*[Image OCR]
Barrier Options (I)

Table 19.1 Categories of Single Barrier Options

Trading Difficulty
Category Description and Risks
Down and out Regular knock- The option dies Low
call out call out-of-the-money
Down and out put Reverse knock- The option dies High
‘ out put in-the-money
Up and out call Reverse knock- The option dies High
out call in-the-money
Up and out put Regular knock- The option dies Low
out put out-of-the-money
Down and incall Regular knock-in The option is born Low
call out-of-the-money
Down and in put Reverse knock-in = The option is born High
put in-the-money
Up and in call Reverse knock-in The option is born High
call in-the-money
Up and in put Regular knock-in The option is born Low

put

out-of-the-money

when it is out-of-the-money. The trigger will be below the current spot if
the latter is below the strike. It will be above the current spot if the latter is
out-of-the-money. Otherwise, the option is called a reverse knock-out. Re-
verse knock-out structures are more difficult to trade than knock-outs be-
cause the option vanishes when it holds a large intrinsic value, a difficulty
that causes traders to call a regular knock-out a vanilla by comparison.

Knock-out options are often unaccommodating to trade because of the
discontinuity in the deltas as the option crosses the trigger. Slippage is
greatest with them, particularly when gaps occur and the trader who needs
to unwind his hedge at some price ends up doing so at a markedly less at-
tractive level. This is one reason that dynamic hedging is said to be costly.
When in possession of a large enough size to impact the market, however,
some traders are known to create considerable profits out of their struc-
tures. These are the cases that create liquidity holes, much to the consterna-
tion of the financial community, whose members see the knock-outs as a
source of large volatility.

Example: If the underlying security trades at 100, assuming no for-
ward curve, a one-month at-the-money option would trade at 1.80% of
face value.
[End OCR]*

## Page 328

*[Image OCR]
314 Trading and Hedging Exotic Options

The same at-the-money call but knocking out 2% below the present
strike would trade for 1.34% (plus a higher commission owing to its
wider bid /offer spread).

From the customer point of view, if the market should rally immedi-
ately, the knock-out will be equivalent to a regular call obtained cheaper.
In a break, the worst-case scenario will be the same. The knock-out
would fare worse if the market went down prior to the rally, a scenario
that tends to occur somewhat more often than the theory would suggest
(Figure 19.1).

Customer Demand for Knock-Out Options. Many customers prefer to
use barrier options for the simple reason that they come cheaper. In addi-
tion, many users of an option prefer to rid themselves of it when it is no
longer needed.

The major users are as follows:

¢ Funds managers holding a large stock market exposure, in order to
lower their hedging expenditure. Funds managers owning stocks al-
ways prefer to own puts as protection. They generally feel, however,
that 5% above the market they no longer need such protection as they
would normally flatten themselves up in such conditions. Such a
profile is similar to that of a risk reversal (or collar) except that they
would not necessarily lose their stocks in a violent rally.

¢ Speculators who believe strongly in trends and serial correlation (or
dependence) in a market. A fund manager prefers to be long as long
as the market does not sell below a certain point. He can thus buy
an out-of-the-money call that knocks out at a certain price if he-is
wrong. The knock-out agrees with the psychology of trend follow-
ers. Its path dependency conforms to their beliefs in the shape of the

0 Asset Price

104 107
-20

Figure 19.1 Constant volatility comparison, P/L a vanilla and a knock-out call.
[End OCR]*

## Page 329

*[Image OCR]
Barrier Options (I) 315

distribution. They believe that the world is going to behave differ-
ently than a random walk and that they can translate their beliefs
into option trades. Most frequently, the barrier would be placed close
to a chart point or some level that chartists consider critical for the
market direction. When chartists lose faith in a trend, they do not
mind canceling their bet.

¢ Corporates who have a contingent exposure but who believe that
they could square it up should the market move their way.

Discontinuity at the Barrier. The delta of a knock-out shows discontinuity
at the barrier. In the earlier example of 100 calls/KO 98 (pronounced 100
calls knock-out 98 in the vernacular), it moves from .66 to zero. The seller of
the option can either jump up in joy for having lost a liability or jump up in
anger at having lost considerable money by being filled at worse than the
stop level he set at 98.00.

Figure 19.2 shows the behavior of the delta of the KO at different asset
price levels. In this example, the trader who is short the knock-out call needs
to buy more contracts than if he were short a vanilla. While the vanilla has
the at-the-money delta close to 50%, the knock-out option has a delta of about
68%. This is to compensate for the fact that in the rally the knock-out and the
vanilla will be equivalent in price, whereas the trader sold the knock-out at a
lower price than the vanilla. As the market goes higher, however, the deltas
and prices converge.

Should the market jump down by 2 points immediately, the 68% deltas
would generate close to $1.36 per unit sold. To be exact, $1.34 because the
final delta in the example is 66%, causing the structure to be hedged at an
average delta of .67 (owing to the small gamma effect of the move). The
trader sold the option at 1.34. But he needs more than 1.34 to cover the trans-
action costs of liquidation at the barrier.

Asset Price

98 100 102 104

Figure 19.2 Delta comparison between a vanilla and a knock-out.
[End OCR]*

## Page 330

*[Image OCR]
316 Trading and Hedging Exotic Options

Trading Rule: The transaction costs of unwinding the hedge

should be taken into account when selling a barrier.

Let us say that some panic choked the market and that the execution slip-
page at the barrier turned out to be .10. This translates into a final price for
the barrier of at least .1 < .66 = 0.066 more than the initial 1.34 should the
option terminate. Therefore the minimum fair value for the operator will
be 1.34 plus .066 times the probability of hitting the barrier (which corre-
sponds to the value of an American binary paying .066 at 98 and expiring at
the same time as the barrier option). It will be roughly 1.39. That is the min-
imum, as more costs will crop up to affect the structure.

As the market approaches the barrier, the deltas of the knock-out and
the vanillas start to diverge: The trader needs to accumulate more and
more delta to make up for the fact that he might lose the option soon and
be freed of an obligation. When the asset becomes close to the barrier,
however, he needs to unwind the deltas as they become superfluous. It is a
double-edged sword. Hitting the barrier is a blessing because the trader
loses a liability (the option), but the way the barrier is touched is signifi-
cant. Should the market gap through the barrier (it often does so), the
hedge would be unwound at a worse level than expected and the trader
would wish he had not hit the barrier. The slippage around the barrier is
meaningful. .

Figure 19.3 shows the 100 call/90. The option price and deltas start re-
sembling the vanilla when the difference widens. It becomes conceivable
that less is at stake when the difference between the strike and the trigger
widens.

88 90 92 94 96 98 100

Figure 19.3. Delta for a 1m out-of-the-money knock-out and a vanilla.
[End OCR]*

## Page 331

*[Image OCR]
Barrier Options (I) 317

Knock-In Options

The knock-in option is an option that comes into existence when a particu-
lar price is hit in the market. A regular knock-in comes into existence as it is
out-of-the-money which simplifies its trigger conditions. A reverse knock-
in comes into the world with intrinsic value (by the forward), which com-
plicates its existence.

Studying knock-in options as their own instrument must be limited be-
cause they are, by construction, nothing but a combination of a vanilla and
a knock-out. This will be shown later in the section.

Customer Demand for Knock-In Options. For the same reason knock-out
options are a device for a trend follower, knock-in options are a device de-
signed for a mean-reverting mentality. They are inherently options that
come alive against the market direction.

The knock-out options can be used in structures. But perhaps the most
important thing for a dynamic hedger is not to focus too much on the uses
of the product and concentrate instead on their hedging techniques.

Example: Using a similar example as before—a one-month 100 call
knocking in at 98—Figure 19.4 shows an interesting graph. The first
time traders encounter the plot of the price with respect to the underly-
ing security, they know they are confronted with two piecewise options
of opposing deltas. ,

This can be confirmed by looking at the delta (see Figure 19.5).
In spite of its scary complexity, the knock-in call presents exactly the
same features as a knock-out: The jump in deltas is of equal magnitude,

0.8
0.6
0.4

0.2

Asset
94 96 98 100 102 104 106

Figure 19.4 Price of a knock-in call.
[End OCR]*

## Page 332

*[Image OCR]
318 Trading and Hedging Exotic Options

0.3
0.2
0.1

Asset
94 96 100 1 04 106
-0.1
-0.2

-0.3

Figure 19.5 Delta of a knock-in call.

although its move is in the opposite sign. This means that the knock-in
could complement the knock-out of same maturity and strike to create a
vanilla. If the reader took Figure 19.5 and subtracted Figure 19.2, he
would see the same graph.

At the barrier, the delta of the knock-in moves up from —.32 to .33,
which means that the operator has a gap delta of .65 to make up for. A
knock-out call, likewise, has a gap delta of a jump down of .65, which means
that the operator long one and long the other would have nothing to do.
This is the basis for the rule of the short barrier.

@ Long the barrier means that the operator benefits from the hitting of
the trigger, either by the decrease in liability (a knock-out) or by an in-
crease in wealth (a knock-in).

A short knock-out is long the barrier. A long knock-in is long the bar-
rier. A short knock-out benefits when the market hits the trigger as the con-
tingent liability extinguishes. A long knock-in benefits in the same manner,
as the triggering moves it from a conditional asset (the underlying option)
to an unconditional one.

By extension, the long knock-in and long knock-out are flat the barrier
and have exposure to the underlying option, hence the arbitrage relationship:

Knock-in (K, t, H) + Knock-out (K, t, H) = Vanilla (K, £)

Where K is the strike price, f time to expiration, and H the barrier level. A
long knock-in of one strike plus long knock-out of the same strike and ma-
turity are equal to a long vanilla. Look at the one-month 100 calls KO 98, 100
calls KI 98, and the 100 calls (Figure 19.6). Below the 98 point, the knock-in
exhibits the same behavior of the vanilla (for a good reason: It is now a
[End OCR]*

## Page 333

*[Image OCR]
Barrier Options (I) 319

Option Price

Asset
94 96 98 100102 104 106

Figure 19.6 Vanilla, knock-in, and knock-out.

vanilla). The knock-out will then be dead. Above 98, the knock-out behaves
like a vanilla, while the knock-in increasingly divorces itself from it.

Effects of Volatility

Figure 19.7 shows the price sensitivity of a knock-out to volatility levels. To
the right, as volatility increases, the bafrier gets flat vega while the vanilla
retains its effect. This is due to the barrier becoming closer in a nonlinear
way as volatility rises. Volatility remains linear for a vanilla. Eventually, the
barrier would dominate: As the barrier nears, by the sheer effect of volatil-
ity becoming higher, the option has a shorter and shorter time to live.

Price
3 Vanill
2.5 ,
2 Barrier
1.5
1
0.5

0.05 0.2 0.15 0.2 0.25 0.3
Volatility

Figure 19.7 Volatility effect on a 1m 100/98 knock-out and a vanilla call.
[End OCR]*

## Page 334

*[Image OCR]
320 Trading and Hedging Exotic Options

Option Price

Vanilla

3
2.5

2 Knock-in
1.5

1
0.5

Honea ~ Volatility

0.05 0.1 0.15 0.2 0.25 0.3
Figure 19.8 Knock-in and a vanilla.
Figure 19.8 shows the opposite effect for a knock-in. The knock-in will

have a convex vega while the knock-out will have a concave one. This is
meaningful for hedging purposes.

Trading Rule: Volatility brings the barrier closer in a nonlinear
way. This occurs because volatility is always non-linear for events

that are away from the center of the distribution (they affect the
out-of-the-money and in-the-money options).

Close to the barrier, the options have different vegas since the barrier
loses in convexity. Far from the barrier (the “other side”), the options would
start converging in behavior, including their sensitivity to volatility. It is
understandable that when the barrier is away from the asset price, the bar-
rier option converges in price to that of the vanilla. Operators thus test bar-
rier pricing systems by placing the barrier at 0 (or .00001) for a knock-out
call and a very high number for a knock-out put and verifying that the price
becomes that of a vanilla.

This leads to the vega linearity rule. ,

Trading Rule: The vegas emanating from the barrier part of a
barrier option are concave for the seller. The combination barrier

vega + vanilla vega will result in a concave vega for the long bar-
rier (see definition) and a convex vega for the short barrier. ,

The extension of the rule is that the concavity/convexity would start vanish-
ing as the trader moves away from the barrier into areas where the vanilla
dominates. Figures 19.9 and 19.10 show the vegas of the two instruments.
[End OCR]*

## Page 335

*[Image OCR]
Barrier Options (I) 321

Vega
10
8
6
4 Vanilla
2
Volatility

0.05 0.1 0.15 0.2 0.25 0.3

Figure 19.9 Vega of a down and in and a vanilla.

A more thorough analysis will involve changing the example into an
out-of-the-money option struck at 105 with the same trigger at 98. The ma-
turity is lengthened to take a 6-month option; volatility remains at 15.7%.

Figure 19.11 shows the shifts in vega from changing the asset price. As
the market moves away from the barrier, the vegas of the knock-out call
start resembling those of a vanilla.

Likewise, the vega of the knock-in call will start resembling those of the
vanilla the closer the trader is to the trigger (see Figure 19.12).

Adding the Drift: Complexity of the Forward Line

Risk Management Rule: A regular knock-out (i.e., triggered when
itis out of the money) will never have a delta higher than one if the

cash-forward line is flat. In an upward sloping forward curve, the
delta of the calls can be higher than one. In 4 downward sloping
curve, that of the put would be, too.

Vega
Vanilla
3
2.5
2
1.5
1
0.5
Volatility

0.05 0.1 0.15 0.2 0.25 0.3

Figure 19.10 Vega of a down and out and a vanilla.
[End OCR]*

## Page 336

*[Image OCR]
322 Trading and Hedging Exotic Options
Vega

25 Vanilla
20
15
10
5
Asset
95 100 105 110 115 120

Figure 19.11 Vegas for a knock-out and a vanilla (6 months to expiration, 100
calls, 98 barrier).

Extensions of the rule are:

¢ Near the barrier, options will retain a high value when the delta is
higher.

e Any delta higher than one will lead to a negative gamma somewhere
in the map. As this rise in delta is to compensate for a strong barrier
payoff, the move into an area where the underlying option dominates
would bring back the option irito normal proportion.

Figure 19.13 shows the effects of a rise in interest rates on the deltas.
Figure 19.14 shows the corresponding gammas. The gammas in Figure 19.14
are so extreme that there is no way to fit the trough in a conventional graph.
In the event of high interest rates, the position starts exhibiting an increase
in the third moment. Taking interest rates out permits a simpler analysis;
they will be incorporated again later.

Vega s

25
20
15
10

Asset
90 100 110 120

Figure 19.12 Vegas of a knock-in and a vanilla.
[End OCR]*

## Page 337

*[Image OCR]
Barrier Options (I) 323

Asset
95 100 105 110 115

Figure 19.13 Rising interest rates’ effect on the delta (strike 98.50, volatility
15.7%, time to expiration 2 years, barrier 91.50).

Risk Reversals

Comparing the vegas of the knock-in and knock-out with those of risk re-
versals would reveal some striking similarities: a vega that increases in one
direction and fades in the other. This prompted a rich literature of option
replication through risk reversals.

Put/Call Symmetry and the Hedging of Barrier Options

This section examines the technique of hedging the barrier option by sym-
metry as exposed in the common literature before warning the trader about
some of the pitfalls of the blind execution of the concept. The exact deriva-
tion and reasoning behind the formula is provided later in the section “The
Reflection Principle.”

Asset

Figure 19.14 The gammas corresponding to Figure 19.13.
[End OCR]*

## Page 338

*[Image OCR]
Option Wizard: Option Symmetry

Every European option has a symmetrical equivalent pair that matches its risk.
The symmetry is established vis-a-vis the forward, not the cash. That amount is
called “distance.” Assuming no skew (i.e., symmetrical volatility on both sides
of the forward}, the distance between puts and calls is computed as follows:

K is one strike price and F the forward. The symmetrical K’ (a put if K is a
call and a call if K is a put) assumes the equality.*

lf)

Ki = —

hence

Asset Price
A Put Cali Symmetry

Kt

100 : K'2
Kt

K2 |

The second, symmetrical strike is set in a way to make the forward become
the geometric average of the two strikes. The preceding chart shows the strikes
K1 and K2 and their corresponding symmetrical K’ 1 and Kk’ 2

In the absence of a skew, the strikes will be set so that the ratio of price
between put and calls will be the square root of the ratio of the strikes:

Put(K’ ) _ VK
Call(k) VK

The reader could, solely for entertainment, test the equality during a mo-
ment of (severe) boredom.

The symmetry should present the following attributes: A long one short the
other in proper ratio (of square root of strikes) should offer a zero gamma zero
vega. So the put-call symmetry in a delta-neutral structure should be neutral
moments until the third moment.

Experienced dynamic hedgers consider the notion of symmetry unstable.
There is an extreme difference between “upside” and “downside” strikes, ex-
treme to the point of causing meaningful behavioral differences if the market
rallies or sells.

*While put/call symmetry has been practiced by traders for a few decades, the first formal
treatment of the subject can be found in Carr (1994).

324
[End OCR]*

## Page 339

*[Image OCR]
Barrier Options (I) 325

This method has been deemed questionable by the author. Its value is
pedagogical as it gives some insight on the risks of barrier structures.

Assuming no drift, the barrier could be somewhat replicated using the
risk reversal. Assume further that a call knock-out is a call but not quite so.
So it could be a long call mitigated with a short out-of-the-money put. The
strike of the out-of-the-money put should be selected to be entirely sym-
metrical to the call at the barrier—in such a way and in such amounts that
the structure would be worth exactly 0 when the market reaches the barrier
level. Assume the trader is dealing with a 105 call KO 98 for 6 months. There
is a structure, such as an out-of-the-money put that, combined with the 105
call, should be able to replicate the barrier option every where except on ter-
mination, in which case the barrier needs to be unwound in the market.
Such a structure is only effective if the volatility does not change at all dur-
ing the operation.

Case 1: Knock-Out Call

Assuming the market trades at 100 and volatility at 15.7%.
Price of 6-month vanilla 105 Strike call: 2.35

Price of the 6-month KO (105/98) call: 1.06. The notation KO (strike,
outstrike) is used in the examples.

The symmetrical put should be of a strike price that, multiplied by the
strike of the call (i.e., 105), satisfies the geometric average 98. It will be
such that K X 105 = 987. K = 91.47

Price of the put symmetrical to the KO with strike 91.47: 1.20

The ratio of puts to calls should be V(105/91.47) = 1.0714 puts
for 1 call.

The trader should verify that at the barrier (i.e.; 98.00), the risk reversal
should be worth 0.

One additional constraint on the replication: The risk reversal
needs to be closed in the market after the barrier is hit. However, the re-
striction needs not to be stringent: A simple execution of a delta neutral-
ity would do immediately after the trigger is hit. The trader would then
have ample time to close the risk reversal. Table 19.2 plots the risk rever-
sal prices.

The thinking goes: The barrier is nothing but some form of risk rever-
sal. It should then react, in its vega, like the risk reversal. The reader
should retain the major benefits of this methodology while remaining
cautious about its pitfalls. An important result is that the barrier call
should be sold cheaper when there is a downside skew than when there is
a flat skew.
[End OCR]*

## Page 340

*[Image OCR]
326 Trading and Hedging Exotic Options

Table 19.2. The Method of Barrier Replication
Ratio: 1.0714

Risk Reversal

Asset (long 1 call short
Price KO Call Put 1.0714 puts)

91 0.00 0.48 4.07

92 0.00 0.59 3.62

93 0.00 0.73 3.20

94 0.00 0.88 2.82

95 0.00 1.06 2.47

96 0.00 1.26 2.16

97 0.00 1.49 1.88

98 0.00 1.74 1.63 0

99 0.53 2.03 1.40 0.53
100 1.06 2.35 1.21 1.06
101 1.59 2.70 1.03 1.59
102 2.14 3.08 0.88 2.14
103 2.69 3.49 0.75 2.69
104 3.27 3.94 0.63 3.27
105 3.85 4.42 0.53 3.85
106 4.46 4.93 0.45 4.46
107 5.08 5.48 0.37 5.08
108 5.72 6.05 0.31 5.72
109 6.39 6.66 — 0.26 6.39
110 7.07 7.30 0.21 7.07

Extending it to knock-in options should be straightforward because KI
call (105/98) + KO call (105/98) = Vanilla call (105). It follows that
K1(105/98) = Vanilla call (105) — KO(105/98). Since the KO call (105/98) was
constructed with: — 1.0714 put (91.47) + call (105) by the previous equality,
one gets:

KIC (105/98) = call (105) + 1.07 x put (91.47) — call (105)
= 1.07 puts (91.47)

Hence the replicating portfolio of the knock-in is the put (in some
ratio) until the barrier is hit and then the call after the barrier is hit.
The perfect replication would then entail executing the risk reversal at
the barrier by swapping the put into a call. Table 19.3 plots the knock-in
replication.

Case 2: Knock-In Call

The option constitutes a switch: At 98, when the passage from.one to the
other takes place, the risk reversal (in a ratio of 1 call to 1.0714 puts) is
[End OCR]*

## Page 341

*[Image OCR]
Table 19.3 Knock-in Replication

Asset KI
Price 105/98

91 0.48
92 0.59
93 0.73
94 0.88
95 1.06
96 1.26
97 1.49
98 1.74
99 1.50
100 1.29
101 1.10
102 0.94
103 0.79
104 0.67
105 0.56
106 0.47
107 0.39
108 0.33
109 0.27
0.22

110

Call
105

0.48
0.59
0.73
0.88
1.06
1.26
1.49
1.74
2.03
2.35
2.70
3.08
3.49
3.94
4.42
4.93
5.48
6.05
6.66
7.30

Ratio: 1.0714

Put
91.47

4.07
3.62
3.20
2.82
2.47
2.16
1.88
1.63
1.40
1.21
1.03
0.88
0.75
0.63
0.53
0.45
0.37
0.31
0.26
0.21

Barrier Options (I) 327

Replication
(long 1.0714 puts
above the barrier,
long 1 call below

the barrier)

0.48
0.59
0.73
0.88
1.06
1.26
1.49
1.74
1.50
1.29
1.10
0.94
0.79
0.67
0.56
0.47
0.39
0.33
0.27
0.22

worth exactly 0. So the knock-in represents a switch from one to the other at

that particular price.

Benefits of the Method. Viewing the barrier as an embedded risk reversal

presents three major results:

1. Pricing the skew. The slope of the skew should be kept in mind
when analyzing the barrier. This is a methodology similar to the bi-
nary option analysis. Later in the chapter, a methodology will be de-
rived to incorporate the skew and take into account its effect on

structures.

2. Pricing the volatility curve. The term structure of volatility matters
significantly, as will be seen later with the concept of stopping time.
It is interesting, however, to see that the decomposition of the barrier
leads to two options of a different duration on the volatility ladder. In
the case of the knock-out call, the put leg of the risk reversal reacts to
[End OCR]*

## Page 342

*[Image OCR]
328 Trading and Hedging Exotic Options

Option Wizard: The Skew Revisited

Volatility
Positive Skew
Higher Strikes have A higher Volatility
At-the-mone'
y Strike
Volatility

Negative Skew
igher Strikes have a lower Volatility

At-the-money Strike

one period while the call leg reacts to another. In the case of knock-
in options, the analysis is even more complex.

3. Hedging. A barrier option, like a risk reversal, is composed of two
polar components. At any time, only one of them dominates when
the market trades in its neighborhood. This holds for barriers (out of
the money when they knock out or in), not so for reverse barriers
where the large payoff is dominating. In the previous examples, the
call comes alive and dominates in the rally and weakens in the sell-
off with the put taking over. The option will therefore show a long
gamma in a rally and a short gamma ina sell-off. Would a risk rever-
sal hedge the structure? In most cases—but not all.

Pitfalls of the Method. The method of hedging through the risk reversal
is grounded on the following assumptions:
[End OCR]*

## Page 343

*[Image OCR]
Barrier Options (I) 329

¢ Stable (preferably constant) skew.

e Flat and constant forward curve (no drift or premium/discount in
the forward curve).

Making the skew unstable would cause a higher imprecision in the
hedge. Initially, whether the operator entered the trade at a positive or nega-
tive skew, the position would be locally matched, owing to the incorporation
of the skew price into it. Later on, the skew, being unstable, could shift and
eventually reverse, causing the trade to become divorced from the skew
slope. This is in addition to the possible ravages of time. This possible divorce
occurs because the skew hedge may not adequately match the duration of the
barrier. Even worse, there might not be a real duration susceptible of hedging
the skew owing to the instability of the stopping time.

Making the interest rates or carry for the asset extremely positive or
negative would invalidate the method as, by the rules expressed earlier in
the chapter, the barrier option might no longer behave exactly like a risk re-
versal above the barrier. The next section provides a digression on the
behavior of the gamma of a structure with a flat forward curve and that of a
structure with an exaggerated interest rate differential.

Gammas of Structures Compared with That of the Risk Reversal. Assum-
ing a flat forward curve, the knock-out gamma would resemble the right half
of a risk reversal (Figure 19.15).

The knock-in gamma would resemble the gamma of one call on one side
of the hill (the right side) and that of the put on the left side of the hill (Fig-
ure 19.16). The risk reversal is shown in Figure 19.17.

Introducing an interest rate differential provides a more complicated
situation, as depicted in Figure 19.18.

So what if a trader were to look at Mexico or an emerging market yield-
ing 45% interest rate differential? Assume he examines a one-year option,
knock-out call on the US dollar, put, on Mexico scaled at 100, with a knock-
out price at 105 and a strike at 102. He can safely use a 50% volatility. This

s

Option Wizard: Where There Is Skew, There Is Skew Instability

Traders notice that the existence of a skew is accompanied by a noticeable
shifting of the third moment (i.e., a strong fifth moment). Distributions with
strongly positive or negative third moment will have strongly positive or nega-
tive fifth moment and a higher fourth moment.

That a market needs a skew in its option volatility surface is a consequence
of asymmetry and a sign of structural instability.
[End OCR]*

## Page 344

*[Image OCR]
330 Trading and Hedging Exotic Options

0.02

0.015

0.005

Asset
90 95 100 105 110 115

Figure 19.15 Knock-out gamma.

0.025

0.015

: + Asset
90 95 100 105 110 115

Figure 19.16 Knock-in gamma.

Asset

Figure 19.17 Risk reversal gamma.

Asset

Figure 19.18 Knock-out gamma with a mild/positive drift.
[End OCR]*

## Page 345

*[Image OCR]
Barrier Options (I} 331

Asset

100 110 120 130 140 150

Figure 19.19 Deltas for a knock-out option on high yielding asset.

will help the reader understand why institutions chronically mismanage
their emerging market positions.

Figures 19.19 and 19.20 show that, in these conditions, the barrier op-
tion behaves like a full risk reversal, not like a half risk reversal as depicted
in Figure 19.15. In other words, its gamma changes in sign to the right of
the barrier, quite a confusing matter.

To price that “Mexico effect” on any path-dependent option, it is nec-
essary to use the Dupire-Derman-Kani techniques explained later in the
book. :

Barrier Decomposition under Skew Environments?

If in Case 1 (the knock-out call) the puts traded at a premium (i.e., the lower
strikes traded at a premium to the higher strikes), the operator could afford
to sell the barrier cheaper.

Asset ‘

Figure 19.20 Gamma of a high yielding position.
[End OCR]*

## Page 346

*[Image OCR]
332 Trading and Hedging Exotic Options

Option Wizard: Long Skew/Short Skew (Last Comment on the Skew)

Long skew means having the vega (and/or gamma) increase in the rally and de-
crease in the sell-off.

Short skew means having the vega (or gamma) decrease in the rally and in-
crease in the sell-off.

Long skew means benefiting from a positive third moment. Short skew
means benefiting from a negative third moment.

Trading Rule: Assuming a flat calendar, in a positive skew (higher

strikes at a premium), the knock-out call option will be cheaper than
in a flat world. The opposite applies to the negative skew. The oppo-
site also applies to the put (assuming the put is a regular knock-out,
not a reverse knock-out).

In general, if the trigger is higher than the strike, the knock-out option is
short the skew, and vice versa.

The skew procedure will be analyzed in two steps. The first one is to ig-
nore first exit time and simplify the world to look at the skew in a spatial di-
mension. After acquiring some familiarity with stopping time, it should be
possible to examine the skew in two dimensions.

Thus this analysis (for pedagogical reasons) takes place in the flat cal-
endar world, which means that options trade at the same volatility regard-
less of expiration.

The trader may attempt to quantify the skew for Case 1.

If the replicating portfolio is priced on the skew curve, the trader would
have to find the portfolio that can be worth 0 at the trigger time. That port-
folio includes selling a put further out of the money. Case 1 would not pro-
duce in a skew environment a call (K) — ratio < put(K’) = 0. Say that the
skew traded at a slope making the 91.47 trade at 1.7 volatility points over the
105 calls. The trader tries the replication of Case 1 (Table 19.4).

Obviously, it does not work. The risk reversal has a residual value of
(0.45) at the barrier. So he needs to scale put/call symmetry to introduce
some skew effect to it. He looks for the put that, combined with the call, sat-
isfies the replicating portfolio value of 0 at the barrier.

The put is 89.92, which he found by iteration; he solved for the strike of
the put that would be worth 1.74, the call value, when multiplied by the
square root of 105/strike. He also found it to trade at 2 volatility points over
the 105 call. The 89.92 put has a value of 1.62 at 98, which multiplied by
V105/89.92 becomes 1.74, the exact value of the call.
[End OCR]*

## Page 347

*[Image OCR]
Barrier Options (I) 333

Table 19.4 Case 1—Replication Attempt under Skew Environment

Ratio: 1.0714
kO Put 2 vols Risk Reversal
Asset (priced at Call higher (long 1 call short
Price flat vols) 15.7 vols than call 1.0714 puts)
94 0.00 0.88 3.29
95 0.00 1.06 2.94
96 0.00 1.26 2.61
97 0.00 1.49 2.31
98 0.00 1.74 2.05 (0.45)
99 0.53 2.03 1.80 (0.10)
100 1.06 2.35 1.58 0.65
101 1.59 2.70 1.38 1.22
102 2.14 3.08 1.21 1.78
103 2.69 3.49 1.05 2.36
104 3.27 3.94 0.91 2.96
105 3.85 4.42 0.79 3.58
106 4.46 4.93 0.68 4.20
107 5.08 5.48 0.58 4.85
108 5.72 6.05 0.50 5.51
109 6.39 6.66 0.43 6.19
110 7.07 7.30 0.36 6.91

The hint: Because the trader sold a put that is further out-of-the-money,
the entire structure would behave differently than before. The first lesson
in skew trading is learning that options that trade at different volatilities
exhibit unequal time decay. The trader solved for a replicating portfolio that
has no residual value today upon reaching the trigger. Tomorrow, however,
there will be a residual value at the trigger, given that the 89.92 put that he
is short, will decay at a higher speed than the 105 call.

Result: The positive skew drops the value of the barrier.

Table 19.5 verifies that the replicating portfolio above the barrier would
provide a residual value that would let a trader executing the trade against
a barrier option priced at fat skew eke out a profit.

This confirms any skew trader’s intuition: The put/call symmetry ob-
tained on a skew would decay favorably. The last column shows the differ-
ence. The difference will peak but reach 0 if the market does not move, as all
the options would then expire worthless.

The dependence of the skew premium on time to expiration shown in
Figures 19.21 and 19.22 shows traders the need for a more complex pricing
tool in the presence of a skew, any skew, in the market. Currently, there is
no known closed formula for that purpose. It also illustrates the need to use
a numerical method factoring the skew to price barrier options.
[End OCR]*

## Page 348

*[Image OCR]
334 Trading and Hedging Exotic Options

Table 19.5 Skew Risk Reversal through Time

3 Months Later Ratio: 1.0714
3-Month Net
Asset 3-Month Risk Net Replication

* Price kO Change Call Put Reversal Change P/L

—_ 0.10 1.86

92 — 0.14 1.53

93 — 0.19 1.24

94 _ 0.27 1.00

95 _— 0.36 0.79

96 _— 0.47 0.62

97 —_— 0.62 0.49
98 —_ 0.80 0.38 0.40 0.40 0.40
99 0.38 (0.15) 1.01 0.29 0.70 0.20 0.35
100 0.77 (0.29) 1.26 0.22 1.03 0.02 0.30
101 1.17 (0.42) 1.55 0.16 1.38 (0.16) 0.26
102 1.60 (0.54) 1.89 0.12 1.76 (0.31) 0.22
103 2.06 (0.64) 2.27 0.09 2.18 (0.44) 0.19
104 2.54 (0.72) 2.71 0.06 2.64 (0.56) 0.17
105 3.06 (0.79) 3.19 0.05 3.13 (0.64) 0.15
106 3.62 (0.84) 3.71 0.03 3.68 (0.71) 0.13
107 4.22 (0.86) 4.28 0.02 4.26 (0.75) 0.11
108 4.85 (0.87) 4.90 0.02 4.88 (0.77) 0.10
109 5.53 (0.86) 5.56 0.61 5.55 (0.77) 0.09
110 6.24 (0.83) 6.26 _ 6.26 (0.74) 0.08

%
0.01
0.005

~0.005
-0.01
-0.015
-0.02

Figure 19.21 Skew premium and the time to expiration.
[End OCR]*

## Page 349

*[Image OCR]
Barrier Options (I) 335

kD
LK >
TRS

LOT
Re

Figure 19.22. Skew premium in time and space.

The T decomposition formula is:

Barrier option _ Barrier option
(skew) ~ (No skew)

+ Expected value the residual
of the replicating portfolio/Conditional on hitting time

The reflection principle and the Girsanov theorem will augment insight
into this method.

The Reflection Principle

Another way of looking at the barrier symmetry is through the reflection
principle (see Figure 19.23). For a random walk, the paths from a point a to

b

-a

Figure 19.23 The reflection principle.
[End OCR]*

## Page 350

*[Image OCR]
336 Trading and Hedging Exotic Options

a point b that do not go through the origin (zero) are equal to the number of
paths from ~a to b.

This principle works on arithmetic paths only. To accommodate finan-
cial markets, it is necessary to trick it by using the logarithm of the prices.

This links back to the notion of option symmetry (see Figure 19.24).
Using the regular example of a market trading at 100, with a 98 knock-out
barrier, the number of paths from 100 to any point higher than it that does
not touch the barrier of 98 is equal to the number of paths from 987/100 to
the same point. Option symmetry can define the distance in logarithmic
terms to accommodate a geometric Brownian motion. So Log (100/98) =
Log (98/96.04).

Example. Binomial Tree (Simplified): From here we can proceed to an
intuitive understanding of the skew with the barrier option. This exam-
ple will also allow the reader to review binomial option pricing. Figure
19.25 shows the paths from 100 leading to in-the-money parts on the
map (higher than 100). It is assumed that there are 20 trading days (one
month) with movement of .99 per day, risk neutral, simply chosen for
pedagogical reasons (ignore lognormality for such an interval and as-
sume a “fudge,” to have the up-probability initially equal to the down-
probability*). Path 1 shows the number of paths from 100 leading to each
terminal node (i.e., the vanilla option). Path 2 shows the number of
paths from 96.04 to the same end nodes (i.e., the difference between the
vanilla and the barrier).

For the vanilla, option pricing is commonly done by multiplying the
intrinsic value on the end node by the risk-neutral probability of occur-
rence. The trader should be careful to consider such risk-neutrality as
an arbitrage derivation (i.e., an option = expected cost of gamma P/L as
explained in Module B). Here it is readily computed as the number of
paths at the node divided by the total number of possible paths times
the payoff.

103

98

96.04

Figure 19.24 Extension to option symmetry.
[End OCR]*

## Page 351

*[Image OCR]
Barrier Options (I) 337
intrinsic Path1 Path2 Net
109.90 9.90 i !
. 20 20
10792 <SSs 792 190 190
106.93 SESS 6.93 i140 1140
105.94 <S OS SSS 5.94 4845 1 4844
104.95 <> GOSS 4.95 15504 20 15484
103.96 LSSSSSISS 3.96 38760 190 38570
102.97 OS ISS IS 2.97 77520 1140 76380
101.98 SS SSS 1.98 125970 4845 121125
100.99 SLSSSSIS I 0.99 167960 15504 152456
rng SSSSSSESSSSS 1 een
99.01 SSeS CSS Se 17750
me SSS
96.04 SSSSSSSSSSSS 0.00 38760 184756
95.05 OS FS ESOS SS 0.00 15504 167960
se SSSSESSSESSS RTS“
92.08 SESESESESESS 0.00 190 38760
92.08 SOS SS SS <> 0.00 20 15504
SOS OS SSS 0.00 1 4845
92.08 SSS SSESLS
<<< 0.00 1140
<SSSSS tin 90
a 0.00 20
0.00 1

Figure 19.25 Sample paths and the reflection principle.

Expected value
at the node

Intrinsic at
the node

Number of paths leading to the node
Total number of possible paths

In Figure 19.25, the total number of paths is 27° = 1,048,576 paths. Ex-
pected payoff at any node is equal to intrinsic value times probability.
Probability equals number of possible paths/total number of paths. So
looking at the column Path 1, expected payoff at node 103.96 = 3.96 x

38760/1048576 = .1463.

The reader can also verify risk neutrality with no drift: Every possible
outcome weighted by its probability adds up to 100. It is called the “condi-
tional expectation of the future asset price at time t,” conditional means that
the expectation is based on information at the present.

The reader can also price the vanilla call struck at 100. It is equal to in-
trinsic times the probability of every level, as shown in Table 19.6.

The value of the vanilla option can be intuitively derived as the sum of
the expected values at every node. The precision increases with the higher

number of steps.

In Figure 19.25, Path 2 shows the number of paths that do not go
through the barrier using the reflection principle. The value of the barrier
option would therefore be the payoff (column labeled “Intrinsic”) times the

column net/total number of paths.
[End OCR]*

## Page 352

*[Image OCR]
338 Trading and Hedging Exotic Options

Table 19.6 Vanilla Valuation

Final Number Probability Expected
Price Intrinsic of Paths (%) Value
109.90 9.90 1 0.00 0.00001
108.91 8.91 20 0.00 0.00017
107.92 7.92 190 0.02 0.00144
106.93 6.93 1,140 0.11 0.00753
105.94 5.94 4,845 0.46 0.02745
104.95 4.95 15,504 1.48 0.07319
103.96 3.96 38,760 3.70 0.14638
102.97 2.97 77,520 7.39 0.21957
101.97 1.98 125,970 12.01 0.23787
100.99 0.99 167,960 16.02 0.15858
100.00 —_ 184,756 17.62 —_
Total 1,048,576 0.872
(all paths)

Repeating the exercise of Table 19.6, the trader derives the price of the
same option but with a knock-out barrier at 98 (see Table 19.7).

This leads to the equality: In the absence of skew and drift, a regular
knock-out option (that knocks out while out of the money), when the asset
trades at S, with a strike K and barrier H is equal to the vanilla option of
the same strike and maturity minus the same vanilla option priced with
the spot at S’ such that S’ is the symmetrical reciprocal of S centered

Table 19.7. Knock-Out Valuation

Path 2
Asset (Reflected Probability Expected
Price Intrinsic Path 1 Paths) Net (%) Value
109.9 9.9 1 1 0.0000 0.000
108.91 8.91 20 20 0.0000 0.000
107.92 7.92 190 190 0.0002 0.001
106.93 6.93 1140 1140 0.0011 0.008
105.94 5.94 4845 1 4844 0.0046 0.027
104.95 4.95 15504 20 15484 0.0148 0.073
103.96 3.96 38760 190 38570 0.0368 0.146
102.97 2.97 77520 1140 76380 0.0728 0.216
101.98 1.98 125970 4845 121125 0.1155 0.229
100.99 0.99 167960 15504 152456 0.1454. 0.144
100 0.00 184756 38760 145996 0.1392 —

Total 1048576 Total 0.844
[End OCR]*

## Page 353

*[Image OCR]
Barrier Options (I) 339

around H (S S' = H?). The amounts need to be adjusted by their adequate
ratio V(S/S').

Example: A 3-month knock-out call strike 100, spot 100, barrier at 98 is
equal to the price of the vanilla minus 100/98 times the same vanilla
priced at 96.04.

This can lead to the notion of risk reversal. Applying put/call parity, as
well as the change of numeraire method, a call with a strike price of 100,
priced with a 96.04 spot will be equal in price to the 96.04 put priced with
the spot at 100. This proves the skew rules shown earlier.

An extension of this notion is that, with a reverse knock-out option or a
regular knock-out that presents a rebate, the rules are more complicated as
one needs to add the payoff of the corresponding American binary.

Girsanov

It is easy to introduce the drift in the preceding framework to mix it with
the reflection principle. This would be done in assuming different probabil-
ities on exactly the same path structure, to compensate for the drift. The
path remaining the same, the probabilities are shifted to alter the payoff
without changing the paths.

Assume that over the same period, the spot needs to earn .99% per
month. It means that we expect that the resulting price after one month
would be 100.99.

The Girsanov theorem allows traders to change the probability of every
outcome upward or downward to allow for the expected return of 100.99. In
the example, the trader will have to shift the final probabilities. The module
{pricing contingent claims} presents a more formal presentation of the
method. Figure 19.26 shows the payoffs.

Repricing the option shows that the value of the package changed (Table
19.8).

s

Effect of Time on Knock-Out Options

With regular knock-outs, the effect of time is not truly perfidious, unlike
reverse knock-outs where it can be harrying. Time makes the underlying
option lose its value faster but at the same time increases discontinuity at
the barrier, which mitigates the effect: What would be knocked out will
have less and less value. Figure 19.27 shows this effect.

With knock-in options, however, the barrier has less effect in dominat-
ing the structure because there are no conflicting options in the structure.
Figure 19.28 illustrates this effect.
[End OCR]*

## Page 354

*[Image OCR]
340 Trading and Hedging Exotic Options

Intrinsic
109.90 3.90

108.91 << .
we SESE
105.94 CS OSS $.94

" oS "
104.95 SOS SS 4.95
103.96 SOS eSS SS 3.96
ms SSSA iS
100.99 CSLSOSSSCS EA 0.99
100.00. << CSSSOSCSSS 2 0.00
99.01 > SSSR 0.00
98.02 Coe OO Oa x 0.00
7 oe
96.04 oo OO ex 0.00
95.08 CS OSS SOS 0.00
94.06 <> SOS OS SSS OS OSS 0.00
93,07 SOSSSLSISLSES 0.08
SSSSESESESS

° OO x

92.08 SSSOSS Se
SSS te

<< .
<> 0.00
<< 0.00

Total

Path

]

20
190
1140
4845
15504
38760
77520
125970
167960
184756
167960
125970
771520
38760
15504
4845
1140
190
20

i

1048576

Path2

1
20
190
1140
4845
15504
38760
77320
125970
167960
184756
167960
125970
77520
38760
15504
4845
1140
190
20

{

1048576

Net

1

20

190
1140
4844
15484
38570
76380
124125
132456

Figure 19.26 The reflection principle with changes of probability measure.

First Exit Time and Its Risk-Neutral Expectation

M@ The expected first exit time (or stopping time) is the time an asset price
is expected to cross a given point in the market, conditional on an
expiration date. Module G (Pricing Contingent Claims) provides the cal-
culation for both the distribution of the first exit time and its expectation.

A barrier option unlike a European binary needs to be hedged using
instruments that have a similar duration. This necessitates the knowledge
at all times of the expected first exit time, which is in this case the time
when the barrier is hit. The difference between the nominal expiration
and the expected stopping time will correspond to the dominance of the
barrier in the structure. When the barrier is weak (by exerting a small in-
fluence on the risks of the structure), the expected exit time is equal to the
nominal duration. When the barrier is strong, it will be considerably

shorter.
[End OCR]*

## Page 355

*[Image OCR]
Barrier Options (I) 341

Table 19.8 Knock-Out Value after Change of Probabilities

Asset Probability Expected
Price Intrinsic Path 1 Path 2 Net (%) Value
110.89 10.89 1 1 0.0000 0.000
109.9 9.90 20 20 0.0000 0.000
108.91 8.91 190 190 0.0002 0.002
107.92 7.92 1140 1140 0.0011 0.009
106.93 6.93 4845 1 4844 0.0046 0.032
105.94 5.94 15504 20 15484 0.0148 0.088
104.95 4.95 38760 190 38570 0.0368 0.182
103.96 3.96 77520 1140 76380 0.0728 0.288
102.97 2.97 125970 4845 121125 0.1155 0.343
101.98 1.98 167960 15504 152456 0.1454 0.288
100.99 0.99 184756 38760 145996 0.1392 0.138

100.00 0.00 1048576 Total 1.369

The expected first exit time can thus be represented as shown in Fig-
ure 19.29.

As volatility increases, the option shortens and the barrier starts domi-
nating. When volatility is low, the underlying option dominates and the ex-
pected first exit time becomes the maturity of the vanilla.

Trading Rule: The certainty that first exit time is shorter than or

equal to the nominal maturity makes any nonarbitrage static repli-
cation of a barrier with combination of vanillas imperfect.

Price

FNM WwW F&F UW

Time (years)
0.2 0.4 0.6 0.8 1

Figure 19.27 Knock-out option and a vanilla with time (100 calls, barrier 98).
[End OCR]*

## Page 356

*[Image OCR]
342 Trading and Hedging Exotic Options

Price

PN WwW bP UW

Time (years)
0.2 0.4 0.6 0.8 1

Figure 19.28 Knock-in option and a vanilla with time (100 calls, barrier 98).

Many simplistic methods of option hedging have been used, such as the sta-
tic replication with a risk reversal presented earlier. Their not taking into
account the fuzziness of the duration can be downright dangerous.

Many software vendors took the simplified homoskedastic (constant
volatility) pricing tools available in the early academic literature and
dumped them on the traders. Using only one volatility when there is a
possibility of the barrier expiring earlier leads to overpricing when the
volatility term structure is upward sloping and underpricing in a back-
wardized volatility curve.

0.25
0.2
5 0.15

= ’
- 0.1

=

0.05
0

on NO OOlUtTLCUCOMDCUCN CUCOUOUCUCUCOOUUCUOULUMDUCClCUD

a » Ss N © st © MN a Qo N o w ©

0) @ oO o oO Oo oO oO ~ - r - -

5 - - hoa r- r - - rc - ad -

Asset Price

Figure 19.29 Risk-neutral expected first exit time: 3-month, 100 call/knock-out
98, volatility 5, 15, and 20%.
[End OCR]*

## Page 357

*[Image OCR]
Barrier Options (I) 343

The consequences of the risk management rule are:

e Any vega hedge against the barrier needs to take into account the
first exit time for the matching of the barrier payoff.

* Hedging the forward exposure stemming from the barrier needs to
be done carefully. First exit can move, and may lengthen or shorten.
A lengthening of the expected first exit would necessitate a length-
ening of the maturity of the hedge. This rebalancing can consume
massive transaction costs. It is therefore recommended to hedge at a
shorter maturity than the nominal one.

¢ First exit time depends on the volatility. Where there is a volatility
curve in the market, positive and negative, the first exit time would
lengthen or shorten.

Issues in Pricing Barrier Options

Since the first exit time is uncertain, it is necessary to know where it is lo-
cated on the term structure for pricing purposes. Currently, most systems
offered to traders present an erroneous system of pricing. There are two
methods of fudging matters:

1. The single volatility fudge. It is less computer intensive but lacks accu-
racy. This author recommends using it only to patch up the flaws in
the pricing formulas. It provides a half-answer.

2. The pricing on a forward-forward curve (the Dupire method). It is more
computer intensive but is generally accurate in determining the real
value and the appropriate hedges. Some people complain about the
burdening implementation and convergence difficulties.

The Single Volatility Fudge

* Step 1. Determine the implied spot volatility curve in the market.

Month 1 2 3 4 5 6 7 8 9
Volatility 9.9 10.25 10.5 10.7. = 10.85 10.95 11.1 11.2 11.3

e Step 2. Find the expected stopping time. Notice that it is a distribu-
tion, not a fixed point in time. Assume the expected stopping time
was derived to be close to 6 months. Therefore, interpolating on Fig-
ure 19.30, the volatility to use is 10.95, which should define that of
the knock-out option.

° Step 3. Price the knock-out option using a synthetic risk reversal.
The major option (the 105 call in the previous example) needs to be
[End OCR]*

## Page 358

*[Image OCR]
344 Trading and Hedging Exotic Options

11.5 5

114
|
10.5 +

10 4

Implied Volatility

©
nn
1 +—--+-

I

+-~ ~ + 4 + + 4
N oO wv wo ie o an

wo

Months to Expiration

Expected Stopping Time

Figure 19.30 Stopping time: Volatility term structure.

priced as an option maturing with the nominal value. The second op-
tion needs to be priced using the parameters of first exit time. Thus a
skew dimension would be added to the risk reversal.

° Step 4. Find the forward-forward between the stopping time and
the end date. Using the formula of Chapter 9, the reader will obtain
FFwd (6,9) = 12%. This number is useful in determining the value of
the knock-in of the same nominal duration.

Finally, the fact that the stopping time is not fixed but has a distribution
makes it necessary to specify the following rule.

Risk Management Rule: The hedges in forward assets that corre-
spond to an option position need to be scattered to match the dis-

tribution of first exit time.

A More Accurate Method:
The Dupire-Derman-Kani Technique

In 1992, Dupire presented the seminal notion of a two-dimensional volatility
process. Between one point and another in space (spot/strike) and time (to
expiration) there is a volatility implied by both the skew and the forward-
forward volatility in the market.> This was followed by breakthrough re-
search by Dupire (1993) and Derman and Kani (1994). They arrived at a
numerical technique that breaks up time into smaller sections between the
nodes, thus constituting a calendar-based time-speed function (one can even
[End OCR]*

## Page 359

*[Image OCR]
Barrier Options (I) 345

call it diagonal). The method entails the use of some form of flexible nodes in
the binomial (or multinomial) tree.

The method of pricing on the tree has serious implications for any path-
dependent structure. Since a barrier depends not just on one single volatility
but on that of every node, the pricing can more precisely establish the value
of the barrier using a two-dimensional volatility surface. In effect, both the
forward-forward problem and the estimation of the stopping time can thus
be skirted.

Independently, in Berkeley, Marc Rubinstein discovered the notion of
implied binomial tree and studied the lateral volatility between two asset
prices derived from option prices. In other words, as Dupire and Derman
and Kani explored the process between a point in space (asset price) and
time (S, t) and (S + AS, t + At), Rubinstein was independently discovering
the implied process betwen (S, t) and (S + AS, f).

A few issues, however, have marred the implementation of these tech-
niques. The primary problem is that a high number of steps is needed to
make the binomial tree (or trinomial) converge to a stable price. This could
be time consuming and many efforts at finding numerical techniques to
speed up the calculations and deliver feasible techniques of option pricing
are under way.

To this date, no commercially available system implements such
a method.

Additional Pricing Complexity: The Variance Ratios

Option traders hedge themselves at a given frequency, say an observable
time interval. They concern themselves with the volatility at the time scale
of their intervention. A trader who hedges himself every few minutes will
not be concerned with tick volatility. A barrier option, however, responds to
the infinitesimal volatility. The two can be markedly different as shown in
Chapter 6.

Barrier options are priced on the assumption of the distribution of the
minimum of the maximum of a Brownian motion that follows the same
volatility as the Brownian motion itself. The most intuitive way to look at
the barrier is to assume the risk-neutral (alas) discounted probabilistic pay-
off of a vanilla option (i.e., S — K for the call and K — S for the put) condi-
tional on the stopping time not being attained during that period. Putting
that condition on the regular option represents the difference from a
vanilla.

The condition of the stopping time, however, may react to the distribu-
tion of the minimum of maximum. The analysis of the Parkinson number (see
Chapter 6) hinted that the distribution of the extremes on a particular day
may be markedly different from what would be expected from a log-normal
(or even normal as it does not matter much within one day) distribution.
[End OCR]*

## Page 360

*[Image OCR]
346 Trading and Hedging Exotic Options

Trading Rule: Ina mean-reverting market, the barrier component

of a barrier option is overpriced. In a trending market, it is under-
priced.

The explanation for this effect is that the knock-out option is more likely to
be knocked out when the market has a high level of intraday negative auto-
correlation and options are priced on the close-to-close volatility. Remem-
ber that the structure is triggered off a price printing on the screen, not a
terminal state for the asset.

Since this is not an econometrics course but general guidelines for op-
tion hedging, the trader must decide on whether the market follows one
particular process or another, or whether one particular condition is more
permanent than the other.

From this author’s experience, every market, including markets in-
volved in a trend (one-week sampled volatility is higher than daily samples)
will exhibit mean reversion within a one-day framework. This is due to the
execution slippage of the participants, to widening of the bid-offer spread
in the event of higher volatility than normal, and several other factors.

Another way to examine it is to use the Parkinson number rule:

Trading Rule: If the close-to-close volatility is lower than the

Parkinson number, the barrier component is underpriced by the
regular method. Otherwise, it could be considered overpriced.

EXERCISE: ADDING THE PUTS

-As an exercise, the reader could redo the previous examples by pricing a 98
put knocking out at 105 and the same one knocking in at 105. The author sub-
scribes to the notion that a put is nothing but a synthetic call and prefers to
spend appropriate time on the differences where they might arise.
[End OCR]*

## Page 361

*[Image OCR]
Chapter 2 0

Barrier Options (IT)

The only thing I like about a knock-out option is that it knocks out.
N. Zeidan

This chapter focuses on reverse barrier options and double barrier options.'

REVERSE BARRIER OPTIONS

Reverse Knock-Out Options

@ A reverse knock-out is a barrier option that terminates in the money,
therefore causing a large discrepancy ketween the values on each side of
the barrier.

Reverse knock-outs are almost decomposable into regular knock-outs plus
an American bet. However, one ends up with the opposite gap delta. Be-
cause a large number of horror stories attend their risk management, it will
prove helpful to study reverse knock-outs separately.

A three-month call is generally worth 3.11% of face at 15.7% volatility.
A reverse knock-out that terminates at 108 with no rebate will be worth
55% only. Strangely enough, the reverse knock-out call will sometimes pre-
sent a negative delta (its price will decrease when the market rallies), which
newcomers to the exotics club find difficult to conceive since ordinary op-
tions gain in value when their intrinsic value increases. In this case, how-
ever, the intrinsic value disappears at one point so the price decreases as the
market rallies close to the barrier. The conflicting pull between the barrier
and the regular underlying vanilla will show a serious conflict throughout
the life of the instrument.

Reverse knockouts appear to be priced abnormally low, with their value
generally moving inversely to time: they decay negatively. Such a feature
makes them an attractive selling instrument to corporates. The same fea-
ture creates nightmares for option traders and sleepless night fraught with
cold sweats, as will be shown later.

347
[End OCR]*

## Page 362

*[Image OCR]
348 Trading and Hedging Exotic Options

The best intuitive way to view reverse knock-out and knock-in options
is to consider them American binary options with a high payout. Since the
payoff is in the form of a deep-in-the-money option with a high intrinsic
and little time value, they look very little like an option when close to the
barrier and more like a bet.

Figures 20.1 and 20.2 show the unusual behavior of a reverse knock-out,
terminating at 10, with 6 months on the clock. The reader will detect a
mildly short gamma for the owner, as well as a short delta.

Figure 20.3 shows the nasty effect of time near the barrier, reminiscent
of that of an American binary.

Figure 20.4 shows the effect of time on a reverse KO (100,105) and 20.5
shows the difference with an American binary: below 100 the trade is short
gamma, therefore will earn time decay (even when there is no carry).

Case Study: The Knock-Out Box

The following real-life trade (only slightly modified) took place in the dol-
lar against French Franc.

As the dollar dropped steadily against the European currencies during the
1990s, many French corporate treasuries found themselves under water on
their speculative positions. It had been fashionable in Parisian circles to call
the turn of the U.S. currency as travelers were invoking economic equilib-
rium theories and other notions that made the dollar look underpriced in
their eyes. It was rare for this author to enter conversations with Parisians
without hearing purchasing-parity comparisons such as the cost of a com-
pact disk sold by Chinatown peddlers and the exact same item sold at

70 +

Vanilla

-30 aL

Figure 20.1 Comparison of the P/L of a vanilla call and a reverse knock-out.
[End OCR]*

## Page 363

*[Image OCR]
Barrier Options (II) 349

Vanilla

-20 £ Asset

Figure 20.2 Comparison of the delta of a vanilla call and a reverse knock-out.

Roissy Airport. So macroeconomists-traders went long the dollar massively
around 5.60 to the French franc, some 12% higher than the rest of the story.
Many corporates without real marks-to-market rules found it easy to carry
a position without the discipline of a stringent stop-loss. However, they
needed at some point to swallow the bullet and show a loss.

As the franc was hovering around 5.00, some derivatives house presented
them with the following gimmick: Here is a chance for you to be “hedged”

6.00 ]

5.00 +

4.00 '

3.00 +

ni
3
o

Figure 20.3 Effect of time on the price of a reverse knock-out.
[End OCR]*

## Page 364

*[Image OCR]
350 Trading and Hedging Exotic Options

2-25

101.5
103.5
105.5 ieee
107.5
109.5

Figure 20.4 Reverse KO (100, 105) at different levels of volatility.

against your dollar exposure. So they sold them the quack medicine de-
scribed below.

Customer is long dollars at 5.60 in the amount of $100 Mil. Therefore
the losses are around $12 Mil since spot USD-FRF is at the round 5.00. The
quack medicine would allow the customer to get it back.

Customer buys a 5.60 knock-out put on the dollar (call on the FF) for a
relatively small price. It knocks out, however, at 4.85 (not very far from
where spot is) and has an entire year to go. The cost of the trade is very
small: around .8% of the face value of the position, which corresponds to
$800,000 (including the bank’s profit margin). At expiration, if the option
has not been knocked out, the trader would make back his $12 Mil as

W0.1-0.2
—0-0.1

wm -0.1-0

fq -0.2--0.1
@ -0.3--0.2
-0.4--0.3
0 -0.5--0.4
w -0.6--0.5
@ -0.7--0.6

+
Theta

2 16
_
oO
es g
rT +

107.5;

Asset

Figure 20.5 A theta that changes in sign.
[End OCR]*

## Page 365

*[Image OCR]
Barrier Options (II) 351

Option Wizard: How Bad Salespeople Fool Their Customers

There is an old trick bad salespeople use to fool their customers into trades
that generate profits for their employer and keep them in the game: taking ad-
vantage of their statistical misperception. It is easier to fool someone with a
distributional confusion.

Covered Writes, for example, allowed brokers to kill the upside of their cus-
tomers and generate steady flows of commissions. A customer is long the stock
and sells calls against it. If the customer loses money on the calls, the broker
could argue that the customer made money on the stock. Should the customer
lose money on the stock, the broker would not fail to point out that the customer
was spared further losses thanks to the income brought by the calls.

Until customers gained in sophistication, covered writes were the best
game in town. As customers started understanding a little more statistics, the
game moved to the fancier exotic options payoffs. The distributional confusion
moved to the notion of path dependency.

The Wall Street firms that survive are those who cater to their clients’
needs.

he would be able to go short the dollar back at 5.60 where the entire matter
would be closed. The trader would be able to start a new life as a dollar bull
again if he so wishes. Great trade. It sold very well, like a Le Monde issue
on election day.

“Fair value” for the trade using a Black-Scholes-Merton type valuation
(one constant volatility) shows the trade to be worth .63% at 11.5% volatility,
6% domestic U.S. rates, and 7% French rates.

The profile of the trade by itself is first examined. Figure 20.6 shows the
value of a 5.60 put on the dollar (call on the FRF) KO at 4.85, with one year

USD-FRF

4.75 5.25 5.5 5.75 6 6.25

Figure 20.6 The salvation trade. Premium in mil dollars.
[End OCR]*

## Page 366

*[Image OCR]
352 Trading and Hedging Exotic Options

P/L

USD-FRF
-5 5.75 6 6.25

Figure 20.7 The position without the salvation trade.

to go and a volatility of 11.5%. The price is expressed in total dollar value of
the premium for a $100 million dollar trade.

Figure 20.7 shows the initial position without the salvation trade while
Figure 20.8 shows the position after. A drop in the bucket, one may say.

It is enlightening to compare it to the position with the Salvation Trade.

So today the trade does not do much for the hapless Frenchman. Tomor-
row, however, there might be some hope. Look at the profile around expira-
tion in Figure 20.8. A zone of break-even can be seen between 4.85 and 5.60
on the overall package (less the premium spent on the Salvation trade).

Trading Rule: Some institutions fool the gullible investors by tak-
ing advantage of their lack of knowledge of probability theory, and

truncating the distribution of profits. The investors, particularly
when under pressure, tend to believe in Santa Claus.

P/L y

USD-FRF

55.75 6 6.25

Figure 20.8 The position with the salvation trade.
[End OCR]*

## Page 367

*[Image OCR]
Barrier Options (II) 353

P/L

USD-FRF
6 6.25

5.25 5.5 5.75

Figure 20.9 But there is some hope.

It does not improve the customer’s condition to see the graph in Figure
20.9. But the broker now has a story to tell in cases of success and failure.

Further analysis of the trade shows that the reason such protection is
cheap is precisely because its first exit time is very short. The trade is en-
tirely dominated by the barrier, so the odds of the 5.60 put remaining an op-
tion are very low.

With time, the value of the option would increase, and, should (for some
reason) the market become stationary, the structure would look like a real
protection. Such an event has an infinitely small probability of surviving the
next year. The salesperson would show the customer the pictures in Figures
20.9 and 20.10 but would fail to tell him that such a picture represents not the
expiration P/L like a regular option but shows the P/L if USD-FRF never
trades below than 4.850001. There is a massive difference between a path in
dependent expiration P/L and a path-dependent one. Most customers have
not been trained nor are they used to looking at such graphs properly.

Finally, the positive time decay from the trade, provided the USD-FRF
remained totally frozen, shows in Figure 20.11.

Option Wizard: Fooled by the Expiration Profile

The expiration profile is a great way to delude a potential customer into enter-
ing a good trade (good for the seller). It has already been established that the
statistical sense of people is rather weak, even in the examination of a regular
risk profile.

When looking at a path-dependent profile “at expiration,” the customer
may not realize that he has mentally correct for the payoff probabilities. This is
difficult when the report is made to look like a static payoff with a small men-
tion at the bottom.
[End OCR]*

## Page 368

*[Image OCR]
354 Trading and Hedging Exotic Options

0.6
Time(Years)

5.4
USD-FRF

5.6

Figure 20.10 Profile until expiration.

Note: This example does not take into account that a trader who lost $12
million would incur interest charges on the losses. In fact, some trades es-
cape such treatment but are penalized by their not getting any interest on
their profits.

Some customers after a while could come back and tell their salesper-
son that perhaps the trade might not always work in all circumstances.
They may be paying $800,000 (actually the bank might be charging them
up to $1,000,000 for the trade under the illusion that every penny above
Black-Scholes-Merton value is profits) just to buy a butterfly. So here is
the answer.

Making the Trade Debit-Free. By entering the trade for no debit, or per-
haps a small credit, the customer might complain very little in the event of
being knocked-out of the trade. So the new idea is to perform, in addition to
the “Great Salvation,” a covered write.

The customer can thus, at no cost, keep performing the trade until the
day, way down the road, when he would break even. A statistician can prob-
ably estimate that day to be a few hundred years away.

P/L (Spot USD-FRF 5.00)

Time (Years)
“2i\ 0.20.40.60.8 1
-6
-8
-10
-12

Figure 20.11 Static (no spot movement) view of the P/L.
[End OCR]*

## Page 369

*[Image OCR]
Barrier Options (II) 355

To turn the trade into a debit trade, the dealer could make the gullible
customer-underwater sell some 5.60 dollar calls for some premium. This
is the late 1990s, however, and this kind of trade belongs to the preexotic
structures of earlier days. Moreover, the customer could then decompose
the trade by himself, trade it on his own and deprive them of the profits. Fi-
nally, then, should the customer sell a 5.60 call on the dollar, the structure
could cause a problem in the event of the KO being terminated and the
other leg staying naked to prevent the renewal of the structure.

What to do then? The package could include a call knock-out that
would vanish also at 4.85. The no-debit rule can be achieved. In addition,
the trader could even get some cushioning income from it. The terms
would be as follows:

Length: one year. Customer buys 5.60 put on dollar KO 4.85, sells 5.60
call on dollar KO 4.85 for no premium. Should, at the end of the 1.5
years, the dollar stay above 4.75, the customer would get back his $12
million losses.

However, should the structure be knocked out, which means that the
dollar dropped below 4.85, the trader would have his global position lose
even more money. That, the derivatives house would not mention.

What in effect the customer was led to do was truncating his distribu-
tion, cutting off any hope of making money back against the possibility of
keep trying the trade forever, until he can spend one entire year without
hitting a barrier.

Figure 20.12 shows the profile of the entire package at expiration. The
principal difference between Figure 20.12 and Figure 20.9 is that, above the
barrier of 485, Figure 20.12 shows perfect breakeven.

P/L

USD-FRF
§.255.55.75 6 6.25

-25

Figure 20.12. The customer truncated his distribution.
[End OCR]*

## Page 370

*[Image OCR]
356 Trading and Hedging Exotic Options

Hedging Reverse Knock-Outs: A Graphical Case Study

Reverse knock-outs are principally difficult to hedge. The following case
study will examine the reverse knock-out from the vantage point of the
trader and illustrate the possible hedges. The presentation will go through
the Greeks and examine their relevance—or lack of it.

The trade will be examined from the standpoint of the seller of the op-
tion, and it will be hedged step by step.

The trader sells $100 million of one-year 5.60 dollar-put FRF call for
$800,000, with a KO feature at 4.85, only 3% away from the market.

The Deltas. The operator is short calls but an inexperienced trader might
be surprised to see the deltas move from positive to negative. This occurs
because close to the barrier the option structure weakens and is dominated
by the bet. The trade becomes cheaper as the market approaches the barrier.
Close to the other end, however, the trade becomes cheaper again as it re-
sembles a long put in a rally. In between the two poles, the trade could be
somewhat confusing.

The delta is expressed in percentage of face value (in dollars, not cur-
rencies). Figure 20.13 shows the changes in the delta of the structure. The
reader can see the drop from the cliff at the barrier price of 4.75.

As time, however, moves closer, the deltas would give a little bit of
heartbeat. Look at it with one day to go (Figure 20.14).

The delta one day from expiration reaches 2000 level, which, for, a $100
million position, will represent serious deltas. The face value of the position
is actually larger for a dollar-based person since it corresponds to being
short the equivalent of 100 million FRF at 5.6, 560 million FRF, which would
be by the two-country paradox constant FRF but at 4.85 would turn

Delta

USD-FRF

6 6.25

Figure 20.13 Delta one year to expiration.
[End OCR]*

## Page 371

*[Image OCR]
Barrier Options (II) 357

Delta

20

USD-FRF
4.75 5.255.55.75 6 6.25

Figure 20.14 Deltas on expiration day (in number of times the position).

into 100 x 5.60/4.85 = $115.4 million. So 2000 deltas would be 20 times
115.4 = $231 billion. The intrinsic value of the option would, by the same
token, be worth $15.5 million, quite an impressive time decay given that a
year before, it is worth just between $600,000 and $800,000, the price at
which the traders sold it.

Figure 20.15 shows the accelerated changes in the value of the option
from the passage of time. In trading, acceleration means instability. Such
instability increases at the barrier, provided the structure does not knock-
out before the end of its life. So this scenario, while perhaps being the most
beastly, is no doubt the least likely of all. There are very small chances for
the market spending one year near a barrier without touching it.

The reader will find in Figure 20.16 two possible scenarios: The change
in value at a USD-FRF rate of 4.76 (close to the barrier) and at 5.55 (away

{Option Value}

PNY WwW PP WwW

Days
50 100 150 200 250 300 350

Figure 20.15 Changes in value with time at the trigger.
[End OCR]*

## Page 372

*[Image OCR]
358 Trading and Hedging Exotic Options

FPN WwW FPF YN KH ~I

4.76
Days to Expiration

50 100 150 200 250 300 350

Figure 20.16 Value through its life at exchange rate 4.76 and exchange rate 5.55.

from the barrier). At 4.76 the trade undergoes a harrying negative time decay
for the seller. At 5.55, however, the trade decays positively. The trader exe-
cutes his delta. He sold the KO put and is therefore long dollars (short FRF).
He sells dollars, buys FRF against the position in 19% of the dollar value (at
5.00 the dollar value is 100 < 5.60/5.00 = $112 mil so 19% is $21.2 mil).

Figures 20.17 through 20.23 will examine the various risks assuming /
dynamic hedging. Figures 20.17 through 20.20 will assume delta neutrality
and the remaining figures will assume, in addition to delta neutrality, vega
neutrality. Given the difference between the local behavior of the Greeks
and their global behavior, the graphs will examine the sensitivity both
using a microscope (for a local sensitivity assuming a small change in the
parameters) and a telescope (assuming large changes).

Figure 20.17 shows the delta sensitivity through the microscope, in the
neighborhood of a USD-FRF of 5.00. It looks like simple long gamma, with
the “V” Shape of the P/L. Figure 20.18 shows the same sensitivity at a

USD-FRF

Figure 20.17 Hedged P/L in the microscope.
[End OCR]*

## Page 373

*[Image OCR]
Barrier Options (II) 359

USD-FRF

Figure 20.18 Hedged P/L in larger increments.

larger scale, with a jump at the barrier. Figure 20.18 shows the real need to
liquidate the delta hedge at the barrier, one of the problems in managing a
barrier option.

Gamma and Vega. Figure 20.19 shows the disturbing gamma, as it is cen-
tered in the neighborhood of 5.30 USD-FRF. Figure 20.20 shows now that, 5
week from expiration, the gamma peaks at a level close to the barrier, and
turns negative in the neighborhood of 5.35. The difference between Figure
20.19 and Figure 20.20 represents the reason traders are scared of reverse
knock-out options, since they cannot be trusted to retain any characteristic
for long.

Before looking at the necessary vega hedges, examine the stability of
the measure. Figures 20.21 and 20.22 show the position at current spot. Fig-
ure 20.21 shows the volatility sensitivity of the position with 1 year left and
Figure 20.22 shows the same sensitivity one month from expiration. It is

USD-~FRF

5.2 5.4 5.6 5.8

Figure 20.19 Gamma of the position (short side — ).
[End OCR]*

## Page 374

*[Image OCR]
360 Trading and Hedging Exotic Options

USD-FRF

Figure 20.20 Gamma 5 weeks before expiration.

Volatility

Figure 20.21 Vega sensitivity one year to go.

Volatility

Figure 20.22 Vega sensitivity one month to go.
[End OCR]*

## Page 375

*[Image OCR]
Barrier Options (I) 361

Volatility

Figure 20.23 Vega neutral P/L.

noticeable that both the vegas and the vega convexity increase with time to
expiration, in the neighborhood of the barrier.

It is worthy mentioning that the vega sign is negative: the owner is short
vega. The trader who sold the option will have to sell vega to offset a long
volatility position.

By looking at the gamma sheet for immediate moves, the trader would
notice that the gamma peaks at around 5.30. So will the vega (for a single op-
tion gamma and vega peak at the same level). So the trader will sell an option
that peaks. around that level. The trader selects a 3-month option struck
around USD-FRF of 5.30, which leaves him with a vega-neutral position.
Figure 20.23, which is a graph of the vega combination, shows that vega neu-
trality will be very transitory for the next move.

It is never pleasant for traders to be in a situation where the vega of the
position increases when volatility comes down and decreases in a rally. It
means that vega neutrality will be a sure loser.

Figure 20.24 shows the vega of the structure unhedged across ‘spot
across volatility levels. Comparing Figure 20.24 to Figure 20.23 would cer-
tainly show it to be more favorable. Perhaps, after all, it would have been
better to avoid selling the 3 month options as a hedge.

Finally, Figure 20.25 shows the hedged exposure across the exchange
rate moves, to show that not just the volatility changes can bother the vega,
but spot moves as well.

To conclude the case study, perhaps it would have been better to leave
the trade without a vega hedge. As a trader the author prefers a dangerous
trade left roughly unhedged than accurately mishedged.
[End OCR]*

## Page 376

*[Image OCR]
362 Trading and Hedging Exotic Options

Volatility

Figure 20.24 Vega of the structure.

USD-FRF

Figure 20.25 A dangerous vega hedge.

DousLE BARRIER OPTIONS

Double barrier options are structures that expire when either of the two
barriers is touched by the market. In structures on a single asset (that isa
one-dimensional option), there can only be two barriers, one higher than
the market and another one lower. With multidimensional options, there
could only be twice as many barriers as dimensions.

Figure 20.26 compares the value of a double barrier with that of a simple
barrier. In both cases the option is a 1 year call struck at 100. The single bar-
rier has an upper outstrike at 110 while the double barrier has an upper out-
strike at 110 and a lower one at 90.

The following are basic rules about double barriers:

¢ Typically, the first exit time is so short that the underlying option’s
characteristics will be of little relevance. At best, the structure would

e
[End OCR]*

## Page 377

*[Image OCR]
Barrier Options (II) 363

Asset Price
85 90 95 100 105 110

Figure 20.26 Value comparison between a reverse and a double barrier.

look like a reverse knockout. At the worst, it would seem to want to
expire imminently.

* Consequently, the best way to study them is to examine American
double bets.

Rebate

The rebate is a simple American binary option that pays a certain sum on
the termination of the barrier. A knock-out option with a rebate is therefore
a straight knock-out plus an American bet of the same strike.

Suppose that a reverse knock-out call of strike 100 and knock-out trig-
ger 109 paid $9 on termination. The structure will simply be additive: one
knock-out call 100 trigger 109 plus one $9 face value bet, which means that
the bet pays $9 if the trigger 109 is reached at any point during its life.

Exercise: Adding the Knock-In

The reader should try, as an exercise, to repeat the study with knock-in op-
tions, using the methodology of the comparison between regular knock-out
and regular knock-in options.

Alternative Barrier Options

These are also called outbarrier options. They belong to both the path-de-
pendent and correlation families of exotic options.

A trader may have an option on one asset with a barrier, typically a
knock-out, in another one.

Example: A SCUD, a second currency underlying, is an option on
one asset in one currency with a barrier struck on the exchange rate.

=
[End OCR]*

## Page 378

*[Image OCR]
364 Trading and Hedging Exotic Options

Examples are a Deutsche mark cap with a dollar-mark barrier, or a
Nikei Index put option with a knock-out on a lower dollar yen.

The SCUD assumes that the customer no longer needs the option
protection when the price of the asset moves in such a way as to create a
profit from the currency alone.

Say that a dollar-based fund is long Japanese stocks. A lower dollar,
stronger yen would be sufficient to create a profit. He can simply pur-
chase a protection against an adverse movement in the Japanese market
that terminates if the currency appreciates. Should the option termi-
nate, the investor can then unwind the entire package for a profit.

Another example is for a U.S.-based investor to borrow in short-term
DM and roll over his borrowing periodically. He can buy a cap for his
borrowing but with a knock-out should the DM depreciate. In the event
of his being knocked out, he will be able to satisfy his short-term DM
liability by buying back the currency for a profit.

The following are key trading issues:

e The correlation between assets makes the alternative barrier resem-
ble a regular one. Independence makes the barrier become a bet on
the other option.

e The structure will end up having two correlated deltas, one positive
in the asset A and one, on the opposing side of the barrier, for the
asset B in which the barrier is struck (see Chapter 22, “Multiasset
Options”). : .

The Exploding Option

An exploding option is one that presents a certain payout if and when a cer-
tain price is reached between the initiation and maturity. It is said then to
“explode.” .

Assume that the underlying is trading at 100. A 102 call “exploding” at
107 pays 5 points when 107 is touched. It offers an alternative to a call
spread or a put spread since the owner of the exploding option can thus
take profits without having to incur transaction costs. The drawback of
such an option is that should the owner change his mind about the posi-
tion before expiration, the spread costs would be greater than those of a
regular option.

An exploding option is equivalent to a reverse knock-out with a rebate
equal to the exploding payoff at the time of the termination of the option.

Figures 20.27 and 20.28 show the difference between a regular call and
an exploding option. The regular call is struck at 100 while the exploding
option is struck at 100 and terminates at 105. Figure 20.27 considers a ter-
mination payoff of 2 while Figure 20.28 considers a payoff of 10.
[End OCR]*

## Page 379

*[Image OCR]
Barrier Options (II) 365

Regular Call

Exploding

+—+—_+—_+—_ +++ +—t+-—+- +} ++ +_ +++ + +++ 4+
oe 2 ee eee ee gee 2 ek 2
2 2 Oo oo 8 a

8S 88 5S FS eS Ssssqs

Figure 20.27 The 100 call that pays $2 on termination compared with the regular
100 call with 90 days until expiration.

Exploding

Regular Call

+ +
” © 2
| aod oOo oa

8 Qo o ao

+
7)
3

Figure 20.28 The 100 call that pays $10 on termination compared to the regular
100 call with 90 days until expiration.

Capped Index Option

A capped index option, called CAPs, is best described as a reverse knock-out
that pays a rebate on termination equal to the difference between strike and
outstrike. It is made to appear to be a call spread or put spread that offers a
payout equal to the difference between the two strikes. One other intricacy is
[End OCR]*

## Page 380

*[Image OCR]
366 ©Trading and Hedging Exotic Options

that unlike most knock-outs that are triggered if touched, this one needs the
settlement price (in some cases) to be at or through the second strike.

CAPs are different from exploding options in that they are only termi-
nated at the official closing price of the asset while the exploding option is
commonly knocked out at any time. It is then said to explode. As explained
in Chapter 18, the at-settlement options act like American binaries (i.e., mo-
notonously long or short vega) except close to expiration day when they be-
come European and can thus present mixed gamma features.

Example: Figure 20.29 compares the P/L value of the CAPs 100-105
(i.e., a 100/105 capped call spread) with that of a regular call struck at
100 with the market at 100. The initial purchase price of the CAP is 2.72
compared with 3.11 for a regular call. Assume that volatility is 15.7%
annualized.

CAPs are commonly difficult to trade compared with other structures be-
cause they only terminate if the underlying assets go through the settlement
price, regardless of what can happen during the trading day. While this fea-
ture may appear of small consequence to the nonhedging user, it carries some
significant risks or some significant opportunities to the hedger.

In a traditional barrier option, the trader knows exactly when to un-
wind his hedge. When a market maker is short the CAP and has a position
in the underlying asset against it, he would work an order close to his CAP
price. In an ordinary barrier, he would get out of his hedge at close to the

Regular Call

CAPs s

Figure 20.29 Comparison between a CAPs 100-105 and a 105 call 90 days to
expiration.
[End OCR]*

## Page 381

*[Image OCR]
Barrier Options (II) 367

Option Wizard: The Legality of the Triggering

Operators often disagree on what constitutes an incontrovertible termination
of a barrier structure. If a trader owns a knock-out option that was terminated
according to the counterparty, he would like to see clear signs that the market
earnestly traded there without manipulation.

Manipulation is possible (some say likely) in illiquid over-the-counter mar-
kets. Listed exchanges do have a procedure in place for preventing two traders
from consummating a trade at a fictitious price. A trade cannot be recorded at
a price lower than existing bids or higher than existing offers. Such protection
is not afforded by the over-the-counter traders who do not have a centralized
place to exchange contracts.

The need for such protection caused banks and dealers to establish elabo-
rate contracts specifying that the barrier-triggering conditions would be met only
with a proof of the trade on more than an independent forum (i.e., more than
one broker). Some firms are usually unhappy about such terms. They consider
that a market if touched condition would be preferable, as in the heat of the bat-
tle, it would be impossible to predict whether the option was knocked out if the
operator has to wait to gather information on the number of times it traded and
where. The operator needs certainty. Knowing immediately whether he has
been knocked out is of the utmost significance. Having to wait for the proof
might be costly as the information lag causes slippage.

The best solution is to require the bank short the knock-out (or long the
knock-in) to leave a stop-loss with a designated list of third parties (consid-
ered reputable from both sides). Thus the party long the barrier (i.e., having to
place a stop-loss order against the market move as opposed to having a limit
order) can satisfy the proof that the barrier was triggered with a simple confir-
mation slip.

knock-out price and close his risk. In this case, he does not know the exact
price, which can, ironically give rise to some opportunity.

Take the example of the 100/105 CAPs structure. The trader is short
$100 million and long $48 million of the underlying asset. At the 105
knock-out price, the market maker has his delta change from 48% to zero
instantaneously. But should that happen during the day, the trader will
not know whether to cover entirely by selling back the $48 million. He
could sell a portion of it and wait to see whether the market goes higher, in
which case he could sell the balance at a higher price if he gains confi-
dence that the structure would be terminated. Should the market head
lower through the barrier he could buy back the hedge for a profit, and
hope for a repeat of the operation. So in this example, the trader has an in-
valuable option in his hand.
[End OCR]*

## Page 382

*[Image OCR]
368 Trading and Hedging Exotic Options

Option Wizard: Barriers and Price Manipulation

There are famous cases of traders defending their barriers, as with the well
publicized case of the Venezuelan par bonds. One firm owning knock-in puts
allegedly tried to buy massive amounts of the bond to satisfy the trigger condi-
tions while the party short the puts left massive sell orders to prevent the price
from going through the limit. A flash screen update caused one party to con-
sider that the option was knocked in while the counterparty thought otherwise.
The Securities and Exchange Commission (SEC) was alerted by the abnormal
trading volume and its concentration among few players.

The lesson from such an event is that in an illiquid market the seller of
the barrier option (i.e., the party benefiting from the hitting and therefore
leaving a stop order) can easily manipulate markets with quantities small
enough to make such manipulation worthwhile. The payoff at the barrier is
large enough to make the risks of the attempt of triggering it very small in
comparison. According to the contamination principle, this represents no
less than a free option.

There are similar situations with cash-settled options where small buying a
few minutes prior to expiration can lead to massive profits, without the penalty
of the slippage in getting out of the trade.

Lesson: One should never get involved in a knock-out option in an illiquid ,
market unless one gets proper compensation for the risk of liquidity holes.

The example shows that the owner of the structure would be at a disad-
vantage if he engaged in dynamic hedging.

READING A Risk MANAGEMENT REPORT

This section includes the following material:
s

¢ A report on the way Greeks are commonly disclosed to trading risk
managers.

e¢ Examples of the Greeks as treated (or mistreated) by a conventional
risk management system.

¢ An example of the pin risks at the barrier.

¢ Anopportunity through the gap report to examine the difference be-
tween mined and clear markets.

The discussion of the barrier option risk precedes the introduction of a
standard risk management report (i-e., the kind used by banks’ risk
[End OCR]*

## Page 383

*[Image OCR]
Barrier Options (II) 369

management departments in their daily monitoring of the risks). Because
few institutions today have positions without barriers, the author believes
in holistic risk management: Have all possible positions on one underlying
cumulated in one report.

Risk managers get a hint of the risks of a particular position through
common scenario analysis reports showing the sensitivity of the options to
all manners of parameters. The principal such report is the spot sensitivity.
It outlines the risks of P/L stemming from the changes in asset price.

Note that some firms generate “shadow” reports (i.e., showing shadow
gamma, as defined in Chapter 8) while others, unfortunately, limit them-
selves to straight reports.

The risk reports of straight option positions are simple to look at as the
risk manager does not need to make any assumption about the potential ac-
tions of traders. With barrier options, managers need to consider the risks
of execution of hedges on hitting the strike.

Two reports are included in this section: One takes into account the liq-
uidation of a barrier option position (Figure 20.30), and the other does not
(Figure 20.31). The example will illustrate some of the concepts we covered
and will show the variety of risks faced by risk managers in dealing with
explosive products. For the sake of simplicity, the position only contains
one delta-neutral position in a flat drift flat term structures of volatility
environment.

It is assumed that the Syldavian currency trades at SYD 1.00 to the dol-
lar (the report will show 100 as per reporting convention). The trader has
only one position in his inventory: He is short 100 million dollar equivalent
of the 3-month call on SYD put on the USD struck at 105.00 with a knock-
out feature at 97.00. Typically, traders can carry positions with thousands
of strikes, but there is no loss of generality from studying an exposure with
a small number of positions.

The report provides the following information:

¢ Variable: Spot. The report changes the spot and does not allow for
any change in other parameters. It is therefore a “straight ganima”
sensitivity, incorporating no correlation structure between the spot
levels and either the interest rates or volatility.

¢ Pair. The unit the trader is concerned with, assuming such position
is entirely stand-alone.

¢ P/L cur. The currency in which the P/L is computed. It matters to
have it in the numeraire currency and to have the delta expressed in
countercurrency (as with the two-country paradox).

¢ The dates. The evolution of the P/L between the dates 8/13/96 and
8/13/96 (in this case the same day). The horizon date could have
been longer.
[End OCR]*

## Page 384

*[Image OCR]
370 Trading and Hedging Exotic Options

Credit Bank of Syldavia - New York Branch

Trader = M.

Dept 332
SPOT From = 8/13/96_-—s Incr (STD) 0.25 Portfolio 1038
SYD_USD to 8/13/96 Min 94.5 Audited ON
USD Center 100.00 Max 105.5 Barrier Analytic
SYD Rounding 000 American Cox-Ross
SYD Barrier OFF Iterations 90 O/E

10

Delta Gamma
(5,301) (20,660) (4,400)
(4,798) (19,570) (4,360)
(4,332) (18,490) = (4,320)
(3,874) (17,420) (4,280)
(3,452) (16,340) (4,320)
(3,056) (15,280) (4,240)
(2,688) (14,230) (4,200)
(2,345) (13,190) (4,160)
(2,028) (12,160) (4,120)
(1,737) (11,150) (4,040)
(1,470) (10,160) (3,960)
(1,228) (9,190) (3,880)
(1,010) (8,240) (3,800)
(816) (7,320) (3,680)
(644) (6,420) (3,600)
(494) (5,550) (3,480)
(366) (4,710) (3,360)
(259) (3,900) (3,240)
(171) (3,130) (3,080)
(102) (2,380) (3,000)
(51) (1,680) (2,800)
(17) (1,020) (2,640)
- (390) (2,520)
2 180 (2,280)
(8) 720 (2,160)
(33) 1,210 (1,960)
(69) 1,650 (1,760)
(115) 2,050 (1,600)
(171) 2,400 (1,400)
(235) 2,700 (1,200)
(306) 2,940 (960)
(382) 3,140 (800)
(462) 3,270 (520)
(545) 3,360 (360)
(630) 35,000 = Error!!!
(1,505) 35,000
(2,379) 35,000
(3,254) 35,000
(4,129) 35,000
(5,004) 35,000
(5,879) 35,000
(6,754) 35,000
(7,629) 35,000
(8,504) 35,000
(9,379) 35,000

NOL A AMANO OO oO
R8SSFL2AXRSRES
’

©

a

a ow

no

mo

eooecoopo ooo o4
e270 D OOOO OOO omg
oooooDODOOOOg
cocoon oo ooo ORs

Figure 20.30 Report 1 with the barrier stop-loss off.
[End OCR]*

## Page 385

*[Image OCR]
Barrier Options (II) 371

Credit Bank of Syldavia - New York Branch

Report 6 Trader M.

Dept 332
Variable SPOT From = 8/13/96—s sIner (STD) 0.25 Portfollo 103B
Palr SYD_USD to 8/13/96 Min 94.5 Audited ON
P&L cur USD Center 100.00 Max 105.5 Barrier Analytic
Delta cur SYD Rounding 000 American Cox-Ross

SYD Barrier STOP ON iterations 90 O/E

Delta Gamma e
105.50 (5,301) (20,660) (4,400)  -1851 172 -1309 1392

105.25 (4,798) (19,570) (4,360) -1833 170 -1276 1356
105.00 (4,332) (18,490) (4,320) ~-1819 169 1243 ©1319
104.75 (3,874) (17,420) (4,280) ~—--1803 167 -1209 1282
104.50 (3,452) (16,340) (4,320) ~—--1783 166 1175 1245
104.25 (3,056) (15,280) (4,240) -——-1761 163 1141 1208
104 (2,688) (14,280) (4,200) _-1736 160 -1106 1170
103.75 (2,345) (13,180) (4,160) ~—--1709 158 1072. 1133
103.5 (2,028) (12,160) (4,120) -1678 154 -1037 1095
103.25 (1,737) (19,150) (4,040) ~—_-1645 151 1002 1057
103.00 (1,470) (10,160) (3,960) ~—--1609 147 967 1019
102.75 (1,228) (9,190) (3,880) -1570 144 -932 981
102.50 (1,010) (8,240) (3,800) -1529 140 896 943
102.25 (816) (7,320) (3,680) -1485 135 860 905
102.00 (644) (6,420) (3,600) —-1389 131 824 866
101.75 (494) (5,550) (3,480) -1337 125 788 828
101.5 (366) (4,710) (3,360) —-1283 121 752 789
101.25 (259) (3,900) (3,240) -1226 116 715-750
101 (171) (3,130) (3,080) -1167 110 678 710
100.75 (102) (2,380) (3,000) -1106 105 640 671
100.50 (51) (1,680) (2,800) -1043 99 603 631
100.25 (17) (1,020) (2,640) —--977 93 564 590
100.00 - (390) (2,520) + -910 87 526 549
99.75 2 180 (2,280) —--841 80 -487 508
99.50 (8) 720 (2,160) -770 74 -447 466
99.25 (33) 1210 (1,960) -697 67 -406 424
99.00 (69) 1,650 (1,760) -624 61 -365 381
98.75 (115) 2,050 (1,600) -548 54 323 337
98.50 (171) 2,400 (1,400) -472 47 -280 292
96.25 (235) 2,700 (1,200) -395 40 237 246
98 (306) 2,940 (960) -317 33 192 199
97.75 (382) 3,140 (800)  -238 25 146 = 152
97.5 (462) 3,270 (620) 17 -98 102
97.25 (545) 3,360 (360) 9 -50 52
97.00 (630) 0 Errorit! 0 0 0
96.75 (630) 0 0 0 0 0
96.50 (630) 0 0 0 9 0
96.25 (630) 0 0 0 0 0
96.00 (630) 0 0 0 0 0
95.75 (630) 0 0 0 0 0
95.50 (630) 0 0 0 0 0
95.25 (630) 0 0 0 0 0
95.00 (630) 0 0 0 0 0
(630) 0 0 0 Oo 0
(630) 0 0 0 ) 0

Figure 20.31 Report 2 with the barrier stop-loss on.
[End OCR]*

## Page 386

*[Image OCR]
372 Trading and Hedging Exotic Options

¢ Incr. In standard deviations (self-explanatory). The report is for .25
standard deviations.

¢ Rounding. 000 means that amounts are in thousands.

e Audited ON. A flag that shows the risk management report was is-
sued from a database audited by the controllers. It signifies that the re-
port produced can be read but cannot be altered by the trader, on a
read-only mode. The trader cannot (in theory) hide positions and take
risks that do not show on the report.

¢ Barrier: Analytic. The pricing of the barrier options is performed
on a simple model without any attempt at using more sophisticated
tools. Analytic means that the system uses a closed-form solution.

¢ American: Cox-Ross (the Rubinstein in Cox-Ross-Rubinstein is unde-
servedly missing because of the narrowness of the field on the report).
Indicates that a binomial model is implemented for American options.
The field Iterations item shows 90 O/E, which means that the tree will
have 90 iterations but will flip between odd and even as a precision
optimizing trick.

¢ Barrier: ON or OFF. This is where the beef is: Some reports assume
that the trader has a stop-loss to unwind the hedge at the barrier (in
this case, 35 mil SYD at 97.00), and Report 2 with the flag ON makes
such assumption. Typically, if the stop-loss is triggered, the trader
has a documented proof that the market traded at such level as re-
quired by the contract and finished off the option. So it appears safe
that, conditional upon a firm order in the market, the risk manager
would be interested in seeing the rest of the position post-stop-loss.

Matters are not that easy, however: What if the barrier stop-loss was not
executed at the desired price? There is the nefarious slippage, as defined in
Chapter 4. Beyond that, there might be the more severe gap risk: the discon-
tinuous moves as the market can snap through the barrier, thus causing liqui-
dation at lower than 97.00. Every point lower would cost $3.5 million dollars
in execution costs. This shows the need for another risk report, the gap risk.
Traders assume that markets in foreign exchange trade on a 24-hour basis, but
there is the risk of the discontinuity caused by the weekend. In addition,
there are political announcements that can cause such gap risks: The news
that the Syldavian king has some cardiovascular problems can hit the screens,
thus causing a large and sudden gapping drop in the currency.

Many risk managers suggested to the author that the stop-loss be
placed before the official knock-out trigger (slightly higher than a lower
barrier and slightly lower than a higher barrier). This is a bad idea as it cre-
ates a negative gamma should the barrier not be reached (a phenomenon
commonly called a “bounce” by traders trying to defend their barriers).
[End OCR]*

## Page 387

*[Image OCR]
Barrier Options (II) 373

At another level, banks would be choked with risk and stopped from
trading unless they eliminated the triggered barriers from the risk reports.
Unlike the regular short option gamma, the barrier is a one-time hit, not a
permanent danger. Assume that the trader sold the $3.5 million SYD on
a stop-loss basis. The trade confirmation would be the proof that the barrier
was knocked out. Should the market rally higher or continue dropping, the
position would not exist at all.

The best solution for a risk manager is to have Report 2 with the barrier
stop-loss ON (not Report 1 with the stop-loss OFF) as well as a gap report
showing the potential execution risks.

This risk management issue illustrates why dynamic hedgers often go
through the difficult task of trying to explain the P/L swings to their boss.
The “I thought you were hedged” is an answer by managers unfamiliar
with the product (or too burned out to try to learn about it) and the best
remedy is the gap risk report.

Risk Management Rule: Traders and managers need to make
a clear distinction on a report between the contractually path-

dependent positions (i.e., barriers) and the securities that are path
dependent owing to dynamic hedging (vanilla options).

A brief explanation of the report columns follows:

¢ P&L (column 2) shows the profits from the origin (Center on top
shows 100). It assumes no shadow gamma, no changes in volatility.
To compute the shadow gamma resulting P/L, one needs to make as-
sumptions about the volatility behavior conditional on a move. As-
sume the trader thinks that at 102 the volatility would be (for the
period concerned) one point higher, an extremely safe assumption.
Being that he is short modified vega (the same as simple vega) in
the amount of $139,000 per point the P/L at 102 would then ,be
— 644,000 —139,000 = -$783,000 at such spot level.

¢ Gamma in this report is conventional except that it blows up at the
barrier where it becomes a Dirac. The report shows an error sign.

* The modified vega (expressed in 100 not 1,000 for reasons of conven-
tion) assumes curve shifts that are weighted in 3-month equivalent.
Whatever amount is seen corresponds to the 3-month move with the
previous month a function of it, according to some empirical observed
weightings. It assumes the move to be one point in the 3-month, .5 in
the 1-year, and so on. Because the model is using an analytical model
in place of the Dupire-Derman-Kani, it is safe to assume that the
[End OCR]*

## Page 388

*[Image OCR]
374 Trading and Hedging Exotic Options

modification will not be very precise for a barrier since the barrier re-
acts to the entire curve.

¢ Rhod shows the exposure to parallel shifts in the numeraire cur-
rency and Rhof shows the shifts in the countercurrency. Both are un-
weighted, an extremely bad idea. Rhod is higher than rhof because it
includes the exposure to the carry of the premium.

Gaps and Gap Reports

M The gap report shows and quantifies the possible execution risks of the
gap delta (the deltas that need to be unloaded on a stop-loss on the ter-
mination of the structure). These are one-time slippage costs.

The second column of the gap report shown in Figure 20.32 shows the
amount of deltas that need to be sold at 97.00. The system at the Credit Bank
of Syldavia estimated the gap risk to be on weekdays $634,000 (Column 3),
which corresponds to a gap of .20 (the risk of having the order filled on av-
erage at 96.80) and three times such amount when the market gaps over the
weekend (Column 4).

The magnitude of the gap risk is a function of the number of options and
will remain independent of the duration of the portfolio or stopping time.
Moreover, the market has a tendency; after a period of high volatility, of
cleaning up some range in the market of stop-losses. This often has some se-
rious impact on volatility. It suggests the notion of mined markets.

m@ A mined market is one that has many knock-outs and knock-in gap-
delta orders placed in it and will therefore experience a high level of
“whipping” and mean reverting volatility around these prices. A
cleared market is one that has no such orders in it.

As the open interest in barrier options builds up, so will the gap orders by
operators. Traders monitor the recent high and low (say for one month), and
when such difference remains narrow for a few weeks, they can expect the

Credit Bank of Syldavia - New York Branch

Report 7 Trader M.
Pair SYD_USD Dept 332
SYD Date 8/13/96 Portfolio 103B
Audited ON

Spot DeltaGap Week-da Week-end

31,700

97.00 634,000 2,536,000

Figure 20.32 The gap report for the previous position.
[End OCR]*

## Page 389

*[Image OCR]
Barrier Options (II) 375

Lightly Mined
Market

103

100

Densely
Mined Market

Figure 20.33 Mined markets.

97

buildup of gap delta orders below the low and above the high. As soon as
the market enters a “mined” zone, a chain of liquidity holes is triggered.
The market then clears out. This phenomenon is called the “driftwood ef-
fect” as all barrier option orders will be pushed out to the recent high and
low range (see Figure 20.33).

Risk Management Rule: The expected slippage costs of gap deltas
need to be increased when the risk manager has adequate informa-

tion to conclude that their level is located in a densely mined market
range.
[End OCR]*

## Page 390

*[Image OCR]
Chapter 2 1

Compound, Choosers, and
Higher Order Options

Hiring an option trader with P/L volatility is the most common manifestation of a
compound option.

Compound options are options that deliver another option. They are partic-
ularly present in the extendible option variety and other structures where
the owner of the structure has, somewhere along the way, the possibility of
extending his choice against the payment of a fee.

Compound options are extremely sensitive to higher derivatives with re-
spect to spot, particularly the fourth moment of the distribution, which
traders call volatility of volatility, in addition to the second moment with re-
spect to volatility (somewhat related). They are commonly mispriced because
of that effect. More than regular options, they depend on the thickness of the
tails of the distribution. This makes a constant volatility type model, written
in early days, dangerous to use. This chapter also gives warning of a more
complex modeling technique that includes the volatility of volatility owing to
the difficulties in both its estimation and its association with the underlying
price behavior.

At the time of this writing, there is no known formula nor publicly
available method to price compound options.correctly (using stochastic
volatility), other than a few numerical techniques this author dabbled with.
So in this presentation, the author will use what poor formula is available
and try to work around it by estimating the appropriate markup.

@ A compound option, or second-order option, is an option on a European
option. It consists in the right to buy or sell a European option (called un-
derlying option) of a given strike and maturity for a predetermined price.

A third-order option is an option on a compound option, and so forth. A
compound option has one final strike price and one final expiration date
(the equivalent to a vanilla). In addition, it will have intermediate strikes
and intermediate dates. So a second-order option will be

Second-order option (®,K,, ®,. K,. tite)

11’ “ final” “final’” 1’ final

376
[End OCR]*

## Page 391

*[Image OCR]
Compound, Choosers, and Higher Order Options 377

tei Asset Price

© 7
o -
—

-

102
105

Figure 21.1 Compound option: OTM calls on 1-year call, 16% volatility.

with K,,.,, and #,,,, the strike price and the time left to expiration of the
vanilla. K, and £, will be the intermediate dates K, the strike price at which
one has the right to buy or sell the option (Ke nav and f,. ,,) at time f.. Also, t,
needs to be shorter than the final date. Note that every strike price will
need the further specification of whether it is a put or a call (i.e., the right to
buy or the right to sell), which is done through the indicator ®.

A higher order option will have the following specifications:

nth order option (®,K,,... DK Dena tina’ a
witht, <t,<...<t <t..

Example: Figure 21.1 shows the prices of a call on call struck for $1, at
16% volatility, assuming a flat forward curve.

Figure 21.2 displays an out-of-the-money put on a call. Figure 21.3 illus-
trates the straddle.

Figure 21.2 Compound option: Put on 1-year call.
[End OCR]*

## Page 392

*[Image OCR]
378 Trading and Hedging Exotic Options

25 -

180D

——- 45D
— 15D

om OO OD
on © Mm

117 3

o
N
-

~~ ot
- ++
- =

105
108 +

Bm N
© 3S
2

Figure 21.3 Straddle on 1-year call.

VEGA CONVEXITY: THE Costs OF DYNAMIC HEDGING

Vega convexity can be best illustrated with the following case study (see
Table 21.1). The underlying option has tenure of 2 years, struck at 100, with
a flat forward, and 15% volatility. Assume for simplicity that the volatility
curve is flat. As the price of the option is initially 8%, the trader buys a 16%
out-of-the-money call on the call and hedges it vega neutral in the underly-
ing security.

To simplify, assume that vega neutrality can be done without the adjust-
ment factors. Adjustment factors, or weightings, add nothing to the exam-
ple. Figure 21.4 shows the instability of vega neutrality. What a trader gets

1.000
0.800
0.600
= 0.400
0.200
0.000
0.208 G5 = A AW
Oo oO [o>] oO
Volatility

Figure 21.4 Long the higher moments: Compound versus vanilla long a 6M call
on ATM 2-year call, 16% strike.
[End OCR]*

## Page 393

*[Image OCR]
Compound, Choosers, and Higher Order Options 379

Table 21.1 Vega Effect
vol Price Vanilla P/L Compound P/L P/L

0.04 0.000
0.05 0.000 2.062 —1.065 0.997
0.06 0.000 1.874 — 1.065 0.809
0.07 0.015 1.687 — 1.050 0.637
0.08 0.061 1.499 ~ 1.004 0.495
0.09 0.129 1.312 —0.936 0.376
0.1 0.210 1.125 —0.855 0.270
0.11 0.305 0.937 —0.760 0.177
0.12 0.420 0.750 —0.645 0.105
0.13 0.564 0.562 —0.501 0.061
0.14 0.710 0.375 —0.355 0.020
0.15 0.878 0.187 —0.187 0.000
0.16 1.065 0.000 0.000 0.000
0.17 1.256 —0.187 0.190 0.003
0.18 1.449 —0.375 0.384 0.009
0.19 1.646 —0.562 0.581 0.019
0.2 1.878 —0.750 0.813 0.063
0.21 2.126 —0.937 1.061 0.124
0.22 2.378 ~1.125 1.313 0.188
0.23 2.634 —1.312 1.569 0.257
0.24 2.893 —1.499 1.828 0.329
0.25 3.157 —1.687 . 2.092 0.405
0.26 3.424 —1.874 2.359 0.485
0.27 3.696 —2.062 2.631 0.569
0.28 3.971 —2.249 2.906 0.657
0.29 4.255 —2.437 3.190 0.754

0.3 4.611 2.624 3.546 0.922

5
when he goes vega neutral against a long compound option is a straddle on
volatility.

Usrs OF COMPOUND OPTIONS: HEDGING BARRIER VEGA
Because barrier options present extreme vega concavity as a package, it

seems sounder to execute the hedge of the vega with the use of a compound,
should these be available at no higher cost in the market.
[End OCR]*

## Page 394

*[Image OCR]
380 Trading and Hedging Exotic Options

CHOOSER OPTIONS

@ Chooser options are options that could turn into either a put or a call at
some predetermined time.

They have the following specifications:

Simple Chooser (K,f.

intermediate’ tinal)

The more complex “Gutspin” Chooser (K, Kit nterme diate’ final)
with K the strike price, t,. mediate {he time until the owner of the structure
would have to decide on whether he prefers to own the put or the call. The
“gutspin” chooser allows for two strike prices, and the owner would cer-
tainly choose the one that is the furthest in the money.

A European chooser option that has the intermediate date equal to the
end date would then be identical to a straddle. At the other extreme, a
chooser that has very little time before the decision deadline would be
priced at the highest of the put or the call.

Figure 21.5 shows a 2-month option (strike = 100, no drift) with a
choosing period that floats between the immediate need for a decision and
a time to decision equal to the end date. At one extreme, as the choosing pe-
riod is equal to zero, the structure will be priced at the value of a 2~month
call or put (they are equal in this case). At the other it will be valued at the
price of the straddle.

Ironically, the chooser resembles a rainbow option because the trader
has to pick one of two assets. It reaches its maximum (i.e., the straddle
price) when the correlation between the two assets (the put and the call) is

Time (years)
0.025 0.05 0.075 0.1 0.125 0.15

Figure 21.5 Chooser option price sensitivity to choosing period.
[End OCR]*

## Page 395

*[Image OCR]
Compound, Choosers, and Higher Order Options 381

~~~ Chooser
Straddie

Qu ' pond fet poet

Asset Price

Figure 21.6 Comparison between a chooser and a straddle.

exactly equal to —1, which is the case close to expiration. The reader should
go through the mental exercise of thinking that a put and a call are two neg-
atively correlated assets.

Figure 21.6 provides a pictorial comparison between a chooser and a
straddle.

There are few interesting things to say about the chooser except that it
comes halfway between a vanilla and a straddle. Its sensitivity to volatility
would be very linear when it is at the money and more convex than an out-
of-the-money vanilla as it becomes deep in the money (a chooser has the
characteristic or never being out of the money in the simple one-strike case).

Figure 21.7 shows the difference in vega convexity between an out-of-
the-money vanilla and the deep-in-the money corresponding chooser. Spot
trades at 100, with the chooser and the vanilla call) struck at 110. Volatility
is the usual 15.7%. Both have 60 days to expiration, with the chooser 30 days
to decision date.

15
12.5 Chooser
10
7.5
5 Vanilla

Volatility
0.1 0.2 0.3 0.4 0.5

Figure 21.7 Price reaction to volatility.
[End OCR]*

## Page 396

*[Image OCR]
382 Trading and Hedging Exotic Options

1.00 —
0.80 - a
w 0.60 —
> : ——— Chooser
> :
0.40 Vanilla
0.20 +.
0.00 fe

Oo oO © @ © OO WM
- NN OO MO YT

Volatility

Figure 21.8 Vega changes with volatility: Convex vanilla versus chooser.

The vega representation shows the additional convexity of the chooser,
though not as exaggerated as with the true compound (see Figure 21.8). The
option has no vega below the 10% volatility: At low volatility, the in-the-
money compound has decided that it was going to be a call and became in-
sensitive to vega. At a higher volatility, the possibility of a switch from put
to call becomes increasingly prominent.

A Few APPLICATIONS OF THE
HIGHER ORDER OPTIONS

¢ Caption, Floortions. They are options on caps and options on
floors, themselves a SDF (smallest decomposable fragment) sum of op-
tions on independent forward Eurodeposit (or other form of interest
rates) instruments called caplets and floorlets. Hence, captions and
floortions are options on baskets (a basket of Euros) and compound
options at the same time. The first characteristic allows traders to
use covariance matrices in the analysis. The second necessitates a
thorough analysis of the volatility of volatility.

¢ The compound element in captions and floortions is not remarkable,
as the volatility of the basket will pull down that of the entire struc-
ture. The back contracts will not have the same sensitivity to volatil-
ity as the front ones. In addition, the first contract will not be very
volatile.

e The multivariate element is not as dominating as with regular inter-
est rate instruments. The correlation between the prices of caplets
and floorlets presents more stable features than that of the underly-
ing forward Eurodeposits.
[End OCR]*

## Page 397

*[Image OCR]
Chapter 22

Multiasset Options

The principal difference between a quant and a trader is that a quant favors a
flawless model based on imperfect assumptions while the trader prefers an imper-
fect model based on flawless assumptions.

Multi-asset options range from the basket option to the rainbow. However,
any structure in a pair not involving the home currency as a numeraire could
be interpreted as a multiasset option. This chapter will cover the essential
risk in most of the categories but will focus most on building the intuition of
their risk management with a simplified dual-asset option structure.

Mathematicians have the habit of solving problems in the lower dimen-
sions, then generalizing to R". The reader should in the same fashion gener-
alize structures from the lower into the higher dimensions. This is where
the quantitative trading starts: While ptire trader judgment and intuition
can be relied on in the lower dimensions, a more mathematical framework
and matrix analysis become essential in the higher dimensions.

Multiasset structures include:

Choice. Options involving a choice between two or more instru-
ments: best of, worst of, rainbow options, etc, A sample will be ana-
lyzed with the rainbow options.

Linear Combinations. Options involving baskets, spreads. This chap-
ter offers a brief examination of baskets. Unlike the product and oth-
ers, asum of two exponential Brownian motions is not an exponential
Brownian. These instruments resemble Asian options in the sense
that they will present some pricing complications in addition to some
minor skew exposure when hedged with other instruments.

Product or Quotients. Options involving a product between two or
more instruments. They are easy to price, often more difficult to
hedge. The chapter includes an analysis of a sample with the Mexi-
can structured note.

A basic understanding of the risk management of multiasset options
can be obtained with the rainbow option. There are indeed many possible

383
[End OCR]*

## Page 398

*[Image OCR]
384 Trading and Hedging Exotic Options

structures but the same dynamic hedging method applies to all of them,
which makes it unnecessary to fill up a phone book with combinations.

CHOICE BETWEEN ASSETS: RAINBOW OPTIONS

M Rainbow options are options that include more than one strike price on
more than one underlying asset.

They are usually specified as having one expiration date and a payoff that is
equivalent to the largest in-the-money portion of any of the strike prices:

Rainbow (®,,K,, ®,, K,,...®,,K,,t)

with ®, indicating whether it is a put or a call (it could be a call on asset 1
and a put on asset 2) and one time to expiration in the simple case (struc-
turers could complicate life a little bit by including a multiple expiration
date structure).

This simple example shows an option on “either or” two assets A and B,
each currently trading at 100, with two strike prices of 100 for asset A and
100 for asset B. The reader will initially consider the sensitivity of the struc-
ture with 30 days to expiration then will extend the maturity to 6 months
for a better analysis. Both assets trade at 15.7% volatility. Assume an initial
50% correlation between the two assets.

Note that one could rescale the assets to a different base than 100 pro-
vided that both the prices and the strikes are multiplied by the same
amount.

The reader can intuitively see that the final payoff (Figure 22.1) covers
more areas than either of the two options each taken separately but some-
what less than the sum of the two independent options (see Figures 22.2
and 22.3).

Using the contamination principle, it is possible to see that the price of
the structure with time until expiration would look like an inflated balloon
that would start sticking to the surface of Figure 22.1 as either volatilities or
time until expiration decrease.

The structure is sensitive, in addition to the usual battery of Greeks, to
correlation. As a matter of fact, it has a correlation vega, usually overlooked
by researchers, as some tend to write off correlation as constant, owing to
improper training in the physical sciences.

Figure 22.4 displays the sensitivity of the structure to correlation. As
correlation is bounded between —1 and 1, the trader does not have to simu-
late very far for a two-asset option structure. For a higher dimension struc-
ture, the trader would need to perform the more difficult matrix analysis.
[End OCR]*

## Page 399

*[Image OCR]
Multiasset Options 385

Asset B

Payoff

Asset A 102

Payoff

Asset B

Figure 22.3 Single asset B expiration payoff.
[End OCR]*

## Page 400

*[Image OCR]
386 Trading and Hedging Exotic Options

Figure 22.4 Sensitivity of the dual strike to correlation.

At a correlation of 1, the fact that the option has two assets becomes ir-
relevant. Because the two assets are valued at the same volatility, the option
structure would trade at the price of either of the two. If there is (for some
reason attributable to the cross-volatility between A and B) a difference be-
tween the volatilities of the assets, the structure would follow the highest of
the two volatilities.

At a correlation of —1, the structure would trade at twice the value of a
regular option because it is guaranteed to be in-the-money in one of the two
assets. For every down-move in one asset, the operator is guaranteed an an-
tithetic up-move in the other so one of the assets would be in the money.

Note. This example uses dual call options. A structure with a call on
one asset and a put on the other (only one of which could be exercised at
maturity) would have opposite results: negative correlation would depress
the price.

@ A correlation vega for a dual asset structure corresponds to the change
in price in the structure that results from a change in correlation.

For a multiasset option including more than two assets, there are many cor-
relation vegas, one for every possible pair. So for a structure that has 4 assets:
. 5

Pre Py Prg
Vega Pog Pog
P34

The diagonal is, of course, 0 since every asset has a 1 correlation with it-
self. Only half the matrix is shown because the correlations are mirror
image with each other (correlation between asset 1 and asset 2 will be equal
to that between asset 2 and asset 1).

There is a sensitivity to each one of these correlations.

This method could be pushed one step further and placed in the context
of a covariance matrix (i.e., the risk of the overall portfolio). The author will
[End OCR]*

## Page 401

*[Image OCR]
Multiasset Options 387

call the covariance matrix for a portfolio, symbol 2, the matrix showing the
covariances between the assets. Traders usually learn to view the & as the
equivalent of the volatility for an asset:

with o,, the covariance between asset i and that of asset j. o,, The total ma-
trix needs to satisfy some restrictions,! and the correlations and volatilities
need to conform to some boundary or the matrix would turn negative, the
equivalent of a volatility being negative, something even experienced op-
tion traders have not yet encountered.

Correlated and Uncorrelated Greeks

The dual asset option presents more than one delta, and some assumptions
need to be made by the hedger (Figure 22.5). A believer in the stability of cor-
relations would necessarily trade the structure differently from a conspiracy
theorist.

It is therefore necessary to build the gradient, also called total delta and
correlated delta, composed of partial deltas..

A,, called partial delta in A, represents the sensitivity of the structure
to changes in the price of asset A assuming that asset B moves according to
its correlation to A.

15
Option 10
5

105 Asset B
90

100
105

Asset A 110 90

Figure 22.5 Time value of the dual strike.
[End OCR]*

## Page 402

*[Image OCR]
388 Trading and Hedging Exotic Options

A,, called partial delta in B, represents the sensitivity of the structure to
changes in the price of asset B assuming that asset A moves according to its
correlation to B.

The gradient, V, also called the correlated delta, represents the sensitiv-
ity of the structure to the changes in the price of the assets A and B assum-
ing that A would move according to its correlation to B and B according to
its correlation to A.

Figure 22.6 shows two simulated runs for asset A. One varies A with-
out a corresponding correlated move in asset B while the other takes B
into account.

Asset A moving alone gives a different result than when the asset A
move is accompanied by a correlated one with B.

The real risk then needs to be seen in two deltas, A, and A, as well as in
the total deltas. It will be useful to compare the deltas of the total structure.

Table 22.1 shows the sensitivity of the option partial delta A to the asset
A alone (asset B remains frozen).

To measure the total, or uncorrelated delta, requires more involved ma-
trix analysis:

Total delta: V7 V
which can be computed for our two-asset position as:
a7 ¢ A
[A A ] A “| *]
4° oan & | LAs

This leads to the notion of partial] gamma.

Price

95 100 105 110
Asset A

Figure 22.6 Correlated and uncorrelated moves. +
[End OCR]*

## Page 403

*[Image OCR]
Table 22.1 Partial Delta

Asset A

95
96
97
98
99
100
101
102
103
104
105

A

A

18
22
27
29
Ol
37
42
A5
46
1
57

Multiasset Options

389

As each structure has four possible deltas, it will also have the following
gammas: Only the realistic correlated gammas will be examined:

GammaAA GammaAB
| Canunad A GammaBB

GammaAA = Changes in the A, stemming from the changes in A

(B moves according to its correlation)

GammaAB = Changes in the A, stemming from the changes in B

(A moves according to its correlation)

GammaBA = Changes in the A, stemming from the changes in A
(B moves according to its correlation). Both GammaAB and

GammaBaA yield the same result.

GammaBB = Changes in the A, stemming from the changes in B

(A moves according to its correlation)

Another option involving choice is the outperformance option.

@ An outperformance option (between two assets) is an option that entitles
the owner to buy or sell one asset against the other at a predetermined

rate. They are generally calls on the maximum, puts on a minimum.”

Outperformance options are a useful tool in the study of numeraire issues.
An interesting way to view them is with index allocation. A fund manager
who does not have a fixed allocation can assume that he has a theoretical
delta that is similar to that of such option. He would then use the delta vec-
tors and gamma matrices to continuously adjust his position.
[End OCR]*

## Page 404

*[Image OCR]
390 Trading and Hedging Exotic Options

An outperformance option can easily be seen as a spread option once it
is specified:

Max(S1, $2) = [S1 + Max(0, $2 — S1)]

which means that an outperformance option is nothing but one asset plus a
spread between two of the assets. As to spread options, it is better to view
them from the vantage point of a basket by seeing that one of the assets has
a negative weight.

Perhaps the mother of all outperformance options is the option on the
bond future. It entitles the short future party to deliver the cheapest of the el-
igible bonds. As such it becomes an option on the minimum of several assets.

LINEAR COMBINATIONS

@ An option on a linear combination of assets has the following specifi-
cations:

nn

Option ( Ss ws,)
1

i=
The category contains basket options, spread options, and any combination

one would dream of. In addition, Asian options can be included in it, as will
be seen. ,

Example

° With n= 2 and w, = w, = 1. The option will be a call, put, or other
(e.g., digital) on an underlying that is the sum of two assets.

¢ With 7 = 2 and w, = 1 w, = —1. The option will be a call, put, or
other on an underlying that is the spread between two assets. Such
option is generally called a Margrabe option.?

¢ With n = 500 and the w, variable. The underlying basket wowld be a
stock index, and so on.

Linear combinations between assets pose the problem of lognormality,
which does not occur with other structures. To illustrate, assume that as-
sets A and B are both lognormal:

A, = A,exp(—% a2 At + o, VAt Z,)
B, = B,exp(—2 G2 At + o, VAt Z,)
with Af the time until expiration of the structure, A, and B, the initial price

of the assets, 0, and o, the volatilities, Z, and Z, Wiener processes inde-
pendent identically distributed with unit variance and 0 mean.
[End OCR]*

## Page 405

*[Image OCR]
Multiasset Options 391

It is easy to see that the sum W =A + B cannot be put in a geometric
Brownian shape

W = A,exp( —2 a? At + a, VAt Z,) + Bexp( —'2 6,’ At + o, VAt Z,)
nor would W = A — B whereas setting W = AB, results in
W = A,Bexp( —% (G,2 + 0,2) At + (6,Z, + 0, Z,) VAR),
or if setting W = A/B gets
W = A,/Bexp(—'2 (2 — o,2)At + (o,Z, — %Z,) WAB),

both of which are geometric Brownian.

Basket Options

Basket options are options on a weighted sum of two or more assets. They
typically have a strike price on the net weighted sum. Basket options are
omnipresent, since often many products are considered a basket of some
subcomponent,

These options will not receive much attention here, as they can be simply
analyzed as options on a product, the product being the underlying basket.
Two hitches—lognormality and the correlation between the underlying
assets—represent some singularities, however.

Lognormality

Problems in hedging basket options can be of some significance when deal-
ing with markets that exhibit a strong skew.

Many operators price basket options as if the underlying basket were a
commodity on its own following a stochastic process similar to that of other
commodities, with its volatility derived from its own time series, This, how-
ever, conflicts with the fact that an average (or any linear combination) of
assets with a lognormal distribution does not follow a lognormal distribu-
tion. So there is a conflict between saying that the SP100 components are
lognormal and that the SP100 as a commodity is. In such cases, operators re-
sort to assuming arbitrarily that the most traded of the basket or the com-
ponents is going to be the lognormal product.

This problem arises in swaps and Eurodollar options. It is well known
that a strip is a basket of Eurodollar contracts strung together and weighted
by some discounting factor. An option on a strip is therefore a basket op-
tion. Which one needs to be lognormal?

A mitigating factor with stock indices and fixed income products is that
when the correlation between the components is high the sum would come
closer to a lognormally distributed asset.
[End OCR]*

## Page 406

*[Image OCR]
392 Trading and Hedging Exotic Options

The following example illustrates the problem.

Example: There are two uncorrelated assets A and B, with prices
S, and S,, independently lognormally distributed with volatilities o,
and o,.

The basket rule indicates that the volatility of the basket S contain-
ing both assets A and B would be the weighted square root of the volatil-
ities minus correlation, in this case zero. Therefore

So = \wKo,S?, + wro5S,
with w, and w, the weights imparted to each (a more general formula
will be presented in the next section).

Figure 22.7 graphically displays why the process, the volatility of which
is known, is not lognormal.

Assume no drift and equal weights of .5, S,, and S, both initially at 100
and 50% volatility. So the operator can expect the resulting sensitivity of a
position that is long a combination of the options of the assets and short the
basket (in ratio satisfying gamma neutrality): It looks like a risk reversal.
Figure 22.7 shows such position at a high volatility. A more acute version of
this problem will be discussed with Asian options.

Correlation Issues

The multiasset structure will be considered on n underlying securities. Ig-
noring the lognormality issue* the operator could consider the structure a
pseudovanilla option and price it by assuming its volatility is that of the
basket. To derive the correlation sensitivity, he would reprice the structure
at different levels of correlation. This is best approximated by multiplying
the vega by the effect of the correlation on the volatility.

Gamma

Basket
Standard
Deviations

Figure 22.7. Gamma of a portfolio long asset options, short the basket options.
[End OCR]*

## Page 407

*[Image OCR]
Multiasset Options 393

Option Wizard: A Pseudovanilla

A pseudovanilla option is an option used for risk management purposes to test
some of the sensitivities of a more complex product.

The pseudovanilla for a basket is the option that trades at the net volatility
of the basket. It could only be used to test some of the sensitivities.

A pseudovanilla for a barrier is, as seen earlier, a risk reversal, provided one
only tested for the skew effect. It could also be (yes, indeed), a calendar spread,
provided one tested only for the sensitivity to the term structure of volatility.

Before gauging the effect of an increase in any of the volatilities over
the overall structure, the trader must get some formulas updated on a
spreadsheet.

Assume that a basket S is a sum using the individual weights w, of as-
sets X, through X_, that is:

S=w,X,+ w,X,+...w,X,
Using o, the volatility of the basket and a; that of asset i, cov(i,j) the covari-

ance between assets i and j, p,, the correlation between assets / and asset j,
the operator has

i
82
which, since the correlation is the ratio of the covariances to the product of
the individual variances, is equivalent to

2
Os

(> w? var(X,)X?, + 2 >) w, w, cov(X,, X,) X; x,)
Pa a

1<]

o,= s > wot X? + 25° WW .p;, 0 ,0,X;X,
i=l i<j
It follows that the correlation vega (as we can call it) of the structure is:
bo, wWWT 0 XX, ’
dp S20,

ij
Example (simplified): A 1-year option on the average of USD-DEM and
USD-JPY, struck at 1.19 (for simplicity, divide yen by 100), which is at-
the-money presently. USD-DEM trades at 1.42 and USD-JPY at 100.
Rates are 5% for the USD, 5% for the DEM and 1% for the JPY. Note: To
compute the forward of a basket, linear weightings of the forwards can
be used. There are no complications (owing to the homogeneity of the
function) as:

Forward (2w,S,, t) = 2 w, Forward (S,, t).

mir
[End OCR]*

## Page 408

*[Image OCR]
394 Trading and Hedging Exotic Options

Hence, the operator can ignore the forward (taking into account no
shadow gamma) and simply use spot for the volatility computation.
The 6-month volatilities of USD-DEM and USD-JPY are 11.85% and
11% respectively. Assume that the weights are each .5. The correlation
between the instruments is .60. Ignoring lognormality for now, the op-
tion on the basket can be priced as a vanilla with the volatility equal to:

(1.19 .5 + 1.42 .5) 71 (.5? 11? 1.19? + .5? 1185? 1.42?
+ 2.5.5.6 11.1185 1.42 1.19) = 10.28%

Should the correlation go to .5, the volatility of the basket would de-
crease to 9.96%. Should correlation go to —1, the effect would be to com-

press the option price.

Table 22.2 shows the simplified example, with the arbitrage aberrations
at the extremes of 0.9 and —0.9.

Table 22.2 Correlation Vega of a Basket Option

Correlation
Correlation Basket Price Vega

1.0 11.50% 4.58

0.9 * 11.21 4.47 0.11

0.8 10.91 4.35 0.11

0.7 10.60 4.23 0.12

0.6 10.29 4.10 * 0.12

0.5 9.96 3.97 0.13

0.4 9.63 3.84 0.13

0.3 9.28 3.70 0.13 ,

0.2 8.91 3.55 0.14

0.1 8.54 3.40 0.15

0.0 8.14 3.24 0.15
—0.1 7.72 3.08 0.16
—0.2 7.28 2.90 0.17
—0.3 6.82 2.71 0.18
—0.4 6.31 2.51 0.20
—0.5 5.77 2.30 0.21
—0.6 5.16 2.05 0.24 °
—0.7 4.48 1.78 0.27
—0.8 3.67 1.46 0.32

-0.9 2.62 1.04 0.41
[End OCR]*

## Page 409

*[Image OCR]
Multiasset Options 395

Risk Management Rule: The Correlation Trap. A basket option
needs to trade below the level imparted by a correlation of 1 and
above the level imparted by a correlation of —1 where either of the

volatilities is independently (or partially independently) variable.
This is a simple extension of the contamination principle. It results
because a correlation becomes convex for the seller when it neigh-
bors 1 and for the buyer when it neighbors —1.

Another related option is a quanto option. This is an option where the
residual from the trade (i.e., the P/L) depends on a foreign rate that may
be correlated with it. This represents a weak correlation effect to take into
account.

COMPOSITE UNDERLYING SECURITIES

@ Acomposite underlying security is a security whose payoff is linked to
a formula related to the price of two or more securities. Typically, oper-
ators bundle in such categories options on assets that are combined in
such a way as to prevent uncorrelated delta neutrality.

Being tailor-made, these securities have no stable specifications. These se-
curities can be ratios or weighted combinations (linear or nonlinear). Linear
combinations have been addressed. The following case study examines an
indexed note (a combination involving a quotient).

Defining the categories is a generally difficult task so each one of them
must be tailor-priced, generally through numerical techniques. Complica-
tions can arise when a sum is involved: A ratio of two log-normal returns is
lognormal, but not the sum, which brings the operator back to the basket
problem.

Often correlations are involved in such instruments. Operators need to
exercise great care, however, in finding the partial derivatives (partial delta,
partial gamma, etc.), as the correlation measures will show serious instability.

Without probing into the firms’ motivations when they issue such in-
struments, this author would warn potential customers not to look at path-
dependent payoffs without a correcting lens as the final term sheets can
be seriously misleading when conditional on some path, as seen with the
French-Franc case study.

QUANTITATIVE CASE STUDY: INDEXED NOTES

This case study affords the reader a study of a composite underlying secu-
rity that can only be decomposed using a correlation analysis.
[End OCR]*

## Page 410

*[Image OCR]
396 Trading and Hedging Exotic Options

Mathematical Note: This indexed note, in addition, illustrates a
classical problem of instruments that are presented in incomplete
markets and that need to be priced on a statistical (expected expi-
ration value), non-arbitrage basis.

In consequence, instantaneous volatility and correlation can no
longer be used, but rather term (expiration) volatility and correla-
tions. We also recommends the use of trader’s sense in describing
the discrete possible states as opposed to continuous-time finance.

Background

A friend of the author was told by some indexed notes salesperson that the
following note “embedded a currency option” making it more valuable in
the eyes of his firm. The friend, seasoned enough to discount salespeo-
ple’s opinions, especially on such matters as option valuation asked the
author to tie the note with the topics covered in the book. This examina-
tion will be limited to an intuitive approach; the attempt at pricing the
note aims only at the general qualitative sensitivity rather than at estab-
lishing the elusive fair value.

Terms of the Note

Mexico issued on December 5, 1995, a USD-denominated note paying the fol-
lowing (using the notation of the term sheet):

Face Value (in USD) < Max (CETES option or LIBOR option) on No-
vember 27, 1996.

e¢ With LIBOR option = FV x (1 + USD LIBOR x Actual days/360)
known on the issue date with certainty. \

¢ LLIBOR = 12 month USD LIBOR rate 2 days before the issue. Assume
. that the LIBOR option will be fixed at 1.056.

CETES option = FV MXN, MXN, /USD, .

¢ MXN, = MXN-USD (i.e., number of pesos per 1 U.S. dollar) 2 days
prior to the issuance of the note. Assume it to be fixed at 7.7.

¢ MXN, = 1+ Max (CETES rate on expiration minus 2 days (annual-
ized actual/360) — .06, 0). The CETES rate is the Mexican govern-
ment note. The analysis ignores that the CETES rate will be
[End OCR]*

## Page 411

*[Image OCR]
Multiasset Options 397

compounded from the issue date, which causes some differences, up
to 30 basis points.

¢ USD, = Peso rate MXN-USD on expiration minus 2 days.

The term sheet explains that the government of the United States of
Mexico would be glad to accommodate the holder in local currency should it
run out of dollars. This can be interpreted as meaning that the owner is not
exempt from the well-known default, or convertibility risk: Governments
have gotten into the bad habit of making the transfer of money illegal when-
ever they face a shortage of foreign reserves. The note may thus end up pay-
ing back in nonconvertible pesos forcing the holder to brush up his Spanish
and elect some peaceful existence near a Mexican golf course.

Where Is the Underlying?

Using revised terminology, the note’s value (in dollars) will be:
exp (—rt) Max (7.7 (1 + c,) X 1/MXN-USD, 1.056) — {Default risk}
or roughly,
exp (—rt) Max (7.7 (1 + “) X 1/MXN-USD — 1.056, 0) + 1 — {Default risk}

since the 1 + LIBOR at inception will be equal to the financing rate exp (1);
r is the continuously compounded rate used in these formulas; c is the
CETES rate. Therefore it could be:

1 + Call on U with a strike price of 0 — {Default risk}.
With

U = 7.7 (1 +¢,) X 1/MXN-USD — 1.056
s
Barring the default risk, this note seems to be a simple call on an underlying
security that we first need to define. The security is principally composed of
the product of the CETES and the currency (in American terms) or the divi-
sion of the CETES and the currency (in peso terms). Is this product a cur-
rency? Not very likely, unless one of the terms (the CETES rate) remains
frozen. Also the note would be a call on the CETES rate if the currency re-
mained frozen. Somehow, the cases where one would move and the other re-
main frozen are limited to spreadsheet exercises. Assume that both CETES
and the currency move in opposite directions when expressed in dollar terms
(the rates go up when the peso weakens) and in the same direction when the
[End OCR]*

## Page 412

*[Image OCR]
398 Trading and Hedging Exotic Options

operator looks at the currency in peso terms (the rates go up when the dollar
rallies against the peso). So the underlying for the currency option will be
conditionally linked to the currency or the CETES.

The potential buyer could therefore start correcting the salesperson. He
should have said, “This note has an embedded option on an instrument that
I cannot quite define but that is related somewhat to a currency.”

Should such a product exist in the market, it would be easy for the
holder of the note to go delta-neutral against it, and apply all the great rules
of dynamic hedging. But the salesperson no doubt would show no intention
to make such a market. It is therefore necessary to study the structure as a
multiasset by design, as defined in Chapter 1. This is similar to a currency
option on a bizarre pair that the owner needs to hedge by triangular decom-
position only, therefore having to look at such matters as correlation.

As to the default risk, it could be best priced by considering it an Amer-
ican binary option on a percentage of the face value, itself correlated to the
currency rate.

Triangular Decomposition

Omit the 1 from the formula for a while. Look at the payoff at expiration
(Figure 22.8).

The figure shows that most of the payoff of the note corresponds to
areas Northeast, which corresponds to a higher CETES and a stronger Mex-
ican peso, something not very likely to occur in real life. It is unnecessary to
price it to see that such an area is not very likely: Operators generally ex-
pect high-yielding currencies to raise their rate when they weaken and

Figure 22.8 Expiration payoff.
[End OCR]*

## Page 413

*[Image OCR]
Multiasset Options 399

lower it during the (rare) times of strength. The embedded currency option
only comes into play in such areas.

Furthermore the expected return from the Mexican peso is far to the
right of the inception rate, because the Mexican forward trades at a discount
(the dollar at a premium).

Assuming 1-year forward rates = CETES (a simplification as the
yield curve is being eliminated for pedagogical reasons)

Mu = interest rate differential. Assume it to be .225% (the differ-
ence between the United States and Mexican rate to use only the in-
formation available on the sheet).

Forward Mexico = Spot exp ( —Mu ft) the equivalent of 9.64 peso to
the dollar. Additional pricing information will be relegated to the
technical note.

¢ The first conclusion is that, as the salesperson initially said, such
note has some form of “optionality” in it (in the absence of default
risk) because it pays nothing in most parts of the map and could pay
something somewhere. By the rule of stochastic dominance it needs
to be worth some premium.

¢ The next step is to evaluate the optionality by computing the proba-
bility of ending up on each spot on the map. Intuitively, the operator
has already established that the raised areas are quite unlikely.

Next, it-is necessary to estimate the probability environment at some
high correlation between the two elements MXN-USD and CETES. Figure
22.9 shows the probability of being on every spot of Figure 22.8 assuming a
75% correlation: .

¢ A warning. One needs to be careful about correlation. This map is
an expiration map and requires the use of longer term correlation.
Using daily correlations will include some noise that will be of
small significance with such an abnormal market as Mexico. An-
other way to look at it using the variance ratio method of analysis is
that there may be a bias linked to the frequency of sampling:
Shorter sampling periods would show a lower correlation measure.
Finally, up-correlation is usually different from down-correlation, a
point discussed in Chapter 15.

¢ Another warning. When deriving fair value for such a security,
one needs to be careful about the horizon. Assuming that the oper-
ator is with either a mean-reverting process or any process with
some heteroskedasticity, one needs to be warned against the risks
[End OCR]*

## Page 414

*[Image OCR]
400 Trading and Hedging Exotic Options

Probability

Density
CETES

MXN-USD

Figure 22.9 Probability map with a 75% correlation.

of assuming that V At, the square root of time notion works (i.e., the
variance is linearly proportional to the time horizon). Therefore one
cannot use the notion of daily measured volatility to price the final
payoff. See the discussion in Chapter 6.

The reader can mentally combine the two maps to examine the condi-
tional nature of the payoff. We do not need additional mathematics to see
that the area Northeast of the mountain will show increased country de-
fault risk, a risk that is not rewarded for, as if it would force the owner of
the note to retire in Mexico. There is indeed a tautological relationship be-
tween higher interest rates in a country, weaker currency, and increased
default risk.

An examination of the sensitivities shows:

¢ The note statically gains in value if either the CETES or the currency
gain in volatility. This can be visible on the map: the mountain in
Figure 22.10 would become thicker and cover more payoff areas.

¢ The note statically gains in value if the correlation weakens ketween
the two. It would also make the area of the mountain less diagonal
(see Figure 22.11).

If Mexico was defined as a biased asset, as in Chapter 15, the operator
would have expectations of a behavior that is conditional on market
regimes: Mexico would have a higher volatility in some panic conditions
that are linked to the position of the currency on some spots over the map.
The definition of biased asset also stated that the down-correlation is higher
than the up-correlation, a result of a negative-Poisson that affects both the
rates and the currency pair.° If the two are governed by a diffusion plus a
negative jump that affects both of them, it becomes conceivable that the
[End OCR]*

## Page 415

*[Image OCR]
Multiasset Options 401

Probability

Density CETES

MXN-USD
Figure 22.10 Higher CETES and currency volatilities.

currency and the rates exhibit some form of relative independence (diffu-
sion, hence lower correlation) outside of periods of panic (jumps, hence
high correlation).

The point to be stressed is that static examination of the Greek deriva-
tives bears no true benefit. This analysis therefore will skip the sensitivity
analysis. ,

The conclusion is that correlation-dependent products, particularly
those in biased assets, are too difficult to listen to casual opinion.

Probability

Density CETES

MXN-USD

Figure 22.11 Lower correlation.
[End OCR]*

## Page 416

*[Image OCR]
402 Trading and Hedging Exotic Options

Pricing note (advanced topic):° USD-MXN is the Mexican rate ex-
pressed using the dollar as the numeraire.

USD-MXN = USD-MXN, exp{(Mu — 50,2) f+ Vtx oy}

with Mu = .225 the interest rate differential and USD-MXN, = 1/7.7. x is
a random variable normally distributed with unit variance and 0 mean.
o,, is the volatility of the MXN-USD dollar pair (it is the same as that of the
USD-MXN).

The operator needs to study the process for MXN-USD not USD-MXN
as the numeraire is (so far) in USD. He has the expected rate period t

E, (rate period t) = | USD-MXN, exp {(Mu,, — .5 0,4) t + Vex Oy} p(x) ax
= USD-MXN, exp(Mu £)

p(x) is the centered normal density function. Remarkably if he chooses
MXN-USD as numeraire, he would have as process:

MXN-USD = MXN-USD, exp( — Mu —.5 0,2) t + Vi xo,

E,(rate period t) = [|| MXN-USD, exp{(Mu, - 5 042) # + Vix oy) p(x) dx

= MXN-USD,exp{( — Mu + 04,7) #

So the expectations for USD-MXN are 9.645, while those of the MXN-USD
are at 35% volatility in the currency, at 10.90. This is another illustration of
the two-country paradox.

As to the CETES rate, the trader uses:

c, = c,exp{(d — 502) t + Vt (px + VA = p”)y) op}

with d as the drift of the CETES rate and p its correlation with the currency.
The fair value of the note would therefore be:

exp (=r!) [ [ (Max (7.7 (1 + ¢,) X 1/MXN-USD — 1.056, 0) + 1) p(x) ply) dx dy

which can be computed using standard pricing techniques.
[End OCR]*

## Page 417

*[Image OCR]
Chapter 23

Minor Exotics: Lookback and
Asian Options

A veteran trader-intellectual friend of the author, waxing philosophical after a few
sips of a remarkable Bordeaux, pointed at the bottle and uttered: “A trader with a
good mind, like good wine, improves with age. A bad trader rapidly turns into
vinegar.”

This chapter presents some options that, despite the plethora of research
publications on their subject, do not represent any meaningful trading nov-
elty on their own. They are called minor exotics as most of the information
has been included in the previous chapters on exotics.

LOOKBACK AND LADDER OPTIONS

™@ Lookback options are options that allow the owner to sell the high or
buy the low over a set period.

As such, they are called “floating strike.” There are other varieties of look-
back options, such as the fixed strike where the owner of the structure owns
an at-the-money option that will be exercised at the maximum over the pe-
riod. Such varieties will not be examined here. 5

The major problem with lookback options is that they do not customar-
ily trade, owing to their high price. They are indeed very costly: as a rule of
thumb, approximately twice the premium of a conventional option.

Example: The structure presented in Figure 23.1 allowed the owner to
sell the high over the period. The high turned out to be 116.92. The
owner became the happy owner of a put struck at 116.92.

There are many ways to view lookback options for risk management.

Exotic option traders grounded in barrier options usually have no real dif-
ficulty understanding the skew exposures stemming from the product.

403
[End OCR]*

## Page 418

*[Image OCR]
404 Trading and Hedging Exotic Options

120

Price 115 Max

110
105
100

95

90
Time

Figure 23.1 Maximum of a sample path.

The principal risk is the gamma. It becomes easy to see that the option
has a gamma that is “one-sided,” as it requires action on one side of the
market and not the other. A lookback option that has reached a maximum,
say 117 (as in Figure 23.1), and dropped below such maximum, would not
really present any major gamma for the seller unless the market reaches a
new high. This is called a one-sided gamma exposure. Such a situation is
reminiscent of barrier options where the gamma changes in style beyond a
barrier. °

The Rollover Option

Lookback options can be viewed for risk management from two angles: the
roll-over replication on one hand and the limit decomposition on the other.
The original pricing by Goldman, Sosin, and Gatto! was inspired by the fol-
lowing replicating strategy.

Assume that the dealer sells a lookback call to the customer; that is, the
right to buy the low over the next year. The dealer would initially buy a one-
year at-the-money call. Should the market rally immediately, the ‘trader
would be relieved as the risk would no longer require dynamic hedging. The
markets, however, have a bad tendency of swinging a little. To guarantee
owning a call struck at the low, the dealer would have to roll his option po-
sition into a lower strike call every time the market dropped. Such a roll in-
cludes a cash expenditure. The additional cost could be easily calculated
with the aid of a stochastic integral. The lookback option would then be the
sum of the two components:

1. The original option.
2. The costs of the call spread “roll-down.”
[End OCR]*

## Page 419

*[Image OCR]
Minor Exotics: Lookback and Asian Options 405

The stochastic integration (in present value) of the cost function corre-
sponds to the markup over an at-the-money option. Such an option is
known as a “strike bonus.”

Using this framework to analyze lookbacks can provide the operator
with the intuition of the real problem with their trading: the skew. A look-
back call will be more valuable with a downside skew (when lower strikes
are more expensive) because its maximum gamma will always be posi-
tioned at the lowest price the market reached over the lifetime of the option.
Likewise, a lookback put in such a skew environment would be less valuable
because its gamma would be located at upper levels. The reader should
know by now that in down-skewed markets the upside is not where the
gamma should be located.

Risk Management Rule: The lookback option has a third moment

exposure that is not possible to hedge with vanilla options.

The reason for the rule is that the gamma of the lookback would always
be maximum at the recorded extreme during its lifetime. A vanilla option
that would initially match the gamma risk would rapidly move away from the
money in a continuous sell-off while the lookback would retain its gamma.

One consolation, however, is that the lookback gamma remains one-
sided. The operator would not be whipped beth ways. Figure 23.2 shows the
exact issue: The gamma of the option follows a bell-shaped curve around
the strike while that of the lookback looks like a cliff with a flat top.

Before looking at the second method of pricing the lookback for hedging
purposes, the reader should examine some form of combination involving
them.

™ Ladder options are lookback options giving the owner the right to sell

the high or buy the low at some set discrete increments. They are also

called discrete lookbacks.
s

Lookback Gamma
Gamma

Vanilla Ga

Figure 23.2 Lookback gamma.
[End OCR]*

## Page 420

*[Image OCR]
406 Trading and Hedging Exotic Options

Example: The same lookback as before allows the owner to sell the
high but in $5 increments. It can sell 105, 110, 115, 120, and so forth. The
high was 116.90 but the owner of the structure was able to sell at 115,
which is not disastrous.

Figure 23.3 shows that the ladder resembles a lookback that is “rounded”
in price. They will come cheaper because their payoff is capped at that of
he lookback. In the previous example, the ladder paid off $1.90 less than the
lookback.

Mathematicians know that applied problems can be approached from
several different disciplines to yield identical results: The same applies to
option valuation. The ladder option can be viewed as a lookback when the
strike bonus part is only calculated off a set price increment. It can also be
seen a strip of knock-in options.

Assume that the option gives the right to sell the asset at the best high
rounded to the lowest 5 points, say 100, 105, 110, 115, and so on for the next
year. Also assume an asset price of 100. Such option can be decomposed into
the following, using the notation KI(strike, trigger):

Long — 1 K1(100,100) (that would be knocked in immediately)
Short 1 K1(100,105)
Long — 1 KI(105,105)
Short —1 K1(105,110)
Long 1 K1(110,110)
Short 1 KI(110,115)
Long 1: K1(115,115)
Short 1 KI(115,120)
Long 1 K1(120,120)

Price
115 ‘
110

105

100

Time

Figure 23.3 Ladder levels.
[End OCR]*

## Page 421

*[Image OCR]
Minor Exotics: Lookback and Asian Options 407

and so forth until the difference between the KI(S,S) and the KI(S,S + 5)
vanishes to a small number.

Now the skew: The reader can calculate the skew for knock-in options
using the method in Chapter 19, thus building a more accurate model for
both ladders and the general lookback.

The lookback option can be viewed as a limiting case of such decompo-
sition as the difference between strikes becomes very narrow. The general
formula is:

Laddermax = limit «0 Sum(KIP(S,,S,) — KIP(S,,S; + €))
from i = 1 to infinity, with € as increment.

The ladder minimum is the same mirror constructed with knock-in
calls. Skew-minded traders will not fail to improve on the formula by using
a volatility function, that is related to the strike price of the barrier. Typi-
cally, this will raise the price of the lookback considerably.

Figure 23.4 shows how the price of a strip of knock-in options will con-
verge away from that of a lookback as the trader increases the size of the dif-
ference between the legs of the knock-in. The graph depicts an at-the-money,
one-year option with no drift and 10% volatility. At the extreme left, is the
price of a lookback, close to twice the price of a vanilla. At the extreme right
is the value of a vanilla.

In addition to the skew, the product shows the “driftwood” effect as
traders tend to call it, which is the same chronic problem as that incurred
with barrier options. This is the concentration of strikes right below the re-
cent lows and above the recent highs. It results in considerable bad gamma,
which is the gamma the operator will start incurring as the cost of covering
it turns onerous.

Finally, it is worthy mentioning another application of the omnipresent
Arcsine law of the random walk that was used for the distribution of the
profits or loss of an individual trader in Chapter 3 (see Figure 23.5). The dis-
tribution of the extrema of any Brownian motion will be such that it will be

maximal very early on or very late in the game.
>

Strip 7.

Wn An TM aM

Increment
10 20 30 40 50

Figure 23.4 Convergence of a KI strip to a vanilla.
[End OCR]*

## Page 422

*[Image OCR]
408 Trading and Hedging Exotic Options

Frequency

Time

Figure 23.5 ArcSine law.

Why?

Starting in the middle of the period makes the concept easier to grasp.
Assume that the option has six months to expiration and that the market
just made new lows. The probability of the new low of being the low of the
next six months is very weak as the market has a high probability of moving
down further. It is as if the trader were standing with an entirely new look-
back option in front of him.

Look at the beginning of the period: The possibility of the market hav-
ing reached one of its extremes is very high as any trend will build up on
the other side. If one throws dice and make three winnings, the cumulative
P/L is three. The trend is such that the probability of the lows (here 0) being
close to the origin is highest. ,

A FOOTNOTE ON BASKET OPTIONS: ASIAN OPTIONS

Asian options are commonly entrusted to junior traders owing to the regu-
larity of their risks and the ease of their management.

™ Asian options refer to options on averages.” Their payoff depends on a
weighted combination of events through a certain period. ,
Asian options include the average floating strike where, as with the lookback,
the owner becomes entitled to a strike price through time, and the fixed
strike, an option that has a known strike and that settles at the weighted av-
erage during a certain time window. There are several types of averages:
Geometric and arithmetic are the most common. The usual Brownian mo-
tion is, this time discretized in equal intervals of one step for the sake of
simplicity (and no drift):
S,=S,_, exp( —% ot + oVtz)

t
[End OCR]*

## Page 423

*[Image OCR]
Minor Exotics: Lookback and Asian Options 409

where o is the volatility, t the time until expiration, and z a centered normal
variate.

¢ Geometric average:
n 1/P
itis
i=]
with the w, such that their product is equal to p. The author has never
seen any trade but it is necessary to analyze it to view the difference

between it and the arithmetic one, as the geometric average is more
natural in the financial markets.
e Arithmetic average:
“ w, S;
=i

t

with the w, such that their sum equals n.

As with the exercise of Chapter 22, the following will help the reader
understand some of their pricing difficulty.
Assume an average of four days. One would have the following process:

Geometric

average = (Sq Sy exp(—2 1/365 0? + o\(1/365) Z,) Syexp(—2 2/365 0?

+ o\(1/365)(z, + z,)) Sgexp(—2 3/365 0?
+ o\(1/365)(z, + z, + z,))
= S, exp( —% 6/365 o? + 0\(1/365)(3 z, + 2z, + z,))"*

The process takes an exponential shape and can easily yield some
closed form results using simple Black-Scholes-Merton.

The arithmetic average, however, will be more illuminating. Look at the
process for the same four days:

Arithmetic

=} 1 2 a pace
average 74(Sy + Sy exp(— 2 1/365 o* + 0\(1/365)z,)

+ S,exp(—'2 2/365 o? + 0\(1/365)(z, + z,))
+ S,exp(—% 3/365 o? + 0\(1/365)(z, + 2, + z,)))

exactly as if the user were looking at a linear combination of options on in-
dependent assets of the same volatility (it is known that z,,z,,...z, are in-
dependent). There is little lognormality to it as the process cannot be
summarized in the form S, = S, exp(something). In other words, one cannot
obtain dS/S a normally distributed variable.
[End OCR]*

## Page 424

*[Image OCR]
410 Trading and Hedging Exotic Options

This complication has caused some ink to flow. As a trader, the author
was rapidly put to sleep by Asians. As an amateur probabilist, he discov-
ered the process to be quizzical.*

The remaining of the chapter will focus on the few points that matter
for hedging. The pricing methods currently in place attempt to fudge the
process of the average 2w,S, by finding some form that can track the bias be-
tween the lognormal distribution and that of the averages.

The reader will be happy to learn that such bias is the skew. Most fudg-
ing methods attempt to replicate the distribution by detecting its moments
and coming up with a lognormal distribution function that satisfies such
moments. Still most traders are satisfied and resort to the more entertain-
ing Monte Carlo engines as will be seen.

By comparing the distribution of the average to that of the underlying
itself, it is possible to see that the ratio of the second moments of the distri-
butions is close to 1/3. So the instantaneous local variance can be miti-
gated with a ratio hedge of equal amounts:

¢ The trader buys an Asian option and sells a corresponding vanilla in
proper gamma neutral ratios of approximately 1.73 to 1. Figure 23.6
compares each position independently.

¢ Figure 23.7 shows the spread between them, as expected, short the
skew.

A helpful way to understand the effect is through the notion of com-
pounding. A sum of exponentials does not compound in the same way as
an exponential of a sum. A sum of exponentials, say exp(1) + exp(m), that
are equal to another exp(a@ + b) will not track the growth when they
are all multiplied by 2. The result is exp(2n) + exp(2m) compared to

P/L

2 oD i «o ~
6®esk ses BSseSses
4 2

o
-
- -

116

oo
rou,
-

Asset Price

Figure 23.6 Asian and European option (locally delta-neutral gamma-neutral
ratio).
[End OCR]*

## Page 425

*[Image OCR]
Minor Exotics: Lookback and Asian Options 411

P/L

-1.500

Asset Price

Figure 23.7. Hedging the Asian option with the European.

exp(2(a + b)). The reader can try it as an exercise to see one of the convex
qualities of the exponential.

When pricing an Asian option, the reader is recommended to use for-
ward volatility and interest rates (i.e., the entire curve) as every bucket mat-
ters. It is also recommended to use a Monte Carlo whenever volatility rises
above 30%. In most other cases, a pricing model based on the usual approx-
imations would do. More precision in the approximation is minuscule com-
pared with the precision lost in the use of the homoskedastic model.

A few final points: .

¢ Bucketing. Asian options require a schedule of bucket vegas and a
distribution of forwards. Strangely, the middle point for forward
hedges resembles visually that of the stopping time.

¢ A Catch. An arithmetic average is not invertible, by Jensen’s inequal-
ity. An average on USD-DEM is not equal to 1/average on DEM-USD.
[End OCR]*

## Page 426

*[Image OCR]
PART IV
MODULES
[End OCR]*

## Page 427

*[Image OCR]
Module A

Brownian Motion on a
Spreadsheet, a Tutorial

This module provides an introduction to the random walk.

@ A random walk for security prices means that a share of the price
changes of the security over a given period of time is random. The share
of the price change that is not expected to be random is called the drift.

Financial instruments, for reasons of efficiency, are assumed to follow a
random process that we will let the reader create on a spreadsheet.

Brownian motion = Random walk + Drift

These exercises will focus on the random component. The drift will be

covered in Module B.
THE CLassicat One-Asset RANDOM WALK

Think of a drunk man on Madison Avenue. Dead drunk, he will have
no memory as to where he was last. He can only move forward and will
keep moving at the same pace. Each one of his steps will be exactly one for-
ward + left or one forward + right, as shown in Figure A.1.

After 10 steps, the following combinations are possible: 10 forward + 10°
left on one extreme and 10 forward + 10 right on the other extreme, with all

the combinations in between.

Figure A.1_| Random walk.

415
[End OCR]*

## Page 428

*[Image OCR]
416 Modules

Securities markets are assumed to follow a similar walk, with one hitch:
The size of the steps increases in function of the asset’s price. Creating a
random walk on an Excel™ spreadsheet will provide a good illustration of
the concept:!

Open a new spreadsheet

Tools — Data analysis > Random number generation
Number of variables = 1

Number of random numbers = 248

Distribution = Normal

Mean = 0

Standard deviation = 1

Output range = B4

— OK

Excel will generate 248 random numbers. The average will be close to 0.
These numbers are then called normally distributed with a mean of 0 anda
standard deviation of 1.

Put the number 100 in cell A3. This will be the initial asset price. Next,
put the volatility of the asset (say .157) in cell A2. This means that the an-
nual standard deviation is V 248 = 15.7% (a daily equivalent of 1% ona 248-
day year). In cell A4, put the following formula:

A3 X EXP(—.5 .$A$24 2 x (1/248) + $A$2 x (1/15. 7) X B4)
and copy it down to A251. This would be the path of daily returns:

S,=S,_, x Exp(-1/2 0 + o X Vi Xx W))

t t

Equating the cells to the preceding formula is accomplished as follows:

S,= A4 ’
S,.,=A3
t = 1/248
Vt = SQRT (1/248) since every line is one day, so Vt = SQRT(1/15.7)

W, is the random number with mean 0 and average positive 1 and negative —1.

Next select (A3:A251) and graph it. Figure A.2 will result.

Repeating the random number generation to get a fresh series of “white
noise” numbers causes new paths to form on the screen.

Increasing the volatility magnifies the movement.
[End OCR]*

## Page 429

*[Image OCR]
Brownian Motion on a Spreadsheet, a Tutorial 417

Brownian Motion

140
120
100

Asset Price

80

GO — ere H :
1 44 87 130 173 216
Days

Figure A.2. Geometric Brownian motion.

SOME QUESTIONS

Question 1 (a question that is invariably asked): How does one go
from the drunk man with equal size steps to the unequal steps W,? (W,
can be any number between minus infinity and infinity, and take on a
wide array of values like .56, 1.03.)

The answer lies at the core of probability laws and involves breaking
up time in infinitely small fractions of ‘40 of a second where such move-
ment takes place in digital form +1 and ~—1.

The sum of the moves after 1 second (composed of 10 moves) will
average 0 but will be spread between —10 and +10. Assuming the +1
and ~—1 came from a “fair” random number generator, the resulting his-
togram of the 1-second moves will present a clearly defined bell shape.
The reader can try that in a spreadsheet by creating a tree, with up-
branch +1 and down-branch —1. There will bé 1024 combinations
called sample paths leading to 11 possible outcomes. After 10 steps, the
outcome can only be an even number of steps. The final proportions are
shown in Table A.1 and plotted in Figure A.3. ’

This represents a heuristic derivation of the central limit theorem
(in its simplest form: DeMoire-Laplace). The sum of a +1, —1 series of
random steps approaches the bell-shaped distribution shown in Figure
A.3 when one increases the number of observations. An apparent con-
straint is that the steps need to be the same size at all times. The law is
simpler to understand in that context, except that it is perhaps the most
misinterpreted law in the history of mathematics.

Question 2 Why is the standard deviation the square root of time?
In the preceding scenario, it was assumed that every step for the +1,
—1 distribution is a unit of time. The standard deviation is the square
[End OCR]*

## Page 430

*[Image OCR]
418 Modules

Table A.1 Number of Paths

Total Number of

Up Down Move Occurrences
10 0 10 1
9 1 8 10
8 2 6 45
7 3 4 120
6 4 2 210
5 5 0 252
4 6 —2 210
3 7 —4 120
2 8 —6 45
1 9 -8 10
0 10 —10 1
Total 1024

root of the sum of the square of the moves. Here the moves all square up
to ( —1)? and (1)?, which equals 1. In addition the mean E(W) = 0

n (W, _— Ww)?
S at

i=1 n

Hence o =

and since all the W? = 1, and f, the time equals n, o = Vt

Therefore the standard deviation of the binary +1, —1 is equal to the
square root of the number of steps. It is apparent that the standard
deviation in this exercise would equal to V10 = 3.16. Roughly two-
thirds of the paths lie between +3.16 and —3.16.

Outcome of Ten Steps
1024 Sample Paths
300 --

200 t

I

100 +

|
i

Number of Paths

ot gM ©
Outcome

Figure A.3 (Almost) bell shaped.
[End OCR]*

## Page 431

*[Image OCR]
Brownian Motion on a Spreadsheet, a Tutorial 419

Option Wizard: The Diffusion

The random walk on a spreadsheet illustrates a diffusion process.

The principal property of a diffusion is that, no matter how small time is
sliced, the function remains jagged. It will always look geometrically like the
following sample path:

The same sample path with more frequent observations would look like
this:

Slicing time in even smaller increments would not make the picture any
“rounder” in any of its segments. Although it is continuous, it is nowhere dif-
ferentiable and does not become so any time. A fashionable term to describe
such jaggedness is the “fractal” structure. Students taking college calculus
learn, through Taylor methods, that any function in a very narrow segment
could be expanded into a polynomial involving its derivatives.

f(x + Ax) = F(x) + f(x) Ax + V2 #7 (x)(Ax)? +... + 1/1 FP (x)(Ax)"

They also learn that, for a function S of time t, when the partitions become
very small, (AS)? - 0 when t > 0. One of the fundamental rules of stochastic
calculus is that no matter how small time is sliced, AS? does not vanish owing
to the random element in it. As a matter of fact AS* > o? At.

For an option trader, such intuition is important: lf there is a possible
smoothness on the curve of the underlying, then the manufacturing costs of
the option through gamma rebalancing would be lowered provided the opera-
tor picked a frequency of hedges that matched such increment. Black-Scholes-
Merton does not allow for such: Even if the trader rebalanced every billionth of
the second, the cost of the option would not go down.
[End OCR]*

## Page 432

*[Image OCR]
420 Modules

A Two-AssET RANDOM WALK: AN INTRODUCTION
TO THE EFFECTS OF CORRELATION

As the single-asset random walk has been used in connection with the
drunk man, think of a drunk bird in space. Its location at any time will be
determined by the elevation and position on the map (north-south and east-
west). So the location of the bird requires three pieces of information. It is
therefore considered in three dimensions.

The two-asset random walk can easily be simulated on the computer,
and it does not require very complex matrix algebra. The hitch is that three
parameters must be estimated: volatility of asset A, volatility of asset B, cor-
relation between the assets.

Note. The reader can ignore matrix algebra for now and take the results
for granted.

As in the example before, the reader can simulate the same Brownian
motion.

The example can be done on a similar spreadsheet as before but with
cells defined as names for ease of computation.

There are two assets A and B:

Go to Excel™

Open a new spreadsheet

Tools — Data analysis > Random number generation
Number of variables = 2

Number of random numbers = 252

Distribution = Normal

Mean = 0

Standard deviation = 1

Output range = B4

— OK

5

As in the previous example, there are two series of independent random
numbers. The first asset A will be independent. The second will need to be
bridged to the first by the corresponding correlation:

Cell Al:Type Voll Cell B1 type 1 ~ Insert — Name —> Define. Name Vol1
Cell A2: Type Vol2 Cell B2 type 1 -> Insert — Name — Define. Name Vol2
Cell A3: Type Correl Cell B3 type 0 > Insert ~ Name — Define. Name Correl

Thus the cells are named. Voll is 100 times the daily volatility. Assume
for the example that it was 1% and enter 1.
[End OCR]*

## Page 433

*[Image OCR]
Brownian Motion on a Spreadsheet, a Tutorial 421

Next create the covariance matrix. It is necessary to create a 2 by 2 ma-
trix with the following:
Cell Ci: type Cov Matrix
Cell C2: type = Vol1*2
Cell D2: type = Correl < Voll x Vol2
Cell C3: type = Correl X Vol2 x Voll
Cell D3: type = Vol2%2
Next, a special matrix called Cholesky is needed to decompose the pre-
vious matrix.? Familiarity with such decomposition is not essential (see
Table A.2).
Cell El: type Cholesky
Cell E2: type = SQRT(C2)
Cell E3: type = C3/E2 Cell F3: type = SQRT(D3 — E342)

Next, give names to the Cholesky matrix:

Name E2 a_.i1
Name E3 a_12
Name F3 a_22
Next, start looking at pairs of returns.
Then, to generate logarithmic returns for the securities that agree with
the correlation matrix. .
Cell C7: type RET A
Cell D7: type RET B
Cell C8: type A8 X a_11 copy down to cell C 261
Cell D8: type a_12 x A8 + a_22 X B8 copy down to cell D261

Table A.2. The Spreadsheet

Cells A B Cc D E F
1 Voll 1 Cov Matrix Cholesky
2 Vol2 1 1 0 1
3 Correl 0 0 1 0 1
[End OCR]*

## Page 434

*[Image OCR]
422 Modules

Correlation = 0

Figure A.4 Daily moves for a two-asset random walk.

Now there are streams of paired returns. Since a22 = 0, returns B are
completely independent from returns A.

Figure A.4 shows a graph of the pairs: The returns are plotted for
columns C8:C261 and D8:D261.

A circle is apparent with the density of the points very high in the cen-
ter and diminishing away from it. Just as traders look at the market for a
distribution of one asset as a bell-shaped curve, they need to look at that of
a pair of assets in concentric circles. The transformation is shown in Figure
A.5 and in Figure A.6.

Next, change the correlations and examine the resulting returns.

Figure A.7 shows that the curves compress toward the center forming
one line as the correlation increases toward one. As the correlation tends to-
ward negative one, again only one line is forming; the returns will be equal
but with opposite signs.

Figure A.8 allows the reader to compare the example with correlations
in the real world.

Between 2 and 3
Standard Deviations

Between 1 and 2

Between 0 and !

Figure A.5 Standard deviation of returns for two assets.
[End OCR]*

## Page 435

*[Image OCR]
Brownian Motion on a Spreadsheet, a Tutorial 423

0.05 ye

Figure A.6 In three dimensions.

5 Year
Benchmark

10 Year Benchmark

Figure A.8 A high correlation: The 5- and 10-year U.S. bonds; percentage change
in bonds (prices: 3/94-5/95).
[End OCR]*

## Page 436

*[Image OCR]
424 Modules

Two Asset Brownian Motion
Correlation = 1

4-90-——
3.00 +
2.00 +
1.00 | ae

- + 0: a {+$—~~
-4}00 -2.00 £4,000.

Figure A.9 Increasing the correlation.

Figure A.9 shows the results of a correlation of 1.

To add a third security, the process is the same. The relation of the third
return to the first two will be similar to that of the second return to the
first.

EXTENSION: A THREE-ASSET RANDOM WALK
As the returns of two assets can be represented, the end arrival of the combi-
nation of three uncorrelated assets can be viewed as a sphere (Figure A.10). A

Figure A.10 One standard deviation for the returns of three uncorrelated assets
[End OCR]*

## Page 437

*[Image OCR]
Brownian Motion on a Spreadsheet, a Tutorial 425

Asset 3

Figure A.11 One standard deviation for the returns of three assets, with one
degenerate.

0.5

Asset 2

Asset ] , 1

s
Figure A.12 One standard deviation when the other two of the three assets are

100% correlated.

small sphere for 1 standard deviation, a second one for 2, and so on, would re-
place the concentric circles.

When one of the assets has no volatility (and is said to be “degenerate”)
it reduces to a two-asset world (Figure A.11).

When two of the assets are perfectly correlated (100%), the result is a
two-asset environment (Figure A.12).

When the three assets are perfectly correlated, the result is a line.
[End OCR]*

## Page 438

*[Image OCR]
Module B

Risk Neutrality Explained

The risk-neutral argument of financial theory is explored in this module. It
is essential for the understanding any arbitrage pricing of a contingent
claim. The reader needs to get a notion of it to comprehend that all proba-
bilities used in the text will be “tricked” into their risk-neutral equivalent.

STEP 1. PROBABILISTIC FAIRNESS, THE
“Fair DICE” AND THE SKEW

In this example, start by assuming an economy without rates.

The expected payoff of the security-is each final price times its proba-
bility. In the example shown in Figure B.1, it will be 101.01 x .499 + 98.99
x .501 = 100 minus the initial price, which sums up to 0. The security is
therefore priced at a level where, based on information about possible
events and their probability, it constitutes a fair game for a trader.

On a binomial tree, the following equality would occur!

p = probability of moving up

q = (1 ~ p) = probability of moving down
S(t + 1) = e“ S(t) the up-price period t +1 ’
S(t + 1) = e4 S(t) the down-price

No Skew, No Interest Rates
Initial Price Move Final Price Probability
exp(0.01) ————> 101.01 49.9%

100 <<
exp(-0.01)————— 98.99 50.1%
Expected Payoff = 0.00 100.00

Figure B.1 No skew, no interest rates.

426
[End OCR]*

## Page 439

*[Image OCR]
Risk Neutrality Explained 427

Skew, No Interest Rates
Initial Price Move Final Price Probability

_—— exp(0.05)~~———e 105.13 16.5%
100

a x9(-0.01 j————> 98.99 83.5%

Expected Payoff = 0.00 100.00

Figure B.2 Skew, no interest rates.

In the absence of all rates (to be introduced later), the user has:
St+1)=pS,+(—-p)S,= S(t)
hence

pew S(t) + (1 — p) e? S(t) = S(t)
so:pe’+(1—-p)e’=1

being satisfied at all times in the economy. This, in a nutshell, is called the
probability fairness or “no free lunch.”

It has been simplified to allow for discrete steps between prices in the
market.

What if there was a skew? Assume that the market can move up by a u
considerable larger than v. To make it probability neutral, the larger size of
the up-steps needs to be compensated with a reduced probability of the
event.

In Figure B.2, each outcome is multiplied by its probability: 105.13 x
165 = 98.99 X .835 = 100.

STEP 2. ADDING THE REAL WORLD:
Tue Risk-NEUTRAL ARGUMENT

The Drift

Assume the asset that the operator is concerned with commanded a rate of
return pw in the economy. It would then be necessary to assume that the ex-
pected return from holding the stock would be » multiplied by the time
horizon, p At, the “drift” return. What if the rate was different from the
risk-free rate in the economy because it includes a premium for “some-
thing” that could be called risk?

The answer lies in the Black-Scholes-Merton breakthrough. There could
not have been an option-pricing formula without the following argument:
Arbitrage derivation of security prices means replication. Option replication
[End OCR]*

## Page 440

*[Image OCR]
428 Modules

eliminates the delta and the exposure to the asset change of direction (i.e.,
the return) and would let the operator only concern himself with the
volatility of the asset not its required rate of return, nor the operator’s risk
preference. Black-Scholes-Merton constituted a self-financing (cash-flow
neutral) delta-neutral portfolio that would attempt to replicate the option
through continuous rebalancing (the buying and selling of the underlying
asset against buying or selling a risk-free bond with the residual cash thus
obtained). The value of the option would then correspond to the replication
costs of the portfolio, themselves depending on the volatility of the asset
and the risk-free rate in the economy. Module G includes a demonstration
of the Black-Scholes-Merton magic trick.

This leads to the following result: For arbitrage trading and option val-
uation, the drift of the underlying non-interest-paying asset needs to be re-
placed by the risk-neutral rate in the economy.

On this tree, one needs to earn the “risk-free” rate of .11 for the period.
There are many ways finance academia has adjusted for it. The first would
be to increase the difference between up-steps and down-steps. The second,
for reasons that will appear handy with the pricing of barrier options, con-
sists in changing the probability of every move (while satisfying that the
sum remains 100%). Figure B.3 shows the risky asset returns.

The operator tricks the distribution by fictitiously changing the proba-
bilities to make believe (for derivative security valuation purposes only)
that the expected payoff is that of a riskless asset. This is called a change of
probability measure. Figure B.4 shows the result. .

Say that r is the risk-free in the economy and that the asset concerned
has an expected future payoff of w.

Risky Asset Payoff
Initial Price Move Final Price Probability
exp(0.014—————> 101.41 50%
—

100
Sn exp(001 )—— 98.99 50%

Expected Future Payoff = 0.20 100.20

Figure B.3. Prechange of measure—risky asset payoff.

Risk Neutral Payoff
Initial Price Move Final Price p*
_—— exp(0.014————- 101.41 46.47%

100
7 ox9¢-0.01) ——> 98.99 53.53%

Expected Future Risk Neutral Payoff = 0.11 100.31

Figure B.4 Postchange of measure—risk-neutral payoff.
[End OCR]*

## Page 441

*[Image OCR]
Risk Neutrality Explained 429

The operator defines the up-and-down prices:

S(t + At) = S(t) e*
S(t + At) = S(t) e4

with the following restriction:
pSjt + At) + (1 — p) S(t + At) = S(t) e™
So at the note ¢ + At, the trader has
pei +(1— pet =er

He creates a new probability measure p*, which should satisfy the fictitious
replication values at the same period:

p*e" + (1 _ p*) ef = ef

Option Wizard: Why Traders Know the Risk-Neutral Argument

The risk-neutral argument is easier for traders to understand once they consider
the simpler put-call parity rules. Assume that the asset is expected to grow at
23%. The asset, however, can be “rolled” (sold and bought back to avoid de-
livery) every day at 11% annualized, the differential between the financing and
its carry, which makes its forward price one day hence lower by the annualized
11%. If the call traded at premium to the “roll” (hence the put at a discount to
the “roll”), then the operator can sell the call, buy the put and own the asset
that would cost him only 11% per day (annualized). Hence all puts and calls
need to be priced at the “roli,” the net of the risk-free rates.

Another way to view it is to see that the forward for an asset is risk neutral
(cash and carry makes it trade at the risk-free rate minus its yield). Given that
neutrality, the puts and calls need to synthetically replicate the forward owing
to the equality for European options:

Call — Put = Forward

“Bullishness” would normally raise the value of the call. But arbitrage
would also raise the value of the put, to satisfy the equality, which is a paradox:
Bearishness would also raise the value of the put. Hence one’s preferences
should not affect the option’s fair value.
[End OCR]*

## Page 442

*[Image OCR]
430 Modules

This gives a hint of the technique of changing the probability measure
to satisfy some purpose. This method, based on the Girsanov theorem’, will
present more applications in option pricing.

An option theory generalization is that the risk-neutral drift used for a
pair is the difference between the two risk-neutral drifts. For a dividend-
paying stock, it becomes the differential of carry (dividend minus the risk-
free rate), for a bond it becomes the difference between its carry and the
financing rate, and so forth.

Risk Management Rule: A dynamic hedger should use risk-neutral

probabilities at all times.
[End OCR]*

## Page 443

*[Image OCR]
Module C

Numeraire Relativity and the
Two-Country Paradox

™@ A base currency for an operator, also called a numeraire, is one in
which his final P/L is expressed. A countercurrency is the one that cor-
responds to the number of units traded.

A numeraire could also be any possible unit. An operator can have his nu-
meraire defined as a stock index or a bond fund, as many unwittingly do.

The numeraire problem! is well known to option traders with experi-
ence in the currencies. Unlike other contracts where the trading is exclu-
sively done in units per dollar (or, if the contract is traded in France, in units
per French franc), trading in currencies is done in pairs where the unit
could be either of the two currencies. The relativity is of the utmost impor-
tance as the hedging amounts will depend on the base currency.

The notion is counterintuitive because in most people’s minds a currency
pair is invertible. A JPY against the USD is the inverse of the USD against the
JPY. However, although the equality will apply to the price, it will not apply
to the profits and losses since the P/L will be caused by currencies diverging
and that such divergence will cause a change in the yardstick.

As will be seen a put on a currency is a call on the countercurrency. Like-
wise, a call on a SP500 can be viewed as a put on cash. The governing notion
that a stock market can only go to zero as expressed in dollar terms can be
stood on its head: Cash can go to infinite. For a SP500-based person, a call on
the index (put on cash) can have a limited potential, while the put on the
index (call on cash) can have an unlimited upside. This notion is largely ig-
nored by fund managers, to the detriment of their hedging precision.

As is the convention, the currencies in this module are expressed in
pairs, with the first currency the countercurrency and the second one the
base currency.

The following symbols are used in the book (and the market):

USD is the U.S. dollar.
DEM is the German mark.
GBP is the British pound (the “G” stands for Great).

431
[End OCR]*

## Page 444

*[Image OCR]
432 Modules

ITL is the Italian lira.
FRF is the French franc.
SP500 is the SP500 index.

USD-DEM will therefore be the DEM units expressed in marks per dollar
(called the European convention). DEM-USD will be the dollars per mark
(American convention). GBP-USD is the number of dollars one pound can
buy. GBP-DEM is the number of German marks per pound sterling.

SP500 will be the SP500 expressed in USD (as implied). SP500-GBP is
the GBP-denominated SP500. However a GBP-SP500 will be the units of
pounds expressed in units of SP500.

Following is an example outlining the importance of the choice of
numeraire.

Example: Assume the operator is trading USD-DEM (as is the con-
vention in the over-the-counter market, since worldwide USD-DEM is
quoted in dollar, not currency, face value). He buys 10 million of US
dollars at 1.40 DEM per USD. The market rallies to 1.50 and he sells $10
million. He will have the pleasant profits of DEM 1,000,000. For a
dollar-based entity, these profits need to be hedged, whereas for
a German company there is no risk since its profits were in its base
currency.

A counterexample: If the trader is dollar-based, he would trade in-
stead 14 million DEM at the inverse rate (DEM-USD), 1/1.40 = .7142.
Should the market drop to .6666, he would sell back 14 million DEM for
a profit of $666,400.

However, the contract is defined in USD-DEM, and the over-the-
counter unit for the currency is in dollars. So mentally, the trader would
go to the market and quote amounts that are constant foreign currency.
In the preceding example, the trader would have had to trade $10 mil-
lion and then exit by trading the dollar amount that represents a con-
stant 14 million DEM, namely $9,333,333.

5

Risk Management Rule: A position in a numeraire is considered
neutral. A residual long the numeraire will be positive P/L, a resid-

ual short the numeraire will be negative P/L. By residual is meant a
position that arose from a trade.

Conversely, the hedger considers any position in a unit that is not a nu-
meraire to be an “open” position, long or short.

A DEM balance for a German is considered square, while the U.S.-based
person needs to have his residual balance in USD. A person from a country
[End OCR]*

## Page 445

*[Image OCR]
Numeraire Relativity and the Two-Country Paradox 433

with hyperinflation might think of anything not denominated in hard assets
as an open position.

EXTENSION: THE TWo-COUNTRY PARADOX

Two traders (one German and one American) were discussing dollar-mark
(USD-DEM). The dollar was trading then at 1.42 and volatility was high. In-
terest rates were equal so the forward traded at “flat” with the cash, namely
1.42. Both had a view on the other’s currency being weaker. Both were some-
how right.

Consider the following (the reader should examine the description of
the Brownian motion that asset prices are assumed to follow): Assume that
assets follow the log normal distribution with expected volatility 20% per
annum.

The risk neutrality forces the operator to have the expected price at any
time in the future, given the spot, equal to the spot. This leads to:

Price period t = up-price period (f + 1) X probability of going up
+ down-price period (¢ + 1) X probability of going down.

Or:
S(t) =S (t+ 1 p+5S,(+1)—-p)

with S, the up-price, S, the down-price, p the probability of going up, 1 — p
that of going down since the sum of both should be equal to 1.
The operator has as the process followed by spot (as explained in Mod-
ule G):
S(t + 1) = S(t) exp(o V2)

S,(t + 1) = S(t) exp( —o0 Vi)
which solves for p
p= (1—4)/(u — 4)

with
u = exp (o Vb) and d = exp (—o V1)

So a tree can be built with two six-month nodes:

1. The asset price at the preceding node < Exp(+ .20 X Sqrt(.5)) with a
probability of 46.47%.

2. The asset price at the preceding node x Exp(— .20 X Sqrt(.5)) with a
probability of 53.52%.
[End OCR]*

## Page 446

*[Image OCR]
434 Modules

Table C.1 Dollar Mark Expectations (for the German)

Volatility 20%
6 Months 1 Year Probability Expected Price

1.8841 p? = .2159 4068

1.6375
1.42 1.42 2p (1 — p) = 0.4975 .7064

1.2327
1.0707 (1 — p)? = 0.2865 3066

Expected

USD-DEM 1.42

The expected final asset price is computed by the final outcome times
its probability. Table C.1 shows the German’s position.

So far so good: The German expects his currency, in the absence of drift
and interest rate differential, to remain the same at the end of one year. To
see at every node what he expects the inverse of dollar-mark to do, one
would take the same table and replace dollar-mark by 1/dollar-mark (see
Table C.2). The first cell will be 1/1.42 = .7042.

So the German expects his currency to remain the same in dollar-mark
terms, but to appreciate considerably in mark-dollar terms. This paradox is
rather unsettling in a global economy.

Table C.3 shows the marks per dollar point of view. Table C.4 shows the
reverse at every node.

The American, too believes that his currency would increase against the
other’s.

Chapters 7 and 17 present an extension of the paradox: A call on dollar-
mark will therefore not have the same delta as a put on mark-dollar. An in-
tuitive way to see it is by considering a move up in dollar-mark to the
infinite. For a dollar-based investor, the move in the currencies is down to
zero, therefore limited.

Table C.2. Mark-Dollar Expectations for the German ’
Move 20%
6 Months 1 Year Probability
1/1.8841 = .5307 p? = .2159 1146
1/1.6375 = .6135
1/1.42 = 0.7042 1/1.42 =0.7042 2p (1 — p) = 0.4975 .3503

1/1.2327 = .8112
1/1.0707 = .9344 (1 — p)? = 0.2865 .2677
Expected
DEM-USD 7327
Equivalent
USD-DEM 1.3648
[End OCR]*

## Page 447

*[Image OCR]
Numeraire Relativity and the Two-Country Paradox 435

Table C.3. Marks per Dollar Expectations (for the American)

Volatility 20%
6 Months 1 Year Probability Expected Price

9344 p? = .2159 .2018

8115
7042 .7042 2p (1 — p) = 0.4975 35

.6113

9307 (1 — p)? = 0.2865 .1520

Expected
DEM-USD 7042

@ Numeraire flipping consists of switching the unit in which the nu-
meraire is expressed from base currency to counterasset t.
A call on S, with strike K, with risk-neutral rate rd and counterasset
rate d can be priced as a put on 1/S with strike 1/K, risk-neutral rate d,
and counterasset rate rd.

Traders need to be warned that numeraire flipping obtains exactly the same
price equivalent (except for Asian and digital options), but the delta will be
different.”

Conclusion

_ Any risk manager/trader should be aware of the real, true numeraire before
proceeding to analyze and measure risks. This issue crops up at either a
higher volatility or in situations where many pairs are traded against each
other without any dominant home currency.

The next section is intended for the mathematical stiekler.

Table C.4 Mark-Dollar Expectations for the German ’
Move 20%
6 Months 1 Year Probability
1.07 p? = .2159 23
1.2328
1/.7042 = 1.42 1.42 2p (1 — p) = 0.4975 .7065
1.6358
1.8843 (1 — p)? = 0.2865 5399
Expected
USD-DEM 1.4775
Equivalent

DEM-USD .6768
[End OCR]*

## Page 448

*[Image OCR]
436 Modules

Mathematical Note?

The preceding situation is a direct extension of Jensen’s inequality: A con-
vex? function of an expectation will be lower than the expectation of the
function.

If ® is convex

P(E[x]) = E[P(x)]

applying for the reciprocal of the defined as (asseti—asset2) =
1/(asset2—asset1) then

1/E(x) = E(1/x)

A more complete method is to use Ito’s lemma and examine the effect of
the change of variable. This would allow the operator to take the drift into
account and get an exact figure. It can be done by creating a 1/x function of
the underlying security and deriving its expectation.

Starting with a Brownian motion:

as pdt + odZ

~with S the security, o the volatility, and Z a Wiener process.
Set u(S) = = 1/S the inverse of the rate (countercurrency) and using Ito’s
lemma:

du = a dt + a as + £2 (asp
as
we have

ou dt — 0

at

ou —1

o= dS = -— dS = — +

98 dS = 5 dS = S (pdt + odZ) »

. 2 (dS)? = < (udt + odZ)?

using the Ito multiplication table:
dt? =0
dtdZ =0
dZ? = dt
[End OCR]*

## Page 449

*[Image OCR]
Numeraire Relativity and the Two-Country Paradox 437

which obtains du
Ta (o? — p)dt — odZ

while dS/S has expectation p dt, dU/U has expectation (o? — p) dt.
Conclusion

Every operator will be subjected to a special risk-neutral stochastic process
as a function of his numeraire.°
[End OCR]*

## Page 450

*[Image OCR]
Module D

Correlation Triangles: A Graphical
Case Study

This section is necessary preparation for the analysis of multiasset options.’
The analysis is limited to implied volatility and correlations stemming from
European options.

Assets can be represented for delivery in a given maturity (the assets
themselves, not the quoted pairs), as points in a Euclidean space. The “dis-
tance” between the points will correspond to the implied volatility between
them. This method can facilitate the understanding of the relationships be-
tween volatilities and the effect the correlation has on all the possible pairs.

The definition of an asset for that purpose is a unit that needs to be
paired with some other unit for it to become tradable. So corn could be an
asset, gold another asset, the US. dollar a third one. One could easily see a
contract defined as an option for corn against gold, a vanilla product for indi-
viduals whose home currency is corn or gold.

It is always easier to analyze currencies in that light because currencies
lend themselves very well to changes in numeraire. A currency trader can
easily see a dollar-yen option on the same footing as a yen-Mongolian Tugrit.
A currency, by definition, is a numeraire, but so could any other unit, includ-
ing baseball tickets, for those who are too centered on the matter and see
everything else translated into baseball ticket equivalence.

¢ The implied volatility of a pair (for a given maturity) is measured
with the distance between the points represented by their coordi-
nates in a Euclidean space.

e Ina two-dimensional universe, the formula of the distance between
two assets with coordinates (x,,x,) and (y,,y,) is

—-—_—
d(x,y) = yy ~ y,) + (x, ~ Y)

It is easy for the mathematical stickler to see that the volatility function

thus defined satisfies the conditions for a metric, or distance function.

Hence, v(x,y) will represent the volatility of the tradable pair x-y, or x in
terms of units of y:

438
[End OCR]*

## Page 451

*[Image OCR]
Correlation Triangles, A Graphical Case Study 439

1. v(x,y) is strictly positive if x is different from y. In addition, v(x,x) =
0: The volatility of the asset expressed in terms of itself as a nu-
meraire is nil. It suffices to see that the volatility of cash in terms of
cash is 0.

2. v(x,y) = v(y,x). The volatility of y in terms of x as a numeraire is equal
to the volatility of y in terms of x as a numeraire.

3. v(x,y) is at all times lower or equal to v(x,z) + v(z,y).

In an dimensional space, the formula of the distance between two as-
sets with coordinates x = (x,,x,,...,%,) and x = (y,,y,,---,Y,) is

(ey) =) > (1)

Start by assuming that one-month DEM volatility is a point in a Euclidean
space (the one-month space) with coordinates {7,12.12}. Also pick USD to be
(0,0). It is always recommended, for the sake of simplicity, to use the point
(0,0) for the numeraire, the commodity in which the P/L is computed.

To conform to market standards, the notation v(x,y) will be used to
write v(x_y), which should read “volatility of the pair x_y”. v(USD-DEM),
the metric between point {0,0} and the point {7,12.12} shown in Figure D.1
will be V7? + 12.12? = 14, which puts the USD-DEM volatility at 14% (ex-
ample is rescaled by 100 for simplicity). By the same token, the v(DEM-
USD), the DEM-USD volatility is seen at 14% as well.

Next, add the JPY. Assume that USD-JPY is trading at 12% volatility so
USD-JPY is a vector of length 12. But there are many possibilities to put it on
the chart as the coordinates could represent an entire circle of radius 12
around the (0,0) point (Figure D.2).

Figure D.1 A two-currency world.
[End OCR]*

## Page 452

*[Image OCR]
440 Modules

DEM 7,12.12
CIRCLE OF
ALL THE POSSIBLE
JPY POINTS

Uso.)

@-i1-10 -9 -8 -7 -6 -5 -4 -3 -2 -1

Figure D.2 Possible JPY points.

The coordinates are arbitrarily set to be (x,, x,), a point that would
be on the circle. There are infinite possibilities as the Vx7 + x? = 12 gives
rise to many combinations. Pick x, = 11, so x, would be Viz — 17,
and the point for JPY would, as a result, be (11,4. 79).

Now there are three points in space: USD = (0,0), DEM = (7,12.12),
JPY = (11,4.79). These three points should be capable of providing
additional information. According to formula (1), v(DEM-JPY) will be
sqrt{(7 — 11)? + (12.12 — 4.79)?} = 8.35.

Note that the same triangle equality shown in Figure D.3 would have
been satisfied with the yen on the other side (i.e., with the coordinates
{—1.35,11.92}).

Before proceeding to the what-if analysis, examine the correlatiqn. It is
well known that the correlation between the volatility segments will be the
cosine of the angle cornering them.

Let the point n = (n,,n,,...m,,) be defined as a numeraire:

cos (b) = (x — n)- (y — n)/ v(x,n) v( yn),

b being the angle between the segments with the dot product, the “inner
product,” (x — n) - (y — n) = X(x, — n)(y, — n, and it is already known that
the metric o(x,n) =V Xx, — ,).

It could be made easier by assuming that X and Y have (x,, X5), (Yy, ¥,) for
coordinates.
[End OCR]*

## Page 453

*[Image OCR]
Correlation Triangles, A Graphical Case Study 441

DEM (7,12. 12)

USD_DEM

JPY(11,4.79)

USD (0,0) USD-JPY

23 4 5 6 7 8 9 10 11 12 13

Figure D.3. Adding the yen.

CORRELATION TRIANGLE RULE

cos(b) = correlation(x-n, y-n) = ((x, — 1,) (y, — 1)
+ (x, — ny) (y, ~ ny))/ V [(x, ~ n,)* + (x, ~ n,)"] v [ty, ~ n,) + (Y, ~ n,)"]

It is easy to see that correlation (x-n, y-n) = correlation (n-x, n-y); in other
words, the correlation between the asset x measured against the numeraire
n and the asset y measured against the numeraire n is equal to the correla-
tion between the asset 1 measured against the numeraire x and the asset n
measured against the numeraire y.

It is now possible to return to the example and calculate the correlations
imparted by the three-currency world. Compute the correlation between
USD-DEM segment and the USD-JPY:

Corr(DEM-USD,JPY-USD) = (7 — 0)(11 — 0)
+ (12.12 — 0)(4.79 — 0)/V(49 + 146.89) V(121 + 22.94) = .8035

Corr(DEM-JPY/USD-DEM) = .52
Corr(DEM-JPY/USD-DEM) = .09

For correlation to remain constant, the volatilities of the sides should in-
crease by the same percentage amount (Figure D.4). This would cause the
[End OCR]*

## Page 454

*[Image OCR]
442 Modules

Figure D.4 Correlation unchanged.

third volatility to increase as well. It is somewhat hard to imagine that
DEM-JPY now increases in volatility because both components gained in
volatility against the dollar. This author spent some precious time trying to
convince fellow colleagues who were reluctant to accept the notion.

Furthermore, when the third leg rotates (as in Figure D.5), the volatili-
ties remain constant and correlation shifts.

To add a currency to this world would necessitate increasing the di-
mensions by 1. Including the GBP requires adding one dimension to all the
others. Using an {x,y,z} coordinate system, the result (displayed in Figure
D.6) is ,

USD(numeraire) = {0,0,0}

DEM = {7,12.12,0}
JPY = {11,4.79,0}
GBP = {0,9.57,2.91)

Figure D.5 Volatility unchanged, correlation decreases.
[End OCR]*

## Page 455

*[Image OCR]
Correlation Triangles, A Graphical Case Study 443

15

Figure D.6 A four-currency world.

What if a user added the SP500? He would now have five units, four dimen-
sions for a possible representation. However, he could eliminate one unit
every time for representation purposes.

The trader represents the subuniverse USD,SP,DEM,JPY (without the
GBP). He removes the third line accordingly for the representation, repre-
senting (n1,n2,n4) to accommodate the fact that even traders can only view
the world in three dimensions.

USD = {0,0,0}
DEM = {7,12.12,0}
JPY = {11,4.79,0}
SP = {0,0,10}

The pair SP500-USD satisfies no correlation with any other segment. The
reader could, as an exercise, verify that. The implication is apparent, as
corr(2 segments) = 0 infers the cos(angle) = 0, which results in a right angle
between SP-USD and JPY-USD, as well as with SP-USD and DEM-USD, as
shown in Figure D.7.

Result: The world of volatilities can be represented as a universe of
points.

Absence of volatility/correlation arbitrage requires that all the volatili-
ties in the market (and the correlations stemming from it) satisfy the same
metric between points. Any shift of a point to the left, right, up, or down
should then result in arbitrage.
[End OCR]*

## Page 456

*[Image OCR]
444 Modules

15

Figure D.7_ Three currencies and the SP500.

CALCULATING AN IMPLIED CORRELATION CURVE

Having the different sides makes it possible to build a correlation curve,
which is the correlation implied by the rules of the triangle (Table D.1).

Table D.1_ Derived Implied Correlations*

DM FF DM/FF CORR
1m 13.40 10.30 4.80 0.95
2m 13.40 10.40 5.00 0.94
3m 13.35 10.45 5.35 0.93
4m 13.35 10.45 5.50 ° 0.92
5m 13.30 10.40 5.60 0.92
6m 13.25 10.45 5.70 0.91
7m 13.15 10.35 5.70 0.91
8m 13.15 10.35 5.70 0.91
9m 13.15 10.45 5.70 0.91
10m 13.10 10.40 5.70 0.91
11m 13.10 10.40 5.70 0.91
12m 13.10 10.50 5.70 0.91

*Source: Tradition Financial Services, October 30,
1995, closing prices.
[End OCR]*

## Page 457

*[Image OCR]
Module E

The Value-at-Risk

Abyssus abyssum invocat (The abyss calls for the abyss)."
Latin Proverb

Below is a presentation of a risk management method that, like portfolio in-
surance, can only work if a small number of people are using it. It is a para-
dox, discussed in Taleb (1997), that states that it can only work (and succeed)
if it is unsuccessful.

@ Value-at-risk? is a method of estimating the maximum loss for a portfolio
at a confidence level assuming a knowledge of the process governing its
components.

The VAR can present some useful short-term hedging tools for traders, as
we examined with the discussion on stacking in Chapter 12. However, it led
to seriously disputed applications by risk management firms that led (per-
haps innocently) their customers to believe they possessed tools to summa-
rize the overall market risks for a position, a unit, a department, or an entire
firm, in one simplified numerical exposure, without standard error.

As the concept is evolving in its applications, the presentation here will
be limited toa rapid intuitive exposition.

Value-at-risk? is growing in use among the banks and corporations both
for adequate risk allocation and appropriate quantification of the risk-
adjusted returns on capital. The idea of disclosing the overall exposure as one
simple quantity appeals to most corporate board members and regulators,
many of whom are uninitiated into the nuances and complexities of financial
market risks. They can easily be impressed by the “scientific” tools used. 5

The condensation of complex factors naturally does not just affect the
accuracy of the measure. Critics of VAR (including the author) argue that
simplification could result in such distortions as to nullify the value of the
measurement. Furthermore, it can lead to charlatanism: Lulling an innocent
investor or business manager into a false sense of security could be a seri-
ous breach of faith. Operators are dealing with unstable parameters, un-
like those of the physical sciences, and risk measurement should not just be
understood to be a vague and imprecise estimate. This approach can easily
lead to distortions. The most nefarious effect of the VAR is that it has al-
lowed people who have never had any exposure to market risks to express
their opinion on the matter.

445
[End OCR]*

## Page 458

*[Image OCR]
446 Modules

In brief, it cannot be used to say, “Within 99.7% (or within 90% or some-
thing of the sort), you are not expected to lose in the next month more than
1 million dollars.” The innocent treasurer or company official would believe
himself to be listening to a scientific statistic similar to statistics on air-
plane crashes. It could, however, be used to say: “ You are expected to lose
no more than 100,000 dollars within the next two hours with a 66% accu-
racy, provided you do not try to liquidate your position and the other simi-
lar firms do not have the same portfolio.”

SIMPLIFIED EXAMPLES

In these examples, the VAR corresponds to the risk for 1 standard devia-
tion (66% of the time), not 3 standard deviations as expressed in the com-
mon literature.

Assume for this series of examples that the universe offered to a trading
desk is composed of the following four items:

USD-DEM: USD against the German mark currency pair as quoted in
the interbank market.

USD-JPY: USD against the Japanese yen currency pair as quoted in the
interbank market.

Treasury Bond futures: As traded ori the Chicago Board of Trade (the ex-
ample translates units into face value exposure).

SP500 Index futures: As quoted on the Chicago Mercantile Exchange.

The example uses data as of May 23, 1995.
Annualized nine-month historical volatilities (percentage) are as follows:

USD-DEM 12.1

USD-JPY 12.3
Bonds 8.5 s
SP500 9.33

Volatility is defined as the standard deviation of the log of the returns
(unweighted).

Table E.1 Correlation Matrices (9 Month)
USD-DEM USD-]PY Bonds SP500

USD-DEM 1.00 0.74 0.30 0.23
USD-JPY 0.74 1.00 0.26 0.30

Bonds 0.30 0.26 1.00 0.45
SP500 0.23 0.30 0.45 1.00
[End OCR]*

## Page 459

*[Image OCR]
The Value at Risk 447

Example 1. No Diversification

The trader has a $20 million face value limit and decides to invest in going
long one of the four instruments, but does not know which one. Table E.2
shows the net exposure that would correspond to each undiversified posi-
tion. The net exposure is defined arbitrarily for one daily standard deviation
of the moves in the market, which should (in theory, unless the distribution
changes) represent 67% of all events. It should also resemble the average daily
swing in the portfolio. Should the trader decide to include wider events, like
3 standard deviations, he could multiply the exposure by 3: It would then
show the total maximum possible loss with a 99.7% confidence. However, the
management of that trader has heard of fat tails and prefers to see 1 standard
deviation as a benchmark for a comparative purpose.

Example 2. A Cross-Position

The trader might desire to get involved in cross-positions. Would he benefit
from a positive correlation?

It is apparent here that the VAR stemming from a cross-currency position
as shown in Table E.3, owing to the higher correlation of changes between
dollar/mark and dollar/yen would correspond to 71% of the risk of a
dollar/mark position.

Obviously, the diversification works, Since he has not fully drawn on
his gross limit, he can add to the position and utilize the balance of $10 mil-
lion in a diversified manner (Table E.4).

The overall position is thus reduced from the undiversified Example 1:
The capital at risk for 1 standard deviation corresponds to at least half the
risks of any of the initial four positions.

Table E.2. The Risk of Each Position

Position ,
(in $ mil) Net Exp

USD-DEM 20 $153,700

USD-JPY 20 $156,240

Bonds 20* $107,971

SP500 20t $118,514

*20 million exposure in Treasury bonds corre-
sponds to 182 contracts, computed by dividing
20,000,000 by the face value of every contract
($100,000) times 110 the price.

*20 million exposure in the SP500 futures corre-
sponds to 77 contracts, which is computed by di-
viding 20,000,000 by 523 (the contract price)
times 500 (the contract is defined as paying 500
time the index).
[End OCR]*

## Page 460

*[Image OCR]
448 Modules

Table E.3. VAR for a Cross-Currency Position

Position
(in $ mil) Net Exp
USD-DEM 5 $38,425
USD-JPY —5 $39,060
Bonds 0 —_
SP500 0 —_—
VAR $27,944

Example 3. Two Possible Trades

USD-DEM and USD-FRF are two highly correlated currencies. The trader
can position himself in one, in the other, or in the cross.

Correlation Matrix:

USD-DEM USD-FRF

USD-DEM 1 0.97
USD-FRE 0.97 1
Volatilities:

vol

USD-DEM 12.1
USD-FRF 11

The Cross-Position Risk:

Position
(in$ mil) Net Exp
USD-DEM —80 $614,798
USD-FRF 80 $558,908
VAR $154,080

s

The risk of a cross-position corresponds to 25% of the risks of each.
The VAR of the equivalent position could have been computed through the

Table E.4 Full Use of the Capital ($20 million)

Position

(in $ mil) Net Exp
USD-DEM 5 $38,425
USD-JPY —5 $39,060
Bonds 5 $26,993
SP500 5 $29,628

VAR $54,868
[End OCR]*

## Page 461

*[Image OCR]
The Value at Risk 449

volatility of the resulting DEM-FRF currency pair: 3.03%. The alternative risk
profile is as follows:

Position

(in $ mil
Vol equivalent) Net Exp
DEM-FRF 3.03 —80 $154,080
VAR $154,080

Problem 1. Assumed Homoskedasticity of Markets. (This argument is de-
veloped in Chapter 15.) The models being used rely heavily on the “normal-
ity” of the markets (i.e., price changes follow a bell-shaped curve, a fact that
is extremely rare). “Fat tails” and “high peaks” are familiar to any option
trader with a thorough understanding of the weakness of the models.

While the normality assumption could be acceptable for some applica-
tions, such as the pricing of at-the-money options, such approximation of
the distribution could be considered inappropriate for the measurement of
“worst-case” risks. Risks are located in the tails, which is where the distri-
bution is the least known.

Assuming that the volatility of markets is not constant, the rule of the a
Vino longer holds. Result: In a heteroskedastic market, the variance is not a
multiple of the time horizon. One-week (5-day) dollar volatility is not 2.23

- times that of one day. :

In addition, correlations not being constant (as dicussed in Chapter 6),

the “joint tails” of several assets are fuzzier to model.

Risk Management Rule: The VAR provides an admirable short-

term hedging tool but is by no means a risk management device.

Problem 2. Liquidity Risks. (This argument is developed in Chapter 4.)
The VAR method makes no allowance for the fact that liquidity could repre-
sent the largest risks in some markets. In some less mature instruments,
liquidity costs become indistinguishable from market risk. The selling of a
large block of securities, particularly in the event of a forced liquidation,
can lead to total market collapse.

Those not interested in the liquidation value of their portfolio need not
be concerned about its market price risk.

Problem 3. Behavior of Parameters at Times of Stress. The essence of
the VAR concept is correlation and diversification. The widespread use of
these techniques leads to the simultaneous breakdown of both at times of
excessive stress in the markets. Typically, low correlation between assets
[End OCR]*

## Page 462

*[Image OCR]
450 Modules

tends to rise in consequence to stressful events, therefore reducing the di-
versification effect. A severe example of that is provided by the “bond
market crash” of 1994 as all the bond markets sank together. Similarly, the
high correlation drops abruptly, sometimes even turning negative, in the
face of such events: The cross-maturity hedges cease to work. The yield
curves shift, often causing nonpredicted deformations, at times of high
volatility.

There is a significant conflict between a multiasset risk model and the
stress-testing approach. One relies on known relationships while the sec-
ond has grounding in a statistics-free world. Relying exclusively on the first
would mislead grossly at times of uncertainty. Relying exclusively on the
second would prevent firms from trading.

Problem 4. Dangers of Generalized Use. Assume that everyone used the
value at risk and that the market moved abruptly. The fact that such a sys-
tem became a benchmark would cause a snowball effect.

The chain would be as follows: In a schematic world of a small number
of homogeneous leveraged players, everyone would end up with close to the
same portfolio constitution and weights owing to the diversification scheme
(the optimal portfolio) {A,B,C} with weights w,, w,, We, in such a way that
they reduced their risks to fit the optimal allocation. They would all invest
more lulled with the knowledge of being comfortably diversified as they
were properly taught by the risk management consultant.

Assume that A went down in price causing the total portfolios to drop
in value. Assume that the volatility of A increased. To maintain a constant
VAR, the weights would need to be adjusted down, so the operator would
have to sell some stocks of B and C. The quantities, though small, would be
large enough to push prices lower and make operators race each other to the
state of near-bankruptcy. The factor would cause the correlation between
the components to rise, weakening the diversification effect in an unex-
pected way. The same effect would take place if one of the weights is nega-
tive and the operators are “hedging.” An interesting parameter in hedges is
that they only work when they are not identified as hedges by the multi-
tude. If most other similar institutions needed to act in a similar manner in
similar circumstances, there would be a dynamical system traders would
need to account for.*

Risk Management Rule: The market will follow the path that will

thwart the higher number of possible hedgers.

Problem 5. Computational Problems. Severe computational issues are as-
sociated with the VAR. Large firms do not cumulate all their position in one
[End OCR]*

## Page 463

*[Image OCR]
The Value at Risk 451

place from where they could be retrieved easily and in a timely way. Miss-
ing one simple position could cause sufficient distortions to invalidate the
measurements. The way the VAR is computed makes it necessary to merge
all the individual positions together, not add up the VAR of different de-
partments and branches.

Again, the value at risk of a sum is not the sum of the VARs. Despite the
sophistication of operators, most active firms experience data-gathering lags.
The complexity of their operations led them to account for units by using a
pyramid of net numbers. The VAR project would necessitate the handling of
every single position in every single branch, not easy for institutions with a
large network.

Another problem arises when the instruments become too numerous,
creating a large matrix. Estimating the covariances will create some round-
ing and estimation errors, and when such errors cumulate, they will pre-
vent the matrices from being positive definite. This makes large matrices
impracticable.°

Problem 6. What Is Volatility? So far, volatility has been discussed as if it
were some observed physical phenomenon. Because it is unstable, traders
need to consider modeling and all the resulting problems. Which volatility?
Implied or historical? Over which horizon? and so on.

Problem 7. Do Not Even Dream of Applying It to Derivatives. The sub-
ject matter of this book stresses the difficult interplay of parameters
traders must face in evaluating the risks of an option position, owing to its
multidimensional nonlinearity. The notion of higher moments also under-
scores that there is no way of tracking the risks of complex portfolios with
one simplified Greek.

4+

Conclusions. Many efforts are under way to correct the inherent defects
of the system. Most promising results, typically in the shape of nonpara-
metric methods for stress testing, are presently outside the public donaain.

Formulas. The problem is being approached from the angle of the correla-
tion matrix instead of the simpler covariance because of conventions in the
software market.

E is the Net exposure = Position (Face value) X 1 Daily stdev. (com-
puted by computing the annualized standard deviation used in option pric-
ing by the square root of 252). M is an m by m correlation Matrix. VAR is the
value at risk, amount risked for n standard deviations.

VAR = n SORT(E™ M E)
[End OCR]*

## Page 464

*[Image OCR]
452 Modules

Example: Say that in a world of four instruments an operator had the
following vector:

Volatilities:

Position:

The trader gets the net exposure E (expected risks per standard
deviation):
P19,
E= 753 [py
P39

PAO,

Then he pulls the correlation matrix:

1 Pip Pig Pra

M = 1 yg Pag
1 P34
1

The lower triangle is the mirror of the upper one since p,, = p.,-

Finally the value at risk is computed by multiplying the transpose of
E (1,4) by M (4,4) and multiplying the resulting (1,4) matrix by E (4,1) to
get one number. The square root of the number will be the VAR.
[End OCR]*

## Page 465

*[Image OCR]
Module F

Probabilistic Rankings in Arbitrage

RANKING OF SECURITIES

The value of a security can be established by ranking. If security A is worth
more than security B and less than security C, a book runner can thus get
some orientation on how to price it. This is particularly easy when the
prices are readily available for the comparative “straddling” securities and
when the comparative securities are close in price to each other.

This instrument pricing and hedging method is based on the rule of
stochastic dominance.! When an instrument or combination of instruments
is at least as valuable as another set, the author will use the sign = to ex-
press the inequality. Instrument A = B means that A in all possible events
(under all parameter changes) is deemed to be at least equal to instrument
B, or that under no event will B be worth more than A. Such rules can be
used to price options given the availability of other prices in the market.
They certainly can be extended to portfolio management.

European Option Rules

The following notation will be used: C(S, t) is a European call expiring pe-
riod t with a strike price S. P(S, t) is a European put. A C(S, t, H,) is a call
with a barrier below the market (down and out) at a price H, and a C(S, t,
H,,) is a call with a barrier higher than market (up and out) at price H,,, also
called a reverse knock-out. A call (S, t, H,, H,,) is a double barrier:

C(S, t) = C(S +4, £)

A 101 call will be worth more than a 102 call of the same expiration under
all circumstances. This rule is obvious:

P(S + a, t) = P(S, t)

A 101 put will be worth more than a 100 put under all circumstances.
The butterfly rules:

C(S — a, t) + C(S +a, 1) =2C(S, 8)

453
[End OCR]*

## Page 466

*[Image OCR]
454 Modules

Two 101 calls are worth less than one 100 call and one 102 call together. It
means that the 100/101/102 butterfly should be worth at least 0. The reason
is that it will have a value of 0 everywhere except in the “eye”; at the center, it
may have a value of 1.

By put/call parity:

P(S — a, t) + P(S + a, t) = 2 PCS, t)

Two puts at 101 should be worth more than one 100 put and one 102 put
together.

These rules may seem simplistic but veteran option traders on the ex-
changes who pride themselves on trading without “sheets” have been able to
operate by these butterfly rules alone. They can price the entire spectrum
using a small number of options.

By extension:

P(S, t) + C(S, t) = P(S — a, t) + C(S + a, t)

a straddle is necessarily more valuable than a strangle.

Calendar Rules
With 7 strickly positive,
P(S,t +7) + C(S, £ + 1) = P(S, t) + C(S, t)

(In present value: It is more appropriate to deduct financing.)

This means that a longer straddle should be worth more than a shorter
one, after taking financing into account. The rule always holds for “serial”
options, options on only 1 future (like a 1-month 2-month option on a 3-
month future). It does not unconditionally apply to options where the under-
lying security S is not “fixed,” and such weakness gets exacerbated with
nonfungible assets (i.e., live cattle. Exceptions can occur when a security can-
not be swapped into a longer maturity). 5

Barrier and Digital Rules

An American digital is worth more than a European one. An American dig-
ital of the “if touched” variety where the bet is satisfied if some event takes
place at all times during the life of the instrument is worth more than a bet
that an event would take place at expiration only.

Example: A note that pays off one dollar if the security goes above 103
at any point in the next year is worth more than a note that pays off one
dollar if the security is above 103 on expiration day:
[End OCR]*

## Page 467

*[Image OCR]
Probabilistic Rankings in Arbitrage 455
C(S, t) = C(S, t, H,,) and = C(S, t, H,)

An option with no barrier is always worth more than an option cum barrier,
since there are possibilities of losing the option before expiration:

C(S, t, H,) = C(S, t, H,, + @)
C(S, t, H,) = C(S, t, Hy, — @)

Any option cum barrier is more valuable than the same option with a bar-
rier closer to the money:

C(S, t, H,,) = C(S, t, H,, Hy)
C(S, t, H,) = C(S, t, H,, Hy)

Any single barrier option is more valuable than a double barrier if they
share a common trigger H.

Correlation Rules

These boundary rules apply for at-the-money options. They work for at-the-
money (forward) straddle prices or implied volatility of options.

The possible currencies are Currency A (used for numeraire), Currency
B, and Currency C. The notation v(A-B) is used to designate the volatility
for the pair A against B.

Triangle Inequality
vo(A-B) + v(B-C) = v(A-C)

Correlation and Volatility. Using p(A-B, A-C), the correlation of the pair
A-B and the pair B-C:

v(B-C)? = v(A-B)? + v(A-C)? —2 p(A-B, A-C) v(A-B) 0(A-C)

Example:

v(USD-DEM) + v(DEM-FRE) = v(USD-FRF)

(This rule seems to be often violated.)

v(USD-FRF) + v(DEM-FRF) = v(USD-DEM)
[End OCR]*

## Page 468

*[Image OCR]
456 Modules

It implies that given three security pairs, implied volatility between
them cannot be such that one of the implied correlations exceeds 1 in ab-
solute value.

Furthermore, by the rule of contamination, no combination between op-
tions should allow for a combination yielding a correlation higher than 1 in
absolute value.

Warning. This does not represent full arbitrage (i.e., a locked-in profit
or loss) since the P/L of combinations with the numeraire can be predicted
with certainty, whereas the “cross” (in this case, DEM-FRF for a dollar-
based person) can yield a P/L that depends on the performance of the com-
ponents with respect to the base currency.

A generalization based on the preceding is that an implied covariance
matrix (implied from at-the-money option prices) needs to be positive defi-
nite (its eigenvalues need to be positive definite). It is an extension of the
boundary rules for a three-currency world. Otherwise an operator would
be able to arbitrage it by getting some combination of options for free (or for
a credit).

o(USD-DEM) = o,
v(USD-FRF) = o,
“Cross” vol (DEM-FRF) = Volatility of FRF-DEM = o.

Thus the following covariance matrix results:

2
o* o
1 Fr],
S= witha, =a, =@+oa-o
2 12 — Sy 1 2 ¢
Oy, 85

o,,and o,, being the covariance between USD-DEM and USD-FRF.

It is easy to check whether the matrix is “clean,” or arbitrage-free. As
volatility cannot possibly be negative, 2 needs to be positive definite, which
requires the eigenvalues to be positive. In this example, the aberrations for
a 2 by 2 matrix can be visible to the naked eye while those of a 20 by,20 ma-
trix would require computer crunching.

In the next example, there are restrictions on one of the elements the
matrix (considering it to be symmetric with covariance 1, 2 = covariance
2, 1). Assume that vol USD-DEM is 14% and DEM-FRF (the cross) is 4%.
Then the restriction is on USD-FRF to be between 10% and 18%. Outside
that band, arbitrages are possible. Below 10%, an operator could buy the
USD-FRF and DEM-FRF volatilities and sell the USD-DEM. He would be
completely bounded as this operation is equivalent to being “short” corre-
lation at 1. Any drop in correlation at any time between the two currencies
would benefit the trader.
[End OCR]*

## Page 469

*[Image OCR]
Probabilistic Rankings in Arbitrage 457

Likewise, “buying” correlation at —1 could be perfected by buying
USD-DEM at 14%, buying DEM-FRF at 4%, and selling USD-FRF at 18%.
Any rise in the correlation above —1 would result in profits.

A more involved matrix can be created by adding a security to the ma-
trix. It becomes much more difficult to see with a naked eye the implica-
tions of an arbitrage-free situation:

2
Gy Fyn F3
_ 2
X=|0,, GF Oy,
O., 0. @

Take the asset 1 = USD-DEM, asset 2 = USD-FRF, asset 3 = USD-CHF,
assets 12 and 21 = DEM-FRF, assets 13 and 31 = DEM-CHF, and the illiquid
assets 23 and 32 FRF-CHF. With vol USD-DEM at 15%, USD-FRF at 12%,
USD-CHF at 18% and crosses USD-FRE at 4%, CHF-FRF at 6% and DEM-CHF
5%, one of the eigenvalues flags a negative number that causes the computer
to beep. The matrix above is faulty. The arbitrage can be constructed through
a combination of options.

Remember that enough currencies are there to allow for some arbitrage,
or at least to construct a trade with a high expected return.

Warning. The rule of the positive eigenvalues does not hold for a com-
parison between implied volatilities of options that are not at-the-money
forward.

CORRELATION CONVEXITY RULES

The next few rules emanate from the contamination principle.

One can easily see from the previous examples that a structure that is
set up short a “high” correlation (close to 1), that is, that should benefit
from correlation dropping, or long a low correlation (close to —1), is
stronger than a structure that is set up in a correlation that is closer to 0,
everything else being equal. Say that a trader in a three-asset world pays
11% for USD-FRF and 4% for DEM-FRF and sells USD-DEM at 14%. He re-
ally disbursed the equivalent of one volatility, or sold correlation at 97%. It
is an application of the following formula:

O cross = yo? + o5 ~ 2p0,9,
from which is derived a correlation p of .97.

In the real world, correlations are unstable. Therefore, a trade that has a
better chance of benefiting from changes in correlation than losing from it
is of higher grade than a trade that presents the opposite characteristic.
[End OCR]*

## Page 470

*[Image OCR]
458 Modules

GENERAL CONVEXITY RULES

Risk Management Rule: Asecurity that is convex with respect toa
nonconstant parameter should, everything else being equal, be

worth more than a security that is not as convex with respect to the
same parameter.

In a flat-yield-curve environment and assuming the back-month matu-
rities have the same volatility, it is preferable to own the bond with the high-
est convexity.

A corollary is that a portfolio of options that presents a symmetric con-
vexity in its vega is worth more than a similar portfolio that presents none.

Example: Two portfolios are long vega in the amount of $100,000 per
volatility point (from 16% down to 15%). Portfolio A’s vega increases as
volatility rallies and decreases in a sell-off, while portfolio B’s vega de-
creases in the rally and increases in a sell-off. On that account alone, the
options that constitute portfolio A are more valuable than those that
constitute portfolio B.

It is preferable to be square vega and own the second derivative of
vega than have a position (through binary options hedged with vanilla)
that loses money when vega moves. These instruments are therefore
deemed “inferior.”

Risk Management Rule: A portfolio of options that presents sym-
metric convexity in its vega with respect to underlying A is worth
more than a similar portfolio of options of the same symmetric con-
vexity with respect to underlying B if underlying B is less het-

eroskedastic than underlying A.

A portfolio of options long vega through shorting a barrier op-
tion is less valuable than a portfolio of options that is long vega
through out-of-the-money options.

Figure F.1 shows the difference between convex and concave vega.

NE — Convex Vega
On
ee
Concave ven

Figure F.1 Convex and concave vega.
[End OCR]*

## Page 471

*[Image OCR]
Module G

Option Pricing

Iro’s LEMMA EXPLAINED

Ito’s lemma is an important tool in option pricing. Without spending an ex-
cessive amount of time on theory, this module will provide a heuristic defi-
nition, couching as much of it as possible in terms practitioners can
understand.

The gist of Ito’s method is that a function of a random walk (not
smooth) is smooth and differentiable. The module of the Brownian motion
on a spreadsheet (Module A) goes into the details of the lack of smooth-
ness.

The random walk has already been intuitively defined:

AW = W(t + At) — W(t) = w(W, BAt + o(W, fAZ (1)

with AZ = VAt U(0, 1), U(0, 1) ani.i.d. process following a normal distribu-
tion! with mean 0 and unit variance. The shortcuts o and w for a(W, t) and
w(W, t) will be used.

Equation (1) means that the change in W, over time At, is composed of a
drift and of a stochastic element the magnitude of which is determined by
the variance. .

It is established that no matter how small one defines the increments At,
the function will not be smooth anywhere. It will remain jagged and
nowhere differentiable. This is analogous to a coastline that no one can ac-
curately measure because of its infinite jaggedness.

Select time in tranches so small that any smaller increment would be 0.
Anything that is multiplied by time would vanish.

It is possible to check that:

E(AW) = pAt since E(AZ) = 0
V(AW) = E{AW) — E(AW)2}2 = E(0 + o2At U2}
so

V(AW) = o2At

459
[End OCR]*

## Page 472

*[Image OCR]
460 Modules

this leads to the following Ito multiplication table:

dt dw
at 0 0
dw 0 o*dt

From here, (1) can be written as what is known as an Ito process:

dW = u(W, dt + o(W, tdZ (2)

at the limit. This differential form should be always seen as a shorthand for
a stochastic integral, not a true partial differential equation.

Let F(W, t) be a security that is function of both W and time. Expanding,
leads to:

aF oF 1 @F 1 #F
F = — dt+ ——dw+ — 24 — (dt)
d ot at aw aw 2 20W? (aw) 2 df? (at)
o°F
+ dW dt 3
aW ar | ) (3)

since:

1. Unlike ordinary calculus, the stochastic expansion does not stop at
dW since dW? does not vanish, like the coastline that remains jagged
at all scales and

2. Anything multiplied by dt vanishes,

so (3) becomes

oF oF 1 &F
dF =  at+ S— aw+ — “= (aw) 4
ot aw aw 2 aw (aw) (4)

expanding dW and (dW)? results in

oF

oF | 1 @F
dF=|" + wy + =
Fe mW, OW

ww wlay . 8E-
> aye OW | dt+ —<—o(W,)dW 5)

It is easy to see how this can lead to the Black-Scholes equatién.
As a first step, consider that dS/S is an Ito process:

& = pdt + odZ (6)
Examine the process of the log dS with the aid of the Ito transform:
2 2
d LogS = 2 + er and since (=) = (ydt + odZ)* = 07(dZ)* = o°dt

o
d Logs = ( - 5) dt + odW
[End OCR]*

## Page 473

*[Image OCR]
Option Pricing 461
S,) a period shorter than period ¢ (usually ft, is the present).

t 2 t
Logs, ~ Logs) = | (u - +] dr + of aw,

o?
= (uw — 5 = 6) + (W, ~ Wy)
This leads to:

S, = 5,.€ (» ~ %) (f ~ to) + a(t — ig) )

Equation (7) is used with most of the option pricing tools in this book. In
addition, (7) satisfies

E,(S,) = $6! '0) | ent 1)!" n(u)du (8)

= S, e (we — °%) (t ~ ig) erntt — to)
0

since the right integral in (8) is the moment-generating function M of the
Gaussian distribution

[ex ®"n(ujdu = M (u) = 27-40)

ot - tg)
hence .

E,(S,) = 5,00 ~ t9) (9)
With the results, operators can proceed to price options, in several ways.

¢ By using the Black-Scholes (1973) argument of delta neutrality lead-
ing us to ignore the utility curve of operators and possible risk pre-
miums establishing that 4 =r. The argument will be displayed
further down.

¢ A more modern method consists in using the Harrison and Kreps
(1979) and Harrison and Pliska (1981) generalization that allows the
extension of the previous argument to any type of contingent claims,
under some conditions of market completeness (as far as all of the in-
strument affecting the derivative are concerned) that can easily be
summarized as: allowing full replication through dynamic (hence static)
hedging.

¢ Other methods, such as Feynman-Kac, which are being used in mod-
ern finance to integrate Ito’s lemma and derive an underlying pro-
cess under risk-neutralized paths.
[End OCR]*

## Page 474

*[Image OCR]
462 Modules

The Feynman-Kac solution of a large class of stochastic differential equa-
‘tions as the probabilistic expectation of a function (under some regularity
conditions on the drift and the variance, see Dana and Jeanblanc-Piqué, 1994
for a clear review) allows operators to use probabilistic methods, rather
than the more burdening (and less intuitive) partial differential equations,
in order to tackle option pricing. More simply, one can price a path indepen-
dent option (and a large class of soft path dependent options such as barri-
ers) as the expectation of the terminal payoff assuming that the asset price
follows a risk neutral diffusion. The solution is then said that

exp (—r (t — £,)) E2 {f(S)}
for a path independent option, and
exp (—r(t — f))) E2{ f(S)/t < 7}

for a soft path dependent one (like a barrier), with Q the risk neutral proba-
bility measure, f(S) the final payoff function, f the time to expiration and +
the stopping time (when hitting the barrier). It facilitates Monte Carlo meth-
ods (of averaging the payoffs in a series of random paths). It presents the spe-
cial advantage of allowing the use of numerical integration, a method which
this author finds extremely flexible to program (owing to the widely avail-
able subroutines). Numerical integration can be lengthier in computer time
but saves in considerable programming labor and presents a smaller error
rate.

The operator can go about his business, provided he can ascertain that
E, (S(t) is risk neutral.

Ito’s Lemma for Two Assets

It is easy to generalize Ito’s lemma to multiple assets (Table G.1)
The example can be speeded up by showing the results of a function of
two Ito processes and time:

oF oF °F
——dw,+ aw, + ———-
aw, ' aw, 7  aW,aw,
with o, and o, the volatility of each asset W, and W, and p the correlation
between their moves.

s
dF =F gee dw.dw
ot 1 2

Table G.1 Two-Asset Ito Multiplication Table

dt dw, dw,
dt 0 0 0
dw, 0 o,’ dt po,o,dt
dw, 0 po,o,dt o,7dt
[End OCR]*

## Page 475

*[Image OCR]
Option Pricing 463
BLACK-SCHOLES EQUATION

The Risk-Neutral Argument

First it is necessary to get the drift out of the equation. Say that the operator
sold a European call C on a dividend paying asset S and bought a bond B
with the proceeds.” The bond pays r, the underlying asset is 1, the asset is
expected to return p rate of return. The call is defined as the Max(S-K, 0)
and expires at period t.

Say that the operator would remain delta neutral at all times. So he
would have on his balance sheet the following portfolio P composed of:

P=—C+0C/faSS+B=0
hence
B=C—90C/as §
and would have B equal of the difference between the cash raised from the
- call and the dollars invested in the stock S. The portfolio earns interest at
the risk-free rate and earns dividends on the stock held. Setting AP = 0 for
the portfolio to be cash-flow insensitive, gives:
—AC + aC/aS X AS + AB=0
so, for infinitesimally small increments and using Ito’s expansion:

—dC/at dt — aC/aS dS — 2 a C/aS7dS? + aC/aS x dS + dB =0

with
dS =S(ydt + o dW)
dS? = S* (wu dt + odZ)* = S*o*dt from the Ito multiplication tables 5
dB = rB, the interest earned on the bond portfolio (or paid, if negative)

dB =(rC — dC/aS Sr — aC/aS S d) dt =r Cdt — dC/dS S (r — d) dt
Thus:
—dC/at dt — 20°C /dS? S*0°dt +r C dt ~ aC/aS S (r.— d) dt = 0
hence:

aC/at + 82C/ dS? S’o? — 1 C + aC/aS S (r — d) =0 (10)
[End OCR]*

## Page 476

*[Image OCR]
464 Modules

One can see dS coming out of the equation, along with w. The only rates left
are the risk-free and the payout rate.

From here on, m will be called the risk-neutral drift, equivalent to r — d.
The operator can proceed and price the differential equation (10) under
boundary conditions. Or he could integrate the diffusion to yield the same
results as before in (7) but with (r — d) in place of w:

C = exp (— rf — t,)) [max (S,, exp(r — d) - & (t — ty)

+ o\(t — t,) x — K,) n(x)d(x)

x being a centered Gaussian random variable. The solution for the integral
from here is tedious but only requires slight manipulations. Users end up
with the value of the call, using ¢, = 0 to simplify as:

C = exp( — dt) S,N(d1) — exp( — rt) K N(d2)

with:
di = [log(S,/K) + (r — d)t]/loVt] + oVt/2
d2 = [log(S,/K) + (r — d)t]/[oVt] — oV't/2

A numeraire change will easily give the put value. The put/call parity rules
can also be used to derive it as a mirror image to the call.

Following are the formulas and techniques used in the book. Most of the
thrust was on numerical methods as the author—thanks to a Pentium chip
received as a birthday gift—did not need to find closed form solutions to
many of the exotic options and the stochastic volatility model.

Most of the numerical methods implemented were part of a Mathemat-
ica™ Nintegrate command that is based on the Konrod quadrature method.
One integrates between —5 and 5 in lieu of positive or negative infinity as the
error beyond becomes very small. These methods are outperforming in cases
where early exercise is excluded and mix rather well with options priced par-
tially analytically. Being still a trader, the author did not have to subscribe to
academic elegance and drew on the comfort of computer power.

STOCHASTIC VOLATILITY MODEL

It is assumed that both the security and the volatility follow a Brownian
motion. This is to illustrate the effect of heteroskedasticity on option prices.
The process is being used for simple purpose: to gauge the value of the port-
folio at expiration assuming the operator hedged both the delta and the
gamma of the options. It is inspired by stochastic volatility models, such as
[End OCR]*

## Page 477

*[Image OCR]
Option Pricing 465

Hull and White? (1987), that did attempt to replicate the volatility in a risk-
neutral way (through the buying and selling of options), akin to the Black-
Scholes method for the asset itself. This formula assumes that there exists a
risk-neutral replication, which allows for Feynman-Kac final payoff pricing.
As the model below assumes independence between asset return and
volatility, a numerical “fudge” is provided that considers that the quadratic
norm of the volatility process

4/1 op
o, = Po, 1074

could be numerically computed as one single evolution between f, and t.

f 1
S, = S,exp (r — d)(t — fy) — > | o2as+ er Sin 734s V(t — t))z
0

4/ 1 t 1 ,
t— j, Jo 7348 = oy, exp(~ZV?E— hy) + V Vt - £2’)

ris the risk-free rate, d the counterrate (foreign rate for a currency, dividend
rate for a stock, etc.), S, the asset price at time t, o(t) the volatility of the asset
at time t; V is the standard deviation of volatility; z and z’ are Wiener pro-
cesses each independent and normally distributed with 0 mean and unit
variance.

For simplification, assume that {, = 0. The no-free lunch rules must be
satisfied:

E,(S,) =| | Sn(z)n(z’)dzdz’ = Se — 4!
E,(o,) = [. on(z’)dz’ = 0,

n(z) is the normal density function of z with mean 0 and variance 1.
Another simplification is to make the expectation uniform with time, in
order not to include a term structure of volatility. This allows the use of a
model for the values from period t only.
The price of the European option follows:

option(S,, K, o, r, d, vvol)
=exp(-rt){ [  Max(®S, — ®K, 0)n(2)n(z\dzdz’
With ® = 1 when the option is a call and ® = —1 when the option is a

put. The double integral could be simplified by integrating a Black-Scholes-
Merton call over different o,,.
[End OCR]*

## Page 478

*[Image OCR]
466 Modules

As usual, numerical integration methods are used, using the Mathemat-
ica™ that this author received for the holidays. The table of the results was
presented in Chapter 15.

MULTIASSET OPTIONS

Take two assets A and B. The risk neutral process will be:

S(t) = S,, exp((r - d, - Soult — t,) +0, yb — fz)

S(t) = Spo expl((r — dy - So(tCt ~ t,) +o, VE=f(pz + V1 = p*2’))

r is the risk-free rate; d, the counterrate for asset A (foreign rate for a cur-
rency, dividend rate for a stock, etc.); S(t) and S,(t) the asset prices A and B
respectively at time f; 0, and o, the volatilities of assets A and B respec-
tively; z and z Wiener processes each independent and normally distrib-
uted with 0 mean and unit variance; p the instantaneous correlation
between A and B.

With the following model, traders can price several varieties of multias-
set options using expectation of final payoff thanks to Feynman-Kac.

To add a third asset C, the same process would take place, but with C de-
pending on a Wiener process.

1
S(t) = Sg exp ((r ~ d, ~ 5 OHNE ~ ty) + 0, yt ~ fy LZ)

where L is the (3,1) vector that is the 3"¢ row of C, the (3,3) lower triangular
Cholesky decomposition such that C C’ becomes the correlation matrix of
the returns of the 3 securities and Z is the vector composed of z, z’ and the
additional z’’’ for asset C.

Rainbow Options ’
They are options on two assets with one strike price:
option(S ,o(t), Spot), Ka, Kp, Oy, Spr 1 dy, Ap)
= exp( — rt) [f. Max(®,S ,(t) — ®, K,, ®,5,(¢) — ®,K,,,0)n(z)n(z )dzdz’
With ®, = 1 when the option is a call and ®, = —1 when the option is a

put, n(z) the normal density function with mean 0 and variance 1. We as-
sume for simplification in the rest of the module that t, = 0.
[End OCR]*

## Page 479

*[Image OCR]
Option Pricing 467

Outperformance Options

option(S ,o(), Spo(t), Og, Og 1 Ayr Ap)

=exp(— rf [ Max(S,(8) ~ $,(0), S,(t) — $,(8))n(e)m(e ded’
Spread Options

option(S ,o(£), Spot), K, Og, Op, 1 Ay, Ap)
= exp( - r#){ ° ” Max((®S,,(t) — ®S,(f)) — ®K, O)n(z)n(z)dzdz’

with K the strike price of the spread defined as S, — S, and ® = 1 when
the option is a call on the spread and ® = —1 when the option is a put on
the spread.

COMPOUND AND CHOOSER ORDER OPTIONS

Unlike before, the formulas introduce a term structure of volatility, interest,
and carry rates. The same univariate process is used as before, with:

K, the strike price of the “mother” options.

®, equals 1 if the mother is a put, —1 if it is a call.
t, is the time to expiration of the mother option.

K, is the strike price of the “daughter” options.

®, equals 1 if the daughter is a put, —lifitisa call.
t, is the time to expiration of the daughter option.

In addition to the conventional models, the following will be included:

o,ando,,r, and r,, d, and d,, respectively the spot volatility and inter-
est rates until the expiration of the daughter and the forward (or for-
ward-forward) between the expiration of the daughter and that of the
mother.

These rates are not the spot rates in the market but need to be derived
using forward-forward break-even formulas, such as the forward volatility
described in Chapter 9.
[End OCR]*

## Page 480

*[Image OCR]
468 Modules

Compound Options

The compound option will be priced as:

Compound(S,, ®, K,, 01,1, 4, BK, F5, 1, d,)

very ~ 20°27

=| Max(®,BS(S,, ®,K, o,, 1, d,) — ,K,, O)n(2)dz

20°27

with BS a Black-Scholes-Merton vanilla and

1 7
S, = S,exp(r, - 4, “5 o? + 0, \t,2)

Chooser Options

Integrate the Max(Put, Call), using BSP as the Black-Scholes-Merton valued
put and BSC the call keeping the same notations as well as the same notion
of forward-forward interest rates and volatilities between the exercise date
of the mother and that of the daughter:

Chooser(S,, K, 04, 14, dy, Fy, Tyr Gy)
= | Max(BSC(S,, K, o,, 1, d,), BSP(S,, K, Gy, 7, d,))n(z)dz

with 1 _
S, = S,exp(r, — 4, “5 a? + o,\t,2)

BARRIER OPTIONS

There are two ways to price barrier options. The first, most intuitive, way, is
to approach it through stopping time and say that the option is the Feynman-
Kac expectation under risk neutral probability Q, times the risk neutral prob-
ability of the option not being knocked-out before expiration (which is
equivalent to the probability of stopping time being longer than expiration).
The second, more convenient, is to use the reflection principle as described in
Chapter 19, which allows for simple results with single barriers.

It is worthy to note that the study of the joint distribution of the triplet
(upper range, lower range, Brownian motion) was investigated early on in
probability theory. Its premonition was attributed to a later (and no less
precursor) Bachelier by Geman and Yor (1996). More recent results for op-
tion pricing were obtained by Kunitomo and Ikeda (1992).

Thanks to the reflection principle and the Girsanov theorem, one can
use closed form solutions for barriers. There are two classes of pricing for-
mulas for barriers:
[End OCR]*

## Page 481

*[Image OCR]
Option Pricing 469

1. Closed form Black-Scholes-Merton based formulas, unfortunately
weak because they assume constant volatility and constant interest
rates between time 0 and expiration, a method that misprices options
when there is either a sloping implied volatility or interest rate curve.

2. Numerical methods, none of which is presently in the public do-
main. They range from heavy-duty Monte Carlo simulators to trees
with accurate local volatility between the nodes.

The Reflection Principle

As described in Chapter 19, the reflection principle allows users to deter-
mine the number of conditional Brownian paths between two points. It will
be used here to compute the risk-neutral probability density of the paths
that reach a given destination without going through the barrier. They will be
the complement (i.e., the sum equals 100%) of:

* The number of paths that reach a destination going through the bar-
rier (the difference between knock-in and knock-out).

¢ The number of paths that do not reach the final destination.

Let W, be defined as a standard Brownian motion (0 drift) and / a given
limit underneath it. The reflection principle (see Karatzas and Shreve, 1991;
Grimmet and Stirzaker, 1992; Lamberton and Lapeyre, 1991) allows the fol-
lowing transformation:

W%, = W, if the barrier is not touched,

and

W, = 21 — W, if the barrier is touched.

Hence, options can be priced using the difference between the two processes
as described in Chapter 19. However, a hitch would appear if there is a drift
which forces us to use the Girsanov theorem.

Girsanov’s Theorem

What Girsanov’s theorem allows users is to create a new fictitious probabil-
ity density to replace the previous one, the new density being the risk-
neutral one.* Most presentations of Girsanov present arduous features
whereas, explained to traders, it appears rather simpler to digest once
stripped of measure theory. A stripped-down version of the theorem is pre-
sented here as it becomes a lot less complicated in situations where the drift
is constant.
[End OCR]*

## Page 482

*[Image OCR]
470 Modules

The principal result is that it eliminates the drift and, instead,
integrates it within the probability distribution by shifting the expecta-
tion accordingly.

Define W, as a standard Brownian motion with (0 drift and unit vari-
ance) under a probability measure P.

Girsanov’s theorem establishes that the process defined by

W = + W,

a standard Brownian motion under probability measure Q, with Radon-
Nikodym derivative:

dQ, /dP = exp (12 Nt — AW',)

A function V of the payoff of a Brownian motion (i.e., a derivative secu-
rity) will be priced so its expectation under probability measure Q is the
expectation under probability measure P times the derivative of the two
measures:

E°(V) = EP(dQ,/aP V)

so
E(V) = EP(exp (‘At — XW’) V)

With these elements, the reader can build the pricing intuition for bar-
rier options.

Pricing Barriers

The results below are inspired by Douady (1996). ©
The trader begins with the call up-and-out and builds the intuition of
the pricing:
%
¢ Call up-and-out CUO = Under probability measure P (risk-neutral),
the discounted, by exp(— rt), expectation of the Max (S — K, 0) as be-
fore with the Black-Scholes-Merton formula, but conditional on the
barrier H, not being touched.

Therefore, the trader should start with S,<H; otherwise the option

would be terminated from the beginning. Hence, using Girsanov, the process
W’, = I/o log(S,/S,) is a Brownian motion without drift under the probability:

Q,= exp( 5 - aw’) P

with A = (r — d)/a — 0/2
[End OCR]*

## Page 483

*[Image OCR]
Option Pricing 471

Hence,

¢ CUO = Under probability measure Q, exp(— rt) times the expecta-

tion of the exp(— ‘Nt — d/alog(S,/S,)) times (Max (S — K, 0), condi-
tional on the barrier H not being touched.

Using the reflection principle, the density of the process (under Q) of
S, is, conditional on the barrier H,, not being touched, equal to that
of the process of H?/S..

This allows jumping to:

¢ CUO = Expected (discounted) price of distribution of the S,>K

minus the expected price of the paths that touch the barrier H:
CUO = exp(— rt) (P1 — P2),
with

Pl = [ exp(-5 Nt + AVE Zz) (S exp(o Vez — K) n(z)dz
k

P= [ exp(-5 Mt + Vet z) (Sp exp(o Vez - K) n(2h - z)dz
k
with
h = (1/(oV¢ ))log (H/S,)
k = (1/(oV¢ })log (K/S,)
Setting ;
N=Ato=(r-d/at+a/2
a= (H/S,)?’°

a’ = (H/S,)*/"
L=NVE-k

d,=hVt—k

d,='Vt-h

d,= Vt -—h

d,=—-dVt-h
d= —-NVt-h
d,=-dAVt-2h+k
d,=—\'Vt-2h+k
[End OCR]*

## Page 484

*[Image OCR]
472 Modules

and N(d,) = cumulative normal distribution of d,

[ exp( — 5 Mt + AVEz) n(z) dz = N( -d,) — N(-d,) = N(@,) ~ NG)
[exp( -F Mt + X'Vtz) n(z) dz = exp{(r — d)t}(N(@,) — N(d,))

[ oxp(-t Nt + AVEz) n(2h — 2) dz = a(N(d,) — N(d,))

[exp(-4 Mt + \’VEz) n(2h — 2) dz = a’ exp((2— d)f)(N(d,) — N(d,))

This leads to the following formulas.
Calls Up-and-Out.

CUO = exp{( —d)t}5,{N(d,) — N(d,) — «'(N(d,)
— N(d,))} — K exp(— rt){N(d,) — N(d,)
— a(N(d.) — N(d,))}

Using the technique, we can extend to all other options.

Calls Up-and-In. Calls up-and-in are synthetically vanilla calls minus
calls up-and-out. Hence:

CUI = exp{(—d)t} S, {N(d,) + a’(N(d,) — N(@,))|
— K exp(— rt){N(d,) + o(N(d,) — N(d,))]

Puts Down-and-Out. Puts down-and-out are priced like the call wp-and-
out but by an exact change of numeraire, which reverses the sign of the
integral. Hence:

PDO = K exp(— r #){N(d,) — N(d,) — a(N(@,) — N(@,))}
— exp(—dt) S{N(d,) ~ N(d,) — a'(N(d,) — N(d,))}
Puts Down-and-In.

PDI = K exp(— r #){1 ~ N(d,) + a(N(d,) — N(d,))}
— exp(—d)t S,{1 — N(d,) — a'(N(d,) — N(d,))}
[End OCR]*

## Page 485

*[Image OCR]
Option Pricing 473
Calls Down-and-Out. Calls down-and-out will be priced in two different
ways: When they are out-of-the-money (i.e., the strike K is higher than H) or
when they knock-out with intrinsic.

Case 1: K =H

CDO, = 1 = exp{(—d)#S,{N(d,) — «(1 — N(d,))}
— Kexp(— rt){N@,) — a(1 — N(d@,))}

Case 2: K =H

CDO, <4 = exp{(—d)t} S, {N(d,) — «(1 — N(d,))}
~ Kexp(— rt)(N(d,) — «(1 — N(d,))}

Calls Down-and-In. The same is true with calls down-and-in:

Case 1: K =H

CDI,» 4 = exp{(—a)t} S, «(1 — N(d,))
— Kexp(— rt) a(1 — NG@,))

Case2:K =H
CDI, ., = exp{(—d)t}S,{N(d,) — N(d,) + a'(1 — N(d,))}
— Kexp(— rt){N(@,) — N(d,) + a(1 — N(d,))}
Puts Up-and-Out.
Case 1: K =H
PUO, .., = K exp(—r t){1 — N(d,) — a(N(d,) — N(d,))} .
— exp(—d)t S,{1 — N(d,) — a’(N(d,))}
Case2: K=H

PUO = K exp(~r £){1 — N@,) — a N(d@,)}

— exp(—d)t S,{1 — N(d,) — «’N(d,)}

K=H

Puts Up-and-In. It is a natural option algebra equivalent to the vanilla
minus the PUO. We know that if 5, > H it will be considered knocked-in
[End OCR]*

## Page 486

*[Image OCR]
474 Modules

and would be valued at exactly the same equivalent price as the vanilla option.
Otherwise it will be:

Case 1: K =H

PUL, .,, = K exp(—r t){N(d,) — N(@,) + eN(d,))}
— exp(—dt)S,{N(d,) — N(d,) + a'N(d,)}

Case 2:K =H

PUI, ., = Klexp( — r t)aN(d,)} — exp(—dt)S,{1 — N(d,) — a’ N(d,)}

Rebates and American Barrier Options. The SDF of a barrier cum rebate is
a barrier without rebate (priced with the formulas above) plus an American
binary option.

Calls Double Knock-Out. Geman and Yor (1996) offer a new methodology
based on excursions theory to provide the Laplace transform of double bar-
rier options, which can be inverted using the Geman-Eydeland technique.
Below is the earlier Kunitomo-Ikeda model. Use a high barrier H,, and a low
barrier H,.

Set:

h,, = 1/0Vt log (H,,/S,)
h, = 1/0V't log (H,/S,)
8 =h, —h,
Case 1:H, < K<H,,

CDBiy cx <Hn = eXP(—4t)S, S TN Vt) - TAM Vt)

n=—e 5

— exp(— rt) K Ss LAV) — J,0.V2)
with

I(x) = exp (— 2nx8) (N(h, + 2n8 — x) — N(k + 2nb’ — x))
J,(x) =exp{2x(nd + h,)} (N(2h, ~ k + 2nd + x) — N(h, + 2nd — x))

Note that the summation converges rapidly and —~ can be replaced by
—12, + by +12.
[End OCR]*

## Page 487

*[Image OCR]
Option Pricing 475

Case 2: K = A,

CDB,, <Hy = exp(—dt)S, S TY’ Vt) _— JN Vt)

n=—o
oo

—exp(—rt)K > LAV) — FAV)

n=—

(same as before) with

I (x) = exp(— 2nxd) (N(h,, + 2n’ — x) — N(h, + (2n — 1)8 — x))
J Ax) = exp{2x(nb + h,,) (N(A,, + (2n + 1)8 + x) — N(h,, + 2nd + x))

Puts Double Knock-Out.
Case 1: H, <K<H,

PDB, <x <Hy = &XP(~ Tt) K S LAV) — J,A0V6)

n= —

~ exp(—dt)S, 5° (NV) - 7,0 V8)
with

I(x) = exp(— 2nx8) (N(k + 2nd — x) — N(h, + (2n — 1)8 — x))
] Ax) = exp{2x(nd + h,) (N(h, + (2n + 1)8 + x) — N(2h,, - k + 2nd + x))

Case 2: H,, =K

PDB,,,, <x = exP(— rt) K S LAV) ~ J,4V2)

n=—o2

- exp(—dt)s, 5° LO’ V2) — J, Vb)

n= —e%

(same as before) with

I(x) = exp(— 2nx8) (N(h,, + 2n8 — x) ~ N(h,, + (2n - 1)8 — x)
J,(x) = exp{2x(nd + h,,) (Ny, + (2n + 1)8 + x) — N(h,, + 2nd + x))

Pricing Double Binary Options. In this book, the author tricked the pre-
ceding formulas for double barriers by pricing for situations where the op-
tion is deep in the money (the payoff becomes equal to intrinsic at trigger
time). Deep calls were selected with strike close to 0.
[End OCR]*

## Page 488

*[Image OCR]
476 Modules

Stopping Time and Its Expectation. The author avoided introducing stop-
ping time in the analysis of barrier options in order to build on the intuitive
approach of the reflection principle. Most papers introduce it in the deriva-
tion of the pricing of any barrier. The following formula is the calculation
for the stopping time and its expectation.

Take as before \ = (r — d)/o — o/2 and the barrier h = 1/0 log H/S,,
with H the barrier. The density of the unconditional stopping time is given
with no drift:

P(t) = (h/ W283) exp(—h?/2t)
and, adding the drift:
P(t) = (h/V2m8°) exp(dh — N2t/2 ~ h2/2t)

As to the expectation of the exit time, that is the distribution of the mini-
mum of t the exit time and the time to expiration T:
E(o7,,) = h/d + (T — b/d) N(A/ VT — XVT)
— exp(2dh)(T + h/d) N(-h/VT — XVT)

With a double barrier, the following can be computed.
h =I/o log H/S,
1=1/o log L/S,

The following formula provides the density of stopping line (starting
inside the range).

Pit) = — exp(-*- t) 2,(-1)" —1y exp(-

nat
(h - 1?

2(h — 1)?

s

i
h-l

(exp())sin nt a4 — exp(Ah)sin nt )
As to its expectation, the user can break it up
E(mint},,) = E(t,,)— Dy

where

E (ty) = (t — 2D) (exp(ah) sinh Ah — exp(Al) sinh XJ) cosh A(h — 1)
— (h exp(AJ) cosh Ah — I exp(Ah) cosh AJ) sinh A (h — 1))/Asinh? A(h — J)
[End OCR]*

## Page 489

*[Image OCR]
Option Pricing 477

or < _ exp(—u,T) _ nah
D,= 2 (-1) ln rr (exp(Al) sin h-l

n

nil

— exp(Ah) sin h

7)
and >>
nee
= a SO +H 2
“= 3h—me

NUMERICAL STOCHASTIC INTEGRATION: A SAMPLE

A Mathematica™ Program

This program illustrates a general option pricing method that will gain in
currency as the computer chip war rages on. The author used such integra-
tion techniques to find the numerical expectation of a stochastic integral.

(*Homoskedastic compound options*)

(*Nassim Taleb*)

gauss[x_]: = Exp[—x42/2]/(Sqrt[2*Pi]);

Gauss[x_]: = (1 + Erf[x/Sqrt[2]])/2;

St[S_,x_,tl_,sig_]: = S Exp [sig Sqrt[t1] x]

(*computing Black-Scholes with no drift*)

di[S_,k_,sig_,t1_]: = (Log[S/k] + sig2(t1)/2)/(sig*Sqrt[t1]);
d2(S_,k_,sig_,t1_]: = (Log[S/k] — sig*2(t1)/2)/(sig*Sqrt[t1]);
call[S_,k_,sig_,tl1_]: = S*Gauss[d1[S,k,sig,t1]] — k*Gauss[d2[S,k,sig,t1]]
put[S_,k_,sig_,t1_]: = call[S,k,sig,t1] — (S — k)

(*compound option*)

callcallpayoff[S_,k_,kopt_,sig_,x_,tl_,tint_]: = Max{call[S Exp[sig
Sqrt[tint] x],k,sig,(t1 — tint)] — kopt,0]
callcall{S_,k_,kopt_,sig_,tl_tint_]: = Nintegrate[ callcallpayoff[S,k,kopt,
sig,x,t1,tint] gauss[x], {x, —4,4}]

(*for precision-minded people increase the integration band to ~6 to 6*)
(‘adding the drift is straight forward”)
[End OCR]*

## Page 490

*[Image OCR]
Notes

PREFACE

1. This represents about 60 transactions per business day and the examination of about
two batches of reports per day (with an average of about 17 different reports repre-
senting the different markets in which I ran option books). The total corresponds to
95,000 trades in the currency options, 30,000 in equity indices, 30,000 in Eurodeposits
(Eurodollars, PIBOR, Euroyen, Euromarks), 1,000 in commodities (live cattle, oil, . . .),
and the rest in swaptions and long bond options. I traded only 2 options on mort-
gages. (I opened and closed the trade disgustingly on the same day).

2. Throughout this book, see Bibliography for full reference citations: See Soros (1987),
Derman (1996).

3. Black-Scholes (1973); Merton (1973).

There are very few books written by practitioners. One notable entry-level text used
by traders is Natenberg (1995). The reader would find the prerequisites to understand
option trading from the standpoint of the “producer” (as opposed to the consumer or
the amateur theorist).

5. This rescaling is compatible with the analytical option pricing formulas, (see Merton,
1973, for the theorems):

Black-Scholes-Merton value V (AS, \K, o, 7, d) = XV (S, K, 9,1, 4)

(Black-Scholes-Merton is homogeneous degree one in the underlying asset price and
the strike price.)

It is also homogeneous degree one in S, K, H,, H,, (H, is the high barrier, H, is a
low barrier ):

V (AS, AK, MH AH, 0, 1, d) = XV (S, K, Hy, Hy, 0,1, d)

INTRODUCTION

1. Leland (1985) provides a clear review of the literature on the issue of discrete timte
dynamic hedging. He shows that the replicating portfolio value, composed of long the
asset and short the option will not be 0 at all times once we start hedging irfre-
quently. The reader will see in the Module F that the Black-Scholes price is derived
from a strategy that would make the P/L stay at 0 the entire time. In reality it is not
possible to buy and sell continuously. So in place of the continuous di we opt for Af,
the time lapse between hedging revisions. The P/L will be, however, 0 in expectation.
The variance of the package, he verified, decreased by V2 when At was cut by half.
Similar results can be obtained using the functional central limit theorems by show-
ing that as the sample paths is divided in infinitely small fragments the variance of
the P/L would (uniformly, no less) converge to 0.

One alarming fact for modern financial theory is that the package has a variance.
It looks just like any other risky security. In other words: option replication has risks,
which requires compensation.

479
[End OCR]*

## Page 491

*[Image OCR]
480

10.

Notes

CHAPTER 1 INTRODUCTION TO THE INSTRUMENTS

Trading them, however, can be done in such a way as to cause dependence on the cor-
relation between parts and make them akin to multiasset options.

The bond future includes a hidden option (the right to choose).

Dixit and Pyndick (1994) study the application of option theory to corporate financial
management.

Most traders think in terms of the forward while the quants think in terms of the cash
since the Black-Scholes (1973) formula views an option price as a function of the cash
price of the asset. Merton’s (1973) approach is more compatible with traders’ meth-
ods. This distortion can lead to mishedging in cases of currency bands.

American options complicate the solution of the Black-Scholes-Merton partial differ-
ential equation under, among other boundary conditions, the early exercise rules of
Q(#), an optimal exercise path where the derivative security is equal to @(S — K). Such
time-dependent boundary condition is causing much ink to flow. Worse, there might
be two boundaries as will be shown.

More complex analysis would use stochastic interest rates and take into account the
possibility of changes in the structure of the interest rate differential. On that, see
Amin and Jarrow (1991) for a treatment of currency option pricing with stochastic in-
terest rates.

There is, ironically, a flood of academic papers attempting to value American options
precisely while parameters affecting their value are imprecise.

A simple application of Ito’s lemma gives these results.

By convention, any rate r in this book will be the continuously compounded zero
coupon rate. :

While Eurodollars are quoted in 100 — yield for all purpose of analysis (like the com-
putation of volatility and correlations), we use the derived yield as the variable.

CHAPTER 2. THE GENERALIZED OPTION

In Durett (1991).

J. Piper, O’Connel, and Piper seminar, Chicago, December 4, 1995. It is remarkable
that some veteran equity traders fear the “bad distribution” to the point of refusing
to touch a put. Such phobias are part of the traders survival toolkit.

More advanced analysis would show that it harbors a measure of shadow gamma and
shadow delta, characteristics that will be discussed later in this book.

CHAPTER 3 MARKET MAKING AND MARKET USING

This case is easily replicated in the real world where the trader buys a one-year at-the-
money European binary option for $2.5 mil paying either zero or $5 mil at expiration if
the SP500 closes at or above 668.00.

See Grossman and Miller (1988).
See Ho and Stroll (1980); Silber (1984); Garman (1976) and O’Hara (1995) for a review.
See D. Guillaume et al. (1995), Roll (1984).
[End OCR]*

## Page 492

*[Image OCR]
10.

11.

12.
13.

Chapter4 481

Sampling of six currency market makers by the author, June 1995.
See Feller I (1968).
A clear way to view it is to compare the P/L of a trader to a Brownian motion with a
drift. The real “skills” of the trader would correspond to the drift.

Let P be the profits of the trader, with P, = 0, Z a Wiener process and o the
volatility of the trader’s P/L:

dP = pdt + odZ

an arithmetic Brownian motion (the P can be allowed to be negative). The reader
should view the » as the “edge”; if he believes he has a 51% against 49% edge on the
market during a given period of time, the drift is 2% of the net amounts involved. The
a is simply the volatility of the market in which the trader has the drift. It becomes
clear that the volatility of the trader will dominate the drift in most cases. Counterin-
tuitively, the o will typically drown the p.

Private conversation with the author, New York, April 1996.

See Dubins and Savage (1965). Also see Billingsley (1986).

A submartingale (the expected wealth next period is higher than that of the
present).

A supermartingale (the expected wealth next period is lower than that of the
present).

See Feller I (1968).

In one of the essays in Taleb (1997), I show that traders subjected to a stop loss pre-
sent a high-right moment, a skewness with a small sample property that weakens the
power of the Sharpe ratio.

CHAPTER 4 LIQUIDITY AND LIQUIDITY HOLES

There is an abundant market microstructure literature. The goal in this chapter is to
portray liquidity from a practitioner’s vantage point.

See the excellent and comprehensive review in Brock and de Lima (1995).

See Brock, Hsieh, and LeBaron (1991).

A Markov process is formally defined as: P(X, = s/X, = Xy X, = Xp... X= Xy—4)
= P(X, =s/X,_, =X, ,) for all n 21. It means that the conditional probability of
the next price X, depends only on the price preceding it, not the sequence of prices
that led there.

At the time of this writing, it is five minutes of suspension if the open is within 500
points from the previous close, then half an hour if it is within 2000 points, then until
the end of the session if the market hits the 3000 points third limit.

See Leland (1992).

Grossman (1988) distinguishes between strategies and securities and shows how an
Arrow-Debreu state of nature cannot be replicated dynamically. Because of the way
trades convey information, a real option will therefore present a different payoff than
a program traded replication.

Implicit bid-offer spread is the true bid offer spread, and is wider than the visible, or
explicit, bid-offer spread. It corresponds to the adjustment of the prices upward or
downward in response to a large order.
[End OCR]*

## Page 493

*[Image OCR]
482 Notes

10.

11.
12.

CHAPTER 5 ARBITRAGE AND THE ARBITRAGEURS

These are applied definitions of arbitrage largely stripped of the dimension in micro-
economics and financial theory. The text will introduce the notion of self-financing
strategies in connection with option values.

Oxford English Dictionary (1971).

If there is a multivariate Ito process then the link is stochastic (see Module G) in
terms of Z, through Z,. If both securities share the same random component, then the
link is deterministic.

It is uncertain whether one can use risk neutral probabilities when pricing the value
of a passive arbitrage.

CHAPTER 6 VOLATILITY AND CORRELATION

The simplification is pedagogical but valid at the limit of very smal] increments. The
process for asset prices dS/S = yp dt + o dz with boundary S(t,) = S, has for solution
log (S,/S,) = (u —ho*)(t — t) +o Vt—#Z, hence S,=S, exp ((w - Yo) (t — ty) +
o Vt — £,Z)) with S the asset price, o the volatility, tf time, and Z a random variable
normally distributed with unit variance and zero mean.

Some assets, according to Raphael Douady, can exhibit overgeometric behavior, with
their volatility increasing in higher levels and decreasing at lower ones.

See Cox and Rubinstein (1985).

Some traders use yp, the risk-neutral drift in a commodity, as a mean, in accordance
with the Black-Scholes-Merton prescription.

See Hamilton (1994) for a presentation of state-space models.

See Parkinson (1980); see Atilgan (1996). Notice that Cox and Rubinstein (1985) calcu-
lates the estimation of volatility based on the average S,,/S,, which biases the estima-
tor downward by a factor of 91%.

Option traders call gamma trading the subsequent delta adjustments emanating from
a long or short gamma position.

See Taleb (1996b).

Lo and McKinley (1988).

There are many potent methods, such as the Brock-Dechert-Scheinckman BDS test of
nonlinearity. The variance ratio is satisfactory enough for traders and risk managers
to gain a knowledge of the structure of the volatility. See Brock, Hsieh, and LeBaron
(1991) for a complete analysis of nonlinear dependence (including a presentation of
the correlation integral).

See Kritzman (1994).

One should note that the variance ratio is another way to look at autocorrelation. It cor-
responds to the summation of the autocovariance function over the corresponding lags.
For a good intuition, see Harvey (1993) and the excellent book by Bloomfield (1976).

CHAPTER 7 ADAPTING BLACK-SCHOLES-
MERTON: THE DELTA

See Black (1988).

If pair (N, A) has a constant volatility (N is a numeraire) and pair (N, B) has a constant
volatility, and pair (B, A) has a constant volatility, then the correlation between (N, A)
[End OCR]*

## Page 494

*[Image OCR]
Chapter 9 483

and (N, B) will be constant. When cross options trade heteroskedasticity means un-
stable correlations.

This is attributable to the issue of discrete-time convergence. Options, therefore, con-
verge in value to Black-Scholes-Merton but not in the higher derivatives.

This is another facet of the discrete-time convergence problem. A binomial leads toa
continuous time valuation of the option. But the delta and higher Greeks wil] not con-
verge. This issue is easier to see with a portfolio of options.

2\aw 6\ au

The Taylor expansion thus can be pushed to include, in addition to the delta, the

gamma—the DgammaDspot (i.e., the stability of the gamma) in the measurement. But

for a portfolio with many options, the expansion needs to cover higher and higher

order derivatives to track the risks. It becomes intuitive that the VAR does not apply

to a book. We will see in Chapter 11 why, since a book is not “compact” and since
higher orders or AU do not vanish, we need more powerful ways to look at the risks.

2 3
Cash equivalent position = ae au+ + (3 C Jaue 4 (4 c au etc.

CHAPTER 8 GAMMA AND SHADOW GAMMA

I owe the coining of the term “shadow gamma” to one of my former coworkers, Lenny
Dundennen.

Sig(x) is the change in volatility resulting from the changes in asset price between x,
and x.

Shadow gamma should not be confused with stochastic volatility techniques. Implied
volatility here is defined to be a direct, deterministic function of the move as predicted
by the trader and added to Black-Scholes-Merton. Traders can thus assert that volatil-
ity would be higher by 100 basis points should the market drop by 1% and establish a
map of volatility changes. Another way to view the difference is that shadow gamma
is concerned with implied volatility (as measured by Black-Scholes-Merton), whereas
stochastic volatility is concerned with a pricing model that tracks more accurately the
distribution of the actual volatility. The latter assumes that volatility is randomly dis-
tributed around some mean, similar to the underlying itself, often correlated with
the asset price.

Organization for Economic Cooperation and Development.
European Rate Mechanism.

CHAPTER 9 VEGA AND THE VOLATILITY SURFACE

It is assumed that one Brownian motion affects the volatility for one maturity and
that the rest of them follow in some known proportion.

Taleb (1996a).

The author’s strong reservations about the value-at-risk system do not apply to the
volatility curve for the following reasons: (1) This method is an aid in short-term
hedging not worst-case scenario analysis. In the “tail” events, volatility buckets mat-
ter less than the gamma exposure getting there. In addition, the cross-vega generally
vanishes in the tails; (2) this could be considered a second-degree value-at-risk method
since the principal instrument is the expectation of the second moment.

This corresponds to the floor traders’ designation of calendar spreads as “horizontal”
and strike spreads as “vertical,” with the mixture called “diagonal.”
[End OCR]*

## Page 495

*[Image OCR]
484

Notes

The local volatility is the volatility between two points with coordinates (S, t) and
(S + AS, t + At). A local volatility between (S, t) and (S, f + Af) is called a forward
volatility, as discussed earlier.

Dupire’s (1996) volatility spot forward is expressed as V(2dF /at)/(d7F/dK*) with K
the strike price, F the derivative security value, and t time to expiration. The second
derivative of the option price with respect to the strike presents what Breeden and
Litzenberger (1978) deemed to be the local risk-neutral density of the underlying
asset.

Some operators build the volatility surface as a function of the delta of the option in-
stead of the spot/strike. We avoided the pitfall of having to determine the delta as it is
well known that the delta is a function of volatility.

The following polynomial function was estimated:
x: time (days), y = In(fwd/K)
Volatility = 11.9257 + 0.110067 Sqrt[x] + 5.77624 y + 350.516 y — / Sqrt[x]
+ 153.843 y? — 1088.23 y? — 1805.65 y! + 21942.8 y5

CHAPTER 10 THETA AND MINOR GREEKS

For the pair numeraire-asset.

It is always preferable to derive the numerical, not the analytical, exposure for the
rho. The analytical exposure would sum up the microscopic derivatives of every op-
tion multiplied by the size. The numerical method reprices the portfolio and can thus
deal with changing derivatives. ,

This section can be skipped at first reading.

It can be computed using the classical Black-Scholes-Merton differential equation
(eliminating the drift):

BV 1 ac, eV
at 27 > 99?

To bridge the alpha with fair value, the trader can consider that the expected theta
given a volatility would be negative where one “pays. up” for the alpha. Expected
theta for a Black-Scholes-Merton hedger is the net of time decay after hedging in a
market of a given volatility. Using @ as time decay, I the gamma, « as the alpha, o as
implied volatility, and o’ as actual over the period, we get the expected theta:

E(®) = a(o) T(o) — a(o’)F(a’) ’

We see that expected time decay is positive if o' > o and negative if ¢' <a.
This section was added because it came to the author’s attention that traders tend to
use these measurements, not because of their risk management value.

Mathematical Note. The percentage difference is inaccurate for large moves because of
the asymmetry in lognormal markets. It is more mathematically correct to use Log
(spot/barrier). The difference becomes larger when the barrier is away from the
money.

Gamma is a special case of convexity (i.e., convexity with respect to the asset price).
However, its precision required presentation in a separate section.

Using the best fit method given the limited number of points, the author found:

Factor = 1.30825 — .03789 x + .0007419 x?, with x the maturity (in years).
[End OCR]*

## Page 496

*[Image OCR]
10.

Chapter 13 485

The mathematical way to price such a convexity is through Ito’s change of variable as
described in Module G. There is an Ito term in the payoff of the Eurodollar. The Eu-
rodollar futures could be interpreted as a function of a random variable with a payoff
equal to dr X exp (r(t — to)} with the random variable r following a Brownian motion.

CHAPTER 11 THE GREEKS AND THEIR BEHAVIOR

This topic is necessary for the study of exotic option structures.

CHAPTER 12. FUNGIBILITY, CONVERGENCE, AND STACKING

Again, this topic is presented from a practitioner’s perspective, not from a theoretical
view point.

The ensuing academic dispute provides a rather interesting, if not entertaining ex-
change for traders. On that, see Culp and Miller (1994), Canter and Edward (1995).

CHAPTER 13. SOME WRINKLES OF OPTION MARKETS

A quarterly option on a future is one whose expiration corresponds to that of a fu-
ture. Otherwise, the option is called a serial option on futures.

There is another type of barrier—the absorbing state—from which one cannot exit. A
trader who has a bad run knows he has hit the absorbing state when clearing firm
representatives show up in the middle of the pit to take away his trading privileges
(sometimes they require the help of security guards to draw the hapless trader out of
the pit),

Mathematical Note. If we bound a random walk to W, < W,,,,. It becomes clear that as
W, becomes close to W,,,, volatility needs to drop for the market to remain in a “fair”
game. Another solution is for the market to have excessive skew.

Flood and Garber (1992) demonstrate how stabilizing exchange rates lead to an in-
creased interest rates volatility.

Confirmed by two major brokerage firms, Tradition Financial Services and BCMG,
about their volume in New York, Singapore, London, and Tokyo. :

CHAPTER 15 BEWARE THE DISTRIBUTION

Heteroskedasticity for a normal distribution leads to a fourth moment higher than 3

times the square of the second moment.

See Hull and White (1987). To make the volatility process risk-neutral, we can safely

assume the existence of a volatility future contract.

I studied in Taleb (1997) the dynamics of a generalized switching regime process.
Let X, = Log(P,/P, _ ,,), the natural log-returns between time ¢ — A, and time # of

an observed variable P,.
[End OCR]*

## Page 497

*[Image OCR]
486

Notes

Let X,, ~ N(m,, o?), in other words the associated price P,, follows a geometric
Brownian motion with mean m1 and variance o?, since dP,/P, = m,dt+o,dW,Wa
unit variance-zero mean Wiener process.
and X,, ~ N(m,, 0,°)

The random variable X, is then said to follow a Brownian mixture of X, and X, if:

¢ X, = X,, with probability p
* X, = X,, with probability (1 — p)

This note simplifies the Markov chain by assuming that p (and 1 — p) are the er-
godic probabilities (i.e., the long term average spent in any of the two states).

The moment generating function defined as ®,(s) = fe p(x) dx becomes the
mixture of the m. g. f. of X, and X,. ®,(s) = p D, ,(s) + (1 — p) ®,,(s). See Feller (1971)
for the theorem.

Since X, and X, are both normally distributed:

®, ,(s) = exp(m,s + '2 0? s*)
and

®,,(s) = exp(m, s + 2 0} s?)
then

Denoting ®’ " the nth derivative of ©, and w, the n moment we obtain the fol-
lowing:
© ny = (0) = mp + m,(1 - p)o,

* my = ©'2(0) = pa? + (1 — p)o} + pm? + (1 ~ p) m?
© p, = 0'4(0) = 3po} + 3(1 — p)of + pm! t+ (1 — p)m3 + 6p m2o? + 6 m}(1 — po?

From which we see the “fatness of the tails”. When o, # a, and pe(0,1), »,/p3 > 3.
The constant elasticity of variance linking asset prices to their volatility seems to ex-
plain some of this behavior. However, the analysis is static. It is concerned with ab-
solute levels, not changes in levels. See Cox (1975); Beckers (1980); or MacBeth and
Merville (1980).

CHAPTER 16 OprtTiON TRADING CONCEPTS

See Derman (1995).

The matching cannot be perfect. It is necessary, therefore, to establish a cost function
(typically the quadratic norm of the difference) and minimize it. In addition, some
weights can be introduced to reflect the importance of the parameter risk. ,
To use market maker jargon, “below sheets.” “Sheets” are spreadsheets with theoret-
ical value that market makers carry on the exchanges. Below sheets means below the-
oretical value (assuming no shift in parameters).

This is analogous in probability theory to the Brownian bridge where possible paths
are examined conditional on a final arrival point.

Assume, for simplicity, that the forward trades at flat with the spot and that the
owner of an option incurs negligible carrying costs.

CHAPTER 17. BINARY OPTIONS: EUROPEAN STYLE

This resembles the Breeder-Litzenberger (1978) argument, where the infinitely
narrow butterfly (the call spread being purchased against another one being sold)
[End OCR]*

## Page 498

*[Image OCR]
Pe Ny

Chapter 22 487

is C(k + h) ~— C(k) — ( —C(k — h) + C(k)) = Clk +h) + Clk — A) - 2C(K). At the limit,
the entity becomes the second derivative of the option price with respect to the
strike.
The bet is the derivative with respect to the strike. Since
ac ac ac aC a0(K)
= + =— +
dC = aK aR + Boy FN) = ox AK + a5 aK

dK

CHAPTER 18 BINARY OPTIONS: AMERICAN STYLE

Howard Savery’s rule of thumb is that the American digital will be roughly twice the
European. The reason is that in at the money American will be worth $1 (it termi-
nates) while an European is worth $.50. This relationship wiil “contaminate” to lower
prices.

CHAPTER 19 BARRIER OPTIONS (I)

The reader should be acquainted with liquidity holes (Chapter 4) before proceeding.
Advanced Topic. (It is recommended that the reader practice the skew decomposition
method of the binary option prior to further reading of this section.)

See Grimmet and Stirzaker (1992); Karatzas and Shreve (1991).

The stickler can use an up node of exp {(a — Yeo)t} and a down node of exp {(— af)} sat-
isfying both risk neutrality and lognormality. Again there is no point complicating
such exercise with true lognormal moves. Alan Brace told the author: “communication
is more important than accuracy.”

See Dupire (1992, 1993, 1994, 1996), Rubinstein (1994), and Derman and Kani (1994).

CHAPTER 20 BARRIER OpTIONs (ID

A thorough understanding of American binary options is necessary before proceeding.

CHAPTER 22 MULTIASSET OPTIONS

The matrix needs to be symmetric and positive definite.

Lee Stulz (1982).

See Margrabe (1978).

Mathematical Note. Since the exercise aims at finding sensitivity, not prices, it is possi-
ble to set aside the lognormality issue. The correlation vega does not vary markedly
after an adjustment in the stochastic process. Only the odd moments (the ones linked
to the skew) will be truly affected.

See Taleb (1997).

The devaluation risk is excluded in pricing; therefore the operator would focus on the
option without the component.
[End OCR]*

## Page 499

*[Image OCR]
488 Notes

FP ey

CHAPTER 23. MINOR Exotics: LOOKBACK AND
ASIAN OPTIONS

See Goldman, Sosin, and Gatto (1979).

Given the fluidity of market designations, the term Asian may bear some other mean-
ing in some circles.

For an interesting approach to the problem, see Geman and Yor (1993) where the aver-
age of a geometric Brownian is decomposed in a series of subordinated Bessel processes.

MopDULE A BROWNIAN MOTION ON A \
SPREADSHEET, A TUTORIAL

This applies to any other spreadsheet program, provided the cells are respected.

Module D provides an explanation for the restrictions on the matrix, the equivalent of
volatility having to be positive.

Mopwutet B_ Risk NEUTRALITY EXPLAINED

See Cox, Ross, and Rubinstein (1979).
See Harrison and Kreps (1979), Harrison and Pliska (1981).

MoputLeE C NUMERAIRE RELATIVITY AND THE
Two-CountTry PARADOX

See Geman, E] Karoui, and Rochet (1995).

Such a method allows the pricing of puts as calls with the numeraire flipped.
This section can be omitted without any loss of the substance of the paradox.
For a definition of a convex function, see Chapter 1.

Changes of numeraire are used with term structure models as “glue” where every
maturity uses the preceding forward as a base for the stochastic process.

MopuLeE D- CORRELATION TRIANGLES:
A GRAPHICAL CASE STUDY

See Geman and Souveton (1996); Margrabe (1993).

MopDuLE E THe VALUE-AT-RISK

A trader translation: “when it goes it keeps going.”

See Markowitz (1959) for the breakthrough in minimum variance portfolio selec-
tion and the birth of modern portfolio theory. See Ingersoll (1986), Huang and
[End OCR]*

## Page 500

*[Image OCR]
ModuleG 489

Litzenberger (1988), or the more accessible Elton and Gruber (1995) for a descrip-
tion of the notion of mean variance.

Similar to the VAR is an excellent technique called “generalized pairs trading,” a short
term trading method where, conditional upon the return vector R, = {r,,+,,...,7,} at
time f of securities a vector X, = {X,, X,,..-, X,,} of security prices where a multivari-
ate diffusion is assumed, one can find a linear combination a = {a,, a,,..., a} maxi-
mizing the conditional probability of making a profit on a portfolio a” (X, — X,,,,)
where a’ is the transpose a, if the market diffused properly between (f — Af)
and (t + Af). In other words, the trader can thus detect which securities at times f
behaved as an outlier with respect to the covariance matrix between securities as
know at period t — At and take action, hoping that this would correct by next period
t + At. If two assets assumed to behave alike (that is highly correlated) parted way,
there might be a trade (assuming no new information) with the certainty of a profit
potential higher than 97%. The trader can buy one and sell the other if he attributes
the statistical divergence to some liquidity factor. Unlike the VAR, such technique,
handled by the professional, is impervious to a large amount of tail events as it fore-
casts the “body” of the distribution, not the rare occurrences.

This rule was inspired by a conversation with Raphael Douady.

There will be a combination of assets with negative volatility.

MopDvULE F PROBABILISTIC RANKINGS IN ARBITRAGE

See Rothchild and Stiglitz (1970) for the presentation of the difference between the
“risks” of two instruments.

MopuLe G) OPTION PRICING

More formally, W is a Wiener process. Uncertainty about the process is represented
by a filtered probability space “quadruple,” {Q, P, F, F,} where Q is the probability
space, P the objective probability measure, F a o-algebra (or, in this instance, a Borel
tribe, which consists in a collection of all subsets of 1 which includes the empty set,
all countable unions and all complements of subsets in (), F, the filtration, the set of
information about the process known at time ¢ (which includes information known at
all previous times), such that F,CF, whenever f = s.

See Billingsley (1985) for a pedagogical presentation of measure theory, Dothan
(1989) and the excellent and comprehensive Duffie (1996) for applications to fi-
nance. The reader will see a clear presentation of stochastic integration in Oksendal
(1995).

See Black and Scholes (1973); Merton (1973).

For further discussion of stochastic volatility techniques, see Scott (1987) and Bates
(1993).
The results here are inspired by Douady (1996).
[End OCR]*

## Page 501

*[Image OCR]
Bibliography

Essential Books for Traders

Cox, J., & Rubinstein, M.' (1985). Option Markets, Englewood Cliffs, NJ: Prentice Hall.

The best book ever written on options; after more than a decade the book has not a
speck of dust. It is written with rare intuitive depth of market sense. Imperative read-
ing for anyone who enters a trading room.

Natenberg, S. (1995). Option Volatility and Pricing Strategies (2nd ed.), Chicago: Probus.

Unencumbered with theory, it portrays option trading from the vantage point of a
down-to-earth floor trader. It provides an essential introductory text for book runners
and risk managers.

Baird, A. (1994). Option Market Making, New York: Wiley.

More advanced than Natenberg, a compact reading for a market maker. The reader
will find in it some introductory preparation for the concepts of risk management of
an inventory.

Hull, J. (1993). Option Futures and Other Derivative Securities (2nd ed.), Englewood Cliffs, NJ:
Prentice Hall.
Clear, pedagogical, and up to date.

Recommended Books

Abramowitz, M., & Stegun, N. C. (1970). Handbook of Mathematical Functions, New York:
Dover.

Bachelier, L. (1990). Theorie de la speculation, Annales de l’Ecole Normale Superieure, Paris:
Gauthier-Villars.

Beck, P., & Sydsaeter, K. (1991). Economist’s Mathematical Manual (2nd ed.), Heidelberg,
Germany: Springer Verlag.

Billingsley, P. (1986). Probability and Measure, New York: Wiley.

Bloomfield, P. (1976). Fourier Analysis of Time Series: An Introduction, New York: Wiley.

Brock, W., Hsieh, D., & LeBaron, B. (1991). Nonlinear Dynamics, Chaos and Instability, Boston:
MIT Press.

Burghardt, G., Belton, T., Lane, M., Luce, G., & McVey, R. (1991). Eurodollar Futures and Op-
tions, Chicago: Probus.

Dana, R. A., & Jeanblanc-Piqué, M. (1994). Marchés financiers en temps continu, Paris: Eco-
nomica.

DeRosa, D. (1992). Options on Foreign Exchange, Chicago: Probus.

DeRosa, D. (1996). Managing Foreign Exchange Risk (2nd ed.), Homewood, IL: Irwin.

Dixit, A., & Pyndick, R. (1994). Investment Under Uncertainty, Princeton, NJ: Princeton Uni-
versity Press.

Dothan, M. (1990). Prices in Financial Markets, New York: Oxford University Press.

"The author owes his trading education to a smuggled copy of the manuscript for this book.

490
[End OCR]*

## Page 502

*[Image OCR]
Bibliography 491

Dubins, L., & Savage, L. (1965). How to Gamble If You Must, New York: McGraw-Hill.

Duffie, D. (1988). Security Markets Stochastic Models, New York: Academic Press.

Duffie, D. (1996). Dynamic Asset Pricing Theory, Princeton, NJ: Princeton University Press.

Durett, R. (1991). Probability: Theories and Examples, Pacific Grove, CA: Wadsworth and
Brooks/Coles.

Elton, E., & Gruber, M. (1995). Modern Portfolio Theory and Investment Analysis (5th ed.),
New York: Wiley.

Enders, W. (1995). Applied Econometric Time Series, New York: Wiley.

Feller, W. (1968). An Introduction to Probability Theory and Its Application (Vol. I; 3rd ed.),
New York: Wiley.

Feller, W. (1971). An Introduction to Probability Theory and Its Application (Vol. I]; 2nd ed.),
New York: Wiley.

Gastineau, G. L. (1992). Dictionary of Financial Risk Management, Chicago: Probus.

Grabbe, J. O. (1996). International Financial Markets (3rd ed.), Englewood Cliffs, NJ: Prentice
Hall.

Grimmett, G. R., & Stirzaker, D. R. (1992). Probability and Random Processes (2nd ed.), Ox-
ford: Oxford University Press.

Hamilton, J. D. (1994). Times Series Analysis, Princeton, NJ: Princeton University Press.

Harvey, A. C. (1993). Time Series Models (2nd ed.), London: Harvester Wheatsheaf.

Huang, C. F., & Litzenberger, R. (1988). Foundations for Financial Economics, Englewood
Cliffs, NJ: Prentice Hall.

Ingersoll, J. (1986). The Theory of Financial Decision Making, Savage, MD: Rovoman & Little-
field.

Jarrow, R. (1995). Modelling Fixed Income Securities» and Interest Rate Options, New York:
McGraw-Hill.

Jarrow, R., & Rudd, A. (1983). Option Pricing, Homewood, IL: Irwin.

Karatzas, I., & Shreve, S. (1991). Brownian Motion and Stochastic Calculus (2nd ed.), New
York: Springer Verlag.

Kennedy, P. (1992). A Guide to Econometrics (3rd ed.), Cambridge, MA: MIT Press.

Lamberton, D., & Lapeyre, B. (1991). Introduction au calcul stochastique applique a la finance,
Paris: Ellipses. °

Levy, P. (1948). Processus stochastiques et mouvement brownien, Paris: Gauthier-Villars.

Markowitz, H. (1959). Portfolio Selection: Efficient Diversification of Investments, New York:
Wiley.

Merton, R. C. (1990), Continuous-Time Finance, Cambridge, MA: Blackwell.

Nelken, I. (Ed.). (1995). The Handbook of Exotic Options, Homewood, IL: Irwin.

Oksendal, B. (1995). Stochastic Differential Equations (4th ed.), Berlin: Springer.

Oxford English Dictionary. (1971). Oxford: Oxford University Press.

Pesaran, N. H., & Potter, S. M. (Eds.). (1993). Nonlinear Dynamics, Chaos, and Econometrics,
London: Wiley.

Ray, C. (1993). The Bond Market, Trading and Risk Management, Homewood, IL: Irwin.

Russell, B. (1945). History of Western Philosophy, New York: Simon & Schuster.

Schwager, J. D. (1992). The New Market Wizards, Conversations with America’s Top Traders,
New York: Harper Business.

Shiller, R. J. (1989). Market Volatility, Cambridge, MA: MIT Press.

Shimko, D. (1992). Continuous Time Finance: A Primer, Miami, FL: Kolb Publishing.

5
[End OCR]*

## Page 503

*[Image OCR]
492 Bibliography

Soros, G. (1994). The Alchemy of Finance (2nd ed.), New York: Wiley.

Stigum, M. (1990). The Money Market (3rd ed.), Homewood, IL: Irwin.

Taylor, H. M., & Karlin, 5. (1981). A Second Course in Stochastic Processes, New York: Aca-
demic Press.

Taylor, H. M., & Karlin, S. (1994). An Introduction to Stochastic Modeling, New York: Aca-
demic Press.

Varian, H. (1993). Economic and Financial Modeling with Mathematica, Santa Clara, CA: Telos.

Wilmott, P., Dewynne, J., & Howison, S. (1993). Option Pricing, Mathematical Models and
Computation, Oxford: Oxford Financial Press.

Wolfram, S. (1991). Mathematica, a System for Doing Mathematics by Computer (2nd ed.), Read-
ing, MA: Addison-Wesley.

Recommended Articles

Amin, K., & Jarrow, R. (1991). “Pricing Foreign Currency Options under Stochastic Inter-
est Rates,” Journal of International Money and Finance, 10, 310-329.

Atilgan, T. (1996). A General Review of Volatility Estimation Techniques, unpublished manu-
script.

Avellaneda, M., Levy, A., & Paras, A. (1995). “Pricing and Hedging Derivative Securities in
Markets with Uncertain Volatilities,” Applied Mathematical Finance, 2, 73-88.

Avellaneda, M., & Paras, A. (1994). “Dynamic Hedging Portfolios for Derivative Securi-
ties in the Presence of Large Transaction Costs,” Applied Mathematical Finance, 1,
165-193.

Bailey, W. (1987, December). “An Empirical Investigation of the Market for Comex Gold Fu-
tures Options,” Journal of Finance, 42(5), 1187-1194.

Bailie, R. T., & Bollerslev, T. (1989). “The Message in Daily Exchange Rates: A Conditional
Variance Tale,” Journal of Finance, 44(1), 167-182.

Bates, D. S. (1993). Jumps and Stochastic Volatility: Exchange Rate Processes Implicit in PHLX
Deutschemark Options, National Bureau of Economic Research, Working Paper No. 4596.

Bauman, W. S., & Miller, R. (1994). “Can Managed Portfolio Performance be Predicted?”
Journal of Portfolio Management, 20(4), 31-40. .

Beckers, S., (1980). “The Constant Elasticity of Variance Model and Its Implications for Op-
tion Pricing,” Journal of Finance, 35(3), 661-673.

Black, F. (1975). “Fact and Fantasy in the Use of Options,” Financial Analysts Journal, 31(4),
36-41, 61-72. s

Black, F. (1976). “Pricing of Commodity Contracts,” Journal of Financial Economics, 3.

Black, F. (1988, August). “The Holes in Black-Scholes,” in New Frontiers in Options, RISK.

Black, F., & Scholes, M. (1973). “The Pricing of Options and Corporate Liabilities,” Journal
of Political Economy, 81, 637-654.

Bollerslev, T. (1986). “Generalized Autoregressive Conditional Heteroskedasticity,” Journal
of Econometrics, 31, 307-327.

Bollerslev, T. (1987, August). “A Conditional Heteroskedastic Times Series Model for Spec-
ulative Prices and Rates of Return,” Review of Econometric Studies.

Bowie, J., & Carr, P. (1994, August). “Static Simplicity,” RISK.

Boyle, P. P. (1977). “Options: A Monte Carlo Approach,” Journal of Financial Economics, 4(3),
323-338.
[End OCR]*

## Page 504

*[Image OCR]
Bibliography 493

Boyle, P. P. (1989, March). “The Quality Option and Timing Option in Futures Contracts,”
Journal of Finance, 44(1), 101-114.

Boyle, P. P., Evnine, J., & Gibbs, S. (1989). “Multivariate Contingent Claims,” Review of Fi-
nancial Studies, 2(2), 241-250.

Breeden, D., & Litzenberger, R. (1978). “Prices of State-Contingent Claims Implicit in Op-
tion Prices, Journal of Business, 51, 621-651.

Brennan, M. J., & Schwartz, E. S. (1978, September). “Finite Difference Methods and Jump
Processes Arising in the Pricing of Contingent Claims: A Synthesis,” Journal of Finan-
cial and Quantitative Analysis.

Brock, W., & de Lima, P. J. F. (1995). Nonlinear Time Series, Complexity Theory, and Finance,
Working Paper.

Carr, P. (1988, December). “The Valuation of Sequential Exchange Opportunities,” Journal
of Finance.

Carr, P. (1994). “European Put Call Symmetry,” Cornell University, Working Paper.

Carr, P. (1995). Static Hedging of Path Dependent Options, Working Paper.

Carr, P., (1995). Two Extensions to Barrier Option Valuation, Working Paper.

Cox, J. (1975). Constant Elasticity of Variance Diffusions, Working Paper, Stanford University.

Cox, J. C., Ingersoll, J., & Ross, S. (1985a). “An Intertemporal General Equilibrium Model of
Asset Prices,” Econometrica, 53, 363-384.

Cox, J. C., Ingersoll, J., & Ross, S. (1985b). “A Theory of the Term Structure of Interest
Rates,” Econometrica, 53, 385-407.

Cox, J. C., & Ross, S. (1976). “The Valuation of Options for Alternative Stochastic
Processes,” Journal of Financial Economics, 3.

Cox, J. C., Ross, S. A., & Rubinstein, M. (1979). “Option Pricing: A Simplified Approach,”
Journal of Financial Economics, 7.

Culp, C. L., & Miller, M. H. (1994). “Hedging a Flow of Commodity Deliveries with Fu-
tures: Lessons from Metallgessenshaft, Derivatives Quarterly, 1.

Dacorogna, M.-M., Muller, U., Embrechts, P., & Samorodnitsky, G. (1995). Moment Condi-
tions for HARCH(k) Models, Working Paper, Olsen & Associates, Zurich, Switzerland.

Derman, E., Ergener, D., & Kani, I. (1995, Summer). “Static Option Replication,” Journal of
Derivatives. :

Derman, E., & Kani, I. (1994, February). “Riding on a Smile,” RISK, 32-39.

Derman, E. (1996). “Valuing Models and Modelling Value,” Journal of Portfolio Management,
22(3), 106~114.

Douady, R. (1995). Mouvement browniens cylindriques et optimisation de Heath-Jarrow-Morton,
Paris: Ecole Normale Superieure Working Paper.

Douady, R. (1995). Options a@ limites, Working Paper.

Douady, R. (1996). De la distribution du temps de sortie et de son esperance, Working Paper.

Dubourg, N., & Douady, R. (1995). “Energy Optimization and Optimal Hedging under
Proportional Transaction Costs,” Societe Generale, Paris.

Dupire, B. (1992). “Arbitrage Pricing with Stochastic Volatility,” in Proceedings of the A.F.F.I.
Conference of June 1992.

Dupire, B. (1993, September). “Model Art,” RISK.

Dupire, B. (1994, January). “Pricing with a Smile,” RISK, 7, 18-20.

Dupire, B., (1996). A Unified Theory of Volatility, Working Paper, Paribas Capital Markets.
[End OCR]*

## Page 505

*[Image OCR]
494 Bibliography

Edwards, F., & Canter, M. (1995). “The Collapse of Metalgessellschaft: Unhedgeable
Risks, Poor Hedging Strategy, or Just Bad Luck?” Journal of Futures Markets, 15(3),
211-264.

El Karoui, N., & Geman, H. (1994). “A Probabilistic Approach to the Valuation of General
Floating Rate Notes with an Application to Interest Rate Swaps,” Advances in Options
and Futures Research.

Engle, R. (1982). “Autoregressive Conditional Heteroskedasticity with Estimates of the
Variance of UK Inflation,” Econometrica, 50.

Engle, R., & Rosenberg, J. (1995, Summer). “Garch Gamma,” Journal of Derivatives.

Flood, R. P., & Garber, P. M. (1992). “The Linkage between Speculative Attack and Target
Zone Models of Exchange Rates: Some Extended Results,” in P. Krugman & M. Miller
(Eds.), Exchange Rate Targets and Currency Bands, Cambridge, England: Cambridge Uni-
versity Press.

Fung, W., & Hsieh, D. (1996). Global Yield Curve Event Risk, Working Paper, Fuqua School of
Business, Duke University.

Garman, M. (1976). “Market Microstructure,” Journal of Finance, 3, 257-275.

Garman, M. (1989), “Immunizing Foreign Exchange Contracts against Swap Rate and
Volatility Risks,” Journal of International Financial Management and Accounting, 1.

Garman, M. (1992, December). “Spread the Load,” RISK.

Garman, M., & Klass, M. J. (1980). “On Estimation of Security Price Volatilities from His-
torical Data,” Journal of Business, 53(1), 67-78.

Garman, M., & Kohlhagen, S. W. (1983). “Foreign Currency Option Values,” Journal of Inter-
national Money and Finance, 2(3), 231-237.

Geman, H., El Karoui, N., & Rochet, J. C. (1995). “Changes of Numeraire, Changes of Prob-
ability Measure, and Option Pricing,” Journal of Applied Probability, 32, 443-458.

Geman, H., & Souveton, R. (1996). “No Arbitrage between Economies and Correlation Risk
Management,” forthcoming in Computational Economics.

Geman, H., & Yor, M. (1993). “Bessel Processes, Asian Options, and Perpetuities,” Mathe-
matical Finance, 3(4), 349-375.

Geman, H., & Yor, M. (1996). “Pricing and Hedging Double Barrier Options: A Probabilis-
tic Approach,” forthcoming, Mathematical Finance.

Geske, R. (1979). “The Valuation of Compound Options,” Journal of Financial Economics,
7(1), 63-82.

Geske, R., & Johnson, H. E. (1984, December). “The American Put Option Valued Analyti-
cally,” Journal of Finance, 39(5), 1511-1524.

Geske, R., & Roll, R. (1984). “On Valuing American Call Options with the Black-ScHoles Eu-
ropean Formula,” Journal of Finance, 39(2), 443-455.

Geske, R., & Shastri, K. (1985). “Valuation by Approximation: A Comparison of Alternative
Option Valuation Techniques,” Journal of Financial and Quantitative Analysis, 20(1),
45-72.

Goldman, B. M., Sosin, H. B., & Gatto, M. A. (1979). “Path Dependent Options: ‘Buy at the
Low, Sell at the High,’” Journal of Finance, 34(5), 1111-1127.

Grossman, S. (1987). “Insurance Seen and Unseen: The Impact on Markets,” Journal of Port-
folio Management, 14(4), 5-6.

Grossman, S. (1988). “An Analysis of the Implication for Stock and Future Price Volatility
of Program Trading and Dynamic Hedging Strategies,” Journal of Business, 61(3).
Grossman, S. J., & Miller, M. H. (1988). “Liquidity and Market Structure,” Journal of Fi-

nance, 43, 617~633.
[End OCR]*

## Page 506

*[Image OCR]
Bibliography 495

Guillaume, D., Dacorogna, M., Dave, R., Muller, U., Olsen, R., & Pictet, O. (1995, March).
From the Bird’s Eye to the Microscope, Working Paper, Olsen & Associates, Zurich,
Switzerland.

Hamilton, J. (1993). “Autoregressive Conditional Heteroskedasticity and Changes in
Regime,” U. C. San Diego, Working Paper.

Hamilton, J. (1994). “Rational Expectations and the Economic Consequence of Changes in
Regime,” U. C. San Diego, Working Paper.

Hamilton, J., & Susmel, R. (1995). “Specification Testing in Markov Switching Time Series
Models,” U. C. San Diego, Working Paper.

Harrison, M., & Kreps, D. (1979). “Martingales and Arbitrage in Multiperiod Securities
Markets,” Journal of Economic Theory, 20, 381-408.

Harrison, M., & Pliska, S. (1981). “Martingales and Stochastic Integrals in the Theory of
Continuous Trading,” Stochastic Processes and Their Applications.

Heath, D., Jarrow, R., & Morton, A. (1987). Bond Pricing and the Term Structure of Interest
Rates, Working Paper, Cornell University.

Heath, D., Jarrow, R., & Morton, A. (1990). “Bond Option Pricing and the Term Structure of
Interest Rates: A Discrete Time Approximation,” Journal of Financial and Quantitative
Analysis, 25, 419-440.

Heath, D., Jarrow, R., & Morton, A. (1992). “Bond Pricing and the Term Structure of Inter-
est Rates: A New Methodology for Contingent Claims Valuation,” Econometrica, 60,
77-106.

Heynen, R., & Kat, H. (1994, June). “Crossing Barriers,” RISK, 46-51.

Ho, T., & Stoll, H. (1980). “On Dealer Markets under Competition,” Journal of Finance, 35,
259-267.

Ho, T., & Stoll, H. (1983). “The Dynamics of Dealer Market under Competition,” Journal of
Finance, 38, 1053-1074.

Hull, J., & White, A. (1987). “The Pricing of Options on Assets with Stochastic Volatili-
ties,” Journal of Finance, 42(2), 281-300.

Hutchinson, J., Lo, A. W., & Poggio, T. (1994, July). “A Non-Parametric Approach to Pricing
and Hedging Derivative Securities via Learning Networks,” Journal of Finance, 49(1),
27-36.

Irvin, S. H., Zulauf, C., & Ward, B. (1994, Winter). “The predictability of Managed Futures
Returns,” Journal of Derivatives.

Jackwerth, J. C., & Rubinstein, M. (1995). “Implied Probability Distributions: Empirical
Analysis,” Working Paper, U. C. Berkeley.

Jamshidian, F. (1989, March). “An Exact Bond Option Formula,” Journal of Finance, 4411),
205-210.

Johnson, H. (1987). “Options on the Maximum or the Minimum of Several Assets,” Journal
of Financial and Quantitative Analysis, 22(3), 277-284.

Johnson, H., Shanno, D. (1987, June). “Option Pricing When the Variance Is Changing,”
Journal of Financial and Quantitative Analysis.

Johnson, H., & Stulz, R. (1987). “The Pricing of Options with Default Risk,” Journal of Fi-
nance, 42(2), 267-280.

Kemna, A. G. Z., & Vorst, A. C. F. (1990). “A Pricing Method for Options Based on Average
Values,” Journal of Banking and Finance, 14, 113-130.

Kim, D., & Kon, S. (1994). “Alternative Models for the Conditional Heteroskedasticity of
Stock Returns,” Journal of Business, 67(4), 563-598.

Kritzman, M. (1994). “About Serial Dependence,” Financial Analysts Journal, 50(2), 19-22.
[End OCR]*

## Page 507

*[Image OCR]
496 Bibliography

Kunitomo, N., & Ikeda, M. (1992). “Pricing Options with Curved Boundaries,” Mathemati-
cal Finance, 2, 275-282.

Kuserk, G., & Locke, P. (1993). “Scalper Behavior in an Auction Market: An Analysis of
Scalpers in Futures Markets,” Journal of Futures Markets, 13, 409-431.

Kuserk, G., & Locke, P. (1994, Summer). “Market Maker Competition on Futures Ex-
changes,” Journal of Derivatives, 56-66.

Leland, H. (1985). “Option Pricing and Replication with Transactions Costs,” Journal of Fi-
nance, 40 (5), 1285-1301.

Leland, H. (1992, December). “The Lessons of History,” RISK.

Levy, E. (1992). “Pricing European Average Rate Currency Options,” Journal of International
Money and Finance, 11(4), 474-491.

Lo, A., & MacKinley, A. C. (1988). “Stock Market Prices Do Not Follow Random Walks:
Evidence from a Simple Specification Test,” Review of Financial Studies, 1, 41-66.

MacBeth, J., & Merville, L. (1980, May). “Tests of the Black-Scholes and Cox Call Option
Valuation Models,” Journal of Finance, 35(2), 285-303.

Margrabe, W. (1978). “The Value of an Option to Exchange One Asset for Another,” Journal
of Finance, 33(1), 177-186.

Margrabe, W. (1993, Fall). “Triangular Equilibrium and Arbitrage in the Market for Op-
tions to Exchange Two Assets,” The Journal of Derivatives, 1.

Merton, R. (1973, Spring). “Theory of Rational Option Pricing,” Bell Journal of Economics.

Merton, R. (1976). “Option Pricing When Underlying Stock Returns Are Discontinuous,”
Journal of Financial Economics, 3(1/2), 125-144.

Merton, R. (1982). “On the Mathematics and Economic Assumptions of Continuous-Time
Models,” in W. F. Sharpe & C. M. Cootner (Eds.), Financial Economics: Essays in Honor of
Paul Cootner, Englewood Cliffs, NJ: Prentice Hall.

Nelson, D. B. (1990). “ARCH Models as Diffusion Approximations,” Journal of Econometrics,
45, 7-38.

Parkinson, M. (1980). “The Extreme Value Method for Estimating the Variance of the Rate
of Return,” Journal of Business, 1.

Pechtl, A. (1995, June). “Classified Information,” RISK.

Peterson, D. R., & Tucker, A. L. (1988). “Implied Spot.Rates as Predictors of Currency Re-
turns: A Note,” Journal of Finance, 43(1), 247-258.

Rabinovitch, R. (1989). “Pricing Stock and Bond Options When the Default-Free Rate Is
Stochastic,” Journal of Financial and Quantitative Analysis, 24(4), 447-458.

Rendleman, R. J., Jr., & Bartter, B. J. (1979, December). “Two-State Option Pricing,< Journal
of Finance.

Roll, R. (1984). “A Simple Implicit Measure of the Effective Bid-Ask Spread in An Efficient
Market,” Journal of Finance, 39(4), 1127~1139.

Ross, S. A. (1989). “Information and Volatility: The No-Arbitrage Martingale Approach to
Timing and Resolution Irrelevancy,” Journal of Finance, 44(1), 1-18.

Rothchild, M., & Stiglitz, J. (1970). “Increasing Risk I: A Definition,” Journal of Economic
Theory, 2, 225-243.

Rubinstein, M. (1983, March). “Displaced Diffusion Option Pricing,” Journal of Finance.

Rubinstein, M. (1984). “A Simple Formula for the Expected Rate of Return of an Option
over a Finite Holding Period,” Journal of Finance, 39(5), 1503-1509.

Rubinstein, M. (1985). “Nonparametric Tests of Alternative Option Pricing Models Using
CBOE Reported Trades,” Journal of Finance, 40(2), 455-480.
[End OCR]*

## Page 508

*[Image OCR]
Bibliography 497

Rubinstein, M. (1991). Exotic Options, Finance Working Paper No. 220, U. C. Berkeley.

Rubinstein, M. (1994, June). “Implied Binomial Trees,” Journal of Finance, 49(3), 771-818.

Scott, L. O. (1987, December). “Option Pricing When the Variance Changes Randomly:
Theory, Estimation and Application,” Journal of Financial and Quantitative Analysis.

Shimko, D. (1993, April). “Bounds of Probability,” RISK.

Silber, W. (1984). “Market making Behavior in an Auction Market: An Analysis of Scalpers
in a Futures Market,” Journal of Finance, 39, 937-953.

Stroll, H. R. (1989). “Inferring the Components of the Bid-Ask Spread: Theory and Empiri-
cal Tests,” Journal of Finance, 44, 115-134.

Stulz, R. M. (1982). “Options on the Minimum or Maximum of Two Risky Assets: Analysis
and Applications,” Journal of Financial Economics, 10, 161-185.

Taleb, N. (1996). “Dynamic Hedging in Heteroskedastic Environments,” presented at the
Institute for Advanced Studies, Princeton, NJ.

Taleb, N. (1996). “Parkinson Statistic, Mean Reversion and Numerical Pricing of Barrier
Options Using a Brownian Bridge,” Working Paper.

Taleb, N. (1997). “Essays in Applied Option Theory,” Working Paper.

Turnbull, 5. M., & Wakeman, L. M. (1991, September). “A Quick Algorithm for Pricing Euro-
pean Average Options,” Journal of Financial and Quantitative Analysis.

Vasicek, O. (1977). “An Equilibrium Theory of the Term Structure,” Journal of Financial Eco-
nomics, 5, 177-188.

Whaley, R. E. (1986). “Valuation of American Futures Options: Theory and Empirical
Tests,” Journal of Finance, 41(1), 127-150.

Whalley, A. E., & Wilmott, P. (1993). “Counting the Costs,” RISK, 6, 10.
[End OCR]*

## Page 509

*[Image OCR]
Index

Note: Bold numbers indicate the page on
which the definition of the term is found.
Bold italic page numbers are Option Wizards.

Accounting, nonparallel, 249-250
Accounting systems, and arbitrage, 85-86
Advanced shadow gamma, 142
Adverse selection, 60-61
Alpha, 113, 178-181, 178. See also Greeks
American options, 18, 24-29, 37, 231, 295-311
binary, 295-311
and bucketing, 231
forward delta, 22
hard, 21, 25-26, 25
soft, 21, 24-25, 24
that nobody ever exercises, 29
Amortizing interest rate securities, 85, 309
Arbitrage, 80-87, 80, 453-458
and accounting systems, 85-86
credit arbitrage, 86
deterministic relationships, 82
distributional, 291
duration of, 84~85, 84
and fungibility, 212
legal arbitrage, 87
mechanical vs. behavioral stability, 81-82
nonmarket forms of, 86-87
orders/levels of (Table 5.1), 81
passive, 83-84, 83
probabilistic rankings in, 453-458
squeeze, 84
tax arbitrage, 86
and variance of returns, 87
Arbs (market participant), 48
ARCH (autoregressive conditional
heteroskedasticity), 102-103, 142, 241
E-GARCH (exponential), 102
GARCH (generalized), 102-103
GARCH gamma, 142
N-GARCH (nonlinear), 102
SWARCH (switching regime), 102
ArcSine law, 66, 407, 408
Aristotle, 13
Arrow-Debreu price, 15
Asian options, 35, 47, 271, 272, 408-411, 408
Asymmetry option hedging costs, 78, 203
Asymptotic delta, 206, 207
Asymptotic vega test, 201
Augmented volatility, 77
Autocorrelation of price changes, 58, 59, 79

Bachelier, Louis, 91, 111
Band, currency, 225
Barrier(s), 36, 43
high barrier, 43
low barrier, 43
Barrier(s), market, see Market(s), barriers
Barrier options, 79, 308, 309, 312-346, 347-375,
467-477
categories of (Table 19.1), 313
and delta, 117, 126-127, 128
double, 362-368, 362
drift, adding (complexity of the forward line),
321-323

Dupire-Derman-Kani technique, 344-345
exercise (adding the puts), 346
first exit time, 339
Girsanov, 339, 469-470
knock-in options, 317-319
knock-out option, 312-316, 312
and liquidity vacuum, 72-73
and path dependence, 47
pricing, 343, 470-477
calls down-and-in, 473
calls down-and-out, 473
calls up-and-in, 472
calls up-and-out, 472
and Parkinson number, 104
puts down-and-in, 472
puts down-and-out, 472
puts up-and-in, 473-474
puts up-and-out, 473
put-call parity rule, 252-253
put/call symmetry and hedging of, 323-331
reflection principle, 336-339, 469
regular, 312-346
reverse, 347-361
knock-out, 347
risk reversals, 323
risks, primary/secondary (Table 1.4), 35
rules, ranking of securities, 454
single volatility fudge, 343-344
skew environments, barrier decomposition
wnder, 331-335
stopping time, 339
time, effect of, on knock-out options, 339
variance ratios, 345-346
volatility, effects of, 319-321
Basket(s)/basket options, 13, 82, 383, 391
Asian options, 408-411
correlation vega of (Table 22.2), 394
in diagram, classification of derivatives (Figure
1.1), 10
rule, 54
and stacking, 220 -
Behavioral stability, 81
Bermudan option, 18
Bet options, 273. See also Binary options
(American); Binary options (European)
Bets, higher moment, 264-265
Betspreads, 292-294
Biased assets, 73, 142, 143, 201, 248, 400
and convergence, 216-217
simplified two-regime world for (Table 15.2),
249
and skew, 245-252
Binary option, 274
American, 295-311, 295
double, 307-311
put-call parity rule, 253
single, 295~307
trading and hedging, 272
European, 273-294, 274
bet contract, 278
forward and spot bets, 278-279
spot bet, 278
conclusion (statistical trading vs. dynamic
hedging), 289-290

499
[End OCR]*

## Page 510

*[Image OCR]
500 Index

Binary option (Continued)
contingent premium options, 290-292, 291
hedging with a vanilla, 275
mixed convexity of (Figure 17.1), 275
put-call parity rule, 253
skew, pricing with the, 279-289
difference between binary and delta (delta
paradox), 284-286
Dirac delta, 284-287
first hedging consequences, 286
formal pricing on the skew, 281
gamma for a bet, 287-289
skew paradox, 282-283
trading and hedging (things to understand), 272
Binomial tree, 175, 176-177, 256, 336, 345
Black Scholes derivation, 461~464
Black-Scholes-Merton, 1, 2, 77, 88, 107, 111, 139,
146, 177, 197, 198, 239, 259, 279-280, 428
as an almost nonparametric pricing system,
109-114
assumptions, and the real world, 109-114
Greeks:
corrections to for use of, 112-113
delta, 35, 112, 115, 115-131. See also Delta
gamma, 11, 112, 132, 132-146. See also Gamma
minor Greeks, 171-181. See also Alpha; Omega;
Rho
theta, 112, 167~171, 167. See also Theta
vega, 112, 147-166, 147. See alse Vega
Bleed, 36, 191-197, 191, 228
Bollerslev, T., 102
Bonds with negative / positive carry, and hard early
exercise, 26
Book, 148, 149
trading, vs. trading an option, 148-149
Book runners vs. price takers, 48-50
Boost, 189, 190
Borel-Cantelli Lemma, 64
Brownian motion, 73, 191, 345, 415-425, 464
arithmetic, 92
extrema, 407
geometric, 89, 111, 128, 130-131, 284, 336
vs, arithmetic, 89-92
memoryless, 73
on a spreadsheet, 415-425
Bucket(s)/bucketing, 36, 127, 174, 229-232, 229
American and path-dependent options, 231
and Asian options, 411
and delta, 127
European options, 229-231
forward or forward-forward, 231-232
static straight, 229-231
straight, 229
Bucket vega, 157
Burnout, trader, 60
Butterfly/ies, 292
convergence, 214
profile (Figure 1.4), 23
put/call, 23
stack, 218, 219, 221

Calendar, smile, 27-29

Calendar rules, and ranking of securities, 454

Calendar spreading, 272

Call option, 18, 472-473

Caplets, 34, 382

Capped index option, 365

Caps, 34, 40

Caption, 382

Carry, 295. See also Convergence
correlation between interest rates and, 252
“hogs,” 216

Case studies:
at-settlement, 306-307
betspreads, 292-294
indexed notes, 395-402
knock-out box, 348-355
multiasset bets, 294
nasty path (Table 16.4), 269
National Vega Bank, 298-299
path dependence of a regular option, 265-270
Syldavian elections, 145-146
worst-case scenario for a delta hedger, 270
Chaos theorists, 246
Charlatanism, 445
Chicago Board Options Exchange (CBOE), 38
Choice (multiasset options), 383, 384-390
rainbow options, 384
Chooser options, 380-382, 380, 467, 468
Circuit breakers, market, 74
Cleared market, 374
Cocteau, Jean, 1
Commodities market, and primary/secondary risks
(Table 1.3), 34
Commoditized products, 50-53
Commodity Futures Trading Commission, 57
Communication problems, traders/quants, 5
Compact distribution, 202
Compound options, 203, 376-379, 376, 467, 468
and primary/secondary risks (Table 1.4), 35
put-call parity rule, 253
uses of; hedging barrier vega, 379
vega linearity; costs of dynamic hedging, 378-379
Condors, 292
Contamination principle, 9-10, 14-15, 20, 295, 457
Contingent claims, 16-20
Contingent premium options, 290-292, 291
Continuous/discontinuous payoffs, 41-42
Convergence, 83, 213-217
~ and biased assets, 216-217
butterfly, 214
convergence trading, 213
and convexity, 216
first-order, 214
leveis of convergence trading, 216
mapping, 215-216
second-order, 214
and volatility, 216
Convexity, 11, 14-15, 113, 183-190, 183. See also
Contamination principle
of a bond to interest rates (simplified), 184
changes (third order curve shift), 162
and convergence, 216
correlation rules, 457
general rules, 458
interest rate (modified), 188 >
middlebrow, 185
modified, 183, 185, 187, 188
for a nonlinear derivative, 185
of an option to the underlying asset, 184
vega, 184, 303-305
Correlated siblings, 45
Correlation, 88-101, 88, 239
actual, 88
calculating historical, 92-94
constant (“no such thing”), 97~99
convexity rules, 457
delta, 113
family of exotic options, 363
filtering, 95~97
implied, 88
interest rates and carry, 252
Parkinson number, 101-107
rules, and ranking of securities, 455
[End OCR]*

## Page 511

*[Image OCR]
triangles, 88, 438-444
variance ratio method, 101-107
vega, 386
Countercurrency, 431
Covariance bucket vega, 153
Covariance matrix, 271
Covered writes, 351
tock delta for, 205
Cox, J., 114, 177, 188, 372
Crash of 1987, 27, 74, 76, 97
Credit:
arbitrage, 86
and futures/forwards, 30
rating, arbitraged, 85
risk, 311
rule, 264
Cross-currency, 82
Cubic spline, 172
Currency(ies). See also Correlation, triangles;
Numeraire
as assets, 250-251
band, 225
book, and classification of risks, 33
cash position, mapping convergence for, 215
exposure, 82, 83
forward, mapping convergence for, 215
market, and primary/secondary risks (Table 1.3), 34
Curvature, 183. See also Convexity
Customer flows, 54

D, 191
Dana, R., 462
Dangerous trades; unhedged vs. mishedged, 361
Ddelta-Dspot, 191
Ddeltadvol, see Stability ratio (Ddeltadvol)
Decay, 179
exponential, 95
trader, 52
Decomposition, and replication, need to
distinguish between, 256
Delta, 36, 112, 115, 115-131
adjustments, general option, 104
and barrier options, 126-127
bleed, 191-197, 191, 228
and bucketing, 127
characteristics of, 116
confusion (by cash or by forward), 123
continuous time, 116-121
of the delta (second moment), 202
difference between binary and, 284
first moment, 202
forward, for European option, 22
of the gamma (third moment; skew), 202
gap, 300
and hedge ratio, 116-121
hedging reverse knock-outs, 356-361
for linear instruments, 123-126
lock, 204-205, 204
modified, 36, 118, 120
paradox, 284
as risk measure, 121-123
and shadow gamma, 140
soft vs. hard, 262-263
hard delta, 262
soft delta, 262
in the value at risk, 127
and volatility/extreme volatility, 127-131
DeMoire-Laplace, 417
Derivatives, 9-12, 9
classification of (Figure 1.1), 10
linear/nonlinear and convex /concave (Figures
1.2A-D), 12

Index 501

noncontingent time-dependent nonlinear
derivatives, 16
nonlinear derivative, 9
test of local linearity, 11
time-dependent linear derivatives, 13-16, 13
and VAR (not applicable), 451
Derman, E., 344, 345, 373
Derman’s method, 256
DeRosa, D., 123
Deterministic relationships, 82
DgammaDspot/Dgamma-Dspot, 118, 191
Digital derivatives, 273. See also Binary options
(European)
Digital (discontinuous) and ramp (continuous)
payoffs, 41
Dimension of the structure, 43-45, 43
Dirac Delta, 285, 286-287, 373
Distribution:
compact, 202
“peware the,” 238-255
skew, see Skew
tails, see Tails
Dollar variance test (Figure 17.14), 290
Douady, R., 470
Double barrier option, see Barrier options, double
Double binary options, see Binary options
(American), double
Double bubble, 190
Double knock-out, 474
Drift:
adding (complexity of the forward line),
321-323
constant/known (assumption of Black-Scholes-
Merton), 111
and risk neutrality, 427-430
Driftwood effect, 407
Dubins, L., 65, 66
Dubins-Savage optimal strategy, 65-66
Dupire, B., 164, 344, 345, 373
Dupire-Derman-Kani technique, 344-345, 373
Duration of arbitrage, 84
Dvega-Dspot, 191
Dynamic hedging, 258
with Black-Scholes-Merton, 109
introduction, 1-5
real world (principles of), 1-3
vs. statistical trading, 289-290
worst-case scenario, delta hedger, 270

Early exercise, 24-29
Edge/skilled traders, 64-65
E-GARCH, see ARCH (autoregressive conditional
heteroskedasticity)
Empirical volatility weightings, 151-152
Engle, R., 102, 142
Equities market, 26, 33, 34
Eurodollars, 15, 91
and convexity, 188-190
correlation matrix (Table 12.4), 219
future, mapping convergence for, 215
European options, 18, 20, 22, 24, 26
and Asian options, 410, 411
barrier option, 40
and bucketing, 229-231
double bet, 40
forward delta, 231
rules, and ranking of securities, 453
simple vega of, 147
Exotic options, see Asian options; Barrier options,
double; Barrier options, regular; Barrier
options, reverse; Binary options (American);
Binary options (European); Chooser options;

s
[End OCR]*

## Page 512

*[Image OCR]
502 Index

Exotic options (Continued)
Compound options; Higher-order options;
Lookback options; Multiasset options
trading and hedging, 271-272
Expiration pin risks, 21, 368, 222-228, 222, 290
Expiration profile, 353
Expiration of vanilla option, and bleed, 196-197
Exploding option, 364-365
Exponential decay, 95
Extrema, 407

Fair dice, 58, 284, 426
and the Dubins-Savage optimal strategy, 65-66
Feller, W., 246
Feynman, Richard, 5
Feynman-Kac, 461-462
Filtering, 95
First exit time, 174, 339. See also Option(s),
duration; Stopping time
First-order volatility trade, 263
Fixed income instruments/ market:
derivatives book, and classification of risks, 33
mapping convergence for, 215
and primary/secondary risks (Table 1.3), 34
Fixed strike, 403
Flat position, 226-228
Floating rate agreements (FRAs), Eurodollars, 15
Floating strike, 403
Floor(s), 34, 40
Floorlets, 34, 382
Floortion, 382
Fooling customers, 351, 353
Foreign cross-currency pairs, and delta, 124
Foreign currency forwards, and delta, 124
Forward delta, European option, 22
Forward implied volatility, 154
Forwards /futures/forward-forwards, 15, 29-32
bucketing, forward or forward-forward, 231-232
correlation between future and financing
(advanced issue), 31-32
delta, 123-126
forward contract, 29
forward-forward, 32
forwards and futures, 19
future, 29
Fractal structure, 419
Franchises, 54
Frictionless markets (assumption of Black-Scholes-
Merton), 110-111
Fung, Bill, 65
Fungibility, 208-213, 208
and changes in rules of game, 212-213
classification of products by degree of (Table
12.1), 211
and option arbitrage, 212
ranking of, 209-210
and term structure of prices, cash-and-carry line,
210-211
Future, 29. See also Forwards /futures/forward-
forwards
delta for, 125~126

Gamma, 11, 112, 132, 132-146
alpha, in definition of, 178, 179
of the back month, correction for, 136-138
for a bet, 285-289
of betspreads, 294
bleed, 191-197, 191, 228
characteristics of short and long (Table 4.1), 78
and delta, continuous-time, 118
delta of the (third moment), 202
down-gamma, 133, 135, 136
of the gamma (fourth moment), 202

GARCH, 142
hedging reverse knock-outs, 356-361
imperfections for a book, 133-136
“letting the gammas run,” 104
lookback, 405
modified, 36
rebalancing, 1
and risk reversal, 135, 136
second moment, 202
shadow, 127, 132, 138-142, 139
advanced, 142-143, 142
simple, 132-133
in time and space (Figure 17.13), 289
up-gamma, 133, 134, 135, 136
and vega, 149-150
Gap(s), 374-375
delia, 300
gap report, 374
risk, 372
GARCH, 102-103. See also ARCH (autoregressive
conditional heteroskedasticity)
E-GARCH (exponential), 102
GARCH gamma, 142
N-GARCH (nonlinear), 102
Garman, M., 57, 103
Garman-Klass estimator, 103
Gatto, M., 405
Geman, H., 468, 474
Geometric Brownian motion, see Brownian motion,
geometric
Girsanov, 34, 295, 339, 430, 469-470
Goldman, B., 405
Gradient, 387
Greater fool theory, 62
Greeks, 10, 11, 34, 181-182
changes in time and space (behavior), 191-207
definition /shortcomings/modification for use of
Black-Scholes-Merton, 112-113
minor, 171-181
state, to match on every, 257-258
stealth and health, 182-183
Greenspan, Alan, 238
Grossman, Sanford, 57, 76
Gulf War, 130

Hard /soft:
American options, 21, 24-26, 24, 25
deltas, soft vs. hard, 262-263
hard-path-dependent option, 46
omega (option duration); hard rule and soft rule,
174-175
optionality, 20
soft-path-dependent option, 46
Harrison, M., 461 y
Health, 183
Heath, D., 37, 114, 174
Hedge ratio, and delta (parting ways), 115
Heteroskedasticity, 102, 217. See also ARCH
(autoregressive conditional heteroskedasticity)
Higher order options, 45, 382
trading and hedging, 272
Histogram, 242
Holes, liquidity, see Liquidity
Homogeneity, 36, 38~41, 272
Homoskedasticity of markets, 449
Hull, J., 114, 176
Hull and White model, 239

Index-amortizing swaps, 85

Index replication, and stacking, 221
Ingersoll, J., 188, 202

Instruments, 9-37

Insurance, portfolio, 75-77
[End OCR]*

## Page 513

*[Image OCR]
Interest rates and carry, correlation between, 252

Intrinsic value, 18

Invisible passive replicator, 110

Invisible risk, 34

Ito Process (assumption of Black-Scholes-Merton),
110

Ito’s lemma explained, 459-462

Jarrow, R., 37, 114, 174
Jeanblanc-Piqué, M., 462
Jensen’s inequality, 411
Jump diffusion, 114

Kalman filter, 95

Kani, I, 344, 345, 373

Kappa, 147

Karatzas, I., 199

Klass, M., 103

Knock-in options, 317~319

Knock-out option, 82, 312-316, 312, 474-475
Kreps, D., 461

Ladder options, 405-408, 405
Legal arbitrage, 87
Leland, H., 77
LIBOR-square (London Interbank Offer Rate), 16
Lifestyles of traders, comparative, 61
LIFFE (London Interbank Financial Future
Exchange), 29
Limit decomposable /decomposition, 40, 110
Limit orders, 70, 78
Linear combinations, 390-395, 390
multiasset options, 383
Linear instruments:
delta for, 123-126
derivatives, 9. See also Derivatives
Liquidity, 68-79
barrier options and liquidity vacuum, 72-73
Black-Scholes, 73-74
holes, 68-69, 68, 73, 300, 313
graphic illustration of (Figure 4.1), 71
“mother of all” (SP500), 74
one-way liquidity traps, 73
and option pricing, 77-79
portfolio insurance, 75-77
risk, 34
management, 70
and VAR, 449
slippage, 68
reverse, 74-75
stop orders and path of illiquidity, 70-72
limit orders, 70
stop-loss orders, 70
and triple witching hour, 75
Live cattle prices, 213
Lo, A., 105
Local (market participant), 48, 79
Local time, 198
Lock delta, 204
Lognormality, 89, 284, 391
Long the barrier, 317
Lookback options, 104, 270, 271, 403-405, 403
and path dependence, 47
and primary/secondary risks (Table 1.4), 35
Lottery effect, 239

McKinley, A. Craig, 105
Macromanagement, 3
Manufacturer/user (categories), 11
Margrabe option, 390
Market(s):

barriers, 224, 224-226

cleared, 374

Index 503

completeness, 461
failures, and limits, 74
making, see Market making
mean-reverting/trending, 346
microstructure (branch of financial theory), 57
mined, 374
vs. Clear, 368
neutrality, 132
stack, market-neutral, 218
Market making, 48-67
adverse selection, 60~61
and autocorrelation of price changes, 58
book runners vs. price takers, 48-50
participants, 49
arbs, 48
local, 48
paper, 48
and price for immediacy, 57-58
products, commoditized and nonstandardized,
50~53
and profitability, illusion of, 58-60
proprietary departments, 54-56
signaling, 60, 61
tacit rules (market making), 56-57
Markov assumption/component of Brownian
motion, 73
Marks-to-market differences, futures / forwards,
30-31
Martingale/submartingale, 58
Mathematica™ program, 477
Mean-reverting market, 346
Mechanical stability, 81
Memory, ills of, 73-74
Memoryless Brownian motion, 73
Merton’s jump diffusion, 114
Micromanagement, 3
Microstructure, market, 57
Middlebrow:
convexity, 185
swaps head trader, 254
Miller, Merton, 57
Mined market, 374
Mirror image, 22, 78
Models, 109, 114
Modified delta, see Delta, modified
Modified gamma, see Gamma, modified
Modified theta, see Theta, modified
Modified vega, see Vega, modified
Moments, 202-204, 264-265
Monkeys on a typewriter, 64~67
Monte Carlo, 175, 462, 245, 410, 411
Morton, A., 37, 114, 174
Multiasset bets, 294
Multiasset options, 44, 383-402, 466-467
basket options, 391 s
choice, 384-390
composite underlying securities, 395
correlated and uncorrelated Greeks, 387-389
correlation issues, 392
and correlation triangles, 438-444
correlation vega, 386
linear combinations, 390-395, 390
lognormality, 391
outperformance option, 389
structure by construction, 44
trading and hedging, 272
Multidimensional options, 35

Neutral spreading, 258

N-GARCH, see ARCH (autoregressive conditional
heteroskedasticity)

Noncontingent time-dependent nonlinear
derivatives, see Derivatives’
[End OCR]*

## Page 514

*[Image OCR]
504 Index

Nonlinear derivative, 9
Nonparallel accounting, 249-250
Nonstandardized products, 50-53
Non-time-homogeneous risks/structures, 39,
40-41, 272
Numeraire, 124, 128, 383, 431
conclusion, 435
flipping, 435
mathematical note, 436-437
relativity, 431-437
and two-country paradox, 433-435
Numerica! stochastic integration, 477

O'Connell, Marty, 115
O'Hara, M,, 61
Omega, 113, 174-178, 174
Option(s), 16-20. See also American options; Binary
options (American); Binary options
(European); European options; Option pricing
analysis, dimensions of, for understanding and
pricing, 38-47
on a basket, 40
duration, 174. See also Omega
equivalence, basic rules of, 20-24
replication, 256, 256-258
risks defined (Figure 1.6), 36
simple, 18-20
on a spread, 40
symmetry, 324, 336
trading concepts, 256-269
vanilla/exotic, 38
Option, generalized, 38-47
barriers (step 3), 43
dimension of the structure and number of assets
(step 4), 43-45
dimension of the structure, 43
multiasset structure by construction, 44
homogeneity of structure (step 1), 38-41
time-homogeneous structure, 38
order of the options (step 5), 45-46
higher order options, 45
Path dependence (step 6), 46-47
hard-path-dependent option, 46
path-dependent options, 46
path-independent options, 46
soft-path-dependent option, 46
payoff type, continuous and discontinuous (step
2), 41-42
digital (discontinuous) and ramp (continuous)
payoffs, 41
Optionality, 20
Option pricing, 239, 459-477
Black Scholes equation, 463-464
barrier options, 468-477
compound and chooser order options, 467-468
Ito’s lemma, 459-462
and liquidity, 77-79
multiasset options, 466-467
simple explanation of, 239
stochastic volatility model, 464-466
Option wizard:
barriers and price manipulation, 368
Black-Scholes-Merton (essence of), 110
bleed, forward and backward, 200
carry vs. convergence, 215
comparative lifestyles of traders, 61
contamination (or convexity) principle, 14-15
contamination principle and barrier options, 295
contamination principle and LIBOR square, 17
convexity, middlebrow vs. modified, 185
correlation and volatility, 89
culture shocks, 5

curve and surface shifts, 162-163
delta and the probability of being in the money,
128
derivative trade, first, 13
difficult boss (another war story), 192
diffusion (random walk on a spreadsheet), 419
Dirac Delta, 285
European bet, pricing, 284
European Monetary System, war story, 227
fooled by the expiration profile, 353
gap delta, 300
“GARCH in their head,” 102
the Greeks, 10
hedger’s viewpoint, 10
how bad salespeople fool their customers, 351
how to get into trouble hedging, 45
interpolation, joys of, 172
Ito calculus, modicum of, 198-199
legality of the triggering, 367
linear and nonlinear securities, 11
looking at a graph through time, 301
market making and the burnout of traders, 60
option strategies, basic forms of, 261
option symmetry, 324
parametric and nonparametric tests, 204
Pareto-Levy, graphical representation of, 246
“Perhaps Bachelier Was Right,” 91
pitfalls for risk managers, 134
pseudovanilla, 393
replication map, 280
risk-neutral argument, why traders know, 429
risk reversal (revisited), 275
rubber time explained, 98
simple/complex trades/ products, 261
simplicity rule, 26
smallest decomposable fragment of a structure, 40
“Thank God It’s Friday,” 193
. the half billion-dollar delta, 117
tick, dividing a, 277
too much hedging is bad for you, 311
trader decay, 52
trees, discrete states of parameters on, 176-177
vexed question, 254
why good models die, 114
Order, 45-46
Outbarrier options, 363
Outperformance options, 389, 467
Outstrike, 312

Paper (market participant), 48, 79

Parallel shift, 162, 163

Parameter convexity, 36

Parkinson number, 101-107, 345

Partial deltas, 387-388, 389

Path dependency, 36,46—47, 248
and bucketing, 231
and bucket vega, 157
of exotic options, 363
hard-path-dependent option, 46
path-dependent options, 46
path-independent options, 46
of a regular option, 265-270
soft-path-dependent option, 46

Payoff, type of (continuous/discontinuous), 41-42

PIBOR, 130

Pin risks, see Expiration pin risks

Pit, SP500, 76

Pit chemistry, 76

Pit etiquette, 79

Pliska, S., 461

Poisson jump size, 114

Popper, Sir Karl, 9
[End OCR]*

## Page 515

*[Image OCR]
Portfolio insurance, 75-77
Premature exercise, rules of, 175
Price for immediacy, 57-58
Price takers, vs. book runners, 48-50
Primary/secondary exposures, 228
Probabilistic fairness, 426
Probabilistic methods, 462
, Product or quotients (multiasset options), 383
Proprietary positions, 54-56
Proust, Marcel, 191
Put, 18, 19, 472-474
Put-call parity rules, 20-24

advanced, 252-255

European option, 20

and risk-neutral argument, 429
Put/call symmetry and hedging, 323-331

Quants/traders, 5, 383
Quasi-linear securities, 10, 12
Quick-and-dirty hedge, 216

Rainbow options, 383, 384-389, 384, 466
dual or multiple strikes, put-call parity rule, 253
Ramp, see Digital (discontinuous) and ramp
(continuous) payoffs
Random volatility, 238-242
Random walk, 89, 415
classical one-asset, 415-416
three-asset, 424-425
two-asset, 420-424
Ranking of securities, 453-457
Ratio backspread, 264, 265
Ratio write, lock delta for, 206, 207
Rebalancing the gamma, 1
Rebates, 363, 474
Recursive option replication (Figure 16.2), 258
Reflection principle, 336-339, 340, 469
Rent, 11, 178
Replication, option, see Option(s), replication
Report reading, see Risk(s)/risk management,
report reading
Reversal, risk, see Risk reversal
Reverse assets, 251
Reverse slippage, see Slippage, reverse
Rho, 36, 171-174, 171, 178
modified, 173
Rhod, 374
Rhof, 374
Rhop, 173
Rhol, 113, 173, 178, 228
Rho2, 113, 173, 178, 228
Risk(s)/risk management, 3-4, 32-35
and adverse selection and signaling, 60-62
commoditized products, 51-53
framework (Figure 1.1), 4
and liquidity, 70
macromanagement, 3
measuring, 109
method of squares for, 164-166
micromanagement, 3
primary/secondary risks, distinction, 32-35
report reading, 368-375
Risk-neutrality, 128, 426-430, 463-464
Risk reversal, 135, 268, 275, 314, 323
Rolled, 429
Rollover option, lookback options, 404-405
Rosenberg, J., 142
Ross, S., 188, 372
Rotation (second order curve/surface shift), 162,
163
Rubber trees, 225
Rubinstein, Marc, 177, 345, 372

Index 505

Rudd, A., 114
Russell, B., 13

Savage, L., 65, 66
Savery, Howard, 273
Schwager’s New Market Wizards, 62
SCUD (second currency underlying), 363-364
SDF, see Smallest decomposable fragment (SDF)
Secondary exposure, 228
Second-order volatility trade, 264
Shadow reports, 369
Sharpe ratios, 66
Shreve, S., 199
Signaling, 60, 61
Skew, 164, 202, 238, 245-252, 245, 426-427
barrier decomposition under skew environments,
331-335
and biased assets, 248-249
correlation between interest rates and carry, 252
currencies as assets, 250-251
nonparallel accounting, 249-250
paradox, 282-283
pricing European binary options with, 279-289
reverse assets, 251
value linked to price, 250
volatility regimes, 251
Slippage, 68-69, 68, 74-75, 300
Smallest decomposable fragment (SDF), 40, 307, 382
Smart stack, 221
Smile, 27-29, 164, 239, 241
Soft, see Hard/soft
Sosin, H., 405
Spatial dimension, 36
SP500 pit, 76
Spot, 369
bet, 278-279, 278
Ddelta-Dspot, 191
DgammaDspot/Dgamma-Dspot, 118, 191
Dvega-Dspot, 191
grid of expected/spot dependence, 141
Spread(s), 264
classification of derivatives (Figure 1.1), 10
options, 467
Spreading:
advantages of, 259-260
neutral, 258
Spreadsheet, Brownian motion on, 415~425
Squares, method of, 164-166
Squeeze, 84 .
Stability ratio (Ddeltadvol), 134, 200-201
Ddeltadvol, 200
Stacking techniques, 217-221
Standard deviation, 417-418, 419, 420, 422, 424, 425
as function of time (Figure 6.1), 89 ‘
States, 257-258
Static put/call parity (Table 1.1), 21
Static replication, 256
Static straight bucketing, see Bucket(s)/bucketing
Statistical trading vs. dynamic hedging, 289-290
Statistical value of track records, 64-65
Stealth, 183 .
Sticky strikes, 223-224, 223
Stochastic parameter tree, 176
Stochastic volatility model, 464-466
Stock or stock index, and delta, 125
Stop-loss orders, 70
Stop orders, 78
Stopping time, 174, 256, 272, 339, 344, 345, 476
Straight bucketing, see Bucket(s)/bucketing
Strangle write, lock delta for, 206
Strike, fixed/floating, 403
Strike, sticky, see Sticky strikes
[End OCR]*

## Page 516

*[Image OCR]
506 Index

Strike bonus, 405, 406
Strike topography, 233-235
Surface shifts, 162-163. See also Volatility, surface
Swaps, 15, 16, 35
classification of derivatives (Figure 1.1), 10
Swaption, 40
SWARCH, see ARCH (autoregressive conditional
heteroskedasticity)
Swiss Bank Corporation, 259
Syldavian elections (imaginary currency), 143,
145-146, 369, 370-371, 374
Synthetic securities, 12-13, 12

Tails, 238-245, 238
histogram, 242
Tax arbitrage, 86
Thales of Miletus, 13
Theoretical edge, 260
Theta, 112, 167-171, 167
for a bet, 169
in definition of alpha, 178
interest carry, and self-financing strategies,
169-170
modified, 36, 167-169, 169
shadow, 170
weakness of, 171
Third derivatives, 191
Third moment, 275
Thresholds, 241
Time:
dimension, 36
economic vs. actual, 98
effect of, on knock-out options, 339
rubber, 98
Time-dependent linear derivatives, 13-16
Time-homogeneous structure, 38
Topography, 36, 232-237
position topography, 232
static topography, 233
Trader(s), 60, 61, 64, 65
Transaction costs, 78-79
Traps, one-way liquidity, 73
Trees, 176-177
binomial, 175, 176-177, 256, 336, 345
rubber, 225
stochastic parameter, 276
Trending market, 346
Triangle(s), 88, 455. See also Correlation, triangles
Triangular decomposition, 398-402
Trigger, 312-313
Trinomial tree, 176 :
Two-country paradox, 128, 431-437

Upstairs traders, 49
User/manufacturer (categories), 11

Vacuum, liquidity, 72-73
Value-at-risk (VAR), 45, 127, 445, 445-452
simplified examples, 446-449
Value trading vs. greater fool theory, 62
VAR, see Value-at-risk (VAR)
Variance ratio method, 101-107, 345-346
Vasicek, O., 188
Vega, 112, 147-166, 147
asymptotic vega test, 201
bucket, 113, 153, 158-161
convexity, 184
decomposition of, by squares (Figure 9.9), 166
effect (Table 21.1), 379
and gamma, 149-150
hedging reverse knock-outs, 356-361

linearity; costs of dynamic hedging, 378-379

modified, 36, 147, 150-151, 150

multifactor, 158-159

stability skew, 36

in tables usually memorized by market makers,
181-182

with time (Figure 9.1), 148

and volatility (Figure 9.2), 149

and volatility surface; smile and skew, 164

vs. gamma, 260-262

Venezuelan par bonds, 368
Volatility, 88-107, 88

actual, 88

and alpha, 179, 180

augmented, 77

barrier options, effects of, 319-321

betting, 263-265

bleed, with changes in, 195-196

calculating historical, 92-94

constant (assumption of Black-Scholes-Merton),
lil

constant (“no such thing”), 97-99

and convergence, 216

and delta, 127-131

extreme, and delta, 96, 127-131

filtering, 95-97

forward implied, 154-161, 154

fudging (for Black-Scholes-Merton), 114

as function of ratio of forward to strike (Figure
9.8), 166

grid of expected/spot dependence, 141

impact of volatility of (Table 15.1), 240

implied, 88, 164, 438

forward, 154-161, 154

map of, at different price levels (Table 8.4), 141

map of volatility and interest rate differentials at
different price levels (Table 8.5), 144

. Parkinson number, 101-107

random, 238-242

regimes, 251

rules (Table 16.1), 263

sensitivity of an American digital (Figure 18.2),
297

single volatility fudge, 343~344

surface, 114, 147, 163, 164

trading, initiation to, 260-262

vanilla, 77

and VAR, 451

variance ratio method, 101~107

and vega, see Vega

weightings, simple, 151-153

War stories:
crash of 1987, 27
difficult boss, 192
“hedging increases your risks,” 311
Weekends, and trading, 98, 193
Weights/weightings, 174, 181. See also Volatility,
weightings, simple
Whalley, A., 77
White, A., 114
Wilmott, P., 77
Worst-case scenario, delta hedger, 270

Yass, Jeff, 62
Yor, M., 468, 474

Zeidan, N., 347
Zeta, 147
[End OCR]*
