package nodeping_test

import (
	"github.com/silinternational/nodeping-go-client"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"strings"
)

var mockCheck = `
{
  "_id":"2018090614528ABCD-MNOPQRST","customer_id":"2018090614528ABCD","label":"Example1","interval":5,
  "notifications":[
    {"AAAA5":
      {"schedule":"All","delay":5}}         
  ], 
  "runlocations":false,"type":"HTTP","status":"assigned","modified":1543260813141,"enable":"active","public":false,"dep":false,
  "parameters":
    {"target":"https://example1.org/","ipv6":false,"follow":false,"threshold":30,"sens":2},
  "created":1539715596937,"queue":"aaaaaaaa10","uuid":"aaaaaaa8-aaa4-aaa4-aaa4-aaaaaaaaaa11","firstdown":0,"state":1
}
`

var mockCheckList = `
{
  "2018090614528ABCD-MNOPQRST":
  {"_id":"2018090614528ABCD-MNOPQRST","customer_id":"2018090614528ABCD","label":"Example1","interval":5,
    "notifications":[
      {"AAAA5":
        {"schedule":"All","delay":5}}         
    ], 
    "runlocations":false,"type":"HTTP","status":"assigned","modified":1543260813141,"enable":"active","public":false,"dep":false,
    "parameters":
      {"target":"https://example1.org/","ipv6":false,"follow":false,"threshold":30,"sens":2},
    "created":1539715596937,"queue":"aaaaaaaa10","uuid":"aaaaaaa8-aaa4-aaa4-aaa4-aaaaaaaaaa11","firstdown":0,"state":1
  },
  "2018090614528ABCD-NOPQRSTU":
  {"_id":"2018090614528ABCD-NOPQRSTU","customer_id":"2018090614528ABCD","label":"Example2","interval":3,
    "notifications":[
      {"2018090614528ABCD-B-BBBB5":
        {"schedule":"All","delay":10}}
    ],
    "runlocations":false,"type":"HTTP","status":"assigned","modified":1543937160541,"enable":"active","public":false,"dep":false,
    "parameters":
      {"target":"https://example2.org/","ipv6":false,"follow":false,"threshold":30,"sens":2},
     "created":1539715552868,"queue":"bbbbbbbb10","uuid":"bbbbbbb8-bbb4-bbb4-bbb4-bbbbbbbbbb11","firstdown":0,"state":1
  },
  "2018090614528ABCD-OPQRSTUV":
  {"_id":"2018090614528ABCD-OPQRSTUV","customer_id":"2018090614528ABCD","label":"Example3","interval":1,
    "notifications":[
      {"2018090614528ABCD-C-CCCCC6":
        {"schedule":"All","delay":5}}
    ],
    "runlocations":false,"type":"HTTP","status":"assigned","modified":1539715504508,"enable":"active","public":false,"dep":false,
    "parameters":
      {"target":"https://example3.org/home","ipv6":false,"follow":false,"threshold":30,"sens":2},
    "created":1539715504508,"queue":"cccccccc10","uuid":"ccccccc8-ccc4-ccc4-ccc4-cccccccccc11","firstdown":0,"state":1
  },
  "2018090614528ABCD-PQRSTUVW":
  {"_id":"2018090614528ABCD-PQRSTUVW","customer_id":"2018090614528ABCD","label":"Example4","interval":1,
    "notifications":[
      {"2018090614528ABCD-D-DDDD5":
        {"schedule":"All","delay":5}},
      {"EEEE5":
        {"schedule":"All","delay":5}}
    ],
    "runlocations":false,"type":"HTTP","status":"assigned","modified":1543260719724,"enable":"active","public":false,"dep":false,
    "parameters":
       {"target":"https://example4.org/check","ipv6":false,"follow":false,"threshold":30,"sens":2},
    "created":1539715451787,"queue":"dddddddd10","uuid":"ddddddd8-ddd4-ddd4-ddd4-dddddddddd11","firstdown":0,"state":1
  }
}
`

