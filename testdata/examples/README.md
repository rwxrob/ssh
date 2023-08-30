# Testable example

This directory contains a testable example implementation of `ssh.RunOnAny` that can be used to validate the package interactively. This example requires the `podman` container engine and command to be installed on the local system.

Here are the steps to conduct these interactive tests in real-time. In this way, one can monitor the status of a given server and bring it up or down on demand and note how this affects the `ssh.Controller` (inside of `runonany`).

## A note about `run` script

The [`run`](run) bash script is designed to be used with simple bash tab completion which can be easily enabled as follows:

```sh
complete -C ./run ./run
```

Some may prefer to add a `run` script/function/alias to avoid the `./` local directory.

## Start up all servers

The [server](server) directory contains an `ssh-server` (`ubuntu:latest`) image used to bring up SSH servers on different ports.

First build the client and server images.

```sh
./run build-images
```

Then start the three servers listening on ports `2221-2223`.

```sh
./run start-servers
```

## Check the servers with `ssh`

Confirm server SSH connection is working with `ssh`. First you will need to note the host IP of the `podman` container engine and export it. (This can be obtained any number of ways.)

```sh
export BUILDIP=192.168.1.6
```

Now check the servers (or just do individual):

```sh
./run check-servers
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

## Build the `runonany` Go binary and `client` container

The [`runonany`](client/runonany/main.go) binary is a simple program that encapsulates an `ssh.Controller` configured in the [`runonany.yaml`](client/runonany/runonany.yaml) YAML file.

The `runonany` can be run individually (installed on the local system) or from the [`client`](client) container, which automatically builds the latest and runs the `runonany` program with the default command `hostname` on a loop when started.

## Interactively stop and start SSH server containers

The containerized `ssh-server` images can be stopped and started while monitoring the live status.. To stop a server simply use `podman stop ssh-server{1,3}`. To start a server use the `./run start-server {1,3}` command.

Here are some things to validate:

* Random servers selected are between 1-3.
* Stop one server and note random selected no longer include.
* Stop two servers and not only single server selected.
* Stop all servers and not only the `ERR` section returned.
* Start one server after stopping and note recovery.
* Start two servers and note recovery.
* Start all servers and note recovery.
