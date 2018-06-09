package bun

import (
	"context"
	"log"
	"path/filepath"
	"regexp"
)

type Bundle struct {
	Path  string
	Hosts map[string]Host // IP to Host map
}

func NewBundle(ctx context.Context, path string) (Bundle, error) {
	b := Bundle{Hosts: make(map[string]Host)}
	var err error
	b.Path, err = filepath.Abs(path)
	if err != nil {
		log.Printf("Error occurred while detecting absolute path: %v", err)
		return b, err
	}
	fs := ctx.Value("fs").(FileSystem)
	infos, err := fs.ReadDir(b.Path)
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
		switch groups[5] {
		case "master":
			host.Type = Master
		case "agent":
			host.Type = Agent
		case "agent_public":
			host.Type = PublicAgent
		default:
			panic("dcosbundle.NewBundle: unknown host type: " + groups[5])
		}
		host.IP = groups[1]
		host.Path = filepath.Join(b.Path, info.Name())
		b.Hosts[host.IP] = host
	}
	return b, nil
}

// OpenFile opens a file in a root directory of the bundle. The caller is
// responsible for closing the file.
func (b Bundle) OpenFile(fileType string) (File, error) {
	return OpenFile(b.Path, fileType)
}
