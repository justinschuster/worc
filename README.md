# worc

A Linux CLI addon manager for World of Warcraft

### DESIGN

* go 1.22.2
* cli based

### Libraries
* [cobra 1.8.1](https://github.com/spf13/cobra)
* [viper 1.8.1](https://github.com/spf13/viper)

### Application paths
* cobra: /home/justin/projects/worc/
* battle.net: ~/Games/battlenet/
* wow: ~/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/

### Requirements
* Battle.net launcher installed through [Lutris](https://lutris.net/)

### Install
* clone this repository
* go build
* edit your .bashrc or zshrc with:
    * alias worc=<path-to-worc>
    * ex: alias worc='~/projects/worc/worc'
