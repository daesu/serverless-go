FROM node:14.3.0-alpine AS serverless_go
RUN npm install -g serverless
WORKDIR /opt/app

COPY --from=golang:1.16.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"