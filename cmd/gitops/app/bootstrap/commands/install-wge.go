package commands

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	"github.com/fluxcd/pkg/apis/meta"
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
	"github.com/weaveworks/weave-gitops-enterprise/cmd/gitops/app/bootstrap/domain"
	"github.com/weaveworks/weave-gitops-enterprise/cmd/gitops/app/bootstrap/utils"
	"github.com/weaveworks/weave-gitops/cmd/gitops/config"
	"github.com/weaveworks/weave-gitops/pkg/runner"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	domainMsg             = "Please select the domain to be used"
	clusterDomainMsg      = "Please enter your cluster domain"
	externalDNSWarningMsg = `Please make sure to have the external DNS service installed in your cluster, or you have a domain that points to your cluster.
For more information about external DNS, please refer to: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring.html
`
	wgeInstallMsg          = "All set installing WGE v%s, This may take a few minutes...\n"
	installSuccessMsg      = "\nWGE v%s is installed successfully\n\n✅ You can visit the UI at https://%s/\n"
	localInstallSuccessMsg = "\nWGE v%s is installed successfully\n\n✅ You can visit the UI at http://localhost:8000/\n"
)

const (
	wgeHelmRepoCommitMsg              = "Add WGE HelmRepository YAML file"
	wgeHelmReleaseCommitMsg           = "Add WGE HelmRelease YAML file"
	wgeChartName                      = "mccp"
	wgeHelmRepositoryName             = "weave-gitops-enterprise-charts"
	wgeHelmReleaseName                = "weave-gitops-enterprise"
	wgeDefaultNamespace               = "flux-system"
	domainTypelocalhost               = "localhost"
	domainTypeExternalDNS             = "external DNS"
	wgeHelmrepoFileName               = "wge-hrepo.yaml"
	wgeHelmReleaseFileName            = "wge-hrelease.yaml"
	wgeChartUrl                       = "https://charts.dev.wkp.weave.works/releases/charts-v3"
	clusterControllerFullOverrideName = "cluster"
	clusterControllerImage            = "docker.io/weaveworks/cluster-controller"
	clusterControllerImageTag         = "v1.5.2"
)

// InstallWge installs weave gitops enterprise chart.
func InstallWge(opts config.Options, version string) (string, error) {
	domainTypes := []string{
		domainTypelocalhost,
		domainTypeExternalDNS,
	}

	domainType, err := utils.GetSelectInput(domainMsg, domainTypes)
	if err != nil {
		return "", err
	}

	userDomain := domainTypelocalhost
	if domainType == domainTypeExternalDNS {
		utils.Warning(externalDNSWarningMsg)
		userDomain, err = utils.GetStringInput(clusterDomainMsg, "")
		if err != nil {
			return "", err
		}
	}

	utils.Info(wgeInstallMsg, version)

	pathInRepo, err := utils.CloneRepo()
	if err != nil {
		return "", err
	}

	defer func() {
		err = utils.CleanupRepo()
		if err != nil {
			utils.Warning(utils.RepoCleanupMsg)
		}
	}()

	wgehelmRepo, err := constructWgeHelmRepository()
	if err != nil {
		return "", err
	}

	err = utils.CreateFileToRepo(wgeHelmrepoFileName, wgehelmRepo, pathInRepo, wgeHelmRepoCommitMsg)
	if err != nil {
		return "", err
	}

	gitOpsSetsValues := map[string]interface{}{
		"enabled": true,
		"controllerManager": map[string]interface{}{
			"manager": map[string]interface{}{
				"args": []string{
					"--health-probe-bind-address=:8081",
					"--metrics-bind-address=127.0.0.1:8080",
					"--leader-elect",
					"--enabled-generators=GitRepository,Cluster,PullRequests,List,APIClient,Matrix,Config",
				},
			},
		},
	}

	clusterController := domain.ClusterController{
		Enabled:          true,
		FullNameOverride: clusterControllerFullOverrideName,
		ControllerManager: domain.ClusterControllerManager{
			Manager: domain.ClusterControllerManagerManager{
				Image: domain.ClusterControllerImage{
					Repository: clusterControllerImage,
					Tag:        clusterControllerImageTag,
				},
			},
		}}

	values := domain.ValuesFile{
		Ingress: constructIngressValues(userDomain),
		TLS: map[string]interface{}{
			"enabled": false,
		},
		GitOpsSets:        gitOpsSetsValues,
		EnablePipelines:   true,
		ClusterController: clusterController,
	}

	wgeHelmRelease, err := constructWGEhelmRelease(values, version)
	if err != nil {
		return "", err
	}

	if err := utils.CreateFileToRepo(wgeHelmReleaseFileName, wgeHelmRelease, pathInRepo, wgeHelmReleaseCommitMsg); err != nil {
		return "", err
	}

	if err := utils.ReconcileFlux(wgeHelmReleaseName); err != nil {
		return "", err
	}

	return userDomain, nil
}

func constructWgeHelmRepository() (string, error) {
	wgeHelmRepo := sourcev1.HelmRepository{
		ObjectMeta: v1.ObjectMeta{
			Name:      wgeHelmRepositoryName,
			Namespace: wgeDefaultNamespace,
		},
		Spec: sourcev1.HelmRepositorySpec{
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
			"external-dns.alpha.kubernetes.io/hostname":                     userDomain,
			"service.beta.kubernetes.io/aws-load-balancer-backend-protocol": "http",
			"service.beta.kubernetes.io/aws-load-balancer-type":             "nlb",
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

func constructWGEhelmRelease(valuesFile domain.ValuesFile, chartVersion string) (string, error) {
	valuesBytes, err := json.Marshal(valuesFile)
	if err != nil {
		return "", err
	}

	wgeHelmRelease := helmv2.HelmRelease{
		ObjectMeta: v1.ObjectMeta{
			Name:      wgeHelmReleaseName,
			Namespace: wgeDefaultNamespace,
		}, Spec: helmv2.HelmReleaseSpec{
			Chart: helmv2.HelmChartTemplate{
				Spec: helmv2.HelmChartTemplateSpec{
					Chart:             wgeChartName,
					ReconcileStrategy: sourcev1.ReconcileStrategyChartVersion,
					SourceRef: helmv2.CrossNamespaceObjectReference{
						Name:      wgeHelmRepositoryName,
						Namespace: wgeDefaultNamespace,
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

// CheckUIDomain display the message to be for external dns or localhost.
func CheckUIDomain(opts config.Options, userDomain string, wgeVersion string) error {
	if !strings.Contains(userDomain, domainTypelocalhost) {
		utils.Info(installSuccessMsg, wgeVersion, userDomain)
		return nil
	}

	utils.Info(localInstallSuccessMsg, wgeVersion)

	var runner runner.CLIRunner
	out, err := runner.Run("kubectl", "-n", "flux-system", "port-forward", "svc/clusters-service", "8000:8000")
	if err != nil {
		return fmt.Errorf("%s%s", err.Error(), string(out))
	}

	return nil
}
