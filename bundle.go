package bun

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

// HostType represent different types of the hosts.
type HostType string

// File is a safe way to access bundle files.
type File io.ReadCloser

const (
	// Master host type
	Master HostType = "master"
	// Agent host type
	Agent = "agent"
	// PublicAgent host type
	PublicAgent = "public agent"
)

// Host represents a host in a DC/OS cluster.
type Host struct {
	Type HostType
	IP   string
	fileOwner
}

// Bundle describes DC/OS diagnostics bundle.
type Bundle struct {
	Hosts        map[string]Host // IP to Host map
	Masters      map[string]Host
	Agents       map[string]Host
	PublicAgents map[string]Host
	fileOwner
}

// NewBundle creates new Bundle
func NewBundle(path string) (Bundle, error) {
	b := Bundle{
		Hosts:        make(map[string]Host),
		Masters:      make(map[string]Host),
		Agents:       make(map[string]Host),
		PublicAgents: make(map[string]Host),
	}
	var err error
	b.Path, err = filepath.Abs(path)
	if err != nil {
		log.Printf("bun.NewBundle: cannot determine absolute path: %v", err)
		return b, err
	}
	infos, err := ioutil.ReadDir(b.Path)
	if err != nil {
		return b, err
	}
	const restr = `^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))_(agent_public|agent|master)$`
	re := regexp.MustCompile(restr)
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		groups := re.FindStringSubmatch(info.Name())
		if groups == nil {
			continue
		}
		host := Host{}
		host.IP = groups[1]
		host.Path = filepath.Join(b.Path, info.Name())
		switch groups[5] {
		case "master":
			host.Type = Master
			b.Masters[host.IP] = host
		case "agent":
			host.Type = Agent
			b.Agents[host.IP] = host
		case "agent_public":
			host.Type = PublicAgent
			b.PublicAgents[host.IP] = host
		default:
			panic(fmt.Sprintf("bun.NewBundle: unknown host type: %v", groups[5]))
		}
		b.Hosts[host.IP] = host
	}
	return b, nil
}
