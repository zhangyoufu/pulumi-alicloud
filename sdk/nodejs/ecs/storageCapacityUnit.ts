// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Storage Capacity Unit resource.
 *
 * For information about ECS Storage Capacity Unit and how to use it, see [What is Storage Capacity Unit](https://www.alibabacloud.com/help/en/doc-detail/161157.html).
 *
 * > **NOTE:** Available in v1.155.0+.
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
 * const _default = new alicloud.ecs.StorageCapacityUnit("default", {
 *     capacity: 20,
 *     description: "tftestdescription",
 *     storageCapacityUnitName: "tftestname",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ECS Storage Capacity Unit can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/storageCapacityUnit:StorageCapacityUnit example <id>
 * ```
 */
export class StorageCapacityUnit extends pulumi.CustomResource {
    /**
     * Get an existing StorageCapacityUnit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageCapacityUnitState, opts?: pulumi.CustomResourceOptions): StorageCapacityUnit {
        return new StorageCapacityUnit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/storageCapacityUnit:StorageCapacityUnit';

    /**
     * Returns true if the given object is an instance of StorageCapacityUnit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageCapacityUnit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageCapacityUnit.__pulumiType;
    }

    /**
     * The capacity of the Storage Capacity Unit. Unit: GiB. Valid values: `20`, `40`, `100`, `200`, `500`, `1024`, `2048`, `5120`, `10240`, `20480`, and `51200`.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * The description of the Storage Capacity Unit. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The validity period of the Storage Capacity Unit. Default value: `1`.
     * * When PeriodUnit is set to Month, Valid values: `1`, `2`, `3`, `6`.
     * * When PeriodUnit is set to Year, Valid values: `1`, `3`, `5`.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * The unit of the validity period of the Storage Capacity Unit. Default value: `Month`. Valid values: `Month`, `Year`.
     */
    public readonly periodUnit!: pulumi.Output<string>;
    /**
     * The time when the Storage Capacity Unit takes effect. It cannot be earlier than or more than six months later than the time when the Storage Capacity Unit is created. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. **NOTE:** This parameter is empty by default. The Storage Capacity Unit immediately takes effect after it is created.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * The status of Storage Capacity Unit.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of the Storage Capacity Unit.
     */
    public readonly storageCapacityUnitName!: pulumi.Output<string>;

    /**
     * Create a StorageCapacityUnit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageCapacityUnitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageCapacityUnitArgs | StorageCapacityUnitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StorageCapacityUnitState | undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageCapacityUnitName"] = state ? state.storageCapacityUnitName : undefined;
        } else {
            const args = argsOrState as StorageCapacityUnitArgs | undefined;
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["storageCapacityUnitName"] = args ? args.storageCapacityUnitName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StorageCapacityUnit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StorageCapacityUnit resources.
 */
export interface StorageCapacityUnitState {
    /**
     * The capacity of the Storage Capacity Unit. Unit: GiB. Valid values: `20`, `40`, `100`, `200`, `500`, `1024`, `2048`, `5120`, `10240`, `20480`, and `51200`.
     */
    capacity?: pulumi.Input<number>;
    /**
     * The description of the Storage Capacity Unit. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The validity period of the Storage Capacity Unit. Default value: `1`.
     * * When PeriodUnit is set to Month, Valid values: `1`, `2`, `3`, `6`.
     * * When PeriodUnit is set to Year, Valid values: `1`, `3`, `5`.
     */
    period?: pulumi.Input<number>;
    /**
     * The unit of the validity period of the Storage Capacity Unit. Default value: `Month`. Valid values: `Month`, `Year`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The time when the Storage Capacity Unit takes effect. It cannot be earlier than or more than six months later than the time when the Storage Capacity Unit is created. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. **NOTE:** This parameter is empty by default. The Storage Capacity Unit immediately takes effect after it is created.
     */
    startTime?: pulumi.Input<string>;
    /**
     * The status of Storage Capacity Unit.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the Storage Capacity Unit.
     */
    storageCapacityUnitName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StorageCapacityUnit resource.
 */
export interface StorageCapacityUnitArgs {
    /**
     * The capacity of the Storage Capacity Unit. Unit: GiB. Valid values: `20`, `40`, `100`, `200`, `500`, `1024`, `2048`, `5120`, `10240`, `20480`, and `51200`.
     */
    capacity: pulumi.Input<number>;
    /**
     * The description of the Storage Capacity Unit. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The validity period of the Storage Capacity Unit. Default value: `1`.
     * * When PeriodUnit is set to Month, Valid values: `1`, `2`, `3`, `6`.
     * * When PeriodUnit is set to Year, Valid values: `1`, `3`, `5`.
     */
    period?: pulumi.Input<number>;
    /**
     * The unit of the validity period of the Storage Capacity Unit. Default value: `Month`. Valid values: `Month`, `Year`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The time when the Storage Capacity Unit takes effect. It cannot be earlier than or more than six months later than the time when the Storage Capacity Unit is created. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. **NOTE:** This parameter is empty by default. The Storage Capacity Unit immediately takes effect after it is created.
     */
    startTime?: pulumi.Input<string>;
    /**
     * The name of the Storage Capacity Unit.
     */
    storageCapacityUnitName?: pulumi.Input<string>;
}
