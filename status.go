package unofficialnest

import (
    "encoding/json"
)

// Status contains a report about the devices, structures, and users associated
// with an account.
type Status struct {
    nest      *NestSession
    Device    map[string]Device         `json:"device"`
    Schedule  map[string]DeviceSchedule `json:"schedule"`
    Shared    map[string]DeviceShared   `json:"shared"`
    Structure map[string]Structure      `json:"structure"`
    Where     map[string]StructureWhere `json:"where"`
}

// GetStatus gets a Status from the Nest API and returns it.
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
    status.nest = nest

    for id, dev := range status.Device {
        dev.status = &status
        status.Device[id] = dev
    }
    for id, shared := range status.Shared {
        shared.status = &status
        shared.serialNumber = id
        status.Shared[id] = shared
    }
    for _, where := range status.Where {
        where.populateWhereMap()
    }
    for id, structure := range status.Structure {
        structure.status = &status
        structure.uuid = id
        status.Structure[id] = structure
    }

    return
}
