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

type loginResponse struct {
    User        string      `json:"user"`
    UserID      string      `json:"user_id"`
    Email       string      `json:"email"`
    AccessToken string      `json:"access_token"`
    ExpiresIn   string      `json:"expires_in"`
    URLs        serviceURLs `json:"urls"`
}

type serviceURLs struct {
    TransportURL       string `json:"transport_url"`
    DirectTransportURL string `json:"direct_transport_url"`
    RubyAPIURL         string `json:"rubyapi_url"`
    WeatherURL         string `json:"weather_url"`
    LogUploadURL       string `json:"log_upload_url"`
    SupportURL         string `json:"support_url"`
}

func (nest *NestSession) Login() error {
    client := nest.makeClient()
    req, err := nest.makePost(
        "https://home.nest.com", "/user/login",
        url.Values{
            "username": {nest.Email},
            "password": {nest.Password},
        },
        false,
    )
    if err != nil {
        return err
    }

    res, err := client.Do(req)
    if err != nil {
        return err
    }

    var lr loginResponse
    if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
        return err
    }
    nest.serviceURLs = lr.URLs
    nest.user = lr.User
    nest.userID = lr.UserID
    nest.accessToken = lr.AccessToken
    nest.accessExpires, err = time.Parse(expiresFormat, lr.ExpiresIn)
    if err != nil {
        return err
    }
    return nil
}

func (nest *NestSession) requireLogin() error {
    if nest.accessToken != "" && time.Now().Before(nest.accessExpires) {
        // we have an unexpired access token, do nothing
        return nil
    }
    return nest.Login()
}

func (nest *NestSession) getUser() (user string, err error) {
    err = nest.requireLogin()
    user = nest.user
    return
}
