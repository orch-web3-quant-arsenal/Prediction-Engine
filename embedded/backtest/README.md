# Prediction-Engine

![GitHub stars](https://img.shields.io/github/stars/Archsec-Emman/Prediction-Engine?style=social)
![GitHub forks](https://img.shields.io/github/forks/Archsec-Emman/Prediction-Engine?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/Archsec-Emman/Prediction-Engine?style=social)

[![Licensing: Mixed](https://img.shields.io/badge/licensing-MIT%20%2B%20LGPL--3.0--or--later-blue.svg)](NOTICE)
[![Ruff](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/charliermarsh/ruff/main/assets/badge/v2.json)](https://github.com/astral-sh/ruff)
[![uv](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/astral-sh/uv/main/assets/badge/v0.json)](https://github.com/astral-sh/uv)
![Python](https://img.shields.io/badge/python-3.12%2B-3776AB?logo=python&logoColor=white)
![Rust](https://img.shields.io/badge/rust-1.93.1-CE422B?logo=rust&logoColor=white)
![Rust Edition](https://img.shields.io/badge/edition-2024-CE422B?logo=rust&logoColor=white)
![NautilusTrader](https://img.shields.io/badge/NautilusTrader-1.226.0-1E3A5F)
![GitHub last commit](https://img.shields.io/github/last-commit/Archsec-Emman/Prediction-Engine)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/Archsec-Emman/Prediction-Engine)
![GitHub code size](https://img.shields.io/github/languages/code-size/Archsec-Emman/Prediction-Engine)
![GitHub top language](https://img.shields.io/github/languages/top/Archsec-Emman/Prediction-Engine)
![GitHub open issues](https://img.shields.io/github/issues/Archsec-Emman/Prediction-Engine)
![GitHub contributors](https://img.shields.io/github/contributors/Archsec-Emman/Prediction-Engine)
![GitHub pull requests](https://img.shields.io/github/issues-pr/Archsec-Emman/Prediction-Engine)
![GitHub closed issues](https://img.shields.io/github/issues-closed/Archsec-Emman/Prediction-Engine)
![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/Archsec-Emman/Prediction-Engine)

**New in Version 4.1-alpha:**
- Live sandbox plumbing for Polymarket BTC 5min markets
- Example runner showing how to use live BTC 5min hooks
- (strategy & model *now* included)
- open sourced some /private files (felt like it)

**New in Version 4:**
- Nautilus 1.226.0
- Rust-native data conversion
- Faster staged data loading
- Improved materialized caches
- Unified cache/local/archive/API message bus

**New in Version 3:**
- Telonex vendor support
- Local Telonex download script
- Many bug fixes & accuracy improvements
- Book replay order book deltas with trade ticks

**New in Version 2:**
- Nautilus via PyPI in lieu of a subtree
- Better backtest runner classes via EXPERIMENT objects
- IPython notebook support (.ipynb files)
- Joint portfolio multi replay runners
- Growing support for statistical optimizers
- New aggregate charts
- Massive improvements charting gen speed
- an attempt at a Tree-structured Parzen Estimator via Optuna

Looking for the old version? That was renamed to [Version 1](https://github.com/Archsec-Emman/Prediction-Engine/tree/v1)

Backtesting framework for prediction market strategies on
[Polymarket](https://polymarket.com), built on top of
[NautilusTrader](https://github.com/nautechsystems/nautilus_trader) with custom
exchange adapters. [Limitless.exchange](https://limitless.exchange) and
[Opinion.trade](https://opinion.trade) are planned next; [Kalshi](https://kalshi.com) support depends on access to L2 historical book data. Current Kalshi components are research and fee-modeling plumbing, not a public runnable backtest path. Plotting inspired by [minitrade](https://github.com/dodid/minitrade). This repo is still in active development.


Fantastic single & multi-market charting. Featuring: equity (total & individual markets), profit / loss ticks, P&L periodic bars, market allocation, YES price (with green buy and red sell fills), drawdown, sharpe (with above/below shading), cash / equity, monthly returns, and cumulative brier advantage.
![Charting preview](https://raw.githubusercontent.com/Archsec-Emman/Prediction-Engine/main/docs/assets/charting-preview.jpeg)

**If you find any bugs, unexpected behavior, or missing simulation features, PLEASE post an [issue](https://github.com/Archsec-Emman/Prediction-Engine/issues/new) or [discussion](https://github.com/Archsec-Emman/Prediction-Engine/discussions/new/choose).**

Detailed guides have been filed away in the [docs index](https://Archsec-Emman.github.io/Prediction-Engine/) for better organization and long-term sustainability.

## Table of Contents

- [Docs Index](https://Archsec-Emman.github.io/Prediction-Engine/)
  - [Start Here](https://Archsec-Emman.github.io/Prediction-Engine/#start-here)
  - [Core Framework](https://Archsec-Emman.github.io/Prediction-Engine/#core-framework)
  - [Advanced / Experiments](https://Archsec-Emman.github.io/Prediction-Engine/#advanced-experiments)
  - [Project](https://Archsec-Emman.github.io/Prediction-Engine/#project)
  - [Acknowledgements](https://Archsec-Emman.github.io/Prediction-Engine/#acknowledgements)
- [Setup](https://Archsec-Emman.github.io/Prediction-Engine/setup/)
  - [Prerequisites](https://Archsec-Emman.github.io/Prediction-Engine/setup/#prerequisites)
  - [Install](https://Archsec-Emman.github.io/Prediction-Engine/setup/#install)
  - [First Run](https://Archsec-Emman.github.io/Prediction-Engine/setup/#first-run)
  - [Timing And Cache Defaults](https://Archsec-Emman.github.io/Prediction-Engine/setup/#timing-and-cache-defaults)
  - [Extension Architecture](https://Archsec-Emman.github.io/Prediction-Engine/setup/#extension-architecture)
- [Backtests And Runners](https://Archsec-Emman.github.io/Prediction-Engine/backtests/)
  - [Repo Layout](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#repo-layout)
  - [Archived Private Research](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#archived-private-research)
  - [Runner Contract](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#runner-contract)
  - [HTML And Report Modes](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#html-and-report-modes)
  - [Optimization Runners](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#optimization-runners)
  - [Designing Good Runner Files](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#designing-good-runner-files)
  - [Multi-Market Strategy Configs](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#multi-market-strategy-configs)
  - [Running Backtests](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#running-backtests)
  - [Notebook Runners](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#notebook-runners)
  - [Editing Runner Inputs](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#editing-runner-inputs)
  - [Data Vendor Notes](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#data-vendor-notes)
    - [Native Vendors](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#native-vendors)
    - [PMXT](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#pmxt)
    - [Telonex](https://Archsec-Emman.github.io/Prediction-Engine/backtests/#telonex)
- [Sandbox And Live Runners](https://Archsec-Emman.github.io/Prediction-Engine/live/)
  - [Current Scope](https://Archsec-Emman.github.io/Prediction-Engine/live/#current-scope)
  - [Directory Contract](https://Archsec-Emman.github.io/Prediction-Engine/live/#directory-contract)
  - [Sandbox Runner Contract](https://Archsec-Emman.github.io/Prediction-Engine/live/#sandbox-runner-contract)
  - [Shared Live Helpers](https://Archsec-Emman.github.io/Prediction-Engine/live/#shared-live-helpers)
  - [BTC 5m Sandbox Plumbing](https://Archsec-Emman.github.io/Prediction-Engine/live/#btc-5m-sandbox-plumbing)
  - [Example BTC Snapshot Runner](https://Archsec-Emman.github.io/Prediction-Engine/live/#example-btc-snapshot-runner)
  - [Archived Strategy Boundary](https://Archsec-Emman.github.io/Prediction-Engine/live/#archived-strategy-boundary)
  - [Model And Parameter Placement](https://Archsec-Emman.github.io/Prediction-Engine/live/#model-and-parameter-placement)
  - [Running Sandbox](https://Archsec-Emman.github.io/Prediction-Engine/live/#running-sandbox)
  - [Public Polymarket Data](https://Archsec-Emman.github.io/Prediction-Engine/live/#public-polymarket-data)
  - [Path To Live Polymarket Trading](https://Archsec-Emman.github.io/Prediction-Engine/live/#path-to-live-polymarket-trading)
- [Data Loading](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/)
  - [Mental Model](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#mental-model)
  - [Staged Loading](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#staged-loading)
  - [PMXT Flow](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#pmxt-flow)
  - [Telonex Flow](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#telonex-flow)
  - [Caching](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#caching)
  - [Downloading Local Data](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#downloading-local-data)
  - [Progress And Timing](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#progress-and-timing)
  - [Failure Semantics](https://Archsec-Emman.github.io/Prediction-Engine/data-loading/#failure-semantics)
- [Polymarket Account Ledger Replay](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/)
  - [Runner And Notebook](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#runner-and-notebook)
  - [What The Strategy Does](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#what-the-strategy-does)
  - [Why Exact Reproduction Fails](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#why-exact-reproduction-fails)
  - [Copy-Trading Interpretation](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#copy-trading-interpretation)
  - [External Source Check](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#external-source-check)
  - [Observed Result](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#observed-result)
  - [Latest Terminal Output](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#latest-terminal-output)
  - [How To Use This Experiment](https://Archsec-Emman.github.io/Prediction-Engine/account-ledger-replay/#how-to-use-this-experiment)
- [Research](https://Archsec-Emman.github.io/Prediction-Engine/research/)
  - [Overview](https://Archsec-Emman.github.io/Prediction-Engine/research/#overview)
  - [Warm PMXT Cache Before Notebook Runs](https://Archsec-Emman.github.io/Prediction-Engine/research/#warm-pmxt-cache-before-notebook-runs)
  - [Scoring](https://Archsec-Emman.github.io/Prediction-Engine/research/#scoring)
  - [Joint-Portfolio Mode](https://Archsec-Emman.github.io/Prediction-Engine/research/#joint-portfolio-mode)
  - [Samplers](https://Archsec-Emman.github.io/Prediction-Engine/research/#samplers)
    - [Random Grid (`sampler="random"`)](https://Archsec-Emman.github.io/Prediction-Engine/research/#random-grid-samplerrandom)
    - [TPE (`sampler="tpe"`)](https://Archsec-Emman.github.io/Prediction-Engine/research/#tpe-samplertpe)
  - [Caveats](https://Archsec-Emman.github.io/Prediction-Engine/research/#caveats)
  - [Notebook Output Persistence](https://Archsec-Emman.github.io/Prediction-Engine/research/#notebook-output-persistence)
- [Execution Modeling](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/)
  - [Fees](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#fees)
    - [Maker Rebates](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#maker-rebates)
  - [Slippage](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#slippage)
  - [Passive Orders And Queue Position](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#passive-orders-and-queue-position)
  - [Latency](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#latency)
  - [Limits](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#limits)
  - [Vendor L2 Behavior](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#vendor-l2-behavior)
    - [PMXT](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#pmxt)
    - [Telonex](https://Archsec-Emman.github.io/Prediction-Engine/execution-modeling/#telonex)
- [Data Vendors And Local Mirrors](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/)
  - [PMXT](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#pmxt)
    - [Runner Source Modes](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#runner-source-modes)
    - [Lower-Level Loader Env Vars](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#lower-level-loader-env-vars)
    - [What Works Today](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#what-works-today)
    - [Supported Local File Layout](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#supported-local-file-layout)
    - [Required Parquet Columns](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#required-parquet-columns)
    - [Legacy JSON Payload Shape](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#legacy-json-payload-shape)
  - [Telonex](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#telonex)
    - [Download Local Telonex Files](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#download-local-telonex-files)
  - [What Is Not Plug-And-Play Yet](https://Archsec-Emman.github.io/Prediction-Engine/data-vendors/#what-is-not-plug-and-play-yet)
- [Vendor Fetch Sources And Timing](https://Archsec-Emman.github.io/Prediction-Engine/vendor-fetch-sources/)
  - [PMXT](https://Archsec-Emman.github.io/Prediction-Engine/vendor-fetch-sources/#pmxt)
  - [Example Output](https://Archsec-Emman.github.io/Prediction-Engine/vendor-fetch-sources/#example-output)
  - [Telonex](https://Archsec-Emman.github.io/Prediction-Engine/vendor-fetch-sources/#telonex)
  - [Timing Expectations By Source](https://Archsec-Emman.github.io/Prediction-Engine/vendor-fetch-sources/#timing-expectations-by-source)
  - [How To See This Output](https://Archsec-Emman.github.io/Prediction-Engine/vendor-fetch-sources/#how-to-see-this-output)
- [Plotting](https://Archsec-Emman.github.io/Prediction-Engine/plotting/)
  - [Scaling Model](https://Archsec-Emman.github.io/Prediction-Engine/plotting/#scaling-model)
  - [Downsampling](https://Archsec-Emman.github.io/Prediction-Engine/plotting/#downsampling)
  - [Output Types](https://Archsec-Emman.github.io/Prediction-Engine/plotting/#output-types)
  - [Output Paths](https://Archsec-Emman.github.io/Prediction-Engine/plotting/#output-paths)
  - [Example Summary Output](https://Archsec-Emman.github.io/Prediction-Engine/plotting/#example-summary-output)
  - [Multi-Market References](https://Archsec-Emman.github.io/Prediction-Engine/plotting/#multi-market-references)
- [Testing](https://Archsec-Emman.github.io/Prediction-Engine/testing/)
  - [Standard Repo Gate](https://Archsec-Emman.github.io/Prediction-Engine/testing/#standard-repo-gate)
  - [Useful Smoke Checks](https://Archsec-Emman.github.io/Prediction-Engine/testing/#useful-smoke-checks)
  - [Docs Validation](https://Archsec-Emman.github.io/Prediction-Engine/testing/#docs-validation)
- [Project Status](https://Archsec-Emman.github.io/Prediction-Engine/project-status/)
  - [Roadmap](https://Archsec-Emman.github.io/Prediction-Engine/project-status/#roadmap)
  - [Known Issues](https://Archsec-Emman.github.io/Prediction-Engine/project-status/#known-issues)
  - [Recently Fixed](https://Archsec-Emman.github.io/Prediction-Engine/project-status/#recently-fixed)
- [License Notes](https://Archsec-Emman.github.io/Prediction-Engine/license/)
  - [Scope](https://Archsec-Emman.github.io/Prediction-Engine/license/#scope)
  - [NautilusTrader Attribution](https://Archsec-Emman.github.io/Prediction-Engine/license/#nautilustrader-attribution)
  - [Practical Meaning](https://Archsec-Emman.github.io/Prediction-Engine/license/#practical-meaning)


## Star History

<a href="https://www.star-history.com/?repos=Archsec-Emman%2FPrediction-Engine&type=date&legend=top-left">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/image?repos=Archsec-Emman/Prediction-Engine&type=date&theme=dark&legend=top-left" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/image?repos=Archsec-Emman/Prediction-Engine&type=date&legend=top-left" />
   <img alt="Star History Chart" src="https://api.star-history.com/image?repos=Archsec-Emman/Prediction-Engine&type=date&legend=top-left" />
 </picture>
</a>
