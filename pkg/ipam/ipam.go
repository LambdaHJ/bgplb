/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ipam

import (
	"errors"
	"net"

	goipam "github.com/metal-stack/go-ipam"
)

type IPAMManager struct {
	ipam     goipam.Ipamer
	cidrs    []string
	cidrList []*net.IPNet
}

func NewIPAMManager() *IPAMManager {
	i := goipam.New()
	m := make([]string, 0)
	cidrList := make([]*net.IPNet, 0)
	return &IPAMManager{ipam: i, cidrs: m, cidrList: cidrList}
}

func (im *IPAMManager) NewCidr(cidr string) error {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	_, err = im.ipam.NewPrefix(cidr)
	if err != nil {
		return err
	}
	im.cidrList = append(im.cidrList, ipnet)
	im.cidrs = append(im.cidrs, cidr)
	return nil
}

func (im *IPAMManager) AddUsedIP(ip string) bool {
	if cidr := im.getCidrOfIP(ip); cidr != "" {
		im.ipam.AcquireSpecificIP(cidr, ip)
		return true
	}
	return false
}

func (im *IPAMManager) AcquireSpecificIP(ip string) bool {
	if cidr := im.getCidrOfIP(ip); cidr != "" {
		if _, err := im.ipam.AcquireSpecificIP(cidr, ip); err == nil {
			return true
		}
	}

	return false
}

func (im *IPAMManager) AcquireIP() (string, error) {
	for i := range im.cidrs {
		if ip, err := im.ipam.AcquireIP(im.cidrs[i]); err == nil {
			return ip.IP.String(), err
		} else if err == goipam.ErrNoIPAvailable {
			continue
		}
	}

	return "", errors.New("get ip failed")
}

func (im *IPAMManager) ReleaseIP(ip string) error {

	if cidr := im.getCidrOfIP(ip); cidr != "" {
		if err := im.ipam.ReleaseIPFromPrefix(cidr, ip); err == nil || err == goipam.ErrNotFound {
			return nil
		} else {
			return err
		}
	}

	return nil
}

func (im *IPAMManager) getCidrOfIP(ip string) string {
	IP := net.ParseIP(ip)
	for i := range im.cidrList {
		if im.cidrList[i].Contains(IP) {
			return im.cidrs[i]
		}
	}

	return ""
}
