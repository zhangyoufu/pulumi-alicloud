// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudauth;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudauth.FaceConfigArgs;
import com.pulumi.alicloud.cloudauth.inputs.FaceConfigState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Cloudauth Face Config resource.
 * 
 * For information about Cloudauth Face Config and how to use it, see [What is Face Config](https://help.aliyun.com/zh/id-verification/cloudauth/product-overview/end-of-integration-announcement-on-id-verification).
 * 
 * &gt; **NOTE:** Available since v1.137.0.
 * 
 * &gt; **NOTE:** In order to provide you with more perfect product capabilities, the real person certification service has stopped access, it is recommended that you use the upgraded version of the [real person certification financial real person certification service](https://help.aliyun.com/zh/id-verification/product-overview/what-is-id-verification-for-financial-services). Users that have access to real person authentication are not affected.
 * 
 * ## Import
 * 
 * Cloudauth Face Config can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudauth/faceConfig:FaceConfig example &lt;lang&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudauth/faceConfig:FaceConfig")
public class FaceConfig extends com.pulumi.resources.CustomResource {
    /**
     * Scene name.
     * 
     */
    @Export(name="bizName", refs={String.class}, tree="[0]")
    private Output<String> bizName;

    /**
     * @return Scene name.
     * 
     */
    public Output<String> bizName() {
        return this.bizName;
    }
    /**
     * Scene type. **NOTE:** The biz_type cannot exceed 32 characters and can only use English letters, numbers and dashes (-).
     * 
     */
    @Export(name="bizType", refs={String.class}, tree="[0]")
    private Output<String> bizType;

    /**
     * @return Scene type. **NOTE:** The biz_type cannot exceed 32 characters and can only use English letters, numbers and dashes (-).
     * 
     */
    public Output<String> bizType() {
        return this.bizType;
    }
    /**
     * Last Modified Date.
     * 
     */
    @Export(name="gmtModified", refs={String.class}, tree="[0]")
    private Output<String> gmtModified;

    /**
     * @return Last Modified Date.
     * 
     */
    public Output<String> gmtModified() {
        return this.gmtModified;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FaceConfig(java.lang.String name) {
        this(name, FaceConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FaceConfig(java.lang.String name, FaceConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FaceConfig(java.lang.String name, FaceConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudauth/faceConfig:FaceConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FaceConfig(java.lang.String name, Output<java.lang.String> id, @Nullable FaceConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudauth/faceConfig:FaceConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static FaceConfigArgs makeArgs(FaceConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FaceConfigArgs.Empty : args;
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
    public static FaceConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable FaceConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FaceConfig(name, id, state, options);
    }
}
