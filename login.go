package unofficialnest

import (
    "encoding/json"
    "net/url"
    "time"
)

type Credentials struct {
    Email    string
    Password string
}

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

func (nest *NestSession) Login() error {
    client := MakeClient()
    req, err := MakePost(
        "https://home.nest.com/user/login",
        url.Values{
            "username": {nest.Email},
            "password": {nest.Password},
        },
    )
    if err != nil {
        return err
    }

    res, err := client.Do(req)
    if err != nil {
        return err
    }

    var lr LoginResponse
    if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
        return err
    }
    nest.ServiceURLs = lr.URLs
    nest.User = lr.User
    nest.UserID = lr.UserID
    nest.AccessToken = lr.AccessToken
    nest.AccessExpires, err = time.Parse(expiresFormat, lr.ExpiresIn)
    if err != nil {
        return err
    }
    return nil
}

func (nest *NestSession) RequireLogin() error {
    if nest.AccessToken != "" && time.Now().Before(nest.AccessExpires) {
        // we have an unexpired access token, do nothing
        return nil
    }
    return nest.Login()
}
