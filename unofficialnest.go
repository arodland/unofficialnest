package unofficialnest

import (
    "encoding/json"
    "time"
)

type NestSession struct {
    serviceURLs
    Credentials
    user          string
    userID        string
    accessToken   string
    accessExpires time.Time
}

func NewSession(creds Credentials) *NestSession {
    return &NestSession{
        Credentials: creds,
    }
}

func (nest *NestSession) GetStatusRaw() (interface{}, error) {
    client := nest.makeClient()
    req, err := nest.makeGet("", "/v2/mobile/"+nest.user, nil, true)
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
