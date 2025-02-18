#!/bin/sh
cd /metrics
/metrics/bin/docker-metrics &
cd -
/bin/prometheus "$@"
