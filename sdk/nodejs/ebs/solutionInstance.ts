// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a EBS Solution Instance resource.
 *
 * For information about EBS Solution Instance and how to use it, see [What is Solution Instance](https://www.alibabacloud.com/help/en/).
 *
 * > **NOTE:** Available since v1.216.0.
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
 * const zoneId = config.get("zoneId") || "cn-shanghai-l";
 * const regionId = config.get("regionId") || "cn-shanghai";
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({});
 * const defaultSolutionInstance = new alicloud.ebs.SolutionInstance("defaultSolutionInstance", {
 *     solutionInstanceName: name,
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     description: "description",
 *     solutionId: "mysql",
 *     parameters: [
 *         {
 *             parameterKey: "zoneId",
 *             parameterValue: zoneId,
 *         },
 *         {
 *             parameterKey: "ecsType",
 *             parameterValue: "ecs.c6.large",
 *         },
 *         {
 *             parameterKey: "ecsImageId",
 *             parameterValue: "CentOS_7",
 *         },
 *         {
 *             parameterKey: "internetMaxBandwidthOut",
 *             parameterValue: "100",
 *         },
 *         {
 *             parameterKey: "internetChargeType",
 *             parameterValue: "PayByTraffic",
 *         },
 *         {
 *             parameterKey: "ecsPassword",
 *             parameterValue: "Ebs12345",
 *         },
 *         {
 *             parameterKey: "sysDiskType",
 *             parameterValue: "cloud_essd",
 *         },
 *         {
 *             parameterKey: "sysDiskPerformance",
 *             parameterValue: "PL0",
 *         },
 *         {
 *             parameterKey: "sysDiskSize",
 *             parameterValue: "40",
 *         },
 *         {
 *             parameterKey: "dataDiskType",
 *             parameterValue: "cloud_essd",
 *         },
 *         {
 *             parameterKey: "dataDiskPerformance",
 *             parameterValue: "PL0",
 *         },
 *         {
 *             parameterKey: "dataDiskSize",
 *             parameterValue: "40",
 *         },
 *         {
 *             parameterKey: "mysqlVersion",
 *             parameterValue: "MySQL80",
 *         },
 *         {
 *             parameterKey: "mysqlUser",
 *             parameterValue: "root",
 *         },
 *         {
 *             parameterKey: "mysqlPassword",
 *             parameterValue: "Ebs12345",
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * EBS Solution Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ebs/solutionInstance:SolutionInstance example <id>
 * ```
 */
export class SolutionInstance extends pulumi.CustomResource {
    /**
     * Get an existing SolutionInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SolutionInstanceState, opts?: pulumi.CustomResourceOptions): SolutionInstance {
        return new SolutionInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ebs/solutionInstance:SolutionInstance';

    /**
     * Returns true if the given object is an instance of SolutionInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SolutionInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SolutionInstance.__pulumiType;
    }

    /**
     * Solution Instance Creation Time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Solution Instance Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Solution Instance Creation Parameters. See `parameters` below.
     */
    public readonly parameters!: pulumi.Output<outputs.ebs.SolutionInstanceParameter[] | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Solution ID.
     */
    public readonly solutionId!: pulumi.Output<string>;
    /**
     * Solution Instance Name.
     */
    public readonly solutionInstanceName!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a SolutionInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SolutionInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SolutionInstanceArgs | SolutionInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SolutionInstanceState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["solutionId"] = state ? state.solutionId : undefined;
            resourceInputs["solutionInstanceName"] = state ? state.solutionInstanceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SolutionInstanceArgs | undefined;
            if ((!args || args.solutionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'solutionId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["solutionId"] = args ? args.solutionId : undefined;
            resourceInputs["solutionInstanceName"] = args ? args.solutionInstanceName : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SolutionInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SolutionInstance resources.
 */
export interface SolutionInstanceState {
    /**
     * Solution Instance Creation Time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Solution Instance Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Solution Instance Creation Parameters. See `parameters` below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.ebs.SolutionInstanceParameter>[]>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Solution ID.
     */
    solutionId?: pulumi.Input<string>;
    /**
     * Solution Instance Name.
     */
    solutionInstanceName?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SolutionInstance resource.
 */
export interface SolutionInstanceArgs {
    /**
     * Solution Instance Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Solution Instance Creation Parameters. See `parameters` below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.ebs.SolutionInstanceParameter>[]>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Solution ID.
     */
    solutionId: pulumi.Input<string>;
    /**
     * Solution Instance Name.
     */
    solutionInstanceName?: pulumi.Input<string>;
}
