# CurApi
Simple api for converting currencies without authorization
Exchange Rates parsed from [www.xe.com](https://www.xe.com/) using [goquery](https://github.com/PuerkitoBio/goquery)

## Build
```bash
go build -o build/curapi-server cmd/main.go
```

## Usage
```bash
./build/curapi --port <PORT [default=8080]>
```
## Api Reference
```
/convert?
  from=<FROM CURRENCY CODE>
 &to=<CONVERT TO CURRENCIES COMMA SEPARATED>
 &amount=<AMOUNT OF FROM CURRENCY>
```
- from string (currency code)
- to []string (currencies codes comma separated)
- amount float64 (amount of from currency)
