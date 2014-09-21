# Fileconsul
Fileconsul is sharing files(configuration file, Service/Check definition scripts, handler scripts and more) in a consul cluster.

## Usage
Run `fileconsul -h` to see the usage help:

```
$ fileconsul -h
NAME:
   fileconsul - Sharing files in a consul cluster.

USAGE:
   fileconsul [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   status   Show status of local files
   pull     Pull files from a consul cluster
   push     Push file to a consul cluster
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --version, -v    print the version
   --help, -h       show help
```

### Status Command
Run `fileconsul status -h` to see the usage help:

```
$ fileconsul status -h
NAME:
   status - Show status of local files

USAGE:
   command status [command options] [arguments...]

DESCRIPTION:
   Show the difference between local files and remote files that is stored in K/V Store of a consul cluster.

OPTIONS:
   --addr 'localhost:8500'  consul HTTP API address with port
   --dc 'dc1'               consul datacenter, uses local if blank
   --prefix 'fileconsul'    reading file status from Consul's K/V store with the given prefix
   --basepath '.'           base directory path of target files
```

### Pull Command
Run `fileconsul pull -h` to see the usage help:

```
$ fileconsul pull -h
NAME:
   pull - Pull files from a consul cluster

USAGE:
   command pull [command options] [arguments...]

DESCRIPTION:
   Pull remote files from K/V Store of a consul cluster.

OPTIONS:
   --addr 'localhost:8500'  consul HTTP API address with port
   --dc 'dc1'               consul datacenter, uses local if blank
   --prefix 'fileconsul'    reading file status from Consul's K/V store with the given prefix
   --basepath '.'           base directory path of target files
```

### Push Command
Run `fileconsul push -h` to see the usage help:

```
$ fileconsul push -h
NAME:
   push - Push file to a consul cluster

USAGE:
   command push [command options] [arguments...]

DESCRIPTION:
   Push remote files to K/V Store of a consul cluster.

OPTIONS:
   --addr 'localhost:8500'  consul HTTP API address with port
   --dc 'dc1'               consul datacenter, uses local if blank
   --prefix 'fileconsul'    reading file status from Consul's K/V store with the given prefix
   --basepath '.'           base directory path of target files
```

## Example
Use the demo environment for fileconsul
### Setup the demo environment used by Docker
```
$ git clone https://github.com/foostan/fileconsul.git
$ cd fileconsul/demo
$ docker build -t fileconsul .
```

### Run containers and a consul agent
#### Consul server

```
$ docker run -h server -i -t fileconsul
root@server:/# consul agent -data-dir=/tmp/consul -server -bootstrap-expect 1 &
==> WARNING: BootstrapExpect Mode is specified as 1; this is the same as Bootstrap mode.
==> WARNING: Bootstrap mode enabled! Do not enable unless necessary
==> WARNING: It is highly recommended to set GOMAXPROCS higher than 1
==> Starting Consul agent...
==> Starting Consul agent RPC...
==> Consul agent running!
         Node name: 'server'
        Datacenter: 'dc1'
            Server: true (bootstrap: true)
       Client Addr: 127.0.0.1 (HTTP: 8500, DNS: 8600, RPC: 8400)
      Cluster Addr: 172.17.0.6 (LAN: 8301, WAN: 8302)
    Gossip encrypt: false, RPC-TLS: false, TLS-Incoming: false

--- omitted below ---
```

#### Consul client

```
$ docker run -h client -i -t fileconsul
root@client:/# consul agent -data-dir=/tmp/consul -join 172.17.0.6 &
==> WARNING: It is highly recommended to set GOMAXPROCS higher than 1
==> Starting Consul agent...
==> Starting Consul agent RPC...
==> Joining cluster...
    Join completed. Synced with 1 initial agents
==> Consul agent running!
         Node name: 'client'
        Datacenter: 'dc1'
            Server: false (bootstrap: false)
       Client Addr: 127.0.0.1 (HTTP: 8500, DNS: 8600, RPC: 8400)
      Cluster Addr: 172.17.0.5 (LAN: 8301, WAN: 8302)
    Gossip encrypt: false, RPC-TLS: false, TLS-Incoming: false

--- omitted below ---
```

### Use fileconsul

#### Check status and push on server
```
root@server:/# cd /consul/share/
root@server:/consul/share# fileconsul status
Changes to be pushed:
  (use "fileconsul pull [command options]" to reset all files)
	new file:	bin/ntp
	new file:	bin/apache2
	new file:	bin/loadavg
	new file:	config/service/apache2.json
	new file:	config/service/ntp.json
	new file:	config/agent/client.json
	new file:	config/agent/server.json
	new file:	config/check/loadavg.json
root@server:/consul/share# fileconsul push
push new file:	bin/ntp
push new file:	bin/apache2
push new file:	bin/loadavg
push new file:	config/service/apache2.json
push new file:	config/service/ntp.json
push new file:	config/agent/client.json
push new file:	config/agent/server.json
push new file:	config/check/loadavg.json
```

#### Edit a file and push on client
```
root@server:/# cd /consul/share/
root@client:/consul/share# vi bin/apache2   # edit a file
root@client:/consul/share# fileconsul status
Changes to be pushed:
  (use "fileconsul pull [command options]" to reset all files)
	modified:	bin/apache2
root@client:/consul/share# fileconsul push
push modified file:	bin/apache2
```

#### Check status and pull on server
```
root@server:/consul/share# fileconsul status
Changes to be pushed:
  (use "fileconsul pull [command options]" to reset all files)
	modified:	bin/apache2
root@server:/consul/share# fileconsul pull
Synchronize remote files:
	modified:	bin/apache2
Already up-to-date.
```

## Contributing

1. Fork it ( https://github.com/[my-github-username]/fileconsul/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
