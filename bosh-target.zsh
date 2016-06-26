#!/bin/bash -eu
function bosh_target {
    # check for $HOME/bosh_config
    if [ ! -f $HOME/.bosh_config ]; then
        exit 1
    fi
    
    # read bosh config
    local bosh_config="$(< $HOME/.bosh_config)"

    # read current "target"
    local target_ip=$(echo $bosh_config | grep "^target:" | sed "s/target: //")
    
    # find the target's alias within "aliases -> target" by matching on the target_ip
    local target_alias_name="$(echo $bosh_config | grep -m 1 "[^\t]: ${target_ip}$" | sed "s/\:.*$//" | sed "s/^[ \t]*//")"
    echo "BOSH -> $target_alias_name "
}
