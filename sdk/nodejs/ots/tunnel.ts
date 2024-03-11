// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an OTS tunnel resource.
 *
 * For information about OTS tunnel and how to use it, see [Tunnel overview](https://www.alibabacloud.com/help/en/tablestore/latest/tunnel-service-overview).
 *
 * > **NOTE:** Available since v1.172.0.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultRandomInteger = new random.RandomInteger("defaultRandomInteger", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultInstance = new alicloud.ots.Instance("defaultInstance", {
 *     description: name,
 *     accessedBy: "Any",
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * const defaultTable = new alicloud.ots.Table("defaultTable", {
 *     instanceName: defaultInstance.name,
 *     tableName: "tf_example",
 *     timeToLive: -1,
 *     maxVersion: 1,
 *     enableSse: true,
 *     sseKeyType: "SSE_KMS_SERVICE",
 *     primaryKeys: [
 *         {
 *             name: "pk1",
 *             type: "Integer",
 *         },
 *         {
 *             name: "pk2",
 *             type: "String",
 *         },
 *         {
 *             name: "pk3",
 *             type: "Binary",
 *         },
 *     ],
 * });
 * const defaultTunnel = new alicloud.ots.Tunnel("defaultTunnel", {
 *     instanceName: defaultInstance.name,
 *     tableName: defaultTable.tableName,
 *     tunnelName: "tf_example",
 *     tunnelType: "BaseAndStream",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * OTS tunnel can be imported using id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ots/tunnel:Tunnel foo <instance_name>:<table_name>:<tunnel_name>
 * ```
 */
export class Tunnel extends pulumi.CustomResource {
    /**
     * Get an existing Tunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TunnelState, opts?: pulumi.CustomResourceOptions): Tunnel {
        return new Tunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ots/tunnel:Tunnel';

    /**
     * Returns true if the given object is an instance of Tunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tunnel.__pulumiType;
    }

    /**
     * The channels of OTS tunnel. Each element contains the following attributes:
     */
    public /*out*/ readonly channels!: pulumi.Output<outputs.ots.TunnelChannel[]>;
    /**
     * The creation time of the Tunnel.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * Whether the tunnel has expired.
     */
    public /*out*/ readonly expired!: pulumi.Output<boolean>;
    /**
     * The name of the OTS instance in which table will located.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The name of the OTS table. If changed, a new table would be created.
     */
    public readonly tableName!: pulumi.Output<string>;
    /**
     * The tunnel id of the OTS which could not be changed.
     */
    public /*out*/ readonly tunnelId!: pulumi.Output<string>;
    /**
     * The name of the OTS tunnel. If changed, a new tunnel would be created.
     */
    public readonly tunnelName!: pulumi.Output<string>;
    /**
     * The latest consumption time of the tunnel, unix time in nanosecond.
     */
    public /*out*/ readonly tunnelRpo!: pulumi.Output<number>;
    /**
     * The stage of OTS tunnel, valid values: `InitBaseDataAndStreamShard`, `ProcessBaseData`, `ProcessStream`.
     */
    public /*out*/ readonly tunnelStage!: pulumi.Output<string>;
    /**
     * The type of the OTS tunnel. Only `BaseAndStream`, `BaseData` or `Stream` is allowed.
     */
    public readonly tunnelType!: pulumi.Output<string>;

    /**
     * Create a Tunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TunnelArgs | TunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TunnelState | undefined;
            resourceInputs["channels"] = state ? state.channels : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["expired"] = state ? state.expired : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["tableName"] = state ? state.tableName : undefined;
            resourceInputs["tunnelId"] = state ? state.tunnelId : undefined;
            resourceInputs["tunnelName"] = state ? state.tunnelName : undefined;
            resourceInputs["tunnelRpo"] = state ? state.tunnelRpo : undefined;
            resourceInputs["tunnelStage"] = state ? state.tunnelStage : undefined;
            resourceInputs["tunnelType"] = state ? state.tunnelType : undefined;
        } else {
            const args = argsOrState as TunnelArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            if ((!args || args.tunnelName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tunnelName'");
            }
            if ((!args || args.tunnelType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tunnelType'");
            }
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
            resourceInputs["tunnelName"] = args ? args.tunnelName : undefined;
            resourceInputs["tunnelType"] = args ? args.tunnelType : undefined;
            resourceInputs["channels"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expired"] = undefined /*out*/;
            resourceInputs["tunnelId"] = undefined /*out*/;
            resourceInputs["tunnelRpo"] = undefined /*out*/;
            resourceInputs["tunnelStage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tunnel resources.
 */
export interface TunnelState {
    /**
     * The channels of OTS tunnel. Each element contains the following attributes:
     */
    channels?: pulumi.Input<pulumi.Input<inputs.ots.TunnelChannel>[]>;
    /**
     * The creation time of the Tunnel.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Whether the tunnel has expired.
     */
    expired?: pulumi.Input<boolean>;
    /**
     * The name of the OTS instance in which table will located.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The name of the OTS table. If changed, a new table would be created.
     */
    tableName?: pulumi.Input<string>;
    /**
     * The tunnel id of the OTS which could not be changed.
     */
    tunnelId?: pulumi.Input<string>;
    /**
     * The name of the OTS tunnel. If changed, a new tunnel would be created.
     */
    tunnelName?: pulumi.Input<string>;
    /**
     * The latest consumption time of the tunnel, unix time in nanosecond.
     */
    tunnelRpo?: pulumi.Input<number>;
    /**
     * The stage of OTS tunnel, valid values: `InitBaseDataAndStreamShard`, `ProcessBaseData`, `ProcessStream`.
     */
    tunnelStage?: pulumi.Input<string>;
    /**
     * The type of the OTS tunnel. Only `BaseAndStream`, `BaseData` or `Stream` is allowed.
     */
    tunnelType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tunnel resource.
 */
export interface TunnelArgs {
    /**
     * The name of the OTS instance in which table will located.
     */
    instanceName: pulumi.Input<string>;
    /**
     * The name of the OTS table. If changed, a new table would be created.
     */
    tableName: pulumi.Input<string>;
    /**
     * The name of the OTS tunnel. If changed, a new tunnel would be created.
     */
    tunnelName: pulumi.Input<string>;
    /**
     * The type of the OTS tunnel. Only `BaseAndStream`, `BaseData` or `Stream` is allowed.
     */
    tunnelType: pulumi.Input<string>;
}
