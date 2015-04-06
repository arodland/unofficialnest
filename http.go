package unofficialnest

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "net/url"
    "strings"
)

const (
    defaultUserAgent = "Nest/3.0.15 (iOS) os=6.0 platform=iPad3,1"
    expiresFormat    = "Mon, 02-Jan-2006 15:04:05 MST"
)

func (nest *NestSession) makeClient() http.Client {
    return http.Client{}
}

func (nest *NestSession) makeRequest(method, host, path string, body io.Reader, authenticated bool) (req *http.Request, err error) {
    if authenticated {
        err = nest.requireLogin()
        if err != nil {
            return
        }
        if host == "" {
            host = nest.TransportURL
        }
    }

    req, err = http.NewRequest(method, host+path, body)
    if err != nil {
        return
    }
    req.Header.Add("User-Agent", defaultUserAgent)
    if authenticated {
        err = nest.authenticate(req)
    }
    return
}

func (nest *NestSession) makePost(host, path string, params interface{}, authenticated bool) (req *http.Request, err error) {
    var body io.Reader
    var ct string
    if params != nil {
        if urlValues, ok := params.(url.Values); ok {
            body = strings.NewReader(urlValues.Encode())
            ct = "application/x-www-form-urlencoded; charset=utf-8"
        } else {
            var js []byte
            js, err = json.Marshal(params)
            if err != nil {
                return
            }
            body = bytes.NewBuffer(js)
            ct = "application/json"
        }
    }
    req, err = nest.makeRequest("POST", host, path, body, authenticated)
    if err != nil {
        return
    }
    req.Header.Add("Content-Type", ct)
    return
}

func (nest *NestSession) makeGet(host, path string, params url.Values, authenticated bool) (req *http.Request, err error) {
    qs := params.Encode()
    if qs != "" {
        path = path + "?" + qs
    }
    req, err = nest.makeRequest("GET", host, path, nil, authenticated)
    return
}

func (nest *NestSession) authenticate(req *http.Request) error {
    err := nest.requireLogin()
    if err != nil {
        return err
    }
    req.Header.Add("X-nl-user-id", nest.userID)
    req.Header.Add("X-nl-protocol-version", "1")
    req.Header.Add("Authorization", "Basic "+nest.accessToken)
    req.Header.Add("Accept-Language", "en")
    return nil
}
