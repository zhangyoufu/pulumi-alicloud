// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gpdb

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "alicloud:gpdb/account:Account":
		r = &Account{}
	case "alicloud:gpdb/backupPolicy:BackupPolicy":
		r = &BackupPolicy{}
	case "alicloud:gpdb/connection:Connection":
		r = &Connection{}
	case "alicloud:gpdb/dbInstancePlan:DbInstancePlan":
		r = &DbInstancePlan{}
	case "alicloud:gpdb/dbResourceGroup:DbResourceGroup":
		r = &DbResourceGroup{}
	case "alicloud:gpdb/elasticInstance:ElasticInstance":
		r = &ElasticInstance{}
	case "alicloud:gpdb/externalDataService:ExternalDataService":
		r = &ExternalDataService{}
	case "alicloud:gpdb/instance:Instance":
		r = &Instance{}
	case "alicloud:gpdb/remoteAdbDataSource:RemoteAdbDataSource":
		r = &RemoteAdbDataSource{}
	case "alicloud:gpdb/streamingDataService:StreamingDataService":
		r = &StreamingDataService{}
	case "alicloud:gpdb/streamingDataSource:StreamingDataSource":
		r = &StreamingDataSource{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/backupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/connection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/dbInstancePlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/dbResourceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/elasticInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/externalDataService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/remoteAdbDataSource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/streamingDataService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"gpdb/streamingDataSource",
		&module{version},
	)
}
