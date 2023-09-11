// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Serverless App Engine (SAE) GreyTagRoute resource.
 *
 * For information about Serverless App Engine (SAE) GreyTagRoute and how to use it, see [What is GreyTagRoute](https://www.alibabacloud.com/help/en/sae/latest/create-grey-tag-route).
 *
 * > **NOTE:** Available since v1.160.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultRandomInteger = new random.RandomInteger("defaultRandomInteger", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultNamespace = new alicloud.sae.Namespace("defaultNamespace", {
 *     namespaceId: pulumi.all([defaultRegions, defaultRandomInteger.result]).apply(([defaultRegions, result]) => `${defaultRegions.regions?.[0]?.id}:example${result}`),
 *     namespaceName: name,
 *     namespaceDescription: name,
 *     enableMicroRegistration: false,
 * });
 * const defaultApplication = new alicloud.sae.Application("defaultApplication", {
 *     appDescription: name,
 *     appName: name,
 *     namespaceId: defaultNamespace.id,
 *     imageUrl: defaultRegions.then(defaultRegions => `registry-vpc.${defaultRegions.regions?.[0]?.id}.aliyuncs.com/sae-demo-image/consumer:1.0`),
 *     packageType: "Image",
 *     securityGroupId: defaultSecurityGroup.id,
 *     vpcId: defaultNetwork.id,
 *     vswitchId: defaultSwitch.id,
 *     timezone: "Asia/Beijing",
 *     replicas: 5,
 *     cpu: 500,
 *     memory: 2048,
 * });
 * const defaultGreyTagRoute = new alicloud.sae.GreyTagRoute("defaultGreyTagRoute", {
 *     greyTagRouteName: name,
 *     description: name,
 *     appId: defaultApplication.id,
 *     scRules: [{
 *         items: [{
 *             type: "param",
 *             name: "tfexample",
 *             operator: "rawvalue",
 *             value: "example",
 *             cond: "==",
 *         }],
 *         path: "/tf/example",
 *         condition: "AND",
 *     }],
 *     dubboRules: [{
 *         items: [{
 *             cond: "==",
 *             expr: ".key1",
 *             index: 1,
 *             operator: "rawvalue",
 *             value: "value1",
 *         }],
 *         condition: "OR",
 *         group: "DUBBO",
 *         methodName: "example",
 *         serviceName: "com.example.service",
 *         version: "1.0.0",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Serverless App Engine (SAE) GreyTagRoute can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:sae/greyTagRoute:GreyTagRoute example <id>
 * ```
 */
export class GreyTagRoute extends pulumi.CustomResource {
    /**
     * Get an existing GreyTagRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GreyTagRouteState, opts?: pulumi.CustomResourceOptions): GreyTagRoute {
        return new GreyTagRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:sae/greyTagRoute:GreyTagRoute';

    /**
     * Returns true if the given object is an instance of GreyTagRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GreyTagRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GreyTagRoute.__pulumiType;
    }

    /**
     * The ID  of the SAE Application.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The description of GreyTagRoute.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The grayscale rule created for Dubbo Application. See `dubboRules` below.
     */
    public readonly dubboRules!: pulumi.Output<outputs.sae.GreyTagRouteDubboRule[] | undefined>;
    /**
     * The name of GreyTagRoute.
     */
    public readonly greyTagRouteName!: pulumi.Output<string>;
    /**
     * The grayscale rule created for SpringCloud Application. See `scRules` below.
     */
    public readonly scRules!: pulumi.Output<outputs.sae.GreyTagRouteScRule[] | undefined>;

    /**
     * Create a GreyTagRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GreyTagRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GreyTagRouteArgs | GreyTagRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GreyTagRouteState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dubboRules"] = state ? state.dubboRules : undefined;
            resourceInputs["greyTagRouteName"] = state ? state.greyTagRouteName : undefined;
            resourceInputs["scRules"] = state ? state.scRules : undefined;
        } else {
            const args = argsOrState as GreyTagRouteArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.greyTagRouteName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'greyTagRouteName'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dubboRules"] = args ? args.dubboRules : undefined;
            resourceInputs["greyTagRouteName"] = args ? args.greyTagRouteName : undefined;
            resourceInputs["scRules"] = args ? args.scRules : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GreyTagRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GreyTagRoute resources.
 */
export interface GreyTagRouteState {
    /**
     * The ID  of the SAE Application.
     */
    appId?: pulumi.Input<string>;
    /**
     * The description of GreyTagRoute.
     */
    description?: pulumi.Input<string>;
    /**
     * The grayscale rule created for Dubbo Application. See `dubboRules` below.
     */
    dubboRules?: pulumi.Input<pulumi.Input<inputs.sae.GreyTagRouteDubboRule>[]>;
    /**
     * The name of GreyTagRoute.
     */
    greyTagRouteName?: pulumi.Input<string>;
    /**
     * The grayscale rule created for SpringCloud Application. See `scRules` below.
     */
    scRules?: pulumi.Input<pulumi.Input<inputs.sae.GreyTagRouteScRule>[]>;
}

/**
 * The set of arguments for constructing a GreyTagRoute resource.
 */
export interface GreyTagRouteArgs {
    /**
     * The ID  of the SAE Application.
     */
    appId: pulumi.Input<string>;
    /**
     * The description of GreyTagRoute.
     */
    description?: pulumi.Input<string>;
    /**
     * The grayscale rule created for Dubbo Application. See `dubboRules` below.
     */
    dubboRules?: pulumi.Input<pulumi.Input<inputs.sae.GreyTagRouteDubboRule>[]>;
    /**
     * The name of GreyTagRoute.
     */
    greyTagRouteName: pulumi.Input<string>;
    /**
     * The grayscale rule created for SpringCloud Application. See `scRules` below.
     */
    scRules?: pulumi.Input<pulumi.Input<inputs.sae.GreyTagRouteScRule>[]>;
}
