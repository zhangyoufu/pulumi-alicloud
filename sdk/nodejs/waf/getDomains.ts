// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a WAF datasource to retrieve domains.
 *
 * For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
 *
 * > **NOTE:** Available in 1.86.0+ .
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.waf.getDomains({
 *     instanceId: "waf-cf-xxxxx",
 * });
 * ```
 */
export function getDomains(args: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:waf/getDomains:getDomains", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * Default to false and only output `id`, `domainName`. Set it to true can output more details.
     */
    enableDetails?: boolean;
    /**
     * A list of WAF domain names. Each item is domain name.
     */
    ids?: string[];
    /**
     * The Id of waf instance to which waf domain belongs.
     */
    instanceId: string;
    /**
     * A regex string to filter results by domain name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management.
     */
    resourceGroupId?: string;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    /**
     * A list of Domains. Each element contains the following attributes:
     */
    readonly domains: outputs.waf.GetDomainsDomain[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of WAF domain self ID, value as `domainName`.
     */
    readonly ids: string[];
    readonly instanceId: string;
    readonly nameRegex?: string;
    /**
     * A list of WAF domain names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management.
     */
    readonly resourceGroupId?: string;
}
/**
 * Provides a WAF datasource to retrieve domains.
 *
 * For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
 *
 * > **NOTE:** Available in 1.86.0+ .
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.waf.getDomains({
 *     instanceId: "waf-cf-xxxxx",
 * });
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
     * Default to false and only output `id`, `domainName`. Set it to true can output more details.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of WAF domain names. Each item is domain name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Id of waf instance to which waf domain belongs.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A regex string to filter results by domain name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management.
     */
    resourceGroupId?: pulumi.Input<string>;
}
