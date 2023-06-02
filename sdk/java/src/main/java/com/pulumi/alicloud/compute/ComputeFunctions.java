// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.compute;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesPlainArgs;
import com.pulumi.alicloud.compute.outputs.GetNestServiceInstancesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class ComputeFunctions {
    /**
     * This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.205.0+.
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
     * import com.pulumi.alicloud.compute.ComputeFunctions;
     * import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
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
     *         final var ids = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId1&#34;, ids.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *         final var nameRegex = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .nameRegex(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId2&#34;, nameRegex.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetNestServiceInstancesResult> getNestServiceInstances() {
        return getNestServiceInstances(GetNestServiceInstancesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.205.0+.
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
     * import com.pulumi.alicloud.compute.ComputeFunctions;
     * import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
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
     *         final var ids = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId1&#34;, ids.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *         final var nameRegex = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .nameRegex(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId2&#34;, nameRegex.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetNestServiceInstancesResult> getNestServiceInstancesPlain() {
        return getNestServiceInstancesPlain(GetNestServiceInstancesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.205.0+.
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
     * import com.pulumi.alicloud.compute.ComputeFunctions;
     * import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
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
     *         final var ids = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId1&#34;, ids.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *         final var nameRegex = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .nameRegex(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId2&#34;, nameRegex.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetNestServiceInstancesResult> getNestServiceInstances(GetNestServiceInstancesArgs args) {
        return getNestServiceInstances(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.205.0+.
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
     * import com.pulumi.alicloud.compute.ComputeFunctions;
     * import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
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
     *         final var ids = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId1&#34;, ids.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *         final var nameRegex = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .nameRegex(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId2&#34;, nameRegex.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetNestServiceInstancesResult> getNestServiceInstancesPlain(GetNestServiceInstancesPlainArgs args) {
        return getNestServiceInstancesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.205.0+.
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
     * import com.pulumi.alicloud.compute.ComputeFunctions;
     * import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
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
     *         final var ids = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId1&#34;, ids.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *         final var nameRegex = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .nameRegex(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId2&#34;, nameRegex.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetNestServiceInstancesResult> getNestServiceInstances(GetNestServiceInstancesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:compute/getNestServiceInstances:getNestServiceInstances", TypeShape.of(GetNestServiceInstancesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.205.0+.
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
     * import com.pulumi.alicloud.compute.ComputeFunctions;
     * import com.pulumi.alicloud.compute.inputs.GetNestServiceInstancesArgs;
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
     *         final var ids = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId1&#34;, ids.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *         final var nameRegex = ComputeFunctions.getNestServiceInstances(GetNestServiceInstancesArgs.builder()
     *             .nameRegex(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheisId2&#34;, nameRegex.applyValue(getNestServiceInstancesResult -&gt; getNestServiceInstancesResult.serviceInstances()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetNestServiceInstancesResult> getNestServiceInstancesPlain(GetNestServiceInstancesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:compute/getNestServiceInstances:getNestServiceInstances", TypeShape.of(GetNestServiceInstancesResult.class), args, Utilities.withVersion(options));
    }
}
