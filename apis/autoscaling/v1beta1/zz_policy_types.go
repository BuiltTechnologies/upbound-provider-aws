/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomizedCapacityMetricSpecificationObservation struct {
}

type CustomizedCapacityMetricSpecificationParameters struct {

	// List of up to 10 structures that defines custom capacity metric in predictive scaling policy
	// +kubebuilder:validation:Required
	MetricDataQueries []MetricDataQueriesParameters `json:"metricDataQueries" tf:"metric_data_queries,omitempty"`
}

type CustomizedLoadMetricSpecificationMetricDataQueriesObservation struct {
}

type CustomizedLoadMetricSpecificationMetricDataQueriesParameters struct {

	// Math expression used on the returned metric. You must specify either expression or metric_stat, but not both.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Short name for the metric used in predictive scaling policy.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Human-readable label for this metric or expression.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Structure that defines CloudWatch metric to be used in predictive scaling policy. You must specify either expression or metric_stat, but not both.
	// +kubebuilder:validation:Optional
	MetricStat []MetricDataQueriesMetricStatParameters `json:"metricStat,omitempty" tf:"metric_stat,omitempty"`

	// Boolean that indicates whether to return the timestamps and raw data values of this metric, the default it true
	// +kubebuilder:validation:Optional
	ReturnData *bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

type CustomizedLoadMetricSpecificationObservation struct {
}

type CustomizedLoadMetricSpecificationParameters struct {

	// List of up to 10 structures that defines custom load metric in predictive scaling policy
	// +kubebuilder:validation:Required
	MetricDataQueries []CustomizedLoadMetricSpecificationMetricDataQueriesParameters `json:"metricDataQueries" tf:"metric_data_queries,omitempty"`
}

type CustomizedMetricSpecificationObservation struct {
}

type CustomizedMetricSpecificationParameters struct {

	// Dimensions of the metric.
	// +kubebuilder:validation:Optional
	MetricDimension []MetricDimensionParameters `json:"metricDimension,omitempty" tf:"metric_dimension,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// Statistic of the metric.
	// +kubebuilder:validation:Required
	Statistic *string `json:"statistic" tf:"statistic,omitempty"`

	// Unit of the metrics to return.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedScalingMetricSpecificationMetricDataQueriesMetricStatObservation struct {
}

type CustomizedScalingMetricSpecificationMetricDataQueriesMetricStatParameters struct {

	// Structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
	// +kubebuilder:validation:Required
	Metric []MetricDataQueriesMetricStatMetricParameters `json:"metric" tf:"metric,omitempty"`

	// Statistic of the metrics to return.
	// +kubebuilder:validation:Required
	Stat *string `json:"stat" tf:"stat,omitempty"`

	// Unit of the metrics to return.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedScalingMetricSpecificationMetricDataQueriesObservation struct {
}

type CustomizedScalingMetricSpecificationMetricDataQueriesParameters struct {

	// Math expression used on the returned metric. You must specify either expression or metric_stat, but not both.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Short name for the metric used in predictive scaling policy.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Human-readable label for this metric or expression.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Structure that defines CloudWatch metric to be used in predictive scaling policy. You must specify either expression or metric_stat, but not both.
	// +kubebuilder:validation:Optional
	MetricStat []CustomizedScalingMetricSpecificationMetricDataQueriesMetricStatParameters `json:"metricStat,omitempty" tf:"metric_stat,omitempty"`

	// Boolean that indicates whether to return the timestamps and raw data values of this metric, the default it true
	// +kubebuilder:validation:Optional
	ReturnData *bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

type CustomizedScalingMetricSpecificationObservation struct {
}

type CustomizedScalingMetricSpecificationParameters struct {

	// List of up to 10 structures that defines custom scaling metric in predictive scaling policy
	// +kubebuilder:validation:Required
	MetricDataQueries []CustomizedScalingMetricSpecificationMetricDataQueriesParameters `json:"metricDataQueries" tf:"metric_data_queries,omitempty"`
}

type DimensionsObservation struct {
}

type DimensionsParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Value of the dimension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MetricDataQueriesMetricStatMetricObservation struct {
}

type MetricDataQueriesMetricStatMetricParameters struct {

	// Dimensions of the metric.
	// +kubebuilder:validation:Optional
	Dimensions []MetricStatMetricDimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type MetricDataQueriesMetricStatObservation struct {
}

type MetricDataQueriesMetricStatParameters struct {

	// Structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
	// +kubebuilder:validation:Required
	Metric []MetricStatMetricParameters `json:"metric" tf:"metric,omitempty"`

	// Statistic of the metrics to return.
	// +kubebuilder:validation:Required
	Stat *string `json:"stat" tf:"stat,omitempty"`

	// Unit of the metrics to return.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricDataQueriesObservation struct {
}

type MetricDataQueriesParameters struct {

	// Math expression used on the returned metric. You must specify either expression or metric_stat, but not both.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Short name for the metric used in predictive scaling policy.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Human-readable label for this metric or expression.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Structure that defines CloudWatch metric to be used in predictive scaling policy. You must specify either expression or metric_stat, but not both.
	// +kubebuilder:validation:Optional
	MetricStat []MetricStatParameters `json:"metricStat,omitempty" tf:"metric_stat,omitempty"`

	// Boolean that indicates whether to return the timestamps and raw data values of this metric, the default it true
	// +kubebuilder:validation:Optional
	ReturnData *bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

type MetricDimensionObservation struct {
}

type MetricDimensionParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Value of the dimension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MetricDimensionsObservation struct {
}

type MetricDimensionsParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Value of the dimension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MetricObservation struct {
}

type MetricParameters struct {

	// Dimensions of the metric.
	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type MetricSpecificationObservation struct {
}

type MetricSpecificationParameters struct {

	// Customized capacity metric specification. The field is only valid when you use customized_load_metric_specification
	// +kubebuilder:validation:Optional
	CustomizedCapacityMetricSpecification []CustomizedCapacityMetricSpecificationParameters `json:"customizedCapacityMetricSpecification,omitempty" tf:"customized_capacity_metric_specification,omitempty"`

	// Customized load metric specification.
	// +kubebuilder:validation:Optional
	CustomizedLoadMetricSpecification []CustomizedLoadMetricSpecificationParameters `json:"customizedLoadMetricSpecification,omitempty" tf:"customized_load_metric_specification,omitempty"`

	// Customized scaling metric specification.
	// +kubebuilder:validation:Optional
	CustomizedScalingMetricSpecification []CustomizedScalingMetricSpecificationParameters `json:"customizedScalingMetricSpecification,omitempty" tf:"customized_scaling_metric_specification,omitempty"`

	// Predefined load metric specification.
	// +kubebuilder:validation:Optional
	PredefinedLoadMetricSpecification []PredefinedLoadMetricSpecificationParameters `json:"predefinedLoadMetricSpecification,omitempty" tf:"predefined_load_metric_specification,omitempty"`

	// Metric pair specification from which Amazon EC2 Auto Scaling determines the appropriate scaling metric and load metric to use.
	// +kubebuilder:validation:Optional
	PredefinedMetricPairSpecification []PredefinedMetricPairSpecificationParameters `json:"predefinedMetricPairSpecification,omitempty" tf:"predefined_metric_pair_specification,omitempty"`

	// Predefined scaling metric specification.
	// +kubebuilder:validation:Optional
	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationParameters `json:"predefinedScalingMetricSpecification,omitempty" tf:"predefined_scaling_metric_specification,omitempty"`

	// Target value for the metric.
	// +kubebuilder:validation:Required
	TargetValue *float64 `json:"targetValue" tf:"target_value,omitempty"`
}

type MetricStatMetricDimensionsObservation struct {
}

type MetricStatMetricDimensionsParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Value of the dimension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MetricStatMetricObservation struct {
}

type MetricStatMetricParameters struct {

	// Dimensions of the metric.
	// +kubebuilder:validation:Optional
	Dimensions []MetricDimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type MetricStatObservation struct {
}

type MetricStatParameters struct {

	// Structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
	// +kubebuilder:validation:Required
	Metric []MetricParameters `json:"metric" tf:"metric,omitempty"`

	// Statistic of the metrics to return.
	// +kubebuilder:validation:Required
	Stat *string `json:"stat" tf:"stat,omitempty"`

	// Unit of the metrics to return.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type PolicyObservation struct {

	// ARN assigned by AWS to the scaling policy.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Short name for the metric used in predictive scaling policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyParameters struct {

	// Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity.
	// +kubebuilder:validation:Optional
	AdjustmentType *string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`

	// Name of the autoscaling group.
	// +kubebuilder:validation:Required
	AutoscalingGroupName *string `json:"autoscalingGroupName" tf:"autoscaling_group_name,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	// +kubebuilder:validation:Optional
	Cooldown *float64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// Whether the scaling policy is enabled or disabled. Default: true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	// +kubebuilder:validation:Optional
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup,omitempty"`

	// Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	// +kubebuilder:validation:Optional
	MetricAggregationType *string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type,omitempty"`

	// Minimum value to scale by when adjustment_type is set to PercentChangeInCapacity.
	// +kubebuilder:validation:Optional
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude,omitempty"`

	// Policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
	// +kubebuilder:validation:Optional
	PredictiveScalingConfiguration []PredictiveScalingConfigurationParameters `json:"predictiveScalingConfiguration,omitempty" tf:"predictive_scaling_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Number of instances by which to scale. adjustment_type determines the interpretation of this number (e.g., as an absolute number or as a percentage of the existing Auto Scaling group size). A positive increment adds to the current capacity and a negative value removes from the current capacity.
	// +kubebuilder:validation:Optional
	ScalingAdjustment *float64 `json:"scalingAdjustment,omitempty" tf:"scaling_adjustment,omitempty"`

	// Set of adjustments that manage
	// group scaling. These have the following structure:
	// +kubebuilder:validation:Optional
	StepAdjustment []StepAdjustmentParameters `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`

	// Target tracking policy. These have the following structure:
	// +kubebuilder:validation:Optional
	TargetTrackingConfiguration []TargetTrackingConfigurationParameters `json:"targetTrackingConfiguration,omitempty" tf:"target_tracking_configuration,omitempty"`
}

type PredefinedLoadMetricSpecificationObservation struct {
}

type PredefinedLoadMetricSpecificationParameters struct {

	// Metric type. Valid values are ASGTotalCPUUtilization, ASGTotalNetworkIn, ASGTotalNetworkOut, or ALBTargetGroupRequestCount.
	// +kubebuilder:validation:Required
	PredefinedMetricType *string `json:"predefinedMetricType" tf:"predefined_metric_type,omitempty"`

	// Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group.
	// +kubebuilder:validation:Required
	ResourceLabel *string `json:"resourceLabel" tf:"resource_label,omitempty"`
}

type PredefinedMetricPairSpecificationObservation struct {
}

type PredefinedMetricPairSpecificationParameters struct {

	// Which metrics to use. There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric. For example, if the metric type is ASGCPUUtilization, the Auto Scaling group's total CPU metric is used as the load metric, and the average CPU metric is used for the scaling metric. Valid values are ASGCPUUtilization, ASGNetworkIn, ASGNetworkOut, or ALBRequestCount.
	// +kubebuilder:validation:Required
	PredefinedMetricType *string `json:"predefinedMetricType" tf:"predefined_metric_type,omitempty"`

	// Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group.
	// +kubebuilder:validation:Required
	ResourceLabel *string `json:"resourceLabel" tf:"resource_label,omitempty"`
}

type PredefinedMetricSpecificationObservation struct {
}

type PredefinedMetricSpecificationParameters struct {

	// Describes a scaling metric for a predictive scaling policy. Valid values are ASGAverageCPUUtilization, ASGAverageNetworkIn, ASGAverageNetworkOut, or ALBRequestCountPerTarget.
	// +kubebuilder:validation:Required
	PredefinedMetricType *string `json:"predefinedMetricType" tf:"predefined_metric_type,omitempty"`

	// Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group.
	// +kubebuilder:validation:Optional
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedScalingMetricSpecificationObservation struct {
}

type PredefinedScalingMetricSpecificationParameters struct {

	// Describes a scaling metric for a predictive scaling policy. Valid values are ASGAverageCPUUtilization, ASGAverageNetworkIn, ASGAverageNetworkOut, or ALBRequestCountPerTarget.
	// +kubebuilder:validation:Required
	PredefinedMetricType *string `json:"predefinedMetricType" tf:"predefined_metric_type,omitempty"`

	// Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group.
	// +kubebuilder:validation:Required
	ResourceLabel *string `json:"resourceLabel" tf:"resource_label,omitempty"`
}

type PredictiveScalingConfigurationObservation struct {
}

type PredictiveScalingConfigurationParameters struct {

	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity of the Auto Scaling group. Valid values are HonorMaxCapacity or IncreaseMaxCapacity. Default is HonorMaxCapacity.
	// +kubebuilder:validation:Optional
	MaxCapacityBreachBehavior *string `json:"maxCapacityBreachBehavior,omitempty" tf:"max_capacity_breach_behavior,omitempty"`

	// Size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity. Valid range is 0 to 100. If set to 0, Amazon EC2 Auto Scaling may scale capacity higher than the maximum capacity to equal but not exceed forecast capacity.
	// +kubebuilder:validation:Optional
	MaxCapacityBuffer *string `json:"maxCapacityBuffer,omitempty" tf:"max_capacity_buffer,omitempty"`

	// This structure includes the metrics and target utilization to use for predictive scaling.
	// +kubebuilder:validation:Required
	MetricSpecification []MetricSpecificationParameters `json:"metricSpecification" tf:"metric_specification,omitempty"`

	// Predictive scaling mode. Valid values are ForecastAndScale and ForecastOnly. Default is ForecastOnly.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Amount of time, in seconds, by which the instance launch time can be advanced. Minimum is 0.
	// +kubebuilder:validation:Optional
	SchedulingBufferTime *string `json:"schedulingBufferTime,omitempty" tf:"scheduling_buffer_time,omitempty"`
}

type StepAdjustmentObservation struct {
}

type StepAdjustmentParameters struct {

	// Lower bound for the
	// difference between the alarm threshold and the CloudWatch metric.
	// Without a value, AWS will treat this bound as negative infinity.
	// +kubebuilder:validation:Optional
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`

	// Upper bound for the
	// difference between the alarm threshold and the CloudWatch metric.
	// Without a value, AWS will treat this bound as positive infinity. The upper bound
	// must be greater than the lower bound.
	// +kubebuilder:validation:Optional
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`

	// Number of instances by which to scale. adjustment_type determines the interpretation of this number (e.g., as an absolute number or as a percentage of the existing Auto Scaling group size). A positive increment adds to the current capacity and a negative value removes from the current capacity.
	// +kubebuilder:validation:Required
	ScalingAdjustment *float64 `json:"scalingAdjustment" tf:"scaling_adjustment,omitempty"`
}

type TargetTrackingConfigurationObservation struct {
}

type TargetTrackingConfigurationParameters struct {

	// Customized metric. Conflicts with predefined_metric_specification.
	// +kubebuilder:validation:Optional
	CustomizedMetricSpecification []CustomizedMetricSpecificationParameters `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification,omitempty"`

	// Whether scale in by the target tracking policy is disabled.
	// +kubebuilder:validation:Optional
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// Predefined metric. Conflicts with customized_metric_specification.
	// +kubebuilder:validation:Optional
	PredefinedMetricSpecification []PredefinedMetricSpecificationParameters `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification,omitempty"`

	// Target value for the metric.
	// +kubebuilder:validation:Required
	TargetValue *float64 `json:"targetValue" tf:"target_value,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Provides an AutoScaling Scaling Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
