package cmd

import (
	_ "embed"
	"fmt"
	"github.com/qauzy/trat/pkg/base"
	"github.com/qauzy/trat/pkg/rest"
	"github.com/qauzy/trat/pkg/rest/model"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed banner.txt
var banner string

func Start(cfg *model.StartConfig) {
	fmt.Println(banner)
	srv, listener, err := rest.BuildServer(cfg)
	if err != nil {
		panic(err)
	}
	downloadCfg, err := rest.Downloader.GetConfig()
	if err != nil {
		panic(err)
	}
	if downloadCfg.FirstLoad {
		// Set default download dir, in docker, it will be ${exe}/Downloads, else it will be ${user}/Downloads
		var downloadDir string
		if base.InDocker == "true" {
			downloadDir = filepath.Join(filepath.Dir(cfg.StorageDir), "Downloads")
		} else {
			userDir, err := os.UserHomeDir()
			if err == nil {
				downloadDir = filepath.Join(userDir, "Downloads")
			}
		}
		if downloadDir != "" {
			downloadCfg.DownloadDir = downloadDir
			rest.Downloader.PutConfig(downloadCfg)
		}
	}
	fmt.Printf("Server start success on http://%s\n", listener.Addr().String())
	if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
