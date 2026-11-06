# Enterprise Architecture & Security Auditor (CLI)

A high-performance architectural and structural code auditor built in Go (Golang). Designed for engineering teams and systems architects to automatically detect deep concurrency risks, resource leaks, and architectural anti-patterns in complex codebases.

## Features
- **Concurrency Risk Analysis:** Detects uncontrolled goroutine spawns and potential resource/memory leaks.
- **Automated Reporting:** Instantly generates machine-readable JSON compliance reports (`audit_report.json`).
- **Zero Dependencies:** Compiles into a single, lightning-fast native binary.

## Usage
```bash
arch_auditor <path_to_project>cat << 'EOF' > README.md
# Enterprise Architecture & Security Auditor (CLI)

A high-performance architectural and structural code auditor built in Go (Golang). Designed for engineering teams and systems architects to automatically detect deep concurrency risks, resource leaks, and architectural anti-patterns in complex codebases.

## Features
- **Concurrency Risk Analysis:** Detects uncontrolled goroutine spawns and potential resource/memory leaks.
- **Automated Reporting:** Instantly generates machine-readable JSON compliance reports (`audit_report.json`).
- **Zero Dependencies:** Compiles into a single, lightning-fast native binary.

## Usage
```bash
arch_auditor <path_to_project>
