// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Eipanycast Anycast Eip Address Attachment resource.
 *
 * For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://www.alibabacloud.com/help/en/anycast-eip/latest/api-eipanycast-2020-03-09-associateanycasteipaddress).
 *
 * > **NOTE:** Available since v1.113.0.
 *
 * > **NOTE:** The following regions support currently while Slb instance support bound.
 * [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultZones = alicloud.slb.getZones({
 *     availableSlbAddressType: "vpc",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.0.0.0/8",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.1.0.0/16",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", {
 *     addressType: "intranet",
 *     vswitchId: defaultSwitch.id,
 *     loadBalancerName: name,
 *     loadBalancerSpec: "slb.s1.small",
 *     masterZoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultAnycastEipAddress = new alicloud.eipanycast.AnycastEipAddress("defaultAnycastEipAddress", {
 *     anycastEipAddressName: name,
 *     serviceLocation: "ChineseMainland",
 * });
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultAnycastEipAddressAttachment = new alicloud.eipanycast.AnycastEipAddressAttachment("defaultAnycastEipAddressAttachment", {
 *     bindInstanceId: defaultApplicationLoadBalancer.id,
 *     bindInstanceType: "SlbInstance",
 *     bindInstanceRegionId: defaultRegions.then(defaultRegions => defaultRegions.regions?.[0]?.id),
 *     anycastId: defaultAnycastEipAddress.id,
 * });
 * ```
 *
 * Multiple Usage
 *
 * > **NOTE:**  Anycast EIP supports binding cloud resource instances in multiple regions. Only one cloud resource instance is supported as the default origin station, and the rest are normal origin stations. When no access point is specified or an access point is added, the access request is forwarded to the default origin by default.  If you are bound for the first time, the Default value of the binding mode is **Default * *. /li> li> If you are not binding for the first time, you can set the binding mode to **Default**, and the new Default origin will take effect. The original Default origin will be changed to a common origin.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const beijing = new alicloud.Provider("beijing", {region: "cn-beijing"});
 * const hangzhou = new alicloud.Provider("hangzhou", {region: "cn-hangzhou"});
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultVpc = new alicloud.vpc.Network("defaultVpc", {
 *     vpcName: name,
 *     cidrBlock: "192.168.0.0/16",
 * }, {
 *     provider: "alicloud.beijing",
 * });
 * const defaultVsw = new alicloud.vpc.Switch("defaultVsw", {
 *     vpcId: defaultVpc.id,
 *     cidrBlock: "192.168.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * }, {
 *     provider: "alicloud.beijing",
 * });
 * const defaultuBsECI = new alicloud.ecs.SecurityGroup("defaultuBsECI", {vpcId: defaultVpc.id}, {
 *     provider: "alicloud.beijing",
 * });
 * const default9KDlN7 = new alicloud.ecs.Instance("default9KDlN7", {
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     instanceName: name,
 *     securityGroups: [defaultuBsECI.id],
 *     availabilityZone: defaultVsw.zoneId,
 *     instanceChargeType: "PostPaid",
 *     systemDiskCategory: "cloud_efficiency",
 *     vswitchId: defaultVsw.id,
 * }, {
 *     provider: "alicloud.beijing",
 * });
 * const defaultXkpFRs = new alicloud.eipanycast.AnycastEipAddress("defaultXkpFRs", {serviceLocation: "ChineseMainland"}, {
 *     provider: "alicloud.hangzhou",
 * });
 * const defaultVpc2 = new alicloud.vpc.Network("defaultVpc2", {
 *     vpcName: `${name}6`,
 *     cidrBlock: "192.168.0.0/16",
 * }, {
 *     provider: "alicloud.hangzhou",
 * });
 * const default2Zones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const default2Images = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const default2InstanceTypes = default2Zones.then(default2Zones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: default2Zones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultdsVsw2 = new alicloud.vpc.Switch("defaultdsVsw2", {
 *     vpcId: defaultVpc2.id,
 *     cidrBlock: "192.168.0.0/24",
 *     zoneId: default2Zones.then(default2Zones => default2Zones.zones?.[1]?.id),
 * }, {
 *     provider: "alicloud.hangzhou",
 * });
 * const defaultuBsECI2 = new alicloud.ecs.SecurityGroup("defaultuBsECI2", {vpcId: defaultVpc2.id}, {
 *     provider: "alicloud.hangzhou",
 * });
 * const defaultEcs2 = new alicloud.ecs.Instance("defaultEcs2", {
 *     imageId: default2Images.then(default2Images => default2Images.images?.[0]?.id),
 *     instanceType: default2InstanceTypes.then(default2InstanceTypes => default2InstanceTypes.instanceTypes?.[0]?.id),
 *     instanceName: name,
 *     securityGroups: [defaultuBsECI2.id],
 *     availabilityZone: defaultdsVsw2.zoneId,
 *     instanceChargeType: "PostPaid",
 *     systemDiskCategory: "cloud_efficiency",
 *     vswitchId: defaultdsVsw2.id,
 * }, {
 *     provider: "alicloud.hangzhou",
 * });
 * const defaultEfYBJY = new alicloud.eipanycast.AnycastEipAddressAttachment("defaultEfYBJY", {
 *     bindInstanceId: default9KDlN7.networkInterfaceId,
 *     bindInstanceType: "NetworkInterface",
 *     bindInstanceRegionId: "cn-beijing",
 *     anycastId: defaultXkpFRs.id,
 *     associationMode: "Default",
 * }, {
 *     provider: "alicloud.beijing",
 * });
 * const normal = new alicloud.eipanycast.AnycastEipAddressAttachment("normal", {
 *     bindInstanceId: defaultEcs2.networkInterfaceId,
 *     bindInstanceType: "NetworkInterface",
 *     bindInstanceRegionId: "cn-hangzhou",
 *     anycastId: defaultEfYBJY.anycastId,
 * }, {
 *     provider: "alicloud.hangzhou",
 * });
 * ```
 *
 * ## Import
 *
 * Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example <anycast_id>:<bind_instance_id>:<bind_instance_region_id>:<bind_instance_type>
 * ```
 */
export class AnycastEipAddressAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AnycastEipAddressAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnycastEipAddressAttachmentState, opts?: pulumi.CustomResourceOptions): AnycastEipAddressAttachment {
        return new AnycastEipAddressAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment';

    /**
     * Returns true if the given object is an instance of AnycastEipAddressAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnycastEipAddressAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnycastEipAddressAttachment.__pulumiType;
    }

    /**
     * The ID of the Anycast EIP instance.
     */
    public readonly anycastId!: pulumi.Output<string>;
    /**
     * Binding mode, value:
     * - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
     * - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
     */
    public readonly associationMode!: pulumi.Output<string>;
    /**
     * The ID of the cloud resource instance to be bound.
     */
    public readonly bindInstanceId!: pulumi.Output<string>;
    /**
     * The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
     */
    public readonly bindInstanceRegionId!: pulumi.Output<string>;
    /**
     * The type of the cloud resource instance to be bound. Value:
     * - **SlbInstance**: a private network SLB instance.
     * - **NetworkInterface**: ENI.
     */
    public readonly bindInstanceType!: pulumi.Output<string>;
    /**
     * Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
     */
    public /*out*/ readonly bindTime!: pulumi.Output<string>;
    /**
     * The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
     */
    public readonly popLocations!: pulumi.Output<outputs.eipanycast.AnycastEipAddressAttachmentPopLocation[] | undefined>;
    /**
     * The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
     */
    public readonly privateIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AnycastEipAddressAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnycastEipAddressAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnycastEipAddressAttachmentArgs | AnycastEipAddressAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnycastEipAddressAttachmentState | undefined;
            resourceInputs["anycastId"] = state ? state.anycastId : undefined;
            resourceInputs["associationMode"] = state ? state.associationMode : undefined;
            resourceInputs["bindInstanceId"] = state ? state.bindInstanceId : undefined;
            resourceInputs["bindInstanceRegionId"] = state ? state.bindInstanceRegionId : undefined;
            resourceInputs["bindInstanceType"] = state ? state.bindInstanceType : undefined;
            resourceInputs["bindTime"] = state ? state.bindTime : undefined;
            resourceInputs["popLocations"] = state ? state.popLocations : undefined;
            resourceInputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AnycastEipAddressAttachmentArgs | undefined;
            if ((!args || args.anycastId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'anycastId'");
            }
            if ((!args || args.bindInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindInstanceId'");
            }
            if ((!args || args.bindInstanceRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindInstanceRegionId'");
            }
            if ((!args || args.bindInstanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindInstanceType'");
            }
            resourceInputs["anycastId"] = args ? args.anycastId : undefined;
            resourceInputs["associationMode"] = args ? args.associationMode : undefined;
            resourceInputs["bindInstanceId"] = args ? args.bindInstanceId : undefined;
            resourceInputs["bindInstanceRegionId"] = args ? args.bindInstanceRegionId : undefined;
            resourceInputs["bindInstanceType"] = args ? args.bindInstanceType : undefined;
            resourceInputs["popLocations"] = args ? args.popLocations : undefined;
            resourceInputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
            resourceInputs["bindTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnycastEipAddressAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnycastEipAddressAttachment resources.
 */
export interface AnycastEipAddressAttachmentState {
    /**
     * The ID of the Anycast EIP instance.
     */
    anycastId?: pulumi.Input<string>;
    /**
     * Binding mode, value:
     * - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
     * - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
     */
    associationMode?: pulumi.Input<string>;
    /**
     * The ID of the cloud resource instance to be bound.
     */
    bindInstanceId?: pulumi.Input<string>;
    /**
     * The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
     */
    bindInstanceRegionId?: pulumi.Input<string>;
    /**
     * The type of the cloud resource instance to be bound. Value:
     * - **SlbInstance**: a private network SLB instance.
     * - **NetworkInterface**: ENI.
     */
    bindInstanceType?: pulumi.Input<string>;
    /**
     * Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
     */
    bindTime?: pulumi.Input<string>;
    /**
     * The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
     */
    popLocations?: pulumi.Input<pulumi.Input<inputs.eipanycast.AnycastEipAddressAttachmentPopLocation>[]>;
    /**
     * The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnycastEipAddressAttachment resource.
 */
export interface AnycastEipAddressAttachmentArgs {
    /**
     * The ID of the Anycast EIP instance.
     */
    anycastId: pulumi.Input<string>;
    /**
     * Binding mode, value:
     * - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
     * - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
     */
    associationMode?: pulumi.Input<string>;
    /**
     * The ID of the cloud resource instance to be bound.
     */
    bindInstanceId: pulumi.Input<string>;
    /**
     * The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
     */
    bindInstanceRegionId: pulumi.Input<string>;
    /**
     * The type of the cloud resource instance to be bound. Value:
     * - **SlbInstance**: a private network SLB instance.
     * - **NetworkInterface**: ENI.
     */
    bindInstanceType: pulumi.Input<string>;
    /**
     * The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
     */
    popLocations?: pulumi.Input<pulumi.Input<inputs.eipanycast.AnycastEipAddressAttachmentPopLocation>[]>;
    /**
     * The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
     */
    privateIpAddress?: pulumi.Input<string>;
}
