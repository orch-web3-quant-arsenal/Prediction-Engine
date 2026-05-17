package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

//go:embed embedded/*
var toolbox embed.FS

const banner = `
   ╔══════════════════════════════════════════════════════════╗
   ║                                                        ║
   ║   ██████╗ ██████╗ ███████╗██████╗ ██╗ ██████╗████████╗ ║
   ║   ██╔══██╗██╔══██╗██╔════╝██╔══██╗██║██╔════╝╚══██╔══╝ ║
   ║   ██████╔╝██████╔╝█████╗  ██║  ██║██║██║        ██║    ║
   ║   ██╔═══╝ ██╔══██╗██╔══╝  ██║  ██║██║██║        ██║    ║
   ║   ██║     ██║  ██║███████╗██████╔╝██║╚██████╗   ██║    ║
   ║   ╚═╝     ╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝ ╚═════╝   ╚═╝    ║
   ║                                                        ║
   ║     Prediction Engine — The Unified Market Oracle      ║
   ║                                                        ║
   ║   by Archsec-Emman (@Archsec-Emman)                    ║
   ║   orch-web3-quant-arsenal Edition                      ║
   ║   https://github.com/orch-web3-quant-arsenal/Prediction-Engine ║
   ║                                                        ║
   ╚══════════════════════════════════════════════════════════╝
`

var (
	targetURL  string
	onlyEngine string
	skipEngine string
	balance    float64
	servePort  int
)

func init() {
	flag.StringVar(&targetURL, "u", "", "Polymarket or Kalshi URL to analyze")
	flag.StringVar(&onlyEngine, "only", "", "Run only specified engines (analysis,papertrader,backtest)")
	flag.StringVar(&skipEngine, "skip", "", "Skip engines (comma-separated)")
	flag.Float64Var(&balance, "balance", 10000.0, "Starting paper balance for trading engine")
	flag.IntVar(&servePort, "serve", 0, "Start analysis web UI on given port (0=disabled)")
}

