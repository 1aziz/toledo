package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		engineVersions, err := container.GetEngineVersions(ctx, &container.GetEngineVersionsArgs{})
		if err != nil {
			return err
		}
		masterVersion := engineVersions.LatestMasterVersion

		cluster, err := container.NewCluster(ctx, "demo-cluster", &container.ClusterArgs{
			Location:           pulumi.String("europe-west3"),
			EnableAutopilot:    pulumi.Bool(true),
			Network:            pulumi.String("default"),
			MinMasterVersion:   pulumi.String(masterVersion),
			NodeVersion:        pulumi.String(masterVersion),
			DeletionProtection: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}

		ctx.Export("kubeconfig", generateKubeconfig(cluster.Endpoint, cluster.Name, cluster.MasterAuth))

		return nil
	})
}

func generateKubeconfig(clusterEndpoint pulumi.StringOutput, clusterName pulumi.StringOutput,
	clusterMasterAuth container.ClusterMasterAuthOutput) pulumi.StringOutput {
	context := pulumi.Sprintf("demo_%s", clusterName)

	return pulumi.Sprintf(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: %s
    server: https://%s
  name: %s
contexts:
- context:
    cluster: %s
    user: %s
  name: %s
current-context: %s
kind: Config
preferences: {}
users:
- name: %s
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: gke-gcloud-auth-plugin
      installHint: Install gke-gcloud-auth-plugin for use with kubectl by following
        https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
      provideClusterInfo: true
`,
		clusterMasterAuth.ClusterCaCertificate().Elem(),
		clusterEndpoint, context, context, context, context, context, context)
}
