// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.AddressPoolArgs;
import com.pulumi.alicloud.dns.inputs.AddressPoolState;
import com.pulumi.alicloud.dns.outputs.AddressPoolAddress;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Alidns Address Pool resource.
 * 
 * For information about Alidns Address Pool and how to use it, see [What is Address Pool](https://www.alibabacloud.com/help/doc-detail/189621.html).
 * 
 * &gt; **NOTE:** Available since v1.152.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.cms.AlarmContactGroup;
 * import com.pulumi.alicloud.cms.AlarmContactGroupArgs;
 * import com.pulumi.alicloud.dns.GtmInstance;
 * import com.pulumi.alicloud.dns.GtmInstanceArgs;
 * import com.pulumi.alicloud.dns.inputs.GtmInstanceAlertConfigArgs;
 * import com.pulumi.alicloud.dns.AddressPool;
 * import com.pulumi.alicloud.dns.AddressPoolArgs;
 * import com.pulumi.alicloud.dns.inputs.AddressPoolAddressArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var domainName = config.get(&#34;domainName&#34;).orElse(&#34;alicloud-provider.com&#34;);
 *         final var default = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultAlarmContactGroup = new AlarmContactGroup(&#34;defaultAlarmContactGroup&#34;, AlarmContactGroupArgs.builder()        
 *             .alarmContactGroupName(name)
 *             .build());
 * 
 *         var defaultGtmInstance = new GtmInstance(&#34;defaultGtmInstance&#34;, GtmInstanceArgs.builder()        
 *             .instanceName(name)
 *             .paymentType(&#34;Subscription&#34;)
 *             .period(1)
 *             .renewalStatus(&#34;ManualRenewal&#34;)
 *             .packageEdition(&#34;standard&#34;)
 *             .healthCheckTaskCount(100)
 *             .smsNotificationCount(1000)
 *             .publicCnameMode(&#34;SYSTEM_ASSIGN&#34;)
 *             .ttl(60)
 *             .cnameType(&#34;PUBLIC&#34;)
 *             .resourceGroupId(default_.groups()[0].id())
 *             .alertGroups(defaultAlarmContactGroup.alarmContactGroupName())
 *             .publicUserDomainName(domainName)
 *             .alertConfigs(GtmInstanceAlertConfigArgs.builder()
 *                 .smsNotice(true)
 *                 .noticeType(&#34;ADDR_ALERT&#34;)
 *                 .emailNotice(true)
 *                 .dingtalkNotice(true)
 *                 .build())
 *             .build());
 * 
 *         var defaultAddressPool = new AddressPool(&#34;defaultAddressPool&#34;, AddressPoolArgs.builder()        
 *             .addressPoolName(name)
 *             .instanceId(defaultGtmInstance.id())
 *             .lbaStrategy(&#34;RATIO&#34;)
 *             .type(&#34;IPV4&#34;)
 *             .addresses(AddressPoolAddressArgs.builder()
 *                 .attributeInfo(&#34;&#34;&#34;
 *     {
 *       &#34;lineCodeRectifyType&#34;: &#34;RECTIFIED&#34;,
 *       &#34;lineCodes&#34;: [&#34;os_namerica_us&#34;]
 *     }
 *                 &#34;&#34;&#34;)
 *                 .remark(&#34;address_remark&#34;)
 *                 .address(&#34;1.1.1.1&#34;)
 *                 .mode(&#34;SMART&#34;)
 *                 .lbaWeight(1)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Alidns Address Pool can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dns/addressPool:AddressPool example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/addressPool:AddressPool")
public class AddressPool extends com.pulumi.resources.CustomResource {
    /**
     * The name of the address pool.
     * 
     */
    @Export(name="addressPoolName", refs={String.class}, tree="[0]")
    private Output<String> addressPoolName;

    /**
     * @return The name of the address pool.
     * 
     */
    public Output<String> addressPoolName() {
        return this.addressPoolName;
    }
    /**
     * The address lists of the Address Pool. See `address` below for details.
     * 
     */
    @Export(name="addresses", refs={List.class,AddressPoolAddress.class}, tree="[0,1]")
    private Output<List<AddressPoolAddress>> addresses;

    /**
     * @return The address lists of the Address Pool. See `address` below for details.
     * 
     */
    public Output<List<AddressPoolAddress>> addresses() {
        return this.addresses;
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
     * 
     */
    @Export(name="lbaStrategy", refs={String.class}, tree="[0]")
    private Output<String> lbaStrategy;

    /**
     * @return The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
     * 
     */
    public Output<String> lbaStrategy() {
        return this.lbaStrategy;
    }
    /**
     * The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AddressPool(String name) {
        this(name, AddressPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AddressPool(String name, AddressPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AddressPool(String name, AddressPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/addressPool:AddressPool", name, args == null ? AddressPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AddressPool(String name, Output<String> id, @Nullable AddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/addressPool:AddressPool", name, state, makeResourceOptions(options, id));
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
    public static AddressPool get(String name, Output<String> id, @Nullable AddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AddressPool(name, id, state, options);
    }
}
