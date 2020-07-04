# Pricetopus
![](https://github.com/ay0o/pricetopus/workflows/.github/workflows/ci.yaml/badge.svg)

Pricetopus checks whether the price of a product is lower or equal to what you're expecting, and if so, notifies you by email.

Check `internal/parser/selector.go` for the supported sites.

## Configuration
Before use, you need to set the following environment variables:

- PRICETOPUS_EMAIL_SERVER
- PRICETOPUS_EMAIL_SERVER_PORT
- PRICETOPUS_EMAIL_USER
- PRICETOPUS_EMAIL_PASSWORD
- PRICETOPUS_EMAIL_TO
- PRICETOPUS_PRODUCT_URL
- PRICETOPUS_PRODUCT_PRICE

## Usage

You can either build from source and run, or use the container available at [Docker Hub](https://cloud.docker.com/repository/docker/ay0o/pricetopus).

### AWS ECS with Fargate
A Terraform module is provided in `deployments/terraform` to launch the service in AWS ECS using Fargate.

### Kubernetes
A job and a cronjob are provided as well, in case you prefer to launch pricetopus in a Kubernetes cluster.

## Contributing

If you would like to add a new site to Pricetopus, check the content of `internal/parser/selector.go` and add the site to the selectors map.
