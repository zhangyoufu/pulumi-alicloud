// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud Serverless App Engine (SAE) Application Load Balancer Attachment resource.
 *
 * For information about Serverless App Engine (SAE) Load Balancer Internet Attachment and how to use it, see [alicloud.sae.LoadBalancerInternet](https://www.alibabacloud.com/help/en/sae/latest/bindslb).
 *
 * > **NOTE:** Available since v1.164.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultRandomInteger = new random.RandomInteger("defaultRandomInteger", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultNamespace = new alicloud.sae.Namespace("defaultNamespace", {
 *     namespaceId: pulumi.all([defaultRegions, defaultRandomInteger.result]).apply(([defaultRegions, result]) => `${defaultRegions.regions?.[0]?.id}:example${result}`),
 *     namespaceName: name,
 *     namespaceDescription: name,
 *     enableMicroRegistration: false,
 * });
 * const defaultApplication = new alicloud.sae.Application("defaultApplication", {
 *     appDescription: name,
 *     appName: name,
 *     namespaceId: defaultNamespace.id,
 *     imageUrl: "registry-vpc.cn-hangzhou.aliyuncs.com/lxepoo/apache-php5",
 *     packageType: "Image",
 *     jdk: "Open JDK 8",
 *     securityGroupId: defaultSecurityGroup.id,
 *     vpcId: defaultNetwork.id,
 *     vswitchId: defaultSwitch.id,
 *     timezone: "Asia/Beijing",
 *     replicas: 5,
 *     cpu: 500,
 *     memory: 2048,
 * });
 * const defaultApplicationLoadBalancer = new alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", {
 *     loadBalancerName: name,
 *     vswitchId: defaultSwitch.id,
 *     loadBalancerSpec: "slb.s2.small",
 *     addressType: "internet",
 * });
 * const defaultLoadBalancerInternet = new alicloud.sae.LoadBalancerInternet("defaultLoadBalancerInternet", {
 *     appId: defaultApplication.id,
 *     internetSlbId: defaultApplicationLoadBalancer.id,
 *     internets: [{
 *         protocol: "TCP",
 *         port: 80,
 *         targetPort: 8080,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * The resource can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:sae/loadBalancerInternet:LoadBalancerInternet example <id>
 * ```
 */
export class LoadBalancerInternet extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerInternet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerInternetState, opts?: pulumi.CustomResourceOptions): LoadBalancerInternet {
        return new LoadBalancerInternet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:sae/loadBalancerInternet:LoadBalancerInternet';

    /**
     * Returns true if the given object is an instance of LoadBalancerInternet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerInternet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerInternet.__pulumiType;
    }

    /**
     * The target application ID that needs to be bound to the SLB.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Use designated public network SLBs that have been purchased to support non-shared instances.
     */
    public /*out*/ readonly internetIp!: pulumi.Output<string>;
    /**
     * The internet SLB ID.
     */
    public readonly internetSlbId!: pulumi.Output<string | undefined>;
    /**
     * The bound private network SLB. See `internet` below.
     */
    public readonly internets!: pulumi.Output<outputs.sae.LoadBalancerInternetInternet[]>;

    /**
     * Create a LoadBalancerInternet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerInternetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerInternetArgs | LoadBalancerInternetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerInternetState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["internetIp"] = state ? state.internetIp : undefined;
            resourceInputs["internetSlbId"] = state ? state.internetSlbId : undefined;
            resourceInputs["internets"] = state ? state.internets : undefined;
        } else {
            const args = argsOrState as LoadBalancerInternetArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.internets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internets'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["internetSlbId"] = args ? args.internetSlbId : undefined;
            resourceInputs["internets"] = args ? args.internets : undefined;
            resourceInputs["internetIp"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancerInternet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerInternet resources.
 */
export interface LoadBalancerInternetState {
    /**
     * The target application ID that needs to be bound to the SLB.
     */
    appId?: pulumi.Input<string>;
    /**
     * Use designated public network SLBs that have been purchased to support non-shared instances.
     */
    internetIp?: pulumi.Input<string>;
    /**
     * The internet SLB ID.
     */
    internetSlbId?: pulumi.Input<string>;
    /**
     * The bound private network SLB. See `internet` below.
     */
    internets?: pulumi.Input<pulumi.Input<inputs.sae.LoadBalancerInternetInternet>[]>;
}

/**
 * The set of arguments for constructing a LoadBalancerInternet resource.
 */
export interface LoadBalancerInternetArgs {
    /**
     * The target application ID that needs to be bound to the SLB.
     */
    appId: pulumi.Input<string>;
    /**
     * The internet SLB ID.
     */
    internetSlbId?: pulumi.Input<string>;
    /**
     * The bound private network SLB. See `internet` below.
     */
    internets: pulumi.Input<pulumi.Input<inputs.sae.LoadBalancerInternetInternet>[]>;
}
