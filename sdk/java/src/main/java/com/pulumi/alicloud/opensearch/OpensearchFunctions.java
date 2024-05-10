// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.opensearch;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsPlainArgs;
import com.pulumi.alicloud.opensearch.outputs.GetAppGroupsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class OpensearchFunctions {
    /**
     * This data source provides the Open Search App Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.opensearch.AppGroup;
     * import com.pulumi.alicloud.opensearch.AppGroupArgs;
     * import com.pulumi.alicloud.opensearch.inputs.AppGroupQuotaArgs;
     * import com.pulumi.alicloud.opensearch.OpensearchFunctions;
     * import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
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
     *         final var name = config.get("name").orElse("tf_testacc");
     *         var defaultAppGroup = new AppGroup("defaultAppGroup", AppGroupArgs.builder()        
     *             .appGroupName(name)
     *             .paymentType("PayAsYouGo")
     *             .type("standard")
     *             .quota(AppGroupQuotaArgs.builder()
     *                 .docSize(1)
     *                 .computeResource(20)
     *                 .spec("opensearch.share.common")
     *                 .build())
     *             .build());
     * 
     *         final var default = OpensearchFunctions.getAppGroups(GetAppGroupsArgs.builder()
     *             .ids(defaultAppGroup.id())
     *             .build());
     * 
     *         ctx.export("appGroups", default_.applyValue(default_ -> default_.groups()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAppGroupsResult> getAppGroups() {
        return getAppGroups(GetAppGroupsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Open Search App Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.opensearch.AppGroup;
     * import com.pulumi.alicloud.opensearch.AppGroupArgs;
     * import com.pulumi.alicloud.opensearch.inputs.AppGroupQuotaArgs;
     * import com.pulumi.alicloud.opensearch.OpensearchFunctions;
     * import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
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
     *         final var name = config.get("name").orElse("tf_testacc");
     *         var defaultAppGroup = new AppGroup("defaultAppGroup", AppGroupArgs.builder()        
     *             .appGroupName(name)
     *             .paymentType("PayAsYouGo")
     *             .type("standard")
     *             .quota(AppGroupQuotaArgs.builder()
     *                 .docSize(1)
     *                 .computeResource(20)
     *                 .spec("opensearch.share.common")
     *                 .build())
     *             .build());
     * 
     *         final var default = OpensearchFunctions.getAppGroups(GetAppGroupsArgs.builder()
     *             .ids(defaultAppGroup.id())
     *             .build());
     * 
     *         ctx.export("appGroups", default_.applyValue(default_ -> default_.groups()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAppGroupsResult> getAppGroupsPlain() {
        return getAppGroupsPlain(GetAppGroupsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Open Search App Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.opensearch.AppGroup;
     * import com.pulumi.alicloud.opensearch.AppGroupArgs;
     * import com.pulumi.alicloud.opensearch.inputs.AppGroupQuotaArgs;
     * import com.pulumi.alicloud.opensearch.OpensearchFunctions;
     * import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
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
     *         final var name = config.get("name").orElse("tf_testacc");
     *         var defaultAppGroup = new AppGroup("defaultAppGroup", AppGroupArgs.builder()        
     *             .appGroupName(name)
     *             .paymentType("PayAsYouGo")
     *             .type("standard")
     *             .quota(AppGroupQuotaArgs.builder()
     *                 .docSize(1)
     *                 .computeResource(20)
     *                 .spec("opensearch.share.common")
     *                 .build())
     *             .build());
     * 
     *         final var default = OpensearchFunctions.getAppGroups(GetAppGroupsArgs.builder()
     *             .ids(defaultAppGroup.id())
     *             .build());
     * 
     *         ctx.export("appGroups", default_.applyValue(default_ -> default_.groups()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAppGroupsResult> getAppGroups(GetAppGroupsArgs args) {
        return getAppGroups(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Open Search App Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.opensearch.AppGroup;
     * import com.pulumi.alicloud.opensearch.AppGroupArgs;
     * import com.pulumi.alicloud.opensearch.inputs.AppGroupQuotaArgs;
     * import com.pulumi.alicloud.opensearch.OpensearchFunctions;
     * import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
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
     *         final var name = config.get("name").orElse("tf_testacc");
     *         var defaultAppGroup = new AppGroup("defaultAppGroup", AppGroupArgs.builder()        
     *             .appGroupName(name)
     *             .paymentType("PayAsYouGo")
     *             .type("standard")
     *             .quota(AppGroupQuotaArgs.builder()
     *                 .docSize(1)
     *                 .computeResource(20)
     *                 .spec("opensearch.share.common")
     *                 .build())
     *             .build());
     * 
     *         final var default = OpensearchFunctions.getAppGroups(GetAppGroupsArgs.builder()
     *             .ids(defaultAppGroup.id())
     *             .build());
     * 
     *         ctx.export("appGroups", default_.applyValue(default_ -> default_.groups()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAppGroupsResult> getAppGroupsPlain(GetAppGroupsPlainArgs args) {
        return getAppGroupsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Open Search App Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.opensearch.AppGroup;
     * import com.pulumi.alicloud.opensearch.AppGroupArgs;
     * import com.pulumi.alicloud.opensearch.inputs.AppGroupQuotaArgs;
     * import com.pulumi.alicloud.opensearch.OpensearchFunctions;
     * import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
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
     *         final var name = config.get("name").orElse("tf_testacc");
     *         var defaultAppGroup = new AppGroup("defaultAppGroup", AppGroupArgs.builder()        
     *             .appGroupName(name)
     *             .paymentType("PayAsYouGo")
     *             .type("standard")
     *             .quota(AppGroupQuotaArgs.builder()
     *                 .docSize(1)
     *                 .computeResource(20)
     *                 .spec("opensearch.share.common")
     *                 .build())
     *             .build());
     * 
     *         final var default = OpensearchFunctions.getAppGroups(GetAppGroupsArgs.builder()
     *             .ids(defaultAppGroup.id())
     *             .build());
     * 
     *         ctx.export("appGroups", default_.applyValue(default_ -> default_.groups()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAppGroupsResult> getAppGroups(GetAppGroupsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:opensearch/getAppGroups:getAppGroups", TypeShape.of(GetAppGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Open Search App Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.opensearch.AppGroup;
     * import com.pulumi.alicloud.opensearch.AppGroupArgs;
     * import com.pulumi.alicloud.opensearch.inputs.AppGroupQuotaArgs;
     * import com.pulumi.alicloud.opensearch.OpensearchFunctions;
     * import com.pulumi.alicloud.opensearch.inputs.GetAppGroupsArgs;
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
     *         final var name = config.get("name").orElse("tf_testacc");
     *         var defaultAppGroup = new AppGroup("defaultAppGroup", AppGroupArgs.builder()        
     *             .appGroupName(name)
     *             .paymentType("PayAsYouGo")
     *             .type("standard")
     *             .quota(AppGroupQuotaArgs.builder()
     *                 .docSize(1)
     *                 .computeResource(20)
     *                 .spec("opensearch.share.common")
     *                 .build())
     *             .build());
     * 
     *         final var default = OpensearchFunctions.getAppGroups(GetAppGroupsArgs.builder()
     *             .ids(defaultAppGroup.id())
     *             .build());
     * 
     *         ctx.export("appGroups", default_.applyValue(default_ -> default_.groups()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAppGroupsResult> getAppGroupsPlain(GetAppGroupsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:opensearch/getAppGroups:getAppGroups", TypeShape.of(GetAppGroupsResult.class), args, Utilities.withVersion(options));
    }
}
