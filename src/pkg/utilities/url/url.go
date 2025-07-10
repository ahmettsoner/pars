package url

import (
	"net/url"
	"strings"
)

// EnsureProtocol ekler (http veya https yoksa)
func EnsureProtocol(u string) string {
	if strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://") {
		return u
	}
	return "https://" + u
}

// IsValidURL geçerli URL mi kontrol eder
func IsValidURL(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	return err == nil && parsed.Scheme != "" && parsed.Host != ""
}

// NormalizeURL - trailing slash temizler, lowercase yapar
func NormalizeURL(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	parsed.Host = strings.ToLower(parsed.Host)
	parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	return parsed.String(), nil
}

// ExtractDomain sadece domain bilgisini döner
func ExtractDomain(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return parsed.Hostname(), nil
}

// AddQueryParam - yeni parametre ekler
func AddQueryParam(u, key, value string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	q := parsed.Query()
	q.Add(key, value)
	parsed.RawQuery = q.Encode()
	return parsed.String(), nil
}

// RemoveQueryParam - parametreyi siler
func RemoveQueryParam(u, key string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	q := parsed.Query()
	q.Del(key)
	parsed.RawQuery = q.Encode()
	return parsed.String(), nil
}

// GetQueryParams - map[string][]string olarak query'yi döner
func GetQueryParams(u string) (url.Values, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return parsed.Query(), nil
}

// JoinURL - base ve path birleştirir
func JoinURL(base, path string) (string, error) {
	parsed, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	parsed.Path = strings.TrimSuffix(parsed.Path, "/") + "/" + strings.TrimPrefix(path, "/")
	return parsed.String(), nil
}
