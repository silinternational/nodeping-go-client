package nodeping_test

import (
	"github.com/silinternational/nodeping-go-client"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// Confirm error when Token is not provided.
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
