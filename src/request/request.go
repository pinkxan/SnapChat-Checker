package request

import (
	"errors"
	"fmt"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

type RequestsClient struct {
	client tls_client.HttpClient
}

func Client(proxyUrl string) RequestsClient {
	options := []tls_client.HttpClientOption{
		tls_client.WithClientProfile(tls_client.NikeIosMobile),
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithProxyUrl(proxyUrl),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		panic(err)
	}

	return RequestsClient{
		client: client,
	}
}

func (c *RequestsClient) SendRequest(method string, url string, body io.Reader, headers http.Header) (*http.Response, string) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	request.Header = headers

	response, err := c.client.Do(request)
	if err != nil {
		if strings.Contains(err.Error(), "dial tcp: lookup gcp.api.snapchat.com: no such host") {
			return c.SendRequest(method, url, body, headers)
		}

		fmt.Println(err.Error())
		panic(errors.New("no worky"))
	}
	defer response.Body.Close()

	readBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return response, string(readBytes)
}
