// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud Function Compute Trigger resource. Based on trigger, execute your code in response to events in Alibaba Cloud.
 *  For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/en/fc/developer-reference/api-fc-open-2021-04-06-createtrigger).
 *
 * > **NOTE:** The resource requires a provider field 'account_id'. See account_id.
 *
 * > **NOTE:** Available since v1.93.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const default = alicloud.getAccount({});
 * const defaultGetRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultInteger = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultProject = new alicloud.log.Project("default", {projectName: `example-value-${defaultInteger.result}`});
 * const defaultStore = new alicloud.log.Store("default", {
 *     projectName: defaultProject.name,
 *     logstoreName: "example-value",
 * });
 * const sourceStore = new alicloud.log.Store("source_store", {
 *     projectName: defaultProject.name,
 *     logstoreName: "example-source-store",
 * });
 * const defaultRole = new alicloud.ram.Role("default", {
 *     name: `fcservicerole-${defaultInteger.result}`,
 *     document: `  {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "fc.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *   }
 * `,
 *     description: "this is a example",
 *     force: true,
 * });
 * const defaultRolePolicyAttachment = new alicloud.ram.RolePolicyAttachment("default", {
 *     roleName: defaultRole.name,
 *     policyName: "AliyunLogFullAccess",
 *     policyType: "System",
 * });
 * const defaultService = new alicloud.fc.Service("default", {
 *     name: `example-value-${defaultInteger.result}`,
 *     description: "example-value",
 *     role: defaultRole.arn,
 *     logConfig: {
 *         project: defaultProject.name,
 *         logstore: defaultStore.name,
 *         enableInstanceMetrics: true,
 *         enableRequestMetrics: true,
 *     },
 * });
 * const defaultBucket = new alicloud.oss.Bucket("default", {bucket: `terraform-example-${defaultInteger.result}`});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const defaultBucketObject = new alicloud.oss.BucketObject("default", {
 *     bucket: defaultBucket.id,
 *     key: "index.py",
 *     content: `import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'`,
 * });
 * const defaultFunction = new alicloud.fc.Function("default", {
 *     service: defaultService.name,
 *     name: "terraform-example",
 *     description: "example",
 *     ossBucket: defaultBucket.id,
 *     ossKey: defaultBucketObject.key,
 *     memorySize: 512,
 *     runtime: "python3.10",
 *     handler: "hello.handler",
 * });
 * const defaultTrigger = new alicloud.fc.Trigger("default", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example",
 *     role: defaultRole.arn,
 *     sourceArn: pulumi.all([defaultGetRegions, _default, defaultProject.name]).apply(([defaultGetRegions, _default, name]) => `acs:log:${defaultGetRegions.regions?.[0]?.id}:${_default.id}:project/${name}`),
 *     type: "log",
 *     config: pulumi.interpolate`    {
 *         "sourceConfig": {
 *             "logstore": "${sourceStore.name}",
 *             "startTime": null
 *         },
 *         "jobConfig": {
 *             "maxRetryTime": 3,
 *             "triggerInterval": 60
 *         },
 *         "functionParameter": {
 *             "a": "b",
 *             "c": "d"
 *         },
 *         "logConfig": {
 *              "project": "${defaultProject.name}",
 *             "logstore": "${defaultStore.name}"
 *         },
 *         "targetConfig": null,
 *         "enable": true
 *     }
 *   
 * `,
 * });
 * ```
 *
 * MNS topic trigger:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const default = alicloud.getAccount({});
 * const defaultGetRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultInteger = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultTopic = new alicloud.mns.Topic("default", {name: `example-value-${defaultInteger.result}`});
 * const defaultRole = new alicloud.ram.Role("default", {
 *     name: `fcservicerole-${defaultInteger.result}`,
 *     document: `  {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "mns.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *   }
 * `,
 *     description: "this is a example",
 *     force: true,
 * });
 * const defaultRolePolicyAttachment = new alicloud.ram.RolePolicyAttachment("default", {
 *     roleName: defaultRole.name,
 *     policyName: "AliyunMNSNotificationRolePolicy",
 *     policyType: "System",
 * });
 * const defaultService = new alicloud.fc.Service("default", {
 *     name: `example-value-${defaultInteger.result}`,
 *     description: "example-value",
 *     internetAccess: false,
 * });
 * const defaultBucket = new alicloud.oss.Bucket("default", {bucket: `terraform-example-${defaultInteger.result}`});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const defaultBucketObject = new alicloud.oss.BucketObject("default", {
 *     bucket: defaultBucket.id,
 *     key: "index.py",
 *     content: `import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'`,
 * });
 * const defaultFunction = new alicloud.fc.Function("default", {
 *     service: defaultService.name,
 *     name: `terraform-example-${defaultInteger.result}`,
 *     description: "example",
 *     ossBucket: defaultBucket.id,
 *     ossKey: defaultBucketObject.key,
 *     memorySize: 512,
 *     runtime: "python3.10",
 *     handler: "hello.handler",
 * });
 * const defaultTrigger = new alicloud.fc.Trigger("default", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example",
 *     role: defaultRole.arn,
 *     sourceArn: pulumi.all([defaultGetRegions, _default, defaultTopic.name]).apply(([defaultGetRegions, _default, name]) => `acs:mns:${defaultGetRegions.regions?.[0]?.id}:${_default.id}:/topics/${name}`),
 *     type: "mns_topic",
 *     configMns: `  {
 *     "filterTag":"exampleTag",
 *     "notifyContentFormat":"STREAM",
 *     "notifyStrategy":"BACKOFF_RETRY"
 *   }
 * `,
 * });
 * ```
 *
 * CDN events trigger:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const default = alicloud.getAccount({});
 * const defaultInteger = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultDomainNew = new alicloud.cdn.DomainNew("default", {
 *     domainName: `example${defaultInteger.result}.tf.com`,
 *     cdnType: "web",
 *     scope: "overseas",
 *     sources: [{
 *         content: "1.1.1.1",
 *         type: "ipaddr",
 *         priority: 20,
 *         port: 80,
 *         weight: 10,
 *     }],
 * });
 * const defaultService = new alicloud.fc.Service("default", {
 *     name: `example-value-${defaultInteger.result}`,
 *     description: "example-value",
 *     internetAccess: false,
 * });
 * const defaultRole = new alicloud.ram.Role("default", {
 *     name: `fcservicerole-${defaultInteger.result}`,
 *     document: `    {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "cdn.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *   }
 * `,
 *     description: "this is a example",
 *     force: true,
 * });
 * const defaultPolicy = new alicloud.ram.Policy("default", {
 *     policyName: `fcservicepolicy-${defaultInteger.result}`,
 *     policyDocument: pulumi.interpolate`    {
 *         "Version": "1",
 *         "Statement": [
 *         {
 *             "Action": [
 *             "fc:InvokeFunction"
 *             ],
 *         "Resource": [
 *             "acs:fc:*:*:services/${defaultService.name}/functions/*",
 *             "acs:fc:*:*:services/${defaultService.name}.*&#47;functions/*"
 *         ],
 *         "Effect": "Allow"
 *         }
 *         ]
 *     }
 * `,
 *     description: "this is a example",
 *     force: true,
 * });
 * const defaultRolePolicyAttachment = new alicloud.ram.RolePolicyAttachment("default", {
 *     roleName: defaultRole.name,
 *     policyName: defaultPolicy.policyName,
 *     policyType: "Custom",
 * });
 * const defaultBucket = new alicloud.oss.Bucket("default", {bucket: `terraform-example-${defaultInteger.result}`});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const defaultBucketObject = new alicloud.oss.BucketObject("default", {
 *     bucket: defaultBucket.id,
 *     key: "index.py",
 *     content: `import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'`,
 * });
 * const defaultFunction = new alicloud.fc.Function("default", {
 *     service: defaultService.name,
 *     name: `terraform-example-${defaultInteger.result}`,
 *     description: "example",
 *     ossBucket: defaultBucket.id,
 *     ossKey: defaultBucketObject.key,
 *     memorySize: 512,
 *     runtime: "python3.10",
 *     handler: "hello.handler",
 * });
 * const defaultTrigger = new alicloud.fc.Trigger("default", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example",
 *     role: defaultRole.arn,
 *     sourceArn: _default.then(_default => `acs:cdn:*:${_default.id}`),
 *     type: "cdn_events",
 *     config: pulumi.interpolate`      {"eventName":"LogFileCreated",
 *      "eventVersion":"1.0.0",
 *      "notes":"cdn events trigger",
 *      "filter":{
 *         "domain": ["${defaultDomainNew.domainName}"]
 *         }
 *     }
 * `,
 * });
 * ```
 *
 * EventBridge trigger:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const default = alicloud.getAccount({});
 * const defaultGetRegions = alicloud.getRegions({
 *     current: true,
 * });
 * const defaultInteger = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const serviceLinkedRole = new alicloud.eventbridge.ServiceLinkedRole("service_linked_role", {productName: "AliyunServiceRoleForEventBridgeSendToFC"});
 * const defaultService = new alicloud.fc.Service("default", {
 *     name: `example-value-${defaultInteger.result}`,
 *     description: "example-value",
 *     internetAccess: false,
 * });
 * const defaultBucket = new alicloud.oss.Bucket("default", {bucket: `terraform-example-${defaultInteger.result}`});
 * // If you upload the function by OSS Bucket, you need to specify path can't upload by content.
 * const defaultBucketObject = new alicloud.oss.BucketObject("default", {
 *     bucket: defaultBucket.id,
 *     key: "index.py",
 *     content: `import logging 
 * def handler(event, context): 
 * logger = logging.getLogger() 
 * logger.info('hello world') 
 * return 'hello world'`,
 * });
 * const defaultFunction = new alicloud.fc.Function("default", {
 *     service: defaultService.name,
 *     name: "terraform-example",
 *     description: "example",
 *     ossBucket: defaultBucket.id,
 *     ossKey: defaultBucketObject.key,
 *     memorySize: 512,
 *     runtime: "python3.10",
 *     handler: "hello.handler",
 * });
 * const ossTrigger = new alicloud.fc.Trigger("oss_trigger", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example-oss",
 *     type: "eventbridge",
 *     config: `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": {
 *           "source":[
 *             "acs.oss"
 *             ],
 *             "type":[
 *               "oss:BucketCreated:PutBucket"
 *             ]
 *         },
 *         "eventSourceConfig": {
 *             "eventSourceType": "Default"
 *         }
 *     }
 * `,
 * });
 * const mnsTrigger = new alicloud.fc.Trigger("mns_trigger", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example-mns",
 *     type: "eventbridge",
 *     config: `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "MNS",
 *             "eventSourceParameters": {
 *                 "sourceMNSParameters": {
 *                     "RegionId": "cn-hangzhou",
 *                     "QueueName": "mns-queue",
 *                     "IsBase64Decode": true
 *                 }
 *             }
 *         }
 *     }
 * `,
 * });
 * const defaultInstance = new alicloud.rocketmq.Instance("default", {
 *     instanceName: `terraform-example-${defaultInteger.result}`,
 *     remark: "terraform-example",
 * });
 * const defaultGroup = new alicloud.rocketmq.Group("default", {
 *     groupName: "GID-example",
 *     instanceId: defaultInstance.id,
 *     remark: "terraform-example",
 * });
 * const defaultTopic = new alicloud.rocketmq.Topic("default", {
 *     topicName: "mytopic",
 *     instanceId: defaultInstance.id,
 *     messageType: 0,
 *     remark: "terraform-example",
 * });
 * const rocketmqTrigger = new alicloud.fc.Trigger("rocketmq_trigger", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example-rocketmq",
 *     type: "eventbridge",
 *     config: pulumi.all([defaultGetRegions, defaultInstance.id, defaultGroup.groupName, defaultTopic.topicName]).apply(([defaultGetRegions, id, groupName, topicName]) => `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "RocketMQ",
 *             "eventSourceParameters": {
 *                 "sourceRocketMQParameters": {
 *                     "RegionId": "${defaultGetRegions.regions?.[0]?.id}",
 *                     "InstanceId": "${id}",
 *                     "GroupID": "${groupName}",
 *                     "Topic": "${topicName}",
 *                     "Timestamp": 1686296162,
 *                     "Tag": "example-tag",
 *                     "Offset": "CONSUME_FROM_LAST_OFFSET"
 *                 }
 *             }
 *         }
 *     }
 * `),
 * });
 * const defaultInstance2 = new alicloud.amqp.Instance("default", {
 *     instanceName: `terraform-example-${defaultInteger.result}`,
 *     instanceType: "professional",
 *     maxTps: "1000",
 *     queueCapacity: "50",
 *     supportEip: true,
 *     maxEipTps: "128",
 *     paymentType: "Subscription",
 *     period: 1,
 * });
 * const defaultVirtualHost = new alicloud.amqp.VirtualHost("default", {
 *     instanceId: defaultInstance2.id,
 *     virtualHostName: "example-VirtualHost",
 * });
 * const defaultQueue = new alicloud.amqp.Queue("default", {
 *     instanceId: defaultVirtualHost.instanceId,
 *     queueName: "example-queue",
 *     virtualHostName: defaultVirtualHost.virtualHostName,
 * });
 * const rabbitmqTrigger = new alicloud.fc.Trigger("rabbitmq_trigger", {
 *     service: defaultService.name,
 *     "function": defaultFunction.name,
 *     name: "terraform-example-rabbitmq",
 *     type: "eventbridge",
 *     config: pulumi.all([defaultGetRegions, defaultInstance2.id, defaultVirtualHost.virtualHostName, defaultQueue.queueName]).apply(([defaultGetRegions, id, virtualHostName, queueName]) => `    {
 *         "triggerEnable": false,
 *         "asyncInvocationType": false,
 *         "eventRuleFilterPattern": "{}",
 *         "eventSourceConfig": {
 *             "eventSourceType": "RabbitMQ",
 *             "eventSourceParameters": {
 *                 "sourceRabbitMQParameters": {
 *                     "RegionId": "${defaultGetRegions.regions?.[0]?.id}",
 *                     "InstanceId": "${id}",
 *                     "VirtualHostName": "${virtualHostName}",
 *                     "QueueName": "${queueName}"
 *                 }
 *             }
 *         }
 *     }
 * `),
 * });
 * ```
 *
 * ## Module Support
 *
 * You can use to the existing fc module
 * to create several triggers quickly.
 *
 * ## Import
 *
 * Function Compute trigger can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
 * ```
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    public readonly config!: pulumi.Output<string | undefined>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    public readonly configMns!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute function name.
     */
    public readonly function!: pulumi.Output<string>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute service name.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    public readonly sourceArn!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute trigger ID.
     */
    public /*out*/ readonly triggerId!: pulumi.Output<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
     *
     * > **NOTE:** Config does not support modification when type is mns_topic.
     * > **NOTE:** type = cdn_events, available in 1.47.0+.
     * > **NOTE:** type = eventbridge, available in 1.173.0+.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["configMns"] = state ? state.configMns : undefined;
            resourceInputs["function"] = state ? state.function : undefined;
            resourceInputs["lastModified"] = state ? state.lastModified : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["sourceArn"] = state ? state.sourceArn : undefined;
            resourceInputs["triggerId"] = state ? state.triggerId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if ((!args || args.function === undefined) && !opts.urn) {
                throw new Error("Missing required property 'function'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["configMns"] = args ? args.configMns : undefined;
            resourceInputs["function"] = args ? args.function : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["sourceArn"] = args ? args.sourceArn : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["lastModified"] = undefined /*out*/;
            resourceInputs["triggerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Trigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    config?: pulumi.Input<string>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    configMns?: pulumi.Input<string>;
    /**
     * The Function Compute function name.
     */
    function?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    lastModified?: pulumi.Input<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    role?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service?: pulumi.Input<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The Function Compute trigger ID.
     */
    triggerId?: pulumi.Input<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
     *
     * > **NOTE:** Config does not support modification when type is mns_topic.
     * > **NOTE:** type = cdn_events, available in 1.47.0+.
     * > **NOTE:** type = eventbridge, available in 1.173.0+.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    config?: pulumi.Input<string>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    configMns?: pulumi.Input<string>;
    /**
     * The Function Compute function name.
     */
    function: pulumi.Input<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    role?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service: pulumi.Input<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
     *
     * > **NOTE:** Config does not support modification when type is mns_topic.
     * > **NOTE:** type = cdn_events, available in 1.47.0+.
     * > **NOTE:** type = eventbridge, available in 1.173.0+.
     */
    type: pulumi.Input<string>;
}
