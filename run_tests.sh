#!/bin/bash

set -euo pipefail

_base_dir=$(dirname "$(readlink -f "${BASH_SOURCE[0]}")")
source "$_base_dir/vendor/github.com/reconquest/test-runner.bash/test-runner.bash"

go build -o "tests:treetrunks"

:cleanup() {
    rm "./tests:treetrunks"
}

trap ":cleanup" EXIT

test-runner:run "${@}"
