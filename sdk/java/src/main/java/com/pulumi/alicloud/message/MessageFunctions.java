// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
import com.pulumi.alicloud.message.inputs.GetServiceQueuesPlainArgs;
import com.pulumi.alicloud.message.inputs.GetServiceSubscriptionsArgs;
import com.pulumi.alicloud.message.inputs.GetServiceSubscriptionsPlainArgs;
import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
import com.pulumi.alicloud.message.inputs.GetServiceTopicsPlainArgs;
import com.pulumi.alicloud.message.outputs.GetServiceQueuesResult;
import com.pulumi.alicloud.message.outputs.GetServiceSubscriptionsResult;
import com.pulumi.alicloud.message.outputs.GetServiceTopicsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class MessageFunctions {
    /**
     * This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
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
     *         final var ids = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId1&#34;, ids.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *         final var name = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .queueName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId2&#34;, name.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceQueuesResult> getServiceQueues() {
        return getServiceQueues(GetServiceQueuesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
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
     *         final var ids = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId1&#34;, ids.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *         final var name = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .queueName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId2&#34;, name.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceQueuesResult> getServiceQueuesPlain() {
        return getServiceQueuesPlain(GetServiceQueuesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
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
     *         final var ids = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId1&#34;, ids.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *         final var name = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .queueName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId2&#34;, name.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceQueuesResult> getServiceQueues(GetServiceQueuesArgs args) {
        return getServiceQueues(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
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
     *         final var ids = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId1&#34;, ids.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *         final var name = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .queueName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId2&#34;, name.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceQueuesResult> getServiceQueuesPlain(GetServiceQueuesPlainArgs args) {
        return getServiceQueuesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
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
     *         final var ids = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId1&#34;, ids.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *         final var name = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .queueName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId2&#34;, name.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceQueuesResult> getServiceQueues(GetServiceQueuesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:message/getServiceQueues:getServiceQueues", TypeShape.of(GetServiceQueuesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceQueuesArgs;
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
     *         final var ids = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId1&#34;, ids.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *         final var name = MessageFunctions.getServiceQueues(GetServiceQueuesArgs.builder()
     *             .queueName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;queueId2&#34;, name.applyValue(getServiceQueuesResult -&gt; getServiceQueuesResult.queues()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceQueuesResult> getServiceQueuesPlain(GetServiceQueuesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:message/getServiceQueues:getServiceQueues", TypeShape.of(GetServiceQueuesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceSubscriptionsArgs;
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
     *         final var ids = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId1&#34;, ids.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *         final var name = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId2&#34;, name.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceSubscriptionsResult> getServiceSubscriptions(GetServiceSubscriptionsArgs args) {
        return getServiceSubscriptions(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceSubscriptionsArgs;
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
     *         final var ids = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId1&#34;, ids.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *         final var name = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId2&#34;, name.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceSubscriptionsResult> getServiceSubscriptionsPlain(GetServiceSubscriptionsPlainArgs args) {
        return getServiceSubscriptionsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceSubscriptionsArgs;
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
     *         final var ids = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId1&#34;, ids.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *         final var name = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId2&#34;, name.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceSubscriptionsResult> getServiceSubscriptions(GetServiceSubscriptionsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:message/getServiceSubscriptions:getServiceSubscriptions", TypeShape.of(GetServiceSubscriptionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceSubscriptionsArgs;
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
     *         final var ids = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId1&#34;, ids.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *         final var name = MessageFunctions.getServiceSubscriptions(GetServiceSubscriptionsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;subscriptionId2&#34;, name.applyValue(getServiceSubscriptionsResult -&gt; getServiceSubscriptionsResult.subscriptions()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceSubscriptionsResult> getServiceSubscriptionsPlain(GetServiceSubscriptionsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:message/getServiceSubscriptions:getServiceSubscriptions", TypeShape.of(GetServiceSubscriptionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
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
     *         final var ids = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId1&#34;, ids.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *         final var name = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId2&#34;, name.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceTopicsResult> getServiceTopics() {
        return getServiceTopics(GetServiceTopicsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
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
     *         final var ids = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId1&#34;, ids.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *         final var name = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId2&#34;, name.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceTopicsResult> getServiceTopicsPlain() {
        return getServiceTopicsPlain(GetServiceTopicsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
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
     *         final var ids = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId1&#34;, ids.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *         final var name = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId2&#34;, name.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceTopicsResult> getServiceTopics(GetServiceTopicsArgs args) {
        return getServiceTopics(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
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
     *         final var ids = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId1&#34;, ids.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *         final var name = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId2&#34;, name.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceTopicsResult> getServiceTopicsPlain(GetServiceTopicsPlainArgs args) {
        return getServiceTopicsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
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
     *         final var ids = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId1&#34;, ids.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *         final var name = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId2&#34;, name.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceTopicsResult> getServiceTopics(GetServiceTopicsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:message/getServiceTopics:getServiceTopics", TypeShape.of(GetServiceTopicsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.188.0+.
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
     * import com.pulumi.alicloud.message.MessageFunctions;
     * import com.pulumi.alicloud.message.inputs.GetServiceTopicsArgs;
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
     *         final var ids = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .ids(&#34;example_id&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId1&#34;, ids.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *         final var name = MessageFunctions.getServiceTopics(GetServiceTopicsArgs.builder()
     *             .topicName(&#34;tf-example&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;topicId2&#34;, name.applyValue(getServiceTopicsResult -&gt; getServiceTopicsResult.topics()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceTopicsResult> getServiceTopicsPlain(GetServiceTopicsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:message/getServiceTopics:getServiceTopics", TypeShape.of(GetServiceTopicsResult.class), args, Utilities.withVersion(options));
    }
}
