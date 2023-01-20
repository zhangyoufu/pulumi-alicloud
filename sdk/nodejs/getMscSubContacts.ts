// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This data source provides the Message Center Contacts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.getMscSubContacts({});
 * export const mscSubContactId1 = ids.then(ids => ids.contacts?.[0]?.id);
 * const nameRegex = alicloud.getMscSubContacts({
 *     nameRegex: "^my-Contact",
 * });
 * export const mscSubContactId2 = nameRegex.then(nameRegex => nameRegex.contacts?.[0]?.id);
 * ```
 */
export function getMscSubContacts(args?: GetMscSubContactsArgs, opts?: pulumi.InvokeOptions): Promise<GetMscSubContactsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:index/getMscSubContacts:getMscSubContacts", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMscSubContacts.
 */
export interface GetMscSubContactsArgs {
    /**
     * A list of Contact IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Contact name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getMscSubContacts.
 */
export interface GetMscSubContactsResult {
    readonly contacts: outputs.GetMscSubContactsContact[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Message Center Contacts of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.getMscSubContacts({});
 * export const mscSubContactId1 = ids.then(ids => ids.contacts?.[0]?.id);
 * const nameRegex = alicloud.getMscSubContacts({
 *     nameRegex: "^my-Contact",
 * });
 * export const mscSubContactId2 = nameRegex.then(nameRegex => nameRegex.contacts?.[0]?.id);
 * ```
 */
export function getMscSubContactsOutput(args?: GetMscSubContactsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMscSubContactsResult> {
    return pulumi.output(args).apply((a: any) => getMscSubContacts(a, opts))
}

/**
 * A collection of arguments for invoking getMscSubContacts.
 */
export interface GetMscSubContactsOutputArgs {
    /**
     * A list of Contact IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Contact name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
