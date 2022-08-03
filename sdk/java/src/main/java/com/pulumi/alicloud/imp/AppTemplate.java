// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.imp;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.imp.AppTemplateArgs;
import com.pulumi.alicloud.imp.inputs.AppTemplateState;
import com.pulumi.alicloud.imp.outputs.AppTemplateConfigList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Apsara Agile Live (IMP) App Template resource.
 * 
 * For information about Apsara Agile Live (IMP) App Template and how to use it, see [What is App Template](https://help.aliyun.com/document_detail/270121.html).
 * 
 * &gt; **NOTE:** Available in v1.137.0+.
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
 * import com.pulumi.alicloud.imp.AppTemplate;
 * import com.pulumi.alicloud.imp.AppTemplateArgs;
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
 *         var example = new AppTemplate(&#34;example&#34;, AppTemplateArgs.builder()        
 *             .appTemplateName(&#34;example_value&#34;)
 *             .componentLists(            
 *                 &#34;component.live&#34;,
 *                 &#34;component.liveRecord&#34;)
 *             .integrationMode(&#34;paasSDK&#34;)
 *             .scene(&#34;business&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Apsara Agile Live (IMP) App Template can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:imp/appTemplate:AppTemplate example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:imp/appTemplate:AppTemplate")
public class AppTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The name of the resource.
     * 
     */
    @Export(name="appTemplateName", type=String.class, parameters={})
    private Output<String> appTemplateName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> appTemplateName() {
        return this.appTemplateName;
    }
    /**
     * List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
     * 
     */
    @Export(name="componentLists", type=List.class, parameters={String.class})
    private Output<List<String>> componentLists;

    /**
     * @return List of components. Its element valid values: [&#34;component.live&#34;,&#34;component.liveRecord&#34;,&#34;component.liveBeauty&#34;,&#34;component.rtc&#34;,&#34;component.rtcRecord&#34;,&#34;component.im&#34;,&#34;component.whiteboard&#34;,&#34;component.liveSecurity&#34;,&#34;component.chatSecurity&#34;].
     * 
     */
    public Output<List<String>> componentLists() {
        return this.componentLists;
    }
    /**
     * Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
     * 
     */
    @Export(name="configLists", type=List.class, parameters={AppTemplateConfigList.class})
    private Output<List<AppTemplateConfigList>> configLists;

    /**
     * @return Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
     * 
     */
    public Output<List<AppTemplateConfigList>> configLists() {
        return this.configLists;
    }
    /**
     * Integration mode. Valid values:
     * * paasSDK: Integrated SDK.
     * * standardRoom: Model Room.
     * 
     */
    @Export(name="integrationMode", type=String.class, parameters={})
    private Output</* @Nullable */ String> integrationMode;

    /**
     * @return Integration mode. Valid values:
     * * paasSDK: Integrated SDK.
     * * standardRoom: Model Room.
     * 
     */
    public Output<Optional<String>> integrationMode() {
        return Codegen.optional(this.integrationMode);
    }
    /**
     * Application Template scenario. Valid values: [&#34;business&#34;, &#34;classroom&#34;].
     * 
     */
    @Export(name="scene", type=String.class, parameters={})
    private Output</* @Nullable */ String> scene;

    /**
     * @return Application Template scenario. Valid values: [&#34;business&#34;, &#34;classroom&#34;].
     * 
     */
    public Output<Optional<String>> scene() {
        return Codegen.optional(this.scene);
    }
    /**
     * Application template usage status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Application template usage status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppTemplate(String name) {
        this(name, AppTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppTemplate(String name, AppTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppTemplate(String name, AppTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:imp/appTemplate:AppTemplate", name, args == null ? AppTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppTemplate(String name, Output<String> id, @Nullable AppTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:imp/appTemplate:AppTemplate", name, state, makeResourceOptions(options, id));
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
    public static AppTemplate get(String name, Output<String> id, @Nullable AppTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppTemplate(name, id, state, options);
    }
}
