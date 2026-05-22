# logslice

Fast log filtering utility with time-range and field-based queries.

## Installation

```bash
go install github.com/yourusername/logslice@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/logslice.git && cd logslice && go build ./...
```

## Usage

Filter logs by time range and field values:

```bash
# Filter logs between two timestamps
logslice --from "2024-01-15T08:00:00Z" --to "2024-01-15T09:00:00Z" app.log

# Filter by field value
logslice --field level=error app.log

# Combine time range and field filters
logslice --from "2024-01-15T08:00:00Z" --to "2024-01-15T09:00:00Z" --field service=api --field level=warn app.log

# Read from stdin
cat app.log | logslice --from "2024-01-15T08:00:00Z" --field level=error
```

### Flags

| Flag | Description |
|------|-------------|
| `--from` | Start of time range (RFC3339) |
| `--to` | End of time range (RFC3339) |
| `--field` | Field filter in `key=value` format (repeatable) |
| `--format` | Log format: `json`, `logfmt` (default: `json`) |

## License

MIT © 2024 [yourusername](https://github.com/yourusername)