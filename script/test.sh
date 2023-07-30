#!/bin/bash

set -eu

. import-env .env
make test
