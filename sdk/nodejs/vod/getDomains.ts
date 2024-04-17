// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vod Domains of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
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
 * const defaultDomain = new alicloud.vod.Domain("default", {
 *     domainName: "your_domain_name",
 *     scope: "domestic",
 *     sources: [{
 *         sourceType: "domain",
 *         sourceContent: "your_source_content",
 *         sourcePort: "80",
 *     }],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 * const default = alicloud.vod.getDomainsOutput({
 *     ids: [defaultDomain.id],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 * export const vodDomain = _default.apply(_default => _default.domains?.[0]);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vod/getDomains:getDomains", {
        "domainSearchType": args.domainSearchType,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * The search method. Valid values:
     */
    domainSearchType?: string;
    /**
     * A list of Domain IDs. Its element value is same as Domain Name.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Domain name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the resource.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    readonly domainSearchType?: string;
    readonly domains: outputs.vod.GetDomainsDomain[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: any};
}
/**
 * This data source provides the Vod Domains of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
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
 * const defaultDomain = new alicloud.vod.Domain("default", {
 *     domainName: "your_domain_name",
 *     scope: "domestic",
 *     sources: [{
 *         sourceType: "domain",
 *         sourceContent: "your_source_content",
 *         sourcePort: "80",
 *     }],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 * const default = alicloud.vod.getDomainsOutput({
 *     ids: [defaultDomain.id],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 * export const vodDomain = _default.apply(_default => _default.domains?.[0]);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDomainsOutput(args?: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply((a: any) => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    /**
     * The search method. Valid values:
     */
    domainSearchType?: pulumi.Input<string>;
    /**
     * A list of Domain IDs. Its element value is same as Domain Name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Domain name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
