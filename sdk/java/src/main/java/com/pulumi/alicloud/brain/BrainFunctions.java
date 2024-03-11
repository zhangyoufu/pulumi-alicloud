// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.brain;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.brain.inputs.GetIndustrialPidLoopsArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialPidLoopsPlainArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsPlainArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsPlainArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
import com.pulumi.alicloud.brain.inputs.GetIndustrialSericePlainArgs;
import com.pulumi.alicloud.brain.outputs.GetIndustrialPidLoopsResult;
import com.pulumi.alicloud.brain.outputs.GetIndustrialPidOrganizationsResult;
import com.pulumi.alicloud.brain.outputs.GetIndustrialPidProjectsResult;
import com.pulumi.alicloud.brain.outputs.GetIndustrialSericeResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class BrainFunctions {
    /**
     * This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.117.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidLoopsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidLoops(GetIndustrialPidLoopsArgs.builder()
     *             .pidProjectId(&#34;856c6b8f-ca63-40a4-xxxx-xxxx&#34;)
     *             .ids(&#34;742a3d4e-d8b0-47c8-xxxx-xxxx&#34;)
     *             .nameRegex(&#34;tf-testACC&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidLoopId&#34;, example.applyValue(getIndustrialPidLoopsResult -&gt; getIndustrialPidLoopsResult.loops()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidLoopsResult> getIndustrialPidLoops(GetIndustrialPidLoopsArgs args) {
        return getIndustrialPidLoops(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.117.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidLoopsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidLoops(GetIndustrialPidLoopsArgs.builder()
     *             .pidProjectId(&#34;856c6b8f-ca63-40a4-xxxx-xxxx&#34;)
     *             .ids(&#34;742a3d4e-d8b0-47c8-xxxx-xxxx&#34;)
     *             .nameRegex(&#34;tf-testACC&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidLoopId&#34;, example.applyValue(getIndustrialPidLoopsResult -&gt; getIndustrialPidLoopsResult.loops()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidLoopsResult> getIndustrialPidLoopsPlain(GetIndustrialPidLoopsPlainArgs args) {
        return getIndustrialPidLoopsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.117.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidLoopsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidLoops(GetIndustrialPidLoopsArgs.builder()
     *             .pidProjectId(&#34;856c6b8f-ca63-40a4-xxxx-xxxx&#34;)
     *             .ids(&#34;742a3d4e-d8b0-47c8-xxxx-xxxx&#34;)
     *             .nameRegex(&#34;tf-testACC&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidLoopId&#34;, example.applyValue(getIndustrialPidLoopsResult -&gt; getIndustrialPidLoopsResult.loops()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidLoopsResult> getIndustrialPidLoops(GetIndustrialPidLoopsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:brain/getIndustrialPidLoops:getIndustrialPidLoops", TypeShape.of(GetIndustrialPidLoopsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.117.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidLoopsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidLoops(GetIndustrialPidLoopsArgs.builder()
     *             .pidProjectId(&#34;856c6b8f-ca63-40a4-xxxx-xxxx&#34;)
     *             .ids(&#34;742a3d4e-d8b0-47c8-xxxx-xxxx&#34;)
     *             .nameRegex(&#34;tf-testACC&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidLoopId&#34;, example.applyValue(getIndustrialPidLoopsResult -&gt; getIndustrialPidLoopsResult.loops()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidLoopsResult> getIndustrialPidLoopsPlain(GetIndustrialPidLoopsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:brain/getIndustrialPidLoops:getIndustrialPidLoops", TypeShape.of(GetIndustrialPidLoopsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Brain Industrial Pid Organizations of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidOrganizationId&#34;, example.applyValue(getIndustrialPidOrganizationsResult -&gt; getIndustrialPidOrganizationsResult.organizations()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidOrganizationsResult> getIndustrialPidOrganizations() {
        return getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Organizations of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidOrganizationId&#34;, example.applyValue(getIndustrialPidOrganizationsResult -&gt; getIndustrialPidOrganizationsResult.organizations()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidOrganizationsResult> getIndustrialPidOrganizationsPlain() {
        return getIndustrialPidOrganizationsPlain(GetIndustrialPidOrganizationsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Organizations of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidOrganizationId&#34;, example.applyValue(getIndustrialPidOrganizationsResult -&gt; getIndustrialPidOrganizationsResult.organizations()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidOrganizationsResult> getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs args) {
        return getIndustrialPidOrganizations(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Organizations of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidOrganizationId&#34;, example.applyValue(getIndustrialPidOrganizationsResult -&gt; getIndustrialPidOrganizationsResult.organizations()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidOrganizationsResult> getIndustrialPidOrganizationsPlain(GetIndustrialPidOrganizationsPlainArgs args) {
        return getIndustrialPidOrganizationsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Organizations of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidOrganizationId&#34;, example.applyValue(getIndustrialPidOrganizationsResult -&gt; getIndustrialPidOrganizationsResult.organizations()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidOrganizationsResult> getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:brain/getIndustrialPidOrganizations:getIndustrialPidOrganizations", TypeShape.of(GetIndustrialPidOrganizationsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Brain Industrial Pid Organizations of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidOrganizationsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidOrganizations(GetIndustrialPidOrganizationsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidOrganizationId&#34;, example.applyValue(getIndustrialPidOrganizationsResult -&gt; getIndustrialPidOrganizationsResult.organizations()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidOrganizationsResult> getIndustrialPidOrganizationsPlain(GetIndustrialPidOrganizationsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:brain/getIndustrialPidOrganizations:getIndustrialPidOrganizations", TypeShape.of(GetIndustrialPidOrganizationsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Brain Industrial Pid Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidProjects(GetIndustrialPidProjectsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidProjectId&#34;, example.applyValue(getIndustrialPidProjectsResult -&gt; getIndustrialPidProjectsResult.projects()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidProjectsResult> getIndustrialPidProjects() {
        return getIndustrialPidProjects(GetIndustrialPidProjectsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidProjects(GetIndustrialPidProjectsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidProjectId&#34;, example.applyValue(getIndustrialPidProjectsResult -&gt; getIndustrialPidProjectsResult.projects()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidProjectsResult> getIndustrialPidProjectsPlain() {
        return getIndustrialPidProjectsPlain(GetIndustrialPidProjectsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidProjects(GetIndustrialPidProjectsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidProjectId&#34;, example.applyValue(getIndustrialPidProjectsResult -&gt; getIndustrialPidProjectsResult.projects()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidProjectsResult> getIndustrialPidProjects(GetIndustrialPidProjectsArgs args) {
        return getIndustrialPidProjects(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidProjects(GetIndustrialPidProjectsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidProjectId&#34;, example.applyValue(getIndustrialPidProjectsResult -&gt; getIndustrialPidProjectsResult.projects()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidProjectsResult> getIndustrialPidProjectsPlain(GetIndustrialPidProjectsPlainArgs args) {
        return getIndustrialPidProjectsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Brain Industrial Pid Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidProjects(GetIndustrialPidProjectsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidProjectId&#34;, example.applyValue(getIndustrialPidProjectsResult -&gt; getIndustrialPidProjectsResult.projects()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialPidProjectsResult> getIndustrialPidProjects(GetIndustrialPidProjectsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:brain/getIndustrialPidProjects:getIndustrialPidProjects", TypeShape.of(GetIndustrialPidProjectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Brain Industrial Pid Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.113.0+.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialPidProjectsArgs;
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
     *         final var example = BrainFunctions.getIndustrialPidProjects(GetIndustrialPidProjectsArgs.builder()
     *             .ids(&#34;3e74e684-cbb5-xxxx&#34;)
     *             .nameRegex(&#34;tf-testAcc&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstBrainIndustrialPidProjectId&#34;, example.applyValue(getIndustrialPidProjectsResult -&gt; getIndustrialPidProjectsResult.projects()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialPidProjectsResult> getIndustrialPidProjectsPlain(GetIndustrialPidProjectsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:brain/getIndustrialPidProjects:getIndustrialPidProjects", TypeShape.of(GetIndustrialPidProjectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
     * 
     * &gt; **NOTE:** Available in v1.115.0+
     * 
     * &gt; **NOTE:** The Brain Industrial service is not support in the international site.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
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
     *         final var open = BrainFunctions.getIndustrialSerice(GetIndustrialSericeArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialSericeResult> getIndustrialSerice() {
        return getIndustrialSerice(GetIndustrialSericeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
     * 
     * &gt; **NOTE:** Available in v1.115.0+
     * 
     * &gt; **NOTE:** The Brain Industrial service is not support in the international site.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
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
     *         final var open = BrainFunctions.getIndustrialSerice(GetIndustrialSericeArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialSericeResult> getIndustrialSericePlain() {
        return getIndustrialSericePlain(GetIndustrialSericePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
     * 
     * &gt; **NOTE:** Available in v1.115.0+
     * 
     * &gt; **NOTE:** The Brain Industrial service is not support in the international site.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
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
     *         final var open = BrainFunctions.getIndustrialSerice(GetIndustrialSericeArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialSericeResult> getIndustrialSerice(GetIndustrialSericeArgs args) {
        return getIndustrialSerice(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
     * 
     * &gt; **NOTE:** Available in v1.115.0+
     * 
     * &gt; **NOTE:** The Brain Industrial service is not support in the international site.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
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
     *         final var open = BrainFunctions.getIndustrialSerice(GetIndustrialSericeArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialSericeResult> getIndustrialSericePlain(GetIndustrialSericePlainArgs args) {
        return getIndustrialSericePlain(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
     * 
     * &gt; **NOTE:** Available in v1.115.0+
     * 
     * &gt; **NOTE:** The Brain Industrial service is not support in the international site.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
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
     *         final var open = BrainFunctions.getIndustrialSerice(GetIndustrialSericeArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetIndustrialSericeResult> getIndustrialSerice(GetIndustrialSericeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:brain/getIndustrialSerice:getIndustrialSerice", TypeShape.of(GetIndustrialSericeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
     * 
     * &gt; **NOTE:** Available in v1.115.0+
     * 
     * &gt; **NOTE:** The Brain Industrial service is not support in the international site.
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
     * import com.pulumi.alicloud.brain.BrainFunctions;
     * import com.pulumi.alicloud.brain.inputs.GetIndustrialSericeArgs;
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
     *         final var open = BrainFunctions.getIndustrialSerice(GetIndustrialSericeArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetIndustrialSericeResult> getIndustrialSericePlain(GetIndustrialSericePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:brain/getIndustrialSerice:getIndustrialSerice", TypeShape.of(GetIndustrialSericeResult.class), args, Utilities.withVersion(options));
    }
}
