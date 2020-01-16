package cloudability

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
// 	"github.com/hashicorp/terraform-plugin-sdk/terraform"
// 	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
// 	"os"
// 	"testing"
// )

// func TestAccMasterAccount(t *testing.T) {
// 	var account *cloudability.Account
// 	payerAccountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
// 	cloudabilityApikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	awsRegion := os.Getenv("AWS_DEFAULT_REGION")

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckMasterAccountDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccMasterAccount(cloudabilityApikey, awsRegion, payerAccountId),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckMasterAccountResourceExists("cloudability_master_account.aws_payer_account", account),
// 					// testAccCheckExampleWidgetValues(widget, rName),
// 					resource.TestCheckResourceAttr("cloudability_master_account.aws_payer_account", "vendor_account_id", payerAccountId),
// 					resource.TestCheckResourceAttr("cloudability_master_account.aws_payer_account", "vendor_key", "aws"),
// 				),
// 			},
// 		},
// 	})
// }

// // testAccCheckMasterAccountDestroy verifies the MasterAccount resource has been destroyed.
// func testAccCheckMasterAccountDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*cloudability.Client)
// 	// loop through the resources in state, verifying each resource is destroyed
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "cloudability_master_account" {
// 			continue
// 		}
// 		account, err := client.Vendors().GetAccount(rs.Primary.Attributes["vendor_key"], rs.Primary.ID)
// 		if err != nil {
// 			// Ignore 404 errors (No account found)
// 			var apiError cloudability.APIError
// 			jsonErr := json.Unmarshal([]byte(err.Error()), &apiError)
// 			if jsonErr != nil || apiError.Error.Code != 404 {
// 				return err
// 			}
// 		} else if account != nil {
// 			if account.Verification != nil || account.Authorization != nil {
// 				return fmt.Errorf("Master Account (%s) still exists.", rs.Primary.ID)
// 			}
// 		}
// 	}
// 	return nil
// }

// // testAccCheckMasterAccountResourceExists uses the Cloudability SDK to retrieve the Account.
// func testAccCheckMasterAccountResourceExists(resourceName string, account *cloudability.Account) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		// retrieve the resource by name from state
// 		rs, ok := s.RootModule().Resources[resourceName]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", resourceName)
// 		}
// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("MasterAccount ID is not set")
// 		}
// 		// retrieve the client from the test provider
// 		client := testAccProvider.Meta().(*cloudability.Client)
// 		var err error
// 		account, err = client.Vendors().GetAccount(rs.Primary.Attributes["vendor_key"], rs.Primary.ID)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

// // // testAccMasterAccount returns a configuration for an MasterAccount
// // func testAccMasterAccount(cloudabilityApikey string, awsRegion string, payerAccountId string) string {
// // 	return fmt.Sprintf(`
// // module "cloudability_master_account" {
// // 	source = "github.com/skyscrapr/terraform-cloudability-modules//cloudability-master-account"
// // 	aws_payer_account_id = "%s"
// // 	aws_region = "%s"
// // 	cloudability_apikey = "%s"
// // }`, payerAccountId, awsRegion, cloudabilityApikey)
// // }

// func testAccMasterAccount(cloudabilityApikey string, awsRegion string, payerAccountId string) string {
// 	return fmt.Sprintf(`
// provider cloudability {
// 	apikey = "%s"
// }

// resource "aws_s3_bucket" "cloudability" {
// 	bucket_prefix = "cloudability-"
// 	acl    = "private"
// 	force_destroy = true
// }

// data "aws_iam_policy_document" "cloudability" {
// 	statement {
// 		actions = [
// 			"s3:GetBucketAcl",
// 			"s3:GetBucketPolicy"
// 		]
// 		principals {
// 			type = "Service"
// 			identifiers = ["billingreports.amazonaws.com"]
// 		}
// 		resources = [
// 			"${aws_s3_bucket.cloudability.arn}"
// 		]
// 	}
// 	statement {
// 		actions = [
// 			"s3:PutObject"
// 		]
// 		principals {
// 			type = "Service"
// 			identifiers = ["billingreports.amazonaws.com"]
// 		}
// 		resources = [
// 			"${aws_s3_bucket.cloudability.arn}/*"
// 		]
// 	}
// }

// resource "aws_s3_bucket_policy" "cloudability" {
// 	bucket = aws_s3_bucket.cloudability.id
// 	policy = data.aws_iam_policy_document.cloudability.json
// }

// resource "aws_cur_report_definition" "cloudability" {
// 	report_name                = "Cloudability"
// 	time_unit                  = "HOURLY"
// 	format                     = "textORcsv"
// 	compression                = "GZIP"
// 	additional_schema_elements = ["RESOURCES"]
// 	s3_bucket                  = aws_s3_bucket.cloudability.bucket
// 	s3_region                  = "%s"
// 	s3_prefix                  = "CostAndUsageReports"
// 	depends_on = [
// 		aws_s3_bucket_policy.cloudability
// 	]
// }

// resource "cloudability_master_account" "aws_payer_account" {
//     vendor_account_id = "%s"
//     bucket_name = aws_s3_bucket.cloudability.bucket
//     report_name = "Cloudability"
// 	report_prefix = "CostAndUsageReports"
// 	depends_on = [
// 		aws_cur_report_definition.cloudability
// 	]
// }

// resource "aws_iam_role" "cloudability_role" {
// 	name               = "CloudabilityRole"
// 	path               = "/"
// 	assume_role_policy = <<EOF
// {
// 	"Version": "2012-10-17",
// 	"Statement": [
// 	  {
// 		"Action": "sts:AssumeRole",
// 		"Principal": {
// 		  "AWS": "arn:aws:iam::165736516723:user/cloudability"
// 		},
// 		"Effect": "Allow",
// 		"Condition": {
// 		  "StringEquals": {
// 			"sts:ExternalId": "${cloudability_master_account.aws_payer_account.external_id}"
// 		  }
// 		}
// 	  }
// 	]
// }
// EOF
// }

// resource "aws_iam_policy" "cloudability_verification_policy" {
// 	name   = "CloudabilityVerificationPolicy"
// 	path   = "/"
// 	policy = <<EOF
// {
// 	"Version": "2012-10-17",
// 	"Statement": [
// 	  {
// 		"Sid": "VerifyRolePermissions",
// 		"Effect": "Allow",
// 		"Action": "iam:SimulatePrincipalPolicy",
// 		"Resource": "${aws_iam_role.cloudability_role.arn}"
// 	  }
// 	]
// }
// EOF
// }

// resource "aws_iam_policy" "cloudability_monitor_resources_policy" {
// 	name   = "CloudabilityMonitorResourcesPolicy"
// 	path   = "/"
// 	policy = <<EOF
// {
// 	"Version": "2012-10-17",
// 	"Statement": [
// 	  	{
// 			"Effect": "Allow",
// 			"Action": [
// 			"cloudwatch:GetMetricStatistics",
// 			"dynamodb:DescribeTable",
// 			"dynamodb:ListTables",
// 			"ec2:DescribeImages",
// 			"ec2:DescribeInstances",
// 			"ec2:DescribeRegions",
// 			"ec2:DescribeReservedInstances",
// 			"ec2:DescribeReservedInstancesModifications",
// 			"ec2:DescribeSnapshots",
// 			"ec2:DescribeVolumes",
// 			"ec2:GetReservedInstancesExchangeQuote",
// 			"ecs:DescribeClusters",
// 			"ecs:DescribeContainerInstances",
// 			"ecs:ListClusters",
// 			"ecs:ListContainerInstances",
// 			"elasticache:DescribeCacheClusters",
// 			"elasticache:DescribeReservedCacheNodes",
// 			"elasticache:ListTagsForResource",
// 			"elasticmapreduce:DescribeCluster",
// 			"elasticmapreduce:ListClusters",
// 			"elasticmapreduce:ListInstances",
// 			"rds:DescribeDBClusters",
// 			"rds:DescribeDBInstances",
// 			"rds:DescribeReservedDBInstances",
// 			"rds:ListTagsForResource",
// 			"redshift:DescribeClusters",
// 			"redshift:DescribeReservedNodes",
// 			"redshift:DescribeTags",
// 			"savingsplans:DescribeSavingsPlans"
// 			],
// 			"Resource": "*"
// 	  	}
// 	]
// }
// EOF
// }

// resource "aws_iam_policy" "cloudability_master_payer_policy" {
// 	name   = "CloudabilityMasterPayerPolicy"
// 	path   = "/"
// 	policy = <<EOF
// {
// 	"Version": "2012-10-17",
// 	"Statement": [
// 	  	{
// 			"Effect": "Allow",
// 			"Action": [
// 				"s3:ListBucket",
// 				"s3:GetObject"
// 			],
// 			"Resource": [
// 				"${aws_s3_bucket.cloudability.arn}",
// 				"${aws_s3_bucket.cloudability.arn}/*"
// 			]
// 	  	},
// 	  	{
// 			"Effect": "Allow",
// 			"Action": [
// 				"organizations:ListAccounts"
// 			],
// 			"Resource": "*"
// 		}
// 	]
// }
// EOF
// }

// resource "aws_iam_role_policy_attachment" "cloudability_master_payer_policy" {
// 	role       = aws_iam_role.cloudability_role.id
// 	policy_arn = aws_iam_policy.cloudability_master_payer_policy.arn
// }

// resource "aws_iam_role_policy_attachment" "cloudability_verification_policy" {
// 	role       = aws_iam_role.cloudability_role.id
// 	policy_arn = aws_iam_policy.cloudability_verification_policy.arn
// }

// resource "aws_iam_role_policy_attachment" "cloudability_monitor_resources_policy" {
// 	role       = aws_iam_role.cloudability_role.id
// 	policy_arn = aws_iam_policy.cloudability_monitor_resources_policy.arn
// }

// data "cloudability_account_verification" "aws_payer_account" {
// 	vendor_account_id = "%s"
// 	dependency = aws_iam_role.cloudability_role.id
// }
// `, cloudabilityApikey, awsRegion, payerAccountId, payerAccountId)
// }

// func TestResourceMasterAccountRead(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
// 	d.Set("vendor_key", "aws")
// 	c := config.Client()
// 	err := resourceMasterAccountRead(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceMasterAccountCreate(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	accountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
// 	d.Set("vendor_account_id", accountId)
// 	d.Set("vendor_key", "aws")
// 	d.Set("type", "aws_role")
// 	d.Set("bucket_name", fmt.Sprintf("cloudability-%s", accountId))
// 	d.Set("report_name", "Cloudability")
// 	d.Set("report_prefix", "CostAndUsageReports")

// 	c := config.Client()
// 	err := resourceMasterAccountCreate(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceMasterAccountDelete(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
// 	d.Set("vendor_key", "aws")
// 	c := config.Client()
// 	err := resourceMasterAccountDelete(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }
