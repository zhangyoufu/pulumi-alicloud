// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eais;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eais.inputs.GetInstancesArgs;
import com.pulumi.alicloud.eais.inputs.GetInstancesPlainArgs;
import com.pulumi.alicloud.eais.outputs.GetInstancesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class EaisFunctions {
    /**
     * This data source provides the Eais Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.137.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * 
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eais.EaisFunctions;
     * import com.pulumi.alicloud.actiontrail.inputs.GetInstancesArgs;
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
     *         final var ids = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .id(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId1&#34;, ids.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *         final var nameRegex = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .nameRegex(&#34;^my-Instance&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId2&#34;, nameRegex.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetInstancesResult> getInstances() {
        return getInstances(GetInstancesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eais Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.137.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * 
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eais.EaisFunctions;
     * import com.pulumi.alicloud.actiontrail.inputs.GetInstancesArgs;
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
     *         final var ids = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .id(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId1&#34;, ids.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *         final var nameRegex = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .nameRegex(&#34;^my-Instance&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId2&#34;, nameRegex.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain() {
        return getInstancesPlain(GetInstancesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eais Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.137.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * 
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eais.EaisFunctions;
     * import com.pulumi.alicloud.actiontrail.inputs.GetInstancesArgs;
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
     *         final var ids = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .id(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId1&#34;, ids.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *         final var nameRegex = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .nameRegex(&#34;^my-Instance&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId2&#34;, nameRegex.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args) {
        return getInstances(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eais Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.137.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * 
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eais.EaisFunctions;
     * import com.pulumi.alicloud.actiontrail.inputs.GetInstancesArgs;
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
     *         final var ids = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .id(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId1&#34;, ids.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *         final var nameRegex = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .nameRegex(&#34;^my-Instance&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId2&#34;, nameRegex.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args) {
        return getInstancesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eais Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.137.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * 
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eais.EaisFunctions;
     * import com.pulumi.alicloud.actiontrail.inputs.GetInstancesArgs;
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
     *         final var ids = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .id(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId1&#34;, ids.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *         final var nameRegex = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .nameRegex(&#34;^my-Instance&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId2&#34;, nameRegex.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eais/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Eais Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.137.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * 
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eais.EaisFunctions;
     * import com.pulumi.alicloud.actiontrail.inputs.GetInstancesArgs;
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
     *         final var ids = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .id(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId1&#34;, ids.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *         final var nameRegex = EaisFunctions.getInstances(GetInstancesArgs.builder()
     *             .nameRegex(&#34;^my-Instance&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eaisInstanceId2&#34;, nameRegex.applyValue(getInstancesResult -&gt; getInstancesResult.instances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eais/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
}
