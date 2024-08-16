// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Max Compute Project resource.
 *
 * MaxCompute project .
 *
 * For information about Max Compute Project and how to use it, see [What is Project](https://www.alibabacloud.com/help/en/maxcompute/).
 *
 * > **NOTE:** Available since v1.77.0.
 *
 * > **NOTE:** Field `name`, `specificationType`, `orderType` has been removed from provider version 1.227.1.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf_example";
 * const _default = new alicloud.maxcompute.Project("default", {
 *     defaultQuota: "默认后付费Quota",
 *     projectName: name,
 *     comment: name,
 *     productType: "PayAsYouGo",
 * });
 * ```
 *
 * ## Import
 *
 * Max Compute Project can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:maxcompute/project:Project example <id>
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:maxcompute/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Project description information. The length is 1 to 256 English or Chinese characters. The default value is blank.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Represents the creation time of the project
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Used to implement computing resource allocation. If the calculation Quota is not specified, the default Quota resource will be consumed by jobs initiated by the project. For more information about computing resource usage, see [Computing Resource Usage](https://www.alibabacloud.com/help/en/maxcompute/user-guide/use-of-computing-resources).
     */
    public readonly defaultQuota!: pulumi.Output<string | undefined>;
    /**
     * IP whitelist See `ipWhiteList` below.
     */
    public readonly ipWhiteList!: pulumi.Output<outputs.maxcompute.ProjectIpWhiteList | undefined>;
    /**
     * Logical deletion, value: (true/false) true: In this case, the project status will be changed to 'DELETING' and completely deleted after 14 days. false: immediately deleted, that is, completely deleted, permanently unrecoverable.
     */
    public readonly isLogical!: pulumi.Output<string | undefined>;
    /**
     * Project owner
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
     */
    public readonly productType!: pulumi.Output<string | undefined>;
    /**
     * The name begins with a letter, containing letters, digits, and underscores (_). It can be 3 to 28 characters in length and is globally unique.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * Project base attributes See `properties` below.
     */
    public readonly properties!: pulumi.Output<outputs.maxcompute.ProjectProperties>;
    /**
     * Security-related attributes See `securityProperties` below.
     */
    public readonly securityProperties!: pulumi.Output<outputs.maxcompute.ProjectSecurityProperties>;
    /**
     * The project status. Default value: AVAILABLE. Value: (AVAILABLE/READONLY/FROZEN/DELETING)
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The tag of the resource
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Project type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["defaultQuota"] = state ? state.defaultQuota : undefined;
            resourceInputs["ipWhiteList"] = state ? state.ipWhiteList : undefined;
            resourceInputs["isLogical"] = state ? state.isLogical : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["productType"] = state ? state.productType : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["properties"] = state ? state.properties : undefined;
            resourceInputs["securityProperties"] = state ? state.securityProperties : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["defaultQuota"] = args ? args.defaultQuota : undefined;
            resourceInputs["ipWhiteList"] = args ? args.ipWhiteList : undefined;
            resourceInputs["isLogical"] = args ? args.isLogical : undefined;
            resourceInputs["productType"] = args ? args.productType : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["securityProperties"] = args ? args.securityProperties : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Project description information. The length is 1 to 256 English or Chinese characters. The default value is blank.
     */
    comment?: pulumi.Input<string>;
    /**
     * Represents the creation time of the project
     */
    createTime?: pulumi.Input<string>;
    /**
     * Used to implement computing resource allocation. If the calculation Quota is not specified, the default Quota resource will be consumed by jobs initiated by the project. For more information about computing resource usage, see [Computing Resource Usage](https://www.alibabacloud.com/help/en/maxcompute/user-guide/use-of-computing-resources).
     */
    defaultQuota?: pulumi.Input<string>;
    /**
     * IP whitelist See `ipWhiteList` below.
     */
    ipWhiteList?: pulumi.Input<inputs.maxcompute.ProjectIpWhiteList>;
    /**
     * Logical deletion, value: (true/false) true: In this case, the project status will be changed to 'DELETING' and completely deleted after 14 days. false: immediately deleted, that is, completely deleted, permanently unrecoverable.
     */
    isLogical?: pulumi.Input<string>;
    /**
     * Project owner
     */
    owner?: pulumi.Input<string>;
    /**
     * Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
     */
    productType?: pulumi.Input<string>;
    /**
     * The name begins with a letter, containing letters, digits, and underscores (_). It can be 3 to 28 characters in length and is globally unique.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Project base attributes See `properties` below.
     */
    properties?: pulumi.Input<inputs.maxcompute.ProjectProperties>;
    /**
     * Security-related attributes See `securityProperties` below.
     */
    securityProperties?: pulumi.Input<inputs.maxcompute.ProjectSecurityProperties>;
    /**
     * The project status. Default value: AVAILABLE. Value: (AVAILABLE/READONLY/FROZEN/DELETING)
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Project type
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Project description information. The length is 1 to 256 English or Chinese characters. The default value is blank.
     */
    comment?: pulumi.Input<string>;
    /**
     * Used to implement computing resource allocation. If the calculation Quota is not specified, the default Quota resource will be consumed by jobs initiated by the project. For more information about computing resource usage, see [Computing Resource Usage](https://www.alibabacloud.com/help/en/maxcompute/user-guide/use-of-computing-resources).
     */
    defaultQuota?: pulumi.Input<string>;
    /**
     * IP whitelist See `ipWhiteList` below.
     */
    ipWhiteList?: pulumi.Input<inputs.maxcompute.ProjectIpWhiteList>;
    /**
     * Logical deletion, value: (true/false) true: In this case, the project status will be changed to 'DELETING' and completely deleted after 14 days. false: immediately deleted, that is, completely deleted, permanently unrecoverable.
     */
    isLogical?: pulumi.Input<string>;
    /**
     * Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
     */
    productType?: pulumi.Input<string>;
    /**
     * The name begins with a letter, containing letters, digits, and underscores (_). It can be 3 to 28 characters in length and is globally unique.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Project base attributes See `properties` below.
     */
    properties?: pulumi.Input<inputs.maxcompute.ProjectProperties>;
    /**
     * Security-related attributes See `securityProperties` below.
     */
    securityProperties?: pulumi.Input<inputs.maxcompute.ProjectSecurityProperties>;
    /**
     * The project status. Default value: AVAILABLE. Value: (AVAILABLE/READONLY/FROZEN/DELETING)
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
