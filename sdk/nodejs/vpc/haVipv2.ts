// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Vpc Ha Vip resource. Highly available virtual IP
 *
 * For information about Vpc Ha Vip and how to use it, see [What is Ha Vip](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createhavip).
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
 * const name = config.get("name") || "tf-testacc-example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultVpc = new alicloud.vpc.Network("defaultVpc", {
 *     description: "tf-test-acc-vpc",
 *     vpcName: name,
 *     cidrBlock: "192.168.0.0/16",
 * });
 * const defaultVswitch = new alicloud.vpc.Switch("defaultVswitch", {
 *     vpcId: defaultVpc.id,
 *     cidrBlock: "192.168.0.0/21",
 *     vswitchName: `${name}1`,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     description: "tf-testacc-vswitch",
 * });
 * const defaultRg = new alicloud.resourcemanager.ResourceGroup("defaultRg", {
 *     displayName: "tf-testacc-rg819",
 *     resourceGroupName: `${name}2`,
 * });
 * const changeRg = new alicloud.resourcemanager.ResourceGroup("changeRg", {
 *     displayName: "tf-testacc-changerg670",
 *     resourceGroupName: `${name}3`,
 * });
 * const defaultHaVipv2 = new alicloud.vpc.HaVipv2("defaultHaVipv2", {
 *     description: "test",
 *     vswitchId: defaultVswitch.id,
 *     haVipName: name,
 *     ipAddress: "192.168.1.101",
 *     resourceGroupId: defaultRg.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Vpc Ha Vip can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/haVipv2:HaVipv2 example <id>
 * ```
 */
export class HaVipv2 extends pulumi.CustomResource {
    /**
     * Get an existing HaVipv2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HaVipv2State, opts?: pulumi.CustomResourceOptions): HaVipv2 {
        return new HaVipv2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/haVipv2:HaVipv2';

    /**
     * Returns true if the given object is an instance of HaVipv2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HaVipv2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HaVipv2.__pulumiType;
    }

    /**
     * EIP bound to HaVip.
     */
    public /*out*/ readonly associatedEipAddresses!: pulumi.Output<string[]>;
    /**
     * The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
     */
    public /*out*/ readonly associatedInstanceType!: pulumi.Output<string>;
    /**
     * An ECS instance that is bound to HaVip.
     */
    public /*out*/ readonly associatedInstances!: pulumi.Output<string[]>;
    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the HaVip instance. The length is 2 to 256 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource.
     */
    public /*out*/ readonly haVipId!: pulumi.Output<string>;
    /**
     * The name of the HaVip instance.
     */
    public readonly haVipName!: pulumi.Output<string>;
    /**
     * Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
     *
     * @deprecated Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
     */
    public readonly havipName!: pulumi.Output<string>;
    /**
     * The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The primary instance ID bound to HaVip.
     */
    public /*out*/ readonly masterInstanceId!: pulumi.Output<string>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The status of this resource instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags of HaVip.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VPC ID to which the HaVip instance belongs.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The switch ID to which the HaVip instance belongs.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a HaVipv2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HaVipv2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HaVipv2Args | HaVipv2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HaVipv2State | undefined;
            resourceInputs["associatedEipAddresses"] = state ? state.associatedEipAddresses : undefined;
            resourceInputs["associatedInstanceType"] = state ? state.associatedInstanceType : undefined;
            resourceInputs["associatedInstances"] = state ? state.associatedInstances : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["haVipId"] = state ? state.haVipId : undefined;
            resourceInputs["haVipName"] = state ? state.haVipName : undefined;
            resourceInputs["havipName"] = state ? state.havipName : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["masterInstanceId"] = state ? state.masterInstanceId : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as HaVipv2Args | undefined;
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["haVipName"] = args ? args.haVipName : undefined;
            resourceInputs["havipName"] = args ? args.havipName : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["associatedEipAddresses"] = undefined /*out*/;
            resourceInputs["associatedInstanceType"] = undefined /*out*/;
            resourceInputs["associatedInstances"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["haVipId"] = undefined /*out*/;
            resourceInputs["masterInstanceId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HaVipv2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HaVipv2 resources.
 */
export interface HaVipv2State {
    /**
     * EIP bound to HaVip.
     */
    associatedEipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
     */
    associatedInstanceType?: pulumi.Input<string>;
    /**
     * An ECS instance that is bound to HaVip.
     */
    associatedInstances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the HaVip instance. The length is 2 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the resource.
     */
    haVipId?: pulumi.Input<string>;
    /**
     * The name of the HaVip instance.
     */
    haVipName?: pulumi.Input<string>;
    /**
     * Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
     *
     * @deprecated Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
     */
    havipName?: pulumi.Input<string>;
    /**
     * The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The primary instance ID bound to HaVip.
     */
    masterInstanceId?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of this resource instance.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags of HaVip.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID to which the HaVip instance belongs.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The switch ID to which the HaVip instance belongs.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HaVipv2 resource.
 */
export interface HaVipv2Args {
    /**
     * The description of the HaVip instance. The length is 2 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the HaVip instance.
     */
    haVipName?: pulumi.Input<string>;
    /**
     * Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
     *
     * @deprecated Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
     */
    havipName?: pulumi.Input<string>;
    /**
     * The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The tags of HaVip.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The switch ID to which the HaVip instance belongs.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    vswitchId: pulumi.Input<string>;
}
