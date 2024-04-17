// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.apigateway.AppAttachmentArgs;
import com.pulumi.alicloud.apigateway.inputs.AppAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.alicloud.apigateway.Group;
 * import com.pulumi.alicloud.apigateway.GroupArgs;
 * import com.pulumi.alicloud.apigateway.Api;
 * import com.pulumi.alicloud.apigateway.ApiArgs;
 * import com.pulumi.alicloud.apigateway.inputs.ApiRequestConfigArgs;
 * import com.pulumi.alicloud.apigateway.inputs.ApiHttpServiceConfigArgs;
 * import com.pulumi.alicloud.apigateway.inputs.ApiRequestParameterArgs;
 * import com.pulumi.alicloud.apigateway.App;
 * import com.pulumi.alicloud.apigateway.AppArgs;
 * import com.pulumi.alicloud.apigateway.AppAttachment;
 * import com.pulumi.alicloud.apigateway.AppAttachmentArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform_example&#34;);
 *         var example = new Group(&#34;example&#34;, GroupArgs.builder()        
 *             .name(name)
 *             .description(name)
 *             .build());
 * 
 *         var exampleApi = new Api(&#34;exampleApi&#34;, ApiArgs.builder()        
 *             .groupId(example.id())
 *             .name(name)
 *             .description(name)
 *             .authType(&#34;APP&#34;)
 *             .forceNonceCheck(false)
 *             .requestConfig(ApiRequestConfigArgs.builder()
 *                 .protocol(&#34;HTTP&#34;)
 *                 .method(&#34;GET&#34;)
 *                 .path(&#34;/example/path&#34;)
 *                 .mode(&#34;MAPPING&#34;)
 *                 .build())
 *             .serviceType(&#34;HTTP&#34;)
 *             .httpServiceConfig(ApiHttpServiceConfigArgs.builder()
 *                 .address(&#34;http://apigateway-backend.alicloudapi.com:8080&#34;)
 *                 .method(&#34;GET&#34;)
 *                 .path(&#34;/web/cloudapi&#34;)
 *                 .timeout(12)
 *                 .aoneName(&#34;cloudapi-openapi&#34;)
 *                 .build())
 *             .requestParameters(ApiRequestParameterArgs.builder()
 *                 .name(&#34;example&#34;)
 *                 .type(&#34;STRING&#34;)
 *                 .required(&#34;OPTIONAL&#34;)
 *                 .in(&#34;QUERY&#34;)
 *                 .inService(&#34;QUERY&#34;)
 *                 .nameService(&#34;exampleservice&#34;)
 *                 .build())
 *             .stageNames(            
 *                 &#34;RELEASE&#34;,
 *                 &#34;TEST&#34;)
 *             .build());
 * 
 *         var exampleApp = new App(&#34;exampleApp&#34;, AppArgs.builder()        
 *             .name(name)
 *             .description(name)
 *             .build());
 * 
 *         var exampleAppAttachment = new AppAttachment(&#34;exampleAppAttachment&#34;, AppAttachmentArgs.builder()        
 *             .apiId(exampleApi.apiId())
 *             .groupId(example.id())
 *             .appId(exampleApp.id())
 *             .stageName(&#34;PRE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="alicloud:apigateway/appAttachment:AppAttachment")
public class AppAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The api_id that app apply to access.
     * 
     */
    @Export(name="apiId", refs={String.class}, tree="[0]")
    private Output<String> apiId;

    /**
     * @return The api_id that app apply to access.
     * 
     */
    public Output<String> apiId() {
        return this.apiId;
    }
    /**
     * The app that apply to the authorization.
     * 
     */
    @Export(name="appId", refs={String.class}, tree="[0]")
    private Output<String> appId;

    /**
     * @return The app that apply to the authorization.
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * The group that the api belongs to.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The group that the api belongs to.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * Stage that the app apply to access.
     * 
     */
    @Export(name="stageName", refs={String.class}, tree="[0]")
    private Output<String> stageName;

    /**
     * @return Stage that the app apply to access.
     * 
     */
    public Output<String> stageName() {
        return this.stageName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppAttachment(String name) {
        this(name, AppAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppAttachment(String name, AppAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppAttachment(String name, AppAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:apigateway/appAttachment:AppAttachment", name, args == null ? AppAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppAttachment(String name, Output<String> id, @Nullable AppAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:apigateway/appAttachment:AppAttachment", name, state, makeResourceOptions(options, id));
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
    public static AppAttachment get(String name, Output<String> id, @Nullable AppAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppAttachment(name, id, state, options);
    }
}
