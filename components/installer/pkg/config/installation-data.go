package config

import (
	"github.com/kyma-project/kyma/components/installer/pkg/apis/installer/v1alpha1"
)

// InstallationContext describes properties of K8S Installation object that triggers installation process
type InstallationContext struct {
	Name      string
	Namespace string
}

// InstallationData describes all installation attributes
type InstallationData struct {
	Context                    InstallationContext
	ExternalPublicIP           string
	Domain                     string
	KymaVersion                string
	URL                        string
	AzureBrokerTenantID        string
	AzureBrokerClientID        string
	AzureBrokerSubscriptionID  string
	AzureBrokerClientSecret    string
	ClusterTLSKey              string
	ClusterTLSCert             string
	RemoteEnvCa                string
	RemoteEnvCaKey             string
	RemoteEnvIP                string
	K8sApiserverURL            string
	K8sApiserverCa             string
	UITestUser                 string
	UITestPassword             string
	AdminGroup                 string
	EtcdBackupABSContainerName string
	EnableEtcdBackupOperator   string
	EtcdBackupABSAccount       string
	EtcdBackupABSKey           string
	Components                 map[string]v1alpha1.KymaComponent
	IsLocalInstallation        bool
	VictorOpsApiKey            string
	VictorOpsRoutingKey        string
	SlackChannel               string
	SlackApiUrl                string
}

// NewInstallationData .
func NewInstallationData(installation *v1alpha1.Installation, installationConfig *installationConfig) (*InstallationData, error) {

	ctx := InstallationContext{
		Name:      installation.Name,
		Namespace: installation.Namespace,
	}

	res := &InstallationData{
		Context:                    ctx,
		ExternalPublicIP:           installationConfig.ExternalPublicIP,
		Domain:                     installationConfig.Domain,
		KymaVersion:                installation.Spec.KymaVersion,
		URL:                        installation.Spec.URL,
		AzureBrokerTenantID:        installationConfig.AzureBrokerTenantID,
		AzureBrokerClientID:        installationConfig.AzureBrokerClientID,
		AzureBrokerSubscriptionID:  installationConfig.AzureBrokerSubscriptionID,
		AzureBrokerClientSecret:    installationConfig.AzureBrokerClientSecret,
		ClusterTLSKey:              installationConfig.ClusterTLSKey,
		ClusterTLSCert:             installationConfig.ClusterTLSCert,
		RemoteEnvCa:                installationConfig.RemoteEnvCa,
		RemoteEnvCaKey:             installationConfig.RemoteEnvCaKey,
		RemoteEnvIP:                installationConfig.RemoteEnvIP,
		K8sApiserverURL:            installationConfig.K8sApiserverUrl,
		K8sApiserverCa:             installationConfig.K8sApiserverCa,
		UITestUser:                 installationConfig.UITestUser,
		UITestPassword:             installationConfig.UITestPassword,
		AdminGroup:                 installationConfig.AdminGroup,
		EtcdBackupABSContainerName: installationConfig.EtcdBackupABSContainerName,
		EnableEtcdBackupOperator:   installationConfig.EnableEtcdBackupOperator,
		EtcdBackupABSAccount:       installationConfig.EtcdBackupABSAccount,
		EtcdBackupABSKey:           installationConfig.EtcdBackupABSKey,
		Components:                 convertToMap(installation.Spec.Components),
		IsLocalInstallation:        installationConfig.IsLocalInstallation,
		VictorOpsApiKey:            installationConfig.VictorOpsApiKey,
		VictorOpsRoutingKey:        installationConfig.VictorOpsRoutingKey,
		SlackChannel:               installationConfig.SlackChannel,
		SlackApiUrl:                installationConfig.SlackApiUrl,
	}
	return res, nil
}

// ShouldInstallComponent returns true if the provided component is on the list of desired components
func (installationData *InstallationData) ShouldInstallComponent(componentName string) bool {
	_, found := installationData.Components[componentName]
	return found
}

func convertToMap(cList []v1alpha1.KymaComponent) map[string]v1alpha1.KymaComponent {
	output := make(map[string]v1alpha1.KymaComponent, len(cList))
	for _, c := range cList {
		if c.Name != "" {
			output[c.Name] = c
		}
	}
	return output
}
