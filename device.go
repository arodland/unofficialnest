package unofficialnest

import (
    "fmt"
    "time"
)

// Device contains information about a Nest device (Thermostat or Protect).
// If you can't find what you're looking for in here, it's probably in DeviceShared
// instead; find it by looking in status.Shared with this device's SerialNumber.
type Device struct {
    status       *Status
    Timestamp    int64  `json:"$timestamp"`
    Version      int    `json:"$version"`
    SerialNumber string `json:"serial_number"`

    BatteryLevel    float64 `json:"battery_level"`
    CurrentHumidity float64 `json:"current_humidity"`
    ErrorCode       string  `json:"error_code"`
    FanMode         string  `json:"fan_mode"`
    HasAltHeat      bool    `json:"has_alt_heat"`
    HasEmerHeat     bool    `json:"has_emer_heat"`
    HasFan          bool    `json:"has_fan"`
    Leaf            bool    `json:"leaf"`
    RSSI            float64 `json:"rssi"`
    WhereID         string  `json:"where_id"`
}

type DeviceSchedule struct {
    status    *Status
    TimeStamp int64 `json:"$timestamp"`
    Version   int   `json:"$version"`
    Days      map[string]ScheduleDay
}

type ScheduleDay map[string]ScheduleEntry

type ScheduleEntry struct {
    EntryType string  `json:"entry_type"`
    Type      string  `json:"type"`
    Time      int     `json:"time"`
    temp      float64 `json:"temp"`
}

// DeviceShared contains information about the state of a Nest device (Thermostat
// or Protect), including the temperature and whether it is heating, cooling,
// blowing the fan, etc.
type DeviceShared struct {
    status       *Status
    serialNumber string
    Timestamp    int64 `json:"$timestamp"`
    Version      int   `json:"$version"`

    AutoAway int `json:"auto_away"`

    CanCool bool `json:"can_cool"`
    CanHeat bool `json:"can_heat"`

    CurrentTemperature float64 `json:"current_temperature"`
    TargetTemperature  float64 `json:"target_temperature"`

    ACState     bool `json:"hvac_ac_state"`
    CoolX2State bool `json:"hvac_cool_x2_state"`

    FanState bool `json:"hvac_fan_state"`

    AltHeatState   bool `json:"hvac_alt_heat_state"`
    AltHeatX2State bool `json:"hvac_alt_heat_x2_state"`

    EmerHeatState bool `json:"hvac_emer_heat_state"`

    HeaterState bool `json:"hvac_heater_state"`
    HeatX2State bool `json:"hvac_heat_x2_state"`
    HeatX3State bool `json:"hvac_heat_x3_state"`
}

// GetTimestamp converts the device's $timestamp field into a go time.Time.
func (device *Device) GetTimestamp() time.Time {
    return time.Unix(
        device.Timestamp/1000,
        (device.Timestamp%1000)*1000000,
    )
}

func (device *Device) Update(payload interface{}) error {
    id := device.SerialNumber
    nest := device.status.nest
    client := nest.makeClient()
    req, err := nest.makePost("", "/v2/put/device." + id, payload, true)
    if err != nil {
        return err
    }
    req.Header.Add("X-nl-base-version", fmt.Sprintf("%d",device.Version))
    _, err = client.Do(req)
    return err
}

func (shared *DeviceShared) Update(payload interface{}) error {
    id := shared.serialNumber
    nest := shared.status.nest
    client := nest.makeClient()
    req, err := nest.makePost("", "/v2/put/shared." + id, payload, true)
    if err != nil {
        return err
    }
    req.Header.Add("X-nl-base-version", fmt.Sprintf("%d",shared.Version))
    _, err = client.Do(req)
    return err
}

func (shared *DeviceShared) SetTargetTemperature(tempC float64) error {
    return shared.Update(map[string]interface{}{
        "target_change_pending": true,
        "target_temperature": tempC,
    })
}

func (device *Device) SetFanMode(fanMode string) error {
    return device.Update(map[string]interface{}{
        "fan_mode": fanMode,
    })
}
