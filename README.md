# Toledo

Toledo is a simple GitOps boilerplate designed to help you get started with configuring Kubernetes resources using GitOps practices. This repository provides sample Kubernetes configurations and includes a demo application, `halo-toledo`, for experimenting with core Cloud Native concepts and tools.

## Key Features

- **Hands-on GitOps Learning**: Discover how to automate deployments and manage Kubernetes cluster configurations using **Flux**.
- **Networking Exploration**: Experiment with advanced Kubernetes networking concepts using **Cilium**, including service discovery, load balancing, and service mesh functionality.
- **Task Automation**: Simplify cluster bootstrapping with **Taskfile**, a task runner that standardizes and automates common setup processes.

## Who Is This For?

This repository is ideal for:

- Beginners to GitOps and Kubernetes who want a practical, hands-on playground.
- Developers exploring Kubernetes networking and Cloud Native tools.
- Teams looking for a starting point to integrate GitOps practices into their workflows.

## Included Tools and Technologies

1. **Cilium**
   Delve into service mesh capabilities, such as:

   - Service discovery
   - Load balancing
   - Advanced networking

2. **Flux**
   Master GitOps by automating deployments and managing your Kubernetes cluster configuration directly from a Git repository.

3. **Taskfile**
   Simplify and streamline cluster bootstrapping and other repetitive tasks with predefined commands.

## Benefits of Using Toledo

By working through this repository, you will:

- Gain practical experience with GitOps workflows and tooling.
- Understand Kubernetes networking fundamentals.
- Learn how to set up and manage a Cloud Native environment efficiently.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/1aziz/toledo.git
   ```
2. Install the required tools (e.g., Flux, Cilium, Taskfile).
3. Follow the provided instructions in the respective directories to set up your Kubernetes cluster and start experimenting.

## Get Started (Local)

### 1. Bootstrap the Cluster

To set up and configure a local Kubernetes cluster, follow these steps:

1. **Install `Taskfile`**:
   [`Taskfile`](https://taskfile.dev/) is required for automating cluster setup. Install it before proceeding.

2. **Run the Bootstrap Command**:
   Execute the following command to bootstrap a new local cluster:

   ```sh
   task kind:bootstrap
   ```

This command performs the following actions:

```
kind:bootstrap
├── kind:create
│   ├── Deletes an existing Kind cluster (`toledo-local`) if it exists
│   ├── Creates a new Kind cluster using `kind-cluster-config.yaml`
│   └── Verifies that the cluster context is set correctly
├── apply:cilium
│   ├── Adds the Cilium Helm repository
│   └── Installs the Cilium networking plugin (version 1.16.5) into the cluster
├── apply:regcred
│   └── Creates a Docker registry credential secret (`regcred`) for pulling images
└── bootstrap
    └── Configures Flux for GitHub:
        │── Installs Flux CD into the cluster
        ├── Authenticates using `GITHUB_TOKEN`
        ├── Configures Flux to sync from a GitHub repository
        └── Uses specified branch, repository, and cluster configuration path
```

### 2. Explore the Repository

#### **`halo-toledo` Demo App**

- **Location**:

  - Source code: `./halo-toledo/src`
  - Kubernetes configurations: `./halo-toledo/deploy`

- **What to Do**:
  - Review the app's source code and deployment configurations to understand how applications are deployed using GitOps.
  - Experiment with deploying and updating the app within your local cluster.

#### **Bring Your Own Applications**

This repository is designed to be extended with your own applications.

- **Steps**:

  1. Create a new directory under the `./apps` folder:
     Example: `./apps/my-new-app`
  2. Inside your new app directory, create the following structure:
     ```
     ./apps/my-new-app
     ├── src    # Application source code
     └── deploy # Kubernetes deployment configurations (e.g., Deployments, Services)
     ```

- **Benefits**:
  - Practice deploying real-world applications using GitOps principles.
  - Experiment with custom configurations and tools in a Cloud Native environment.
  - Tailor the repository to your specific learning or project needs.

## GitOps

This repository leverages [Flux CD](https://fluxcd.io/) to manage and synchronize Kubernetes resources using GitOps principles. Flux continuously ensures that the state of your Kubernetes cluster matches the desired state defined in the Git repository.

### Synced Resource Categories

Flux is configured to manage two primary types of resources:

1. **Infrastructure**:

   - Includes cluster-wide resources such as:
     - **cert-manager**: For managing SSL/TLS certificates.
     - **Cilium**: For advanced Kubernetes networking and service mesh capabilities.
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
