// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.chatbot;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
import com.pulumi.alicloud.chatbot.inputs.GetAgentsPlainArgs;
import com.pulumi.alicloud.chatbot.outputs.GetAgentsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class ChatbotFunctions {
    /**
     * This data source provides the Chatbot Agents of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.203.0+.
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
     * import com.pulumi.alicloud.chatbot.ChatbotFunctions;
     * import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
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
     *         final var nameRegex = ChatbotFunctions.getAgents(GetAgentsArgs.builder()
     *             .nameRegex(&#34;^my-Agent&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;alicloudChatbotAgentsId1&#34;, nameRegex.applyValue(getAgentsResult -&gt; getAgentsResult.agents()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAgentsResult> getAgents() {
        return getAgents(GetAgentsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Chatbot Agents of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.203.0+.
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
     * import com.pulumi.alicloud.chatbot.ChatbotFunctions;
     * import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
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
     *         final var nameRegex = ChatbotFunctions.getAgents(GetAgentsArgs.builder()
     *             .nameRegex(&#34;^my-Agent&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;alicloudChatbotAgentsId1&#34;, nameRegex.applyValue(getAgentsResult -&gt; getAgentsResult.agents()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAgentsResult> getAgentsPlain() {
        return getAgentsPlain(GetAgentsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Chatbot Agents of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.203.0+.
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
     * import com.pulumi.alicloud.chatbot.ChatbotFunctions;
     * import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
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
     *         final var nameRegex = ChatbotFunctions.getAgents(GetAgentsArgs.builder()
     *             .nameRegex(&#34;^my-Agent&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;alicloudChatbotAgentsId1&#34;, nameRegex.applyValue(getAgentsResult -&gt; getAgentsResult.agents()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAgentsResult> getAgents(GetAgentsArgs args) {
        return getAgents(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Chatbot Agents of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.203.0+.
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
     * import com.pulumi.alicloud.chatbot.ChatbotFunctions;
     * import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
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
     *         final var nameRegex = ChatbotFunctions.getAgents(GetAgentsArgs.builder()
     *             .nameRegex(&#34;^my-Agent&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;alicloudChatbotAgentsId1&#34;, nameRegex.applyValue(getAgentsResult -&gt; getAgentsResult.agents()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAgentsResult> getAgentsPlain(GetAgentsPlainArgs args) {
        return getAgentsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Chatbot Agents of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.203.0+.
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
     * import com.pulumi.alicloud.chatbot.ChatbotFunctions;
     * import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
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
     *         final var nameRegex = ChatbotFunctions.getAgents(GetAgentsArgs.builder()
     *             .nameRegex(&#34;^my-Agent&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;alicloudChatbotAgentsId1&#34;, nameRegex.applyValue(getAgentsResult -&gt; getAgentsResult.agents()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAgentsResult> getAgents(GetAgentsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:chatbot/getAgents:getAgents", TypeShape.of(GetAgentsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Chatbot Agents of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.203.0+.
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
     * import com.pulumi.alicloud.chatbot.ChatbotFunctions;
     * import com.pulumi.alicloud.chatbot.inputs.GetAgentsArgs;
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
     *         final var nameRegex = ChatbotFunctions.getAgents(GetAgentsArgs.builder()
     *             .nameRegex(&#34;^my-Agent&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;alicloudChatbotAgentsId1&#34;, nameRegex.applyValue(getAgentsResult -&gt; getAgentsResult.agents()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAgentsResult> getAgentsPlain(GetAgentsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:chatbot/getAgents:getAgents", TypeShape.of(GetAgentsResult.class), args, Utilities.withVersion(options));
    }
}
