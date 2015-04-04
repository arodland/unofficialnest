package unofficialnest

import (
    "io"
    "net/http"
    "net/url"
    "strings"
)

const (
    defaultUserAgent = "Nest/3.0.15 (iOS) os=6.0 platform=iPad3,1"
    expiresFormat    = "Mon, 02-Jan-2006 15:04:05 MST"
)

func MakeClient() http.Client {
    return http.Client{}
}

func MakeRequest(method string, uri string, body io.Reader) (req *http.Request, err error) {
    req, err = http.NewRequest(method, uri, body)
    if err != nil {
        return
    }
    req.Header.Add("User-Agent", defaultUserAgent)
    return
}

func MakePost(uri string, params url.Values) (req *http.Request, err error) {
    body := strings.NewReader(params.Encode())
    req, err = MakeRequest("POST", uri, body)
    if err != nil {
        return
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
    return
}

func MakeGet(uri string, params url.Values) (req *http.Request, err error) {
    qs := params.Encode()
    if qs != "" {
        uri = uri + "?" + qs
    }
    req, err = MakeRequest("GET", uri, nil)
    return
}

func (nest *NestSession) Authenticate(req *http.Request) error {
    err := nest.RequireLogin()
    if err != nil {
        return err
    }
    req.Header.Add("X-nl-user-id", nest.UserID)
    req.Header.Add("X-nl-protocol-version", "1")
    req.Header.Add("Authorization", "Basic "+nest.AccessToken)
    req.Header.Add("Accept-Language", "en")
    return nil
}
