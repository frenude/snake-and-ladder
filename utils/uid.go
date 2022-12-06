package utils

import "github.com/rs/xid"

func GetId() string {
	guid := xid.New()
	return guid.String()
}
