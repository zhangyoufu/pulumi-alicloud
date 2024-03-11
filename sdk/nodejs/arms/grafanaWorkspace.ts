// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ARMS Grafana Workspace resource.
 *
 * For information about ARMS Grafana Workspace and how to use it, see [What is Grafana Workspace](https://next.api.alibabacloud.com/document/ARMS/2019-08-08/ListGrafanaWorkspace).
 *
 * > **NOTE:** Available since v1.215.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({});
 * const defaultGrafanaWorkspace = new alicloud.arms.GrafanaWorkspace("defaultGrafanaWorkspace", {
 *     grafanaVersion: "9.0.x",
 *     description: name,
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     grafanaWorkspaceEdition: "standard",
 *     grafanaWorkspaceName: name,
 *     tags: {
 *         Created: "tf",
 *         For: "example",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ARMS Grafana Workspace can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:arms/grafanaWorkspace:GrafanaWorkspace example <id>
 * ```
 */
export class GrafanaWorkspace extends pulumi.CustomResource {
    /**
     * Get an existing GrafanaWorkspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrafanaWorkspaceState, opts?: pulumi.CustomResourceOptions): GrafanaWorkspace {
        return new GrafanaWorkspace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:arms/grafanaWorkspace:GrafanaWorkspace';

    /**
     * Returns true if the given object is an instance of GrafanaWorkspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GrafanaWorkspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GrafanaWorkspace.__pulumiType;
    }

    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The version of the grafana.
     */
    public readonly grafanaVersion!: pulumi.Output<string | undefined>;
    /**
     * The edition of the grafana.
     */
    public readonly grafanaWorkspaceEdition!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly grafanaWorkspaceName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tag of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a GrafanaWorkspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GrafanaWorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrafanaWorkspaceArgs | GrafanaWorkspaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrafanaWorkspaceState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["grafanaVersion"] = state ? state.grafanaVersion : undefined;
            resourceInputs["grafanaWorkspaceEdition"] = state ? state.grafanaWorkspaceEdition : undefined;
            resourceInputs["grafanaWorkspaceName"] = state ? state.grafanaWorkspaceName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as GrafanaWorkspaceArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["grafanaVersion"] = args ? args.grafanaVersion : undefined;
            resourceInputs["grafanaWorkspaceEdition"] = args ? args.grafanaWorkspaceEdition : undefined;
            resourceInputs["grafanaWorkspaceName"] = args ? args.grafanaWorkspaceName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GrafanaWorkspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GrafanaWorkspace resources.
 */
export interface GrafanaWorkspaceState {
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * The version of the grafana.
     */
    grafanaVersion?: pulumi.Input<string>;
    /**
     * The edition of the grafana.
     */
    grafanaWorkspaceEdition?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    grafanaWorkspaceName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a GrafanaWorkspace resource.
 */
export interface GrafanaWorkspaceArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * The version of the grafana.
     */
    grafanaVersion?: pulumi.Input<string>;
    /**
     * The edition of the grafana.
     */
    grafanaWorkspaceEdition?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    grafanaWorkspaceName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
