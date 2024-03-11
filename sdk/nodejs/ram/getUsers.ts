// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of RAM users in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available since v1.0.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultGroup = new alicloud.ram.Group("defaultGroup", {
 *     comments: "group comments",
 *     force: true,
 * });
 * const defaultUser = new alicloud.ram.User("defaultUser", {
 *     displayName: "displayname",
 *     mobile: "86-18888888888",
 *     email: "hello.uuu@aaa.com",
 *     comments: "yoyoyo",
 * });
 * const defaultGroupMembership = new alicloud.ram.GroupMembership("defaultGroupMembership", {
 *     groupName: defaultGroup.name,
 *     userNames: [defaultUser.name],
 * });
 * const defaultPolicy = new alicloud.ram.Policy("defaultPolicy", {
 *     policyName: "ram-policy-example",
 *     policyDocument: `			{
 * 				"Statement": [
 * 				 {
 * 					"Action": [
 * 					"oss:ListObjects",
 * 					"oss:ListObjects"
 * 			  		],
 * 			  		"Effect": "Deny",
 * 			  		"Resource": [
 * 						"acs:oss:*:*:mybucket",
 * 						"acs:oss:*:*:mybucket/*"
 * 			  		]
 * 				 }
 * 		  		],
 * 				"Version": "1"
 * 			}
 * `,
 *     description: "this is a policy example",
 *     force: true,
 * });
 * const defaultUserPolicyAttachment = new alicloud.ram.UserPolicyAttachment("defaultUserPolicyAttachment", {
 *     policyName: defaultPolicy.policyName,
 *     userName: defaultUser.name,
 *     policyType: defaultPolicy.type,
 * });
 * const usersDs = alicloud.ram.getUsersOutput({
 *     outputFile: "users.txt",
 *     groupName: defaultGroup.name,
 *     policyName: defaultPolicy.policyName,
 *     policyType: "Custom",
 *     nameRegex: defaultUser.name,
 * });
 * export const firstUserId = usersDs.apply(usersDs => usersDs.users?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ram/getUsers:getUsers", {
        "groupName": args.groupName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "policyName": args.policyName,
        "policyType": args.policyType,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * Filter results by a specific group name. Returned users are in the specified group.
     */
    groupName?: string;
    /**
     * A list of ram user IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter resulting users by their names.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * Filter results by a specific policy name. If you set this parameter without setting `policyType`, the later will be automatically set to `System`. Returned users are attached to the specified policy.
     */
    policyName?: string;
    /**
     * Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policyName` as well.
     */
    policyType?: string;
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    readonly groupName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of ram user IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of ram user's name.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly policyName?: string;
    readonly policyType?: string;
    /**
     * A list of users. Each element contains the following attributes:
     */
    readonly users: outputs.ram.GetUsersUser[];
}
/**
 * This data source provides a list of RAM users in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available since v1.0.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultGroup = new alicloud.ram.Group("defaultGroup", {
 *     comments: "group comments",
 *     force: true,
 * });
 * const defaultUser = new alicloud.ram.User("defaultUser", {
 *     displayName: "displayname",
 *     mobile: "86-18888888888",
 *     email: "hello.uuu@aaa.com",
 *     comments: "yoyoyo",
 * });
 * const defaultGroupMembership = new alicloud.ram.GroupMembership("defaultGroupMembership", {
 *     groupName: defaultGroup.name,
 *     userNames: [defaultUser.name],
 * });
 * const defaultPolicy = new alicloud.ram.Policy("defaultPolicy", {
 *     policyName: "ram-policy-example",
 *     policyDocument: `			{
 * 				"Statement": [
 * 				 {
 * 					"Action": [
 * 					"oss:ListObjects",
 * 					"oss:ListObjects"
 * 			  		],
 * 			  		"Effect": "Deny",
 * 			  		"Resource": [
 * 						"acs:oss:*:*:mybucket",
 * 						"acs:oss:*:*:mybucket/*"
 * 			  		]
 * 				 }
 * 		  		],
 * 				"Version": "1"
 * 			}
 * `,
 *     description: "this is a policy example",
 *     force: true,
 * });
 * const defaultUserPolicyAttachment = new alicloud.ram.UserPolicyAttachment("defaultUserPolicyAttachment", {
 *     policyName: defaultPolicy.policyName,
 *     userName: defaultUser.name,
 *     policyType: defaultPolicy.type,
 * });
 * const usersDs = alicloud.ram.getUsersOutput({
 *     outputFile: "users.txt",
 *     groupName: defaultGroup.name,
 *     policyName: defaultPolicy.policyName,
 *     policyType: "Custom",
 *     nameRegex: defaultUser.name,
 * });
 * export const firstUserId = usersDs.apply(usersDs => usersDs.users?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUsersOutput(args?: GetUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsersResult> {
    return pulumi.output(args).apply((a: any) => getUsers(a, opts))
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * Filter results by a specific group name. Returned users are in the specified group.
     */
    groupName?: pulumi.Input<string>;
    /**
     * A list of ram user IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter resulting users by their names.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Filter results by a specific policy name. If you set this parameter without setting `policyType`, the later will be automatically set to `System`. Returned users are attached to the specified policy.
     */
    policyName?: pulumi.Input<string>;
    /**
     * Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policyName` as well.
     */
    policyType?: pulumi.Input<string>;
}
