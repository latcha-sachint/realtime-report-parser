## Report Parser in Go

This project was created purely out of personal interest and is not associated with any client-specific or internal requirements. While the data used is MIQ-specific, the primary goal was to experiment with writing `Go` code.

### Features

The program processes a CSV file and parses it into a specified format with the following columns:

- **vin**: `string`
- **dealer_code**: `int`
- **created**: `time.Time` (equivalent to `DateTime` in JavaScript)
- **overall_severity**: `string`
- **delivery_status**: `string`
- **lead_id**: `string`

The parsing logic includes basic validation checks and avoids unnecessary complexity.

### Performance Benchmark

The program was benchmarked against other languages/runtimes using similar logic. Results are as follows:

| Language             | Execution Time |
|----------------------|----------------|
| Go                   | 30.38ms        |
| Python               | 359.44ms       |
| Deno (TS/JS runtime) | 170.72ms       |

No external dependencies were used; the implementation relies solely on Go's standard library.

---

### Running the Program

Using the `go` CLI:
```bash
go run .
```

Using `make`:
```bash
make run
```

---

### Building the Binary

Using the `go` CLI:
```bash
go build -o bin/main
```

Using `make`:
```bash
make build
```

---

### Testing the Program

Run tests:
```bash
go test
```

Check code coverage:
```bash
go test -cover
```