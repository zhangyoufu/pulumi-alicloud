// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Vpc Vpc resource. A VPC instance creates a VPC. You can fully control your own VPC, such as selecting IP address ranges, configuring routing tables, and gateways. You can use Alibaba cloud resources such as cloud servers, apsaradb for RDS, and load balancer in your own VPC.
 *
 * > **NOTE:** Available since v1.0.0.
 *
 * > **NOTE:** This resource will auto build a router and a route table while it uses `alicloud.vpc.Network` to build a vpc resource.
 *
 * > **NOTE:** Currently, the IPv4 / IPv6 dual-stack VPC function is under public testing. Only the following regions support IPv4 / IPv6 dual-stack VPC: `cn-hangzhou`, `cn-shanghai`, `cn-shenzhen`, `cn-beijing`, `cn-huhehaote`, `cn-hongkong` and `ap-southeast-1`, and need to apply for public beta qualification. To use, please [submit an application](https://www.alibabacloud.com/help/en/vpc/getting-started/create-a-vpc-with-an-ipv6-cidr-block).
 *
 * ## Module Support
 *
 * You can use the existing vpc module
 * to create a VPC and several VSwitches one-click.
 *
 * For information about Vpc Vpc and how to use it, see [What is Vpc](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/what-is-a-vpc).
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const _default = new alicloud.vpc.Network("default", {
 *     ipv6Isp: "BGP",
 *     description: "test",
 *     cidrBlock: "10.0.0.0/8",
 *     vpcName: name,
 *     enableIpv6: true,
 * });
 * ```
 *
 * ## Import
 *
 * Vpc Vpc can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/network:Network example <id>
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The status of ClassicLink function.
     */
    public readonly classicLinkEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The creation time of the VPC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The VPC description. Defaults to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     */
    public readonly enableIpv6!: pulumi.Output<boolean | undefined>;
    /**
     * The IPv6 CIDR block of the VPC.
     */
    public /*out*/ readonly ipv6CidrBlock!: pulumi.Output<string>;
    /**
     * The IPv6 CIDR block information of the VPC.
     */
    public /*out*/ readonly ipv6CidrBlocks!: pulumi.Output<outputs.vpc.NetworkIpv6CidrBlock[]>;
    /**
     * The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * > **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     */
    public readonly ipv6Isp!: pulumi.Output<string | undefined>;
    /**
     * Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the resource group to which the VPC belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The route table ID of the router created by default on VPC creation.
     */
    public /*out*/ readonly routeTableId!: pulumi.Output<string>;
    /**
     * The ID of the router created by default on VPC creation.
     */
    public /*out*/ readonly routerId!: pulumi.Output<string>;
    /**
     * (Deprecated since v1.206.0+) Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead.
     *
     * @deprecated Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead.
     */
    public /*out*/ readonly routerTableId!: pulumi.Output<string>;
    /**
     * Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondaryCidrBlocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     *
     * @deprecated Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     */
    public readonly secondaryCidrBlocks!: pulumi.Output<string[]>;
    /**
     * The status of the VPC. Valid values:  **Pending**: The VPC is being configured. **Available**: The VPC is available.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags of Vpc.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A list of user CIDRs.
     */
    public readonly userCidrs!: pulumi.Output<string[]>;
    /**
     * The name of the VPC. Defaults to null.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly vpcName!: pulumi.Output<string>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["classicLinkEnabled"] = state ? state.classicLinkEnabled : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["enableIpv6"] = state ? state.enableIpv6 : undefined;
            resourceInputs["ipv6CidrBlock"] = state ? state.ipv6CidrBlock : undefined;
            resourceInputs["ipv6CidrBlocks"] = state ? state.ipv6CidrBlocks : undefined;
            resourceInputs["ipv6Isp"] = state ? state.ipv6Isp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
            resourceInputs["routerId"] = state ? state.routerId : undefined;
            resourceInputs["routerTableId"] = state ? state.routerTableId : undefined;
            resourceInputs["secondaryCidrBlocks"] = state ? state.secondaryCidrBlocks : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userCidrs"] = state ? state.userCidrs : undefined;
            resourceInputs["vpcName"] = state ? state.vpcName : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["classicLinkEnabled"] = args ? args.classicLinkEnabled : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["enableIpv6"] = args ? args.enableIpv6 : undefined;
            resourceInputs["ipv6Isp"] = args ? args.ipv6Isp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["secondaryCidrBlocks"] = args ? args.secondaryCidrBlocks : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userCidrs"] = args ? args.userCidrs : undefined;
            resourceInputs["vpcName"] = args ? args.vpcName : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["ipv6CidrBlock"] = undefined /*out*/;
            resourceInputs["ipv6CidrBlocks"] = undefined /*out*/;
            resourceInputs["routeTableId"] = undefined /*out*/;
            resourceInputs["routerId"] = undefined /*out*/;
            resourceInputs["routerTableId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Network.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The status of ClassicLink function.
     */
    classicLinkEnabled?: pulumi.Input<boolean>;
    /**
     * The creation time of the VPC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The VPC description. Defaults to null.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     */
    enableIpv6?: pulumi.Input<boolean>;
    /**
     * The IPv6 CIDR block of the VPC.
     */
    ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * The IPv6 CIDR block information of the VPC.
     */
    ipv6CidrBlocks?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkIpv6CidrBlock>[]>;
    /**
     * The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * > **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     */
    ipv6Isp?: pulumi.Input<string>;
    /**
     * Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the VPC belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The route table ID of the router created by default on VPC creation.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The ID of the router created by default on VPC creation.
     */
    routerId?: pulumi.Input<string>;
    /**
     * (Deprecated since v1.206.0+) Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead.
     *
     * @deprecated Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead.
     */
    routerTableId?: pulumi.Input<string>;
    /**
     * Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondaryCidrBlocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     *
     * @deprecated Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     */
    secondaryCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the VPC. Valid values:  **Pending**: The VPC is being configured. **Available**: The VPC is available.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags of Vpc.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A list of user CIDRs.
     */
    userCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the VPC. Defaults to null.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    vpcName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The status of ClassicLink function.
     */
    classicLinkEnabled?: pulumi.Input<boolean>;
    /**
     * The VPC description. Defaults to null.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
     * - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Whether to enable the IPv6 network segment. Value:
     * - **false** (default): not enabled.
     * - **true**: on.
     */
    enableIpv6?: pulumi.Input<boolean>;
    /**
     * The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * > **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     */
    ipv6Isp?: pulumi.Input<string>;
    /**
     * Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the VPC belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondaryCidrBlocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
     *
     * @deprecated Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
     */
    secondaryCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags of Vpc.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A list of user CIDRs.
     */
    userCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the VPC. Defaults to null.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    vpcName?: pulumi.Input<string>;
}
