package unofficialnest

import (
    "encoding/json"
    "net/url"
    "time"
)

type LoginResponse struct {
    User        string      `json:"user"`
    UserID      string      `json:"user_id"`
    Email       string      `json:"email"`
    AccessToken string      `json:"access_token"`
    ExpiresIn   string      `json:"expires_in"`
    URLs        ServiceURLs `json:"urls"`
}

type ServiceURLs struct {
    TransportURL       string `json:"transport_url"`
    DirectTransportURL string `json:"direct_transport_url"`
    RubyAPIURL         string `json:"rubyapi_url"`
    WeatherURL         string `json:"weather_url"`
    LogUploadURL       string `json:"log_upload_url"`
    SupportURL         string `json:"support_url"`
}

func (nest *NestSession) Login(username string, password string) (*LoginResponse, error) {
    client := MakeClient()
    req, err := MakePost(
        "https://home.nest.com/user/login",
        url.Values{
            "username": {username},
            "password": {password},
        },
    )
    if err != nil {
        return nil, err
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    var lr LoginResponse
    if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
        return nil, err
    }
    nest.ServiceURLs = lr.URLs
    nest.User = lr.User
    nest.UserID = lr.UserID
    nest.AccessToken = lr.AccessToken
    nest.AccessExpires, err = time.Parse(expiresFormat, lr.ExpiresIn)
    if err != nil {
        return nil, err
    }
    return &lr, nil
}
