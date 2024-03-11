// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a NLB Load Balancer resource.
 *
 * For information about NLB Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createloadbalancer).
 *
 * > **NOTE:** Available since v1.191.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({});
 * const defaultZones = alicloud.nlb.getZones({});
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const default1 = new alicloud.vpc.Switch("default1", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.1.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[1]?.id),
 * });
 * const defaultLoadBalancer = new alicloud.nlb.LoadBalancer("defaultLoadBalancer", {
 *     loadBalancerName: name,
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     loadBalancerType: "Network",
 *     addressType: "Internet",
 *     addressIpVersion: "Ipv4",
 *     vpcId: defaultNetwork.id,
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 *     zoneMappings: [
 *         {
 *             vswitchId: defaultSwitch.id,
 *             zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *         },
 *         {
 *             vswitchId: default1.id,
 *             zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[1]?.id),
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * NLB Load Balancer can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:nlb/loadBalancer:LoadBalancer example <id>
 * ```
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nlb/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * Protocol version. Value:
     * - **ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     */
    public readonly addressIpVersion!: pulumi.Output<string>;
    /**
     * The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     */
    public readonly addressType!: pulumi.Output<string>;
    /**
     * The ID of the shared bandwidth package associated with the public network instance.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * Resource creation time, using Greenwich Mean Time, formating' yyyy-MM-ddTHH:mm:ssZ '.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     */
    public readonly crossZoneEnabled!: pulumi.Output<boolean>;
    /**
     * Delete protection. See `deletionProtectionConfig` below.
     */
    public readonly deletionProtectionConfig!: pulumi.Output<outputs.nlb.LoadBalancerDeletionProtectionConfig>;
    /**
     * Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     */
    public readonly deletionProtectionEnabled!: pulumi.Output<boolean>;
    /**
     * The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
     */
    public readonly deletionProtectionReason!: pulumi.Output<string>;
    /**
     * The domain name of the NLB instance.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     */
    public readonly ipv6AddressType!: pulumi.Output<string>;
    /**
     * The business status of the NLB instance.
     */
    public /*out*/ readonly loadBalancerBusinessStatus!: pulumi.Output<string>;
    /**
     * The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     */
    public readonly loadBalancerName!: pulumi.Output<string | undefined>;
    /**
     * Load balancing type. Only value: **network**, which indicates network-based load balancing.
     */
    public readonly loadBalancerType!: pulumi.Output<string>;
    /**
     * Modify protection. See `modificationProtectionConfig` below.
     */
    public readonly modificationProtectionConfig!: pulumi.Output<outputs.nlb.LoadBalancerModificationProtectionConfig>;
    /**
     * The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
     */
    public readonly modificationProtectionReason!: pulumi.Output<string>;
    /**
     * Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     */
    public readonly modificationProtectionStatus!: pulumi.Output<string>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The security group to which the network-based SLB instance belongs.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * ON.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * List of labels.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the network-based SLB instance.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
     */
    public readonly zoneMappings!: pulumi.Output<outputs.nlb.LoadBalancerZoneMapping[]>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["addressType"] = state ? state.addressType : undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["crossZoneEnabled"] = state ? state.crossZoneEnabled : undefined;
            resourceInputs["deletionProtectionConfig"] = state ? state.deletionProtectionConfig : undefined;
            resourceInputs["deletionProtectionEnabled"] = state ? state.deletionProtectionEnabled : undefined;
            resourceInputs["deletionProtectionReason"] = state ? state.deletionProtectionReason : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["ipv6AddressType"] = state ? state.ipv6AddressType : undefined;
            resourceInputs["loadBalancerBusinessStatus"] = state ? state.loadBalancerBusinessStatus : undefined;
            resourceInputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            resourceInputs["loadBalancerType"] = state ? state.loadBalancerType : undefined;
            resourceInputs["modificationProtectionConfig"] = state ? state.modificationProtectionConfig : undefined;
            resourceInputs["modificationProtectionReason"] = state ? state.modificationProtectionReason : undefined;
            resourceInputs["modificationProtectionStatus"] = state ? state.modificationProtectionStatus : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneMappings"] = state ? state.zoneMappings : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            if ((!args || args.addressType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressType'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.zoneMappings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneMappings'");
            }
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["addressType"] = args ? args.addressType : undefined;
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["crossZoneEnabled"] = args ? args.crossZoneEnabled : undefined;
            resourceInputs["deletionProtectionConfig"] = args ? args.deletionProtectionConfig : undefined;
            resourceInputs["deletionProtectionEnabled"] = args ? args.deletionProtectionEnabled : undefined;
            resourceInputs["deletionProtectionReason"] = args ? args.deletionProtectionReason : undefined;
            resourceInputs["ipv6AddressType"] = args ? args.ipv6AddressType : undefined;
            resourceInputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            resourceInputs["loadBalancerType"] = args ? args.loadBalancerType : undefined;
            resourceInputs["modificationProtectionConfig"] = args ? args.modificationProtectionConfig : undefined;
            resourceInputs["modificationProtectionReason"] = args ? args.modificationProtectionReason : undefined;
            resourceInputs["modificationProtectionStatus"] = args ? args.modificationProtectionStatus : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zoneMappings"] = args ? args.zoneMappings : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["loadBalancerBusinessStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * Protocol version. Value:
     * - **ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     */
    addressType?: pulumi.Input<string>;
    /**
     * The ID of the shared bandwidth package associated with the public network instance.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Resource creation time, using Greenwich Mean Time, formating' yyyy-MM-ddTHH:mm:ssZ '.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     */
    crossZoneEnabled?: pulumi.Input<boolean>;
    /**
     * Delete protection. See `deletionProtectionConfig` below.
     */
    deletionProtectionConfig?: pulumi.Input<inputs.nlb.LoadBalancerDeletionProtectionConfig>;
    /**
     * Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     */
    deletionProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
     */
    deletionProtectionReason?: pulumi.Input<string>;
    /**
     * The domain name of the NLB instance.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     */
    ipv6AddressType?: pulumi.Input<string>;
    /**
     * The business status of the NLB instance.
     */
    loadBalancerBusinessStatus?: pulumi.Input<string>;
    /**
     * The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     */
    loadBalancerName?: pulumi.Input<string>;
    /**
     * Load balancing type. Only value: **network**, which indicates network-based load balancing.
     */
    loadBalancerType?: pulumi.Input<string>;
    /**
     * Modify protection. See `modificationProtectionConfig` below.
     */
    modificationProtectionConfig?: pulumi.Input<inputs.nlb.LoadBalancerModificationProtectionConfig>;
    /**
     * The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
     */
    modificationProtectionReason?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     */
    modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The security group to which the network-based SLB instance belongs.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ON.
     */
    status?: pulumi.Input<string>;
    /**
     * List of labels.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the network-based SLB instance.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
     */
    zoneMappings?: pulumi.Input<pulumi.Input<inputs.nlb.LoadBalancerZoneMapping>[]>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * Protocol version. Value:
     * - **ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     */
    addressType: pulumi.Input<string>;
    /**
     * The ID of the shared bandwidth package associated with the public network instance.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     */
    crossZoneEnabled?: pulumi.Input<boolean>;
    /**
     * Delete protection. See `deletionProtectionConfig` below.
     */
    deletionProtectionConfig?: pulumi.Input<inputs.nlb.LoadBalancerDeletionProtectionConfig>;
    /**
     * Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     */
    deletionProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
     */
    deletionProtectionReason?: pulumi.Input<string>;
    /**
     * The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     */
    ipv6AddressType?: pulumi.Input<string>;
    /**
     * The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     */
    loadBalancerName?: pulumi.Input<string>;
    /**
     * Load balancing type. Only value: **network**, which indicates network-based load balancing.
     */
    loadBalancerType?: pulumi.Input<string>;
    /**
     * Modify protection. See `modificationProtectionConfig` below.
     */
    modificationProtectionConfig?: pulumi.Input<inputs.nlb.LoadBalancerModificationProtectionConfig>;
    /**
     * The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
     */
    modificationProtectionReason?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     */
    modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The security group to which the network-based SLB instance belongs.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of labels.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the network-based SLB instance.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
     */
    zoneMappings: pulumi.Input<pulumi.Input<inputs.nlb.LoadBalancerZoneMapping>[]>;
}
