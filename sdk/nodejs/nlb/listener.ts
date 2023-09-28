// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a NLB Listener resource.
 *
 * For information about NLB Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createlistener-nl).
 *
 * > **NOTE:** Available since v1.191.0.
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
 * const name = config.get("name") || "tf-example";
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({});
 * const defaultZones = alicloud.nlb.getZones({});
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
 * const default1 = new alicloud.vpc.Switch("default1", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.1.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[1]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultLoadBalancer = new alicloud.nlb.LoadBalancer("defaultLoadBalancer", {
 *     loadBalancerName: name,
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     loadBalancerType: "Network",
 *     addressType: "Internet",
 *     addressIpVersion: "Ipv4",
 *     vpcId: defaultNetwork.id,
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 *     zoneMappings: [
 *         {
 *             vswitchId: defaultSwitch.id,
 *             zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *         },
 *         {
 *             vswitchId: default1.id,
 *             zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[1]?.id),
 *         },
 *     ],
 * });
 * const defaultServerGroup = new alicloud.nlb.ServerGroup("defaultServerGroup", {
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     serverGroupName: name,
 *     serverGroupType: "Instance",
 *     vpcId: defaultNetwork.id,
 *     scheduler: "Wrr",
 *     protocol: "TCP",
 *     connectionDrain: true,
 *     connectionDrainTimeout: 60,
 *     addressIpVersion: "Ipv4",
 *     healthCheck: {
 *         healthCheckEnabled: true,
 *         healthCheckType: "TCP",
 *         healthCheckConnectPort: 0,
 *         healthyThreshold: 2,
 *         unhealthyThreshold: 2,
 *         healthCheckConnectTimeout: 5,
 *         healthCheckInterval: 10,
 *         httpCheckMethod: "GET",
 *         healthCheckHttpCodes: [
 *             "http_2xx",
 *             "http_3xx",
 *             "http_4xx",
 *         ],
 *     },
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * const defaultListener = new alicloud.nlb.Listener("defaultListener", {
 *     listenerProtocol: "TCP",
 *     listenerPort: 80,
 *     listenerDescription: name,
 *     loadBalancerId: defaultLoadBalancer.id,
 *     serverGroupId: defaultServerGroup.id,
 *     idleTimeout: 900,
 *     proxyProtocolEnabled: true,
 *     secSensorEnabled: true,
 *     cps: 10000,
 *     mss: 0,
 * });
 * ```
 *
 * ## Import
 *
 * NLB Listener can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:nlb/listener:Listener example <id>
 * ```
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:nlb/listener:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     */
    public readonly alpnEnabled!: pulumi.Output<boolean>;
    /**
     * The ALPN policy.
     */
    public readonly alpnPolicy!: pulumi.Output<string | undefined>;
    /**
     * The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     */
    public readonly caCertificateIds!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies whether to enable mutual authentication.
     */
    public readonly caEnabled!: pulumi.Output<boolean>;
    /**
     * The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     */
    public readonly certificateIds!: pulumi.Output<string[] | undefined>;
    /**
     * The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     */
    public readonly cps!: pulumi.Output<number | undefined>;
    /**
     * Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     */
    public readonly endPort!: pulumi.Output<number | undefined>;
    /**
     * The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     */
    public readonly idleTimeout!: pulumi.Output<number>;
    /**
     * Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     */
    public readonly listenerDescription!: pulumi.Output<string | undefined>;
    /**
     * Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     */
    public readonly listenerPort!: pulumi.Output<number>;
    /**
     * The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     */
    public readonly listenerProtocol!: pulumi.Output<string>;
    /**
     * The ID of the network-based server load balancer instance.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     */
    public readonly mss!: pulumi.Output<number | undefined>;
    /**
     * Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     */
    public readonly proxyProtocolEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies whether to enable fine-grained monitoring.
     */
    public readonly secSensorEnabled!: pulumi.Output<boolean>;
    /**
     * The ID of the security policy. System security policies and custom security policies are supported. 
     * System security policies valid values: `tlsCipherPolicy10` (default), `tls_cipher_policy_1_1,` `tlsCipherPolicy12`, `tlsCipherPolicy12Strict`, and `tlsCipherPolicy12StrictWith13`.
     * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * The ID of the server group.
     */
    public readonly serverGroupId!: pulumi.Output<string>;
    /**
     * Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     */
    public readonly startPort!: pulumi.Output<number | undefined>;
    /**
     * The status of the resource. Valid values: `Running`, `Stopped`.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerState | undefined;
            resourceInputs["alpnEnabled"] = state ? state.alpnEnabled : undefined;
            resourceInputs["alpnPolicy"] = state ? state.alpnPolicy : undefined;
            resourceInputs["caCertificateIds"] = state ? state.caCertificateIds : undefined;
            resourceInputs["caEnabled"] = state ? state.caEnabled : undefined;
            resourceInputs["certificateIds"] = state ? state.certificateIds : undefined;
            resourceInputs["cps"] = state ? state.cps : undefined;
            resourceInputs["endPort"] = state ? state.endPort : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["listenerDescription"] = state ? state.listenerDescription : undefined;
            resourceInputs["listenerPort"] = state ? state.listenerPort : undefined;
            resourceInputs["listenerProtocol"] = state ? state.listenerProtocol : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["mss"] = state ? state.mss : undefined;
            resourceInputs["proxyProtocolEnabled"] = state ? state.proxyProtocolEnabled : undefined;
            resourceInputs["secSensorEnabled"] = state ? state.secSensorEnabled : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            resourceInputs["serverGroupId"] = state ? state.serverGroupId : undefined;
            resourceInputs["startPort"] = state ? state.startPort : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if ((!args || args.listenerPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerPort'");
            }
            if ((!args || args.listenerProtocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerProtocol'");
            }
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            if ((!args || args.serverGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverGroupId'");
            }
            resourceInputs["alpnEnabled"] = args ? args.alpnEnabled : undefined;
            resourceInputs["alpnPolicy"] = args ? args.alpnPolicy : undefined;
            resourceInputs["caCertificateIds"] = args ? args.caCertificateIds : undefined;
            resourceInputs["caEnabled"] = args ? args.caEnabled : undefined;
            resourceInputs["certificateIds"] = args ? args.certificateIds : undefined;
            resourceInputs["cps"] = args ? args.cps : undefined;
            resourceInputs["endPort"] = args ? args.endPort : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["listenerDescription"] = args ? args.listenerDescription : undefined;
            resourceInputs["listenerPort"] = args ? args.listenerPort : undefined;
            resourceInputs["listenerProtocol"] = args ? args.listenerProtocol : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["mss"] = args ? args.mss : undefined;
            resourceInputs["proxyProtocolEnabled"] = args ? args.proxyProtocolEnabled : undefined;
            resourceInputs["secSensorEnabled"] = args ? args.secSensorEnabled : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            resourceInputs["serverGroupId"] = args ? args.serverGroupId : undefined;
            resourceInputs["startPort"] = args ? args.startPort : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Listener.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     */
    alpnEnabled?: pulumi.Input<boolean>;
    /**
     * The ALPN policy.
     */
    alpnPolicy?: pulumi.Input<string>;
    /**
     * The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     */
    caCertificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to enable mutual authentication.
     */
    caEnabled?: pulumi.Input<boolean>;
    /**
     * The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     */
    cps?: pulumi.Input<number>;
    /**
     * Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     */
    endPort?: pulumi.Input<number>;
    /**
     * The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     */
    listenerDescription?: pulumi.Input<string>;
    /**
     * Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     */
    listenerPort?: pulumi.Input<number>;
    /**
     * The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     */
    listenerProtocol?: pulumi.Input<string>;
    /**
     * The ID of the network-based server load balancer instance.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     */
    mss?: pulumi.Input<number>;
    /**
     * Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     */
    proxyProtocolEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable fine-grained monitoring.
     */
    secSensorEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the security policy. System security policies and custom security policies are supported. 
     * System security policies valid values: `tlsCipherPolicy10` (default), `tls_cipher_policy_1_1,` `tlsCipherPolicy12`, `tlsCipherPolicy12Strict`, and `tlsCipherPolicy12StrictWith13`.
     * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * The ID of the server group.
     */
    serverGroupId?: pulumi.Input<string>;
    /**
     * Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     */
    startPort?: pulumi.Input<number>;
    /**
     * The status of the resource. Valid values: `Running`, `Stopped`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     */
    alpnEnabled?: pulumi.Input<boolean>;
    /**
     * The ALPN policy.
     */
    alpnPolicy?: pulumi.Input<string>;
    /**
     * The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     */
    caCertificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to enable mutual authentication.
     */
    caEnabled?: pulumi.Input<boolean>;
    /**
     * The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     */
    cps?: pulumi.Input<number>;
    /**
     * Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     */
    endPort?: pulumi.Input<number>;
    /**
     * The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     */
    listenerDescription?: pulumi.Input<string>;
    /**
     * Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     */
    listenerPort: pulumi.Input<number>;
    /**
     * The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     */
    listenerProtocol: pulumi.Input<string>;
    /**
     * The ID of the network-based server load balancer instance.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     */
    mss?: pulumi.Input<number>;
    /**
     * Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     */
    proxyProtocolEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable fine-grained monitoring.
     */
    secSensorEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the security policy. System security policies and custom security policies are supported. 
     * System security policies valid values: `tlsCipherPolicy10` (default), `tls_cipher_policy_1_1,` `tlsCipherPolicy12`, `tlsCipherPolicy12Strict`, and `tlsCipherPolicy12StrictWith13`.
     * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * The ID of the server group.
     */
    serverGroupId: pulumi.Input<string>;
    /**
     * Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     */
    startPort?: pulumi.Input<number>;
    /**
     * The status of the resource. Valid values: `Running`, `Stopped`.
     */
    status?: pulumi.Input<string>;
}
