package main

import (
	api "github.com/ALT-F4-LLC/vorpal/sdk/go/pkg/api/artifact"
	"github.com/ALT-F4-LLC/vorpal/sdk/go/pkg/artifact"
	"github.com/ALT-F4-LLC/vorpal/sdk/go/pkg/config"
)

var system = []api.ArtifactSystem{api.ArtifactSystem_AARCH64_DARWIN}

func RequiredPackages(ctx *config.ConfigContext) []*string {
	pkgs := []*string{}

	neovimStep, err := artifact.Shell(ctx, nil, []string{}, `
        ARCH=$(uname -m)
        if [ "$ARCH" = "arm64" ]; then
            URL="https://github.com/neovim/neovim/releases/download/nightly/nvim-macos-arm64.tar.gz"
        else
            URL="https://github.com/neovim/neovim/releases/download/nightly/nvim-macos-x86_64.tar.gz"
        fi

        curl -L "$URL" | tar xz
        cp -r nvim-macos-*/  $VORPAL_OUTPUT/
    `, nil)
	if err != nil {
		panic(err)
	}

	neovim, err := artifact.NewArtifact("neovim", []*api.ArtifactStep{neovimStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}

	pkgs = append(pkgs, neovim)

	return pkgs
}
