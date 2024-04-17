// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > **NOTE:** Available in v1.156.0+.
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
 * const defaultMscSubContract = new alicloud.MscSubContract("default", {
 *     contactName: "example_value",
 *     position: "CEO",
 *     email: "123@163.com",
 *     mobile: "153xxxxx906",
 * });
 * const default = defaultMscSubContract.id.apply(id => alicloud.getMscSubContactVerificationMessageOutput({
 *     contactId: id,
 *     type: 1,
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMscSubContactVerificationMessage(args: GetMscSubContactVerificationMessageArgs, opts?: pulumi.InvokeOptions): Promise<GetMscSubContactVerificationMessageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:index/getMscSubContactVerificationMessage:getMscSubContactVerificationMessage", {
        "contactId": args.contactId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getMscSubContactVerificationMessage.
 */
export interface GetMscSubContactVerificationMessageArgs {
    /**
     * The ID of the Contact.
     */
    contactId: string;
    /**
     * How a user receives verification messages. Valid values : `1`, `2`.
     */
    type: number;
}

/**
 * A collection of values returned by getMscSubContactVerificationMessage.
 */
export interface GetMscSubContactVerificationMessageResult {
    readonly contactId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The sending status of the message. Valid values : `Success`, `Failed`.
     */
    readonly status: string;
    readonly type: number;
}
/**
 * > **NOTE:** Available in v1.156.0+.
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
 * const defaultMscSubContract = new alicloud.MscSubContract("default", {
 *     contactName: "example_value",
 *     position: "CEO",
 *     email: "123@163.com",
 *     mobile: "153xxxxx906",
 * });
 * const default = defaultMscSubContract.id.apply(id => alicloud.getMscSubContactVerificationMessageOutput({
 *     contactId: id,
 *     type: 1,
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMscSubContactVerificationMessageOutput(args: GetMscSubContactVerificationMessageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMscSubContactVerificationMessageResult> {
    return pulumi.output(args).apply((a: any) => getMscSubContactVerificationMessage(a, opts))
}

/**
 * A collection of arguments for invoking getMscSubContactVerificationMessage.
 */
export interface GetMscSubContactVerificationMessageOutputArgs {
    /**
     * The ID of the Contact.
     */
    contactId: pulumi.Input<string>;
    /**
     * How a user receives verification messages. Valid values : `1`, `2`.
     */
    type: pulumi.Input<number>;
}
