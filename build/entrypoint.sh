#! /usr/bin/env bash
set -e

APP_NAME=flange

# build the binary with a certain version
build() {
    echo "--> Building ${APP_NAME} (version ${VERSION} ~ ${SHORTHASH})"
    if [[ ! -d "/go/out/${SHORTHASH}" ]]; then
        echo "Directory /${SHORTHASH} does not exist. Creating..."
        mkdir -p /go/out/${SHORTHASH} && \
        chown -R ${UID}:${UID} /go/out/${SHORTHASH}
    fi

    EXIT_CODE=0
    go get && \
    go install && \
    cp /go/bin/${APP_NAME} /go/out/${SHORTHASH}/${APP_NAME} && \
    chown ${UID}:${UID} /go/out/${SHORTHASH}/${APP_NAME} || EXIT_CODE=1

    if [ ${EXIT_CODE} -eq 0 ]
    then
        echo "[BUILD SUCCESSFUL]"
    fi
    exit ${EXIT_CODE}
}

vet() {
    EXIT_CODE=0
    echo " --> Vetting ${APP_NAME} (version ${VERSION} ~ ${SHORTHASH})"
    exit ${EXIT_CODE}
}

test() {
    EXIT_CODE=0
    echo " --> Retrieving dependencies"
    go get -t -d
    echo " --> Testing ${APP_NAME} (version ${VERSION} ~ ${SHORTHASH})"
    exit ${EXIT_CODE}
}

case $1 in

build)
    shift
    build $@
;;

vet)
    shift
    vet $@
;;

test)
    shift
    test $@
;;

*)
    >&2 echo "'${1}' is not supported."
    exit 1
;;

esac