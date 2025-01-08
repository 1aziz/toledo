#!/bin/bash
source .env

gum style --border rounded --margin "1" --padding "1 2" --border-foreground 212 "Hello, there! Welcome to $(gum style --foreground '#FFFF00' 'Toledo')."
KIND="Local"; GKE="GKE"

PLATFORM=$(gum choose "kind" "gke" --header "Where to create your Toledo cluster?");

check_and_export_variable() {
    if [ -z "${!1}" ]; then
        VALUE=$(gum input --prompt "$1: " --prompt.foreground '#FFFF00'); export "$1"="$VALUE"
    fi
}

check_and_export_variable "GH_REPO"
check_and_export_variable "GH_USER"
check_and_export_variable "GH_TOKEN"
check_and_export_variable "PLATFORM"

if [ $PLATFORM == 'kind' ]; then
  gum confirm "Any existing kind cluster with name 'toledo-local' will be deleted, are you sure?" || exit 1
  gum spin -s moon --title "Creating the local cluster 'toledo-local' with $(gum style --foreground '#04B575' 'Kind').." --show-error -- task kind:create
fi

gum spin -s moon --title "Appling default $(gum style --foreground '#04B575' 'regcred').." --show-error -- task apply:regcred
gum spin -s moon --title "Bootstapping the cluster with $(gum style --foreground '#04B575' 'Flux').." --show-error -- task flux_bootstrap

test $? -ne 0 || echo "$(gum style --foreground '#FFFF00' 'Your cluster is bootstrapped! :tada: :rocket:')" | gum format -t emoji
