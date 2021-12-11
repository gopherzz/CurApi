# CurApi
Simple api for converting currencies without authorization
Exchange Rates parsed from [www.xe.com](https://www.xe.com/) using [goquery](https://github.com/PuerkitoBio/goquery)

## Usage
```
/convert?
  from=<FROM CURRENCY CODE>
 &to=<CONVERT TO CURRENCIES COMMA SEPARATED>
 &amount=<AMOUNT OF FROM CURRENCY>
```
- from string (currency code)
- to []string (currencies codes comma separated)
- amount float64 (amount of from currency)
