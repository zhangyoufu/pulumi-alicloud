// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.videosurveillance;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
import com.pulumi.alicloud.videosurveillance.inputs.GetServicePlainArgs;
import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsPlainArgs;
import com.pulumi.alicloud.videosurveillance.outputs.GetServiceResult;
import com.pulumi.alicloud.videosurveillance.outputs.GetSystemGroupsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class VideosurveillanceFunctions {
    /**
     * Using this data source can open Video Surveillance System service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Video Surveillance System and how to use it, see [What is VS](https://help.aliyun.com/product/108765.html).
     * 
     * &gt; **NOTE:** Available in v1.116.0+
     * 
     * &gt; **NOTE:** The Video Surveillance System service is not support in the international site.
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
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
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
     *         final var open = VideosurveillanceFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceResult> getService() {
        return getService(GetServiceArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Video Surveillance System service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Video Surveillance System and how to use it, see [What is VS](https://help.aliyun.com/product/108765.html).
     * 
     * &gt; **NOTE:** Available in v1.116.0+
     * 
     * &gt; **NOTE:** The Video Surveillance System service is not support in the international site.
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
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
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
     *         final var open = VideosurveillanceFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain() {
        return getServicePlain(GetServicePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Video Surveillance System service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Video Surveillance System and how to use it, see [What is VS](https://help.aliyun.com/product/108765.html).
     * 
     * &gt; **NOTE:** Available in v1.116.0+
     * 
     * &gt; **NOTE:** The Video Surveillance System service is not support in the international site.
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
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
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
     *         final var open = VideosurveillanceFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args) {
        return getService(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Video Surveillance System service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Video Surveillance System and how to use it, see [What is VS](https://help.aliyun.com/product/108765.html).
     * 
     * &gt; **NOTE:** Available in v1.116.0+
     * 
     * &gt; **NOTE:** The Video Surveillance System service is not support in the international site.
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
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
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
     *         final var open = VideosurveillanceFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args) {
        return getServicePlain(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Video Surveillance System service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Video Surveillance System and how to use it, see [What is VS](https://help.aliyun.com/product/108765.html).
     * 
     * &gt; **NOTE:** Available in v1.116.0+
     * 
     * &gt; **NOTE:** The Video Surveillance System service is not support in the international site.
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
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
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
     *         final var open = VideosurveillanceFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:videosurveillance/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open Video Surveillance System service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Video Surveillance System and how to use it, see [What is VS](https://help.aliyun.com/product/108765.html).
     * 
     * &gt; **NOTE:** Available in v1.116.0+
     * 
     * &gt; **NOTE:** The Video Surveillance System service is not support in the international site.
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
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetServiceArgs;
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
     *         final var open = VideosurveillanceFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:videosurveillance/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Video Surveillance System Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.135.0+.
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
     * import com.pulumi.alicloud.videosurveillance.SystemGroup;
     * import com.pulumi.alicloud.videosurveillance.SystemGroupArgs;
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
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
     *         var defaultSystemGroup = new SystemGroup(&#34;defaultSystemGroup&#34;, SystemGroupArgs.builder()        
     *             .groupName(&#34;groupname&#34;)
     *             .inProtocol(&#34;rtmp&#34;)
     *             .outProtocol(&#34;flv&#34;)
     *             .playDomain(&#34;your_plan_domain&#34;)
     *             .pushDomain(&#34;your_push_domain&#34;)
     *             .build());
     * 
     *         final var defaultSystemGroups = VideosurveillanceFunctions.getSystemGroups(GetSystemGroupsArgs.builder()
     *             .ids(defaultSystemGroup.id())
     *             .build());
     * 
     *         ctx.export(&#34;vsGroup&#34;, defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult).applyValue(defaultSystemGroups -&gt; defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult.ids()[0])));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSystemGroupsResult> getSystemGroups() {
        return getSystemGroups(GetSystemGroupsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Video Surveillance System Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.135.0+.
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
     * import com.pulumi.alicloud.videosurveillance.SystemGroup;
     * import com.pulumi.alicloud.videosurveillance.SystemGroupArgs;
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
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
     *         var defaultSystemGroup = new SystemGroup(&#34;defaultSystemGroup&#34;, SystemGroupArgs.builder()        
     *             .groupName(&#34;groupname&#34;)
     *             .inProtocol(&#34;rtmp&#34;)
     *             .outProtocol(&#34;flv&#34;)
     *             .playDomain(&#34;your_plan_domain&#34;)
     *             .pushDomain(&#34;your_push_domain&#34;)
     *             .build());
     * 
     *         final var defaultSystemGroups = VideosurveillanceFunctions.getSystemGroups(GetSystemGroupsArgs.builder()
     *             .ids(defaultSystemGroup.id())
     *             .build());
     * 
     *         ctx.export(&#34;vsGroup&#34;, defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult).applyValue(defaultSystemGroups -&gt; defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult.ids()[0])));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSystemGroupsResult> getSystemGroupsPlain() {
        return getSystemGroupsPlain(GetSystemGroupsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Video Surveillance System Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.135.0+.
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
     * import com.pulumi.alicloud.videosurveillance.SystemGroup;
     * import com.pulumi.alicloud.videosurveillance.SystemGroupArgs;
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
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
     *         var defaultSystemGroup = new SystemGroup(&#34;defaultSystemGroup&#34;, SystemGroupArgs.builder()        
     *             .groupName(&#34;groupname&#34;)
     *             .inProtocol(&#34;rtmp&#34;)
     *             .outProtocol(&#34;flv&#34;)
     *             .playDomain(&#34;your_plan_domain&#34;)
     *             .pushDomain(&#34;your_push_domain&#34;)
     *             .build());
     * 
     *         final var defaultSystemGroups = VideosurveillanceFunctions.getSystemGroups(GetSystemGroupsArgs.builder()
     *             .ids(defaultSystemGroup.id())
     *             .build());
     * 
     *         ctx.export(&#34;vsGroup&#34;, defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult).applyValue(defaultSystemGroups -&gt; defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult.ids()[0])));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSystemGroupsResult> getSystemGroups(GetSystemGroupsArgs args) {
        return getSystemGroups(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Video Surveillance System Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.135.0+.
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
     * import com.pulumi.alicloud.videosurveillance.SystemGroup;
     * import com.pulumi.alicloud.videosurveillance.SystemGroupArgs;
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
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
     *         var defaultSystemGroup = new SystemGroup(&#34;defaultSystemGroup&#34;, SystemGroupArgs.builder()        
     *             .groupName(&#34;groupname&#34;)
     *             .inProtocol(&#34;rtmp&#34;)
     *             .outProtocol(&#34;flv&#34;)
     *             .playDomain(&#34;your_plan_domain&#34;)
     *             .pushDomain(&#34;your_push_domain&#34;)
     *             .build());
     * 
     *         final var defaultSystemGroups = VideosurveillanceFunctions.getSystemGroups(GetSystemGroupsArgs.builder()
     *             .ids(defaultSystemGroup.id())
     *             .build());
     * 
     *         ctx.export(&#34;vsGroup&#34;, defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult).applyValue(defaultSystemGroups -&gt; defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult.ids()[0])));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSystemGroupsResult> getSystemGroupsPlain(GetSystemGroupsPlainArgs args) {
        return getSystemGroupsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Video Surveillance System Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.135.0+.
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
     * import com.pulumi.alicloud.videosurveillance.SystemGroup;
     * import com.pulumi.alicloud.videosurveillance.SystemGroupArgs;
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
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
     *         var defaultSystemGroup = new SystemGroup(&#34;defaultSystemGroup&#34;, SystemGroupArgs.builder()        
     *             .groupName(&#34;groupname&#34;)
     *             .inProtocol(&#34;rtmp&#34;)
     *             .outProtocol(&#34;flv&#34;)
     *             .playDomain(&#34;your_plan_domain&#34;)
     *             .pushDomain(&#34;your_push_domain&#34;)
     *             .build());
     * 
     *         final var defaultSystemGroups = VideosurveillanceFunctions.getSystemGroups(GetSystemGroupsArgs.builder()
     *             .ids(defaultSystemGroup.id())
     *             .build());
     * 
     *         ctx.export(&#34;vsGroup&#34;, defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult).applyValue(defaultSystemGroups -&gt; defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult.ids()[0])));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSystemGroupsResult> getSystemGroups(GetSystemGroupsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:videosurveillance/getSystemGroups:getSystemGroups", TypeShape.of(GetSystemGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Video Surveillance System Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.135.0+.
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
     * import com.pulumi.alicloud.videosurveillance.SystemGroup;
     * import com.pulumi.alicloud.videosurveillance.SystemGroupArgs;
     * import com.pulumi.alicloud.videosurveillance.VideosurveillanceFunctions;
     * import com.pulumi.alicloud.videosurveillance.inputs.GetSystemGroupsArgs;
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
     *         var defaultSystemGroup = new SystemGroup(&#34;defaultSystemGroup&#34;, SystemGroupArgs.builder()        
     *             .groupName(&#34;groupname&#34;)
     *             .inProtocol(&#34;rtmp&#34;)
     *             .outProtocol(&#34;flv&#34;)
     *             .playDomain(&#34;your_plan_domain&#34;)
     *             .pushDomain(&#34;your_push_domain&#34;)
     *             .build());
     * 
     *         final var defaultSystemGroups = VideosurveillanceFunctions.getSystemGroups(GetSystemGroupsArgs.builder()
     *             .ids(defaultSystemGroup.id())
     *             .build());
     * 
     *         ctx.export(&#34;vsGroup&#34;, defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult).applyValue(defaultSystemGroups -&gt; defaultSystemGroups.applyValue(getSystemGroupsResult -&gt; getSystemGroupsResult.ids()[0])));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSystemGroupsResult> getSystemGroupsPlain(GetSystemGroupsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:videosurveillance/getSystemGroups:getSystemGroups", TypeShape.of(GetSystemGroupsResult.class), args, Utilities.withVersion(options));
    }
}
