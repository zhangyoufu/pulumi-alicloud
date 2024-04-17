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
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.nas.getZones({
 *     fileSystemType: "standard",
 * });
 * const foo = new alicloud.nas.FileSystem("foo", {
 *     protocolType: "NFS",
 *     storageType: "Performance",
 *     description: "terraform-example",
 *     encryptType: 1,
 *     zoneId: example.then(example => example.zones?.[0]?.zoneId),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.nas.getZones({
 *     fileSystemType: "extreme",
 * });
 * const foo = new alicloud.nas.FileSystem("foo", {
 *     fileSystemType: "extreme",
 *     protocolType: "NFS",
 *     zoneId: example.then(example => example.zones?.[0]?.zoneId),
 *     storageType: "standard",
 *     capacity: 100,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.nas.getZones({
 *     fileSystemType: "cpfs",
 * });
 * const exampleNetwork = new alicloud.vpc.Network("example", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("example", {
 *     vswitchName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: exampleNetwork.id,
 *     zoneId: example.then(example => example.zones?.[1]?.zoneId),
 * });
 * const exampleFileSystem = new alicloud.nas.FileSystem("example", {
 *     protocolType: "cpfs",
 *     storageType: "advance_200",
 *     fileSystemType: "cpfs",
 *     capacity: 3600,
 *     zoneId: example.then(example => example.zones?.[1]?.zoneId),
 *     vpcId: exampleNetwork.id,
 *     vswitchId: exampleSwitch.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Nas File System can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:nas/fileSystem:FileSystem foo 1337849c59
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
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * The capacity of the file system. The `capacity` is required when the `fileSystemType` is `extreme`.
     * Unit: gib; **Note**: The minimum value is 100.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * The File System description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
     * * Valid values:
     */
    public readonly encryptType!: pulumi.Output<number | undefined>;
    /**
     * the type of the file system. 
     * Valid values:
     * `standard` (Default),
     * `extreme`,
     * `cpfs`.
     */
    public readonly fileSystemType!: pulumi.Output<string | undefined>;
    /**
     * The id of the KMS key. The `kmsKeyId` is required when the `encryptType` is `2`.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `fileSystemType` is `standard`),
     * `cpfs` (Available when the `fileSystemType` is `cpfs`).
     */
    public readonly protocolType!: pulumi.Output<string>;
    /**
     * The storage type of the file System. 
     * * Valid values:
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The id of the VPC. The `vpcId` is required when the `fileSystemType` is `cpfs`.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * The id of the vSwitch. The `vswitchId` is required when the `fileSystemType` is `cpfs`.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocolType` and `storageType` configuration.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileSystemArgs | FileSystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FileSystemState | undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encryptType"] = state ? state.encryptType : undefined;
            resourceInputs["fileSystemType"] = state ? state.fileSystemType : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["protocolType"] = state ? state.protocolType : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as FileSystemArgs | undefined;
            if ((!args || args.protocolType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocolType'");
            }
            if ((!args || args.storageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageType'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encryptType"] = args ? args.encryptType : undefined;
            resourceInputs["fileSystemType"] = args ? args.fileSystemType : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["protocolType"] = args ? args.protocolType : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FileSystem resources.
 */
export interface FileSystemState {
    /**
     * The capacity of the file system. The `capacity` is required when the `fileSystemType` is `extreme`.
     * Unit: gib; **Note**: The minimum value is 100.
     */
    capacity?: pulumi.Input<number>;
    /**
     * The File System description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
     * * Valid values:
     */
    encryptType?: pulumi.Input<number>;
    /**
     * the type of the file system. 
     * Valid values:
     * `standard` (Default),
     * `extreme`,
     * `cpfs`.
     */
    fileSystemType?: pulumi.Input<string>;
    /**
     * The id of the KMS key. The `kmsKeyId` is required when the `encryptType` is `2`.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `fileSystemType` is `standard`),
     * `cpfs` (Available when the `fileSystemType` is `cpfs`).
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The storage type of the file System. 
     * * Valid values:
     */
    storageType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The id of the VPC. The `vpcId` is required when the `fileSystemType` is `cpfs`.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The id of the vSwitch. The `vswitchId` is required when the `fileSystemType` is `cpfs`.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocolType` and `storageType` configuration.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    /**
     * The capacity of the file system. The `capacity` is required when the `fileSystemType` is `extreme`.
     * Unit: gib; **Note**: The minimum value is 100.
     */
    capacity?: pulumi.Input<number>;
    /**
     * The File System description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
     * * Valid values:
     */
    encryptType?: pulumi.Input<number>;
    /**
     * the type of the file system. 
     * Valid values:
     * `standard` (Default),
     * `extreme`,
     * `cpfs`.
     */
    fileSystemType?: pulumi.Input<string>;
    /**
     * The id of the KMS key. The `kmsKeyId` is required when the `encryptType` is `2`.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `fileSystemType` is `standard`),
     * `cpfs` (Available when the `fileSystemType` is `cpfs`).
     */
    protocolType: pulumi.Input<string>;
    /**
     * The storage type of the file System. 
     * * Valid values:
     */
    storageType: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The id of the VPC. The `vpcId` is required when the `fileSystemType` is `cpfs`.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The id of the vSwitch. The `vswitchId` is required when the `fileSystemType` is `cpfs`.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocolType` and `storageType` configuration.
     */
    zoneId?: pulumi.Input<string>;
}
