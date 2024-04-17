// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an EDAS deploy group resource, see [What is EDAS Deploy Group](https://www.alibabacloud.com/help/en/edas/developer-reference/api-edas-2017-08-01-insertdeploygroup).
 *
 * > **NOTE:** Available since v1.82.0.
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
 * const name = config.get("name") || "tf-example";
 * const defaultInteger = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const default = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: `${name}-${defaultInteger.result}`,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultCluster = new alicloud.edas.Cluster("default", {
 *     clusterName: `${name}-${defaultInteger.result}`,
 *     clusterType: 2,
 *     networkMode: 2,
 *     logicalRegionId: _default.then(_default => _default.regions?.[0]?.id),
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultApplication = new alicloud.edas.Application("default", {
 *     applicationName: `${name}-${defaultInteger.result}`,
 *     clusterId: defaultCluster.id,
 *     packageType: "JAR",
 * });
 * const defaultDeployGroup = new alicloud.edas.DeployGroup("default", {
 *     appId: defaultApplication.id,
 *     groupName: `${name}-${defaultInteger.result}`,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * EDAS deploy group can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:edas/deployGroup:DeployGroup group app_id:group_name:group_id
 * ```
 */
export class DeployGroup extends pulumi.CustomResource {
    /**
     * Get an existing DeployGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeployGroupState, opts?: pulumi.CustomResourceOptions): DeployGroup {
        return new DeployGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:edas/deployGroup:DeployGroup';

    /**
     * Returns true if the given object is an instance of DeployGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeployGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeployGroup.__pulumiType;
    }

    /**
     * The ID of the application that you want to deploy.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The name of the instance group that you want to create.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
     */
    public /*out*/ readonly groupType!: pulumi.Output<number>;

    /**
     * Create a DeployGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeployGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeployGroupArgs | DeployGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeployGroupState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["groupType"] = state ? state.groupType : undefined;
        } else {
            const args = argsOrState as DeployGroupArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["groupType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeployGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployGroup resources.
 */
export interface DeployGroupState {
    /**
     * The ID of the application that you want to deploy.
     */
    appId?: pulumi.Input<string>;
    /**
     * The name of the instance group that you want to create.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
     */
    groupType?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DeployGroup resource.
 */
export interface DeployGroupArgs {
    /**
     * The ID of the application that you want to deploy.
     */
    appId: pulumi.Input<string>;
    /**
     * The name of the instance group that you want to create.
     */
    groupName: pulumi.Input<string>;
}
