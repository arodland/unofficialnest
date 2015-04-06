// unoficialnest provides an interface to the "unofficial" Nest Labs API for 
// thermostats and Nest Protect devices.  Although Nest provides an official
// API through their developer program, there is information that isn't exposed
// through the official API, including whether the heating or cooling are
// actually on. The unofficial API is the endpoint used by Nest's own web and
// mobile apps to interact with the servers; this code impersonates a mobile
// device to get information and change settings for Nest devices.
package unofficialnest

import (
    "encoding/json"
    "time"
)

// NestSession is the main point of interaction with the API. It is created
// with a set of Credentials and contains state used for interacting with the service.
type NestSession struct {
    serviceURLs
    Credentials
    user          string
    userID        string
    accessToken   string
    accessExpires time.Time
}

// NewSession creates a NestSession from a set of Credentials.
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
