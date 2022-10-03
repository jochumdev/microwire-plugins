#!/bin/bash

function go_mods() {
	local OLD_SHA=$(git rev-list -n 1 ${1})
	local NEW_SHA=$(git rev-list -n 1 ${2})

	local cfiles=$(git diff --name-only ${OLD_SHA} ${NEW_SHA})
	for f in ${cfiles}; do
		if [[ -f "$(dirname ${f})/go.mod" ]]; then
			echo "$(dirname ${f})"
		fi
		if [[ -f "$(dirname ${f}/../go.mod)" ]]; then
			echo $(dirname ${f}/../)
		fi
		if [[ -f "$(dirname ${f}/../../go.mod)" ]]; then
			echo $(dirname ${f}/../../)
		fi
	done
}


function last_tags() {
    local alltags=$(git tag --sort=-creatordate)
    for t in ${alltags}; do
        if [[ ${t} != *"/"* ]]; then
            echo ${t}
        fi
    done

}

function release() {
    local new_tag="$(git tag --points-at HEAD)"
    local last_tag=$(last_tags | head -n 2 | tail -n 1)
    local changed_pkgs=$(go_mods ${new_tag} ${last_tag})
    for pkg in ${changes}; do
        hub release create -m "${pkg}/${new_tag}" "${pkg}/{$new_tag}"; 
    done
}

release