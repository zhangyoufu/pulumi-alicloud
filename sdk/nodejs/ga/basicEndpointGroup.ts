// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Basic Endpoint Group resource.
 *
 * For information about Global Accelerator (GA) Basic Endpoint Group and how to use it, see [What is Basic Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/createbasicendpointgroup).
 *
 * > **NOTE:** Available in v1.194.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "default-NODELETING",
 * });
 * const defaultSwitches = defaultNetworks.then(defaultNetworks => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?.[0],
 * }));
 * const defaultApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", {
 *     loadBalancerSpec: "slb.s2.small",
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?.[0]),
 * });
 * const defaultBasicAccelerator = new alicloud.ga.BasicAccelerator("defaultBasicAccelerator", {
 *     duration: 1,
 *     pricingCycle: "Month",
 *     bandwidthBillingType: "CDT",
 *     autoPay: true,
 *     autoUseCoupon: "true",
 *     autoRenew: false,
 *     autoRenewDuration: 1,
 * });
 * const defaultBasicEndpointGroup = new alicloud.ga.BasicEndpointGroup("defaultBasicEndpointGroup", {
 *     acceleratorId: defaultBasicAccelerator.id,
 *     endpointGroupRegion: "cn-beijing",
 *     endpointType: "SLB",
 *     endpointAddress: defaultApplicationLoadBalancer.id,
 *     endpointSubAddress: "192.168.0.1",
 *     basicEndpointGroupName: "example_value",
 *     description: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator (GA) Basic Endpoint Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/basicEndpointGroup:BasicEndpointGroup example <id>
 * ```
 */
export class BasicEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing BasicEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasicEndpointGroupState, opts?: pulumi.CustomResourceOptions): BasicEndpointGroup {
        return new BasicEndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/basicEndpointGroup:BasicEndpointGroup';

    /**
     * Returns true if the given object is an instance of BasicEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasicEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasicEndpointGroup.__pulumiType;
    }

    /**
     * The ID of the basic GA instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
     */
    public readonly basicEndpointGroupName!: pulumi.Output<string | undefined>;
    /**
     * The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The address of the endpoint.
     */
    public readonly endpointAddress!: pulumi.Output<string>;
    /**
     * The ID of the region where you want to create the endpoint group.
     */
    public readonly endpointGroupRegion!: pulumi.Output<string>;
    /**
     * The sub address of the endpoint.
     */
    public readonly endpointSubAddress!: pulumi.Output<string>;
    /**
     * The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
     */
    public readonly endpointType!: pulumi.Output<string>;
    /**
     * The status of the Basic Endpoint Group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a BasicEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasicEndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasicEndpointGroupArgs | BasicEndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BasicEndpointGroupState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["basicEndpointGroupName"] = state ? state.basicEndpointGroupName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpointAddress"] = state ? state.endpointAddress : undefined;
            resourceInputs["endpointGroupRegion"] = state ? state.endpointGroupRegion : undefined;
            resourceInputs["endpointSubAddress"] = state ? state.endpointSubAddress : undefined;
            resourceInputs["endpointType"] = state ? state.endpointType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BasicEndpointGroupArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.endpointGroupRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointGroupRegion'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["basicEndpointGroupName"] = args ? args.basicEndpointGroupName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointAddress"] = args ? args.endpointAddress : undefined;
            resourceInputs["endpointGroupRegion"] = args ? args.endpointGroupRegion : undefined;
            resourceInputs["endpointSubAddress"] = args ? args.endpointSubAddress : undefined;
            resourceInputs["endpointType"] = args ? args.endpointType : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BasicEndpointGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasicEndpointGroup resources.
 */
export interface BasicEndpointGroupState {
    /**
     * The ID of the basic GA instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
     */
    basicEndpointGroupName?: pulumi.Input<string>;
    /**
     * The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
     */
    description?: pulumi.Input<string>;
    /**
     * The address of the endpoint.
     */
    endpointAddress?: pulumi.Input<string>;
    /**
     * The ID of the region where you want to create the endpoint group.
     */
    endpointGroupRegion?: pulumi.Input<string>;
    /**
     * The sub address of the endpoint.
     */
    endpointSubAddress?: pulumi.Input<string>;
    /**
     * The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
     */
    endpointType?: pulumi.Input<string>;
    /**
     * The status of the Basic Endpoint Group.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BasicEndpointGroup resource.
 */
export interface BasicEndpointGroupArgs {
    /**
     * The ID of the basic GA instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
     */
    basicEndpointGroupName?: pulumi.Input<string>;
    /**
     * The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
     */
    description?: pulumi.Input<string>;
    /**
     * The address of the endpoint.
     */
    endpointAddress?: pulumi.Input<string>;
    /**
     * The ID of the region where you want to create the endpoint group.
     */
    endpointGroupRegion: pulumi.Input<string>;
    /**
     * The sub address of the endpoint.
     */
    endpointSubAddress?: pulumi.Input<string>;
    /**
     * The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
     */
    endpointType?: pulumi.Input<string>;
}
