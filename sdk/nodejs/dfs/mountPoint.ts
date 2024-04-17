// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DFS Mount Point resource.
 *
 * For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/en/aibaba-cloud-storage-services/latest/apsara-file-storage-for-hdfs).
 *
 * > **NOTE:** Available since v1.140.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const default = alicloud.dfs.getZones({});
 * const defaultInteger = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultVPC = new alicloud.vpc.Network("DefaultVPC", {
 *     cidrBlock: "172.16.0.0/12",
 *     vpcName: name,
 * });
 * const defaultVSwitch = new alicloud.vpc.Switch("DefaultVSwitch", {
 *     description: "example",
 *     vpcId: defaultVPC.id,
 *     cidrBlock: "172.16.0.0/24",
 *     vswitchName: name,
 *     zoneId: _default.then(_default => _default.zones?.[0]?.zoneId),
 * });
 * const defaultAccessGroup = new alicloud.dfs.AccessGroup("DefaultAccessGroup", {
 *     description: "AccessGroup resource manager center example",
 *     networkType: "VPC",
 *     accessGroupName: `${name}-${defaultInteger.result}`,
 * });
 * const updateAccessGroup = new alicloud.dfs.AccessGroup("UpdateAccessGroup", {
 *     description: "Second AccessGroup resource manager center example",
 *     networkType: "VPC",
 *     accessGroupName: `${name}-update-${defaultInteger.result}`,
 * });
 * const defaultFs = new alicloud.dfs.FileSystem("DefaultFs", {
 *     spaceCapacity: 1024,
 *     description: "for mountpoint  example",
 *     storageType: "STANDARD",
 *     zoneId: _default.then(_default => _default.zones?.[0]?.zoneId),
 *     protocolType: "HDFS",
 *     dataRedundancyType: "LRS",
 *     fileSystemName: `${name}-${defaultInteger.result}`,
 * });
 * const defaultMountPoint = new alicloud.dfs.MountPoint("default", {
 *     vpcId: defaultVPC.id,
 *     description: "mountpoint example",
 *     networkType: "VPC",
 *     vswitchId: defaultVSwitch.id,
 *     fileSystemId: defaultFs.id,
 *     accessGroupId: defaultAccessGroup.id,
 *     status: "Active",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * DFS Mount Point can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dfs/mountPoint:MountPoint example <file_system_id>:<mount_point_id>
 * ```
 */
export class MountPoint extends pulumi.CustomResource {
    /**
     * Get an existing MountPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MountPointState, opts?: pulumi.CustomResourceOptions): MountPoint {
        return new MountPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dfs/mountPoint:MountPoint';

    /**
     * Returns true if the given object is an instance of MountPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MountPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MountPoint.__pulumiType;
    }

    /**
     * The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
     */
    public readonly accessGroupId!: pulumi.Output<string>;
    /**
     * The mount point alias prefix, which specifies the mount point alias prefix.
     */
    public readonly aliasPrefix!: pulumi.Output<string | undefined>;
    /**
     * The creation time of the Mount point resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the Mount point.  No more than 32 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Unique file system identifier, used to retrieve specified file system resources.
     */
    public readonly fileSystemId!: pulumi.Output<string>;
    /**
     * The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
     */
    public /*out*/ readonly mountPointId!: pulumi.Output<string>;
    /**
     * The network type of the Mount point.  Only VPC (VPC) is supported.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * VSwitch ID, which specifies the VSwitch resource used to create the mount point.
     */
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a MountPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MountPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MountPointArgs | MountPointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MountPointState | undefined;
            resourceInputs["accessGroupId"] = state ? state.accessGroupId : undefined;
            resourceInputs["aliasPrefix"] = state ? state.aliasPrefix : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fileSystemId"] = state ? state.fileSystemId : undefined;
            resourceInputs["mountPointId"] = state ? state.mountPointId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as MountPointArgs | undefined;
            if ((!args || args.accessGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessGroupId'");
            }
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            if ((!args || args.networkType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkType'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["accessGroupId"] = args ? args.accessGroupId : undefined;
            resourceInputs["aliasPrefix"] = args ? args.aliasPrefix : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["mountPointId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MountPoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MountPoint resources.
 */
export interface MountPointState {
    /**
     * The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * The mount point alias prefix, which specifies the mount point alias prefix.
     */
    aliasPrefix?: pulumi.Input<string>;
    /**
     * The creation time of the Mount point resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the Mount point.  No more than 32 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique file system identifier, used to retrieve specified file system resources.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
     */
    mountPointId?: pulumi.Input<string>;
    /**
     * The network type of the Mount point.  Only VPC (VPC) is supported.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * VSwitch ID, which specifies the VSwitch resource used to create the mount point.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MountPoint resource.
 */
export interface MountPointArgs {
    /**
     * The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
     */
    accessGroupId: pulumi.Input<string>;
    /**
     * The mount point alias prefix, which specifies the mount point alias prefix.
     */
    aliasPrefix?: pulumi.Input<string>;
    /**
     * The description of the Mount point.  No more than 32 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique file system identifier, used to retrieve specified file system resources.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * The network type of the Mount point.  Only VPC (VPC) is supported.
     */
    networkType: pulumi.Input<string>;
    /**
     * Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
     */
    vpcId: pulumi.Input<string>;
    /**
     * VSwitch ID, which specifies the VSwitch resource used to create the mount point.
     */
    vswitchId: pulumi.Input<string>;
}
