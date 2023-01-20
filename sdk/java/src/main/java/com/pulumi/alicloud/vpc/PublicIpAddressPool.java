// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.PublicIpAddressPoolArgs;
import com.pulumi.alicloud.vpc.inputs.PublicIpAddressPoolState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Public Ip Address Pool resource.
 * 
 * For information about VPC Public Ip Address Pool and how to use it, see [What is Public Ip Address Pool](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createpublicipaddresspool).
 * 
 * &gt; **NOTE:** Available in v1.186.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.vpc.PublicIpAddressPool;
 * import com.pulumi.alicloud.vpc.PublicIpAddressPoolArgs;
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
 *         var default_ = new PublicIpAddressPool(&#34;default&#34;, PublicIpAddressPoolArgs.builder()        
 *             .description(&#34;example_value&#34;)
 *             .isp(&#34;BGP_PRO&#34;)
 *             .publicIpAddressPoolName(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Public Ip Address Pool can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/publicIpAddressPool:PublicIpAddressPool example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/publicIpAddressPool:PublicIpAddressPool")
public class PublicIpAddressPool extends com.pulumi.resources.CustomResource {
    /**
     * The description of the VPC Public IP address pool.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the VPC Public IP address pool.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
     * 
     */
    @Export(name="isp", type=String.class, parameters={})
    private Output<String> isp;

    /**
     * @return The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
     * 
     */
    public Output<String> isp() {
        return this.isp;
    }
    /**
     * The name of the VPC Public IP address pool.
     * 
     */
    @Export(name="publicIpAddressPoolName", type=String.class, parameters={})
    private Output</* @Nullable */ String> publicIpAddressPoolName;

    /**
     * @return The name of the VPC Public IP address pool.
     * 
     */
    public Output<Optional<String>> publicIpAddressPoolName() {
        return Codegen.optional(this.publicIpAddressPoolName);
    }
    /**
     * The status of the VPC Public IP address pool.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the VPC Public IP address pool.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PublicIpAddressPool(String name) {
        this(name, PublicIpAddressPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PublicIpAddressPool(String name, @Nullable PublicIpAddressPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PublicIpAddressPool(String name, @Nullable PublicIpAddressPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/publicIpAddressPool:PublicIpAddressPool", name, args == null ? PublicIpAddressPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PublicIpAddressPool(String name, Output<String> id, @Nullable PublicIpAddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/publicIpAddressPool:PublicIpAddressPool", name, state, makeResourceOptions(options, id));
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
    public static PublicIpAddressPool get(String name, Output<String> id, @Nullable PublicIpAddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PublicIpAddressPool(name, id, state, options);
    }
}
