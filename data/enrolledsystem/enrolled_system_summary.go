package enrolledsystem

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data/enrolmentkey"
	"github.com/enclave-networks/go-enclaveapi/data/tag"
)

// A basic model representing a single system.
type EnrolledSystemSummary struct {
	SystemId                     SystemId
	Description                  string
	Type                         int
	State                        int
	ConnectedAt                  time.Time
	LastSeen                     time.Time
	EnrolledAt                   time.Time
	EnrolmentKeyId               enrolmentkey.EnrolmentKeyId
	EnrolmentKeyDescription      string
	IsEnabled                    bool
	ConnectedFrom                string
	Hostname                     string
	PlatformType                 string
	OSVersion                    string
	EnclaveVersion               string
	Tags                         []tag.TagReference
	DisconnectedRetentionMinutes int
}
