// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Nas File System resource.
 *
 * After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.
 *
 * For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)
 *
 * > **NOTE:** Available in v1.33.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const foo = new alicloud.nas.FileSystem("foo", {
 *     description: "tf-testAccNasConfig",
 *     protocolType: "NFS",
 *     storageType: "Performance",
 * });
 * ```
 */
export class FileSystem extends pulumi.CustomResource {
    /**
     * Get an existing FileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileSystemState, opts?: pulumi.CustomResourceOptions): FileSystem {
        return new FileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nas/fileSystem:FileSystem';

    /**
     * Returns true if the given object is an instance of FileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileSystem.__pulumiType;
    }

    /**
     * The File System description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
     */
    public readonly protocolType!: pulumi.Output<string>;
    /**
     * The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
     */
    public readonly storageType!: pulumi.Output<string>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileSystemArgs | FileSystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FileSystemState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["protocolType"] = state ? state.protocolType : undefined;
            inputs["storageType"] = state ? state.storageType : undefined;
        } else {
            const args = argsOrState as FileSystemArgs | undefined;
            if (!args || args.protocolType === undefined) {
                throw new Error("Missing required property 'protocolType'");
            }
            if (!args || args.storageType === undefined) {
                throw new Error("Missing required property 'storageType'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["protocolType"] = args ? args.protocolType : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FileSystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FileSystem resources.
 */
export interface FileSystemState {
    /**
     * The File System description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
     */
    readonly protocolType?: pulumi.Input<string>;
    /**
     * The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
     */
    readonly storageType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    /**
     * The File System description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
     */
    readonly protocolType: pulumi.Input<string>;
    /**
     * The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
     */
    readonly storageType: pulumi.Input<string>;
}
