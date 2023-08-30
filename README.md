# Simplified pure Go secure shell (SSH) multi-host client

[![GoDoc](https://godoc.org/github.com/rwxrob/ssh?status.svg)](https://godoc.org/github.com/rwxrob/ssh)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
# Testable example

## Testable `runasany` example

This directory also contains a testable example implementation of `ssh.RunOnAny` as the command [`runonany`](cmd/runonany/main.go) that can be used to validate this specific functionality interactively. This example requires the `podman` container engine and command to be installed on the local system.

Here are the steps to conduct an interactive test. In this way, one can monitor the status of a given server and bring it up or down on demand and note how this affects the `ssh.Controller` (inside of `runonany`).

### A note about `build` script

The [`build`](build) bash script is designed to be used with simple bash tab completion. First add a `build` script like the following in your `PATH`.

```sh
#!/bin/sh
exec ./build "$@"
```

If you prefer you can do the above with `./build` instead and avoid the script. This just avoids the anti-pattern of adding `./` to your `PATH`.

Then add the following to `~/.bashrc`.

```sh
complete -C build build
```

### Build the SSH server image

The [testdata/server](testdata/server) directory contains an `ssh-server` (`ubuntu:latest`) image used to bring up SSH servers on different ports. This directory contains everything for the `ssh-client` container image.

```sh
build server
```

### Start up the three SSH server containers

Then start the three servers listening on ports `2221-2223`.

```sh
build start-ssh-servers
```

### Set the `RUNONANY_TARGET` environment variable

The containers all share the same underlying host (and IP address) but they don't know about it. We use the `RUNONANY_TARGET` environment variable to communicate this to the running containers.

```sh
export RUNONANY_TARGET=192.168.1.6
```

### Check the servers by running `ssh` from each client

Confirm server SSH connection is working with `ssh`. First you will need to note the host IP of the `podman` container engine and export it. (This can be obtained any number of ways.)

Now check the servers (or just do individual):

```sh
build check-servers
```

You should see the `ssh hostname` output of each command.

```
Warning: Permanently added '[192.168.1.6]:2221' (ED25519) to the list of known hosts.
ssh-server1
Warning: Permanently added '[192.168.1.6]:2222' (ED25519) to the list of known hosts.
ssh-server2
Warning: Permanently added '[192.168.1.6]:2223' (ED25519) to the list of known hosts.
ssh-server3
```

Note that the primary distinction is the port number. These servers all share the same `user` and credentials. They even share the same `authorized_hosts` key (which we ignore here deliberately for testing).

### Build and run `runonany` Go binary and `client` container

The [`runonany`](cmd/runonany/main.go) binary is a simple program that encapsulates an `ssh.Controller` configured in the [`runonany.yaml`](testdata/runonany.yaml) YAML file.

```sh
build client
build watch-client-runonany
```

This will update every two seconds.

## Interactively stop and start SSH server containers

The containerized `ssh-server` images can be stopped and started while monitoring the live status using commands similar to the following:

```sh
build stop-ssh-server 2
build start-ssh-server 2
build stop-ssh-servers
build start-ssh-servers
```

It is useful to do these commands from one TMUX pane while running `build watch-client-runonany` from another to see the change in `ssh.Controller.Clients` status.

Here are some things to validate:

* Random servers selected are between 1-3.
* Stop one server and note random selected no longer include.
* Stop two servers and not only single server selected.
* Stop all servers and not only the `ERR` section returned.
* Start one server after stopping and note recovery.
* Start two servers and note recovery.
* Start all servers and note recovery.
