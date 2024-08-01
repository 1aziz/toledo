package main

import (
	"encoding/base64"
	"fmt"

	kubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")

		engineVersions, err := container.GetEngineVersions(ctx, &container.GetEngineVersionsArgs{})
		if err != nil {
			return err
		}
		masterVersion := engineVersions.LatestMasterVersion

		cluster, err := container.NewCluster(ctx, "toledo", &container.ClusterArgs{
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

		k8sProvider, err := kubernetes.NewProvider(ctx, "k8sprovider", &kubernetes.ProviderArgs{
			Kubeconfig: generateKubeconfig(cluster.Endpoint, cluster.Name, cluster.MasterAuth),
		}, pulumi.DependsOn([]pulumi.Resource{cluster}))
		if err != nil {
			return err
		}

		gh_token := conf.RequireSecret("gh_token")
		crDockerConfig := conf.RequireSecret("doToken").ApplyT(
			func(token string) string {
				return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(
					"{\"auths\":{\"%s\":{\"auth\":\"%s\"}}}",
					"crUrl",
					base64.StdEncoding.EncodeToString([]byte(token+gh_token.ToStringOutput().ElementType().String())),
				)))
			},
		).(pulumi.StringOutput)

		_, err = corev1.NewSecret(ctx, "cr-secret",
			&corev1.SecretArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String("cr-secret"),
				},
				Data: pulumi.StringMap{
					".dockerconfigjson": crDockerConfig,
				},
				Type: pulumi.String("kubernetes.io/dockerconfigjson"),
			},
			pulumi.Provider(k8sProvider),
		)
		if err != nil {
			return err
		}

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
