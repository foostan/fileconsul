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

## Contributing

1. Fork it ( https://github.com/[my-github-username]/fileconsul/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
