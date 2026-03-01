package main

import (
	"fmt"

	api "github.com/ALT-F4-LLC/vorpal/sdk/go/pkg/api/artifact"
	"github.com/ALT-F4-LLC/vorpal/sdk/go/pkg/artifact"
	"github.com/ALT-F4-LLC/vorpal/sdk/go/pkg/config"
)

var Systems = []api.ArtifactSystem{
	api.ArtifactSystem_AARCH64_DARWIN,
}

func main() {
	ctx:= config.GetContext()
	fmt.Println("Building dotfiles...")

	artifact.NewUserEnvironment("mvaldes", Systems).
		WithSymlinks(map[string]string{
			"$HOME/git/dotfiles/.config/aerospace":   "$HOME/.config/aerospace",
			"$HOME/git/dotfiles/.config/atuin":       "$HOME/.config/atuin",
			"$HOME/git/dotfiles/.config/bat":         "$HOME/.config/bat",
			"$HOME/git/dotfiles/.config/direnv":      "$HOME/.config/direnv",
			"$HOME/git/dotfiles/.config/ghostty":     "$HOME/.config/ghostty",
			"$HOME/git/dotfiles/.config/glazewm":     "$HOME/.config/glazewm",
			"$HOME/git/dotfiles/.config/lazygit":     "$HOME/.config/lazygit",
			"$HOME/git/dotfiles/.config/nvim":        "$HOME/.config/nvim",
			"$HOME/git/dotfiles/.config/opencode":    "$HOME/.config/opencode",
			"$HOME/git/dotfiles/.config/starship.toml": "$HOME/.config/starship.toml",
			"$HOME/git/dotfiles/.config/tmux":        "$HOME/.config/tmux",
			"$HOME/git/dotfiles/.config/wezterm":     "$HOME/.config/wezterm",
			"$HOME/git/dotfiles/.config/zed":         "$HOME/.config/zed",
			"$HOME/git/dotfiles/.config/zsh":         "$HOME/.config/zsh",
		}).
		WithEnvironments([]string{"CLAUDE_TEST=1", "FOO=BAR"}).
		WithArtifacts(RequiredPackages(ctx)).
		Build(ctx)

	// Run the build
	ctx.Run()
}
