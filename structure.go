package unofficialnest

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
