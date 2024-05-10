// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
import com.pulumi.alicloud.mns.inputs.GetQueuesPlainArgs;
import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
import com.pulumi.alicloud.mns.inputs.GetServicePlainArgs;
import com.pulumi.alicloud.mns.inputs.GetTopicSubscriptionsArgs;
import com.pulumi.alicloud.mns.inputs.GetTopicSubscriptionsPlainArgs;
import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
import com.pulumi.alicloud.mns.inputs.GetTopicsPlainArgs;
import com.pulumi.alicloud.mns.outputs.GetQueuesResult;
import com.pulumi.alicloud.mns.outputs.GetServiceResult;
import com.pulumi.alicloud.mns.outputs.GetTopicSubscriptionsResult;
import com.pulumi.alicloud.mns.outputs.GetTopicsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class MnsFunctions {
    /**
     * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
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
     *         final var queues = MnsFunctions.getQueues(GetQueuesArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstQueueId", queues.applyValue(getQueuesResult -> getQueuesResult.queues()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetQueuesResult> getQueues() {
        return getQueues(GetQueuesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
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
     *         final var queues = MnsFunctions.getQueues(GetQueuesArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstQueueId", queues.applyValue(getQueuesResult -> getQueuesResult.queues()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetQueuesResult> getQueuesPlain() {
        return getQueuesPlain(GetQueuesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
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
     *         final var queues = MnsFunctions.getQueues(GetQueuesArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstQueueId", queues.applyValue(getQueuesResult -> getQueuesResult.queues()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetQueuesResult> getQueues(GetQueuesArgs args) {
        return getQueues(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
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
     *         final var queues = MnsFunctions.getQueues(GetQueuesArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstQueueId", queues.applyValue(getQueuesResult -> getQueuesResult.queues()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetQueuesResult> getQueuesPlain(GetQueuesPlainArgs args) {
        return getQueuesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
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
     *         final var queues = MnsFunctions.getQueues(GetQueuesArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstQueueId", queues.applyValue(getQueuesResult -> getQueuesResult.queues()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetQueuesResult> getQueues(GetQueuesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:mns/getQueues:getQueues", TypeShape.of(GetQueuesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetQueuesArgs;
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
     *         final var queues = MnsFunctions.getQueues(GetQueuesArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstQueueId", queues.applyValue(getQueuesResult -> getQueuesResult.queues()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetQueuesResult> getQueuesPlain(GetQueuesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:mns/getQueues:getQueues", TypeShape.of(GetQueuesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open MNS service automatically. If the service has been opened, it will return opened.
     * 
     * For information about MNS and how to use it, see [What is MNS](https://www.alibabacloud.com/help/en/product/27412.htm).
     * 
     * &gt; **NOTE:** Available in v1.118.0+
     * 
     * &gt; **NOTE:** The MNS service is not support in the international site.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
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
     *         final var open = MnsFunctions.getService(GetServiceArgs.builder()
     *             .enable("On")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceResult> getService() {
        return getService(GetServiceArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open MNS service automatically. If the service has been opened, it will return opened.
     * 
     * For information about MNS and how to use it, see [What is MNS](https://www.alibabacloud.com/help/en/product/27412.htm).
     * 
     * &gt; **NOTE:** Available in v1.118.0+
     * 
     * &gt; **NOTE:** The MNS service is not support in the international site.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
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
     *         final var open = MnsFunctions.getService(GetServiceArgs.builder()
     *             .enable("On")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain() {
        return getServicePlain(GetServicePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open MNS service automatically. If the service has been opened, it will return opened.
     * 
     * For information about MNS and how to use it, see [What is MNS](https://www.alibabacloud.com/help/en/product/27412.htm).
     * 
     * &gt; **NOTE:** Available in v1.118.0+
     * 
     * &gt; **NOTE:** The MNS service is not support in the international site.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
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
     *         final var open = MnsFunctions.getService(GetServiceArgs.builder()
     *             .enable("On")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args) {
        return getService(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open MNS service automatically. If the service has been opened, it will return opened.
     * 
     * For information about MNS and how to use it, see [What is MNS](https://www.alibabacloud.com/help/en/product/27412.htm).
     * 
     * &gt; **NOTE:** Available in v1.118.0+
     * 
     * &gt; **NOTE:** The MNS service is not support in the international site.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
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
     *         final var open = MnsFunctions.getService(GetServiceArgs.builder()
     *             .enable("On")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args) {
        return getServicePlain(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open MNS service automatically. If the service has been opened, it will return opened.
     * 
     * For information about MNS and how to use it, see [What is MNS](https://www.alibabacloud.com/help/en/product/27412.htm).
     * 
     * &gt; **NOTE:** Available in v1.118.0+
     * 
     * &gt; **NOTE:** The MNS service is not support in the international site.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
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
     *         final var open = MnsFunctions.getService(GetServiceArgs.builder()
     *             .enable("On")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:mns/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open MNS service automatically. If the service has been opened, it will return opened.
     * 
     * For information about MNS and how to use it, see [What is MNS](https://www.alibabacloud.com/help/en/product/27412.htm).
     * 
     * &gt; **NOTE:** Available in v1.118.0+
     * 
     * &gt; **NOTE:** The MNS service is not support in the international site.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetServiceArgs;
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
     *         final var open = MnsFunctions.getService(GetServiceArgs.builder()
     *             .enable("On")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:mns/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides a list of MNS topic subscriptions in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_subscriptions.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicSubscriptionsArgs;
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
     *         final var subscriptions = MnsFunctions.getTopicSubscriptions(GetTopicSubscriptionsArgs.builder()
     *             .topicName("topic_name")
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicSubscriptionId", subscriptions.applyValue(getTopicSubscriptionsResult -> getTopicSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTopicSubscriptionsResult> getTopicSubscriptions(GetTopicSubscriptionsArgs args) {
        return getTopicSubscriptions(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS topic subscriptions in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_subscriptions.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicSubscriptionsArgs;
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
     *         final var subscriptions = MnsFunctions.getTopicSubscriptions(GetTopicSubscriptionsArgs.builder()
     *             .topicName("topic_name")
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicSubscriptionId", subscriptions.applyValue(getTopicSubscriptionsResult -> getTopicSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTopicSubscriptionsResult> getTopicSubscriptionsPlain(GetTopicSubscriptionsPlainArgs args) {
        return getTopicSubscriptionsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS topic subscriptions in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_subscriptions.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicSubscriptionsArgs;
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
     *         final var subscriptions = MnsFunctions.getTopicSubscriptions(GetTopicSubscriptionsArgs.builder()
     *             .topicName("topic_name")
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicSubscriptionId", subscriptions.applyValue(getTopicSubscriptionsResult -> getTopicSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTopicSubscriptionsResult> getTopicSubscriptions(GetTopicSubscriptionsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:mns/getTopicSubscriptions:getTopicSubscriptions", TypeShape.of(GetTopicSubscriptionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides a list of MNS topic subscriptions in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_subscriptions.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicSubscriptionsArgs;
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
     *         final var subscriptions = MnsFunctions.getTopicSubscriptions(GetTopicSubscriptionsArgs.builder()
     *             .topicName("topic_name")
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicSubscriptionId", subscriptions.applyValue(getTopicSubscriptionsResult -> getTopicSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTopicSubscriptionsResult> getTopicSubscriptionsPlain(GetTopicSubscriptionsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:mns/getTopicSubscriptions:getTopicSubscriptions", TypeShape.of(GetTopicSubscriptionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides a list of MNS topics in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_topics.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
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
     *         final var topics = MnsFunctions.getTopics(GetTopicsArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicId", topics.applyValue(getTopicsResult -> getTopicsResult.topics()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTopicsResult> getTopics() {
        return getTopics(GetTopicsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS topics in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_topics.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
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
     *         final var topics = MnsFunctions.getTopics(GetTopicsArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicId", topics.applyValue(getTopicsResult -> getTopicsResult.topics()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTopicsResult> getTopicsPlain() {
        return getTopicsPlain(GetTopicsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS topics in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_topics.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
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
     *         final var topics = MnsFunctions.getTopics(GetTopicsArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicId", topics.applyValue(getTopicsResult -> getTopicsResult.topics()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTopicsResult> getTopics(GetTopicsArgs args) {
        return getTopics(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS topics in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_topics.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
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
     *         final var topics = MnsFunctions.getTopics(GetTopicsArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicId", topics.applyValue(getTopicsResult -> getTopicsResult.topics()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTopicsResult> getTopicsPlain(GetTopicsPlainArgs args) {
        return getTopicsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides a list of MNS topics in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_topics.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
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
     *         final var topics = MnsFunctions.getTopics(GetTopicsArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicId", topics.applyValue(getTopicsResult -> getTopicsResult.topics()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetTopicsResult> getTopics(GetTopicsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:mns/getTopics:getTopics", TypeShape.of(GetTopicsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides a list of MNS topics in an Alibaba Cloud account according to the specified parameters.
     * 
     * &gt; **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_topics.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.mns.MnsFunctions;
     * import com.pulumi.alicloud.mns.inputs.GetTopicsArgs;
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
     *         final var topics = MnsFunctions.getTopics(GetTopicsArgs.builder()
     *             .namePrefix("tf-")
     *             .build());
     * 
     *         ctx.export("firstTopicId", topics.applyValue(getTopicsResult -> getTopicsResult.topics()[0].id()));
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetTopicsResult> getTopicsPlain(GetTopicsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:mns/getTopics:getTopics", TypeShape.of(GetTopicsResult.class), args, Utilities.withVersion(options));
    }
}
