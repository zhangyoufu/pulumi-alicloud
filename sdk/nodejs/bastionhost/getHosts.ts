// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Bastionhost Hosts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.bastionhost.getHosts({
 *     instanceId: "example_value",
 *     ids: [
 *         "1",
 *         "2",
 *     ],
 * });
 * export const bastionhostHostId1 = ids.then(ids => ids.hosts[0].id);
 * const nameRegex = alicloud.bastionhost.getHosts({
 *     instanceId: "example_value",
 *     nameRegex: "^my-Host",
 * });
 * export const bastionhostHostId2 = nameRegex.then(nameRegex => nameRegex.hosts[0].id);
 * ```
 */
export function getHosts(args: GetHostsArgs, opts?: pulumi.InvokeOptions): Promise<GetHostsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:bastionhost/getHosts:getHosts", {
        "enableDetails": args.enableDetails,
        "hostAddress": args.hostAddress,
        "hostName": args.hostName,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "osType": args.osType,
        "outputFile": args.outputFile,
        "source": args.source,
        "sourceInstanceId": args.sourceInstanceId,
        "sourceInstanceState": args.sourceInstanceState,
    }, opts);
}

/**
 * A collection of arguments for invoking getHosts.
 */
export interface GetHostsArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    readonly enableDetails?: boolean;
    /**
     * The host address.
     */
    readonly hostAddress?: string;
    /**
     * Specify the new create a host name of the supports up to 128 characters.
     */
    readonly hostName?: string;
    /**
     * A list of Host IDs.
     */
    readonly ids?: string[];
    /**
     * Specify the new create a host where the Bastion host ID of.
     */
    readonly instanceId: string;
    /**
     * A regex string to filter results by Host name.
     */
    readonly nameRegex?: string;
    /**
     * Specify the new create the host's operating system. Valid values: Linux Windows.
     */
    readonly osType?: string;
    readonly outputFile?: string;
    /**
     * Specify the new create a host of source. Valid values: Local: localhost Ecs:ECS instance Rds:RDS exclusive cluster host.
     */
    readonly source?: string;
    /**
     * Specify the newly created ECS instance ID or dedicated cluster host ID.
     */
    readonly sourceInstanceId?: string;
    /**
     * The source instance state.
     */
    readonly sourceInstanceState?: string;
}

/**
 * A collection of values returned by getHosts.
 */
export interface GetHostsResult {
    readonly enableDetails?: boolean;
    readonly hostAddress?: string;
    readonly hostName?: string;
    readonly hosts: outputs.bastionhost.GetHostsHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly osType?: string;
    readonly outputFile?: string;
    readonly source?: string;
    readonly sourceInstanceId?: string;
    readonly sourceInstanceState?: string;
}
