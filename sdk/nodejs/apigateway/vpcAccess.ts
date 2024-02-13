// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleZones = alicloud.getZones({
 *     availableResourceCreation: "Instance",
 * });
 * const exampleInstanceTypes = exampleZones.then(exampleZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: exampleZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("exampleSwitch", {
 *     vswitchName: "terraform-example",
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: exampleNetwork.id,
 *     zoneId: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.id),
 * });
 * const exampleSecurityGroup = new alicloud.ecs.SecurityGroup("exampleSecurityGroup", {
 *     description: "New security group",
 *     vpcId: exampleNetwork.id,
 * });
 * const exampleImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     owners: "system",
 * });
 * const exampleInstance = new alicloud.ecs.Instance("exampleInstance", {
 *     availabilityZone: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.id),
 *     instanceName: "terraform-example",
 *     imageId: exampleImages.then(exampleImages => exampleImages.images?.[0]?.id),
 *     instanceType: exampleInstanceTypes.then(exampleInstanceTypes => exampleInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroups: [exampleSecurityGroup.id],
 *     vswitchId: exampleSwitch.id,
 * });
 * const exampleVpcAccess = new alicloud.apigateway.VpcAccess("exampleVpcAccess", {
 *     vpcId: exampleNetwork.id,
 *     instanceId: exampleInstance.id,
 *     port: 8080,
 * });
 * ```
 *
 * ## Import
 *
 * Api gateway app can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:apigateway/vpcAccess:VpcAccess example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080"
 * ```
 */
export class VpcAccess extends pulumi.CustomResource {
    /**
     * Get an existing VpcAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcAccessState, opts?: pulumi.CustomResourceOptions): VpcAccess {
        return new VpcAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:apigateway/vpcAccess:VpcAccess';

    /**
     * Returns true if the given object is an instance of VpcAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcAccess.__pulumiType;
    }

    /**
     * ID of the instance in VPC (ECS/Server Load Balance).
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of the vpc authorization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the port corresponding to the instance.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The vpc id of the vpc authorization.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcAccessArgs | VpcAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcAccessState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcAccessArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcAccess resources.
 */
export interface VpcAccessState {
    /**
     * ID of the instance in VPC (ECS/Server Load Balance).
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the vpc authorization.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the port corresponding to the instance.
     */
    port?: pulumi.Input<number>;
    /**
     * The vpc id of the vpc authorization.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcAccess resource.
 */
export interface VpcAccessArgs {
    /**
     * ID of the instance in VPC (ECS/Server Load Balance).
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of the vpc authorization.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the port corresponding to the instance.
     */
    port: pulumi.Input<number>;
    /**
     * The vpc id of the vpc authorization.
     */
    vpcId: pulumi.Input<string>;
}
