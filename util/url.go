package util

import (
	"errors"
	"golang.org/x/exp/slices"
	"net/url"
	"sort"
	"strings"
)

func UrlSign(urlString string, secret string) (string, error) {
	u, err := url.Parse(urlString)
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}

	keys := make([]string, 0, len(q))
	for k, _ := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	queryStrings := make([]string, 0, len(keys))
	for _, v := range keys {
		queryStrings = append(queryStrings, v+"="+q.Get(v))
	}

	sortedQueryString := strings.Join(queryStrings, "&")
	sign := HmacSha256(sortedQueryString, secret)

	return sign, nil
}

func UrlSignCheck(urlString string, secret string, exclude []string) (bool, error) {
	u, err := url.Parse(urlString)
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return false, err
	}

	signOrigin := q.Get("sign")
	if signOrigin == "" {
		return false, errors.New("sign is empty")
	}

	exclude = append(exclude, "sign")

	keys := make([]string, 0, len(q))
	for k, _ := range q {
		//exclude项不参与签名
		if !slices.Contains(exclude, k) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	queryStrings := make([]string, 0, len(keys))
	for _, v := range keys {
		queryStrings = append(queryStrings, v+"="+q.Get(v))
	}
	sortedQueryString := strings.Join(queryStrings, "&")

	return CheckMAC(sortedQueryString, signOrigin, secret), nil
}

// UrlAddQueryString 添加QueryString到url
func UrlAddQueryString(urlOrigin string, queryMap map[string]interface{}) string {
	u, err := url.Parse(urlOrigin)
	if err != nil {
		return urlOrigin
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return urlOrigin
	}

	for k, v := range queryMap {
		q.Set(k, v.(string))
	}

	u.RawQuery = q.Encode()
	return u.String()
}

// IsUrl 是否是url
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
