// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecd Ad Connector Office Sites of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.176.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.eds.getAdConnectorOfficeSites({});
 * export const ecdAdConnectorOfficeSiteId1 = ids.then(ids => ids.sites?.[0]?.id);
 * const nameRegex = alicloud.eds.getAdConnectorOfficeSites({
 *     nameRegex: "^my-AdConnectorOfficeSite",
 * });
 * export const ecdAdConnectorOfficeSiteId2 = nameRegex.then(nameRegex => nameRegex.sites?.[0]?.id);
 * ```
 */
export function getAdConnectorOfficeSites(args?: GetAdConnectorOfficeSitesArgs, opts?: pulumi.InvokeOptions): Promise<GetAdConnectorOfficeSitesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:eds/getAdConnectorOfficeSites:getAdConnectorOfficeSites", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getAdConnectorOfficeSites.
 */
export interface GetAdConnectorOfficeSitesArgs {
    /**
     * A list of Ad Connector Office Site IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Ad Connector Office Site name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The workspace status.
     */
    status?: string;
}

/**
 * A collection of values returned by getAdConnectorOfficeSites.
 */
export interface GetAdConnectorOfficeSitesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly sites: outputs.eds.GetAdConnectorOfficeSitesSite[];
    readonly status?: string;
}
/**
 * This data source provides the Ecd Ad Connector Office Sites of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.176.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.eds.getAdConnectorOfficeSites({});
 * export const ecdAdConnectorOfficeSiteId1 = ids.then(ids => ids.sites?.[0]?.id);
 * const nameRegex = alicloud.eds.getAdConnectorOfficeSites({
 *     nameRegex: "^my-AdConnectorOfficeSite",
 * });
 * export const ecdAdConnectorOfficeSiteId2 = nameRegex.then(nameRegex => nameRegex.sites?.[0]?.id);
 * ```
 */
export function getAdConnectorOfficeSitesOutput(args?: GetAdConnectorOfficeSitesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAdConnectorOfficeSitesResult> {
    return pulumi.output(args).apply((a: any) => getAdConnectorOfficeSites(a, opts))
}

/**
 * A collection of arguments for invoking getAdConnectorOfficeSites.
 */
export interface GetAdConnectorOfficeSitesOutputArgs {
    /**
     * A list of Ad Connector Office Site IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Ad Connector Office Site name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The workspace status.
     */
    status?: pulumi.Input<string>;
}
