// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ddos

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
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
	case "alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold":
		r = &BasicDefenseThreshold{}
	case "alicloud:ddos/bgpIp:BgpIp":
		r = &BgpIp{}
	case "alicloud:ddos/ddosBgpInstance:DdosBgpInstance":
		r = &DdosBgpInstance{}
	case "alicloud:ddos/ddosCooInstance:DdosCooInstance":
		r = &DdosCooInstance{}
	case "alicloud:ddos/domainResource:DomainResource":
		r = &DomainResource{}
	case "alicloud:ddos/port:Port":
		r = &Port{}
	case "alicloud:ddos/schedulerRule:SchedulerRule":
		r = &SchedulerRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/basicDefenseThreshold",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/bgpIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/ddosBgpInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/ddosCooInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/domainResource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/port",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ddos/schedulerRule",
		&module{version},
	)
}
