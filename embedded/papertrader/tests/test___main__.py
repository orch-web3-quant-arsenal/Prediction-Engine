"""Test pe_trader.__main__ module."""

import subprocess
import sys


def test_main_module_runs():
    """python -m pe_trader executes without import errors."""
    result = subprocess.run(
        [sys.executable, "-m", "pe_trader", "--help"],
        capture_output=True, text=True, timeout=10,
    )
    assert result.returncode == 0
    assert "pe-trader" in result.stdout or "Usage" in result.stdout
