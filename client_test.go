package nodeping_test

import (
	"github.com/silinternational/nodeping-go-client"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// Confirm error when Token is not provided
	_, err := nodeping.New(nodeping.ClientConfig{})
	if err == nil {
		t.Error("Error not thrown when Token was not provided")
	}

	// Test defaults with empty config
	client, err := nodeping.New(nodeping.ClientConfig{Token: "abc123"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, nodeping.BaseURL, client.Config.BaseURL)
	assert.Equal(t, "abc123", client.Config.Token)
}

func TestListChecks(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_, err = client.ListChecks()
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}
}

func TestGetCheck(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}

	check, err := client.GetCheck(checks[0].ID)
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}

	if check.ID != checks[0].ID {
		t.Errorf("Did not get back expected check, got: %+v", check)
	}
}

func TestListContactGroups(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	cgs, err := client.ListContactGroups()
	if err != nil {
		t.Error(err)
	}

	t.Logf("CGs: %+v", cgs)

	if client.Error.Error != "" {
		t.Error(client.Error)
	}
}

func TestGetResultUptime(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}

	uptimes, err := client.GetUptime(checks[0].ID, 0, 0)
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}

	if len(uptimes) < 1 {
		t.Error("Did not get back any uptime data.")
	}

}

func TestGetResultUptimeWithParams(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}

	start := int64(1291161600000)  // Dec 1, 2010
	end := int64(1922313600000)  // Dec 1, 2030
	uptimes, err := client.GetUptime(checks[0].ID, start, end)
	if err != nil {
		t.Error(err)
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}

	if len(uptimes) < 1 {
		t.Error("Did not get back any uptime data.")
	}

}
