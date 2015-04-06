package unofficialnest

import (
    "encoding/json"
)

type Status struct {
    Device    map[string]Device         `json:"device"`
    Schedule  map[string]DeviceSchedule `json:"schedule"`
    Shared    map[string]DeviceShared   `json:"shared"`
    Structure map[string]Structure      `json:"structure"`
    Where     map[string]StructureWhere `json:"where"`
}

func (nest *NestSession) GetStatus() (status Status, err error) {
    user, err := nest.getUser()
    if err != nil {
        return
    }
    client := nest.makeClient()
    req, err := nest.makeGet("", "/v2/mobile/"+user, nil, true)
    if err != nil {
        return
    }
    res, err := client.Do(req)
    if err != nil {
        return
    }

    err = json.NewDecoder(res.Body).Decode(&status)

    for _, sw := range status.Where {
        sw.populateWhereMap()
    }

    return
}
