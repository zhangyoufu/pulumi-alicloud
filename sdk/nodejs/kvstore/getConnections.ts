// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can query the public IP of the specified KVStore DBInstance.
 *
 * > **NOTE:** Available in v1.101.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const example = alicloud.kvstore.getConnections({
 *     ids: "r-wer123456",
 * });
 * export const connectionString = example.then(example => example.connections?.[0]?.connectionString);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getConnections(args: GetConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:kvstore/getConnections:getConnections", {
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getConnections.
 */
export interface GetConnectionsArgs {
    /**
     * A list of KVStore DBInstance ids, only support one item.
     */
    ids: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getConnections.
 */
export interface GetConnectionsResult {
    /**
     * Public network details of the specified resource. contains the following attributes:
     */
    readonly connections: outputs.kvstore.GetConnectionsConnection[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of KVStore DBInstance ids.
     */
    readonly ids: string;
    readonly outputFile?: string;
}
/**
 * This data source can query the public IP of the specified KVStore DBInstance.
 *
 * > **NOTE:** Available in v1.101.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const example = alicloud.kvstore.getConnections({
 *     ids: "r-wer123456",
 * });
 * export const connectionString = example.then(example => example.connections?.[0]?.connectionString);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getConnectionsOutput(args: GetConnectionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionsResult> {
    return pulumi.output(args).apply((a: any) => getConnections(a, opts))
}

/**
 * A collection of arguments for invoking getConnections.
 */
export interface GetConnectionsOutputArgs {
    /**
     * A list of KVStore DBInstance ids, only support one item.
     */
    ids: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
