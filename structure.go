package unofficialnest

import (
    "fmt"
    "time"
)

// Structure is the building containing one or more Nest devices. It
// contains information about the building itself, and the devices within it.
type Structure struct {
    uuid      string
    status    *Status
    Timestamp int64 `json:"$timestamp"`
    Version   int   `json:"$version"`

    Away    bool     `json:"away"`
    Devices []string `json:"devices"`
}

type StructureWhere struct {
    Timestamp int64   `json:"$timestamp"`
    Version   int     `json:"$version"`
    Wheres    []Where `json:"wheres"`
    WhereMap  map[string]*Where
}

type Where struct {
    Name    string `json:"name"`
    WhereID string `json:"where_id"`
}

func (s *StructureWhere) populateWhereMap() {
    s.WhereMap = make(map[string]*Where)
    for i, where := range s.Wheres {
        s.WhereMap[where.WhereID] = &s.Wheres[i]
    }
}

func (structure *Structure) Update(payload interface{}) error {
    id := structure.uuid
    nest := structure.status.nest
    client := nest.makeClient()
    req, err := nest.makePost("", "/v2/put/structure." + id, payload, true)
    if err != nil {
        return err
    }
    req.Header.Add("X-nl-base-version", fmt.Sprintf("%d", structure.Version))
    _, err = client.Do(req)
    return err
}

func (structure *Structure) SetAway(away bool) error {
    return structure.Update(map[string]interface{}{
        "away": away,
        "away_setter": 0,
        "away_timestamp": time.Now().UTC().Unix() * 1000,
    })
}
