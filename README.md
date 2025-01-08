# Toledo

Toledo is a simple GitOps boilerplate designed to help you get started with configuring Kubernetes resources using GitOps practices. This repository provides sample Kubernetes configurations and includes a demo application, `halo-toledo`, designed for experimenting with fundamental Cloud Native concepts and tools."

## Key Features

- **Hands-on GitOps Learning**: Discover how to automate deployments and manage Kubernetes cluster configurations using **Flux**.
- **Task Automation**: Simplify cluster bootstrapping with **Taskfile**, a task runner that standardizes and automates common setup processes.

### Included Tools and Technologies

1. **Flux**
   Master GitOps by automating deployments and managing your Kubernetes cluster configuration directly from a Git repository.

2. **Taskfile**
   Simplify cluster bootstrapping and streamline repetitive tasks with predefined commands. In future, we can add more tasks, for example to support new types of clusters.

## Who Is This For?

This repository is ideal for:

- Beginners to GitOps and Kubernetes who want a practical, hands-on playground.
- Those exploring Kubernetes networking and Cloud Native tools.
- Those looking for a starting point to integrate GitOps practices into their workflows.

## Get Started (Local)

### Prerequisites

Before you proceed with the steps below, ensure you have the following tools installed and configured on your system:

1. **Flux CD**

   - Flux is a set of continuous and progressive delivery solutions for Kubernetes.
   - [Installation Guide](https://fluxcd.io/docs/installation/)

2. **Taskfile**

   - A task runner, like `make`, for easily managing and automating commands.
   - [Installation Guide](https://taskfile.dev/#/installation)

3. **Kind**
   - Kind (Kubernetes IN Docker) is a tool for running local Kubernetes clusters using Docker.
   - [Installation Guide](https://kind.sigs.k8s.io/docs/user/quick-start/)

### Bootstrap the Cluster

To set up and configure a _local_ Kubernetes cluster, run the bootstrap script `sh bootstrap` and follow the steps!

This command performs the following actions listed in the `Taskfile`:

```
kind:bootstrap
├── kind:create
│   ├── Deletes an existing Kind cluster (`toledo-local`) if it exists
│   ├── Creates a new Kind cluster using `kind-cluster-config.yaml`
│   ├── Verifies that the cluster context is set correctly
│   └── Creates a Docker registry credential secret (`regcred`) for pulling images
└── bootstrap
    └── Configures Flux for GitHub:
        │── Installs Flux CD into the cluster
        ├── Authenticates with `GITHUB_TOKEN`
        ├── Configures Flux to sync from a GitHub repository
        └── Uses the specified branch, repository, and cluster configuration path
```

## Explore the Repository

### **`halo-toledo` Demo App**

- **Location**:

  - Source code: `./halo-toledo/src`
  - Kubernetes configurations: `./halo-toledo/deploy`

- **What to Do**:
  - Review the app's source code and deployment configurations to understand how applications are deployed using GitOps.
  - Experiment with deploying and updating the app within your local cluster.

### **Bring Your Own Applications**

This repository is designed to be extended with your own applications. You can fork the repo and tailor it to your specific learning or project need (e.g. to build a home lab).

- **Steps**:

  1. Create a new directory under the `./apps` folder:
     Example: `./apps/my-new-app`
  2. Inside your new app directory, create the following structure:
     ```
     ./apps/my-new-app
     ├── src    # Application source code
     └── deploy # Kubernetes deployment configurations (e.g., Deployments, Services)
     ```

## GitOps

This repository leverages [Flux CD](https://fluxcd.io/) to manage and synchronize Kubernetes resources using GitOps principles. Flux continuously ensures that the state of your Kubernetes cluster matches the desired state defined in the Git repository.

### Synced Resource Categories

Flux is configured to manage two primary types of resources:

1. **Infrastructure**:

   - Includes cluster-wide resources such as:
     - **cert-manager**: For managing SSL/TLS certificates.
     - **ingress-nginx**: For handling HTTP and HTTPS traffic to cluster workloads.

2. **Applications**:
   - Covers all workload-specific configurations, such as:
     - Deployments
     - Services
     - ConfigMaps
     - Secrets

### Configuration Location

All configuration files are organized under the `./k8s` directory:

- **Infrastructure configurations**: Define cluster-level tools and services.
- **Application configurations**: Define workload deployments and their supporting resources.

By maintaining these configurations in Git, Flux ensures a reliable, version-controlled, and auditable process for managing your Kubernetes cluster.
