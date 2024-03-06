# Go-Esperanto

A short program with RESTful API to randomly give an Esperanto word and translation. Mostly a proof of concept and a toy.

Esperanto is the world's most widely spoken constructed international auxiliary language intended to be a universal second language for international communication.

Phrasebook converted from Paul Denisowski's ESPDIC Project at <http://www.denisowski.org/Esperanto/ESPDIC/espdic.txt>

## Versioning

Currently at v2 of the module which has several breaking changes:

- Auth is required at all end points
- Removed the ability to add to the flat file list
- Converted from Gin to Chi

## Configuration

A phrasebook is included. If you want to use a different one, you will need to replace the contents in a similarly names file.

The following environment variables can be set:

- GO_EO_AUTHTOKEN - The authtoken to be used for authentication of the API POST endpoints; read as the "auth" key in either the
post body or the query string. If not specified, a random MD5 token will be generated on each startup
- GO_EO_API_PORT - The port to listen on, defaults to 8081
- GO_EO_PHRASEBOOK_DIR - The directory to find the phrasebook, (differs between development and docker)

Currently, only supports flat file storage; eventually may expand to other storage mechanisms.

## Running

We have several different options for setting up and running the repo. We offer a [Task](https://taskfile.dev/) file and a Make file. Task is a cross-platform task runner. Make is often already available in some systems, so may be easier.

### Task

We use Task for all of our actions:

```bash
% task
task: Available tasks for this project:
* build:              Build the local Go image
* docker-build:       Build Docker image
* docker-push:        Pushes the Docker image
* docker-run:         Run the Docker image; will build it
* run:                Run the generated binary locally
* test:               Runt the tests
* vendor:             Updates the vendor directory
```

## Usage

Get entire dictionary:

```bash
curl -H "X-API-TOKEN:randomtokenforapi" http://localhost:8081/

[{"esperanto":"Gxis la revido","english":"Goodbye"}, ....]
```

Get random phrase:

```bash
curl -H "X-API-TOKEN:randomtokenforapi" http://localhost:8081/random

{"esperanto":"Gxis la revido","english":"Goodbye"}
```

## Todo

- Add other data storage
