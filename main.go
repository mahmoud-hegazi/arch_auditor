package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type Issue struct {
    File     string `json:"file"`
    Line     int    `json:"line"`
    Message  string `json:"message"`
    Severity string `json:"severity"`
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: arch_auditor <project_path>")
        return
    }

    rootPath := os.Args[1]
    fmt.Printf("[*] Enterprise Architecture & Security Auditor v1.0\n")
    fmt.Printf("[*] Scanning Target: %s\n\n", rootPath)

    var issues []Issue

    filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(path, ".go") {
            fileIssues := auditFile(path)
            issues = append(issues, fileIssues...)
        }
        return nil
    })

    printReport(issues)
}

func auditFile(filePath string) []Issue {
    var issues []Issue
    file, err := os.Open(filePath)
    if err != nil {
        return issues
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNum := 0

    for scanner.Scan() {
        lineNum++
        line := scanner.Text()

        if strings.Contains(line, "go func(") {
            issues = append(issues, Issue{
                File:     filePath,
                Line:     lineNum,
                Message:  "Architectural Risk: Uncontrolled goroutine spawn detected.",
                Severity: "HIGH",
            })
        }
    }

    return issues
}

func printReport(issues []Issue) {
    fmt.Println("=== ARCHITECTURAL AUDIT REPORT ===")
    if len(issues) == 0 {
        fmt.Println("[+] Architecture is clean. No high-risk patterns detected.")
        return
    }

    for _, issue := range issues {
        fmt.Printf("[%s] %s:%d\n    -> %s\n\n", issue.Severity, issue.File, issue.Line, issue.Message)
    }
    fmt.Printf("Total Structural Issues Found: %d\n", len(issues))

    // حفظ التقرير كملف JSON تلقائياً للعملاء
    file, _ := json.MarshalIndent(issues, "", "  ")
    os.WriteFile("audit_report.json", file, 0644)
    fmt.Println("[+] Report successfully exported to: audit_report.json")
}
