// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.waf;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.waf.InstanceArgs;
import com.pulumi.alicloud.waf.inputs.InstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a WAF Instance resource to create instance in the Web Application Firewall.
 * 
 * For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
 * 
 * &gt; **NOTE:** Available in 1.83.0+ .
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.waf.Instance;
 * import com.pulumi.alicloud.waf.InstanceArgs;
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
 *         var default_ = new Instance(&#34;default&#34;, InstanceArgs.builder()        
 *             .bigScreen(&#34;0&#34;)
 *             .exclusiveIpPackage(&#34;1&#34;)
 *             .extBandwidth(&#34;50&#34;)
 *             .extDomainPackage(&#34;1&#34;)
 *             .logStorage(&#34;3&#34;)
 *             .logTime(&#34;180&#34;)
 *             .packageCode(&#34;version_3&#34;)
 *             .period(1)
 *             .prefessionalService(&#34;false&#34;)
 *             .resourceGroupId(&#34;rs-abc12345&#34;)
 *             .subscriptionType(&#34;Subscription&#34;)
 *             .wafLog(&#34;false&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * WAF instance can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:waf/instance:Instance default waf-cn-132435
 * ```
 * 
 */
@ResourceType(type="alicloud:waf/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Specify whether big screen is supported. Valid values: [&#34;0&#34;, &#34;1&#34;]. &#34;0&#34; for false and &#34;1&#34; for true.
     * 
     */
    @Export(name="bigScreen", type=String.class, parameters={})
    private Output<String> bigScreen;

    /**
     * @return Specify whether big screen is supported. Valid values: [&#34;0&#34;, &#34;1&#34;]. &#34;0&#34; for false and &#34;1&#34; for true.
     * 
     */
    public Output<String> bigScreen() {
        return this.bigScreen;
    }
    /**
     * Specify the number of exclusive WAF IP addresses.
     * 
     */
    @Export(name="exclusiveIpPackage", type=String.class, parameters={})
    private Output<String> exclusiveIpPackage;

    /**
     * @return Specify the number of exclusive WAF IP addresses.
     * 
     */
    public Output<String> exclusiveIpPackage() {
        return this.exclusiveIpPackage;
    }
    /**
     * The extra bandwidth. Unit: Mbit/s.
     * 
     */
    @Export(name="extBandwidth", type=String.class, parameters={})
    private Output<String> extBandwidth;

    /**
     * @return The extra bandwidth. Unit: Mbit/s.
     * 
     */
    public Output<String> extBandwidth() {
        return this.extBandwidth;
    }
    /**
     * The number of extra domains.
     * 
     */
    @Export(name="extDomainPackage", type=String.class, parameters={})
    private Output<String> extDomainPackage;

    /**
     * @return The number of extra domains.
     * 
     */
    public Output<String> extDomainPackage() {
        return this.extDomainPackage;
    }
    /**
     * Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
     * 
     */
    @Export(name="logStorage", type=String.class, parameters={})
    private Output<String> logStorage;

    /**
     * @return Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].
     * 
     */
    public Output<String> logStorage() {
        return this.logStorage;
    }
    /**
     * Log storage period. Unit: day. Valid values: [180, 360].
     * 
     */
    @Export(name="logTime", type=String.class, parameters={})
    private Output<String> logTime;

    /**
     * @return Log storage period. Unit: day. Valid values: [180, 360].
     * 
     */
    public Output<String> logTime() {
        return this.logTime;
    }
    /**
     * Type of configuration change. Valid value: Upgrade.
     * 
     */
    @Export(name="modifyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> modifyType;

    /**
     * @return Type of configuration change. Valid value: Upgrade.
     * 
     */
    public Output<Optional<String>> modifyType() {
        return Codegen.optional(this.modifyType);
    }
    /**
     * Subscription plan:
     * * China site customers can purchase the following versions of China Mainland region, valid values: [&#34;version_3&#34;, &#34;version_4&#34;, &#34;version_5&#34;].
     * * China site customers can purchase the following versions of International region, valid values: [&#34;version_pro_asia&#34;, &#34;version_business_asia&#34;, &#34;version_enterprise_asia&#34;]
     * * International site customers can purchase the following versions of China Mainland region: [&#34;version_pro_china&#34;, &#34;version_business_china&#34;, &#34;version_enterprise_china&#34;]
     * * International site customers can purchase the following versions of International region: [&#34;version_pro&#34;, &#34;version_business&#34;, &#34;version_enterprise&#34;].
     * 
     */
    @Export(name="packageCode", type=String.class, parameters={})
    private Output<String> packageCode;

    /**
     * @return Subscription plan:
     * * China site customers can purchase the following versions of China Mainland region, valid values: [&#34;version_3&#34;, &#34;version_4&#34;, &#34;version_5&#34;].
     * * China site customers can purchase the following versions of International region, valid values: [&#34;version_pro_asia&#34;, &#34;version_business_asia&#34;, &#34;version_enterprise_asia&#34;]
     * * International site customers can purchase the following versions of China Mainland region: [&#34;version_pro_china&#34;, &#34;version_business_china&#34;, &#34;version_enterprise_china&#34;]
     * * International site customers can purchase the following versions of International region: [&#34;version_pro&#34;, &#34;version_business&#34;, &#34;version_enterprise&#34;].
     * 
     */
    public Output<String> packageCode() {
        return this.packageCode;
    }
    /**
     * Service time of Web Application Firewall.
     * 
     */
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    /**
     * @return Service time of Web Application Firewall.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Specify whether professional service is supported. Valid values: [&#34;true&#34;, &#34;false&#34;]
     * 
     */
    @Export(name="prefessionalService", type=String.class, parameters={})
    private Output<String> prefessionalService;

    /**
     * @return Specify whether professional service is supported. Valid values: [&#34;true&#34;, &#34;false&#34;]
     * 
     */
    public Output<String> prefessionalService() {
        return this.prefessionalService;
    }
    /**
     * The instance region ID.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output</* @Nullable */ String> region;

    /**
     * @return The instance region ID.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * Renewal period of WAF service. Unit: month
     * 
     */
    @Export(name="renewPeriod", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> renewPeriod;

    /**
     * @return Renewal period of WAF service. Unit: month
     * 
     */
    public Output<Optional<Integer>> renewPeriod() {
        return Codegen.optional(this.renewPeriod);
    }
    /**
     * Renewal status of WAF service. Valid values:
     * * AutoRenewal: The service time of WAF is renewed automatically.
     * * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: &#34;On&#34; and &#34;Off&#34;. Default to &#34;Off&#34;.
     * 
     */
    @Export(name="renewalStatus", type=String.class, parameters={})
    private Output</* @Nullable */ String> renewalStatus;

    /**
     * @return Renewal status of WAF service. Valid values:
     * * AutoRenewal: The service time of WAF is renewed automatically.
     * * ManualRenewal (default): The service time of WAF is renewed manually.Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: &#34;On&#34; and &#34;Off&#34;. Default to &#34;Off&#34;.
     * 
     */
    public Output<Optional<String>> renewalStatus() {
        return Codegen.optional(this.renewalStatus);
    }
    /**
     * The resource group ID.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The resource group ID.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The status of the instance.
     * 
     */
    @Export(name="status", type=Integer.class, parameters={})
    private Output<Integer> status;

    /**
     * @return The status of the instance.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * Subscription of WAF service. Valid values: [&#34;Subscription&#34;, &#34;PayAsYouGo&#34;].
     * 
     */
    @Export(name="subscriptionType", type=String.class, parameters={})
    private Output<String> subscriptionType;

    /**
     * @return Subscription of WAF service. Valid values: [&#34;Subscription&#34;, &#34;PayAsYouGo&#34;].
     * 
     */
    public Output<String> subscriptionType() {
        return this.subscriptionType;
    }
    /**
     * Specify whether Log service is supported. Valid values: [&#34;true&#34;, &#34;false&#34;]
     * 
     */
    @Export(name="wafLog", type=String.class, parameters={})
    private Output<String> wafLog;

    /**
     * @return Specify whether Log service is supported. Valid values: [&#34;true&#34;, &#34;false&#34;]
     * 
     */
    public Output<String> wafLog() {
        return this.wafLog;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:waf/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:waf/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
