package nodeping

//
type NodePingError struct {
	Error string `json:"error"`
}

type CheckNotification struct {
	Contactkey1 struct {
		Delay    int    `json:"delay"`
		Schedule string `json:"schedule"`
	} `json:"contactkey1"`
}

type CheckRequest struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Target        string              `json:"target"`
	Label         string              `json:"label"`
	Interval      int                 `json:"interval"`
	Enabled       string              `json:"enabled"`
	Public        string              `json:"public"`
	Runlocations  []string            `json:"runlocations"`
	Homeloc       bool                `json:"homeloc"`
	Threshold     int                 `json:"threshold"`
	Sens          int                 `json:"sens"`
	Notifications []CheckNotification `json:"notifications"`
	Dep           string              `json:"dep"`
	Contentstring string              `json:"contentstring"`
	Follow        bool                `json:"follow"`
	Data          string              `json:"data"`
	Method        string              `json:"method"`
	Statuscode    string              `json:"statuscode"`
	Ipv6          bool                `json:"ipv6"`
}

type CheckResponse struct {
	ID            string        `json:"_id"`
	Rev           string        `json:"_rev"`
	CustomerID    string        `json:"customer_id"`
	Label         string        `json:"label"`
	Interval      int           `json:"interval"`
	Notifications []interface{} `json:"notifications"`
	Type          string        `json:"type"`
	Status        string        `json:"status"`
	Modified      int64         `json:"modified"`
	Enable        string        `json:"enable"`
	Public        bool          `json:"public"`
	Parameters    struct {
		Target    string `json:"target"`
		Threshold int    `json:"threshold"`
		Sens      int    `json:"sens"`
	} `json:"parameters"`
	Created   int64  `json:"created"`
	Queue     string `json:"queue"`
	UUID      string `json:"uuid"`
	State     int    `json:"state"`
	Firstdown int64  `json:"firstdown"`
}
