# Go-Esperanto

A short program with RESTful API to randomly give an Esperanto word and translation. Mostly a proof of concept and a toy.

Phrasebook converted from Paul Denisowski's ESPDIC Project at <http://www.denisowski.org/Esperanto/ESPDIC/espdic.txt>

## Versioning

Currently at v2 of the module which has several breaking changes:

- Auth is required at all end points
- Removed the ability to add to the flat file list
- Converted from Gin to Chi

## Configuration

A phrasebook is included. If you want to use a different on, you will need to replace the contents in a similarly names file.

The following environment variables can be set:

- GO_EO_AUTHTOKEN - The authtoken to be used for authentication of the API POST endpoints; read as the "auth" key in either the
post body or the query string. If not specified, a random MD5 token will be generated on each startup
- GO_EO_API_PORT - The port to listen on, defaults to 8081
- GO_EO_PHRASEBOOK_DIR - The directory to find the phrasebook, (differs between development and docker)

Currently, only supports flat file storage; eventually may expand to other storage mechanisms.

## Running

`go build -mod=vendor .`

`GO_EO_API_PORT=8081 GO_EO_AUTHTOKEN=randomtokenforapi ./go-esperanto`

```bash
curl http://localhost:8081/

[{"esperanto":"Gxis la revido","english":"Goodbye"}, ....]
```

```bash
curl http://localhost:8081/random

{"esperanto":"Gxis la revido","english":"Goodbye"}
```

## Todo

- Add other data storage
