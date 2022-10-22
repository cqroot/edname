#!/usr/bin/env bash

TESTBASE='./internal/renamer/testdata'

mkdir -p ${TESTBASE}

touch \
    ${TESTBASE}/.test_file_a \
    ${TESTBASE}/.test_file_b \
    ${TESTBASE}/test_file_a \
    ${TESTBASE}/test_file_b

mkdir \
    ${TESTBASE}/.test_dir_a \
    ${TESTBASE}/.test_dir_b \
    ${TESTBASE}/test_dir_a \
    ${TESTBASE}/test_dir_b
