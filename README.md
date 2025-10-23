# unique-ips

`unique-ips` is a HTTP server that you can send log lines and it will parse them and count the number of unique IPs found.

## 📌 Description
The HTTP server will run on:
- Port 5000 with the endpoint /logs - this will receive the log line to be processed
- Port 9102 with the endpoint /metrics - will expose the Prometheus counter `unique_ip_addresses_total`

## 🚀 Installation
```sh
# Clone the repository
$ git clone https://github.com/paulofeitor/unique-ips.git
$ cd unique-ips

# Run the server
$ go run main.go
```

## 📖 Usage
Start the service and push logs (port 5000)
```
POST /logs
{"timestamp": "2020-06-24T15:27:00.123456Z", "ip": "83.150.59.250", "url": ... }
```

Scrape metrics manually or through Prometheus (port 9102)
```
GET /metrics
```

## 📊 To be considered
- **Async mode** - to improve requests handling, we could perhaps add an async mode
- **Regex mode** - instead of parsing/binding the full payload, fetch what we need without converting into struct
- **Basic HTTP metrics** - maybe would be benefitial to understand performance issues

## 🛠 Technologies Used
- **Go** for CLI development
- **Echo** for HTTP server framework
- **Prometheus** for Prometheus library
- **Gobench** for HTTP Benchmarks
