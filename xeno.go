package xeno

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	APIEndpointBase = "https://www.xeno-canto.org/api/2/recordings?query="
)

type Response struct {
	NumRecordings string       `json:"numRecordings,omitempty"`
	NumSpecies    string       `json:"numSpecies,omitempty"`
	Page          int          `json:"page,omitempty"`
	NumPages      int          `json:"numPages,omitempty"`
	Recordings    []Recordings `json:"recordings,omitempty"`
}

type Recordings struct {
	Id           string   `json:"id,omitempty"`
	Gen          string   `json:"gen,omitempty"`
	Sp           string   `json:"sp,omitempty"`
	Ssp          string   `json:"ssp,omitempty"`
	Group        string   `json:"group,omitempty"`
	En           string   `json:"en,omitempty"`
	Rec          string   `json:"rec,omitempty"`
	Cnt          string   `json:"cnt,omitempty"`
	Loc          string   `json:"loc,omitempty"`
	Lat          string   `json:"lat,omitempty"`
	Lng          string   `json:"lng,omitempty"`
	Alt          string   `json:"alt,omitempty"`
	Type         string   `json:"type,omitempty"`
	Sex          string   `json:"sex,omitempty"`
	Stage        string   `json:"stage,omitempty"`
	Method       string   `json:"method,omitempty"`
	Url          string   `json:"url,omitempty"`
	File         string   `json:"file,omitempty"`
	FileName     string   `json:"file-name,omitempty"`
	Sono         Sono     `json:"sono,omitempty"`
	Osci         Osci     `json:"osci,omitempty"`
	Lic          string   `json:"lic,omitempty"`
	Q            string   `json:"q,omitempty"`
	Length       string   `json:"length,omitempty"`
	Time         string   `json:"time,omitempty"`
	Date         string   `json:"date,omitempty"`
	Uploaded     string   `json:"uploaded,omitempty"`
	Also         []string `json:"also,omitempty"`
	Rmk          string   `json:"rmk,omitempty"`
	BirdSeen     string   `json:"bird-seen,omitempty"`
	AnimalSeen   string   `json:"animal-seen,omitempty"`
	PlaybackUsed string   `json:"playback-used,omitempty"`
	Temp         string   `json:"temp,omitempty"`
	Regnr        string   `json:"regnr,omitempty"`
	Auto         string   `json:"auto,omitempty"`
	Dvc          string   `json:"dvc,omitempty"`
	Mic          string   `json:"mic,omitempty"`
	Smp          string   `json:"smp,omitempty"`
}

type Sono struct {
	Small string `json:"small,omitempty"`
	Med   string `json:"med,omitempty"`
	Large string `json:"large,omitempty"`
	Full  string `json:"full,omitempty"`
}

type Osci struct {
	Small string `json:"small,omitempty"`
	Med   string `json:"med,omitempty"`
	Large string `json:"large,omitempty"`
}

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

type ClientOption func(client *Client)

func WithBaseURL(urlStr string) ClientOption {
	return func(client *Client) {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}
		client.baseURL = parsedURL
	}
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e Error) Error() string {
	return e.Message
}

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
	}

	if c.baseURL == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}
		c.baseURL = u
	}

	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) get(ctx context.Context, query string, result interface{}) error {
	xenoURL := c.baseURL.String() + query
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, xenoURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e struct {
		E Error `json:"error"`
	}
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	if e.E.Message == "" {
		e.E.Message = fmt.Sprintf("unexpected HTTP %d: %s (empty error)",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return e.E
}

func (c *Client) Get(ctx context.Context, query string, opts ...RequestOption) (*Response, error) {
	query = url.QueryEscape(query)
	var r Response
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		query += "&" + params
	}
	err := c.get(ctx, query, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
