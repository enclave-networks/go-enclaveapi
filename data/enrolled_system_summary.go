package data

import "time"

type EnrolledSystemSummary struct {
	SystemId                     *string
	Description                  *string
	Type                         *int
	State                        *int
	ConnectedAt                  *time.Time
	LastSeen                     *time.Time
	EnrolledAt                   *time.Time
	EnrolmentKeyId               *int
	EnrolmentKeyDescription      *string
	IsEnabled                    *bool
	ConnectedFrom                *string
	Hostname                     *string
	PlatformType                 *string
	OSVersion                    *string
	EnclaveVersion               *string
	Tags                         []*TagReference
	DisconnectedRetentionMinutes *int
}
