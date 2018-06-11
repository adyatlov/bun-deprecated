# bun

DC/OS [diagnostics bundle](https://docs.mesosphere.com/1.11/cli/command-reference/dcos-node/dcos-node-diagnostics-create/) analysis tool

## Installation

[Download binaries from the release page](https://github.com/adyatlov/bun/releases) or install from source:

```bash
$ go get github.com/adyatlov/bun/bun
```

## Usage

```bash
$ mkdir bundle
$ unzip bundle.zip -d bundle
$ cd bundle
bundle$ bun check
PROBLEM: dcos-version - Versions are different
Details:
master 172.20.0.22 has DC/OS version 1.11.0
master 172.20.0.24 has DC/OS version 1.11.0
agent 172.20.0.25 has DC/OS version 1.11.0
public agent 172.20.0.26 has DC/OS version 1.11.0
agent 172.20.0.27 has DC/OS version 1.11.0
agent 172.20.0.29 has DC/OS version 1.11.0
agent 172.20.0.21 has DC/OS version 1.10.1
master 172.20.0.23 has DC/OS version 1.11.0
agent 172.20.0.28 has DC/OS version 1.11.0

PROBLEM: health - Some DC/OS systemd units are not healthy.
Details:
172.20.0.21 dcos-docker-gc.service: health = 1

OK: node-count - Masters: 3, Agents: 5, Public Agents: 1, Total: 9
```

## Feedback

Please, report bugs and share your ideas for new features via the [issue page](https://github.com/adyatlov/bun/issues).

## Contributing

[Pull requests](https://github.com/adyatlov/bun/pulls) are welcome.
