// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.inputs.GetAlertResourceArgs;
import com.pulumi.alicloud.log.inputs.GetAlertResourcePlainArgs;
import com.pulumi.alicloud.log.inputs.GetProjectsArgs;
import com.pulumi.alicloud.log.inputs.GetProjectsPlainArgs;
import com.pulumi.alicloud.log.inputs.GetServiceArgs;
import com.pulumi.alicloud.log.inputs.GetServicePlainArgs;
import com.pulumi.alicloud.log.inputs.GetStoresArgs;
import com.pulumi.alicloud.log.inputs.GetStoresPlainArgs;
import com.pulumi.alicloud.log.outputs.GetAlertResourceResult;
import com.pulumi.alicloud.log.outputs.GetProjectsResult;
import com.pulumi.alicloud.log.outputs.GetServiceResult;
import com.pulumi.alicloud.log.outputs.GetStoresResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class LogFunctions {
    /**
     * Using this data source can init SLS Alert resources automatically.
     * 
     * For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
     * 
     * &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.219.0`. Please use new resource alicloud_log_alert_resource.
     * 
     * &gt; **NOTE:** Available since v1.161.0.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetAlertResourceArgs;
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
     *         final var exampleUser = LogFunctions.getAlertResource(GetAlertResourceArgs.builder()
     *             .lang(&#34;cn&#34;)
     *             .type(&#34;user&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAlertResourceResult> getAlertResource(GetAlertResourceArgs args) {
        return getAlertResource(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can init SLS Alert resources automatically.
     * 
     * For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
     * 
     * &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.219.0`. Please use new resource alicloud_log_alert_resource.
     * 
     * &gt; **NOTE:** Available since v1.161.0.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetAlertResourceArgs;
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
     *         final var exampleUser = LogFunctions.getAlertResource(GetAlertResourceArgs.builder()
     *             .lang(&#34;cn&#34;)
     *             .type(&#34;user&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAlertResourceResult> getAlertResourcePlain(GetAlertResourcePlainArgs args) {
        return getAlertResourcePlain(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can init SLS Alert resources automatically.
     * 
     * For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
     * 
     * &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.219.0`. Please use new resource alicloud_log_alert_resource.
     * 
     * &gt; **NOTE:** Available since v1.161.0.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetAlertResourceArgs;
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
     *         final var exampleUser = LogFunctions.getAlertResource(GetAlertResourceArgs.builder()
     *             .lang(&#34;cn&#34;)
     *             .type(&#34;user&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAlertResourceResult> getAlertResource(GetAlertResourceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:log/getAlertResource:getAlertResource", TypeShape.of(GetAlertResourceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can init SLS Alert resources automatically.
     * 
     * For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
     * 
     * &gt; **DEPRECATED:**  This resource  has been deprecated from version `1.219.0`. Please use new resource alicloud_log_alert_resource.
     * 
     * &gt; **NOTE:** Available since v1.161.0.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetAlertResourceArgs;
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
     *         final var exampleUser = LogFunctions.getAlertResource(GetAlertResourceArgs.builder()
     *             .lang(&#34;cn&#34;)
     *             .type(&#34;user&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAlertResourceResult> getAlertResourcePlain(GetAlertResourcePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:log/getAlertResource:getAlertResource", TypeShape.of(GetAlertResourceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Log Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
     * 
     */
    public static Output<GetProjectsResult> getProjects() {
        return getProjects(GetProjectsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Log Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
     * 
     */
    public static CompletableFuture<GetProjectsResult> getProjectsPlain() {
        return getProjectsPlain(GetProjectsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Log Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
     * 
     */
    public static Output<GetProjectsResult> getProjects(GetProjectsArgs args) {
        return getProjects(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Log Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
     * 
     */
    public static CompletableFuture<GetProjectsResult> getProjectsPlain(GetProjectsPlainArgs args) {
        return getProjectsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Log Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
     * 
     */
    public static Output<GetProjectsResult> getProjects(GetProjectsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:log/getProjects:getProjects", TypeShape.of(GetProjectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Log Projects of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
     * 
     */
    public static CompletableFuture<GetProjectsResult> getProjectsPlain(GetProjectsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:log/getProjects:getProjects", TypeShape.of(GetProjectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).
     * 
     * &gt; **NOTE:** Available in v1.96.0+
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetServiceArgs;
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
     *         final var open = LogFunctions.getService(GetServiceArgs.builder()
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
     * Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).
     * 
     * &gt; **NOTE:** Available in v1.96.0+
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetServiceArgs;
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
     *         final var open = LogFunctions.getService(GetServiceArgs.builder()
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
     * Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).
     * 
     * &gt; **NOTE:** Available in v1.96.0+
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetServiceArgs;
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
     *         final var open = LogFunctions.getService(GetServiceArgs.builder()
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
     * Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).
     * 
     * &gt; **NOTE:** Available in v1.96.0+
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetServiceArgs;
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
     *         final var open = LogFunctions.getService(GetServiceArgs.builder()
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
     * Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).
     * 
     * &gt; **NOTE:** Available in v1.96.0+
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetServiceArgs;
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
     *         final var open = LogFunctions.getService(GetServiceArgs.builder()
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
        return Deployment.getInstance().invoke("alicloud:log/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can enable Log service automatically. If the service has been enabled, it will return `Opened`.
     * 
     * For information about Log service and how to use it, see [What is Log Service](https://www.alibabacloud.com/help/product/28958.htm).
     * 
     * &gt; **NOTE:** Available in v1.96.0+
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetServiceArgs;
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
     *         final var open = LogFunctions.getService(GetServiceArgs.builder()
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
        return Deployment.getInstance().invokeAsync("alicloud:log/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Log Stores of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetStoresArgs;
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
     *         final var example = LogFunctions.getStores(GetStoresArgs.builder()
     *             .project(&#34;the_project_name&#34;)
     *             .ids(&#34;the_store_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstLogStoreId&#34;, example.applyValue(getStoresResult -&gt; getStoresResult.stores()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetStoresResult> getStores(GetStoresArgs args) {
        return getStores(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Log Stores of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetStoresArgs;
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
     *         final var example = LogFunctions.getStores(GetStoresArgs.builder()
     *             .project(&#34;the_project_name&#34;)
     *             .ids(&#34;the_store_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstLogStoreId&#34;, example.applyValue(getStoresResult -&gt; getStoresResult.stores()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetStoresResult> getStoresPlain(GetStoresPlainArgs args) {
        return getStoresPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Log Stores of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetStoresArgs;
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
     *         final var example = LogFunctions.getStores(GetStoresArgs.builder()
     *             .project(&#34;the_project_name&#34;)
     *             .ids(&#34;the_store_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstLogStoreId&#34;, example.applyValue(getStoresResult -&gt; getStoresResult.stores()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetStoresResult> getStores(GetStoresArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:log/getStores:getStores", TypeShape.of(GetStoresResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Log Stores of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.126.0+.
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
     * import com.pulumi.alicloud.log.LogFunctions;
     * import com.pulumi.alicloud.log.inputs.GetStoresArgs;
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
     *         final var example = LogFunctions.getStores(GetStoresArgs.builder()
     *             .project(&#34;the_project_name&#34;)
     *             .ids(&#34;the_store_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstLogStoreId&#34;, example.applyValue(getStoresResult -&gt; getStoresResult.stores()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetStoresResult> getStoresPlain(GetStoresPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:log/getStores:getStores", TypeShape.of(GetStoresResult.class), args, Utilities.withVersion(options));
    }
}
