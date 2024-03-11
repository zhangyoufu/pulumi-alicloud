// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **DEPRECATED:** This resource has been renamed to alicloud.ecs.EcsNetworkInterface from version 1.123.1.
 *
 * Provides an ECS Elastic Network Interface resource.
 *
 * For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).
 *
 * > **NOTE** Only one of privateIps or privateIpsCount can be specified when assign private IPs.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "networkInterfaceName";
 * const vpc = new alicloud.vpc.Network("vpc", {
 *     vpcName: name,
 *     cidrBlock: "192.168.0.0/24",
 * });
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const vswitch = new alicloud.vpc.Switch("vswitch", {
 *     cidrBlock: "192.168.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vpcId: vpc.id,
 * });
 * const group = new alicloud.ecs.SecurityGroup("group", {vpcId: vpc.id});
 * const defaultNetworkInterface = new alicloud.vpc.NetworkInterface("defaultNetworkInterface", {
 *     networkInterfaceName: name,
 *     vswitchId: vswitch.id,
 *     securityGroupIds: [group.id],
 *     privateIp: "192.168.0.2",
 *     privateIpsCount: 3,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ENI can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/networkInterface:NetworkInterface eni eni-abc1234567890000
 * ```
 */
export class NetworkInterface extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceState, opts?: pulumi.CustomResourceOptions): NetworkInterface {
        return new NetworkInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/networkInterface:NetworkInterface';

    /**
     * Returns true if the given object is an instance of NetworkInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterface.__pulumiType;
    }

    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly ipv4PrefixCount!: pulumi.Output<number>;
    public readonly ipv4Prefixes!: pulumi.Output<string[]>;
    public readonly ipv6AddressCount!: pulumi.Output<number>;
    public readonly ipv6Addresses!: pulumi.Output<string[]>;
    /**
     * (Available in 1.54.0+) The MAC address of an ENI.
     */
    public /*out*/ readonly mac!: pulumi.Output<string>;
    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
     */
    public readonly name!: pulumi.Output<string>;
    public readonly networkInterfaceName!: pulumi.Output<string>;
    public readonly primaryIpAddress!: pulumi.Output<string>;
    /**
     * The primary private IP of the ENI.
     *
     * @deprecated Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
     */
    public readonly privateIp!: pulumi.Output<string>;
    public readonly privateIpAddresses!: pulumi.Output<string[]>;
    /**
     * List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
     */
    public readonly privateIps!: pulumi.Output<string[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
     */
    public readonly privateIpsCount!: pulumi.Output<number>;
    public readonly queueNumber!: pulumi.Output<number>;
    /**
     * The Id of resource group which the network interface belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    public readonly secondaryPrivateIpAddressCount!: pulumi.Output<number>;
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * A list of security group ids to associate with.
     *
     * @deprecated Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VSwitch to create the ENI in.
     */
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceArgs | NetworkInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipv4PrefixCount"] = state ? state.ipv4PrefixCount : undefined;
            resourceInputs["ipv4Prefixes"] = state ? state.ipv4Prefixes : undefined;
            resourceInputs["ipv6AddressCount"] = state ? state.ipv6AddressCount : undefined;
            resourceInputs["ipv6Addresses"] = state ? state.ipv6Addresses : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkInterfaceName"] = state ? state.networkInterfaceName : undefined;
            resourceInputs["primaryIpAddress"] = state ? state.primaryIpAddress : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["privateIpAddresses"] = state ? state.privateIpAddresses : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["privateIpsCount"] = state ? state.privateIpsCount : undefined;
            resourceInputs["queueNumber"] = state ? state.queueNumber : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["secondaryPrivateIpAddressCount"] = state ? state.secondaryPrivateIpAddressCount : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceArgs | undefined;
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipv4PrefixCount"] = args ? args.ipv4PrefixCount : undefined;
            resourceInputs["ipv4Prefixes"] = args ? args.ipv4Prefixes : undefined;
            resourceInputs["ipv6AddressCount"] = args ? args.ipv6AddressCount : undefined;
            resourceInputs["ipv6Addresses"] = args ? args.ipv6Addresses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkInterfaceName"] = args ? args.networkInterfaceName : undefined;
            resourceInputs["primaryIpAddress"] = args ? args.primaryIpAddress : undefined;
            resourceInputs["privateIp"] = args ? args.privateIp : undefined;
            resourceInputs["privateIpAddresses"] = args ? args.privateIpAddresses : undefined;
            resourceInputs["privateIps"] = args ? args.privateIps : undefined;
            resourceInputs["privateIpsCount"] = args ? args.privateIpsCount : undefined;
            resourceInputs["queueNumber"] = args ? args.queueNumber : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["secondaryPrivateIpAddressCount"] = args ? args.secondaryPrivateIpAddressCount : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["mac"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterface resources.
 */
export interface NetworkInterfaceState {
    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    description?: pulumi.Input<string>;
    ipv4PrefixCount?: pulumi.Input<number>;
    ipv4Prefixes?: pulumi.Input<pulumi.Input<string>[]>;
    ipv6AddressCount?: pulumi.Input<number>;
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Available in 1.54.0+) The MAC address of an ENI.
     */
    mac?: pulumi.Input<string>;
    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
     */
    name?: pulumi.Input<string>;
    networkInterfaceName?: pulumi.Input<string>;
    primaryIpAddress?: pulumi.Input<string>;
    /**
     * The primary private IP of the ENI.
     *
     * @deprecated Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
     */
    privateIp?: pulumi.Input<string>;
    privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
     */
    privateIpsCount?: pulumi.Input<number>;
    queueNumber?: pulumi.Input<number>;
    /**
     * The Id of resource group which the network interface belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    secondaryPrivateIpAddressCount?: pulumi.Input<number>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of security group ids to associate with.
     *
     * @deprecated Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch to create the ENI in.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterface resource.
 */
export interface NetworkInterfaceArgs {
    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    description?: pulumi.Input<string>;
    ipv4PrefixCount?: pulumi.Input<number>;
    ipv4Prefixes?: pulumi.Input<pulumi.Input<string>[]>;
    ipv6AddressCount?: pulumi.Input<number>;
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
     */
    name?: pulumi.Input<string>;
    networkInterfaceName?: pulumi.Input<string>;
    primaryIpAddress?: pulumi.Input<string>;
    /**
     * The primary private IP of the ENI.
     *
     * @deprecated Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
     */
    privateIp?: pulumi.Input<string>;
    privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
     */
    privateIpsCount?: pulumi.Input<number>;
    queueNumber?: pulumi.Input<number>;
    /**
     * The Id of resource group which the network interface belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    secondaryPrivateIpAddressCount?: pulumi.Input<number>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of security group ids to associate with.
     *
     * @deprecated Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch to create the ENI in.
     */
    vswitchId: pulumi.Input<string>;
}
