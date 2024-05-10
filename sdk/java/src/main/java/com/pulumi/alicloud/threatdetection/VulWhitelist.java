// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.VulWhitelistArgs;
import com.pulumi.alicloud.threatdetection.inputs.VulWhitelistState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Vul Whitelist resource.
 * 
 * For information about Threat Detection Vul Whitelist and how to use it, see [What is Vul Whitelist](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
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
 * import com.pulumi.alicloud.threatdetection.VulWhitelist;
 * import com.pulumi.alicloud.threatdetection.VulWhitelistArgs;
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
 *         var default_ = new VulWhitelist("default", VulWhitelistArgs.builder()        
 *             .whitelist("[{\"aliasName\":\"RHSA-2021:2260: libwebp 安全更新\",\"name\":\"RHSA-2021:2260: libwebp 安全更新\",\"type\":\"cve\"}]")
 *             .targetInfo("{\"type\":\"GroupId\",\"uuids\":[],\"groupIds\":[10782678]}")
 *             .reason("tf-example-reason")
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
 * Threat Detection Vul Whitelist can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/vulWhitelist:VulWhitelist example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/vulWhitelist:VulWhitelist")
public class VulWhitelist extends com.pulumi.resources.CustomResource {
    /**
     * Reason for adding whitelist.
     * 
     */
    @Export(name="reason", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> reason;

    /**
     * @return Reason for adding whitelist.
     * 
     */
    public Output<Optional<String>> reason() {
        return Codegen.optional(this.reason);
    }
    /**
     * Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
     * 
     */
    @Export(name="targetInfo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> targetInfo;

    /**
     * @return Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
     * 
     */
    public Output<Optional<String>> targetInfo() {
        return Codegen.optional(this.targetInfo);
    }
    /**
     * Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
     * 
     */
    @Export(name="whitelist", refs={String.class}, tree="[0]")
    private Output<String> whitelist;

    /**
     * @return Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
     * 
     */
    public Output<String> whitelist() {
        return this.whitelist;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VulWhitelist(String name) {
        this(name, VulWhitelistArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VulWhitelist(String name, VulWhitelistArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VulWhitelist(String name, VulWhitelistArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/vulWhitelist:VulWhitelist", name, args == null ? VulWhitelistArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VulWhitelist(String name, Output<String> id, @Nullable VulWhitelistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/vulWhitelist:VulWhitelist", name, state, makeResourceOptions(options, id));
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
    public static VulWhitelist get(String name, Output<String> id, @Nullable VulWhitelistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VulWhitelist(name, id, state, options);
    }
}
