// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Wafv3 Domain available to the user.[What is Domain](https://www.alibabacloud.com/help/en/web-application-firewall/latest/api-doc-waf-openapi-2021-10-01-api-doc-createdomain)
 *
 * > **NOTE:** Available in 1.200.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.wafv3.getDomains({
 *     domain: "zctest12.wafqax.top",
 *     instanceId: "waf_v3prepaid_public_cn-*****",
 * });
 * export const alicloudWafv3DomainExampleId = _default.then(_default => _default.domains?.[0]?.id);
 * ```
 */
export function getDomains(args: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:wafv3/getDomains:getDomains", {
        "backend": args.backend,
        "domain": args.domain,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
     */
    backend?: string;
    /**
     * The name of the domain name to query.
     */
    domain?: string;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of domain IDs.
     */
    ids?: string[];
    /**
     * The WAF instance ID.
     */
    instanceId: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    readonly backend?: string;
    /**
     * The name of the domain.
     */
    readonly domain?: string;
    /**
     * A list of Domain Entries. Each element contains the following attributes:
     */
    readonly domains: outputs.wafv3.GetDomainsDomain[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId: string;
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
}
/**
 * This data source provides Wafv3 Domain available to the user.[What is Domain](https://www.alibabacloud.com/help/en/web-application-firewall/latest/api-doc-waf-openapi-2021-10-01-api-doc-createdomain)
 *
 * > **NOTE:** Available in 1.200.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.wafv3.getDomains({
 *     domain: "zctest12.wafqax.top",
 *     instanceId: "waf_v3prepaid_public_cn-*****",
 * });
 * export const alicloudWafv3DomainExampleId = _default.then(_default => _default.domains?.[0]?.id);
 * ```
 */
export function getDomainsOutput(args: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply((a: any) => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    /**
     * The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
     */
    backend?: pulumi.Input<string>;
    /**
     * The name of the domain name to query.
     */
    domain?: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of domain IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The WAF instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
}
