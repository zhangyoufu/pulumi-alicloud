// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a AnalyticDB for PostgreSQL instance resource supports replica set instances only. the AnalyticDB for PostgreSQL provides stable, reliable, and automatic scalable database services.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/api-gpdb-2016-05-03-createdbinstance)
 *
 * > **NOTE:** Available since v1.47.0.
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
 * const default = alicloud.resourcemanager.getResourceGroups({});
 * const defaultGetZones = alicloud.gpdb.getZones({});
 * const defaultGetNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 * });
 * const defaultGetSwitches = Promise.all([defaultGetNetworks, defaultGetZones]).then(([defaultGetNetworks, defaultGetZones]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultGetNetworks.ids?.[0],
 *     zoneId: defaultGetZones.ids?.[0],
 * }));
 * const defaultInstance = new alicloud.gpdb.Instance("default", {
 *     dbInstanceCategory: "HighAvailability",
 *     dbInstanceClass: "gpdb.group.segsdx1",
 *     dbInstanceMode: "StorageElastic",
 *     description: name,
 *     engine: "gpdb",
 *     engineVersion: "6.0",
 *     zoneId: defaultGetZones.then(defaultGetZones => defaultGetZones.ids?.[0]),
 *     instanceNetworkType: "VPC",
 *     instanceSpec: "2C16G",
 *     paymentType: "PayAsYouGo",
 *     segStorageType: "cloud_essd",
 *     segNodeNum: 4,
 *     storageSize: 50,
 *     vpcId: defaultGetNetworks.then(defaultGetNetworks => defaultGetNetworks.ids?.[0]),
 *     vswitchId: defaultGetSwitches.then(defaultGetSwitches => defaultGetSwitches.ids?.[0]),
 *     ipWhitelists: [{
 *         securityIpList: "127.0.0.1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * AnalyticDB for PostgreSQL can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:gpdb/instance:Instance example <id>
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:gpdb/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Field `availabilityZone` has been deprecated from provider version 1.187.0. New field `zoneId` instead.
     *
     * @deprecated Field 'availability_zone' has been deprecated from version 1.187.0. Use 'zone_id' instead.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * (Available since v1.196.0) The connection string of the instance.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Whether to load the sample dataset after the instance is created. Valid values: `true`, `false`.
     */
    public readonly createSampleData!: pulumi.Output<boolean>;
    /**
     * The db instance category. Valid values: `Basic`, `HighAvailability`.
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    public readonly dbInstanceCategory!: pulumi.Output<string>;
    /**
     * The db instance class. see [Instance specifications](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/instance-types).
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    public readonly dbInstanceClass!: pulumi.Output<string | undefined>;
    /**
     * The db instance mode. Valid values: `StorageElastic`, `Serverless`, `Classic`.
     */
    public readonly dbInstanceMode!: pulumi.Output<string>;
    /**
     * The description of the instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the encryption key.
     * > **NOTE:** If `encryptionType` is set to `CloudDisk`, you must specify an encryption key that resides in the same region as the cloud disk that is specified by EncryptionType. Otherwise, leave this parameter empty.
     */
    public readonly encryptionKey!: pulumi.Output<string | undefined>;
    /**
     * The encryption type. Valid values: `CloudDisk`.
     * > **NOTE:** Disk encryption cannot be disabled after it is enabled.
     */
    public readonly encryptionType!: pulumi.Output<string | undefined>;
    /**
     * The database engine used by the instance. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/api-gpdb-2016-05-03-createdbinstance) `EngineVersion`.
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * The version of the database engine used by the instance.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Field `instanceChargeType` has been deprecated from provider version 1.187.0. New field `paymentType` instead.
     *
     * @deprecated Field `instanceChargeType` has been deprecated from version 1.187.0. Use `paymentType` instead.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * The number of nodes. Valid values: `2`, `4`, `8`, `12`, `16`, `24`, `32`, `64`, `96`, `128`.
     */
    public readonly instanceGroupCount!: pulumi.Output<number | undefined>;
    /**
     * The network type of the instance. Valid values: `VPC`.
     */
    public readonly instanceNetworkType!: pulumi.Output<string>;
    /**
     * The specification of segment nodes.
     * * When `dbInstanceCategory` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
     * * When `dbInstanceCategory` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
     * * When `dbInstanceCategory` is `Serverless`, Valid values: `4C16G`, `8C32G`.
     * > **NOTE:** This parameter must be passed to create a storage elastic mode instance and a serverless version instance.
     */
    public readonly instanceSpec!: pulumi.Output<string | undefined>;
    /**
     * The ip whitelist. See `ipWhitelist` below.
     * Default to creating a whitelist group with the group name "default" and securityIpList "127.0.0.1".
     */
    public readonly ipWhitelists!: pulumi.Output<outputs.gpdb.InstanceIpWhitelist[]>;
    /**
     * The end time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 03:00Z. start time should be later than end time.
     */
    public readonly maintainEndTime!: pulumi.Output<string>;
    /**
     * The start time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 02:00Z.
     */
    public readonly maintainStartTime!: pulumi.Output<string>;
    /**
     * The amount of coordinator node resources. Valid values: `2`, `4`, `8`, `16`, `32`.
     */
    public readonly masterCu!: pulumi.Output<number>;
    /**
     * The number of Master nodes. **NOTE:** Field `masterNodeNum` has been deprecated from provider version 1.213.0.
     *
     * @deprecated Field `masterNodeNum` has been deprecated from provider version 1.213.0.
     */
    public readonly masterNodeNum!: pulumi.Output<number | undefined>;
    /**
     * The billing method of the instance. Valid values: `Subscription`, `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The duration that you will buy the resource, in month. required when `paymentType` is `Subscription`. Valid values: `Year`, `Month`.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * (Available since v1.196.0) The connection port of the instance.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * The private ip address. **NOTE:** Field `privateIpAddress` has been deprecated from provider version 1.213.0.
     *
     * @deprecated Field `privateIpAddress` has been deprecated from provider version 1.213.0.
     */
    public readonly privateIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The ID of the enterprise resource group to which the instance belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Resource management mode. Valid values: `resourceGroup`, `resourceQueue`.
     */
    public readonly resourceManagementMode!: pulumi.Output<string>;
    /**
     * Field `securityIpList` has been deprecated from provider version 1.187.0. New field `ipWhitelist` instead.
     *
     * @deprecated Field 'security_ip_list' has been deprecated from version 1.187.0. Use 'ip_whitelist' instead.
     */
    public readonly securityIpLists!: pulumi.Output<string[] | undefined>;
    /**
     * Calculate the number of nodes. Valid values: `2` to `512`. The value range of the high-availability version of the storage elastic mode is `4` to `512`, and the value must be a multiple of `4`. The value range of the basic version of the storage elastic mode is `2` to `512`, and the value must be a multiple of `2`. The-Serverless version has a value range of `2` to `512`. The value must be a multiple of `2`.
     * > **NOTE:** This parameter must be passed in to create a storage elastic mode instance and a Serverless version instance. During the public beta of the Serverless version (from 0101, 2022 to 0131, 2022), a maximum of 12 compute nodes can be created.
     */
    public readonly segNodeNum!: pulumi.Output<number>;
    /**
     * The seg storage type. Valid values: `cloudEssd`, `cloudEfficiency`.
     * > **NOTE:** This parameter must be passed in to create a storage elastic mode instance. Storage Elastic Mode Basic Edition instances only support ESSD cloud disks.
     */
    public readonly segStorageType!: pulumi.Output<string | undefined>;
    /**
     * Enable or disable SSL. Valid values: `0` and `1`.
     */
    public readonly sslEnabled!: pulumi.Output<number>;
    /**
     * The status of the instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The storage capacity. Unit: GB. Valid values: `50` to `4000`.
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    public readonly storageSize!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The used time. When the parameter `period` is `Year`, the `usedTime` value is `1` to `3`. When the parameter `period` is `Month`, the `usedTime` value is `1` to `9`.
     */
    public readonly usedTime!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to enable vector engine optimization. Default value: `disabled`. Valid values: `enabled` and `disabled`.
     */
    public readonly vectorConfigurationStatus!: pulumi.Output<string>;
    /**
     * The vpc ID of the resource.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The vswitch id.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The zone ID of the instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["createSampleData"] = state ? state.createSampleData : undefined;
            resourceInputs["dbInstanceCategory"] = state ? state.dbInstanceCategory : undefined;
            resourceInputs["dbInstanceClass"] = state ? state.dbInstanceClass : undefined;
            resourceInputs["dbInstanceMode"] = state ? state.dbInstanceMode : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            resourceInputs["encryptionType"] = state ? state.encryptionType : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceGroupCount"] = state ? state.instanceGroupCount : undefined;
            resourceInputs["instanceNetworkType"] = state ? state.instanceNetworkType : undefined;
            resourceInputs["instanceSpec"] = state ? state.instanceSpec : undefined;
            resourceInputs["ipWhitelists"] = state ? state.ipWhitelists : undefined;
            resourceInputs["maintainEndTime"] = state ? state.maintainEndTime : undefined;
            resourceInputs["maintainStartTime"] = state ? state.maintainStartTime : undefined;
            resourceInputs["masterCu"] = state ? state.masterCu : undefined;
            resourceInputs["masterNodeNum"] = state ? state.masterNodeNum : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["resourceManagementMode"] = state ? state.resourceManagementMode : undefined;
            resourceInputs["securityIpLists"] = state ? state.securityIpLists : undefined;
            resourceInputs["segNodeNum"] = state ? state.segNodeNum : undefined;
            resourceInputs["segStorageType"] = state ? state.segStorageType : undefined;
            resourceInputs["sslEnabled"] = state ? state.sslEnabled : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["usedTime"] = state ? state.usedTime : undefined;
            resourceInputs["vectorConfigurationStatus"] = state ? state.vectorConfigurationStatus : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.dbInstanceMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceMode'");
            }
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["createSampleData"] = args ? args.createSampleData : undefined;
            resourceInputs["dbInstanceCategory"] = args ? args.dbInstanceCategory : undefined;
            resourceInputs["dbInstanceClass"] = args ? args.dbInstanceClass : undefined;
            resourceInputs["dbInstanceMode"] = args ? args.dbInstanceMode : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            resourceInputs["encryptionType"] = args ? args.encryptionType : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceGroupCount"] = args ? args.instanceGroupCount : undefined;
            resourceInputs["instanceNetworkType"] = args ? args.instanceNetworkType : undefined;
            resourceInputs["instanceSpec"] = args ? args.instanceSpec : undefined;
            resourceInputs["ipWhitelists"] = args ? args.ipWhitelists : undefined;
            resourceInputs["maintainEndTime"] = args ? args.maintainEndTime : undefined;
            resourceInputs["maintainStartTime"] = args ? args.maintainStartTime : undefined;
            resourceInputs["masterCu"] = args ? args.masterCu : undefined;
            resourceInputs["masterNodeNum"] = args ? args.masterNodeNum : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["resourceManagementMode"] = args ? args.resourceManagementMode : undefined;
            resourceInputs["securityIpLists"] = args ? args.securityIpLists : undefined;
            resourceInputs["segNodeNum"] = args ? args.segNodeNum : undefined;
            resourceInputs["segStorageType"] = args ? args.segStorageType : undefined;
            resourceInputs["sslEnabled"] = args ? args.sslEnabled : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["usedTime"] = args ? args.usedTime : undefined;
            resourceInputs["vectorConfigurationStatus"] = args ? args.vectorConfigurationStatus : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["connectionString"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Field `availabilityZone` has been deprecated from provider version 1.187.0. New field `zoneId` instead.
     *
     * @deprecated Field 'availability_zone' has been deprecated from version 1.187.0. Use 'zone_id' instead.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * (Available since v1.196.0) The connection string of the instance.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * Whether to load the sample dataset after the instance is created. Valid values: `true`, `false`.
     */
    createSampleData?: pulumi.Input<boolean>;
    /**
     * The db instance category. Valid values: `Basic`, `HighAvailability`.
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    dbInstanceCategory?: pulumi.Input<string>;
    /**
     * The db instance class. see [Instance specifications](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/instance-types).
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    dbInstanceClass?: pulumi.Input<string>;
    /**
     * The db instance mode. Valid values: `StorageElastic`, `Serverless`, `Classic`.
     */
    dbInstanceMode?: pulumi.Input<string>;
    /**
     * The description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the encryption key.
     * > **NOTE:** If `encryptionType` is set to `CloudDisk`, you must specify an encryption key that resides in the same region as the cloud disk that is specified by EncryptionType. Otherwise, leave this parameter empty.
     */
    encryptionKey?: pulumi.Input<string>;
    /**
     * The encryption type. Valid values: `CloudDisk`.
     * > **NOTE:** Disk encryption cannot be disabled after it is enabled.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * The database engine used by the instance. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/api-gpdb-2016-05-03-createdbinstance) `EngineVersion`.
     */
    engine?: pulumi.Input<string>;
    /**
     * The version of the database engine used by the instance.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Field `instanceChargeType` has been deprecated from provider version 1.187.0. New field `paymentType` instead.
     *
     * @deprecated Field `instanceChargeType` has been deprecated from version 1.187.0. Use `paymentType` instead.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The number of nodes. Valid values: `2`, `4`, `8`, `12`, `16`, `24`, `32`, `64`, `96`, `128`.
     */
    instanceGroupCount?: pulumi.Input<number>;
    /**
     * The network type of the instance. Valid values: `VPC`.
     */
    instanceNetworkType?: pulumi.Input<string>;
    /**
     * The specification of segment nodes.
     * * When `dbInstanceCategory` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
     * * When `dbInstanceCategory` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
     * * When `dbInstanceCategory` is `Serverless`, Valid values: `4C16G`, `8C32G`.
     * > **NOTE:** This parameter must be passed to create a storage elastic mode instance and a serverless version instance.
     */
    instanceSpec?: pulumi.Input<string>;
    /**
     * The ip whitelist. See `ipWhitelist` below.
     * Default to creating a whitelist group with the group name "default" and securityIpList "127.0.0.1".
     */
    ipWhitelists?: pulumi.Input<pulumi.Input<inputs.gpdb.InstanceIpWhitelist>[]>;
    /**
     * The end time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 03:00Z. start time should be later than end time.
     */
    maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 02:00Z.
     */
    maintainStartTime?: pulumi.Input<string>;
    /**
     * The amount of coordinator node resources. Valid values: `2`, `4`, `8`, `16`, `32`.
     */
    masterCu?: pulumi.Input<number>;
    /**
     * The number of Master nodes. **NOTE:** Field `masterNodeNum` has been deprecated from provider version 1.213.0.
     *
     * @deprecated Field `masterNodeNum` has been deprecated from provider version 1.213.0.
     */
    masterNodeNum?: pulumi.Input<number>;
    /**
     * The billing method of the instance. Valid values: `Subscription`, `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. required when `paymentType` is `Subscription`. Valid values: `Year`, `Month`.
     */
    period?: pulumi.Input<string>;
    /**
     * (Available since v1.196.0) The connection port of the instance.
     */
    port?: pulumi.Input<string>;
    /**
     * The private ip address. **NOTE:** Field `privateIpAddress` has been deprecated from provider version 1.213.0.
     *
     * @deprecated Field `privateIpAddress` has been deprecated from provider version 1.213.0.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the enterprise resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Resource management mode. Valid values: `resourceGroup`, `resourceQueue`.
     */
    resourceManagementMode?: pulumi.Input<string>;
    /**
     * Field `securityIpList` has been deprecated from provider version 1.187.0. New field `ipWhitelist` instead.
     *
     * @deprecated Field 'security_ip_list' has been deprecated from version 1.187.0. Use 'ip_whitelist' instead.
     */
    securityIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Calculate the number of nodes. Valid values: `2` to `512`. The value range of the high-availability version of the storage elastic mode is `4` to `512`, and the value must be a multiple of `4`. The value range of the basic version of the storage elastic mode is `2` to `512`, and the value must be a multiple of `2`. The-Serverless version has a value range of `2` to `512`. The value must be a multiple of `2`.
     * > **NOTE:** This parameter must be passed in to create a storage elastic mode instance and a Serverless version instance. During the public beta of the Serverless version (from 0101, 2022 to 0131, 2022), a maximum of 12 compute nodes can be created.
     */
    segNodeNum?: pulumi.Input<number>;
    /**
     * The seg storage type. Valid values: `cloudEssd`, `cloudEfficiency`.
     * > **NOTE:** This parameter must be passed in to create a storage elastic mode instance. Storage Elastic Mode Basic Edition instances only support ESSD cloud disks.
     */
    segStorageType?: pulumi.Input<string>;
    /**
     * Enable or disable SSL. Valid values: `0` and `1`.
     */
    sslEnabled?: pulumi.Input<number>;
    /**
     * The status of the instance.
     */
    status?: pulumi.Input<string>;
    /**
     * The storage capacity. Unit: GB. Valid values: `50` to `4000`.
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    storageSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The used time. When the parameter `period` is `Year`, the `usedTime` value is `1` to `3`. When the parameter `period` is `Month`, the `usedTime` value is `1` to `9`.
     */
    usedTime?: pulumi.Input<string>;
    /**
     * Specifies whether to enable vector engine optimization. Default value: `disabled`. Valid values: `enabled` and `disabled`.
     */
    vectorConfigurationStatus?: pulumi.Input<string>;
    /**
     * The vpc ID of the resource.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The vswitch id.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Field `availabilityZone` has been deprecated from provider version 1.187.0. New field `zoneId` instead.
     *
     * @deprecated Field 'availability_zone' has been deprecated from version 1.187.0. Use 'zone_id' instead.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Whether to load the sample dataset after the instance is created. Valid values: `true`, `false`.
     */
    createSampleData?: pulumi.Input<boolean>;
    /**
     * The db instance category. Valid values: `Basic`, `HighAvailability`.
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    dbInstanceCategory?: pulumi.Input<string>;
    /**
     * The db instance class. see [Instance specifications](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/instance-types).
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    dbInstanceClass?: pulumi.Input<string>;
    /**
     * The db instance mode. Valid values: `StorageElastic`, `Serverless`, `Classic`.
     */
    dbInstanceMode: pulumi.Input<string>;
    /**
     * The description of the instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the encryption key.
     * > **NOTE:** If `encryptionType` is set to `CloudDisk`, you must specify an encryption key that resides in the same region as the cloud disk that is specified by EncryptionType. Otherwise, leave this parameter empty.
     */
    encryptionKey?: pulumi.Input<string>;
    /**
     * The encryption type. Valid values: `CloudDisk`.
     * > **NOTE:** Disk encryption cannot be disabled after it is enabled.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * The database engine used by the instance. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/latest/api-gpdb-2016-05-03-createdbinstance) `EngineVersion`.
     */
    engine: pulumi.Input<string>;
    /**
     * The version of the database engine used by the instance.
     */
    engineVersion: pulumi.Input<string>;
    /**
     * Field `instanceChargeType` has been deprecated from provider version 1.187.0. New field `paymentType` instead.
     *
     * @deprecated Field `instanceChargeType` has been deprecated from version 1.187.0. Use `paymentType` instead.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The number of nodes. Valid values: `2`, `4`, `8`, `12`, `16`, `24`, `32`, `64`, `96`, `128`.
     */
    instanceGroupCount?: pulumi.Input<number>;
    /**
     * The network type of the instance. Valid values: `VPC`.
     */
    instanceNetworkType?: pulumi.Input<string>;
    /**
     * The specification of segment nodes.
     * * When `dbInstanceCategory` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
     * * When `dbInstanceCategory` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
     * * When `dbInstanceCategory` is `Serverless`, Valid values: `4C16G`, `8C32G`.
     * > **NOTE:** This parameter must be passed to create a storage elastic mode instance and a serverless version instance.
     */
    instanceSpec?: pulumi.Input<string>;
    /**
     * The ip whitelist. See `ipWhitelist` below.
     * Default to creating a whitelist group with the group name "default" and securityIpList "127.0.0.1".
     */
    ipWhitelists?: pulumi.Input<pulumi.Input<inputs.gpdb.InstanceIpWhitelist>[]>;
    /**
     * The end time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 03:00Z. start time should be later than end time.
     */
    maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 02:00Z.
     */
    maintainStartTime?: pulumi.Input<string>;
    /**
     * The amount of coordinator node resources. Valid values: `2`, `4`, `8`, `16`, `32`.
     */
    masterCu?: pulumi.Input<number>;
    /**
     * The number of Master nodes. **NOTE:** Field `masterNodeNum` has been deprecated from provider version 1.213.0.
     *
     * @deprecated Field `masterNodeNum` has been deprecated from provider version 1.213.0.
     */
    masterNodeNum?: pulumi.Input<number>;
    /**
     * The billing method of the instance. Valid values: `Subscription`, `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. required when `paymentType` is `Subscription`. Valid values: `Year`, `Month`.
     */
    period?: pulumi.Input<string>;
    /**
     * The private ip address. **NOTE:** Field `privateIpAddress` has been deprecated from provider version 1.213.0.
     *
     * @deprecated Field `privateIpAddress` has been deprecated from provider version 1.213.0.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the enterprise resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Resource management mode. Valid values: `resourceGroup`, `resourceQueue`.
     */
    resourceManagementMode?: pulumi.Input<string>;
    /**
     * Field `securityIpList` has been deprecated from provider version 1.187.0. New field `ipWhitelist` instead.
     *
     * @deprecated Field 'security_ip_list' has been deprecated from version 1.187.0. Use 'ip_whitelist' instead.
     */
    securityIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Calculate the number of nodes. Valid values: `2` to `512`. The value range of the high-availability version of the storage elastic mode is `4` to `512`, and the value must be a multiple of `4`. The value range of the basic version of the storage elastic mode is `2` to `512`, and the value must be a multiple of `2`. The-Serverless version has a value range of `2` to `512`. The value must be a multiple of `2`.
     * > **NOTE:** This parameter must be passed in to create a storage elastic mode instance and a Serverless version instance. During the public beta of the Serverless version (from 0101, 2022 to 0131, 2022), a maximum of 12 compute nodes can be created.
     */
    segNodeNum?: pulumi.Input<number>;
    /**
     * The seg storage type. Valid values: `cloudEssd`, `cloudEfficiency`.
     * > **NOTE:** This parameter must be passed in to create a storage elastic mode instance. Storage Elastic Mode Basic Edition instances only support ESSD cloud disks.
     */
    segStorageType?: pulumi.Input<string>;
    /**
     * Enable or disable SSL. Valid values: `0` and `1`.
     */
    sslEnabled?: pulumi.Input<number>;
    /**
     * The storage capacity. Unit: GB. Valid values: `50` to `4000`.
     * > **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
     */
    storageSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The used time. When the parameter `period` is `Year`, the `usedTime` value is `1` to `3`. When the parameter `period` is `Month`, the `usedTime` value is `1` to `9`.
     */
    usedTime?: pulumi.Input<string>;
    /**
     * Specifies whether to enable vector engine optimization. Default value: `disabled`. Valid values: `enabled` and `disabled`.
     */
    vectorConfigurationStatus?: pulumi.Input<string>;
    /**
     * The vpc ID of the resource.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The vswitch id.
     */
    vswitchId: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}
