// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of DNS Domains in an Alibaba Cloud account according to the specified filters.
 *
 * ## Example Usage
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
        "groupId": args.groupId,
        "groupNameRegex": args.groupNameRegex,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "keyWord": args.keyWord,
        "lang": args.lang,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "searchMode": args.searchMode,
        "starmark": args.starmark,
        "tags": args.tags,
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
     * Domain group ID, if not filled, the default is all groups.
     */
    readonly groupId?: string;
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
    /**
     * The keywords are searched according to the `%KeyWord%` mode, which is not case sensitive.
     */
    readonly keyWord?: string;
    /**
     * User language.
     */
    readonly lang?: string;
    readonly outputFile?: string;
    /**
     * The Id of resource group which the dns belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * Search mode, `LIKE` fuzzy search, `EXACT` exact search.
     */
    readonly searchMode?: string;
    /**
     * Whether to query the domain name star.
     */
    readonly starmark?: boolean;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
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
    /**
     * Id of group that contains the domain.
     */
    readonly groupId?: string;
    readonly groupNameRegex?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of domain IDs.
     */
    readonly ids: string[];
    /**
     * Cloud analysis product ID of the domain.
     */
    readonly instanceId?: string;
    readonly keyWord?: string;
    readonly lang?: string;
    /**
     * A list of domain names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The Id of resource group which the dns belongs.
     */
    readonly resourceGroupId?: string;
    readonly searchMode?: string;
    readonly starmark?: boolean;
    readonly tags?: {[key: string]: any};
    /**
     * Cloud resolution version ID.
     */
    readonly versionCode?: string;
}
