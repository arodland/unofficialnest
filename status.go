package unofficialnest

import (
    "encoding/json"
    "net/url"
)

type Status struct {
    Devices map[string]Device `json:"device"`
}

type Device struct {
    BatteryLevel float64 `json:"battery_level"`
}

func (nest *NestSession) GetStatus() (status Status, err error) {
    err = nest.RequireLogin()
    if err != nil {
        return
    }

    client := MakeClient()
    req, err := MakeGet(
        nest.TransportURL+"/v2/mobile/"+nest.User,
        url.Values{},
    )
    if err != nil {
        return
    }
    err = nest.Authenticate(req)
    if err != nil {
        return
    }

    res, err := client.Do(req)
    if err != nil {
        return
    }

    err = json.NewDecoder(res.Body).Decode(&status)
    return
}
