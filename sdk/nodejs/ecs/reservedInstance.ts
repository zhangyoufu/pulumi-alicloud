// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Reserved Instance resource.
 *
 * > **NOTE:** Available in 1.65.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.ecs.getInstanceTypes({
 *     instanceTypeFamily: "ecs.g6",
 * });
 * const defaultReservedInstance = new alicloud.ecs.ReservedInstance("default", {
 *     instanceType: _default.then(_default => _default.instanceTypes?.[0]?.id),
 *     instanceAmount: 1,
 *     periodUnit: "Month",
 *     offeringType: "All Upfront",
 *     reservedInstanceName: "terraform-example",
 *     description: "ReservedInstance",
 *     zoneId: _default.then(_default => _default.instanceTypes?.[0]?.availabilityZones?.[0]),
 *     scope: "Zone",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * reservedInstance can be imported using id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/reservedInstance:ReservedInstance default ecsri-uf6df4xm0h3licit****
 * ```
 */
export class ReservedInstance extends pulumi.CustomResource {
    /**
     * Get an existing ReservedInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReservedInstanceState, opts?: pulumi.CustomResourceOptions): ReservedInstance {
        return new ReservedInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/reservedInstance:ReservedInstance';

    /**
     * Returns true if the given object is an instance of ReservedInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReservedInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReservedInstance.__pulumiType;
    }

    /**
     * Indicates the sharing status of the reserved instance when the AllocationType parameter is set to Shared. Valid values: `allocated`: The reserved instance is allocated to another account. `beAllocated`: The reserved instance is allocated by another account.
     */
    public /*out*/ readonly allocationStatus!: pulumi.Output<string>;
    /**
     * The auto-renewal term of the reserved instance. This parameter takes effect only when AutoRenew is set to true. Valid values: 1, 12, 36, and 60. Default value when `periodUnit` is set to Month: 1 Default value when `periodUnit` is set to Year: 12
     */
    public readonly autoRenewPeriod!: pulumi.Output<number>;
    /**
     * The time when the reserved instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the RI. 2 to 256 English or Chinese characters. It cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The time when the reserved instance expires.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).
     */
    public readonly instanceAmount!: pulumi.Output<number>;
    /**
     * Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Field `name` has been deprecated from provider version 1.194.0. New field `reservedInstanceName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.194.0. New field 'reserved_instance_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Payment type of the RI. Default value: `All Upfront`. Valid values:
     * - `No Upfront`: No upfront payment.
     * - `Partial Upfront`: A portion of upfront payment.
     * - `All Upfront`: Full upfront payment.
     */
    public readonly offeringType!: pulumi.Output<string>;
    /**
     * Details about the lock status of the reserved instance.
     */
    public /*out*/ readonly operationLocks!: pulumi.Output<outputs.ecs.ReservedInstanceOperationLock[]>;
    /**
     * The validity period of the reserved instance. Default value: `1`. **NOTE:** From version 1.183.0, `period` can be set to `5`, when `periodUnit` is `Year`.
     * - When `periodUnit` is `Year`, Valid values: `1`, `3`, `5`.
     * - When `periodUnit` is `Month`, Valid values: `1`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The unit of the validity period of the reserved instance. Valid value: `Month`, `Year`. Default value: `Year`. **NOTE:** From version 1.183.0, `periodUnit` can be set to `Month`.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The operating system type of the image used by the instance. Optional values: `Windows`, `Linux`. Default is `Linux`.
     */
    public readonly platform!: pulumi.Output<string>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`,`Normal`.
     */
    public readonly renewalStatus!: pulumi.Output<string>;
    /**
     * Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.
     */
    public readonly reservedInstanceName!: pulumi.Output<string>;
    /**
     * Resource group ID.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Scope of the RI. Optional values: `Region`: region-level, `Zone`: zone-level. Default is `Region`.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * The time when the reserved instance took effect.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * The status of the reserved instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).
     */
    public readonly zoneId!: pulumi.Output<string | undefined>;

    /**
     * Create a ReservedInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReservedInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReservedInstanceArgs | ReservedInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReservedInstanceState | undefined;
            resourceInputs["allocationStatus"] = state ? state.allocationStatus : undefined;
            resourceInputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["instanceAmount"] = state ? state.instanceAmount : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["offeringType"] = state ? state.offeringType : undefined;
            resourceInputs["operationLocks"] = state ? state.operationLocks : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            resourceInputs["reservedInstanceName"] = state ? state.reservedInstanceName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ReservedInstanceArgs | undefined;
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceAmount"] = args ? args.instanceAmount : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["offeringType"] = args ? args.offeringType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["platform"] = args ? args.platform : undefined;
            resourceInputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            resourceInputs["reservedInstanceName"] = args ? args.reservedInstanceName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["allocationStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["operationLocks"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReservedInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReservedInstance resources.
 */
export interface ReservedInstanceState {
    /**
     * Indicates the sharing status of the reserved instance when the AllocationType parameter is set to Shared. Valid values: `allocated`: The reserved instance is allocated to another account. `beAllocated`: The reserved instance is allocated by another account.
     */
    allocationStatus?: pulumi.Input<string>;
    /**
     * The auto-renewal term of the reserved instance. This parameter takes effect only when AutoRenew is set to true. Valid values: 1, 12, 36, and 60. Default value when `periodUnit` is set to Month: 1 Default value when `periodUnit` is set to Year: 12
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The time when the reserved instance was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the RI. 2 to 256 English or Chinese characters. It cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The time when the reserved instance expires.
     */
    expiredTime?: pulumi.Input<string>;
    /**
     * Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).
     */
    instanceAmount?: pulumi.Input<number>;
    /**
     * Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.194.0. New field `reservedInstanceName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.194.0. New field 'reserved_instance_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Payment type of the RI. Default value: `All Upfront`. Valid values:
     * - `No Upfront`: No upfront payment.
     * - `Partial Upfront`: A portion of upfront payment.
     * - `All Upfront`: Full upfront payment.
     */
    offeringType?: pulumi.Input<string>;
    /**
     * Details about the lock status of the reserved instance.
     */
    operationLocks?: pulumi.Input<pulumi.Input<inputs.ecs.ReservedInstanceOperationLock>[]>;
    /**
     * The validity period of the reserved instance. Default value: `1`. **NOTE:** From version 1.183.0, `period` can be set to `5`, when `periodUnit` is `Year`.
     * - When `periodUnit` is `Year`, Valid values: `1`, `3`, `5`.
     * - When `periodUnit` is `Month`, Valid values: `1`.
     */
    period?: pulumi.Input<number>;
    /**
     * The unit of the validity period of the reserved instance. Valid value: `Month`, `Year`. Default value: `Year`. **NOTE:** From version 1.183.0, `periodUnit` can be set to `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The operating system type of the image used by the instance. Optional values: `Windows`, `Linux`. Default is `Linux`.
     */
    platform?: pulumi.Input<string>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`,`Normal`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.
     */
    reservedInstanceName?: pulumi.Input<string>;
    /**
     * Resource group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Scope of the RI. Optional values: `Region`: region-level, `Zone`: zone-level. Default is `Region`.
     */
    scope?: pulumi.Input<string>;
    /**
     * The time when the reserved instance took effect.
     */
    startTime?: pulumi.Input<string>;
    /**
     * The status of the reserved instance.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReservedInstance resource.
 */
export interface ReservedInstanceArgs {
    /**
     * The auto-renewal term of the reserved instance. This parameter takes effect only when AutoRenew is set to true. Valid values: 1, 12, 36, and 60. Default value when `periodUnit` is set to Month: 1 Default value when `periodUnit` is set to Year: 12
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * Description of the RI. 2 to 256 English or Chinese characters. It cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).
     */
    instanceAmount?: pulumi.Input<number>;
    /**
     * Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).
     */
    instanceType: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.194.0. New field `reservedInstanceName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.194.0. New field 'reserved_instance_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Payment type of the RI. Default value: `All Upfront`. Valid values:
     * - `No Upfront`: No upfront payment.
     * - `Partial Upfront`: A portion of upfront payment.
     * - `All Upfront`: Full upfront payment.
     */
    offeringType?: pulumi.Input<string>;
    /**
     * The validity period of the reserved instance. Default value: `1`. **NOTE:** From version 1.183.0, `period` can be set to `5`, when `periodUnit` is `Year`.
     * - When `periodUnit` is `Year`, Valid values: `1`, `3`, `5`.
     * - When `periodUnit` is `Month`, Valid values: `1`.
     */
    period?: pulumi.Input<number>;
    /**
     * The unit of the validity period of the reserved instance. Valid value: `Month`, `Year`. Default value: `Year`. **NOTE:** From version 1.183.0, `periodUnit` can be set to `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The operating system type of the image used by the instance. Optional values: `Windows`, `Linux`. Default is `Linux`.
     */
    platform?: pulumi.Input<string>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`,`Normal`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.
     */
    reservedInstanceName?: pulumi.Input<string>;
    /**
     * Resource group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Scope of the RI. Optional values: `Region`: region-level, `Zone`: zone-level. Default is `Region`.
     */
    scope?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).
     */
    zoneId?: pulumi.Input<string>;
}
