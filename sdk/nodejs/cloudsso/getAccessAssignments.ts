// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cloud Sso Access Assignments of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.193.0+.
 *
 * > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
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
 * const ids = alicloud.cloudsso.getAccessAssignments({
 *     directoryId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const cloudSsoAccessAssignmentId1 = ids.then(ids => ids.assignments?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAccessAssignments(args: GetAccessAssignmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessAssignmentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cloudsso/getAccessAssignments:getAccessAssignments", {
        "accessConfigurationId": args.accessConfigurationId,
        "directoryId": args.directoryId,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "principalType": args.principalType,
        "targetId": args.targetId,
        "targetType": args.targetType,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessAssignments.
 */
export interface GetAccessAssignmentsArgs {
    /**
     * Access configuration ID.
     */
    accessConfigurationId?: string;
    /**
     * Directory ID.
     */
    directoryId: string;
    /**
     * A list of Access Assignment IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Create the identity type of the access assignment, which can be a user or a user group.
     */
    principalType?: string;
    /**
     * The ID of the target to create the resource range.
     */
    targetId?: string;
    /**
     * The type of the resource range target to be accessed. Only a single RD primary account or member account can be specified in the first phase.
     */
    targetType?: string;
}

/**
 * A collection of values returned by getAccessAssignments.
 */
export interface GetAccessAssignmentsResult {
    readonly accessConfigurationId?: string;
    readonly assignments: outputs.cloudsso.GetAccessAssignmentsAssignment[];
    readonly directoryId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly principalType?: string;
    readonly targetId?: string;
    readonly targetType?: string;
}
/**
 * This data source provides the Cloud Sso Access Assignments of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.193.0+.
 *
 * > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
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
 * const ids = alicloud.cloudsso.getAccessAssignments({
 *     directoryId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const cloudSsoAccessAssignmentId1 = ids.then(ids => ids.assignments?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAccessAssignmentsOutput(args: GetAccessAssignmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessAssignmentsResult> {
    return pulumi.output(args).apply((a: any) => getAccessAssignments(a, opts))
}

/**
 * A collection of arguments for invoking getAccessAssignments.
 */
export interface GetAccessAssignmentsOutputArgs {
    /**
     * Access configuration ID.
     */
    accessConfigurationId?: pulumi.Input<string>;
    /**
     * Directory ID.
     */
    directoryId: pulumi.Input<string>;
    /**
     * A list of Access Assignment IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Create the identity type of the access assignment, which can be a user or a user group.
     */
    principalType?: pulumi.Input<string>;
    /**
     * The ID of the target to create the resource range.
     */
    targetId?: pulumi.Input<string>;
    /**
     * The type of the resource range target to be accessed. Only a single RD primary account or member account can be specified in the first phase.
     */
    targetType?: pulumi.Input<string>;
}
