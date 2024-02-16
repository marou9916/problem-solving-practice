#!/usr/bin/bash

get_latest_go_version() {
    local latest_version=$(curl -s https://golang.org/dl/ | grep -oE 'id="downloadBox">Go [0-9]+\.[0-9]+\.[0-9]+' | head -n 1 | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
    echo "$latest_version"
}

get_current_go_version() {
    local current_version=$(go version | grep -oE 'go[0-9]+\.[0-9]+\.[0-9]+' | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
    echo "$current_version"
}

install_latest_go() {
    local latest_version=$1
    echo "Downloading and installing Go version $latest_version..."
    curl -O "https://golang.org/dl/go$latest_version.$(uname -s)-$(uname -m).tar.gz"
    sudo tar -C /usr/local -xzf "go$latest_version.$(uname -s)-$(uname -m).tar.gz"
    echo "Go has been updated successfully."
}

while true; do
    latest_version=$(get_latest_go_version)
    current_version=$(get_current_go_version)

    if [[ "$latest_version" != "$current_version" ]]; then
        install_latest_go "$latest_version"
    else
        echo "Go is already up to date."
    fi

    # Vérification toutes les 24 heures
    sleep 86400
done
