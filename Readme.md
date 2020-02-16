# Simple metrics HTTP client for Linux

[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://choosealicense.com/licenses/mit/)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](mailto:dev@nikitaisakov.com?subject=[GitHub]%20Ask%20me%20anything)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/razzzila-dev/golang-server-metrix-client/issues)

HTTP server for Linux machines which returns the base metrics:
- CPU
- RAM
- Disk

A simple proxy that makes basic authentication and returns results of execution sh calls.


## Installation
### Requirements:
 - Unix machine with `df`, `free`, `tr`, `awk`, `grep` packages and `/proc/stat` file
 - GoLang 1.3
 
### Package install
Download archive and unzip it or clone this repository using `git clone`.
Change working directory to the package root folder and run:
```BASH
go mod vendor
go build
```

## Configure
In `config.yml` file stored basic configurations:
 - url_prefix
 - port
 - key  

| Param name | Description                                                                                                               | Default                    | Examples                                                 |
|------------|---------------------------------------------------------------------------------------------------------------------------|----------------------------|----------------------------------------------------------|
| key        | API key. Currently used as an authentication mechanism. String value. Recommended length 16+. !!!SHOULD BE URI ENCODED!!! | N2XMMQrt99Kk8JrT           | uXA4UM2EAmrY9aMp69vlfNaB69fuRjr 6VrrTePmhnDztsh W6vlfNaB |
| port       | PORT which will be listened.                                                                                              | 32642                      | 4034 3333 58232                                          |
| url_prefix | A part of URI which goes right after host and in front of API key. !!!SHOULD BE URI ENCODED!!!                            | droplet-metrics/raz-client | hidden/path                                              |

## Usage
Just run built file `./golang-server-metrics-client`.

Or run as service.
To doing so, first change settings in the `./systemd-config.sample`.
Most important here is `ExecStart` field which points to the built file.
Also you can change `User` field.
```BASH
cp ./systemd-config.sample /etc/systemd/system/metrics-http-client.service
systemctl status metrics-http-client
systemctl stop metrics-http-client
systemctl start metrics-http-client
```

## TO-DO
In plans to refactor a little bit and add next functionality:

CPU
- getFullCPUMetrics(Time) /* return all from below */
- getAverageCPUMetric(Time)
- getLowestCPUMetric(Time)
- getHighestCPUMetric(Time)

CPU Each Core
- getFullEachCoreMetrics() /* return all from below */
- getEachCoreMetrics()
- getEachCoreAverateMetrics(Time)
- getEachCoreLowestMetrics(Time)
- getEachCoreHighestMetrics(Time)

RAM
- getAverageRAM(Time)
- getLowestRAM(Time)
- getHighestRAM(Time)

Add SSL.

## License
[MIT](https://choosealicense.com/licenses/mit/)
