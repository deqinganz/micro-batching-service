# Micro-Batching Service

This is a simple micro-batching service that uses Micro-Batching library to process jobs in batches, and exposes a REST API to interact with the library.

## Build and Run

```bash
go build .
go run .
```

## Usage

Configurations can be set in `config.json` file, can also be updated via endpoints when the program runs

### Add Job

To add a job to the queue:

```bash 
curl http://localhost:8080/job -H "Content-Type:application/json" -d '{
    "type": "UPDATE_USER_INFO",
    "name": "update user name to John",
    "params": {
        "userId": "123",
        "name": "John"
    }
}'
```

or

```bash
curl http://localhost:8080/job -H "Content-Type:application/json" -d '{
    "type": "BALANCE_UPDATE",
    "name": "user1 to $50",
    "params": {
        "userId": "1",
        "amount": 50
    }
}'
```

### Set Frequency

To call BatchProcessor every 5 seconds via the `/batch-frequency` endpoint:

```bash
curl http://localhost:8080/batch-frequency -H "Content-Type:application/json" -d '{"frequency":10}'
```

### Get Frequency

To get the current frequency of BatchProcessor:

```bash
curl http://localhost:8080/batch-frequency
```

### Set Batch Size

To set the batch size of BatchProcessor via the `/batch-size` endpoint:

```bash
curl http://localhost:8080/batch-size -H "Content-Type:application/json" -d '{"batch-size":10}'
```

### Get Batch Size

To get the current batch size of BatchProcessor:

```bash
curl http://localhost:8080/batch-size
```

### Set on/off for Preprocessing

To turn on preprocessing via the `/preprocess` endpoint:

```bash
curl http://localhost:8080/preprocess -H "Content-Type:application/json" -d '{"preprocessing":true}'
```