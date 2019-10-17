# Excelize Infinite Loop Repro

I'm running into an issue with excelize v2.0.2 never returning false from the rows iterator. To reproduce:

## Working v2.0.1

```bash
go get github.com/360EntSecGroup-Skylar/excelize/v2@v2.0.1
go run main.go
```

Notice that after running the total row count is 11.

## Broken v2.0.2

```bash
go get github.com/360EntSecGroup-Skylar/excelize/v2@v2.0.2
go run main.go
```

Notice that the program never exits.
