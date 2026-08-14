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

## Go version

Every module declares `go 1.25`, and the dependencies below need it. The lab installs a
current Go in step 2 rather than relying on the lab image's, which is older:

```
curl -sSL https://go.dev/dl/go1.26.6.linux-amd64.tar.gz -o /tmp/go.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go.tar.gz
export PATH=/usr/local/go/bin:$PATH
```

That version is pinned deliberately. Installing "the latest" of anything makes a lab
irreproducible - two learners on different days get different software, and a failure
cannot be distinguished from an upstream change. Bump it on purpose, then re-run the
lab's test to confirm nothing broke.

Note the lab image's own Go cannot download a newer toolchain automatically: it reports
`toolchain not available`, so `GOTOOLCHAIN=auto` is not a substitute for installing one.

## Dependencies

All pinned, all current as of this writing:

| dependency | version | note |
| --- | --- | --- |
| `gin-gonic/gin` | v1.12.0 | needs Go 1.25 |
| `tidwall/gjson` | v1.19.0 | needs Go 1.23 |
| `quic-go/quic-go` | v0.59.1 | indirect, via gin's HTTP/3 support |

`quic-go` is pinned above what gin resolves to on its own. gin v1.12 brought in HTTP/3,
and the version it selects, v0.59.0, carries a reachable advisory (GO-2026-5676) that
v0.59.1 fixes. `govulncheck ./...` reports no affected code paths with these versions.

Build images use `golang:1.26-alpine3.22`.

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
