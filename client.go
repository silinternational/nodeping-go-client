package nodeping

import (
	"fmt"
	"gopkg.in/resty.v1"
)

const BaseURL = "https://api.nodeping.com/api/1"
const Version = "0.0.1"

// ClientConfig type includes configuration options for NodePing client
type ClientConfig struct {
	BaseURL    string
	Token      string
	CustomerID string
}

var Client NodePingClient

// NodePingClient holds config and provides methods for various api calls
type NodePingClient struct {
	Config ClientConfig
	Error  NodePingError
	R      *resty.Request
}

// Initialize new NodePingClient
func New(config ClientConfig) (*NodePingClient, error) {
	if config.Token == "" {
		return &NodePingClient{}, fmt.Errorf("token is required in ClientConfig")
	}
	Client.Config.Token = config.Token

	Client.Config.BaseURL = BaseURL
	if config.BaseURL != "" {
		Client.Config.BaseURL = config.BaseURL
	}

	Client.Config.CustomerID = config.CustomerID

	resty.SetHostURL(Client.Config.BaseURL)
	resty.SetBasicAuth(Client.Config.Token, "")
	resty.SetHeader("user-agent", "silinternational/nodeping-go-client "+Version)
	Client.R = resty.R()
	Client.R.SetError(&Client.Error)

	return &Client, nil
}

// ListChecks retrieves all the "Checks" in NodePing
func (c *NodePingClient) ListChecks() ([]CheckResponse, error) {
	path := "/checks"
	if c.Config.CustomerID != "" {
		path = fmt.Sprintf("/checks/%s", c.Config.CustomerID)
	}
	var listObj map[string]CheckResponse
	_, err := c.R.SetResult(&listObj).Get(path)
	errChk := CheckForError(err, c)
	if errChk != nil {
		return []CheckResponse{}, errChk
	}

	var list []CheckResponse
	for _, item := range listObj {
		list = append(list, item)
	}

	return list, nil
}

// GetCheck retrieves data about one Check using its id
func (c *NodePingClient) GetCheck(id string) (CheckResponse, error) {
	path := fmt.Sprintf("/checks/%s", id)
	var check CheckResponse
	_, err := c.R.SetResult(&check).Get(path)
	errChk := CheckForError(err, c)
	if errChk != nil {
		return CheckResponse{}, errChk
	}

	return check, nil
}

// GetUptime retrieves the uptime entries for a certain check within an optional date range (by Timestamp with microseconds)
func (c *NodePingClient) GetUptime(id string, start, end int64) (map[string]UptimeResponse, error) {
	path := fmt.Sprintf("/results/uptime/%s", id)
	queryParams := ""
	queryParamDelimiter := ""

	if start > 0 {
		queryParams = fmt.Sprintf("start=%d", start)
		queryParamDelimiter = "&"
	}

	if end > 0 {
		queryParams = fmt.Sprintf("%s%send=%d", queryParams, queryParamDelimiter, end)
	}

	if len(queryParams) > 0 {
		c.R.SetQueryString(queryParams)
	}

	var listObj map[string]UptimeResponse
	_, err := c.R.SetResult(&listObj).Get(path)
	errChk := CheckForError(err, c)
	if errChk != nil {
		return map[string]UptimeResponse{}, errChk
	}

	return listObj, nil
}

// ListContactGroups retrieves the list of Contact Groups
func (c *NodePingClient) ListContactGroups() (map[string]ContactGroupResponse, error) {
	path := "/contactgroups"
	if c.Config.CustomerID != "" {
		path = fmt.Sprintf("/checks/%s", c.Config.CustomerID)
	}
	var listObj map[string]ContactGroupResponse
	_, err := c.R.SetResult(&listObj).Get(path)
	errChk := CheckForError(err, c)
	if errChk != nil {
		return map[string]ContactGroupResponse{}, errChk
	}

	return listObj, nil
}

func CheckForError(err error, client *NodePingClient) error {
	if err != nil {
		return err
	}
	if client.Error.Error != "" {
		return fmt.Errorf(client.Error.Error)
	}
	return nil
}
