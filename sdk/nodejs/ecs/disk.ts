// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Cloud disk can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/disk:Disk example d-abc12345678
 * ```
 */
export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskState, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/disk:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    public readonly advancedFeatures!: pulumi.Output<string | undefined>;
    /**
     * The Zone to create the disk in.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    public readonly category!: pulumi.Output<string | undefined>;
    public readonly dedicatedBlockStorageClusterId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     */
    public readonly deleteAutoSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean | undefined>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly diskName!: pulumi.Output<string>;
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     */
    public readonly enableAutoSnapshot!: pulumi.Output<boolean | undefined>;
    public readonly encryptAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:                                                       
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     */
    public readonly performanceLevel!: pulumi.Output<string>;
    /**
     * The Id of resource group which the disk belongs.
     * > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    public readonly size!: pulumi.Output<number | undefined>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The disk status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly storageSetId!: pulumi.Output<string | undefined>;
    public readonly storageSetPartitionNumber!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskArgs | DiskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiskState | undefined;
            inputs["advancedFeatures"] = state ? state.advancedFeatures : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["category"] = state ? state.category : undefined;
            inputs["dedicatedBlockStorageClusterId"] = state ? state.dedicatedBlockStorageClusterId : undefined;
            inputs["deleteAutoSnapshot"] = state ? state.deleteAutoSnapshot : undefined;
            inputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskName"] = state ? state.diskName : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["enableAutoSnapshot"] = state ? state.enableAutoSnapshot : undefined;
            inputs["encryptAlgorithm"] = state ? state.encryptAlgorithm : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["paymentType"] = state ? state.paymentType : undefined;
            inputs["performanceLevel"] = state ? state.performanceLevel : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["storageSetId"] = state ? state.storageSetId : undefined;
            inputs["storageSetPartitionNumber"] = state ? state.storageSetPartitionNumber : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as DiskArgs | undefined;
            inputs["advancedFeatures"] = args ? args.advancedFeatures : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["dedicatedBlockStorageClusterId"] = args ? args.dedicatedBlockStorageClusterId : undefined;
            inputs["deleteAutoSnapshot"] = args ? args.deleteAutoSnapshot : undefined;
            inputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diskName"] = args ? args.diskName : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["enableAutoSnapshot"] = args ? args.enableAutoSnapshot : undefined;
            inputs["encryptAlgorithm"] = args ? args.encryptAlgorithm : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["paymentType"] = args ? args.paymentType : undefined;
            inputs["performanceLevel"] = args ? args.performanceLevel : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["storageSetId"] = args ? args.storageSetId : undefined;
            inputs["storageSetPartitionNumber"] = args ? args.storageSetPartitionNumber : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Disk.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Disk resources.
 */
export interface DiskState {
    readonly advancedFeatures?: pulumi.Input<string>;
    /**
     * The Zone to create the disk in.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    readonly category?: pulumi.Input<string>;
    readonly dedicatedBlockStorageClusterId?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     */
    readonly deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     */
    readonly deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    readonly diskName?: pulumi.Input<string>;
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     */
    readonly enableAutoSnapshot?: pulumi.Input<boolean>;
    readonly encryptAlgorithm?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
     */
    readonly name?: pulumi.Input<string>;
    readonly paymentType?: pulumi.Input<string>;
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:                                                       
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     */
    readonly performanceLevel?: pulumi.Input<string>;
    /**
     * The Id of resource group which the disk belongs.
     * > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The disk status.
     */
    readonly status?: pulumi.Input<string>;
    readonly storageSetId?: pulumi.Input<string>;
    readonly storageSetPartitionNumber?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    readonly advancedFeatures?: pulumi.Input<string>;
    /**
     * The Zone to create the disk in.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    readonly category?: pulumi.Input<string>;
    readonly dedicatedBlockStorageClusterId?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     */
    readonly deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     */
    readonly deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    readonly diskName?: pulumi.Input<string>;
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     */
    readonly enableAutoSnapshot?: pulumi.Input<boolean>;
    readonly encryptAlgorithm?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
     */
    readonly name?: pulumi.Input<string>;
    readonly paymentType?: pulumi.Input<string>;
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:                                                       
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     */
    readonly performanceLevel?: pulumi.Input<string>;
    /**
     * The Id of resource group which the disk belongs.
     * > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    readonly snapshotId?: pulumi.Input<string>;
    readonly storageSetId?: pulumi.Input<string>;
    readonly storageSetPartitionNumber?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
    readonly zoneId?: pulumi.Input<string>;
}
