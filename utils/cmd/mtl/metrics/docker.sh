#!/bin/sh
cd /metrics
/metrics/bin/metrics &
cd -
/bin/prometheus "$@"
