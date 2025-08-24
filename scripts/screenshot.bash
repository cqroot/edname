#!/usr/bin/env bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
WORK_DIR=${SCRIPT_DIR}/dir

rm -f "${SCRIPT_DIR}/edname.cast"
rm -rf "${WORK_DIR}"

mkdir "${WORK_DIR}"
for i in $(seq -w 0 5); do
    touch "${WORK_DIR}/testfile${i}"
done

cd "${WORK_DIR}" || exit 1
../asciinema.exp
