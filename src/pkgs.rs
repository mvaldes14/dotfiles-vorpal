use anyhow::Result;
use vorpal_sdk::{
    api::artifact::ArtifactSystem::{self, Aarch64Darwin},
    artifact::{self, step, Artifact},
    context::ConfigContext,
};

pub const SYSTEMS: [ArtifactSystem; 1] = [Aarch64Darwin];

async fn shell_artifact(
    context: &mut ConfigContext,
    name: &str,
    artifacts: &[String],
    script: &str,
) -> Result<String> {
    let shell_step = step::shell(context, artifacts, &[], script.to_string(), &[]).await?;

    Artifact::new(name, vec![shell_step], SYSTEMS.to_vec())
        .build(context)
        .await
}

pub async fn required_packages(context: &mut ConfigContext) -> Result<Vec<String>> {
    let mut packages = Vec::new();

    packages.push(
        shell_artifact(
            context,
            "neovim",
            &[],
            r#"
        ARCH=$(uname -m)
        if [ "$ARCH" = "arm64" ]; then
            URL="https://github.com/neovim/neovim/releases/download/nightly/nvim-macos-arm64.tar.gz"
        else
            URL="https://github.com/neovim/neovim/releases/download/nightly/nvim-macos-x86_64.tar.gz"
        fi

        curl -L "$URL" | tar xz
        cp -r nvim-macos-*/  $VORPAL_OUTPUT/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "bat",
            &[],
            r#"
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/sharkdp/bat/releases/latest | sed 's|.*/||')
        curl -sfL "https://github.com/sharkdp/bat/releases/download/${VERSION}/bat-${VERSION}-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp bat-*/bat $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "fd",
            &[],
            r#"
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/sharkdp/fd/releases/latest | sed 's|.*/||')
        curl -sfL "https://github.com/sharkdp/fd/releases/download/${VERSION}/fd-${VERSION}-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp fd-*/fd $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "fzf",
            &[],
            r#"
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/junegunn/fzf/releases/latest | sed 's|.*/||')
        VER="${VERSION#v}"
        curl -sfL "https://github.com/junegunn/fzf/releases/download/${VERSION}/fzf-${VER}-darwin_arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp fzf $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "go-task",
            &[],
            r#"
        curl -sfL "https://github.com/go-task/task/releases/latest/download/task_darwin_arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp task $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "helm",
            &[],
            r#"
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/helm/helm/releases/latest | sed 's|.*/||')
        curl -sfL "https://get.helm.sh/helm-${VERSION}-darwin-arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp darwin-arm64/helm $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "k9s",
            &[],
            r#"
        curl -sfL "https://github.com/derailed/k9s/releases/latest/download/k9s_Darwin_arm64.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp k9s $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "ripgrep",
            &[],
            r#"
        VERSION=$(curl -sfL -o /dev/null -w '%{url_effective}' https://github.com/BurntSushi/ripgrep/releases/latest | sed 's|.*/||')
        curl -sfL "https://github.com/BurntSushi/ripgrep/releases/download/${VERSION}/ripgrep-${VERSION}-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp ripgrep-*/rg $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    packages.push(
        shell_artifact(
            context,
            "starship",
            &[],
            r#"
        curl -sfL "https://github.com/starship/starship/releases/latest/download/starship-aarch64-apple-darwin.tar.gz" | tar xz
        mkdir -p $VORPAL_OUTPUT/bin
        cp starship $VORPAL_OUTPUT/bin/
    "#,
        )
        .await?,
    );

    let bun = artifact::bun::Bun::new().build(context).await?;
    let npm_packages = shell_artifact(
        context,
        "npm-packages",
        std::slice::from_ref(&bun),
        r#"
        mkdir -p "$VORPAL_OUTPUT/bin"
        BUN_INSTALL="$VORPAL_OUTPUT" bun install -g \
            obsidian-headless \
            @doist/todoist-cli
        cp "$(which bun)" "$VORPAL_OUTPUT/bin/bun"
    "#,
    )
    .await?;
    packages.push(npm_packages);

    Ok(packages)
}