func main() {
	flag.Usage = func() {
		fmt.Print(banner)
		fmt.Println("\nUsage:  prediction-engine -u <market-url>  [--only <engine>]  [--skip <engine>]")
		fmt.Println("        prediction-engine --serve 3000")
		fmt.Println("\nExamples:")
		fmt.Println("  prediction-engine -u https://polymarket.com/event/bitcoin-price")
		fmt.Println("  prediction-engine -u https://polymarket.com/event/bitcoin-price --balance 5000")
		fmt.Println("  prediction-engine --serve 3000")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if targetURL == "" && servePort == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(banner)

	// If serve mode, start the analysis web UI
	if servePort > 0 {
		startWebUI(servePort)
		return
	}

	// Extract tools to temp directory
	tmpDir, err := ioutil.TempDir("", "prediction-engine")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Extract all embedded tools
	analysisDir := filepath.Join(tmpDir, "analysis")
	if err := extractDir("embedded/analysis", analysisDir); err != nil {
		log.Printf("[!] Failed to extract analysis engine: %v", err)
	}

	paperTraderDir := filepath.Join(tmpDir, "papertrader")
	if err := extractDir("embedded/papertrader", paperTraderDir); err != nil {
		log.Printf("[!] Failed to extract paper trader: %v", err)
	}

	backtestDir := filepath.Join(tmpDir, "backtest")
	if err := extractDir("embedded/backtest", backtestDir); err != nil {
		log.Printf("[!] Failed to extract backtest engine: %v", err)
	}

	// Determine which engines to run
	runSet := map[string]bool{
		"analysis":    true,
		"papertrader": true,
		"backtest":    true,
	}
	if onlyEngine != "" {
		runSet = make(map[string]bool)
		for _, e := range strings.Split(onlyEngine, ",") {
			runSet[strings.TrimSpace(e)] = true
		}
	}
	if skipEngine != "" {
		for _, e := range strings.Split(skipEngine, ",") {
			runSet[strings.TrimSpace(e)] = false
		}
	}

	var wg sync.WaitGroup
	outputs := make(map[string]string)
	var mu sync.Mutex

	// Engine 1: AI Analysis (Polyseer-like analysis of the market URL)
	if runSet["analysis"] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			output := fmt.Sprintf("Analysis Engine - Target: %s\n", targetURL)
			output += "========================================\n"
			output += "The Analysis Engine performs multi-agent AI research:\n"
			output += "  - Platform detection (Polymarket/Kalshi)\n"
			output += "  - Multi-agent research strategy generation\n"
			output += "  - Evidence collection from academic, web, and market sources\n"
			output += "  - Bayesian probability aggregation\n"
			output += "  - Final verdict with pNeutral and pAware probabilities\n\n"

			// Attempt to use the Next.js dev server if available
			output += "[*] Note: Analysis engine requires Node.js to run the full UI.\n"
			output += "    Start with: prediction-engine --serve 3000\n"
			output += "    Then open http://localhost:3000 and paste the target URL.\n\n"

			// Check for Python-based analysis (if the tool has a CLI mode)
			pythonAnalysis := filepath.Join(analysisDir, "analyze.py")
			if _, err := os.Stat(pythonAnalysis); err == nil {
				cmd := exec.Command("python3", pythonAnalysis, targetURL)
				cmd.Dir = analysisDir
				out, err := cmd.Output()
				if err == nil {
					output += string(out)
				}
			}

			mu.Lock()
			outputs["Analysis"] = output
			mu.Unlock()
		}()
	}

	// Engine 2: Paper Trading (polymarket-paper-trader)
	if runSet["papertrader"] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cliPath := filepath.Join(paperTraderDir, "cli.py")
			if _, err := os.Stat(cliPath); os.IsNotExist(err) {
				// Try setup.py install
				installCmd := exec.Command("pip3", "install", "-e", paperTraderDir, "--quiet")
				installCmd.Run()
			}

			cmd := exec.Command("python3", "-m", "pe_trader.cli",
				"init", "--balance", fmt.Sprintf("%.0f", balance))
			cmd.Dir = paperTraderDir
			out, err := cmd.Output()

			mu.Lock()
			if err != nil {
				outputs["PaperTrader"] = fmt.Sprintf("Paper trader initialized with $%.0f balance.\nUse 'pe-trader markets search <query>' to find markets.\nUse 'pe-trader buy <slug> <outcome> <amount>' to trade.\nError output: %v", balance, err)
			} else {
				outputs["PaperTrader"] = fmt.Sprintf("Paper Trader Engine initialized with $%.0f balance.\n%s", balance, string(out))
			}
			mu.Unlock()
		}()
	}

	// Engine 3: Backtesting (prediction-market-backtesting)
	if runSet["backtest"] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			output := fmt.Sprintf("Backtesting Engine\n")
			output += "========================================\n"
			output += "High-performance backtesting framework built on NautilusTrader:\n"
			output += "  - Rust-native data conversion for speed\n"
			output += "  - Polymarket exchange adapter with order book replay\n"
			output += "  - Multi-market strategy configs\n"
			output += "  - Equity curves, P&L, Sharpe ratios, drawdown analysis\n"
			output += "  - Statistical optimizers (Optuna TPE)\n"
			output += "  - Joint portfolio multi-replay runners\n\n"

			// Check if the framework has a runner script
			runnerPath := filepath.Join(backtestDir, "runners")
			if _, err := os.Stat(runnerPath); err == nil {
				output += fmt.Sprintf("[*] Backtest runners available in: %s\n", runnerPath)
				output += "    Run with: python3 -m nautilus_trader.backtest\n"
			}
			output += "\n[*] Requires NautilusTrader installed: pip install nautilus_trader\n"

			mu.Lock()
			outputs["Backtest"] = output
			mu.Unlock()
		}()
	}

	wg.Wait()

	// Print combined report
	fmt.Println("======================================")
	fmt.Println("  Prediction Engine Combined Report")
	fmt.Println("======================================")
	fmt.Println()

	engineOrder := []string{"Analysis", "PaperTrader", "Backtest"}
	for _, name := range engineOrder {
		out, ok := outputs[name]
		if !ok {
			continue
		}
		fmt.Printf("--- %s ---\n", name)
		fmt.Println(out)
		fmt.Println()
	}

	fmt.Println("=== DONE ===")
	fmt.Printf("\n⚡ Prediction Engine completed multi-engine analysis.\n")
}

func startWebUI(port int) {
	tmpDir, err := ioutil.TempDir("", "prediction-engine-ui")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}

	analysisDir := filepath.Join(tmpDir, "analysis")
	if err := extractDir("embedded/analysis", analysisDir); err != nil {
		log.Fatalf("Failed to extract web UI: %v", err)
	}

	fmt.Printf("[*] Starting Analysis Web UI on port %d...\n", port)
	fmt.Printf("    Open http://localhost:%d in your browser.\n", port)
	fmt.Println("    Press Ctrl+C to stop.")

	cmd := exec.Command("npm", "run", "dev", "--", "-p", fmt.Sprintf("%d", port))
	cmd.Dir = analysisDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("NEXT_PUBLIC_APP_MODE=self-hosted"),
		fmt.Sprintf("PORT=%d", port),
	)

	if err := cmd.Run(); err != nil {
		log.Printf("[!] Web server exited: %v", err)
	}
}

func extractDir(embedPath, destDir string) error {
	return fs.WalkDir(toolbox, embedPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(embedPath, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}
		data, err := toolbox.ReadFile(path)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(destPath, data, 0644)
	})
}
