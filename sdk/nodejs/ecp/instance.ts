// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Elastic Cloud Phone (ECP) Instance resource.
 *
 * For information about Elastic Cloud Phone (ECP) Instance and how to use it,
 * see [What is Instance](https://www.alibabacloud.com/help/en/cloudphone/latest/api-cloudphone-2020-12-30-runinstances).
 *
 * > **NOTE:** Available since v1.158.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultInteger = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const default = alicloud.ecp.getZones({});
 * const defaultGetInstanceTypes = alicloud.ecp.getInstanceTypes({});
 * const countSize = _default.then(_default => _default.zones).length;
 * const zoneId = Promise.all([_default, countSize]).then(([_default, countSize]) => _default.zones[countSize - 1].zoneId);
 * const instanceTypeCountSize = defaultGetInstanceTypes.then(defaultGetInstanceTypes => defaultGetInstanceTypes.instanceTypes).length;
 * const instanceType = Promise.all([defaultGetInstanceTypes, instanceTypeCountSize]).then(([defaultGetInstanceTypes, instanceTypeCountSize]) => defaultGetInstanceTypes.instanceTypes[instanceTypeCountSize - 1].instanceType);
 * const defaultGetNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 * });
 * const defaultGetSwitches = defaultGetNetworks.then(defaultGetNetworks => alicloud.vpc.getSwitches({
 *     vpcId: defaultGetNetworks.ids?.[0],
 *     zoneId: zoneId,
 * }));
 * const group = new alicloud.ecs.SecurityGroup("group", {
 *     name: name,
 *     vpcId: defaultGetNetworks.then(defaultGetNetworks => defaultGetNetworks.ids?.[0]),
 * });
 * const defaultKeyPair = new alicloud.ecp.KeyPair("default", {
 *     keyPairName: `${name}-${defaultInteger.result}`,
 *     publicKeyBody: "ssh-rsa AAAAB3Nza12345678qwertyuudsfsg",
 * });
 * const defaultInstance = new alicloud.ecp.Instance("default", {
 *     instanceName: name,
 *     description: name,
 *     keyPairName: defaultKeyPair.keyPairName,
 *     securityGroupId: group.id,
 *     vswitchId: defaultGetSwitches.then(defaultGetSwitches => defaultGetSwitches.ids?.[0]),
 *     imageId: "android_9_0_0_release_2851157_20211201.vhd",
 *     instanceType: defaultGetInstanceTypes.then(defaultGetInstanceTypes => defaultGetInstanceTypes.instanceTypes?.[1]?.instanceType),
 *     vncPassword: "Ecp123",
 *     paymentType: "PayAsYouGo",
 *     force: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Elastic Cloud Phone (ECP) Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ecp/instance:Instance example <id>
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecp/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The auto pay.
     */
    public readonly autoPay!: pulumi.Output<boolean | undefined>;
    /**
     * The auto renew.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * Description of the instance. 2 to 256 English or Chinese characters in length and cannot
     * start with `http://` and `https`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The eip bandwidth.
     */
    public readonly eipBandwidth!: pulumi.Output<number | undefined>;
    /**
     * The force.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * The ID Of The Image.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The name of the instance. It must be 2 to 128 characters in length and must start with an
     * uppercase letter or Chinese. It cannot start with http:// or https. It can contain Chinese, English, numbers,
     * half-width colons (:), underscores (_), half-width periods (.), or dashes (-). The default value is the InstanceId of
     * the instance.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * Instance Type.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The name of the key pair of the mobile phone instance.
     */
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    /**
     * The payment type.Valid values: `PayAsYouGo`,`Subscription`
     */
    public readonly paymentType!: pulumi.Output<string | undefined>;
    /**
     * The period. It is valid when `periodUnit` is 'Year'. Valid value: `1`, `2`, `3`, `4`, `5`. It
     * is valid when `periodUnit` is 'Month'. Valid value: `1`, `2`, `3`, `5`
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * The duration unit that you will buy the resource. Valid value: `Year`,`Month`. Default
     * to `Month`.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The selected resolution for the cloud mobile phone instance.
     */
    public readonly resolution!: pulumi.Output<string>;
    /**
     * The ID of the security group. The security group is the same as that of the
     * ECS instance.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Instance status. Valid values: `Running`, `Stopped`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Cloud mobile phone VNC password. The password must be six characters in length and must
     * contain only uppercase, lowercase English letters and Arabic numerals.
     */
    public readonly vncPassword!: pulumi.Output<string | undefined>;
    /**
     * The vswitch id.
     */
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["autoPay"] = state ? state.autoPay : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["eipBandwidth"] = state ? state.eipBandwidth : undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["resolution"] = state ? state.resolution : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vncPassword"] = state ? state.vncPassword : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["autoPay"] = args ? args.autoPay : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eipBandwidth"] = args ? args.eipBandwidth : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["resolution"] = args ? args.resolution : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vncPassword"] = args?.vncPassword ? pulumi.secret(args.vncPassword) : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["vncPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The auto pay.
     */
    autoPay?: pulumi.Input<boolean>;
    /**
     * The auto renew.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * Description of the instance. 2 to 256 English or Chinese characters in length and cannot
     * start with `http://` and `https`.
     */
    description?: pulumi.Input<string>;
    /**
     * The eip bandwidth.
     */
    eipBandwidth?: pulumi.Input<number>;
    /**
     * The force.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The ID Of The Image.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The name of the instance. It must be 2 to 128 characters in length and must start with an
     * uppercase letter or Chinese. It cannot start with http:// or https. It can contain Chinese, English, numbers,
     * half-width colons (:), underscores (_), half-width periods (.), or dashes (-). The default value is the InstanceId of
     * the instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance Type.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The name of the key pair of the mobile phone instance.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The payment type.Valid values: `PayAsYouGo`,`Subscription`
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The period. It is valid when `periodUnit` is 'Year'. Valid value: `1`, `2`, `3`, `4`, `5`. It
     * is valid when `periodUnit` is 'Month'. Valid value: `1`, `2`, `3`, `5`
     */
    period?: pulumi.Input<string>;
    /**
     * The duration unit that you will buy the resource. Valid value: `Year`,`Month`. Default
     * to `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The selected resolution for the cloud mobile phone instance.
     */
    resolution?: pulumi.Input<string>;
    /**
     * The ID of the security group. The security group is the same as that of the
     * ECS instance.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * Instance status. Valid values: `Running`, `Stopped`.
     */
    status?: pulumi.Input<string>;
    /**
     * Cloud mobile phone VNC password. The password must be six characters in length and must
     * contain only uppercase, lowercase English letters and Arabic numerals.
     */
    vncPassword?: pulumi.Input<string>;
    /**
     * The vswitch id.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The auto pay.
     */
    autoPay?: pulumi.Input<boolean>;
    /**
     * The auto renew.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * Description of the instance. 2 to 256 English or Chinese characters in length and cannot
     * start with `http://` and `https`.
     */
    description?: pulumi.Input<string>;
    /**
     * The eip bandwidth.
     */
    eipBandwidth?: pulumi.Input<number>;
    /**
     * The force.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The ID Of The Image.
     */
    imageId: pulumi.Input<string>;
    /**
     * The name of the instance. It must be 2 to 128 characters in length and must start with an
     * uppercase letter or Chinese. It cannot start with http:// or https. It can contain Chinese, English, numbers,
     * half-width colons (:), underscores (_), half-width periods (.), or dashes (-). The default value is the InstanceId of
     * the instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance Type.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The name of the key pair of the mobile phone instance.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The payment type.Valid values: `PayAsYouGo`,`Subscription`
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The period. It is valid when `periodUnit` is 'Year'. Valid value: `1`, `2`, `3`, `4`, `5`. It
     * is valid when `periodUnit` is 'Month'. Valid value: `1`, `2`, `3`, `5`
     */
    period?: pulumi.Input<string>;
    /**
     * The duration unit that you will buy the resource. Valid value: `Year`,`Month`. Default
     * to `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The selected resolution for the cloud mobile phone instance.
     */
    resolution?: pulumi.Input<string>;
    /**
     * The ID of the security group. The security group is the same as that of the
     * ECS instance.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * Instance status. Valid values: `Running`, `Stopped`.
     */
    status?: pulumi.Input<string>;
    /**
     * Cloud mobile phone VNC password. The password must be six characters in length and must
     * contain only uppercase, lowercase English letters and Arabic numerals.
     */
    vncPassword?: pulumi.Input<string>;
    /**
     * The vswitch id.
     */
    vswitchId: pulumi.Input<string>;
}
