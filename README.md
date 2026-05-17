# 🔮 Prediction Engine

### The Unified Prediction Market Oracle

**Prediction Engine** combines three independent prediction‑market tools into a single, self‑contained binary — AI‑driven market analysis, high‑fidelity paper trading, and institutional‑grade strategy backtesting. Deploy all three engines in parallel against any Polymarket or Kalshi event, and receive a combined intelligence report that covers *why* a market moves, *how* to trade it, and *what* strategies have historically worked.

---

## 💡 How Prediction Engine Works

The tool embeds three complementary engines. Each attacks the prediction market problem from a different angle.

### Engine 1 — AI‑Driven Market Analysis

This engine performs deep, multi‑agent research on any market URL:

- **Platform Detection** — automatically identifies Polymarket or Kalshi events and fetches real‑time data
- **Multi‑Agent Orchestration** — a Planner agent generates research strategies with sub‑claims and search seeds; Research agents query academic papers, news, market data, and proprietary intelligence
- **Evidence Classification** — each source is graded A (primary) through D (speculative) with verifiability, consistency, independence, and recency scores
- **Bayesian Probability Aggregation** — log‑likelihood ratios combined with correlation adjustments produce two probabilities: `pNeutral` (objective) and `pAware` (market‑informed)
- **Final Verdict** — a structured analyst‑grade report with evidence weights, confidence intervals, and key insights

The analysis engine uses large language models (GPT‑4o / GPT‑5) for reasoning and Valyu for real‑time research access. Self‑hosted mode requires only API keys and runs entirely on your machine.

### Engine 2 — High‑Fidelity Paper Trading

This engine simulates real Polymarket trades with institutional accuracy:

- **Level‑by‑Level Order Book Execution** — orders walk the real ask/bid book, consuming liquidity at each price level
- **Exact Fee Model** — uses Polymarket's actual fee formula: `bps/10000 × min(price, 1-price) × shares`
- **Slippage Tracking** — every trade records fill quality in basis points vs. the midpoint
- **Limit Order State Machine** — full lifecycle for GTC (good‑til‑cancelled) and GTD (good‑til‑date) orders
- **Strategy Backtesting** — replay any strategy against historical price snapshots
- **Multi‑Outcome Markets** — supports binary (YES/NO) and multi‑outcome markets
- **MCP Server** — 26 Model Context Protocol tools for AI agents to trade autonomously

Start with $10,000 paper balance and trade real order books with zero risk. Your paper P&L matches real P&L within the spread.

### Engine 3 — Institutional‑Grade Backtesting

This engine provides production‑grade strategy backtesting built on NautilusTrader:

- **Rust‑Native Data Conversion** — high‑performance data pipeline with Parquet‑based catalog storage
- **Polymarket Exchange Adapter** — custom adapters with order book deltas, trade ticks, and book replay
- **Multi‑Market Strategy Configs** — design and test strategies across any number of markets simultaneously
- **Statistical Optimizers** — Tree‑structured Parzen Estimator (TPE) via Optuna for hyper‑parameter optimization
- **Rich Charting** — equity curves, individual market P&L, drawdown, Sharpe ratio, monthly returns, cumulative Brier advantage
- **Joint Portfolio Runners** — run multiple strategies with isolated accounts and compare head‑to‑head
- **Live Sandbox Plumbing** — BTC 5‑minute market hooks for near‑live testing

Supports Polymarket natively, with Telonex data vendor integration and planned Limitless.exchange / Opinion.trade support.

**All three engines run concurrently.** You get AI analysis, trading simulation, and strategy backtesting in a single command.

---

## 🧰 Features

| Feature | Description |
|---------|-------------|
| 🔬 **AI Market Analysis** | Multi-agent LLM research with Bayesian probability |
| 📊 **Paper Trading** | Level-by-level order book simulation with real prices |
| ⚡ **High‑Performance Backtesting** | Rust-native engine via NautilusTrader |
| 🧠 **Engine Selection** | Use `--only` or `--skip` to pick specific engines |
| 📋 **Unified Reporting** | Clean, sectioned output across all three engines |
| 🔒 **Self‑Contained Binary** | No external dependencies at runtime — all engines embedded |
| 🪶 **Static Binary** | Runs on Linux, macOS, and Windows without installers |
| 🌐 **Web UI Mode** | Start the analysis dashboard with `--serve <port>` |

---

## 🚀 Installation

### From Pre‑compiled Binary
Download from [Releases](https://github.com/orch-web3-quant-arsenal/Prediction-Engine/releases)

### Build from Source (Go 1.20+ required)
```bash
git clone https://github.com/orch-web3-quant-arsenal/Prediction-Engine
cd Prediction-Engine
go build -o Prediction-Engine
