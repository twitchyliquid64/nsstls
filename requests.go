package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var transport *http.Transport
var baseURL string

type user struct {
	UID       int    `json:"uid"`
	Username  string `json:"username"`
	Directory string `json:"directory"`
	Shell     string `json:"shell"`

	GID   int    `json:"group-id"`
	Gecos string `json:"full-name"`
}

type scalarResponse struct {
	User  user   `json:"user"`
	Error string `json:"error"`
}

type scalarRequest struct {
	Username string `json:"username"`
	UID      int    `json:"uid"`
	Token    string `json:"token"`
}

func getUserByName(user, token string) (*scalarResponse, error) {
	url := baseURL + "/userByName"
	if isDebugMode {
		info("API-GETUSERBYNAME", fmt.Sprintf("Making request to %q", url))
	}

	b, err := json.Marshal(scalarRequest{Username: user, Token: token})
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	client := &http.Client{Transport: transport}
	resp, err := client.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if isDebugMode {
		info("API-GETUSERBYNAME", fmt.Sprintf("Response: code=%d(%s),length=%d,content-type=%s", resp.StatusCode, resp.Status, resp.ContentLength, resp.Header.Get("Content-Type")))
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var response scalarResponse
	buf.Reset()
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	return &response, nil
}

func getUserByUID(uid int, token string) (*scalarResponse, error) {
	url := baseURL + "/userByUID"
	if isDebugMode {
		info("API-GETUSERBYUID", fmt.Sprintf("Making request to %q", url))
	}

	b, err := json.Marshal(scalarRequest{UID: uid, Token: token})
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)

	client := &http.Client{Transport: transport}
	resp, err := client.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if isDebugMode {
		info("API-GETUSERBYUID", fmt.Sprintf("Response: code=%d(%s),length=%d,content-type=%s", resp.StatusCode, resp.Status, resp.ContentLength, resp.Header.Get("Content-Type")))
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var response scalarResponse
	buf.Reset()
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	return &response, nil
}
