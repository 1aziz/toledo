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

## Get Started (Local installation with Kind)

> [!IMPORTANT]
> Before you proceed with the steps below, ensure you have the following tools installed and configured on your machine:
> - [Gum](https://github.com/charmbracelet/gum),
> - [Flux CD](https://fluxcd.io/docs/installation/),
> - [Taskfile](https://taskfile.dev/#/installation)
> - [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

In the future, we can simplify the prerequisites by incorporating the installation of the Flux CD CLI and Kind directly into the Taskfile.

### Bootstrap the Cluster

To set up and configure a _local_ Kubernetes cluster, run the bootstrap script `sh bootstrap` and follow the steps in terminal.

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

The configuration files are structured under the `./k8s` directory for better organization and manageability:

- **Infrastructure Configurations**: These files define cluster-level tools and services, such as monitoring, networking, and security.
- **Application Configurations**: These files define workload deployments, including application services and their supporting resources like ConfigMaps and Secrets.

### Adding New Applications

The `halo-toledo` application, which uses an image stored on GitHub Packages, serves as an example of how to configure new applications. 

- **Source Code**: Located at `./apps/halo-toledo/src`.
- **Kubernetes Manifests**: Found under `./apps/halo-toledo/deploy`.
- **Kustomization**: Available in `k8s/apps/toledo`.

You can follow a similar approach to add your own applications. Below are the steps:

#### Steps to Add a New Application:

1. **Create a Directory**  
   Create a new directory under the `./apps` folder.  
   Example: `./apps/my-new-app`

2. **Set Up Directory Structure**  
   Inside your new application directory, organize the following structure:
   ```
   ./apps/my-new-app
      ├── src # Application source code
      └── deploy # Kubernetes deployment configurations (e.g., Deployment, Service)
   ```

4. **Prepare Kubernetes Manifests**  
- Ensure your application manifests and container image are ready.
- Configure the cluster with an appropriate image pull secret to fetch your application's image.

4. **Add `Kustomization`**  
Create a new directory under the `./k8s/apps` folder to store your app’s `Kustomization` file.  
Example: Refer to `k8s/apps/toledo` for guidance.
