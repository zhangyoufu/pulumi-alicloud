// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.wafv3;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.wafv3.inputs.GetDomainsArgs;
import com.pulumi.alicloud.wafv3.inputs.GetDomainsPlainArgs;
import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
import com.pulumi.alicloud.wafv3.inputs.GetInstancesPlainArgs;
import com.pulumi.alicloud.wafv3.outputs.GetDomainsResult;
import com.pulumi.alicloud.wafv3.outputs.GetInstancesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class Wafv3Functions {
    /**
     * This data source provides the Wafv3 Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available since v1.200.0.
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
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
     * import com.pulumi.alicloud.wafv3.inputs.GetDomainsArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         final var ids = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId1&#34;, ids.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
     *         final var defaultGetDomains = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .domain(&#34;zctest12.wafqax.top&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId2&#34;, defaultGetDomains.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
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
     * This data source provides the Wafv3 Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available since v1.200.0.
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
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
     * import com.pulumi.alicloud.wafv3.inputs.GetDomainsArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         final var ids = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId1&#34;, ids.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
     *         final var defaultGetDomains = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .domain(&#34;zctest12.wafqax.top&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId2&#34;, defaultGetDomains.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
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
     * This data source provides the Wafv3 Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available since v1.200.0.
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
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
     * import com.pulumi.alicloud.wafv3.inputs.GetDomainsArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         final var ids = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId1&#34;, ids.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
     *         final var defaultGetDomains = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .domain(&#34;zctest12.wafqax.top&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId2&#34;, defaultGetDomains.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDomainsResult> getDomains(GetDomainsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:wafv3/getDomains:getDomains", TypeShape.of(GetDomainsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Wafv3 Domains of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available since v1.200.0.
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
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
     * import com.pulumi.alicloud.wafv3.inputs.GetDomainsArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         final var ids = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId1&#34;, ids.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
     *         final var defaultGetDomains = Wafv3Functions.getDomains(GetDomainsArgs.builder()
     *             .instanceId(default_.ids()[0])
     *             .domain(&#34;zctest12.wafqax.top&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;wafv3DomainsId2&#34;, defaultGetDomains.applyValue(getDomainsResult -&gt; getDomainsResult.domains()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDomainsResult> getDomainsPlain(GetDomainsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:wafv3/getDomains:getDomains", TypeShape.of(GetDomainsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides Wafv3 Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/web-application-firewall/latest/what-is-waf)
     * 
     * &gt; **NOTE:** Available in 1.200.0+
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         ctx.export(&#34;alicloudWafv3InstanceExampleId&#34;, default_.instances()[0].id());
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetInstancesResult> getInstances() {
        return getInstances(GetInstancesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides Wafv3 Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/web-application-firewall/latest/what-is-waf)
     * 
     * &gt; **NOTE:** Available in 1.200.0+
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         ctx.export(&#34;alicloudWafv3InstanceExampleId&#34;, default_.instances()[0].id());
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain() {
        return getInstancesPlain(GetInstancesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides Wafv3 Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/web-application-firewall/latest/what-is-waf)
     * 
     * &gt; **NOTE:** Available in 1.200.0+
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         ctx.export(&#34;alicloudWafv3InstanceExampleId&#34;, default_.instances()[0].id());
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args) {
        return getInstances(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides Wafv3 Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/web-application-firewall/latest/what-is-waf)
     * 
     * &gt; **NOTE:** Available in 1.200.0+
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         ctx.export(&#34;alicloudWafv3InstanceExampleId&#34;, default_.instances()[0].id());
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args) {
        return getInstancesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides Wafv3 Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/web-application-firewall/latest/what-is-waf)
     * 
     * &gt; **NOTE:** Available in 1.200.0+
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         ctx.export(&#34;alicloudWafv3InstanceExampleId&#34;, default_.instances()[0].id());
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:wafv3/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides Wafv3 Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/web-application-firewall/latest/what-is-waf)
     * 
     * &gt; **NOTE:** Available in 1.200.0+
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.wafv3.Wafv3Functions;
     * import com.pulumi.alicloud.wafv3.inputs.GetInstancesArgs;
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
     *         final var default = Wafv3Functions.getInstances();
     * 
     *         ctx.export(&#34;alicloudWafv3InstanceExampleId&#34;, default_.instances()[0].id());
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:wafv3/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
}
