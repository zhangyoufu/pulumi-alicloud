// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Control Policy Attachments of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.120.0+.
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
 * const example = alicloud.resourcemanager.getControlPolicyAttachments({
 *     targetId: "example_value",
 * });
 * export const firstResourceManagerControlPolicyAttachmentId = example.then(example => example.attachments?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getControlPolicyAttachments(args: GetControlPolicyAttachmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetControlPolicyAttachmentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:resourcemanager/getControlPolicyAttachments:getControlPolicyAttachments", {
        "language": args.language,
        "outputFile": args.outputFile,
        "policyType": args.policyType,
        "targetId": args.targetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getControlPolicyAttachments.
 */
export interface GetControlPolicyAttachmentsArgs {
    /**
     * The language. Valid value `zh-CN`, `en`, and `ja`. Default value `zh-CN`
     */
    language?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The type of policy.
     */
    policyType?: string;
    /**
     * The Id of target.
     */
    targetId: string;
}

/**
 * A collection of values returned by getControlPolicyAttachments.
 */
export interface GetControlPolicyAttachmentsResult {
    readonly attachments: outputs.resourcemanager.GetControlPolicyAttachmentsAttachment[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly language?: string;
    readonly outputFile?: string;
    readonly policyType?: string;
    readonly targetId: string;
}
/**
 * This data source provides the Resource Manager Control Policy Attachments of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.120.0+.
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
 * const example = alicloud.resourcemanager.getControlPolicyAttachments({
 *     targetId: "example_value",
 * });
 * export const firstResourceManagerControlPolicyAttachmentId = example.then(example => example.attachments?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getControlPolicyAttachmentsOutput(args: GetControlPolicyAttachmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetControlPolicyAttachmentsResult> {
    return pulumi.output(args).apply((a: any) => getControlPolicyAttachments(a, opts))
}

/**
 * A collection of arguments for invoking getControlPolicyAttachments.
 */
export interface GetControlPolicyAttachmentsOutputArgs {
    /**
     * The language. Valid value `zh-CN`, `en`, and `ja`. Default value `zh-CN`
     */
    language?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The type of policy.
     */
    policyType?: pulumi.Input<string>;
    /**
     * The Id of target.
     */
    targetId: pulumi.Input<string>;
}
