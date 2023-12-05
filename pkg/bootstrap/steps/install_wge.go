package steps

import (
	"encoding/json"
	"fmt"
	"time"

	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	"github.com/fluxcd/pkg/apis/meta"
	sourcev1beta2 "github.com/fluxcd/source-controller/api/v1beta2"
	"github.com/weaveworks/weave-gitops-enterprise/pkg/bootstrap/utils"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/cli-utils/pkg/object"
)

const (
	externalDNSWarningMsg = `please make sure to have the external DNS service installed in your cluster, or you have a domain that points to your cluster.`
	clusterDomainMsg      = "please enter your cluster domain"

	wgeInstallMsg = "installing v%s ... It may take a few minutes."
)

const (
	wgeHelmRepoCommitMsg              = "Add WGE HelmRepository YAML file"
	wgeHelmReleaseCommitMsg           = "Add WGE HelmRelease YAML file"
	wgeChartName                      = "mccp"
	wgeHelmRepositoryName             = "weave-gitops-enterprise-charts"
	WgeHelmReleaseName                = "weave-gitops-enterprise"
	WGEDefaultNamespace               = "flux-system"
	WGEDefaultRepoName                = "flux-system"
	wgeHelmrepoFileName               = "wge-hrepo.yaml"
	wgeHelmReleaseFileName            = "wge-hrelease.yaml"
	wgeChartUrl                       = "https://charts.dev.wkp.weave.works/releases/charts-v3"
	clusterControllerFullOverrideName = "cluster"
	clusterControllerImageName        = "docker.io/weaveworks/cluster-controller"
	clusterControllerImageTag         = "v1.5.2"
	gitopssetsEnabledGenerators       = "GitRepository,Cluster,PullRequests,List,APIClient,Matrix,Config"
	gitopssetsBindAddress             = "127.0.0.1:8080"
	gitopssetsHealthBindAddress       = ":8081"
)

var Components = []string{"cluster-controller-manager",
	"weave-gitops-enterprise-mccp-cluster-bootstrap-controller",
	"weave-gitops-enterprise-mccp-cluster-service"}

var getUserDomain = StepInput{
	Name:         inUserDomain,
	Type:         stringInput,
	Msg:          clusterDomainMsg,
	DefaultValue: "",
	Enabled:      isUserDomainEnabled,
}

// NewInstallWGEStep step to install Weave GitOps Enterprise
func NewInstallWGEStep(config Config) BootstrapStep {
	inputs := []StepInput{}

	if config.UserDomain == "" {
		inputs = append(inputs, getUserDomain)
	}

	return BootstrapStep{
		Name:  "install Weave GitOps Enterprise",
		Input: inputs,
		Step:  installWge,
	}
}

// InstallWge installs weave gitops enterprise chart.
func installWge(input []StepInput, c *Config) ([]StepOutput, error) {
	var ingressValues map[string]interface{}
	switch c.DomainType {
	case domainTypeLocalhost:
		c.UserDomain = domainTypeLocalhost
	case domainTypeExternalDNS:
		if c.UserDomain == "" {
			for _, param := range input {
				if param.Name == inUserDomain {
					userDomain, ok := param.Value.(string)
					if !ok {
						return []StepOutput{}, fmt.Errorf("unexpected error occurred. UserDomain not found")
					}
					c.UserDomain = userDomain
				}
			}
			ingressValues = constructIngressValues(c.UserDomain)
		}
	default:
		return []StepOutput{}, fmt.Errorf("unsupported domain type:%s", c.DomainType)
	}

	c.Logger.Actionf(wgeInstallMsg, c.WGEVersion)

	wgehelmRepo, err := constructWgeHelmRepository()
	if err != nil {
		return []StepOutput{}, err
	}
	c.Logger.Actionf("rendered HelmRepository file")

	gitOpsSetsValues := map[string]interface{}{
		"enabled": true,
		"controllerManager": map[string]interface{}{
			"manager": map[string]interface{}{
				"args": []string{
					fmt.Sprintf("--health-probe-bind-address=%s", gitopssetsHealthBindAddress),
					fmt.Sprintf("--metrics-bind-address=%s", gitopssetsBindAddress),
					"--leader-elect",
					fmt.Sprintf("--enabled-generators=%s", gitopssetsEnabledGenerators),
				},
			},
		},
	}

	clusterController := clusterController{
		Enabled:          true,
		FullNameOverride: clusterControllerFullOverrideName,
		ControllerManager: clusterControllerManager{
			Manager: clusterControllerManagerManager{
				Image: clusterControllerImage{
					Repository: clusterControllerImageName,
					Tag:        clusterControllerImageTag,
				},
			},
		}}

	values := valuesFile{
		Ingress: ingressValues,
		TLS: map[string]interface{}{
			"enabled": false,
		},
		GitOpsSets:        gitOpsSetsValues,
		EnablePipelines:   true,
		ClusterController: clusterController,
	}

	wgeHelmRelease, err := constructWGEhelmRelease(values, c.WGEVersion)
	if err != nil {
		return []StepOutput{}, err
	}
	c.Logger.Actionf("rendered HelmRelease file")

	helmrepoFile := fileContent{
		Name:      wgeHelmrepoFileName,
		Content:   wgehelmRepo,
		CommitMsg: wgeHelmRepoCommitMsg,
	}
	helmreleaseFile := fileContent{
		Name:      wgeHelmReleaseFileName,
		Content:   wgeHelmRelease,
		CommitMsg: wgeHelmReleaseCommitMsg,
	}

	if !c.SkipComponentCheck {
		// Wait for the components to be healthy

		c.Logger.Waitingf("waiting for components to be healthy")
		err = reportComponentsHealth(c, Components, WGEDefaultNamespace, 5*time.Minute)
		if err != nil {
			return []StepOutput{}, err
		}
	}

	return []StepOutput{
		{
			Name:  wgeHelmrepoFileName,
			Type:  typeFile,
			Value: helmrepoFile,
		},
		{
			Name:  wgeHelmReleaseFileName,
			Type:  typeFile,
			Value: helmreleaseFile,
		},
	}, nil
}

