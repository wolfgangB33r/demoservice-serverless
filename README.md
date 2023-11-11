# Dockerized Demoservice

A dockerized tracing demo service built for running in serverless context

## Set the callee config

```bash
export DEMO_SERVICE_CALLEES='{ "Callees" : [ { "Adr" : "http://www.example.com", "Count" : 2 }, { "Adr" : "http://www.orf.at", "Count" : 1 } ] }'
echo $DEMO_SERVICE_CALLEES
```