// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cdn.inputs.GetBlockedRegionsArgs;
import com.pulumi.alicloud.cdn.inputs.GetBlockedRegionsPlainArgs;
import com.pulumi.alicloud.cdn.inputs.GetIpInfoArgs;
import com.pulumi.alicloud.cdn.inputs.GetIpInfoPlainArgs;
import com.pulumi.alicloud.cdn.inputs.GetRealTimeLogDeliveriesArgs;
import com.pulumi.alicloud.cdn.inputs.GetRealTimeLogDeliveriesPlainArgs;
import com.pulumi.alicloud.cdn.inputs.GetServiceArgs;
import com.pulumi.alicloud.cdn.inputs.GetServicePlainArgs;
import com.pulumi.alicloud.cdn.outputs.GetBlockedRegionsResult;
import com.pulumi.alicloud.cdn.outputs.GetIpInfoResult;
import com.pulumi.alicloud.cdn.outputs.GetRealTimeLogDeliveriesResult;
import com.pulumi.alicloud.cdn.outputs.GetServiceResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class CdnFunctions {
    /**
     * This data source provides the Cdn blocked regions.
     * 
     * &gt; **NOTE:** Available in v1.173.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetBlockedRegionsArgs;
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
     *         final var example = CdnFunctions.getBlockedRegions(GetBlockedRegionsArgs.builder()
     *             .language(&#34;zh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetBlockedRegionsResult> getBlockedRegions(GetBlockedRegionsArgs args) {
        return getBlockedRegions(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Cdn blocked regions.
     * 
     * &gt; **NOTE:** Available in v1.173.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetBlockedRegionsArgs;
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
     *         final var example = CdnFunctions.getBlockedRegions(GetBlockedRegionsArgs.builder()
     *             .language(&#34;zh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetBlockedRegionsResult> getBlockedRegionsPlain(GetBlockedRegionsPlainArgs args) {
        return getBlockedRegionsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Cdn blocked regions.
     * 
     * &gt; **NOTE:** Available in v1.173.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetBlockedRegionsArgs;
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
     *         final var example = CdnFunctions.getBlockedRegions(GetBlockedRegionsArgs.builder()
     *             .language(&#34;zh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetBlockedRegionsResult> getBlockedRegions(GetBlockedRegionsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cdn/getBlockedRegions:getBlockedRegions", TypeShape.of(GetBlockedRegionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Cdn blocked regions.
     * 
     * &gt; **NOTE:** Available in v1.173.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetBlockedRegionsArgs;
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
     *         final var example = CdnFunctions.getBlockedRegions(GetBlockedRegionsArgs.builder()
     *             .language(&#34;zh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetBlockedRegionsResult> getBlockedRegionsPlain(GetBlockedRegionsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cdn/getBlockedRegions:getBlockedRegions", TypeShape.of(GetBlockedRegionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the function of verifying whether an IP is a CDN node.
     * 
     * &gt; **NOTE:** Available in v1.153.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetIpInfoArgs;
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
     *         final var ipTest = CdnFunctions.getIpInfo(GetIpInfoArgs.builder()
     *             .ip(&#34;114.114.114.114&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetIpInfoResult> getIpInfo(GetIpInfoArgs args) {
        return getIpInfo(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the function of verifying whether an IP is a CDN node.
     * 
     * &gt; **NOTE:** Available in v1.153.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetIpInfoArgs;
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
     *         final var ipTest = CdnFunctions.getIpInfo(GetIpInfoArgs.builder()
     *             .ip(&#34;114.114.114.114&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetIpInfoResult> getIpInfoPlain(GetIpInfoPlainArgs args) {
        return getIpInfoPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the function of verifying whether an IP is a CDN node.
     * 
     * &gt; **NOTE:** Available in v1.153.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetIpInfoArgs;
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
     *         final var ipTest = CdnFunctions.getIpInfo(GetIpInfoArgs.builder()
     *             .ip(&#34;114.114.114.114&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetIpInfoResult> getIpInfo(GetIpInfoArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cdn/getIpInfo:getIpInfo", TypeShape.of(GetIpInfoResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the function of verifying whether an IP is a CDN node.
     * 
     * &gt; **NOTE:** Available in v1.153.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetIpInfoArgs;
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
     *         final var ipTest = CdnFunctions.getIpInfo(GetIpInfoArgs.builder()
     *             .ip(&#34;114.114.114.114&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetIpInfoResult> getIpInfoPlain(GetIpInfoPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cdn/getIpInfo:getIpInfo", TypeShape.of(GetIpInfoResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.134.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetRealTimeLogDeliveriesArgs;
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
     *         final var example = CdnFunctions.getRealTimeLogDeliveries(GetRealTimeLogDeliveriesArgs.builder()
     *             .domain(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;cdnRealTimeLogDelivery1&#34;, example.applyValue(getRealTimeLogDeliveriesResult -&gt; getRealTimeLogDeliveriesResult.deliveries()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetRealTimeLogDeliveriesResult> getRealTimeLogDeliveries(GetRealTimeLogDeliveriesArgs args) {
        return getRealTimeLogDeliveries(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.134.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetRealTimeLogDeliveriesArgs;
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
     *         final var example = CdnFunctions.getRealTimeLogDeliveries(GetRealTimeLogDeliveriesArgs.builder()
     *             .domain(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;cdnRealTimeLogDelivery1&#34;, example.applyValue(getRealTimeLogDeliveriesResult -&gt; getRealTimeLogDeliveriesResult.deliveries()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetRealTimeLogDeliveriesResult> getRealTimeLogDeliveriesPlain(GetRealTimeLogDeliveriesPlainArgs args) {
        return getRealTimeLogDeliveriesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.134.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetRealTimeLogDeliveriesArgs;
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
     *         final var example = CdnFunctions.getRealTimeLogDeliveries(GetRealTimeLogDeliveriesArgs.builder()
     *             .domain(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;cdnRealTimeLogDelivery1&#34;, example.applyValue(getRealTimeLogDeliveriesResult -&gt; getRealTimeLogDeliveriesResult.deliveries()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetRealTimeLogDeliveriesResult> getRealTimeLogDeliveries(GetRealTimeLogDeliveriesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cdn/getRealTimeLogDeliveries:getRealTimeLogDeliveries", TypeShape.of(GetRealTimeLogDeliveriesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.134.0+.
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
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.cdn.inputs.GetRealTimeLogDeliveriesArgs;
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
     *         final var example = CdnFunctions.getRealTimeLogDeliveries(GetRealTimeLogDeliveriesArgs.builder()
     *             .domain(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;cdnRealTimeLogDelivery1&#34;, example.applyValue(getRealTimeLogDeliveriesResult -&gt; getRealTimeLogDeliveriesResult.deliveries()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetRealTimeLogDeliveriesResult> getRealTimeLogDeliveriesPlain(GetRealTimeLogDeliveriesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cdn/getRealTimeLogDeliveries:getRealTimeLogDeliveries", TypeShape.of(GetRealTimeLogDeliveriesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can enable CDN service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about CDN and how to use it, see [What is CDN](https://www.alibabacloud.com/help/product/27099.htm).
     * 
     * &gt; **NOTE:** Available in v1.98.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.apigateway.inputs.GetServiceArgs;
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
     *         final var open = CdnFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .internetChargeType(&#34;PayByTraffic&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetServiceResult> getService() {
        return getService(GetServiceArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can enable CDN service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about CDN and how to use it, see [What is CDN](https://www.alibabacloud.com/help/product/27099.htm).
     * 
     * &gt; **NOTE:** Available in v1.98.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.apigateway.inputs.GetServiceArgs;
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
     *         final var open = CdnFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .internetChargeType(&#34;PayByTraffic&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain() {
        return getServicePlain(GetServicePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can enable CDN service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about CDN and how to use it, see [What is CDN](https://www.alibabacloud.com/help/product/27099.htm).
     * 
     * &gt; **NOTE:** Available in v1.98.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.apigateway.inputs.GetServiceArgs;
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
     *         final var open = CdnFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .internetChargeType(&#34;PayByTraffic&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args) {
        return getService(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can enable CDN service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about CDN and how to use it, see [What is CDN](https://www.alibabacloud.com/help/product/27099.htm).
     * 
     * &gt; **NOTE:** Available in v1.98.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.apigateway.inputs.GetServiceArgs;
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
     *         final var open = CdnFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .internetChargeType(&#34;PayByTraffic&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args) {
        return getServicePlain(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can enable CDN service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about CDN and how to use it, see [What is CDN](https://www.alibabacloud.com/help/product/27099.htm).
     * 
     * &gt; **NOTE:** Available in v1.98.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.apigateway.inputs.GetServiceArgs;
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
     *         final var open = CdnFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .internetChargeType(&#34;PayByTraffic&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cdn/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can enable CDN service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about CDN and how to use it, see [What is CDN](https://www.alibabacloud.com/help/product/27099.htm).
     * 
     * &gt; **NOTE:** Available in v1.98.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.cdn.CdnFunctions;
     * import com.pulumi.alicloud.apigateway.inputs.GetServiceArgs;
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
     *         final var open = CdnFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .internetChargeType(&#34;PayByTraffic&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cdn/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
}
