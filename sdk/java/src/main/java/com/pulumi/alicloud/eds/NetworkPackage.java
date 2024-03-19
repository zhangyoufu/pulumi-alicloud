// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.NetworkPackageArgs;
import com.pulumi.alicloud.eds.inputs.NetworkPackageState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a ECD Network Package resource.
 * 
 * For information about ECD Network Package and how to use it, see [What is Network Package](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createnetworkpackage).
 * 
 * &gt; **NOTE:** Available since v1.142.0.
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
 * import com.pulumi.alicloud.eds.EdsFunctions;
 * import com.pulumi.alicloud.eds.inputs.GetSimpleOfficeSitesArgs;
 * import com.pulumi.alicloud.eds.NetworkPackage;
 * import com.pulumi.alicloud.eds.NetworkPackageArgs;
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
 *         final var defaultSimpleOfficeSites = EdsFunctions.getSimpleOfficeSites(GetSimpleOfficeSitesArgs.builder()
 *             .status(&#34;REGISTERED&#34;)
 *             .nameRegex(&#34;default&#34;)
 *             .build());
 * 
 *         var defaultNetworkPackage = new NetworkPackage(&#34;defaultNetworkPackage&#34;, NetworkPackageArgs.builder()        
 *             .bandwidth(10)
 *             .officeSiteId(defaultSimpleOfficeSites.applyValue(getSimpleOfficeSitesResult -&gt; getSimpleOfficeSitesResult.ids()[0]))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ECD Network Package can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/networkPackage:NetworkPackage example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/networkPackage:NetworkPackage")
public class NetworkPackage extends com.pulumi.resources.CustomResource {
    /**
     * The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The internet charge type  of  package.
     * 
     */
    @Export(name="internetChargeType", refs={String.class}, tree="[0]")
    private Output<String> internetChargeType;

    /**
     * @return The internet charge type  of  package.
     * 
     */
    public Output<String> internetChargeType() {
        return this.internetChargeType;
    }
    /**
     * The ID of office site.
     * 
     */
    @Export(name="officeSiteId", refs={String.class}, tree="[0]")
    private Output<String> officeSiteId;

    /**
     * @return The ID of office site.
     * 
     */
    public Output<String> officeSiteId() {
        return this.officeSiteId;
    }
    /**
     * The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkPackage(String name) {
        this(name, NetworkPackageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkPackage(String name, NetworkPackageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkPackage(String name, NetworkPackageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/networkPackage:NetworkPackage", name, args == null ? NetworkPackageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkPackage(String name, Output<String> id, @Nullable NetworkPackageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/networkPackage:NetworkPackage", name, state, makeResourceOptions(options, id));
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
    public static NetworkPackage get(String name, Output<String> id, @Nullable NetworkPackageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkPackage(name, id, state, options);
    }
}
