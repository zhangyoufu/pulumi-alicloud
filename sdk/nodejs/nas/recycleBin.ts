// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Network Attached Storage (NAS) Recycle Bin resource.
 *
 * For information about Network Attached Storage (NAS) Recycle Bin and how to use it, see [What is Recycle Bin](https://www.alibabacloud.com/help/en/doc-detail/264185.html).
 *
 * > **NOTE:** Available in v1.155.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleZones = alicloud.nas.getZones({
 *     fileSystemType: "standard",
 * });
 * const exampleFileSystem = new alicloud.nas.FileSystem("exampleFileSystem", {
 *     protocolType: "NFS",
 *     storageType: "Performance",
 *     description: "terraform-example",
 *     encryptType: 1,
 *     zoneId: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.zoneId),
 * });
 * const exampleRecycleBin = new alicloud.nas.RecycleBin("exampleRecycleBin", {
 *     fileSystemId: exampleFileSystem.id,
 *     reservedDays: 3,
 * });
 * ```
 *
 * ## Import
 *
 * Network Attached Storage (NAS) Recycle Bin can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:nas/recycleBin:RecycleBin example <file_system_id>
 * ```
 */
export class RecycleBin extends pulumi.CustomResource {
    /**
     * Get an existing RecycleBin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecycleBinState, opts?: pulumi.CustomResourceOptions): RecycleBin {
        return new RecycleBin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nas/recycleBin:RecycleBin';

    /**
     * Returns true if the given object is an instance of RecycleBin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecycleBin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecycleBin.__pulumiType;
    }

    /**
     * The ID of the file system for which you want to enable the recycle bin feature.
     */
    public readonly fileSystemId!: pulumi.Output<string>;
    /**
     * The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
     */
    public readonly reservedDays!: pulumi.Output<number>;
    /**
     * The status of the recycle bin.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a RecycleBin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecycleBinArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecycleBinArgs | RecycleBinState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecycleBinState | undefined;
            resourceInputs["fileSystemId"] = state ? state.fileSystemId : undefined;
            resourceInputs["reservedDays"] = state ? state.reservedDays : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RecycleBinArgs | undefined;
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            resourceInputs["reservedDays"] = args ? args.reservedDays : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RecycleBin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecycleBin resources.
 */
export interface RecycleBinState {
    /**
     * The ID of the file system for which you want to enable the recycle bin feature.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
     */
    reservedDays?: pulumi.Input<number>;
    /**
     * The status of the recycle bin.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RecycleBin resource.
 */
export interface RecycleBinArgs {
    /**
     * The ID of the file system for which you want to enable the recycle bin feature.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
     */
    reservedDays?: pulumi.Input<number>;
}
