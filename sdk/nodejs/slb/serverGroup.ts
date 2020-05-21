// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A virtual server group contains several ECS instances. The virtual server group can help you to define multiple listening dimension,
 * and to meet the personalized requirements of domain name and URL forwarding.
 *
 * > **NOTE:** One ECS instance can be added into multiple virtual server groups.
 *
 * > **NOTE:** One virtual server group can be attached with multiple listeners in one load balancer.
 *
 * > **NOTE:** One Classic and Internet load balancer, its virtual server group can add Classic and VPC ECS instances.
 *
 * > **NOTE:** One Classic and Intranet load balancer, its virtual server group can only add Classic ECS instances.
 *
 * > **NOTE:** One VPC load balancer, its virtual server group can only add the same VPC ECS instances.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "slbservergroupvpc";
 *
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableDiskCategory: "cloudEfficiency",
 *     availableResourceCreation: "VSwitch",
 * }, { async: true }));
 * const defaultInstanceTypes = defaultZones.apply(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }, { async: true }));
 * const defaultImages = pulumi.output(alicloud.ecs.getImages({
 *     mostRecent: true,
 *     nameRegex: "^ubuntu_18.*64",
 *     owners: "system",
 * }, { async: true }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/16",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("default", {
 *     vpcId: defaultNetwork.id,
 * });
 * const instance: alicloud.ecs.Instance[] = [];
 * for (let i = 0; i < 2; i++) {
 *     instance.push(new alicloud.ecs.Instance(`instance-${i}`, {
 *         availabilityZone: defaultZones.zones[0].id,
 *         imageId: defaultImages.images[0].id,
 *         instanceChargeType: "PostPaid",
 *         instanceName: name,
 *         instanceType: defaultInstanceTypes.instanceTypes[0].id,
 *         internetChargeType: "PayByTraffic",
 *         internetMaxBandwidthOut: 10,
 *         securityGroups: defaultSecurityGroup.id,
 *         systemDiskCategory: "cloudEfficiency",
 *         vswitchId: defaultSwitch.id,
 *     }));
 * }
 * const defaultLoadBalancer = new alicloud.slb.LoadBalancer("default", {
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultServerGroup = new alicloud.slb.ServerGroup("default", {
 *     loadBalancerId: defaultLoadBalancer.id,
 *     servers: [
 *         {
 *             port: 100,
 *             serverIds: [
 *                 instance[0].id,
 *                 instance[1].id,
 *             ],
 *             weight: 10,
 *         },
 *         {
 *             port: 80,
 *             serverIds: instance.map(v => v.id),
 *             weight: 100,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Block servers
 *
 * The servers mapping supports the following:
 *
 * * `serverIds` - (Required) A list backend server ID (ECS instance ID).
 * * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
 * * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
 * * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
 */
export class ServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerGroupState, opts?: pulumi.CustomResourceOptions): ServerGroup {
        return new ServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/serverGroup:ServerGroup';

    /**
     * Returns true if the given object is an instance of ServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerGroup.__pulumiType;
    }

    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    public readonly deleteProtectionValidation!: pulumi.Output<boolean | undefined>;
    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of ECS instances to be added. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     */
    public readonly servers!: pulumi.Output<outputs.slb.ServerGroupServer[]>;

    /**
     * Create a ServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerGroupArgs | ServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerGroupState | undefined;
            inputs["deleteProtectionValidation"] = state ? state.deleteProtectionValidation : undefined;
            inputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as ServerGroupArgs | undefined;
            if (!args || args.loadBalancerId === undefined) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            inputs["deleteProtectionValidation"] = args ? args.deleteProtectionValidation : undefined;
            inputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["servers"] = args ? args.servers : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServerGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerGroup resources.
 */
export interface ServerGroupState {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    readonly deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     */
    readonly loadBalancerId?: pulumi.Input<string>;
    /**
     * Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of ECS instances to be added. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     */
    readonly servers?: pulumi.Input<pulumi.Input<inputs.slb.ServerGroupServer>[]>;
}

/**
 * The set of arguments for constructing a ServerGroup resource.
 */
export interface ServerGroupArgs {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    readonly deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The Load Balancer ID which is used to launch a new virtual server group.
     */
    readonly loadBalancerId: pulumi.Input<string>;
    /**
     * Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of ECS instances to be added. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
     */
    readonly servers?: pulumi.Input<pulumi.Input<inputs.slb.ServerGroupServer>[]>;
}
