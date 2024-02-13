// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a snat resource.
 *
 * > **NOTE:** Available since v1.119.0.
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
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultNatGateway = new alicloud.vpc.NatGateway("defaultNatGateway", {
 *     vpcId: defaultNetwork.id,
 *     natGatewayName: name,
 *     paymentType: "PayAsYouGo",
 *     vswitchId: defaultSwitch.id,
 *     natType: "Enhanced",
 * });
 * const defaultEipAddress = new alicloud.ecs.EipAddress("defaultEipAddress", {addressName: name});
 * const defaultEipAssociation = new alicloud.ecs.EipAssociation("defaultEipAssociation", {
 *     allocationId: defaultEipAddress.id,
 *     instanceId: defaultNatGateway.id,
 * });
 * const defaultSnatEntry = new alicloud.vpc.SnatEntry("defaultSnatEntry", {
 *     snatTableId: defaultNatGateway.snatTableIds,
 *     sourceVswitchId: defaultSwitch.id,
 *     snatIp: defaultEipAddress.ipAddress,
 * });
 * ```
 *
 * ## Import
 *
 * Snat Entry can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/snatEntry:SnatEntry foo stb-1aece3:snat-232ce2
 * ```
 */
export class SnatEntry extends pulumi.CustomResource {
    /**
     * Get an existing SnatEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnatEntryState, opts?: pulumi.CustomResourceOptions): SnatEntry {
        return new SnatEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/snatEntry:SnatEntry';

    /**
     * Returns true if the given object is an instance of SnatEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnatEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnatEntry.__pulumiType;
    }

    /**
     * The id of the snat entry on the server.
     */
    public /*out*/ readonly snatEntryId!: pulumi.Output<string>;
    /**
     * The name of snat entry.
     */
    public readonly snatEntryName!: pulumi.Output<string | undefined>;
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidthPackages`.
     */
    public readonly snatIp!: pulumi.Output<string>;
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes "snatTableIds".
     */
    public readonly snatTableId!: pulumi.Output<string>;
    /**
     * The private network segment of Ecs. This parameter and the `sourceVswitchId` parameter are mutually exclusive and cannot appear at the same time.
     */
    public readonly sourceCidr!: pulumi.Output<string>;
    /**
     * The vswitch ID.
     */
    public readonly sourceVswitchId!: pulumi.Output<string>;
    /**
     * (Available since v1.119.1) The status of snat entry.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a SnatEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnatEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnatEntryArgs | SnatEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnatEntryState | undefined;
            resourceInputs["snatEntryId"] = state ? state.snatEntryId : undefined;
            resourceInputs["snatEntryName"] = state ? state.snatEntryName : undefined;
            resourceInputs["snatIp"] = state ? state.snatIp : undefined;
            resourceInputs["snatTableId"] = state ? state.snatTableId : undefined;
            resourceInputs["sourceCidr"] = state ? state.sourceCidr : undefined;
            resourceInputs["sourceVswitchId"] = state ? state.sourceVswitchId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SnatEntryArgs | undefined;
            if ((!args || args.snatIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snatIp'");
            }
            if ((!args || args.snatTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snatTableId'");
            }
            resourceInputs["snatEntryName"] = args ? args.snatEntryName : undefined;
            resourceInputs["snatIp"] = args ? args.snatIp : undefined;
            resourceInputs["snatTableId"] = args ? args.snatTableId : undefined;
            resourceInputs["sourceCidr"] = args ? args.sourceCidr : undefined;
            resourceInputs["sourceVswitchId"] = args ? args.sourceVswitchId : undefined;
            resourceInputs["snatEntryId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnatEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnatEntry resources.
 */
export interface SnatEntryState {
    /**
     * The id of the snat entry on the server.
     */
    snatEntryId?: pulumi.Input<string>;
    /**
     * The name of snat entry.
     */
    snatEntryName?: pulumi.Input<string>;
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidthPackages`.
     */
    snatIp?: pulumi.Input<string>;
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes "snatTableIds".
     */
    snatTableId?: pulumi.Input<string>;
    /**
     * The private network segment of Ecs. This parameter and the `sourceVswitchId` parameter are mutually exclusive and cannot appear at the same time.
     */
    sourceCidr?: pulumi.Input<string>;
    /**
     * The vswitch ID.
     */
    sourceVswitchId?: pulumi.Input<string>;
    /**
     * (Available since v1.119.1) The status of snat entry.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnatEntry resource.
 */
export interface SnatEntryArgs {
    /**
     * The name of snat entry.
     */
    snatEntryName?: pulumi.Input<string>;
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidthPackages`.
     */
    snatIp: pulumi.Input<string>;
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes "snatTableIds".
     */
    snatTableId: pulumi.Input<string>;
    /**
     * The private network segment of Ecs. This parameter and the `sourceVswitchId` parameter are mutually exclusive and cannot appear at the same time.
     */
    sourceCidr?: pulumi.Input<string>;
    /**
     * The vswitch ID.
     */
    sourceVswitchId?: pulumi.Input<string>;
}
