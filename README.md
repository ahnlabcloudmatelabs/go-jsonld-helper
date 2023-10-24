<div align="center">

![cloudmate logo](https://avatars.githubusercontent.com/u/69299682?s=200&v=4)

# JSON-LD Helper

<small style="opacity: 0.7;">by Cloudmate</small>

---

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

[![Go Test](https://github.com/cloudmatelabs/go-jsonld-helper/actions/workflows/gotest.yml/badge.svg)](https://github.com/cloudmatelabs/go-jsonld-helper/actions/workflows/gotest.yml)

</div>

## Why use this library?

When receive JSON-LD,

```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "name": "juunini",
  "type": "Person",
  "id": "juunini"
}
```

is equals

```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "as:name": "juunini",
  "type": "Person",
  "@id": "juunini"
}
```

and it also equals

```json
[
  {
    "https://www.w3.org/ns/activitystreams#name": [
      {
        "@value": "juunini"
      }
    ],
    "@id": "juunini",
    "https://www.w3.org/ns/activitystreams#type": [
      {
        "@value": "Person"
      }
    ]
  }
]
```

when receive any shape of JSON-LD, we can parse and use same interface.

## Installation

```sh
go get github.com/cloudmatelabs/go-jsonld-helper
```

## Usage

```go
import (
  jsonld_helper "github.com/cloudmatelabs/go-jsonld-helper"
)

var doc = map[string]any{
	"@context": []any{
		"https://www.w3.org/ns/activitystreams",
		map[string]any{
			"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
		},
	},
	"@id":     "https://mastodon.social/users/juunini",
	"as:type": "Person",
	"url":     "https://mastodon.social/@juunini",
	"as:image": map[string]any{
		"@type":        "Image",
		"as:mediaType": "image/png",
		"url":          "https://files.mastodon.social/accounts/headers/109/408/471/076/954/889/original/f4158a0d06a05763.png",
	},
	"manuallyApprovesFollowers": "true",
}

jsonld, err := jsonld_helper.ParseJsonLD(doc, nil)
if err != nil {
  // handle error
}

imageType := jsonld.ReadKey("image").ReadKey("mediaType").StringOrElse("")
// image/png
imageURL := jsonld.ReadKey("image").ReadKey("url").StringOrThrow()
// https://files.mastodon.social/accounts/headers/109/408/471/076/954/889/original/f4158a0d06a05763.png
id := jsonld.ReadKey("id").Get()
id := jsonld.ReadKey("@id").Get()
// https://mastodon.social/users/juunini
messageType := jsonld.ReadKey("type").Get()
messageType := jsonld.ReadKey("@type").Get()
// Person
manuallyApprovesFollowers := jsonld.ReadKey("manuallyApprovesFollowers").BooleanOrElse(false)
// true
```

## Methods and Properties

```ts
import "github.com/piprate/json-gold/ld"

jsonld_helper.ParseJsonLD (doc any, options *ld.JsonLdOptions) (JsonLDReader, error)

.Value() any
.Length() int

.ReadKey(string) JsonLDReader
.ReadIndex(int) JsonLDReader

.Get() any
.GetOrElse(any) any
.GetOrThrow(error) (any, error)

.StringOrElse(string) string
.StringOrThrow(error) (string, error)

.BoolOrElse(bool) bool
.BoolOrThrow(error) (bool, error)

.IntOrElse(int) int
.IntOrThrow(error) (int, error)

.FloatOrElse(float64) float64
.FloatOrThrow(error) (float64, error)
```

## License

[MIT](LICENSE)

But, this library use [json-gold].  
[json-gold] is licensed under the [Apache License 2.0](https://github.com/kazarena/json-gold/blob/master/LICENSE)

[json-gold]: https://github.com/kazarena/json-gold
