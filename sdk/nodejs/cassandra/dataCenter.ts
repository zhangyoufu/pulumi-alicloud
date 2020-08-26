// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cassandra dataCenter resource supports replica set dataCenters only. The Cassandra provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/product/49055.htm).
 *
 * > **NOTE:**  Available in 1.88.0+.
 *
 * > **NOTE:**  Create a cassandra dataCenter need a clusterId,so need create a cassandra cluster first.
 *
 * > **NOTE:**  The following regions support create Vpc network Cassandra cluster.
 * The official website mark  more regions. Or you can call [DescribeRegions](https://help.aliyun.com/document_detail/157540.html).
 *
 * > **NOTE:**  Create Cassandra dataCenter or change dataCenter type and storage would cost 30 minutes. Please make full preparation.
 *
 * ## Example Usage
 */
export class DataCenter extends pulumi.CustomResource {
    /**
     * Get an existing DataCenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataCenterState, opts?: pulumi.CustomResourceOptions): DataCenter {
        return new DataCenter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cassandra/dataCenter:DataCenter';

    /**
     * Returns true if the given object is an instance of DataCenter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataCenter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataCenter.__pulumiType;
    }

    /**
     * Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when payType = Subscription.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when payType = Subscription. Unit: month.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * Cassandra cluster id of dataCenter-2 belongs to.
     */
    public readonly clusterId!: pulumi.Output<string>;
    public /*out*/ readonly dataCenterId!: pulumi.Output<string>;
    /**
     * Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     */
    public readonly dataCenterName!: pulumi.Output<string | undefined>;
    /**
     * User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
     * - Custom storage space; value range: [160, 2000].
     * - 80-GB increments.
     */
    public readonly diskSize!: pulumi.Output<number | undefined>;
    /**
     * The disk type of Cassandra dataCenter-2. Valid values are `cloudSsd`, `cloudEfficiency`, `localHddPro`, `localSsdPro`, localDisk size is fixed.
     */
    public readonly diskType!: pulumi.Output<string | undefined>;
    public readonly enablePublic!: pulumi.Output<boolean | undefined>;
    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The node count of Cassandra dataCenter-2, default to 2.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
     */
    public readonly payType!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<number | undefined>;
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    public /*out*/ readonly publicPoints!: pulumi.Output<string[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The vswitchId of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The Zone to launch the Cassandra dataCenter-2. If vswitchId is not empty, this zoneId can be "" or consistent.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a DataCenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataCenterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataCenterArgs | DataCenterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataCenterState | undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["dataCenterId"] = state ? state.dataCenterId : undefined;
            inputs["dataCenterName"] = state ? state.dataCenterName : undefined;
            inputs["diskSize"] = state ? state.diskSize : undefined;
            inputs["diskType"] = state ? state.diskType : undefined;
            inputs["enablePublic"] = state ? state.enablePublic : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["payType"] = state ? state.payType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["periodUnit"] = state ? state.periodUnit : undefined;
            inputs["publicPoints"] = state ? state.publicPoints : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as DataCenterArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            if (!args || args.nodeCount === undefined) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if (!args || args.payType === undefined) {
                throw new Error("Missing required property 'payType'");
            }
            if (!args || args.vswitchId === undefined) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["dataCenterName"] = args ? args.dataCenterName : undefined;
            inputs["diskSize"] = args ? args.diskSize : undefined;
            inputs["diskType"] = args ? args.diskType : undefined;
            inputs["enablePublic"] = args ? args.enablePublic : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["payType"] = args ? args.payType : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["periodUnit"] = args ? args.periodUnit : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["dataCenterId"] = undefined /*out*/;
            inputs["publicPoints"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataCenter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataCenter resources.
 */
export interface DataCenterState {
    /**
     * Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when payType = Subscription.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when payType = Subscription. Unit: month.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * Cassandra cluster id of dataCenter-2 belongs to.
     */
    readonly clusterId?: pulumi.Input<string>;
    readonly dataCenterId?: pulumi.Input<string>;
    /**
     * Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     */
    readonly dataCenterName?: pulumi.Input<string>;
    /**
     * User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
     * - Custom storage space; value range: [160, 2000].
     * - 80-GB increments.
     */
    readonly diskSize?: pulumi.Input<number>;
    /**
     * The disk type of Cassandra dataCenter-2. Valid values are `cloudSsd`, `cloudEfficiency`, `localHddPro`, `localSsdPro`, localDisk size is fixed.
     */
    readonly diskType?: pulumi.Input<string>;
    readonly enablePublic?: pulumi.Input<boolean>;
    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The node count of Cassandra dataCenter-2, default to 2.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
     */
    readonly payType?: pulumi.Input<string>;
    readonly period?: pulumi.Input<number>;
    readonly periodUnit?: pulumi.Input<string>;
    readonly publicPoints?: pulumi.Input<pulumi.Input<string>[]>;
    readonly status?: pulumi.Input<string>;
    /**
     * The vswitchId of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the Cassandra dataCenter-2. If vswitchId is not empty, this zoneId can be "" or consistent.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataCenter resource.
 */
export interface DataCenterArgs {
    /**
     * Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when payType = Subscription.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when payType = Subscription. Unit: month.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * Cassandra cluster id of dataCenter-2 belongs to.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
     */
    readonly dataCenterName?: pulumi.Input<string>;
    /**
     * User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
     * - Custom storage space; value range: [160, 2000].
     * - 80-GB increments.
     */
    readonly diskSize?: pulumi.Input<number>;
    /**
     * The disk type of Cassandra dataCenter-2. Valid values are `cloudSsd`, `cloudEfficiency`, `localHddPro`, `localSsdPro`, localDisk size is fixed.
     */
    readonly diskType?: pulumi.Input<string>;
    readonly enablePublic?: pulumi.Input<boolean>;
    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
     */
    readonly instanceType: pulumi.Input<string>;
    /**
     * The node count of Cassandra dataCenter-2, default to 2.
     */
    readonly nodeCount: pulumi.Input<number>;
    /**
     * The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
     */
    readonly payType: pulumi.Input<string>;
    readonly period?: pulumi.Input<number>;
    readonly periodUnit?: pulumi.Input<string>;
    /**
     * The vswitchId of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
     */
    readonly vswitchId: pulumi.Input<string>;
    /**
     * The Zone to launch the Cassandra dataCenter-2. If vswitchId is not empty, this zoneId can be "" or consistent.
     */
    readonly zoneId?: pulumi.Input<string>;
}
