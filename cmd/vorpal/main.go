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
		WithSymlinks(map[string]string{"$HOME/git/dotfiles/.config": "$HOME/.config"}).
		WithEnvironments([]string{"CLAUDE_TEST=1", "FOO=BAR"}).
		WithArtifacts(RequiredPackages(ctx)).
		Build(ctx)

	// Run the build
	ctx.Run()
}
