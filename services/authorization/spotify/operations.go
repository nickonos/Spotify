package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (op Operations) GetAccessToken(ctx context.Context, code string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "POST", op.auth_url+TOKEN_ENDPOINT, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", op.client_id, op.client_secret))))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Form.Add("code", code)
	req.Form.Add("grant_type", "authorization_code")
	req.Form.Add("redirect_uri", "/")

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
