// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Amqp Bindings of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.135.0+.
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
 * const examples = alicloud.amqp.getBindings({
 *     instanceId: "amqp-cn-xxxxx",
 *     virtualHostName: "my-vh",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBindings(args: GetBindingsArgs, opts?: pulumi.InvokeOptions): Promise<GetBindingsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:amqp/getBindings:getBindings", {
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "virtualHostName": args.virtualHostName,
    }, opts);
}

/**
 * A collection of arguments for invoking getBindings.
 */
export interface GetBindingsArgs {
    /**
     * Instance Id.
     */
    instanceId: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Virtualhost Name.
     */
    virtualHostName: string;
}

/**
 * A collection of values returned by getBindings.
 */
export interface GetBindingsResult {
    readonly bindings: outputs.amqp.GetBindingsBinding[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId: string;
    readonly outputFile?: string;
    readonly virtualHostName: string;
}
/**
 * This data source provides the Amqp Bindings of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.135.0+.
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
 * const examples = alicloud.amqp.getBindings({
 *     instanceId: "amqp-cn-xxxxx",
 *     virtualHostName: "my-vh",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBindingsOutput(args: GetBindingsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBindingsResult> {
    return pulumi.output(args).apply((a: any) => getBindings(a, opts))
}

/**
 * A collection of arguments for invoking getBindings.
 */
export interface GetBindingsOutputArgs {
    /**
     * Instance Id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Virtualhost Name.
     */
    virtualHostName: pulumi.Input<string>;
}
