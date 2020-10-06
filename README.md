# es-health-check

This is a small project to monitor a ElasticSearch cluster health check. Basically it just query the health check url and search by the string `green` from the body response.

## Pre req
In order to run this project you must have:

 - Golang
 - Serverless
 - AWS Credentials

## 🚧 Adjustments
Make sure to update your `serverless.yml`:

```bash
vpc:
    securityGroupIds:
      - sg-000000
    subnetIds:
      - subnet-00000
      - subnet-00000
```
OBS: This section is only required if you need to run in a specif VPC

```bash
ES_ENDPOINT: "http://your-elastic-search-url:9200/_cat/health"
```
Set to your ES url.


## 🚀 Deploy

Make the app build:

```bash
make build
```

Deploying to aws lambda:
```bash
make clean deploy
```

## TODO

 - Integrate with slack
 - Integrate with pagerduty
