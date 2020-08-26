// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of snapshot according to the specified filters in an Alibaba Cloud account.
 *
 * For information about snapshot and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).
 *
 * > **NOTE:**  Available in 1.40.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const snapshots = pulumi.output(alicloud.ecs.getSnapshots({
 *     ids: ["s-123456890abcdef"],
 *     nameRegex: "tf-testAcc-snapshot",
 * }, { async: true }));
 * ```
 * ## Argument Reference
 *
 * The following arguments are supported:
 *
 * * `instanceId` - (Optional) The specified instance ID.
 * * `diskId` - (Optional) The specified disk ID.
 * * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values:
 *   * true: Encrypted snapshots.
 *   * false: No encryption attribute limit.
 *   
 *   Default value: false.
 * * `ids` - (Optional)  A list of snapshot IDs.
 * * `nameRegex` - (Optional) A regex string to filter results by snapshot name.
 * * `status` - (Optional) The specified snapshot status.
 *   * The snapshot status. Optional values:
 *   * progressing: The snapshots are being created.
 *   * accomplished: The snapshots are ready to use.
 *   * failed: The snapshot creation failed.
 *   * all: All status.
 *   
 *   Default value: all.
 *
 * * `type` - (Optional) The snapshot category. Optional values:
 *   * auto: Auto snapshots.
 *   * user: Manual snapshots.
 *   * all: Auto and manual snapshots.
 *   
 *   Default value: all.
 * * `sourceDiskType` - (Optional) The type of source disk:
 *   * System: The snapshots are created for system disks.
 *   * Data: The snapshots are created for data disks.
 *
 * * `usage` - (Optional) The usage of the snapshot:
 *   * image: The snapshots are used to create custom images.
 *   * disk: The snapshots are used to CreateDisk.
 *   * mage_disk: The snapshots are used to create custom images and data disks.
 *   * none: The snapshots are not used yet.
 * * `tags` - (Optional) A map of tags assigned to snapshots.
 * * `outputFile` - (Optional) The name of output file that saves the filter results.
 */
export function getSnapshots(args?: GetSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ecs/getSnapshots:getSnapshots", {
        "diskId": args.diskId,
        "encrypted": args.encrypted,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "sourceDiskType": args.sourceDiskType,
        "status": args.status,
        "tags": args.tags,
        "type": args.type,
        "usage": args.usage,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshots.
 */
export interface GetSnapshotsArgs {
    readonly diskId?: string;
    /**
     * Whether the snapshot is encrypted or not.
     */
    readonly encrypted?: boolean;
    /**
     * A list of snapshot IDs.
     */
    readonly ids?: string[];
    readonly instanceId?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Source disk attribute. Value range:
     * * System
     * * Data
     */
    readonly sourceDiskType?: string;
    /**
     * The snapshot status. Value range:
     * * progressing
     * * accomplished
     * * failed
     */
    readonly status?: string;
    /**
     * A map of tags assigned to the snapshot.
     */
    readonly tags?: {[key: string]: any};
    readonly type?: string;
    /**
     * Whether the snapshots are used to create resources or not. Value range:
     * * image
     * * disk
     * * imageDisk
     * * none
     */
    readonly usage?: string;
}

/**
 * A collection of values returned by getSnapshots.
 */
export interface GetSnapshotsResult {
    readonly diskId?: string;
    /**
     * Whether the snapshot is encrypted or not.
     */
    readonly encrypted?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of snapshot IDs.
     */
    readonly ids: string[];
    readonly instanceId?: string;
    readonly nameRegex?: string;
    /**
     * A list of snapshots names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of snapshots. Each element contains the following attributes:
     */
    readonly snapshots: outputs.ecs.GetSnapshotsSnapshot[];
    /**
     * Source disk attribute. Value range:
     * * System
     * * Data
     */
    readonly sourceDiskType?: string;
    /**
     * The snapshot status. Value range:
     * * progressing
     * * accomplished
     * * failed
     */
    readonly status?: string;
    /**
     * A map of tags assigned to the snapshot.
     */
    readonly tags?: {[key: string]: any};
    readonly type?: string;
    /**
     * Whether the snapshots are used to create resources or not. Value range:
     * * image
     * * disk
     * * imageDisk
     * * none
     */
    readonly usage?: string;
}
