// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a OOS Parameter resource.
 *
 * For information about OOS Parameter and how to use it, see [What is Parameter](https://www.alibabacloud.com/help/en/doc-detail/183408.html).
 *
 * > **NOTE:** Available in v1.147.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.resourcemanager.getResourceGroups({});
 * const example = new alicloud.oos.Parameter("example", {
 *     parameterName: "my-Parameter",
 *     type: "String",
 *     value: "example_value",
 *     description: "example_value",
 *     tags: {
 *         Created: "TF",
 *         For: "OosParameter",
 *     },
 *     resourceGroupId: _default.then(_default => _default.groups?.[0]?.id),
 * });
 * ```
 *
 * ## Import
 *
 * OOS Parameter can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:oos/parameter:Parameter example <parameter_name>
 * ```
 */
export class Parameter extends pulumi.CustomResource {
    /**
     * Get an existing Parameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParameterState, opts?: pulumi.CustomResourceOptions): Parameter {
        return new Parameter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oos/parameter:Parameter';

    /**
     * Returns true if the given object is an instance of Parameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Parameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Parameter.__pulumiType;
    }

    /**
     * The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
     * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
     * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
     * * `MinLength`: The minimum length of the common parameter.
     * * `MaxLength`: The maximum length of the common parameter.
     */
    public readonly constraints!: pulumi.Output<string | undefined>;
    /**
     * The description of the common parameter. The description must be `1` to `200` characters in length.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     */
    public readonly parameterName!: pulumi.Output<string>;
    /**
     * The ID of the Resource Group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The data type of the common parameter. Valid values: `String` and `StringList`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value of the common parameter. The value must be `1` to `4096` characters in length.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Parameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParameterArgs | ParameterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ParameterState | undefined;
            resourceInputs["constraints"] = state ? state.constraints : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["parameterName"] = state ? state.parameterName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ParameterArgs | undefined;
            if ((!args || args.parameterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameterName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["constraints"] = args ? args.constraints : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["parameterName"] = args ? args.parameterName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Parameter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Parameter resources.
 */
export interface ParameterState {
    /**
     * The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
     * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
     * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
     * * `MinLength`: The minimum length of the common parameter.
     * * `MaxLength`: The maximum length of the common parameter.
     */
    constraints?: pulumi.Input<string>;
    /**
     * The description of the common parameter. The description must be `1` to `200` characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     */
    parameterName?: pulumi.Input<string>;
    /**
     * The ID of the Resource Group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The data type of the common parameter. Valid values: `String` and `StringList`.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the common parameter. The value must be `1` to `4096` characters in length.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Parameter resource.
 */
export interface ParameterArgs {
    /**
     * The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
     * * `AllowedValues`: The value that is allowed for the common parameter. It must be an array string.
     * * `AllowedPattern`: The pattern that is allowed for the common parameter. It must be a regular expression.
     * * `MinLength`: The minimum length of the common parameter.
     * * `MaxLength`: The maximum length of the common parameter.
     */
    constraints?: pulumi.Input<string>;
    /**
     * The description of the common parameter. The description must be `1` to `200` characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
     */
    parameterName: pulumi.Input<string>;
    /**
     * The ID of the Resource Group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The data type of the common parameter. Valid values: `String` and `StringList`.
     */
    type: pulumi.Input<string>;
    /**
     * The value of the common parameter. The value must be `1` to `4096` characters in length.
     */
    value: pulumi.Input<string>;
}
