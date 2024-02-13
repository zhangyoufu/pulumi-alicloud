// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ens.KeyPairArgs;
import com.pulumi.alicloud.ens.inputs.KeyPairState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a ENS Key Pair resource.
 * 
 * For information about ENS Key Pair and how to use it, see [What is Key Pair](https://www.alibabacloud.com/help/en/ens/latest/createkeypair).
 * 
 * &gt; **NOTE:** Available since v1.133.0.
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
 * import com.pulumi.alicloud.ens.KeyPair;
 * import com.pulumi.alicloud.ens.KeyPairArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var example = new KeyPair(&#34;example&#34;, KeyPairArgs.builder()        
 *             .keyPairName(name)
 *             .version(&#34;2017-11-10&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ENS Key Pair can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ens/keyPair:KeyPair example &lt;key_pair_name&gt;:&lt;version&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ens/keyPair:KeyPair")
public class KeyPair extends com.pulumi.resources.CustomResource {
    /**
     * The name of the key pair.
     * 
     */
    @Export(name="keyPairName", refs={String.class}, tree="[0]")
    private Output<String> keyPairName;

    /**
     * @return The name of the key pair.
     * 
     */
    public Output<String> keyPairName() {
        return this.keyPairName;
    }
    /**
     * The version number.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The version number.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KeyPair(String name) {
        this(name, KeyPairArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KeyPair(String name, KeyPairArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KeyPair(String name, KeyPairArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/keyPair:KeyPair", name, args == null ? KeyPairArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KeyPair(String name, Output<String> id, @Nullable KeyPairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/keyPair:KeyPair", name, state, makeResourceOptions(options, id));
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
    public static KeyPair get(String name, Output<String> id, @Nullable KeyPairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KeyPair(name, id, state, options);
    }
}
