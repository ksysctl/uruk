# Uruk
Example of Go project.

# Getting started

## Requirements
- [Docker](https://docs.docker.com/engine/installation/)
- [Docker compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install)


Set up Go paths.
```bash
mkdir -p ~/Projects/go/
mkdir -p ~/Library/go/
```

~/.bash_profile
```bash
# Set up GOROOT
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

# Set up GOPATH for libraries 
export GOPATH=~/Library/go
export PATH=$PATH:$GOPATH/bin

# Set up GOPATH for projects
export GOPATH=$GOPATH:~/Projects/go
```

Set up project.
```bash
cd ~/Projects/go/
mkdir -p ~/Projects/go/github.com/ksysctl/uruk
git clone git@github.com:ksysctl/uruk.git github.com/ksysctl/uruk
 ```

## Development

Run project build and bring services up.
```bash
make build
```

Run rebuild project.
```bash
make rebuild
```

Bring services up.
```bash
make up
```

Bring services down.
```bash
make down
```

Show the services logs.
```bash
make logs
```

Clean and format source code.
```bash
make clean
```

Update dependencies.
```bash
make deps
```