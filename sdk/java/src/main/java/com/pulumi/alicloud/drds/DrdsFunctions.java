// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.drds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
import com.pulumi.alicloud.drds.inputs.GetInstancesPlainArgs;
import com.pulumi.alicloud.drds.outputs.GetInstancesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class DrdsFunctions {
    /**
     * The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
     * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
     * 
     * &gt; **NOTE:** Available in 1.35.0+.
     * 
     * ## Example Usage
     * 
     *  &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.drds.DrdsFunctions;
     * import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
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
     *         final var drdsInstancesDs = DrdsFunctions.getInstances(GetInstancesArgs.builder()
     *             .ids(&#34;drdsabc123456&#34;)
     *             .nameRegex(&#34;drds-\\d+&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstDbInstanceId&#34;, drdsInstancesDs.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
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
     * The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
     * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
     * 
     * &gt; **NOTE:** Available in 1.35.0+.
     * 
     * ## Example Usage
     * 
     *  &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.drds.DrdsFunctions;
     * import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
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
     *         final var drdsInstancesDs = DrdsFunctions.getInstances(GetInstancesArgs.builder()
     *             .ids(&#34;drdsabc123456&#34;)
     *             .nameRegex(&#34;drds-\\d+&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstDbInstanceId&#34;, drdsInstancesDs.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
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
     * The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
     * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
     * 
     * &gt; **NOTE:** Available in 1.35.0+.
     * 
     * ## Example Usage
     * 
     *  &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.drds.DrdsFunctions;
     * import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
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
     *         final var drdsInstancesDs = DrdsFunctions.getInstances(GetInstancesArgs.builder()
     *             .ids(&#34;drdsabc123456&#34;)
     *             .nameRegex(&#34;drds-\\d+&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstDbInstanceId&#34;, drdsInstancesDs.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
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
     * The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
     * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
     * 
     * &gt; **NOTE:** Available in 1.35.0+.
     * 
     * ## Example Usage
     * 
     *  &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.drds.DrdsFunctions;
     * import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
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
     *         final var drdsInstancesDs = DrdsFunctions.getInstances(GetInstancesArgs.builder()
     *             .ids(&#34;drdsabc123456&#34;)
     *             .nameRegex(&#34;drds-\\d+&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstDbInstanceId&#34;, drdsInstancesDs.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
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
     * The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
     * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
     * 
     * &gt; **NOTE:** Available in 1.35.0+.
     * 
     * ## Example Usage
     * 
     *  &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.drds.DrdsFunctions;
     * import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
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
     *         final var drdsInstancesDs = DrdsFunctions.getInstances(GetInstancesArgs.builder()
     *             .ids(&#34;drdsabc123456&#34;)
     *             .nameRegex(&#34;drds-\\d+&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstDbInstanceId&#34;, drdsInstancesDs.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:drds/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
     * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
     * 
     * &gt; **NOTE:** Available in 1.35.0+.
     * 
     * ## Example Usage
     * 
     *  &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.drds.DrdsFunctions;
     * import com.pulumi.alicloud.drds.inputs.GetInstancesArgs;
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
     *         final var drdsInstancesDs = DrdsFunctions.getInstances(GetInstancesArgs.builder()
     *             .ids(&#34;drdsabc123456&#34;)
     *             .nameRegex(&#34;drds-\\d+&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstDbInstanceId&#34;, drdsInstancesDs.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:drds/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
}
