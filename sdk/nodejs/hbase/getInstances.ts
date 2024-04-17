// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
 * Filters support regular expression for the instance name, ids or availability_zone.
 *
 * > **NOTE:**  Available in 1.67.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const hbase = alicloud.hbase.getInstances({
 *     nameRegex: "tf_testAccHBase",
 *     availabilityZone: "cn-shenzhen-b",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:hbase/getInstances:getInstances", {
        "availabilityZone": args.availabilityZone,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Instance availability zone.
     */
    availabilityZone?: string;
    /**
     * The ids list of HBase instances
     */
    ids?: string[];
    /**
     * A regex string to apply to the instance name.
     */
    nameRegex?: string;
    /**
     * The name of file that can save the collection of instances after running `pulumi preview`.
     */
    outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly availabilityZone?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ids list of HBase instances
     */
    readonly ids: string[];
    /**
     * A list of HBase instances. Its every element contains the following attributes:
     */
    readonly instances: outputs.hbase.GetInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * The names list of HBase instances
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}
/**
 * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
 * Filters support regular expression for the instance name, ids or availability_zone.
 *
 * > **NOTE:**  Available in 1.67.0+
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const hbase = alicloud.hbase.getInstances({
 *     nameRegex: "tf_testAccHBase",
 *     availabilityZone: "cn-shenzhen-b",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * Instance availability zone.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The ids list of HBase instances
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to apply to the instance name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The name of file that can save the collection of instances after running `pulumi preview`.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
