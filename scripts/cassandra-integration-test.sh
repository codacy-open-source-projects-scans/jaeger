#!/bin/bash

# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

set -euxf -o pipefail

export CASSANDRA_USERNAME="cassandra"
export CASSANDRA_PASSWORD="cassandra"
success="false"

usage() {
  echo $"Usage: $0 <cassandra_version> <schema_version>"
  exit 1
}

check_arg() {
  if [ ! $# -eq 3 ]; then
    echo "ERROR: need exactly three arguments, <cassandra_version> <schema_version> <jaeger_version>"
    usage
  fi
}

setup_cassandra() {
  local compose_file=$1
  docker compose -f "$compose_file" up -d
}

dump_logs() {
  local compose_file=$1
  echo "::group::🚧 🚧 🚧 Cassandra logs"
  docker compose -f "${compose_file}" logs
  echo "::endgroup::"
}

teardown_cassandra() {
  local compose_file=$1
   if [[ "$success" == "false" ]]; then
    dump_logs "${compose_file}"
  fi
  docker compose -f "$compose_file" down
}

apply_schema() {
  local image=cassandra-schema
  local schema_dir=plugin/storage/cassandra/
  local schema_version=$1
  local keyspace=$2
  local params=(
    --rm
    --env CQLSH_HOST=localhost
    --env CQLSH_PORT=9042
    --env "TEMPLATE=/cassandra-schema/${schema_version}.cql.tmpl"
    --env "KEYSPACE=${keyspace}"
    --env "CASSANDRA_USERNAME=${CASSANDRA_USERNAME}"
    --env "CASSANDRA_PASSWORD=${CASSANDRA_PASSWORD}"
    --network host
  )
  docker build -t ${image} ${schema_dir}
  docker run "${params[@]}" ${image}
}

run_integration_test() {
  local version=$1
  local major_version=${version%%.*}
  local schema_version=$2
  local jaegerVersion=$3
  local primaryKeyspace="jaeger_v1_dc1"
  local archiveKeyspace="jaeger_v1_dc1_archive"
  local compose_file="docker-compose/cassandra/v$major_version/docker-compose.yaml"

  setup_cassandra "${compose_file}"

  # shellcheck disable=SC2064
  trap "teardown_cassandra ${compose_file}" EXIT

  apply_schema "$schema_version" "$primaryKeyspace"
  apply_schema "$schema_version" "$archiveKeyspace"

  if [ "${jaegerVersion}" = "v1" ]; then
    STORAGE=cassandra make storage-integration-test
  elif [ "${jaegerVersion}" == "v2" ]; then
    STORAGE=cassandra make jaeger-v2-storage-integration-test
  else
    echo "Unknown jaeger version $jaegerVersion. Valid options are v1 or v2"
    exit 1
  fi
  success="true"
}


main() {
  check_arg "$@"

  echo "Executing integration test for $1 with schema $2.cql.tmpl"
  run_integration_test "$1" "$2" "$3"
}

main "$@"
