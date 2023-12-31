# MTA Hosting Optimizer Service Documentation

## Problem Statement

Currently, there are about 35 physical servers hosting 482 mail transfer agents (MTAs), each with a dedicated public IP address. To optimize the utilization of servers and reduce costs, we need to identify and uncover inefficient servers hosting only a few active MTAs.

## Solution Overview

The MTA Hosting Optimizer Service is designed to analyze server information stored in Redis, identify servers with few active MTAs (as per a configurable threshold), and return the results.

## How to get started?
This structured format allows users to expand each section for detailed instructions as needed.

Follow these steps to set up and run the project:


<details>
  <summary><strong>Step 1: Clone the Project</strong></summary>
Clone the project repository to your local machine.

```bash
git clone <repository-url>
```
</details>
<details>
  <summary><strong>Step 2: Start Docker Compose</strong></summary>
Navigate to the project directory and start the Docker Compose services in detached mode.

```bash
cd <project-directory>
docker-compose up -d
```
</details>
<details>
  <summary><strong>Step 3: Create an Environment File</strong></summary>
Create an environment file (.env) from the provided example (.env.example) and configure it with your environment-specific settings.

```bash
cp .env.example .env
# Edit the .env file to set your configurations
```
</details>
<details>
  <summary><strong>Step 4: Run the Application</strong></summary>
Execute the main application using Go.

```bash
go run main.go
```
</details>
<details>
  <summary><strong>Step 5: Check Service Health</strong></summary>
Check the service health by making a GET request to /healthCheckup. If the response is "OKAY!", the service is running perfectly.
</details>
<details>
  <summary><strong>Step 6: Set Threshold (if needed)</strong></summary>
If a threshold value is not specified in your .env file, set it to your desired value.
</details>
<details>
  <summary><strong>Step 7: Hit /hostname</strong></summary>
Make a GET request to /hostname to get the desired result.
</details>

## Implementation Details

### Configuration

The service uses an environment variable `THRESHOLD` to set the maximum allowed active MTAs on a server. If this variable is not set, it defaults to a threshold of 1.

```go
func GetThresholdFromEnv() int {
    // Retrieve the threshold value from the environment variable THRESHOLD
    // Default value is 1 if THRESHOLD is not set
}
```

### Data Retrieval
The service retrieves server information from Redis using the model.GetServerInformation() function.

```go
ipConfigs, err := model.GetServerInformation()
```

### Processing
The service processes the server information, counts active MTAs per server, and identifies servers with few active MTAs, based on the threshold.

```go
for i, ipConfig := range ipConfigs {
    // Process server information
}

var result []Result
for hostname, activeIP := range activeIPs {
    if len(activeIP) <= threshold {
        result = append(result, Result{Hostname: hostname})
    }
}
```

### Execution
The service execution is triggered by a client's API request. It calculates the optimization results and returns them.

```go
func GetActiveMTAsAboveThreshold() ([]Result, error) {
    // Entry point for service execution
}
```

### API Endpoint
The service exposes an HTTP/REST endpoint to retrieve hostnames having less than or equal to the configured active IP addresses.

- Endpoint: /hostnames
- HTTP Method: GET
- Parameters: None
- Response Format: JSON
- Sample Request: GET http://localhost:8080/hostnames
- Sample response:

```json
[
    {"Hostname": "mta-prod-1"},
    {"Hostname": "mta-prod-3"}
]
```
### Usage

1. Configure the threshold value by setting the THRESHOLD environment variable.
2. Start the service.
3. Make a GET request to /hostnames to retrieve the inefficient servers.

## Performance Testing Result

### Hardware Information

- CPU: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz

### Benchmark Results

#### KeyDB Set Benchmark

- Benchmark Name: BenchmarkKeyDBSet-12
- Number of Operations: 1
- Average Time per Operation: 2,020,434,800 ns/op
- Average Memory Allocation: 82,320 bytes/op
- Average Number of Allocations: 500 allocs/op

#### Redis Set Benchmark

- Benchmark Name: BenchmarkRedisSet-12
- Number of Operations: 1
- Average Time per Operation: 2,004,819,600 ns/op
- Average Memory Allocation: 69,632 bytes/op
- Average Number of Allocations: 429 allocs/op

### Performance Results

Based on the performance testing results, the KeyDB and Redis benchmarks have been executed. KeyDB shows a slightly higher average time per operation compared to Redis. Memory allocation and the number of allocations are also slightly higher in the KeyDB benchmark.

Please note that these benchmarks provide insights into the performance of KeyDB and Redis in your specific use case. The choice between KeyDB and Redis may depend on your application's requirements and trade-offs between performance and resource usage.

For further analysis or specific recommendations, please consult with your development and infrastructure teams.


## Conclusion
The MTA Hosting Optimizer Service efficiently identifies inefficient servers hosting only a few active MTAs, helping optimize server utilization and reduce costs.