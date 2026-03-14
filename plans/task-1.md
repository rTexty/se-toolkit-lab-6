# Task 1: Call an LLM from Code (Golang Implementation)

## Objective
Implement a basic CLI agent in Go that takes a question as an argument, calls an LLM via API, and returns a JSON response containing the answer and an empty list of tool calls.

## Architecture & Tools
- **Language**: Go 1.24
- **Libraries**: standard library (`net/http`, `encoding/json`, `os`, `flag`). Minimal dependencies.
- **Provider**: Based on `.env.agent.secret` (Qwen Code API or other OpenAI-compatible).

## Data Flow
1. Parse command-line argument for the user's question.
2. Read environment variables (`LLM_API_KEY`, `LLM_API_BASE`, `LLM_MODEL`).
3. Construct a JSON payload for the OpenAI Chat Completions API format.
4. Send an HTTP POST request to the LLM API.
5. Parse the JSON response.
6. Print the required JSON output format to `stdout`.
7. Any debug/logging information must be printed to `stderr` to not pollute the JSON output.

## Output Format
```json
{
  "answer": "LLM response here",
  "tool_calls": []
}
```

## Error Handling
- Missing CLI arguments -> Write error to `stderr` and exit code 1.
- Missing environment variables -> Write error to `stderr` and exit code 1.
- HTTP Request failures -> Write error to `stderr` and exit code 1.

## Testing
- Unit/Regression test: `agent_test.go` calling the function logic and verifying the JSON structure.
