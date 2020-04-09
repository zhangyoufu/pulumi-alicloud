// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides information about [router interfaces](https://www.alibabacloud.com/help/doc-detail/52412.htm)
// that connect VPCs together.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/router_interfaces.html.markdown.
func GetRouterInterfaces(ctx *pulumi.Context, args *GetRouterInterfacesArgs, opts ...pulumi.InvokeOption) (*GetRouterInterfacesResult, error) {
	var rv GetRouterInterfacesResult
	err := ctx.Invoke("alicloud:vpc/getRouterInterfaces:getRouterInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterInterfaces.
type GetRouterInterfacesArgs struct {
	// A list of router interface IDs.
	Ids []string `pulumi:"ids"`
	// A regex string used to filter by router interface name.
	NameRegex *string `pulumi:"nameRegex"`
	// ID of the peer router interface.
	OppositeInterfaceId *string `pulumi:"oppositeInterfaceId"`
	// Account ID of the owner of the peer router interface.
	OppositeInterfaceOwnerId *string `pulumi:"oppositeInterfaceOwnerId"`
	OutputFile               *string `pulumi:"outputFile"`
	// Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and
	// `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `routerType` is set to `VBR`.
	Role *string `pulumi:"role"`
	// ID of the VRouter located in the local region.
	RouterId *string `pulumi:"routerId"`
	// Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
	RouterType *string `pulumi:"routerType"`
	// Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
	Specification *string `pulumi:"specification"`
	// Expected status. Valid values are `Active`, `Inactive` and `Idle`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getRouterInterfaces.
type GetRouterInterfacesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of router interface IDs.
	Ids []string `pulumi:"ids"`
	// A list of router interfaces. Each element contains the following attributes:
	Interfaces []GetRouterInterfacesInterface `pulumi:"interfaces"`
	NameRegex  *string                        `pulumi:"nameRegex"`
	// A list of router interface names.
	Names []string `pulumi:"names"`
	// Peer router interface ID.
	OppositeInterfaceId *string `pulumi:"oppositeInterfaceId"`
	// Account ID of the owner of the peer router interface.
	OppositeInterfaceOwnerId *string `pulumi:"oppositeInterfaceOwnerId"`
	OutputFile               *string `pulumi:"outputFile"`
	// Router interface role. Possible values: `InitiatingSide` and `AcceptingSide`.
	Role *string `pulumi:"role"`
	// ID of the VRouter located in the local region.
	RouterId *string `pulumi:"routerId"`
	// Router type in the local region. Possible values: `VRouter` and `VBR`.
	RouterType *string `pulumi:"routerType"`
	// Router interface specification. Possible values: `Small.1`, `Middle.1`, `Large.2`, ...etc.
	Specification *string `pulumi:"specification"`
	// Router interface status. Possible values: `Active`, `Inactive` and `Idle`.
	Status *string `pulumi:"status"`
}
