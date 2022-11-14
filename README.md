## What is Query?

This is a coding practice for myself to provide a text processor with a
SQL-styled syntax. The goal is to practice a) how to write a SQL syntax
parser by hand, and b) how to design an API to allow manipulation of
data.

## Quick start

A conceptual code example looks like below. It tries to download a CSV
file from an S3 bucket, then find lines matching a condition.

```go
package main

import (
	"time"

	"github.com/fuzhouch/query"
	"github.com/fuzhouch/query/format/csv"
	"github.com/fuzhouch/query/media/s3"
)

func main() {
	media := s3.New(s3.WithKeyID("..."),
		s3.WithSecret("...."),
		s3.WithSSL(true),
		s3.WithBucket("mybucket"))
	defer media.Close()

	file := csv.New("/address-book.csv")
	defer file.Close()

	source := query.From(media, file)
	expr := query.Select("*").Where(query.Eq(query.Field("name"), "fuzhouch"))

	scanner := expr.Scanner()
	for scanner.Next() {
		data := scanner.Text()
        fmt.Printf("%v\n", data)
	}
```
