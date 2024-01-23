package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Operations struct {
	url           string
	auth_url      string
	client_id     string
	client_secret string
}

const TOKEN_ENDPOINT = "/api/token"

const USER_ENDPOINT = "/me"

func NewOperations() (Operations, error) {
	url := os.Getenv("SPOTIFY_URL")
	auth_url := os.Getenv("SPOTIFY_AUTH_URL")
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if url == "" || client_id == "" || client_secret == "" || auth_url == "" {
		return Operations{}, errors.New("missing environement variables")
	}

	return Operations{
		url:           url,
		auth_url:      auth_url,
		client_id:     client_id,
		client_secret: client_secret,
	}, nil
}

type GetAccessTokenResponse struct {
	AccesToken   string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type GetAccessTokenBody struct {
	Code        string `json:"code"`
	GrantType   string `json:"grant_type"`
	RedirectURI string `json:"redirect_uri"`
}

func (op Operations) GetAccessToken(ctx context.Context, code string) (string, error) {
	client := &http.Client{}

	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:5173/auth/callback")
	// b, err := json.Marshal(GetAccessTokenBody{
	// 	Code:        code,
	// 	GrantType:   "authorization_code",
	// 	RedirectURI: "http://localhost:5173/auth/callback",
	// })
	// if err != nil {
	// 	return "", err
	// }

	req, err := http.NewRequestWithContext(ctx, "POST", op.auth_url+TOKEN_ENDPOINT, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", op.client_id, op.client_secret))))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", errors.New("unexpected response code " + res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var body GetAccessTokenResponse
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return "", nil
	}

	return body.AccesToken, nil
}

type GetUserResponse struct {
	Email string `json:"email"`
}

func (op Operations) GetUserEmail(ctx context.Context, auth_token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", op.url+USER_ENDPOINT, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+auth_token)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", errors.New("unexpected response code " + res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var body GetUserResponse
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return "", nil
	}

	return body.Email, nil
}
