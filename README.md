# Gin-based RESTful, Prometheus-Scrapable System Metrics Service

This project provides a RESTful service that returns system status data.  The
items returned are:

    - Kernel IO Wait Percentage
    - Kernel Free Memory
    - The number of processes running on the system

The service exposes a single endpoint that returns this data:

    /metrics

## Building

In the top-level directory, the following commands can be run to build the
executable or a Docker image.

    - Building the executable:

        make

    - Building a Docker image:

        docker build -t prometheus/metrics_agent:latest .

## Testing

The unit tests for this project can be run with the following command in the
top-level directory:

    go test

## Running

- Standalone application:

    The service can be run as a standalone application after a successful
    build of the executable with the following command:

        agent

    On the local system, the service can then be accessed with the following
    curl command:

        curl http://127.0.0.1:9091/metrics

- Containerized application:

    The service can be run as a containerized application after a successful
    Docker image build with the following command:

            docker run -d \
                --rm \
                --name metrics \
                -e "PROMETHEUS_AGENT_PROC=/mnt/host/proc" \
                -e "PROMETHEUS_AGENT_LISTEN=:9091" \
                -p 127.0.0.1:9091:9091 \
                -v /proc:/mnt/host/proc:ro \
                prometheus/metrics_agent:latest

    As was the case for the standalone application, the containerized instance
    of the service can be accessed with the following curl command from the
    host system:

        curl http://127.0.0.1:9091/metrics

