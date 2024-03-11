// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Security Group resource.
 *
 * For information about Security Group and how to use it, see [What is Security Group](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-createsecuritygroup).
 *
 * > **NOTE:** Available since v1.0.0.
 *
 * > **NOTE:** `alicloud.ecs.SecurityGroup` is used to build and manage a security group, and `alicloud.ecs.SecurityGroupRule` can define ingress or egress rules for it.
 *
 * > **NOTE:** From version 1.7.2, `alicloud.ecs.SecurityGroup` has supported to segregate different ECS instance in which the same security group.
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
 * const _default = new alicloud.ecs.SecurityGroup("default", {description: "New security group"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * Basic Usage for VPC
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const vpc = new alicloud.vpc.Network("vpc", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "10.1.0.0/21",
 * });
 * const group = new alicloud.ecs.SecurityGroup("group", {vpcId: vpc.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Module Support
 *
 * You can use the existing security-group module
 * to create a security group and add several rules one-click.
 *
 * ## Import
 *
 * Security Group can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecs/securityGroup:SecurityGroup example sg-abc123456
 * ```
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * The security group description. Defaults to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Field `innerAccess` has been deprecated from provider version 1.55.3. New field `innerAccessPolicy` instead.
     *
     * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
     *
     * @deprecated Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
     */
    public readonly innerAccess!: pulumi.Output<boolean>;
    /**
     * The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
     */
    public readonly innerAccessPolicy!: pulumi.Output<string>;
    /**
     * The name of the security group. Defaults to null.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resourceGroupId` can be modified.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The type of the security group. Valid values:
     */
    public readonly securityGroupType!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["innerAccess"] = state ? state.innerAccess : undefined;
            resourceInputs["innerAccessPolicy"] = state ? state.innerAccessPolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityGroupType"] = state ? state.securityGroupType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["innerAccess"] = args ? args.innerAccess : undefined;
            resourceInputs["innerAccessPolicy"] = args ? args.innerAccessPolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityGroupType"] = args ? args.securityGroupType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * The security group description. Defaults to null.
     */
    description?: pulumi.Input<string>;
    /**
     * Field `innerAccess` has been deprecated from provider version 1.55.3. New field `innerAccessPolicy` instead.
     *
     * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
     *
     * @deprecated Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
     */
    innerAccess?: pulumi.Input<boolean>;
    /**
     * The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
     */
    innerAccessPolicy?: pulumi.Input<string>;
    /**
     * The name of the security group. Defaults to null.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resourceGroupId` can be modified.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The type of the security group. Valid values:
     */
    securityGroupType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * The security group description. Defaults to null.
     */
    description?: pulumi.Input<string>;
    /**
     * Field `innerAccess` has been deprecated from provider version 1.55.3. New field `innerAccessPolicy` instead.
     *
     * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
     *
     * @deprecated Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
     */
    innerAccess?: pulumi.Input<boolean>;
    /**
     * The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
     */
    innerAccessPolicy?: pulumi.Input<string>;
    /**
     * The name of the security group. Defaults to null.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resourceGroupId` can be modified.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The type of the security group. Valid values:
     */
    securityGroupType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}
