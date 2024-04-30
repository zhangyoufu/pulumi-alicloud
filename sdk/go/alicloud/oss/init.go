// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

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
	case "alicloud:oss/bucket:Bucket":
		r = &Bucket{}
	case "alicloud:oss/bucketAcl:BucketAcl":
		r = &BucketAcl{}
	case "alicloud:oss/bucketCors:BucketCors":
		r = &BucketCors{}
	case "alicloud:oss/bucketHttpsConfig:BucketHttpsConfig":
		r = &BucketHttpsConfig{}
	case "alicloud:oss/bucketLogging:BucketLogging":
		r = &BucketLogging{}
	case "alicloud:oss/bucketObject:BucketObject":
		r = &BucketObject{}
	case "alicloud:oss/bucketPolicy:BucketPolicy":
		r = &BucketPolicy{}
	case "alicloud:oss/bucketReferer:BucketReferer":
		r = &BucketReferer{}
	case "alicloud:oss/bucketReplication:BucketReplication":
		r = &BucketReplication{}
	case "alicloud:oss/bucketRequestPayment:BucketRequestPayment":
		r = &BucketRequestPayment{}
	case "alicloud:oss/bucketServerSideEncryption:BucketServerSideEncryption":
		r = &BucketServerSideEncryption{}
	case "alicloud:oss/bucketVersioning:BucketVersioning":
		r = &BucketVersioning{}
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
		"oss/bucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketCors",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketHttpsConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketLogging",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketObject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketReferer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketReplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketRequestPayment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketServerSideEncryption",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oss/bucketVersioning",
		&module{version},
	)
}
