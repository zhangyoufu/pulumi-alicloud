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
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Public Ip Address Pool resource.
 * 
 * For information about VPC Public Ip Address Pool and how to use it, see [What is Public Ip Address Pool](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createpublicipaddresspool).
 * 
 * &gt; **NOTE:** Available since v1.186.0.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
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
 *         final var config = ctx.config();
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status("OK")
 *             .build());
 * 
 *         var defaultPublicIpAddressPool = new PublicIpAddressPool("defaultPublicIpAddressPool", PublicIpAddressPoolArgs.builder()
 *             .description(name)
 *             .publicIpAddressPoolName(name)
 *             .isp("BGP")
 *             .resourceGroupId(default_.ids()[0])
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
 * VPC Public Ip Address Pool can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/publicIpAddressPool:PublicIpAddressPool example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/publicIpAddressPool:PublicIpAddressPool")
public class PublicIpAddressPool extends com.pulumi.resources.CustomResource {
    /**
     * The name of the VPC Public IP address pool.
     * 
     */
    @Export(name="bizType", refs={String.class}, tree="[0]")
    private Output<String> bizType;

    /**
     * @return The name of the VPC Public IP address pool.
     * 
     */
    public Output<String> bizType() {
        return this.bizType;
    }
    /**
     * The creation time of the resource
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the resource
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether there is a free IP address.
     * 
     */
    @Export(name="ipAddressRemaining", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ipAddressRemaining;

    /**
     * @return Whether there is a free IP address.
     * 
     */
    public Output<Boolean> ipAddressRemaining() {
        return this.ipAddressRemaining;
    }
    /**
     * The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
     * 
     */
    @Export(name="isp", refs={String.class}, tree="[0]")
    private Output<String> isp;

    /**
     * @return The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
     * 
     */
    public Output<String> isp() {
        return this.isp;
    }
    @Export(name="publicIpAddressPoolId", refs={String.class}, tree="[0]")
    private Output<String> publicIpAddressPoolId;

    public Output<String> publicIpAddressPoolId() {
        return this.publicIpAddressPoolId;
    }
    /**
     * The name of the VPC Public IP address pool.
     * 
     */
    @Export(name="publicIpAddressPoolName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> publicIpAddressPoolName;

    /**
     * @return The name of the VPC Public IP address pool.
     * 
     */
    public Output<Optional<String>> publicIpAddressPoolName() {
        return Codegen.optional(this.publicIpAddressPoolName);
    }
    /**
     * The resource group ID of the VPC Public IP address pool.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The resource group ID of the VPC Public IP address pool.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Security protection level.
     * - If the configuration is empty, the default value is DDoS protection (Basic edition).
     * - `AntiDDoS_Enhanced` indicates DDoS protection (enhanced version).
     * 
     */
    @Export(name="securityProtectionTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityProtectionTypes;

    /**
     * @return Security protection level.
     * - If the configuration is empty, the default value is DDoS protection (Basic edition).
     * - `AntiDDoS_Enhanced` indicates DDoS protection (enhanced version).
     * 
     */
    public Output<Optional<List<String>>> securityProtectionTypes() {
        return Codegen.optional(this.securityProtectionTypes);
    }
    /**
     * The status of the VPC Public IP address pool.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the VPC Public IP address pool.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags of PrefixList.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tags of PrefixList.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The total number of public IP address pools.
     * 
     */
    @Export(name="totalIpNum", refs={Integer.class}, tree="[0]")
    private Output<Integer> totalIpNum;

    /**
     * @return The total number of public IP address pools.
     * 
     */
    public Output<Integer> totalIpNum() {
        return this.totalIpNum;
    }
    /**
     * The number of used IP addresses in the public IP address pool.
     * 
     */
    @Export(name="usedIpNum", refs={Integer.class}, tree="[0]")
    private Output<Integer> usedIpNum;

    /**
     * @return The number of used IP addresses in the public IP address pool.
     * 
     */
    public Output<Integer> usedIpNum() {
        return this.usedIpNum;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PublicIpAddressPool(java.lang.String name) {
        this(name, PublicIpAddressPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PublicIpAddressPool(java.lang.String name, @Nullable PublicIpAddressPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PublicIpAddressPool(java.lang.String name, @Nullable PublicIpAddressPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/publicIpAddressPool:PublicIpAddressPool", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PublicIpAddressPool(java.lang.String name, Output<java.lang.String> id, @Nullable PublicIpAddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/publicIpAddressPool:PublicIpAddressPool", name, state, makeResourceOptions(options, id), false);
    }

    private static PublicIpAddressPoolArgs makeArgs(@Nullable PublicIpAddressPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PublicIpAddressPoolArgs.Empty : args;
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
    public static PublicIpAddressPool get(java.lang.String name, Output<java.lang.String> id, @Nullable PublicIpAddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PublicIpAddressPool(name, id, state, options);
    }
}