func constructWgeHelmRepository() (string, error) {
	wgeHelmRepo := sourcev1beta2.HelmRepository{
		ObjectMeta: v1.ObjectMeta{
			Name:      wgeHelmRepositoryName,
			Namespace: WGEDefaultNamespace,
		},
		Spec: sourcev1beta2.HelmRepositorySpec{
			URL: wgeChartUrl,
			Interval: v1.Duration{
				Duration: time.Minute,
			},
			SecretRef: &meta.LocalObjectReference{
				Name: entitlementSecretName,
			},
		},
	}

	return utils.CreateHelmRepositoryYamlString(wgeHelmRepo)
}

func constructIngressValues(userDomain string) map[string]interface{} {
	ingressValues := map[string]interface{}{
		"annotations": map[string]string{
			"external-dns.alpha.kubernetes.io/hostname": userDomain,
		},
		"className": "public-nginx",
		"enabled":   true,
		"hosts": []map[string]interface{}{
			{
				"host": userDomain,
				"paths": []map[string]string{
					{
						"path":     "/",
						"pathType": "ImplementationSpecific",
					},
				},
			},
		},
	}

	return ingressValues
}

func constructWGEhelmRelease(valuesFile valuesFile, chartVersion string) (string, error) {
	valuesBytes, err := json.Marshal(valuesFile)
	if err != nil {
		return "", err
	}

	wgeHelmRelease := helmv2.HelmRelease{
		ObjectMeta: v1.ObjectMeta{
			Name:      WgeHelmReleaseName,
			Namespace: WGEDefaultNamespace,
		}, Spec: helmv2.HelmReleaseSpec{
			Chart: helmv2.HelmChartTemplate{
				Spec: helmv2.HelmChartTemplateSpec{
					Chart:             wgeChartName,
					ReconcileStrategy: sourcev1beta2.ReconcileStrategyChartVersion,
					SourceRef: helmv2.CrossNamespaceObjectReference{
						Name:      wgeHelmRepositoryName,
						Namespace: WGEDefaultNamespace,
					},
					Version: chartVersion,
				},
			},
			Install: &helmv2.Install{
				CRDs: helmv2.CreateReplace,
			},
			Upgrade: &helmv2.Upgrade{
				CRDs: helmv2.CreateReplace,
			},
			Interval: v1.Duration{
				Duration: time.Hour,
			},
			Values: &apiextensionsv1.JSON{Raw: valuesBytes},
		},
	}

	return utils.CreateHelmReleaseYamlString(wgeHelmRelease)
}

func isUserDomainEnabled(input []StepInput, c *Config) bool {
	if c.DomainType == domainTypeExternalDNS {
		c.Logger.L().Info(externalDNSWarningMsg)
		return true
	}
	return false
}

func reportComponentsHealth(c *Config, componentNames []string, namespace string, timeout time.Duration) error {
	// Initialize the status checker
	checker, err := utils.NewStatusChecker(c.KubernetesClient, 5*time.Second, timeout, c.Logger)
	if err != nil {
		return err
	}

	// Construct a list of resources to check
	var identifiers []object.ObjMetadata
	for _, name := range componentNames {
		identifiers = append(identifiers, object.ObjMetadata{
			Namespace: namespace,
			Name:      name,
			GroupKind: schema.GroupKind{Group: "apps", Kind: "Deployment"},
		})
	}

	// Perform the health check
	if err := checker.Assess(identifiers...); err != nil {
		return err
	}

	return nil
}
