# AoS 

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/aos.svg)](https://pkg.go.dev/github.com/brittonhayes/aos)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/aos)](https://goreportcard.com/report/github.com/brittonhayes/aos)

> An unnoficial AoS REST API built with Go and Sqlite.

## ⚡ Quick Start

Get started with docker-compose to seed the database and start the API server.

```sh
docker-compose up --build --force-recreate --remove-orphans --detach
```

or with [task](https://taskfile.dev/)

```sh
task up
```

## 📖 Documentation

API documentation is available in YAML format within the repository. The OpenAPI spec is used to generate the transport logic thanks to [goapi-gen](https://github.com/discord-gophers/goapi-gen).

- [View the OpenAPI Spec](https://github.com/brittonhayes/aos/blob/main/api/openapi.yaml)

## 🗄️ Adding or Editing Data

All application data used to seed the database is editable in the [fixtures/](https://github.com/brittonhayes/aos/blob/main/fixtures/) directory. The API is built to be self-seeding and read-only, so if there is a need to add more data, simply add it to the fixtures files, rebuild and redeploy the API.

### Example - Adding a new alliance

To add a new entry to the database, just add a new object to the appropriate yaml fixtures file. In this case, `fixtures/alliances.yaml`. This processes is the same for all data types.

```diff
# fixtures/alliances.yaml
    - model: GrandAlliance
      rows:
        - id: order
          name: Order
          description: The forces of Order strive to bring unity and stability to the Mortal Realms. Composed of various factions, they fight against the forces of Chaos.
+       - id: chaos
+         name: Chaos
+         description: The forces of Chaos seek to corrupt and dominate the Mortal Realms. Made up of daemons, monsters, and twisted beings, they spread destruction wherever theygo.
```

- [View the data](https://github.com/brittonhayes/aos/blob/main/fixtures/fixtures.yaml)

## API Endpoints

*The API is read-only*

✅=Available
🚧=Under Construction

- ✅ `/cities` - Get all cities
- ✅ `/cities/{id}` - Get a city by ID
- ✅ `/grand-alliances` - Get all grand alliances
- ✅ `/grand-alliances/{id}` - Get a grand alliance by ID
- ✅ `/grand-strategies/` - Get all grand strategies
- ✅ `/grand-strategies/{id}` - Get a grand strategy by ID
- ✅ `/units` - Get all units
- ✅ `/units/{id}` - Get a unit by ID
- ✅ `/warscrolls/` - Get all warscrolls
- ✅ `/warscrolls/{id}` - Get a warscroll by ID
- ✅ `/graphql` - GraphQL playground
- ✅ `/query` - GraphQL query endpoint

## 🔎 Querying

The API supports GraphQL queries. The GraphQL playground is available at `/graphql` and the query endpoint is available at `/query`.

### Example - Get all units

```graphql
query {
  units {
    id
    name
    description
    grandAlliance
    points
  }
}
```

### Example - Get all units, filtering for a specific name

```graphql
query {
  units(filter: { name: "Lord" }) {
    id
    name
    description
    grandAlliance
    points
  }
}
```

## 📦 Go Client

A Go client is available for the API. More examples are available in the [example/](https://github.com/brittonhayes/aos/tree/main/example) directory.

```go
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/brittonhayes/aos/client"
)

func main() {
	//	Setup a context with a timeout so we're covered in case of a slow response
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new http client
	c := client.NewClient(&http.Client{}, "http://aos-api.localhost/query", nil)

	// Get all allegiances
	resp, err := c.GetAllegiances(ctx, client.AllegianceFilters{})
	if err != nil {
		panic(err)
	}

	// List the allegiances
	for _, a := range resp.Allegiances {
		println(a.Name)
	}
}
```

## 📈 Monitoring

Application observability is instrumented with OpenTelemetry. Telemetry is available in [Grafana](https://grafana.com/grafana/) and [Prometheus](https://prometheus.io/). Application tracing is powered by [Grafana Tempo](https://grafana.com/oss/tempo/). All application services are behind [Traefik](https://doc.traefik.io/traefik/) reverse proxy.

When running with docker-compose, the following services are available:

- API is available at [http://aos.localhost:8090](http://aos.localhost:8090)
- Documentation is available at [http://aos.localhost:8090/docs](http://aos.localhost:8090/docs)
- Traefik Dashboard is available at [http://localhost:8080](http://localhost:8080)
- Prometheus is available at [http://prometheus.localhost:8090](http://prometheus.localhost:8090)
- Grafana is available at [http://grafana.localhost:8090](http://grafana.localhost:8090)

## 🙋 FAQ

- **Q: Where is X {unit,alliance,etc}?** - A: Waiting for you to add it! See the [fixtures/](https://github.com/brittonhayes/aos/blob/main/fixtures/) directory for more information.

## ⚖️ Copyright and Data Ownership

This project provides a self-hosted API solution for Warhammer Age of Sigmar data. We are in no way affiliated with Games Workshop, and all Warhammer Age of Sigmar data is the sole property of Games Workshop. 

As the operator of your own instance of this API, you are responsible for:
1. Reviewing and complying with Games Workshop's [Intellectual Property Guidelines](https://www.warhammer.com/en-GB/legal#IntellectualPropertyGuidlines)
2. Ensuring your usage falls under the `Celebrating the Hobby` section of the agreement
3. Not commercializing the data in any way
4. Operating the API in accordance with Games Workshop's guidelines

This project aims to help users engage with the wonderful world of Warhammer in a programmatic way through a REST interface rather than the usual PDF format. Any changes to the data in your instance must respect the licensing rules documented by Games Workshop.

For more information, view our [Contributing Guidelines](https://github.com/brittonhayes/aos/blob/main/CONTRIBUTING.md).

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->