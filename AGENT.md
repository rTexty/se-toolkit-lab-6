# Agent Documentation

## Architecture
The agent is single-file Golang CLI application (`agent.go`). It uses standard library `net/http` to do requests and `encoding/json` to serialize requests/responses. Because minimal dependencies are preferred, godotenv is used for loading secrets out of `.env.agent.secret`.

## Features
- Task 1: Basic functionality implemented. Connects to `chat/completions` API via OpenAI-compatible endpoints. Supports configuring `LLM_API_KEY`, `LLM_API_BASE`, `LLM_MODEL`.
- Prints strictly JSON payload `{ "answer": "...", "tool_calls": [] }` on stdout for correct task parsing. Any logs/errors are routed to stderr.

## How to run
1. Set up `.env.agent.secret` (defining `LLM_API_KEY`, `LLM_API_BASE`, `LLM_MODEL`)
2. Execute the agent:
   ```bash
   go run agent.go "What is the capital of France?"
   ```

## Testing
Run the tests using standard Go tooling:
```bash
go test -v
```
