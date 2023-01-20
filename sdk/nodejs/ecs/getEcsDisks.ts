// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecs Disks of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ecs.getEcsDisks({
 *     ids: ["d-artgdsvdvxxxx"],
 *     nameRegex: "tf-test",
 * });
 * export const firstEcsDiskId = example.then(example => example.disks?.[0]?.id);
 * ```
 */
export function getEcsDisks(args?: GetEcsDisksArgs, opts?: pulumi.InvokeOptions): Promise<GetEcsDisksResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecs/getEcsDisks:getEcsDisks", {
        "additionalAttributes": args.additionalAttributes,
        "autoSnapshotPolicyId": args.autoSnapshotPolicyId,
        "availabilityZone": args.availabilityZone,
        "category": args.category,
        "deleteAutoSnapshot": args.deleteAutoSnapshot,
        "deleteWithInstance": args.deleteWithInstance,
        "diskName": args.diskName,
        "diskType": args.diskType,
        "dryRun": args.dryRun,
        "enableAutoSnapshot": args.enableAutoSnapshot,
        "enableAutomatedSnapshotPolicy": args.enableAutomatedSnapshotPolicy,
        "enableShared": args.enableShared,
        "encrypted": args.encrypted,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "kmsKeyId": args.kmsKeyId,
        "nameRegex": args.nameRegex,
        "operationLocks": args.operationLocks,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "paymentType": args.paymentType,
        "portable": args.portable,
        "resourceGroupId": args.resourceGroupId,
        "snapshotId": args.snapshotId,
        "status": args.status,
        "tags": args.tags,
        "type": args.type,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcsDisks.
 */
export interface GetEcsDisksArgs {
    /**
     * Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
     */
    additionalAttributes?: string[];
    /**
     * Query cloud disks based on the automatic snapshot policy ID.
     */
    autoSnapshotPolicyId?: string;
    /**
     * Availability zone of the disk.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    availabilityZone?: string;
    /**
     * Disk category.
     */
    category?: string;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released.
     */
    deleteAutoSnapshot?: boolean;
    /**
     * Indicates whether the disk is released together with the instance.
     */
    deleteWithInstance?: boolean;
    /**
     * The disk name.
     */
    diskName?: string;
    /**
     * The disk type.
     */
    diskType?: string;
    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     */
    dryRun?: boolean;
    /**
     * Whether the disk implements an automatic snapshot policy.
     */
    enableAutoSnapshot?: boolean;
    /**
     * Whether the disk implements an automatic snapshot policy.
     */
    enableAutomatedSnapshotPolicy?: boolean;
    /**
     * Whether it is shared block storage.
     */
    enableShared?: boolean;
    /**
     * Indicate whether the disk is encrypted or not.
     */
    encrypted?: string;
    /**
     * A list of Disk IDs.
     */
    ids?: string[];
    /**
     * The instance ID of the disk mount.
     */
    instanceId?: string;
    /**
     * The kms key id.
     */
    kmsKeyId?: string;
    /**
     * A regex string to filter results by Disk name.
     */
    nameRegex?: string;
    operationLocks?: inputs.ecs.GetEcsDisksOperationLock[];
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * Payment method for disk.
     */
    paymentType?: string;
    /**
     * Whether the disk is unmountable.
     */
    portable?: boolean;
    /**
     * The Id of resource group.
     */
    resourceGroupId?: string;
    /**
     * Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
     */
    snapshotId?: string;
    /**
     * Current status.
     */
    status?: string;
    /**
     * A map of tags assigned to the disk.
     */
    tags?: {[key: string]: any};
    /**
     * Disk type.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
     */
    type?: string;
    /**
     * The zone id.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getEcsDisks.
 */
export interface GetEcsDisksResult {
    readonly additionalAttributes?: string[];
    readonly autoSnapshotPolicyId?: string;
    /**
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    readonly availabilityZone?: string;
    readonly category?: string;
    readonly deleteAutoSnapshot?: boolean;
    readonly deleteWithInstance?: boolean;
    readonly diskName?: string;
    readonly diskType?: string;
    readonly disks: outputs.ecs.GetEcsDisksDisk[];
    readonly dryRun?: boolean;
    readonly enableAutoSnapshot?: boolean;
    readonly enableAutomatedSnapshotPolicy?: boolean;
    readonly enableShared?: boolean;
    readonly encrypted?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId?: string;
    readonly kmsKeyId?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly operationLocks?: outputs.ecs.GetEcsDisksOperationLock[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    readonly paymentType?: string;
    readonly portable?: boolean;
    readonly resourceGroupId?: string;
    readonly snapshotId?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: any};
    readonly totalCount: number;
    /**
     * @deprecated Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
     */
    readonly type?: string;
    readonly zoneId?: string;
}
/**
 * This data source provides the Ecs Disks of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ecs.getEcsDisks({
 *     ids: ["d-artgdsvdvxxxx"],
 *     nameRegex: "tf-test",
 * });
 * export const firstEcsDiskId = example.then(example => example.disks?.[0]?.id);
 * ```
 */
export function getEcsDisksOutput(args?: GetEcsDisksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEcsDisksResult> {
    return pulumi.output(args).apply((a: any) => getEcsDisks(a, opts))
}

/**
 * A collection of arguments for invoking getEcsDisks.
 */
export interface GetEcsDisksOutputArgs {
    /**
     * Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
     */
    additionalAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Query cloud disks based on the automatic snapshot policy ID.
     */
    autoSnapshotPolicyId?: pulumi.Input<string>;
    /**
     * Availability zone of the disk.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Disk category.
     */
    category?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released.
     */
    deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The disk name.
     */
    diskName?: pulumi.Input<string>;
    /**
     * The disk type.
     */
    diskType?: pulumi.Input<string>;
    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Whether the disk implements an automatic snapshot policy.
     */
    enableAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Whether the disk implements an automatic snapshot policy.
     */
    enableAutomatedSnapshotPolicy?: pulumi.Input<boolean>;
    /**
     * Whether it is shared block storage.
     */
    enableShared?: pulumi.Input<boolean>;
    /**
     * Indicate whether the disk is encrypted or not.
     */
    encrypted?: pulumi.Input<string>;
    /**
     * A list of Disk IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The instance ID of the disk mount.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The kms key id.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Disk name.
     */
    nameRegex?: pulumi.Input<string>;
    operationLocks?: pulumi.Input<pulumi.Input<inputs.ecs.GetEcsDisksOperationLockArgs>[]>;
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * Payment method for disk.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Whether the disk is unmountable.
     */
    portable?: pulumi.Input<boolean>;
    /**
     * The Id of resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Current status.
     */
    status?: pulumi.Input<string>;
    /**
     * A map of tags assigned to the disk.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Disk type.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
     */
    type?: pulumi.Input<string>;
    /**
     * The zone id.
     */
    zoneId?: pulumi.Input<string>;
}
