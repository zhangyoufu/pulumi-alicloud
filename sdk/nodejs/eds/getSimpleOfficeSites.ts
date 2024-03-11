// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecd Simple Office Sites of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.140.0+.
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
 * const default = alicloud.eds.getSimpleOfficeSites({
 *     ids: ["example_id"],
 *     status: "REGISTERED",
 * });
 * export const desktopAccessType = _default.then(_default => _default.sites?.[0]?.desktopAccessType);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSimpleOfficeSites(args?: GetSimpleOfficeSitesArgs, opts?: pulumi.InvokeOptions): Promise<GetSimpleOfficeSitesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:eds/getSimpleOfficeSites:getSimpleOfficeSites", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getSimpleOfficeSites.
 */
export interface GetSimpleOfficeSitesArgs {
    /**
     * A list of Simple Office Site IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Simple Office Site name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Workspace State. Possible Values: Registering: Registered in the Registered: Registered.
     */
    status?: string;
}

/**
 * A collection of values returned by getSimpleOfficeSites.
 */
export interface GetSimpleOfficeSitesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly sites: outputs.eds.GetSimpleOfficeSitesSite[];
    readonly status?: string;
}
/**
 * This data source provides the Ecd Simple Office Sites of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.140.0+.
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
 * const default = alicloud.eds.getSimpleOfficeSites({
 *     ids: ["example_id"],
 *     status: "REGISTERED",
 * });
 * export const desktopAccessType = _default.then(_default => _default.sites?.[0]?.desktopAccessType);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSimpleOfficeSitesOutput(args?: GetSimpleOfficeSitesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSimpleOfficeSitesResult> {
    return pulumi.output(args).apply((a: any) => getSimpleOfficeSites(a, opts))
}

/**
 * A collection of arguments for invoking getSimpleOfficeSites.
 */
export interface GetSimpleOfficeSitesOutputArgs {
    /**
     * A list of Simple Office Site IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Simple Office Site name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Workspace State. Possible Values: Registering: Registered in the Registered: Registered.
     */
    status?: pulumi.Input<string>;
}
