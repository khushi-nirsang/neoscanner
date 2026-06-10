# NeoScanner

NeoScanner is a fast MVP vulnerability scanner written in Go.

## Run

```powershell
go run . -u https://example.com -o reports/results.json
go run . -l targets.txt -c 50 -s medium -o reports/results.json
```

Both JSON and HTML reports are generated from the `--output` path.
