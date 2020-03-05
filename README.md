# Apache Beam Golang UDF

Run UDFs (User Defined Functions) on Apache Beam Golang SDK.

Go Modules activation (if necessary):
```sh
export GO111MODULE=on
```

CSV parse example:
```sh
# Direct
go run examples/parse/csv/csv.go

# Direct with internet files
go run examples/parse/csv/csv.go --location="http://localhost:8081/"

# Direct with GCS files
go run examples/parse/csv/csv.go --location=gs://apache-beam-golang-udf

# Dataflow with GCS files
go run examples/parse/csv/csv.go --location=gs://apache-beam-golang-udf \
    --max_num_workers=1 \
    --num_workers=1 \
    --project=marcelo-henrique-neppel \
    --runner=dataflow \
    --staging_location=gs://apache-beam-golang-udf/bin \
    --temp_location=gs://apache-beam-golang-udf/temp \
    --worker_harness_container_image=apachebeam/go_sdk:latest \
    --worker_machine_type=n1-standard-1

# Flink with GCS files
./gradlew :runners:flink:1.9:job-server:runShadow # For using the embedded cluster

./gradlew :runners:flink:1.9:job-server:runShadow -PflinkMasterUrl=localhost:8081 # For using a separate cluster

go run examples/parse/csv/csv.go --location=gs://apache-beam-golang-udf \
    --endpoint=localhost:8099 \
    --runner=flink
```

JSON parse example:
```sh
# Direct
go run examples/parse/json/json.go

# Direct with internet files
go run examples/parse/json/json.go --location="http://localhost:8081/"

# Direct with GCS files
go run examples/parse/json/json.go --location=gs://apache-beam-golang-udf

# Dataflow with GCS files
go run examples/parse/json/json.go --location=gs://apache-beam-golang-udf \
    --max_num_workers=1 \
    --num_workers=1 \
    --project=marcelo-henrique-neppel \
    --runner=dataflow \
    --staging_location=gs://apache-beam-golang-udf/bin \
    --temp_location=gs://apache-beam-golang-udf/temp \
    --worker_harness_container_image=apachebeam/go_sdk:latest \
    --worker_machine_type=n1-standard-1

# Flink with GCS files
./gradlew :runners:flink:1.9:job-server:runShadow # For using the embedded cluster

./gradlew :runners:flink:1.9:job-server:runShadow -PflinkMasterUrl=localhost:8081 # For using a separate cluster

go run examples/parse/json/json.go --location=gs://apache-beam-golang-udf \
    --endpoint=localhost:8099 \
    --runner=flink
```

XML parse example:
```sh
# Direct
go run examples/parse/xml/xml.go

# Direct with internet files
go run examples/parse/xml/xml.go --location="http://localhost:8081/"

# Direct with GCS files
go run examples/parse/xml/xml.go --location=gs://apache-beam-golang-udf

# Dataflow with GCS files
go run examples/parse/xml/xml.go --location=gs://apache-beam-golang-udf \
    --max_num_workers=1 \
    --num_workers=1 \
    --project=marcelo-henrique-neppel \
    --runner=dataflow \
    --staging_location=gs://apache-beam-golang-udf/bin \
    --temp_location=gs://apache-beam-golang-udf/temp \
    --worker_harness_container_image=apachebeam/go_sdk:latest \
    --worker_machine_type=n1-standard-1

# Flink with GCS files
./gradlew :runners:flink:1.9:job-server:runShadow # For using the embedded cluster

./gradlew :runners:flink:1.9:job-server:runShadow -PflinkMasterUrl=localhost:8081 # For using a separate cluster

go run examples/parse/xml/xml.go --location=gs://apache-beam-golang-udf \
    --endpoint=localhost:8099 \
    --runner=flink
```

On examples using GCS files, please upload the example files to one of your buckets first.
