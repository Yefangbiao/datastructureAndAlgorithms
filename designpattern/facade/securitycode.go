package main

import (
	"fmt"
)

type securityCode struct {
	password string
}

func (s *securityCode) checkSecurityCode(checkPass string) error {
	if checkPass != s.password {
		return fmt.Errorf("error: securityCode is wrong\n")
	}
	return nil
}
