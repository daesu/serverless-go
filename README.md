# serverless-skeleton

Skeleton monorepo focused on go lambda jobs on AWS.

This repo is architected to contain a number of sub-projects that are responsible
for handling and processing various messages/events used within the Hub Platform Service.

The following table outlines the sub-projects:

| Name | Directory  | Purpose                      |
| :--- | :--------- | :--------------------------- |
| api  | [api](api) | Creates an ApiGateway on AWS |

## Setup

Create a `.env` file by copying the contents of `.env.example` and adding values.

You can setup the local development environment with the following command:

```
make setup-dev
```

Build all sub-projects and move any binaries to `bin`

```
make build
```

Provision all required AWS resources, create lambda functions and upload binaries.

```
make deploy
```
