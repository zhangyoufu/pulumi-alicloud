// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the domain extensions associated with a server load balancer listener.
 *
 * > **NOTE:** Available in 1.60.0+
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const foo = pulumi.output(alicloud.slb.getDomainExtensions({
 *     frontendPort: Number.parseFloat("fake-port"),
 *     ids: ["fake-de-id"],
 *     loadBalancerId: "fake-lb-id",
 * }, { async: true }));
 * ```
 */
export function getDomainExtensions(args: GetDomainExtensionsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainExtensionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:slb/getDomainExtensions:getDomainExtensions", {
        "frontendPort": args.frontendPort,
        "ids": args.ids,
        "loadBalancerId": args.loadBalancerId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomainExtensions.
 */
export interface GetDomainExtensionsArgs {
    /**
     * The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
     */
    readonly frontendPort: number;
    /**
     * IDs of the SLB domain extensions.
     */
    readonly ids?: string[];
    /**
     * The ID of the SLB instance.
     */
    readonly loadBalancerId: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getDomainExtensions.
 */
export interface GetDomainExtensionsResult {
    /**
     * A list of SLB domain extension. Each element contains the following attributes:
     */
    readonly extensions: outputs.slb.GetDomainExtensionsExtension[];
    readonly frontendPort: number;
    readonly ids: string[];
    readonly loadBalancerId: string;
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
