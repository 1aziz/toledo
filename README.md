# Toledo

This repository provides a simple GitOps boilerplate to help you get started with configuring resources on Kubernetes using GitOps practices. It includes sample Kubernetes configurations and a demo `halo-toledo` application to experiment with core Cloud Native concepts and tools, including:

- **Cilium**: Explore service mesh functionalities like service discovery and load balancing.
- **Flux**: Learn how to use GitOps for automating deployments and managing your Kubernetes cluster configuration.

This playground is ideal for those who are new to GitOps and Kubernetes networking and want to explore these functionalities in a hands-on environment.

In addition, in this repo, we're using/experimenting the following technologies:

- **Taskfile**: Defines tasks for bootstrapping Kubernetes clusters.

By working through this repository, you'll gain practical experience with GitOps principles and essential Cloud Native tools.

## Get started (local)

### 1. Bootstrap cluster

To setup and configure a local cluster, you'll first need to make sure you've installed [`Taskfile`](https://taskfile.dev/).

Then, you can run the following command to bootstrap a new local cluster:

```sh
task bootstrap
```

The command above ensures `kind` ensures:

1. `kind` and `helm` are installed on your machine.
2. Creates a new Kind cluster using the configurations declared in `kind-cluster-config.yaml` in the repo.
3. Installs Cilium and Flux controllers
4. Boostraps GitOps to sync all the configs under `k8s` directory

Once you've created and bootstrapped your cluster, you can start experimenting! Here's how:

- **Explore the `halo-toledo` Demo App**: This sample application provides a starting point for understanding how to deploy applications using GitOps. Dive into its source code (`./halo-toledo/src`) and Kubernetes configurations (`./halo-toledo/deploy`) to see how it's structured and deployed.

- **Bring Your Own Applications**: This repository is designed to be extended! You can fork it and add your own applications. To do this:
  - Create a new directory inside the `./apps` folder, for example: `./apps/my-new-app`
  - Inside your new app directory, create two subfolders: - `src`: This folder will contain your application's source code. - `deploy`: This folder will hold your Kubernetes deployment configurations (e.g., Deployments, Services).
    Benefits of Customization:

By adding your own applications, you can:

- Practice deploying and managing real-world applications with GitOps.
- Experiment with different configurations and tools within the GitOps workflow.
- Tailor the playground to your specific learning goals.

## GitOps

We're using Flux CD (https://fluxcd.io/) to sync the following two categories of resources on Kubernetes:

1. **Infrastructure**: All cluster-wide resources, such as cert-manager, Cilium, or ingress-nginx.
2. **Apps**: All configs related to our actual workloads.

You can find these configs under `./k8s`.
