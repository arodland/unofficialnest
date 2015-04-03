package unofficialnest

import (
    "encoding/json"
    "net/url"
    "time"
)

type NestSession struct {
    ServiceURLs
    User          string
    UserID        string
    AccessToken   string
    AccessExpires time.Time
}

func NewSession() *NestSession {
    return &NestSession{}
}

func (nest *NestSession) GetStatus() (interface{}, error) {
    client := MakeClient()
    req, err := MakeGet(
        nest.TransportURL+"/v2/mobile/"+nest.User,
        url.Values{},
    )
    if err != nil {
        return nil, err
    }
    nest.Authenticate(req)

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
