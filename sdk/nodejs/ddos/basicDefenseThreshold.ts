// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Ddos Basic defense threshold resource.
 *
 * For information about Ddos Basic Antiddos and how to use it, see [What is Defense Threshold](https://www.alibabacloud.com/help/en/ddos-protection/latest/modifydefensethreshold).
 *
 * > **NOTE:** Available since v1.168.0.
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
 * const name = config.get("name") || "tf-example";
 * const defaultEipAddress = new alicloud.ecs.EipAddress("defaultEipAddress", {
 *     addressName: name,
 *     isp: "BGP",
 *     internetChargeType: "PayByBandwidth",
 *     paymentType: "PayAsYouGo",
 * });
 * const defaultBasicDefenseThreshold = new alicloud.ddos.BasicDefenseThreshold("defaultBasicDefenseThreshold", {
 *     instanceId: defaultEipAddress.id,
 *     ddosType: "defense",
 *     instanceType: "eip",
 *     bps: 390,
 *     pps: 90000,
 * });
 * ```
 *
 * ## Import
 *
 * Ddos Basic Antiddos can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold example <instance_id>:<instance_type>:<ddos_type>
 * ```
 */
export class BasicDefenseThreshold extends pulumi.CustomResource {
    /**
     * Get an existing BasicDefenseThreshold resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasicDefenseThresholdState, opts?: pulumi.CustomResourceOptions): BasicDefenseThreshold {
        return new BasicDefenseThreshold(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold';

    /**
     * Returns true if the given object is an instance of BasicDefenseThreshold.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasicDefenseThreshold {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasicDefenseThreshold.__pulumiType;
    }

    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     */
    public readonly bps!: pulumi.Output<number>;
    /**
     * The type of the threshold to query. Valid values: `defense`,`blackhole`.
     */
    public readonly ddosType!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The Internet IP address.
     */
    public readonly internetIp!: pulumi.Output<string>;
    /**
     * Whether it is the system default threshold. Value:
     */
    public readonly isAuto!: pulumi.Output<boolean>;
    /**
     * The maximum traffic scrubbing threshold. Unit: Mbit/s.
     */
    public /*out*/ readonly maxBps!: pulumi.Output<number>;
    /**
     * The maximum packet scrubbing threshold. Unit: pps.
     */
    public /*out*/ readonly maxPps!: pulumi.Output<number>;
    /**
     * The current message number cleaning threshold. Unit: pps.
     */
    public readonly pps!: pulumi.Output<number>;

    /**
     * Create a BasicDefenseThreshold resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasicDefenseThresholdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasicDefenseThresholdArgs | BasicDefenseThresholdState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BasicDefenseThresholdState | undefined;
            resourceInputs["bps"] = state ? state.bps : undefined;
            resourceInputs["ddosType"] = state ? state.ddosType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetIp"] = state ? state.internetIp : undefined;
            resourceInputs["isAuto"] = state ? state.isAuto : undefined;
            resourceInputs["maxBps"] = state ? state.maxBps : undefined;
            resourceInputs["maxPps"] = state ? state.maxPps : undefined;
            resourceInputs["pps"] = state ? state.pps : undefined;
        } else {
            const args = argsOrState as BasicDefenseThresholdArgs | undefined;
            if ((!args || args.ddosType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ddosType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["bps"] = args ? args.bps : undefined;
            resourceInputs["ddosType"] = args ? args.ddosType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetIp"] = args ? args.internetIp : undefined;
            resourceInputs["isAuto"] = args ? args.isAuto : undefined;
            resourceInputs["pps"] = args ? args.pps : undefined;
            resourceInputs["maxBps"] = undefined /*out*/;
            resourceInputs["maxPps"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BasicDefenseThreshold.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasicDefenseThreshold resources.
 */
export interface BasicDefenseThresholdState {
    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     */
    bps?: pulumi.Input<number>;
    /**
     * The type of the threshold to query. Valid values: `defense`,`blackhole`.
     */
    ddosType?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The Internet IP address.
     */
    internetIp?: pulumi.Input<string>;
    /**
     * Whether it is the system default threshold. Value:
     */
    isAuto?: pulumi.Input<boolean>;
    /**
     * The maximum traffic scrubbing threshold. Unit: Mbit/s.
     */
    maxBps?: pulumi.Input<number>;
    /**
     * The maximum packet scrubbing threshold. Unit: pps.
     */
    maxPps?: pulumi.Input<number>;
    /**
     * The current message number cleaning threshold. Unit: pps.
     */
    pps?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BasicDefenseThreshold resource.
 */
export interface BasicDefenseThresholdArgs {
    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     */
    bps?: pulumi.Input<number>;
    /**
     * The type of the threshold to query. Valid values: `defense`,`blackhole`.
     */
    ddosType: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The Internet IP address.
     */
    internetIp?: pulumi.Input<string>;
    /**
     * Whether it is the system default threshold. Value:
     */
    isAuto?: pulumi.Input<boolean>;
    /**
     * The current message number cleaning threshold. Unit: pps.
     */
    pps?: pulumi.Input<number>;
}
