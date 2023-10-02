# Gin-based REST Prometheus-Scrapable System Metrics Service

This project provides a RESTful service that returns system status data.  The items returned are:

    - Kernel IO Wait Percentage
    - Kernel Free Memory
    - The number of processes running on the system

The service exposes a single endpoint that returns this data:

    `/metrics`

## Building

In the top-level directory, the following commands can be run to build the
executable or a Docker image.

    - Building the executable:

        `make`

    - Building a Docker image:

        `docker build -t prometheus/metrics_agent:latest .`

## Running

- Standalone application:

    The service can be run as a standalone application after a successful
    build of the executable with the following command:

        `agent`

    On the local system, the service can then be accessed with the following
    curl command:

        `curl http://127.0.0.1:9091/metrics`

- Containerized application:

    The service can be run as a containerized application after a successful
    Docker image build with the following command:

        `
            docker run -d \
                --rm \
                --name metrics
                -p 9091:localhost:9091
                prometheus/metrics_agent:latest
        `

