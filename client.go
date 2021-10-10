package pride

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	c      *http.Client
	apiKey string
}

type Config struct {
	c      *http.Client
	apiKey string
}

type HTTPError struct {
	Status  int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Message)
}

func New(conf *Config) *Client {
	var httpClient *http.Client
	if conf.c == nil {
		httpClient = &http.Client{Timeout: time.Second * 30}
	} else {
		httpClient = conf.c
	}
	return &Client{c: httpClient, apiKey: conf.apiKey}
}

func (c *Client) do(req *http.Request) (*response, error) {
	if c.apiKey != "" {
		req.Header.Set("X-API-Key", c.apiKey)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r response

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return nil, err
	}

	if r.Code != 200 {
		return nil, &HTTPError{r.Code, r.Msg}
	}

	return &r, nil
}

func (c *Client) doBinary(req *http.Request) ([]byte, error) {
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &HTTPError{resp.StatusCode, resp.Status}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetFlags() (map[string]Flag, error) {
	req, err := http.NewRequest("GET", GetFlagsURL, nil)
	if err != nil {
		return nil, err
	}

	r, err := c.do(req)
	if err != nil {
		return nil, err
	}

	flags := make(map[string]Flag)
	err = json.Unmarshal(r.Data, &flags)
	if err != nil {
		return nil, err
	}

	return flags, nil
}

func (c *Client) GetFlag(name string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(GetFlagsByNameURL, name), nil)
	if err != nil {
		return nil, err
	}

	return c.doBinary(req)
}

func (c *Client) PutImage(r io.Reader, contentType string) (*ImageData, error) {
	req, err := http.NewRequest("PUT", PutImageURL, r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := c.do(req)

	var data ImageData
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return nil, err
	}

	return &data, err
}

func (c *Client) EditImage(id string, params *EditImageParams) (*EditImageResponse, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(EditImageURL, id), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var r EditImageResponse
	if err := json.Unmarshal(resp.Data, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) SearchFlags(query string, limit int) (map[string]Flag, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(SearchFlagsURL, query, limit), nil)
	if err != nil {
		return nil, err
	}

	r, err := c.do(req)
	if err != nil {
		return nil, err
	}

	flags := make(map[string]Flag)
	err = json.Unmarshal(r.Data, &flags)
	if err != nil {
		return nil, err
	}

	return flags, nil
}
