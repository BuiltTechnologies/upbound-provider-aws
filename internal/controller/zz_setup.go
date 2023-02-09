/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	analyzer "github.com/upbound/provider-aws/internal/controller/accessanalyzer/analyzer"
	alternatecontact "github.com/upbound/provider-aws/internal/controller/account/alternatecontact"
	app "github.com/upbound/provider-aws/internal/controller/amplify/app"
	backendenvironment "github.com/upbound/provider-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/upbound/provider-aws/internal/controller/amplify/branch"
	webhook "github.com/upbound/provider-aws/internal/controller/amplify/webhook"
	account "github.com/upbound/provider-aws/internal/controller/apigateway/account"
	apikey "github.com/upbound/provider-aws/internal/controller/apigateway/apikey"
	authorizer "github.com/upbound/provider-aws/internal/controller/apigateway/authorizer"
	basepathmapping "github.com/upbound/provider-aws/internal/controller/apigateway/basepathmapping"
	clientcertificate "github.com/upbound/provider-aws/internal/controller/apigateway/clientcertificate"
	deployment "github.com/upbound/provider-aws/internal/controller/apigateway/deployment"
	documentationpart "github.com/upbound/provider-aws/internal/controller/apigateway/documentationpart"
	documentationversion "github.com/upbound/provider-aws/internal/controller/apigateway/documentationversion"
	domainname "github.com/upbound/provider-aws/internal/controller/apigateway/domainname"
	gatewayresponse "github.com/upbound/provider-aws/internal/controller/apigateway/gatewayresponse"
	integration "github.com/upbound/provider-aws/internal/controller/apigateway/integration"
	integrationresponse "github.com/upbound/provider-aws/internal/controller/apigateway/integrationresponse"
	method "github.com/upbound/provider-aws/internal/controller/apigateway/method"
	methodresponse "github.com/upbound/provider-aws/internal/controller/apigateway/methodresponse"
	methodsettings "github.com/upbound/provider-aws/internal/controller/apigateway/methodsettings"
	model "github.com/upbound/provider-aws/internal/controller/apigateway/model"
	requestvalidator "github.com/upbound/provider-aws/internal/controller/apigateway/requestvalidator"
	resource "github.com/upbound/provider-aws/internal/controller/apigateway/resource"
	restapi "github.com/upbound/provider-aws/internal/controller/apigateway/restapi"
	restapipolicy "github.com/upbound/provider-aws/internal/controller/apigateway/restapipolicy"
	stage "github.com/upbound/provider-aws/internal/controller/apigateway/stage"
	usageplan "github.com/upbound/provider-aws/internal/controller/apigateway/usageplan"
	usageplankey "github.com/upbound/provider-aws/internal/controller/apigateway/usageplankey"
	vpclink "github.com/upbound/provider-aws/internal/controller/apigateway/vpclink"
	policy "github.com/upbound/provider-aws/internal/controller/appautoscaling/policy"
	scheduledaction "github.com/upbound/provider-aws/internal/controller/appautoscaling/scheduledaction"
	target "github.com/upbound/provider-aws/internal/controller/appautoscaling/target"
	application "github.com/upbound/provider-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/upbound/provider-aws/internal/controller/appconfig/configurationprofile"
	deploymentappconfig "github.com/upbound/provider-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/upbound/provider-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/upbound/provider-aws/internal/controller/appconfig/environment"
	hostedconfigurationversion "github.com/upbound/provider-aws/internal/controller/appconfig/hostedconfigurationversion"
	flow "github.com/upbound/provider-aws/internal/controller/appflow/flow"
	eventintegration "github.com/upbound/provider-aws/internal/controller/appintegrations/eventintegration"
	gatewayroute "github.com/upbound/provider-aws/internal/controller/appmesh/gatewayroute"
	mesh "github.com/upbound/provider-aws/internal/controller/appmesh/mesh"
	route "github.com/upbound/provider-aws/internal/controller/appmesh/route"
	virtualgateway "github.com/upbound/provider-aws/internal/controller/appmesh/virtualgateway"
	virtualnode "github.com/upbound/provider-aws/internal/controller/appmesh/virtualnode"
	virtualrouter "github.com/upbound/provider-aws/internal/controller/appmesh/virtualrouter"
	virtualservice "github.com/upbound/provider-aws/internal/controller/appmesh/virtualservice"
	autoscalingconfigurationversion "github.com/upbound/provider-aws/internal/controller/apprunner/autoscalingconfigurationversion"
	connection "github.com/upbound/provider-aws/internal/controller/apprunner/connection"
	service "github.com/upbound/provider-aws/internal/controller/apprunner/service"
	vpcconnector "github.com/upbound/provider-aws/internal/controller/apprunner/vpcconnector"
	directoryconfig "github.com/upbound/provider-aws/internal/controller/appstream/directoryconfig"
	fleet "github.com/upbound/provider-aws/internal/controller/appstream/fleet"
	fleetstackassociation "github.com/upbound/provider-aws/internal/controller/appstream/fleetstackassociation"
	imagebuilder "github.com/upbound/provider-aws/internal/controller/appstream/imagebuilder"
	stack "github.com/upbound/provider-aws/internal/controller/appstream/stack"
	user "github.com/upbound/provider-aws/internal/controller/appstream/user"
	userstackassociation "github.com/upbound/provider-aws/internal/controller/appstream/userstackassociation"
	apicache "github.com/upbound/provider-aws/internal/controller/appsync/apicache"
	apikeyappsync "github.com/upbound/provider-aws/internal/controller/appsync/apikey"
	datasource "github.com/upbound/provider-aws/internal/controller/appsync/datasource"
	function "github.com/upbound/provider-aws/internal/controller/appsync/function"
	graphqlapi "github.com/upbound/provider-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/upbound/provider-aws/internal/controller/appsync/resolver"
	database "github.com/upbound/provider-aws/internal/controller/athena/database"
	datacatalog "github.com/upbound/provider-aws/internal/controller/athena/datacatalog"
	namedquery "github.com/upbound/provider-aws/internal/controller/athena/namedquery"
	workgroup "github.com/upbound/provider-aws/internal/controller/athena/workgroup"
	grouptag "github.com/upbound/provider-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/upbound/provider-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/upbound/provider-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/upbound/provider-aws/internal/controller/autoscaling/notification"
	policyautoscaling "github.com/upbound/provider-aws/internal/controller/autoscaling/policy"
	schedule "github.com/upbound/provider-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/upbound/provider-aws/internal/controller/autoscalingplans/scalingplan"
	schedulingpolicy "github.com/upbound/provider-aws/internal/controller/batch/schedulingpolicy"
	budget "github.com/upbound/provider-aws/internal/controller/budgets/budget"
	budgetaction "github.com/upbound/provider-aws/internal/controller/budgets/budgetaction"
	voiceconnector "github.com/upbound/provider-aws/internal/controller/chime/voiceconnector"
	voiceconnectorgroup "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorterminationcredentials"
	environmentec2 "github.com/upbound/provider-aws/internal/controller/cloud9/environmentec2"
	environmentmembership "github.com/upbound/provider-aws/internal/controller/cloud9/environmentmembership"
	resourcecloudcontrol "github.com/upbound/provider-aws/internal/controller/cloudcontrol/resource"
	stackcloudformation "github.com/upbound/provider-aws/internal/controller/cloudformation/stack"
	stackset "github.com/upbound/provider-aws/internal/controller/cloudformation/stackset"
	domain "github.com/upbound/provider-aws/internal/controller/cloudsearch/domain"
	domainserviceaccesspolicy "github.com/upbound/provider-aws/internal/controller/cloudsearch/domainserviceaccesspolicy"
	eventdatastore "github.com/upbound/provider-aws/internal/controller/cloudtrail/eventdatastore"
	trail "github.com/upbound/provider-aws/internal/controller/cloudtrail/trail"
	compositealarm "github.com/upbound/provider-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/upbound/provider-aws/internal/controller/cloudwatch/dashboard"
	metricalarm "github.com/upbound/provider-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/upbound/provider-aws/internal/controller/cloudwatch/metricstream"
	apidestination "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/apidestination"
	archive "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/archive"
	bus "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/bus"
	buspolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/buspolicy"
	connectioncloudwatchevents "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/connection"
	permission "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/permission"
	rule "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/rule"
	targetcloudwatchevents "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/target"
	definition "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
	approvalruletemplate "github.com/upbound/provider-aws/internal/controller/codecommit/approvalruletemplate"
	approvalruletemplateassociation "github.com/upbound/provider-aws/internal/controller/codecommit/approvalruletemplateassociation"
	repository "github.com/upbound/provider-aws/internal/controller/codecommit/repository"
	trigger "github.com/upbound/provider-aws/internal/controller/codecommit/trigger"
	codepipeline "github.com/upbound/provider-aws/internal/controller/codepipeline/codepipeline"
	webhookcodepipeline "github.com/upbound/provider-aws/internal/controller/codepipeline/webhook"
	connectioncodestarconnections "github.com/upbound/provider-aws/internal/controller/codestarconnections/connection"
	host "github.com/upbound/provider-aws/internal/controller/codestarconnections/host"
	notificationrule "github.com/upbound/provider-aws/internal/controller/codestarnotifications/notificationrule"
	awsconfigurationrecorderstatus "github.com/upbound/provider-aws/internal/controller/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/upbound/provider-aws/internal/controller/configservice/configrule"
	configurationaggregator "github.com/upbound/provider-aws/internal/controller/configservice/configurationaggregator"
	configurationrecorder "github.com/upbound/provider-aws/internal/controller/configservice/configurationrecorder"
	conformancepack "github.com/upbound/provider-aws/internal/controller/configservice/conformancepack"
	deliverychannel "github.com/upbound/provider-aws/internal/controller/configservice/deliverychannel"
	remediationconfiguration "github.com/upbound/provider-aws/internal/controller/configservice/remediationconfiguration"
	botassociation "github.com/upbound/provider-aws/internal/controller/connect/botassociation"
	contactflow "github.com/upbound/provider-aws/internal/controller/connect/contactflow"
	contactflowmodule "github.com/upbound/provider-aws/internal/controller/connect/contactflowmodule"
	hoursofoperation "github.com/upbound/provider-aws/internal/controller/connect/hoursofoperation"
	instance "github.com/upbound/provider-aws/internal/controller/connect/instance"
	lambdafunctionassociation "github.com/upbound/provider-aws/internal/controller/connect/lambdafunctionassociation"
	queue "github.com/upbound/provider-aws/internal/controller/connect/queue"
	quickconnect "github.com/upbound/provider-aws/internal/controller/connect/quickconnect"
	routingprofile "github.com/upbound/provider-aws/internal/controller/connect/routingprofile"
	securityprofile "github.com/upbound/provider-aws/internal/controller/connect/securityprofile"
	userhierarchystructure "github.com/upbound/provider-aws/internal/controller/connect/userhierarchystructure"
	reportdefinition "github.com/upbound/provider-aws/internal/controller/cur/reportdefinition"
	dataset "github.com/upbound/provider-aws/internal/controller/dataexchange/dataset"
	revision "github.com/upbound/provider-aws/internal/controller/dataexchange/revision"
	pipeline "github.com/upbound/provider-aws/internal/controller/datapipeline/pipeline"
	cluster "github.com/upbound/provider-aws/internal/controller/dax/cluster"
	parametergroup "github.com/upbound/provider-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/dax/subnetgroup"
	appdeploy "github.com/upbound/provider-aws/internal/controller/deploy/app"
	deploymentconfig "github.com/upbound/provider-aws/internal/controller/deploy/deploymentconfig"
	deploymentgroup "github.com/upbound/provider-aws/internal/controller/deploy/deploymentgroup"
	graph "github.com/upbound/provider-aws/internal/controller/detective/graph"
	invitationaccepter "github.com/upbound/provider-aws/internal/controller/detective/invitationaccepter"
	member "github.com/upbound/provider-aws/internal/controller/detective/member"
	devicepool "github.com/upbound/provider-aws/internal/controller/devicefarm/devicepool"
	instanceprofile "github.com/upbound/provider-aws/internal/controller/devicefarm/instanceprofile"
	networkprofile "github.com/upbound/provider-aws/internal/controller/devicefarm/networkprofile"
	project "github.com/upbound/provider-aws/internal/controller/devicefarm/project"
	testgridproject "github.com/upbound/provider-aws/internal/controller/devicefarm/testgridproject"
	upload "github.com/upbound/provider-aws/internal/controller/devicefarm/upload"
	bgppeer "github.com/upbound/provider-aws/internal/controller/directconnect/bgppeer"
	connectiondirectconnect "github.com/upbound/provider-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/upbound/provider-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/upbound/provider-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/upbound/provider-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/upbound/provider-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/upbound/provider-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/upbound/provider-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/upbound/provider-aws/internal/controller/directconnect/transitvirtualinterface"
	lifecyclepolicy "github.com/upbound/provider-aws/internal/controller/dlm/lifecyclepolicy"
	certificate "github.com/upbound/provider-aws/internal/controller/dms/certificate"
	endpoint "github.com/upbound/provider-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/upbound/provider-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/upbound/provider-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/upbound/provider-aws/internal/controller/dms/replicationtask"
	directory "github.com/upbound/provider-aws/internal/controller/ds/directory"
	contributorinsights "github.com/upbound/provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "github.com/upbound/provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/upbound/provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/upbound/provider-aws/internal/controller/dynamodb/table"
	tableitem "github.com/upbound/provider-aws/internal/controller/dynamodb/tableitem"
	tag "github.com/upbound/provider-aws/internal/controller/dynamodb/tag"
	ami "github.com/upbound/provider-aws/internal/controller/ec2/ami"
	amicopy "github.com/upbound/provider-aws/internal/controller/ec2/amicopy"
	amilaunchpermission "github.com/upbound/provider-aws/internal/controller/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/upbound/provider-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/upbound/provider-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/upbound/provider-aws/internal/controller/ec2/carriergateway"
	customergateway "github.com/upbound/provider-aws/internal/controller/ec2/customergateway"
	defaultnetworkacl "github.com/upbound/provider-aws/internal/controller/ec2/defaultnetworkacl"
	defaultroutetable "github.com/upbound/provider-aws/internal/controller/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/upbound/provider-aws/internal/controller/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/upbound/provider-aws/internal/controller/ec2/defaultsubnet"
	defaultvpc "github.com/upbound/provider-aws/internal/controller/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/upbound/provider-aws/internal/controller/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/upbound/provider-aws/internal/controller/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/upbound/provider-aws/internal/controller/ec2/ebssnapshotimport"
	ebsvolume "github.com/upbound/provider-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/upbound/provider-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/upbound/provider-aws/internal/controller/ec2/eip"
	eipassociation "github.com/upbound/provider-aws/internal/controller/ec2/eipassociation"
	flowlog "github.com/upbound/provider-aws/internal/controller/ec2/flowlog"
	hostec2 "github.com/upbound/provider-aws/internal/controller/ec2/host"
	instanceec2 "github.com/upbound/provider-aws/internal/controller/ec2/instance"
	internetgateway "github.com/upbound/provider-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/upbound/provider-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/upbound/provider-aws/internal/controller/ec2/launchtemplate"
	mainroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlist"
	managedprefixlistentry "github.com/upbound/provider-aws/internal/controller/ec2/managedprefixlistentry"
	natgateway "github.com/upbound/provider-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/upbound/provider-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/upbound/provider-aws/internal/controller/ec2/networkaclrule"
	networkinsightspath "github.com/upbound/provider-aws/internal/controller/ec2/networkinsightspath"
	networkinterface "github.com/upbound/provider-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/upbound/provider-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/upbound/provider-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/upbound/provider-aws/internal/controller/ec2/placementgroup"
	routeec2 "github.com/upbound/provider-aws/internal/controller/ec2/route"
	routetable "github.com/upbound/provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/upbound/provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/upbound/provider-aws/internal/controller/ec2/securitygrouprule"
	serialconsoleaccess "github.com/upbound/provider-aws/internal/controller/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/upbound/provider-aws/internal/controller/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/upbound/provider-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/upbound/provider-aws/internal/controller/ec2/spotfleetrequest"
	spotinstancerequest "github.com/upbound/provider-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/upbound/provider-aws/internal/controller/ec2/subnet"
	subnetcidrreservation "github.com/upbound/provider-aws/internal/controller/ec2/subnetcidrreservation"
	trafficmirrorfilter "github.com/upbound/provider-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/upbound/provider-aws/internal/controller/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/upbound/provider-aws/internal/controller/ec2/transitgateway"
	transitgatewayconnect "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewayprefixlistreference "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/upbound/provider-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/upbound/provider-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/upbound/provider-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/upbound/provider-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/upbound/provider-aws/internal/controller/ec2/vpcipam"
	vpcipampool "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampool"
	vpcipampoolcidr "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/upbound/provider-aws/internal/controller/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/upbound/provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/upbound/provider-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/upbound/provider-aws/internal/controller/ec2/vpnconnection"
	vpnconnectionroute "github.com/upbound/provider-aws/internal/controller/ec2/vpnconnectionroute"
	vpngateway "github.com/upbound/provider-aws/internal/controller/ec2/vpngateway"
	vpngatewayattachment "github.com/upbound/provider-aws/internal/controller/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/upbound/provider-aws/internal/controller/ec2/vpngatewayroutepropagation"
	lifecyclepolicyecr "github.com/upbound/provider-aws/internal/controller/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/upbound/provider-aws/internal/controller/ecr/pullthroughcacherule"
	registrypolicy "github.com/upbound/provider-aws/internal/controller/ecr/registrypolicy"
	registryscanningconfiguration "github.com/upbound/provider-aws/internal/controller/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/upbound/provider-aws/internal/controller/ecr/replicationconfiguration"
	repositoryecr "github.com/upbound/provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/ecr/repositorypolicy"
	repositoryecrpublic "github.com/upbound/provider-aws/internal/controller/ecrpublic/repository"
	repositorypolicyecrpublic "github.com/upbound/provider-aws/internal/controller/ecrpublic/repositorypolicy"
	accountsettingdefault "github.com/upbound/provider-aws/internal/controller/ecs/accountsettingdefault"
	capacityprovider "github.com/upbound/provider-aws/internal/controller/ecs/capacityprovider"
	clusterecs "github.com/upbound/provider-aws/internal/controller/ecs/cluster"
	clustercapacityproviders "github.com/upbound/provider-aws/internal/controller/ecs/clustercapacityproviders"
	serviceecs "github.com/upbound/provider-aws/internal/controller/ecs/service"
	taskdefinition "github.com/upbound/provider-aws/internal/controller/ecs/taskdefinition"
	accesspoint "github.com/upbound/provider-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/upbound/provider-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/upbound/provider-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/internal/controller/efs/mounttarget"
	addon "github.com/upbound/provider-aws/internal/controller/eks/addon"
	clustereks "github.com/upbound/provider-aws/internal/controller/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/internal/controller/eks/clusterauth"
	fargateprofile "github.com/upbound/provider-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/cluster"
	parametergroupelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/upbound/provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroupelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/subnetgroup"
	userelasticache "github.com/upbound/provider-aws/internal/controller/elasticache/user"
	usergroup "github.com/upbound/provider-aws/internal/controller/elasticache/usergroup"
	applicationelasticbeanstalk "github.com/upbound/provider-aws/internal/controller/elasticbeanstalk/application"
	configurationtemplate "github.com/upbound/provider-aws/internal/controller/elasticbeanstalk/configurationtemplate"
	domainelasticsearch "github.com/upbound/provider-aws/internal/controller/elasticsearch/domain"
	pipelineelastictranscoder "github.com/upbound/provider-aws/internal/controller/elastictranscoder/pipeline"
	preset "github.com/upbound/provider-aws/internal/controller/elastictranscoder/preset"
	appcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/elb/appcookiestickinesspolicy"
	attachment "github.com/upbound/provider-aws/internal/controller/elb/attachment"
	backendserverpolicy "github.com/upbound/provider-aws/internal/controller/elb/backendserverpolicy"
	elb "github.com/upbound/provider-aws/internal/controller/elb/elb"
	lbcookiestickinesspolicy "github.com/upbound/provider-aws/internal/controller/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/upbound/provider-aws/internal/controller/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/upbound/provider-aws/internal/controller/elb/listenerpolicy"
	policyelb "github.com/upbound/provider-aws/internal/controller/elb/policy"
	proxyprotocolpolicy "github.com/upbound/provider-aws/internal/controller/elb/proxyprotocolpolicy"
	lb "github.com/upbound/provider-aws/internal/controller/elbv2/lb"
	lblistener "github.com/upbound/provider-aws/internal/controller/elbv2/lblistener"
	lblistenerrule "github.com/upbound/provider-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/upbound/provider-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/upbound/provider-aws/internal/controller/elbv2/lbtargetgroupattachment"
	securityconfiguration "github.com/upbound/provider-aws/internal/controller/emr/securityconfiguration"
	deliverystream "github.com/upbound/provider-aws/internal/controller/firehose/deliverystream"
	backup "github.com/upbound/provider-aws/internal/controller/fsx/backup"
	datarepositoryassociation "github.com/upbound/provider-aws/internal/controller/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/upbound/provider-aws/internal/controller/fsx/lustrefilesystem"
	ontapfilesystem "github.com/upbound/provider-aws/internal/controller/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/upbound/provider-aws/internal/controller/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/upbound/provider-aws/internal/controller/fsx/windowsfilesystem"
	vault "github.com/upbound/provider-aws/internal/controller/glacier/vault"
	vaultlock "github.com/upbound/provider-aws/internal/controller/glacier/vaultlock"
	accelerator "github.com/upbound/provider-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/upbound/provider-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/upbound/provider-aws/internal/controller/globalaccelerator/listener"
	licenseassociation "github.com/upbound/provider-aws/internal/controller/grafana/licenseassociation"
	detector "github.com/upbound/provider-aws/internal/controller/guardduty/detector"
	filter "github.com/upbound/provider-aws/internal/controller/guardduty/filter"
	memberguardduty "github.com/upbound/provider-aws/internal/controller/guardduty/member"
	accesskey "github.com/upbound/provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/upbound/provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/provider-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/upbound/provider-aws/internal/controller/iam/group"
	groupmembership "github.com/upbound/provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofileiam "github.com/upbound/provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/iam/openidconnectprovider"
	policyiam "github.com/upbound/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/upbound/provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/upbound/provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/provider-aws/internal/controller/iam/signingcertificate"
	useriam "github.com/upbound/provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/upbound/provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/upbound/provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/upbound/provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/upbound/provider-aws/internal/controller/iam/virtualmfadevice"
	containerrecipe "github.com/upbound/provider-aws/internal/controller/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/upbound/provider-aws/internal/controller/imagebuilder/distributionconfiguration"
	image "github.com/upbound/provider-aws/internal/controller/imagebuilder/image"
	imagepipeline "github.com/upbound/provider-aws/internal/controller/imagebuilder/imagepipeline"
	imagerecipe "github.com/upbound/provider-aws/internal/controller/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/upbound/provider-aws/internal/controller/imagebuilder/infrastructureconfiguration"
	assessmenttarget "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/upbound/provider-aws/internal/controller/inspector/resourcegroup"
	certificateiot "github.com/upbound/provider-aws/internal/controller/iot/certificate"
	indexingconfiguration "github.com/upbound/provider-aws/internal/controller/iot/indexingconfiguration"
	loggingoptions "github.com/upbound/provider-aws/internal/controller/iot/loggingoptions"
	policyiot "github.com/upbound/provider-aws/internal/controller/iot/policy"
	policyattachment "github.com/upbound/provider-aws/internal/controller/iot/policyattachment"
	provisioningtemplate "github.com/upbound/provider-aws/internal/controller/iot/provisioningtemplate"
	rolealias "github.com/upbound/provider-aws/internal/controller/iot/rolealias"
	thing "github.com/upbound/provider-aws/internal/controller/iot/thing"
	thinggroup "github.com/upbound/provider-aws/internal/controller/iot/thinggroup"
	thinggroupmembership "github.com/upbound/provider-aws/internal/controller/iot/thinggroupmembership"
	thingprincipalattachment "github.com/upbound/provider-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/upbound/provider-aws/internal/controller/iot/thingtype"
	topicrule "github.com/upbound/provider-aws/internal/controller/iot/topicrule"
	clusterkafka "github.com/upbound/provider-aws/internal/controller/kafka/cluster"
	configuration "github.com/upbound/provider-aws/internal/controller/kafka/configuration"
	keyspace "github.com/upbound/provider-aws/internal/controller/keyspaces/keyspace"
	tablekeyspaces "github.com/upbound/provider-aws/internal/controller/keyspaces/table"
	streamkinesis "github.com/upbound/provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/upbound/provider-aws/internal/controller/kinesis/streamconsumer"
	applicationkinesisanalytics "github.com/upbound/provider-aws/internal/controller/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/upbound/provider-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/upbound/provider-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/upbound/provider-aws/internal/controller/kinesisvideo/stream"
	alias "github.com/upbound/provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/upbound/provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/upbound/provider-aws/internal/controller/kms/externalkey"
	grant "github.com/upbound/provider-aws/internal/controller/kms/grant"
	key "github.com/upbound/provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/upbound/provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/upbound/provider-aws/internal/controller/kms/replicakey"
	aliaslambda "github.com/upbound/provider-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/upbound/provider-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/upbound/provider-aws/internal/controller/lambda/eventsourcemapping"
	functionlambda "github.com/upbound/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	functionurl "github.com/upbound/provider-aws/internal/controller/lambda/functionurl"
	invocation "github.com/upbound/provider-aws/internal/controller/lambda/invocation"
	layerversion "github.com/upbound/provider-aws/internal/controller/lambda/layerversion"
	layerversionpermission "github.com/upbound/provider-aws/internal/controller/lambda/layerversionpermission"
	permissionlambda "github.com/upbound/provider-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/upbound/provider-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	association "github.com/upbound/provider-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/upbound/provider-aws/internal/controller/licensemanager/licenseconfiguration"
	domainlightsail "github.com/upbound/provider-aws/internal/controller/lightsail/domain"
	instancelightsail "github.com/upbound/provider-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/upbound/provider-aws/internal/controller/lightsail/instancepublicports"
	keypairlightsail "github.com/upbound/provider-aws/internal/controller/lightsail/keypair"
	staticip "github.com/upbound/provider-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/upbound/provider-aws/internal/controller/lightsail/staticipattachment"
	accountmacie2 "github.com/upbound/provider-aws/internal/controller/macie2/account"
	classificationjob "github.com/upbound/provider-aws/internal/controller/macie2/classificationjob"
	customdataidentifier "github.com/upbound/provider-aws/internal/controller/macie2/customdataidentifier"
	findingsfilter "github.com/upbound/provider-aws/internal/controller/macie2/findingsfilter"
	invitationacceptermacie2 "github.com/upbound/provider-aws/internal/controller/macie2/invitationaccepter"
	membermacie2 "github.com/upbound/provider-aws/internal/controller/macie2/member"
	queuemediaconvert "github.com/upbound/provider-aws/internal/controller/mediaconvert/queue"
	channel "github.com/upbound/provider-aws/internal/controller/mediapackage/channel"
	container "github.com/upbound/provider-aws/internal/controller/mediastore/container"
	containerpolicy "github.com/upbound/provider-aws/internal/controller/mediastore/containerpolicy"
	acl "github.com/upbound/provider-aws/internal/controller/memorydb/acl"
	clustermemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/cluster"
	parametergroupmemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/memorydb/snapshot"
	subnetgroupmemorydb "github.com/upbound/provider-aws/internal/controller/memorydb/subnetgroup"
	firewallpolicy "github.com/upbound/provider-aws/internal/controller/networkfirewall/firewallpolicy"
	rulegroup "github.com/upbound/provider-aws/internal/controller/networkfirewall/rulegroup"
	connectionnetworkmanager "github.com/upbound/provider-aws/internal/controller/networkmanager/connection"
	customergatewayassociation "github.com/upbound/provider-aws/internal/controller/networkmanager/customergatewayassociation"
	device "github.com/upbound/provider-aws/internal/controller/networkmanager/device"
	globalnetwork "github.com/upbound/provider-aws/internal/controller/networkmanager/globalnetwork"
	link "github.com/upbound/provider-aws/internal/controller/networkmanager/link"
	linkassociation "github.com/upbound/provider-aws/internal/controller/networkmanager/linkassociation"
	site "github.com/upbound/provider-aws/internal/controller/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/upbound/provider-aws/internal/controller/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/upbound/provider-aws/internal/controller/networkmanager/transitgatewayregistration"
	domainopensearch "github.com/upbound/provider-aws/internal/controller/opensearch/domain"
	domainpolicy "github.com/upbound/provider-aws/internal/controller/opensearch/domainpolicy"
	domainsamloptions "github.com/upbound/provider-aws/internal/controller/opensearch/domainsamloptions"
	applicationopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/application"
	customlayer "github.com/upbound/provider-aws/internal/controller/opsworks/customlayer"
	ecsclusterlayer "github.com/upbound/provider-aws/internal/controller/opsworks/ecsclusterlayer"
	ganglialayer "github.com/upbound/provider-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/upbound/provider-aws/internal/controller/opsworks/haproxylayer"
	instanceopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/upbound/provider-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/upbound/provider-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/nodejsapplayer"
	permissionopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/upbound/provider-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/upbound/provider-aws/internal/controller/opsworks/rdsdbinstance"
	stackopsworks "github.com/upbound/provider-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/upbound/provider-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/upbound/provider-aws/internal/controller/opsworks/userprofile"
	accountorganizations "github.com/upbound/provider-aws/internal/controller/organizations/account"
	delegatedadministrator "github.com/upbound/provider-aws/internal/controller/organizations/delegatedadministrator"
	organization "github.com/upbound/provider-aws/internal/controller/organizations/organization"
	organizationalunit "github.com/upbound/provider-aws/internal/controller/organizations/organizationalunit"
	policyorganizations "github.com/upbound/provider-aws/internal/controller/organizations/policy"
	policyattachmentorganizations "github.com/upbound/provider-aws/internal/controller/organizations/policyattachment"
	apppinpoint "github.com/upbound/provider-aws/internal/controller/pinpoint/app"
	smschannel "github.com/upbound/provider-aws/internal/controller/pinpoint/smschannel"
	providerconfig "github.com/upbound/provider-aws/internal/controller/providerconfig"
	ledger "github.com/upbound/provider-aws/internal/controller/qldb/ledger"
	streamqldb "github.com/upbound/provider-aws/internal/controller/qldb/stream"
	groupquicksight "github.com/upbound/provider-aws/internal/controller/quicksight/group"
	userquicksight "github.com/upbound/provider-aws/internal/controller/quicksight/user"
	resourceshare "github.com/upbound/provider-aws/internal/controller/ram/resourceshare"
	clusterrds "github.com/upbound/provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/upbound/provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpoint "github.com/upbound/provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstance "github.com/upbound/provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergroup "github.com/upbound/provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/upbound/provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshot "github.com/upbound/provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/upbound/provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/upbound/provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscriptionrds "github.com/upbound/provider-aws/internal/controller/rds/eventsubscription"
	globalcluster "github.com/upbound/provider-aws/internal/controller/rds/globalcluster"
	instancerds "github.com/upbound/provider-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/upbound/provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/upbound/provider-aws/internal/controller/rds/optiongroup"
	parametergrouprds "github.com/upbound/provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/upbound/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/upbound/provider-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/upbound/provider-aws/internal/controller/rds/securitygroup"
	snapshotrds "github.com/upbound/provider-aws/internal/controller/rds/snapshot"
	subnetgrouprds "github.com/upbound/provider-aws/internal/controller/rds/subnetgroup"
	clusterredshift "github.com/upbound/provider-aws/internal/controller/redshift/cluster"
	eventsubscriptionredshift "github.com/upbound/provider-aws/internal/controller/redshift/eventsubscription"
	parametergroupredshift "github.com/upbound/provider-aws/internal/controller/redshift/parametergroup"
	scheduledactionredshift "github.com/upbound/provider-aws/internal/controller/redshift/scheduledaction"
	snapshotcopygrant "github.com/upbound/provider-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/upbound/provider-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/upbound/provider-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroupredshift "github.com/upbound/provider-aws/internal/controller/redshift/subnetgroup"
	groupresourcegroups "github.com/upbound/provider-aws/internal/controller/resourcegroups/group"
	delegationset "github.com/upbound/provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/upbound/provider-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/upbound/provider-aws/internal/controller/route53/hostedzonednssec"
	record "github.com/upbound/provider-aws/internal/controller/route53/record"
	trafficpolicy "github.com/upbound/provider-aws/internal/controller/route53/trafficpolicy"
	trafficpolicyinstance "github.com/upbound/provider-aws/internal/controller/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/upbound/provider-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/upbound/provider-aws/internal/controller/route53/zone"
	clusterroute53recoverycontrolconfig "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/upbound/provider-aws/internal/controller/route53recoverycontrolconfig/safetyrule"
	cell "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/cell"
	readinesscheck "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/upbound/provider-aws/internal/controller/route53recoveryreadiness/resourceset"
	endpointroute53resolver "github.com/upbound/provider-aws/internal/controller/route53resolver/endpoint"
	ruleroute53resolver "github.com/upbound/provider-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/upbound/provider-aws/internal/controller/route53resolver/ruleassociation"
	bucket "github.com/upbound/provider-aws/internal/controller/s3/bucket"
	bucketaccelerateconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/upbound/provider-aws/internal/controller/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/upbound/provider-aws/internal/controller/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/upbound/provider-aws/internal/controller/s3/bucketlogging"
	bucketmetric "github.com/upbound/provider-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/upbound/provider-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/upbound/provider-aws/internal/controller/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/upbound/provider-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/upbound/provider-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/upbound/provider-aws/internal/controller/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/upbound/provider-aws/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/upbound/provider-aws/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/upbound/provider-aws/internal/controller/s3/object"
	objectcopy "github.com/upbound/provider-aws/internal/controller/s3/objectcopy"
	accesspoints3control "github.com/upbound/provider-aws/internal/controller/s3control/accesspoint"
	accesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/upbound/provider-aws/internal/controller/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/upbound/provider-aws/internal/controller/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/upbound/provider-aws/internal/controller/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/upbound/provider-aws/internal/controller/s3control/objectlambdaaccesspointpolicy"
	appsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/upbound/provider-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/upbound/provider-aws/internal/controller/sagemaker/coderepository"
	devicesagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/device"
	devicefleet "github.com/upbound/provider-aws/internal/controller/sagemaker/devicefleet"
	domainsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/domain"
	endpointconfiguration "github.com/upbound/provider-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/upbound/provider-aws/internal/controller/sagemaker/featuregroup"
	imagesagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/image"
	imageversion "github.com/upbound/provider-aws/internal/controller/sagemaker/imageversion"
	modelsagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/model"
	modelpackagegroup "github.com/upbound/provider-aws/internal/controller/sagemaker/modelpackagegroup"
	modelpackagegrouppolicy "github.com/upbound/provider-aws/internal/controller/sagemaker/modelpackagegrouppolicy"
	notebookinstance "github.com/upbound/provider-aws/internal/controller/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/upbound/provider-aws/internal/controller/sagemaker/notebookinstancelifecycleconfiguration"
	studiolifecycleconfig "github.com/upbound/provider-aws/internal/controller/sagemaker/studiolifecycleconfig"
	userprofilesagemaker "github.com/upbound/provider-aws/internal/controller/sagemaker/userprofile"
	workforce "github.com/upbound/provider-aws/internal/controller/sagemaker/workforce"
	workteam "github.com/upbound/provider-aws/internal/controller/sagemaker/workteam"
	discoverer "github.com/upbound/provider-aws/internal/controller/schemas/discoverer"
	registry "github.com/upbound/provider-aws/internal/controller/schemas/registry"
	schema "github.com/upbound/provider-aws/internal/controller/schemas/schema"
	secret "github.com/upbound/provider-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretversion"
	accountsecurityhub "github.com/upbound/provider-aws/internal/controller/securityhub/account"
	actiontarget "github.com/upbound/provider-aws/internal/controller/securityhub/actiontarget"
	findingaggregator "github.com/upbound/provider-aws/internal/controller/securityhub/findingaggregator"
	insight "github.com/upbound/provider-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/upbound/provider-aws/internal/controller/securityhub/inviteaccepter"
	membersecurityhub "github.com/upbound/provider-aws/internal/controller/securityhub/member"
	productsubscription "github.com/upbound/provider-aws/internal/controller/securityhub/productsubscription"
	standardssubscription "github.com/upbound/provider-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/upbound/provider-aws/internal/controller/serverlessrepo/cloudformationstack"
	budgetresourceassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/upbound/provider-aws/internal/controller/servicecatalog/constraint"
	portfolio "github.com/upbound/provider-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/upbound/provider-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/upbound/provider-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/upbound/provider-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/upbound/provider-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/upbound/provider-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
	httpnamespace "github.com/upbound/provider-aws/internal/controller/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/upbound/provider-aws/internal/controller/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/upbound/provider-aws/internal/controller/servicediscovery/publicdnsnamespace"
	serviceservicediscovery "github.com/upbound/provider-aws/internal/controller/servicediscovery/service"
	servicequota "github.com/upbound/provider-aws/internal/controller/servicequotas/servicequota"
	activereceiptruleset "github.com/upbound/provider-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/upbound/provider-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/upbound/provider-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/upbound/provider-aws/internal/controller/ses/domainidentity"
	domainmailfrom "github.com/upbound/provider-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/upbound/provider-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/upbound/provider-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/upbound/provider-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/upbound/provider-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/upbound/provider-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/upbound/provider-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/upbound/provider-aws/internal/controller/ses/receiptruleset"
	template "github.com/upbound/provider-aws/internal/controller/ses/template"
	activity "github.com/upbound/provider-aws/internal/controller/sfn/activity"
	statemachine "github.com/upbound/provider-aws/internal/controller/sfn/statemachine"
	signingjob "github.com/upbound/provider-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/upbound/provider-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/upbound/provider-aws/internal/controller/signer/signingprofilepermission"
	domainsimpledb "github.com/upbound/provider-aws/internal/controller/simpledb/domain"
	platformapplication "github.com/upbound/provider-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/upbound/provider-aws/internal/controller/sns/smspreferences"
	topic "github.com/upbound/provider-aws/internal/controller/sns/topic"
	topicpolicy "github.com/upbound/provider-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/upbound/provider-aws/internal/controller/sns/topicsubscription"
	queuesqs "github.com/upbound/provider-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/upbound/provider-aws/internal/controller/sqs/queuepolicy"
	activation "github.com/upbound/provider-aws/internal/controller/ssm/activation"
	associationssm "github.com/upbound/provider-aws/internal/controller/ssm/association"
	document "github.com/upbound/provider-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindowtask"
	parameter "github.com/upbound/provider-aws/internal/controller/ssm/parameter"
	patchbaseline "github.com/upbound/provider-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/upbound/provider-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/upbound/provider-aws/internal/controller/ssm/resourcedatasync"
	domainswf "github.com/upbound/provider-aws/internal/controller/swf/domain"
	databasetimestreamwrite "github.com/upbound/provider-aws/internal/controller/timestreamwrite/database"
	tabletimestreamwrite "github.com/upbound/provider-aws/internal/controller/timestreamwrite/table"
	bytematchset "github.com/upbound/provider-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/upbound/provider-aws/internal/controller/waf/geomatchset"
	ipset "github.com/upbound/provider-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/upbound/provider-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/upbound/provider-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/upbound/provider-aws/internal/controller/waf/regexpatternset"
	rulewaf "github.com/upbound/provider-aws/internal/controller/waf/rule"
	sizeconstraintset "github.com/upbound/provider-aws/internal/controller/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/upbound/provider-aws/internal/controller/waf/sqlinjectionmatchset"
	webacl "github.com/upbound/provider-aws/internal/controller/waf/webacl"
	xssmatchset "github.com/upbound/provider-aws/internal/controller/waf/xssmatchset"
	bytematchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/bytematchset"
	geomatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/geomatchset"
	ipsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/ipset"
	ratebasedrulewafregional "github.com/upbound/provider-aws/internal/controller/wafregional/ratebasedrule"
	regexmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/regexmatchset"
	regexpatternsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/regexpatternset"
	rulewafregional "github.com/upbound/provider-aws/internal/controller/wafregional/rule"
	sizeconstraintsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/sizeconstraintset"
	sqlinjectionmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/sqlinjectionmatchset"
	webaclwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/webacl"
	xssmatchsetwafregional "github.com/upbound/provider-aws/internal/controller/wafregional/xssmatchset"
	ipsetwafv2 "github.com/upbound/provider-aws/internal/controller/wafv2/ipset"
	regexpatternsetwafv2 "github.com/upbound/provider-aws/internal/controller/wafv2/regexpatternset"
	directoryworkspaces "github.com/upbound/provider-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/upbound/provider-aws/internal/controller/workspaces/ipgroup"
	encryptionconfig "github.com/upbound/provider-aws/internal/controller/xray/encryptionconfig"
	groupxray "github.com/upbound/provider-aws/internal/controller/xray/group"
	samplingrule "github.com/upbound/provider-aws/internal/controller/xray/samplingrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		alternatecontact.Setup,
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		webhook.Setup,
		account.Setup,
		apikey.Setup,
		authorizer.Setup,
		basepathmapping.Setup,
		clientcertificate.Setup,
		deployment.Setup,
		documentationpart.Setup,
		documentationversion.Setup,
		domainname.Setup,
		gatewayresponse.Setup,
		integration.Setup,
		integrationresponse.Setup,
		method.Setup,
		methodresponse.Setup,
		methodsettings.Setup,
		model.Setup,
		requestvalidator.Setup,
		resource.Setup,
		restapi.Setup,
		restapipolicy.Setup,
		stage.Setup,
		usageplan.Setup,
		usageplankey.Setup,
		vpclink.Setup,
		policy.Setup,
		scheduledaction.Setup,
		target.Setup,
		application.Setup,
		configurationprofile.Setup,
		deploymentappconfig.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		hostedconfigurationversion.Setup,
		flow.Setup,
		eventintegration.Setup,
		gatewayroute.Setup,
		mesh.Setup,
		route.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		virtualservice.Setup,
		autoscalingconfigurationversion.Setup,
		connection.Setup,
		service.Setup,
		vpcconnector.Setup,
		directoryconfig.Setup,
		fleet.Setup,
		fleetstackassociation.Setup,
		imagebuilder.Setup,
		stack.Setup,
		user.Setup,
		userstackassociation.Setup,
		apicache.Setup,
		apikeyappsync.Setup,
		datasource.Setup,
		function.Setup,
		graphqlapi.Setup,
		resolver.Setup,
		database.Setup,
		datacatalog.Setup,
		namedquery.Setup,
		workgroup.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policyautoscaling.Setup,
		schedule.Setup,
		scalingplan.Setup,
		schedulingpolicy.Setup,
		budget.Setup,
		budgetaction.Setup,
		voiceconnector.Setup,
		voiceconnectorgroup.Setup,
		voiceconnectorlogging.Setup,
		voiceconnectororigination.Setup,
		voiceconnectorstreaming.Setup,
		voiceconnectortermination.Setup,
		voiceconnectorterminationcredentials.Setup,
		environmentec2.Setup,
		environmentmembership.Setup,
		resourcecloudcontrol.Setup,
		stackcloudformation.Setup,
		stackset.Setup,
		domain.Setup,
		domainserviceaccesspolicy.Setup,
		eventdatastore.Setup,
		trail.Setup,
		compositealarm.Setup,
		dashboard.Setup,
		metricalarm.Setup,
		metricstream.Setup,
		apidestination.Setup,
		archive.Setup,
		bus.Setup,
		buspolicy.Setup,
		connectioncloudwatchevents.Setup,
		permission.Setup,
		rule.Setup,
		targetcloudwatchevents.Setup,
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
		approvalruletemplate.Setup,
		approvalruletemplateassociation.Setup,
		repository.Setup,
		trigger.Setup,
		codepipeline.Setup,
		webhookcodepipeline.Setup,
		connectioncodestarconnections.Setup,
		host.Setup,
		notificationrule.Setup,
		awsconfigurationrecorderstatus.Setup,
		configrule.Setup,
		configurationaggregator.Setup,
		configurationrecorder.Setup,
		conformancepack.Setup,
		deliverychannel.Setup,
		remediationconfiguration.Setup,
		botassociation.Setup,
		contactflow.Setup,
		contactflowmodule.Setup,
		hoursofoperation.Setup,
		instance.Setup,
		lambdafunctionassociation.Setup,
		queue.Setup,
		quickconnect.Setup,
		routingprofile.Setup,
		securityprofile.Setup,
		userhierarchystructure.Setup,
		reportdefinition.Setup,
		dataset.Setup,
		revision.Setup,
		pipeline.Setup,
		cluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		appdeploy.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
		graph.Setup,
		invitationaccepter.Setup,
		member.Setup,
		devicepool.Setup,
		instanceprofile.Setup,
		networkprofile.Setup,
		project.Setup,
		testgridproject.Setup,
		upload.Setup,
		bgppeer.Setup,
		connectiondirectconnect.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
		lifecyclepolicy.Setup,
		certificate.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		directory.Setup,
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		tag.Setup,
		ami.Setup,
		amicopy.Setup,
		amilaunchpermission.Setup,
		availabilityzonegroup.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		customergateway.Setup,
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		ebsdefaultkmskey.Setup,
		ebsencryptionbydefault.Setup,
		ebssnapshot.Setup,
		ebssnapshotcopy.Setup,
		ebssnapshotimport.Setup,
		ebsvolume.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		flowlog.Setup,
		hostec2.Setup,
		instanceec2.Setup,
		internetgateway.Setup,
		keypair.Setup,
		launchtemplate.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		managedprefixlistentry.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkinsightspath.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		placementgroup.Setup,
		routeec2.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		serialconsoleaccess.Setup,
		snapshotcreatevolumepermission.Setup,
		spotdatafeedsubscription.Setup,
		spotfleetrequest.Setup,
		spotinstancerequest.Setup,
		subnet.Setup,
		subnetcidrreservation.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilterrule.Setup,
		transitgateway.Setup,
		transitgatewayconnect.Setup,
		transitgatewayconnectpeer.Setup,
		transitgatewaymulticastdomain.Setup,
		transitgatewaymulticastdomainassociation.Setup,
		transitgatewaymulticastgroupmember.Setup,
		transitgatewaymulticastgroupsource.Setup,
		transitgatewaypeeringattachment.Setup,
		transitgatewaypeeringattachmentaccepter.Setup,
		transitgatewayprefixlistreference.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		volumeattachment.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		vpcdhcpoptionsassociation.Setup,
		vpcendpoint.Setup,
		vpcendpointconnectionnotification.Setup,
		vpcendpointroutetableassociation.Setup,
		vpcendpointsecuritygroupassociation.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpcipam.Setup,
		vpcipampool.Setup,
		vpcipampoolcidr.Setup,
		vpcipampoolcidrallocation.Setup,
		vpcipamscope.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionaccepter.Setup,
		vpcpeeringconnectionoptions.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
		vpngateway.Setup,
		vpngatewayattachment.Setup,
		vpngatewayroutepropagation.Setup,
		lifecyclepolicyecr.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repositoryecr.Setup,
		repositorypolicy.Setup,
		repositoryecrpublic.Setup,
		repositorypolicyecrpublic.Setup,
		accountsettingdefault.Setup,
		capacityprovider.Setup,
		clusterecs.Setup,
		clustercapacityproviders.Setup,
		serviceecs.Setup,
		taskdefinition.Setup,
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		addon.Setup,
		clustereks.Setup,
		clusterauth.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		clusterelasticache.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		subnetgroupelasticache.Setup,
		userelasticache.Setup,
		usergroup.Setup,
		applicationelasticbeanstalk.Setup,
		configurationtemplate.Setup,
		domainelasticsearch.Setup,
		pipelineelastictranscoder.Setup,
		preset.Setup,
		appcookiestickinesspolicy.Setup,
		attachment.Setup,
		backendserverpolicy.Setup,
		elb.Setup,
		lbcookiestickinesspolicy.Setup,
		lbsslnegotiationpolicy.Setup,
		listenerpolicy.Setup,
		policyelb.Setup,
		proxyprotocolpolicy.Setup,
		lb.Setup,
		lblistener.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		securityconfiguration.Setup,
		deliverystream.Setup,
		backup.Setup,
		datarepositoryassociation.Setup,
		lustrefilesystem.Setup,
		ontapfilesystem.Setup,
		ontapstoragevirtualmachine.Setup,
		windowsfilesystem.Setup,
		vault.Setup,
		vaultlock.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		licenseassociation.Setup,
		detector.Setup,
		filter.Setup,
		memberguardduty.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofileiam.Setup,
		openidconnectprovider.Setup,
		policyiam.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		useriam.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
		containerrecipe.Setup,
		distributionconfiguration.Setup,
		image.Setup,
		imagepipeline.Setup,
		imagerecipe.Setup,
		infrastructureconfiguration.Setup,
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
		certificateiot.Setup,
		indexingconfiguration.Setup,
		loggingoptions.Setup,
		policyiot.Setup,
		policyattachment.Setup,
		provisioningtemplate.Setup,
		rolealias.Setup,
		thing.Setup,
		thinggroup.Setup,
		thinggroupmembership.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		topicrule.Setup,
		clusterkafka.Setup,
		configuration.Setup,
		keyspace.Setup,
		tablekeyspaces.Setup,
		streamkinesis.Setup,
		streamconsumer.Setup,
		applicationkinesisanalytics.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		streamkinesisvideo.Setup,
		alias.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
		aliaslambda.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		functionlambda.Setup,
		functioneventinvokeconfig.Setup,
		functionurl.Setup,
		invocation.Setup,
		layerversion.Setup,
		layerversionpermission.Setup,
		permissionlambda.Setup,
		provisionedconcurrencyconfig.Setup,
		association.Setup,
		licenseconfiguration.Setup,
		domainlightsail.Setup,
		instancelightsail.Setup,
		instancepublicports.Setup,
		keypairlightsail.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		accountmacie2.Setup,
		classificationjob.Setup,
		customdataidentifier.Setup,
		findingsfilter.Setup,
		invitationacceptermacie2.Setup,
		membermacie2.Setup,
		queuemediaconvert.Setup,
		channel.Setup,
		container.Setup,
		containerpolicy.Setup,
		acl.Setup,
		clustermemorydb.Setup,
		parametergroupmemorydb.Setup,
		snapshot.Setup,
		subnetgroupmemorydb.Setup,
		firewallpolicy.Setup,
		rulegroup.Setup,
		connectionnetworkmanager.Setup,
		customergatewayassociation.Setup,
		device.Setup,
		globalnetwork.Setup,
		link.Setup,
		linkassociation.Setup,
		site.Setup,
		transitgatewayconnectpeerassociation.Setup,
		transitgatewayregistration.Setup,
		domainopensearch.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		applicationopsworks.Setup,
		customlayer.Setup,
		ecsclusterlayer.Setup,
		ganglialayer.Setup,
		haproxylayer.Setup,
		instanceopsworks.Setup,
		javaapplayer.Setup,
		memcachedlayer.Setup,
		mysqllayer.Setup,
		nodejsapplayer.Setup,
		permissionopsworks.Setup,
		phpapplayer.Setup,
		railsapplayer.Setup,
		rdsdbinstance.Setup,
		stackopsworks.Setup,
		staticweblayer.Setup,
		userprofile.Setup,
		accountorganizations.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policyorganizations.Setup,
		policyattachmentorganizations.Setup,
		apppinpoint.Setup,
		smschannel.Setup,
		providerconfig.Setup,
		ledger.Setup,
		streamqldb.Setup,
		groupquicksight.Setup,
		userquicksight.Setup,
		resourceshare.Setup,
		clusterrds.Setup,
		clusteractivitystream.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clusterroleassociation.Setup,
		clustersnapshot.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscriptionrds.Setup,
		globalcluster.Setup,
		instancerds.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergrouprds.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshotrds.Setup,
		subnetgrouprds.Setup,
		clusterredshift.Setup,
		eventsubscriptionredshift.Setup,
		parametergroupredshift.Setup,
		scheduledactionredshift.Setup,
		snapshotcopygrant.Setup,
		snapshotschedule.Setup,
		snapshotscheduleassociation.Setup,
		subnetgroupredshift.Setup,
		groupresourcegroups.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		record.Setup,
		trafficpolicy.Setup,
		trafficpolicyinstance.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		clusterroute53recoverycontrolconfig.Setup,
		controlpanel.Setup,
		routingcontrol.Setup,
		safetyrule.Setup,
		cell.Setup,
		readinesscheck.Setup,
		recoverygroup.Setup,
		resourceset.Setup,
		endpointroute53resolver.Setup,
		ruleroute53resolver.Setup,
		ruleassociation.Setup,
		bucket.Setup,
		bucketaccelerateconfiguration.Setup,
		bucketacl.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketcorsconfiguration.Setup,
		bucketintelligenttieringconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketlogging.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreplicationconfiguration.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		object.Setup,
		objectcopy.Setup,
		accesspoints3control.Setup,
		accesspointpolicy.Setup,
		accountpublicaccessblock.Setup,
		multiregionaccesspoint.Setup,
		multiregionaccesspointpolicy.Setup,
		objectlambdaaccesspoint.Setup,
		objectlambdaaccesspointpolicy.Setup,
		appsagemaker.Setup,
		appimageconfig.Setup,
		coderepository.Setup,
		devicesagemaker.Setup,
		devicefleet.Setup,
		domainsagemaker.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
		imagesagemaker.Setup,
		imageversion.Setup,
		modelsagemaker.Setup,
		modelpackagegroup.Setup,
		modelpackagegrouppolicy.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		studiolifecycleconfig.Setup,
		userprofilesagemaker.Setup,
		workforce.Setup,
		workteam.Setup,
		discoverer.Setup,
		registry.Setup,
		schema.Setup,
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
		accountsecurityhub.Setup,
		actiontarget.Setup,
		findingaggregator.Setup,
		insight.Setup,
		inviteaccepter.Setup,
		membersecurityhub.Setup,
		productsubscription.Setup,
		standardssubscription.Setup,
		cloudformationstack.Setup,
		budgetresourceassociation.Setup,
		constraint.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
		httpnamespace.Setup,
		privatednsnamespace.Setup,
		publicdnsnamespace.Setup,
		serviceservicediscovery.Setup,
		servicequota.Setup,
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
		activity.Setup,
		statemachine.Setup,
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
		domainsimpledb.Setup,
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
		queuesqs.Setup,
		queuepolicy.Setup,
		activation.Setup,
		associationssm.Setup,
		document.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		resourcedatasync.Setup,
		domainswf.Setup,
		databasetimestreamwrite.Setup,
		tabletimestreamwrite.Setup,
		bytematchset.Setup,
		geomatchset.Setup,
		ipset.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rulewaf.Setup,
		sizeconstraintset.Setup,
		sqlinjectionmatchset.Setup,
		webacl.Setup,
		xssmatchset.Setup,
		bytematchsetwafregional.Setup,
		geomatchsetwafregional.Setup,
		ipsetwafregional.Setup,
		ratebasedrulewafregional.Setup,
		regexmatchsetwafregional.Setup,
		regexpatternsetwafregional.Setup,
		rulewafregional.Setup,
		sizeconstraintsetwafregional.Setup,
		sqlinjectionmatchsetwafregional.Setup,
		webaclwafregional.Setup,
		xssmatchsetwafregional.Setup,
		ipsetwafv2.Setup,
		regexpatternsetwafv2.Setup,
		directoryworkspaces.Setup,
		ipgroup.Setup,
		encryptionconfig.Setup,
		groupxray.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
