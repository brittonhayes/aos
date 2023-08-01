# AoS 

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/aos.svg)](https://pkg.go.dev/github.com/brittonhayes/aos)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/aos)](https://goreportcard.com/report/github.com/brittonhayes/aos)

> An unnoficial AoS REST API built with Go and Sqlite.

## ✨ Try the API (hosted)

- API: https://aos-api.com/
- Documentation: https://aos-api.com/docs

## ⚡ Quick Start (self-hosted)

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

*The API is read-only so all endpoints are GET requests.*

✅=Available
🚧=Under Construction

- ✅ `/api/cities` - Get all cities
- ✅ `/api/cities/{id}` - Get a city by ID
- ✅ `/api/grand-alliances` - Get all grand alliances
- ✅ `/api/grand-alliances/{id}` - Get a grand alliance by ID
- ✅ `/api/grand-strategies/` - Get all grand strategies
- ✅ `/api/grand-strategies/{id}` - Get a grand strategy by ID
- 🚧 `/api/units` - Get all units
- 🚧 `/api/units/{id}` - Get a unit by ID
- 🚧 `/api/warscrolls/` - Get all warscrolls
- 🚧 `/api/warscrolls/{id}` - Get a warscroll by ID

## 📈 Monitoring (self-hosted)

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

Any changes to the data hosted by this repository must respect the licensing rules documented by Games Workshop
here [Intellectual Property Policy](https://www.games-workshop.com/en-US/Intellectual-Property-Policy).

We are in no way affiliated with Games Workshop and the Warhammer Age of Sigmar data is the sole property of Games
Workshop. We are abiding by their `Celebrating the Hobby` section of the agreement and not commercializing this data in
any way. This API is purely to help allow users to engage with the wonderful world of Warhammer in a programattic way
through a REST interface rather than the usual PDF.

If you consume the data served through this API, be aware that you are also obligated to respect
[Games Workshop's](https://www.games-workshop.com) [intellectual property guidelines](https://www.games-workshop.com/en-US/Intellectual-Property-Guidelines).

For more information, view our [Contributing Guidelines](https://github.com/brittonhayes/aos/blob/main/CONTRIBUTING.md).

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->