# go-esperanto
A short program with RESTful API to randomly give an Esperanto word and translation.

Phrasebook converted from Paul Denisowski's ESPDIC Project at http://www.denisowski.org/Esperanto/ESPDIC/espdic.txt


The following environment variables can be set:
- GO_EO_AUTHTOKEN - The authtoken to be used for authentication of the API POST endpoints; read as the "auth" key in either the
post body or the query string. If not specified, a random MD5 token will be generated on each startup
- GO_EO_API_PORT - The port to listen on, defaults to 8081
- GIN_MODE=release - Sets API into release mode (uses Gin for API layer)
- GO_EO_PHRASEBOOK_DIR - The directory to find the phrasebook, (differs between development and docker)

Currently, only supports flat file storage; eventually will expand to other storage mechanisms

# Running
GO_EO_API_PORT=8081 GO_EO_AUTHTOKEN=randomtokenforapi GIN_MODE=release go run main.go

curl http://localhost:8081/

[{"esperanto":"Gxis la revido","english":"Goodbye"}, ....] 

curl http://localhost:8081/random

{"esperanto":"Gxis la revido","english":"Goodbye"}

curl http://localhost:8081/ -d "esperanto=viro&english=man&auth=randomtokenforapi"

{"esperanto":"viro","english":"man"}

# Testing
GIN_MODE=release go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html

# Todo
- Add other data storage
- Add docker support