// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecs Invocations of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.168.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.ecs.getEcsInvocations({
 *     ids: ["example-id"],
 * });
 * export const ecsInvocationId1 = ids.then(ids => ids.invocations?.[0]?.id);
 * ```
 */
export function getEcsInvocations(args?: GetEcsInvocationsArgs, opts?: pulumi.InvokeOptions): Promise<GetEcsInvocationsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecs/getEcsInvocations:getEcsInvocations", {
        "commandId": args.commandId,
        "contentEncoding": args.contentEncoding,
        "ids": args.ids,
        "invokeStatus": args.invokeStatus,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcsInvocations.
 */
export interface GetEcsInvocationsArgs {
    /**
     * The ID of the command.
     */
    commandId?: string;
    /**
     * The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
     */
    contentEncoding?: string;
    /**
     * A list of Invocation IDs.
     */
    ids?: string[];
    /**
     * The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocationStatus` response parameter for the overall execution state.
     */
    invokeStatus?: string;
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getEcsInvocations.
 */
export interface GetEcsInvocationsResult {
    readonly commandId?: string;
    readonly contentEncoding?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly invocations: outputs.ecs.GetEcsInvocationsInvocation[];
    readonly invokeStatus?: string;
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides the Ecs Invocations of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.168.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.ecs.getEcsInvocations({
 *     ids: ["example-id"],
 * });
 * export const ecsInvocationId1 = ids.then(ids => ids.invocations?.[0]?.id);
 * ```
 */
export function getEcsInvocationsOutput(args?: GetEcsInvocationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEcsInvocationsResult> {
    return pulumi.output(args).apply((a: any) => getEcsInvocations(a, opts))
}

/**
 * A collection of arguments for invoking getEcsInvocations.
 */
export interface GetEcsInvocationsOutputArgs {
    /**
     * The ID of the command.
     */
    commandId?: pulumi.Input<string>;
    /**
     * The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
     */
    contentEncoding?: pulumi.Input<string>;
    /**
     * A list of Invocation IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocationStatus` response parameter for the overall execution state.
     */
    invokeStatus?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}
