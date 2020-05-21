// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
     */
    public readonly innerAccess!: pulumi.Output<boolean>;
    /**
     * Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
     */
    public readonly innerAccessPolicy!: pulumi.Output<string>;
    /**
     * The name of the security group. Defaults to null.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Id of resource group which the securityGroup belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The type of the security group. Valid values:
     * `normal`: basic security group.
     * `enterprise`: advanced security group For more information.
     */
    public readonly securityGroupType!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VPC ID.	
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["innerAccess"] = state ? state.innerAccess : undefined;
            inputs["innerAccessPolicy"] = state ? state.innerAccessPolicy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["securityGroupType"] = state ? state.securityGroupType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["innerAccess"] = args ? args.innerAccess : undefined;
            inputs["innerAccessPolicy"] = args ? args.innerAccessPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["securityGroupType"] = args ? args.securityGroupType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecurityGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * The security group description. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
     * @deprecated Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
     */
    readonly innerAccess?: pulumi.Input<boolean>;
    /**
     * Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
     */
    readonly innerAccessPolicy?: pulumi.Input<string>;
    /**
     * The name of the security group. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the securityGroup belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The type of the security group. Valid values:
     * `normal`: basic security group.
     * `enterprise`: advanced security group For more information.
     */
    readonly securityGroupType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID.	
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * The security group description. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
     * @deprecated Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
     */
    readonly innerAccess?: pulumi.Input<boolean>;
    /**
     * Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
     */
    readonly innerAccessPolicy?: pulumi.Input<string>;
    /**
     * The name of the security group. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the securityGroup belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The type of the security group. Valid values:
     * `normal`: basic security group.
     * `enterprise`: advanced security group For more information.
     */
    readonly securityGroupType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID.	
     */
    readonly vpcId?: pulumi.Input<string>;
}
