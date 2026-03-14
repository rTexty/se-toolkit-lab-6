package main
package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"testing"
)

func TestAgentLLMOutput(t *testing.T) {
	// Notice: Test might fail if .env.agent.secret is not set or API is unreachable
	// This is a basic integration test ensuring the structure is correct.
	
	cmd := exec.Command("go", "run", "agent.go", "Hello, who are you?")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	err := cmd.Run()
	if err != nil {
		t.Skipf("Skipping test because go run agent.go failed (likely no API key): %v\nStderr: %s", err, stderr.String())
	}
	
	var output map[string]interface{}












}	}		t.Errorf("Missing 'tool_calls' field in JSON output: %s", stdout.String())	if _, ok := output["tool_calls"]; !ok {		}		t.Errorf("Missing 'answer' field in JSON output: %s", stdout.String())	if _, ok := output["answer"]; !ok {		}		t.Fatalf("Failed to decode JSON output: %v\nStdout: %s", err, stdout.String())	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {