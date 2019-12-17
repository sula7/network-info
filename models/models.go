package models

type (
	IPAddresses struct {
		List []IP
	}

	IP struct {
		Address string `json:"ip"`
	}
)
