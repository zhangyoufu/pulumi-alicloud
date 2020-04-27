// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of DNS Domains in an Alibaba Cloud account according to the specified filters.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const domainsDs = pulumi.output(alicloud.dns.getDomains({
 *     domainNameRegex: "^hegu",
 *     outputFile: "domains.txt",
 * }, { async: true }));
 * 
 * export const firstDomainId = domainsDs.domains[0].domainId;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/dns_domains.html.markdown.
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:dns/getDomains:getDomains", {
        "aliDomain": args.aliDomain,
        "domainNameRegex": args.domainNameRegex,
        "groupNameRegex": args.groupNameRegex,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "versionCode": args.versionCode,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * Specifies whether the domain is from Alibaba Cloud or not.
     */
    readonly aliDomain?: boolean;
    /**
     * A regex string to filter results by the domain name. 
     */
    readonly domainNameRegex?: string;
    /**
     * A regex string to filter results by the group name.
     */
    readonly groupNameRegex?: string;
    /**
     * - A list of domain IDs.
     */
    readonly ids?: string[];
    /**
     * Cloud analysis product ID.
     */
    readonly instanceId?: string;
    readonly outputFile?: string;
    /**
     * The Id of resource group which the dns belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * Cloud analysis version code.
     */
    readonly versionCode?: string;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    /**
     * Indicates whether the domain is an Alibaba Cloud domain.
     */
    readonly aliDomain?: boolean;
    readonly domainNameRegex?: string;
    /**
     * A list of domains. Each element contains the following attributes:
     */
    readonly domains: outputs.dns.GetDomainsDomain[];
    readonly groupNameRegex?: string;
    /**
     * A list of domain IDs.
     */
    readonly ids: string[];
    /**
     * Cloud analysis product ID of the domain.
     */
    readonly instanceId?: string;
    /**
     * A list of domain names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The Id of resource group which the dns belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * Cloud analysis version code of the domain.
     */
    readonly versionCode?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
