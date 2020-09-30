// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN Route Service resource. The virtual border routers (VBRs) and Cloud Connect Network (CCN) instances attached to Cloud Enterprise Network (CEN) instances can access the cloud services deployed in VPCs through the CEN instances.
 *
 * For information about CEN Route Service and how to use it, see [What is Route Service](https://www.alibabacloud.com/help/en/doc-detail/106671.htm).
 *
 * > **NOTE:** Available in v1.99.0+.
 *
 * > **NOTE:** Ensure that at least one VPC in the selected region is attached to the CEN instance.
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
 * const name = config.requireObject("name");
 * const exampleNetworks = alicloud.vpc.getNetworks({
 *     isExample: true,
 * });
 * const exampleInstance = new alicloud.cen.Instance("exampleInstance", {});
 * const vpc = new alicloud.cen.InstanceAttachment("vpc", {
 *     instanceId: exampleInstance.id,
 *     childInstanceId: exampleNetworks.then(exampleNetworks => exampleNetworks.vpcs[0].id),
 *     childInstanceType: "VPC",
 *     childInstanceRegionId: exampleNetworks.then(exampleNetworks => exampleNetworks.vpcs[0].regionId),
 * });
 * const _this = new alicloud.cen.RouteService("this", {
 *     accessRegionId: exampleNetworks.then(exampleNetworks => exampleNetworks.vpcs[0].regionId),
 *     hostRegionId: exampleNetworks.then(exampleNetworks => exampleNetworks.vpcs[0].regionId),
 *     hostVpcId: exampleNetworks.then(exampleNetworks => exampleNetworks.vpcs[0].id),
 *     cenId: vpc.instanceId,
 *     host: "100.118.28.52/32",
 * });
 * ```
 */
export class RouteService extends pulumi.CustomResource {
    /**
     * Get an existing RouteService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteServiceState, opts?: pulumi.CustomResourceOptions): RouteService {
        return new RouteService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/routeService:RouteService';

    /**
     * Returns true if the given object is an instance of RouteService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteService.__pulumiType;
    }

    /**
     * The region of the network instances that access the cloud services.
     */
    public readonly accessRegionId!: pulumi.Output<string>;
    /**
     * The ID of the CEN instance.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The description of the cloud service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain name or IP address of the cloud service.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The region of the cloud service.
     */
    public readonly hostRegionId!: pulumi.Output<string>;
    /**
     * The VPC associated with the cloud service.
     */
    public readonly hostVpcId!: pulumi.Output<string>;
    /**
     * The status of the cloud service.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a RouteService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteServiceArgs | RouteServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteServiceState | undefined;
            inputs["accessRegionId"] = state ? state.accessRegionId : undefined;
            inputs["cenId"] = state ? state.cenId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["hostRegionId"] = state ? state.hostRegionId : undefined;
            inputs["hostVpcId"] = state ? state.hostVpcId : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RouteServiceArgs | undefined;
            if (!args || args.accessRegionId === undefined) {
                throw new Error("Missing required property 'accessRegionId'");
            }
            if (!args || args.cenId === undefined) {
                throw new Error("Missing required property 'cenId'");
            }
            if (!args || args.host === undefined) {
                throw new Error("Missing required property 'host'");
            }
            if (!args || args.hostRegionId === undefined) {
                throw new Error("Missing required property 'hostRegionId'");
            }
            if (!args || args.hostVpcId === undefined) {
                throw new Error("Missing required property 'hostVpcId'");
            }
            inputs["accessRegionId"] = args ? args.accessRegionId : undefined;
            inputs["cenId"] = args ? args.cenId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["hostRegionId"] = args ? args.hostRegionId : undefined;
            inputs["hostVpcId"] = args ? args.hostVpcId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RouteService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteService resources.
 */
export interface RouteServiceState {
    /**
     * The region of the network instances that access the cloud services.
     */
    readonly accessRegionId?: pulumi.Input<string>;
    /**
     * The ID of the CEN instance.
     */
    readonly cenId?: pulumi.Input<string>;
    /**
     * The description of the cloud service.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain name or IP address of the cloud service.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * The region of the cloud service.
     */
    readonly hostRegionId?: pulumi.Input<string>;
    /**
     * The VPC associated with the cloud service.
     */
    readonly hostVpcId?: pulumi.Input<string>;
    /**
     * The status of the cloud service.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteService resource.
 */
export interface RouteServiceArgs {
    /**
     * The region of the network instances that access the cloud services.
     */
    readonly accessRegionId: pulumi.Input<string>;
    /**
     * The ID of the CEN instance.
     */
    readonly cenId: pulumi.Input<string>;
    /**
     * The description of the cloud service.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain name or IP address of the cloud service.
     */
    readonly host: pulumi.Input<string>;
    /**
     * The region of the cloud service.
     */
    readonly hostRegionId: pulumi.Input<string>;
    /**
     * The VPC associated with the cloud service.
     */
    readonly hostVpcId: pulumi.Input<string>;
}
