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

	batStep, err := artifact.Shell(ctx, nil, []string{}, `
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/sharkdp/bat/releases/latest | sed 's|.*/||')
        curl -sfL "https://github.com/sharkdp/bat/releases/download/${VERSION}/bat-${VERSION}-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp bat-*/bat $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	bat, err := artifact.NewArtifact("bat", []*api.ArtifactStep{batStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, bat)

	fdStep, err := artifact.Shell(ctx, nil, []string{}, `
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/sharkdp/fd/releases/latest | sed 's|.*/||')
        curl -sfL "https://github.com/sharkdp/fd/releases/download/${VERSION}/fd-${VERSION}-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp fd-*/fd $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	fd, err := artifact.NewArtifact("fd", []*api.ArtifactStep{fdStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, fd)

	fzfStep, err := artifact.Shell(ctx, nil, []string{}, `
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/junegunn/fzf/releases/latest | sed 's|.*/||')
        VER="${VERSION#v}"
        curl -sfL "https://github.com/junegunn/fzf/releases/download/${VERSION}/fzf-${VER}-darwin_arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp fzf $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	fzf, err := artifact.NewArtifact("fzf", []*api.ArtifactStep{fzfStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, fzf)

	goTaskStep, err := artifact.Shell(ctx, nil, []string{}, `
        curl -sfL "https://github.com/go-task/task/releases/latest/download/task_darwin_arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp task $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	goTask, err := artifact.NewArtifact("go-task", []*api.ArtifactStep{goTaskStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, goTask)

	helmStep, err := artifact.Shell(ctx, nil, []string{}, `
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/helm/helm/releases/latest | sed 's|.*/||')
        curl -sfL "https://get.helm.sh/helm-${VERSION}-darwin-arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp darwin-arm64/helm $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	helm, err := artifact.NewArtifact("helm", []*api.ArtifactStep{helmStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, helm)

	k9sStep, err := artifact.Shell(ctx, nil, []string{}, `
        curl -sfL "https://github.com/derailed/k9s/releases/latest/download/k9s_Darwin_arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp k9s $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	k9s, err := artifact.NewArtifact("k9s", []*api.ArtifactStep{k9sStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, k9s)

	ripgrepStep, err := artifact.Shell(ctx, nil, []string{}, `
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/BurntSushi/ripgrep/releases/latest | sed 's|.*/||')
        curl -sfL "https://github.com/BurntSushi/ripgrep/releases/download/${VERSION}/ripgrep-${VERSION}-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp ripgrep-*/rg $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	ripgrep, err := artifact.NewArtifact("ripgrep", []*api.ArtifactStep{ripgrepStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, ripgrep)

	starshipStep, err := artifact.Shell(ctx, nil, []string{}, `
        curl -sfL "https://github.com/starship/starship/releases/latest/download/starship-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp starship $VORPAL_OUTPUT/bin/
    `, nil)
	if err != nil {
		panic(err)
	}
	starship, err := artifact.NewArtifact("starship", []*api.ArtifactStep{starshipStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, starship)

	bun, err := artifact.Bun(ctx)
	if err != nil {
		panic(err)
	}

	npmPkgsStep, err := artifact.Shell(ctx, []*string{bun}, []string{}, `
        mkdir -p "$VORPAL_OUTPUT/bin"
        BUN_INSTALL="$VORPAL_OUTPUT" bun install -g \
            obsidian-headless \
            @doist/todoist-cli
        cp "$(which bun)" "$VORPAL_OUTPUT/bin/bun"
    `, nil)
	if err != nil {
		panic(err)
	}
	npmPkgs, err := artifact.NewArtifact("npm-packages", []*api.ArtifactStep{npmPkgsStep}, system).Build(ctx)
	if err != nil {
		panic(err)
	}
	pkgs = append(pkgs, npmPkgs)

	return pkgs
}
