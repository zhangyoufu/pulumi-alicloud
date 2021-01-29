// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup":
		r, err = NewAutoProvisioningGroup(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/copyImage:CopyImage":
		r, err = NewCopyImage(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/dedicatedHost:DedicatedHost":
		r, err = NewDedicatedHost(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/disk:Disk":
		r, err = NewDisk(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/diskAttachment:DiskAttachment":
		r, err = NewDiskAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/eip:Eip":
		r, err = NewEip(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/eipAssociation:EipAssociation":
		r, err = NewEipAssociation(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/image:Image":
		r, err = NewImage(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/imageCopy:ImageCopy":
		r, err = NewImageCopy(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/imageExport:ImageExport":
		r, err = NewImageExport(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/imageImport:ImageImport":
		r, err = NewImageImport(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/imageSharePermission:ImageSharePermission":
		r, err = NewImageSharePermission(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/instance:Instance":
		r, err = NewInstance(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/keyPair:KeyPair":
		r, err = NewKeyPair(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/keyPairAttachment:KeyPairAttachment":
		r, err = NewKeyPairAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/launchTemplate:LaunchTemplate":
		r, err = NewLaunchTemplate(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/reservedInstance:ReservedInstance":
		r, err = NewReservedInstance(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/securityGroup:SecurityGroup":
		r, err = NewSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/securityGroupRule:SecurityGroupRule":
		r, err = NewSecurityGroupRule(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/snapshot:Snapshot":
		r, err = NewSnapshot(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:ecs/snapshotPolicy:SnapshotPolicy":
		r, err = NewSnapshotPolicy(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/autoProvisioningGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/copyImage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/dedicatedHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/disk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/diskAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/eip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/eipAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/image",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/imageCopy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/imageExport",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/imageImport",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/imageSharePermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/keyPair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/keyPairAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/launchTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/reservedInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/securityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/securityGroupRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/snapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ecs/snapshotPolicy",
		&module{version},
	)
}
