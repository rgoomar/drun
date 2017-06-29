# drun (Docker Runner)

[![Build Status](https://travis-ci.org/rgoomar/drun.svg?branch=master)](https://travis-ci.org/rgoomar/drun)

## About

This is a simple CLI to run a container and selective commands with Docker.

The primary purpose for this tool is to help create a consistent development experience.

## Notes

* Currently this will mount the directory you are executing from to `/app` on the container
    * This is useful for building projects locally
* This is my first Go project, so there's room for improvement

## Installation

### Standalone Binary

1. Download binary from [releases](https://github.com/rgoomar/drun/releases/)
1. Move the binary to `/usr/local/bin/drun`
    - `mv drun-[arch] /usr/loca/bin/drun`

### Building with Go

```bash
go get github.com/rgoomar/drun
cd $GOPATH/src/github.com/rgoomar/drun
# if you want to use $GOPATH/bin
go install
# if you want it to be local to the directory
go build
```

## Configuration

You can add a `drun.json` file in the root of your project or wherever you going to run it.

*Config Options*

`image` - Specify Docker image

`defaultCommand` - Default command to run

*Example*
```json
{
    "image": "openjdk:8",
    "defaultCommand": "./gradlew build"
}
```

## Usage

### Flags

`--image` - Specify an image to run

`[COMMANDS]` - Commands to execute in container

### Run based on configuration defined

With this configuration:
```json
{
    "image": "openjdk:8",
    "defaultCommand": "./gradlew build"
}
```

You can run:
```bash
drun
```

### Run with image specified in configuration only

With this configuration:
```bash
{
    "image": "openjdk:8",
    "defaultCommand": "./gradlew build"
}
```

You can run:
```bash
drun ./gradlew build
```

### Run without configuration

```bash
drun --image node:8-alpine node --version
```
