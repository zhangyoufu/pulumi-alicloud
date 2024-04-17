// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vod;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
import com.pulumi.alicloud.vod.inputs.GetDomainsPlainArgs;
import com.pulumi.alicloud.vod.outputs.GetDomainsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class VodFunctions {
    /**
     * This data source provides the Vod Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.vod.Domain;
     * import com.pulumi.alicloud.vod.DomainArgs;
     * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
     * import com.pulumi.alicloud.vod.VodFunctions;
     * import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
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
     *         var defaultDomain = new Domain(&#34;defaultDomain&#34;, DomainArgs.builder()        
     *             .domainName(&#34;your_domain_name&#34;)
     *             .scope(&#34;domestic&#34;)
     *             .sources(DomainSourceArgs.builder()
     *                 .sourceType(&#34;domain&#34;)
     *                 .sourceContent(&#34;your_source_content&#34;)
     *                 .sourcePort(&#34;80&#34;)
     *                 .build())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         final var default = VodFunctions.getDomains(GetDomainsArgs.builder()
     *             .ids(defaultDomain.id())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         ctx.export(&#34;vodDomain&#34;, default_.applyValue(default_ -&gt; default_.domains()[0]));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDomainsResult> getDomains() {
        return getDomains(GetDomainsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Vod Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.vod.Domain;
     * import com.pulumi.alicloud.vod.DomainArgs;
     * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
     * import com.pulumi.alicloud.vod.VodFunctions;
     * import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
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
     *         var defaultDomain = new Domain(&#34;defaultDomain&#34;, DomainArgs.builder()        
     *             .domainName(&#34;your_domain_name&#34;)
     *             .scope(&#34;domestic&#34;)
     *             .sources(DomainSourceArgs.builder()
     *                 .sourceType(&#34;domain&#34;)
     *                 .sourceContent(&#34;your_source_content&#34;)
     *                 .sourcePort(&#34;80&#34;)
     *                 .build())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         final var default = VodFunctions.getDomains(GetDomainsArgs.builder()
     *             .ids(defaultDomain.id())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         ctx.export(&#34;vodDomain&#34;, default_.applyValue(default_ -&gt; default_.domains()[0]));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDomainsResult> getDomainsPlain() {
        return getDomainsPlain(GetDomainsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Vod Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.vod.Domain;
     * import com.pulumi.alicloud.vod.DomainArgs;
     * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
     * import com.pulumi.alicloud.vod.VodFunctions;
     * import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
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
     *         var defaultDomain = new Domain(&#34;defaultDomain&#34;, DomainArgs.builder()        
     *             .domainName(&#34;your_domain_name&#34;)
     *             .scope(&#34;domestic&#34;)
     *             .sources(DomainSourceArgs.builder()
     *                 .sourceType(&#34;domain&#34;)
     *                 .sourceContent(&#34;your_source_content&#34;)
     *                 .sourcePort(&#34;80&#34;)
     *                 .build())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         final var default = VodFunctions.getDomains(GetDomainsArgs.builder()
     *             .ids(defaultDomain.id())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         ctx.export(&#34;vodDomain&#34;, default_.applyValue(default_ -&gt; default_.domains()[0]));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDomainsResult> getDomains(GetDomainsArgs args) {
        return getDomains(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Vod Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.vod.Domain;
     * import com.pulumi.alicloud.vod.DomainArgs;
     * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
     * import com.pulumi.alicloud.vod.VodFunctions;
     * import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
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
     *         var defaultDomain = new Domain(&#34;defaultDomain&#34;, DomainArgs.builder()        
     *             .domainName(&#34;your_domain_name&#34;)
     *             .scope(&#34;domestic&#34;)
     *             .sources(DomainSourceArgs.builder()
     *                 .sourceType(&#34;domain&#34;)
     *                 .sourceContent(&#34;your_source_content&#34;)
     *                 .sourcePort(&#34;80&#34;)
     *                 .build())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         final var default = VodFunctions.getDomains(GetDomainsArgs.builder()
     *             .ids(defaultDomain.id())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         ctx.export(&#34;vodDomain&#34;, default_.applyValue(default_ -&gt; default_.domains()[0]));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDomainsResult> getDomainsPlain(GetDomainsPlainArgs args) {
        return getDomainsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Vod Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.vod.Domain;
     * import com.pulumi.alicloud.vod.DomainArgs;
     * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
     * import com.pulumi.alicloud.vod.VodFunctions;
     * import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
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
     *         var defaultDomain = new Domain(&#34;defaultDomain&#34;, DomainArgs.builder()        
     *             .domainName(&#34;your_domain_name&#34;)
     *             .scope(&#34;domestic&#34;)
     *             .sources(DomainSourceArgs.builder()
     *                 .sourceType(&#34;domain&#34;)
     *                 .sourceContent(&#34;your_source_content&#34;)
     *                 .sourcePort(&#34;80&#34;)
     *                 .build())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         final var default = VodFunctions.getDomains(GetDomainsArgs.builder()
     *             .ids(defaultDomain.id())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         ctx.export(&#34;vodDomain&#34;, default_.applyValue(default_ -&gt; default_.domains()[0]));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDomainsResult> getDomains(GetDomainsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:vod/getDomains:getDomains", TypeShape.of(GetDomainsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Vod Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.vod.Domain;
     * import com.pulumi.alicloud.vod.DomainArgs;
     * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
     * import com.pulumi.alicloud.vod.VodFunctions;
     * import com.pulumi.alicloud.vod.inputs.GetDomainsArgs;
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
     *         var defaultDomain = new Domain(&#34;defaultDomain&#34;, DomainArgs.builder()        
     *             .domainName(&#34;your_domain_name&#34;)
     *             .scope(&#34;domestic&#34;)
     *             .sources(DomainSourceArgs.builder()
     *                 .sourceType(&#34;domain&#34;)
     *                 .sourceContent(&#34;your_source_content&#34;)
     *                 .sourcePort(&#34;80&#34;)
     *                 .build())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         final var default = VodFunctions.getDomains(GetDomainsArgs.builder()
     *             .ids(defaultDomain.id())
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;key1&#34;, &#34;value1&#34;),
     *                 Map.entry(&#34;key2&#34;, &#34;value2&#34;)
     *             ))
     *             .build());
     * 
     *         ctx.export(&#34;vodDomain&#34;, default_.applyValue(default_ -&gt; default_.domains()[0]));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDomainsResult> getDomainsPlain(GetDomainsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:vod/getDomains:getDomains", TypeShape.of(GetDomainsResult.class), args, Utilities.withVersion(options));
    }
}
