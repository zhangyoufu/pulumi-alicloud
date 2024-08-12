// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.DdosBgpInstanceArgs;
import com.pulumi.alicloud.dns.inputs.DdosBgpInstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Anti-DDoS Advanced instance resource. &#34;Ddosbgp&#34; is the short term of this product.
 * 
 * &gt; **NOTE:** The endpoint of bssopenapi used only support &#34;business.aliyuncs.com&#34; at present.
 * 
 * &gt; **NOTE:** Available since v1.183.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Ddosbgp instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dns/ddosBgpInstance:DdosBgpInstance example ddosbgp-abc123456
 * ```
 * 
 * @deprecated
 * alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance
 * 
 */
@Deprecated /* alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance */
@ResourceType(type="alicloud:dns/ddosBgpInstance:DdosBgpInstance")
public class DdosBgpInstance extends com.pulumi.resources.CustomResource {
    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
     * 
     */
    @Export(name="baseBandwidth", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> baseBandwidth;

    /**
     * @return Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
     * 
     */
    public Output<Optional<Integer>> baseBandwidth() {
        return Codegen.optional(this.baseBandwidth);
    }
    /**
     * IP count of the instance. Valid values: 100.
     * 
     */
    @Export(name="ipCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> ipCount;

    /**
     * @return IP count of the instance. Valid values: 100.
     * 
     */
    public Output<Integer> ipCount() {
        return this.ipCount;
    }
    /**
     * IP version of the instance. Valid values: IPv4,IPv6.
     * 
     */
    @Export(name="ipType", refs={String.class}, tree="[0]")
    private Output<String> ipType;

    /**
     * @return IP version of the instance. Valid values: IPv4,IPv6.
     * 
     */
    public Output<String> ipType() {
        return this.ipType;
    }
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the instance. This name can have a string of 1 to 63 characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Normal defend bandwidth of the instance. The unit is Gbps.
     * 
     */
    @Export(name="normalBandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> normalBandwidth;

    /**
     * @return Normal defend bandwidth of the instance. The unit is Gbps.
     * 
     */
    public Output<Integer> normalBandwidth() {
        return this.normalBandwidth;
    }
    /**
     * The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify &#34;period&#34;.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify &#34;period&#34;.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DdosBgpInstance(java.lang.String name) {
        this(name, DdosBgpInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DdosBgpInstance(java.lang.String name, DdosBgpInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DdosBgpInstance(java.lang.String name, DdosBgpInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/ddosBgpInstance:DdosBgpInstance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DdosBgpInstance(java.lang.String name, Output<java.lang.String> id, @Nullable DdosBgpInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/ddosBgpInstance:DdosBgpInstance", name, state, makeResourceOptions(options, id), false);
    }

    private static DdosBgpInstanceArgs makeArgs(DdosBgpInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DdosBgpInstanceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static DdosBgpInstance get(java.lang.String name, Output<java.lang.String> id, @Nullable DdosBgpInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DdosBgpInstance(name, id, state, options);
    }
}