var mockContactGroupList = `
{
 "2018090614528ABCD-A-AAAA5":
    {"type":"group","customer_id":"2018090614528ABCD","name":"CGList1","members":["AAAA5","BBBB5","CCCC5","DDDD5","EEEE5"]},
 "2018090614528Abcd-B-BBBB5":
    {"type":"group","customer_id":"2018090614528ABCD","name":"CGList2","members":["FFFF5"]}
}
`

var mockUptime = `
{
  "2018-10":{"enabled":1315092551,"down":82790,"uptime":99.010},
  "2018-11":{"enabled":2592000000,"down":89391,"uptime":99.011},
  "2018-12":{"enabled":837810368,"down":80892,"uptime":99.012},
  "total":{"enabled":4744902919,"down":253073,"uptime":99.011}
}
`

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
	assert.Equal(t, "", client.MockResults)
}

func TestListChecks(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.ListChecks()
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
	}
}


func TestListChecksMock(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})

	if err != nil {
		t.Error(err)
		return
	}

	client.MockResults = mockCheckList

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	checkCount := len(checks)
	expectedCount := 4
	if checkCount != expectedCount {
		t.Errorf("Got wrong number of checks in list.  Expected %d, but got %d", expectedCount, checkCount)
		return
	}

	results := checks[0].Label
	expected := "Example"
	if ! strings.HasPrefix(results, expected) {
		t.Errorf("Got wrong Label on check.  Expected it to start with %s, but got %s", expected, results)
		return
	}
}

func TestGetCheck(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	check, err := client.GetCheck(checks[0].ID)
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	if check.ID != checks[0].ID {
		t.Errorf("Did not get back expected check, got: %+v", check)
	}
}


func TestGetCheckMock(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	client.MockResults = mockCheck
    id := "2018090614528ABCD-MNOPQRST"
	check, err := client.GetCheck(id)
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	expected := "Example1"
	if check.Label != expected {
		t.Errorf("Did not get back expected check. Expected it to have label: %s. \n  But got: \n%+v", expected, check)
	}
}


func TestListContactGroups(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	cgs, err := client.ListContactGroups()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("CGs: %+v", cgs)

	if client.Error.Error != "" {
		t.Error(client.Error)
	}
}

func TestListContactGroupsMock(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	client.MockResults = mockContactGroupList

	cgs, err := client.ListContactGroups()
	if err != nil {
		t.Error(err)
		return
	}

	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	groupsCount := len(cgs)
	expectedCount := 2
	if groupsCount != expectedCount {
		t.Errorf("Got wrong number of contact groups. Expected %d, but got %d", expectedCount, groupsCount)
		return
	}
	results := cgs["2018090614528ABCD-A-AAAA5"]
	expected := "CGList1"
	if results.Name != expected {
		t.Errorf("Did not get back expected contact group. Expected it to have name: %s. \n  But got: \n%+v", expected, results)
	}
}

func TestGetResultUptime(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	uptimes, err := client.GetUptime(checks[0].ID, 0, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	if len(uptimes) < 1 {
		t.Error("Did not get back any uptime data.")
	}

}

func TestGetResultUptimeWithParams(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	checks, err := client.ListChecks()
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	start := int64(1291161600000)  // Dec 1, 2010
	end := int64(1922313600000)  // Dec 1, 2030
	uptimes, err := client.GetUptime(checks[0].ID, start, end)
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	if len(uptimes) < 1 {
		t.Error("Did not get back any uptime data.")
	}

}


func TestGetResultUptimeMock(t *testing.T) {
	client, err := nodeping.New(nodeping.ClientConfig{Token: os.Getenv("NODEPING_TOKEN")})
	if err != nil {
		t.Error(err)
		return
	}

	client.MockResults = mockUptime

	uptimes, err := client.GetUptime("2018090614528ABCD", 0, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if client.Error.Error != "" {
		t.Error(client.Error)
		return
	}

	uptimeCount := len(uptimes)
	expectedCount := 4

	if uptimeCount != expectedCount {
		t.Errorf("Got wrong number of uptime data. Expected %d, but got %d", expectedCount, uptimeCount)
		return
	}

	results := uptimes["total"]
    expected := int64(253073)

    if results.Down != expected {
    	t.Errorf("Got wrong total down value. Expected %d, but got %d", expected, results.Down)
	}

}
