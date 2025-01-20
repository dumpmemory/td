package main

import (
	"fmt"
	"net/netip"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/gotd/td/telegram/dcs"
)

type CIDRElement struct {
	CIDR string `json:"cidr" yaml:"cidr"`
}

type Ports struct {
	Ports []Port `json:"ports" yaml:"ports"`
}

type Port struct {
	Port string `json:"port" yaml:"port"`
}

type Element struct {
	ToCIDRSet []CIDRElement `json:"toCIDRSet" yaml:"toCIDRSet"`
	ToPorts   []Ports       `json:"toPorts" yaml:"toPorts"`
}

/*

- toCIDRSet:
   - cidr: 192.169.0.1/32
  toPorts:
   - ports:
	   - port: "443"
*/

func main() {
	list := dcs.Prod()
	var out []CIDRElement
	met := make(map[string]struct{})
	for _, v := range list.Options {
		ip, err := netip.ParseAddr(v.IPAddress)
		if err != nil {
			panic(err)
		}
		if v.Port != 443 {
			panic("port != 443")
		}
		prefix := netip.PrefixFrom(ip, ip.BitLen())
		cidr := prefix.String()
		if _, ok := met[cidr]; ok {
			continue
		}
		met[cidr] = struct{}{}
		out = append(out, CIDRElement{CIDR: cidr})
	}
	fmt.Println("# telegram datacenters list")
	fmt.Println("# generated by cmd/dcsidr")
	e := yaml.NewEncoder(os.Stdout)
	e.SetIndent(2)
	if err := e.Encode(Element{
		ToCIDRSet: out,
		ToPorts: []Ports{
			{
				Ports: []Port{
					{Port: "443"},
				},
			},
		},
	}); err != nil {
		panic(err)
	}
}
