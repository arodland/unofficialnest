package unofficialnest

import (
    "encoding/json"
    "time"
)

type NestSession struct {
    ServiceURLs
    Credentials
    User          string
    UserID        string
    AccessToken   string
    AccessExpires time.Time
}

func NewSession(creds Credentials) *NestSession {
    return &NestSession{
        Credentials: creds,
    }
}

func (nest *NestSession) GetStatusRaw() (interface{}, error) {
    client := nest.MakeClient()
    req, err := nest.MakeGet("", "/v2/mobile/"+nest.User, nil, true)
    if err != nil {
        return nil, err
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    var out interface{}
    err = json.NewDecoder(res.Body).Decode(&out)
    if err != nil {
        return nil, err
    }
    return out, nil
}
