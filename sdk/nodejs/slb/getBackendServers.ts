// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the server load balancer backend servers related to a server load balancer..
 * 
 * > **NOTE:** Available in 1.53.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const sampleDs = alicloud_slb_sample_slb.id.apply(id => alicloud.SlbBeckendServers({
 *     loadBalancerId: id,
 * }, { async: true }));
 * 
 * export const firstSlbBackendServerId = sampleDs.backendServers.0.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_backend_servers.html.markdown.
 */
export function getBackendServers(args: GetBackendServersArgs, opts?: pulumi.InvokeOptions): Promise<GetBackendServersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:slb/getBackendServers:getBackendServers", {
        "ids": args.ids,
        "loadBalancerId": args.loadBalancerId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackendServers.
 */
export interface GetBackendServersArgs {
    /**
     * List of attached ECS instance IDs.
     */
    readonly ids?: string[];
    /**
     * ID of the SLB with attachments.
     */
    readonly loadBalancerId: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getBackendServers.
 */
export interface GetBackendServersResult {
    readonly backendServers: outputs.slb.GetBackendServersBackendServer[];
    readonly ids: string[];
    readonly loadBalancerId: string;
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
