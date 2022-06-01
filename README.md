Here are some examples of running [Go](https://golang.org/) applications in containers and on Kubernetes. The examples range from a simple Hello World application to basic microservice with [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) actions.

This source code is referenced by the Katacoda scenario [Go Apps to Kubernetes](https://katacoda.com/javajon/courses/kubernetes-containers).

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
