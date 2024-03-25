# Toledo

This repository provides a simple GitOps boilerplate to help you learn how to configure resources on Kubernetes using GitOps practices. It includes sample Kubernetes configs and dummy applications to experiment with certain Cloud Native concepts and tools, including Cilium, Flux.

This playground is ideal for those who are new to GitOps (and Kubernetes networking) and want to explore its functionalities in a hands-on environment.

## GitOps

We're using Flux CD (https://fluxcd.io/) to sync the following two categories of resources on Kubernetes:

1. **Infrastructure**: All cluster-wide resources, such as cert-manager, Cilium, or ingress-nginx.
2. **Apps**: All configs related to our actual workloads.

You can find these configs under `./k8s`.

## Get started (local)

### 1. Create cluster

**Prerequisites:** Make sure you have `kind` installed before proceeding. You can find installation instructions on the kind website (https://kind.io).

To start, create a kind cluster with the configs stored in the `kind-cluster-config.yaml` file:

```sh
kind create cluster --config kind-cluster-config.yaml --name toledo-local
```

The command above creates a cluster with 3 worker nodes and **no** CNI plugin installed.

### 2. Install Cilium

Once the cluster is created, we need to [install Cilium using Helm](https://docs.cilium.io/en/stable/installation/k8s-install-helm/) (to be later managed using Flux CD).

### 3. Bootstrap Flux CD

You can customize the Flux CD bootstrap process based on your specific needs by following the full instructions provided by Flux CD (https://fluxcd.io/flux/cmd/flux_bootstrap/).

Once the cluster has been created and bootstrapepd, you can experiment with the included dummy applications or add additional infrastructure resources.
