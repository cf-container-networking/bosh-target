# bosh-target
This is a bash script that will print the alias of the BOSH director that your BOSH CLI is currently targeting.
The script was created with the intention of printing the BOSH director alias in your terminal along with your
current path (see screenshot). There are a few ways to do this and instructions for how to use this with the 
popular oh-my-zsh are listed below.

### Screenshot
![bosh-target in action](https://raw.githubusercontent.com/kkallday/bosh-target/master/screenshot.png)

### Setup with an oh-my-zsh theme

1. Clone this repo anywhere and either:
    * copy the file into $HOME/.oh-my-zsh/custom OR...
    * sym link `bosh-target.zsh` in the repo into `$HOME/.oh-my-zsh/custom` directory
2. Determine the zsh theme you are using by finding the `ZSH_THEME` value in the `~/.zshrc` file
3. Copy the theme file from `~/.oh-my-zsh/themes/THEME-NAME.zsh-theme` to `~/.oh-my-zsh/custom/themes/THEME-NAME.zsh-theme`
4. Open `~/.oh-my-zsh/custom/themes/THEME-NAME.zsh-theme`
5. Edit the `PROMPT` variable to call the `bosh_target` function
    * The following example will put the BOSH target before all other prompts:
      `PROMPT='%{$fg_bold[red]%}$(bosh_target) ...'`
6. Open a new terminal window or tab and you should see the BOSH director alias in your terminal prompt
