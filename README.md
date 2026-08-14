Here are some examples of running [Go](https://golang.org/) applications in containers and on Kubernetes. The examples range from a simple Hello World application to basic microservice with [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) actions.

This source code is used by the O'Reilly interactive lab
[Kubernetes Containers: Go Apps in Containers on Kubernetes](https://learning.oreilly.com/scenarios/~/9781492090144/),
which builds each of these applications into a container image, pushes them to a private
registry, and runs them on a Kubernetes cluster. The lab runs in a live cluster
provisioned for you in the browser, and is available to O'Reilly Learning subscribers as
part of the subscription.

## The three examples

| folder | what it shows |
| --- | --- |
| `hello` | the smallest possible Go container, built in one stage |
| `basic` | a single-file app calling a public API, built multi-stage into `scratch` |
| `restafarian-gin` | a REST microservice with CRUD handlers, using the Gin framework |

## Go version, and why the `go` directive is what it is

The `go` line in each `go.mod` is a **minimum**, not a pin: it says the oldest toolchain
that can build the module. It is deliberately kept at `1.22`, because the lab runs
`go run` directly on the lab machine and that machine's Go is what has to satisfy it.
Raising the directive above the lab's Go makes the toolchain try to download a newer one,
which the lab environment cannot do:

```
go: downloading go1.24 (linux/amd64)
go: download go1.24 for linux/amd64: toolchain not available
```

That is not a reason to hold the dependencies back. The Dockerfiles build with
`golang:1.25-alpine3.22`, because a newer toolchain builds an older floor happily. Only
the directive has to suit the lab.

Dependencies are held to versions that work at that floor:

| dependency | version | needs Go |
| --- | --- | --- |
| `gin-gonic/gin` | v1.10.1 | 1.20 |
| `tidwall/gjson` | v1.18.0 | 1.12 |

Newer ones exist and are not usable here yet: gin v1.11 needs Go 1.23 and v1.12 needs
1.25, gjson v1.19 needs 1.23. If the lab image moves to a newer Go, these can move with
it.

## Running locally

```
cd restafarian-gin && go run *.go     # serves on :8080
curl http://localhost:8080/ideas/v1/
```

## Weather Apps

The `basic` folder demonstrates a small Go app that obtains the current weather in London. A The Open-Meteo project offers a Weather Forecast API and provides a [REST URL builder](https://open-meteo.com/en/docs).

Given a request like this:

`curl https://api.open-meteo.com/v1/forecast?latitude=51.5002&longitude=-0.1262&daily=weathercode&timezone=Europe%2FLondon`

The result would be:

```json
{
  "longitude": -0.120000124,
  "daily_units": {
    "time": "iso8601",
    "weathercode": "wmo code"
  },
  "latitude": 51.5,
  "generationtime_ms": 0.29206275939941406,
  "elevation": 13.1953125,
  "daily": {
    "time": [
      "2022-06-01",
      "2022-06-02",
      "2022-06-03",
      "2022-06-04",
      "2022-06-05",
      "2022-06-06",
      "2022-06-07"
    ],
    "weathercode": [
      80,
      45,
      3,
      3,
      80,
      63,
      61
    ]
  },
  "utc_offset_seconds": 3600
}
```

The demonstration app simply makes this call to discover the somewhat predictable drizzly conditions in London.
