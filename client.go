package nodeping

import (
	"fmt"
	"gopkg.in/resty.v1"
)

const BaseURL = "https://api.nodeping.com/api/1"
const Version = "0.0.1"

// ClientConfig type includes configuration options for NodePing client.
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

func CheckForError(err error, client *NodePingClient) error {
	if err != nil {
		return err
	}
	if client.Error.Error != "" {
		return fmt.Errorf(client.Error.Error)
	}
	return nil
}
