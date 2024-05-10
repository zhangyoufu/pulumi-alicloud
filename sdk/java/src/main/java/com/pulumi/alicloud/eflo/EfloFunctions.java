// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eflo;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
import com.pulumi.alicloud.eflo.inputs.GetSubnetsPlainArgs;
import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
import com.pulumi.alicloud.eflo.inputs.GetVpdsPlainArgs;
import com.pulumi.alicloud.eflo.outputs.GetSubnetsResult;
import com.pulumi.alicloud.eflo.outputs.GetVpdsResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class EfloFunctions {
    /**
     * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
     * 
     * &gt; **NOTE:** Available in 1.204.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
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
     *         final var default = EfloFunctions.getSubnets(GetSubnetsArgs.builder()
     *             .nameRegex(defaultAlicloudEfloSubnet.name())
     *             .subnetName("SubnetTestForTerraform")
     *             .vpdId(vpdId)
     *             .zoneId(zoneId)
     *             .build());
     * 
     *         ctx.export("alicloudEfloSubnetExampleId", default_.subnets()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSubnetsResult> getSubnets() {
        return getSubnets(GetSubnetsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
     * 
     * &gt; **NOTE:** Available in 1.204.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
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
     *         final var default = EfloFunctions.getSubnets(GetSubnetsArgs.builder()
     *             .nameRegex(defaultAlicloudEfloSubnet.name())
     *             .subnetName("SubnetTestForTerraform")
     *             .vpdId(vpdId)
     *             .zoneId(zoneId)
     *             .build());
     * 
     *         ctx.export("alicloudEfloSubnetExampleId", default_.subnets()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSubnetsResult> getSubnetsPlain() {
        return getSubnetsPlain(GetSubnetsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
     * 
     * &gt; **NOTE:** Available in 1.204.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
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
     *         final var default = EfloFunctions.getSubnets(GetSubnetsArgs.builder()
     *             .nameRegex(defaultAlicloudEfloSubnet.name())
     *             .subnetName("SubnetTestForTerraform")
     *             .vpdId(vpdId)
     *             .zoneId(zoneId)
     *             .build());
     * 
     *         ctx.export("alicloudEfloSubnetExampleId", default_.subnets()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSubnetsResult> getSubnets(GetSubnetsArgs args) {
        return getSubnets(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
     * 
     * &gt; **NOTE:** Available in 1.204.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
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
     *         final var default = EfloFunctions.getSubnets(GetSubnetsArgs.builder()
     *             .nameRegex(defaultAlicloudEfloSubnet.name())
     *             .subnetName("SubnetTestForTerraform")
     *             .vpdId(vpdId)
     *             .zoneId(zoneId)
     *             .build());
     * 
     *         ctx.export("alicloudEfloSubnetExampleId", default_.subnets()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSubnetsResult> getSubnetsPlain(GetSubnetsPlainArgs args) {
        return getSubnetsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
     * 
     * &gt; **NOTE:** Available in 1.204.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
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
     *         final var default = EfloFunctions.getSubnets(GetSubnetsArgs.builder()
     *             .nameRegex(defaultAlicloudEfloSubnet.name())
     *             .subnetName("SubnetTestForTerraform")
     *             .vpdId(vpdId)
     *             .zoneId(zoneId)
     *             .build());
     * 
     *         ctx.export("alicloudEfloSubnetExampleId", default_.subnets()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSubnetsResult> getSubnets(GetSubnetsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eflo/getSubnets:getSubnets", TypeShape.of(GetSubnetsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides Eflo Subnet available to the user.[What is Subnet](https://help.aliyun.com/document_detail/604977.html)
     * 
     * &gt; **NOTE:** Available in 1.204.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetSubnetsArgs;
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
     *         final var default = EfloFunctions.getSubnets(GetSubnetsArgs.builder()
     *             .nameRegex(defaultAlicloudEfloSubnet.name())
     *             .subnetName("SubnetTestForTerraform")
     *             .vpdId(vpdId)
     *             .zoneId(zoneId)
     *             .build());
     * 
     *         ctx.export("alicloudEfloSubnetExampleId", default_.subnets()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSubnetsResult> getSubnetsPlain(GetSubnetsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eflo/getSubnets:getSubnets", TypeShape.of(GetSubnetsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides Eflo Vpd available to the user.[What is Vpd](https://help.aliyun.com/document_detail/604976.html)
     * 
     * &gt; **NOTE:** Available in 1.201.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
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
     *         final var default = EfloFunctions.getVpds(GetVpdsArgs.builder()
     *             .ids(defaultAlicloudEfloVpd.id())
     *             .nameRegex(defaultAlicloudEfloVpd.name())
     *             .vpdName("RMC-Terraform-Test")
     *             .build());
     * 
     *         ctx.export("alicloudEfloVpdExampleId", default_.vpds()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpdsResult> getVpds() {
        return getVpds(GetVpdsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Vpd available to the user.[What is Vpd](https://help.aliyun.com/document_detail/604976.html)
     * 
     * &gt; **NOTE:** Available in 1.201.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
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
     *         final var default = EfloFunctions.getVpds(GetVpdsArgs.builder()
     *             .ids(defaultAlicloudEfloVpd.id())
     *             .nameRegex(defaultAlicloudEfloVpd.name())
     *             .vpdName("RMC-Terraform-Test")
     *             .build());
     * 
     *         ctx.export("alicloudEfloVpdExampleId", default_.vpds()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpdsResult> getVpdsPlain() {
        return getVpdsPlain(GetVpdsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Vpd available to the user.[What is Vpd](https://help.aliyun.com/document_detail/604976.html)
     * 
     * &gt; **NOTE:** Available in 1.201.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
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
     *         final var default = EfloFunctions.getVpds(GetVpdsArgs.builder()
     *             .ids(defaultAlicloudEfloVpd.id())
     *             .nameRegex(defaultAlicloudEfloVpd.name())
     *             .vpdName("RMC-Terraform-Test")
     *             .build());
     * 
     *         ctx.export("alicloudEfloVpdExampleId", default_.vpds()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpdsResult> getVpds(GetVpdsArgs args) {
        return getVpds(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Vpd available to the user.[What is Vpd](https://help.aliyun.com/document_detail/604976.html)
     * 
     * &gt; **NOTE:** Available in 1.201.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
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
     *         final var default = EfloFunctions.getVpds(GetVpdsArgs.builder()
     *             .ids(defaultAlicloudEfloVpd.id())
     *             .nameRegex(defaultAlicloudEfloVpd.name())
     *             .vpdName("RMC-Terraform-Test")
     *             .build());
     * 
     *         ctx.export("alicloudEfloVpdExampleId", default_.vpds()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpdsResult> getVpdsPlain(GetVpdsPlainArgs args) {
        return getVpdsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides Eflo Vpd available to the user.[What is Vpd](https://help.aliyun.com/document_detail/604976.html)
     * 
     * &gt; **NOTE:** Available in 1.201.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
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
     *         final var default = EfloFunctions.getVpds(GetVpdsArgs.builder()
     *             .ids(defaultAlicloudEfloVpd.id())
     *             .nameRegex(defaultAlicloudEfloVpd.name())
     *             .vpdName("RMC-Terraform-Test")
     *             .build());
     * 
     *         ctx.export("alicloudEfloVpdExampleId", default_.vpds()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpdsResult> getVpds(GetVpdsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eflo/getVpds:getVpds", TypeShape.of(GetVpdsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides Eflo Vpd available to the user.[What is Vpd](https://help.aliyun.com/document_detail/604976.html)
     * 
     * &gt; **NOTE:** Available in 1.201.0+
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
     * import com.pulumi.alicloud.eflo.EfloFunctions;
     * import com.pulumi.alicloud.eflo.inputs.GetVpdsArgs;
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
     *         final var default = EfloFunctions.getVpds(GetVpdsArgs.builder()
     *             .ids(defaultAlicloudEfloVpd.id())
     *             .nameRegex(defaultAlicloudEfloVpd.name())
     *             .vpdName("RMC-Terraform-Test")
     *             .build());
     * 
     *         ctx.export("alicloudEfloVpdExampleId", default_.vpds()[0].id());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpdsResult> getVpdsPlain(GetVpdsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eflo/getVpds:getVpds", TypeShape.of(GetVpdsResult.class), args, Utilities.withVersion(options));
    }
}
