#!/usr/bin/env bash

source "${PROJECT_ABSOLUTE_PATH}"/env
start=$(date +%s)

pushd "${PROJECT_ABSOLUTE_PATH}"/certificates || exit

rm -rf ./keys ./certs ../configs/certs ./keys_local ./certs_local ../configs/certs_local
mkdir ./keys ./certs ./keys_local ./certs_local ../configs/certs_local
while IFS="" read -r component || [ -n "$component" ]; do
  if ! test -z "${component}"; then
    go run ./generate_cert.go --rsa-bits 1024 --ca --start-date "Jan 1 00:00:00 2020" --duration=100000h --host "${component}"
    mv cert.pem ./certs_local/"${component}".pem
    mv key.pem ./keys_local/"${component}".pem
  fi
done <endpoints_local

cp ./certs_local/* ../configs/certs_local
cp ./certs_local/* ./certs
cp ./keys_local/* ./keys

popd || exit

end=$(date +%s)

echo Making certificates in $(expr $end - $start) seconds.
