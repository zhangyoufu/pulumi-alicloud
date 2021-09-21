// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Bastion Host Host resource.
 *
 * For information about Bastion Host Host and how to use it, see [What is Host](https://www.alibabacloud.com/help/en/doc-detail/201330.htm).
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.bastionhost.Host("example", {
 *     activeAddressType: "Private",
 *     hostName: "example_value",
 *     hostPrivateAddress: "172.16.0.10",
 *     instanceId: "bastionhost-cn-tl3xxxxxxx",
 *     osType: "Linux",
 *     source: "Local",
 * });
 * ```
 *
 * ## Import
 *
 * Bastion Host Host can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:bastionhost/host:Host example <instance_id>:<host_id>
 * ```
 */
export class Host extends pulumi.CustomResource {
    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostState, opts?: pulumi.CustomResourceOptions): Host {
        return new Host(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:bastionhost/host:Host';

    /**
     * Returns true if the given object is an instance of Host.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Host {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Host.__pulumiType;
    }

    /**
     * Specify the new create a host of address types. Valid values: Public: the IP address of a Public network Private: Private network address.
     */
    public readonly activeAddressType!: pulumi.Output<string>;
    /**
     * Specify a host of notes, supports up to 500 characters.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The host ID.
     */
    public /*out*/ readonly hostId!: pulumi.Output<string>;
    /**
     * Specify the new create a host name of the supports up to 128 characters.
     */
    public readonly hostName!: pulumi.Output<string>;
    /**
     * Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS.
     */
    public readonly hostPrivateAddress!: pulumi.Output<string | undefined>;
    /**
     * Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
     */
    public readonly hostPublicAddress!: pulumi.Output<string | undefined>;
    /**
     * Specify the new create a host where the Bastion host ID of.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The instance region id.
     */
    public readonly instanceRegionId!: pulumi.Output<string | undefined>;
    /**
     * Specify the new create the host's operating system. Valid values: Linux Windows.
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * Specify the new create a host of source. Valid values: 
     * * Local: localhost
     * * Ecs:ECS instance
     * * Rds:RDS exclusive cluster host.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Specify the newly created ECS instance ID or dedicated cluster host ID.
     */
    public readonly sourceInstanceId!: pulumi.Output<string | undefined>;

    /**
     * Create a Host resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostArgs | HostState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostState | undefined;
            inputs["activeAddressType"] = state ? state.activeAddressType : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["hostId"] = state ? state.hostId : undefined;
            inputs["hostName"] = state ? state.hostName : undefined;
            inputs["hostPrivateAddress"] = state ? state.hostPrivateAddress : undefined;
            inputs["hostPublicAddress"] = state ? state.hostPublicAddress : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["instanceRegionId"] = state ? state.instanceRegionId : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["sourceInstanceId"] = state ? state.sourceInstanceId : undefined;
        } else {
            const args = argsOrState as HostArgs | undefined;
            if ((!args || args.activeAddressType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'activeAddressType'");
            }
            if ((!args || args.hostName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.osType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'osType'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            inputs["activeAddressType"] = args ? args.activeAddressType : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["hostName"] = args ? args.hostName : undefined;
            inputs["hostPrivateAddress"] = args ? args.hostPrivateAddress : undefined;
            inputs["hostPublicAddress"] = args ? args.hostPublicAddress : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["instanceRegionId"] = args ? args.instanceRegionId : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["sourceInstanceId"] = args ? args.sourceInstanceId : undefined;
            inputs["hostId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Host.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Host resources.
 */
export interface HostState {
    /**
     * Specify the new create a host of address types. Valid values: Public: the IP address of a Public network Private: Private network address.
     */
    readonly activeAddressType?: pulumi.Input<string>;
    /**
     * Specify a host of notes, supports up to 500 characters.
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * The host ID.
     */
    readonly hostId?: pulumi.Input<string>;
    /**
     * Specify the new create a host name of the supports up to 128 characters.
     */
    readonly hostName?: pulumi.Input<string>;
    /**
     * Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS.
     */
    readonly hostPrivateAddress?: pulumi.Input<string>;
    /**
     * Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
     */
    readonly hostPublicAddress?: pulumi.Input<string>;
    /**
     * Specify the new create a host where the Bastion host ID of.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The instance region id.
     */
    readonly instanceRegionId?: pulumi.Input<string>;
    /**
     * Specify the new create the host's operating system. Valid values: Linux Windows.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * Specify the new create a host of source. Valid values: 
     * * Local: localhost
     * * Ecs:ECS instance
     * * Rds:RDS exclusive cluster host.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * Specify the newly created ECS instance ID or dedicated cluster host ID.
     */
    readonly sourceInstanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Host resource.
 */
export interface HostArgs {
    /**
     * Specify the new create a host of address types. Valid values: Public: the IP address of a Public network Private: Private network address.
     */
    readonly activeAddressType: pulumi.Input<string>;
    /**
     * Specify a host of notes, supports up to 500 characters.
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Specify the new create a host name of the supports up to 128 characters.
     */
    readonly hostName: pulumi.Input<string>;
    /**
     * Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS.
     */
    readonly hostPrivateAddress?: pulumi.Input<string>;
    /**
     * Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
     */
    readonly hostPublicAddress?: pulumi.Input<string>;
    /**
     * Specify the new create a host where the Bastion host ID of.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The instance region id.
     */
    readonly instanceRegionId?: pulumi.Input<string>;
    /**
     * Specify the new create the host's operating system. Valid values: Linux Windows.
     */
    readonly osType: pulumi.Input<string>;
    /**
     * Specify the new create a host of source. Valid values: 
     * * Local: localhost
     * * Ecs:ECS instance
     * * Rds:RDS exclusive cluster host.
     */
    readonly source: pulumi.Input<string>;
    /**
     * Specify the newly created ECS instance ID or dedicated cluster host ID.
     */
    readonly sourceInstanceId?: pulumi.Input<string>;
}
