// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package azure

import (
	"flag"

	"github.com/confidential-containers/cloud-api-adaptor/pkg/adaptor/cloud"
)

var azurecfg Config

type Manager struct{}

func init() {
	cloud.AddCloud("azure", &Manager{})
}

func (_ *Manager) ParseCmd(flags *flag.FlagSet) {
	flags.StringVar(&azurecfg.ClientId, "clientid", "", "Client Id, defaults to `AZURE_CLIENT_ID`")
	flags.StringVar(&azurecfg.ClientSecret, "secret", "", "Client Secret, defaults to `AZURE_CLIENT_SECRET`")
	flags.StringVar(&azurecfg.TenantId, "tenantid", "", "Tenant Id, defaults to `AZURE_TENANT_ID`")
	flags.StringVar(&azurecfg.ResourceGroupName, "resourcegroup", "", "Resource Group")
	flags.StringVar(&azurecfg.Zone, "zone", "", "Zone")
	flags.StringVar(&azurecfg.Region, "region", "", "Region")
	flags.StringVar(&azurecfg.SubnetId, "subnetid", "", "Network Subnet Id")
	flags.StringVar(&azurecfg.SecurityGroupId, "securitygroupid", "", "Security Group Id")
	flags.StringVar(&azurecfg.Size, "instance-size", "Standard_DC2as_v5", "Instance size")
	flags.StringVar(&azurecfg.ImageId, "imageid", "", "Image Id")
	flags.StringVar(&azurecfg.SubscriptionId, "subscriptionid", "", "Subscription ID")
	flags.StringVar(&azurecfg.SSHKeyPath, "ssh-key-path", "$HOME/.ssh/id_rsa.pub", "Path to SSH public key")
	flags.StringVar(&azurecfg.SSHUserName, "ssh-username", "peerpod", "SSH User Name")
	flags.BoolVar(&azurecfg.DisableCVM, "disable-cvm", false, "Use non-CVMs for peer pods")
	// Add a List parameter to indicate differet type of instance sizes to be used for the Pod VMs
	flags.Var(&azurecfg.InstanceSizes, "instance-sizes", "Instance sizes to be used for the Pod VMs, comma separated")
	// Add a key value list parameter to indicate custom tags to be used for the Pod VMs
	flags.Var(&azurecfg.Tags, "tags", "Custom tags (key=value pairs) to be used for the Pod VMs, comma separated")
	// Add a flag to disable cloud config and use userdata via metadata service
	flags.BoolVar(&azurecfg.DisableCloudConfig, "disable-cloud-config", false, "Disable cloud config and use userdata via metadata service")
	flags.BoolVar(&azurecfg.EnableSecureBoot, "enable-secure-boot", false, "Enable secure boot for the VMs")
}

func (_ *Manager) LoadEnv() {
	cloud.DefaultToEnv(&azurecfg.ClientId, "AZURE_CLIENT_ID", "")
	cloud.DefaultToEnv(&azurecfg.ClientSecret, "AZURE_CLIENT_SECRET", "")
	cloud.DefaultToEnv(&azurecfg.TenantId, "AZURE_TENANT_ID", "")
	cloud.DefaultToEnv(&azurecfg.SubscriptionId, "AZURE_SUBSCRIPTION_ID", "")
	cloud.DefaultToEnv(&azurecfg.Region, "AZURE_REGION", "")
	cloud.DefaultToEnv(&azurecfg.ResourceGroupName, "AZURE_RESOURCE_GROUP", "")
	cloud.DefaultToEnv(&azurecfg.Size, "AZURE_INSTANCE_SIZE", "Standard_DC2as_v5")
}

func (_ *Manager) NewProvider() (cloud.Provider, error) {
	return NewProvider(&azurecfg)
}

func (_ *Manager) GetConfig() (config *Config) {
	return &azurecfg
}
