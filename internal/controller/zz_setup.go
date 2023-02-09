/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	alternatecontact "github.com/upbound/provider-aws/internal/controller/account/alternatecontact"
	grouptag "github.com/upbound/provider-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/upbound/provider-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/upbound/provider-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/upbound/provider-aws/internal/controller/autoscaling/notification"
	policy "github.com/upbound/provider-aws/internal/controller/autoscaling/policy"
	schedule "github.com/upbound/provider-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/upbound/provider-aws/internal/controller/autoscalingplans/scalingplan"
	schedulingpolicy "github.com/upbound/provider-aws/internal/controller/batch/schedulingpolicy"
	voiceconnector "github.com/upbound/provider-aws/internal/controller/chime/voiceconnector"
	voiceconnectorgroup "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/upbound/provider-aws/internal/controller/chime/voiceconnectorterminationcredentials"
	environmentec2 "github.com/upbound/provider-aws/internal/controller/cloud9/environmentec2"
	environmentmembership "github.com/upbound/provider-aws/internal/controller/cloud9/environmentmembership"
	resource "github.com/upbound/provider-aws/internal/controller/cloudcontrol/resource"
	stack "github.com/upbound/provider-aws/internal/controller/cloudformation/stack"
	stackset "github.com/upbound/provider-aws/internal/controller/cloudformation/stackset"
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
	connection "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/connection"
	permission "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/permission"
	rule "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/rule"
	target "github.com/upbound/provider-aws/internal/controller/cloudwatchevents/target"
	definition "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/upbound/provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
	reportdefinition "github.com/upbound/provider-aws/internal/controller/cur/reportdefinition"
	lifecyclepolicy "github.com/upbound/provider-aws/internal/controller/dlm/lifecyclepolicy"
	certificate "github.com/upbound/provider-aws/internal/controller/dms/certificate"
	endpoint "github.com/upbound/provider-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/upbound/provider-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/upbound/provider-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/upbound/provider-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/upbound/provider-aws/internal/controller/dms/replicationtask"
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
	host "github.com/upbound/provider-aws/internal/controller/ec2/host"
	instance "github.com/upbound/provider-aws/internal/controller/ec2/instance"
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
	route "github.com/upbound/provider-aws/internal/controller/ec2/route"
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
	repository "github.com/upbound/provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/upbound/provider-aws/internal/controller/ecr/repositorypolicy"
	accesspoint "github.com/upbound/provider-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/upbound/provider-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/upbound/provider-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/upbound/provider-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/upbound/provider-aws/internal/controller/efs/mounttarget"
	addon "github.com/upbound/provider-aws/internal/controller/eks/addon"
	cluster "github.com/upbound/provider-aws/internal/controller/eks/cluster"
	clusterauth "github.com/upbound/provider-aws/internal/controller/eks/clusterauth"
	fargateprofile "github.com/upbound/provider-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/upbound/provider-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/upbound/provider-aws/internal/controller/eks/nodegroup"
	domain "github.com/upbound/provider-aws/internal/controller/elasticsearch/domain"
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
	deliverystream "github.com/upbound/provider-aws/internal/controller/firehose/deliverystream"
	accesskey "github.com/upbound/provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/upbound/provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/upbound/provider-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/upbound/provider-aws/internal/controller/iam/group"
	groupmembership "github.com/upbound/provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofile "github.com/upbound/provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/upbound/provider-aws/internal/controller/iam/openidconnectprovider"
	policyiam "github.com/upbound/provider-aws/internal/controller/iam/policy"
	role "github.com/upbound/provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/upbound/provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/upbound/provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/upbound/provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/upbound/provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/upbound/provider-aws/internal/controller/iam/signingcertificate"
	user "github.com/upbound/provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/upbound/provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/upbound/provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/upbound/provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/upbound/provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/upbound/provider-aws/internal/controller/iam/virtualmfadevice"
	assessmenttarget "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/upbound/provider-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/upbound/provider-aws/internal/controller/inspector/resourcegroup"
	streamkinesis "github.com/upbound/provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/upbound/provider-aws/internal/controller/kinesis/streamconsumer"
	application "github.com/upbound/provider-aws/internal/controller/kinesisanalytics/application"
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
	function "github.com/upbound/provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/upbound/provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	functionurl "github.com/upbound/provider-aws/internal/controller/lambda/functionurl"
	invocation "github.com/upbound/provider-aws/internal/controller/lambda/invocation"
	layerversion "github.com/upbound/provider-aws/internal/controller/lambda/layerversion"
	layerversionpermission "github.com/upbound/provider-aws/internal/controller/lambda/layerversionpermission"
	permissionlambda "github.com/upbound/provider-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/upbound/provider-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	domainlightsail "github.com/upbound/provider-aws/internal/controller/lightsail/domain"
	instancelightsail "github.com/upbound/provider-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/upbound/provider-aws/internal/controller/lightsail/instancepublicports"
	keypairlightsail "github.com/upbound/provider-aws/internal/controller/lightsail/keypair"
	staticip "github.com/upbound/provider-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/upbound/provider-aws/internal/controller/lightsail/staticipattachment"
	providerconfig "github.com/upbound/provider-aws/internal/controller/providerconfig"
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
	parametergroup "github.com/upbound/provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/upbound/provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/upbound/provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/upbound/provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/upbound/provider-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/upbound/provider-aws/internal/controller/rds/securitygroup"
	snapshot "github.com/upbound/provider-aws/internal/controller/rds/snapshot"
	subnetgroup "github.com/upbound/provider-aws/internal/controller/rds/subnetgroup"
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
	discoverer "github.com/upbound/provider-aws/internal/controller/schemas/discoverer"
	registry "github.com/upbound/provider-aws/internal/controller/schemas/registry"
	schema "github.com/upbound/provider-aws/internal/controller/schemas/schema"
	secret "github.com/upbound/provider-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/upbound/provider-aws/internal/controller/secretsmanager/secretversion"
	account "github.com/upbound/provider-aws/internal/controller/securityhub/account"
	actiontarget "github.com/upbound/provider-aws/internal/controller/securityhub/actiontarget"
	findingaggregator "github.com/upbound/provider-aws/internal/controller/securityhub/findingaggregator"
	insight "github.com/upbound/provider-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/upbound/provider-aws/internal/controller/securityhub/inviteaccepter"
	member "github.com/upbound/provider-aws/internal/controller/securityhub/member"
	productsubscription "github.com/upbound/provider-aws/internal/controller/securityhub/productsubscription"
	standardssubscription "github.com/upbound/provider-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/upbound/provider-aws/internal/controller/serverlessrepo/cloudformationstack"
	servicequota "github.com/upbound/provider-aws/internal/controller/servicequotas/servicequota"
	signingjob "github.com/upbound/provider-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/upbound/provider-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/upbound/provider-aws/internal/controller/signer/signingprofilepermission"
	domainsimpledb "github.com/upbound/provider-aws/internal/controller/simpledb/domain"
	platformapplication "github.com/upbound/provider-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/upbound/provider-aws/internal/controller/sns/smspreferences"
	topic "github.com/upbound/provider-aws/internal/controller/sns/topic"
	topicpolicy "github.com/upbound/provider-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/upbound/provider-aws/internal/controller/sns/topicsubscription"
	queue "github.com/upbound/provider-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/upbound/provider-aws/internal/controller/sqs/queuepolicy"
	activation "github.com/upbound/provider-aws/internal/controller/ssm/activation"
	association "github.com/upbound/provider-aws/internal/controller/ssm/association"
	document "github.com/upbound/provider-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/upbound/provider-aws/internal/controller/ssm/maintenancewindowtask"
	parameter "github.com/upbound/provider-aws/internal/controller/ssm/parameter"
	patchbaseline "github.com/upbound/provider-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/upbound/provider-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/upbound/provider-aws/internal/controller/ssm/resourcedatasync"
	domainswf "github.com/upbound/provider-aws/internal/controller/swf/domain"
	database "github.com/upbound/provider-aws/internal/controller/timestreamwrite/database"
	tabletimestreamwrite "github.com/upbound/provider-aws/internal/controller/timestreamwrite/table"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alternatecontact.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policy.Setup,
		schedule.Setup,
		scalingplan.Setup,
		schedulingpolicy.Setup,
		voiceconnector.Setup,
		voiceconnectorgroup.Setup,
		voiceconnectorlogging.Setup,
		voiceconnectororigination.Setup,
		voiceconnectorstreaming.Setup,
		voiceconnectortermination.Setup,
		voiceconnectorterminationcredentials.Setup,
		environmentec2.Setup,
		environmentmembership.Setup,
		resource.Setup,
		stack.Setup,
		stackset.Setup,
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
		connection.Setup,
		permission.Setup,
		rule.Setup,
		target.Setup,
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
		reportdefinition.Setup,
		lifecyclepolicy.Setup,
		certificate.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
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
		host.Setup,
		instance.Setup,
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
		route.Setup,
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
		repository.Setup,
		repositorypolicy.Setup,
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		addon.Setup,
		cluster.Setup,
		clusterauth.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		domain.Setup,
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
		deliverystream.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofile.Setup,
		openidconnectprovider.Setup,
		policyiam.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		user.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
		streamkinesis.Setup,
		streamconsumer.Setup,
		application.Setup,
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
		function.Setup,
		functioneventinvokeconfig.Setup,
		functionurl.Setup,
		invocation.Setup,
		layerversion.Setup,
		layerversionpermission.Setup,
		permissionlambda.Setup,
		provisionedconcurrencyconfig.Setup,
		domainlightsail.Setup,
		instancelightsail.Setup,
		instancepublicports.Setup,
		keypairlightsail.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		providerconfig.Setup,
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
		parametergroup.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
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
		discoverer.Setup,
		registry.Setup,
		schema.Setup,
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
		account.Setup,
		actiontarget.Setup,
		findingaggregator.Setup,
		insight.Setup,
		inviteaccepter.Setup,
		member.Setup,
		productsubscription.Setup,
		standardssubscription.Setup,
		cloudformationstack.Setup,
		servicequota.Setup,
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
		domainsimpledb.Setup,
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
		queue.Setup,
		queuepolicy.Setup,
		activation.Setup,
		association.Setup,
		document.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		resourcedatasync.Setup,
		domainswf.Setup,
		database.Setup,
		tabletimestreamwrite.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
