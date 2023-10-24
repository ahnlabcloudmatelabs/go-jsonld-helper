package jsonld_helper_test

import (
	"testing"

	jsonld_helper "github.com/cloudmatelabs/go-jsonld-helper"
)

func Test_Get_WhenGivenKeyCanNotFind_ShouldReturnNil(t *testing.T) {
	givenKey := "notFoundKey"

	if result := jsonld.ReadKey(givenKey).Get(); result != nil {
		t.Fatal("should return nil:", result)
	}

	if result := jsonld.ReadIndex(0).ReadKey(givenKey).Get(); result != nil {
		t.Fatal("should return nil:", result)
	}
}

func Test_Get_WhenGivenIndexResultIsObject_ShouldReturnNil(t *testing.T) {
	if result := jsonld.ReadIndex(0).ReadIndex(0).Get(); result != nil {
		t.Fatal("should return nil:", result)
	}
}

func Test_Get_WhenGivenKeyCanFind_ShouldReturnValue(t *testing.T) {
	givenKey := "name"

	if result := jsonld.ReadKey(givenKey).Get(); result != "지상 최강의 개발자 쥬니니" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadIndex(0).ReadKey(givenKey).Get(); result != "지상 최강의 개발자 쥬니니" {
		t.Fatal("failed", result)
	}
}

func Test_Get_WhenGivenKeyIsPreDefinedKey_ShouldReturnValue(t *testing.T) {
	if result := jsonld.ReadKey("id").Get(); result != "acct:juunini@snippet.cloudmt.co.kr" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("@id").Get(); result != "acct:juunini@snippet.cloudmt.co.kr" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("type").Get(); result != "Person" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("@type").Get(); result != "Person" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadIndex(0).ReadKey("id").Get(); result != "acct:juunini@snippet.cloudmt.co.kr" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadIndex(0).ReadKey("@id").Get(); result != "acct:juunini@snippet.cloudmt.co.kr" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadIndex(0).ReadKey("type").Get(); result != "Person" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadIndex(0).ReadKey("@type").Get(); result != "Person" {
		t.Fatal("failed", result)
	}
}

func Test_Get_WhenGivenIndexAndKeyCanFind_ShouldReturnValue(t *testing.T) {
	if result := jsonld.ReadKey("attachment").ReadIndex(0).ReadKey("value").Get(); result != "juunini" {
		t.Fatal("failed", result)
	}
}

func Test_GetOrElse_WhenGivenKeyCanNotFind_ShouldReturnDefaultValue(t *testing.T) {
	givenKey := "notFoundKey"
	givenDefaultValue := "defaultValue"

	if result := jsonld.ReadKey(givenKey).GetOrElse(givenDefaultValue); result != givenDefaultValue {
		t.Fatal("failed", result)
	}

	if result := jsonld.ReadIndex(0).ReadKey(givenKey).GetOrElse(givenDefaultValue); result != givenDefaultValue {
		t.Fatal("failed", result)
	}
}

func Test_GetOrElse_WhenGivenKeyCanFind_ShouldReturnValue(t *testing.T) {
	givenKey := "name"
	givenDefaultValue := "defaultValue"

	if result := jsonld.ReadKey(givenKey).GetOrElse("지상 최강의 개발자 쥬니니"); result == givenDefaultValue {
		t.Fatal("failed", result)
	}

	if result := jsonld.ReadIndex(0).ReadKey(givenKey).GetOrElse("지상 최강의 개발자 쥬니니"); result == givenDefaultValue {
		t.Fatal("failed", result)
	}
}

func Test_GetOrElse_WhenGivenIndexCanNotFind_ShouldReturnDefaultValue(t *testing.T) {
	givenIndex := 1
	givenDefaultValue := "defaultValue"

	if result := jsonld.ReadIndex(givenIndex).GetOrElse(givenDefaultValue); result != givenDefaultValue {
		t.Fatal("failed", result)
	}

	if result := jsonld.ReadIndex(0).ReadIndex(0).GetOrElse(givenDefaultValue); result != givenDefaultValue {
		t.Fatal("failed", result)
	}
}

func Test_GetOrThrow_WhenGivenKeyCanNotFind_ShouldThrowError(t *testing.T) {
	givenKey := "notFoundKey"

	if result, err := jsonld.ReadKey(givenKey).GetOrThrow(nil); err == nil || err.Error() != "not found key: notFoundKey" {
		t.Fatal("failed", result, err)
	}

	if result, err := jsonld.ReadIndex(0).ReadKey(givenKey).GetOrThrow(nil); err == nil || err.Error() != "not found key: notFoundKey" {
		t.Fatal("failed", result, err)
	}
}

func Test_GetOrThrow_WhenGivenKeyCanFind_ShouldReturnValue(t *testing.T) {
	givenKey := "name"

	if result, err := jsonld.ReadKey(givenKey).GetOrThrow(nil); err != nil || result != "지상 최강의 개발자 쥬니니" {
		t.Fatal("failed", result, err)
	}

	if result, err := jsonld.ReadIndex(0).ReadKey(givenKey).GetOrThrow(nil); err != nil || result != "지상 최강의 개발자 쥬니니" {
		t.Fatal("failed", result, err)
	}
}

func Test_READMEJsonLDTest(t *testing.T) {
	jsonld, _ := jsonld_helper.ParseJsonLD(readmeJsonLD, nil)

	if result := jsonld.ReadKey("id").Get(); result != "https://mastodon.social/users/juunini" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("@id").Get(); result != "https://mastodon.social/users/juunini" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("@id").StringOrElse(""); result != "https://mastodon.social/users/juunini" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("type").Get(); result != "Person" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("@type").Get(); result != "Person" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("@type").StringOrElse(""); result != "Person" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("url").Get(); result != "https://mastodon.social/@juunini" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("url").StringOrElse(""); result != "https://mastodon.social/@juunini" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("type").Get(); result != "Image" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("@type").Get(); result != "Image" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("@type").StringOrElse(""); result != "Image" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("mediaType").Get(); result != "image/png" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("mediaType").StringOrElse(""); result != "image/png" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("url").Get(); result != "https://files.mastodon.social/accounts/headers/109/408/471/076/954/889/original/f4158a0d06a05763.png" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("image").ReadKey("url").StringOrElse(""); result != "https://files.mastodon.social/accounts/headers/109/408/471/076/954/889/original/f4158a0d06a05763.png" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("manuallyApprovesFollowers").Get(); result != "true" {
		t.Fatal("failed", result)
	}
	if result := jsonld.ReadKey("manuallyApprovesFollowers").BoolOrElse(false); result != true {
		t.Fatal("failed", result)
	}
}

var jsonld, _ = jsonld_helper.ParseJsonLD(givenJsonLD, nil)
var givenJsonLD = map[string]any{
	"@context": []any{
		"activitystreams.json",
		map[string]any{
			"schema":        "http://schema.org#",
			"PropertyValue": "schema:PropertyValue",
			"value":         "schema:value",
		},
	},
	"as:type": "Person",
	"@id":     "acct:juunini@snippet.cloudmt.co.kr",
	"name":    "지상 최강의 개발자 쥬니니",
	"attachment": []map[string]any{
		{
			"type":  "PropertyValue",
			"name":  "GitHub",
			"value": "juunini",
		},
	},
}
var readmeJsonLD = map[string]any{
	"@context": []any{
		"activitystreams.json",
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
