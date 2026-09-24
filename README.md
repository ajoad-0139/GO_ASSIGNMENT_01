# Rental Property API

A read-only REST API built with Go and Beego that serves rental property
data loaded from a local JSON file into memory at startup.

## Requirements

- Go 1.22+
- [bee tool](https://github.com/beego/bee) (optional, for `bee run`)

## Setup

```bash
git clone <https://github.com/ajoad-0139/GO_ASSIGNMENT_01.git>
cd GO_ASSIGNMENT_01
cd rental-property-api
go mod tidy
```

Make sure the data file exists at the path configured by `sourcepath` in
`conf/app.conf` (default: `rental_properties.json` in the data folder which is in the project root).

## Configuration

All runtime settings live in `conf/app.conf`:

```ini
appname = rental-property-api
httpport = 3000
runmode = dev
sourcepath = "data/rental_properties.json"
copyrequestbody = true
autorender = false
EnableDocs = true
```

## Run

```bash
bee run
```

or, without the bee tool:

```bash
go run main.go
```

The server starts on `http://localhost:3000` (or whatever `httpport` is set
to in `conf/app.conf`).

## Run Tests

```bash
cd services
go test ./... -v
go vet ./...
```

## Swagger Documentation

**Match port number as App.config**
```bash
    "http://localhost:3000/swagger/index.html"
```

## Sample curl Commands

**List all properties**
```bash
curl "http://localhost:3000/v1/properties"
```

**Filter by price range and property type**
```bash
curl "http://localhost:3000/v1/properties?min_price=50&max_price=150&property_type=Hotel"
```

**Filter by feed and published status**
```bash
curl "http://localhost:3000/v1/properties?feed=11&published=false"
```

**Filter by amenities (OR logic)**
```bash
curl "http://localhost:3000/v1/properties?amenities=Internet,Parking"
```

**Combined AND + amenities OR filter, with limit**
```bash
curl "http://localhost:3000/v1/properties?feed=11&amenities=Internet,Parking&limit=5"
```

**Get a single property by ID**
```bash
curl "http://localhost:3000/v1/properties/BC-1000001"
```

**Not found example**
```bash
curl "http://localhost:3000/v1/properties/does-not-exist"
```