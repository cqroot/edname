#!/usr/bin/env bash

SCRIPT_PATH=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
PROJ_PATH=$(dirname "${SCRIPT_PATH}")
TEST_DATA_PATH="${PROJ_PATH}/internal/renamer/testdata"

rm -rf ${TEST_DATA_PATH}
mkdir -p ${TEST_DATA_PATH}

touch \
	${TEST_DATA_PATH}/.test_file_a \
	${TEST_DATA_PATH}/.test_file_b \
	${TEST_DATA_PATH}/test_file_a \
	${TEST_DATA_PATH}/test_file_b

mkdir \
	${TEST_DATA_PATH}/.test_dir_a \
	${TEST_DATA_PATH}/.test_dir_b \
	${TEST_DATA_PATH}/test_dir_a \
	${TEST_DATA_PATH}/test_dir_b
