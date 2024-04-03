// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cloud Monitor Service Hybrid Double Writes of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available since v1.220.0.
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
 * const defaultAccount = alicloud.getAccount({});
 * const source = new alicloud.cms.Namespace("source", {namespace: "your-source-namespace"});
 * const defaultNamespace = new alicloud.cms.Namespace("defaultNamespace", {namespace: "your-namespace"});
 * const defaultServiceHybridDoubleWrite = new alicloud.cloudmonitor.ServiceHybridDoubleWrite("defaultServiceHybridDoubleWrite", {
 *     sourceNamespace: source.id,
 *     sourceUserId: defaultAccount.then(defaultAccount => defaultAccount.id),
 *     namespace: defaultNamespace.id,
 *     userId: defaultAccount.then(defaultAccount => defaultAccount.id),
 * });
 * const ids = alicloud.cloudmonitor.getServiceHybridDoubleWritesOutput({
 *     ids: [defaultServiceHybridDoubleWrite.id],
 * });
 * export const cloudMonitorServiceHybridDoubleWritesId1 = ids.apply(ids => ids.hybridDoubleWrites?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceHybridDoubleWrites(args?: GetServiceHybridDoubleWritesArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceHybridDoubleWritesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cloudmonitor/getServiceHybridDoubleWrites:getServiceHybridDoubleWrites", {
        "ids": args.ids,
        "namespace": args.namespace,
        "outputFile": args.outputFile,
        "sourceNamespace": args.sourceNamespace,
        "sourceUserId": args.sourceUserId,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceHybridDoubleWrites.
 */
export interface GetServiceHybridDoubleWritesArgs {
    /**
     * A list of Hybrid Double Write IDs.
     */
    ids?: string[];
    /**
     * Target Namespace.
     */
    namespace?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Source Namespace.
     */
    sourceNamespace?: string;
    /**
     * Source UserId.
     */
    sourceUserId?: string;
    /**
     * Target UserId.
     */
    userId?: string;
}

/**
 * A collection of values returned by getServiceHybridDoubleWrites.
 */
export interface GetServiceHybridDoubleWritesResult {
    /**
     * A list of Hybrid Double Writes. Each element contains the following attributes:
     */
    readonly hybridDoubleWrites: outputs.cloudmonitor.GetServiceHybridDoubleWritesHybridDoubleWrite[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    /**
     * Target Namespace.
     */
    readonly namespace?: string;
    readonly outputFile?: string;
    /**
     * Source Namespace.
     */
    readonly sourceNamespace?: string;
    /**
     * Source UserId.
     */
    readonly sourceUserId?: string;
    /**
     * Target UserId.
     */
    readonly userId?: string;
}
/**
 * This data source provides the Cloud Monitor Service Hybrid Double Writes of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available since v1.220.0.
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
 * const defaultAccount = alicloud.getAccount({});
 * const source = new alicloud.cms.Namespace("source", {namespace: "your-source-namespace"});
 * const defaultNamespace = new alicloud.cms.Namespace("defaultNamespace", {namespace: "your-namespace"});
 * const defaultServiceHybridDoubleWrite = new alicloud.cloudmonitor.ServiceHybridDoubleWrite("defaultServiceHybridDoubleWrite", {
 *     sourceNamespace: source.id,
 *     sourceUserId: defaultAccount.then(defaultAccount => defaultAccount.id),
 *     namespace: defaultNamespace.id,
 *     userId: defaultAccount.then(defaultAccount => defaultAccount.id),
 * });
 * const ids = alicloud.cloudmonitor.getServiceHybridDoubleWritesOutput({
 *     ids: [defaultServiceHybridDoubleWrite.id],
 * });
 * export const cloudMonitorServiceHybridDoubleWritesId1 = ids.apply(ids => ids.hybridDoubleWrites?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceHybridDoubleWritesOutput(args?: GetServiceHybridDoubleWritesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceHybridDoubleWritesResult> {
    return pulumi.output(args).apply((a: any) => getServiceHybridDoubleWrites(a, opts))
}

/**
 * A collection of arguments for invoking getServiceHybridDoubleWrites.
 */
export interface GetServiceHybridDoubleWritesOutputArgs {
    /**
     * A list of Hybrid Double Write IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Target Namespace.
     */
    namespace?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Source Namespace.
     */
    sourceNamespace?: pulumi.Input<string>;
    /**
     * Source UserId.
     */
    sourceUserId?: pulumi.Input<string>;
    /**
     * Target UserId.
     */
    userId?: pulumi.Input<string>;
}
