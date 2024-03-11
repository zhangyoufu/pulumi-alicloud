// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Monitor Group Instances resource.
 *
 * For information about Cloud Monitor Service Monitor Group Instances and how to use it, see [What is Monitor Group Instances](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createmonitorgroupinstances).
 *
 * > **NOTE:** Available since v1.115.0.
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
 * const name = config.get("name") || "tf_example";
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "192.168.0.0/16",
 * });
 * const defaultMonitorGroup = new alicloud.cms.MonitorGroup("defaultMonitorGroup", {monitorGroupName: name});
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const example = new alicloud.cms.MonitorGroupInstances("example", {
 *     groupId: defaultMonitorGroup.id,
 *     instances: [{
 *         instanceId: defaultNetwork.id,
 *         instanceName: name,
 *         regionId: defaultRegions.then(defaultRegions => defaultRegions.regions?.[0]?.id),
 *         category: "vpc",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Monitor Service Monitor Group Instances can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cms/monitorGroupInstances:MonitorGroupInstances example <group_id>
 * ```
 */
export class MonitorGroupInstances extends pulumi.CustomResource {
    /**
     * Get an existing MonitorGroupInstances resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorGroupInstancesState, opts?: pulumi.CustomResourceOptions): MonitorGroupInstances {
        return new MonitorGroupInstances(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/monitorGroupInstances:MonitorGroupInstances';

    /**
     * Returns true if the given object is an instance of MonitorGroupInstances.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitorGroupInstances {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitorGroupInstances.__pulumiType;
    }

    /**
     * The id of Cms Group.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * Instance information added to the Cms Group. See `instances` below.
     */
    public readonly instances!: pulumi.Output<outputs.cms.MonitorGroupInstancesInstance[]>;

    /**
     * Create a MonitorGroupInstances resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorGroupInstancesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitorGroupInstancesArgs | MonitorGroupInstancesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitorGroupInstancesState | undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
        } else {
            const args = argsOrState as MonitorGroupInstancesArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.instances === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instances'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["instances"] = args ? args.instances : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MonitorGroupInstances.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MonitorGroupInstances resources.
 */
export interface MonitorGroupInstancesState {
    /**
     * The id of Cms Group.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Instance information added to the Cms Group. See `instances` below.
     */
    instances?: pulumi.Input<pulumi.Input<inputs.cms.MonitorGroupInstancesInstance>[]>;
}

/**
 * The set of arguments for constructing a MonitorGroupInstances resource.
 */
export interface MonitorGroupInstancesArgs {
    /**
     * The id of Cms Group.
     */
    groupId: pulumi.Input<string>;
    /**
     * Instance information added to the Cms Group. See `instances` below.
     */
    instances: pulumi.Input<pulumi.Input<inputs.cms.MonitorGroupInstancesInstance>[]>;
}
