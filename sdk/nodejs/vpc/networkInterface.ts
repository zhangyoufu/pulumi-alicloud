// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * ENI can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/networkInterface:NetworkInterface eni eni-abc1234567890000
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["mac"] = state ? state.mac : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkInterfaceName"] = state ? state.networkInterfaceName : undefined;
            inputs["primaryIpAddress"] = state ? state.primaryIpAddress : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["privateIpAddresses"] = state ? state.privateIpAddresses : undefined;
            inputs["privateIps"] = state ? state.privateIps : undefined;
            inputs["privateIpsCount"] = state ? state.privateIpsCount : undefined;
            inputs["queueNumber"] = state ? state.queueNumber : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["secondaryPrivateIpAddressCount"] = state ? state.secondaryPrivateIpAddressCount : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceArgs | undefined;
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaceName"] = args ? args.networkInterfaceName : undefined;
            inputs["primaryIpAddress"] = args ? args.primaryIpAddress : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["privateIpAddresses"] = args ? args.privateIpAddresses : undefined;
            inputs["privateIps"] = args ? args.privateIps : undefined;
            inputs["privateIpsCount"] = args ? args.privateIpsCount : undefined;
            inputs["queueNumber"] = args ? args.queueNumber : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["secondaryPrivateIpAddressCount"] = args ? args.secondaryPrivateIpAddressCount : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["mac"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterface resources.
 */
export interface NetworkInterfaceState {
    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * (Available in 1.54.0+) The MAC address of an ENI.
     */
    readonly mac?: pulumi.Input<string>;
    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
     */
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaceName?: pulumi.Input<string>;
    readonly primaryIpAddress?: pulumi.Input<string>;
    /**
     * The primary private IP of the ENI.
     *
     * @deprecated Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
     */
    readonly privateIp?: pulumi.Input<string>;
    readonly privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
     */
    readonly privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
     */
    readonly privateIpsCount?: pulumi.Input<number>;
    readonly queueNumber?: pulumi.Input<number>;
    /**
     * The Id of resource group which the network interface belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    readonly secondaryPrivateIpAddressCount?: pulumi.Input<number>;
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of security group ids to associate with.
     *
     * @deprecated Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch to create the ENI in.
     */
    readonly vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterface resource.
 */
export interface NetworkInterfaceArgs {
    /**
     * Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
     */
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaceName?: pulumi.Input<string>;
    readonly primaryIpAddress?: pulumi.Input<string>;
    /**
     * The primary private IP of the ENI.
     *
     * @deprecated Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
     */
    readonly privateIp?: pulumi.Input<string>;
    readonly privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
     */
    readonly privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
     *
     * @deprecated Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
     */
    readonly privateIpsCount?: pulumi.Input<number>;
    readonly queueNumber?: pulumi.Input<number>;
    /**
     * The Id of resource group which the network interface belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    readonly secondaryPrivateIpAddressCount?: pulumi.Input<number>;
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of security group ids to associate with.
     *
     * @deprecated Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch to create the ENI in.
     */
    readonly vswitchId: pulumi.Input<string>;
}
