use anyhow::Result;
use dotfiles::pkgs::{required_packages, SYSTEMS};
use vorpal_sdk::{artifact, context::get_context};

#[tokio::main]
async fn main() -> Result<()> {
    let mut context = get_context().await?;
    println!("Building dotfiles...");

    let packages = required_packages(&mut context).await?;

    artifact::UserEnvironment::new("mvaldes", SYSTEMS.to_vec())
        .with_artifacts(packages)
        .with_symlinks(vec![
            ("$HOME/git/dotfiles/.config/atuin", "$HOME/.config/atuin"),
            ("$HOME/git/dotfiles/.config/bat", "$HOME/.config/bat"),
            ("$HOME/git/dotfiles/.config/direnv", "$HOME/.config/direnv"),
            (
                "$HOME/git/dotfiles/.config/ghostty",
                "$HOME/.config/ghostty",
            ),
            (
                "$HOME/git/dotfiles/.config/lazygit",
                "$HOME/.config/lazygit",
            ),
            ("$HOME/git/dotfiles/.config/nvim", "$HOME/.config/nvim"),
            (
                "$HOME/git/dotfiles/.config/opencode",
                "$HOME/.config/opencode",
            ),
            (
                "$HOME/git/dotfiles/.config/starship.toml",
                "$HOME/.config/starship.toml",
            ),
            ("$HOME/git/dotfiles/.config/tmux", "$HOME/.config/tmux"),
            (
                "$HOME/git/dotfiles/.config/wezterm",
                "$HOME/.config/wezterm",
            ),
            ("$HOME/git/dotfiles/.config/zed", "$HOME/.config/zed"),
            ("$HOME/git/dotfiles/.config/zsh", "$HOME/.config/zsh"),
            ("$HOME/git/dotfiles/claude", "$HOME/.claude"),
            ("$HOME/git/dotfiles/.gitconfig", "$HOME/.gitconfig"),
        ])
        .build(&mut context)
        .await?;

    context.run().await
}
