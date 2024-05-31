// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudfirewall.NatFirewallControlPolicyArgs;
import com.pulumi.alicloud.cloudfirewall.inputs.NatFirewallControlPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Firewall Nat Firewall Control Policy resource. Nat firewall access control policy.
 * 
 * For information about Cloud Firewall Nat Firewall Control Policy and how to use it, see [What is Nat Firewall Control Policy](https://www.alibabacloud.com/help/en/cloud-firewall/developer-reference/api-cloudfw-2017-12-07-createnatfirewallcontrolpolicy).
 * 
 * &gt; **NOTE:** Available since v1.224.0.
 * 
 * ## Import
 * 
 * Cloud Firewall Nat Firewall Control Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudfirewall/natFirewallControlPolicy:NatFirewallControlPolicy example &lt;acl_uuid&gt;:&lt;nat_gateway_id&gt;:&lt;direction&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudfirewall/natFirewallControlPolicy:NatFirewallControlPolicy")
public class NatFirewallControlPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The method (action) of access traffic passing through Cloud Firewall in the security access control policy. Valid values:
     * - **accept**: Release
     * - **drop**: Refused
     * - **log**: Observation.
     * 
     */
    @Export(name="aclAction", refs={String.class}, tree="[0]")
    private Output<String> aclAction;

    /**
     * @return The method (action) of access traffic passing through Cloud Firewall in the security access control policy. Valid values:
     * - **accept**: Release
     * - **drop**: Refused
     * - **log**: Observation.
     * 
     */
    public Output<String> aclAction() {
        return this.aclAction;
    }
    /**
     * The unique ID of the security access control policy.
     * &gt; **NOTE:**  To modify a security access control policy, you need to provide the unique ID of the policy. You can call the DescribeNatFirewallControlPolicy interface to obtain the ID.
     * 
     */
    @Export(name="aclUuid", refs={String.class}, tree="[0]")
    private Output<String> aclUuid;

    /**
     * @return The unique ID of the security access control policy.
     * &gt; **NOTE:**  To modify a security access control policy, you need to provide the unique ID of the policy. You can call the DescribeNatFirewallControlPolicy interface to obtain the ID.
     * 
     */
    public Output<String> aclUuid() {
        return this.aclUuid;
    }
    /**
     * The list of application types supported by the access control policy.
     * 
     */
    @Export(name="applicationNameLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> applicationNameLists;

    /**
     * @return The list of application types supported by the access control policy.
     * 
     */
    public Output<List<String>> applicationNameLists() {
        return this.applicationNameLists;
    }
    /**
     * The time when the policy was created.
     * 
     */
    @Export(name="createTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> createTime;

    /**
     * @return The time when the policy was created.
     * 
     */
    public Output<Integer> createTime() {
        return this.createTime;
    }
    /**
     * The description of the access control policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the access control policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The destination port of traffic access in the access control policy. Value:
     * - When the protocol type is set to ICMP, the value of DestPort is null.
     * &gt; **NOTE:**  When the protocol type is ICMP, access control on the destination port is not supported.
     * - When the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) IS group, the value of DestPort is null.
     * &gt; **NOTE:**  When you select group (destination port address book) for the destination port type of the access control policy, you do not need to set a specific destination port number. All ports that need to be controlled by this access control policy are included in the destination port address book.
     * - When the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) is port, the value of DestPort is the destination port number.
     * 
     */
    @Export(name="destPort", refs={String.class}, tree="[0]")
    private Output<String> destPort;

    /**
     * @return The destination port of traffic access in the access control policy. Value:
     * - When the protocol type is set to ICMP, the value of DestPort is null.
     * &gt; **NOTE:**  When the protocol type is ICMP, access control on the destination port is not supported.
     * - When the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) IS group, the value of DestPort is null.
     * &gt; **NOTE:**  When you select group (destination port address book) for the destination port type of the access control policy, you do not need to set a specific destination port number. All ports that need to be controlled by this access control policy are included in the destination port address book.
     * - When the protocol type is TCP, UDP, or ANY, and the destination port type (DestPortType) is port, the value of DestPort is the destination port number.
     * 
     */
    public Output<String> destPort() {
        return this.destPort;
    }
    /**
     * The address book name of the destination port of the access traffic in the access control policy.
     * &gt; **NOTE:**  When DestPortType is set to group, you need to set the destination port address book name.
     * 
     */
    @Export(name="destPortGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destPortGroup;

    /**
     * @return The address book name of the destination port of the access traffic in the access control policy.
     * &gt; **NOTE:**  When DestPortType is set to group, you need to set the destination port address book name.
     * 
     */
    public Output<Optional<String>> destPortGroup() {
        return Codegen.optional(this.destPortGroup);
    }
    /**
     * The destination port type of the access traffic in the security access control policy.
     * - **port**: port
     * - **group**: Port Address Book.
     * 
     */
    @Export(name="destPortType", refs={String.class}, tree="[0]")
    private Output<String> destPortType;

    /**
     * @return The destination port type of the access traffic in the security access control policy.
     * - **port**: port
     * - **group**: Port Address Book.
     * 
     */
    public Output<String> destPortType() {
        return this.destPortType;
    }
    /**
     * The destination address segment in the access control policy. Valid values:
     * - When DestinationType is net, Destination is the Destination CIDR. For example: 1.2.XX.XX/24
     * - When DestinationType IS group, Destination is the name of the Destination address book. For example: db_group
     * - When DestinationType is domain, Destination is the Destination domain name. For example: * .aliyuncs.com
     * - When DestinationType is location, Destination is the Destination region. For example: \[&#34;BJ11&#34;, &#34;ZB&#34;\].
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output<String> destination;

    /**
     * @return The destination address segment in the access control policy. Valid values:
     * - When DestinationType is net, Destination is the Destination CIDR. For example: 1.2.XX.XX/24
     * - When DestinationType IS group, Destination is the name of the Destination address book. For example: db_group
     * - When DestinationType is domain, Destination is the Destination domain name. For example: * .aliyuncs.com
     * - When DestinationType is location, Destination is the Destination region. For example: \[&#34;BJ11&#34;, &#34;ZB&#34;\].
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * The destination address type in the access control policy. Valid values:
     * - **net**: Destination Network segment (CIDR address)
     * - **group**: Destination Address Book
     * - **domain**: the destination domain name.
     * 
     */
    @Export(name="destinationType", refs={String.class}, tree="[0]")
    private Output<String> destinationType;

    /**
     * @return The destination address type in the access control policy. Valid values:
     * - **net**: Destination Network segment (CIDR address)
     * - **group**: Destination Address Book
     * - **domain**: the destination domain name.
     * 
     */
    public Output<String> destinationType() {
        return this.destinationType;
    }
    /**
     * The traffic direction of the access control policy. Valid values:
     * - **out**: Internal and external traffic access control.
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output<String> direction;

    /**
     * @return The traffic direction of the access control policy. Valid values:
     * - **out**: Internal and external traffic access control.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * The domain name resolution method of the access control policy. The policy is enabled by default after it is created. Valid values:
     * - **0**: Based on FQDN
     * - **1**: DNS-based dynamic resolution
     * - **2**: dynamic resolution based on FQDN and DNS.
     * 
     */
    @Export(name="domainResolveType", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> domainResolveType;

    /**
     * @return The domain name resolution method of the access control policy. The policy is enabled by default after it is created. Valid values:
     * - **0**: Based on FQDN
     * - **1**: DNS-based dynamic resolution
     * - **2**: dynamic resolution based on FQDN and DNS.
     * 
     */
    public Output<Optional<Integer>> domainResolveType() {
        return Codegen.optional(this.domainResolveType);
    }
    /**
     * The end time of the policy validity period of the access control policy. Expresses using the second-level timestamp format. Must be full or half time and at least half an hour greater than the start time.
     * &gt; **NOTE:**  When RepeatType is set to permit, EndTime is null. When the RepeatType is None, Daily, Weekly, or Monthly, EndTime must have a value and you need to set the end time.
     * 
     */
    @Export(name="endTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> endTime;

    /**
     * @return The end time of the policy validity period of the access control policy. Expresses using the second-level timestamp format. Must be full or half time and at least half an hour greater than the start time.
     * &gt; **NOTE:**  When RepeatType is set to permit, EndTime is null. When the RepeatType is None, Daily, Weekly, or Monthly, EndTime must have a value and you need to set the end time.
     * 
     */
    public Output<Optional<Integer>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * Supported IP address version. Value:
     * - **4** (default): indicates the IPv4 address.
     * 
     */
    @Export(name="ipVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipVersion;

    /**
     * @return Supported IP address version. Value:
     * - **4** (default): indicates the IPv4 address.
     * 
     */
    public Output<Optional<String>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * The ID of the NAT gateway instance.
     * 
     */
    @Export(name="natGatewayId", refs={String.class}, tree="[0]")
    private Output<String> natGatewayId;

    /**
     * @return The ID of the NAT gateway instance.
     * 
     */
    public Output<String> natGatewayId() {
        return this.natGatewayId;
    }
    /**
     * The priority for the access control policy to take effect. The priority number increases sequentially from 1, and the smaller the priority number, the higher the priority.
     * 
     */
    @Export(name="newOrder", refs={String.class}, tree="[0]")
    private Output<String> newOrder;

    /**
     * @return The priority for the access control policy to take effect. The priority number increases sequentially from 1, and the smaller the priority number, the higher the priority.
     * 
     */
    public Output<String> newOrder() {
        return this.newOrder;
    }
    /**
     * The security protocol type for traffic access in the access control policy. Valid values:
     * - ANY (indicates that all protocol types are queried)
     * - TCP
     * - UDP
     * - ICMP.
     * 
     */
    @Export(name="proto", refs={String.class}, tree="[0]")
    private Output<String> proto;

    /**
     * @return The security protocol type for traffic access in the access control policy. Valid values:
     * - ANY (indicates that all protocol types are queried)
     * - TCP
     * - UDP
     * - ICMP.
     * 
     */
    public Output<String> proto() {
        return this.proto;
    }
    /**
     * The enabled status of the access control policy. The policy is enabled by default after it is created. Value:
     * - **true**: Enable access control policy
     * - **false**: Do not enable access control policies.
     * 
     */
    @Export(name="release", refs={String.class}, tree="[0]")
    private Output<String> release;

    /**
     * @return The enabled status of the access control policy. The policy is enabled by default after it is created. Value:
     * - **true**: Enable access control policy
     * - **false**: Do not enable access control policies.
     * 
     */
    public Output<String> release() {
        return this.release;
    }
    /**
     * Collection of recurring dates for the policy validity period of the access control policy.
     * - When RepeatType is &#39;Permanent&#39;, &#39;None&#39;, &#39;Daily&#39;, RepeatDays is an empty collection. For example:[]
     * - When RepeatType is Weekly, RepeatDays cannot be empty. For example:[&#34;0&#34;, &#34;6&#34;]. When the RepeatType is set to Weekly, RepeatDays cannot be repeated.
     * - RepeatDays cannot be empty when RepeatType is &#39;Monthly. For example:[1, 31]. When RepeatType is set to Monthly, RepeatDays cannot be repeated.
     * 
     */
    @Export(name="repeatDays", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> repeatDays;

    /**
     * @return Collection of recurring dates for the policy validity period of the access control policy.
     * - When RepeatType is &#39;Permanent&#39;, &#39;None&#39;, &#39;Daily&#39;, RepeatDays is an empty collection. For example:[]
     * - When RepeatType is Weekly, RepeatDays cannot be empty. For example:[&#34;0&#34;, &#34;6&#34;]. When the RepeatType is set to Weekly, RepeatDays cannot be repeated.
     * - RepeatDays cannot be empty when RepeatType is &#39;Monthly. For example:[1, 31]. When RepeatType is set to Monthly, RepeatDays cannot be repeated.
     * 
     */
    public Output<Optional<List<Integer>>> repeatDays() {
        return Codegen.optional(this.repeatDays);
    }
    /**
     * The recurring end time of the policy validity period of the access control policy. For example: 23:30, it must be the whole point or half point time, and at least half an hour greater than the repeat start time.
     * &gt; **NOTE:**  When RepeatType is set to normal or None, RepeatEndTime is null. When the RepeatType is Daily, Weekly, or Monthly, the RepeatEndTime must have a value, and you need to set the repeat end time.
     * 
     */
    @Export(name="repeatEndTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repeatEndTime;

    /**
     * @return The recurring end time of the policy validity period of the access control policy. For example: 23:30, it must be the whole point or half point time, and at least half an hour greater than the repeat start time.
     * &gt; **NOTE:**  When RepeatType is set to normal or None, RepeatEndTime is null. When the RepeatType is Daily, Weekly, or Monthly, the RepeatEndTime must have a value, and you need to set the repeat end time.
     * 
     */
    public Output<Optional<String>> repeatEndTime() {
        return Codegen.optional(this.repeatEndTime);
    }
    /**
     * The recurring start time of the policy validity period of the access control policy. For example: 08:00, it must be the whole point or half point time, and at least half an hour less than the repeat end time.
     * &gt; **NOTE:**  When RepeatType is set to permit or None, RepeatStartTime is empty. When the RepeatType is Daily, Weekly, or Monthly, the RepeatStartTime must have a value and you need to set the repeat start time.
     * 
     */
    @Export(name="repeatStartTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repeatStartTime;

    /**
     * @return The recurring start time of the policy validity period of the access control policy. For example: 08:00, it must be the whole point or half point time, and at least half an hour less than the repeat end time.
     * &gt; **NOTE:**  When RepeatType is set to permit or None, RepeatStartTime is empty. When the RepeatType is Daily, Weekly, or Monthly, the RepeatStartTime must have a value and you need to set the repeat start time.
     * 
     */
    public Output<Optional<String>> repeatStartTime() {
        return Codegen.optional(this.repeatStartTime);
    }
    /**
     * The type of repetition for the policy validity period of the access control policy. Value:
     * - **Permit** (default): Always
     * - **None**: Specify a single time
     * - **Daily**: Daily
     * - **Weekly**: Weekly
     * - **Monthly**: Monthly.
     * 
     */
    @Export(name="repeatType", refs={String.class}, tree="[0]")
    private Output<String> repeatType;

    /**
     * @return The type of repetition for the policy validity period of the access control policy. Value:
     * - **Permit** (default): Always
     * - **None**: Specify a single time
     * - **Daily**: Daily
     * - **Weekly**: Weekly
     * - **Monthly**: Monthly.
     * 
     */
    public Output<String> repeatType() {
        return this.repeatType;
    }
    /**
     * The source address in the access control policy. Valid values:
     * - When **SourceType** is set to &#39;net&#39;, Source is the Source CIDR address. For example: 10.2.4.0/24
     * - When **SourceType** is set to &#39;group&#39;, Source is the name of the Source address book. For example: db_group.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return The source address in the access control policy. Valid values:
     * - When **SourceType** is set to &#39;net&#39;, Source is the Source CIDR address. For example: 10.2.4.0/24
     * - When **SourceType** is set to &#39;group&#39;, Source is the name of the Source address book. For example: db_group.
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * The source address type in the access control policy. Valid values:
     * - **net**: the source network segment (CIDR address)
     * - **group**: source address book
     * 
     */
    @Export(name="sourceType", refs={String.class}, tree="[0]")
    private Output<String> sourceType;

    /**
     * @return The source address type in the access control policy. Valid values:
     * - **net**: the source network segment (CIDR address)
     * - **group**: source address book
     * 
     */
    public Output<String> sourceType() {
        return this.sourceType;
    }
    /**
     * The start time of the policy validity period of the access control policy. Expresses using the second-level timestamp format. It must be a full or half hour and at least half an hour less than the end time.
     * &gt; **NOTE:**  When RepeatType is set to normal, StartTime is null. When the RepeatType is None, Daily, Weekly, or Monthly, StartTime must have a value and you need to set the start time.
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startTime;

    /**
     * @return The start time of the policy validity period of the access control policy. Expresses using the second-level timestamp format. It must be a full or half hour and at least half an hour less than the end time.
     * &gt; **NOTE:**  When RepeatType is set to normal, StartTime is null. When the RepeatType is None, Daily, Weekly, or Monthly, StartTime must have a value and you need to set the start time.
     * 
     */
    public Output<Optional<Integer>> startTime() {
        return Codegen.optional(this.startTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NatFirewallControlPolicy(String name) {
        this(name, NatFirewallControlPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NatFirewallControlPolicy(String name, NatFirewallControlPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NatFirewallControlPolicy(String name, NatFirewallControlPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/natFirewallControlPolicy:NatFirewallControlPolicy", name, args == null ? NatFirewallControlPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NatFirewallControlPolicy(String name, Output<String> id, @Nullable NatFirewallControlPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/natFirewallControlPolicy:NatFirewallControlPolicy", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static NatFirewallControlPolicy get(String name, Output<String> id, @Nullable NatFirewallControlPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NatFirewallControlPolicy(name, id, state, options);
    }
}
