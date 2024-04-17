// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Nas File System resource.
 *
 * For information about ECD Nas File System and how to use it, see [What is Nas File System](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/api-reference-for-easy-use-1).
 *
 * > **NOTE:** Available since v1.141.0.
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
 * const _default = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultSimpleOfficeSite = new alicloud.eds.SimpleOfficeSite("default", {
 *     cidrBlock: "172.16.0.0/12",
 *     enableAdminAccess: false,
 *     desktopAccessType: "Internet",
 *     officeSiteName: `${name}-${_default.result}`,
 * });
 * const example = new alicloud.eds.NasFileSystem("example", {
 *     nasFileSystemName: name,
 *     officeSiteId: defaultSimpleOfficeSite.id,
 *     description: name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ECD Nas File System can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/nasFileSystem:NasFileSystem example <id>
 * ```
 */
export class NasFileSystem extends pulumi.CustomResource {
    /**
     * Get an existing NasFileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NasFileSystemState, opts?: pulumi.CustomResourceOptions): NasFileSystem {
        return new NasFileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/nasFileSystem:NasFileSystem';

    /**
     * Returns true if the given object is an instance of NasFileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NasFileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NasFileSystem.__pulumiType;
    }

    /**
     * The description of nas file system.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The filesystem id of nas file system.
     */
    public readonly fileSystemId!: pulumi.Output<string>;
    /**
     * The domain of mount target.
     */
    public readonly mountTargetDomain!: pulumi.Output<string>;
    /**
     * The name of nas file system.
     */
    public readonly nasFileSystemName!: pulumi.Output<string | undefined>;
    /**
     * The ID of office site.
     */
    public readonly officeSiteId!: pulumi.Output<string>;
    /**
     * The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
     */
    public readonly reset!: pulumi.Output<boolean | undefined>;
    /**
     * The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a NasFileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NasFileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NasFileSystemArgs | NasFileSystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NasFileSystemState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fileSystemId"] = state ? state.fileSystemId : undefined;
            resourceInputs["mountTargetDomain"] = state ? state.mountTargetDomain : undefined;
            resourceInputs["nasFileSystemName"] = state ? state.nasFileSystemName : undefined;
            resourceInputs["officeSiteId"] = state ? state.officeSiteId : undefined;
            resourceInputs["reset"] = state ? state.reset : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as NasFileSystemArgs | undefined;
            if ((!args || args.officeSiteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'officeSiteId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            resourceInputs["mountTargetDomain"] = args ? args.mountTargetDomain : undefined;
            resourceInputs["nasFileSystemName"] = args ? args.nasFileSystemName : undefined;
            resourceInputs["officeSiteId"] = args ? args.officeSiteId : undefined;
            resourceInputs["reset"] = args ? args.reset : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NasFileSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NasFileSystem resources.
 */
export interface NasFileSystemState {
    /**
     * The description of nas file system.
     */
    description?: pulumi.Input<string>;
    /**
     * The filesystem id of nas file system.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * The domain of mount target.
     */
    mountTargetDomain?: pulumi.Input<string>;
    /**
     * The name of nas file system.
     */
    nasFileSystemName?: pulumi.Input<string>;
    /**
     * The ID of office site.
     */
    officeSiteId?: pulumi.Input<string>;
    /**
     * The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
     */
    reset?: pulumi.Input<boolean>;
    /**
     * The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NasFileSystem resource.
 */
export interface NasFileSystemArgs {
    /**
     * The description of nas file system.
     */
    description?: pulumi.Input<string>;
    /**
     * The filesystem id of nas file system.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * The domain of mount target.
     */
    mountTargetDomain?: pulumi.Input<string>;
    /**
     * The name of nas file system.
     */
    nasFileSystemName?: pulumi.Input<string>;
    /**
     * The ID of office site.
     */
    officeSiteId: pulumi.Input<string>;
    /**
     * The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
     */
    reset?: pulumi.Input<boolean>;
}
