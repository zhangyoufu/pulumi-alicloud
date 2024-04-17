// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Data Works Folder resource.
 *
 * For information about Data Works Folder and how to use it, see [What is Folder](https://help.aliyun.com/document_detail/173940.html).
 *
 * > **NOTE:** Available in v1.131.0+.
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
 * const example = new alicloud.dataworks.Folder("example", {
 *     projectId: "320687",
 *     folderPath: "Business Flow/tfTestAcc/folderDi/tftest1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Data Works Folder can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dataworks/folder:Folder example <folder_id>:<$.ProjectId>
 * ```
 */
export class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderState, opts?: pulumi.CustomResourceOptions): Folder {
        return new Folder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dataworks/folder:Folder';

    /**
     * Returns true if the given object is an instance of Folder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Folder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Folder.__pulumiType;
    }

    public /*out*/ readonly folderId!: pulumi.Output<string>;
    /**
     * Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
     */
    public readonly folderPath!: pulumi.Output<string>;
    /**
     * The ID of the project.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    public readonly projectIdentifier!: pulumi.Output<string | undefined>;

    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FolderArgs | FolderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FolderState | undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["folderPath"] = state ? state.folderPath : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["projectIdentifier"] = state ? state.projectIdentifier : undefined;
        } else {
            const args = argsOrState as FolderArgs | undefined;
            if ((!args || args.folderPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderPath'");
            }
            resourceInputs["folderPath"] = args ? args.folderPath : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["projectIdentifier"] = args ? args.projectIdentifier : undefined;
            resourceInputs["folderId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Folder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Folder resources.
 */
export interface FolderState {
    folderId?: pulumi.Input<string>;
    /**
     * Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
     */
    folderPath?: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    projectIdentifier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Folder resource.
 */
export interface FolderArgs {
    /**
     * Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
     */
    folderPath: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    projectId?: pulumi.Input<string>;
    projectIdentifier?: pulumi.Input<string>;
}
