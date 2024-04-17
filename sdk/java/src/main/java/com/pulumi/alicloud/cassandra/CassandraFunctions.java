// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cassandra.inputs.GetBackupPlansArgs;
import com.pulumi.alicloud.cassandra.inputs.GetBackupPlansPlainArgs;
import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
import com.pulumi.alicloud.cassandra.inputs.GetClustersPlainArgs;
import com.pulumi.alicloud.cassandra.inputs.GetDataCentersArgs;
import com.pulumi.alicloud.cassandra.inputs.GetDataCentersPlainArgs;
import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
import com.pulumi.alicloud.cassandra.inputs.GetZonesPlainArgs;
import com.pulumi.alicloud.cassandra.outputs.GetBackupPlansResult;
import com.pulumi.alicloud.cassandra.outputs.GetClustersResult;
import com.pulumi.alicloud.cassandra.outputs.GetDataCentersResult;
import com.pulumi.alicloud.cassandra.outputs.GetZonesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class CassandraFunctions {
    /**
     * This data source provides the Cassandra Backup Plans of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.128.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetBackupPlansArgs;
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
     *         final var example = CassandraFunctions.getBackupPlans(GetBackupPlansArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstCassandraBackupPlanId&#34;, example.applyValue(getBackupPlansResult -&gt; getBackupPlansResult.plans()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetBackupPlansResult> getBackupPlans(GetBackupPlansArgs args) {
        return getBackupPlans(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Cassandra Backup Plans of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.128.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetBackupPlansArgs;
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
     *         final var example = CassandraFunctions.getBackupPlans(GetBackupPlansArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstCassandraBackupPlanId&#34;, example.applyValue(getBackupPlansResult -&gt; getBackupPlansResult.plans()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetBackupPlansResult> getBackupPlansPlain(GetBackupPlansPlainArgs args) {
        return getBackupPlansPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Cassandra Backup Plans of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.128.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetBackupPlansArgs;
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
     *         final var example = CassandraFunctions.getBackupPlans(GetBackupPlansArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstCassandraBackupPlanId&#34;, example.applyValue(getBackupPlansResult -&gt; getBackupPlansResult.plans()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetBackupPlansResult> getBackupPlans(GetBackupPlansArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cassandra/getBackupPlans:getBackupPlans", TypeShape.of(GetBackupPlansResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Cassandra Backup Plans of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.128.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetBackupPlansArgs;
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
     *         final var example = CassandraFunctions.getBackupPlans(GetBackupPlansArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstCassandraBackupPlanId&#34;, example.applyValue(getBackupPlansResult -&gt; getBackupPlansResult.plans()[0].id()));
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetBackupPlansResult> getBackupPlansPlain(GetBackupPlansPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cassandra/getBackupPlans:getBackupPlans", TypeShape.of(GetBackupPlansResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
     * Filters support regular expression for the cluster name, ids or tags.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
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
     *         final var cassandra = CassandraFunctions.getClusters(GetClustersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClustersResult> getClusters() {
        return getClusters(GetClustersArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
     * Filters support regular expression for the cluster name, ids or tags.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
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
     *         final var cassandra = CassandraFunctions.getClusters(GetClustersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClustersResult> getClustersPlain() {
        return getClustersPlain(GetClustersPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
     * Filters support regular expression for the cluster name, ids or tags.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
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
     *         final var cassandra = CassandraFunctions.getClusters(GetClustersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClustersResult> getClusters(GetClustersArgs args) {
        return getClusters(args, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
     * Filters support regular expression for the cluster name, ids or tags.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
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
     *         final var cassandra = CassandraFunctions.getClusters(GetClustersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClustersResult> getClustersPlain(GetClustersPlainArgs args) {
        return getClustersPlain(args, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
     * Filters support regular expression for the cluster name, ids or tags.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
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
     *         final var cassandra = CassandraFunctions.getClusters(GetClustersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClustersResult> getClusters(GetClustersArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cassandra/getClusters:getClusters", TypeShape.of(GetClustersResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
     * Filters support regular expression for the cluster name, ids or tags.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetClustersArgs;
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
     *         final var cassandra = CassandraFunctions.getClusters(GetClustersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClustersResult> getClustersPlain(GetClustersPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cassandra/getClusters:getClusters", TypeShape.of(GetClustersResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.cassandra.getDataCenters` data source provides a collection of Cassandra Data Centers available in Alicloud account.
     * Filters support regular expression for the cluster name or ids.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetDataCentersArgs;
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
     *         final var cassandra = CassandraFunctions.getDataCenters(GetDataCentersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra_dc&#34;)
     *             .clusterId(&#34;cds-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDataCentersResult> getDataCenters(GetDataCentersArgs args) {
        return getDataCenters(args, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.cassandra.getDataCenters` data source provides a collection of Cassandra Data Centers available in Alicloud account.
     * Filters support regular expression for the cluster name or ids.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetDataCentersArgs;
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
     *         final var cassandra = CassandraFunctions.getDataCenters(GetDataCentersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra_dc&#34;)
     *             .clusterId(&#34;cds-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDataCentersResult> getDataCentersPlain(GetDataCentersPlainArgs args) {
        return getDataCentersPlain(args, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.cassandra.getDataCenters` data source provides a collection of Cassandra Data Centers available in Alicloud account.
     * Filters support regular expression for the cluster name or ids.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetDataCentersArgs;
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
     *         final var cassandra = CassandraFunctions.getDataCenters(GetDataCentersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra_dc&#34;)
     *             .clusterId(&#34;cds-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDataCentersResult> getDataCenters(GetDataCentersArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cassandra/getDataCenters:getDataCenters", TypeShape.of(GetDataCentersResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.cassandra.getDataCenters` data source provides a collection of Cassandra Data Centers available in Alicloud account.
     * Filters support regular expression for the cluster name or ids.
     * 
     * &gt; **NOTE:**  Available in 1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetDataCentersArgs;
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
     *         final var cassandra = CassandraFunctions.getDataCenters(GetDataCentersArgs.builder()
     *             .nameRegex(&#34;tf_testAccCassandra_dc&#34;)
     *             .clusterId(&#34;cds-xxxxx&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDataCentersResult> getDataCentersPlain(GetDataCentersPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cassandra/getDataCenters:getDataCenters", TypeShape.of(GetDataCentersResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides availability zones for Cassandra that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.cassandra.Cluster;
     * import com.pulumi.alicloud.cassandra.ClusterArgs;
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
     *         // Declare the data source
     *         final var zonesIds = CassandraFunctions.getZones();
     * 
     *         // Create an Cassandra cluster with the first matched zone
     *         var cassandra = new Cluster(&#34;cassandra&#34;, ClusterArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetZonesResult> getZones() {
        return getZones(GetZonesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for Cassandra that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.cassandra.Cluster;
     * import com.pulumi.alicloud.cassandra.ClusterArgs;
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
     *         // Declare the data source
     *         final var zonesIds = CassandraFunctions.getZones();
     * 
     *         // Create an Cassandra cluster with the first matched zone
     *         var cassandra = new Cluster(&#34;cassandra&#34;, ClusterArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain() {
        return getZonesPlain(GetZonesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for Cassandra that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.cassandra.Cluster;
     * import com.pulumi.alicloud.cassandra.ClusterArgs;
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
     *         // Declare the data source
     *         final var zonesIds = CassandraFunctions.getZones();
     * 
     *         // Create an Cassandra cluster with the first matched zone
     *         var cassandra = new Cluster(&#34;cassandra&#34;, ClusterArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetZonesResult> getZones(GetZonesArgs args) {
        return getZones(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for Cassandra that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.cassandra.Cluster;
     * import com.pulumi.alicloud.cassandra.ClusterArgs;
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
     *         // Declare the data source
     *         final var zonesIds = CassandraFunctions.getZones();
     * 
     *         // Create an Cassandra cluster with the first matched zone
     *         var cassandra = new Cluster(&#34;cassandra&#34;, ClusterArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain(GetZonesPlainArgs args) {
        return getZonesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for Cassandra that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.cassandra.Cluster;
     * import com.pulumi.alicloud.cassandra.ClusterArgs;
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
     *         // Declare the data source
     *         final var zonesIds = CassandraFunctions.getZones();
     * 
     *         // Create an Cassandra cluster with the first matched zone
     *         var cassandra = new Cluster(&#34;cassandra&#34;, ClusterArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetZonesResult> getZones(GetZonesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:cassandra/getZones:getZones", TypeShape.of(GetZonesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides availability zones for Cassandra that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.88.0+.
     * 
     * &gt; **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
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
     * import com.pulumi.alicloud.cassandra.CassandraFunctions;
     * import com.pulumi.alicloud.cassandra.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.cassandra.Cluster;
     * import com.pulumi.alicloud.cassandra.ClusterArgs;
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
     *         // Declare the data source
     *         final var zonesIds = CassandraFunctions.getZones();
     * 
     *         // Create an Cassandra cluster with the first matched zone
     *         var cassandra = new Cluster(&#34;cassandra&#34;, ClusterArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain(GetZonesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:cassandra/getZones:getZones", TypeShape.of(GetZonesResult.class), args, Utilities.withVersion(options));
    }
}
