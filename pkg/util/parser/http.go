package parser

import "net/url"

func ParseURL(rawurl string) (string, error) {
	reqUrl, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return reqUrl.String(), nil
}
