# <img src="./logo-dark.svg" width="50px"/> Completerr

*Note: This project is in an alpha stage. Your code contributions, bug reports and feature requests are greatly appreciated! Planned features and known issues are tracked in the issues tab. Feel free to contribute*

Completerr connects to your Sonarr and Radarr and periodically triggers searches for missing items in your library.

### Motive

The goal of the project is to build a service that will periodically search for missing items to gradually backfill items that may be missing for whatever reason or to facilitate a workflow where large lists are added/maintained via plex meta manager or another mechanism but not searched all at once.

## Getting started

Completerr is distrubuted as a docker container. You can start include it using a docker-compose file that looks like this:

```yaml
version: "3.7"
services:
  completerr:
    image: ghcr.io/completerr/completerr:main
    ports:
      - 8080:8080
    volumes:
      - ./config:/config
```

You'll need a `config.yaml` file located in the mounted folder, which should look like:

```yaml
sonarr:
  api_key: "YOUR-SONARR-API-KEY"
  url: "http://YOUR-SONARR-URL:8989/"
  library_sync_interval: 300 # How frequently to sync completerr library with sonarr library
  search:
    interval: 60 # How frequently to search (minutes)
    count: 100 # How many items to queue searches for each run
    backoff_days: 30  # How many days to wait before searching the same item a second time
radarr:
  api_key: "YOUR-RADARR-API-KEY"
  url: "http://YOUR-RADARR-URL:8989/"
  library_sync_interval: 60 # How frequently to sync completerr library with sonarr library
  search:
    interval: 60 # How frequently to search (minutes)
    count: 30 # How many items to queue searches for each run
    backoff_days: 30 # How many days to wait before searching the same item a second time
```

After running `docker-compose up` the service should be available on port `http://localhost:8080` (assuming you are running locally)
