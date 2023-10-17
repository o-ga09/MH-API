# API Gateway CLI

```bash
gcloud api-gateway apis create mh-api-gateway-v2
gcloud api-gateway api-configs create mh-api-config --api=mh-api-gateway-v2 --openapi-spec=./doc/openapi/apigateway.yml
gcloud api-gateway gateways create --api=mh-api-gateway-v2 --api-config=mh-api-config --location asia-northeast1 --project=mh-api-389212
gcloud api-gateway gateways update mh-api-v2 --project=mh-api-389212 --api-config=mh-api-config --api=mh-api-gateway-v2 --location=asia-northeast1
```