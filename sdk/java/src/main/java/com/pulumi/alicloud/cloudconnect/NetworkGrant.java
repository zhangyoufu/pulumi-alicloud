// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudconnect;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudconnect.NetworkGrantArgs;
import com.pulumi.alicloud.cloudconnect.inputs.NetworkGrantState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Connect Network Grant resource. If the CEN instance to be attached belongs to another account, authorization by the CEN instance is required.
 * 
 * For information about Cloud Connect Network Grant and how to use it, see [What is Cloud Connect Network Grant](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/grantinstancetocbn).
 * 
 * &gt; **NOTE:** Available since v1.63.0.
 * 
 * &gt; **NOTE:** Only the following regions support create Cloud Connect Network Grant. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
 * import com.pulumi.alicloud.cloudconnect.Network;
 * import com.pulumi.alicloud.cloudconnect.NetworkArgs;
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cloudconnect.NetworkGrant;
 * import com.pulumi.alicloud.cloudconnect.NetworkGrantArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var cenUid = config.get(&#34;cenUid&#34;).orElse(123456789);
 *         // Method 2: Use the target cen account&#39;s access_key, secret_key
 *         // provider &#34;alicloud&#34; {
 *         //   region     = &#34;cn-hangzhou&#34;
 *         //   access_key = &#34;access_key&#34;
 *         //   secret_key = &#34;secret_key&#34;
 *         //   alias      = &#34;cen_account&#34;
 *         // }
 *         var default_ = new Network(&#34;default&#34;, NetworkArgs.builder()        
 *             .name(name)
 *             .description(name)
 *             .cidrBlock(&#34;192.168.0.0/24&#34;)
 *             .isDefault(true)
 *             .build());
 * 
 *         var cen = new Instance(&#34;cen&#34;, InstanceArgs.builder()        
 *             .cenInstanceName(name)
 *             .build());
 * 
 *         var defaultNetworkGrant = new NetworkGrant(&#34;defaultNetworkGrant&#34;, NetworkGrantArgs.builder()        
 *             .ccnId(default_.id())
 *             .cenId(cen.id())
 *             .cenUid(cenUid)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * The Cloud Connect Network Grant can be imported using the instance_id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudconnect/networkGrant:NetworkGrant example ccn-abc123456:cen-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudconnect/networkGrant:NetworkGrant")
public class NetworkGrant extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the CCN instance.
     * 
     */
    @Export(name="ccnId", refs={String.class}, tree="[0]")
    private Output<String> ccnId;

    /**
     * @return The ID of the CCN instance.
     * 
     */
    public Output<String> ccnId() {
        return this.ccnId;
    }
    /**
     * The ID of the CEN instance.
     * 
     */
    @Export(name="cenId", refs={String.class}, tree="[0]")
    private Output<String> cenId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Output<String> cenId() {
        return this.cenId;
    }
    /**
     * The ID of the account to which the CEN instance belongs.
     * 
     */
    @Export(name="cenUid", refs={String.class}, tree="[0]")
    private Output<String> cenUid;

    /**
     * @return The ID of the account to which the CEN instance belongs.
     * 
     */
    public Output<String> cenUid() {
        return this.cenUid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkGrant(String name) {
        this(name, NetworkGrantArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkGrant(String name, NetworkGrantArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkGrant(String name, NetworkGrantArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudconnect/networkGrant:NetworkGrant", name, args == null ? NetworkGrantArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkGrant(String name, Output<String> id, @Nullable NetworkGrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudconnect/networkGrant:NetworkGrant", name, state, makeResourceOptions(options, id));
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
    public static NetworkGrant get(String name, Output<String> id, @Nullable NetworkGrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkGrant(name, id, state, options);
    }
}
