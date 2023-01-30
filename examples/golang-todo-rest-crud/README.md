# ToDo API proof of concept

[![Go Report Card](https://goreportcard.com/badge/github.com/chelodoz/todo-api-golang)](https://goreportcard.com/report/github.com/chelodoz/todo-api-golang)
[![ci](https://github.com/chelodoz/todo-api-golang/actions/workflows/ci.yml/badge.svg)](https://github.com/chelodoz/todo-api-golang/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/chelodoz/todo-api-golang/branch/main/graph/badge.svg)](https://codecov.io/gh/chelodoz/todo-api-golang)

- [Introduction](#introduction)
- [Environment](#environment)
- [Docker](#docker)
- [Run](#run)
- [Make File](#make-file)
  - [Unit test coverage](#unit-test-coverage)
  - [Integration tests](#integration-test)
  - [Swagger](#swagger)
  - [Generate mocks](#generate-mocks)
- [Graceful shutdown](#graceful-shutdown)
- [Pre commit](#pre-commit)
- [Project Layout](#project-layout)
- [API Definition](#api-definition)
  - [Create Note](#create-note)
    - [Create Note Request](#create-note-request)
    - [Create Note Response](#create-note-response)
  - [Get Note](#get-note)
    - [Get Note Request](#get-note-request)
    - [Get Note Response](#get-note-response)
    - [Get Notes Request](#get-notes-request)
    - [Get Notes Response](#get-notes-response)
  - [Update Note](#update-note)
    - [Update Note Request](#update-note-request)
    - [Update Note Response](#update-note-response)
- [Logging](#logging)
  - [Log levels](#log-levels)
  - [Log body](#log-body)
  - [Example of Create Note Request](#example-of-create-note-request )
  - [Example of Create Note Request Logging](#example-of-create-note-request-logging)
  - [Example of Create Note Response Logging](#example-of-create-note-response-logging)
  - [Example of Fatal level error when .env file is missing](#example-of-fatal-level-error-when-env-file-is-missing)
  - [Example of Error level error when 500 error occurs](#example-of-error-level-error-when-500-error-occurs)
- [Rate Limiting](#rate-limiting)
  - [Get Health Rate Limit](#get-health-rate-limit)
    - [Get Health Request Rate Limit](#get-health-request-rate-limit)
    - [Get Health Response Rate Limit](#get-health-response-rate-limit)
- [Trace](#trace)
  - [Get Health Trace](#get-health-trace)
    - [Get Health Request Without Trace Id](#get-health-request-without-trace-id)
    - [Get Health Response With Trace Id](#get-health-response-with-trace-id)
    - [Get Health Request With Trace Id](#get-health-request-with-trace-id)
    - [Get Health Response With System Trace Id](#get-health-response-with-system-trace-id)
- [Go cover view](#go-cover-view)

## Introduction

Welcome! 👋

The end goal of this project is to make a simple proof of concept of a RESTful API with Go using gorilla/mux.

If you're have not encountered Go before, you should visit this website [here](https://golang.org/doc/install).

This project was developed using go v1.9.4

## Environment

For the handling of environment variables and the reading of .env/.json/.yaml files, the [Viper]("https://github.com/spf13/viper") library was used.

The `.env.example` file is included in the root directory to provide development environment variables change it to `.env` to make it work.

## Docker

For the creation of the mongo db and mongodb express in docker the following repository was taken as a reference [NaN Labs Devops reference](https://github.com/nanlabs/devops-reference/tree/main/examples/docker/mongodb).

## Run

To run the code, you will need docker and docker-compose installed on your machine. In the project root, run `docker compose up --build -d` or `make dcbuild` if you have make file installed to create and start all the containers..

You can run it `manually without docker` using the command `go run ./cmd/todo` or `make run`, to make it work, the environment variable `MONGO_HOST=localhost` must be changed in the .env file to `localhost` instead of the mongo container name.

Note that if you plan to run all in containers the env variable must be the container name `MONGO_HOST=mongodb`

After that, you have a RESTful API that is running at `http://127.0.0.1:8080`.

## Make File

The [Make File]("https://linuxhint.com/install-make-ubuntu/") library was used to make a list of useful commands.

- `make build`                  run go build
- `make run`                    run go run
- `make unittest`               run go unit tests with coverage
- `make integrationtest`        run go integration tests package
- `make check_swagger_install`  define a dependency to install go-swagger cli
- `make swagger`                run go generate to generate swagger .yaml and json
- `make check_mockery_install`  define a dependency to install mockery cli
- `make mocks`                  run go generate to generate mocks of interfaces with go generate tag
- `make dcbuild`                run all the containers

## Unit test coverage

Use the command `make unittest` to run all the unit tests including coverage

## Integration test

In order to perform integrations tests using real requests the following package was used [Dockertest]("https://github.com/ory/dockertest").
The idea behind this is to create two containers one for the api and the other for mongo db completely separate from the development containers, run the integration tests by making calls to the test api and then delete both containers.

Use the command `make integrationtest` to run all the integration tests

## Swagger

For the api documentation [Go swagger]("https://goswagger.io/") was the choice, using design first approach the documentation can be generated through code notes.

Run `go install github.com/go-swagger/go-swagger/cmd/swagger` to install mockery CLI.

Use the command `make swagger` to generate the /docs/swagger.yaml and third_party/swagger-ui-4.11.1/swagger.json files from the go-swagger models.

## Generate mocks

For generating mocks [Mockery]("https://github.com/vektra/mockery") package was used.

Run `go install github.com/vektra/mockery/v2@latest` to install mockery CLI.

Use the command `make mocks` to generate the mocks of the interfaces in /internal/todo/note folder.

## Graceful shutdown

A `graceful shutdown in a process` is when the OS (operating system) can safely shut down its processes and close all connections, taking as much time as needed.

To be able to achieve that, one has to listen to [Termination signals]("https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html") that are sent to the application by the process manager, and act accordingly. A delay of 30 seconds was implemented at the moment of listening for a termination signal in order to shut down the server.

## Pre commit

To maintain high code quality we opted to use [Pre commit]("https://pre-commit.com/") which allows to run hooks to automatically point out problems in code such as missing semicolons, trailing whitespace, and debug statements. It can also be configured to run tests, linter, dependency checking and other commands.

It can be installed using python running `pip install pre-commit`

Check the `.pre-commit-config.yaml` file to see the hooks

More details on [Pre Commit Golang]("https://github.com/dnephin/pre-commit-golang")

## Project Layout

The project uses the following project layout:

```text
.
├── cmd                main applications
│   └── todo             api server setup
├── docs               api documentation
├── test               non-unit tests
│   └── integration      integration tests
├── internal           private application and library code
│   ├── config           configuration library
│   ├── platform         provide support for databases, authentication
│   │     └── mongo         mongo client
│   ├── ratelimit        api rate limiting
│   ├── todo             todo related features
│   │     └── note          note related features
│   └── trace          package for generating request and trace ids
├── pkg                public library code
│   ├── apierror         standard api errors
│   ├── encode           encode and decode helpers
│   ├── health           health check definition
│   └── logs             logs setup
└── third_party        third party libraries
     └── swagger-ui        static files from swagger ui

```

The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout) and [Package Oriented Design]("https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html").

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example,
the `todo` directory contains the application logic related with the todo feature.

Within each feature package, code are organized in layers (handlers, service, repository), following the dependency guidelines as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## API Definition

### Create Note

#### Create Note Request

```js
POST api/v1/notes
```

```json
{
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
}
 ```

#### Create Note Response

```js
201 Created
```

```json
{
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
    "status":"To Do",
}
 ```

### Update Note

#### Update Note Request

 ```js
PATCH /api/v1/notes/{noteId}
```

```json
{
    "name": "Go shopping",
    "description":"Buy groceries for the week",
    "status" : "In Progress"
}
```

#### Update Note Response

```js
204 No Content
```

### Get Note

#### Get Note Request

```js
GET /api/v1/notes/{noteId}
```

#### Get Note Response

```js
200 Ok
```

```json
{
    "name": "Go shopping",
    "description":"Buy groceries for the week",
    "status" : "In Progress"
}
```

#### Get Notes Request

```js
GET api/v1/notes
```

#### Get Notes Response

```js
200 Ok
```

```json
[
  {
    "name": "Go shopping",
    "description":"Buy groceries for the week",
    "status" : "To Do"
  },
  {
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
    "status" : "In Progress"
  },
]
```

### Get Health

#### Get Health Request

```js
GET /api/v1/health
```

#### Get Health Response

```js
200 Ok
```

```json
{
    "status" : "Healthy"
}
```

### Get Swagger UI

#### Get Swagger UI Request

```js
GET /api/v1/swagger/
```

## Logging

Logging plays a vital role in identifying issues, evaluating performances, and knowing the process status within the application. For this reason it was decided to select structured logging, using Uber's library [Zap](https://github.com/uber-go/zap).

### Log levels

For simplicity it was decided to use three types of log levels

- `Info`  Generally useful information to log.
- `Error` Anything that can potentially cause application oddities including 50x server errors.
- `Fatal` Any error that is forcing a shutdown of the service or application to prevent data loss.

### Log body

The log itself is a json structure with the following keys

- `level`       define log level
- `ts`          define log timestamp
- `caller`      define the file where the log was called
- `msg`         define the main info of the log
- `method`      define the request method
- `url`         define the resource called in the api
- `statusCode`  define the response status code
- `duration`    define the duration of the request in nanoseconds
- `detail`     define extra error information
- `stacktrace`  define extra trace information

Logging was included at the beginning and end of the requests in order to maintain traceability.

### Example of Create Note Request

```js
POST api/v1/notes
```

```json
{
    "name": "Go to the bank",
    "description":"Schedule an appointment to the bank",
}
 ```

### Example of Create Note Request Logging

```json
{
    "level": "info",
    "ts": 1674960805.3700233,
    "caller": "todo/middleware.go:64",
    "msg": "Start http request",
    "method": "POST",
    "url": "/api/v1/notes"
}
 ```

### Example of Create Note Response Logging

```json
{
    "level": "info",
    "ts": 1674960805.3718014,
    "caller": "todo/middleware.go:87",
    "msg": "Finish http request",
    "method": "POST",
    "url": "/api/v1/notes",
    "statusCode": 201,
    "duration": 0.0017783
}
 ```

### Example of Fatal level error when .env file is missing

```json
{
    "level": "fatal",
    "ts": 1674963617.935728,
    "caller": "todo/main.go:29",
    "msg": "Cannot load config",
    "detail": "Config File \".env\" Not Found in \"[C:\\\\Users\\\\User\\\\Documents\\\\todo-api-golang C:\\\\Users\\\\User\\\\Documents\\\\todo-api-golang\\\\cmd\\\\todo]\"",
    "stacktrace": "main.main\n\tC:/Users/User/Documents/todo-api-golang/cmd/todo/main.go:29\nruntime.main\n\tC:/Program Files/Go/src/runtime/proc.go:250"
}
 ```

### Example of Error level error when 500 error occurs

In this case a body of the response is included to provide additional information in order to identify the issue.

```json
{
    "level": "error",
    "ts": 1674965353.375559,
    "caller": "todo/middleware.go:84",
    "msg": "Finish http request",
    "method": "POST",
    "url": "/api/v1/notes",
    "statusCode": 500,
    "duration": 0.0005219,
    "body": "{\"type\":\"INTERNAL\",\"message\":\"Internal server error.\",\"code\":500,\"detail\":\"error creating note id\"}\n",
    "stacktrace": "todo-api-golang/internal/todo.LogMiddleware.func1.1\n\tC:/Users/Chelo/Documents/todo-api-golang/internal/todo/middleware.go:84\nnet/http.HandlerFunc.ServeHTTP\n\tC:/Program Files/Go/src/net/http/server.go:2109\ngithub.com/gorilla/mux.(*Router).ServeHTTP\n\tC:/Users/Chelo/go/pkg/mod/github.com/gorilla/mux@v1.8.0/mux.go:210\nnet/http.serverHandler.ServeHTTP\n\tC:/Program Files/Go/src/net/http/server.go:2947\nnet/http.(*conn).serve\n\tC:/Program Files/Go/src/net/http/server.go:1991"
}
 ```

## Rate Limiting

Rate limiting is a technique used to control the number of requests a user can make to an API over a given period of time. Using the library [Toolbooth]("https://github.com/didip/tollbooth") which provides a simple API to perform rate limiting.

It can be configured by the following client variables

- `HTTP_RATE_LIMIT=3`
- `HTTP_RATE_INTERVAL=second`
- `INTEGRATION_HTTP_RATE_LIMIT=100`
- `INTEGRATION_HTTP_RATE_INTERVAL=minute`

The interval has the values as listed below

- `second` default value
- `minute`
- `hour`

### Get Health Rate Limit

#### Get Health Request Rate Limit

```js
GET /api/v1/health
```

#### Get Health Response Rate Limit

```js
429 Too Many Requests
```

```text
HTTP/1.1 429 Too Many Requests
Content-Type: application/json
Ratelimit-Limit: 3
Ratelimit-Remaining: 0
Ratelimit-Reset: 1
X-Rate-Limit-Duration: 1
X-Rate-Limit-Limit: 3.00
X-Rate-Limit-Request-Remote-Addr: 172.25.0.1:54948
```

```json
{
  "type":"TOO_MANY_REQUEST",
  "message":"You have reached maximum request limit.",
  "code":429
}
```

## Trace

To maintain the traceability of a request it is always useful to use unique identifiers when generating logs and calling other services. That is why the use of `X-Request-Id` and `X-Trace-Id` was implemented, when entering a request to the api will automatically generate a unique identifier `X-Request-Id`, then it will be verified in the header if there is any trace identifier `X-Trace-Id`, if there is one it will be kept the same if not it will be assigned the `X-Request-Id`, both identifiers will be transmitted in the context of the request, and both will be returned in the request response as headers.

### Get Health Trace

#### Get Health Request Without Trace Id

```js
GET /api/v1/health
```

#### Get Health Response With Trace Id

```js
200 Ok
```

```text
X-Request-Id: 3d54f2f9-4418-4e50-90c6-d5209dc25d5d
X-Trace-Id: 3d54f2f9-4418-4e50-90c6-d5209dc25d5d
```

```json
{
    "status" : "Healthy"
}
```

#### Get Health Request With Trace Id

```js
GET /api/v1/health
```

```text
X-Trace-Id: trace id from other system
```

#### Get Health Response With System Trace Id

```js
200 Ok
```

```text
X-Request-Id: 3d54f2f9-4418-4e50-90c6-d5209dc25d5d
X-Trace-Id: trace id from other system
```

```json
{
    "status" : "Healthy"
}
```

## Go cover view

<details> <summary> todo-api-golang/internal/todo/note/handler.go </summary>

```
    1: // note package include application logic related with the note sub feature.
    2: package note
    3:
    4: import (
    5:  "errors"
    6:  "net/http"
    7:  "todo-api-golang/pkg/apierror"
    8:  "todo-api-golang/pkg/encode"
    9:
   10:  "github.com/google/uuid"
   11:
   12:  "github.com/go-playground/validator"
   13: )
   14:
   15: // Handler is a common interface to perform operations related to notes and http.
   16: type Handler interface {
   17:  Create(rw http.ResponseWriter, r *http.Request)
   18:  GetById(rw http.ResponseWriter, r *http.Request)
   19:  GetAll(rw http.ResponseWriter, r *http.Request)
   20:  Update(rw http.ResponseWriter, r *http.Request)
   21: }
   22:
   23: // handler is the implementation of the operations related to notes.
   24: type handler struct {
   25:  service  Service
   26:  validate *validator.Validate
   27: }
   28:
   29: // NewHandler creates a note handler which have operations related to notes.
O  30: func NewHandler(service Service, validator *validator.Validate) Handler {
O  31:  return &handler{
O  32:          service:  service,
O  33:          validate: validator,
O  34:  }
O  35: }
   36:
   37: // swagger:route POST /notes Notes CreateNoteRequestWrapper
   38: // Creates a new note
   39: //
   40: // Create a new note in a database
   41: //
   42: // responses:
   43: // 201: CreateNoteResponse
   44: // 400: ValidationErrorResponseWrapper
   45: // 422: ErrorResponseWrapper
   46: // 500: ErrorResponseWrapper
   47:
   48: // Create handles POST requests and create a note into the data store.
O  49: func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
O  50:  var createNoteRequest CreateNoteRequest
O  51:
O  52:  if err := encode.ReadRequestBody(r, &createNoteRequest); err != nil {
O  53:          encode.WriteError(rw, apierror.NewUnprocessableEntity())
O  54:          return
O  55:  }
   56:
O  57:  if err := h.validate.Struct(&createNoteRequest); err != nil {
O  58:          encode.WriteError(rw, apierror.NewValidationBadRequest(err.(validator.ValidationErrors)))
O  59:          return
O  60:  }
   61:
O  62:  newNote := &Note{
O  63:          Name:        createNoteRequest.Name,
O  64:          Description: createNoteRequest.Description,
O  65:  }
O  66:
O  67:  note, err := h.service.Create(newNote, r.Context())
O  68:
X  69:  if err != nil {
X  70:          encode.WriteError(rw, apierror.NewInternal(err.Error()))
X  71:          return
X  72:  }
   73:
O  74:  noteResponse := &CreateNoteResponse{
O  75:          ID:          note.ID.String(),
O  76:          Name:        note.Name,
O  77:          Description: note.Description,
O  78:          Status:      note.Status,
O  79:  }
O  80:
O  81:  encode.WriteResponse(rw, http.StatusCreated, noteResponse)
   82: }
   83:
   84: // swagger:route GET /notes Notes Notes
   85: // Returns a list of notes
   86: //
   87: // Returns a list of notes from the database
   88: // responses:
   89: // 200: GetNotesResponse
   90: // 500: ErrorResponseWrapper
   91:
   92: // GetAll handles GET requests and returns all the notes from the data store.
O  93: func (h *handler) GetAll(rw http.ResponseWriter, r *http.Request) {
O  94:  notes, err := h.service.GetAll(r.Context())
O  95:
O  96:  if err != nil {
O  97:          switch {
O  98:          case errors.Is(err, ErrFoundingNote):
O  99:                  encode.WriteResponse(rw, http.StatusOK, &GetNotesResponse{})
O 100:          default:
O 101:                  encode.WriteError(rw, apierror.NewInternal(err.Error()))
  102:          }
O 103:          return
  104:  }
  105:
O 106:  var notesResponse GetNotesResponse
O 107:
O 108:  for _, note := range notes {
O 109:          noteResponse := GetNoteResponse{
O 110:                  ID:          note.ID.String(),
O 111:                  Name:        note.Name,
O 112:                  Description: note.Description,
O 113:                  Status:      note.Status,
O 114:          }
O 115:          notesResponse = append(notesResponse, noteResponse)
O 116:  }
  117:
O 118:  encode.WriteResponse(rw, http.StatusOK, &notesResponse)
  119: }
  120:
  121: // swagger:route GET /notes/{noteId} Notes NoteIdQueryParamWrapper
  122: // Returns a single note
  123: //
  124: // Returns a single note from the database
  125: // responses:
  126: // 200: GetNoteResponse
  127: // 500: ErrorResponseWrapper
  128:
  129: // GetNote handles GET/{noteId} requests and returns a note from the data store.
O 130: func (h *handler) GetById(rw http.ResponseWriter, r *http.Request) {
O 131:  noteId, err := encode.GetUriParam(r, "noteId")
X 132:  if err != nil {
X 133:          encode.WriteError(rw, apierror.NewBadRequest(ErrInvalidNoteId.Error()))
X 134:          return
X 135:  }
  136:
O 137:  uid, err := uuid.Parse(noteId)
O 138:  if err != nil {
O 139:          encode.WriteError(rw, apierror.NewBadRequest(ErrInvalidNoteId.Error()))
O 140:          return
O 141:  }
  142:
O 143:  note, err := h.service.GetById(uid, r.Context())
O 144:
O 145:  if err != nil {
O 146:          switch {
O 147:          case errors.Is(err, ErrFoundingNote):
O 148:                  encode.WriteError(rw, apierror.NewNotFound())
O 149:          default:
O 150:                  encode.WriteError(rw, apierror.NewInternal(err.Error()))
  151:          }
O 152:          return
  153:  }
  154:
O 155:  noteResponse := &GetNoteResponse{
O 156:          ID:          note.ID.String(),
O 157:          Name:        note.Name,
O 158:          Description: note.Description,
O 159:          Status:      note.Status,
O 160:  }
O 161:
O 162:  encode.WriteResponse(rw, http.StatusOK, noteResponse)
  163: }
  164:
  165: // swagger:route PATCH /notes/{noteId} Notes UpdateNoteRequestWrapper
  166: // Update an existing note
  167: //
  168: // Update a new note in a database
  169: //
  170: // responses:
  171: // 204: NoContentResponseWrapper
  172: // 400: ValidationErrorResponseWrapper
  173: // 422: ErrorResponseWrapper
  174: // 500: ErrorResponseWrapper
  175:
  176: // Update handles PATCH requests and updates a note into the data store.
O 177: func (h *handler) Update(rw http.ResponseWriter, r *http.Request) {
O 178:  var updateNoteRequest UpdateNoteRequest
O 179:  noteId, err := encode.GetUriParam(r, "noteId")
X 180:  if err != nil {
X 181:          encode.WriteError(rw, apierror.NewBadRequest(ErrInvalidNoteId.Error()))
X 182:          return
X 183:  }
  184:
O 185:  uid, err := uuid.Parse(noteId)
O 186:  if err != nil {
O 187:          encode.WriteError(rw, apierror.NewBadRequest(ErrInvalidNoteId.Error()))
O 188:          return
O 189:  }
  190:
O 191:  if err := encode.ReadRequestBody(r, &updateNoteRequest); err != nil {
O 192:          encode.WriteError(rw, apierror.NewUnprocessableEntity())
O 193:          return
O 194:  }
  195:
O 196:  if err := h.validate.Struct(&updateNoteRequest); err != nil {
O 197:          encode.WriteError(rw, apierror.NewValidationBadRequest(err.(validator.ValidationErrors)))
O 198:          return
O 199:  }
  200:
O 201:  updatedNote := &Note{
O 202:          ID:          uid,
O 203:          Name:        updateNoteRequest.Name,
O 204:          Description: updateNoteRequest.Description,
O 205:          Status:      updateNoteRequest.Status,
O 206:  }
O 207:
O 208:  _, err = h.service.Update(updatedNote, r.Context())
O 209:
O 210:  if err != nil {
O 211:          switch {
O 212:          case errors.Is(err, ErrFoundingNote):
O 213:                  encode.WriteError(rw, apierror.NewNotFound())
O 214:          default:
O 215:                  encode.WriteError(rw, apierror.NewInternal(err.Error()))
  216:          }
O 217:          return
  218:  }
  219:
O 220:  rw.WriteHeader(http.StatusNoContent)
  221: }

```

</details>

<details> <summary> todo-api-golang/internal/todo/note/note.go </summary>

```
   1: // note package include application logic related with the note sub feature.
   2: package note
   3:
   4: import (
   5:   "errors"
   6:   "time"
   7:
   8:   "github.com/google/uuid"
   9: )
  10:
  11: // Note defines the fields of a note.
  12: type Note struct {
  13:   ID          uuid.UUID `bson:"_id,omitempty"`
  14:   Name        string    `bson:"name,omitempty"`
  15:   Description string    `bson:"description,omitempty"`
  16:   Status      Status    `bson:"status,omitempty"`
  17:   CreatedAt   time.Time `bson:"createdAt,omitempty"`
  18:   UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
  19: }
  20:
  21: // List of all possible errors managed by the note package.
  22: var (
  23:   ErrInvalidNoteId  = errors.New("error invalid id")
  24:   ErrCreatingNoteId = errors.New("error creating note id")
  25:   ErrDecodingNote   = errors.New("error decoding note")
  26:   ErrUpdatingNote   = errors.New("error updating note")
  27:   ErrCreatingNote   = errors.New("error creating note")
  28:   ErrFoundingNote   = errors.New("error founding note")
  29: )
  30:
  31: // swagger:enum Status
  32: // State is a type that defines all possible states of a note.
  33: type Status string
  34:
  35: const (
  36:   Todo       Status = "To Do"
  37:   InProgress Status = "In Progress"
  38:   Done       Status = "Done"
  39: )
  40:
  41: // IsValid defines when a status is a valid value.
O 42: func (s Status) IsValid() bool {
O 43:   switch s {
O 44:   case Todo, InProgress, Done:
O 45:           return true
O 46:   default:
O 47:           return false
  48:   }
  49: }

```

</details>

<details> <summary> todo-api-golang/internal/todo/note/repository.go </summary>

```
    1: // note package include application logic related with the note sub feature.
    2: package note
    3:
    4: import (
    5:  "context"
    6:  "time"
    7:  "todo-api-golang/internal/config"
    8:
    9:  "github.com/google/uuid"
   10:
   11:  "go.mongodb.org/mongo-driver/bson"
   12:  "go.mongodb.org/mongo-driver/mongo"
   13:  "go.mongodb.org/mongo-driver/mongo/options"
   14:
   15:  cmongo "todo-api-golang/internal/platform/mongo"
   16: )
   17:
   18: // Repository is a common interface to perform operations related to notes and infrastructure layer.
   19: //
   20: //go:generate mockery --name=Repository --output=note --inpackage=true --filename=repository_mock.go
   21: type Repository interface {
   22:  Create(note *Note, ctx context.Context) (*Note, error)
   23:  GetById(id uuid.UUID, ctx context.Context) (*Note, error)
   24:  GetAll(ctx context.Context) ([]Note, error)
   25:  Update(note *Note, ctx context.Context) (*Note, error)
   26: }
   27:
   28: // repository represents the repository used for interacting with notes records.
   29: type repository struct {
   30:  client cmongo.ClientHelper
   31:  config *config.Config
   32: }
   33:
   34: // NewRepository instantiates the note repository.
O  35: func NewRepository(client cmongo.ClientHelper, config *config.Config) Repository {
O  36:  return &repository{
O  37:          client: client,
O  38:          config: config,
O  39:  }
O  40: }
   41:
   42: // getCollection retrieve a mongo collection.
O  43: func (r *repository) getCollection() cmongo.CollectionHelper {
O  44:  database := r.client.Database(r.config.MongoDatabase)
O  45:
O  46:  return database.Collection(r.config.MongoCollection)
O  47: }
   48:
   49: // Create inserts a new note record.
O  50: func (r *repository) Create(note *Note, ctx context.Context) (*Note, error) {
O  51:
O  52:  collection := r.getCollection()
O  53:
O  54:  id, err := uuid.NewRandom()
X  55:  if err != nil {
X  56:          return nil, ErrCreatingNoteId
X  57:  }
O  58:  note.ID = id
O  59:  note.CreatedAt = time.Now().UTC()
O  60:  note.Status = "To Do"
O  61:
O  62:  _, err = collection.InsertOne(ctx, note)
O  63:
X  64:  if err != nil {
X  65:          return nil, ErrCreatingNote
X  66:  }
   67:
O  68:  return note, nil
   69: }
   70:
   71: // GetById returns the requested note by searching its id.
O  72: func (r *repository) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
O  73:  var note Note
O  74:
O  75:  collection := r.getCollection()
O  76:  filter := bson.M{"_id": id}
O  77:
O  78:  err := collection.FindOne(ctx, filter).Decode(&note)
X  79:  if err != nil {
X  80:          return nil, ErrFoundingNote
X  81:  }
   82:
O  83:  return &note, nil
   84: }
   85:
   86: // GetAll returns all notes.
X  87: func (r *repository) GetAll(ctx context.Context) ([]Note, error) {
X  88:  var notes []Note
X  89:
X  90:  findOptions := options.Find()
X  91:  findOptions.SetLimit(100)
X  92:
X  93:  collection := r.getCollection()
X  94:
X  95:  cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
X  96:  if err != nil {
X  97:          return nil, ErrFoundingNote
X  98:  }
   99:
X 100:  for cur.Next(ctx) {
X 101:          var note Note
X 102:          if err := cur.Decode(&note); err != nil {
X 103:                  return nil, ErrDecodingNote
X 104:          }
  105:
X 106:          notes = append(notes, note)
  107:  }
X 108:  cur.Close(ctx)
X 109:
X 110:  if notes == nil {
X 111:          return nil, ErrFoundingNote
X 112:  }
  113:
X 114:  return notes, nil
  115: }
  116:
  117: // Update updates the existing record with new value.
O 118: func (r *repository) Update(note *Note, ctx context.Context) (*Note, error) {
O 119:  collection := r.getCollection()
O 120:  filter := bson.M{"_id": note.ID}
O 121:
O 122:  note.UpdatedAt = time.Now().UTC()
O 123:
O 124:  update := bson.M{
O 125:          "$set": note,
O 126:  }
O 127:
O 128:  result, err := collection.UpdateOne(ctx, filter, update)
O 129:
X 130:  if result.(*mongo.UpdateResult).MatchedCount == 0 {
X 131:          return nil, ErrFoundingNote
X 132:  }
  133:
X 134:  if err != nil {
X 135:          return nil, ErrUpdatingNote
X 136:  }
  137:
O 138:  return note, nil
  139: }

```

</details>

<details> <summary> todo-api-golang/internal/todo/note/repository_mock.go </summary>

```
    1: // Code generated by mockery v2.16.0. DO NOT EDIT.
    2:
    3: package note
    4:
    5: import (
    6:  context "context"
    7:
    8:  uuid "github.com/google/uuid"
    9:  mock "github.com/stretchr/testify/mock"
   10: )
   11:
   12: // MockRepository is an autogenerated mock type for the Repository type
   13: type MockRepository struct {
   14:  mock.Mock
   15: }
   16:
   17: // Create provides a mock function with given fields: note, ctx
O  18: func (_m *MockRepository) Create(note *Note, ctx context.Context) (*Note, error) {
O  19:  ret := _m.Called(note, ctx)
O  20:
O  21:  var r0 *Note
X  22:  if rf, ok := ret.Get(0).(func(*Note, context.Context) *Note); ok {
X  23:          r0 = rf(note, ctx)
O  24:  } else {
O  25:          if ret.Get(0) != nil {
O  26:                  r0 = ret.Get(0).(*Note)
O  27:          }
   28:  }
   29:
O  30:  var r1 error
X  31:  if rf, ok := ret.Get(1).(func(*Note, context.Context) error); ok {
X  32:          r1 = rf(note, ctx)
O  33:  } else {
O  34:          r1 = ret.Error(1)
O  35:  }
   36:
O  37:  return r0, r1
   38: }
   39:
   40: // GetAll provides a mock function with given fields: ctx
O  41: func (_m *MockRepository) GetAll(ctx context.Context) ([]Note, error) {
O  42:  ret := _m.Called(ctx)
O  43:
O  44:  var r0 []Note
X  45:  if rf, ok := ret.Get(0).(func(context.Context) []Note); ok {
X  46:          r0 = rf(ctx)
O  47:  } else {
O  48:          if ret.Get(0) != nil {
O  49:                  r0 = ret.Get(0).([]Note)
O  50:          }
   51:  }
   52:
O  53:  var r1 error
X  54:  if rf, ok := ret.Get(1).(func(context.Context) error); ok {
X  55:          r1 = rf(ctx)
O  56:  } else {
O  57:          r1 = ret.Error(1)
O  58:  }
   59:
O  60:  return r0, r1
   61: }
   62:
   63: // GetById provides a mock function with given fields: id, ctx
O  64: func (_m *MockRepository) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
O  65:  ret := _m.Called(id, ctx)
O  66:
O  67:  var r0 *Note
X  68:  if rf, ok := ret.Get(0).(func(uuid.UUID, context.Context) *Note); ok {
X  69:          r0 = rf(id, ctx)
O  70:  } else {
O  71:          if ret.Get(0) != nil {
O  72:                  r0 = ret.Get(0).(*Note)
O  73:          }
   74:  }
   75:
O  76:  var r1 error
X  77:  if rf, ok := ret.Get(1).(func(uuid.UUID, context.Context) error); ok {
X  78:          r1 = rf(id, ctx)
O  79:  } else {
O  80:          r1 = ret.Error(1)
O  81:  }
   82:
O  83:  return r0, r1
   84: }
   85:
   86: // Update provides a mock function with given fields: note, ctx
O  87: func (_m *MockRepository) Update(note *Note, ctx context.Context) (*Note, error) {
O  88:  ret := _m.Called(note, ctx)
O  89:
O  90:  var r0 *Note
X  91:  if rf, ok := ret.Get(0).(func(*Note, context.Context) *Note); ok {
X  92:          r0 = rf(note, ctx)
O  93:  } else {
O  94:          if ret.Get(0) != nil {
O  95:                  r0 = ret.Get(0).(*Note)
O  96:          }
   97:  }
   98:
O  99:  var r1 error
X 100:  if rf, ok := ret.Get(1).(func(*Note, context.Context) error); ok {
X 101:          r1 = rf(note, ctx)
O 102:  } else {
O 103:          r1 = ret.Error(1)
O 104:  }
  105:
O 106:  return r0, r1
  107: }
  108:
  109: type mockConstructorTestingTNewMockRepository interface {
  110:  mock.TestingT
  111:  Cleanup(func())
  112: }
  113:
  114: // NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
O 115: func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
O 116:  mock := &MockRepository{}
O 117:  mock.Mock.Test(t)
O 118:
O 119:  t.Cleanup(func() { mock.AssertExpectations(t) })
  120:
O 121:  return mock
  122: }

```

</details>

<details> <summary> todo-api-golang/internal/todo/note/service.go </summary>

```
   1: // note package include application logic related with the note sub feature.
   2: package note
   3:
   4: import (
   5:   "context"
   6:
   7:   "github.com/google/uuid"
   8: )
   9:
  10: // Service is a common interface to perform operations related to notes.
  11: //
  12: //go:generate mockery --name=Service --output=note --inpackage=true --filename=service_mock.go
  13: type Service interface {
  14:   Create(note *Note, ctx context.Context) (*Note, error)
  15:   GetById(id uuid.UUID, ctx context.Context) (*Note, error)
  16:   GetAll(ctx context.Context) ([]Note, error)
  17:   Update(note *Note, ctx context.Context) (*Note, error)
  18: }
  19:
  20: // service represents the service used for interacting with notes.
  21: type service struct {
  22:   repository Repository
  23: }
  24:
  25: // NewService instantiates the note service.
O 26: func NewService(repository Repository) Service {
O 27:   return &service{
O 28:           repository: repository,
O 29:   }
O 30: }
  31:
  32: // Create handles create note business logic.
O 33: func (s *service) Create(note *Note, ctx context.Context) (*Note, error) {
O 34:   newNote, err := s.repository.Create(note, ctx)
O 35:
O 36:   if err != nil {
O 37:           return nil, err
O 38:   }
  39:
O 40:   return newNote, nil
  41: }
  42:
  43: // Create handles the note update business logic.
O 44: func (s *service) Update(note *Note, ctx context.Context) (*Note, error) {
O 45:   updatedNote, err := s.repository.Update(note, ctx)
O 46:
O 47:   if err != nil {
O 48:           return nil, err
O 49:   }
  50:
O 51:   return updatedNote, nil
  52: }
  53:
  54: // GetById handles the business logic requesting a note.
O 55: func (s *service) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
O 56:
O 57:   note, err := s.repository.GetById(id, ctx)
O 58:
O 59:   if err != nil {
O 60:           return nil, err
O 61:   }
  62:
O 63:   return note, nil
  64: }
  65:
  66: // GetAll handles the business logic requesting multiple notes.
O 67: func (s *service) GetAll(ctx context.Context) ([]Note, error) {
O 68:
O 69:   notes, err := s.repository.GetAll(ctx)
O 70:
O 71:   if err != nil {
O 72:           return nil, err
O 73:   }
  74:
O 75:   return notes, nil
  76: }

```

</details>

<details> <summary> todo-api-golang/internal/todo/note/service_mock.go </summary>

```
    1: // Code generated by mockery v2.16.0. DO NOT EDIT.
    2:
    3: package note
    4:
    5: import (
    6:  context "context"
    7:
    8:  uuid "github.com/google/uuid"
    9:  mock "github.com/stretchr/testify/mock"
   10: )
   11:
   12: // MockService is an autogenerated mock type for the Service type
   13: type MockService struct {
   14:  mock.Mock
   15: }
   16:
   17: // Create provides a mock function with given fields: note, ctx
O  18: func (_m *MockService) Create(note *Note, ctx context.Context) (*Note, error) {
O  19:  ret := _m.Called(note, ctx)
O  20:
O  21:  var r0 *Note
X  22:  if rf, ok := ret.Get(0).(func(*Note, context.Context) *Note); ok {
X  23:          r0 = rf(note, ctx)
O  24:  } else {
O  25:          if ret.Get(0) != nil {
O  26:                  r0 = ret.Get(0).(*Note)
O  27:          }
   28:  }
   29:
O  30:  var r1 error
X  31:  if rf, ok := ret.Get(1).(func(*Note, context.Context) error); ok {
X  32:          r1 = rf(note, ctx)
O  33:  } else {
O  34:          r1 = ret.Error(1)
O  35:  }
   36:
O  37:  return r0, r1
   38: }
   39:
   40: // GetAll provides a mock function with given fields: ctx
O  41: func (_m *MockService) GetAll(ctx context.Context) ([]Note, error) {
O  42:  ret := _m.Called(ctx)
O  43:
O  44:  var r0 []Note
X  45:  if rf, ok := ret.Get(0).(func(context.Context) []Note); ok {
X  46:          r0 = rf(ctx)
O  47:  } else {
O  48:          if ret.Get(0) != nil {
O  49:                  r0 = ret.Get(0).([]Note)
O  50:          }
   51:  }
   52:
O  53:  var r1 error
X  54:  if rf, ok := ret.Get(1).(func(context.Context) error); ok {
X  55:          r1 = rf(ctx)
O  56:  } else {
O  57:          r1 = ret.Error(1)
O  58:  }
   59:
O  60:  return r0, r1
   61: }
   62:
   63: // GetById provides a mock function with given fields: id, ctx
O  64: func (_m *MockService) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
O  65:  ret := _m.Called(id, ctx)
O  66:
O  67:  var r0 *Note
X  68:  if rf, ok := ret.Get(0).(func(uuid.UUID, context.Context) *Note); ok {
X  69:          r0 = rf(id, ctx)
O  70:  } else {
O  71:          if ret.Get(0) != nil {
O  72:                  r0 = ret.Get(0).(*Note)
O  73:          }
   74:  }
   75:
O  76:  var r1 error
X  77:  if rf, ok := ret.Get(1).(func(uuid.UUID, context.Context) error); ok {
X  78:          r1 = rf(id, ctx)
O  79:  } else {
O  80:          r1 = ret.Error(1)
O  81:  }
   82:
O  83:  return r0, r1
   84: }
   85:
   86: // Update provides a mock function with given fields: note, ctx
O  87: func (_m *MockService) Update(note *Note, ctx context.Context) (*Note, error) {
O  88:  ret := _m.Called(note, ctx)
O  89:
O  90:  var r0 *Note
X  91:  if rf, ok := ret.Get(0).(func(*Note, context.Context) *Note); ok {
X  92:          r0 = rf(note, ctx)
O  93:  } else {
O  94:          if ret.Get(0) != nil {
O  95:                  r0 = ret.Get(0).(*Note)
O  96:          }
   97:  }
   98:
O  99:  var r1 error
X 100:  if rf, ok := ret.Get(1).(func(*Note, context.Context) error); ok {
X 101:          r1 = rf(note, ctx)
O 102:  } else {
O 103:          r1 = ret.Error(1)
O 104:  }
  105:
O 106:  return r0, r1
  107: }
  108:
  109: type mockConstructorTestingTNewMockService interface {
  110:  mock.TestingT
  111:  Cleanup(func())
  112: }
  113:
  114: // NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to
assert the mocks expectations.
O 115: func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
O 116:  mock := &MockService{}
O 117:  mock.Mock.Test(t)
O 118:
O 119:  t.Cleanup(func() { mock.AssertExpectations(t) })
  120:
O 121:  return mock
  122: }

```

</details>

<details> <summary> todo-api-golang/internal/todo/note/validator.go </summary>

```
   1: // note package include application logic related with the note sub feature.
   2: package note
   3:
   4: import "github.com/go-playground/validator"
   5:
   6: // CustomValidator is a common interface for custom validators.
   7: type CustomValidator interface {
   8:   IsValid() bool
   9: }
  10:
  11: // ValidateEnum is a custom validator to validate if the input is within the list of valid values.
O 12: func ValidateEnum(fl validator.FieldLevel) bool {
O 13:   value := fl.Field().Interface().(CustomValidator)
O 14:   return value.IsValid()
O 15: }

```

</details>

<details> <summary> todo-api-golang/pkg/health/health.go </summary>

```
   1: // health package include a reusable health check handler.
   2: package health
   3:
   4: import (
   5:   "log"
   6:   "net/http"
   7: )
   8:
   9: type HealthResponse struct {
  10:   // example: Healthy
  11:   Status string `json:"status"`
  12: }
  13:
  14: // swagger:route GET /health Health Health
  15: // Check health of the api
  16: //
  17: // Check health of the api
  18: //
  19: // responses:
  20: // 200: HealthResponseWrapper
  21:
  22: // HealthCheck return a Healthy message in the response.
O 23: func HealthCheck(w http.ResponseWriter, r *http.Request) {
O 24:   w.Header().Set("Content-Type", "application/json")
O 25:   w.WriteHeader(http.StatusOK)
O 26:   _, err := w.Write([]byte(`{"status":"Healthy"}`))
X 27:   if err != nil {
X 28:           log.Printf("Write failed: %v", err)
X 29:   }
  30: }
  31:
  32: // Returns Healthy if the api is working
  33: // swagger:response HealthResponseWrapper
  34: type HealthResponseWrapper struct {
  35:   // in: body
  36:   Body HealthResponse
  37: }

```

</details>
