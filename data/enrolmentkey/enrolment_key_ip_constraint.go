package enrolmentkey

import (
	"fmt"
	"net"
	"strings"
)

type EnrolmentKeyIpConstraint struct {
	Range       string
	Description string
}

func CreateEnrolmentKeyIpConstraintFromCidr(cidrNotation string) (*EnrolmentKeyIpConstraint, error) {
	if len(strings.Trim(cidrNotation, " ")) == 0 {
		err := fmt.Errorf("empty string cidrNotation found, please input a correct notation")
		return nil, err
	}

	if !strings.Contains(cidrNotation, "/") {
		err := fmt.Errorf("incorrect CIDR format")
		return nil, err
	}

	return &EnrolmentKeyIpConstraint{
		Range: cidrNotation,
	}, nil
}

func CreateEnrolmentKeyIpConstraintFromIpAddr(ipAddress *net.IPAddr) (*EnrolmentKeyIpConstraint, error) {
	if ipAddress == nil {
		err := fmt.Errorf("ipAddress cannot be nil")
		return nil, err
	}

	return &EnrolmentKeyIpConstraint{
		Range: ipAddress.IP.String(),
	}, nil
}

func CreateEnrolmentKeyIpConstraintFromIpAddrRange(start *net.IPAddr, end *net.IPAddr) (*EnrolmentKeyIpConstraint, error) {
	if start == nil {
		err := fmt.Errorf("start ipAddr cannot be nil")
		return nil, err
	}

	if end == nil {
		err := fmt.Errorf("end ipAddr cannot be nil")
		return nil, err
	}

	return &EnrolmentKeyIpConstraint{
		Range: fmt.Sprintf("%s - %s", start.IP.String(), end.IP.String()),
	}, nil
}
