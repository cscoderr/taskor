# Taskor

**A concurrent CLI task runner built with Go**

`taskor` is a lightweight command-line task runner written in Go.
It executes shell commands concurrently using a worker pool, supports graceful shutdown, and prints clean, structured output for each job.

This project was built to deeply understand Go’s concurrency model (goroutines, channels, context, WaitGroup) by building something real.

--- 

## Features

* **Concurrent Worker Pool**: Efficiently manage task execution using a fixed or dynamic set of workers.
* **Context-Aware**: Built-in support for timeouts and cancellations using Go `context`.
* **Graceful Shutdown**: Listens for `SIGINT` (Ctrl+C) and `SIGTERM` to stop tasks safely.
* **JSON-Driven**: Define your entire workflow in a simple, portable JSON format.
* **Isolation**: Captures `stdout` and `stderr` separately for every job to keep your terminal clean.
	
---

## Installation

```bash
# Clone the repository
git clone https://github.com/cscoderr/taskor.git
cd taskor

# Build the binary
go build -o taskor
```

## Usage

**1. Create a jobs file**

Create a JSON file describing the jobs you want to run:

```json
{
  "Print Hello": "echo Hello World",
  "Print Name": "echo Tomiwa Idowu",
  "Ping Google": "ping -c 3 google.com",
  "Check Disk": "df -h",
  "Show Date": "date"
}
```
Each key is the job name, and each value is the shell command to execute.

**2. Run taskor**

```bash
./taskor -jobs jobs.json -workers 4
```

**Flags:**
| Flag | Description | Default |
|---------|-----------|---------|
| `-jobs` | Path to JSON jobs file | `required` |
| `-workers` | Number of concurrent workers | `NumCPU()` |

## Sample Output

```bash
======================================================================
JOB #1: Print Hello
COMMAND: [-c echo Hello World]
----------------------------------------------------------------------
Status : Done
Output :
Hello World
======================================================================

======================================================================
JOB #2: Ping Google
COMMAND: [-c ping -c 3 google.com]
----------------------------------------------------------------------
Status : Done
Output :
PING google.com ...
======================================================================
```

## Testing

Run all tests:

```bash
go test ./...
```
