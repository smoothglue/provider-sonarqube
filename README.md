# Provider SonarQube

`provider-sonarqube` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
SonarQube API.

## Getting Started

### Installation

Install the provider using the following `Provider` resource:

```shell
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-sonarqube
spec:
  package: ghcr.io/smoothglue/provider-sonarqube:v0.1.0
EOF
```

### Configuration

Each SonarQube instance you wish to manage will need a separate `ProviderConfig` defined.

Here is an example:
```yaml
apiVersion: sonarqube.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
---
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "host": "https://127.0.0.1:9000",
      "token": "deadbeef",
      "tls_insecure_skip_verify": true
    }
```

### Usage

There are examples of the various resources in the `examples` folder  in this repository.

You can see the API reference [here](https://doc.crds.dev/github.com/smoothglue/provider-sonarqube).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/smoothglue/provider-sonarqube/issues).
