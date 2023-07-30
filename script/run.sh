#!/bin/bash

set -eu

. import-env .env
./gscp $@
