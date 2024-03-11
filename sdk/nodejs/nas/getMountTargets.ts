// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides MountTargets available to the user.
 *
 * > **NOTE**: Available in 1.35.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.nas.getMountTargets({
 *     fileSystemId: "1a2sc4d",
 *     accessGroupName: "tf-testAccNasConfig",
 * });
 * export const theFirstMountTargetDomain = example.then(example => example.targets?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMountTargets(args: GetMountTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:nas/getMountTargets:getMountTargets", {
        "accessGroupName": args.accessGroupName,
        "fileSystemId": args.fileSystemId,
        "ids": args.ids,
        "mountTargetDomain": args.mountTargetDomain,
        "networkType": args.networkType,
        "outputFile": args.outputFile,
        "status": args.status,
        "type": args.type,
        "vpcId": args.vpcId,
        "vswitchId": args.vswitchId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMountTargets.
 */
export interface GetMountTargetsArgs {
    /**
     * Filter results by a specific AccessGroupName.
     */
    accessGroupName?: string;
    /**
     * The ID of the FileSystem that owns the MountTarget.
     */
    fileSystemId: string;
    /**
     * A list of MountTargetDomain.
     */
    ids?: string[];
    /**
     * Field `mountTargetDomain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
     *
     * @deprecated Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
     */
    mountTargetDomain?: string;
    /**
     * Filter results by a specific NetworkType.
     */
    networkType?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
     */
    status?: string;
    /**
     * Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
     */
    type?: string;
    /**
     * Filter results by a specific VpcId.
     */
    vpcId?: string;
    /**
     * Filter results by a specific VSwitchId.
     */
    vswitchId?: string;
}

/**
 * A collection of values returned by getMountTargets.
 */
export interface GetMountTargetsResult {
    /**
     * AccessGroup of The MountTarget.
     */
    readonly accessGroupName?: string;
    readonly fileSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of MountTargetDomain.
     */
    readonly ids: string[];
    /**
     * MountTargetDomain of the MountTarget.
     *
     * @deprecated Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
     */
    readonly mountTargetDomain?: string;
    /**
     * (Available 1.95.0+) NetworkType of The MountTarget.
     */
    readonly networkType?: string;
    readonly outputFile?: string;
    /**
     * (Available 1.95.0+) The status of the mount target.
     */
    readonly status?: string;
    /**
     * A list of MountTargetDomains. Each element contains the following attributes:
     */
    readonly targets: outputs.nas.GetMountTargetsTarget[];
    /**
     * Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
     */
    readonly type?: string;
    /**
     * VpcId of The MountTarget.
     */
    readonly vpcId?: string;
    /**
     * VSwitchId of The MountTarget.
     */
    readonly vswitchId?: string;
}
/**
 * This data source provides MountTargets available to the user.
 *
 * > **NOTE**: Available in 1.35.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.nas.getMountTargets({
 *     fileSystemId: "1a2sc4d",
 *     accessGroupName: "tf-testAccNasConfig",
 * });
 * export const theFirstMountTargetDomain = example.then(example => example.targets?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMountTargetsOutput(args: GetMountTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMountTargetsResult> {
    return pulumi.output(args).apply((a: any) => getMountTargets(a, opts))
}

/**
 * A collection of arguments for invoking getMountTargets.
 */
export interface GetMountTargetsOutputArgs {
    /**
     * Filter results by a specific AccessGroupName.
     */
    accessGroupName?: pulumi.Input<string>;
    /**
     * The ID of the FileSystem that owns the MountTarget.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * A list of MountTargetDomain.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Field `mountTargetDomain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
     *
     * @deprecated Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
     */
    mountTargetDomain?: pulumi.Input<string>;
    /**
     * Filter results by a specific NetworkType.
     */
    networkType?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
     */
    status?: pulumi.Input<string>;
    /**
     * Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
     */
    type?: pulumi.Input<string>;
    /**
     * Filter results by a specific VpcId.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Filter results by a specific VSwitchId.
     */
    vswitchId?: pulumi.Input<string>;
}
