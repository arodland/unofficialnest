package unofficialnest

import (
    "time"
)

type Device struct {
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

type DeviceShared struct {
    Timestamp int64 `json:"$timestamp"`
    Version   int   `json:"$version"`

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

func (device *Device) GetTimestamp() time.Time {
    return time.Unix(
        device.Timestamp/1000,
        (device.Timestamp%1000)*1000000,
    )
}
