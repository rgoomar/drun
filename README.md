# drun (Docker Runner)

## About

This is a simple CLI to run a container and selective commands with Docker.

The primary purpose for this tool is to help create a consistent development experience.

## Installation

TBD

## Configuration

You can add a `drun.json` file in the root of your project or wherever you going to run it.

*Config Options*

`image` - Specify Docker image

`defaultCommand` - Default command to run

*Example*
```
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
```
{
    "image": "openjdk:8",
    "defaultCommand": "./gradlew build"
}
```

You can run:
```
drun
```

### Run with image specified in configuration only

With this configuration:
```
{
    "image": "openjdk:8",
    "defaultCommand": "./gradlew build"
}
```

You can run:
```
drun ./gradlew build
```

### Run without configuration

```
drun --image node:8-alpine node --version
```
