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

// func TestAccLinkedAccount(t *testing.T) {
// 	var account *cloudability.Account
// 	payerAccountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
// 	accountId := os.Getenv("CLOUDABILITY_ACCOUNTID")
// 	cloudabilityApikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	awsRegion := os.Getenv("AWS_DEFAULT_REGION")

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckLinkedAccountDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccLinkedAccount(cloudabilityApikey, awsRegion, payerAccountId, accountId),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckLinkedAccountResourceExists("cloudability_master_account.aws_payer_account", account),
// 					// TODO: Complete this
// 					// testAccCheckExampleWidgetValues(widget, rName),
// 					resource.TestCheckResourceAttr("cloudability_master_account.aws_payer_account", "vendor_account_id", payerAccountId),
// 					resource.TestCheckResourceAttr("cloudability_master_account.aws_payer_account", "vendor_account_name", "9492-3762-0074"),

// 					resource.TestCheckResourceAttr("cloudability_master_account.aws_account", "vendor_account_id", accountId),
// 					resource.TestCheckResourceAttr("cloudability_master_account.aws_account", "vendor_account_name", "CloudabilityLinkedTest"),
// 				),
// 			},
// 		},
// 	})
// }

// // testAccCheckLinkedAccountDestroy verifies the LinkedAccount resource has been destroyed.
// func testAccCheckLinkedAccountDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*cloudability.Client)
// 	// loop through the resources in state, verifying each resource is destroyed
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "cloudability_linked_account" {
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
// 				return fmt.Errorf("Linked Account (%s) still exists.", rs.Primary.ID)
// 			}
// 		}
// 	}
// 	return nil
// }

// // testAccCheckLinkedAccountResourceExists uses the Cloudability SDK to retrieve the Account.
// func testAccCheckLinkedAccountResourceExists(resourceName string, account *cloudability.Account) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		// retrieve the resource by name from state
// 		rs, ok := s.RootModule().Resources[resourceName]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", resourceName)
// 		}
// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("LinkedAccount ID is not set")
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

// func testAccLinkedAccount(cloudabilityApikey string, awsRegion string, payerAccountId string, accountId string) string {
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

// // func TestResourceLinkedAccountRead(t *testing.T) {
// // 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// // 	config := Config{
// // 		ApiKey: apikey,
// // 	}
// // 	resource := resourceLinkedAccount()
// // 	d := resource.Data(nil)
// // 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
// // 	d.Set("vendor_key", "aws")
// // 	c := config.Client()
// // 	err := resourceLinkedAccountRead(d, c)
// // 	if err != nil {
// // 		t.Errorf(err.Error())
// // 	}
// // }

// // func TestResourceLinkedAccountCreate(t *testing.T) {
// // 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// // 	config := Config{
// // 		ApiKey: apikey,
// // 	}
// // 	resource := resourceLinkedAccount()
// // 	d := resource.Data(nil)
// // 	accountId := os.Getenv("CLOUDABILITY_ACCOUNTID")
// // 	d.Set("vendor_account_id", accountId)
// // 	d.Set("vendor_key", "aws")
// // 	d.Set("type", "aws_role")

// // 	c := config.Client()
// // 	err := resourceLinkedAccountCreate(d, c)
// // 	if err != nil {
// // 		t.Errorf(err.Error())
// // 	}
// // }

// // func TestResourceLinkedAccountDelete(t *testing.T) {
// // 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// // 	config := Config{
// // 		ApiKey: apikey,
// // 	}
// // 	resource := resourceLinkedAccount()
// // 	d := resource.Data(nil)
// // 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
// // 	d.Set("vendor_key", "aws")
// // 	c := config.Client()
// // 	err := resourceLinkedAccountDelete(d, c)
// // 	if err != nil {
// // 		t.Errorf(err.Error())
// // 	}
// // }
