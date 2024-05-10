// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.DdosCooInstanceArgs;
import com.pulumi.alicloud.dns.inputs.DdosCooInstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a BGP-line Anti-DDoS Pro(DdosCoo) Instance resource.
 * 
 * For information about BGP-line Anti-DDoS Pro(DdosCoo) Instance and how to use it, see [What is Anti-DDoS Pro Instance](https://www.alibabacloud.com/help/en/ddos-protection/latest/create-an-anti-ddos-pro-or-anti-ddos-premium-instance-by-calling-an-api-operation).
 * 
 * &gt; **NOTE:** Available since v1.37.0.
 * 
 * &gt; **NOTE:** The endpoint of bssopenapi used only support &#34;business.aliyuncs.com&#34; at present.
 * 
 * &gt; **NOTE:** From version 1.214.0, if `product_type` is set to `ddoscoo` or `ddoscoo_intl`, the provider `region` should be set to `cn-hangzhou`, and if `product_type` is set to `ddosDip`, the provider `region` should be set to `ap-southeast-1`.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ddos.DdosCooInstance;
 * import com.pulumi.alicloud.ddos.DdosCooInstanceArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         var default_ = new DdosCooInstance("default", DdosCooInstanceArgs.builder()        
 *             .name(name)
 *             .baseBandwidth("30")
 *             .bandwidth("30")
 *             .serviceBandwidth("100")
 *             .portCount("50")
 *             .domainCount("50")
 *             .productType("ddoscoo")
 *             .period("1")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * DdosCoo instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dns/ddosCooInstance:DdosCooInstance example ddoscoo-cn-123456
 * ```
 * 
 * @deprecated
 * alicloud.dns.DdosCooInstance has been deprecated in favor of alicloud.ddos.DdosCooInstance
 * 
 */
@Deprecated /* alicloud.dns.DdosCooInstance has been deprecated in favor of alicloud.ddos.DdosCooInstance */
@ResourceType(type="alicloud:dns/ddosCooInstance:DdosCooInstance")
public class DdosCooInstance extends com.pulumi.resources.CustomResource {
    /**
     * The IP version of the IP address. Default value: `Ipv4`. Valid values: `Ipv4`, `Ipv6`. **NOTE:** `address_type` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    @Export(name="addressType", refs={String.class}, tree="[0]")
    private Output<String> addressType;

    /**
     * @return The IP version of the IP address. Default value: `Ipv4`. Valid values: `Ipv4`, `Ipv6`. **NOTE:** `address_type` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    public Output<String> addressType() {
        return this.addressType;
    }
    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: `30`, `60`, `100`, `300`, `400`, `500`, `600`. The unit is Gbps. Only support upgrade. **NOTE:** `bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    @Export(name="bandwidth", refs={String.class}, tree="[0]")
    private Output<String> bandwidth;

    /**
     * @return Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: `30`, `60`, `100`, `300`, `400`, `500`, `600`. The unit is Gbps. Only support upgrade. **NOTE:** `bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    public Output<String> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The mitigation plan of the instance. Valid values:
     * 
     */
    @Export(name="bandwidthMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bandwidthMode;

    /**
     * @return The mitigation plan of the instance. Valid values:
     * 
     */
    public Output<Optional<String>> bandwidthMode() {
        return Codegen.optional(this.bandwidthMode);
    }
    /**
     * Base defend bandwidth of the instance. Valid values: `30`, `60`, `100`, `300`, `400`, `500`, `600`. The unit is Gbps. Only support upgrade. **NOTE:** `base_bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    @Export(name="baseBandwidth", refs={String.class}, tree="[0]")
    private Output<String> baseBandwidth;

    /**
     * @return Base defend bandwidth of the instance. Valid values: `30`, `60`, `100`, `300`, `400`, `500`, `600`. The unit is Gbps. Only support upgrade. **NOTE:** `base_bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    public Output<String> baseBandwidth() {
        return this.baseBandwidth;
    }
    /**
     * Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     * 
     */
    @Export(name="domainCount", refs={String.class}, tree="[0]")
    private Output<String> domainCount;

    /**
     * @return Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     * 
     */
    public Output<String> domainCount() {
        return this.domainCount;
    }
    /**
     * The mitigation plan of the instance. Default value: `coop`. Valid values:
     * 
     */
    @Export(name="editionSale", refs={String.class}, tree="[0]")
    private Output<String> editionSale;

    /**
     * @return The mitigation plan of the instance. Default value: `coop`. Valid values:
     * 
     */
    public Output<String> editionSale() {
        return this.editionSale;
    }
    /**
     * The function plan of the instance. Valid values:
     * 
     */
    @Export(name="functionVersion", refs={String.class}, tree="[0]")
    private Output<String> functionVersion;

    /**
     * @return The function plan of the instance. Valid values:
     * 
     */
    public Output<String> functionVersion() {
        return this.functionVersion;
    }
    /**
     * (Available since v1.212.0) The IP address of the instance.
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return (Available since v1.212.0) The IP address of the instance.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Name of the instance. This name can have a string of `1` to `64` characters.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the instance. This name can have a string of `1` to `64` characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The clean bandwidth provided by the instance. **NOTE:** `normal_bandwidth` is valid only when `product_type` is set to `ddosDip`.
     * 
     */
    @Export(name="normalBandwidth", refs={String.class}, tree="[0]")
    private Output<String> normalBandwidth;

    /**
     * @return The clean bandwidth provided by the instance. **NOTE:** `normal_bandwidth` is valid only when `product_type` is set to `ddosDip`.
     * 
     */
    public Output<String> normalBandwidth() {
        return this.normalBandwidth;
    }
    /**
     * The clean QPS provided by the instance. **NOTE:** `normal_qps` is valid only when `product_type` is set to `ddosDip`.
     * 
     */
    @Export(name="normalQps", refs={String.class}, tree="[0]")
    private Output<String> normalQps;

    /**
     * @return The clean QPS provided by the instance. **NOTE:** `normal_qps` is valid only when `product_type` is set to `ddosDip`.
     * 
     */
    public Output<String> normalQps() {
        return this.normalQps;
    }
    /**
     * The duration that you will buy DdosCoo instance (in month). Valid values: [1~9], `12`, `24`, `36`. Default value: `1`. At present, the provider does not support modify `period`.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The duration that you will buy DdosCoo instance (in month). Valid values: [1~9], `12`, `24`, `36`. Default value: `1`. At present, the provider does not support modify `period`.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     * 
     */
    @Export(name="portCount", refs={String.class}, tree="[0]")
    private Output<String> portCount;

    /**
     * @return Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     * 
     */
    public Output<String> portCount() {
        return this.portCount;
    }
    /**
     * The mitigation plan of the instance. Valid values:
     * 
     */
    @Export(name="productPlan", refs={String.class}, tree="[0]")
    private Output<String> productPlan;

    /**
     * @return The mitigation plan of the instance. Valid values:
     * 
     */
    public Output<String> productPlan() {
        return this.productPlan;
    }
    /**
     * The product type for purchasing DDOSCOO instances used to differ different account type. Default value: `ddoscoo`. Valid values:
     * 
     */
    @Export(name="productType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> productType;

    /**
     * @return The product type for purchasing DDOSCOO instances used to differ different account type. Default value: `ddoscoo`. Valid values:
     * 
     */
    public Output<Optional<String>> productType() {
        return Codegen.optional(this.productType);
    }
    /**
     * Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade. **NOTE:** `service_bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    @Export(name="serviceBandwidth", refs={String.class}, tree="[0]")
    private Output<String> serviceBandwidth;

    /**
     * @return Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade. **NOTE:** `service_bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
     * 
     */
    public Output<String> serviceBandwidth() {
        return this.serviceBandwidth;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DdosCooInstance(String name) {
        this(name, DdosCooInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DdosCooInstance(String name, DdosCooInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DdosCooInstance(String name, DdosCooInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/ddosCooInstance:DdosCooInstance", name, args == null ? DdosCooInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DdosCooInstance(String name, Output<String> id, @Nullable DdosCooInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/ddosCooInstance:DdosCooInstance", name, state, makeResourceOptions(options, id));
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
    public static DdosCooInstance get(String name, Output<String> id, @Nullable DdosCooInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DdosCooInstance(name, id, state, options);
    }
}
