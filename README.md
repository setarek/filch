## Filch

A Golang tool for Persian and Arabic text manipulation.

This package can:

- Convert Arabic alphabet like ``` ي - أ - إ - ئ - ك  ```  to Persian equivalent.
- Convert English and Arabic numbers into Persian numbers.
- Convert Persian or Arabic numbers into English.



## Installation



``` go get github.com/setarek/filch ```



## How?



Convert Arabic alphabet to Persian:

``` go
s := "موسيقي كلاسیك"
filch.ConvertArabicToPersian(s)
// موسیقی کلاسیک
```



Convert English/Arabic numbers to Persian:

```go
s1 := "کلاس شماره 5"
s2 := "کلاس شماره ٥"
filch.ConvertToPersianNumber(s1)
// کلاس شماره ۵
filch.ConvertToPersianNumber(s2)
// کلاس شماره ۵
```



Convert Persian/Arabic numbers to English:

```go
s := "۰۹۱۲۳۴۵۶۷۸۹"
filch.ConvertToEnglishNumber(s)
// 09123456789
```



Combination of these functions:

```go
s := "پلاك شماره ئ 12"
filch.ConvertArabicToPersian(filch.ConvertToPersianNumber(s))
// پلاک شماره ی ۱۲
```

