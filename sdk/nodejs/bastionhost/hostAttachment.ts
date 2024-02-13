// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Bastion Host Host Attachment resource to add host into one host group.
 *
 * > **NOTE:** Available since v1.135.0.
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
 * const name = config.get("name") || "tf_example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetworks.ids?.[0],
 *     zoneId: defaultZones.zones?.[0]?.id,
 * }));
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?.[0])});
 * const defaultInstance = new alicloud.bastionhost.Instance("defaultInstance", {
 *     description: name,
 *     licenseCode: "bhah_ent_50_asset",
 *     planCode: "cloudbastion",
 *     storage: "5",
 *     bandwidth: "5",
 *     period: 1,
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?.[0]),
 *     securityGroupIds: [defaultSecurityGroup.id],
 * });
 * const defaultHostGroup = new alicloud.bastionhost.HostGroup("defaultHostGroup", {
 *     hostGroupName: name,
 *     instanceId: defaultInstance.id,
 * });
 * const defaultHost = new alicloud.bastionhost.Host("defaultHost", {
 *     instanceId: defaultInstance.id,
 *     hostName: name,
 *     activeAddressType: "Private",
 *     hostPrivateAddress: "172.16.0.10",
 *     osType: "Linux",
 *     source: "Local",
 * });
 * const defaultHostAttachment = new alicloud.bastionhost.HostAttachment("defaultHostAttachment", {
 *     hostGroupId: defaultHostGroup.hostGroupId,
 *     hostId: defaultHost.hostId,
 *     instanceId: defaultInstance.id,
 * });
 * ```
 *
 * ## Import
 *
 * Bastion Host Host Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:bastionhost/hostAttachment:HostAttachment example <instance_id>:<host_group_id>:<host_id>
 * ```
 */
export class HostAttachment extends pulumi.CustomResource {
    /**
     * Get an existing HostAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostAttachmentState, opts?: pulumi.CustomResourceOptions): HostAttachment {
        return new HostAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:bastionhost/hostAttachment:HostAttachment';

    /**
     * Returns true if the given object is an instance of HostAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostAttachment.__pulumiType;
    }

    /**
     * Specifies the added to the host group ID.
     */
    public readonly hostGroupId!: pulumi.Output<string>;
    /**
     * Specified to be part of a host group of host ID.
     */
    public readonly hostId!: pulumi.Output<string>;
    /**
     * The bastion host instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a HostAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostAttachmentArgs | HostAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostAttachmentState | undefined;
            resourceInputs["hostGroupId"] = state ? state.hostGroupId : undefined;
            resourceInputs["hostId"] = state ? state.hostId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as HostAttachmentArgs | undefined;
            if ((!args || args.hostGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostGroupId'");
            }
            if ((!args || args.hostId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["hostGroupId"] = args ? args.hostGroupId : undefined;
            resourceInputs["hostId"] = args ? args.hostId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostAttachment resources.
 */
export interface HostAttachmentState {
    /**
     * Specifies the added to the host group ID.
     */
    hostGroupId?: pulumi.Input<string>;
    /**
     * Specified to be part of a host group of host ID.
     */
    hostId?: pulumi.Input<string>;
    /**
     * The bastion host instance id.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostAttachment resource.
 */
export interface HostAttachmentArgs {
    /**
     * Specifies the added to the host group ID.
     */
    hostGroupId: pulumi.Input<string>;
    /**
     * Specified to be part of a host group of host ID.
     */
    hostId: pulumi.Input<string>;
    /**
     * The bastion host instance id.
     */
    instanceId: pulumi.Input<string>;
}
